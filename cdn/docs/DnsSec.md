# DnsSec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] [default to false]
**Ds** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDnsSec

`func NewDnsSec() *DnsSec`

NewDnsSec instantiates a new DnsSec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsSecWithDefaults

`func NewDnsSecWithDefaults() *DnsSec`

NewDnsSecWithDefaults instantiates a new DnsSec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DnsSec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DnsSec) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DnsSec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DnsSec) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDs

`func (o *DnsSec) GetDs() string`

GetDs returns the Ds field if non-nil, zero value otherwise.

### GetDsOk

`func (o *DnsSec) GetDsOk() (*string, bool)`

GetDsOk returns a tuple with the Ds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDs

`func (o *DnsSec) SetDs(v string)`

SetDs sets Ds field to given value.

### HasDs

`func (o *DnsSec) HasDs() bool`

HasDs returns a boolean if a field has been set.

### SetDsNil

`func (o *DnsSec) SetDsNil(b bool)`

 SetDsNil sets the value for Ds to be an explicit nil

### UnsetDs
`func (o *DnsSec) UnsetDs()`

UnsetDs ensures that no value is present for Ds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


