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


// FileAPIService FileAPI service
type FileAPIService service

type ApiChannelsChannelFilesFileHeadRequest struct {
	ctx context.Context
	ApiService *FileAPIService
	channel string
	file string
}

func (r ApiChannelsChannelFilesFileHeadRequest) Execute() (*http.Response, error) {
	return r.ApiService.ChannelsChannelFilesFileHeadExecute(r)
}

/*
ChannelsChannelFilesFileHead Get upload offset. See https://tus.io/ for more detail.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channel The Id of channel
 @param file The Id of file
 @return ApiChannelsChannelFilesFileHeadRequest
*/
func (a *FileAPIService) ChannelsChannelFilesFileHead(ctx context.Context, channel string, file string) ApiChannelsChannelFilesFileHeadRequest {
	return ApiChannelsChannelFilesFileHeadRequest{
		ApiService: a,
		ctx: ctx,
		channel: channel,
		file: file,
	}
}

// Execute executes the request
func (a *FileAPIService) ChannelsChannelFilesFileHeadExecute(r ApiChannelsChannelFilesFileHeadRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodHead
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FileAPIService.ChannelsChannelFilesFileHead")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel}/files/{file}"
	localVarPath = strings.Replace(localVarPath, "{"+"channel"+"}", url.PathEscape(parameterValueToString(r.channel, "channel")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"file"+"}", url.PathEscape(parameterValueToString(r.file, "file")), -1)

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

type ApiChannelsChannelFilesFilePatchRequest struct {
	ctx context.Context
	ApiService *FileAPIService
	channel string
	file string
	tusResumable *string
	uploadOffset *int32
	contentType *string
}

// version of tus.io
func (r ApiChannelsChannelFilesFilePatchRequest) TusResumable(tusResumable string) ApiChannelsChannelFilesFilePatchRequest {
	r.tusResumable = &tusResumable
	return r
}

// request and response header indicates a byte offset within a resource.      * For uploading entire file in one request, set this to &#39;0&#39;
func (r ApiChannelsChannelFilesFilePatchRequest) UploadOffset(uploadOffset int32) ApiChannelsChannelFilesFilePatchRequest {
	r.uploadOffset = &uploadOffset
	return r
}

// Request content type
func (r ApiChannelsChannelFilesFilePatchRequest) ContentType(contentType string) ApiChannelsChannelFilesFilePatchRequest {
	r.contentType = &contentType
	return r
}

func (r ApiChannelsChannelFilesFilePatchRequest) Execute() (*http.Response, error) {
	return r.ApiService.ChannelsChannelFilesFilePatchExecute(r)
}

/*
ChannelsChannelFilesFilePatch Upload and apply bytes to a file. See https://tus.io/ for more detail.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channel The Id of channel
 @param file The Id of file
 @return ApiChannelsChannelFilesFilePatchRequest
*/
func (a *FileAPIService) ChannelsChannelFilesFilePatch(ctx context.Context, channel string, file string) ApiChannelsChannelFilesFilePatchRequest {
	return ApiChannelsChannelFilesFilePatchRequest{
		ApiService: a,
		ctx: ctx,
		channel: channel,
		file: file,
	}
}

// Execute executes the request
func (a *FileAPIService) ChannelsChannelFilesFilePatchExecute(r ApiChannelsChannelFilesFilePatchRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FileAPIService.ChannelsChannelFilesFilePatch")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel}/files/{file}"
	localVarPath = strings.Replace(localVarPath, "{"+"channel"+"}", url.PathEscape(parameterValueToString(r.channel, "channel")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"file"+"}", url.PathEscape(parameterValueToString(r.file, "file")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tusResumable == nil {
		return nil, reportError("tusResumable is required and must be specified")
	}
	if r.uploadOffset == nil {
		return nil, reportError("uploadOffset is required and must be specified")
	}
	if r.contentType == nil {
		return nil, reportError("contentType is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "tus-resumable", r.tusResumable, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "upload-offset", r.uploadOffset, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
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

type ApiChannelsChannelFilesGetRequest struct {
	ctx context.Context
	ApiService *FileAPIService
	channel string
	filter *string
}

// Filter result
func (r ApiChannelsChannelFilesGetRequest) Filter(filter string) ApiChannelsChannelFilesGetRequest {
	r.filter = &filter
	return r
}

func (r ApiChannelsChannelFilesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ChannelsChannelFilesGetExecute(r)
}

/*
ChannelsChannelFilesGet Return all draft files of channel.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channel The Id of channel
 @return ApiChannelsChannelFilesGetRequest
*/
func (a *FileAPIService) ChannelsChannelFilesGet(ctx context.Context, channel string) ApiChannelsChannelFilesGetRequest {
	return ApiChannelsChannelFilesGetRequest{
		ApiService: a,
		ctx: ctx,
		channel: channel,
	}
}

// Execute executes the request
func (a *FileAPIService) ChannelsChannelFilesGetExecute(r ApiChannelsChannelFilesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FileAPIService.ChannelsChannelFilesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel}/files"
	localVarPath = strings.Replace(localVarPath, "{"+"channel"+"}", url.PathEscape(parameterValueToString(r.channel, "channel")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
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

type ApiChannelsChannelFilesPostRequest struct {
	ctx context.Context
	ApiService *FileAPIService
	channel string
	tusResumable *string
	uploadLength *int32
	uploadMetadata *string
}

// version of tus.io
func (r ApiChannelsChannelFilesPostRequest) TusResumable(tusResumable string) ApiChannelsChannelFilesPostRequest {
	r.tusResumable = &tusResumable
	return r
}

// To indicate the size of entire upload in bytes
func (r ApiChannelsChannelFilesPostRequest) UploadLength(uploadLength int32) ApiChannelsChannelFilesPostRequest {
	r.uploadLength = &uploadLength
	return r
}

// To add additional metadata to the upload creation request.      * MUST contain &#39;filename&#39; and &#39;filetype&#39;. From all available fields only these two fields will be used.      * MUST consist of one or more comma-separated key-value pairs.      * The key and value MUST be separated by a space.      * The key MUST NOT contain spaces and commas and MUST NOT be empty.      * The key SHOULD be ASCII encoded and the value MUST be Base64 encoded.
func (r ApiChannelsChannelFilesPostRequest) UploadMetadata(uploadMetadata string) ApiChannelsChannelFilesPostRequest {
	r.uploadMetadata = &uploadMetadata
	return r
}

func (r ApiChannelsChannelFilesPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ChannelsChannelFilesPostExecute(r)
}

/*
ChannelsChannelFilesPost Request a new upload file. See https://tus.io/ for more detail.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channel The Id of channel
 @return ApiChannelsChannelFilesPostRequest
*/
func (a *FileAPIService) ChannelsChannelFilesPost(ctx context.Context, channel string) ApiChannelsChannelFilesPostRequest {
	return ApiChannelsChannelFilesPostRequest{
		ApiService: a,
		ctx: ctx,
		channel: channel,
	}
}

// Execute executes the request
func (a *FileAPIService) ChannelsChannelFilesPostExecute(r ApiChannelsChannelFilesPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FileAPIService.ChannelsChannelFilesPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/channels/{channel}/files"
	localVarPath = strings.Replace(localVarPath, "{"+"channel"+"}", url.PathEscape(parameterValueToString(r.channel, "channel")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tusResumable == nil {
		return nil, reportError("tusResumable is required and must be specified")
	}
	if r.uploadLength == nil {
		return nil, reportError("uploadLength is required and must be specified")
	}
	if r.uploadMetadata == nil {
		return nil, reportError("uploadMetadata is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "tus-resumable", r.tusResumable, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "upload-length", r.uploadLength, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "upload-metadata", r.uploadMetadata, "")
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

type ApiFilesFileDeleteRequest struct {
	ctx context.Context
	ApiService *FileAPIService
	file string
}

func (r ApiFilesFileDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.FilesFileDeleteExecute(r)
}

/*
FilesFileDelete Remove the specified file.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param file The Id of file
 @return ApiFilesFileDeleteRequest
*/
func (a *FileAPIService) FilesFileDelete(ctx context.Context, file string) ApiFilesFileDeleteRequest {
	return ApiFilesFileDeleteRequest{
		ApiService: a,
		ctx: ctx,
		file: file,
	}
}

// Execute executes the request
func (a *FileAPIService) FilesFileDeleteExecute(r ApiFilesFileDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FileAPIService.FilesFileDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/files/{file}"
	localVarPath = strings.Replace(localVarPath, "{"+"file"+"}", url.PathEscape(parameterValueToString(r.file, "file")), -1)

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

type ApiFilesFileGetRequest struct {
	ctx context.Context
	ApiService *FileAPIService
	file string
}

func (r ApiFilesFileGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.FilesFileGetExecute(r)
}

/*
FilesFileGet Return the specified file.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param file The Id of file
 @return ApiFilesFileGetRequest
*/
func (a *FileAPIService) FilesFileGet(ctx context.Context, file string) ApiFilesFileGetRequest {
	return ApiFilesFileGetRequest{
		ApiService: a,
		ctx: ctx,
		file: file,
	}
}

// Execute executes the request
func (a *FileAPIService) FilesFileGetExecute(r ApiFilesFileGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FileAPIService.FilesFileGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/files/{file}"
	localVarPath = strings.Replace(localVarPath, "{"+"File"+"}", url.PathEscape(parameterValueToString(r.file, "file")), -1)

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
