provider "nutanix" {
  username = var.prism_central_username
  password = var.prism_central_password
  endpoint = var.prism_central_endpoint
  insecure = var.insecure
  port = var.prism_central_port
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
  insecure = var.insecure
  prism_central_port = var.prism_central_port
  wait_timeout = var.wait_timeout
  karbon_cluster_name = var.karbon_cluster_name
  kubeconfig_filename = var.kubeconfig_filename
}

resource "null_resource" "download_manifests" {
  triggers = {
    always_run = timestamp()
    kubeconfig_filename = var.kubeconfig_filename
    knative_version = var.knative_version
  }

  provisioner "local-exec" {
    command = <<-EOT
      mkdir manifests;
      wget -P manifests https://github.com/knative/serving/releases/download/v${self.triggers.knative_version}/serving-crds.yaml;
      wget -P manifests https://github.com/knative/serving/releases/download/v${self.triggers.knative_version}/serving-core.yaml;
      wget -P manifests https://github.com/knative/eventing/releases/download/v${self.triggers.knative_version}/eventing-crds.yaml;
      wget -P manifests https://github.com/knative/eventing/releases/download/v${self.triggers.knative_version}/eventing-core.yaml;
      wget -P manifests https://github.com/knative/eventing/releases/download/v${self.triggers.knative_version}/in-memory-channel.yaml;
      wget -P manifests https://github.com/knative/eventing/releases/download/v${self.triggers.knative_version}/mt-channel-broker.yaml;
      wget -P manifests https://github.com/knative/net-istio/releases/download/v${self.triggers.knative_version}/istio.yaml;
      wget -P manifests https://github.com/knative/net-istio/releases/download/v${self.triggers.knative_version}/net-istio.yaml
    EOT
  }

  depends_on = [
    null_resource.now,
    module.karbon_kube_config
  ]
}

resource "null_resource" "install_manifests" {
  triggers = {
    always_run = timestamp()
    kubeconfig_filename = var.kubeconfig_filename
    knative_version = var.knative_version
  }
  provisioner "local-exec" {
    command = <<-EOT
      n=0
        until [ $n -ge 2 ]
        do
          kubectl --kubeconfig=${self.triggers.kubeconfig_filename} apply -f manifests && break
          n=$[$n+1]
          sleep 10
        done
    EOT
  }
  provisioner "local-exec" {
    when = destroy
    command = <<-EOT
      kubectl --kubeconfig=${self.triggers.kubeconfig_filename} delete -f manifests --ignore-not-found=true
      rm -rf manifests
    EOT
  }

  depends_on = [
    null_resource.download_manifests,
  ]
}

resource "null_resource" "setup_otel_collector" {
  triggers = {
    always_run = timestamp()
    kubeconfig_filename = var.kubeconfig_filename
    knative_version = var.knative_version
  }

  provisioner "local-exec" {
    command = <<-EOT
        kubectl create namespace metrics
        kubectl apply -f https://raw.githubusercontent.com/knative/docs/main/docs/admin/collecting-metrics/collector.yaml
        kubectl patch --namespace knative-serving configmap/config-observability --type merge --patch '{"data":{"metrics.backend-destination":"opencensus","request-metrics-backend-destination":"opencensus","metrics.opencensus-address":"otel-collector.metrics:55678"}}'
        kubectl patch --namespace knative-eventing configmap/config-observability --type merge --patch '{"data":{"metrics.backend-destination":"opencensus","metrics.opencensus-address":"otel-collector.metrics:55678"}}'
    EOT
  }

  provisioner "local-exec" {
    when = destroy
    command = <<-EOT
      kubectl --kubeconfig=${self.triggers.kubeconfig_filename} delete namespace metrics
    EOT
  }

  depends_on = [
    null_resource.install_manifests,
  ]
}

resource "null_resource" "helm_setup" {
  provisioner "local-exec" {
    command = <<EOT
    helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
    helm repo update
    EOT
  }
}

resource "null_resource" "install_prometheus" {
  triggers = {
    kubeconfig_filename             = var.kubeconfig_filename
  }

  provisioner "local-exec" {
    command = "helm install --kubeconfig=${self.triggers.kubeconfig_filename} -f monitoring/prometheus.yaml metrics prometheus-community/kube-prometheus-stack"
  }

  depends_on = [
    null_resource.install_manifests,
    null_resource.helm_setup,
    null_resource.setup_otel_collector,
  ]
}
