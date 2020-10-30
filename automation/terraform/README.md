# Deploying a Service Domain with KPS

Enter the cloud provider directory of your choice. You will see a file called `variables.tf` where you can configure values such as provider info, cloud info, image config, etc.

### Workflow

#### Download Terraform Plugins
<pre>terraform init</pre>

#### Create Deployment Plan
<pre>terraform plan</pre>

#### Apply the Infrastructure Plan
<pre>terraform apply</pre>

#### Teardown Deployed Infrastructure
<pre>terraform destroy</pre>


## References

### Nutanix's Terraform Provider
https://www.youtube.com/watch?v=V8_Lu1mxV6g 

### Terraform Documentation
https://www.terraform.io/docs/providers/nutanix/index.html 