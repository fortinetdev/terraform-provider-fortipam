
/*
 * FortiPAM
 *
 * REST API for configuring FortiPAM objects and settings
 *
 * API version: v1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package fpam_go

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type SecretdatabaseApiService service

/*
SecretdatabaseApiService Delete this specific resource.  Access Group: secretgrp
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id mkey
 * @param optional nil or *SecretdatabaseApiCmdbSecretDatabaseIdDeleteOpts - Optional Parameters:
     * @param "Vdom" (optional.Interface of []string) -  Specify the Virtual Domain(s) from which results are returned or changes are applied to. If this parameter is not provided, the management VDOM will be used. If the admin does not have access to the VDOM, a permission error will be returned. The URL parameter is one of: vdom&#x3D;root (Single VDOM) vdom&#x3D;vdom1,vdom2 (Multiple VDOMs) vdom&#x3D;* (All VDOMs) 

@return interface{}
*/

type SecretdatabaseApiCmdbSecretDatabaseIdDeleteOpts struct { 
	Vdom optional.Interface
}

func (a *SecretdatabaseApiService) CmdbSecretDatabaseIdDelete(ctx context.Context, id int32, localVarOptionals *SecretdatabaseApiCmdbSecretDatabaseIdDeleteOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cmdb/secret/database/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Vdom.IsSet() {
		localVarQueryParams.Add("vdom", parameterToString(localVarOptionals.Vdom.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
SecretdatabaseApiService Select a specific entry from a CLI table.  Access Group: secretgrp
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id mkey
 * @param optional nil or *SecretdatabaseApiCmdbSecretDatabaseIdGetOpts - Optional Parameters:
     * @param "Datasource" (optional.Bool) -  Enable to include datasource information for each linked object.
     * @param "WithMeta" (optional.Bool) -  Enable to include meta information about each object (type id, references, etc).
     * @param "Skip" (optional.Bool) -  Enable to call CLI skip operator to hide skipped properties.
     * @param "Format" (optional.Interface of []string) -  List of property names to include in results, separated by | (i.e. policyid|srcintf).
     * @param "Action" (optional.String) -  stats: Return CMDB aggregated statistics. default: Return the CLI default values for this object type. schema: Return the CLI schema for this object type. revision: Return the CMDB revision for this object type. transaction-list: List all configuration transaction(s). 
     * @param "Vdom" (optional.Interface of []string) -  Specify the Virtual Domain(s) from which results are returned or changes are applied to. If this parameter is not provided, the management VDOM will be used. If the admin does not have access to the VDOM, a permission error will be returned. The URL parameter is one of: vdom&#x3D;root (Single VDOM) vdom&#x3D;vdom1,vdom2 (Multiple VDOMs) vdom&#x3D;* (All VDOMs) 
     * @param "SecretEncoding" (optional.String) -  Credential encryption 
     * @param "Fieldname" (optional.String) -  Filter results by field name 

@return interface{}
*/

type SecretdatabaseApiCmdbSecretDatabaseIdGetOpts struct { 
	Datasource optional.Bool
	WithMeta optional.Bool
	Skip optional.Bool
	Format optional.Interface
	Action optional.String
	Vdom optional.Interface
	SecretEncoding optional.String
	Fieldname optional.String
}

func (a *SecretdatabaseApiService) CmdbSecretDatabaseIdGet(ctx context.Context, id int32, localVarOptionals *SecretdatabaseApiCmdbSecretDatabaseIdGetOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cmdb/secret/database/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Datasource.IsSet() {
		localVarQueryParams.Add("datasource", parameterToString(localVarOptionals.Datasource.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WithMeta.IsSet() {
		localVarQueryParams.Add("with_meta", parameterToString(localVarOptionals.WithMeta.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Skip.IsSet() {
		localVarQueryParams.Add("skip", parameterToString(localVarOptionals.Skip.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Format.IsSet() {
		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format.Value(), "pipes"))
	}
	if localVarOptionals != nil && localVarOptionals.Action.IsSet() {
		localVarQueryParams.Add("action", parameterToString(localVarOptionals.Action.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Vdom.IsSet() {
		localVarQueryParams.Add("vdom", parameterToString(localVarOptionals.Vdom.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.SecretEncoding.IsSet() {
		localVarQueryParams.Add("secret_encoding", parameterToString(localVarOptionals.SecretEncoding.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Fieldname.IsSet() {
		localVarQueryParams.Add("fieldname", parameterToString(localVarOptionals.Fieldname.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
SecretdatabaseApiService Update this specific resource.  Access Group: secretgrp
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body Possible parameters to go in the body for the request
 * @param id mkey
 * @param optional nil or *SecretdatabaseApiCmdbSecretDatabaseIdPutOpts - Optional Parameters:
     * @param "Vdom" (optional.Interface of []string) -  Specify the Virtual Domain(s) from which results are returned or changes are applied to. If this parameter is not provided, the management VDOM will be used. If the admin does not have access to the VDOM, a permission error will be returned. The URL parameter is one of: vdom&#x3D;root (Single VDOM) vdom&#x3D;vdom1,vdom2 (Multiple VDOMs) vdom&#x3D;* (All VDOMs) 
     * @param "Action" (optional.String) -  If supported, an action can be specified. _move_: Move this specific resource. When *action&#x3D;move* is set, one of the extra parameters (*before*, *after*) must be provided. __*Note:*__ If this parameter is provided when not supported, the action will be ignored and an “invalid request” error will be returned. 
     * @param "Before" (optional.String) -  If *action&#x3D;move*, use *before* to specify the ID of the resource that this resource will be moved before. For example, to move &#x60;object 1&#x60; to before &#x60;object 2&#x60;, use: __action&#x3D;move&amp;before&#x3D;2__ __*Note:*__ This parameter can only be used when the *action* parameter is set to *move*. 
     * @param "After" (optional.String) -  If *action&#x3D;move*, use *after* to specify the ID of the resource that this resource will be moved after. For example, to move &#x60;object 1&#x60; to after &#x60;object 3&#x60;, use: __action&#x3D;move&amp;after&#x3D;3__ __*Note:*__ This parameter can only be used when the *action* parameter is set to *move*. 

@return interface{}
*/

type SecretdatabaseApiCmdbSecretDatabaseIdPutOpts struct { 
	Vdom optional.Interface
	Action optional.String
	Before optional.String
	After optional.String
}

func (a *SecretdatabaseApiService) CmdbSecretDatabaseIdPut(ctx context.Context, body Body, id int32, localVarOptionals *SecretdatabaseApiCmdbSecretDatabaseIdPutOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cmdb/secret/database/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Vdom.IsSet() {
		localVarQueryParams.Add("vdom", parameterToString(localVarOptionals.Vdom.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Action.IsSet() {
		localVarQueryParams.Add("action", parameterToString(localVarOptionals.Action.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Before.IsSet() {
		localVarQueryParams.Add("before", parameterToString(localVarOptionals.Before.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.After.IsSet() {
		localVarQueryParams.Add("after", parameterToString(localVarOptionals.After.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
SecretdatabaseApiService Create object(s) in this table.  Access Group: secretgrp
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body Possible parameters to go in the body for the request
 * @param optional nil or *SecretdatabaseApiCmdbSecretDatabasePostOpts - Optional Parameters:
     * @param "Vdom" (optional.Interface of []string) -  Specify the Virtual Domain(s) from which results are returned or changes are applied to. If this parameter is not provided, the management VDOM will be used. If the admin does not have access to the VDOM, a permission error will be returned. The URL parameter is one of: vdom&#x3D;root (Single VDOM) vdom&#x3D;vdom1,vdom2 (Multiple VDOMs) vdom&#x3D;* (All VDOMs) 
     * @param "Action" (optional.String) -  If supported, an action can be specified. _clone_: Clone this specific resource. When *action&#x3D;clone* is set, the extra parameters *nkey* must be provided. __*Note:*__ If this parameter is provided when not supported, the action will be ignored and an “invalid request” error will be returned. 
     * @param "Nkey" (optional.String) -   If *action&#x3D;clone*, use *nkey* to specify the ID for the new resource to be created. For example, to clone &#x60;address1&#x60; to &#x60;address1_clone&#x60;, use: __action&#x3D;clone&amp;nkey&#x3D;address1_clone__ __*Note:*__ This parameter can only be used when the *action* parameter is set to *clone*. 

@return interface{}
*/

type SecretdatabaseApiCmdbSecretDatabasePostOpts struct { 
	Vdom optional.Interface
	Action optional.String
	Nkey optional.String
}

func (a *SecretdatabaseApiService) CmdbSecretDatabasePost(ctx context.Context, body Body1, localVarOptionals *SecretdatabaseApiCmdbSecretDatabasePostOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cmdb/secret/database"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Vdom.IsSet() {
		localVarQueryParams.Add("vdom", parameterToString(localVarOptionals.Vdom.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Action.IsSet() {
		localVarQueryParams.Add("action", parameterToString(localVarOptionals.Action.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Nkey.IsSet() {
		localVarQueryParams.Add("nkey", parameterToString(localVarOptionals.Nkey.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
