# V1SignatureCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastProbeTime** | Pointer to **string** | Last time the condition was checked. | [optional] 
**LastTransitionTime** | Pointer to **string** | Last time the condition transit from one status to another. | [optional] 
**Message** | Pointer to **string** | Human readable message indicating details about last transition. | [optional] 
**Reason** | Pointer to **string** | (brief) reason for the condition&#39;s last transition. | [optional] 
**Status** | **string** | Status of the condition, one of True, False, Unknown. | 
**Type** | **string** | Type of signature condition, Complete or Failed. | 

## Methods

### NewV1SignatureCondition

`func NewV1SignatureCondition(status string, type_ string, ) *V1SignatureCondition`

NewV1SignatureCondition instantiates a new V1SignatureCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SignatureConditionWithDefaults

`func NewV1SignatureConditionWithDefaults() *V1SignatureCondition`

NewV1SignatureConditionWithDefaults instantiates a new V1SignatureCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastProbeTime

`func (o *V1SignatureCondition) GetLastProbeTime() string`

GetLastProbeTime returns the LastProbeTime field if non-nil, zero value otherwise.

### GetLastProbeTimeOk

`func (o *V1SignatureCondition) GetLastProbeTimeOk() (*string, bool)`

GetLastProbeTimeOk returns a tuple with the LastProbeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProbeTime

`func (o *V1SignatureCondition) SetLastProbeTime(v string)`

SetLastProbeTime sets LastProbeTime field to given value.

### HasLastProbeTime

`func (o *V1SignatureCondition) HasLastProbeTime() bool`

HasLastProbeTime returns a boolean if a field has been set.

### GetLastTransitionTime

`func (o *V1SignatureCondition) GetLastTransitionTime() string`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *V1SignatureCondition) GetLastTransitionTimeOk() (*string, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *V1SignatureCondition) SetLastTransitionTime(v string)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *V1SignatureCondition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMessage

`func (o *V1SignatureCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1SignatureCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1SignatureCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1SignatureCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *V1SignatureCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V1SignatureCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V1SignatureCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V1SignatureCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *V1SignatureCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1SignatureCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1SignatureCondition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *V1SignatureCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1SignatureCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1SignatureCondition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


