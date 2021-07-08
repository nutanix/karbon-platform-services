# kps_api.LogApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**application_log_entries_get_v2**](LogApi.md#application_log_entries_get_v2) | **GET** /v1.0/logs/applications/{id} | Lists applications log entries specific to an application.
[**application_log_entries_list_v2**](LogApi.md#application_log_entries_list_v2) | **GET** /v1.0/logs/applications | Lists application log entries.
[**edge_log_entries_get_v2**](LogApi.md#edge_log_entries_get_v2) | **GET** /v1.0/logs/edges/{id} | Lists infrastructure log entries for an edge.
[**edge_log_entries_list_v2**](LogApi.md#edge_log_entries_list_v2) | **GET** /v1.0/logs/edges | Lists infrastructure log entries for edges.
[**log_entries_list_v2**](LogApi.md#log_entries_list_v2) | **GET** /v1.0/logs/entries | Lists log entries.
[**log_entry_delete_v2**](LogApi.md#log_entry_delete_v2) | **DELETE** /v1.0/logs/entries/{id} | Delete log entry by ID.
[**log_request_download_v2**](LogApi.md#log_request_download_v2) | **POST** /v1.0/logs/requestdownload | Request log download.
[**log_request_upload_v2**](LogApi.md#log_request_upload_v2) | **POST** /v1.0/logs/requestupload | Request log upload.
[**log_stream_endpoints**](LogApi.md#log_stream_endpoints) | **POST** /v1.0/logs/stream/endpoints | Get the endpoints to stream logs for a given container from an edge.

# **application_log_entries_get_v2**
> LogEntriesListPayload application_log_entries_get_v2(id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)

Lists applications log entries specific to an application.

Retrieve application log entries specific to an application. Use filter on edge ID and batch ID to get the application log specific to an edge and a batch.

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
api_instance = kps_api.LogApi(kps_api.ApiClient(configuration))
id = 'id_example' # str | ID of the entity
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Lists applications log entries specific to an application.
    api_response = api_instance.application_log_entries_get_v2(id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling LogApi->application_log_entries_get_v2: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| ID of the entity | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **page_index** | **int**| 0-based index of the page to fetch results. | [optional] 
 **page_size** | **int**| Item count of each page. | [optional] 
 **order_by** | [**list[str]**](str.md)| Specify result order. Zero or more entries with format: &amp;ltkey&gt; [desc] where orderByKeys lists allowed keys in each response. | [optional] 
 **filter** | **str**| Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE &#x27;foo%&#x27;. Supported filter keys are the same as order by keys. | [optional] 

### Return type

[**LogEntriesListPayload**](LogEntriesListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **application_log_entries_list_v2**
> LogEntriesListPayload application_log_entries_list_v2(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)

Lists application log entries.

Retrieve all the application log entries. Use filter on edge ID and batch ID to get the application log specific to an edge and a batch.

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
api_instance = kps_api.LogApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Lists application log entries.
    api_response = api_instance.application_log_entries_list_v2(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling LogApi->application_log_entries_list_v2: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **page_index** | **int**| 0-based index of the page to fetch results. | [optional] 
 **page_size** | **int**| Item count of each page. | [optional] 
 **order_by** | [**list[str]**](str.md)| Specify result order. Zero or more entries with format: &amp;ltkey&gt; [desc] where orderByKeys lists allowed keys in each response. | [optional] 
 **filter** | **str**| Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE &#x27;foo%&#x27;. Supported filter keys are the same as order by keys. | [optional] 

### Return type

[**LogEntriesListPayload**](LogEntriesListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **edge_log_entries_get_v2**
> LogEntriesListPayload edge_log_entries_get_v2(id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)

Lists infrastructure log entries for an edge.

Retrieve infrastructure log entries specific to an edge. Use filter on batch ID to get logs entries specific to a batch.

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
api_instance = kps_api.LogApi(kps_api.ApiClient(configuration))
id = 'id_example' # str | ID of the entity
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Lists infrastructure log entries for an edge.
    api_response = api_instance.edge_log_entries_get_v2(id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling LogApi->edge_log_entries_get_v2: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| ID of the entity | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **page_index** | **int**| 0-based index of the page to fetch results. | [optional] 
 **page_size** | **int**| Item count of each page. | [optional] 
 **order_by** | [**list[str]**](str.md)| Specify result order. Zero or more entries with format: &amp;ltkey&gt; [desc] where orderByKeys lists allowed keys in each response. | [optional] 
 **filter** | **str**| Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE &#x27;foo%&#x27;. Supported filter keys are the same as order by keys. | [optional] 

### Return type

[**LogEntriesListPayload**](LogEntriesListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **edge_log_entries_list_v2**
> LogEntriesListPayload edge_log_entries_list_v2(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)

Lists infrastructure log entries for edges.

Retrieve all infrastructure log entries.

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
api_instance = kps_api.LogApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Lists infrastructure log entries for edges.
    api_response = api_instance.edge_log_entries_list_v2(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling LogApi->edge_log_entries_list_v2: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **page_index** | **int**| 0-based index of the page to fetch results. | [optional] 
 **page_size** | **int**| Item count of each page. | [optional] 
 **order_by** | [**list[str]**](str.md)| Specify result order. Zero or more entries with format: &amp;ltkey&gt; [desc] where orderByKeys lists allowed keys in each response. | [optional] 
 **filter** | **str**| Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE &#x27;foo%&#x27;. Supported filter keys are the same as order by keys. | [optional] 

### Return type

[**LogEntriesListPayload**](LogEntriesListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **log_entries_list_v2**
> LogEntriesListPayload log_entries_list_v2(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)

Lists log entries.

Retrieve all log entries.

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
api_instance = kps_api.LogApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Lists log entries.
    api_response = api_instance.log_entries_list_v2(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling LogApi->log_entries_list_v2: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **page_index** | **int**| 0-based index of the page to fetch results. | [optional] 
 **page_size** | **int**| Item count of each page. | [optional] 
 **order_by** | [**list[str]**](str.md)| Specify result order. Zero or more entries with format: &amp;ltkey&gt; [desc] where orderByKeys lists allowed keys in each response. | [optional] 
 **filter** | **str**| Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE &#x27;foo%&#x27;. Supported filter keys are the same as order by keys. | [optional] 

### Return type

[**LogEntriesListPayload**](LogEntriesListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **log_entry_delete_v2**
> DeleteDocumentResponse log_entry_delete_v2(id, authorization)

Delete log entry by ID.

Deletes the log entry with the given id.

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
api_instance = kps_api.LogApi(kps_api.ApiClient(configuration))
id = 'id_example' # str | ID of the entity
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Delete log entry by ID.
    api_response = api_instance.log_entry_delete_v2(id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling LogApi->log_entry_delete_v2: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| ID of the entity | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**DeleteDocumentResponse**](DeleteDocumentResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **log_request_download_v2**
> LogDownloadPayload log_request_download_v2(body, authorization)

Request log download.

Generates the log download URL.

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
api_instance = kps_api.LogApi(kps_api.ApiClient(configuration))
body = kps_api.RequestLogDownloadPayload() # RequestLogDownloadPayload | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Request log download.
    api_response = api_instance.log_request_download_v2(body, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling LogApi->log_request_download_v2: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RequestLogDownloadPayload**](RequestLogDownloadPayload.md)|  | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**LogDownloadPayload**](LogDownloadPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **log_request_upload_v2**
> list[LogUploadPayload] log_request_upload_v2(body, authorization)

Request log upload.

Request edges to upload logs to cloud storage.

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
api_instance = kps_api.LogApi(kps_api.ApiClient(configuration))
body = kps_api.RequestLogUploadPayload() # RequestLogUploadPayload | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Request log upload.
    api_response = api_instance.log_request_upload_v2(body, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling LogApi->log_request_upload_v2: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RequestLogUploadPayload**](RequestLogUploadPayload.md)|  | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**list[LogUploadPayload]**](LogUploadPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **log_stream_endpoints**
> LogStreamResponsePayload log_stream_endpoints(body, authorization)

Get the endpoints to stream logs for a given container from an edge.

Get the endpoints to stream logs for a given container from an edge.

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
api_instance = kps_api.LogApi(kps_api.ApiClient(configuration))
body = kps_api.LogStream() # LogStream | A description of the log streaming request.
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Get the endpoints to stream logs for a given container from an edge.
    api_response = api_instance.log_stream_endpoints(body, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling LogApi->log_stream_endpoints: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LogStream**](LogStream.md)| A description of the log streaming request. | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**LogStreamResponsePayload**](LogStreamResponsePayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

