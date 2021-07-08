# Node

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**description** | **str** | Node description | [optional] 
**gateway** | **str** | Node Gateway IP address | 
**id** | **str** | ID of the entity Maximum character length is 64 for project, category, and runtime environment, 36 for other entity types. | [optional] 
**ip_address** | **str** | Node IP Address | 
**is_bootstrap_master** | **bool** |  | [optional] 
**name** | **str** | Node name. Maximum length edge name is determined by kubernetes. Name length limited to 60 and contraints are defined here https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/util/validation/validation.go | 
**role** | [**NodeRole**](NodeRole.md) |  | [optional] 
**serial_number** | **str** | Node serial number | 
**subnet** | **str** | Node subnet mask | 
**svc_domain_id** | **str** | ID of the service domain this entity belongs to | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

