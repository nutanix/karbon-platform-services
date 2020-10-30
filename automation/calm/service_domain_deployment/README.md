## Using this Blueprint
The Calm Blueprint to to deploy a Single or Multinode Service Domains on Nutanix AHV is maintained at: https://github.com/nutanix/blueprints/tree/master/kps-service-domain-deployment

### Getting Started
1. Download the required Service Domain Image named "Service Domain VM QCOW2 File for AHV" from the Nutanix Support Portal here: https://portal.nutanix.com/page/downloads?product=karbonplatformservices.
1. After logging into the Karbon Platform Services Cloud Management Console, create an API Token by clicking on your username in the upper right corner and selecting **Manage API Keys**.
1. Launch the Blueprint.

Note that creating a Multinode Service Domain requires the following additional information:
  - Additional free VirtualIP for the Service Domain
  - Prism Element Credentials
  - Data Services IP
  - Storage Container

Credits to @davehocking and @wolfganghuse for the initial development and updates.