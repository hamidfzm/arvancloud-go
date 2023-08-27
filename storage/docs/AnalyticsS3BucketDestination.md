# AnalyticsS3BucketDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | [**AnalyticsS3ExportFileFormat**](AnalyticsS3ExportFileFormat.md) |  | 
**BucketAccountId** | Pointer to **string** |  | [optional] 
**Bucket** | **string** |  | 
**Prefix** | Pointer to **string** |  | [optional] 

## Methods

### NewAnalyticsS3BucketDestination

`func NewAnalyticsS3BucketDestination(format AnalyticsS3ExportFileFormat, bucket string, ) *AnalyticsS3BucketDestination`

NewAnalyticsS3BucketDestination instantiates a new AnalyticsS3BucketDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsS3BucketDestinationWithDefaults

`func NewAnalyticsS3BucketDestinationWithDefaults() *AnalyticsS3BucketDestination`

NewAnalyticsS3BucketDestinationWithDefaults instantiates a new AnalyticsS3BucketDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *AnalyticsS3BucketDestination) GetFormat() AnalyticsS3ExportFileFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *AnalyticsS3BucketDestination) GetFormatOk() (*AnalyticsS3ExportFileFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *AnalyticsS3BucketDestination) SetFormat(v AnalyticsS3ExportFileFormat)`

SetFormat sets Format field to given value.


### GetBucketAccountId

`func (o *AnalyticsS3BucketDestination) GetBucketAccountId() string`

GetBucketAccountId returns the BucketAccountId field if non-nil, zero value otherwise.

### GetBucketAccountIdOk

`func (o *AnalyticsS3BucketDestination) GetBucketAccountIdOk() (*string, bool)`

GetBucketAccountIdOk returns a tuple with the BucketAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketAccountId

`func (o *AnalyticsS3BucketDestination) SetBucketAccountId(v string)`

SetBucketAccountId sets BucketAccountId field to given value.

### HasBucketAccountId

`func (o *AnalyticsS3BucketDestination) HasBucketAccountId() bool`

HasBucketAccountId returns a boolean if a field has been set.

### GetBucket

`func (o *AnalyticsS3BucketDestination) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *AnalyticsS3BucketDestination) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *AnalyticsS3BucketDestination) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetPrefix

`func (o *AnalyticsS3BucketDestination) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *AnalyticsS3BucketDestination) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *AnalyticsS3BucketDestination) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *AnalyticsS3BucketDestination) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


