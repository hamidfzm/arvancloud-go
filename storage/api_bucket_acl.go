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


// BucketACLAPIService BucketACLAPI service
type BucketACLAPIService service

type ApiGetBucketAclRequest struct {
	ctx context.Context
	ApiService *BucketACLAPIService
	bucket string
	acl *bool
	xAmzSecurityToken *string
	xAmzExpectedBucketOwner *string
}

func (r ApiGetBucketAclRequest) Acl(acl bool) ApiGetBucketAclRequest {
	r.acl = &acl
	return r
}

// This parameter is currently not supported and is not required.
func (r ApiGetBucketAclRequest) XAmzSecurityToken(xAmzSecurityToken string) ApiGetBucketAclRequest {
	r.xAmzSecurityToken = &xAmzSecurityToken
	return r
}

// The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error.
func (r ApiGetBucketAclRequest) XAmzExpectedBucketOwner(xAmzExpectedBucketOwner string) ApiGetBucketAclRequest {
	r.xAmzExpectedBucketOwner = &xAmzExpectedBucketOwner
	return r
}

func (r ApiGetBucketAclRequest) Execute() (*GetBucketAclOutput, *http.Response, error) {
	return r.ApiService.GetBucketAclExecute(r)
}

/*
GetBucketAcl Method for GetBucketAcl

<p>This implementation of the <code>GET</code> action uses the <code>acl</code> subresource to return the access control list (ACL) of a bucket. To use <code>GET</code> to return the ACL of the bucket, you must have <code>READ_ACP</code> access to the bucket. If <code>READ_ACP</code> permission is granted to the anonymous user, you can return the ACL of the bucket without using an authorization header.</p> <p class="title"> <b>Related Resources</b> </p> <ul> <li> <p> <a href="https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjects.html">ListObjects</a> </p> </li> </ul>

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bucket Specifies the S3 bucket whose ACL is being requested.
 @return ApiGetBucketAclRequest
*/
func (a *BucketACLAPIService) GetBucketAcl(ctx context.Context, bucket string) ApiGetBucketAclRequest {
	return ApiGetBucketAclRequest{
		ApiService: a,
		ctx: ctx,
		bucket: bucket,
	}
}

// Execute executes the request
//  @return GetBucketAclOutput
func (a *BucketACLAPIService) GetBucketAclExecute(r ApiGetBucketAclRequest) (*GetBucketAclOutput, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetBucketAclOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BucketACLAPIService.GetBucketAcl")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{Bucket}#acl"
	localVarPath = strings.Replace(localVarPath, "{"+"Bucket"+"}", url.PathEscape(parameterValueToString(r.bucket, "bucket")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.acl == nil {
		return localVarReturnValue, nil, reportError("acl is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "acl", r.acl, "")
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

type ApiPutBucketAclRequest struct {
	ctx context.Context
	ApiService *BucketACLAPIService
	bucket string
	acl *bool
	putBucketAclRequest *PutBucketAclRequest
	xAmzSecurityToken *string
	xAmzAcl *string
	contentMD5 *string
	xAmzGrantFullControl *string
	xAmzGrantRead *string
	xAmzGrantReadAcp *string
	xAmzGrantWrite *string
	xAmzGrantWriteAcp *string
	xAmzExpectedBucketOwner *string
}

func (r ApiPutBucketAclRequest) Acl(acl bool) ApiPutBucketAclRequest {
	r.acl = &acl
	return r
}

func (r ApiPutBucketAclRequest) PutBucketAclRequest(putBucketAclRequest PutBucketAclRequest) ApiPutBucketAclRequest {
	r.putBucketAclRequest = &putBucketAclRequest
	return r
}

// This parameter is currently not supported and is not required.
func (r ApiPutBucketAclRequest) XAmzSecurityToken(xAmzSecurityToken string) ApiPutBucketAclRequest {
	r.xAmzSecurityToken = &xAmzSecurityToken
	return r
}

// The canned ACL to apply to the bucket.
func (r ApiPutBucketAclRequest) XAmzAcl(xAmzAcl string) ApiPutBucketAclRequest {
	r.xAmzAcl = &xAmzAcl
	return r
}

// &lt;p&gt;The base64-encoded 128-bit MD5 digest of the data. This header must be used as a message integrity check to verify that the request body was not corrupted in transit. For more information, go to &lt;a href&#x3D;\&quot;http://www.ietf.org/rfc/rfc1864.txt\&quot;&gt;RFC 1864.&lt;/a&gt; &lt;/p&gt; &lt;p&gt;For requests made using the Amazon Web Services Command Line Interface (CLI) or Amazon Web Services SDKs, this field is calculated automatically.&lt;/p&gt;
func (r ApiPutBucketAclRequest) ContentMD5(contentMD5 string) ApiPutBucketAclRequest {
	r.contentMD5 = &contentMD5
	return r
}

// Allows grantee the read, write, read ACP, and write ACP permissions on the bucket.
func (r ApiPutBucketAclRequest) XAmzGrantFullControl(xAmzGrantFullControl string) ApiPutBucketAclRequest {
	r.xAmzGrantFullControl = &xAmzGrantFullControl
	return r
}

// Allows grantee to list the objects in the bucket.
func (r ApiPutBucketAclRequest) XAmzGrantRead(xAmzGrantRead string) ApiPutBucketAclRequest {
	r.xAmzGrantRead = &xAmzGrantRead
	return r
}

// Allows grantee to read the bucket ACL.
func (r ApiPutBucketAclRequest) XAmzGrantReadAcp(xAmzGrantReadAcp string) ApiPutBucketAclRequest {
	r.xAmzGrantReadAcp = &xAmzGrantReadAcp
	return r
}

// &lt;p&gt;Allows grantee to create new objects in the bucket.&lt;/p&gt; &lt;p&gt;For the bucket and object owners of existing objects, also allows deletions and overwrites of those objects.&lt;/p&gt;
func (r ApiPutBucketAclRequest) XAmzGrantWrite(xAmzGrantWrite string) ApiPutBucketAclRequest {
	r.xAmzGrantWrite = &xAmzGrantWrite
	return r
}

// Allows grantee to write the ACL for the applicable bucket.
func (r ApiPutBucketAclRequest) XAmzGrantWriteAcp(xAmzGrantWriteAcp string) ApiPutBucketAclRequest {
	r.xAmzGrantWriteAcp = &xAmzGrantWriteAcp
	return r
}

// The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error.
func (r ApiPutBucketAclRequest) XAmzExpectedBucketOwner(xAmzExpectedBucketOwner string) ApiPutBucketAclRequest {
	r.xAmzExpectedBucketOwner = &xAmzExpectedBucketOwner
	return r
}

func (r ApiPutBucketAclRequest) Execute() (*http.Response, error) {
	return r.ApiService.PutBucketAclExecute(r)
}

/*
PutBucketAcl Method for PutBucketAcl

<p>Sets the permissions on an existing bucket using access control lists (ACL). For more information, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/S3_ACLs_UsingACLs.html">Using ACLs</a>. To set the ACL of a bucket, you must have <code>WRITE_ACP</code> permission.</p> <p>You can use one of the following two ways to set a bucket's permissions:</p> <ul> <li> <p>Specify the ACL in the request body</p> </li> <li> <p>Specify permissions using request headers</p> </li> </ul> <note> <p>You cannot specify access permission using both the body and the request headers.</p> </note> <p>Depending on your application needs, you may choose to set the ACL on a bucket using either the request body or the headers. For example, if you have an existing application that updates a bucket ACL using the request body, then you can continue to use that approach.</p> <p> <b>Access Permissions</b> </p> <p>You can set access permissions using one of the following methods:</p> <ul> <li> <p>Specify a canned ACL with the <code>x-amz-acl</code> request header. ArvanCloud S3 supports a set of predefined ACLs, known as <i>canned ACLs</i>. Each canned ACL has a predefined set of grantees and permissions. Specify the canned ACL name as the value of <code>x-amz-acl</code>. If you use this header, you cannot use other access control-specific headers in your request. For more information, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL">Canned ACL</a>.</p> </li> <li> <p>Specify access permissions explicitly with the <code>x-amz-grant-read</code>, <code>x-amz-grant-read-acp</code>, <code>x-amz-grant-write-acp</code>, and <code>x-amz-grant-full-control</code> headers. When using these headers, you specify explicit access permissions and grantees (Amazon Web Services accounts or ArvanCloud S3 groups) who will receive the permission. If you use these ACL-specific headers, you cannot use the <code>x-amz-acl</code> header to set a canned ACL. These parameters map to the set of permissions that ArvanCloud S3 supports in an ACL. For more information, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html">Access Control List (ACL) Overview</a>.</p> <p>You specify each grantee as a type=value pair, where the type is one of the following:</p> <ul> <li> <p> <code>id</code> – if the value specified is the canonical user ID of an Amazon Web Services account</p> </li> <li> <p> <code>uri</code> – if you are granting permissions to a predefined group</p> </li> <li> <p> <code>emailAddress</code> – if the value specified is the email address of an Amazon Web Services account</p> <note> <p>

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bucket The bucket to which to apply the ACL.
 @return ApiPutBucketAclRequest
*/
func (a *BucketACLAPIService) PutBucketAcl(ctx context.Context, bucket string) ApiPutBucketAclRequest {
	return ApiPutBucketAclRequest{
		ApiService: a,
		ctx: ctx,
		bucket: bucket,
	}
}

// Execute executes the request
func (a *BucketACLAPIService) PutBucketAclExecute(r ApiPutBucketAclRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BucketACLAPIService.PutBucketAcl")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{Bucket}#acl"
	localVarPath = strings.Replace(localVarPath, "{"+"Bucket"+"}", url.PathEscape(parameterValueToString(r.bucket, "bucket")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.acl == nil {
		return nil, reportError("acl is required and must be specified")
	}
	if r.putBucketAclRequest == nil {
		return nil, reportError("putBucketAclRequest is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "acl", r.acl, "")
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
	if r.xAmzAcl != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-acl", r.xAmzAcl, "")
	}
	if r.contentMD5 != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-MD5", r.contentMD5, "")
	}
	if r.xAmzGrantFullControl != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-grant-full-control", r.xAmzGrantFullControl, "")
	}
	if r.xAmzGrantRead != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-grant-read", r.xAmzGrantRead, "")
	}
	if r.xAmzGrantReadAcp != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-grant-read-acp", r.xAmzGrantReadAcp, "")
	}
	if r.xAmzGrantWrite != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-grant-write", r.xAmzGrantWrite, "")
	}
	if r.xAmzGrantWriteAcp != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-grant-write-acp", r.xAmzGrantWriteAcp, "")
	}
	if r.xAmzExpectedBucketOwner != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-expected-bucket-owner", r.xAmzExpectedBucketOwner, "")
	}
	// body params
	localVarPostBody = r.putBucketAclRequest
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
