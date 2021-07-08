# kps_api.KubernetesClusterApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**kubernetes_cluster_installer**](KubernetesClusterApi.md#kubernetes_cluster_installer) | **GET** /v1.0/kubernetescluster-installer | Get the kubernetes clusters helm installer.
[**kubernetes_clusters_create**](KubernetesClusterApi.md#kubernetes_clusters_create) | **POST** /v1.0/kubernetesclusters | Create a kubernetes cluster.
[**kubernetes_clusters_delete**](KubernetesClusterApi.md#kubernetes_clusters_delete) | **DELETE** /v1.0/kubernetesclusters/{id} | Delete a kubernetes cluster as specified by its ID.
[**kubernetes_clusters_get**](KubernetesClusterApi.md#kubernetes_clusters_get) | **GET** /v1.0/kubernetesclusters/{id} | Get single kubernetes cluster.
[**kubernetes_clusters_list**](KubernetesClusterApi.md#kubernetes_clusters_list) | **GET** /v1.0/kubernetesclusters | Get all kubernetes clusters.
[**kubernetes_clusters_update**](KubernetesClusterApi.md#kubernetes_clusters_update) | **PUT** /v1.0/kubernetesclusters/{id} | Update a kubernetes cluster by its ID.

# **kubernetes_cluster_installer**
> KubernetesClusterInstaller kubernetes_cluster_installer(authorization)

Get the kubernetes clusters helm installer.

 Gets the kubernetes cluster helm installer.

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
api_instance = kps_api.KubernetesClusterApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Get the kubernetes clusters helm installer.
    api_response = api_instance.kubernetes_cluster_installer(authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KubernetesClusterApi->kubernetes_cluster_installer: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**KubernetesClusterInstaller**](KubernetesClusterInstaller.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **kubernetes_clusters_create**
> CreateDocumentResponseV2 kubernetes_clusters_create(body, authorization)

Create a kubernetes cluster.

Create a kubernetes cluster.

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
api_instance = kps_api.KubernetesClusterApi(kps_api.ApiClient(configuration))
body = kps_api.KubernetesCluster() # KubernetesCluster | Describes the kubernetes cluster creation request.
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Create a kubernetes cluster.
    api_response = api_instance.kubernetes_clusters_create(body, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KubernetesClusterApi->kubernetes_clusters_create: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**KubernetesCluster**](KubernetesCluster.md)| Describes the kubernetes cluster creation request. | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**CreateDocumentResponseV2**](CreateDocumentResponseV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **kubernetes_clusters_delete**
> DeleteDocumentResponseV2 kubernetes_clusters_delete(id, authorization)

Delete a kubernetes cluster as specified by its ID.

Deletes a kubernetes cluster by its ID {id}.

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
api_instance = kps_api.KubernetesClusterApi(kps_api.ApiClient(configuration))
id = 'id_example' # str | ID of the entity
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Delete a kubernetes cluster as specified by its ID.
    api_response = api_instance.kubernetes_clusters_delete(id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KubernetesClusterApi->kubernetes_clusters_delete: %s\n" % e)
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

# **kubernetes_clusters_get**
> KubernetesCluster kubernetes_clusters_get(id, authorization)

Get single kubernetes cluster.

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
api_instance = kps_api.KubernetesClusterApi(kps_api.ApiClient(configuration))
id = 'id_example' # str | ID of the entity
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Get single kubernetes cluster.
    api_response = api_instance.kubernetes_clusters_get(id, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KubernetesClusterApi->kubernetes_clusters_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| ID of the entity | 
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**KubernetesCluster**](KubernetesCluster.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **kubernetes_clusters_list**
> KubernetesClustersListResponsePayload kubernetes_clusters_list(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)

Get all kubernetes clusters.

Retrieves a list of all kubernetes clusters.

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
api_instance = kps_api.KubernetesClusterApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
page_index = 789 # int | 0-based index of the page to fetch results. (optional)
page_size = 789 # int | Item count of each page. (optional)
order_by = ['order_by_example'] # list[str] | Specify result order. Zero or more entries with format: &ltkey> [desc] where orderByKeys lists allowed keys in each response. (optional)
filter = 'filter_example' # str | Specify result filter. Format is similar to a SQL WHERE clause. For example, to filter object by name with prefix foo, use: name LIKE 'foo%'. Supported filter keys are the same as order by keys. (optional)

try:
    # Get all kubernetes clusters.
    api_response = api_instance.kubernetes_clusters_list(authorization, page_index=page_index, page_size=page_size, order_by=order_by, filter=filter)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KubernetesClusterApi->kubernetes_clusters_list: %s\n" % e)
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

[**KubernetesClustersListResponsePayload**](KubernetesClustersListResponsePayload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **kubernetes_clusters_update**
> UpdateDocumentResponseV2 kubernetes_clusters_update(body, authorization, id)

Update a kubernetes cluster by its ID.

Updates a kubernetes cluster by its ID {id}.

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
api_instance = kps_api.KubernetesClusterApi(kps_api.ApiClient(configuration))
body = kps_api.KubernetesCluster() # KubernetesCluster | 
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.
id = 'id_example' # str | ID of the entity

try:
    # Update a kubernetes cluster by its ID.
    api_response = api_instance.kubernetes_clusters_update(body, authorization, id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KubernetesClusterApi->kubernetes_clusters_update: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**KubernetesCluster**](KubernetesCluster.md)|  | 
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

