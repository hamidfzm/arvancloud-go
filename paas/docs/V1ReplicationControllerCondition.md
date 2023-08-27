# V1ReplicationControllerCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTransitionTime** | Pointer to **string** | The last time the condition transitioned from one status to another. | [optional] 
**Message** | Pointer to **string** | A human readable message indicating details about the transition. | [optional] 
**Reason** | Pointer to **string** | The reason for the condition&#39;s last transition. | [optional] 
**Status** | **string** | Status of the condition, one of True, False, Unknown. | 
**Type** | **string** | Type of replication controller condition. | 

## Methods

### NewV1ReplicationControllerCondition

`func NewV1ReplicationControllerCondition(status string, type_ string, ) *V1ReplicationControllerCondition`

NewV1ReplicationControllerCondition instantiates a new V1ReplicationControllerCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ReplicationControllerConditionWithDefaults

`func NewV1ReplicationControllerConditionWithDefaults() *V1ReplicationControllerCondition`

NewV1ReplicationControllerConditionWithDefaults instantiates a new V1ReplicationControllerCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTransitionTime

`func (o *V1ReplicationControllerCondition) GetLastTransitionTime() string`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *V1ReplicationControllerCondition) GetLastTransitionTimeOk() (*string, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *V1ReplicationControllerCondition) SetLastTransitionTime(v string)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *V1ReplicationControllerCondition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMessage

`func (o *V1ReplicationControllerCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1ReplicationControllerCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1ReplicationControllerCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1ReplicationControllerCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *V1ReplicationControllerCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V1ReplicationControllerCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V1ReplicationControllerCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V1ReplicationControllerCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *V1ReplicationControllerCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1ReplicationControllerCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1ReplicationControllerCondition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *V1ReplicationControllerCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1ReplicationControllerCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1ReplicationControllerCondition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


