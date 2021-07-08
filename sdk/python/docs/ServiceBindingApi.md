# kps_api.ServiceBindingApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**service_binding_create**](ServiceBindingApi.md#service_binding_create) | **POST** /v1.0/servicebindings | Create a Service Binding.
[**service_binding_delete**](ServiceBindingApi.md#service_binding_delete) | **DELETE** /v1.0/servicebindings/{svcBindingId} | Delete a Service Binding.
[**service_binding_get**](ServiceBindingApi.md#service_binding_get) | **GET** /v1.0/servicebindings/{svcBindingId} | Get a Service Binding.
[**service_binding_list**](ServiceBindingApi.md#service_binding_list) | **GET** /v1.0/servicebindings | List Service Bindings.
[**service_binding_status_list**](ServiceBindingApi.md#service_binding_status_list) | **GET** /v1.0/servicebindings/{svcBindingId}/status | Get the status of Service Binding.

# **service_binding_create**
> CreateDocumentResponseV2 service_binding_create(body, authorization)

Create a Service Binding.

Create a Service Binding

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
api_instance = kps_api.ServiceBindingApi(kps_api.ApiClient(configuration))
body = kps_api.ServiceBindingParam() # ServiceBindingParam | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Create a Service Binding.
    api_response = api_instance.service_binding_create(body, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceBindingApi->service_binding_create: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ServiceBindingParam**](ServiceBindingParam.md)|  | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**CreateDocumentResponseV2**](CreateDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_binding_delete**
> DeleteDocumentResponseV2 service_binding_delete(svc_binding_id, authorization)

Delete a Service Binding.

Delete a Service Binding

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
api_instance = kps_api.ServiceBindingApi(kps_api.ApiClient(configuration))
svc_binding_id = 'svc_binding_id_example' # str | Service Binding ID
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Delete a Service Binding.
    api_response = api_instance.service_binding_delete(svc_binding_id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceBindingApi->service_binding_delete: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svc_binding_id** | **str**| Service Binding ID | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**DeleteDocumentResponseV2**](DeleteDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_binding_get**
> ServiceBinding service_binding_get(svc_binding_id, authorization)

Get a Service Binding.

Get a Service Binding

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
api_instance = kps_api.ServiceBindingApi(kps_api.ApiClient(configuration))
svc_binding_id = 'svc_binding_id_example' # str | Service Binding ID
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Get a Service Binding.
    api_response = api_instance.service_binding_get(svc_binding_id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceBindingApi->service_binding_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svc_binding_id** | **str**| Service Binding ID | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**ServiceBinding**](ServiceBinding.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_binding_list**
> ServiceBindingListPayload service_binding_list(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter, svc_class_id=svc_class_id, bind_resource_type=bind_resource_type, bind_resource_id=bind_resource_id)

List Service Bindings.

List Service Bindings

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
api_instance = kps_api.ServiceBindingApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)
svc_class_id = 'svc_class_id_example' # str | Service Class ID (optional)
bind_resource_type = 'bind_resource_type_example' # str | Bind resource type (optional)
bind_resource_id = 'bind_resource_id_example' # str | Bind resource ID (optional)

try:
    # List Service Bindings.
    api_response = api_instance.service_binding_list(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter, svc_class_id=svc_class_id, bind_resource_type=bind_resource_type, bind_resource_id=bind_resource_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceBindingApi->service_binding_list: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **page_index** | **int**| 0-based index of the page to fetch results. | [optional] 
 **page_size** | **int**| Item count of each page. | [optional] 
 **order_by** | [**list[str]**](str.md)| Specify result order. Zero or more entries with format: &amp;ltkey&gt; [desc] where orderByKeys lists allowed keys in each response. | [optional] 
 **filter** | **str**| Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE &#x27;foo%&#x27;. Supported filter keys are the same as order by keys. | [optional] 
 **svc_class_id** | **str**| Service Class ID | [optional] 
 **bind_resource_type** | **str**| Bind resource type | [optional] 
 **bind_resource_id** | **str**| Bind resource ID | [optional] 

### Return type

[**ServiceBindingListPayload**](ServiceBindingListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_binding_status_list**
> ServiceBindingStatusListPayload service_binding_status_list(svc_binding_id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter, svc_domain_id=svc_domain_id, svc_instance_id=svc_instance_id)

Get the status of Service Binding.

Get the status of Service Binding on Service Domains

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
api_instance = kps_api.ServiceBindingApi(kps_api.ApiClient(configuration))
svc_binding_id = 'svc_binding_id_example' # str | Service Binding ID
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)
svc_domain_id = 'svc_domain_id_example' # str | Service Domain ID (optional)
svc_instance_id = 'svc_instance_id_example' # str | Service Instance ID (optional)

try:
    # Get the status of Service Binding.
    api_response = api_instance.service_binding_status_list(svc_binding_id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter, svc_domain_id=svc_domain_id, svc_instance_id=svc_instance_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceBindingApi->service_binding_status_list: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svc_binding_id** | **str**| Service Binding ID | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **page_index** | **int**| 0-based index of the page to fetch results. | [optional] 
 **page_size** | **int**| Item count of each page. | [optional] 
 **order_by** | [**list[str]**](str.md)| Specify result order. Zero or more entries with format: &amp;ltkey&gt; [desc] where orderByKeys lists allowed keys in each response. | [optional] 
 **filter** | **str**| Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE &#x27;foo%&#x27;. Supported filter keys are the same as order by keys. | [optional] 
 **svc_domain_id** | **str**| Service Domain ID | [optional] 
 **svc_instance_id** | **str**| Service Instance ID | [optional] 

### Return type

[**ServiceBindingStatusListPayload**](ServiceBindingStatusListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

