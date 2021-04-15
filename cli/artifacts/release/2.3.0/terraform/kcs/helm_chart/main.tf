provider "helm" {
  kubernetes {
    config_path = var.kubeconfig_filename
  }
}

resource "helm_release" "kps_helm_release" {
  name  = var.helm_release["name"]
  chart = var.helm_release["chart"]
  values = [
    file(var.helm_release["values"])
  ]
}
