# WafPreset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Packages** | Pointer to [**[]WafPresetPackagesInner**](WafPresetPackagesInner.md) |  | [optional] 

## Methods

### NewWafPreset

`func NewWafPreset() *WafPreset`

NewWafPreset instantiates a new WafPreset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafPresetWithDefaults

`func NewWafPresetWithDefaults() *WafPreset`

NewWafPresetWithDefaults instantiates a new WafPreset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WafPreset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WafPreset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WafPreset) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WafPreset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WafPreset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WafPreset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WafPreset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WafPreset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *WafPreset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WafPreset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WafPreset) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WafPreset) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPackages

`func (o *WafPreset) GetPackages() []WafPresetPackagesInner`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *WafPreset) GetPackagesOk() (*[]WafPresetPackagesInner, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *WafPreset) SetPackages(v []WafPresetPackagesInner)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *WafPreset) HasPackages() bool`

HasPackages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


