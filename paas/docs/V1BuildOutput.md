# V1BuildOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageLabels** | Pointer to [**[]V1ImageLabel**](V1ImageLabel.md) | imageLabels define a list of labels that are applied to the resulting image. If there are multiple labels with the same name then the last one in the list is used. | [optional] 
**PushSecret** | Pointer to [**V1LocalObjectReference**](V1LocalObjectReference.md) |  | [optional] 
**To** | Pointer to [**V1ObjectReference**](V1ObjectReference.md) |  | [optional] 

## Methods

### NewV1BuildOutput

`func NewV1BuildOutput() *V1BuildOutput`

NewV1BuildOutput instantiates a new V1BuildOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BuildOutputWithDefaults

`func NewV1BuildOutputWithDefaults() *V1BuildOutput`

NewV1BuildOutputWithDefaults instantiates a new V1BuildOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageLabels

`func (o *V1BuildOutput) GetImageLabels() []V1ImageLabel`

GetImageLabels returns the ImageLabels field if non-nil, zero value otherwise.

### GetImageLabelsOk

`func (o *V1BuildOutput) GetImageLabelsOk() (*[]V1ImageLabel, bool)`

GetImageLabelsOk returns a tuple with the ImageLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageLabels

`func (o *V1BuildOutput) SetImageLabels(v []V1ImageLabel)`

SetImageLabels sets ImageLabels field to given value.

### HasImageLabels

`func (o *V1BuildOutput) HasImageLabels() bool`

HasImageLabels returns a boolean if a field has been set.

### GetPushSecret

`func (o *V1BuildOutput) GetPushSecret() V1LocalObjectReference`

GetPushSecret returns the PushSecret field if non-nil, zero value otherwise.

### GetPushSecretOk

`func (o *V1BuildOutput) GetPushSecretOk() (*V1LocalObjectReference, bool)`

GetPushSecretOk returns a tuple with the PushSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushSecret

`func (o *V1BuildOutput) SetPushSecret(v V1LocalObjectReference)`

SetPushSecret sets PushSecret field to given value.

### HasPushSecret

`func (o *V1BuildOutput) HasPushSecret() bool`

HasPushSecret returns a boolean if a field has been set.

### GetTo

`func (o *V1BuildOutput) GetTo() V1ObjectReference`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *V1BuildOutput) GetToOk() (*V1ObjectReference, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *V1BuildOutput) SetTo(v V1ObjectReference)`

SetTo sets To field to given value.

### HasTo

`func (o *V1BuildOutput) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


