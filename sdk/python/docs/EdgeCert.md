# EdgeCert

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ca_certificate** | **str** | Root CA certificate for the tenant. | 
**certificate** | **str** | Certificate for the edge using old/fixed root CA. | 
**client_certificate** | **str** | Certificate for mqtt client on the edge | 
**client_private_key** | **str** | Encrypted private key corresponding to the client certificate. | 
**edge_certificate** | **str** | Certificate for the edge using per-tenant root CA. | 
**edge_id** | **str** | ID of the edge this entity belongs to | 
**edge_private_key** | **str** | Encrypted private key using per-tenant root CA. | 
**id** | **str** | ID of the entity Maximum character length is 64 for project, category, and runtime environment, 36 for other entity types. | [optional] 
**locked** | **bool** |  | [optional] 
**private_key** | **str** | Encrypted private key using old/fixed root CA. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

