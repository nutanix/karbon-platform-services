# ConnectorClass

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**config_parameter_schema** | [**ConnectorParametersSchema**](ConnectorParametersSchema.md) |  | [optional] 
**connector_version** | **str** | External version of a connector. It is possible to have multiple connectors with the same name, but different versions. | 
**description** | **str** |  | [optional] 
**id** | **str** | ID of the entity Maximum character length is 64 for project, category, and runtime environment, 36 for other entity types. | [optional] 
**min_svc_domain_version** | **str** |  | [optional] 
**name** | **str** |  | 
**static_parameter_schema** | [**ConnectorParametersSchema**](ConnectorParametersSchema.md) |  | [optional] 
**stream_parameter_schema** | [**ConnectorParametersSchema**](ConnectorParametersSchema.md) |  | [optional] 
**type** | [**ConnectorClassType**](ConnectorClassType.md) |  | 
**yaml_data** | **str** | The YAML content for the application. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

