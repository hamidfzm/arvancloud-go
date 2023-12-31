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
	"os"
)


// WatermarkAPIService WatermarkAPI service
type WatermarkAPIService service

type ApiChannelsChannelWatermarksGetRequest struct {
	ctx context.Context
	ApiService *WatermarkAPIService
	channel string
	filter *string
	page *int32
	perPage *int32
	secureIp *string
	secureExpireTime *int32
}

// Filter result
func (r ApiChannelsChannelWatermarksGetRequest) Filter(filter string) ApiChannelsChannelWatermarksGetRequest {
	r.filter = &filter
	return r
}

// Page number
func (r ApiChannelsChannelWatermarksGetRequest) Page(page int32) ApiChannelsChannelWatermarksGetRequest {
	r.page = &page
	return r
}

// Page limit for query
func (r ApiChannelsChannelWatermarksGetRequest) PerPage(perPage int32) ApiChannelsChannelWatermarksGetRequest {
	r.perPage = &perPage
	return r
}

// The IP address for generate secure links for. If channel is secure default is request IP
func (r ApiChannelsChannelWatermarksGetRequest) SecureIp(secureIp string) ApiChannelsChannelWatermarksGetRequest {
	r.secureIp = &secureIp
	return r
}

// The Unix Timestamp for expire secure links.      *          If channel is secure default is 24 hours later from now
func (r ApiChannelsChannelWatermarksGetRequest) SecureExpireTime(secureExpireTime int32) ApiChannelsChannelWatermarksGetRequest {
	r.secureExpireTime = &secureExpireTime
	return r
}

func (r ApiChannelsChannelWatermarksGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ChannelsChannelWatermarksGetExecute(r)
}

/*
ChannelsChannelWatermarksGet Return all channel's watermarks.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channel The Id of channel
 @return ApiChannelsChannelWatermarksGetRequest
*/
func (a *WatermarkAPIService) ChannelsChannelWatermarksGet(ctx context.Context, channel string) ApiChannelsChannelWatermarksGetRequest {
	return ApiChannelsChannelWatermarksGetRequest{
		ApiService: a,
		ctx: ctx,
		channel: channel,
	}
}

// Execute executes the request
func (a *WatermarkAPIService) ChannelsChannelWatermarksGetExecute(r ApiChannelsChannelWatermarksGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WatermarkAPIService.ChannelsChannelWatermarksGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel}/watermarks"
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
	if r.secureIp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "secure_ip", r.secureIp, "")
	}
	if r.secureExpireTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "secure_expire_time", r.secureExpireTime, "")
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

type ApiChannelsChannelWatermarksPostRequest struct {
	ctx context.Context
	ApiService *WatermarkAPIService
	channel string
	title *string
	watermark *os.File
	description *string
}

// Title of watermark
func (r ApiChannelsChannelWatermarksPostRequest) Title(title string) ApiChannelsChannelWatermarksPostRequest {
	r.title = &title
	return r
}

// Watermark file
func (r ApiChannelsChannelWatermarksPostRequest) Watermark(watermark *os.File) ApiChannelsChannelWatermarksPostRequest {
	r.watermark = watermark
	return r
}

// Description of watermark
func (r ApiChannelsChannelWatermarksPostRequest) Description(description string) ApiChannelsChannelWatermarksPostRequest {
	r.description = &description
	return r
}

func (r ApiChannelsChannelWatermarksPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ChannelsChannelWatermarksPostExecute(r)
}

/*
ChannelsChannelWatermarksPost Store a newly created Watermark.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channel The Id of channel
 @return ApiChannelsChannelWatermarksPostRequest
*/
func (a *WatermarkAPIService) ChannelsChannelWatermarksPost(ctx context.Context, channel string) ApiChannelsChannelWatermarksPostRequest {
	return ApiChannelsChannelWatermarksPostRequest{
		ApiService: a,
		ctx: ctx,
		channel: channel,
	}
}

// Execute executes the request
func (a *WatermarkAPIService) ChannelsChannelWatermarksPostExecute(r ApiChannelsChannelWatermarksPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WatermarkAPIService.ChannelsChannelWatermarksPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel}/watermarks"
	localVarPath = strings.Replace(localVarPath, "{"+"channel"+"}", url.PathEscape(parameterValueToString(r.channel, "channel")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.title == nil {
		return nil, reportError("title is required and must be specified")
	}
	if r.watermark == nil {
		return nil, reportError("watermark is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	parameterAddToHeaderOrQuery(localVarFormParams, "title", r.title, "")
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "description", r.description, "")
	}
	var watermarkLocalVarFormFileName string
	var watermarkLocalVarFileName     string
	var watermarkLocalVarFileBytes    []byte

	watermarkLocalVarFormFileName = "watermark"


	watermarkLocalVarFile := r.watermark

	if watermarkLocalVarFile != nil {
		fbs, _ := io.ReadAll(watermarkLocalVarFile)

		watermarkLocalVarFileBytes = fbs
		watermarkLocalVarFileName = watermarkLocalVarFile.Name()
		watermarkLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: watermarkLocalVarFileBytes, fileName: watermarkLocalVarFileName, formFileName: watermarkLocalVarFormFileName})
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

type ApiWatermarksWatermarkDeleteRequest struct {
	ctx context.Context
	ApiService *WatermarkAPIService
	watermark string
}

func (r ApiWatermarksWatermarkDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.WatermarksWatermarkDeleteExecute(r)
}

/*
WatermarksWatermarkDelete Remove the specified watermark.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param watermark The Id of watermark
 @return ApiWatermarksWatermarkDeleteRequest
*/
func (a *WatermarkAPIService) WatermarksWatermarkDelete(ctx context.Context, watermark string) ApiWatermarksWatermarkDeleteRequest {
	return ApiWatermarksWatermarkDeleteRequest{
		ApiService: a,
		ctx: ctx,
		watermark: watermark,
	}
}

// Execute executes the request
func (a *WatermarkAPIService) WatermarksWatermarkDeleteExecute(r ApiWatermarksWatermarkDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WatermarkAPIService.WatermarksWatermarkDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/watermarks/{watermark}"
	localVarPath = strings.Replace(localVarPath, "{"+"watermark"+"}", url.PathEscape(parameterValueToString(r.watermark, "watermark")), -1)

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

type ApiWatermarksWatermarkGetRequest struct {
	ctx context.Context
	ApiService *WatermarkAPIService
	watermark string
}

func (r ApiWatermarksWatermarkGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.WatermarksWatermarkGetExecute(r)
}

/*
WatermarksWatermarkGet Return the specified watermark.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param watermark The Id of watermark
 @return ApiWatermarksWatermarkGetRequest
*/
func (a *WatermarkAPIService) WatermarksWatermarkGet(ctx context.Context, watermark string) ApiWatermarksWatermarkGetRequest {
	return ApiWatermarksWatermarkGetRequest{
		ApiService: a,
		ctx: ctx,
		watermark: watermark,
	}
}

// Execute executes the request
func (a *WatermarkAPIService) WatermarksWatermarkGetExecute(r ApiWatermarksWatermarkGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WatermarkAPIService.WatermarksWatermarkGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/watermarks/{watermark}"
	localVarPath = strings.Replace(localVarPath, "{"+"watermark"+"}", url.PathEscape(parameterValueToString(r.watermark, "watermark")), -1)

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

type ApiWatermarksWatermarkPatchRequest struct {
	ctx context.Context
	ApiService *WatermarkAPIService
	watermark string
	body *WatermarksWatermarkPatchRequest
}

// Watermark details
func (r ApiWatermarksWatermarkPatchRequest) Body(body WatermarksWatermarkPatchRequest) ApiWatermarksWatermarkPatchRequest {
	r.body = &body
	return r
}

func (r ApiWatermarksWatermarkPatchRequest) Execute() (*http.Response, error) {
	return r.ApiService.WatermarksWatermarkPatchExecute(r)
}

/*
WatermarksWatermarkPatch Update the specified watermark.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param watermark The Id of watermark
 @return ApiWatermarksWatermarkPatchRequest
*/
func (a *WatermarkAPIService) WatermarksWatermarkPatch(ctx context.Context, watermark string) ApiWatermarksWatermarkPatchRequest {
	return ApiWatermarksWatermarkPatchRequest{
		ApiService: a,
		ctx: ctx,
		watermark: watermark,
	}
}

// Execute executes the request
func (a *WatermarkAPIService) WatermarksWatermarkPatchExecute(r ApiWatermarksWatermarkPatchRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WatermarkAPIService.WatermarksWatermarkPatch")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/watermarks/{watermark}"
	localVarPath = strings.Replace(localVarPath, "{"+"watermark"+"}", url.PathEscape(parameterValueToString(r.watermark, "watermark")), -1)

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
