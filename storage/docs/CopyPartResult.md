# CopyPartResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ETag** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### NewCopyPartResult

`func NewCopyPartResult() *CopyPartResult`

NewCopyPartResult instantiates a new CopyPartResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyPartResultWithDefaults

`func NewCopyPartResultWithDefaults() *CopyPartResult`

NewCopyPartResultWithDefaults instantiates a new CopyPartResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetETag

`func (o *CopyPartResult) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *CopyPartResult) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *CopyPartResult) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *CopyPartResult) HasETag() bool`

HasETag returns a boolean if a field has been set.

### GetLastModified

`func (o *CopyPartResult) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *CopyPartResult) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *CopyPartResult) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *CopyPartResult) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


