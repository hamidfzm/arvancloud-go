# CustomPages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderConstruction** | Pointer to [**CustomPage**](CustomPage.md) |  | [optional] 
**FirewallError** | Pointer to [**CustomPage**](CustomPage.md) |  | [optional] 
**WafProtection** | Pointer to [**CustomPage**](CustomPage.md) |  | [optional] 
**RateLimitExceeded** | Pointer to [**CustomPage**](CustomPage.md) |  | [optional] 
**SecureLinkExpired** | Pointer to [**CustomPage**](CustomPage.md) |  | [optional] 
**SecureLinkInvalid** | Pointer to [**CustomPage**](CustomPage.md) |  | [optional] 
**Error500** | Pointer to [**CustomPage**](CustomPage.md) |  | [optional] 
**DdosJs** | Pointer to [**CustomPage**](CustomPage.md) |  | [optional] 
**DdosCaptcha** | Pointer to [**CustomPage**](CustomPage.md) |  | [optional] 

## Methods

### NewCustomPages

`func NewCustomPages() *CustomPages`

NewCustomPages instantiates a new CustomPages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomPagesWithDefaults

`func NewCustomPagesWithDefaults() *CustomPages`

NewCustomPagesWithDefaults instantiates a new CustomPages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderConstruction

`func (o *CustomPages) GetUnderConstruction() CustomPage`

GetUnderConstruction returns the UnderConstruction field if non-nil, zero value otherwise.

### GetUnderConstructionOk

`func (o *CustomPages) GetUnderConstructionOk() (*CustomPage, bool)`

GetUnderConstructionOk returns a tuple with the UnderConstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderConstruction

`func (o *CustomPages) SetUnderConstruction(v CustomPage)`

SetUnderConstruction sets UnderConstruction field to given value.

### HasUnderConstruction

`func (o *CustomPages) HasUnderConstruction() bool`

HasUnderConstruction returns a boolean if a field has been set.

### GetFirewallError

`func (o *CustomPages) GetFirewallError() CustomPage`

GetFirewallError returns the FirewallError field if non-nil, zero value otherwise.

### GetFirewallErrorOk

`func (o *CustomPages) GetFirewallErrorOk() (*CustomPage, bool)`

GetFirewallErrorOk returns a tuple with the FirewallError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallError

`func (o *CustomPages) SetFirewallError(v CustomPage)`

SetFirewallError sets FirewallError field to given value.

### HasFirewallError

`func (o *CustomPages) HasFirewallError() bool`

HasFirewallError returns a boolean if a field has been set.

### GetWafProtection

`func (o *CustomPages) GetWafProtection() CustomPage`

GetWafProtection returns the WafProtection field if non-nil, zero value otherwise.

### GetWafProtectionOk

`func (o *CustomPages) GetWafProtectionOk() (*CustomPage, bool)`

GetWafProtectionOk returns a tuple with the WafProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafProtection

`func (o *CustomPages) SetWafProtection(v CustomPage)`

SetWafProtection sets WafProtection field to given value.

### HasWafProtection

`func (o *CustomPages) HasWafProtection() bool`

HasWafProtection returns a boolean if a field has been set.

### GetRateLimitExceeded

`func (o *CustomPages) GetRateLimitExceeded() CustomPage`

GetRateLimitExceeded returns the RateLimitExceeded field if non-nil, zero value otherwise.

### GetRateLimitExceededOk

`func (o *CustomPages) GetRateLimitExceededOk() (*CustomPage, bool)`

GetRateLimitExceededOk returns a tuple with the RateLimitExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitExceeded

`func (o *CustomPages) SetRateLimitExceeded(v CustomPage)`

SetRateLimitExceeded sets RateLimitExceeded field to given value.

### HasRateLimitExceeded

`func (o *CustomPages) HasRateLimitExceeded() bool`

HasRateLimitExceeded returns a boolean if a field has been set.

### GetSecureLinkExpired

`func (o *CustomPages) GetSecureLinkExpired() CustomPage`

GetSecureLinkExpired returns the SecureLinkExpired field if non-nil, zero value otherwise.

### GetSecureLinkExpiredOk

`func (o *CustomPages) GetSecureLinkExpiredOk() (*CustomPage, bool)`

GetSecureLinkExpiredOk returns a tuple with the SecureLinkExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureLinkExpired

`func (o *CustomPages) SetSecureLinkExpired(v CustomPage)`

SetSecureLinkExpired sets SecureLinkExpired field to given value.

### HasSecureLinkExpired

`func (o *CustomPages) HasSecureLinkExpired() bool`

HasSecureLinkExpired returns a boolean if a field has been set.

### GetSecureLinkInvalid

`func (o *CustomPages) GetSecureLinkInvalid() CustomPage`

GetSecureLinkInvalid returns the SecureLinkInvalid field if non-nil, zero value otherwise.

### GetSecureLinkInvalidOk

`func (o *CustomPages) GetSecureLinkInvalidOk() (*CustomPage, bool)`

GetSecureLinkInvalidOk returns a tuple with the SecureLinkInvalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureLinkInvalid

`func (o *CustomPages) SetSecureLinkInvalid(v CustomPage)`

SetSecureLinkInvalid sets SecureLinkInvalid field to given value.

### HasSecureLinkInvalid

`func (o *CustomPages) HasSecureLinkInvalid() bool`

HasSecureLinkInvalid returns a boolean if a field has been set.

### GetError500

`func (o *CustomPages) GetError500() CustomPage`

GetError500 returns the Error500 field if non-nil, zero value otherwise.

### GetError500Ok

`func (o *CustomPages) GetError500Ok() (*CustomPage, bool)`

GetError500Ok returns a tuple with the Error500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError500

`func (o *CustomPages) SetError500(v CustomPage)`

SetError500 sets Error500 field to given value.

### HasError500

`func (o *CustomPages) HasError500() bool`

HasError500 returns a boolean if a field has been set.

### GetDdosJs

`func (o *CustomPages) GetDdosJs() CustomPage`

GetDdosJs returns the DdosJs field if non-nil, zero value otherwise.

### GetDdosJsOk

`func (o *CustomPages) GetDdosJsOk() (*CustomPage, bool)`

GetDdosJsOk returns a tuple with the DdosJs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosJs

`func (o *CustomPages) SetDdosJs(v CustomPage)`

SetDdosJs sets DdosJs field to given value.

### HasDdosJs

`func (o *CustomPages) HasDdosJs() bool`

HasDdosJs returns a boolean if a field has been set.

### GetDdosCaptcha

`func (o *CustomPages) GetDdosCaptcha() CustomPage`

GetDdosCaptcha returns the DdosCaptcha field if non-nil, zero value otherwise.

### GetDdosCaptchaOk

`func (o *CustomPages) GetDdosCaptchaOk() (*CustomPage, bool)`

GetDdosCaptchaOk returns a tuple with the DdosCaptcha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosCaptcha

`func (o *CustomPages) SetDdosCaptcha(v CustomPage)`

SetDdosCaptcha sets DdosCaptcha field to given value.

### HasDdosCaptcha

`func (o *CustomPages) HasDdosCaptcha() bool`

HasDdosCaptcha returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


