# kps_api.AuditlogApi

All URIs are relative to *//localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**query_audit_logs_v2**](AuditlogApi.md#query_audit_logs_v2) | **POST** /v1.0/auditlogsV2 | Lists audit logs matching the provided filter.

# **query_audit_logs_v2**
> list[AuditLogV2] query_audit_logs_v2(authorization)

Lists audit logs matching the provided filter.

Retrieves all audit logs matching the filter for a tenant.

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
api_instance = kps_api.AuditlogApi(kps_api.ApiClient(configuration))
authorization = 'authorization_example' # str | Format: Bearer <token>, with <token> from login API response.

try:
    # Lists audit logs matching the provided filter.
    api_response = api_instance.query_audit_logs_v2(authorization)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling AuditlogApi->query_audit_logs_v2: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **str**| Format: Bearer &lt;token&gt;, with &lt;token&gt; from login API response. | 

### Return type

[**list[AuditLogV2]**](AuditLogV2.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

