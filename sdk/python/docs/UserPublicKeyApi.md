# kps_api.UserPublicKeyApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**user_public_key_delete**](UserPublicKeyApi.md#user_public_key_delete) | **DELETE** /v1.0/userpublickey | Delete current user public key.
[**user_public_key_get**](UserPublicKeyApi.md#user_public_key_get) | **GET** /v1.0/userpublickey | Get current user public key.
[**user_public_key_list**](UserPublicKeyApi.md#user_public_key_list) | **GET** /v1.0/userpublickeyall | Get all user public keys.
[**user_public_key_update**](UserPublicKeyApi.md#user_public_key_update) | **PUT** /v1.0/userpublickey | Upsert current user public key.

# **user_public_key_delete**
> DeleteDocumentResponseV2 user_public_key_delete(authorization)

Delete current user public key.

Deletes the public key for the current user.

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
api_instance = kps_api.UserPublicKeyApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Delete current user public key.
    api_response = api_instance.user_public_key_delete(authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling UserPublicKeyApi->user_public_key_delete: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**DeleteDocumentResponseV2**](DeleteDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **user_public_key_get**
> UserPublicKey user_public_key_get(authorization)

Get current user public key.

Retrieves the public key for the current user.

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
api_instance = kps_api.UserPublicKeyApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Get current user public key.
    api_response = api_instance.user_public_key_get(authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling UserPublicKeyApi->user_public_key_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**UserPublicKey**](UserPublicKey.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **user_public_key_list**
> list[UserPublicKey] user_public_key_list(authorization)

Get all user public keys.

Retrieves the public keys for all users.

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
api_instance = kps_api.UserPublicKeyApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Get all user public keys.
    api_response = api_instance.user_public_key_list(authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling UserPublicKeyApi->user_public_key_list: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**list[UserPublicKey]**](UserPublicKey.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **user_public_key_update**
> UpdateDocumentResponseV2 user_public_key_update(body, authorization)

Upsert current user public key.

Upserts the public key of the current user.

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
api_instance = kps_api.UserPublicKeyApi(kps_api.ApiClient(configuration))
body = kps_api.UserPublicKeyUpdatePayload() # UserPublicKeyUpdatePayload | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Upsert current user public key.
    api_response = api_instance.user_public_key_update(body, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling UserPublicKeyApi->user_public_key_update: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserPublicKeyUpdatePayload**](UserPublicKeyUpdatePayload.md)|  | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**UpdateDocumentResponseV2**](UpdateDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

