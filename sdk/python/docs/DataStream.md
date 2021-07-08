# DataStream

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**data_ifc_endpoints** | [**list[DataIfcEndpoint]**](DataIfcEndpoint.md) | Data Ifc endpoints connected to this datastream | [optional] 
**aws_cloud_region** | **str** | AWS region. Required if cloudType &#x3D;&#x3D; AWS | [optional] 
**aws_stream_type** | **str** | Type of the DataStream at AWS Cloud. Required if cloudType &#x3D;&#x3D; AWS | [optional] 
**az_stream_type** | **str** | Type of the DataStream at Azure Cloud. Required if cloudType &#x3D;&#x3D; Azure | [optional] 
**cloud_creds_id** | **str** | CloudCreds id. Required if destination &#x3D;&#x3D; Cloud | [optional] 
**cloud_type** | **str** | Cloud type, required if destination &#x3D;&#x3D; Cloud | [optional] 
**data_retention** | [**list[RetentionInfo]**](RetentionInfo.md) | Retention policy for this DataStream. Multiple RetentionInfo are combined using AND semantics. For example, retain data for 1 month AND up to 2 TB of data. | 
**data_type** | **str** | Data type of the DataStream. For example, Temperature, Pressure, Image, Multiple, etc. | 
**description** | **str** | The description of the DataStream | [optional] 
**destination** | **str** | Destination of the DataStream. Either Edge or Cloud or DataInterface. | 
**edge_stream_type** | **str** | Type of the DataStream at Edge. Required if destination &#x3D;&#x3D; Edge | [optional] 
**enable_sampling** | **bool** | Whether to turn sampling on. If true, then samplingInterval should be set as well. | 
**end_point** | **str** | End point of datastream. User specifies the endpoint. | [optional] 
**end_point_uri** | **str** | Endpoint URI Derived from existing fields required false | [optional] 
**gcp_cloud_region** | **str** | GCP region. Required if cloudType &#x3D;&#x3D; GCP | [optional] 
**gcp_stream_type** | **str** | Type of the DataStream at GCP Cloud. Required if cloudType &#x3D;&#x3D; GCP | [optional] 
**id** | **str** | ID of the entity Maximum character length is 64 for project, category, and runtime environment, 36 for other entity types. | [optional] 
**name** | **str** | Name of the DataStream. This is the published output (Kafka topic) name. | 
**origin** | **str** | The origin of the DataStream. Either &#x27;Data Source&#x27; or &#x27;Data Stream&#x27; | 
**origin_id** | **str** | If origin &#x3D;&#x3D; &#x27;Data Stream&#x27;, then originId can be used in place of originSelectors to specify the origin data stream ID if the origin data stream is unique. | [optional] 
**origin_selectors** | [**list[CategoryInfo]**](CategoryInfo.md) | A list of CategoryInfo used as criteria to filter sources applicable to this DataStream. | 
**out_data_ifc** | [**DataSource**](DataSource.md) |  | [optional] 
**project_id** | **str** | ID of parent project. This should be required, but is not marked as such due to backward compatibility. | [optional] 
**sampling_interval** | **float** | Sampling interval in seconds. The sampling interval applies to each mqtt/kafka topic separately. | [optional] 
**size** | **float** | Current size of the DataStream output in GB. | 
**state** | **str** | State of this entity | [optional] 
**transformation_args_list** | [**list[TransformationArgs]**](TransformationArgs.md) | List of transformations (together with their args) to apply to the origin data to produce the destination data. Could be empty if no transformation required. Each entry is the id of the transformation Script to apply to input from origin to produce output to destination. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

