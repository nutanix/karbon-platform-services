provider "nutanix" {
  username     = var.prism_central_username
  password     = var.prism_central_password
  endpoint     = var.prism_central_endpoint
  insecure     = var.insecure
  port         = var.prism_central_port
  wait_timeout = var.wait_timeout
}

resource "null_resource" "now" {
  triggers = {
    always_run = timestamp()
  }
}

# Download kubeconfig for input cluster
module "karbon_kube_config" {
  source = "../karbon_kube_config"

  prism_central_username = var.prism_central_username
  prism_central_password = var.prism_central_password
  prism_central_endpoint = var.prism_central_endpoint
  insecure               = var.insecure
  prism_central_port     = var.prism_central_port
  wait_timeout           = var.wait_timeout
  karbon_cluster_name    = var.karbon_cluster_name
  kubeconfig_filename    = var.kubeconfig_filename
}

# provider "kubernetes" {
#   config_path = var.kubeconfig_filename
# }

# # Get all k8s namespaces
# data "kubernetes_all_namespaces" "allns" {}

# output "all-ns" {
#   value = data.kubernetes_all_namespaces.allns.namespaces
# }

# output "kubeflow-ns-present" {
#   value = contains(data.kubernetes_all_namespaces.allns.namespaces, "kubeflow")
# }

# Configure the GitHub Provider
# provider "github" {}

# data "github_repository" "kubeflow_manifests" {
#   full_name = "kubeflow/manifests"
# }

# locals {
#   http_clone_url = data.github_repository.kubeflow_manifests.http_clone_url
# }

resource "null_resource" "download_manifest_repo" {

  triggers = {
    always_run          = timestamp()
    kubeconfig_filename = var.kubeconfig_filename
    kubeflow_version    = var.kubeflow_version
  }
  provisioner "local-exec" {
    # command = "wget https://github.com/kubeflow/manifests/archive/refs/heads/master.zip && unzip -n ./master.zip"
    command = "wget https://github.com/kubeflow/manifests/archive/refs/tags/v${self.triggers.kubeflow_version}.zip && mkdir -p kubeflow && unzip -n ./v${self.triggers.kubeflow_version}.zip -d ./kubeflow"
  }

  # provisioner "local-exec" {
  #   when = destroy
  #   # command = "rm ./master.zip && rm -rf manifests-master"
  #   command = "rm ./v${self.triggers.kubeflow_version}.zip && rm -rf kubeflow"
  # }
  depends_on = [
    null_resource.now,
    module.karbon_kube_config
  ]
}

# TODO Enable following block after error once following error is resolved----------
# │ Error: Invalid for_each argument
# │ 
# │   on main.tf line 72, in resource "kustomization_resource" "install_kubeflow":
# │   72:   for_each = tolist(data.kustomization_build.kubeflow_manifests.ids)
# │     ├────────────────
# │     │ data.kustomization_build.kubeflow_manifests.ids is a set of string, known only after apply
# │ 
# │ The "for_each" value depends on resource attributes that cannot be determined until apply, so Terraform cannot predict how many instances will be created. To work around this, use the -target argument to
# │ first apply only the resources that the for_each depends on.

# provider "kustomization" {
#   kubeconfig_path = var.kubeconfig_filename
# }

# data "kustomization_build" "kubeflow_manifests" {
#   path = "manifests-master/example"

#   depends_on = [
#     null_resource.download_manifest_repo
#   ]
# }

# resource "kustomization_resource" "install_kubeflow_kustomize" {
#   for_each = tolist(data.kustomization_build.kubeflow_manifests.ids)

#   manifest = data.kustomization_build.kubeflow_manifests.manifests[each.value]
# }
# TODO ----------

resource "null_resource" "install_kubeflow_local_exec" {
  triggers = {
    always_run          = timestamp()
    kubeconfig_filename = var.kubeconfig_filename
    kubeflow_version    = var.kubeflow_version
  }
  provisioner "local-exec" {
    command = "while ! kustomize build kubeflow/manifests-${self.triggers.kubeflow_version}/example | kubectl --kubeconfig=${self.triggers.kubeconfig_filename} apply -f -; do echo 'Retrying to apply resources'; sleep 10; done"
  }

  # provisioner "local-exec" {
  #   when    = destroy
  #   command = "while ! kustomize build kubeflow/manifests-${self.triggers.kubeflow_version}/example | kubectl --kubeconfig=${self.triggers.kubeconfig_filename} delete -f -; do echo 'Retrying to delete resources'; sleep 10; done"
  # }

  depends_on = [
    null_resource.download_manifest_repo
  ]
}
