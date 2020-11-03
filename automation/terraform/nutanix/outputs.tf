output "cluster" {
  value = data.nutanix_clusters.clusters.entities.0.metadata.uuid
}

output "kps_servicedomain_instance_details" {
  value = nutanix_virtual_machine.kps_servicedomain_instance
}
