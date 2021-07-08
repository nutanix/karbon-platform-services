# KubernetesClusterCert

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ca_certificate** | **str** | Root CA certificate for the tenant. | 
**certificate** | **str** | Certificate for the kubernetes cluster using old/fixed root CA. | 
**id** | **str** | ID of the entity Maximum character length is 64 for project, category, and runtime environment, 36 for other entity types. | [optional] 
**kubernetes_cluster_certificate** | **str** | Certificate for the kubernetes cluster using per-tenant root CA. | 
**kubernetes_cluster_id** | **str** | ID of the kubernetes cluster this entity belongs to | 
**kubernetes_cluster_private_key** | **str** | Encrypted private key using per-tenant root CA. | 
**locked** | **bool** |  | [optional] 
**private_key** | **str** | Encrypted private key using old/fixed root CA. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

