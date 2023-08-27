# V1TagEventCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Generation** | **int64** | Generation is the spec tag generation that this status corresponds to | 
**LastTransitionTime** | Pointer to **string** | LastTransitionTIme is the time the condition transitioned from one status to another. | [optional] 
**Message** | Pointer to **string** | Message is a human readable description of the details about last transition, complementing reason. | [optional] 
**Reason** | Pointer to **string** | Reason is a brief machine readable explanation for the condition&#39;s last transition. | [optional] 
**Status** | **string** | Status of the condition, one of True, False, Unknown. | 
**Type** | **string** | Type of tag event condition, currently only ImportSuccess | 

## Methods

### NewV1TagEventCondition

`func NewV1TagEventCondition(generation int64, status string, type_ string, ) *V1TagEventCondition`

NewV1TagEventCondition instantiates a new V1TagEventCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TagEventConditionWithDefaults

`func NewV1TagEventConditionWithDefaults() *V1TagEventCondition`

NewV1TagEventConditionWithDefaults instantiates a new V1TagEventCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneration

`func (o *V1TagEventCondition) GetGeneration() int64`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *V1TagEventCondition) GetGenerationOk() (*int64, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *V1TagEventCondition) SetGeneration(v int64)`

SetGeneration sets Generation field to given value.


### GetLastTransitionTime

`func (o *V1TagEventCondition) GetLastTransitionTime() string`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *V1TagEventCondition) GetLastTransitionTimeOk() (*string, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *V1TagEventCondition) SetLastTransitionTime(v string)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *V1TagEventCondition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMessage

`func (o *V1TagEventCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1TagEventCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1TagEventCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1TagEventCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *V1TagEventCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V1TagEventCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V1TagEventCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V1TagEventCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *V1TagEventCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1TagEventCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1TagEventCondition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *V1TagEventCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1TagEventCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1TagEventCondition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


