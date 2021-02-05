terraform {
  required_providers {
    nutanix = {
      source = "terraform-providers/nutanix"
    }
    nutanixkps = {
      source = "hashicorp.com/nutanix/nutanixkps"
    }
  }
  required_version = ">= 0.13"
}
