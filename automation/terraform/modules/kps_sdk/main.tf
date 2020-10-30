resource "null_resource" "install_kps_sdk" {
  provisioner "local-exec" {
    command = <<EOT

      tar zxvf ${path.module}/kps_sdk_release/kps_api_1.0.1219.tar.gz
      pushd ${path.module}/kps_api
      python3 setup.py install
      python3 -c "import kps_api"
      popd
      EOT
  }
}

resource "null_resource" "kps_sdk_tasks" {
  triggers  =  {
    always_run = "${timestamp()}"
  }

  provisioner "local-exec" {
    command = <<EOT
      mkdir -p ${path.module}/generated
      python3 ${path.module}/python_cli/service_domain_client.py > ${path.module}/generated/service_domain_client_output.txt
      EOT
  }

  depends_on = [
    null_resource.install_kps_sdk,
  ]
}