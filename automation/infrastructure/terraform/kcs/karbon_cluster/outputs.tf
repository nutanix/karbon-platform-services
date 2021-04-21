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
