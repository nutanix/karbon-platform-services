
# Deploy a Karbon Kubernetes cluster:

[Nutanix Kubernetes cluster provider details:] (https://registry.terraform.io/providers/nutanix/nutanix/latest/docs)

## Steps:
1. Clone the repo and cd into the "kcs" directory where the tf files are located
2. Run `terraform init` (this will download the nutanix terraform plugin)
3. Create/download a custom tfvars file if you don't want to be prompted for inputs. [See samples here:] (http://uranus.corp.nutanix.com/~kevin.thomas/terraform/)
4. To validate your configuration run:
  - `terraform fmt -diff=true`
  - `terraform validate`
5. Run `terraform plan -var-file=<custom>.tfvars` or
       `terraform plan` (This will prompt for input)
6. Run `terraform apply -var-file=<custom>.tfvars` or
       `terraform apply` (This will prompt for input)

This will deploy a Single Master, 3 Worker Node KCS cluster and will output the kubeconfig file.

## Steps to deploy KPS on KCS:
1. Using KPS cli run the following commands:
   - `kps create svcdomain -n <NAME> -t SERVICE_DOMAIN_CLUSTER`
   - `kps get svcdomaincredentials -n <NAME> -u -o helm > <PATH_TO_TERRAFORM_FOLDER>/values.yaml` (credentials are only valid for 30 minutes)
2. `terraform init`
3. `terraform apply -auto-approve -var-file kcs.tfvars -target module.karbon_kube_config`
4. `terraform apply -auto-approve -var-file kcs.tfvars` (http://uranus.corp.nutanix.com/~kevin.thomas/terraform/)

This will deploy a Single Master, 3 Worker Node KCS cluster and deploy KPS on the cluster.

### To view the current state of deployment run:
- `terraform show`
### To view the deployed resources run:
-  `terraform state list`

### To destroy the cluster and deploed resources run:
-  `terraform apply -auto-approve -var-file kcs.tfvars -target module.karbon_kube_config`
-  `terraform destroy -auto-approve -var-file kcs.tfvars`

### To retrieve kubeconfig and get resources of any cluster
- `terraform apply -auto-approve -var-file kcs.tfvars -target module.karbon_kube_config`
- `kubectl --kubeconfig kube.config get pod`
