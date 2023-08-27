# V1Taint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | **string** | Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute. | 
**Key** | **string** | Required. The taint key to be applied to a node. | 
**TimeAdded** | Pointer to **string** | TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints. | [optional] 
**Value** | Pointer to **string** | Required. The taint value corresponding to the taint key. | [optional] 

## Methods

### NewV1Taint

`func NewV1Taint(effect string, key string, ) *V1Taint`

NewV1Taint instantiates a new V1Taint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TaintWithDefaults

`func NewV1TaintWithDefaults() *V1Taint`

NewV1TaintWithDefaults instantiates a new V1Taint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *V1Taint) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *V1Taint) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *V1Taint) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetKey

`func (o *V1Taint) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V1Taint) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V1Taint) SetKey(v string)`

SetKey sets Key field to given value.


### GetTimeAdded

`func (o *V1Taint) GetTimeAdded() string`

GetTimeAdded returns the TimeAdded field if non-nil, zero value otherwise.

### GetTimeAddedOk

`func (o *V1Taint) GetTimeAddedOk() (*string, bool)`

GetTimeAddedOk returns a tuple with the TimeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAdded

`func (o *V1Taint) SetTimeAdded(v string)`

SetTimeAdded sets TimeAdded field to given value.

### HasTimeAdded

`func (o *V1Taint) HasTimeAdded() bool`

HasTimeAdded returns a boolean if a field has been set.

### GetValue

`func (o *V1Taint) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V1Taint) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V1Taint) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *V1Taint) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


