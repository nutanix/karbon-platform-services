#https://learn.hashicorp.com/tutorials/terraform/aws-variables

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 2.70"
    }
  }
}

provider "aws" {
    profile = var.profile
    region = var.region
}

data "aws_security_group" "kps_security_group" {
  id = var.security_group
}

data "aws_vpc" "kps_vpc" {
  id = data.aws_security_group.kps_security_group.vpc_id
}

resource "aws_iam_role" "role" {
  name        = "ebs_role_trf"
  force_detach_policies = true
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "policy" {
  name        = "role-policy-trf"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ec2:AttachVolume",
        "ec2:CreateSnapshot",
        "ec2:CreateTags",
        "ec2:CreateVolume",
        "ec2:DeleteSnapshot",
        "ec2:DeleteTags",
        "ec2:DeleteVolume",
        "ec2:DescribeAvailabilityZones",
        "ec2:DescribeInstances",
        "ec2:DescribeSnapshots",
        "ec2:DescribeTags",
        "ec2:DescribeVolumes",
        "ec2:DescribeVolumesModifications",
        "ec2:DetachVolume",
        "ec2:ModifyVolume"
      ],
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_iam_policy_attachment" "policy-attach" {
  name       = "ebs-policy-attachment"
  roles      = [aws_iam_role.role.name]
  policy_arn = aws_iam_policy.policy.arn
}

resource "aws_iam_instance_profile" "instance_profile" {
  name = "instance_profile_trf"
  role = aws_iam_role.role.name
}

resource "aws_instance" "kps_servicedomain_instance" {
  ami = var.amis[var.region]
  instance_type = var.ec2_vm_config["instance_type"]
  security_groups = [data.aws_security_group.kps_security_group.name]
  iam_instance_profile = aws_iam_instance_profile.instance_profile.name
  count = var.instance_info["instance_count"]
  tags = {
    Name = join("-", [var.instance_info["instance_name_prefix"], count.index])
  }
}

resource "aws_ebs_volume" "kps_volume" {
  availability_zone = var.availability_zone
  size = var.data_partition_size_gb
  count = var.instance_info["instance_count"]
}

resource "aws_volume_attachment" "kps_volume_attachment" {
  device_name = "/dev/xvdf"
  volume_id = aws_ebs_volume.kps_volume[count.index].id
  count = var.instance_info["instance_count"]
  instance_id = aws_instance.kps_servicedomain_instance.*.id[count.index]
  force_detach = true
  depends_on = [
    aws_ebs_volume.kps_volume,
    aws_instance.kps_servicedomain_instance
  ]
}

module "service_domain" {
  source = "../modules/service_domain"
  depends_on = [
    aws_volume_attachment.kps_volume_attachment
  ]
  instance_info = var.instance_info
  cloud_info = var.cloud_info
  node_info = var.node_info
  storage_config = var.ebs_storage_config
  service_domain_info = var.service_domain_info
  storage_profile_info = var.storage_profile_info
  kps_servicedomain_instance_details = aws_instance.kps_servicedomain_instance
  create_ebs_storage_profile = var.create_storage_profile
  create_nutanixvolumes_storage_profile = 0
  private_instance_ips = aws_instance.kps_servicedomain_instance[*].private_ip
  public_instance_ips = aws_instance.kps_servicedomain_instance[*].public_ip
}