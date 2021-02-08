variable "instance_info" {
  type = map
}

variable "cloud_info" {
  type = map
}

variable "service_domain_info" {
  type = map
}

variable "node_info" {
  type = map
}

variable "create_nutanixvolumes_storage_profile" {
  type = number
}
variable "create_ebs_storage_profile" {
  type = number
}

variable "storage_config" {
  type = map
}

variable "storage_profile_info" {
  type = map
}

variable "kps_servicedomain_instance_details" {
  type = list
}

variable "private_instance_ips" {
  type = list
}

variable "public_instance_ips" {
  type = list
}

variable "wait_for_onboarding" {
  type = bool
}