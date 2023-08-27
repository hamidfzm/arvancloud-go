# ListMultipartUploadsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** |  | [optional] 
**KeyMarker** | Pointer to **string** |  | [optional] 
**UploadIdMarker** | Pointer to **string** |  | [optional] 
**NextKeyMarker** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**Delimiter** | Pointer to **string** |  | [optional] 
**NextUploadIdMarker** | Pointer to **string** |  | [optional] 
**MaxUploads** | Pointer to **int32** |  | [optional] 
**IsTruncated** | Pointer to **bool** |  | [optional] 
**Uploads** | Pointer to [**Array**](array.md) |  | [optional] 
**CommonPrefixes** | Pointer to [**Array**](array.md) |  | [optional] 
**EncodingType** | Pointer to [**EncodingType**](EncodingType.md) |  | [optional] 

## Methods

### NewListMultipartUploadsOutput

`func NewListMultipartUploadsOutput() *ListMultipartUploadsOutput`

NewListMultipartUploadsOutput instantiates a new ListMultipartUploadsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMultipartUploadsOutputWithDefaults

`func NewListMultipartUploadsOutputWithDefaults() *ListMultipartUploadsOutput`

NewListMultipartUploadsOutputWithDefaults instantiates a new ListMultipartUploadsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *ListMultipartUploadsOutput) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ListMultipartUploadsOutput) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ListMultipartUploadsOutput) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *ListMultipartUploadsOutput) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetKeyMarker

`func (o *ListMultipartUploadsOutput) GetKeyMarker() string`

GetKeyMarker returns the KeyMarker field if non-nil, zero value otherwise.

### GetKeyMarkerOk

`func (o *ListMultipartUploadsOutput) GetKeyMarkerOk() (*string, bool)`

GetKeyMarkerOk returns a tuple with the KeyMarker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyMarker

`func (o *ListMultipartUploadsOutput) SetKeyMarker(v string)`

SetKeyMarker sets KeyMarker field to given value.

### HasKeyMarker

`func (o *ListMultipartUploadsOutput) HasKeyMarker() bool`

HasKeyMarker returns a boolean if a field has been set.

### GetUploadIdMarker

`func (o *ListMultipartUploadsOutput) GetUploadIdMarker() string`

GetUploadIdMarker returns the UploadIdMarker field if non-nil, zero value otherwise.

### GetUploadIdMarkerOk

`func (o *ListMultipartUploadsOutput) GetUploadIdMarkerOk() (*string, bool)`

GetUploadIdMarkerOk returns a tuple with the UploadIdMarker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadIdMarker

`func (o *ListMultipartUploadsOutput) SetUploadIdMarker(v string)`

SetUploadIdMarker sets UploadIdMarker field to given value.

### HasUploadIdMarker

`func (o *ListMultipartUploadsOutput) HasUploadIdMarker() bool`

HasUploadIdMarker returns a boolean if a field has been set.

### GetNextKeyMarker

`func (o *ListMultipartUploadsOutput) GetNextKeyMarker() string`

GetNextKeyMarker returns the NextKeyMarker field if non-nil, zero value otherwise.

### GetNextKeyMarkerOk

`func (o *ListMultipartUploadsOutput) GetNextKeyMarkerOk() (*string, bool)`

GetNextKeyMarkerOk returns a tuple with the NextKeyMarker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextKeyMarker

`func (o *ListMultipartUploadsOutput) SetNextKeyMarker(v string)`

SetNextKeyMarker sets NextKeyMarker field to given value.

### HasNextKeyMarker

`func (o *ListMultipartUploadsOutput) HasNextKeyMarker() bool`

HasNextKeyMarker returns a boolean if a field has been set.

### GetPrefix

`func (o *ListMultipartUploadsOutput) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ListMultipartUploadsOutput) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ListMultipartUploadsOutput) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ListMultipartUploadsOutput) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetDelimiter

`func (o *ListMultipartUploadsOutput) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *ListMultipartUploadsOutput) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *ListMultipartUploadsOutput) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.

### HasDelimiter

`func (o *ListMultipartUploadsOutput) HasDelimiter() bool`

HasDelimiter returns a boolean if a field has been set.

### GetNextUploadIdMarker

`func (o *ListMultipartUploadsOutput) GetNextUploadIdMarker() string`

GetNextUploadIdMarker returns the NextUploadIdMarker field if non-nil, zero value otherwise.

### GetNextUploadIdMarkerOk

`func (o *ListMultipartUploadsOutput) GetNextUploadIdMarkerOk() (*string, bool)`

GetNextUploadIdMarkerOk returns a tuple with the NextUploadIdMarker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUploadIdMarker

`func (o *ListMultipartUploadsOutput) SetNextUploadIdMarker(v string)`

SetNextUploadIdMarker sets NextUploadIdMarker field to given value.

### HasNextUploadIdMarker

`func (o *ListMultipartUploadsOutput) HasNextUploadIdMarker() bool`

HasNextUploadIdMarker returns a boolean if a field has been set.

### GetMaxUploads

`func (o *ListMultipartUploadsOutput) GetMaxUploads() int32`

GetMaxUploads returns the MaxUploads field if non-nil, zero value otherwise.

### GetMaxUploadsOk

`func (o *ListMultipartUploadsOutput) GetMaxUploadsOk() (*int32, bool)`

GetMaxUploadsOk returns a tuple with the MaxUploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUploads

`func (o *ListMultipartUploadsOutput) SetMaxUploads(v int32)`

SetMaxUploads sets MaxUploads field to given value.

### HasMaxUploads

`func (o *ListMultipartUploadsOutput) HasMaxUploads() bool`

HasMaxUploads returns a boolean if a field has been set.

### GetIsTruncated

`func (o *ListMultipartUploadsOutput) GetIsTruncated() bool`

GetIsTruncated returns the IsTruncated field if non-nil, zero value otherwise.

### GetIsTruncatedOk

`func (o *ListMultipartUploadsOutput) GetIsTruncatedOk() (*bool, bool)`

GetIsTruncatedOk returns a tuple with the IsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTruncated

`func (o *ListMultipartUploadsOutput) SetIsTruncated(v bool)`

SetIsTruncated sets IsTruncated field to given value.

### HasIsTruncated

`func (o *ListMultipartUploadsOutput) HasIsTruncated() bool`

HasIsTruncated returns a boolean if a field has been set.

### GetUploads

`func (o *ListMultipartUploadsOutput) GetUploads() Array`

GetUploads returns the Uploads field if non-nil, zero value otherwise.

### GetUploadsOk

`func (o *ListMultipartUploadsOutput) GetUploadsOk() (*Array, bool)`

GetUploadsOk returns a tuple with the Uploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploads

`func (o *ListMultipartUploadsOutput) SetUploads(v Array)`

SetUploads sets Uploads field to given value.

### HasUploads

`func (o *ListMultipartUploadsOutput) HasUploads() bool`

HasUploads returns a boolean if a field has been set.

### GetCommonPrefixes

`func (o *ListMultipartUploadsOutput) GetCommonPrefixes() Array`

GetCommonPrefixes returns the CommonPrefixes field if non-nil, zero value otherwise.

### GetCommonPrefixesOk

`func (o *ListMultipartUploadsOutput) GetCommonPrefixesOk() (*Array, bool)`

GetCommonPrefixesOk returns a tuple with the CommonPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonPrefixes

`func (o *ListMultipartUploadsOutput) SetCommonPrefixes(v Array)`

SetCommonPrefixes sets CommonPrefixes field to given value.

### HasCommonPrefixes

`func (o *ListMultipartUploadsOutput) HasCommonPrefixes() bool`

HasCommonPrefixes returns a boolean if a field has been set.

### GetEncodingType

`func (o *ListMultipartUploadsOutput) GetEncodingType() EncodingType`

GetEncodingType returns the EncodingType field if non-nil, zero value otherwise.

### GetEncodingTypeOk

`func (o *ListMultipartUploadsOutput) GetEncodingTypeOk() (*EncodingType, bool)`

GetEncodingTypeOk returns a tuple with the EncodingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingType

`func (o *ListMultipartUploadsOutput) SetEncodingType(v EncodingType)`

SetEncodingType sets EncodingType field to given value.

### HasEncodingType

`func (o *ListMultipartUploadsOutput) HasEncodingType() bool`

HasEncodingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


