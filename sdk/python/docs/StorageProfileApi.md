# kps_api.StorageProfileApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**storage_profile_create**](StorageProfileApi.md#storage_profile_create) | **POST** /v1.0/servicedomains/{svcDomainId}/storageprofiles | Create a storage profile.
[**storage_profile_update**](StorageProfileApi.md#storage_profile_update) | **PUT** /v1.0/servicedomains/{svcDomainId}/storageprofiles/{id} | Update storage profile.
[**svc_domain_get_storage_profiles**](StorageProfileApi.md#svc_domain_get_storage_profiles) | **GET** /v1.0/servicedomains/{svcDomainId}/storageprofiles | Get storage profiles according to service domain ID.

# **storage_profile_create**
> CreateDocumentResponseV2 storage_profile_create(body, authorization, svc_domain_id)

Create a storage profile.

Create a storage profile on the given service domain ID {svcDomainId}.

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
api_instance = kps_api.StorageProfileApi(kps_api.ApiClient(configuration))
body = kps_api.StorageProfile() # StorageProfile | Description for the storage profile.
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
svc_domain_id = 'svc_domain_id_example' # str | ID for the service domain

try:
    # Create a storage profile.
    api_response = api_instance.storage_profile_create(body, authorization, svc_domain_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling StorageProfileApi->storage_profile_create: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**StorageProfile**](StorageProfile.md)| Description for the storage profile. | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **svc_domain_id** | **str**| ID for the service domain | 

### Return type

[**CreateDocumentResponseV2**](CreateDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **storage_profile_update**
> UpdateDocumentResponseV2 storage_profile_update(body, authorization, id, svc_domain_id)

Update storage profile.

Update the storage profile with {id} on the given service domain ID {svcDomainId}.

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
api_instance = kps_api.StorageProfileApi(kps_api.ApiClient(configuration))
body = kps_api.StorageProfile() # StorageProfile | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
id = 'id_example' # str | ID of the entity
svc_domain_id = 'svc_domain_id_example' # str | ID for the service domain

try:
    # Update storage profile.
    api_response = api_instance.storage_profile_update(body, authorization, id, svc_domain_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling StorageProfileApi->storage_profile_update: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**StorageProfile**](StorageProfile.md)|  | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **id** | **str**| ID of the entity | 
 **svc_domain_id** | **str**| ID for the service domain | 

### Return type

[**UpdateDocumentResponseV2**](UpdateDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **svc_domain_get_storage_profiles**
> StorageProfileListResponsePayload svc_domain_get_storage_profiles(svc_domain_id, authorization, order_by=order_by, filter=filter)

Get storage profiles according to service domain ID.

Retrieves all storage profiles for a service domain with a given ID {svcDomainId}

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
api_instance = kps_api.StorageProfileApi(kps_api.ApiClient(configuration))
svc_domain_id = 'svc_domain_id_example' # str | ID for the service domain
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Get storage profiles according to service domain ID.
    api_response = api_instance.svc_domain_get_storage_profiles(svc_domain_id, authorization, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling StorageProfileApi->svc_domain_get_storage_profiles: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svc_domain_id** | **str**| ID for the service domain | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 
 **order_by** | [**list[str]**](str.md)| Specify result order. Zero or more entries with format: &amp;ltkey&gt; [desc] where orderByKeys lists allowed keys in each response. | [optional] 
 **filter** | **str**| Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE &#x27;foo%&#x27;. Supported filter keys are the same as order by keys. | [optional] 

### Return type

[**StorageProfileListResponsePayload**](StorageProfileListResponsePayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

