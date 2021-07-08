# kps_api.ServiceDomainInfoApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**project_get_service_domains_info**](ServiceDomainInfoApi.md#project_get_service_domains_info) | **GET** /v1.0/projects/{projectId}/servicedomainsinfo | Get all service domain information for a project as specified by project ID.
[**service_domain_info_get**](ServiceDomainInfoApi.md#service_domain_info_get) | **GET** /v1.0/servicedomainsinfo/{svcDomainId} | Get all service domain information by service domain ID.
[**service_domain_info_list**](ServiceDomainInfoApi.md#service_domain_info_list) | **GET** /v1.0/servicedomainsinfo | Get service domain additional information like artifacts.
[**service_domain_info_update**](ServiceDomainInfoApi.md#service_domain_info_update) | **PUT** /v1.0/servicedomainsinfo/{svcDomainId} | Update service domain information by service domain ID.

# **project_get_service_domains_info**
> ServiceDomainInfoListPayload project_get_service_domains_info(project_id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)

Get all service domain information for a project as specified by project ID.

Retrieves all service domain information for a project as specified by project ID.

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
api_instance = kps_api.ServiceDomainInfoApi(kps_api.ApiClient(configuration))
project_id = 'project_id_example' # str | ID for the project
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Get all service domain information for a project as specified by project ID.
    api_response = api_instance.project_get_service_domains_info(project_id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceDomainInfoApi->project_get_service_domains_info: %s\n" % e)
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

[**ServiceDomainInfoListPayload**](ServiceDomainInfoListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_domain_info_get**
> ServiceDomainInfo service_domain_info_get(svc_domain_id, authorization)

Get all service domain information by service domain ID.

Retrieves all service domain additional information for a given service domain ID.

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
api_instance = kps_api.ServiceDomainInfoApi(kps_api.ApiClient(configuration))
svc_domain_id = 'svc_domain_id_example' # str | ID for the service domain
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Get all service domain information by service domain ID.
    api_response = api_instance.service_domain_info_get(svc_domain_id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceDomainInfoApi->service_domain_info_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svc_domain_id** | **str**| ID for the service domain | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**ServiceDomainInfo**](ServiceDomainInfo.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_domain_info_list**
> ServiceDomainInfoListPayload service_domain_info_list(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)

Get service domain additional information like artifacts.

Retrieves all service domain additional information.

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
api_instance = kps_api.ServiceDomainInfoApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Get service domain additional information like artifacts.
    api_response = api_instance.service_domain_info_list(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceDomainInfoApi->service_domain_info_list: %s\n" % e)
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

[**ServiceDomainInfoListPayload**](ServiceDomainInfoListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_domain_info_update**
> UpdateDocumentResponseV2 service_domain_info_update(body, authorization, svc_domain_id)

Update service domain information by service domain ID.

Update service domain additional information for a given service domain ID.

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
api_instance = kps_api.ServiceDomainInfoApi(kps_api.ApiClient(configuration))
body = kps_api.ServiceDomainInfo() # ServiceDomainInfo | Describes parameters used to create or update a ServiceDomainInfo
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
svc_domain_id = 'svc_domain_id_example' # str | ID for the service domain

try:
    # Update service domain information by service domain ID.
    api_response = api_instance.service_domain_info_update(body, authorization, svc_domain_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceDomainInfoApi->service_domain_info_update: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ServiceDomainInfo**](ServiceDomainInfo.md)| Describes parameters used to create or update a ServiceDomainInfo | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **svc_domain_id** | **str**| ID for the service domain | 

### Return type

[**UpdateDocumentResponseV2**](UpdateDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

