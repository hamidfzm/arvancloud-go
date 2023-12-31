# Go API client for storage

<p/>

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2006-03-01
- Package version: 1.1
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import storage "github.com/hamidfzm/arvancloud-go/storage"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), storage.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), storage.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), storage.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), storage.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://s3.ir-thr-at1.arvanstorage.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BucketAPI* | [**CreateBucket**](docs/BucketAPI.md#createbucket) | **Put** /{Bucket} | 
*BucketAPI* | [**DeleteBucket**](docs/BucketAPI.md#deletebucket) | **Delete** /{Bucket} | 
*BucketAPI* | [**DeleteObjects**](docs/BucketAPI.md#deleteobjects) | **Post** /{Bucket}#delete | 
*BucketAPI* | [**DeletePublicAccessBlock**](docs/BucketAPI.md#deletepublicaccessblock) | **Delete** /{Bucket}#publicAccessBlock | 
*BucketAPI* | [**GetPublicAccessBlock**](docs/BucketAPI.md#getpublicaccessblock) | **Get** /{Bucket}#publicAccessBlock | 
*BucketAPI* | [**HeadBucket**](docs/BucketAPI.md#headbucket) | **Head** /{Bucket} | 
*BucketAPI* | [**ListBuckets**](docs/BucketAPI.md#listbuckets) | **Get** /{Bucket} | 
*BucketAPI* | [**ListObjectVersions**](docs/BucketAPI.md#listobjectversions) | **Get** /{Bucket}#versions | 
*BucketAPI* | [**ListObjects**](docs/BucketAPI.md#listobjects) | **Get** /{Bucket}#Listobjects | 
*BucketAPI* | [**ListObjectsV2**](docs/BucketAPI.md#listobjectsv2) | **Get** /{Bucket}#list-type&#x3D;2 | 
*BucketAPI* | [**PutPublicAccessBlock**](docs/BucketAPI.md#putpublicaccessblock) | **Put** /{Bucket}#publicAccessBlock | 
*BucketACLAPI* | [**GetBucketAcl**](docs/BucketACLAPI.md#getbucketacl) | **Get** /{Bucket}#acl | 
*BucketACLAPI* | [**PutBucketAcl**](docs/BucketACLAPI.md#putbucketacl) | **Put** /{Bucket}#acl | 
*BucketCORSAPI* | [**DeleteBucketCors**](docs/BucketCORSAPI.md#deletebucketcors) | **Delete** /{Bucket}#cors | 
*BucketCORSAPI* | [**GetBucketCors**](docs/BucketCORSAPI.md#getbucketcors) | **Get** /{Bucket}#cors | 
*BucketCORSAPI* | [**PutBucketCors**](docs/BucketCORSAPI.md#putbucketcors) | **Put** /{Bucket}#cors | 
*BucketLifecycleConfigurationAPI* | [**DeleteBucketLifecycle**](docs/BucketLifecycleConfigurationAPI.md#deletebucketlifecycle) | **Delete** /{Bucket}#lifecycle | 
*BucketLifecycleConfigurationAPI* | [**GetBucketLifecycleConfiguration**](docs/BucketLifecycleConfigurationAPI.md#getbucketlifecycleconfiguration) | **Get** /{Bucket}#lifecycle | 
*BucketLifecycleConfigurationAPI* | [**PutBucketLifecycleConfiguration**](docs/BucketLifecycleConfigurationAPI.md#putbucketlifecycleconfiguration) | **Put** /{Bucket}#lifecycle | 
*BucketPolicyAPI* | [**DeleteBucketPolicy**](docs/BucketPolicyAPI.md#deletebucketpolicy) | **Delete** /{Bucket}#policy | 
*BucketPolicyAPI* | [**GetBucketPolicy**](docs/BucketPolicyAPI.md#getbucketpolicy) | **Get** /{Bucket}#policy | 
*BucketPolicyAPI* | [**GetBucketPolicyStatus**](docs/BucketPolicyAPI.md#getbucketpolicystatus) | **Get** /{Bucket}#policyStatus | 
*BucketPolicyAPI* | [**PutBucketPolicy**](docs/BucketPolicyAPI.md#putbucketpolicy) | **Put** /{Bucket}#policy | 
*BucketTaggingAPI* | [**DeleteBucketTagging**](docs/BucketTaggingAPI.md#deletebuckettagging) | **Delete** /{Bucket}#tagging | 
*BucketTaggingAPI* | [**GetBucketTagging**](docs/BucketTaggingAPI.md#getbuckettagging) | **Get** /{Bucket}#tagging | 
*BucketTaggingAPI* | [**PutBucketTagging**](docs/BucketTaggingAPI.md#putbuckettagging) | **Put** /{Bucket}#tagging | 
*BucketVersioningAPI* | [**GetBucketVersioning**](docs/BucketVersioningAPI.md#getbucketversioning) | **Get** /{Bucket}#versioning | 
*BucketVersioningAPI* | [**PutBucketVersioning**](docs/BucketVersioningAPI.md#putbucketversioning) | **Put** /{Bucket}#versioning | 
*BucketWebsiteConfigurationAPI* | [**DeleteBucketWebsite**](docs/BucketWebsiteConfigurationAPI.md#deletebucketwebsite) | **Delete** /{Bucket}#website | 
*BucketWebsiteConfigurationAPI* | [**GetBucketWebsite**](docs/BucketWebsiteConfigurationAPI.md#getbucketwebsite) | **Get** /{Bucket}#website | 
*BucketWebsiteConfigurationAPI* | [**PutBucketWebsite**](docs/BucketWebsiteConfigurationAPI.md#putbucketwebsite) | **Put** /{Bucket}#website | 
*MultipartAPI* | [**AbortMultipartUpload**](docs/MultipartAPI.md#abortmultipartupload) | **Delete** /{Bucket}/{Key}#uploadId | 
*MultipartAPI* | [**CompleteMultipartUpload**](docs/MultipartAPI.md#completemultipartupload) | **Post** /{Bucket}/{Key}#uploadId | 
*MultipartAPI* | [**CreateMultipartUpload**](docs/MultipartAPI.md#createmultipartupload) | **Post** /{Bucket}/{Key}#uploads | 
*MultipartAPI* | [**ListMultipartUploads**](docs/MultipartAPI.md#listmultipartuploads) | **Get** /{Bucket}#uploads | 
*MultipartAPI* | [**ListParts**](docs/MultipartAPI.md#listparts) | **Get** /{Bucket}/{Key}#uploadId | 
*MultipartAPI* | [**UploadPart**](docs/MultipartAPI.md#uploadpart) | **Put** /{Bucket}/{Key}#partNumber&amp;uploadId | 
*ObjectAPI* | [**DeleteObject**](docs/ObjectAPI.md#deleteobject) | **Delete** /{Bucket}/{Key} | 
*ObjectAPI* | [**GetObject**](docs/ObjectAPI.md#getobject) | **Get** /{Bucket}/{Key} | 
*ObjectAPI* | [**HeadObject**](docs/ObjectAPI.md#headobject) | **Head** /{Bucket}/{Key} | 
*ObjectAPI* | [**PutObject**](docs/ObjectAPI.md#putobject) | **Put** /{Bucket}/{Key} | 
*ObjectACLAPI* | [**GetObjectAcl**](docs/ObjectACLAPI.md#getobjectacl) | **Get** /{Bucket}/{Key}#acl | 
*ObjectACLAPI* | [**PutObjectAcl**](docs/ObjectACLAPI.md#putobjectacl) | **Put** /{Bucket}/{Key}#acl | 
*ObjectTaggingAPI* | [**DeleteObjectTagging**](docs/ObjectTaggingAPI.md#deleteobjecttagging) | **Delete** /{Bucket}/{Key}#tagging | 
*ObjectTaggingAPI* | [**GetObjectTagging**](docs/ObjectTaggingAPI.md#getobjecttagging) | **Get** /{Bucket}/{Key}#tagging | 
*ObjectTaggingAPI* | [**PutObjectTagging**](docs/ObjectTaggingAPI.md#putobjecttagging) | **Put** /{Bucket}/{Key}#tagging | 


## Documentation For Models

 - [AbortIncompleteMultipartUpload](docs/AbortIncompleteMultipartUpload.md)
 - [AccelerateConfiguration](docs/AccelerateConfiguration.md)
 - [AccessControlPolicy](docs/AccessControlPolicy.md)
 - [AccessControlTranslation](docs/AccessControlTranslation.md)
 - [AnalyticsAndOperator](docs/AnalyticsAndOperator.md)
 - [AnalyticsConfiguration](docs/AnalyticsConfiguration.md)
 - [AnalyticsExportDestination](docs/AnalyticsExportDestination.md)
 - [AnalyticsExportDestinationS3BucketDestination](docs/AnalyticsExportDestinationS3BucketDestination.md)
 - [AnalyticsS3BucketDestination](docs/AnalyticsS3BucketDestination.md)
 - [AnalyticsS3ExportFileFormat](docs/AnalyticsS3ExportFileFormat.md)
 - [ArchiveStatus](docs/ArchiveStatus.md)
 - [Bucket](docs/Bucket.md)
 - [BucketCannedACL](docs/BucketCannedACL.md)
 - [BucketLifecycleConfiguration](docs/BucketLifecycleConfiguration.md)
 - [BucketLocationConstraint](docs/BucketLocationConstraint.md)
 - [BucketLoggingStatus](docs/BucketLoggingStatus.md)
 - [BucketLogsPermission](docs/BucketLogsPermission.md)
 - [BucketVersioningStatus](docs/BucketVersioningStatus.md)
 - [BucketsInner](docs/BucketsInner.md)
 - [CORSConfiguration](docs/CORSConfiguration.md)
 - [CORSRule](docs/CORSRule.md)
 - [CSVInput](docs/CSVInput.md)
 - [CSVOutput](docs/CSVOutput.md)
 - [CommonPrefix](docs/CommonPrefix.md)
 - [CompleteMultipartUploadOutput](docs/CompleteMultipartUploadOutput.md)
 - [CompleteMultipartUploadRequest](docs/CompleteMultipartUploadRequest.md)
 - [CompleteMultipartUploadRequestCompleteMultipartUpload](docs/CompleteMultipartUploadRequestCompleteMultipartUpload.md)
 - [CompleteMultipartUploadRequestMultipartUpload](docs/CompleteMultipartUploadRequestMultipartUpload.md)
 - [CompletedMultipartUpload](docs/CompletedMultipartUpload.md)
 - [CompletedPart](docs/CompletedPart.md)
 - [CompressionType](docs/CompressionType.md)
 - [Condition](docs/Condition.md)
 - [CopyObjectOutput](docs/CopyObjectOutput.md)
 - [CopyObjectOutputCopyObjectResult](docs/CopyObjectOutputCopyObjectResult.md)
 - [CopyObjectRequest](docs/CopyObjectRequest.md)
 - [CopyObjectResult](docs/CopyObjectResult.md)
 - [CopyPartResult](docs/CopyPartResult.md)
 - [CreateBucketConfiguration](docs/CreateBucketConfiguration.md)
 - [CreateBucketRequest](docs/CreateBucketRequest.md)
 - [CreateBucketRequestCreateBucketConfiguration](docs/CreateBucketRequestCreateBucketConfiguration.md)
 - [CreateMultipartUploadOutput](docs/CreateMultipartUploadOutput.md)
 - [CreateMultipartUploadRequest](docs/CreateMultipartUploadRequest.md)
 - [DefaultRetention](docs/DefaultRetention.md)
 - [Delete](docs/Delete.md)
 - [DeleteMarkerEntry](docs/DeleteMarkerEntry.md)
 - [DeleteMarkerReplication](docs/DeleteMarkerReplication.md)
 - [DeleteMarkerReplicationStatus](docs/DeleteMarkerReplicationStatus.md)
 - [DeleteObjectsOutput](docs/DeleteObjectsOutput.md)
 - [DeleteObjectsRequest](docs/DeleteObjectsRequest.md)
 - [DeleteObjectsRequestDelete](docs/DeleteObjectsRequestDelete.md)
 - [DeletedObject](docs/DeletedObject.md)
 - [Destination](docs/Destination.md)
 - [DestinationAccessControlTranslation](docs/DestinationAccessControlTranslation.md)
 - [DestinationEncryptionConfiguration](docs/DestinationEncryptionConfiguration.md)
 - [DestinationMetrics](docs/DestinationMetrics.md)
 - [DestinationReplicationTime](docs/DestinationReplicationTime.md)
 - [EncodingType](docs/EncodingType.md)
 - [Encryption](docs/Encryption.md)
 - [EncryptionConfiguration](docs/EncryptionConfiguration.md)
 - [Error](docs/Error.md)
 - [ErrorDocument](docs/ErrorDocument.md)
 - [Event](docs/Event.md)
 - [ExistingObjectReplication](docs/ExistingObjectReplication.md)
 - [ExistingObjectReplicationStatus](docs/ExistingObjectReplicationStatus.md)
 - [ExpirationStatus](docs/ExpirationStatus.md)
 - [ExpressionType](docs/ExpressionType.md)
 - [FileHeaderInfo](docs/FileHeaderInfo.md)
 - [FilterRule](docs/FilterRule.md)
 - [FilterRuleName](docs/FilterRuleName.md)
 - [GetBucketAclOutput](docs/GetBucketAclOutput.md)
 - [GetBucketCorsOutput](docs/GetBucketCorsOutput.md)
 - [GetBucketLifecycleConfigurationOutput](docs/GetBucketLifecycleConfigurationOutput.md)
 - [GetBucketLifecycleOutput](docs/GetBucketLifecycleOutput.md)
 - [GetBucketPolicyOutput](docs/GetBucketPolicyOutput.md)
 - [GetBucketPolicyStatusOutput](docs/GetBucketPolicyStatusOutput.md)
 - [GetBucketPolicyStatusOutputPolicyStatus](docs/GetBucketPolicyStatusOutputPolicyStatus.md)
 - [GetBucketTaggingOutput](docs/GetBucketTaggingOutput.md)
 - [GetBucketVersioningOutput](docs/GetBucketVersioningOutput.md)
 - [GetBucketWebsiteOutput](docs/GetBucketWebsiteOutput.md)
 - [GetBucketWebsiteOutputErrorDocument](docs/GetBucketWebsiteOutputErrorDocument.md)
 - [GetBucketWebsiteOutputIndexDocument](docs/GetBucketWebsiteOutputIndexDocument.md)
 - [GetBucketWebsiteOutputRedirectAllRequestsTo](docs/GetBucketWebsiteOutputRedirectAllRequestsTo.md)
 - [GetObjectAclOutput](docs/GetObjectAclOutput.md)
 - [GetObjectOutput](docs/GetObjectOutput.md)
 - [GetObjectTaggingOutput](docs/GetObjectTaggingOutput.md)
 - [GlacierJobParameters](docs/GlacierJobParameters.md)
 - [Grant](docs/Grant.md)
 - [GrantGrantee](docs/GrantGrantee.md)
 - [Grantee](docs/Grantee.md)
 - [HeadObjectOutput](docs/HeadObjectOutput.md)
 - [IndexDocument](docs/IndexDocument.md)
 - [Initiator](docs/Initiator.md)
 - [InputSerialization](docs/InputSerialization.md)
 - [InputSerializationCSV](docs/InputSerializationCSV.md)
 - [InputSerializationJSON](docs/InputSerializationJSON.md)
 - [IntelligentTieringAccessTier](docs/IntelligentTieringAccessTier.md)
 - [IntelligentTieringAndOperator](docs/IntelligentTieringAndOperator.md)
 - [IntelligentTieringConfiguration](docs/IntelligentTieringConfiguration.md)
 - [InventoryConfiguration](docs/InventoryConfiguration.md)
 - [InventoryEncryption](docs/InventoryEncryption.md)
 - [InventoryEncryptionSSEKMS](docs/InventoryEncryptionSSEKMS.md)
 - [InventoryFormat](docs/InventoryFormat.md)
 - [InventoryFrequency](docs/InventoryFrequency.md)
 - [InventoryOptionalField](docs/InventoryOptionalField.md)
 - [InventoryS3BucketDestination](docs/InventoryS3BucketDestination.md)
 - [InventoryS3BucketDestinationEncryption](docs/InventoryS3BucketDestinationEncryption.md)
 - [JSONInput](docs/JSONInput.md)
 - [JSONOutput](docs/JSONOutput.md)
 - [JSONType](docs/JSONType.md)
 - [LambdaFunctionConfiguration](docs/LambdaFunctionConfiguration.md)
 - [LifecycleConfiguration](docs/LifecycleConfiguration.md)
 - [LifecycleExpiration](docs/LifecycleExpiration.md)
 - [LifecycleRule](docs/LifecycleRule.md)
 - [LifecycleRuleAndOperator](docs/LifecycleRuleAndOperator.md)
 - [LifecycleRuleExpiration](docs/LifecycleRuleExpiration.md)
 - [LifecycleRuleFilter](docs/LifecycleRuleFilter.md)
 - [LifecycleRuleFilterTag](docs/LifecycleRuleFilterTag.md)
 - [ListBucketsOutput](docs/ListBucketsOutput.md)
 - [ListMultipartUploadsOutput](docs/ListMultipartUploadsOutput.md)
 - [ListObjectVersionsOutput](docs/ListObjectVersionsOutput.md)
 - [ListObjectsOutput](docs/ListObjectsOutput.md)
 - [ListObjectsV2Output](docs/ListObjectsV2Output.md)
 - [ListPartsOutput](docs/ListPartsOutput.md)
 - [ListPartsOutputInitiator](docs/ListPartsOutputInitiator.md)
 - [MFADelete](docs/MFADelete.md)
 - [MFADeleteStatus](docs/MFADeleteStatus.md)
 - [MetadataDirective](docs/MetadataDirective.md)
 - [MetadataEntry](docs/MetadataEntry.md)
 - [Metrics](docs/Metrics.md)
 - [MetricsAndOperator](docs/MetricsAndOperator.md)
 - [MetricsConfiguration](docs/MetricsConfiguration.md)
 - [MetricsEventThreshold](docs/MetricsEventThreshold.md)
 - [MetricsStatus](docs/MetricsStatus.md)
 - [MultipartUpload](docs/MultipartUpload.md)
 - [MultipartUploadInitiator](docs/MultipartUploadInitiator.md)
 - [NoncurrentVersionExpiration](docs/NoncurrentVersionExpiration.md)
 - [NoncurrentVersionTransition](docs/NoncurrentVersionTransition.md)
 - [NotificationConfigurationFilter](docs/NotificationConfigurationFilter.md)
 - [NotificationConfigurationFilterKey](docs/NotificationConfigurationFilterKey.md)
 - [Object](docs/Object.md)
 - [ObjectCannedACL](docs/ObjectCannedACL.md)
 - [ObjectIdentifier](docs/ObjectIdentifier.md)
 - [ObjectLockConfiguration](docs/ObjectLockConfiguration.md)
 - [ObjectLockConfigurationRule](docs/ObjectLockConfigurationRule.md)
 - [ObjectLockEnabled](docs/ObjectLockEnabled.md)
 - [ObjectLockLegalHold](docs/ObjectLockLegalHold.md)
 - [ObjectLockLegalHoldStatus](docs/ObjectLockLegalHoldStatus.md)
 - [ObjectLockMode](docs/ObjectLockMode.md)
 - [ObjectLockRetention](docs/ObjectLockRetention.md)
 - [ObjectLockRetentionMode](docs/ObjectLockRetentionMode.md)
 - [ObjectLockRule](docs/ObjectLockRule.md)
 - [ObjectLockRuleDefaultRetention](docs/ObjectLockRuleDefaultRetention.md)
 - [ObjectOwnership](docs/ObjectOwnership.md)
 - [ObjectStorageClass](docs/ObjectStorageClass.md)
 - [ObjectVersion](docs/ObjectVersion.md)
 - [ObjectVersionStorageClass](docs/ObjectVersionStorageClass.md)
 - [OutputLocation](docs/OutputLocation.md)
 - [OutputLocationS3](docs/OutputLocationS3.md)
 - [OutputSerialization](docs/OutputSerialization.md)
 - [OutputSerializationCSV](docs/OutputSerializationCSV.md)
 - [OutputSerializationJSON](docs/OutputSerializationJSON.md)
 - [OwnerOverride](docs/OwnerOverride.md)
 - [OwnershipControls](docs/OwnershipControls.md)
 - [OwnershipControlsRule](docs/OwnershipControlsRule.md)
 - [Part](docs/Part.md)
 - [Permission](docs/Permission.md)
 - [PolicyStatus](docs/PolicyStatus.md)
 - [Progress](docs/Progress.md)
 - [ProgressEvent](docs/ProgressEvent.md)
 - [ProgressEventDetails](docs/ProgressEventDetails.md)
 - [Protocol](docs/Protocol.md)
 - [PublicAccessBlockConfiguration](docs/PublicAccessBlockConfiguration.md)
 - [PutBucketAccelerateConfigurationRequest](docs/PutBucketAccelerateConfigurationRequest.md)
 - [PutBucketAccelerateConfigurationRequestAccelerateConfiguration](docs/PutBucketAccelerateConfigurationRequestAccelerateConfiguration.md)
 - [PutBucketAclRequest](docs/PutBucketAclRequest.md)
 - [PutBucketAclRequestAccessControlPolicy](docs/PutBucketAclRequestAccessControlPolicy.md)
 - [PutBucketAnalyticsConfigurationRequest](docs/PutBucketAnalyticsConfigurationRequest.md)
 - [PutBucketAnalyticsConfigurationRequestAnalyticsConfiguration](docs/PutBucketAnalyticsConfigurationRequestAnalyticsConfiguration.md)
 - [PutBucketCorsRequest](docs/PutBucketCorsRequest.md)
 - [PutBucketCorsRequestCORSConfiguration](docs/PutBucketCorsRequestCORSConfiguration.md)
 - [PutBucketEncryptionRequest](docs/PutBucketEncryptionRequest.md)
 - [PutBucketIntelligentTieringConfigurationRequest](docs/PutBucketIntelligentTieringConfigurationRequest.md)
 - [PutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration](docs/PutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration.md)
 - [PutBucketInventoryConfigurationRequest](docs/PutBucketInventoryConfigurationRequest.md)
 - [PutBucketInventoryConfigurationRequestInventoryConfiguration](docs/PutBucketInventoryConfigurationRequestInventoryConfiguration.md)
 - [PutBucketLifecycleConfigurationRequest](docs/PutBucketLifecycleConfigurationRequest.md)
 - [PutBucketLifecycleConfigurationRequestLifecycleConfiguration](docs/PutBucketLifecycleConfigurationRequestLifecycleConfiguration.md)
 - [PutBucketLifecycleRequest](docs/PutBucketLifecycleRequest.md)
 - [PutBucketLifecycleRequestLifecycleConfiguration](docs/PutBucketLifecycleRequestLifecycleConfiguration.md)
 - [PutBucketLoggingRequest](docs/PutBucketLoggingRequest.md)
 - [PutBucketLoggingRequestBucketLoggingStatus](docs/PutBucketLoggingRequestBucketLoggingStatus.md)
 - [PutBucketMetricsConfigurationRequest](docs/PutBucketMetricsConfigurationRequest.md)
 - [PutBucketMetricsConfigurationRequestMetricsConfiguration](docs/PutBucketMetricsConfigurationRequestMetricsConfiguration.md)
 - [PutBucketNotificationConfigurationRequest](docs/PutBucketNotificationConfigurationRequest.md)
 - [PutBucketNotificationRequest](docs/PutBucketNotificationRequest.md)
 - [PutBucketOwnershipControlsRequest](docs/PutBucketOwnershipControlsRequest.md)
 - [PutBucketOwnershipControlsRequestOwnershipControls](docs/PutBucketOwnershipControlsRequestOwnershipControls.md)
 - [PutBucketPolicyRequest](docs/PutBucketPolicyRequest.md)
 - [PutBucketReplicationRequest](docs/PutBucketReplicationRequest.md)
 - [PutBucketRequestPaymentRequest](docs/PutBucketRequestPaymentRequest.md)
 - [PutBucketRequestPaymentRequestRequestPaymentConfiguration](docs/PutBucketRequestPaymentRequestRequestPaymentConfiguration.md)
 - [PutBucketTaggingRequest](docs/PutBucketTaggingRequest.md)
 - [PutBucketTaggingRequestTagging](docs/PutBucketTaggingRequestTagging.md)
 - [PutBucketVersioningRequest](docs/PutBucketVersioningRequest.md)
 - [PutBucketVersioningRequestVersioningConfiguration](docs/PutBucketVersioningRequestVersioningConfiguration.md)
 - [PutBucketWebsiteRequest](docs/PutBucketWebsiteRequest.md)
 - [PutBucketWebsiteRequestWebsiteConfiguration](docs/PutBucketWebsiteRequestWebsiteConfiguration.md)
 - [PutBucketWebsiteRequestWebsiteConfigurationErrorDocument](docs/PutBucketWebsiteRequestWebsiteConfigurationErrorDocument.md)
 - [PutBucketWebsiteRequestWebsiteConfigurationIndexDocument](docs/PutBucketWebsiteRequestWebsiteConfigurationIndexDocument.md)
 - [PutBucketWebsiteRequestWebsiteConfigurationRedirectAllRequestsTo](docs/PutBucketWebsiteRequestWebsiteConfigurationRedirectAllRequestsTo.md)
 - [PutObjectAclRequest](docs/PutObjectAclRequest.md)
 - [PutObjectLegalHoldRequest](docs/PutObjectLegalHoldRequest.md)
 - [PutObjectLegalHoldRequestLegalHold](docs/PutObjectLegalHoldRequestLegalHold.md)
 - [PutObjectLockConfigurationRequest](docs/PutObjectLockConfigurationRequest.md)
 - [PutObjectLockConfigurationRequestObjectLockConfiguration](docs/PutObjectLockConfigurationRequestObjectLockConfiguration.md)
 - [PutObjectRequest](docs/PutObjectRequest.md)
 - [PutObjectRetentionRequest](docs/PutObjectRetentionRequest.md)
 - [PutObjectRetentionRequestRetention](docs/PutObjectRetentionRequestRetention.md)
 - [PutObjectTaggingRequest](docs/PutObjectTaggingRequest.md)
 - [PutObjectTaggingRequestTagging](docs/PutObjectTaggingRequestTagging.md)
 - [PutPublicAccessBlockRequest](docs/PutPublicAccessBlockRequest.md)
 - [PutPublicAccessBlockRequestPublicAccessBlockConfiguration](docs/PutPublicAccessBlockRequestPublicAccessBlockConfiguration.md)
 - [QueueConfiguration](docs/QueueConfiguration.md)
 - [QuoteFields](docs/QuoteFields.md)
 - [RecordsEvent](docs/RecordsEvent.md)
 - [Redirect](docs/Redirect.md)
 - [RedirectAllRequestsTo](docs/RedirectAllRequestsTo.md)
 - [ReplicaModifications](docs/ReplicaModifications.md)
 - [ReplicaModificationsStatus](docs/ReplicaModificationsStatus.md)
 - [ReplicationConfiguration](docs/ReplicationConfiguration.md)
 - [ReplicationRule](docs/ReplicationRule.md)
 - [ReplicationRuleAndOperator](docs/ReplicationRuleAndOperator.md)
 - [ReplicationRuleDestination](docs/ReplicationRuleDestination.md)
 - [ReplicationRuleExistingObjectReplication](docs/ReplicationRuleExistingObjectReplication.md)
 - [ReplicationRuleFilter](docs/ReplicationRuleFilter.md)
 - [ReplicationRuleFilterAnd](docs/ReplicationRuleFilterAnd.md)
 - [ReplicationRuleFilterTag](docs/ReplicationRuleFilterTag.md)
 - [ReplicationRuleSourceSelectionCriteria](docs/ReplicationRuleSourceSelectionCriteria.md)
 - [ReplicationRuleStatus](docs/ReplicationRuleStatus.md)
 - [ReplicationStatus](docs/ReplicationStatus.md)
 - [ReplicationTime](docs/ReplicationTime.md)
 - [ReplicationTimeStatus](docs/ReplicationTimeStatus.md)
 - [ReplicationTimeTime](docs/ReplicationTimeTime.md)
 - [ReplicationTimeValue](docs/ReplicationTimeValue.md)
 - [RequestCharged](docs/RequestCharged.md)
 - [RequestPayer](docs/RequestPayer.md)
 - [RequestPaymentConfiguration](docs/RequestPaymentConfiguration.md)
 - [RequestProgress](docs/RequestProgress.md)
 - [RestoreObjectRequest](docs/RestoreObjectRequest.md)
 - [RestoreRequest](docs/RestoreRequest.md)
 - [RestoreRequestGlacierJobParameters](docs/RestoreRequestGlacierJobParameters.md)
 - [RestoreRequestOutputLocation](docs/RestoreRequestOutputLocation.md)
 - [RestoreRequestSelectParameters](docs/RestoreRequestSelectParameters.md)
 - [RestoreRequestType](docs/RestoreRequestType.md)
 - [RoutingRule](docs/RoutingRule.md)
 - [RoutingRuleCondition](docs/RoutingRuleCondition.md)
 - [RoutingRuleRedirect](docs/RoutingRuleRedirect.md)
 - [RoutingRulesInner](docs/RoutingRulesInner.md)
 - [Rule](docs/Rule.md)
 - [RuleExpiration](docs/RuleExpiration.md)
 - [RuleTransition](docs/RuleTransition.md)
 - [S3KeyFilter](docs/S3KeyFilter.md)
 - [S3Location](docs/S3Location.md)
 - [S3LocationTagging](docs/S3LocationTagging.md)
 - [SSEKMS](docs/SSEKMS.md)
 - [ScanRange](docs/ScanRange.md)
 - [SelectObjectContentEventStream](docs/SelectObjectContentEventStream.md)
 - [SelectObjectContentEventStreamProgress](docs/SelectObjectContentEventStreamProgress.md)
 - [SelectObjectContentEventStreamRecords](docs/SelectObjectContentEventStreamRecords.md)
 - [SelectObjectContentEventStreamStats](docs/SelectObjectContentEventStreamStats.md)
 - [SelectObjectContentOutput](docs/SelectObjectContentOutput.md)
 - [SelectObjectContentOutputPayload](docs/SelectObjectContentOutputPayload.md)
 - [SelectObjectContentRequest](docs/SelectObjectContentRequest.md)
 - [SelectObjectContentRequestInputSerialization](docs/SelectObjectContentRequestInputSerialization.md)
 - [SelectObjectContentRequestOutputSerialization](docs/SelectObjectContentRequestOutputSerialization.md)
 - [SelectObjectContentRequestRequestProgress](docs/SelectObjectContentRequestRequestProgress.md)
 - [SelectObjectContentRequestScanRange](docs/SelectObjectContentRequestScanRange.md)
 - [SelectParameters](docs/SelectParameters.md)
 - [SelectParametersInputSerialization](docs/SelectParametersInputSerialization.md)
 - [SelectParametersOutputSerialization](docs/SelectParametersOutputSerialization.md)
 - [ServerSideEncryption](docs/ServerSideEncryption.md)
 - [ServerSideEncryptionByDefault](docs/ServerSideEncryptionByDefault.md)
 - [ServerSideEncryptionConfiguration](docs/ServerSideEncryptionConfiguration.md)
 - [ServerSideEncryptionRule](docs/ServerSideEncryptionRule.md)
 - [ServerSideEncryptionRuleApplyServerSideEncryptionByDefault](docs/ServerSideEncryptionRuleApplyServerSideEncryptionByDefault.md)
 - [SourceSelectionCriteria](docs/SourceSelectionCriteria.md)
 - [SourceSelectionCriteriaReplicaModifications](docs/SourceSelectionCriteriaReplicaModifications.md)
 - [SourceSelectionCriteriaSseKmsEncryptedObjects](docs/SourceSelectionCriteriaSseKmsEncryptedObjects.md)
 - [SseKmsEncryptedObjects](docs/SseKmsEncryptedObjects.md)
 - [SseKmsEncryptedObjectsStatus](docs/SseKmsEncryptedObjectsStatus.md)
 - [Stats](docs/Stats.md)
 - [StatsEvent](docs/StatsEvent.md)
 - [StatsEventDetails](docs/StatsEventDetails.md)
 - [StorageClass](docs/StorageClass.md)
 - [StorageClassAnalysisDataExport](docs/StorageClassAnalysisDataExport.md)
 - [StorageClassAnalysisDataExportDestination](docs/StorageClassAnalysisDataExportDestination.md)
 - [StorageClassAnalysisSchemaVersion](docs/StorageClassAnalysisSchemaVersion.md)
 - [Tag](docs/Tag.md)
 - [TagSetInner](docs/TagSetInner.md)
 - [Tagging](docs/Tagging.md)
 - [TaggingDirective](docs/TaggingDirective.md)
 - [TargetGrant](docs/TargetGrant.md)
 - [TargetGrantGrantee](docs/TargetGrantGrantee.md)
 - [TargetGrantsInner](docs/TargetGrantsInner.md)
 - [Tier](docs/Tier.md)
 - [Tiering](docs/Tiering.md)
 - [TopicConfiguration](docs/TopicConfiguration.md)
 - [Transition](docs/Transition.md)
 - [TransitionStorageClass](docs/TransitionStorageClass.md)
 - [Type](docs/Type.md)
 - [UploadPartCopyOutput](docs/UploadPartCopyOutput.md)
 - [UploadPartCopyOutputCopyPartResult](docs/UploadPartCopyOutputCopyPartResult.md)
 - [UploadPartRequest](docs/UploadPartRequest.md)
 - [UserMetadataInner](docs/UserMetadataInner.md)
 - [VersioningConfiguration](docs/VersioningConfiguration.md)
 - [WebsiteConfiguration](docs/WebsiteConfiguration.md)
 - [WriteGetObjectResponseRequest](docs/WriteGetObjectResponseRequest.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### hmac

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```golang
auth := context.WithValue(
		context.Background(),
		sw.ContextAPIKeys,
		map[string]sw.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



