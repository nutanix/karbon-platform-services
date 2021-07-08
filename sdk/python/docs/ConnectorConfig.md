# ConnectorConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**connector_instance_id** | **str** |  | 
**description** | **str** |  | [optional] 
**exclude_service_domain_ids** | **list[str]** | Service domains to be excluded from the connector config deployment. | [optional] 
**id** | **str** | ID of the entity Maximum character length is 64 for project, category, and runtime environment, 36 for other entity types. | [optional] 
**name** | **str** |  | 
**parameters** | [**ConnectorParametersValues**](ConnectorParametersValues.md) |  | [optional] 
**service_domain_ids** | **list[str]** | Service domains listed according to ID where the connector config is deployed. Only relevant if the parent project EdgeSelectorType value is set to Explicit. | [optional] 
**service_domain_selectors** | [**list[CategoryInfo]**](CategoryInfo.md) | Select service domains according to CategoryInfo. Only relevant if the parent project EdgeSelectorType value is set to Category. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

