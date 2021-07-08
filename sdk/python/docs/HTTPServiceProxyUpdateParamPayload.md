# HTTPServiceProxyUpdateParamPayload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**disable_rewrite_rules** | **bool** | By default, a rewrite rule will be put in place to rewrite service URL path base to / set this flag to true to retain the URL path base. | [optional] 
**duration** | **str** | Duration of the http service proxy. | [optional] 
**headers** | **str** | JSON object representation of HTTP headers to overwrite. May be useful for (https) endpoint that require specific Host field for example. | [optional] 
**name** | **str** |  | [optional] 
**setup_dns** | **bool** | Whether to setup DNS entry for this service. Default is false. Might be useful for services that do not work with URL path. However, bear in mind it may take several minutes for the DNS name to propagate/resolve. | 
**skip_cert_verification** | **bool** | Whether to skip TLS certification verification for endpoint. Only relevant when TLSEndpoint is true. This should be set to true if the endpoint is using a self-signed certificate. | 
**tls_endpoint** | **bool** | Whether the endpoint to proxy to is a TLS endpoint. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

