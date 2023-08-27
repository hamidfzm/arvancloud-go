/*
Arvan VOD

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vod

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ProfileAPIService ProfileAPI service
type ProfileAPIService service

type ApiChannelsChannelProfilesGetRequest struct {
	ctx context.Context
	ApiService *ProfileAPIService
	channel string
	filter *string
	page *int32
	perPage *int32
}

// Filter result
func (r ApiChannelsChannelProfilesGetRequest) Filter(filter string) ApiChannelsChannelProfilesGetRequest {
	r.filter = &filter
	return r
}

// Page number
func (r ApiChannelsChannelProfilesGetRequest) Page(page int32) ApiChannelsChannelProfilesGetRequest {
	r.page = &page
	return r
}

// Page limit
func (r ApiChannelsChannelProfilesGetRequest) PerPage(perPage int32) ApiChannelsChannelProfilesGetRequest {
	r.perPage = &perPage
	return r
}

func (r ApiChannelsChannelProfilesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ChannelsChannelProfilesGetExecute(r)
}

/*
ChannelsChannelProfilesGet Return all channel's profiles.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channel The Id of channel
 @return ApiChannelsChannelProfilesGetRequest
*/
func (a *ProfileAPIService) ChannelsChannelProfilesGet(ctx context.Context, channel string) ApiChannelsChannelProfilesGetRequest {
	return ApiChannelsChannelProfilesGetRequest{
		ApiService: a,
		ctx: ctx,
		channel: channel,
	}
}

// Execute executes the request
func (a *ProfileAPIService) ChannelsChannelProfilesGetExecute(r ApiChannelsChannelProfilesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProfileAPIService.ChannelsChannelProfilesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel}/profiles"
	localVarPath = strings.Replace(localVarPath, "{"+"channel"+"}", url.PathEscape(parameterValueToString(r.channel, "channel")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiChannelsChannelProfilesPostRequest struct {
	ctx context.Context
	ApiService *ProfileAPIService
	channel string
	body *ChannelsChannelProfilesPostRequest
}

// Profile details
func (r ApiChannelsChannelProfilesPostRequest) Body(body ChannelsChannelProfilesPostRequest) ApiChannelsChannelProfilesPostRequest {
	r.body = &body
	return r
}

func (r ApiChannelsChannelProfilesPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ChannelsChannelProfilesPostExecute(r)
}

/*
ChannelsChannelProfilesPost Store a newly created profile.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channel The Id of channel
 @return ApiChannelsChannelProfilesPostRequest
*/
func (a *ProfileAPIService) ChannelsChannelProfilesPost(ctx context.Context, channel string) ApiChannelsChannelProfilesPostRequest {
	return ApiChannelsChannelProfilesPostRequest{
		ApiService: a,
		ctx: ctx,
		channel: channel,
	}
}

// Execute executes the request
func (a *ProfileAPIService) ChannelsChannelProfilesPostExecute(r ApiChannelsChannelProfilesPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProfileAPIService.ChannelsChannelProfilesPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel}/profiles"
	localVarPath = strings.Replace(localVarPath, "{"+"channel"+"}", url.PathEscape(parameterValueToString(r.channel, "channel")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiProfilesProfileDeleteRequest struct {
	ctx context.Context
	ApiService *ProfileAPIService
	profile string
}

func (r ApiProfilesProfileDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ProfilesProfileDeleteExecute(r)
}

/*
ProfilesProfileDelete Remove the specified profile.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param profile The Id of profile
 @return ApiProfilesProfileDeleteRequest
*/
func (a *ProfileAPIService) ProfilesProfileDelete(ctx context.Context, profile string) ApiProfilesProfileDeleteRequest {
	return ApiProfilesProfileDeleteRequest{
		ApiService: a,
		ctx: ctx,
		profile: profile,
	}
}

// Execute executes the request
func (a *ProfileAPIService) ProfilesProfileDeleteExecute(r ApiProfilesProfileDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProfileAPIService.ProfilesProfileDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/profiles/{profile}"
	localVarPath = strings.Replace(localVarPath, "{"+"profile"+"}", url.PathEscape(parameterValueToString(r.profile, "profile")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiProfilesProfileGetRequest struct {
	ctx context.Context
	ApiService *ProfileAPIService
	profile string
}

func (r ApiProfilesProfileGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ProfilesProfileGetExecute(r)
}

/*
ProfilesProfileGet Return the specified profile.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param profile The Id of profile
 @return ApiProfilesProfileGetRequest
*/
func (a *ProfileAPIService) ProfilesProfileGet(ctx context.Context, profile string) ApiProfilesProfileGetRequest {
	return ApiProfilesProfileGetRequest{
		ApiService: a,
		ctx: ctx,
		profile: profile,
	}
}

// Execute executes the request
func (a *ProfileAPIService) ProfilesProfileGetExecute(r ApiProfilesProfileGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProfileAPIService.ProfilesProfileGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/profiles/{profile}"
	localVarPath = strings.Replace(localVarPath, "{"+"profile"+"}", url.PathEscape(parameterValueToString(r.profile, "profile")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiProfilesProfilePatchRequest struct {
	ctx context.Context
	ApiService *ProfileAPIService
	profile string
	body *ProfilesProfilePatchRequest
}

// Profile details
func (r ApiProfilesProfilePatchRequest) Body(body ProfilesProfilePatchRequest) ApiProfilesProfilePatchRequest {
	r.body = &body
	return r
}

func (r ApiProfilesProfilePatchRequest) Execute() (*http.Response, error) {
	return r.ApiService.ProfilesProfilePatchExecute(r)
}

/*
ProfilesProfilePatch Update the specified profile.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param profile The Id of profile
 @return ApiProfilesProfilePatchRequest
*/
func (a *ProfileAPIService) ProfilesProfilePatch(ctx context.Context, profile string) ApiProfilesProfilePatchRequest {
	return ApiProfilesProfilePatchRequest{
		ApiService: a,
		ctx: ctx,
		profile: profile,
	}
}

// Execute executes the request
func (a *ProfileAPIService) ProfilesProfilePatchExecute(r ApiProfilesProfilePatchRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProfileAPIService.ProfilesProfilePatch")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/profiles/{profile}"
	localVarPath = strings.Replace(localVarPath, "{"+"profile"+"}", url.PathEscape(parameterValueToString(r.profile, "profile")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}