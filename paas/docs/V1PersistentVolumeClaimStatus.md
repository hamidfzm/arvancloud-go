# V1PersistentVolumeClaimStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | Pointer to [**[]V1PersistentVolumeAccessMode**](V1PersistentVolumeAccessMode.md) | AccessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 | [optional] 
**Capacity** | Pointer to **map[string]interface{}** | Represents the actual resources of the underlying volume. | [optional] 
**Conditions** | Pointer to [**[]V1PersistentVolumeClaimCondition**](V1PersistentVolumeClaimCondition.md) | Current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to &#39;ResizeStarted&#39;. | [optional] 
**Phase** | Pointer to **string** | Phase represents the current phase of PersistentVolumeClaim. | [optional] 

## Methods

### NewV1PersistentVolumeClaimStatus

`func NewV1PersistentVolumeClaimStatus() *V1PersistentVolumeClaimStatus`

NewV1PersistentVolumeClaimStatus instantiates a new V1PersistentVolumeClaimStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PersistentVolumeClaimStatusWithDefaults

`func NewV1PersistentVolumeClaimStatusWithDefaults() *V1PersistentVolumeClaimStatus`

NewV1PersistentVolumeClaimStatusWithDefaults instantiates a new V1PersistentVolumeClaimStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessModes

`func (o *V1PersistentVolumeClaimStatus) GetAccessModes() []V1PersistentVolumeAccessMode`

GetAccessModes returns the AccessModes field if non-nil, zero value otherwise.

### GetAccessModesOk

`func (o *V1PersistentVolumeClaimStatus) GetAccessModesOk() (*[]V1PersistentVolumeAccessMode, bool)`

GetAccessModesOk returns a tuple with the AccessModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModes

`func (o *V1PersistentVolumeClaimStatus) SetAccessModes(v []V1PersistentVolumeAccessMode)`

SetAccessModes sets AccessModes field to given value.

### HasAccessModes

`func (o *V1PersistentVolumeClaimStatus) HasAccessModes() bool`

HasAccessModes returns a boolean if a field has been set.

### GetCapacity

`func (o *V1PersistentVolumeClaimStatus) GetCapacity() map[string]interface{}`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *V1PersistentVolumeClaimStatus) GetCapacityOk() (*map[string]interface{}, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *V1PersistentVolumeClaimStatus) SetCapacity(v map[string]interface{})`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *V1PersistentVolumeClaimStatus) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetConditions

`func (o *V1PersistentVolumeClaimStatus) GetConditions() []V1PersistentVolumeClaimCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *V1PersistentVolumeClaimStatus) GetConditionsOk() (*[]V1PersistentVolumeClaimCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *V1PersistentVolumeClaimStatus) SetConditions(v []V1PersistentVolumeClaimCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *V1PersistentVolumeClaimStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetPhase

`func (o *V1PersistentVolumeClaimStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *V1PersistentVolumeClaimStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *V1PersistentVolumeClaimStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *V1PersistentVolumeClaimStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


