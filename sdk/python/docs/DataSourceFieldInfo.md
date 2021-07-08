# DataSourceFieldInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**field_type** | **str** | Data type for the field. For example, Temperature, Pressure, Custom, and so on. Specify Custom for the entire topic payload. No special extraction is performed. When you specify Custom, Karbon Platform Services might not perform intelligent operations automatically when you specify other fields like Temperature. In the future custom extraction functions for each field might be allowed. DataSource dataType is derived from fieldType of all fields in the data source. | 
**mqtt_topic** | **str** | Topic for the field. The topic specified depends on the protocol in the data source. Specify the mqqtTopic for the MQTT protocol. For the RTSP protocol, the topic is the server endpoint or named protocol stream in the RSTP URL. | 
**name** | **str** | A unique name within the the data source. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

