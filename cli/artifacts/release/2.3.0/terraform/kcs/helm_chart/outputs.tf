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
