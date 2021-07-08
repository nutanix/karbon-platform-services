# ServiceDomain

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**description** | **str** | EdgeCluster description | [optional] 
**env** | **str** |  | [optional] 
**id** | **str** | ID of the entity Maximum character length is 64 for project, category, and runtime environment, 36 for other entity types. | [optional] 
**labels** | [**list[CategoryInfo]**](CategoryInfo.md) | A list of Category labels for this service domain. | [optional] 
**name** | **str** | Service domain name. Maximum length is limited to 60 characters which must satisfy https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/util/validation/validation.go | 
**profile** | [**ServiceDomainProfile**](ServiceDomainProfile.md) |  | [optional] 
**virtual_ip** | **str** | Virtual IP | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

