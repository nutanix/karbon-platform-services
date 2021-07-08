# kps_api.SSHApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**setup_ssh_tunneling**](SSHApi.md#setup_ssh_tunneling) | **POST** /v1.0/setupsshtunneling | Configure SSH tunneling to the service domain.
[**teardown_ssh_tunneling**](SSHApi.md#teardown_ssh_tunneling) | **POST** /v1.0/teardownsshtunneling | Disable service domain SSH tunneling.

# **setup_ssh_tunneling**
> WstunPayload setup_ssh_tunneling(body, authorization)

Configure SSH tunneling to the service domain.

Configure SSH tunneling to the service domain. Requirements to use this feature: Minimum service domain version of 1.15.0. Remote SSH tunneling feature and CLI access are enabled per account. Service domain profile has SSH enabled.

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
api_instance = kps_api.SSHApi(kps_api.ApiClient(configuration))
body = kps_api.WstunRequest() # WstunRequest | SSH Tunneling setup request param
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Configure SSH tunneling to the service domain.
    api_response = api_instance.setup_ssh_tunneling(body, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SSHApi->setup_ssh_tunneling: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WstunRequest**](WstunRequest.md)| SSH Tunneling setup request param | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**WstunPayload**](WstunPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **teardown_ssh_tunneling**
> teardown_ssh_tunneling(body, authorization)

Disable service domain SSH tunneling.

Shut down SSH tunneling to the service domain. Disables SSH tunneling, including current open sessions.

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
api_instance = kps_api.SSHApi(kps_api.ApiClient(configuration))
body = kps_api.WstunTeardownRequest() # WstunTeardownRequest | SSH Tunneling teardown request param
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Disable service domain SSH tunneling.
    api_instance.teardown_ssh_tunneling(body, authorization)
except ApiException as e:
    print("Exception when calling SSHApi->teardown_ssh_tunneling: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WstunTeardownRequest**](WstunTeardownRequest.md)| SSH Tunneling teardown request param | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

void (empty response body)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

