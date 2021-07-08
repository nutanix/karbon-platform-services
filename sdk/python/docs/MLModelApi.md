# kps_api.MLModelApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**m_l_model_create**](MLModelApi.md#m_l_model_create) | **POST** /v1.0/mlmodels | Create a machine learning model.
[**m_l_model_delete**](MLModelApi.md#m_l_model_delete) | **DELETE** /v1.0/mlmodels/{id} | Delete a machine learning model  by its ID.
[**m_l_model_get**](MLModelApi.md#m_l_model_get) | **GET** /v1.0/mlmodels/{id} | Get machine learning model by its ID.
[**m_l_model_list**](MLModelApi.md#m_l_model_list) | **GET** /v1.0/mlmodels | Lists machine learning models.
[**m_l_model_update**](MLModelApi.md#m_l_model_update) | **PUT** /v1.0/mlmodels/{id} | Update a machine learning model by its ID.
[**m_l_model_version_create**](MLModelApi.md#m_l_model_version_create) | **POST** /v1.0/mlmodels/{id}/versions | Create a new version of the machine learning model by its ID.
[**m_l_model_version_delete**](MLModelApi.md#m_l_model_version_delete) | **DELETE** /v1.0/mlmodels/{id}/versions/{model_version} | Delete the version of the machine learning model by its ID.
[**m_l_model_version_update**](MLModelApi.md#m_l_model_version_update) | **PUT** /v1.0/mlmodels/{id}/versions/{model_version} | Update the version of the machine learning model by its ID.
[**m_l_model_version_url_get**](MLModelApi.md#m_l_model_version_url_get) | **GET** /v1.0/mlmodels/{id}/versions/{model_version}/url | Get a pre-signed URL for the machine learning model according to its ID and version.
[**project_get_ml_models**](MLModelApi.md#project_get_ml_models) | **GET** /v1.0/projects/{projectId}/mlmodels | Lists project machine learning models by project ID.

# **m_l_model_create**
> CreateDocumentResponseV2 m_l_model_create(body, authorization)

Create a machine learning model.

Creates a machine learning model.

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
api_instance = kps_api.MLModelApi(kps_api.ApiClient(configuration))
body = kps_api.MLModelMetadata() # MLModelMetadata | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from the login API response.

try:
    # Create a machine learning model.
    api_response = api_instance.m_l_model_create(body, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling MLModelApi->m_l_model_create: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MLModelMetadata**](MLModelMetadata.md)|  | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from the login API response. | 

### Return type

[**CreateDocumentResponseV2**](CreateDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **m_l_model_delete**
> DeleteDocumentResponseV2 m_l_model_delete(id, authorization)

Delete a machine learning model  by its ID.

Deletes a machine learning model by its given ID.

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
api_instance = kps_api.MLModelApi(kps_api.ApiClient(configuration))
id = 'id_example' # str | ID of the entity
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from the login API response.

try:
    # Delete a machine learning model  by its ID.
    api_response = api_instance.m_l_model_delete(id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling MLModelApi->m_l_model_delete: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| ID of the entity | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from the login API response. | 

### Return type

[**DeleteDocumentResponseV2**](DeleteDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **m_l_model_get**
> MLModel m_l_model_get(id, authorization)

Get machine learning model by its ID.

Retrieves a machine learning model by its given ID.

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
api_instance = kps_api.MLModelApi(kps_api.ApiClient(configuration))
id = 'id_example' # str | ID of the entity
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from the login API response.

try:
    # Get machine learning model by its ID.
    api_response = api_instance.m_l_model_get(id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling MLModelApi->m_l_model_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| ID of the entity | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from the login API response. | 

### Return type

[**MLModel**](MLModel.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **m_l_model_list**
> MLModelListResponsePayload m_l_model_list(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)

Lists machine learning models.

Retrieve all machine learning models.

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
api_instance = kps_api.MLModelApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from the login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Lists machine learning models.
    api_response = api_instance.m_l_model_list(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling MLModelApi->m_l_model_list: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from the login API response. | 
 **page_index** | **int**| 0-based index of the page to fetch results. | [optional] 
 **page_size** | **int**| Item count of each page. | [optional] 
 **order_by** | [**list[str]**](str.md)| Specify result order. Zero or more entries with format: &amp;ltkey&gt; [desc] where orderByKeys lists allowed keys in each response. | [optional] 
 **filter** | **str**| Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE &#x27;foo%&#x27;. Supported filter keys are the same as order by keys. | [optional] 

### Return type

[**MLModelListResponsePayload**](MLModelListResponsePayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **m_l_model_update**
> UpdateDocumentResponseV2 m_l_model_update(body, authorization, id)

Update a machine learning model by its ID.

Updates a machine learning model by its given ID.

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
api_instance = kps_api.MLModelApi(kps_api.ApiClient(configuration))
body = kps_api.MLModelMetadata() # MLModelMetadata | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from the login API response.
id = 'id_example' # str | ID of the entity

try:
    # Update a machine learning model by its ID.
    api_response = api_instance.m_l_model_update(body, authorization, id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling MLModelApi->m_l_model_update: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MLModelMetadata**](MLModelMetadata.md)|  | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from the login API response. | 
 **id** | **str**| ID of the entity | 

### Return type

[**UpdateDocumentResponseV2**](UpdateDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **m_l_model_version_create**
> CreateDocumentResponseV2 m_l_model_version_create(payload, authorization, model_version, id, description=description)

Create a new version of the machine learning model by its ID.

Create a new version of the machine learning model by its given ID.

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
api_instance = kps_api.MLModelApi(kps_api.ApiClient(configuration))
payload = 'payload_example' # file | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from the login API response.
model_version = 789 # int | Model version, a positive integer.
id = 'id_example' # str | ID of the entity
description = 'description_example' # str | Model version description. (optional)

try:
    # Create a new version of the machine learning model by its ID.
    api_response = api_instance.m_l_model_version_create(payload, authorization, model_version, id, description=description)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling MLModelApi->m_l_model_version_create: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | **file**|  | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from the login API response. | 
 **model_version** | **int**| Model version, a positive integer. | 
 **id** | **str**| ID of the entity | 
 **description** | **str**| Model version description. | [optional] 

### Return type

[**CreateDocumentResponseV2**](CreateDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **m_l_model_version_delete**
> DeleteDocumentResponseV2 m_l_model_version_delete(id, model_version, authorization)

Delete the version of the machine learning model by its ID.

Deletes the version of the machine learning model by machine learning model ID.

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
api_instance = kps_api.MLModelApi(kps_api.ApiClient(configuration))
id = 'id_example' # str | ID of the entity
model_version = 789 # int | Model version, a positive integer.
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from the login API response.

try:
    # Delete the version of the machine learning model by its ID.
    api_response = api_instance.m_l_model_version_delete(id, model_version, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling MLModelApi->m_l_model_version_delete: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| ID of the entity | 
 **model_version** | **int**| Model version, a positive integer. | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from the login API response. | 

### Return type

[**DeleteDocumentResponseV2**](DeleteDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **m_l_model_version_update**
> UpdateDocumentResponseV2 m_l_model_version_update(id, model_version, authorization, description=description)

Update the version of the machine learning model by its ID.

Updates the version of the machine learning model by machine learning model ID.

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
api_instance = kps_api.MLModelApi(kps_api.ApiClient(configuration))
id = 'id_example' # str | ID of the entity
model_version = 789 # int | Model version, a positive integer.
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from the login API response.
description = 'description_example' # str | Model version description. (optional)

try:
    # Update the version of the machine learning model by its ID.
    api_response = api_instance.m_l_model_version_update(id, model_version, authorization, description=description)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling MLModelApi->m_l_model_version_update: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| ID of the entity | 
 **model_version** | **int**| Model version, a positive integer. | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from the login API response. | 
 **description** | **str**| Model version description. | [optional] 

### Return type

[**UpdateDocumentResponseV2**](UpdateDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **m_l_model_version_url_get**
> MLModelVersionURLGetResponsePayload m_l_model_version_url_get(id, model_version, authorization, expiration_duration=expiration_duration)

Get a pre-signed URL for the machine learning model according to its ID and version.

Retrieves a pre-signed URL for the machine learning model according to its ID and version.

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
api_instance = kps_api.MLModelApi(kps_api.ApiClient(configuration))
id = 'id_example' # str | ID of the entity
model_version = 789 # int | Model version, a positive integer.
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from the login API response.
expiration_duration = 789 # int | Model URL expiration duration in minutes. (optional)

try:
    # Get a pre-signed URL for the machine learning model according to its ID and version.
    api_response = api_instance.m_l_model_version_url_get(id, model_version, authorization, expiration_duration=expiration_duration)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling MLModelApi->m_l_model_version_url_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| ID of the entity | 
 **model_version** | **int**| Model version, a positive integer. | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from the login API response. | 
 **expiration_duration** | **int**| Model URL expiration duration in minutes. | [optional] 

### Return type

[**MLModelVersionURLGetResponsePayload**](MLModelVersionURLGetResponsePayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **project_get_ml_models**
> MLModelListResponsePayload project_get_ml_models(project_id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)

Lists project machine learning models by project ID.

Retrieves all machine learning models for a project by its given ID.

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
api_instance = kps_api.MLModelApi(kps_api.ApiClient(configuration))
project_id = 'project_id_example' # str | ID for the project
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from the login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Lists project machine learning models by project ID.
    api_response = api_instance.project_get_ml_models(project_id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling MLModelApi->project_get_ml_models: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project_id** | **str**| ID for the project | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from the login API response. | 
 **page_index** | **int**| 0-based index of the page to fetch results. | [optional] 
 **page_size** | **int**| Item count of each page. | [optional] 
 **order_by** | [**list[str]**](str.md)| Specify result order. Zero or more entries with format: &amp;ltkey&gt; [desc] where orderByKeys lists allowed keys in each response. | [optional] 
 **filter** | **str**| Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE &#x27;foo%&#x27;. Supported filter keys are the same as order by keys. | [optional] 

### Return type

[**MLModelListResponsePayload**](MLModelListResponsePayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

