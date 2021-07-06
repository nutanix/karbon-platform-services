output "cluster_name" {
  description = "Name of the deployed Kubernetes Cluster"
  value       = module.karbon_cluster.cluster_name
}

output "kube_apiserver_ip" {
  description = "IP address of kube apiserver"
  value       = module.karbon_cluster.kube_apiserver_ip
}

output "kubernetes_version" {
  description = "Kubernetes version"
  value       = module.karbon_cluster.kubernetes_version
}

output "kubernetes_deployment_type" {
  description = "Type of Kubernetes deployment(single-master, multi-master etc..)"
  value       = module.karbon_cluster.kubernetes_deployment_type
}

output "kubeconfig_filename" {
  description = "Kubernetes config file"
  value       = module.karbon_kube_config.kubeconfig_filename
}

output "kubeconfig_content" {
  description = "Kubernetes config"
  value       = module.karbon_kube_config.kubeconfig_content
  sensitive   = true
}

output "helm_release_name" {
  description = "Helm release name"
  value       = module.helm_chart.helm_release_name
}

output "helm_release_chart" {
  description = "Name of the chart"
  value       = module.helm_chart.helm_release_chart
}

output "helm_release_version" {
  description = "Helm chart revision version"
  value       = module.helm_chart.helm_release_version
}

output "helm_release_status" {
  description = "Status of the helm release"
  value       = module.helm_chart.helm_release_status
}
