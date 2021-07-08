# SoftwareUpdateServiceDomainListPayload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**order_by** | **str** | Specify result order. Zero or more entries with format: &amp;ltkey&gt; [desc] where orderByKeys lists allowed keys in each response. | [optional] 
**order_by_keys** | **list[str]** | Keys that can be used in orderBy. | [optional] 
**page_index** | **int** | 0-based index of the page to fetch results. | 
**page_size** | **int** | Item count of each page. | 
**result** | [**list[SoftwareUpdateServiceDomain]**](SoftwareUpdateServiceDomain.md) | list of service domain stats | 
**total_count** | **int** | Count of all items matching the query. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

