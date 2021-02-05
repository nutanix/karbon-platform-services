#################################################
# AWS Profile Set Up
#################################################
variable "region" {
  description = "AWS region"
  type = string
  default = "us-west-2"
}

variable "availability_zone" {
  description = "AWS availability zone"
  type = string
  default = "us-west-2c"
}

variable "profile" {
  description = "AWS profile"
  type = string
  default = "default"
}

variable "environment" {
  description = "AWS environment"
  type = string
  default = "Dev"
}
#################################################
# AWS EC2 Instance Configuration
#################################################
variable "security_group" {
  description = "AWS Security Group to attach to EC2 instance"
  type = string
  default = "sg-xxxxxxxx"
}

variable "kps_raw_diskimage_version" {
  description = "version of kps raw image"
  type = number
  default = 782
}

variable "instance_info" {
  description = "EC2 instance description"
   type = object({
    instance_count = number
    instance_name_prefix = string
  })
  default = {
    "instance_count" = 3
    "instance_name_prefix" = "kps_instance"    
  }
}

variable "ec2_vm_config" {
  description = "EC2 instance configuration"
  type = object({
    instance_type = string
  })
  default = {
    "instance_type" = "t2.2xlarge"
  }
}

variable "iam_config" {
    description = "AWS IAM configuration"
    type = object({
    aws_iam_role_name = string
    aws_iam_policy_name = string
    aws_iam_instance_profile_name = string
  })
  default = {
    "aws_iam_role_name" = "sam_ebs_role_tf"
    "aws_iam_policy_name" = "sam_role_policy_tf"
    "aws_iam_instance_profile_name" = "sam_instance_profile_tf"
  }
}
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

variable "service_domain_info" {
  description = "KPS service domain information"
  type = object({
    sd_name: string
    sd_description: string
    sd_virtual_ip: string
  })
  default = {
    "sd_name": "awstf1"
    "sd_description": "aws sd created thru tf"
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
  description = "Set to true for synchronous onboarding"
  type = bool
  default = false
}

#################################################
# AWS Storage Profile Configuration
#################################################
variable "create_storage_profile" {
  description = "Number of AWS Storage Profiles to create"
  type = number
  default = 1
}

variable "storage_profile_info" {
  description = "AWS Storage Profile information"
  type = object({
    type: string
    name: string
    isDefault: string
  })
  default = {
    "type": "EBS"
    "name": "tf_sp"
    "isDefault": "true"
  }
}
#################################################
# AWS EC2 EBS Volume Configuration
#################################################
variable "ebs_storage_config" {
  description = "Configuration for AWS EBS Storage Profile to attach to EC2 instance"
  type = object({
    iops_per_gb: string
    type: string
  })
  default = {
    "iops_per_gb": "10"
    "type": "gp2"
  }
}

variable "data_partition_size_gb" {
  description = "Partition size for AWS EBS Volume"
  type = number
  default = 100
}