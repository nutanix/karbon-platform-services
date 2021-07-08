# kps_api.DataPipelineApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**data_pipeline_create**](DataPipelineApi.md#data_pipeline_create) | **POST** /v1.0/datapipelines | Creates a data pipeline.
[**data_pipeline_delete**](DataPipelineApi.md#data_pipeline_delete) | **DELETE** /v1.0/datapipelines/{id} | Deletes data pipeline by its ID.
[**data_pipeline_get**](DataPipelineApi.md#data_pipeline_get) | **GET** /v1.0/datapipelines/{id} | Lists data pipeline by its ID.
[**data_pipeline_list**](DataPipelineApi.md#data_pipeline_list) | **GET** /v1.0/datapipelines | Gets data pipelines.
[**data_pipeline_update**](DataPipelineApi.md#data_pipeline_update) | **PUT** /v1.0/datapipelines/{id} | Updates a data pipeline by its ID
[**get_data_pipeline_containers**](DataPipelineApi.md#get_data_pipeline_containers) | **GET** /v1.0/datapipelines/{id}/containers/{edgeId} | Get containers of a data pipeline specified by Datapipeline ID running on a specific edge.
[**project_get_data_pipelines**](DataPipelineApi.md#project_get_data_pipelines) | **GET** /v1.0/projects/{projectId}/datapipelines | Gets data pipelines for a project.

# **data_pipeline_create**
> CreateDocumentResponseV2 data_pipeline_create(body, authorization)

Creates a data pipeline.

Create a data pipeline.

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
api_instance = kps_api.DataPipelineApi(kps_api.ApiClient(configuration))
body = kps_api.DataPipeline() # DataPipeline | This is a data pipeline creation request description
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Creates a data pipeline.
    api_response = api_instance.data_pipeline_create(body, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DataPipelineApi->data_pipeline_create: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DataPipeline**](DataPipeline.md)| This is a data pipeline creation request description | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**CreateDocumentResponseV2**](CreateDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **data_pipeline_delete**
> DeleteDocumentResponseV2 data_pipeline_delete(id, authorization)

Deletes data pipeline by its ID.

Delete the data pipeline with the given id {id}.

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
api_instance = kps_api.DataPipelineApi(kps_api.ApiClient(configuration))
id = 'id_example' # str | ID of the entity
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Deletes data pipeline by its ID.
    api_response = api_instance.data_pipeline_delete(id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DataPipelineApi->data_pipeline_delete: %s\n" % e)
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

# **data_pipeline_get**
> DataPipeline data_pipeline_get(id, authorization)

Lists data pipeline by its ID.

Retrieves a data pipelines with a given ID {id}.

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
api_instance = kps_api.DataPipelineApi(kps_api.ApiClient(configuration))
id = 'id_example' # str | ID of the entity
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Lists data pipeline by its ID.
    api_response = api_instance.data_pipeline_get(id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DataPipelineApi->data_pipeline_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| ID of the entity | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**DataPipeline**](DataPipeline.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **data_pipeline_list**
> DataPipelineListPayload data_pipeline_list(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)

Gets data pipelines.

Retrieves all data pipelines for a tenant.

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
api_instance = kps_api.DataPipelineApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Gets data pipelines.
    api_response = api_instance.data_pipeline_list(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DataPipelineApi->data_pipeline_list: %s\n" % e)
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

[**DataPipelineListPayload**](DataPipelineListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **data_pipeline_update**
> UpdateDocumentResponseV2 data_pipeline_update(body, authorization, id)

Updates a data pipeline by its ID

Update a data pipeline with a given ID {id}.

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
api_instance = kps_api.DataPipelineApi(kps_api.ApiClient(configuration))
body = kps_api.DataPipeline() # DataPipeline | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
id = 'id_example' # str | ID of the entity

try:
    # Updates a data pipeline by its ID
    api_response = api_instance.data_pipeline_update(body, authorization, id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DataPipelineApi->data_pipeline_update: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DataPipeline**](DataPipeline.md)|  | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **id** | **str**| ID of the entity | 

### Return type

[**UpdateDocumentResponseV2**](UpdateDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_data_pipeline_containers**
> DataPipelineContainers get_data_pipeline_containers(id, edge_id, authorization)

Get containers of a data pipeline specified by Datapipeline ID running on a specific edge.

Gets the containers of a data pipeline with the given ID {id} running on edge with id {edgeId}.

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
api_instance = kps_api.DataPipelineApi(kps_api.ApiClient(configuration))
id = 'id_example' # str | ID of the entity
edge_id = 'edge_id_example' # str | ID for the edge
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Get containers of a data pipeline specified by Datapipeline ID running on a specific edge.
    api_response = api_instance.get_data_pipeline_containers(id, edge_id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DataPipelineApi->get_data_pipeline_containers: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| ID of the entity | 
 **edge_id** | **str**| ID for the edge | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**DataPipelineContainers**](DataPipelineContainers.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **project_get_data_pipelines**
> DataPipelineListPayload project_get_data_pipelines(project_id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)

Gets data pipelines for a project.

Retrieves all data pipelines for a project with a given ID {projectId}.

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
api_instance = kps_api.DataPipelineApi(kps_api.ApiClient(configuration))
project_id = 'project_id_example' # str | ID for the project
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Gets data pipelines for a project.
    api_response = api_instance.project_get_data_pipelines(project_id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DataPipelineApi->project_get_data_pipelines: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **str**| ID for the project | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **page_index** | **int**| 0-based index of the page to fetch results. | [optional] 
 **page_size** | **int**| Item count of each page. | [optional] 
 **order_by** | [**list[str]**](str.md)| Specify result order. Zero or more entries with format: &amp;ltkey&gt; [desc] where orderByKeys lists allowed keys in each response. | [optional] 
 **filter** | **str**| Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE &#x27;foo%&#x27;. Supported filter keys are the same as order by keys. | [optional] 

### Return type

[**DataPipelineListPayload**](DataPipelineListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

