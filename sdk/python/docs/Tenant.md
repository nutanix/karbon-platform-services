# Tenant

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**description** | **str** | Tenant description. Up to 200 characters. | [optional] 
**external_id** | **str** | Unique tenant ID returned by my.nutanix.com. | 
**id** | **str** | Unique ID to identify the tenant, which. can be supplied during create or DB generated. For Nice we will have fixed tenant id such as tenant-id-waldot tenant-id-rocket-blue | 
**name** | **str** | Tenant name. For example, WalDot, Rocket Blue, and so on. Up to 200 characters. | 
**profile** | [**TenantProfile**](TenantProfile.md) |  | [optional] 
**token** | **str** | Unique token for a tenant. Used in authentication. | 
**version** | **float** | Version number of object maintained by DB. Not currently used. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

