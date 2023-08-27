# AnalyticsExportDestinationS3BucketDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | [**AnalyticsS3ExportFileFormat**](AnalyticsS3ExportFileFormat.md) |  | 
**BucketAccountId** | Pointer to **string** |  | [optional] 
**Bucket** | **string** |  | 
**Prefix** | Pointer to **string** |  | [optional] 

## Methods

### NewAnalyticsExportDestinationS3BucketDestination

`func NewAnalyticsExportDestinationS3BucketDestination(format AnalyticsS3ExportFileFormat, bucket string, ) *AnalyticsExportDestinationS3BucketDestination`

NewAnalyticsExportDestinationS3BucketDestination instantiates a new AnalyticsExportDestinationS3BucketDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsExportDestinationS3BucketDestinationWithDefaults

`func NewAnalyticsExportDestinationS3BucketDestinationWithDefaults() *AnalyticsExportDestinationS3BucketDestination`

NewAnalyticsExportDestinationS3BucketDestinationWithDefaults instantiates a new AnalyticsExportDestinationS3BucketDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *AnalyticsExportDestinationS3BucketDestination) GetFormat() AnalyticsS3ExportFileFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *AnalyticsExportDestinationS3BucketDestination) GetFormatOk() (*AnalyticsS3ExportFileFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *AnalyticsExportDestinationS3BucketDestination) SetFormat(v AnalyticsS3ExportFileFormat)`

SetFormat sets Format field to given value.


### GetBucketAccountId

`func (o *AnalyticsExportDestinationS3BucketDestination) GetBucketAccountId() string`

GetBucketAccountId returns the BucketAccountId field if non-nil, zero value otherwise.

### GetBucketAccountIdOk

`func (o *AnalyticsExportDestinationS3BucketDestination) GetBucketAccountIdOk() (*string, bool)`

GetBucketAccountIdOk returns a tuple with the BucketAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketAccountId

`func (o *AnalyticsExportDestinationS3BucketDestination) SetBucketAccountId(v string)`

SetBucketAccountId sets BucketAccountId field to given value.

### HasBucketAccountId

`func (o *AnalyticsExportDestinationS3BucketDestination) HasBucketAccountId() bool`

HasBucketAccountId returns a boolean if a field has been set.

### GetBucket

`func (o *AnalyticsExportDestinationS3BucketDestination) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *AnalyticsExportDestinationS3BucketDestination) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *AnalyticsExportDestinationS3BucketDestination) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetPrefix

`func (o *AnalyticsExportDestinationS3BucketDestination) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *AnalyticsExportDestinationS3BucketDestination) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *AnalyticsExportDestinationS3BucketDestination) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *AnalyticsExportDestinationS3BucketDestination) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


