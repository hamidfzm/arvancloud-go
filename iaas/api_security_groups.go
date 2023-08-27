/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// SecurityGroupsAPIService SecurityGroupsAPI service
type SecurityGroupsAPIService service

type SecurityGroupsAPICreateSecurityGroupRequest struct {
	ctx context.Context
	ApiService *SecurityGroupsAPIService
	region string
	createSecurityGroupRequest *CreateSecurityGroupRequest
}

// Info about the security group to be created
func (r SecurityGroupsAPICreateSecurityGroupRequest) CreateSecurityGroupRequest(createSecurityGroupRequest CreateSecurityGroupRequest) SecurityGroupsAPICreateSecurityGroupRequest {
	r.createSecurityGroupRequest = &createSecurityGroupRequest
	return r
}

func (r SecurityGroupsAPICreateSecurityGroupRequest) Execute() (*SecurityGroup, *http.Response, error) {
	return r.ApiService.CreateSecurityGroupExecute(r)
}

/*
CreateSecurityGroup Method for CreateSecurityGroup

Creates a security group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param region region code
 @return SecurityGroupsAPICreateSecurityGroupRequest
*/
func (a *SecurityGroupsAPIService) CreateSecurityGroup(ctx context.Context, region string) SecurityGroupsAPICreateSecurityGroupRequest {
	return SecurityGroupsAPICreateSecurityGroupRequest{
		ApiService: a,
		ctx: ctx,
		region: region,
	}
}

// Execute executes the request
//  @return SecurityGroup
func (a *SecurityGroupsAPIService) CreateSecurityGroupExecute(r SecurityGroupsAPICreateSecurityGroupRequest) (*SecurityGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SecurityGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityGroupsAPIService.CreateSecurityGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/regions/:region/securities"
	localVarPath = strings.Replace(localVarPath, "{"+"region"+"}", url.PathEscape(parameterValueToString(r.region, "region")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createSecurityGroupRequest == nil {
		return localVarReturnValue, nil, reportError("createSecurityGroupRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createSecurityGroupRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Apikey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SecurityGroupsAPICreateSecurityGroupRuleRequest struct {
	ctx context.Context
	ApiService *SecurityGroupsAPIService
	region string
	id string
	createSecurityGroupRuleRequest *CreateSecurityGroupRuleRequest
}

// Security group rule info
func (r SecurityGroupsAPICreateSecurityGroupRuleRequest) CreateSecurityGroupRuleRequest(createSecurityGroupRuleRequest CreateSecurityGroupRuleRequest) SecurityGroupsAPICreateSecurityGroupRuleRequest {
	r.createSecurityGroupRuleRequest = &createSecurityGroupRuleRequest
	return r
}

func (r SecurityGroupsAPICreateSecurityGroupRuleRequest) Execute() (*MessageResponse, *http.Response, error) {
	return r.ApiService.CreateSecurityGroupRuleExecute(r)
}

/*
CreateSecurityGroupRule Method for CreateSecurityGroupRule

Creates a rule for the specified security group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param region region code
 @param id Security group id
 @return SecurityGroupsAPICreateSecurityGroupRuleRequest
*/
func (a *SecurityGroupsAPIService) CreateSecurityGroupRule(ctx context.Context, region string, id string) SecurityGroupsAPICreateSecurityGroupRuleRequest {
	return SecurityGroupsAPICreateSecurityGroupRuleRequest{
		ApiService: a,
		ctx: ctx,
		region: region,
		id: id,
	}
}

// Execute executes the request
//  @return MessageResponse
func (a *SecurityGroupsAPIService) CreateSecurityGroupRuleExecute(r SecurityGroupsAPICreateSecurityGroupRuleRequest) (*MessageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MessageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityGroupsAPIService.CreateSecurityGroupRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/regions/:region/securities/securitiy-rules/:id"
	localVarPath = strings.Replace(localVarPath, "{"+"region"+"}", url.PathEscape(parameterValueToString(r.region, "region")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createSecurityGroupRuleRequest == nil {
		return localVarReturnValue, nil, reportError("createSecurityGroupRuleRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createSecurityGroupRuleRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Apikey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SecurityGroupsAPIDeleteSecurityGroupRequest struct {
	ctx context.Context
	ApiService *SecurityGroupsAPIService
	region string
}

func (r SecurityGroupsAPIDeleteSecurityGroupRequest) Execute() (*MessageResponse, *http.Response, error) {
	return r.ApiService.DeleteSecurityGroupExecute(r)
}

/*
DeleteSecurityGroup Method for DeleteSecurityGroup

Deletes a security group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param region region code
 @return SecurityGroupsAPIDeleteSecurityGroupRequest
*/
func (a *SecurityGroupsAPIService) DeleteSecurityGroup(ctx context.Context, region string) SecurityGroupsAPIDeleteSecurityGroupRequest {
	return SecurityGroupsAPIDeleteSecurityGroupRequest{
		ApiService: a,
		ctx: ctx,
		region: region,
	}
}

// Execute executes the request
//  @return MessageResponse
func (a *SecurityGroupsAPIService) DeleteSecurityGroupExecute(r SecurityGroupsAPIDeleteSecurityGroupRequest) (*MessageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MessageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityGroupsAPIService.DeleteSecurityGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/regions/:region/securities/:id"
	localVarPath = strings.Replace(localVarPath, "{"+"region"+"}", url.PathEscape(parameterValueToString(r.region, "region")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Apikey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SecurityGroupsAPIDeleteSecurityGroupRuleRequest struct {
	ctx context.Context
	ApiService *SecurityGroupsAPIService
	region string
}

func (r SecurityGroupsAPIDeleteSecurityGroupRuleRequest) Execute() (*SecurityGroup, *http.Response, error) {
	return r.ApiService.DeleteSecurityGroupRuleExecute(r)
}

/*
DeleteSecurityGroupRule Method for DeleteSecurityGroupRule

Deletes a security group rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param region region code
 @return SecurityGroupsAPIDeleteSecurityGroupRuleRequest
*/
func (a *SecurityGroupsAPIService) DeleteSecurityGroupRule(ctx context.Context, region string) SecurityGroupsAPIDeleteSecurityGroupRuleRequest {
	return SecurityGroupsAPIDeleteSecurityGroupRuleRequest{
		ApiService: a,
		ctx: ctx,
		region: region,
	}
}

// Execute executes the request
//  @return SecurityGroup
func (a *SecurityGroupsAPIService) DeleteSecurityGroupRuleExecute(r SecurityGroupsAPIDeleteSecurityGroupRuleRequest) (*SecurityGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SecurityGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityGroupsAPIService.DeleteSecurityGroupRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/regions/:region/securities/securitiy-rules/:id"
	localVarPath = strings.Replace(localVarPath, "{"+"region"+"}", url.PathEscape(parameterValueToString(r.region, "region")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Apikey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SecurityGroupsAPIGetAllSecurityGroupsRequest struct {
	ctx context.Context
	ApiService *SecurityGroupsAPIService
	region string
}

func (r SecurityGroupsAPIGetAllSecurityGroupsRequest) Execute() (*SecurityGroup, *http.Response, error) {
	return r.ApiService.GetAllSecurityGroupsExecute(r)
}

/*
GetAllSecurityGroups Method for GetAllSecurityGroups

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param region region code
 @return SecurityGroupsAPIGetAllSecurityGroupsRequest
*/
func (a *SecurityGroupsAPIService) GetAllSecurityGroups(ctx context.Context, region string) SecurityGroupsAPIGetAllSecurityGroupsRequest {
	return SecurityGroupsAPIGetAllSecurityGroupsRequest{
		ApiService: a,
		ctx: ctx,
		region: region,
	}
}

// Execute executes the request
//  @return SecurityGroup
func (a *SecurityGroupsAPIService) GetAllSecurityGroupsExecute(r SecurityGroupsAPIGetAllSecurityGroupsRequest) (*SecurityGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SecurityGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityGroupsAPIService.GetAllSecurityGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/regions/:region/securities"
	localVarPath = strings.Replace(localVarPath, "{"+"region"+"}", url.PathEscape(parameterValueToString(r.region, "region")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Apikey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SecurityGroupsAPIGetSecurityGroupRulesRequest struct {
	ctx context.Context
	ApiService *SecurityGroupsAPIService
	region string
}

func (r SecurityGroupsAPIGetSecurityGroupRulesRequest) Execute() (*SecurityGroup, *http.Response, error) {
	return r.ApiService.GetSecurityGroupRulesExecute(r)
}

/*
GetSecurityGroupRules Method for GetSecurityGroupRules

Returns rules for a security group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param region region code
 @return SecurityGroupsAPIGetSecurityGroupRulesRequest
*/
func (a *SecurityGroupsAPIService) GetSecurityGroupRules(ctx context.Context, region string) SecurityGroupsAPIGetSecurityGroupRulesRequest {
	return SecurityGroupsAPIGetSecurityGroupRulesRequest{
		ApiService: a,
		ctx: ctx,
		region: region,
	}
}

// Execute executes the request
//  @return SecurityGroup
func (a *SecurityGroupsAPIService) GetSecurityGroupRulesExecute(r SecurityGroupsAPIGetSecurityGroupRulesRequest) (*SecurityGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SecurityGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityGroupsAPIService.GetSecurityGroupRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/regions/:region/securities/security-rules/:id"
	localVarPath = strings.Replace(localVarPath, "{"+"region"+"}", url.PathEscape(parameterValueToString(r.region, "region")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Apikey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SecurityGroupsAPIImportCDNSecurityGroupRequest struct {
	ctx context.Context
	ApiService *SecurityGroupsAPIService
	region string
}

func (r SecurityGroupsAPIImportCDNSecurityGroupRequest) Execute() (*MessageResponse, *http.Response, error) {
	return r.ApiService.ImportCDNSecurityGroupExecute(r)
}

/*
ImportCDNSecurityGroup Method for ImportCDNSecurityGroup

Creates a security group based on Arvan CDN. This allows CDN IPs in security rules of newly created group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param region region code
 @return SecurityGroupsAPIImportCDNSecurityGroupRequest
*/
func (a *SecurityGroupsAPIService) ImportCDNSecurityGroup(ctx context.Context, region string) SecurityGroupsAPIImportCDNSecurityGroupRequest {
	return SecurityGroupsAPIImportCDNSecurityGroupRequest{
		ApiService: a,
		ctx: ctx,
		region: region,
	}
}

// Execute executes the request
//  @return MessageResponse
func (a *SecurityGroupsAPIService) ImportCDNSecurityGroupExecute(r SecurityGroupsAPIImportCDNSecurityGroupRequest) (*MessageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MessageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityGroupsAPIService.ImportCDNSecurityGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/regions/:region/securities/securitiy-rules/cdn"
	localVarPath = strings.Replace(localVarPath, "{"+"region"+"}", url.PathEscape(parameterValueToString(r.region, "region")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Apikey"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
