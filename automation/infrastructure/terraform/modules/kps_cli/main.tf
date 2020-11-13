resource "null_resource" "install_kps_cli" {
  provisioner "local-exec" {
    command = <<EOT
      wget https://kps-cli.s3-us-west-2.amazonaws.com/kps-v1.0.0-rc10-darwin_amd64.tar.gz
      mkdir -p kps && tar zxvf kps-v1.0.0-rc10-darwin_amd64.tar.gz -C kps
      kps/kps config create-context local_user_ctx_1 --email ${var.cloud_info["cloud_user_name"]} --password ${var.cloud_info["cloud_user_pwd"]} -u ${var.cloud_info["cloud_fqdn"]}
      EOT
  }
}

resource "null_resource" "kps_cli_tasks" {
  triggers  =  {
    always_run = "${timestamp()}"
  }

  provisioner "local-exec" {
    command = <<EOT
      mkdir -p ${path.module}/generated
      kps/kps get svcdomain > ${path.module}/generated/servicedomains.txt
      kps/kps get node > ${path.module}/generated/nodes.txt
      kps/kps get project > ${path.module}/generated/projects.txt
      kps/kps get application > ${path.module}/generated/applications.txt

      EOT
  }

  depends_on = [
    null_resource.install_kps_cli,
  ]
}