# ApplicationV2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**app_manifest** | **str** | The kubernetes manifest for the application in YAML format. | 
**data_ifc_endpoints** | [**list[DataIfcEndpoint]**](DataIfcEndpoint.md) | DataIfcEndpoints is a list of endpoints exposed to an application. | [optional] 
**description** | **str** | A description of the application. Maximum length of 200 characters. | [optional] 
**edge_ids** | **list[str]** | Edges listed according to ID where the application is deployed. Only relevant if the parent project EdgeSelectorType value is set to Explicit. | [optional] 
**edge_selectors** | [**list[CategoryInfo]**](CategoryInfo.md) | Select edges according to CategoryInfo. Only relevant if the parent project EdgeSelectorType value is set to Category. | [optional] 
**exclude_edge_ids** | **list[str]** | Edges to be excluded from the application deployment. | [optional] 
**helm_metadata** | [**HelmAppMetadata**](HelmAppMetadata.md) |  | [optional] 
**id** | **str** | ID of the entity Maximum character length is 64 for project, category, and runtime environment, 36 for other entity types. | [optional] 
**name** | **str** | The application name. Maximum length of 200 characters. | 
**only_pre_pull_on_update** | **bool** | Only pre-pull images on service domains w/o doing an actual update. Service domain which have not yet deployed the app will deploy application like usual. Update will commence once this flag is unset. | [optional] 
**origin_selectors** | [**list[CategoryInfo]**](CategoryInfo.md) | OriginSelectors is the list of CategoryInfo used as criteria to feed data into applications. | [optional] 
**packaging_type** | **str** | PackagingType vanilla or helm, nil &#x3D; vanilla | [optional] 
**project_id** | **str** | Parent project ID. Not required (to maintain backward compatibility). | 
**state** | **str** | State of this entity | [optional] 
**user_id** | **str** | Last modified user ID. Only required by edge for privilege check. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

