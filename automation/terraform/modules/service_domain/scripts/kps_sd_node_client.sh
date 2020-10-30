#!/bin/bash -x

# Bash strict mode, stop on any error
set -euo pipefail

sub_get_serial_number() {
  echo "Running 'get_serial_number' command."

  test -n "${NODE_IP}"
  test -n "${NODE_SN_FILE_PATH}"

  mkdir -p $(dirname ${NODE_SN_FILE_PATH})

  # Store all serial numbers in local file
  until curl http://${NODE_IP}:8080/v1/sn; do
      printf '.'
      sleep 5
  done
  resp=$(curl http://${NODE_IP}:8080/v1/sn)
  status=$?
  echo $resp
  case $status in
    0)
        echo "Successfully obtained serial number"
        echo -n ${resp} >  "${NODE_SN_FILE_PATH}" 
        ;;
    *)
        echo "Failed to get serial number"
        echo $status
        exit "$status"
  esac
}

sub_configure_custom_kps_cloud_endpoint() {
  echo "Running 'get_serial_number' command."

  test -n "${NODE_IP}"

  # Configure the cloud fqdn for test edges if not default to karbon.nutanix.com
  resp=$(curl --location --request POST "http://${NODE_IP}:8080/v1/configure" \
  --header "Content-Type: application/json" \
  --data-raw "{
      \"CloudMgmtFQDN\": \"${CLOUD_FQDN}\"
  }")
  status=$?
  echo $resp
  case $status in
    0)
        echo "Successfully configured custom KPS Cloud endpoint"
        ;;
    *)
        echo "Failed to configured custom KPS Cloud endpoint"
        echo $status
        exit "$status"
  esac
}

sub_fetch_ip_address() {
  echo "Running 'fetch_ip_address' command."

  test -n "${NODE_PRIVATE_IP}"
  test -n "${NODE_PUBLIC_IP}"
  test -n "${NODE_PRIVATE_IP_FILE_PATH}"
  test -n "${NODE_PUBLIC_IP_FILE_PATH}"

  mkdir -p $(dirname ${NODE_PRIVATE_IP_FILE_PATH})
  mkdir -p $(dirname ${NODE_PUBLIC_IP_FILE_PATH})

  # Store all VM ips in local file
  echo "${NODE_PRIVATE_IP}" | xargs echo -n > "${NODE_PRIVATE_IP_FILE_PATH}"
  echo "${NODE_PUBLIC_IP}" | xargs echo -n > "${NODE_PUBLIC_IP_FILE_PATH}"

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