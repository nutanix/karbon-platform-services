
provider "nutanix" {
  username     = var.prism_central_username
  password     = var.prism_central_password
  endpoint     = var.prism_central_endpoint
  insecure     = var.insecure
  port         = var.prism_central_port
  wait_timeout = var.wait_timeout
}

data "nutanix_cluster" "ntnxcluster" {
  name = var.prism_element_cluster
}

data "nutanix_subnets" "ntnxsubnet" {}

locals {
  subnet_uuid = [
    for entity in data.nutanix_subnets.ntnxsubnet.entities :
    entity.metadata.uuid
    if entity.cluster_reference.uuid == data.nutanix_cluster.ntnxcluster.id && entity.name == var.prism_element_subnet
  ]
}

# Create a Karbon Kubernetes Cluster
resource "nutanix_karbon_cluster" "cluster" {
  name    = var.karbon_cluster_name
  version = var.kubernetes_version
  storage_class_config {
    name           = "default"
    reclaim_policy = var.pv_reclaim_policy
    volumes_config {
      username                   = var.prism_central_username
      password                   = var.prism_central_password
      prism_element_cluster_uuid = data.nutanix_cluster.ntnxcluster.id
      storage_container          = var.prism_element_container
    }
  }
  cni_config {
    node_cidr_mask_size = var.cni_config["node_cidr_mask_size"]
    pod_ipv4_cidr       = var.cni_config["pod_ipv4_cidr"]
    service_ipv4_cidr   = var.cni_config["service_ipv4_cidr"]
    calico_config {
      ip_pool_config {
        cidr = var.cni_config["pod_ipv4_cidr"]
      }
    }
  }
  worker_node_pool {
    node_os_version = var.worker_node_pool["node_os_version"]
    num_instances   = var.worker_node_pool["num_instances"]
    ahv_config {
      cpu                        = var.worker_node_pool["cpu"]
      memory_mib                 = var.worker_node_pool["memory_mib"]
      network_uuid               = local.subnet_uuid[0]
      prism_element_cluster_uuid = data.nutanix_cluster.ntnxcluster.id
    }
  }
  etcd_node_pool {
    node_os_version = var.etcd_node_pool["node_os_version"]
    num_instances   = var.etcd_node_pool["num_instances"]
    ahv_config {
      cpu                        = var.etcd_node_pool["cpu"]
      memory_mib                 = var.etcd_node_pool["memory_mib"]
      network_uuid               = local.subnet_uuid[0]
      prism_element_cluster_uuid = data.nutanix_cluster.ntnxcluster.id
    }
  }
  master_node_pool {
    node_os_version = var.master_node_pool["node_os_version"]
    num_instances   = var.master_node_pool["num_instances"]
    ahv_config {
      cpu                        = var.master_node_pool["cpu"]
      memory_mib                 = var.master_node_pool["memory_mib"]
      network_uuid               = local.subnet_uuid[0]
      prism_element_cluster_uuid = data.nutanix_cluster.ntnxcluster.id
    }
  }
  dynamic "active_passive_config" {
    for_each = var.active_passive_config["external_ipv4_address"] == "" ? [] : [1]
    content {
      external_ipv4_address = var.active_passive_config["external_ipv4_address"]
    }
  }
}
