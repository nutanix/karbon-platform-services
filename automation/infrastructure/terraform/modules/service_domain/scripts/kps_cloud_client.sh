#!/bin/bash -x

# Bash strict mode, stop on any error
set -euo pipefail

sub_login_to_kps_cloud() {
  echo "Running 'login_to_kps_cloud' command."

  test -n "${CLOUD_FQDN}"
  test -n "${CLOUD_USER_NAME}"
  test -n "${CLOUD_USER_PWD}"
  test -n "${LOGIN_TOKEN_OUTPUT_FILE_PATH}"

    # Get Bearer token by loging in to cloud
  resp=$(curl --location --request POST "https://${CLOUD_FQDN}/v1.0/login" \
  --header "Content-Type: application/json" \
  --data-raw "{
    \"email\": \"${CLOUD_USER_NAME}\",
    \"password\": \"${CLOUD_USER_PWD}\"
  }") 
  status=$?
  echo $resp
  case $status in
    0)
        echo "Successfully logged into KPS Cloud"
        echo ${resp} | jq -r .token | xargs echo -n > ${LOGIN_TOKEN_OUTPUT_FILE_PATH}
        ;;
    *)
        echo "Failed to login to KPS Cloud"
        echo $status
        exit "$status"
  esac
}

sub_add_ebs_storage_profile_to_servicedomain() {
  echo "Running 'add_ebs_storage_profile_to_servicedomain' command."

  test -n "${CLOUD_FQDN}"
  test -n "${BEARER_TOKEN}"
  test -n "${SERVICE_DOMAIN_ID}"

  test -n "${STORAGE_PROFILE_NAME}"
  test -n "${IS_DEFAULT}"

  test -n "${IOPS_PER_GB}"
  test -n "${TYPE}"

  resp=$(curl --location --request POST "https://${CLOUD_FQDN}/v1.0/servicedomains/${SERVICE_DOMAIN_ID}/storageprofiles" \
  --header "Authorization: Bearer ${BEARER_TOKEN}" \
  --header "Content-Type: application/json" \
  --data-raw "{
      \"ebsStorageConfig\": {
          \"encrypted\": \"false\",
          \"iops_per_gb\": \"${IOPS_PER_GB}\",
          \"type\": \"${TYPE}\"
      },
      \"type\": \"EBS\",
      \"name\": \"${STORAGE_PROFILE_NAME}\",
      \"isDefault\": ${IS_DEFAULT}
  }")
  status=$?
  echo $resp
  case $status in
    0)
        echo "Successfully added ebs storage profile to service domain"
        ;;
    *)
        echo "Failed to add ebs storage profile to service domain"
        echo $status
        exit "$status"
  esac
}

sub_remove_storage_profile_from_servicedomain() {
  echo "Running 'remove_storage_profile_from_servicedomain' command."
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