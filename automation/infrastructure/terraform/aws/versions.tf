terraform {
  required_providers {
    nutanix = {
      source = "nutanix/nutanix"
    }
    nutanixkps = {
      source = "nutanix/nutanixkps"
    }
  }
  required_version = ">= 0.14"
}