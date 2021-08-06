# TF_LOG=DEBUG 
terraform init
#terraform plan -var-file=testing.tfvars --auto-approve
terraform apply -var-file=testing.tfvars --auto-approve
# terraform destroy -var-file=testing.tfvars
cp ./*.cfg ~/Downloads/
#terraform show -json | jq -r '.values.root_module.resources[0].values.http_clone_url'