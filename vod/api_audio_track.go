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


// AudioTrackAPIService AudioTrackAPI service
type AudioTrackAPIService service

type ApiAudioTracksAudioTrackDeleteRequest struct {
	ctx context.Context
	ApiService *AudioTrackAPIService
	audioTrack string
}

func (r ApiAudioTracksAudioTrackDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.AudioTracksAudioTrackDeleteExecute(r)
}

/*
AudioTracksAudioTrackDelete Remove the specified audio track.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param audioTrack The Id of audio track
 @return ApiAudioTracksAudioTrackDeleteRequest
*/
func (a *AudioTrackAPIService) AudioTracksAudioTrackDelete(ctx context.Context, audioTrack string) ApiAudioTracksAudioTrackDeleteRequest {
	return ApiAudioTracksAudioTrackDeleteRequest{
		ApiService: a,
		ctx: ctx,
		audioTrack: audioTrack,
	}
}

// Execute executes the request
func (a *AudioTrackAPIService) AudioTracksAudioTrackDeleteExecute(r ApiAudioTracksAudioTrackDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AudioTrackAPIService.AudioTracksAudioTrackDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/audio-tracks/{audio_track}"
	localVarPath = strings.Replace(localVarPath, "{"+"audio_track"+"}", url.PathEscape(parameterValueToString(r.audioTrack, "audioTrack")), -1)

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

type ApiAudioTracksAudioTrackGetRequest struct {
	ctx context.Context
	ApiService *AudioTrackAPIService
	audioTrack string
	secureIp *string
	secureExpireTime *int32
}

// The IP address for generate secure links for. If channel is secure default is request IP
func (r ApiAudioTracksAudioTrackGetRequest) SecureIp(secureIp string) ApiAudioTracksAudioTrackGetRequest {
	r.secureIp = &secureIp
	return r
}

// The Unix Timestamp for expire secure links.      *          If channel is secure default is 24 hours later from now
func (r ApiAudioTracksAudioTrackGetRequest) SecureExpireTime(secureExpireTime int32) ApiAudioTracksAudioTrackGetRequest {
	r.secureExpireTime = &secureExpireTime
	return r
}

func (r ApiAudioTracksAudioTrackGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.AudioTracksAudioTrackGetExecute(r)
}

/*
AudioTracksAudioTrackGet Return the specified audio track.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param audioTrack The Id of audio track
 @return ApiAudioTracksAudioTrackGetRequest
*/
func (a *AudioTrackAPIService) AudioTracksAudioTrackGet(ctx context.Context, audioTrack string) ApiAudioTracksAudioTrackGetRequest {
	return ApiAudioTracksAudioTrackGetRequest{
		ApiService: a,
		ctx: ctx,
		audioTrack: audioTrack,
	}
}

// Execute executes the request
func (a *AudioTrackAPIService) AudioTracksAudioTrackGetExecute(r ApiAudioTracksAudioTrackGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AudioTrackAPIService.AudioTracksAudioTrackGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/audio-tracks/{audio_track}"
	localVarPath = strings.Replace(localVarPath, "{"+"audio_track"+"}", url.PathEscape(parameterValueToString(r.audioTrack, "audioTrack")), -1)

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

type ApiVideosVideoAudioTracksGetRequest struct {
	ctx context.Context
	ApiService *AudioTrackAPIService
	video string
	secureIp *string
	secureExpireTime *int32
}

// The IP address for generate secure links for. If channel is secure default is request IP
func (r ApiVideosVideoAudioTracksGetRequest) SecureIp(secureIp string) ApiVideosVideoAudioTracksGetRequest {
	r.secureIp = &secureIp
	return r
}

// The Unix Timestamp for expire secure links.      *          If channel is secure default is 24 hours later from now
func (r ApiVideosVideoAudioTracksGetRequest) SecureExpireTime(secureExpireTime int32) ApiVideosVideoAudioTracksGetRequest {
	r.secureExpireTime = &secureExpireTime
	return r
}

func (r ApiVideosVideoAudioTracksGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.VideosVideoAudioTracksGetExecute(r)
}

/*
VideosVideoAudioTracksGet Display a listing of the Audio Tracks.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param video The Id of video
 @return ApiVideosVideoAudioTracksGetRequest
*/
func (a *AudioTrackAPIService) VideosVideoAudioTracksGet(ctx context.Context, video string) ApiVideosVideoAudioTracksGetRequest {
	return ApiVideosVideoAudioTracksGetRequest{
		ApiService: a,
		ctx: ctx,
		video: video,
	}
}

// Execute executes the request
func (a *AudioTrackAPIService) VideosVideoAudioTracksGetExecute(r ApiVideosVideoAudioTracksGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AudioTrackAPIService.VideosVideoAudioTracksGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video}/audio-tracks"
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

type ApiVideosVideoAudioTracksPostRequest struct {
	ctx context.Context
	ApiService *AudioTrackAPIService
	video string
	lang *string
	audioTrack *os.File
}

// Track language
func (r ApiVideosVideoAudioTracksPostRequest) Lang(lang string) ApiVideosVideoAudioTracksPostRequest {
	r.lang = &lang
	return r
}

// The Mp3 or ACC Audio file.
func (r ApiVideosVideoAudioTracksPostRequest) AudioTrack(audioTrack *os.File) ApiVideosVideoAudioTracksPostRequest {
	r.audioTrack = audioTrack
	return r
}

func (r ApiVideosVideoAudioTracksPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.VideosVideoAudioTracksPostExecute(r)
}

/*
VideosVideoAudioTracksPost Store a newly created audio track.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param video The Id of video
 @return ApiVideosVideoAudioTracksPostRequest
*/
func (a *AudioTrackAPIService) VideosVideoAudioTracksPost(ctx context.Context, video string) ApiVideosVideoAudioTracksPostRequest {
	return ApiVideosVideoAudioTracksPostRequest{
		ApiService: a,
		ctx: ctx,
		video: video,
	}
}

// Execute executes the request
func (a *AudioTrackAPIService) VideosVideoAudioTracksPostExecute(r ApiVideosVideoAudioTracksPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AudioTrackAPIService.VideosVideoAudioTracksPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/videos/{video}/audio-tracks"
	localVarPath = strings.Replace(localVarPath, "{"+"video"+"}", url.PathEscape(parameterValueToString(r.video, "video")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lang == nil {
		return nil, reportError("lang is required and must be specified")
	}
	if r.audioTrack == nil {
		return nil, reportError("audioTrack is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "lang", r.lang, "")
	var audioTrackLocalVarFormFileName string
	var audioTrackLocalVarFileName     string
	var audioTrackLocalVarFileBytes    []byte

	audioTrackLocalVarFormFileName = "audio_track"


	audioTrackLocalVarFile := r.audioTrack

	if audioTrackLocalVarFile != nil {
		fbs, _ := io.ReadAll(audioTrackLocalVarFile)

		audioTrackLocalVarFileBytes = fbs
		audioTrackLocalVarFileName = audioTrackLocalVarFile.Name()
		audioTrackLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: audioTrackLocalVarFileBytes, fileName: audioTrackLocalVarFileName, formFileName: audioTrackLocalVarFormFileName})
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