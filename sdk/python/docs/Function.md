# Function

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**code** | **str** | The source code for the function script. | 
**description** | **str** | Provide a description for your function code/script. | [optional] 
**environment** | **str** | Runtime environment for the function code/script. | 
**id** | **str** | ID of the entity Maximum character length is 64 for project, category, and runtime environment, 36 for other entity types. | [optional] 
**language** | **str** | Programming language for the function code/script. Supported languages are python and javascript | 
**name** | **str** | Function name. | 
**params** | [**list[ScriptParam]**](ScriptParam.md) | Array of script parameters. | 
**project_id** | **str** | ID of parent project, required for custom (non-builtin) scripts. | [optional] 
**runtime_id** | **str** | ID of the ScriptRuntime to use to run this script | [optional] 
**runtime_tag** | **str** | Docker image tag of the ScriptRuntime to use to run this script. If missing or empty, then backend should treat it as \&quot;latest\&quot; | [optional] 
**type** | **str** | Type of function code/script: Transformation or Function. Transformation takes a data stream as input and produces a different data stream as output. Function takes a data stream as input but has no constraint on output. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

