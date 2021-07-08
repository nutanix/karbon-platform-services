# EdgeCluster

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**description** | **str** | EdgeCluster description | [optional] 
**id** | **str** | ID of the entity Maximum character length is 64 for project, category, and runtime environment, 36 for other entity types. | [optional] 
**labels** | [**list[CategoryInfo]**](CategoryInfo.md) | List of edge Device IDs in cluster | 
**name** | **str** | EdgeCluster name. Maximum length edge name is determined by kubernetes. Name length limited to 60 as node name is the edge cluster name plus a suffix. https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/util/validation/validation.go | 
**virtual_ip** | **str** | Virtual IP | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

