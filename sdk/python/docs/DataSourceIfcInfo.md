# DataSourceIfcInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ifc_class** | **str** | Class of the data source DataInterface or Legacy | 
**ifc_driver_id** | **str** | Driver from which this data source is derived. | 
**ifc_img** | **str** | The docker img that includes the data source | 
**ifc_kind** | **str** | Kind of data source IN, OUT, PIPE (bidirectional) | 
**ifc_ports** | [**list[DataSourceIfcPorts]**](DataSourceIfcPorts.md) | Any ports that will be opened and used by this datasource | [optional] 
**ifc_project_id** | **str** | The project that contains this data source | [optional] 
**ifc_protocol** | **str** | Primary protocol that this data source implements | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

