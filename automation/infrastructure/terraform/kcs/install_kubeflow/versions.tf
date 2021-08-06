terraform {
  required_providers {
    nutanix = {
      source = "nutanix/nutanix"
    }
    # kustomization = {
    #   source  = "kbst/kustomization"
    # }   
    # github = {
    #   source  = "integrations/github"
    # }     
  }
  required_version = ">= 0.14"
}
