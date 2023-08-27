# V1TemplateInstanceCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTransitionTime** | **string** | LastTransitionTime is the last time a condition status transitioned from one state to another. | 
**Message** | **string** | Message is a human readable description of the details of the last transition, complementing reason. | 
**Reason** | **string** | Reason is a brief machine readable explanation for the condition&#39;s last transition. | 
**Status** | **string** | Status of the condition, one of True, False or Unknown. | 
**Type** | **string** | Type of the condition, currently Ready or InstantiateFailure. | 

## Methods

### NewV1TemplateInstanceCondition

`func NewV1TemplateInstanceCondition(lastTransitionTime string, message string, reason string, status string, type_ string, ) *V1TemplateInstanceCondition`

NewV1TemplateInstanceCondition instantiates a new V1TemplateInstanceCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TemplateInstanceConditionWithDefaults

`func NewV1TemplateInstanceConditionWithDefaults() *V1TemplateInstanceCondition`

NewV1TemplateInstanceConditionWithDefaults instantiates a new V1TemplateInstanceCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTransitionTime

`func (o *V1TemplateInstanceCondition) GetLastTransitionTime() string`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *V1TemplateInstanceCondition) GetLastTransitionTimeOk() (*string, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *V1TemplateInstanceCondition) SetLastTransitionTime(v string)`

SetLastTransitionTime sets LastTransitionTime field to given value.


### GetMessage

`func (o *V1TemplateInstanceCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1TemplateInstanceCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1TemplateInstanceCondition) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetReason

`func (o *V1TemplateInstanceCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V1TemplateInstanceCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V1TemplateInstanceCondition) SetReason(v string)`

SetReason sets Reason field to given value.


### GetStatus

`func (o *V1TemplateInstanceCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1TemplateInstanceCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1TemplateInstanceCondition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *V1TemplateInstanceCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1TemplateInstanceCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1TemplateInstanceCondition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


