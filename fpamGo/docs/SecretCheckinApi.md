# \SecretCheckinApi

All URIs are relative to *https://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InternalSecretCheckinPost**](SecretCheckinApi.md#InternalSecretCheckinPost) | **Post** /internal/secret-checkin | Secret checkin.


# **InternalSecretCheckinPost**
> interface{} InternalSecretCheckinPost(ctx, body)
Secret checkin.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body3**](Body3.md)| Possible parameters to go in the body for the request. | 

### Return type

**interface{}**

### Authorization

[APIKeyQueryParam](../README.md#APIKeyQueryParam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

