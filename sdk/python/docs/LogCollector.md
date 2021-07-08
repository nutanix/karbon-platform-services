# LogCollector

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cloud_creds_id** | **str** | CloudCreds id. Destination id for the cloud (should match with the CloudDestinationType) | 
**cloudwatch_details** | [**LogCollectorCloudwatch**](LogCollectorCloudwatch.md) |  | [optional] 
**code** | **str** | A code to modify logs during collection Log stream modifications (script source code) | [optional] 
**dest** | [**LogCollectorDestination**](LogCollectorDestination.md) |  | 
**id** | **str** | ID of the entity Maximum character length is 64 for project, category, and runtime environment, 36 for other entity types. | [optional] 
**kinesis_details** | [**LogCollectorKinesis**](LogCollectorKinesis.md) |  | [optional] 
**name** | **str** | Name of the LogCollector. Visible by UI only | 
**project_id** | **str** | ID of parent project. This should be required for PROJECT log collectors. | [optional] 
**sources** | [**LogCollectorSources**](LogCollectorSources.md) |  | 
**stackdriver_details** | [**LogCollectorStackdriver**](LogCollectorStackdriver.md) |  | [optional] 
**state** | [**LogCollectorStatus**](LogCollectorStatus.md) |  | 
**type** | [**LogCollectorType**](LogCollectorType.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

