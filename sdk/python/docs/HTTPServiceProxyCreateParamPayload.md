# HTTPServiceProxyCreateParamPayload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**disable_rewrite_rules** | **bool** | By default, a rewrite rule will be put in place to rewrite service URL path base to / set this flag to true to retain the URL path base. | 
**duration** | **str** | Duration of the http service proxy. | 
**headers** | **str** | JSON object representation of HTTP headers to overwrite. May be useful for (https) endpoint that require specific Host field for example. | [optional] 
**name** | **str** |  | 
**project_id** | **str** |  | [optional] 
**service_name** | **str** |  | 
**service_namespace** | **str** | Namespace of the http service, required when TYPE &#x3D; SYSTEM | [optional] 
**service_port** | **int** |  | 
**setup_basic_auth** | **bool** | Whether to setup basic auth to protect the endpoint | 
**setup_dns** | **bool** | Whether to setup DNS entry for this service. Default is false. Might be useful for services that do not work with URL path. However, bear in mind it may take several minutes for the DNS name to propagate/resolve. | 
**skip_cert_verification** | **bool** | Whether to skip TLS certification verification for endpoint. Only relevant when TLSEndpoint is true. This should be set to true if the endpoint is using a self-signed certificate. | 
**svc_domain_id** | **str** | ID of Service Domain to create the http service proxy | 
**tls_endpoint** | **bool** | Whether the endpoint to proxy to is a TLS endpoint. | 
**type** | **str** |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

