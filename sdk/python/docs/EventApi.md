# kps_api.EventApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**query_events_v2**](EventApi.md#query_events_v2) | **POST** /v1.0/events | Lists events matching the provided filter.

# **query_events_v2**
> list[Event] query_events_v2(body, authorization, page_index=page_index, page_size=page_size)

Lists events matching the provided filter.

Retrieves all events matching the filter for a tenant.

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
api_instance = kps_api.EventApi(kps_api.ApiClient(configuration))
body = kps_api.EventFilterV2() # EventFilterV2 | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)

try:
    # Lists events matching the provided filter.
    api_response = api_instance.query_events_v2(body, authorization, page_index=page_index, page_size=page_size)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling EventApi->query_events_v2: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EventFilterV2**](EventFilterV2.md)|  | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **page_index** | **int**| 0-based index of the page to fetch results. | [optional] 
 **page_size** | **int**| Item count of each page. | [optional] 

### Return type

[**list[Event]**](Event.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

