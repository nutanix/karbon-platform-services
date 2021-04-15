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
  description = "Kubernetes config"
  type        = string
  default     = "kube.config"
}