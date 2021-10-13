resource "null_resource" "install_monitoring" {
  depends_on = [
    null_resource.now,
    null_resource.install_prometheus_helm
  ]
}

resource "null_resource" "helm_setup" {
  count = var.install_prometheus ? 1 : 0

  provisioner "local-exec" {
    command = <<EOT
    helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
    helm repo update
    EOT
  }
}

resource "null_resource" "create_monitoring_namespace" {
  triggers = {
    kubeconfig_filename             = var.kubeconfig_filename
  }

  count = var.install_prometheus ? 1 : 0

  provisioner "local-exec" {
    command = "kubectl --kubeconfig=${self.triggers.kubeconfig_filename} create namespace kubeflow-monitoring"
  }

  depends_on = [
    null_resource.now,
    module.karbon_kube_config
  ]
}

resource "null_resource" "install_prometheus_helm" {
  triggers = {
    kube_prometheus_stack_version = var.kube_prometheus_stack_version
    kubeconfig_filename           = var.kubeconfig_filename
  }

  count = var.install_prometheus ? 1 : 0

  provisioner "local-exec" {
    command = "helm upgrade --install --kubeconfig=${self.triggers.kubeconfig_filename} -f monitoring.env.yaml --version ${self.triggers.kube_prometheus_stack_version} kubeflow-monitoring prometheus-community/kube-prometheus-stack"
  }

  depends_on = [
    null_resource.now,
    module.karbon_kube_config,
    null_resource.helm_setup,
    null_resource.create_monitoring_namespace
  ]
}
