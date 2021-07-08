# kps_api.ConnectorConfigApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**connector_config_create**](ConnectorConfigApi.md#connector_config_create) | **POST** /v1.0/connector/configs | Create a connector config parameters.
[**connector_config_delete**](ConnectorConfigApi.md#connector_config_delete) | **DELETE** /v1.0/connector/configs/{id} | Delete a specific connector config parameters.
[**connector_config_get**](ConnectorConfigApi.md#connector_config_get) | **GET** /v1.0/connector/configs/{id} | Get a connector config parameters by ID.
[**connector_config_list**](ConnectorConfigApi.md#connector_config_list) | **GET** /v1.0/connector/instances/{id}/configs | Get a connector config parameters for connector instance by ID.
[**connector_config_update**](ConnectorConfigApi.md#connector_config_update) | **PUT** /v1.0/connector/configs/{id} | Update a connector config parameters.

# **connector_config_create**
> CreateDocumentResponseV2 connector_config_create(body, authorization)

Create a connector config parameters.

Create a connector config parameters.

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
api_instance = kps_api.ConnectorConfigApi(kps_api.ApiClient(configuration))
body = kps_api.ConnectorConfig() # ConnectorConfig | Parameters and values used when creating or updating a connector config
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Create a connector config parameters.
    api_response = api_instance.connector_config_create(body, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ConnectorConfigApi->connector_config_create: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConnectorConfig**](ConnectorConfig.md)| Parameters and values used when creating or updating a connector config | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**CreateDocumentResponseV2**](CreateDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **connector_config_delete**
> DeleteDocumentResponseV2 connector_config_delete(id, authorization)

Delete a specific connector config parameters.

Delete a connector config parameters with a given ID {id}.

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
api_instance = kps_api.ConnectorConfigApi(kps_api.ApiClient(configuration))
id = 'id_example' # str | ID of the entity
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Delete a specific connector config parameters.
    api_response = api_instance.connector_config_delete(id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ConnectorConfigApi->connector_config_delete: %s\n" % e)
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

# **connector_config_get**
> ConnectorConfig connector_config_get(id, authorization)

Get a connector config parameters by ID.

Get a connector config parameters according to its given ID {id}.

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
api_instance = kps_api.ConnectorConfigApi(kps_api.ApiClient(configuration))
id = 'id_example' # str | ID of the entity
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Get a connector config parameters by ID.
    api_response = api_instance.connector_config_get(id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ConnectorConfigApi->connector_config_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| ID of the entity | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**ConnectorConfig**](ConnectorConfig.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **connector_config_list**
> ConnectorConfigListResponsePayload connector_config_list(id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)

Get a connector config parameters for connector instance by ID.

Get a connector config parameters according to its instance ID {id}.

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
api_instance = kps_api.ConnectorConfigApi(kps_api.ApiClient(configuration))
id = 'id_example' # str | ID of the entity
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Get a connector config parameters for connector instance by ID.
    api_response = api_instance.connector_config_list(id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ConnectorConfigApi->connector_config_list: %s\n" % e)
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

[**ConnectorConfigListResponsePayload**](ConnectorConfigListResponsePayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **connector_config_update**
> UpdateDocumentResponseV2 connector_config_update(body, authorization, id)

Update a connector config parameters.

Update a connector config parameters.

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
api_instance = kps_api.ConnectorConfigApi(kps_api.ApiClient(configuration))
body = kps_api.ConnectorConfig() # ConnectorConfig | Parameters and values used when creating or updating a connector config
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
id = 'id_example' # str | ID of the entity

try:
    # Update a connector config parameters.
    api_response = api_instance.connector_config_update(body, authorization, id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ConnectorConfigApi->connector_config_update: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConnectorConfig**](ConnectorConfig.md)| Parameters and values used when creating or updating a connector config | 
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

