# NodeInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**artifacts** | **dict(str, object)** | Artifacts is a json object for passing node ip and service ports | [optional] 
**connected** | **bool** |  | [optional] 
**cpu_usage** | **str** | Node CPU usage. | [optional] 
**gpu_info** | **str** | Information about GPUs associated with the node. | [optional] 
**gpu_usage** | **str** | Node GPU Usage. | [optional] 
**health_bits** | **dict(str, bool)** |  | [optional] 
**health_status** | [**NodeHealthStatus**](NodeHealthStatus.md) |  | [optional] 
**healthy** | **bool** | Deprecated. Use healthStatus instead | [optional] 
**id** | **str** | ID of the entity Maximum character length is 64 for project, category, and runtime environment, 36 for other entity types. | [optional] 
**kube_version** | **str** | Node Kubernetes version. | [optional] 
**memory_free_kb** | **str** | Free (available) node memory in KB. | [optional] 
**node_build_num** | **str** | Node build number. | [optional] 
**node_id** | **str** |  | [optional] 
**node_version** | **str** | Node version. | [optional] 
**num_cpu** | **str** | Number of CPUs assigned to the node. | [optional] 
**onboarded** | **bool** |  | [optional] 
**os_version** | **str** | Node OS version | [optional] 
**storage_free_kb** | **str** | Free (available) node storage in KB. | [optional] 
**svc_domain_id** | **str** | ID of the service domain this entity belongs to | 
**total_memory_kb** | **str** | Total node memory in KB. | [optional] 
**total_storage_kb** | **str** | Total node storage capacity in KB. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

