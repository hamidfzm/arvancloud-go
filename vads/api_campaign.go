/*
ArvanCloud Video Advertising Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vads

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// CampaignAPIService CampaignAPI service
type CampaignAPIService service

type ApiCampaignsCampaignDeleteRequest struct {
	ctx context.Context
	ApiService *CampaignAPIService
	campaign string
}

func (r ApiCampaignsCampaignDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.CampaignsCampaignDeleteExecute(r)
}

/*
CampaignsCampaignDelete Remove the specified campaign.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param campaign The Id of campaign
 @return ApiCampaignsCampaignDeleteRequest
*/
func (a *CampaignAPIService) CampaignsCampaignDelete(ctx context.Context, campaign string) ApiCampaignsCampaignDeleteRequest {
	return ApiCampaignsCampaignDeleteRequest{
		ApiService: a,
		ctx: ctx,
		campaign: campaign,
	}
}

// Execute executes the request
func (a *CampaignAPIService) CampaignsCampaignDeleteExecute(r ApiCampaignsCampaignDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignAPIService.CampaignsCampaignDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/campaigns/{campaign}"
	localVarPath = strings.Replace(localVarPath, "{"+"campaign"+"}", url.PathEscape(parameterValueToString(r.campaign, "campaign")), -1)

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

type ApiCampaignsCampaignGetRequest struct {
	ctx context.Context
	ApiService *CampaignAPIService
	campaign string
}

func (r ApiCampaignsCampaignGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.CampaignsCampaignGetExecute(r)
}

/*
CampaignsCampaignGet Return the specified campaign.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param campaign The Id of campaign
 @return ApiCampaignsCampaignGetRequest
*/
func (a *CampaignAPIService) CampaignsCampaignGet(ctx context.Context, campaign string) ApiCampaignsCampaignGetRequest {
	return ApiCampaignsCampaignGetRequest{
		ApiService: a,
		ctx: ctx,
		campaign: campaign,
	}
}

// Execute executes the request
func (a *CampaignAPIService) CampaignsCampaignGetExecute(r ApiCampaignsCampaignGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignAPIService.CampaignsCampaignGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/campaigns/{campaign}"
	localVarPath = strings.Replace(localVarPath, "{"+"campaign"+"}", url.PathEscape(parameterValueToString(r.campaign, "campaign")), -1)

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

type ApiCampaignsCampaignPutRequest struct {
	ctx context.Context
	ApiService *CampaignAPIService
	campaign string
	body *CampaignsCampaignPutRequest
}

// Ad&#39;s ad details
func (r ApiCampaignsCampaignPutRequest) Body(body CampaignsCampaignPutRequest) ApiCampaignsCampaignPutRequest {
	r.body = &body
	return r
}

func (r ApiCampaignsCampaignPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.CampaignsCampaignPutExecute(r)
}

/*
CampaignsCampaignPut Update the specified campaign.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param campaign The Id of campaign
 @return ApiCampaignsCampaignPutRequest
*/
func (a *CampaignAPIService) CampaignsCampaignPut(ctx context.Context, campaign string) ApiCampaignsCampaignPutRequest {
	return ApiCampaignsCampaignPutRequest{
		ApiService: a,
		ctx: ctx,
		campaign: campaign,
	}
}

// Execute executes the request
func (a *CampaignAPIService) CampaignsCampaignPutExecute(r ApiCampaignsCampaignPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignAPIService.CampaignsCampaignPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/campaigns/{campaign}"
	localVarPath = strings.Replace(localVarPath, "{"+"campaign"+"}", url.PathEscape(parameterValueToString(r.campaign, "campaign")), -1)

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

type ApiChannelsChannelCampaignsGetRequest struct {
	ctx context.Context
	ApiService *CampaignAPIService
	channel string
	page *int32
	perPage *int32
}

// Page number
func (r ApiChannelsChannelCampaignsGetRequest) Page(page int32) ApiChannelsChannelCampaignsGetRequest {
	r.page = &page
	return r
}

// Page limit for query
func (r ApiChannelsChannelCampaignsGetRequest) PerPage(perPage int32) ApiChannelsChannelCampaignsGetRequest {
	r.perPage = &perPage
	return r
}

func (r ApiChannelsChannelCampaignsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ChannelsChannelCampaignsGetExecute(r)
}

/*
ChannelsChannelCampaignsGet Return all channel's campaigns.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channel The Id of channel
 @return ApiChannelsChannelCampaignsGetRequest
*/
func (a *CampaignAPIService) ChannelsChannelCampaignsGet(ctx context.Context, channel string) ApiChannelsChannelCampaignsGetRequest {
	return ApiChannelsChannelCampaignsGetRequest{
		ApiService: a,
		ctx: ctx,
		channel: channel,
	}
}

// Execute executes the request
func (a *CampaignAPIService) ChannelsChannelCampaignsGetExecute(r ApiChannelsChannelCampaignsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignAPIService.ChannelsChannelCampaignsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel}/campaigns"
	localVarPath = strings.Replace(localVarPath, "{"+"channel"+"}", url.PathEscape(parameterValueToString(r.channel, "channel")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiChannelsChannelCampaignsPostRequest struct {
	ctx context.Context
	ApiService *CampaignAPIService
	channel string
	body *ChannelsChannelCampaignsPostRequest
}

// Ad&#39;s ad details
func (r ApiChannelsChannelCampaignsPostRequest) Body(body ChannelsChannelCampaignsPostRequest) ApiChannelsChannelCampaignsPostRequest {
	r.body = &body
	return r
}

func (r ApiChannelsChannelCampaignsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ChannelsChannelCampaignsPostExecute(r)
}

/*
ChannelsChannelCampaignsPost Store a newly campaign for specific channel.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channel The Id of channel
 @return ApiChannelsChannelCampaignsPostRequest
*/
func (a *CampaignAPIService) ChannelsChannelCampaignsPost(ctx context.Context, channel string) ApiChannelsChannelCampaignsPostRequest {
	return ApiChannelsChannelCampaignsPostRequest{
		ApiService: a,
		ctx: ctx,
		channel: channel,
	}
}

// Execute executes the request
func (a *CampaignAPIService) ChannelsChannelCampaignsPostExecute(r ApiChannelsChannelCampaignsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CampaignAPIService.ChannelsChannelCampaignsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel}/campaigns"
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
