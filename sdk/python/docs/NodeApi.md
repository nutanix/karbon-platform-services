# kps_api.NodeApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**node_create**](NodeApi.md#node_create) | **POST** /v1.0/nodes | Create a node.
[**node_delete**](NodeApi.md#node_delete) | **DELETE** /v1.0/nodes/{nodeId} | Delete a node as specified by node ID.
[**node_get**](NodeApi.md#node_get) | **GET** /v1.0/nodes/{nodeId} | Get a node as specified by node ID.
[**node_list**](NodeApi.md#node_list) | **GET** /v1.0/nodes | Get all service domain nodes.
[**node_onboarded**](NodeApi.md#node_onboarded) | **POST** /v1.0/nodeonboarded | Update node post onboard info.
[**node_update**](NodeApi.md#node_update) | **PUT** /v1.0/nodes/{nodeId} | Update a node as specified by node ID.
[**project_get_nodes**](NodeApi.md#project_get_nodes) | **GET** /v1.0/projects/{projectId}/nodes | Get all service domain nodes associated with a project by project ID.

# **node_create**
> CreateDocumentResponseV2 node_create(body, authorization)

Create a node.

Create a node.

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
api_instance = kps_api.NodeApi(kps_api.ApiClient(configuration))
body = kps_api.Node() # Node | Parameters and values used when creating a node
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Create a node.
    api_response = api_instance.node_create(body, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling NodeApi->node_create: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Node**](Node.md)| Parameters and values used when creating a node | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**CreateDocumentResponseV2**](CreateDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **node_delete**
> DeleteDocumentResponseV2 node_delete(node_id, authorization)

Delete a node as specified by node ID.

Deletes the node with the given ID  {nodeId}.

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
api_instance = kps_api.NodeApi(kps_api.ApiClient(configuration))
node_id = 'node_id_example' # str | ID for the node
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Delete a node as specified by node ID.
    api_response = api_instance.node_delete(node_id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling NodeApi->node_delete: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **node_id** | **str**| ID for the node | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**DeleteDocumentResponseV2**](DeleteDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **node_get**
> Node node_get(node_id, authorization)

Get a node as specified by node ID.

Retrieves the node with the given ID {nodeId}.

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
api_instance = kps_api.NodeApi(kps_api.ApiClient(configuration))
node_id = 'node_id_example' # str | ID for the node
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Get a node as specified by node ID.
    api_response = api_instance.node_get(node_id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling NodeApi->node_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **node_id** | **str**| ID for the node | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**Node**](Node.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **node_list**
> NodeListPayload node_list(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)

Get all service domain nodes.

Retrieves all service domain nodes for your account.

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
api_instance = kps_api.NodeApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Get all service domain nodes.
    api_response = api_instance.node_list(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling NodeApi->node_list: %s\n" % e)
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

[**NodeListPayload**](NodeListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **node_onboarded**
> UpdateDocumentResponseV2 node_onboarded(body)

Update node post onboard info.

Updates the onboard info by node ID.

### Example
```python
from __future__ import print_function
import time
import kps_api
from kps_api.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = kps_api.NodeApi()
body = kps_api.NodeOnboardInfo() # NodeOnboardInfo | 

try:
    # Update node post onboard info.
    api_response = api_instance.node_onboarded(body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling NodeApi->node_onboarded: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NodeOnboardInfo**](NodeOnboardInfo.md)|  | 

### Return type

[**UpdateDocumentResponseV2**](UpdateDocumentResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **node_update**
> UpdateDocumentResponseV2 node_update(body, authorization, node_id)

Update a node as specified by node ID.

Updates a node by its ID {nodeId}.

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
api_instance = kps_api.NodeApi(kps_api.ApiClient(configuration))
body = kps_api.Node() # Node | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
node_id = 'node_id_example' # str | ID for the node

try:
    # Update a node as specified by node ID.
    api_response = api_instance.node_update(body, authorization, node_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling NodeApi->node_update: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Node**](Node.md)|  | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **node_id** | **str**| ID for the node | 

### Return type

[**UpdateDocumentResponseV2**](UpdateDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **project_get_nodes**
> NodeListPayload project_get_nodes(project_id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)

Get all service domain nodes associated with a project by project ID.

Retrieves all service domain nodes for a project by project ID {projectId}.

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
api_instance = kps_api.NodeApi(kps_api.ApiClient(configuration))
project_id = 'project_id_example' # str | ID for the project
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Get all service domain nodes associated with a project by project ID.
    api_response = api_instance.project_get_nodes(project_id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling NodeApi->project_get_nodes: %s\n" % e)
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

[**NodeListPayload**](NodeListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

