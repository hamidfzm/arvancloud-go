# V1VolumeNodeAffinity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Required** | Pointer to [**V1NodeSelector**](V1NodeSelector.md) |  | [optional] 

## Methods

### NewV1VolumeNodeAffinity

`func NewV1VolumeNodeAffinity() *V1VolumeNodeAffinity`

NewV1VolumeNodeAffinity instantiates a new V1VolumeNodeAffinity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1VolumeNodeAffinityWithDefaults

`func NewV1VolumeNodeAffinityWithDefaults() *V1VolumeNodeAffinity`

NewV1VolumeNodeAffinityWithDefaults instantiates a new V1VolumeNodeAffinity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequired

`func (o *V1VolumeNodeAffinity) GetRequired() V1NodeSelector`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *V1VolumeNodeAffinity) GetRequiredOk() (*V1NodeSelector, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *V1VolumeNodeAffinity) SetRequired(v V1NodeSelector)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *V1VolumeNodeAffinity) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


