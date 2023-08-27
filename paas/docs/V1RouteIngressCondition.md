# V1RouteIngressCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTransitionTime** | Pointer to **string** | RFC 3339 date and time when this condition last transitioned | [optional] 
**Message** | Pointer to **string** | Human readable message indicating details about last transition. | [optional] 
**Reason** | Pointer to **string** | (brief) reason for the condition&#39;s last transition, and is usually a machine and human readable constant | [optional] 
**Status** | **string** | Status is the status of the condition. Can be True, False, Unknown. | 
**Type** | **string** | Type is the type of the condition. Currently only Ready. | 

## Methods

### NewV1RouteIngressCondition

`func NewV1RouteIngressCondition(status string, type_ string, ) *V1RouteIngressCondition`

NewV1RouteIngressCondition instantiates a new V1RouteIngressCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RouteIngressConditionWithDefaults

`func NewV1RouteIngressConditionWithDefaults() *V1RouteIngressCondition`

NewV1RouteIngressConditionWithDefaults instantiates a new V1RouteIngressCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTransitionTime

`func (o *V1RouteIngressCondition) GetLastTransitionTime() string`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *V1RouteIngressCondition) GetLastTransitionTimeOk() (*string, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *V1RouteIngressCondition) SetLastTransitionTime(v string)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *V1RouteIngressCondition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMessage

`func (o *V1RouteIngressCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1RouteIngressCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1RouteIngressCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1RouteIngressCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *V1RouteIngressCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V1RouteIngressCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V1RouteIngressCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V1RouteIngressCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *V1RouteIngressCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1RouteIngressCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1RouteIngressCondition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *V1RouteIngressCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1RouteIngressCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1RouteIngressCondition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


