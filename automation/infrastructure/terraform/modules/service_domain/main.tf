resource "null_resource" "fetch_instance_serialnumber" {
  triggers  =  {
    always_run = "${timestamp()}"
  }
  count = var.instance_info["instance_count"]
  provisioner "local-exec" {
    command = "${path.module}/scripts/kps_sd_node_client.sh get_serial_number"
    environment = {
      NODE_IP = var.public_instance_ips[count.index]
      NODE_SN_FILE_PATH = "${path.module}/generated/sn-${count.index}.txt"
    }
  }
}

resource "null_resource" "fetch_instance_ips" {
  triggers  =  {
    always_run = "${timestamp()}"
  }
  count = var.instance_info["instance_count"]
  provisioner "local-exec" {
    command = "${path.module}/scripts/kps_sd_node_client.sh fetch_ip_address"
    environment = {
      NODE_PRIVATE_IP = var.private_instance_ips[count.index]
      NODE_PUBLIC_IP = var.public_instance_ips[count.index]
      NODE_PRIVATE_IP_FILE_PATH = "${path.module}/generated/private_ip_${count.index}.txt"
      NODE_PUBLIC_IP_FILE_PATH = "${path.module}/generated/public_ip_${count.index}.txt"
    }
  }
}

resource "null_resource" "configure_onboarding_cloud_fqdn" {
  triggers  =  {
    always_run = "${timestamp()}"
  }
  count = var.instance_info["instance_count"]
  provisioner "local-exec" {
    command = "${path.module}/scripts/kps_sd_node_client.sh configure_custom_kps_cloud_endpoint"
    environment = {
      CLOUD_FQDN = var.cloud_info["cloud_fqdn"]
      NODE_IP = var.public_instance_ips[count.index]
    }
  }

  depends_on = [
    null_resource.fetch_instance_serialnumber
  ]
}

#
resource "null_resource" "login_to_kps_cloud" {
  triggers  =  {
    always_run = "${timestamp()}"
  }
  provisioner "local-exec" {
    command = "${path.module}/scripts/kps_cloud_client.sh login_to_kps_cloud"
    environment = {
      CLOUD_FQDN = var.cloud_info["cloud_fqdn"]
      CLOUD_USER_NAME = var.cloud_info["cloud_user_name"]
      CLOUD_USER_PWD = var.cloud_info["cloud_user_pwd"]
      LOGIN_TOKEN_OUTPUT_FILE_PATH = "${path.module}/generated/login-token.txt"
    }
  }

  depends_on = [
    null_resource.fetch_instance_serialnumber
  ]
}

data "local_file" "kps_cloud_login_token" {
  filename = "${path.module}/generated/login-token.txt"

  depends_on = [
    null_resource.login_to_kps_cloud
  ]
}

resource "null_resource" "service_domain" {
  triggers  =  {
    always_run = "${timestamp()}"
  }
  provisioner "local-exec" {
    command = "${path.module}/scripts/kps_cloud_client.sh create_servicedomain"
    environment = {
      CLOUD_FQDN = var.cloud_info["cloud_fqdn"]
      BEARER_TOKEN = data.local_file.kps_cloud_login_token.content
      SERVICE_DOMAIN_DESC = var.service_domain_info["sd_description"]
      SERVICE_DOMAIN_NAME = var.service_domain_info["sd_name"]
      SERVICE_DOMAIN_ID_OUTPUT_FILE_PATH = "${path.module}/generated/new_servicedomain.txt"
      SERVICE_DOMAIN_VIRTUAL_IP = var.service_domain_info["sd_virtual_ip"]
    }
  }

  depends_on = [
    data.local_file.kps_cloud_login_token
  ]
}

data "local_file" "service_domain_id" {
  filename = "${path.module}/generated/new_servicedomain.txt"

  depends_on = [
    null_resource.service_domain
  ]
}


resource "null_resource" "service_domain_destroy" {
  triggers ={
    cloud_fqdn = var.cloud_info["cloud_fqdn"]
    bearer_token = data.local_file.kps_cloud_login_token.content
    service_domain_id = data.local_file.service_domain_id.content
  }
  provisioner "local-exec" {
    when = destroy
    command = "${path.module}/scripts/kps_cloud_client.sh delete_servicedomain"
    environment = {
      CLOUD_FQDN = "${self.triggers.cloud_fqdn}"
      BEARER_TOKEN = "${self.triggers.bearer_token}"
      SERVICE_DOMAIN_ID = "${self.triggers.service_domain_id}"
      SERVICE_DOMAIN_ID_OUTPUT_FILE_PATH = "${path.module}/generated/new_servicedomain.txt"
    }
  }

  depends_on = [
    data.local_file.kps_cloud_login_token,
    data.local_file.service_domain_id
  ]
}

data "local_file" "node_sn" {
  filename = "${path.module}/generated/sn-${count.index}.txt"
  count = var.instance_info["instance_count"]
  depends_on = [
    null_resource.fetch_instance_serialnumber
  ]
}

data "local_file" "node_private_ip" {
  filename = "${path.module}/generated/private_ip_${count.index}.txt"
  count = var.instance_info["instance_count"]
  depends_on = [
    null_resource.fetch_instance_ips
  ]
}

data "local_file" "node_public_ip" {
  filename = "${path.module}/generated/public_ip_${count.index}.txt"
  count = var.instance_info["instance_count"]
  depends_on = [
    null_resource.fetch_instance_ips
  ]
}

resource "null_resource" "service_domain_node" {
  triggers  =  {
    always_run = "${timestamp()}"
  }
  count = var.instance_info["instance_count"]
  provisioner "local-exec" {
    command = "${path.module}/scripts/kps_cloud_client.sh add_node_to_servicedomain"
    environment = {
      CLOUD_FQDN = var.cloud_info["cloud_fqdn"]
      BEARER_TOKEN = data.local_file.kps_cloud_login_token.content
      NODE_NAME = "node-${count.index}1"
      NODE_GATEWAY = var.node_info["node_gateway"]
      NODE_IP = element(data.local_file.node_private_ip.*.content, count.index)
      NODE_SUBNET = var.node_info["node_subnet"]
      NODE_SERIAL_NUMBER = element(data.local_file.node_sn.*.content, count.index)
      SERVICE_DOMAIN_ID = data.local_file.service_domain_id.content
      NODE_ID_OUTPUT_FILE_PATH = "${path.module}/generated/new_node_${count.index}.txt"
    }
  }

  depends_on = [
    data.local_file.kps_cloud_login_token,
    data.local_file.service_domain_id,
    data.local_file.node_sn,
    data.local_file.node_private_ip,
  ]
}

data "local_file" "node_id" {
  filename = "${path.module}/generated/new_node_${count.index}.txt"
  count = var.instance_info["instance_count"]
  depends_on = [
    null_resource.service_domain_node
  ]
}

resource "null_resource" "service_domain_node_destroy" {
  count = var.instance_info["instance_count"]
  triggers ={
    cloud_fqdn = var.cloud_info["cloud_fqdn"]
    bearer_token = data.local_file.kps_cloud_login_token.content
    node_id = element(data.local_file.node_id.*.content, count.index)
  }
  provisioner "local-exec" {
    when = destroy
    command = "${path.module}/scripts/kps_cloud_client.sh remove_node_from_servicedomain"
    environment = {
      CLOUD_FQDN = "${self.triggers.cloud_fqdn}"
      BEARER_TOKEN = "${self.triggers.bearer_token}"
      NODE_ID = "${self.triggers.node_id}"
      NODE_ID_OUTPUT_FILE_PATH = "${path.module}/generated/new_node_${count.index}.txt"
    }
  }

  depends_on = [
    data.local_file.kps_cloud_login_token,
    data.local_file.node_id
  ]
}

resource "null_resource" "nutanixvolumes_storage_profile" {
  triggers  =  {
    always_run = "${timestamp()}"
  }
  // count is used as boolean here
  count = var.create_nutanixvolumes_storage_profile
  provisioner "local-exec" {
    command = "${path.module}/scripts/kps_cloud_client.sh add_nutanixvolumes_storage_profile_to_servicedomain"
    environment = {
      CLOUD_FQDN = var.cloud_info["cloud_fqdn"]
      BEARER_TOKEN = data.local_file.kps_cloud_login_token.content
      SERVICE_DOMAIN_ID = data.local_file.service_domain_id.content
      DATA_SERVICES_IP = var.storage_config["dataServicesIP"]
      DATA_SERVICES_PORT = var.storage_config["dataServicesPort"]
      FLASH_MODE = var.storage_config["flashMode"]
      PE_CLUSTER_VIP = var.storage_config["prismElementClusterVIP"]
      PE_CLUSTER_PORT = var.storage_config["prismElementClusterPort"]
      STORAGE_CONTAINER_NAME = var.storage_config["storageContainerName"]
      PE_USER_NAME = var.storage_config["prismElementUserName"]
      PE_USER_PWD = var.storage_config["prismElementPassword"]
      STORAGE_PROFILE_NAME = var.storage_profile_info["name"]
      IS_DEFAULT = var.storage_profile_info["isDefault"]
    }
  }

  depends_on = [
    data.local_file.kps_cloud_login_token,
    data.local_file.service_domain_id
  ]
}

resource "null_resource" "ebs_storage_profile" {
  triggers  =  {
    always_run = "${timestamp()}"
  }
  // count is used as boolean here
  count = var.create_ebs_storage_profile
  provisioner "local-exec" {
    command = "${path.module}/scripts/kps_cloud_client.sh add_ebs_storage_profile_to_servicedomain"
    environment = {
      CLOUD_FQDN = var.cloud_info["cloud_fqdn"]
      BEARER_TOKEN = data.local_file.kps_cloud_login_token.content
      SERVICE_DOMAIN_ID = data.local_file.service_domain_id.content
      STORAGE_PROFILE_NAME = var.storage_profile_info["name"]
      IS_DEFAULT = var.storage_profile_info["isDefault"]
      ENCRYPTED = var.storage_config["encrypted"]
      IOPS_PER_GB = var.storage_config["iops_per_gb"]
      TYPE = var.storage_config["type"]
    }
  }

  depends_on = [
    data.local_file.kps_cloud_login_token,
    data.local_file.service_domain_id
  ]
}