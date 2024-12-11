# \SecretCheckoutApi

All URIs are relative to *https://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InternalSecretCheckoutPost**](SecretCheckoutApi.md#InternalSecretCheckoutPost) | **Post** /internal/secret-checkout | Secret checkout.


# **InternalSecretCheckoutPost**
> interface{} InternalSecretCheckoutPost(ctx, body)
Secret checkout.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body2**](Body2.md)| Possible parameters to go in the body for the request. | 

### Return type

**interface{}**

### Authorization

[APIKeyQueryParam](../README.md#APIKeyQueryParam)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

