# kps_api.HelmApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**helm_application_create**](HelmApi.md#helm_application_create) | **POST** /v1.0/helm/apps | Create Helm Application.
[**helm_application_update**](HelmApi.md#helm_application_update) | **PUT** /v1.0/helm/apps/{id} | Update Helm Application.
[**helm_template**](HelmApi.md#helm_template) | **POST** /v1.0/helm/template | Run Helm Template.

# **helm_application_create**
> CreateDocumentResponseV2 helm_application_create(chart, values, application, url, authorization)

Create Helm Application.

Create a Helm Chart based Application.  Use Application APIs to list or delete the application created by using this API

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
api_instance = kps_api.HelmApi(kps_api.ApiClient(configuration))
chart = 'chart_example' # file | 
values = 'values_example' # file | 
application = 'application_example' # str | 
url = 'url_example' # str | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Create Helm Application.
    api_response = api_instance.helm_application_create(chart, values, application, url, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling HelmApi->helm_application_create: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chart** | **file**|  | 
 **values** | **file**|  | 
 **application** | [**str**](.md)|  | 
 **url** | [**str**](.md)|  | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**CreateDocumentResponseV2**](CreateDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **helm_application_update**
> UpdateDocumentResponseV2 helm_application_update(chart, values, application, url, authorization, id)

Update Helm Application.

Update a Helm Chart based Application.

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
api_instance = kps_api.HelmApi(kps_api.ApiClient(configuration))
chart = 'chart_example' # file | 
values = 'values_example' # file | 
application = 'application_example' # str | 
url = 'url_example' # str | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
id = 'id_example' # str | ID of the entity

try:
    # Update Helm Application.
    api_response = api_instance.helm_application_update(chart, values, application, url, authorization, id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling HelmApi->helm_application_update: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chart** | **file**|  | 
 **values** | **file**|  | 
 **application** | [**str**](.md)|  | 
 **url** | [**str**](.md)|  | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **id** | **str**| ID of the entity | 

### Return type

[**UpdateDocumentResponseV2**](UpdateDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **helm_template**
> HelmTemplateResponse helm_template(chart, values, release, namespace, url, authorization)

Run Helm Template.

Run Helm Template to render Helm Chart.

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
api_instance = kps_api.HelmApi(kps_api.ApiClient(configuration))
chart = 'chart_example' # file | 
values = 'values_example' # file | 
release = 'release_example' # str | 
namespace = 'namespace_example' # str | 
url = 'url_example' # str | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Run Helm Template.
    api_response = api_instance.helm_template(chart, values, release, namespace, url, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling HelmApi->helm_template: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chart** | **file**|  | 
 **values** | **file**|  | 
 **release** | [**str**](.md)|  | 
 **namespace** | [**str**](.md)|  | 
 **url** | [**str**](.md)|  | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**HelmTemplateResponse**](HelmTemplateResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

