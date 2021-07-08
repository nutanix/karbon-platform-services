# kps_api.AuthApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**login_call_v2**](AuthApi.md#login_call_v2) | **POST** /v1.0/login | Lets the user log in.
[**login_token_v1**](AuthApi.md#login_token_v1) | **POST** /v1.0/login/logintoken | Get a login token
[**short_login_token_v1**](AuthApi.md#short_login_token_v1) | **POST** /v1.0/login/shortlogintoken | Generate a short login token.

# **login_call_v2**
> LoginResponse login_call_v2(body)

Lets the user log in.

Lets the user log in.

### Example
```python
from __future__ import print_function
import time
import kps_api
from kps_api.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = kps_api.AuthApi()
body = kps_api.Credential() # Credential | This is a login credential

try:
    # Lets the user log in.
    api_response = api_instance.login_call_v2(body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling AuthApi->login_call_v2: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Credential**](Credential.md)| This is a login credential | 

### Return type

[**LoginResponse**](LoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **login_token_v1**
> LoginResponse login_token_v1(body, authorization)

Get a login token

Generates a login token equivalent to logging in.

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
api_instance = kps_api.AuthApi(kps_api.ApiClient(configuration))
body = NULL # dict(str, str) | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Get a login token
    api_response = api_instance.login_token_v1(body, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling AuthApi->login_token_v1: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**dict(str, str)**](dict.md)|  | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**LoginResponse**](LoginResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **short_login_token_v1**
> LoginResponse short_login_token_v1(authorization)

Generate a short login token.

Generates a temporary login token valid for a short duration.

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
api_instance = kps_api.AuthApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Generate a short login token.
    api_response = api_instance.short_login_token_v1(authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling AuthApi->short_login_token_v1: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**LoginResponse**](LoginResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

