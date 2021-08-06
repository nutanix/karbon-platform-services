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

variable "kubeconfig_filename" {
  type        = string
  description = "Path to Kubernetes config file"
  default     = "kube.config"
}

variable "kubeflow_version" {
  type        = string
  description = "version of kubeflow"
  default     = "1.3.0"
}
