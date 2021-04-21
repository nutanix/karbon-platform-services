variable "prism_central_username" {
  type        = string
  description = "Administrator user name for Prism Central"
  sensitive   = true
}

variable "prism_central_password" {
  type        = string
  description = "Password for Prism Central"
  sensitive   = true
}

variable "prism_central_endpoint" {
  type        = string
  description = "Prism Central Endpoint(e.g PC ip address)"
}

variable "prism_element_cluster" {
  type        = string
  description = "Prism Element Cluster name"
}

variable "prism_element_container" {
  type        = string
  description = "Prism Element Container name"
}

variable "prism_element_subnet" {
  type        = string
  description = "Prism Element Subnet name"
}

variable "insecure" {
  type    = bool
  default = true
}

variable "prism_central_port" {
  type        = number
  description = "Prism Central Port Number(default: 9440)"
  default     = 9440
}

variable "wait_timeout" {
  type        = number
  description = "Timeout for Prism Central connection"
  default     = 10
}

variable "karbon_cluster_name" {
  type        = string
  description = "Karbon Kubernetes Cluster Name"
  default     = "kcs-cluster"
}

variable "kubernetes_version" {
  type        = string
  description = "Karbon Kubernetes Version"
  default     = "1.18.15-1"
}

variable "pv_reclaim_policy" {
  type        = string
  description = "Karbon PV Reclaim Policy"
  default     = "Delete"
}

variable "cni_config" {
  description = "CNI Plugin Details (default: Calico)"
  type = object({
    node_cidr_mask_size = number
    pod_ipv4_cidr       = string
    service_ipv4_cidr   = string
  })
  default = {
    node_cidr_mask_size = 24
    pod_ipv4_cidr       = "172.20.0.0/16"
    service_ipv4_cidr   = "172.19.0.0/16"
  }
}

variable "worker_node_pool" {
  description = "Worker Node Pool Details"
  type = object({
    node_os_version = string
    num_instances   = number
    cpu             = number
    memory_mib      = number
  })
  default = {
    node_os_version = "ntnx-1.0"
    num_instances   = 3
    cpu             = 8
    memory_mib      = 16384
  }
}

variable "etcd_node_pool" {
  description = "ETCD Node Pool Details"
  type = object({
    node_os_version = string
    num_instances   = number
    cpu             = number
    memory_mib      = number
  })
  default = {
    node_os_version = "ntnx-1.0"
    num_instances   = 3
    cpu             = 4
    memory_mib      = 8192
  }
}

variable "master_node_pool" {
  description = "Master Node Pool Details"
  type = object({
    node_os_version = string
    num_instances   = number
    cpu             = number
    memory_mib      = number
  })
  default = {
    node_os_version = "ntnx-1.0"
    num_instances   = 2
    cpu             = 4
    memory_mib      = 8192
  }
}

variable "active_passive_config" {
  description = "Virtual IP Address"
  type = object({
    external_ipv4_address = string
  })
  default = {
    external_ipv4_address = ""
  }
}

variable "helm_release" {
  description = "Helm release configuration"
  type = object({
    name   = string
    chart  = string
    values = string
  })
  default = {
    name   = "servicedomain"
    chart  = ""
    values = "values.yaml"
  }
}

variable "kubeconfig_filename" {
  type        = string
  description = "Path to Kubernetes config file"
  default     = "kube.config"
}
