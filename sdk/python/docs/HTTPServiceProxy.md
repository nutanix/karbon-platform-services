# HTTPServiceProxy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**dns_url** | **str** | DNS URL of the service proxy endpoint Valid only if setupDNS is set to true when creating the service proxy | 
**duration** | **str** | Duration of the http service proxy. | 
**id** | **str** | ID of the entity Maximum character length is 64 for project, category, and runtime environment, 36 for other entity types. | [optional] 
**name** | **str** | HTTP service proxy name. Unique within (tenant, service domain) | 
**password** | **str** | Password to login to the service when setupBasicAuth&#x3D;true. | [optional] 
**project_id** | **str** | ID of parent project, required when TYPE &#x3D; PROJECT. | [optional] 
**service_name** | **str** | Name of the http service. | 
**service_namespace** | **str** | Namespace of the http service, required when TYPE &#x3D; SYSTEM | [optional] 
**service_port** | **int** | Port of the http service. | 
**svc_domain_id** | **str** | ID of the service domain this entity belongs to | 
**type** | **str** | Service type for this http proxy. | 
**url** | **str** | URL of the service proxy endpoint | 
**username** | **str** | Username to login to the service when setupBasicAuth&#x3D;true. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

