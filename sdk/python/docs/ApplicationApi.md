# kps_api.ApplicationApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**application_create_v2**](ApplicationApi.md#application_create_v2) | **POST** /v1.0/applications | Create an application.
[**application_delete_v2**](ApplicationApi.md#application_delete_v2) | **DELETE** /v1.0/applications/{id} | Delete application specified by the application ID.
[**application_get_v2**](ApplicationApi.md#application_get_v2) | **GET** /v1.0/applications/{id} | Get application by application ID.
[**application_list_v2**](ApplicationApi.md#application_list_v2) | **GET** /v1.0/applications | Get all applications.
[**application_update_v3**](ApplicationApi.md#application_update_v3) | **PUT** /v1.0/applications/{id} | Update a specific application with ID {id}.
[**get_application_containers**](ApplicationApi.md#get_application_containers) | **GET** /v1.0/applications/{id}/containers/{edgeId} | Get containers of an application specified by Application ID running on a specific edge.
[**project_get_applications_v2**](ApplicationApi.md#project_get_applications_v2) | **GET** /v1.0/projects/{projectId}/applications | Get all applications in a project according to project ID.

# **application_create_v2**
> CreateDocumentResponseV2 application_create_v2(body, authorization)

Create an application.

Create an application.

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
api_instance = kps_api.ApplicationApi(kps_api.ApiClient(configuration))
body = kps_api.ApplicationV2() # ApplicationV2 | Describes the application creation request.
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Create an application.
    api_response = api_instance.application_create_v2(body, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ApplicationApi->application_create_v2: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationV2**](ApplicationV2.md)| Describes the application creation request. | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**CreateDocumentResponseV2**](CreateDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **application_delete_v2**
> DeleteDocumentResponseV2 application_delete_v2(authorization, id)

Delete application specified by the application ID.

Deletes the application with the given ID {id}.

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
api_instance = kps_api.ApplicationApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
id = 'id_example' # str | ID of the entity

try:
    # Delete application specified by the application ID.
    api_response = api_instance.application_delete_v2(authorization, id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ApplicationApi->application_delete_v2: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **id** | **str**| ID of the entity | 

### Return type

[**DeleteDocumentResponseV2**](DeleteDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **application_get_v2**
> ApplicationV2 application_get_v2(authorization, id)

Get application by application ID.

Retrieves the application with the given ID {id}.

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
api_instance = kps_api.ApplicationApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
id = 'id_example' # str | ID of the entity

try:
    # Get application by application ID.
    api_response = api_instance.application_get_v2(authorization, id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ApplicationApi->application_get_v2: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **id** | **str**| ID of the entity | 

### Return type

[**ApplicationV2**](ApplicationV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **application_list_v2**
> ApplicationListResponsePayload application_list_v2(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)

Get all applications.

Retrieves a list of all applications.

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
api_instance = kps_api.ApplicationApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Get all applications.
    api_response = api_instance.application_list_v2(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ApplicationApi->application_list_v2: %s\n" % e)
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

[**ApplicationListResponsePayload**](ApplicationListResponsePayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **application_update_v3**
> UpdateDocumentResponseV2 application_update_v3(body, authorization, id)

Update a specific application with ID {id}.

Update a specific application with ID {id}. You cannot change the project associated with the application or the application ID. You can change all other attributes.

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
api_instance = kps_api.ApplicationApi(kps_api.ApiClient(configuration))
body = kps_api.ApplicationV2() # ApplicationV2 | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
id = 'id_example' # str | ID of the entity

try:
    # Update a specific application with ID {id}.
    api_response = api_instance.application_update_v3(body, authorization, id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ApplicationApi->application_update_v3: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationV2**](ApplicationV2.md)|  | 
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

# **get_application_containers**
> ApplicationContainers get_application_containers(authorization, id, edge_id)

Get containers of an application specified by Application ID running on a specific edge.

Gets the containers of an application with the given ID {id} running on edge with id {edgeId}.

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
api_instance = kps_api.ApplicationApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
id = 'id_example' # str | ID of the entity
edge_id = 'edge_id_example' # str | ID for the edge

try:
    # Get containers of an application specified by Application ID running on a specific edge.
    api_response = api_instance.get_application_containers(authorization, id, edge_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ApplicationApi->get_application_containers: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **id** | **str**| ID of the entity | 
 **edge_id** | **str**| ID for the edge | 

### Return type

[**ApplicationContainers**](ApplicationContainers.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **project_get_applications_v2**
> ApplicationListResponsePayload project_get_applications_v2(authorization, project_id, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)

Get all applications in a project according to project ID.

Retrieves a list of all applications in a project with ID {projectId}.

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
api_instance = kps_api.ApplicationApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
project_id = 'project_id_example' # str | ID for the project
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Get all applications in a project according to project ID.
    api_response = api_instance.project_get_applications_v2(authorization, project_id, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ApplicationApi->project_get_applications_v2: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **project_id** | **str**| ID for the project | 
 **page_index** | **int**| 0-based index of the page to fetch results. | [optional] 
 **page_size** | **int**| Item count of each page. | [optional] 
 **order_by** | [**list[str]**](str.md)| Specify result order. Zero or more entries with format: &amp;ltkey&gt; [desc] where orderByKeys lists allowed keys in each response. | [optional] 
 **filter** | **str**| Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE &#x27;foo%&#x27;. Supported filter keys are the same as order by keys. | [optional] 

### Return type

[**ApplicationListResponsePayload**](ApplicationListResponsePayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

