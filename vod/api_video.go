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


// VideoAPIService VideoAPI service
type VideoAPIService service

type ApiChannelsChannelVideosGetRequest struct {
	ctx context.Context
	ApiService *VideoAPIService
	channel string
	filter *string
	page *int32
	perPage *int32
	secureIp *string
	secureExpireTime *int32
	orderBy *string
	sort *string
}

// Filter result
func (r ApiChannelsChannelVideosGetRequest) Filter(filter string) ApiChannelsChannelVideosGetRequest {
	r.filter = &filter
	return r
}

// Page number
func (r ApiChannelsChannelVideosGetRequest) Page(page int32) ApiChannelsChannelVideosGetRequest {
	r.page = &page
	return r
}

// Page limit for query
func (r ApiChannelsChannelVideosGetRequest) PerPage(perPage int32) ApiChannelsChannelVideosGetRequest {
	r.perPage = &perPage
	return r
}

// The IP address for generate secure links for. If channel is secure default is request IP
func (r ApiChannelsChannelVideosGetRequest) SecureIp(secureIp string) ApiChannelsChannelVideosGetRequest {
	r.secureIp = &secureIp
	return r
}

// The Unix Timestamp for expire secure links.      *          If channel is secure default is 24 hours later from now
func (r ApiChannelsChannelVideosGetRequest) SecureExpireTime(secureExpireTime int32) ApiChannelsChannelVideosGetRequest {
	r.secureExpireTime = &secureExpireTime
	return r
}

// Order by ex: title, time
func (r ApiChannelsChannelVideosGetRequest) OrderBy(orderBy string) ApiChannelsChannelVideosGetRequest {
	r.orderBy = &orderBy
	return r
}

// Sort ex: asc ,desc 
func (r ApiChannelsChannelVideosGetRequest) Sort(sort string) ApiChannelsChannelVideosGetRequest {
	r.sort = &sort
	return r
}

func (r ApiChannelsChannelVideosGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ChannelsChannelVideosGetExecute(r)
}

/*
ChannelsChannelVideosGet Return all channel's videos.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channel The Id of channel
 @return ApiChannelsChannelVideosGetRequest
*/
func (a *VideoAPIService) ChannelsChannelVideosGet(ctx context.Context, channel string) ApiChannelsChannelVideosGetRequest {
	return ApiChannelsChannelVideosGetRequest{
		ApiService: a,
		ctx: ctx,
		channel: channel,
	}
}

// Execute executes the request
func (a *VideoAPIService) ChannelsChannelVideosGetExecute(r ApiChannelsChannelVideosGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.ChannelsChannelVideosGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel}/videos"
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
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_by", r.orderBy, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
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

type ApiChannelsChannelVideosPostRequest struct {
	ctx context.Context
	ApiService *VideoAPIService
	channel string
	body *ChannelsChannelVideosPostRequest
}

// Video details
func (r ApiChannelsChannelVideosPostRequest) Body(body ChannelsChannelVideosPostRequest) ApiChannelsChannelVideosPostRequest {
	r.body = &body
	return r
}

func (r ApiChannelsChannelVideosPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ChannelsChannelVideosPostExecute(r)
}

/*
ChannelsChannelVideosPost Store a newly created video.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channel The Id of channel
 @return ApiChannelsChannelVideosPostRequest
*/
func (a *VideoAPIService) ChannelsChannelVideosPost(ctx context.Context, channel string) ApiChannelsChannelVideosPostRequest {
	return ApiChannelsChannelVideosPostRequest{
		ApiService: a,
		ctx: ctx,
		channel: channel,
	}
}

// Execute executes the request
func (a *VideoAPIService) ChannelsChannelVideosPostExecute(r ApiChannelsChannelVideosPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.ChannelsChannelVideosPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel}/videos"
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

type ApiVideosBulkDeleteRequest struct {
	ctx context.Context
	ApiService *VideoAPIService
	ids []string
}

func (r ApiVideosBulkDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.VideosBulkDeleteExecute(r)
}

/*
VideosBulkDelete Remove the multiple video.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ids The Id's of video's
 @return ApiVideosBulkDeleteRequest
*/
func (a *VideoAPIService) VideosBulkDelete(ctx context.Context, ids []string) ApiVideosBulkDeleteRequest {
	return ApiVideosBulkDeleteRequest{
		ApiService: a,
		ctx: ctx,
		ids: ids,
	}
}

// Execute executes the request
func (a *VideoAPIService) VideosBulkDeleteExecute(r ApiVideosBulkDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.VideosBulkDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/bulk"
	localVarPath = strings.Replace(localVarPath, "{"+"ids"+"}", url.PathEscape(parameterValueToString(r.ids, "ids")), -1)

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

type ApiVideosVideoDeleteRequest struct {
	ctx context.Context
	ApiService *VideoAPIService
	video string
}

func (r ApiVideosVideoDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.VideosVideoDeleteExecute(r)
}

/*
VideosVideoDelete Remove the specified video.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param video The Id of video
 @return ApiVideosVideoDeleteRequest
*/
func (a *VideoAPIService) VideosVideoDelete(ctx context.Context, video string) ApiVideosVideoDeleteRequest {
	return ApiVideosVideoDeleteRequest{
		ApiService: a,
		ctx: ctx,
		video: video,
	}
}

// Execute executes the request
func (a *VideoAPIService) VideosVideoDeleteExecute(r ApiVideosVideoDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.VideosVideoDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video}"
	localVarPath = strings.Replace(localVarPath, "{"+"video"+"}", url.PathEscape(parameterValueToString(r.video, "video")), -1)

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

type ApiVideosVideoGetRequest struct {
	ctx context.Context
	ApiService *VideoAPIService
	video string
	secureIp *string
	secureExpireTime *int32
}

// The IP address for generate secure links for. If channel is secure default is request IP
func (r ApiVideosVideoGetRequest) SecureIp(secureIp string) ApiVideosVideoGetRequest {
	r.secureIp = &secureIp
	return r
}

// The Unix Timestamp for expire secure links.      *          If channel is secure default is 24 hours later from now
func (r ApiVideosVideoGetRequest) SecureExpireTime(secureExpireTime int32) ApiVideosVideoGetRequest {
	r.secureExpireTime = &secureExpireTime
	return r
}

func (r ApiVideosVideoGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.VideosVideoGetExecute(r)
}

/*
VideosVideoGet Return the specified video.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param video The Id of video
 @return ApiVideosVideoGetRequest
*/
func (a *VideoAPIService) VideosVideoGet(ctx context.Context, video string) ApiVideosVideoGetRequest {
	return ApiVideosVideoGetRequest{
		ApiService: a,
		ctx: ctx,
		video: video,
	}
}

// Execute executes the request
func (a *VideoAPIService) VideosVideoGetExecute(r ApiVideosVideoGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.VideosVideoGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video}"
	localVarPath = strings.Replace(localVarPath, "{"+"video"+"}", url.PathEscape(parameterValueToString(r.video, "video")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiVideosVideoPatchRequest struct {
	ctx context.Context
	ApiService *VideoAPIService
	video string
	body *VideosVideoPatchRequest
}

// Video details
func (r ApiVideosVideoPatchRequest) Body(body VideosVideoPatchRequest) ApiVideosVideoPatchRequest {
	r.body = &body
	return r
}

func (r ApiVideosVideoPatchRequest) Execute() (*http.Response, error) {
	return r.ApiService.VideosVideoPatchExecute(r)
}

/*
VideosVideoPatch Update the specified video.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param video The Id of video
 @return ApiVideosVideoPatchRequest
*/
func (a *VideoAPIService) VideosVideoPatch(ctx context.Context, video string) ApiVideosVideoPatchRequest {
	return ApiVideosVideoPatchRequest{
		ApiService: a,
		ctx: ctx,
		video: video,
	}
}

// Execute executes the request
func (a *VideoAPIService) VideosVideoPatchExecute(r ApiVideosVideoPatchRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.VideosVideoPatch")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video}"
	localVarPath = strings.Replace(localVarPath, "{"+"video"+"}", url.PathEscape(parameterValueToString(r.video, "video")), -1)

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
