provider "nutanixkps" {
  host = var.cloud_info["cloud_fqdn"]
  username = var.cloud_info["cloud_user_name"]
  password = var.cloud_info["cloud_user_pwd"]
}

provider "nutanix" {
  username = var.provider_info["username"]
  password = var.provider_info["password"]
  endpoint = var.provider_info["endpoint"]
  insecure = var.provider_info["insecure"]
  port     = var.provider_info["port"]
}

data "nutanix_clusters" "clusters" {}

# Create/Upload Nutanix KPS Image

resource "nutanix_image" "kps_servicedomain_image" {
  name        = var.image_config["name"]
  description = var.image_config["description"]
  source_uri  = var.image_config["source_http"]
  depends_on = [
    data.nutanix_clusters.clusters
  ]
}

data "nutanix_image" "kps_servicedomain_image" {
  image_id = nutanix_image.kps_servicedomain_image.id
}

data "nutanix_subnet" "sherlock_net" {
  subnet_name = "sherlock_net"
}

# Create Nutanix KPS SD VM Instance
resource "nutanix_virtual_machine" "kps_servicedomain_instance" {
  provider = nutanix
  cluster_uuid= data.nutanix_clusters.clusters.entities.0.metadata.uuid

  count = var.instance_info["instance_count"]
  name = join("-", [var.instance_info["instance_name_prefix"], count.index])
  description = var.nutanix_vm_spec["description"]

  num_vcpus_per_socket = var.nutanix_vm_spec["num_vcpus_per_socket"]
  num_sockets          = var.nutanix_vm_spec["num_sockets"]
  memory_size_mib      = var.nutanix_vm_spec["memory_size_mib"]

  nic_list {
    subnet_name = data.nutanix_subnet.sherlock_net.name
    subnet_uuid = data.nutanix_subnet.sherlock_net.metadata.uuid
  }

  disk_list {
    # data_source_reference in the Nutanix API refers to where the source for
    # the disk device will come from. Could be a clone of a different VM or a
    # image like we're doing here
    data_source_reference = {
      kind = "image"
      uuid = data.nutanix_image.kps_servicedomain_image.image_id
    }
  }
}

resource "nutanixkps_servicedomain" "kps_servicedomain" {
  name = var.service_domain_info["sd_name"]
  description = var.service_domain_info["sd_description"]
  virtual_ip = var.service_domain_info["sd_virtual_ip"]
}

output "servicedomains" {
  value = nutanixkps_servicedomain.kps_servicedomain
}

data "nutanixkps_vm_config" "kps_vm_config" {
  count = var.instance_info["instance_count"]
  public_ip_address =  nutanix_virtual_machine.kps_servicedomain_instance[count.index].nic_list[0].ip_endpoint_list[0].ip
}

resource "nutanixkps_vm_cloud_config" "kps_vm_cloud_config" {
  count = var.instance_info["instance_count"]
  public_ip_address =  nutanix_virtual_machine.kps_servicedomain_instance[count.index].nic_list[0].ip_endpoint_list[0].ip
  cloud_fqdn = var.cloud_info["cloud_fqdn"]
}

resource "nutanixkps_node" "nodes" {
  count = var.instance_info["instance_count"]
  name = "${var.instance_info["instance_name_prefix"]}-${count.index}"
  description = "Node added to Service Domain through Terraform"
  service_domain_id = nutanixkps_servicedomain.kps_servicedomain.id
  serial_number = data.nutanixkps_vm_config.kps_vm_config[count.index].serial_number
  ip_address = nutanix_virtual_machine.kps_servicedomain_instance[count.index].nic_list[0].ip_endpoint_list[0].ip
  gateway = var.node_info["node_gateway"]
  subnet = var.node_info["node_subnet"]
  role {
    master = true
    worker = true
  }
  wait_for_onboarding = var.wait_for_onboarding
  depends_on = [
    nutanixkps_vm_cloud_config.kps_vm_cloud_config
  ]
}

output "nodes" {
  value = nutanixkps_node.nodes
}

resource "nutanixkps_storageprofile" "nutanixvolumes_storage_profile" {
  // count is used as a boolean here, only effective when creating Nutanix VM based service domain
  count = var.create_storage_profile
  name = var.storage_profile_info["name"]
  service_domain_id = nutanixkps_servicedomain.kps_servicedomain.id
  is_default = var.storage_profile_info["isDefault"]
  nutanix_volumes_config {
    data_services_ip = var.nutanix_volumes_config["dataServicesIP"]
    data_services_port = var.nutanix_volumes_config["dataServicesPort"]
    flash_mode = var.nutanix_volumes_config["flashMode"]
    prism_element_cluster_vip = var.nutanix_volumes_config["prismElementClusterVIP"]
    prism_element_cluster_port = var.nutanix_volumes_config["prismElementClusterPort"]
    prism_element_password = var.nutanix_volumes_config["prismElementPassword"]
    prism_element_username = var.nutanix_volumes_config["prismElementUserName"]
    prism_element_storage_container_name = var.nutanix_volumes_config["storageContainerName"]
  }
}