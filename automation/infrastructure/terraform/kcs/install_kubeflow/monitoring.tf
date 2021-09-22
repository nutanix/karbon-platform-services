resource "null_resource" "install_monitoring" {
  depends_on = [null_resource.now,
    null_resource.helm_setup,
    null_resource.prepare_prometheus,
    null_resource.install_prometheus]
}

resource "null_resource" "helm_setup" {
  count = var.kubeflow_monitoring == "prometheus" ? 1 : 0

  provisioner "local-exec" {
    command = <<EOT
    helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
    helm repo update
    EOT
  }
}

resource "null_resource" "prepare_prometheus" {
  triggers = {
    kubeflow_monitoring            = var.kubeflow_monitoring
    kubeconfig_filename             = var.kubeconfig_filename
  }

  count = var.kubeflow_monitoring == "prometheus" ? 1 : 0

  provisioner "local-exec" {
    command = "kubectl --kubeconfig=${self.triggers.kubeconfig_filename} apply -f monitoring.ns.yaml"
  }

  depends_on = [
    null_resource.now,
    module.karbon_kube_config
  ]
}

resource "null_resource" "install_prometheus" {
  triggers = {
    kubeflow_monitoring            = var.kubeflow_monitoring
    kube_prometheus_stack_version = var.kube_prometheus_stack_version
    kubeconfig_filename             = var.kubeconfig_filename
  }

  count = var.kubeflow_monitoring == "prometheus" ? 1 : 0

  provisioner "local-exec" {
    command = "helm install --kubeconfig=${self.triggers.kubeconfig_filename} -f monitoring.env.yaml --version ${self.triggers.kube_prometheus_stack_version} kubeflow-monitoring prometheus-community/kube-prometheus-stack"
  }

  depends_on = [
    null_resource.now,
    module.karbon_kube_config,
    null_resource.helm_setup,
    null_resource.prepare_prometheus
  ]
}
