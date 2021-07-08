# WorkloadListItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**additional_detail_sample** | [**AdditionalItem**](AdditionalItem.md) |  | [optional] 
**app_label** | **bool** | Define if Pods related to this Workload has the label App | 
**created_at** | **str** | Creation timestamp (in RFC3339 format) | 
**istio_sidecar** | **bool** | Define if Pods related to this Workload has an IstioSidecar deployed | 
**labels** | **dict(str, str)** | Workload labels | [optional] 
**name** | **str** | Name of the workload | 
**pod_count** | **int** | Number of current workload pods | 
**resource_version** | **str** | Kubernetes ResourceVersion | 
**type** | **str** | Type of the workload | 
**version_label** | **bool** | Define if Pods related to this Workload has the label Version | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

