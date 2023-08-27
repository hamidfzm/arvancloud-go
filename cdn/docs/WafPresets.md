# WafPresets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Presets** | Pointer to [**[]WafPreset**](WafPreset.md) |  | [optional] 
**Packages** | Pointer to [**[]WafPackage**](WafPackage.md) |  | [optional] 

## Methods

### NewWafPresets

`func NewWafPresets() *WafPresets`

NewWafPresets instantiates a new WafPresets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafPresetsWithDefaults

`func NewWafPresetsWithDefaults() *WafPresets`

NewWafPresetsWithDefaults instantiates a new WafPresets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPresets

`func (o *WafPresets) GetPresets() []WafPreset`

GetPresets returns the Presets field if non-nil, zero value otherwise.

### GetPresetsOk

`func (o *WafPresets) GetPresetsOk() (*[]WafPreset, bool)`

GetPresetsOk returns a tuple with the Presets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresets

`func (o *WafPresets) SetPresets(v []WafPreset)`

SetPresets sets Presets field to given value.

### HasPresets

`func (o *WafPresets) HasPresets() bool`

HasPresets returns a boolean if a field has been set.

### GetPackages

`func (o *WafPresets) GetPackages() []WafPackage`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *WafPresets) GetPackagesOk() (*[]WafPackage, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *WafPresets) SetPackages(v []WafPackage)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *WafPresets) HasPackages() bool`

HasPackages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


