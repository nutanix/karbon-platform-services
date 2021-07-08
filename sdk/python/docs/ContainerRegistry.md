# ContainerRegistry

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cloud_creds_id** | **str** | Existing cloud profile to use with the container registry profile.  Required if Type &#x3D;&#x3D; AWS || Type &#x3D;&#x3D; GCP | [optional] 
**description** | **str** | Description for the container registry profile. | [optional] 
**email** | **str** | Email address to associate with the container registry profile.  Required for container registry profiles. | [optional] 
**id** | **str** | ID of the entity Maximum character length is 64 for project, category, and runtime environment, 36 for other entity types. | [optional] 
**name** | **str** | Name for the container registry profile. | 
**pwd** | **str** | Password for the container registry profile.  Required for container registry profiles. | [optional] 
**server** | **str** | Provide a server URL to the container registry in the format used by your cloud provider. For example, an Amazon AWS Elastic Container Registry (ECR) URL might be: https://aws_account_id.dkr.ecr.region.amazonaws.com | 
**type** | **str** | Container registry profile type. | 
**user_name** | **str** | Cloud profile user name for use with the container registry profile.  Required for container registry profiles. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

