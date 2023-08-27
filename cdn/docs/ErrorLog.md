# ErrorLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The error message | [optional] 
**Count** | Pointer to **int32** | The error&#39;s count | [optional] 
**Upstreams** | Pointer to [**[]ErrorLogUpstreamsInner**](ErrorLogUpstreamsInner.md) |  | [optional] 

## Methods

### NewErrorLog

`func NewErrorLog() *ErrorLog`

NewErrorLog instantiates a new ErrorLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorLogWithDefaults

`func NewErrorLogWithDefaults() *ErrorLog`

NewErrorLogWithDefaults instantiates a new ErrorLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ErrorLog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ErrorLog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ErrorLog) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ErrorLog) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCount

`func (o *ErrorLog) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ErrorLog) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ErrorLog) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ErrorLog) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetUpstreams

`func (o *ErrorLog) GetUpstreams() []ErrorLogUpstreamsInner`

GetUpstreams returns the Upstreams field if non-nil, zero value otherwise.

### GetUpstreamsOk

`func (o *ErrorLog) GetUpstreamsOk() (*[]ErrorLogUpstreamsInner, bool)`

GetUpstreamsOk returns a tuple with the Upstreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreams

`func (o *ErrorLog) SetUpstreams(v []ErrorLogUpstreamsInner)`

SetUpstreams sets Upstreams field to given value.

### HasUpstreams

`func (o *ErrorLog) HasUpstreams() bool`

HasUpstreams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


