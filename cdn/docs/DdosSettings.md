# DdosSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**ProtectionMode** | Pointer to **string** |  | [optional] 
**CaptchaService** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int32** | Time in seconds for cookie max-age | [optional] 
**HttpsOnly** | Pointer to **bool** | Adds \&quot;SameSite&#x3D;None; Secure\&quot; to set-cookie header | [optional] 
**Preflight** | Pointer to [**DdosPreflight**](DdosPreflight.md) |  | [optional] 

## Methods

### NewDdosSettings

`func NewDdosSettings() *DdosSettings`

NewDdosSettings instantiates a new DdosSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdosSettingsWithDefaults

`func NewDdosSettingsWithDefaults() *DdosSettings`

NewDdosSettingsWithDefaults instantiates a new DdosSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *DdosSettings) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DdosSettings) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DdosSettings) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *DdosSettings) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetProtectionMode

`func (o *DdosSettings) GetProtectionMode() string`

GetProtectionMode returns the ProtectionMode field if non-nil, zero value otherwise.

### GetProtectionModeOk

`func (o *DdosSettings) GetProtectionModeOk() (*string, bool)`

GetProtectionModeOk returns a tuple with the ProtectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionMode

`func (o *DdosSettings) SetProtectionMode(v string)`

SetProtectionMode sets ProtectionMode field to given value.

### HasProtectionMode

`func (o *DdosSettings) HasProtectionMode() bool`

HasProtectionMode returns a boolean if a field has been set.

### GetCaptchaService

`func (o *DdosSettings) GetCaptchaService() string`

GetCaptchaService returns the CaptchaService field if non-nil, zero value otherwise.

### GetCaptchaServiceOk

`func (o *DdosSettings) GetCaptchaServiceOk() (*string, bool)`

GetCaptchaServiceOk returns a tuple with the CaptchaService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptchaService

`func (o *DdosSettings) SetCaptchaService(v string)`

SetCaptchaService sets CaptchaService field to given value.

### HasCaptchaService

`func (o *DdosSettings) HasCaptchaService() bool`

HasCaptchaService returns a boolean if a field has been set.

### GetTtl

`func (o *DdosSettings) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DdosSettings) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DdosSettings) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DdosSettings) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetHttpsOnly

`func (o *DdosSettings) GetHttpsOnly() bool`

GetHttpsOnly returns the HttpsOnly field if non-nil, zero value otherwise.

### GetHttpsOnlyOk

`func (o *DdosSettings) GetHttpsOnlyOk() (*bool, bool)`

GetHttpsOnlyOk returns a tuple with the HttpsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsOnly

`func (o *DdosSettings) SetHttpsOnly(v bool)`

SetHttpsOnly sets HttpsOnly field to given value.

### HasHttpsOnly

`func (o *DdosSettings) HasHttpsOnly() bool`

HasHttpsOnly returns a boolean if a field has been set.

### GetPreflight

`func (o *DdosSettings) GetPreflight() DdosPreflight`

GetPreflight returns the Preflight field if non-nil, zero value otherwise.

### GetPreflightOk

`func (o *DdosSettings) GetPreflightOk() (*DdosPreflight, bool)`

GetPreflightOk returns a tuple with the Preflight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreflight

`func (o *DdosSettings) SetPreflight(v DdosPreflight)`

SetPreflight sets Preflight field to given value.

### HasPreflight

`func (o *DdosSettings) HasPreflight() bool`

HasPreflight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


