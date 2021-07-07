# Deploying infrastructure (e.g Kubernetes Clusters, Helm Charts etc..) with kps cli

## What infrastructure can the cli deploy?
Currently the cli will deploy a KCS(Karbon Container Services) cluster on Nutanix AHV. It will also deploy KPS(Karbon Platform Services) on top of the deployed KCS cluster.

## What are the options to be specified on the cli in order to deploy infrastructure?
The cli primarily uses the `svcdomain` subcommand to create/modify infrastructure(e.g. `kps create/update/delete svcdomain` commands. The `-p`(provider) option *has to be specified* to interact with infrastructure. e.g: `kps create svcdomain -p kcs -n <cluster_name>...` will create a Service Domain on a KCS cluster.

cli options can be specified on the command line completely or all options can be specified in a yaml file like the one displayed below. The values specified in the yaml file will override any command line/default options. 

Invoking create via command line:
- `kps create svcdomain -n <cluster_name> -p kcs -f <provider>-config.yaml`

Sample YAML input file for a KCS cluster:
```
# configuration for a KCS cluster.

#############################################################
# REQUIRED PARAMETERS:
# This section contains required parameters to create a KCS
# cluster. All parameters must have a valid value.
#############################################################
---
kind: kcsconfig
prismCentral:
  # prism central ip address/hostname 
  host: ""  
  password: ""
  username: ""
prismElement:
  cluster: ""
  storageContainer: ""
  subnet: ""
clusterVIP: ""

#############################################################
# OPTIONAL PARAMETERS:
# This section contains optional parameters to customize a
# KCS cluster.
#############################################################
masterPool:
  ahvConfig:
    cpu: ~
    diskMib: ~
    memoryMib: ~
  numInstances: ~
  nodeOsVersion: ~
  clusterVIP: ~
etcdPool:
  ahvConfig:
    cpu: ~
    diskMib: ~
    memoryMib: ~
  numInstances: ~
  nodeOsVersion: ~
workerPool:
  ahvConfig:
    cpu: ~
    diskMib: ~
    memoryMib: ~
  numInstances: ~
  nodeOsVersion: ~
helmConfig:
  name: ~
  chart: ~
kpsVersion: ~
```

The YAML is divided into 2 sections:
* Required parameters:
  All fields in this section should have valid values. Cluster deployment may not be sucessful if incorrect parameters are supllied in this section.
* Optional parameters:
  All fields in this section have default values supplied by the cli. The default values are recommended for deploying a production cluster. In cases where a user wants to supply custom values, a valid value has to supplied for the fields that have to be modified. Values provided in this YAML file will override all defaults. (e.g if you want to increase the number of worker nodes, provide a number in the `numInstances` field of the `workerPool` section of the file and run `kps create svcdomain -p kcs -n <cluster_name> -f kcs-config.yaml` or if you want to update KPS version/helm chart, provide a value for the `kpsVersion` or `chart` in the `helmConfig` section and then run `kps update svcdomain <cluster_name> -f kcs-config.yaml`)


## Sample commands:

### Create a service domain
* `kps create svcdomain -p kcs -n <cluster_name> --pc_username <> --pc_password <> --pc_host 10.X.X.X --pe_cluster <> --pe_storage_container <> --pe_subnet <> --external_ip 10.X.X.X`
* `kps create svcdomain -p kcs -n <cluster_name> -f kcs-config.yaml`

### Update a service domain
* `kps update svcdomain <cluster_name> -p kcs --version 2.3.0`
* `kps update svcdomain <cluster_name> -p kcs --helm_url <https://...>`
* `kps update svcdomain <cluster_name> -p kcs -f kcs-config.yaml`

### Delete a service domain
* `kps delete svcdomain <cluster_name> -p kcs`
