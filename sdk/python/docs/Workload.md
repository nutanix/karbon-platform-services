# Workload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**additional_detail_sample** | [**AdditionalItem**](AdditionalItem.md) |  | [optional] 
**additional_details** | [**list[AdditionalItem]**](AdditionalItem.md) | Additional details to display, such as configured annotations | [optional] 
**app_label** | **bool** | Define if Pods related to this Workload has the label App | 
**available_replicas** | **int** | Number of available replicas | 
**created_at** | **str** | Creation timestamp (in RFC3339 format) | 
**current_replicas** | **int** | Number of current replicas pods that matches controller selector labels | 
**desired_replicas** | **int** | Number of desired replicas defined by the user in the controller Spec | 
**istio_sidecar** | **bool** | Define if Pods related to this Workload has an IstioSidecar deployed | 
**labels** | **dict(str, str)** | Workload labels | [optional] 
**name** | **str** | Name of the workload | 
**pod_count** | **int** | Number of current workload pods | 
**pods** | [**Pods**](Pods.md) |  | [optional] 
**resource_version** | **str** | Kubernetes ResourceVersion | 
**runtimes** | [**list[Runtime]**](Runtime.md) | Runtimes and associated dashboards | [optional] 
**services** | [**Services**](Services.md) |  | [optional] 
**type** | **str** | Type of the workload | 
**version_label** | **bool** | Define if Pods related to this Workload has the label Version | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

