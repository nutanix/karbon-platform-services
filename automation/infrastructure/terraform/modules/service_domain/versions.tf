terraform {
  required_providers {
    nutanixkps = {
      version = "~> 0.1"
      source = "hashicorp.com/nutanix/nutanixkps"
    }
  }
  required_version = ">= 0.13"
}
