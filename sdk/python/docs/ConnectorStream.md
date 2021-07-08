# ConnectorStream

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**connector_instance_id** | **str** |  | 
**description** | **str** |  | [optional] 
**direction** | [**ConnectorStreamDirection**](ConnectorStreamDirection.md) |  | 
**exclude_service_domain_ids** | **list[str]** | Service domains to be excluded from the connector config deployment. | [optional] 
**id** | **str** | ID of the entity Maximum character length is 64 for project, category, and runtime environment, 36 for other entity types. | [optional] 
**labels** | [**list[CategoryInfo]**](CategoryInfo.md) | A list of Category labels for this connector config. \&quot;SOURCE\&quot; connector streams should have at least 1 label. \&quot;SINK\&quot; should not have any labels. | [optional] 
**name** | **str** |  | 
**service_domain_ids** | **list[str]** | Service domains listed according to ID where the connector config is deployed. Only relevant if the parent project EdgeSelectorType value is set to Explicit. | [optional] 
**service_domain_selectors** | [**list[CategoryInfo]**](CategoryInfo.md) | Select service domains according to CategoryInfo. Only relevant if the parent project EdgeSelectorType value is set to Category. | [optional] 
**stream** | [**ConnectorParametersValues**](ConnectorParametersValues.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

