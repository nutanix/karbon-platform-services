resource "null_resource" "fetch_instance_serialnumber" {
  triggers  =  {
    always_run = "${timestamp()}"
  }
  count = var.instance_info["instance_count"]
  provisioner "local-exec" {
    command = "${path.module}/scripts/kps_sd_node_client.sh get_serial_number"
    environment = {
      NODE_IP = var.public_instance_ips[count.index]
      NODE_SN_FILE_PATH = "${path.module}/generated/sn-${count.index}_${terraform.workspace}.txt"
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
      NODE_PRIVATE_IP_FILE_PATH = "${path.module}/generated/private_ip_${count.index}_${terraform.workspace}.txt"
      NODE_PUBLIC_IP_FILE_PATH = "${path.module}/generated/public_ip_${count.index}_${terraform.workspace}.txt"
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
      LOGIN_TOKEN_OUTPUT_FILE_PATH = "${path.module}/generated/login-token_${terraform.workspace}.txt"
    }
  }

  depends_on = [
    null_resource.fetch_instance_serialnumber
  ]
}

data "local_file" "kps_cloud_login_token" {
  filename = "${path.module}/generated/login-token_${terraform.workspace}.txt"

  depends_on = [
    null_resource.login_to_kps_cloud
  ]
}

resource "nutanixkps_servicedomain" "service_domain" {  
  name = var.service_domain_info["sd_name"]
  description = var.service_domain_info["sd_description"]
  virtual_ip = var.service_domain_info["sd_virtual_ip"]
}

output "servicedomains" {
  value = nutanixkps_servicedomain.service_domain
}

data "local_file" "node_sn" {
  filename = "${path.module}/generated/sn-${count.index}_${terraform.workspace}.txt"
  count = var.instance_info["instance_count"]
  depends_on = [
    null_resource.fetch_instance_serialnumber
  ]
}

data "local_file" "node_private_ip" {
  filename = "${path.module}/generated/private_ip_${count.index}_${terraform.workspace}.txt"
  count = var.instance_info["instance_count"]
  depends_on = [
    null_resource.fetch_instance_ips
  ]
}

data "local_file" "node_public_ip" {
  filename = "${path.module}/generated/public_ip_${count.index}_${terraform.workspace}.txt"
  count = var.instance_info["instance_count"]
  depends_on = [
    null_resource.fetch_instance_ips
  ]
}

resource "nutanixkps_node" "nodes" {
  count = var.instance_info["instance_count"]
  name = "${var.instance_info["instance_name_prefix"]}-${count.index}"
  description = "Node added to Service Domain through Terraform"
  service_domain_id = nutanixkps_servicedomain.service_domain.id
  serial_number = upper(element(data.local_file.node_sn.*.content, count.index))
  ip_address = element(data.local_file.node_private_ip.*.content, count.index)
  gateway = var.node_info["node_gateway"]
  subnet = var.node_info["node_subnet"]
  role {
    master = true
    worker = true
  }
  wait_for_onboarding = var.wait_for_onboarding

  depends_on = [
    nutanixkps_servicedomain.service_domain,
    data.local_file.node_sn,
    data.local_file.node_private_ip,
  ]
}

output "nodes" {
  value = nutanixkps_node.nodes
}

resource "nutanixkps_storageprofile" "nutanixvolumes_storage_profile" {
  // count is used as a boolean here, only effective when creating Nutanix VM based service domain
  count = var.create_nutanixvolumes_storage_profile
  name = var.storage_profile_info["name"]
  service_domain_id = nutanixkps_servicedomain.service_domain.id
  is_default = var.storage_profile_info["isDefault"]
  nutanix_volumes_config {
    data_services_ip = var.storage_config["dataServicesIP"]
    data_services_port = var.storage_config["dataServicesPort"]
    flash_mode = var.storage_config["flashMode"]
    prism_element_cluster_vip = var.storage_config["prismElementClusterVIP"]
    prism_element_cluster_port = var.storage_config["prismElementClusterPort"]
    prism_element_password = var.storage_config["prismElementPassword"]
    prism_element_username = var.storage_config["prismElementUserName"]
    prism_element_storage_container_name = var.storage_config["storageContainerName"]
  }

  depends_on = [
    nutanixkps_servicedomain.service_domain
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
      SERVICE_DOMAIN_ID = nutanixkps_servicedomain.service_domain.id
      STORAGE_PROFILE_NAME = var.storage_profile_info["name"]
      IS_DEFAULT = var.storage_profile_info["isDefault"]
      IOPS_PER_GB = var.storage_config["iops_per_gb"]
      TYPE = var.storage_config["type"]
    }
  }

  depends_on = [
    data.local_file.kps_cloud_login_token,
    nutanixkps_servicedomain.service_domain
  ]
}