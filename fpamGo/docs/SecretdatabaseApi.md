# \SecretdatabaseApi

All URIs are relative to *https://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CmdbSecretDatabaseIdDelete**](SecretdatabaseApi.md#CmdbSecretDatabaseIdDelete) | **Delete** /cmdb/secret/database/{id} | Delete this specific resource.  Access Group: secretgrp
[**CmdbSecretDatabaseIdGet**](SecretdatabaseApi.md#CmdbSecretDatabaseIdGet) | **Get** /cmdb/secret/database/{id} | Select a specific entry from a CLI table.  Access Group: secretgrp
[**CmdbSecretDatabaseIdPut**](SecretdatabaseApi.md#CmdbSecretDatabaseIdPut) | **Put** /cmdb/secret/database/{id} | Update this specific resource.  Access Group: secretgrp
[**CmdbSecretDatabasePost**](SecretdatabaseApi.md#CmdbSecretDatabasePost) | **Post** /cmdb/secret/database | Create object(s) in this table.  Access Group: secretgrp


# **CmdbSecretDatabaseIdDelete**
> interface{} CmdbSecretDatabaseIdDelete(ctx, id, optional)
Delete this specific resource.  Access Group: secretgrp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| mkey | 
 **optional** | ***SecretdatabaseApiCmdbSecretDatabaseIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretdatabaseApiCmdbSecretDatabaseIdDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vdom** | [**optional.Interface of []string**](string.md)| Specify the Virtual Domain(s) from which results are returned or changes are applied to. If this parameter is not provided, the management VDOM will be used. If the admin does not have access to the VDOM, a permission error will be returned. The URL parameter is one of: vdom&#x3D;root (Single VDOM) vdom&#x3D;vdom1,vdom2 (Multiple VDOMs) vdom&#x3D;* (All VDOMs)  | 

### Return type

**interface{}**

### Authorization

[APIKeyQueryParam](../README.md#APIKeyQueryParam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CmdbSecretDatabaseIdGet**
> interface{} CmdbSecretDatabaseIdGet(ctx, id, optional)
Select a specific entry from a CLI table.  Access Group: secretgrp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| mkey | 
 **optional** | ***SecretdatabaseApiCmdbSecretDatabaseIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretdatabaseApiCmdbSecretDatabaseIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.Bool**| Enable to include datasource information for each linked object. | 
 **withMeta** | **optional.Bool**| Enable to include meta information about each object (type id, references, etc). | 
 **skip** | **optional.Bool**| Enable to call CLI skip operator to hide skipped properties. | 
 **format** | [**optional.Interface of []string**](string.md)| List of property names to include in results, separated by | (i.e. policyid|srcintf). | 
 **action** | **optional.String**| stats: Return CMDB aggregated statistics. default: Return the CLI default values for this object type. schema: Return the CLI schema for this object type. revision: Return the CMDB revision for this object type. transaction-list: List all configuration transaction(s).  | 
 **vdom** | [**optional.Interface of []string**](string.md)| Specify the Virtual Domain(s) from which results are returned or changes are applied to. If this parameter is not provided, the management VDOM will be used. If the admin does not have access to the VDOM, a permission error will be returned. The URL parameter is one of: vdom&#x3D;root (Single VDOM) vdom&#x3D;vdom1,vdom2 (Multiple VDOMs) vdom&#x3D;* (All VDOMs)  | 
 **secretEncoding** | **optional.String**| Credential encryption  | 
 **fieldname** | **optional.String**| Filter results by field name  | 

### Return type

**interface{}**

### Authorization

[APIKeyQueryParam](../README.md#APIKeyQueryParam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CmdbSecretDatabaseIdPut**
> interface{} CmdbSecretDatabaseIdPut(ctx, body, id, optional)
Update this specific resource.  Access Group: secretgrp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body**](Body.md)| Possible parameters to go in the body for the request | 
  **id** | **int32**| mkey | 
 **optional** | ***SecretdatabaseApiCmdbSecretDatabaseIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretdatabaseApiCmdbSecretDatabaseIdPutOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vdom** | [**optional.Interface of []string**](string.md)| Specify the Virtual Domain(s) from which results are returned or changes are applied to. If this parameter is not provided, the management VDOM will be used. If the admin does not have access to the VDOM, a permission error will be returned. The URL parameter is one of: vdom&#x3D;root (Single VDOM) vdom&#x3D;vdom1,vdom2 (Multiple VDOMs) vdom&#x3D;* (All VDOMs)  | 
 **action** | **optional.String**| If supported, an action can be specified. _move_: Move this specific resource. When *action&#x3D;move* is set, one of the extra parameters (*before*, *after*) must be provided. __*Note:*__ If this parameter is provided when not supported, the action will be ignored and an “invalid request” error will be returned.  | 
 **before** | **optional.String**| If *action&#x3D;move*, use *before* to specify the ID of the resource that this resource will be moved before. For example, to move &#x60;object 1&#x60; to before &#x60;object 2&#x60;, use: __action&#x3D;move&amp;before&#x3D;2__ __*Note:*__ This parameter can only be used when the *action* parameter is set to *move*.  | 
 **after** | **optional.String**| If *action&#x3D;move*, use *after* to specify the ID of the resource that this resource will be moved after. For example, to move &#x60;object 1&#x60; to after &#x60;object 3&#x60;, use: __action&#x3D;move&amp;after&#x3D;3__ __*Note:*__ This parameter can only be used when the *action* parameter is set to *move*.  | 

### Return type

**interface{}**

### Authorization

[APIKeyQueryParam](../README.md#APIKeyQueryParam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CmdbSecretDatabasePost**
> interface{} CmdbSecretDatabasePost(ctx, body, optional)
Create object(s) in this table.  Access Group: secretgrp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body1**](Body1.md)| Possible parameters to go in the body for the request | 
 **optional** | ***SecretdatabaseApiCmdbSecretDatabasePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretdatabaseApiCmdbSecretDatabasePostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vdom** | [**optional.Interface of []string**](string.md)| Specify the Virtual Domain(s) from which results are returned or changes are applied to. If this parameter is not provided, the management VDOM will be used. If the admin does not have access to the VDOM, a permission error will be returned. The URL parameter is one of: vdom&#x3D;root (Single VDOM) vdom&#x3D;vdom1,vdom2 (Multiple VDOMs) vdom&#x3D;* (All VDOMs)  | 
 **action** | **optional.String**| If supported, an action can be specified. _clone_: Clone this specific resource. When *action&#x3D;clone* is set, the extra parameters *nkey* must be provided. __*Note:*__ If this parameter is provided when not supported, the action will be ignored and an “invalid request” error will be returned.  | 
 **nkey** | **optional.String**|  If *action&#x3D;clone*, use *nkey* to specify the ID for the new resource to be created. For example, to clone &#x60;address1&#x60; to &#x60;address1_clone&#x60;, use: __action&#x3D;clone&amp;nkey&#x3D;address1_clone__ __*Note:*__ This parameter can only be used when the *action* parameter is set to *clone*.  | 

### Return type

**interface{}**

### Authorization

[APIKeyQueryParam](../README.md#APIKeyQueryParam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

