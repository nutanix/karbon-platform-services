# kps_api.ServiceInstanceApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**service_instance_create**](ServiceInstanceApi.md#service_instance_create) | **POST** /v1.0/serviceinstances | Create a Service Instance.
[**service_instance_delete**](ServiceInstanceApi.md#service_instance_delete) | **DELETE** /v1.0/serviceinstances/{svcInstanceId} | Delete a Service Instance.
[**service_instance_get**](ServiceInstanceApi.md#service_instance_get) | **GET** /v1.0/serviceinstances/{svcInstanceId} | Get a Service Instance.
[**service_instance_list**](ServiceInstanceApi.md#service_instance_list) | **GET** /v1.0/serviceinstances | List Service Instances.
[**service_instance_status_list**](ServiceInstanceApi.md#service_instance_status_list) | **GET** /v1.0/serviceinstances/{svcInstanceId}/status | Get the status of the Service Instance.
[**service_instance_update**](ServiceInstanceApi.md#service_instance_update) | **PUT** /v1.0/serviceinstances/{svcInstanceId} | Update a Service Instance.

# **service_instance_create**
> CreateDocumentResponseV2 service_instance_create(body, authorization)

Create a Service Instance.

Create a Service Instance

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
api_instance = kps_api.ServiceInstanceApi(kps_api.ApiClient(configuration))
body = kps_api.ServiceInstanceParam() # ServiceInstanceParam | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Create a Service Instance.
    api_response = api_instance.service_instance_create(body, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceInstanceApi->service_instance_create: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ServiceInstanceParam**](ServiceInstanceParam.md)|  | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**CreateDocumentResponseV2**](CreateDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_instance_delete**
> DeleteDocumentResponseV2 service_instance_delete(svc_instance_id, authorization)

Delete a Service Instance.

Delete a Service Instance

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
api_instance = kps_api.ServiceInstanceApi(kps_api.ApiClient(configuration))
svc_instance_id = 'svc_instance_id_example' # str | Service Instance ID
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Delete a Service Instance.
    api_response = api_instance.service_instance_delete(svc_instance_id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceInstanceApi->service_instance_delete: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svc_instance_id** | **str**| Service Instance ID | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**DeleteDocumentResponseV2**](DeleteDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_instance_get**
> ServiceInstance service_instance_get(svc_instance_id, authorization)

Get a Service Instance.

Get a Service Instance

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
api_instance = kps_api.ServiceInstanceApi(kps_api.ApiClient(configuration))
svc_instance_id = 'svc_instance_id_example' # str | Service Instance ID
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Get a Service Instance.
    api_response = api_instance.service_instance_get(svc_instance_id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceInstanceApi->service_instance_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svc_instance_id** | **str**| Service Instance ID | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**ServiceInstance**](ServiceInstance.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_instance_list**
> ServiceInstanceListPayload service_instance_list(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter, scope=scope, type=type, svc_version=svc_version, svc_class_id=svc_class_id, scope_id=scope_id)

List Service Instances.

List Service Instances

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
api_instance = kps_api.ServiceInstanceApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)
scope = 'scope_example' # str | Scope of the Service Class (optional)
type = 'type_example' # str | Type of the Service Class (optional)
svc_version = 'svc_version_example' # str | Version of the Service Class (optional)
svc_class_id = 'svc_class_id_example' # str | Service Class ID (optional)
scope_id = 'scope_id_example' # str | Service Class scope ID (optional)

try:
    # List Service Instances.
    api_response = api_instance.service_instance_list(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter, scope=scope, type=type, svc_version=svc_version, svc_class_id=svc_class_id, scope_id=scope_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceInstanceApi->service_instance_list: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **page_index** | **int**| 0-based index of the page to fetch results. | [optional] 
 **page_size** | **int**| Item count of each page. | [optional] 
 **order_by** | [**list[str]**](str.md)| Specify result order. Zero or more entries with format: &amp;ltkey&gt; [desc] where orderByKeys lists allowed keys in each response. | [optional] 
 **filter** | **str**| Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE &#x27;foo%&#x27;. Supported filter keys are the same as order by keys. | [optional] 
 **scope** | **str**| Scope of the Service Class | [optional] 
 **type** | **str**| Type of the Service Class | [optional] 
 **svc_version** | **str**| Version of the Service Class | [optional] 
 **svc_class_id** | **str**| Service Class ID | [optional] 
 **scope_id** | **str**| Service Class scope ID | [optional] 

### Return type

[**ServiceInstanceListPayload**](ServiceInstanceListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_instance_status_list**
> ServiceInstanceStatusListPayload service_instance_status_list(svc_instance_id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter, svc_domain_id=svc_domain_id)

Get the status of the Service Instance.

Get the status of the Service Instance on Service Domains

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
api_instance = kps_api.ServiceInstanceApi(kps_api.ApiClient(configuration))
svc_instance_id = 'svc_instance_id_example' # str | Service Instance ID
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)
svc_domain_id = 'svc_domain_id_example' # str |  (optional)

try:
    # Get the status of the Service Instance.
    api_response = api_instance.service_instance_status_list(svc_instance_id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter, svc_domain_id=svc_domain_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceInstanceApi->service_instance_status_list: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svc_instance_id** | **str**| Service Instance ID | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **page_index** | **int**| 0-based index of the page to fetch results. | [optional] 
 **page_size** | **int**| Item count of each page. | [optional] 
 **order_by** | [**list[str]**](str.md)| Specify result order. Zero or more entries with format: &amp;ltkey&gt; [desc] where orderByKeys lists allowed keys in each response. | [optional] 
 **filter** | **str**| Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE &#x27;foo%&#x27;. Supported filter keys are the same as order by keys. | [optional] 
 **svc_domain_id** | **str**|  | [optional] 

### Return type

[**ServiceInstanceStatusListPayload**](ServiceInstanceStatusListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_instance_update**
> UpdateDocumentResponseV2 service_instance_update(body, authorization, svc_instance_id)

Update a Service Instance.

Update a Service Instance

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
api_instance = kps_api.ServiceInstanceApi(kps_api.ApiClient(configuration))
body = kps_api.ServiceInstanceParam() # ServiceInstanceParam | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
svc_instance_id = 'svc_instance_id_example' # str | Service Instance ID

try:
    # Update a Service Instance.
    api_response = api_instance.service_instance_update(body, authorization, svc_instance_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceInstanceApi->service_instance_update: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ServiceInstanceParam**](ServiceInstanceParam.md)|  | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **svc_instance_id** | **str**| Service Instance ID | 

### Return type

[**UpdateDocumentResponseV2**](UpdateDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

