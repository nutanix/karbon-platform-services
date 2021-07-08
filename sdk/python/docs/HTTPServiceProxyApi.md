# kps_api.HTTPServiceProxyApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**h_ttp_service_proxy_create**](HTTPServiceProxyApi.md#h_ttp_service_proxy_create) | **POST** /v1.0/httpserviceproxies | Create a HTTP service proxy.
[**h_ttp_service_proxy_delete**](HTTPServiceProxyApi.md#h_ttp_service_proxy_delete) | **DELETE** /v1.0/httpserviceproxies/{id} | Delete HTTP service proxy.
[**h_ttp_service_proxy_get**](HTTPServiceProxyApi.md#h_ttp_service_proxy_get) | **GET** /v1.0/httpserviceproxies/{id} | Get a HTTP service proxy by its ID.
[**h_ttp_service_proxy_list**](HTTPServiceProxyApi.md#h_ttp_service_proxy_list) | **GET** /v1.0/httpserviceproxies | Get all HTTP service proxies.
[**h_ttp_service_proxy_update**](HTTPServiceProxyApi.md#h_ttp_service_proxy_update) | **PUT** /v1.0/httpserviceproxies/{id} | Update a HTTP service proxy by its ID.

# **h_ttp_service_proxy_create**
> HTTPServiceProxyCreateResponsePayload h_ttp_service_proxy_create(body, authorization)

Create a HTTP service proxy.

Create a HTTP service proxy.

### Example
```python
from __future__ import print_function
import time
import kps_api
from kps_api.rest import ApiException
from pprint import pprint

# Configure API key authorization: BearerToken
configuration = kps_api.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = kps_api.HTTPServiceProxyApi(kps_api.ApiClient(configuration))
body = kps_api.HTTPServiceProxyCreateParamPayload() # HTTPServiceProxyCreateParamPayload | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Create a HTTP service proxy.
    api_response = api_instance.h_ttp_service_proxy_create(body, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling HTTPServiceProxyApi->h_ttp_service_proxy_create: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HTTPServiceProxyCreateParamPayload**](HTTPServiceProxyCreateParamPayload.md)|  | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**HTTPServiceProxyCreateResponsePayload**](HTTPServiceProxyCreateResponsePayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **h_ttp_service_proxy_delete**
> DeleteDocumentResponseV2 h_ttp_service_proxy_delete(id, authorization)

Delete HTTP service proxy.

Delete the HTTP service proxy with the given ID {id}.

### Example
```python
from __future__ import print_function
import time
import kps_api
from kps_api.rest import ApiException
from pprint import pprint

# Configure API key authorization: BearerToken
configuration = kps_api.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = kps_api.HTTPServiceProxyApi(kps_api.ApiClient(configuration))
id = 'id_example' # str | ID of the entity
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Delete HTTP service proxy.
    api_response = api_instance.h_ttp_service_proxy_delete(id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling HTTPServiceProxyApi->h_ttp_service_proxy_delete: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| ID of the entity | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**DeleteDocumentResponseV2**](DeleteDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **h_ttp_service_proxy_get**
> HTTPServiceProxy h_ttp_service_proxy_get(id, authorization)

Get a HTTP service proxy by its ID.

Retrieves a HTTP service proxy with the given ID {id}.

### Example
```python
from __future__ import print_function
import time
import kps_api
from kps_api.rest import ApiException
from pprint import pprint

# Configure API key authorization: BearerToken
configuration = kps_api.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = kps_api.HTTPServiceProxyApi(kps_api.ApiClient(configuration))
id = 'id_example' # str | ID of the entity
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Get a HTTP service proxy by its ID.
    api_response = api_instance.h_ttp_service_proxy_get(id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling HTTPServiceProxyApi->h_ttp_service_proxy_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| ID of the entity | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**HTTPServiceProxy**](HTTPServiceProxy.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **h_ttp_service_proxy_list**
> HTTPServiceProxyListPayload h_ttp_service_proxy_list(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter, type=type, project_id=project_id, svc_domain_id=svc_domain_id, name=name, service_name=service_name, service_namespace=service_namespace)

Get all HTTP service proxies.

Retrieves a list of all HTTP service proxies.

### Example
```python
from __future__ import print_function
import time
import kps_api
from kps_api.rest import ApiException
from pprint import pprint

# Configure API key authorization: BearerToken
configuration = kps_api.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = kps_api.HTTPServiceProxyApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)
type = 'type_example' # str | Type of the HTTP Service Proxy (optional)
project_id = 'project_id_example' # str | HTTP Service Proxy Project ID (optional)
svc_domain_id = 'svc_domain_id_example' # str | HTTP Service Proxy Service Domain ID (optional)
name = 'name_example' # str | Name of the HTTP Service Proxy (optional)
service_name = 'service_name_example' # str | ServiceName of the HTTP Service Proxy (optional)
service_namespace = 'service_namespace_example' # str | ServiceNamespace of the HTTP Service Proxy (optional)

try:
    # Get all HTTP service proxies.
    api_response = api_instance.h_ttp_service_proxy_list(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter, type=type, project_id=project_id, svc_domain_id=svc_domain_id, name=name, service_name=service_name, service_namespace=service_namespace)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling HTTPServiceProxyApi->h_ttp_service_proxy_list: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **page_index** | **int**| 0-based index of the page to fetch results. | [optional] 
 **page_size** | **int**| Item count of each page. | [optional] 
 **order_by** | [**list[str]**](str.md)| Specify result order. Zero or more entries with format: &amp;ltkey&gt; [desc] where orderByKeys lists allowed keys in each response. | [optional] 
 **filter** | **str**| Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE &#x27;foo%&#x27;. Supported filter keys are the same as order by keys. | [optional] 
 **type** | **str**| Type of the HTTP Service Proxy | [optional] 
 **project_id** | **str**| HTTP Service Proxy Project ID | [optional] 
 **svc_domain_id** | **str**| HTTP Service Proxy Service Domain ID | [optional] 
 **name** | **str**| Name of the HTTP Service Proxy | [optional] 
 **service_name** | **str**| ServiceName of the HTTP Service Proxy | [optional] 
 **service_namespace** | **str**| ServiceNamespace of the HTTP Service Proxy | [optional] 

### Return type

[**HTTPServiceProxyListPayload**](HTTPServiceProxyListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **h_ttp_service_proxy_update**
> HTTPServiceProxyUpdateResponsePayload h_ttp_service_proxy_update(body, authorization, id)

Update a HTTP service proxy by its ID.

Update a HTTP service proxy with the given ID {id}.

### Example
```python
from __future__ import print_function
import time
import kps_api
from kps_api.rest import ApiException
from pprint import pprint

# Configure API key authorization: BearerToken
configuration = kps_api.Configuration()
configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = kps_api.HTTPServiceProxyApi(kps_api.ApiClient(configuration))
body = kps_api.HTTPServiceProxyUpdateParamPayload() # HTTPServiceProxyUpdateParamPayload | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
id = 'id_example' # str | ID of the entity

try:
    # Update a HTTP service proxy by its ID.
    api_response = api_instance.h_ttp_service_proxy_update(body, authorization, id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling HTTPServiceProxyApi->h_ttp_service_proxy_update: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HTTPServiceProxyUpdateParamPayload**](HTTPServiceProxyUpdateParamPayload.md)|  | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **id** | **str**| ID of the entity | 

### Return type

[**HTTPServiceProxyUpdateResponsePayload**](HTTPServiceProxyUpdateResponsePayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

