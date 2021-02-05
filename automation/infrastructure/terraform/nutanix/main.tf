provider "nutanixkps" {
  host = "samnsnew5.ntnxsherlock.com"
  username = "test@ntnxsherlock.com"
  password = "test"
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

  depends_on = [
    nutanix_image.kps_servicedomain_image
  ]
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
  description = var.nutanix_vm_config["description"]

  num_vcpus_per_socket = var.nutanix_vm_config["num_vcpus_per_socket"]
  num_sockets          = var.nutanix_vm_config["num_sockets"]
  memory_size_mib      = var.nutanix_vm_config["memory_size_mib"]

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

  depends_on = [
    data.nutanix_image.kps_servicedomain_image
  ]
}

module "service_domain" {
  source = "../modules/service_domain"
  depends_on = [
    nutanix_virtual_machine.kps_servicedomain_instance
  ]
  instance_info = var.instance_info
  cloud_info = var.cloud_info
  node_info = var.node_info
  service_domain_info = var.service_domain_info
  storage_profile_info = var.storage_profile_info
  kps_servicedomain_instance_details = nutanix_virtual_machine.kps_servicedomain_instance
  storage_config = var.nutanix_volumes_config
  create_ebs_storage_profile = 0
  create_nutanixvolumes_storage_profile = var.create_storage_profile
  private_instance_ips = nutanix_virtual_machine.kps_servicedomain_instance[*].nic_list[0].ip_endpoint_list[0].ip
  public_instance_ips = nutanix_virtual_machine.kps_servicedomain_instance[*].nic_list[0].ip_endpoint_list[0].ip
  wait_for_onboarding = var.wait_for_onboarding
}