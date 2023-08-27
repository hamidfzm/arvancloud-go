# ListObjectsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsTruncated** | Pointer to **bool** |  | [optional] 
**Marker** | Pointer to **string** |  | [optional] 
**NextMarker** | Pointer to **string** |  | [optional] 
**Contents** | Pointer to [**Array**](array.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**Delimiter** | Pointer to **string** |  | [optional] 
**MaxKeys** | Pointer to **int32** |  | [optional] 
**CommonPrefixes** | Pointer to [**Array**](array.md) |  | [optional] 
**EncodingType** | Pointer to [**EncodingType**](EncodingType.md) |  | [optional] 

## Methods

### NewListObjectsOutput

`func NewListObjectsOutput() *ListObjectsOutput`

NewListObjectsOutput instantiates a new ListObjectsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListObjectsOutputWithDefaults

`func NewListObjectsOutputWithDefaults() *ListObjectsOutput`

NewListObjectsOutputWithDefaults instantiates a new ListObjectsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsTruncated

`func (o *ListObjectsOutput) GetIsTruncated() bool`

GetIsTruncated returns the IsTruncated field if non-nil, zero value otherwise.

### GetIsTruncatedOk

`func (o *ListObjectsOutput) GetIsTruncatedOk() (*bool, bool)`

GetIsTruncatedOk returns a tuple with the IsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTruncated

`func (o *ListObjectsOutput) SetIsTruncated(v bool)`

SetIsTruncated sets IsTruncated field to given value.

### HasIsTruncated

`func (o *ListObjectsOutput) HasIsTruncated() bool`

HasIsTruncated returns a boolean if a field has been set.

### GetMarker

`func (o *ListObjectsOutput) GetMarker() string`

GetMarker returns the Marker field if non-nil, zero value otherwise.

### GetMarkerOk

`func (o *ListObjectsOutput) GetMarkerOk() (*string, bool)`

GetMarkerOk returns a tuple with the Marker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarker

`func (o *ListObjectsOutput) SetMarker(v string)`

SetMarker sets Marker field to given value.

### HasMarker

`func (o *ListObjectsOutput) HasMarker() bool`

HasMarker returns a boolean if a field has been set.

### GetNextMarker

`func (o *ListObjectsOutput) GetNextMarker() string`

GetNextMarker returns the NextMarker field if non-nil, zero value otherwise.

### GetNextMarkerOk

`func (o *ListObjectsOutput) GetNextMarkerOk() (*string, bool)`

GetNextMarkerOk returns a tuple with the NextMarker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextMarker

`func (o *ListObjectsOutput) SetNextMarker(v string)`

SetNextMarker sets NextMarker field to given value.

### HasNextMarker

`func (o *ListObjectsOutput) HasNextMarker() bool`

HasNextMarker returns a boolean if a field has been set.

### GetContents

`func (o *ListObjectsOutput) GetContents() Array`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ListObjectsOutput) GetContentsOk() (*Array, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ListObjectsOutput) SetContents(v Array)`

SetContents sets Contents field to given value.

### HasContents

`func (o *ListObjectsOutput) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetName

`func (o *ListObjectsOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListObjectsOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListObjectsOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListObjectsOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrefix

`func (o *ListObjectsOutput) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ListObjectsOutput) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ListObjectsOutput) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ListObjectsOutput) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetDelimiter

`func (o *ListObjectsOutput) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *ListObjectsOutput) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *ListObjectsOutput) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.

### HasDelimiter

`func (o *ListObjectsOutput) HasDelimiter() bool`

HasDelimiter returns a boolean if a field has been set.

### GetMaxKeys

`func (o *ListObjectsOutput) GetMaxKeys() int32`

GetMaxKeys returns the MaxKeys field if non-nil, zero value otherwise.

### GetMaxKeysOk

`func (o *ListObjectsOutput) GetMaxKeysOk() (*int32, bool)`

GetMaxKeysOk returns a tuple with the MaxKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxKeys

`func (o *ListObjectsOutput) SetMaxKeys(v int32)`

SetMaxKeys sets MaxKeys field to given value.

### HasMaxKeys

`func (o *ListObjectsOutput) HasMaxKeys() bool`

HasMaxKeys returns a boolean if a field has been set.

### GetCommonPrefixes

`func (o *ListObjectsOutput) GetCommonPrefixes() Array`

GetCommonPrefixes returns the CommonPrefixes field if non-nil, zero value otherwise.

### GetCommonPrefixesOk

`func (o *ListObjectsOutput) GetCommonPrefixesOk() (*Array, bool)`

GetCommonPrefixesOk returns a tuple with the CommonPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonPrefixes

`func (o *ListObjectsOutput) SetCommonPrefixes(v Array)`

SetCommonPrefixes sets CommonPrefixes field to given value.

### HasCommonPrefixes

`func (o *ListObjectsOutput) HasCommonPrefixes() bool`

HasCommonPrefixes returns a boolean if a field has been set.

### GetEncodingType

`func (o *ListObjectsOutput) GetEncodingType() EncodingType`

GetEncodingType returns the EncodingType field if non-nil, zero value otherwise.

### GetEncodingTypeOk

`func (o *ListObjectsOutput) GetEncodingTypeOk() (*EncodingType, bool)`

GetEncodingTypeOk returns a tuple with the EncodingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingType

`func (o *ListObjectsOutput) SetEncodingType(v EncodingType)`

SetEncodingType sets EncodingType field to given value.

### HasEncodingType

`func (o *ListObjectsOutput) HasEncodingType() bool`

HasEncodingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


