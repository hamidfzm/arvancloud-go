# V1ComponentCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Condition error code for a component. For example, a health check error code. | [optional] 
**Message** | Pointer to **string** | Message about the condition for a component. For example, information about a health check. | [optional] 
**Status** | **string** | Status of the condition for a component. Valid values for \&quot;Healthy\&quot;: \&quot;True\&quot;, \&quot;False\&quot;, or \&quot;Unknown\&quot;. | 
**Type** | **string** | Type of condition for a component. Valid value: \&quot;Healthy\&quot; | 

## Methods

### NewV1ComponentCondition

`func NewV1ComponentCondition(status string, type_ string, ) *V1ComponentCondition`

NewV1ComponentCondition instantiates a new V1ComponentCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ComponentConditionWithDefaults

`func NewV1ComponentConditionWithDefaults() *V1ComponentCondition`

NewV1ComponentConditionWithDefaults instantiates a new V1ComponentCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *V1ComponentCondition) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V1ComponentCondition) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V1ComponentCondition) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V1ComponentCondition) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *V1ComponentCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1ComponentCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1ComponentCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1ComponentCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *V1ComponentCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1ComponentCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1ComponentCondition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *V1ComponentCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1ComponentCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1ComponentCondition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


