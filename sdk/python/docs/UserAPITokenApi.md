# kps_api.UserAPITokenApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**user_api_token_create**](UserAPITokenApi.md#user_api_token_create) | **POST** /v1.0/userapitokens | Create a user API token.
[**user_api_token_delete**](UserAPITokenApi.md#user_api_token_delete) | **DELETE** /v1.0/userapitokens/{id} | Delete current user API token.
[**user_api_token_get**](UserAPITokenApi.md#user_api_token_get) | **GET** /v1.0/userapitokens | Get current user API tokens.
[**user_api_token_list**](UserAPITokenApi.md#user_api_token_list) | **GET** /v1.0/userapitokensall | Get all user API tokens.
[**user_api_token_update**](UserAPITokenApi.md#user_api_token_update) | **PUT** /v1.0/userapitokens/{id} | Update user API token.

# **user_api_token_create**
> UserApiTokenCreated user_api_token_create(body, authorization)

Create a user API token.

Creates a user API token.

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
api_instance = kps_api.UserAPITokenApi(kps_api.ApiClient(configuration))
body = kps_api.UserApiTokenCreatePayload() # UserApiTokenCreatePayload | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Create a user API token.
    api_response = api_instance.user_api_token_create(body, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling UserAPITokenApi->user_api_token_create: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserApiTokenCreatePayload**](UserApiTokenCreatePayload.md)|  | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**UserApiTokenCreated**](UserApiTokenCreated.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **user_api_token_delete**
> DeleteDocumentResponseV2 user_api_token_delete(id, authorization)

Delete current user API token.

Deletes the API token with the given id for the current user.

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
api_instance = kps_api.UserAPITokenApi(kps_api.ApiClient(configuration))
id = 'id_example' # str | ID of the entity
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Delete current user API token.
    api_response = api_instance.user_api_token_delete(id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling UserAPITokenApi->user_api_token_delete: %s\n" % e)
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

# **user_api_token_get**
> list[UserApiToken] user_api_token_get(authorization)

Get current user API tokens.

Retrieves the API tokens info for the current user.

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
api_instance = kps_api.UserAPITokenApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Get current user API tokens.
    api_response = api_instance.user_api_token_get(authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling UserAPITokenApi->user_api_token_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**list[UserApiToken]**](UserApiToken.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **user_api_token_list**
> list[UserApiToken] user_api_token_list(authorization)

Get all user API tokens.

Retrieves the API tokens info for all users. Must be infra admin for this to work.

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
api_instance = kps_api.UserAPITokenApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Get all user API tokens.
    api_response = api_instance.user_api_token_list(authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling UserAPITokenApi->user_api_token_list: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**list[UserApiToken]**](UserApiToken.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **user_api_token_update**
> UpdateDocumentResponseV2 user_api_token_update(body, authorization, id)

Update user API token.

Update the API token with the given id. Must be current user or infra admin.

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
api_instance = kps_api.UserAPITokenApi(kps_api.ApiClient(configuration))
body = kps_api.UserApiToken() # UserApiToken | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
id = 'id_example' # str | ID of the entity

try:
    # Update user API token.
    api_response = api_instance.user_api_token_update(body, authorization, id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling UserAPITokenApi->user_api_token_update: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserApiToken**](UserApiToken.md)|  | 
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

