# ListObjectsV2Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsTruncated** | Pointer to **bool** |  | [optional] 
**Contents** | Pointer to [**Array**](array.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**Delimiter** | Pointer to **string** |  | [optional] 
**MaxKeys** | Pointer to **int32** |  | [optional] 
**CommonPrefixes** | Pointer to [**Array**](array.md) |  | [optional] 
**EncodingType** | Pointer to [**EncodingType**](EncodingType.md) |  | [optional] 
**KeyCount** | Pointer to **int32** |  | [optional] 
**ContinuationToken** | Pointer to **string** |  | [optional] 
**NextContinuationToken** | Pointer to **string** |  | [optional] 
**StartAfter** | Pointer to **string** |  | [optional] 

## Methods

### NewListObjectsV2Output

`func NewListObjectsV2Output() *ListObjectsV2Output`

NewListObjectsV2Output instantiates a new ListObjectsV2Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListObjectsV2OutputWithDefaults

`func NewListObjectsV2OutputWithDefaults() *ListObjectsV2Output`

NewListObjectsV2OutputWithDefaults instantiates a new ListObjectsV2Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsTruncated

`func (o *ListObjectsV2Output) GetIsTruncated() bool`

GetIsTruncated returns the IsTruncated field if non-nil, zero value otherwise.

### GetIsTruncatedOk

`func (o *ListObjectsV2Output) GetIsTruncatedOk() (*bool, bool)`

GetIsTruncatedOk returns a tuple with the IsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTruncated

`func (o *ListObjectsV2Output) SetIsTruncated(v bool)`

SetIsTruncated sets IsTruncated field to given value.

### HasIsTruncated

`func (o *ListObjectsV2Output) HasIsTruncated() bool`

HasIsTruncated returns a boolean if a field has been set.

### GetContents

`func (o *ListObjectsV2Output) GetContents() Array`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ListObjectsV2Output) GetContentsOk() (*Array, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ListObjectsV2Output) SetContents(v Array)`

SetContents sets Contents field to given value.

### HasContents

`func (o *ListObjectsV2Output) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetName

`func (o *ListObjectsV2Output) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListObjectsV2Output) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListObjectsV2Output) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListObjectsV2Output) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrefix

`func (o *ListObjectsV2Output) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ListObjectsV2Output) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ListObjectsV2Output) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ListObjectsV2Output) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetDelimiter

`func (o *ListObjectsV2Output) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *ListObjectsV2Output) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *ListObjectsV2Output) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.

### HasDelimiter

`func (o *ListObjectsV2Output) HasDelimiter() bool`

HasDelimiter returns a boolean if a field has been set.

### GetMaxKeys

`func (o *ListObjectsV2Output) GetMaxKeys() int32`

GetMaxKeys returns the MaxKeys field if non-nil, zero value otherwise.

### GetMaxKeysOk

`func (o *ListObjectsV2Output) GetMaxKeysOk() (*int32, bool)`

GetMaxKeysOk returns a tuple with the MaxKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxKeys

`func (o *ListObjectsV2Output) SetMaxKeys(v int32)`

SetMaxKeys sets MaxKeys field to given value.

### HasMaxKeys

`func (o *ListObjectsV2Output) HasMaxKeys() bool`

HasMaxKeys returns a boolean if a field has been set.

### GetCommonPrefixes

`func (o *ListObjectsV2Output) GetCommonPrefixes() Array`

GetCommonPrefixes returns the CommonPrefixes field if non-nil, zero value otherwise.

### GetCommonPrefixesOk

`func (o *ListObjectsV2Output) GetCommonPrefixesOk() (*Array, bool)`

GetCommonPrefixesOk returns a tuple with the CommonPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonPrefixes

`func (o *ListObjectsV2Output) SetCommonPrefixes(v Array)`

SetCommonPrefixes sets CommonPrefixes field to given value.

### HasCommonPrefixes

`func (o *ListObjectsV2Output) HasCommonPrefixes() bool`

HasCommonPrefixes returns a boolean if a field has been set.

### GetEncodingType

`func (o *ListObjectsV2Output) GetEncodingType() EncodingType`

GetEncodingType returns the EncodingType field if non-nil, zero value otherwise.

### GetEncodingTypeOk

`func (o *ListObjectsV2Output) GetEncodingTypeOk() (*EncodingType, bool)`

GetEncodingTypeOk returns a tuple with the EncodingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingType

`func (o *ListObjectsV2Output) SetEncodingType(v EncodingType)`

SetEncodingType sets EncodingType field to given value.

### HasEncodingType

`func (o *ListObjectsV2Output) HasEncodingType() bool`

HasEncodingType returns a boolean if a field has been set.

### GetKeyCount

`func (o *ListObjectsV2Output) GetKeyCount() int32`

GetKeyCount returns the KeyCount field if non-nil, zero value otherwise.

### GetKeyCountOk

`func (o *ListObjectsV2Output) GetKeyCountOk() (*int32, bool)`

GetKeyCountOk returns a tuple with the KeyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyCount

`func (o *ListObjectsV2Output) SetKeyCount(v int32)`

SetKeyCount sets KeyCount field to given value.

### HasKeyCount

`func (o *ListObjectsV2Output) HasKeyCount() bool`

HasKeyCount returns a boolean if a field has been set.

### GetContinuationToken

`func (o *ListObjectsV2Output) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *ListObjectsV2Output) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *ListObjectsV2Output) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *ListObjectsV2Output) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### GetNextContinuationToken

`func (o *ListObjectsV2Output) GetNextContinuationToken() string`

GetNextContinuationToken returns the NextContinuationToken field if non-nil, zero value otherwise.

### GetNextContinuationTokenOk

`func (o *ListObjectsV2Output) GetNextContinuationTokenOk() (*string, bool)`

GetNextContinuationTokenOk returns a tuple with the NextContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextContinuationToken

`func (o *ListObjectsV2Output) SetNextContinuationToken(v string)`

SetNextContinuationToken sets NextContinuationToken field to given value.

### HasNextContinuationToken

`func (o *ListObjectsV2Output) HasNextContinuationToken() bool`

HasNextContinuationToken returns a boolean if a field has been set.

### GetStartAfter

`func (o *ListObjectsV2Output) GetStartAfter() string`

GetStartAfter returns the StartAfter field if non-nil, zero value otherwise.

### GetStartAfterOk

`func (o *ListObjectsV2Output) GetStartAfterOk() (*string, bool)`

GetStartAfterOk returns a tuple with the StartAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAfter

`func (o *ListObjectsV2Output) SetStartAfter(v string)`

SetStartAfter sets StartAfter field to given value.

### HasStartAfter

`func (o *ListObjectsV2Output) HasStartAfter() bool`

HasStartAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


