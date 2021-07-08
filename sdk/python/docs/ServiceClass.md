# ServiceClass

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**bindable** | **bool** | Flag to specify if service binding is supported | 
**created_at** | **datetime** |  | [optional] 
**description** | **str** |  | [optional] 
**id** | **str** |  | [optional] 
**min_svc_domain_version** | **str** | Minimum version of the Service Domain supporting this Service Class | 
**name** | **str** |  | 
**schemas** | [**ServiceClassSchemas**](ServiceClassSchemas.md) |  | [optional] 
**scope** | [**ServiceClassScopeType**](ServiceClassScopeType.md) |  | 
**state** | [**ServiceClassStateType**](ServiceClassStateType.md) |  | 
**svc_version** | **str** | Version of the Service Class type | 
**tags** | [**list[ServiceClassTag]**](ServiceClassTag.md) | Tag name can be repeated to hold multiple values. Tags essential &#x3D; yes/no and category &#x3D; some category are required | [optional] 
**type** | **str** | Type of the Service Class e.g Kafka | 
**updated_at** | **datetime** |  | [optional] 
**version** | **float** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

