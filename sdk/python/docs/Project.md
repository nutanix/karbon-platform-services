# Project

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cloud_credential_ids** | **list[str]** | List of cloud profile credential IDs that the project can access. | 
**description** | **str** | Describe the project. | 
**docker_profile_ids** | **list[str]** | List of Docker container registry profile IDs that the project can access. | 
**edge_ids** | **list[str]** | List of edge IDs for edges in this project. Only relevant when edgeSelectorType &#x3D;&#x3D;&#x3D; &#x27;Explicit&#x27; | [optional] 
**edge_selector_type** | **str** | Type of edge selector: Category or Explicit. Specify whether edges belonging to this project are given by edgeIDs (Explicit) or edgeSelectors (Category). | 
**edge_selectors** | [**list[CategoryInfo]**](CategoryInfo.md) | Edge selectors - CategoryInfo list. Only relevant when edgeSelectorType &#x3D;&#x3D;&#x3D; &#x27;Category&#x27; | [optional] 
**id** | **str** | ID of the entity Maximum character length is 64 for project, category, and runtime environment, 36 for other entity types. | [optional] 
**name** | **str** | Project name. | 
**privileged** | **bool** | Privileged projects can use all Kubernetes resources | [optional] 
**users** | [**list[ProjectUserInfo]**](ProjectUserInfo.md) | List of users who can access the project. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

