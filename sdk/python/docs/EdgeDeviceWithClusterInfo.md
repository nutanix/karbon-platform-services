# EdgeDeviceWithClusterInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**bootstrap_master_ssh_public_key** | **str** |  | [optional] 
**cluster_id** | **str** | ID of the cluster this entity belongs to | 
**description** | **str** | EdgeDevice description | [optional] 
**gateway** | **str** | Edge Device Gateway IP address | 
**id** | **str** | ID of the entity Maximum character length is 64 for project, category, and runtime environment, 36 for other entity types. | [optional] 
**ip_address** | **str** | Edge device IP Address | 
**is_bootstrap_master** | **bool** |  | 
**name** | **str** | Edge name. Maximum length edge name is determined by kubernetes. Name length limited to 60 as node name is the edge name plus a suffix. https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/util/validation/validation.go | 
**role** | [**NodeRole**](NodeRole.md) |  | [optional] 
**serial_number** | **str** | Edge device serial number | 
**subnet** | **str** | Edge subnet mask | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

