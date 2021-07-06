# Configure Providers
terraform {
  required_version = ">= 0.13"
  required_providers {
    nutanix = {
      source  = "nutanix/nutanix"
      version = "1.2.0"
    }
    helm = {
      source  = "hashicorp/helm"
      version = "2.0.2"
    }
  }
}
