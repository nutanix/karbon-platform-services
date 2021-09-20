resource "null_resource" "install_monitoring" {
  depends_on = [null_resource.now, null_resource.install_kubeflow_local_exec, null_resource.install_prometheus]
}

resource "null_resource" "install_prometheus" {
  triggers = {
    kubeflow_monitoring          = var.kubeflow_monitoring
    prometheus_operator_version = var.prometheus_operator_version
    kubeconfig_filename           = var.kubeconfig_filename
  }

  count = var.kubeflow_monitoring == "prometheus" ? 1 : 0

  provisioner "local-exec" {
    command = "while ! kustomize build prometheus/ | kubectl --kubeconfig=${self.triggers.kubeconfig_filename} apply -f -; do echo 'Retrying to apply resources'; sleep 10; done"
  }

  depends_on = [null_resource.download_prometheus_repo]
}

resource "null_resource" "download_prometheus_repo" {
  triggers = {
    kubeflow_monitoring          = var.kubeflow_monitoring
    prometheus_operator_version = var.prometheus_operator_version
  }
  count = var.kubeflow_monitoring == "prometheus" ? 1 : 0

  provisioner "local-exec" {
    command = "wget https://github.com/prometheus-operator/prometheus-operator/archive/refs/tags/v${self.triggers.prometheus_operator_version}.zip && unzip -n ./v${self.triggers.prometheus_operator_version}.zip && mv prometheus-operator-${self.triggers.prometheus_operator_version}/ prometheus"
  }



  depends_on = [null_resource.now]
}
