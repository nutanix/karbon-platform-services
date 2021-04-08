output "cluster_name" {
  description = "Name of the deployed Kubernetes Cluster"
  value       = nutanix_karbon_cluster.cluster.name
}

output "kube_apiserver_ip" {
  description = "IP address of kube apiserver"
  value       = nutanix_karbon_cluster.cluster.kubeapi_server_ipv4_address
}

output "kubernetes_version" {
  description = "Kubernetes version"
  value       = nutanix_karbon_cluster.cluster.version
}

output "kubernetes_deployment_type" {
  description = "Type of Kubernetes deployment(single-master, multi-master etc..)"
  value       = nutanix_karbon_cluster.cluster.deployment_type
}

output "helm_release_name" {
  description = "Helm release name"
  value       = helm_release.kps_helm_release.name
}

output "helm_release_chart" {
  description = "Name of the chart"
  value       = helm_release.kps_helm_release.chart
}

output "helm_release_version" {
  description = "Helm chart revision version"
  value       = helm_release.kps_helm_release.version
}

output "helm_release_status" {
  description = "Status of the helm release"
  value       = helm_release.kps_helm_release.status
}
