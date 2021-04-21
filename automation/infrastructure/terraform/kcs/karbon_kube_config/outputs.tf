output "kubeconfig_filename" {
  description = "Path to Kubernetes config file"
  value       = local_file.kubeconfig.filename
}

output "kubeconfig_content" {
  description = "Kubernetes config"
  value       = local_file.kubeconfig.sensitive_content
  sensitive   = true
}
