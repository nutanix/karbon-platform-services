output "public_ip" {
  value       = aws_instance.kps_servicedomain_instance[*].public_ip
  description = "The public IP of the server"
}

output "ami" {
  value = aws_instance.kps_servicedomain_instance[*].ami
}

output "id" {
  description = "List of IDs of instances"
  value       = aws_instance.kps_servicedomain_instance[*].id
}