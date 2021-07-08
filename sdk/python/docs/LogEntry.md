# LogEntry

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**batch_id** | **str** | ID that identifies logs from different edge as the same batch. | [optional] 
**edge_id** | **str** | ID of the edge this entity belongs to | 
**error_message** | **str** | Error message - optional, should be populated when status &#x3D;&#x3D; &#x27;FAILED&#x27; | [optional] 
**id** | **str** | ID of the entity Maximum character length is 64 for project, category, and runtime environment, 36 for other entity types. | [optional] 
**location** | **str** | Location or object key for the log in the bucket. | [optional] 
**status** | [**LogUploadStatus**](LogUploadStatus.md) |  | [optional] 
**tags** | [**list[LogTag]**](LogTag.md) | Tags carry the properties of the log | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

