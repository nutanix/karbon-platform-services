# AuditLogV2Filter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**end_time** | **datetime** | Search for events by this earlier timestamp (inclusive). | [optional] 
**from_document** | **int** |  | [optional] 
**group_by** | **str** |  | [optional] 
**page_size** | **int** |  | [optional] 
**scopes** | **list[str]** |  | [optional] 
**start_time** | **datetime** | Search for events by this later timestamp (inclusive) | [optional] 
**terms_key_value** | [**dict(str, AuditLogV2MultipleValues)**](AuditLogV2MultipleValues.md) | TenantID must be provided in order to search audit logs | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

