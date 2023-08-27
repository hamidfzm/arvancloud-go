# V1ContainerStateWaiting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message regarding why the container is not yet running. | [optional] 
**Reason** | Pointer to **string** | (brief) reason the container is not yet running. | [optional] 

## Methods

### NewV1ContainerStateWaiting

`func NewV1ContainerStateWaiting() *V1ContainerStateWaiting`

NewV1ContainerStateWaiting instantiates a new V1ContainerStateWaiting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ContainerStateWaitingWithDefaults

`func NewV1ContainerStateWaitingWithDefaults() *V1ContainerStateWaiting`

NewV1ContainerStateWaitingWithDefaults instantiates a new V1ContainerStateWaiting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *V1ContainerStateWaiting) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1ContainerStateWaiting) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1ContainerStateWaiting) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1ContainerStateWaiting) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *V1ContainerStateWaiting) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V1ContainerStateWaiting) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V1ContainerStateWaiting) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V1ContainerStateWaiting) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


