provider "nutanix" {
  username     = var.prism_central_username
  password     = var.prism_central_password
  endpoint     = var.prism_central_endpoint
  insecure     = var.insecure
  port         = var.prism_central_port
  wait_timeout = var.wait_timeout
}

# Fake resource for triggering fetching of latest kube.config
resource "null_resource" "now" {
  triggers = {
    always_run = timestamp()
  }
}

# Get Kubeconfig by cluster name
data "nutanix_karbon_cluster_kubeconfig" "configbyname" {
  depends_on = [
    null_resource.now
  ]
  karbon_cluster_name = var.karbon_cluster_name
}

# Generate kube.config
resource "local_file" "kubeconfig" {
  depends_on = [
    null_resource.now
  ]
  sensitive_content = templatefile("${path.module}/kube.config.tpl", {
    cluster_name           = data.nutanix_karbon_cluster_kubeconfig.configbyname.name
    host                   = data.nutanix_karbon_cluster_kubeconfig.configbyname.cluster_url
    cluster_ca_certificate = data.nutanix_karbon_cluster_kubeconfig.configbyname.cluster_ca_certificate
    token                  = data.nutanix_karbon_cluster_kubeconfig.configbyname.access_token
  })
  filename             = var.kubeconfig_filename
  file_permission      = "0600"
  directory_permission = "0600"
}