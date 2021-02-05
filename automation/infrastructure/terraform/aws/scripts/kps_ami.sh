#!/bin/bash -x

# Bash strict mode, stop on any error
set -euo pipefail

# Function to create container json to reference RAW file source
create_container_json() {
  echo "Creating container json..."
  
  KEY=$1
  cat > containers.json << EOF
  {
      "Description": "sherlock raw image",
      "Format": "RAW",
      "UserBucket": {
          "S3Bucket": "sherlock-raw-images",
          "S3Key": "$KEY"
      }
  }
EOF
}

# Function to upload RAW file to Amazon S3 bucket 
raw_upload_s3() {
  echo "Uploading sherlock raw file to S3..."

  if aws s3api head-bucket --bucket sherlock-raw-images 2>/dev/null
  then
    # Create S3 bucket 
    aws s3 mb s3://sherlock-raw-images
  fi

  # Upload raw file to bucket 
  aws s3 cp ../raw/${key} s3://sherlock-raw-images
}

# Function to create role and attach necessary policy 
create_vmimport() {
  echo "Creating IAM role..."
  {
    # Create a role named vmimport and grant VM Import/Export access to it
    aws iam create-role --role-name vmimport --assume-role-policy-document "file://scripts/trust-policy.json"

    # Attach the put-role policy to the sherlock-vmimport role
    aws iam put-role-policy --role-name vmimport --policy-name vmimport --policy-document "file://scripts/role-policy.json"
  } || {
    echo "VMImport role already exists"
  }
}

# Function to create Amazon Machine Image from RAW file 
sub_create_ami() {
  echo "Creating Amazon Machine Image..."

  test -n "${VERSION}"
  test -n "${SNAPSHOT_ID_OUTPUT_FILE_PATH}"
  test -n "${AMI_ID_OUTPUT_FILE_PATH}"

  mkdir generated

  KEY=$(echo sherlock-aws_${VERSION}.raw)
  create_container_json ${KEY}
  # raw_upload_s3 ${KEY}
  create_vmimport

  # Import snapshot as ami image
  import_resp=$(aws ec2 import-snapshot --description "My Sherlock AMI" --disk-container "file://containers.json")
  echo ${import_resp}

  # Grab import id from json response body
  import_id=$(jq -r '.ImportTaskId' <<<"${import_resp}")
  echo ImportId: ${import_id}

  # Waiting till snapshot is created to grab id
  task_resp=$(aws ec2 describe-import-snapshot-tasks --import-task-ids ${import_id} | jq -r '.ImportSnapshotTasks')
  status=$(jq -r '.[0].SnapshotTaskDetail.Status' <<< "${task_resp}")
  curr_progress=0

  # Awaiting completion of import task 
  while [ ${status} = "active" ]
  do
    task_resp=$(aws ec2 describe-import-snapshot-tasks --import-task-ids ${import_id} | jq -r '.ImportSnapshotTasks')
    status=$(jq -r .[0].SnapshotTaskDetail.Status <<< "${task_resp}")
    progress=$(jq -r .[0].SnapshotTaskDetail.Progress <<< "${task_resp}")
  done

  echo ${task_resp}

  echo "Writing Snapshot id to file..."
  snap_id=$(jq -r .[0].SnapshotTaskDetail.SnapshotId <<< "${task_resp}")
  echo -n ${snap_id} > "${SNAPSHOT_ID_OUTPUT_FILE_PATH}"

  # Register AMI 
  resp=$(aws ec2 register-image --virtualization-type hvm \
  --name "Sherlock Service Domain Image" \
  --architecture x86_64 \
  --root-device-name "/dev/sda1" \
  --block-device-mappings "[
      {
          \"DeviceName\": \"/dev/sda1\",
          \"Ebs\": {
              \"SnapshotId\": \"${snap_id}\"
          }
      }
  ]")

  echo "Writing AMI id to file..."
  ami_id=$(jq -r .ImageId <<< "${resp}")
  echo -n ${ami_id} > "${AMI_ID_OUTPUT_FILE_PATH}"
  aws ec2 modify-image-attribute --image-id ${ami_id} --launch-permission "Add=[{Group=all}]"
}

# Function to delete AMI and Snapshot 
sub_delete_ami() {

  test -n "${SNAPSHOT_ID_OUTPUT_FILE_PATH}"
  test -n "${AMI_ID_OUTPUT_FILE_PATH}"

  snap_id=$(cat "${SNAPSHOT_ID_OUTPUT_FILE_PATH}")
  ami_id=$(cat "${AMI_ID_OUTPUT_FILE_PATH}")
  {
    echo "Deleting Amazon Machine Image..."
    aws ec2 deregister-image --image-id ${ami_id}

    echo "Deleting AWS Snapshot..."
    aws ec2 delete-snapshot --snapshot-id ${snap_id}
  } || {
    echo "AMI and Snapshot already destroyed"
  }
}

subcommand=$1
case $subcommand in
"" | "-h" | "--help")
  sub_help
  ;;
*)
  shift
  sub_${subcommand} $@
  if [ $? = 127 ]; then
    echo "Error: '$subcommand' is not a known subcommand." >&2
    echo "       Run '$ProgName --help' for a list of known subcommands." >&2
    exit 1
  fi
  ;;
esac