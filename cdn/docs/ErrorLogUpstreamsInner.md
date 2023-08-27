# ErrorLogUpstreamsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The upstream&#39;s address | [optional] 
**Count** | Pointer to **int32** | Error count per upstream | [optional] 

## Methods

### NewErrorLogUpstreamsInner

`func NewErrorLogUpstreamsInner() *ErrorLogUpstreamsInner`

NewErrorLogUpstreamsInner instantiates a new ErrorLogUpstreamsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorLogUpstreamsInnerWithDefaults

`func NewErrorLogUpstreamsInnerWithDefaults() *ErrorLogUpstreamsInner`

NewErrorLogUpstreamsInnerWithDefaults instantiates a new ErrorLogUpstreamsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ErrorLogUpstreamsInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ErrorLogUpstreamsInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ErrorLogUpstreamsInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ErrorLogUpstreamsInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCount

`func (o *ErrorLogUpstreamsInner) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ErrorLogUpstreamsInner) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ErrorLogUpstreamsInner) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ErrorLogUpstreamsInner) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


