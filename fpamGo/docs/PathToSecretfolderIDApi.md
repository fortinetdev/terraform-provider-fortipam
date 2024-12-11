# \PathToSecretfolderIDApi

All URIs are relative to *https://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UtilityIdPathGet**](PathToSecretfolderIDApi.md#UtilityIdPathGet) | **Get** /utility/id/{path} | Get secret/folder ID by providing full path  Access Group: secretgrp


# **UtilityIdPathGet**
> interface{} UtilityIdPathGet(ctx, type_, path)
Get secret/folder ID by providing full path  Access Group: secretgrp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| Indicate request is to retrieve a secret or folder  | 
  **path** | **string**| Full path to a secret or folder  | 

### Return type

**interface{}**

### Authorization

[APIKeyQueryParam](../README.md#APIKeyQueryParam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

