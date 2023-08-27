# V1SELinuxContextStrategyOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeLinuxOptions** | Pointer to [**V1SELinuxOptions**](V1SELinuxOptions.md) |  | [optional] 
**Type** | Pointer to **string** | Type is the strategy that will dictate what SELinux context is used in the SecurityContext. | [optional] 

## Methods

### NewV1SELinuxContextStrategyOptions

`func NewV1SELinuxContextStrategyOptions() *V1SELinuxContextStrategyOptions`

NewV1SELinuxContextStrategyOptions instantiates a new V1SELinuxContextStrategyOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SELinuxContextStrategyOptionsWithDefaults

`func NewV1SELinuxContextStrategyOptionsWithDefaults() *V1SELinuxContextStrategyOptions`

NewV1SELinuxContextStrategyOptionsWithDefaults instantiates a new V1SELinuxContextStrategyOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeLinuxOptions

`func (o *V1SELinuxContextStrategyOptions) GetSeLinuxOptions() V1SELinuxOptions`

GetSeLinuxOptions returns the SeLinuxOptions field if non-nil, zero value otherwise.

### GetSeLinuxOptionsOk

`func (o *V1SELinuxContextStrategyOptions) GetSeLinuxOptionsOk() (*V1SELinuxOptions, bool)`

GetSeLinuxOptionsOk returns a tuple with the SeLinuxOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeLinuxOptions

`func (o *V1SELinuxContextStrategyOptions) SetSeLinuxOptions(v V1SELinuxOptions)`

SetSeLinuxOptions sets SeLinuxOptions field to given value.

### HasSeLinuxOptions

`func (o *V1SELinuxContextStrategyOptions) HasSeLinuxOptions() bool`

HasSeLinuxOptions returns a boolean if a field has been set.

### GetType

`func (o *V1SELinuxContextStrategyOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1SELinuxContextStrategyOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1SELinuxContextStrategyOptions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1SELinuxContextStrategyOptions) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


