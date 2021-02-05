#https://learn.hashicorp.com/tutorials/terraform/aws-variables

provider "nutanixkps" {
  host = var.cloud_info["cloud_fqdn"]
  username = var.cloud_info["cloud_user_name"]
  password = var.cloud_info["cloud_user_pwd"]
}

provider "aws" {
    profile = var.profile
    region = var.region
}

data "aws_security_group" "kps_security_group" {
  id = var.security_group
}

data "aws_vpc" "kps_vpc" {
  id = data.aws_security_group.kps_security_group.vpc_id
}

resource "null_resource" "kps_ami" {
  provisioner "local-exec" {
    command = "${path.module}/scripts/kps_ami.sh create_ami"
    environment = {
      VERSION = var.kps_raw_diskimage_version
      SNAPSHOT_ID_OUTPUT_FILE_PATH = "${path.module}/generated/snapshot_id.txt"
      AMI_ID_OUTPUT_FILE_PATH = "${path.module}/generated/ami_id.txt"
    }
  }
  provisioner "local-exec" {
    when = destroy
    command = "${path.module}/scripts/kps_ami.sh delete_ami"
    environment = {
      SNAPSHOT_ID_OUTPUT_FILE_PATH = "${path.module}/generated/snapshot_id.txt"
      AMI_ID_OUTPUT_FILE_PATH = "${path.module}/generated/ami_id.txt"
    }
  }
}

data "local_file" "ami_id" {
  filename = "${path.module}/generated/ami_id.txt"
  depends_on = [
    null_resource.kps_ami
  ]
}

resource "aws_iam_role" "role" {
  name        = var.iam_config["aws_iam_role_name"]
  force_detach_policies = true
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "policy" {
  name        = var.iam_config["aws_iam_policy_name"]
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ec2:AttachVolume",
        "ec2:CreateSnapshot",
        "ec2:CreateTags",
        "ec2:CreateVolume",
        "ec2:DeleteSnapshot",
        "ec2:DeleteTags",
        "ec2:DeleteVolume",
        "ec2:DescribeAvailabilityZones",
        "ec2:DescribeInstances",
        "ec2:DescribeSnapshots",
        "ec2:DescribeTags",
        "ec2:DescribeVolumes",
        "ec2:DescribeVolumesModifications",
        "ec2:DetachVolume",
        "ec2:ModifyVolume"
      ],
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_iam_policy_attachment" "policy-attach" {
  name       = "ebs-policy-attachment"
  roles      = [aws_iam_role.role.name]
  policy_arn = aws_iam_policy.policy.arn
}

resource "aws_iam_instance_profile" "instance_profile" {
  name = var.iam_config["aws_iam_instance_profile_name"]
  role = aws_iam_role.role.name
}

resource "aws_instance" "kps_servicedomain_instance" {
  ami = data.local_file.ami_id.content
  instance_type = var.ec2_vm_config["instance_type"]
  security_groups = [data.aws_security_group.kps_security_group.name]
  iam_instance_profile = aws_iam_instance_profile.instance_profile.name
  availability_zone = var.availability_zone
  count = var.instance_info["instance_count"]
  tags = {
    Name = join("-", [var.instance_info["instance_name_prefix"], count.index])
  }
}

resource "aws_ebs_volume" "kps_volume" {
  availability_zone = var.availability_zone
  size = var.data_partition_size_gb
  count = var.instance_info["instance_count"]
}

resource "aws_volume_attachment" "kps_volume_attachment" {
  device_name = "/dev/xvdf"
  volume_id = aws_ebs_volume.kps_volume[count.index].id
  count = var.instance_info["instance_count"]
  instance_id = aws_instance.kps_servicedomain_instance.*.id[count.index]
  force_detach = true
  depends_on = [
    aws_ebs_volume.kps_volume,
    aws_instance.kps_servicedomain_instance
  ]
}

module "service_domain" {
  source = "../modules/service_domain"
  depends_on = [
    aws_volume_attachment.kps_volume_attachment
  ]
  instance_info = var.instance_info
  cloud_info = var.cloud_info
  node_info = var.node_info
  storage_config = var.ebs_storage_config
  service_domain_info = var.service_domain_info
  storage_profile_info = var.storage_profile_info
  kps_servicedomain_instance_details = aws_instance.kps_servicedomain_instance
  create_ebs_storage_profile = var.create_storage_profile
  create_nutanixvolumes_storage_profile = 0
  private_instance_ips = aws_instance.kps_servicedomain_instance[*].private_ip
  public_instance_ips = aws_instance.kps_servicedomain_instance[*].public_ip
  wait_for_onboarding = var.wait_for_onboarding
}