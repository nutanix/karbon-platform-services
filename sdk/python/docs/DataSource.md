# DataSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**auth_type** | **str** | Authentication type used by the sensor. | 
**connection** | **str** | Sensor connection type: Secure or Unsecure | 
**edge_id** | **str** | ID of the edge this entity belongs to | 
**fields** | [**list[DataSourceFieldInfo]**](DataSourceFieldInfo.md) | User defined fields to extract data from the topic payload. | 
**id** | **str** | ID of the entity Maximum character length is 64 for project, category, and runtime environment, 36 for other entity types. | [optional] 
**ifc_info** | [**DataSourceIfcInfo**](DataSourceIfcInfo.md) |  | [optional] 
**name** | **str** |  | 
**protocol** | **str** | Protocol used by the Sensor. | 
**selectors** | [**list[DataSourceFieldSelector]**](DataSourceFieldSelector.md) | A list of DataSourceFieldSelector users assigned to the data source. Allows a user to use Category selectors to identify the data pipeline source. Selectors with different category IDs are combined with the AND operator, while selectors with the same category ID are combined with the OR operator. | 
**type** | **str** | Data source type: Sensor or Gateway | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

