#################################################
# Karbon Platform Services Configuration
#################################################
variable "cloud_info" {
  description = "KPS cloud information for user"
  type = object({
    cloud_fqdn = string
    cloud_user_name = string
    cloud_user_pwd = string
  })
  default = {
    "cloud_fqdn" = "karbon.nutanix.com"
    "cloud_user_name" = "<cloud_username>"
    "cloud_user_pwd" = "<cloud_password>"
  }
}

#################################################
# AHV Configuration 
#################################################
variable "provider_info" {
    description = "Provider information for AHV service domain creation"
    type = object({
        username = string
        password = string
        endpoint = string
        insecure = bool
        port     = number
    })
    default = {
        "username" = "<username>"
        "password" = "<password>"
        "endpoint" = "<endpoint>"
        "insecure" = true
        "port"     = 9440
    }
}

variable "instance_info" {
    description = "AHV instance information"
    type = object({
        instance_count = number
        instance_name_prefix = string
    })
    default = {
        instance_count = 1
        instance_name_prefix = "kps_instance"
    }
}

variable "nutanix_vm_config" {
    description = "AHV  Virtual Machine configuration"
    type = object({
        description = string
        num_vcpus_per_socket = number
        num_sockets          = number
        memory_size_mib      = number
    })
    default = {
        "description" = ""
        "num_vcpus_per_socket" = 2
        "num_sockets"          = 1
        "memory_size_mib"      = 4096
    }
}

variable "image_config" {
    description = "Nutanix KPS Service Domain Node OS Image configuration"
    type = object({
        name        = string
        description = string
        source_path  = string
        source_http = string
    })
    default = {
        "name"        = "kps_servicedomain_image"
        "description" = "kps_servicedomain_image"
        "source_path"  = "<path to qcow2 file>"
        "source_http" = "<http uri to qcow file: Avoids the need to download and upload the file>"
    }
}

#################################################
# Karbon Platform Services Configuration
#################################################
variable "cloud_info" {
    description = "KPS cloud information"
    type = object({
        cloud_fqdn = string
        cloud_user_name = string
        cloud_user_pwd = string
    })
    default = {
        "cloud_fqdn" = "karbon.nutanix.com"
        "cloud_user_name" = "<cloud_username>"
        "cloud_user_pwd" = "<cloud_password>"
    }
}

variable "service_domain_info" {
    description = "KPS service domain information"
    type = object({
        sd_name: string
        sd_description: string
        sd_virtual_ip: string
    })
    default = {
        "sd_name": "sd1.0"
        "sd_description": "sd created thru tf"
        "sd_virtual_ip": "x.x.x.x"
    }
}

variable "node_info" {
    description = "KPS node information"
    type = object({
        node_gateway: string
        node_subnet: string
    })
    default = {
        "node_gateway": "x.x.x.x"
        "node_subnet": "x.x.x.x"
    }
}

variable "wait_for_onboarding" {
    description = "Set to true to wait synchronously for node onboarding to complete."
    type = bool
    default = false
}

#################################################
# Nutanix Storage Profile Configuration
#################################################
variable "create_storage_profile" {
    description = "Number of storage profiles to create, value an be 1 or 0 for yes/no"
    type = number
    default = 1
}

variable "nutanix_volumes_config" {
    description = "AHV Volume configuration"
    type = object({
        dataServicesIP: string
        dataServicesPort: string
        flashMode: string
        prismElementClusterPort: string
        prismElementClusterVIP: string
        storageContainerName: string
        prismElementUserName: string
        prismElementPassword: string
    })
    default = {
        "dataServicesIP": "x.x.x.x"
        "dataServicesPort": "3260"
        "flashMode": "false"
        "prismElementClusterPort": "9440"
        "prismElementClusterVIP": "x.x.x.x"
        "storageContainerName": "<container name>"
        "prismElementUserName": "<username>"
        "prismElementPassword": "<password>"
    }
}

variable "storage_profile_info" {
    description = "AHV Storage Profile configuration"
    type = object({
        type: string
        name: string
        isDefault: string
    })
    default = {
        "type": "NutanixVolumes"
        "name": "tf_storage_profile"
        "isDefault": "true"
    }
}