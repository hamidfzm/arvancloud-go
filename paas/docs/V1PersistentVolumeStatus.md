# V1PersistentVolumeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | A human-readable message indicating details about why the volume is in this state. | [optional] 
**Phase** | Pointer to **string** | Phase indicates if a volume is available, bound to a claim, or released by a claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#phase | [optional] 
**Reason** | Pointer to **string** | Reason is a brief CamelCase string that describes any failure and is meant for machine parsing and tidy display in the CLI. | [optional] 

## Methods

### NewV1PersistentVolumeStatus

`func NewV1PersistentVolumeStatus() *V1PersistentVolumeStatus`

NewV1PersistentVolumeStatus instantiates a new V1PersistentVolumeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PersistentVolumeStatusWithDefaults

`func NewV1PersistentVolumeStatusWithDefaults() *V1PersistentVolumeStatus`

NewV1PersistentVolumeStatusWithDefaults instantiates a new V1PersistentVolumeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *V1PersistentVolumeStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1PersistentVolumeStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1PersistentVolumeStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1PersistentVolumeStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPhase

`func (o *V1PersistentVolumeStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *V1PersistentVolumeStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *V1PersistentVolumeStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *V1PersistentVolumeStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetReason

`func (o *V1PersistentVolumeStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V1PersistentVolumeStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V1PersistentVolumeStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V1PersistentVolumeStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


