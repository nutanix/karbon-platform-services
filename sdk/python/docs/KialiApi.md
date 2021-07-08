# kps_api.KialiApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**app_dashboard**](KialiApi.md#app_dashboard) | **GET** /v1.0/kiali/namespaces/{namespace}/apps/{app}/dashboard | 
[**app_details**](KialiApi.md#app_details) | **GET** /v1.0/kiali/namespaces/{namespace}/apps/{app} | 
[**app_health**](KialiApi.md#app_health) | **GET** /v1.0/kiali/namespaces/{namespace}/apps/{app}/health | 
[**app_list**](KialiApi.md#app_list) | **GET** /v1.0/kiali/namespaces/{namespace}/apps | 
[**app_metrics**](KialiApi.md#app_metrics) | **GET** /v1.0/kiali/namespaces/{namespace}/apps/{app}/metrics | 
[**error_traces**](KialiApi.md#error_traces) | **GET** /v1.0/kiali/namespaces/{namespace}/services/{service}/errortraces | 
[**get_config**](KialiApi.md#get_config) | **GET** /v1.0/kiali/config | 
[**get_status**](KialiApi.md#get_status) | **GET** /v1.0/kiali/status | 
[**graph_app**](KialiApi.md#graph_app) | **GET** /v1.0/kiali/namespaces/{namespace}/applications/{app}/graph | 
[**graph_app_version**](KialiApi.md#graph_app_version) | **GET** /v1.0/kiali/namespaces/{namespace}/applications/{app}/versions/{version}/graph | 
[**graph_namespaces**](KialiApi.md#graph_namespaces) | **GET** /v1.0/kiali/namespaces/graph | The backing JSON for a namespaces graph.
[**graph_service**](KialiApi.md#graph_service) | **GET** /v1.0/kiali/namespaces/{namespace}/services/{service}/graph | The backing JSON for a service node detail graph.
[**graph_workload**](KialiApi.md#graph_workload) | **GET** /v1.0/kiali/namespaces/{namespace}/workloads/{workload}/graph | The backing JSON for a workload node detail graph.
[**istio_config_list**](KialiApi.md#istio_config_list) | **GET** /v1.0/kiali/namespaces/{namespace}/istio | 
[**istio_status**](KialiApi.md#istio_status) | **GET** /v1.0/kiali/istio/status | 
[**namespace_health**](KialiApi.md#namespace_health) | **GET** /v1.0/kiali/namespaces/{namespace}/health | 
[**namespace_list**](KialiApi.md#namespace_list) | **GET** /v1.0/kiali/namespaces | 
[**namespace_metrics**](KialiApi.md#namespace_metrics) | **GET** /v1.0/kiali/namespaces/{namespace}/metrics | 
[**namespace_tls**](KialiApi.md#namespace_tls) | **GET** /v1.0/kiali/namespaces/{namespace}/tls | 
[**namespace_validations**](KialiApi.md#namespace_validations) | **GET** /v1.0/kiali/namespaces/{namespace}/validations | 
[**pod_details**](KialiApi.md#pod_details) | **GET** /v1.0/kiali/namespaces/{namespace}/pods/{pod} | 
[**pod_logs**](KialiApi.md#pod_logs) | **GET** /v1.0/kiali/namespaces/{namespace}/pods/{pod}/logs | 
[**root**](KialiApi.md#root) | **GET** /v1.0/kiali | 
[**service_dashboard**](KialiApi.md#service_dashboard) | **GET** /v1.0/kiali/namespaces/{namespace}/services/{service}/dashboard | 
[**service_details**](KialiApi.md#service_details) | **GET** /v1.0/kiali/namespaces/{namespace}/services/{service} | 
[**service_health**](KialiApi.md#service_health) | **GET** /v1.0/kiali/namespaces/{namespace}/services/{service}/health | 
[**service_list**](KialiApi.md#service_list) | **GET** /v1.0/kiali/namespaces/{namespace}/services | 
[**service_metrics**](KialiApi.md#service_metrics) | **GET** /v1.0/kiali/namespaces/{namespace}/services/{service}/metrics | 
[**spans_list**](KialiApi.md#spans_list) | **GET** /v1.0/kiali/namespaces/{namespace}/services/{service}/spans | 
[**traces_detail**](KialiApi.md#traces_detail) | **GET** /v1.0/kiali/namespaces/{namespace}/services/{service}/traces | 
[**workload_dashboard**](KialiApi.md#workload_dashboard) | **GET** /v1.0/kiali/namespaces/{namespace}/workloads/{workload}/dashboard | 
[**workload_details**](KialiApi.md#workload_details) | **GET** /v1.0/kiali/namespaces/{namespace}/workloads/{workload} | 
[**workload_health**](KialiApi.md#workload_health) | **GET** /v1.0/kiali/namespaces/{namespace}/workloads/{workload}/health | 
[**workload_list**](KialiApi.md#workload_list) | **GET** /v1.0/kiali/namespaces/{namespace}/workloads | 
[**workload_metrics**](KialiApi.md#workload_metrics) | **GET** /v1.0/kiali/namespaces/{namespace}/workloads/{workload}/metrics | 

# **app_dashboard**
> MonitoringDashboard app_dashboard(app, namespace, service_domain, authorization, avg=avg, by_labels=by_labels, direction=direction, duration=duration, quantiles=quantiles, rate_func=rate_func, rate_interval=rate_interval, request_protocol=request_protocol, reporter=reporter, step=step)



Endpoint to fetch dashboard to be displayed, related to a single app

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
app = 'app_example' # str | The app name (label value).
namespace = 'namespace_example' # str | The namespace name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.
avg = true # bool | Flag for fetching histogram average. Default is true. (optional)
by_labels = ['by_labels_example'] # list[str] | List of labels to use for grouping metrics (via Prometheus 'by' clause). (optional)
direction = 'direction_example' # str | Traffic direction: 'inbound' or 'outbound'. (optional)
duration = 789 # int | Duration of the query period, in seconds. (optional)
quantiles = ['quantiles_example'] # list[str] | List of quantiles to fetch. Fetch no quantiles when empty. Ex: [0.5, 0.95, 0.99]. (optional)
rate_func = 'rate_func_example' # str | Prometheus function used to calculate rate: 'rate' or 'irate'. (optional)
rate_interval = 'rate_interval_example' # str | Interval used for rate and histogram calculation. (optional)
request_protocol = 'request_protocol_example' # str | Desired request protocol for the telemetry: For example, 'http' or 'grpc'. (optional)
reporter = 'reporter_example' # str | Istio telemetry reporter: 'source' or 'destination'. (optional)
step = 789 # int | Step between [graph] datapoints, in seconds. (optional)

try:
    api_response = api_instance.app_dashboard(app, namespace, service_domain, authorization, avg=avg, by_labels=by_labels, direction=direction, duration=duration, quantiles=quantiles, rate_func=rate_func, rate_interval=rate_interval, request_protocol=request_protocol, reporter=reporter, step=step)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->app_dashboard: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **str**| The app name (label value). | 
 **namespace** | **str**| The namespace name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 
 **avg** | **bool**| Flag for fetching histogram average. Default is true. | [optional] 
 **by_labels** | [**list[str]**](str.md)| List of labels to use for grouping metrics (via Prometheus &#x27;by&#x27; clause). | [optional] 
 **direction** | **str**| Traffic direction: &#x27;inbound&#x27; or &#x27;outbound&#x27;. | [optional] 
 **duration** | **int**| Duration of the query period, in seconds. | [optional] 
 **quantiles** | [**list[str]**](str.md)| List of quantiles to fetch. Fetch no quantiles when empty. Ex: [0.5, 0.95, 0.99]. | [optional] 
 **rate_func** | **str**| Prometheus function used to calculate rate: &#x27;rate&#x27; or &#x27;irate&#x27;. | [optional] 
 **rate_interval** | **str**| Interval used for rate and histogram calculation. | [optional] 
 **request_protocol** | **str**| Desired request protocol for the telemetry: For example, &#x27;http&#x27; or &#x27;grpc&#x27;. | [optional] 
 **reporter** | **str**| Istio telemetry reporter: &#x27;source&#x27; or &#x27;destination&#x27;. | [optional] 
 **step** | **int**| Step between [graph] datapoints, in seconds. | [optional] 

### Return type

[**MonitoringDashboard**](MonitoringDashboard.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **app_details**
> App app_details(app, namespace, service_domain, authorization)



Endpoint to get the app details

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
app = 'app_example' # str | The app name (label value).
namespace = 'namespace_example' # str | The namespace name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.app_details(app, namespace, service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->app_details: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **str**| The app name (label value). | 
 **namespace** | **str**| The namespace name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**App**](App.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **app_health**
> AppHealth app_health(app, namespace, service_domain, authorization)



Get health associated to the given app

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
app = 'app_example' # str | The app name (label value).
namespace = 'namespace_example' # str | The namespace name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.app_health(app, namespace, service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->app_health: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **str**| The app name (label value). | 
 **namespace** | **str**| The namespace name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**AppHealth**](AppHealth.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **app_list**
> AppList app_list(namespace, service_domain, authorization)



Endpoint to get the list of apps for a namespace

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.app_list(namespace, service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->app_list: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**AppList**](AppList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **app_metrics**
> KialiMetrics app_metrics(app, namespace, service_domain, authorization, avg=avg, by_labels=by_labels, direction=direction, duration=duration, filters=filters, quantiles=quantiles, rate_func=rate_func, rate_interval=rate_interval, request_protocol=request_protocol, reporter=reporter, step=step, version=version)



Endpoint to fetch metrics to be displayed, related to a single app

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
app = 'app_example' # str | The app name (label value).
namespace = 'namespace_example' # str | The namespace name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.
avg = true # bool | Flag for fetching histogram average. Default is true. (optional)
by_labels = ['by_labels_example'] # list[str] | List of labels to use for grouping metrics (via Prometheus 'by' clause). (optional)
direction = 'direction_example' # str | Traffic direction: 'inbound' or 'outbound'. (optional)
duration = 789 # int | Duration of the query period, in seconds. (optional)
filters = ['filters_example'] # list[str] | List of metrics to fetch. Fetch all metrics when empty. List entries are Kiali internal metric names. (optional)
quantiles = ['quantiles_example'] # list[str] | List of quantiles to fetch. Fetch no quantiles when empty. Ex: [0.5, 0.95, 0.99]. (optional)
rate_func = 'rate_func_example' # str | Prometheus function used to calculate rate: 'rate' or 'irate'. (optional)
rate_interval = 'rate_interval_example' # str | Interval used for rate and histogram calculation. (optional)
request_protocol = 'request_protocol_example' # str | Desired request protocol for the telemetry: For example, 'http' or 'grpc'. (optional)
reporter = 'reporter_example' # str | Istio telemetry reporter: 'source' or 'destination'. (optional)
step = 789 # int | Step between [graph] datapoints, in seconds. (optional)
version = 'version_example' # str | Filters metrics by the specified version. (optional)

try:
    api_response = api_instance.app_metrics(app, namespace, service_domain, authorization, avg=avg, by_labels=by_labels, direction=direction, duration=duration, filters=filters, quantiles=quantiles, rate_func=rate_func, rate_interval=rate_interval, request_protocol=request_protocol, reporter=reporter, step=step, version=version)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->app_metrics: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **str**| The app name (label value). | 
 **namespace** | **str**| The namespace name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 
 **avg** | **bool**| Flag for fetching histogram average. Default is true. | [optional] 
 **by_labels** | [**list[str]**](str.md)| List of labels to use for grouping metrics (via Prometheus &#x27;by&#x27; clause). | [optional] 
 **direction** | **str**| Traffic direction: &#x27;inbound&#x27; or &#x27;outbound&#x27;. | [optional] 
 **duration** | **int**| Duration of the query period, in seconds. | [optional] 
 **filters** | [**list[str]**](str.md)| List of metrics to fetch. Fetch all metrics when empty. List entries are Kiali internal metric names. | [optional] 
 **quantiles** | [**list[str]**](str.md)| List of quantiles to fetch. Fetch no quantiles when empty. Ex: [0.5, 0.95, 0.99]. | [optional] 
 **rate_func** | **str**| Prometheus function used to calculate rate: &#x27;rate&#x27; or &#x27;irate&#x27;. | [optional] 
 **rate_interval** | **str**| Interval used for rate and histogram calculation. | [optional] 
 **request_protocol** | **str**| Desired request protocol for the telemetry: For example, &#x27;http&#x27; or &#x27;grpc&#x27;. | [optional] 
 **reporter** | **str**| Istio telemetry reporter: &#x27;source&#x27; or &#x27;destination&#x27;. | [optional] 
 **step** | **int**| Step between [graph] datapoints, in seconds. | [optional] 
 **version** | **str**| Filters metrics by the specified version. | [optional] 

### Return type

[**KialiMetrics**](KialiMetrics.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **error_traces**
> error_traces(namespace, service, service_domain, authorization)



Endpoint to get the number of traces in error for a given service

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
service = 'service_example' # str | The service name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_instance.error_traces(namespace, service, service_domain, authorization)
except ApiException as e:
    print("Exception when calling KialiApi->error_traces: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **service** | **str**| The service name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

void (empty response body)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_config**
> StatusInfo get_config(service_domain, authorization)



Endpoint to get the config of Kiali

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.get_config(service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->get_config: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**StatusInfo**](StatusInfo.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_status**
> StatusInfo get_status(service_domain, authorization)



Endpoint to get the status of Kiali

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.get_status(service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->get_status: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**StatusInfo**](StatusInfo.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **graph_app**
> GraphConfig graph_app(app, namespace, service_domain, authorization, appenders=appenders, duration=duration, graph_type=graph_type, group_by=group_by, inject_service_nodes=inject_service_nodes, query_time=query_time)



The backing JSON for an app node detail graph. (supported graphTypes: app | versionedApp)

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
app = 'app_example' # str | The app name (label value).
namespace = 'namespace_example' # str | The namespace name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.
appenders = 'appenders_example' # str | Comma-separated list of Appenders to run. Available appenders: [deadNode, istio, responseTime, securityPolicy, serviceEntry, sidecarsCheck, unusedNode]. (optional)
duration = 'duration_example' # str | Query time-range duration (Golang string duration). (optional)
graph_type = 'graph_type_example' # str | Graph type. Available graph types: [app, service, versionedApp, workload]. (optional)
group_by = 'group_by_example' # str | App box grouping characteristic. Available groupings: [app, none, version]. (optional)
inject_service_nodes = 'inject_service_nodes_example' # str | Flag for injecting the requested service node between source and destination nodes. (optional)
query_time = 'query_time_example' # str | Unix time (seconds) for query such that time range is [queryTime-duration..queryTime]. Default is now. (optional)

try:
    api_response = api_instance.graph_app(app, namespace, service_domain, authorization, appenders=appenders, duration=duration, graph_type=graph_type, group_by=group_by, inject_service_nodes=inject_service_nodes, query_time=query_time)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->graph_app: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **str**| The app name (label value). | 
 **namespace** | **str**| The namespace name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 
 **appenders** | **str**| Comma-separated list of Appenders to run. Available appenders: [deadNode, istio, responseTime, securityPolicy, serviceEntry, sidecarsCheck, unusedNode]. | [optional] 
 **duration** | **str**| Query time-range duration (Golang string duration). | [optional] 
 **graph_type** | **str**| Graph type. Available graph types: [app, service, versionedApp, workload]. | [optional] 
 **group_by** | **str**| App box grouping characteristic. Available groupings: [app, none, version]. | [optional] 
 **inject_service_nodes** | **str**| Flag for injecting the requested service node between source and destination nodes. | [optional] 
 **query_time** | **str**| Unix time (seconds) for query such that time range is [queryTime-duration..queryTime]. Default is now. | [optional] 

### Return type

[**GraphConfig**](GraphConfig.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **graph_app_version**
> GraphConfig graph_app_version(app, version, namespace, service_domain, authorization, appenders=appenders, duration=duration, graph_type=graph_type, group_by=group_by, inject_service_nodes=inject_service_nodes, query_time=query_time)



The backing JSON for a versioned app node detail graph. (supported graphTypes: app | versionedApp)

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
app = 'app_example' # str | The app name (label value).
version = 'version_example' # str | The app version (label value).
namespace = 'namespace_example' # str | The namespace name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.
appenders = 'appenders_example' # str | Comma-separated list of Appenders to run. Available appenders: [deadNode, istio, responseTime, securityPolicy, serviceEntry, sidecarsCheck, unusedNode]. (optional)
duration = 'duration_example' # str | Query time-range duration (Golang string duration). (optional)
graph_type = 'graph_type_example' # str | Graph type. Available graph types: [app, service, versionedApp, workload]. (optional)
group_by = 'group_by_example' # str | App box grouping characteristic. Available groupings: [app, none, version]. (optional)
inject_service_nodes = 'inject_service_nodes_example' # str | Flag for injecting the requested service node between source and destination nodes. (optional)
query_time = 'query_time_example' # str | Unix time (seconds) for query such that time range is [queryTime-duration..queryTime]. Default is now. (optional)

try:
    api_response = api_instance.graph_app_version(app, version, namespace, service_domain, authorization, appenders=appenders, duration=duration, graph_type=graph_type, group_by=group_by, inject_service_nodes=inject_service_nodes, query_time=query_time)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->graph_app_version: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **str**| The app name (label value). | 
 **version** | **str**| The app version (label value). | 
 **namespace** | **str**| The namespace name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 
 **appenders** | **str**| Comma-separated list of Appenders to run. Available appenders: [deadNode, istio, responseTime, securityPolicy, serviceEntry, sidecarsCheck, unusedNode]. | [optional] 
 **duration** | **str**| Query time-range duration (Golang string duration). | [optional] 
 **graph_type** | **str**| Graph type. Available graph types: [app, service, versionedApp, workload]. | [optional] 
 **group_by** | **str**| App box grouping characteristic. Available groupings: [app, none, version]. | [optional] 
 **inject_service_nodes** | **str**| Flag for injecting the requested service node between source and destination nodes. | [optional] 
 **query_time** | **str**| Unix time (seconds) for query such that time range is [queryTime-duration..queryTime]. Default is now. | [optional] 

### Return type

[**GraphConfig**](GraphConfig.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **graph_namespaces**
> GraphConfig graph_namespaces(namespaces, service_domain, authorization, appenders=appenders, duration=duration, graph_type=graph_type, group_by=group_by, inject_service_nodes=inject_service_nodes, query_time=query_time)

The backing JSON for a namespaces graph.

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespaces = 'namespaces_example' # str | Comma-separated list of namespaces to include in the graph. The namespaces must be accessible to the client.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.
appenders = 'appenders_example' # str | Comma-separated list of Appenders to run. Available appenders: [deadNode, istio, responseTime, securityPolicy, serviceEntry, sidecarsCheck, unusedNode]. (optional)
duration = 'duration_example' # str | Query time-range duration (Golang string duration). (optional)
graph_type = 'graph_type_example' # str | Graph type. Available graph types: [app, service, versionedApp, workload]. (optional)
group_by = 'group_by_example' # str | App box grouping characteristic. Available groupings: [app, none, version]. (optional)
inject_service_nodes = 'inject_service_nodes_example' # str | Flag for injecting the requested service node between source and destination nodes. (optional)
query_time = 'query_time_example' # str | Unix time (seconds) for query such that time range is [queryTime-duration..queryTime]. Default is now. (optional)

try:
    # The backing JSON for a namespaces graph.
    api_response = api_instance.graph_namespaces(namespaces, service_domain, authorization, appenders=appenders, duration=duration, graph_type=graph_type, group_by=group_by, inject_service_nodes=inject_service_nodes, query_time=query_time)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->graph_namespaces: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespaces** | **str**| Comma-separated list of namespaces to include in the graph. The namespaces must be accessible to the client. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 
 **appenders** | **str**| Comma-separated list of Appenders to run. Available appenders: [deadNode, istio, responseTime, securityPolicy, serviceEntry, sidecarsCheck, unusedNode]. | [optional] 
 **duration** | **str**| Query time-range duration (Golang string duration). | [optional] 
 **graph_type** | **str**| Graph type. Available graph types: [app, service, versionedApp, workload]. | [optional] 
 **group_by** | **str**| App box grouping characteristic. Available groupings: [app, none, version]. | [optional] 
 **inject_service_nodes** | **str**| Flag for injecting the requested service node between source and destination nodes. | [optional] 
 **query_time** | **str**| Unix time (seconds) for query such that time range is [queryTime-duration..queryTime]. Default is now. | [optional] 

### Return type

[**GraphConfig**](GraphConfig.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **graph_service**
> GraphConfig graph_service(namespace, service, service_domain, authorization, appenders=appenders, duration=duration, graph_type=graph_type, group_by=group_by, query_time=query_time)

The backing JSON for a service node detail graph.

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
service = 'service_example' # str | The service name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.
appenders = 'appenders_example' # str | Comma-separated list of Appenders to run. Available appenders: [deadNode, istio, responseTime, securityPolicy, serviceEntry, sidecarsCheck, unusedNode]. (optional)
duration = 'duration_example' # str | Query time-range duration (Golang string duration). (optional)
graph_type = 'graph_type_example' # str | Graph type. Available graph types: [app, service, versionedApp, workload]. (optional)
group_by = 'group_by_example' # str | App box grouping characteristic. Available groupings: [app, none, version]. (optional)
query_time = 'query_time_example' # str | Unix time (seconds) for query such that time range is [queryTime-duration..queryTime]. Default is now. (optional)

try:
    # The backing JSON for a service node detail graph.
    api_response = api_instance.graph_service(namespace, service, service_domain, authorization, appenders=appenders, duration=duration, graph_type=graph_type, group_by=group_by, query_time=query_time)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->graph_service: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **service** | **str**| The service name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 
 **appenders** | **str**| Comma-separated list of Appenders to run. Available appenders: [deadNode, istio, responseTime, securityPolicy, serviceEntry, sidecarsCheck, unusedNode]. | [optional] 
 **duration** | **str**| Query time-range duration (Golang string duration). | [optional] 
 **graph_type** | **str**| Graph type. Available graph types: [app, service, versionedApp, workload]. | [optional] 
 **group_by** | **str**| App box grouping characteristic. Available groupings: [app, none, version]. | [optional] 
 **query_time** | **str**| Unix time (seconds) for query such that time range is [queryTime-duration..queryTime]. Default is now. | [optional] 

### Return type

[**GraphConfig**](GraphConfig.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **graph_workload**
> GraphConfig graph_workload(namespace, workload, service_domain, authorization, appenders=appenders, duration=duration, graph_type=graph_type, group_by=group_by, inject_service_nodes=inject_service_nodes, query_time=query_time)

The backing JSON for a workload node detail graph.

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
workload = 'workload_example' # str | The workload name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.
appenders = 'appenders_example' # str | Comma-separated list of Appenders to run. Available appenders: [deadNode, istio, responseTime, securityPolicy, serviceEntry, sidecarsCheck, unusedNode]. (optional)
duration = 'duration_example' # str | Query time-range duration (Golang string duration). (optional)
graph_type = 'graph_type_example' # str | Graph type. Available graph types: [app, service, versionedApp, workload]. (optional)
group_by = 'group_by_example' # str | App box grouping characteristic. Available groupings: [app, none, version]. (optional)
inject_service_nodes = 'inject_service_nodes_example' # str | Flag for injecting the requested service node between source and destination nodes. (optional)
query_time = 'query_time_example' # str | Unix time (seconds) for query such that time range is [queryTime-duration..queryTime]. Default is now. (optional)

try:
    # The backing JSON for a workload node detail graph.
    api_response = api_instance.graph_workload(namespace, workload, service_domain, authorization, appenders=appenders, duration=duration, graph_type=graph_type, group_by=group_by, inject_service_nodes=inject_service_nodes, query_time=query_time)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->graph_workload: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **workload** | **str**| The workload name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 
 **appenders** | **str**| Comma-separated list of Appenders to run. Available appenders: [deadNode, istio, responseTime, securityPolicy, serviceEntry, sidecarsCheck, unusedNode]. | [optional] 
 **duration** | **str**| Query time-range duration (Golang string duration). | [optional] 
 **graph_type** | **str**| Graph type. Available graph types: [app, service, versionedApp, workload]. | [optional] 
 **group_by** | **str**| App box grouping characteristic. Available groupings: [app, none, version]. | [optional] 
 **inject_service_nodes** | **str**| Flag for injecting the requested service node between source and destination nodes. | [optional] 
 **query_time** | **str**| Unix time (seconds) for query such that time range is [queryTime-duration..queryTime]. Default is now. | [optional] 

### Return type

[**GraphConfig**](GraphConfig.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **istio_config_list**
> IstioConfigList istio_config_list(namespace, service_domain, authorization)



Endpoint to get the list of Istio Config of a namespace

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.istio_config_list(namespace, service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->istio_config_list: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**IstioConfigList**](IstioConfigList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **istio_status**
> IstioComponentStatus istio_status(service_domain, authorization)



Get the status of each components needed in the control plane

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.istio_status(service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->istio_status: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**IstioComponentStatus**](IstioComponentStatus.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **namespace_health**
> NamespaceAppHealth namespace_health(namespace, service_domain, authorization)



Get health for all objects in the given namespace

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.namespace_health(namespace, service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->namespace_health: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**NamespaceAppHealth**](NamespaceAppHealth.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **namespace_list**
> list[Namespace] namespace_list(service_domain, authorization)



Endpoint to get the list of the available namespaces

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.namespace_list(service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->namespace_list: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**list[Namespace]**](Namespace.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **namespace_metrics**
> KialiMetrics namespace_metrics(namespace, service_domain, authorization)



Endpoint to fetch metrics to be displayed, related to a namespace

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.namespace_metrics(namespace, service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->namespace_metrics: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**KialiMetrics**](KialiMetrics.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **namespace_tls**
> MTLSStatus namespace_tls(namespace, service_domain, authorization)



Get TLS status for the given namespace

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.namespace_tls(namespace, service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->namespace_tls: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**MTLSStatus**](MTLSStatus.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **namespace_validations**
> IstioValidationSummary namespace_validations(namespace, service_domain, authorization)



Get validation summary for all objects in the given namespace

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.namespace_validations(namespace, service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->namespace_validations: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**IstioValidationSummary**](IstioValidationSummary.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **pod_details**
> Workload pod_details(namespace, pod, service_domain, authorization)



Endpoint to get pod details

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
pod = 'pod_example' # str | The pod name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.pod_details(namespace, pod, service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->pod_details: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **pod** | **str**| The pod name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**Workload**](Workload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **pod_logs**
> Workload pod_logs(namespace, pod, service_domain, authorization, container=container, since_time=since_time)



Endpoint to get pod logs

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
pod = 'pod_example' # str | The pod name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.
container = 'container_example' # str | The pod container name. Optional for single-container pod. Otherwise required. (optional)
since_time = 'since_time_example' # str | The start time for fetching logs. UNIX time in seconds. Default is all logs. (optional)

try:
    api_response = api_instance.pod_logs(namespace, pod, service_domain, authorization, container=container, since_time=since_time)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->pod_logs: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **pod** | **str**| The pod name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 
 **container** | **str**| The pod container name. Optional for single-container pod. Otherwise required. | [optional] 
 **since_time** | **str**| The start time for fetching logs. UNIX time in seconds. Default is all logs. | [optional] 

### Return type

[**Workload**](Workload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **root**
> StatusInfo root(service_domain, authorization)



Endpoint to get the status of Kiali

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.root(service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->root: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**StatusInfo**](StatusInfo.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_dashboard**
> MonitoringDashboard service_dashboard(namespace, service, service_domain, authorization, avg=avg, by_labels=by_labels, direction=direction, duration=duration, quantiles=quantiles, rate_func=rate_func, rate_interval=rate_interval, request_protocol=request_protocol, reporter=reporter, step=step)



Endpoint to fetch dashboard to be displayed, related to a single service

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
service = 'service_example' # str | The service name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.
avg = true # bool | Flag for fetching histogram average. Default is true. (optional)
by_labels = ['by_labels_example'] # list[str] | List of labels to use for grouping metrics (via Prometheus 'by' clause). (optional)
direction = 'direction_example' # str | Traffic direction: 'inbound' or 'outbound'. (optional)
duration = 789 # int | Duration of the query period, in seconds. (optional)
quantiles = ['quantiles_example'] # list[str] | List of quantiles to fetch. Fetch no quantiles when empty. Ex: [0.5, 0.95, 0.99]. (optional)
rate_func = 'rate_func_example' # str | Prometheus function used to calculate rate: 'rate' or 'irate'. (optional)
rate_interval = 'rate_interval_example' # str | Interval used for rate and histogram calculation. (optional)
request_protocol = 'request_protocol_example' # str | Desired request protocol for the telemetry: For example, 'http' or 'grpc'. (optional)
reporter = 'reporter_example' # str | Istio telemetry reporter: 'source' or 'destination'. (optional)
step = 789 # int | Step between [graph] datapoints, in seconds. (optional)

try:
    api_response = api_instance.service_dashboard(namespace, service, service_domain, authorization, avg=avg, by_labels=by_labels, direction=direction, duration=duration, quantiles=quantiles, rate_func=rate_func, rate_interval=rate_interval, request_protocol=request_protocol, reporter=reporter, step=step)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->service_dashboard: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **service** | **str**| The service name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 
 **avg** | **bool**| Flag for fetching histogram average. Default is true. | [optional] 
 **by_labels** | [**list[str]**](str.md)| List of labels to use for grouping metrics (via Prometheus &#x27;by&#x27; clause). | [optional] 
 **direction** | **str**| Traffic direction: &#x27;inbound&#x27; or &#x27;outbound&#x27;. | [optional] 
 **duration** | **int**| Duration of the query period, in seconds. | [optional] 
 **quantiles** | [**list[str]**](str.md)| List of quantiles to fetch. Fetch no quantiles when empty. Ex: [0.5, 0.95, 0.99]. | [optional] 
 **rate_func** | **str**| Prometheus function used to calculate rate: &#x27;rate&#x27; or &#x27;irate&#x27;. | [optional] 
 **rate_interval** | **str**| Interval used for rate and histogram calculation. | [optional] 
 **request_protocol** | **str**| Desired request protocol for the telemetry: For example, &#x27;http&#x27; or &#x27;grpc&#x27;. | [optional] 
 **reporter** | **str**| Istio telemetry reporter: &#x27;source&#x27; or &#x27;destination&#x27;. | [optional] 
 **step** | **int**| Step between [graph] datapoints, in seconds. | [optional] 

### Return type

[**MonitoringDashboard**](MonitoringDashboard.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_details**
> ServiceDetails service_details(namespace, service, service_domain, authorization)



Endpoint to get the details of a given service

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
service = 'service_example' # str | The service name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.service_details(namespace, service, service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->service_details: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **service** | **str**| The service name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**ServiceDetails**](ServiceDetails.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_health**
> ServiceHealth service_health(namespace, service, service_domain, authorization)



Get health associated to the given service

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
service = 'service_example' # str | The service name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.service_health(namespace, service, service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->service_health: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **service** | **str**| The service name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**ServiceHealth**](ServiceHealth.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_list**
> ServiceList service_list(namespace, service_domain, authorization)



Endpoint to get the details of a given service

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.service_list(namespace, service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->service_list: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**ServiceList**](ServiceList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **service_metrics**
> KialiMetrics service_metrics(namespace, service, service_domain, authorization, avg=avg, by_labels=by_labels, direction=direction, duration=duration, filters=filters, quantiles=quantiles, rate_func=rate_func, rate_interval=rate_interval, request_protocol=request_protocol, reporter=reporter, step=step, version=version)



Endpoint to fetch metrics to be displayed, related to a single service

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
service = 'service_example' # str | The service name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.
avg = true # bool | Flag for fetching histogram average. Default is true. (optional)
by_labels = ['by_labels_example'] # list[str] | List of labels to use for grouping metrics (via Prometheus 'by' clause). (optional)
direction = 'direction_example' # str | Traffic direction: 'inbound' or 'outbound'. (optional)
duration = 789 # int | Duration of the query period, in seconds. (optional)
filters = ['filters_example'] # list[str] | List of metrics to fetch. Fetch all metrics when empty. List entries are Kiali internal metric names. (optional)
quantiles = ['quantiles_example'] # list[str] | List of quantiles to fetch. Fetch no quantiles when empty. Ex: [0.5, 0.95, 0.99]. (optional)
rate_func = 'rate_func_example' # str | Prometheus function used to calculate rate: 'rate' or 'irate'. (optional)
rate_interval = 'rate_interval_example' # str | Interval used for rate and histogram calculation. (optional)
request_protocol = 'request_protocol_example' # str | Desired request protocol for the telemetry: For example, 'http' or 'grpc'. (optional)
reporter = 'reporter_example' # str | Istio telemetry reporter: 'source' or 'destination'. (optional)
step = 789 # int | Step between [graph] datapoints, in seconds. (optional)
version = 'version_example' # str | Filters metrics by the specified version. (optional)

try:
    api_response = api_instance.service_metrics(namespace, service, service_domain, authorization, avg=avg, by_labels=by_labels, direction=direction, duration=duration, filters=filters, quantiles=quantiles, rate_func=rate_func, rate_interval=rate_interval, request_protocol=request_protocol, reporter=reporter, step=step, version=version)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->service_metrics: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **service** | **str**| The service name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 
 **avg** | **bool**| Flag for fetching histogram average. Default is true. | [optional] 
 **by_labels** | [**list[str]**](str.md)| List of labels to use for grouping metrics (via Prometheus &#x27;by&#x27; clause). | [optional] 
 **direction** | **str**| Traffic direction: &#x27;inbound&#x27; or &#x27;outbound&#x27;. | [optional] 
 **duration** | **int**| Duration of the query period, in seconds. | [optional] 
 **filters** | [**list[str]**](str.md)| List of metrics to fetch. Fetch all metrics when empty. List entries are Kiali internal metric names. | [optional] 
 **quantiles** | [**list[str]**](str.md)| List of quantiles to fetch. Fetch no quantiles when empty. Ex: [0.5, 0.95, 0.99]. | [optional] 
 **rate_func** | **str**| Prometheus function used to calculate rate: &#x27;rate&#x27; or &#x27;irate&#x27;. | [optional] 
 **rate_interval** | **str**| Interval used for rate and histogram calculation. | [optional] 
 **request_protocol** | **str**| Desired request protocol for the telemetry: For example, &#x27;http&#x27; or &#x27;grpc&#x27;. | [optional] 
 **reporter** | **str**| Istio telemetry reporter: &#x27;source&#x27; or &#x27;destination&#x27;. | [optional] 
 **step** | **int**| Step between [graph] datapoints, in seconds. | [optional] 
 **version** | **str**| Filters metrics by the specified version. | [optional] 

### Return type

[**KialiMetrics**](KialiMetrics.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **spans_list**
> list[Span] spans_list(namespace, service, service_domain, authorization)



Endpoint to get Jaeger spans for a given service

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
service = 'service_example' # str | The service name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.spans_list(namespace, service, service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->spans_list: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **service** | **str**| The service name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**list[Span]**](Span.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **traces_detail**
> list[Trace] traces_detail(namespace, service, service_domain, authorization)



Endpoint to get a specific trace of a given service

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
service = 'service_example' # str | The service name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.traces_detail(namespace, service, service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->traces_detail: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **service** | **str**| The service name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**list[Trace]**](Trace.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **workload_dashboard**
> MonitoringDashboard workload_dashboard(namespace, workload, service_domain, authorization, avg=avg, by_labels=by_labels, direction=direction, duration=duration, quantiles=quantiles, rate_func=rate_func, rate_interval=rate_interval, request_protocol=request_protocol, reporter=reporter, step=step)



Endpoint to fetch dashboard to be displayed, related to a single workload

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
workload = 'workload_example' # str | The workload name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.
avg = true # bool | Flag for fetching histogram average. Default is true. (optional)
by_labels = ['by_labels_example'] # list[str] | List of labels to use for grouping metrics (via Prometheus 'by' clause). (optional)
direction = 'direction_example' # str | Traffic direction: 'inbound' or 'outbound'. (optional)
duration = 789 # int | Duration of the query period, in seconds. (optional)
quantiles = ['quantiles_example'] # list[str] | List of quantiles to fetch. Fetch no quantiles when empty. Ex: [0.5, 0.95, 0.99]. (optional)
rate_func = 'rate_func_example' # str | Prometheus function used to calculate rate: 'rate' or 'irate'. (optional)
rate_interval = 'rate_interval_example' # str | Interval used for rate and histogram calculation. (optional)
request_protocol = 'request_protocol_example' # str | Desired request protocol for the telemetry: For example, 'http' or 'grpc'. (optional)
reporter = 'reporter_example' # str | Istio telemetry reporter: 'source' or 'destination'. (optional)
step = 789 # int | Step between [graph] datapoints, in seconds. (optional)

try:
    api_response = api_instance.workload_dashboard(namespace, workload, service_domain, authorization, avg=avg, by_labels=by_labels, direction=direction, duration=duration, quantiles=quantiles, rate_func=rate_func, rate_interval=rate_interval, request_protocol=request_protocol, reporter=reporter, step=step)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->workload_dashboard: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **workload** | **str**| The workload name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 
 **avg** | **bool**| Flag for fetching histogram average. Default is true. | [optional] 
 **by_labels** | [**list[str]**](str.md)| List of labels to use for grouping metrics (via Prometheus &#x27;by&#x27; clause). | [optional] 
 **direction** | **str**| Traffic direction: &#x27;inbound&#x27; or &#x27;outbound&#x27;. | [optional] 
 **duration** | **int**| Duration of the query period, in seconds. | [optional] 
 **quantiles** | [**list[str]**](str.md)| List of quantiles to fetch. Fetch no quantiles when empty. Ex: [0.5, 0.95, 0.99]. | [optional] 
 **rate_func** | **str**| Prometheus function used to calculate rate: &#x27;rate&#x27; or &#x27;irate&#x27;. | [optional] 
 **rate_interval** | **str**| Interval used for rate and histogram calculation. | [optional] 
 **request_protocol** | **str**| Desired request protocol for the telemetry: For example, &#x27;http&#x27; or &#x27;grpc&#x27;. | [optional] 
 **reporter** | **str**| Istio telemetry reporter: &#x27;source&#x27; or &#x27;destination&#x27;. | [optional] 
 **step** | **int**| Step between [graph] datapoints, in seconds. | [optional] 

### Return type

[**MonitoringDashboard**](MonitoringDashboard.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **workload_details**
> Workload workload_details(namespace, workload, service_domain, authorization)



Endpoint to get the workload details

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
workload = 'workload_example' # str | The workload name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.workload_details(namespace, workload, service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->workload_details: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **workload** | **str**| The workload name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**Workload**](Workload.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **workload_health**
> WorkloadHealth workload_health(namespace, workload, service_domain, authorization)



Get health associated to the given workload

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
workload = 'workload_example' # str | The workload name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.workload_health(namespace, workload, service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->workload_health: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **workload** | **str**| The workload name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**WorkloadHealth**](WorkloadHealth.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **workload_list**
> WorkloadList workload_list(namespace, service_domain, authorization)



Endpoint to get the list of workloads for a namespace

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.

try:
    api_response = api_instance.workload_list(namespace, service_domain, authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->workload_list: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 

### Return type

[**WorkloadList**](WorkloadList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **workload_metrics**
> KialiMetrics workload_metrics(namespace, workload, service_domain, authorization, avg=avg, by_labels=by_labels, direction=direction, duration=duration, filters=filters, quantiles=quantiles, rate_func=rate_func, rate_interval=rate_interval, request_protocol=request_protocol, reporter=reporter, step=step, version=version)



Endpoint to fetch metrics to be displayed, related to a single workload

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
api_instance = kps_api.KialiApi(kps_api.ApiClient(configuration))
namespace = 'namespace_example' # str | The namespace name.
workload = 'workload_example' # str | The workload name.
service_domain = 'service_domain_example' # str | ID of ServiceDomain to access.
authorization = 'authorization_example' # str | Format: Bearer &lt;token>, with &lt;token> from login API response.
avg = true # bool | Flag for fetching histogram average. Default is true. (optional)
by_labels = ['by_labels_example'] # list[str] | List of labels to use for grouping metrics (via Prometheus 'by' clause). (optional)
direction = 'direction_example' # str | Traffic direction: 'inbound' or 'outbound'. (optional)
duration = 789 # int | Duration of the query period, in seconds. (optional)
filters = ['filters_example'] # list[str] | List of metrics to fetch. Fetch all metrics when empty. List entries are Kiali internal metric names. (optional)
quantiles = ['quantiles_example'] # list[str] | List of quantiles to fetch. Fetch no quantiles when empty. Ex: [0.5, 0.95, 0.99]. (optional)
rate_func = 'rate_func_example' # str | Prometheus function used to calculate rate: 'rate' or 'irate'. (optional)
rate_interval = 'rate_interval_example' # str | Interval used for rate and histogram calculation. (optional)
request_protocol = 'request_protocol_example' # str | Desired request protocol for the telemetry: For example, 'http' or 'grpc'. (optional)
reporter = 'reporter_example' # str | Istio telemetry reporter: 'source' or 'destination'. (optional)
step = 789 # int | Step between [graph] datapoints, in seconds. (optional)
version = 'version_example' # str | Filters metrics by the specified version. (optional)

try:
    api_response = api_instance.workload_metrics(namespace, workload, service_domain, authorization, avg=avg, by_labels=by_labels, direction=direction, duration=duration, filters=filters, quantiles=quantiles, rate_func=rate_func, rate_interval=rate_interval, request_protocol=request_protocol, reporter=reporter, step=step, version=version)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling KialiApi->workload_metrics: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **str**| The namespace name. | 
 **workload** | **str**| The workload name. | 
 **service_domain** | **str**| ID of ServiceDomain to access. | 
 **authorization** | **str**| Format: Bearer &amp;lt;token&gt;, with &amp;lt;token&gt; from login API response. | 
 **avg** | **bool**| Flag for fetching histogram average. Default is true. | [optional] 
 **by_labels** | [**list[str]**](str.md)| List of labels to use for grouping metrics (via Prometheus &#x27;by&#x27; clause). | [optional] 
 **direction** | **str**| Traffic direction: &#x27;inbound&#x27; or &#x27;outbound&#x27;. | [optional] 
 **duration** | **int**| Duration of the query period, in seconds. | [optional] 
 **filters** | [**list[str]**](str.md)| List of metrics to fetch. Fetch all metrics when empty. List entries are Kiali internal metric names. | [optional] 
 **quantiles** | [**list[str]**](str.md)| List of quantiles to fetch. Fetch no quantiles when empty. Ex: [0.5, 0.95, 0.99]. | [optional] 
 **rate_func** | **str**| Prometheus function used to calculate rate: &#x27;rate&#x27; or &#x27;irate&#x27;. | [optional] 
 **rate_interval** | **str**| Interval used for rate and histogram calculation. | [optional] 
 **request_protocol** | **str**| Desired request protocol for the telemetry: For example, &#x27;http&#x27; or &#x27;grpc&#x27;. | [optional] 
 **reporter** | **str**| Istio telemetry reporter: &#x27;source&#x27; or &#x27;destination&#x27;. | [optional] 
 **step** | **int**| Step between [graph] datapoints, in seconds. | [optional] 
 **version** | **str**| Filters metrics by the specified version. | [optional] 

### Return type

[**KialiMetrics**](KialiMetrics.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

