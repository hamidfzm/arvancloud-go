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


// ObjectACLAPIService ObjectACLAPI service
type ObjectACLAPIService service

type ApiGetObjectAclRequest struct {
	ctx context.Context
	ApiService *ObjectACLAPIService
	bucket string
	key string
	acl *bool
	xAmzSecurityToken *string
	versionId *string
	xAmzRequestPayer *string
	xAmzExpectedBucketOwner *string
}

func (r ApiGetObjectAclRequest) Acl(acl bool) ApiGetObjectAclRequest {
	r.acl = &acl
	return r
}

// This parameter is currently not supported and is not required.
func (r ApiGetObjectAclRequest) XAmzSecurityToken(xAmzSecurityToken string) ApiGetObjectAclRequest {
	r.xAmzSecurityToken = &xAmzSecurityToken
	return r
}

// VersionId used to reference a specific version of the object.
func (r ApiGetObjectAclRequest) VersionId(versionId string) ApiGetObjectAclRequest {
	r.versionId = &versionId
	return r
}

// 
func (r ApiGetObjectAclRequest) XAmzRequestPayer(xAmzRequestPayer string) ApiGetObjectAclRequest {
	r.xAmzRequestPayer = &xAmzRequestPayer
	return r
}

// The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error.
func (r ApiGetObjectAclRequest) XAmzExpectedBucketOwner(xAmzExpectedBucketOwner string) ApiGetObjectAclRequest {
	r.xAmzExpectedBucketOwner = &xAmzExpectedBucketOwner
	return r
}

func (r ApiGetObjectAclRequest) Execute() (*GetObjectAclOutput, *http.Response, error) {
	return r.ApiService.GetObjectAclExecute(r)
}

/*
GetObjectAcl Method for GetObjectAcl

<p>Returns the access control list (ACL) of an object. To use this operation, you must have <code>READ_ACP</code> access to the object.</p> <p>

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bucket <p>The bucket name that contains the object for which to get the ACL information. </p> <p>
 @param key The key of the object for which to get the ACL information.
 @return ApiGetObjectAclRequest
*/
func (a *ObjectACLAPIService) GetObjectAcl(ctx context.Context, bucket string, key string) ApiGetObjectAclRequest {
	return ApiGetObjectAclRequest{
		ApiService: a,
		ctx: ctx,
		bucket: bucket,
		key: key,
	}
}

// Execute executes the request
//  @return GetObjectAclOutput
func (a *ObjectACLAPIService) GetObjectAclExecute(r ApiGetObjectAclRequest) (*GetObjectAclOutput, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetObjectAclOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectACLAPIService.GetObjectAcl")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{Bucket}/{Key}#acl"
	localVarPath = strings.Replace(localVarPath, "{"+"Bucket"+"}", url.PathEscape(parameterValueToString(r.bucket, "bucket")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"Key"+"}", url.PathEscape(parameterValueToString(r.key, "key")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.key) < 1 {
		return localVarReturnValue, nil, reportError("key must have at least 1 elements")
	}
	if r.acl == nil {
		return localVarReturnValue, nil, reportError("acl is required and must be specified")
	}

	if r.versionId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "versionId", r.versionId, "")
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
	if r.xAmzRequestPayer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-request-payer", r.xAmzRequestPayer, "")
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
		if localVarHTTPResponse.StatusCode == 480 {
			var v interface{}
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

type ApiPutObjectAclRequest struct {
	ctx context.Context
	ApiService *ObjectACLAPIService
	bucket string
	key string
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
	xAmzRequestPayer *string
	versionId *string
	xAmzExpectedBucketOwner *string
}

func (r ApiPutObjectAclRequest) Acl(acl bool) ApiPutObjectAclRequest {
	r.acl = &acl
	return r
}

func (r ApiPutObjectAclRequest) PutBucketAclRequest(putBucketAclRequest PutBucketAclRequest) ApiPutObjectAclRequest {
	r.putBucketAclRequest = &putBucketAclRequest
	return r
}

// This parameter is currently not supported and is not required.
func (r ApiPutObjectAclRequest) XAmzSecurityToken(xAmzSecurityToken string) ApiPutObjectAclRequest {
	r.xAmzSecurityToken = &xAmzSecurityToken
	return r
}

// The canned ACL to apply to the object. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL\&quot;&gt;Canned ACL&lt;/a&gt;.
func (r ApiPutObjectAclRequest) XAmzAcl(xAmzAcl string) ApiPutObjectAclRequest {
	r.xAmzAcl = &xAmzAcl
	return r
}

// &lt;p&gt;The base64-encoded 128-bit MD5 digest of the data. This header must be used as a message integrity check to verify that the request body was not corrupted in transit. For more information, go to &lt;a href&#x3D;\&quot;http://www.ietf.org/rfc/rfc1864.txt\&quot;&gt;RFC 1864.&amp;gt;&lt;/a&gt; &lt;/p&gt; &lt;p&gt;For requests made using the Amazon Web Services Command Line Interface (CLI) or Amazon Web Services SDKs, this field is calculated automatically.&lt;/p&gt;
func (r ApiPutObjectAclRequest) ContentMD5(contentMD5 string) ApiPutObjectAclRequest {
	r.contentMD5 = &contentMD5
	return r
}

// &lt;p&gt;Allows grantee the read, write, read ACP, and write ACP permissions on the bucket.&lt;/p&gt; &lt;/p&gt;
func (r ApiPutObjectAclRequest) XAmzGrantFullControl(xAmzGrantFullControl string) ApiPutObjectAclRequest {
	r.xAmzGrantFullControl = &xAmzGrantFullControl
	return r
}

// &lt;p&gt;Allows grantee to list the objects in the bucket.&lt;/p&gt; &lt;/p&gt;
func (r ApiPutObjectAclRequest) XAmzGrantRead(xAmzGrantRead string) ApiPutObjectAclRequest {
	r.xAmzGrantRead = &xAmzGrantRead
	return r
}

// &lt;p&gt;Allows grantee to read the bucket ACL.&lt;/p&gt; &lt;/p&gt;
func (r ApiPutObjectAclRequest) XAmzGrantReadAcp(xAmzGrantReadAcp string) ApiPutObjectAclRequest {
	r.xAmzGrantReadAcp = &xAmzGrantReadAcp
	return r
}

// &lt;p&gt;Allows grantee to create new objects in the bucket.&lt;/p&gt; &lt;p&gt;For the bucket and object owners of existing objects, also allows deletions and overwrites of those objects.&lt;/p&gt;
func (r ApiPutObjectAclRequest) XAmzGrantWrite(xAmzGrantWrite string) ApiPutObjectAclRequest {
	r.xAmzGrantWrite = &xAmzGrantWrite
	return r
}

// &lt;p&gt;Allows grantee to write the ACL for the applicable bucket.&lt;/p&gt; &lt;/p&gt;
func (r ApiPutObjectAclRequest) XAmzGrantWriteAcp(xAmzGrantWriteAcp string) ApiPutObjectAclRequest {
	r.xAmzGrantWriteAcp = &xAmzGrantWriteAcp
	return r
}

// 
func (r ApiPutObjectAclRequest) XAmzRequestPayer(xAmzRequestPayer string) ApiPutObjectAclRequest {
	r.xAmzRequestPayer = &xAmzRequestPayer
	return r
}

// VersionId used to reference a specific version of the object.
func (r ApiPutObjectAclRequest) VersionId(versionId string) ApiPutObjectAclRequest {
	r.versionId = &versionId
	return r
}

// The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error.
func (r ApiPutObjectAclRequest) XAmzExpectedBucketOwner(xAmzExpectedBucketOwner string) ApiPutObjectAclRequest {
	r.xAmzExpectedBucketOwner = &xAmzExpectedBucketOwner
	return r
}

func (r ApiPutObjectAclRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.PutObjectAclExecute(r)
}

/*
PutObjectAcl Method for PutObjectAcl

<p>Uses the <code>acl</code> subresource to set the access control list (ACL) permissions for a new or existing object in an S3 bucket. You must have <code>WRITE_ACP</code> permission to set the ACL of an object. For more information, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#permissions">What permissions can I grant?</a> in the <i>User Guide</i>.</p> </p> <p>Depending on your application needs, you can choose to set the ACL on an object using either the request body or the headers. For example, if you have an existing application that updates a bucket ACL using the request body, you can continue to use that approach. For more information, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html">Access Control List (ACL) Overview</a> in the <i>S3 User Guide</i>.</p> <p> <b>Access Permissions</b> </p> <p>You can set access permissions using one of the following methods:</p> <ul> <li> <p>Specify a canned ACL with the <code>x-amz-acl</code> request header. ArvanCloud S3 supports a set of predefined ACLs, known as canned ACLs. Each canned ACL has a predefined set of grantees and permissions. Specify the canned ACL name as the value of <code>x-amz-ac</code>l. If you use this header, you cannot use other access control-specific headers in your request. For more information, see <a href="https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL">Canned ACL</a>.</p> </li> <li> <p>Specify access permissions explicitly with the <code>x-amz-grant-read</code>, <code>x-amz-grant-read-acp</code>, <code>x-amz-grant-write-acp</code>, and <code>x-amz-grant-full-control</code> headers. When using these headers, you specify explicit access permissions and grantees (Amazon Web Services accounts or ArvanCloud S3 groups) who will receive the permission. If you use these ACL-specific headers, you cannot use <code>x-amz-acl</code> header to set a canned ACL. <p>You specify each grantee as a type=value pair, where the type is one of the following:</p> <ul> <li> <p> <code>id</code> – if the value specified is the canonical user ID of an Amazon Web Services account</p> </li> <li> <p> <code>uri</code> – if you are granting permissions to a predefined group</p> </li> <li> <p> <code>emailAddress</code> – if the value specified is the email address of an Amazon Web Services account</p> <note> <p>

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bucket <p>The bucket name that contains the object to which you want to attach the ACL. </p> <p>
 @param key <p>Key for which the PUT action was initiated.</p> <p>W
 @return ApiPutObjectAclRequest
*/
func (a *ObjectACLAPIService) PutObjectAcl(ctx context.Context, bucket string, key string) ApiPutObjectAclRequest {
	return ApiPutObjectAclRequest{
		ApiService: a,
		ctx: ctx,
		bucket: bucket,
		key: key,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ObjectACLAPIService) PutObjectAclExecute(r ApiPutObjectAclRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectACLAPIService.PutObjectAcl")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{Bucket}/{Key}#acl"
	localVarPath = strings.Replace(localVarPath, "{"+"Bucket"+"}", url.PathEscape(parameterValueToString(r.bucket, "bucket")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"Key"+"}", url.PathEscape(parameterValueToString(r.key, "key")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.key) < 1 {
		return localVarReturnValue, nil, reportError("key must have at least 1 elements")
	}
	if r.acl == nil {
		return localVarReturnValue, nil, reportError("acl is required and must be specified")
	}
	if r.putBucketAclRequest == nil {
		return localVarReturnValue, nil, reportError("putBucketAclRequest is required and must be specified")
	}

	if r.versionId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "versionId", r.versionId, "")
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
	localVarHTTPHeaderAccepts := []string{"text/xml"}

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
	if r.xAmzRequestPayer != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-amz-request-payer", r.xAmzRequestPayer, "")
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
		if localVarHTTPResponse.StatusCode == 480 {
			var v interface{}
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
