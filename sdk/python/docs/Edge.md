# Edge

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**connected** | **bool** | Determines if the edge is currently connected to XI IoT management services. | [optional] 
**description** | **str** | Edge description | [optional] 
**edge_devices** | **float** | Number of devices (nodes) in this edge | 
**gateway** | **str** | Edge Gateway IP address | 
**id** | **str** | ID of the entity Maximum character length is 64 for project, category, and runtime environment, 36 for other entity types. | [optional] 
**ip_address** | **str** | Edge IP Address | 
**labels** | [**list[CategoryInfo]**](CategoryInfo.md) | A list of Category labels for this edge. | [optional] 
**name** | **str** | Edge name. Maximum length edge name is determined by kubernetes. Name length limited to 60 as node name is the edge name plus a suffix. https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/util/validation/validation.go | 
**role** | [**NodeRole**](NodeRole.md) |  | [optional] 
**serial_number** | **str** | Edge serial number | 
**short_id** | **str** | ShortID is the unique ID for the given edge. This ID must be unique for each edge, for the given tenant. | [optional] 
**storage_capacity** | **float** | Edge storage capacity in GB | 
**storage_usage** | **float** | Edge storage usage in GB | 
**subnet** | **str** | Edge subnet mask | 
**type** | **str** | Edge type. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

