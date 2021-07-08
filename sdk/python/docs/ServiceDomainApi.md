# kps_api.ServiceDomainApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**project_get_service_domains**](ServiceDomainApi.md#project_get_service_domains) | **GET** /v1.0/projects/{projectId}/servicedomains | Get all service domains associated with a project by project ID.
[**service_domain_create**](ServiceDomainApi.md#service_domain_create) | **POST** /v1.0/servicedomains | Create service domain.
[**service_domain_delete**](ServiceDomainApi.md#service_domain_delete) | **DELETE** /v1.0/servicedomains/{svcDomainId} | Delete a service domain as specified by its ID.
[**service_domain_get**](ServiceDomainApi.md#service_domain_get) | **GET** /v1.0/servicedomains/{svcDomainId} | Get a service domain by its ID.
[**service_domain_get_effective_profile**](ServiceDomainApi.md#service_domain_get_effective_profile) | **GET** /v1.0/servicedomains/{svcDomainId}/effectiveprofile | Get a service domain effective profile by ID.
[**service_domain_get_nodes**](ServiceDomainApi.md#service_domain_get_nodes) | **GET** /v1.0/servicedomains/{svcDomainId}/nodes | Retrieves all nodes for a service domain by service domain ID {svcDomainId}.
[**service_domain_get_nodes_info**](ServiceDomainApi.md#service_domain_get_nodes_info) | **GET** /v1.0/servicedomains/{svcDomainId}/nodesinfo | Get nodes info for a service domain by service domain ID.
[**service_domain_list**](ServiceDomainApi.md#service_domain_list) | **GET** /v1.0/servicedomains | Get service domains.
[**service_domain_update**](ServiceDomainApi.md#service_domain_update) | **PUT** /v1.0/servicedomains/{svcDomainId} | Update a service domain by its ID.

# **project_get_service_domains**
> ServiceDomainListPayload project_get_service_domains(project_id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)

Get all service domains associated with a project by project ID.

Retrieves all service domains for a project by project ID {projectId}.

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
api_instance = kps_api.ServiceDomainApi(kps_api.ApiClient(configuration))
project_id = 'project_id_example' # str | ID for the project
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Get all service domains associated with a project by project ID.
    api_response = api_instance.project_get_service_domains(project_id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceDomainApi->project_get_service_domains: %s\n" % e)
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

[**ServiceDomainListPayload**](ServiceDomainListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_domain_create**
> CreateDocumentResponseV2 service_domain_create(body, authorization)

Create service domain.

Create a service domain.

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
api_instance = kps_api.ServiceDomainApi(kps_api.ApiClient(configuration))
body = kps_api.ServiceDomain() # ServiceDomain | Parameters and values used when creating a service domain
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Create service domain.
    api_response = api_instance.service_domain_create(body, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceDomainApi->service_domain_create: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ServiceDomain**](ServiceDomain.md)| Parameters and values used when creating a service domain | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**CreateDocumentResponseV2**](CreateDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_domain_delete**
> DeleteDocumentResponseV2 service_domain_delete(svc_domain_id, authorization)

Delete a service domain as specified by its ID.

Deletes the service domain with the given ID  {svcDomainId}.

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
api_instance = kps_api.ServiceDomainApi(kps_api.ApiClient(configuration))
svc_domain_id = 'svc_domain_id_example' # str | ID for the service domain
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Delete a service domain as specified by its ID.
    api_response = api_instance.service_domain_delete(svc_domain_id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceDomainApi->service_domain_delete: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svc_domain_id** | **str**| ID for the service domain | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**DeleteDocumentResponseV2**](DeleteDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_domain_get**
> ServiceDomain service_domain_get(svc_domain_id, authorization)

Get a service domain by its ID.

Retrieves the service domain with the given ID {svcDomainId}.

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
api_instance = kps_api.ServiceDomainApi(kps_api.ApiClient(configuration))
svc_domain_id = 'svc_domain_id_example' # str | ID for the service domain
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Get a service domain by its ID.
    api_response = api_instance.service_domain_get(svc_domain_id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceDomainApi->service_domain_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svc_domain_id** | **str**| ID for the service domain | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**ServiceDomain**](ServiceDomain.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_domain_get_effective_profile**
> ServiceDomainProfile service_domain_get_effective_profile(svc_domain_id, authorization)

Get a service domain effective profile by ID.

Retrieves the service domain effective profile with the given ID {svcDomainId}.

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
api_instance = kps_api.ServiceDomainApi(kps_api.ApiClient(configuration))
svc_domain_id = 'svc_domain_id_example' # str | ID for the service domain
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Get a service domain effective profile by ID.
    api_response = api_instance.service_domain_get_effective_profile(svc_domain_id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceDomainApi->service_domain_get_effective_profile: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svc_domain_id** | **str**| ID for the service domain | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**ServiceDomainProfile**](ServiceDomainProfile.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_domain_get_nodes**
> NodeListPayload service_domain_get_nodes(svc_domain_id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)

Retrieves all nodes for a service domain by service domain ID {svcDomainId}.

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
api_instance = kps_api.ServiceDomainApi(kps_api.ApiClient(configuration))
svc_domain_id = 'svc_domain_id_example' # str | ID for the service domain
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Retrieves all nodes for a service domain by service domain ID {svcDomainId}.
    api_response = api_instance.service_domain_get_nodes(svc_domain_id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceDomainApi->service_domain_get_nodes: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svc_domain_id** | **str**| ID for the service domain | 
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

# **service_domain_get_nodes_info**
> NodeInfoListPayload service_domain_get_nodes_info(svc_domain_id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)

Get nodes info for a service domain by service domain ID.

Retrieves all nodes info for a service domain by service domain ID {svcDomainId}.

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
api_instance = kps_api.ServiceDomainApi(kps_api.ApiClient(configuration))
svc_domain_id = 'svc_domain_id_example' # str | ID for the service domain
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Get nodes info for a service domain by service domain ID.
    api_response = api_instance.service_domain_get_nodes_info(svc_domain_id, authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceDomainApi->service_domain_get_nodes_info: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svc_domain_id** | **str**| ID for the service domain | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **page_index** | **int**| 0-based index of the page to fetch results. | [optional] 
 **page_size** | **int**| Item count of each page. | [optional] 
 **order_by** | [**list[str]**](str.md)| Specify result order. Zero or more entries with format: &amp;ltkey&gt; [desc] where orderByKeys lists allowed keys in each response. | [optional] 
 **filter** | **str**| Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE &#x27;foo%&#x27;. Supported filter keys are the same as order by keys. | [optional] 

### Return type

[**NodeInfoListPayload**](NodeInfoListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_domain_list**
> ServiceDomainListPayload service_domain_list(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)

Get service domains.

Retrieves all service domains associated with your account.

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
api_instance = kps_api.ServiceDomainApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Get service domains.
    api_response = api_instance.service_domain_list(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceDomainApi->service_domain_list: %s\n" % e)
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

[**ServiceDomainListPayload**](ServiceDomainListPayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_domain_update**
> UpdateDocumentResponseV2 service_domain_update(body, authorization, svc_domain_id)

Update a service domain by its ID.

Updates a service domain by its ID {svcDomainId}.

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
api_instance = kps_api.ServiceDomainApi(kps_api.ApiClient(configuration))
body = kps_api.ServiceDomain() # ServiceDomain | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
svc_domain_id = 'svc_domain_id_example' # str | ID for the service domain

try:
    # Update a service domain by its ID.
    api_response = api_instance.service_domain_update(body, authorization, svc_domain_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ServiceDomainApi->service_domain_update: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ServiceDomain**](ServiceDomain.md)|  | 
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

