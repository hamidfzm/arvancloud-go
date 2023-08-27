/*
ArvanCloud S3 Services

<p/>

API version: 2006-03-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// BucketVersioningAPIService BucketVersioningAPI service
type BucketVersioningAPIService service

type ApiGetBucketVersioningRequest struct {
	ctx context.Context
	ApiService *BucketVersioningAPIService
	bucket string
	versioning *bool
	xAmzSecurityToken *string
	xAmzExpectedBucketOwner *string
}

func (r ApiGetBucketVersioningRequest) Versioning(versioning bool) ApiGetBucketVersioningRequest {
	r.versioning = &versioning
	return r
}

// This parameter is currently not supported and is not required.
func (r ApiGetBucketVersioningRequest) XAmzSecurityToken(xAmzSecurityToken string) ApiGetBucketVersioningRequest {
	r.xAmzSecurityToken = &xAmzSecurityToken
	return r
}

// The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error.
func (r ApiGetBucketVersioningRequest) XAmzExpectedBucketOwner(xAmzExpectedBucketOwner string) ApiGetBucketVersioningRequest {
	r.xAmzExpectedBucketOwner = &xAmzExpectedBucketOwner
	return r
}

func (r ApiGetBucketVersioningRequest) Execute() (*GetBucketVersioningOutput, *http.Response, error) {
	return r.ApiService.GetBucketVersioningExecute(r)
}

/*
GetBucketVersioning Method for GetBucketVersioning

<p>Returns the versioning state of a bucket.</p> <p>To retrieve the versioning state of a bucket, you must be the bucket owner.</p> <p>This implementation also returns the MFA Delete status of the versioning state. If the MFA Delete status is <code>enabled</code>, the bucket owner must use an authentication device to change the versioning state of the bucket.</p>

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bucket The name of the bucket for which to get the versioning information.
 @return ApiGetBucketVersioningRequest
*/
func (a *BucketVersioningAPIService) GetBucketVersioning(ctx context.Context, bucket string) ApiGetBucketVersioningRequest {
	return ApiGetBucketVersioningRequest{
		ApiService: a,
		ctx: ctx,
		bucket: bucket,
	}
}

// Execute executes the request
//  @return GetBucketVersioningOutput
func (a *BucketVersioningAPIService) GetBucketVersioningExecute(r ApiGetBucketVersioningRequest) (*GetBucketVersioningOutput, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetBucketVersioningOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BucketVersioningAPIService.GetBucketVersioning")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{Bucket}#versioning"
	localVarPath = strings.Replace(localVarPath, "{"+"Bucket"+"}", url.PathEscape(parameterValueToString(r.bucket, "bucket")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.versioning == nil {
		return localVarReturnValue, nil, reportError("versioning is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "versioning", r.versioning, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xAmzSecurityToken != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-security-token", r.xAmzSecurityToken, "")
	}
	if r.xAmzExpectedBucketOwner != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-expected-bucket-owner", r.xAmzExpectedBucketOwner, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["hmac"]; ok {
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

type ApiPutBucketVersioningRequest struct {
	ctx context.Context
	ApiService *BucketVersioningAPIService
	bucket string
	versioning *bool
	putBucketVersioningRequest *PutBucketVersioningRequest
	xAmzSecurityToken *string
	contentMD5 *string
	xAmzMfa *string
	xAmzExpectedBucketOwner *string
}

func (r ApiPutBucketVersioningRequest) Versioning(versioning bool) ApiPutBucketVersioningRequest {
	r.versioning = &versioning
	return r
}

func (r ApiPutBucketVersioningRequest) PutBucketVersioningRequest(putBucketVersioningRequest PutBucketVersioningRequest) ApiPutBucketVersioningRequest {
	r.putBucketVersioningRequest = &putBucketVersioningRequest
	return r
}

// This parameter is currently not supported and is not required.
func (r ApiPutBucketVersioningRequest) XAmzSecurityToken(xAmzSecurityToken string) ApiPutBucketVersioningRequest {
	r.xAmzSecurityToken = &xAmzSecurityToken
	return r
}

// &lt;p&gt;&amp;gt;The base64-encoded 128-bit MD5 digest of the data. You must use this header as a message integrity check to verify that the request body was not corrupted in transit. For more information, see &lt;a href&#x3D;\&quot;http://www.ietf.org/rfc/rfc1864.txt\&quot;&gt;RFC 1864&lt;/a&gt;.&lt;/p&gt; &lt;p&gt;For requests made using the Amazon Web Services Command Line Interface (CLI) or Amazon Web Services SDKs, this field is calculated automatically.&lt;/p&gt;
func (r ApiPutBucketVersioningRequest) ContentMD5(contentMD5 string) ApiPutBucketVersioningRequest {
	r.contentMD5 = &contentMD5
	return r
}

// The concatenation of the authentication device&#39;s serial number, a space, and the value that is displayed on your authentication device.
func (r ApiPutBucketVersioningRequest) XAmzMfa(xAmzMfa string) ApiPutBucketVersioningRequest {
	r.xAmzMfa = &xAmzMfa
	return r
}

// The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error.
func (r ApiPutBucketVersioningRequest) XAmzExpectedBucketOwner(xAmzExpectedBucketOwner string) ApiPutBucketVersioningRequest {
	r.xAmzExpectedBucketOwner = &xAmzExpectedBucketOwner
	return r
}

func (r ApiPutBucketVersioningRequest) Execute() (*http.Response, error) {
	return r.ApiService.PutBucketVersioningExecute(r)
}

/*
PutBucketVersioning Method for PutBucketVersioning

<p>Sets the versioning state of an existing bucket. To set the versioning state, you must be the bucket owner.</p> <p>You can set the versioning state with one of the following values:</p> <p> <b>Enabled</b>—Enables versioning for the objects in the bucket. All objects added to the bucket receive a unique version ID.</p> <p> <b>Suspended</b>—Disables versioning for the objects in the bucket. All objects added to the bucket receive the version ID null.</p> <p>If the versioning state has never been set on a bucket, it has no versioning state; a <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketVersioning.html">GetBucketVersioning</a> request does not return a versioning state value.</p> <p>If the bucket owner enables MFA Delete in the bucket versioning configuration, the bucket owner must include the <code>x-amz-mfa request</code> header and the <code>Status</code> and the <code>MfaDelete</code> request elements in a request to set the versioning state of the bucket.</p> <important> <p>If you have an object expiration lifecycle policy in your non-versioned bucket and you want to maintain the same permanent delete behavior when you enable versioning, you must add a noncurrent expiration policy. The noncurrent expiration lifecycle policy will manage the deletes of the noncurrent object versions in the version-enabled bucket. (A version-enabled bucket maintains one current and zero or more noncurrent object versions.) For more information, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html#lifecycle-and-other-bucket-config">Lifecycle and Versioning</a>.</p> </important> <p class="title"> <b>Related Resources</b> </p> <ul> <li> <p> <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html">CreateBucket</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html">DeleteBucket</a> </p> </li> <li> <p> <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketVersioning.html">GetBucketVersioning</a> </p> </li> </ul>

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bucket The bucket name.
 @return ApiPutBucketVersioningRequest
*/
func (a *BucketVersioningAPIService) PutBucketVersioning(ctx context.Context, bucket string) ApiPutBucketVersioningRequest {
	return ApiPutBucketVersioningRequest{
		ApiService: a,
		ctx: ctx,
		bucket: bucket,
	}
}

// Execute executes the request
func (a *BucketVersioningAPIService) PutBucketVersioningExecute(r ApiPutBucketVersioningRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BucketVersioningAPIService.PutBucketVersioning")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{Bucket}#versioning"
	localVarPath = strings.Replace(localVarPath, "{"+"Bucket"+"}", url.PathEscape(parameterValueToString(r.bucket, "bucket")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.versioning == nil {
		return nil, reportError("versioning is required and must be specified")
	}
	if r.putBucketVersioningRequest == nil {
		return nil, reportError("putBucketVersioningRequest is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "versioning", r.versioning, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/xml"}

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
	if r.xAmzSecurityToken != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-security-token", r.xAmzSecurityToken, "")
	}
	if r.contentMD5 != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-MD5", r.contentMD5, "")
	}
	if r.xAmzMfa != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-mfa", r.xAmzMfa, "")
	}
	if r.xAmzExpectedBucketOwner != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-expected-bucket-owner", r.xAmzExpectedBucketOwner, "")
	}
	// body params
	localVarPostBody = r.putBucketVersioningRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["hmac"]; ok {
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