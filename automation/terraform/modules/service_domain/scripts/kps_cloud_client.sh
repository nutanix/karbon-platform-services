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

sub_create_servicedomain() {
  echo "Running 'create_servicedomain' command."

  # Ensure all required environment variables are present
  test -n "${CLOUD_FQDN}"
  test -n "${BEARER_TOKEN}"
  test -n "${SERVICE_DOMAIN_DESC}"
  test -n "${SERVICE_DOMAIN_NAME}"
  test -n "${SERVICE_DOMAIN_ID_OUTPUT_FILE_PATH}"
  test -n "${SERVICE_DOMAIN_VIRTUAL_IP}"

  # Create Service Domain
  resp=$(curl --location --request POST "https://${CLOUD_FQDN}/v1.0/servicedomains" \
  --header "Authorization: Bearer ${BEARER_TOKEN}" \
  --header "Content-Type: application/json" \
  --data-raw "{
    \"description\": \"${SERVICE_DOMAIN_DESC}\",
    \"name\": \"${SERVICE_DOMAIN_NAME}\",
    \"virtualIp\": \"${SERVICE_DOMAIN_VIRTUAL_IP}\"
  }")
  status=$?
  echo $resp
  case $status in
    0)
        echo "Successfully created service domain"
        echo ${resp} | jq -r .id | xargs echo -n > "${SERVICE_DOMAIN_ID_OUTPUT_FILE_PATH}"
        ;;
    *)
        echo "Failed to create service domain"
        echo $status
        exit "$status"
  esac
}

sub_delete_servicedomain() {
  echo "Running 'delete_servicedomain' command."

  # Ensure all required environment variables are present
  test -n "${CLOUD_FQDN}"
  test -n "${BEARER_TOKEN}"
  test -n "${SERVICE_DOMAIN_ID}"
  test -n "${SERVICE_DOMAIN_ID_OUTPUT_FILE_PATH}"

  resp=$(curl --location --request DELETE "https://${CLOUD_FQDN}/v1.0/servicedomains/${SERVICE_DOMAIN_ID}" \
    --header "Authorization: Bearer ${BEARER_TOKEN}") 
  status=$?
  echo $resp
  case $status in
    0)
        echo "Successfully deleted service domain"
        echo ${resp} | jq -r .id | xargs echo -n > "${SERVICE_DOMAIN_ID_OUTPUT_FILE_PATH}"
        ;;
    *)
        echo "Failed to delete service domain"
        echo $status
        exit "$status"
  esac
}

sub_add_node_to_servicedomain() {
  echo "Running 'add_node_to_servicedomain' command."

  # Ensure all required environment variables are present
  test -n "${CLOUD_FQDN}"
  test -n "${BEARER_TOKEN}"
  test -n "${NODE_NAME}"
  test -n "${NODE_GATEWAY}"
  test -n "${NODE_IP}"
  test -n "${NODE_SUBNET}"
  test -n "${NODE_SERIAL_NUMBER}"
  test -n "${SERVICE_DOMAIN_ID}"
  test -n "${NODE_ID_OUTPUT_FILE_PATH}"

  # Add node to service domain
  resp=$(curl --location --request POST "https://${CLOUD_FQDN}/v1.0/nodes" \
  --header "Authorization: Bearer ${BEARER_TOKEN}" \
  --header "Content-Type: application/json" \
  --data-raw "{
    \"description\": \"${NODE_NAME}\",
    \"gateway\": \"${NODE_GATEWAY}\",
    \"ipAddress\": \"${NODE_IP}\",
    \"isBootstrapMaster\": true,
    \"name\": \"${NODE_NAME}\",
    \"role\": {
      \"master\": true,
      \"worker\": true
    },
    \"serialNumber\": \"${NODE_SERIAL_NUMBER}\",
    \"subnet\": \"${NODE_SUBNET}\",
    \"svcDomainId\": \"${SERVICE_DOMAIN_ID}\"
  }")
  status=$?
  echo $resp
  case $status in
    0)
        echo "Successfully added node to service domain"
        echo ${resp} | jq -r .id | xargs echo -n > "${NODE_ID_OUTPUT_FILE_PATH}"
        ;;
    *)
        echo "Failed to add node to service domain"
        echo $status
        exit "$status"
  esac
}

sub_remove_node_from_servicedomain() {
  echo "Running 'remove_node_from_servicedomain' command."

  # Ensure all required environment variables are present
  test -n "${CLOUD_FQDN}"
  test -n "${BEARER_TOKEN}"
  test -n "${NODE_ID}"

  resp=$(curl --location --request DELETE "https://${CLOUD_FQDN}/v1.0/nodes/${NODE_ID}" \
    --header "Authorization: Bearer ${BEARER_TOKEN}")
  status=$?
  echo $resp
  case $status in
    0)
        echo "Successfully remove node from service domain"
        ;;
    *)
        echo "Failed to remove node from service domain"
        echo $status
        exit "$status"
  esac
}

sub_add_nutanixvolumes_storage_profile_to_servicedomain() {
  echo "Running 'add_nutanixvolumes_storage_profile_to_servicedomain' command."

  test -n "${CLOUD_FQDN}"
  test -n "${BEARER_TOKEN}"
  test -n "${SERVICE_DOMAIN_ID}"
  test -n "${DATA_SERVICES_IP}"
  test -n "${DATA_SERVICES_PORT}"
  test -n "${FLASH_MODE}"
  test -n "${PE_CLUSTER_VIP}"
  test -n "${PE_CLUSTER_PORT}"
  test -n "${STORAGE_CONTAINER_NAME}"
  test -n "${PE_USER_NAME}"
  test -n "${PE_USER_PWD}"
  test -n "${STORAGE_PROFILE_NAME}"
  test -n "${IS_DEFAULT}"

  resp=$(curl --location --request POST "https://${CLOUD_FQDN}/v1.0/servicedomains/${SERVICE_DOMAIN_ID}/storageprofiles" \
  --header "Authorization: Bearer ${BEARER_TOKEN}" \
  --header "Content-Type: application/json" \
  --data-raw "{
      \"nutanixVolumesConfig\": {
          \"dataServicesIP\": \"${DATA_SERVICES_IP}\",
          \"dataServicesPort\": ${DATA_SERVICES_PORT},
          \"flashMode\": ${FLASH_MODE},
          \"prismElementClusterVIP\": \"${PE_CLUSTER_VIP}\",
          \"prismElementClusterPort\": ${PE_CLUSTER_PORT},
          \"storageContainerName\": \"${STORAGE_CONTAINER_NAME}\",
          \"prismElementUserName\": \"${PE_USER_NAME}\",
          \"prismElementPassword\": \"${PE_USER_PWD}\"
      },
      \"type\": \"NutanixVolumes\",
      \"name\": \"${STORAGE_PROFILE_NAME}\",
      \"isDefault\": ${IS_DEFAULT}
  }")
  status=$?
  echo $resp
  case $status in
    0)
        echo "Successfully added nutanix storage profile to service domain"
        ;;
    *)
        echo "Failed to add nutanix storage profile to service domain"
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

  test -n "${ENCRYPTED}"
  test -n "${IOPS_PER_GB}"
  test -n "${TYPE}"

  resp=$(curl --location --request POST "https://${CLOUD_FQDN}/v1.0/servicedomains/${SERVICE_DOMAIN_ID}/storageprofiles" \
  --header "Authorization: Bearer ${BEARER_TOKEN}" \
  --header "Content-Type: application/json" \
  --data-raw "{
      \"ebsStorageConfig\": {
          \"encrypted\": \"${ENCRYPTED}\",
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