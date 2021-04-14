variable "helm_release" {
  description = "Helm release configuration"
  type = object({
    name    = string
    chart   = string
    values  = string
  })
  default = {
    name = "servicedomain"
    # temporary location until we publish helm charts to a well known address
    chart = "http://uranus.corp.nutanix.com/~kevin.thomas/helm/servicedomain-2.3.0.tgz"
    values = "values.yaml"
  }
}

variable "kubeconfig_filename" {
  description = "Kubernetes config"
  type = string
  default = "kube.config"
}