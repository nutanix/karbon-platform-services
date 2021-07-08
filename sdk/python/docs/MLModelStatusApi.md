# kps_api.MLModelStatusApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**m_l_model_status_get**](MLModelStatusApi.md#m_l_model_status_get) | **GET** /v1.0/mlmodelstatuses/{id} | Get ML model status by model ID.
[**m_l_model_status_list**](MLModelStatusApi.md#m_l_model_status_list) | **GET** /v1.0/mlmodelstatuses | Get status for all ML models.

# **m_l_model_status_get**
> MLModelStatusListPayload m_l_model_status_get(id, authorization, page_index=page_index, page_size=page_size)

Get ML model status by model ID.

Retrieve status for an ML model with the given ID {id}.

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
api_instance = kps_api.MLModelStatusApi(kps_api.ApiClient(configuration))
id = 'id_example' # str | ID of the entity
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)

try:
    # Get ML model status by model ID.
    api_response = api_instance.m_l_model_status_get(id, authorization, page_index=page_index, page_size=page_size)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling MLModelStatusApi->m_l_model_status_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| ID of the entity | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **page_index** | **int**| 0-based index of the page to fetch results. | [optional] 
 **page_size** | **int**| Item count of each page. | [optional] 

### Return type

[**MLModelStatusListPayload**](MLModelStatusListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **m_l_model_status_list**
> MLModelStatusListPayload m_l_model_status_list(authorization, page_index=page_index, page_size=page_size)

Get status for all ML models.

Retrieves status for all ML models.

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
api_instance = kps_api.MLModelStatusApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)

try:
    # Get status for all ML models.
    api_response = api_instance.m_l_model_status_list(authorization, page_index=page_index, page_size=page_size)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling MLModelStatusApi->m_l_model_status_list: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **page_index** | **int**| 0-based index of the page to fetch results. | [optional] 
 **page_size** | **int**| Item count of each page. | [optional] 

### Return type

[**MLModelStatusListPayload**](MLModelStatusListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

