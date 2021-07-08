# kps_api.ServiceClassApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**service_class_get**](ServiceClassApi.md#service_class_get) | **GET** /v1.0/serviceclasses/{svcClassId} | Get a Service Class.
[**service_class_list**](ServiceClassApi.md#service_class_list) | **GET** /v1.0/serviceclasses | List Service Classes.

# **service_class_get**
> ServiceClass service_class_get(svc_class_id, authorization)

Get a Service Class.

Get a Service Class

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
api_instance = kps_api.ServiceClassApi(kps_api.ApiClient(configuration))
svc_class_id = 'svc_class_id_example' # str | Service Class ID
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Get a Service Class.
    api_response = api_instance.service_class_get(svc_class_id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceClassApi->service_class_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svc_class_id** | **str**| Service Class ID | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**ServiceClass**](ServiceClass.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_class_list**
> ServiceClassListPayload service_class_list(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter, scope=scope, type=type, svc_version=svc_version, tags=tags)

List Service Classes.

List Service Classes

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
api_instance = kps_api.ServiceClassApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)
scope = 'scope_example' # str | Scope of the Service Class (optional)
type = 'type_example' # str | Type of the Service Class (optional)
svc_version = 'svc_version_example' # str | Version of the Service Class (optional)
tags = ['tags_example'] # list[str] | Tags on the Service Class (optional)

try:
    # List Service Classes.
    api_response = api_instance.service_class_list(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter, scope=scope, type=type, svc_version=svc_version, tags=tags)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceClassApi->service_class_list: %s\n" % e)
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
 **tags** | [**list[str]**](str.md)| Tags on the Service Class | [optional] 

### Return type

[**ServiceClassListPayload**](ServiceClassListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

