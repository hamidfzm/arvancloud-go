# FirewallActionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rlimit** | Pointer to **bool** |  | [optional] 
**Challenge** | Pointer to **bool** |  | [optional] 
**Waf** | Pointer to **bool** |  | [optional] 
**Mode** | Pointer to **int32** | The mode of mitigation (1: Cookie, 2: Javascript, 3: Captcha) | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 
**HttpsOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewFirewallActionDetails

`func NewFirewallActionDetails() *FirewallActionDetails`

NewFirewallActionDetails instantiates a new FirewallActionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallActionDetailsWithDefaults

`func NewFirewallActionDetailsWithDefaults() *FirewallActionDetails`

NewFirewallActionDetailsWithDefaults instantiates a new FirewallActionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRlimit

`func (o *FirewallActionDetails) GetRlimit() bool`

GetRlimit returns the Rlimit field if non-nil, zero value otherwise.

### GetRlimitOk

`func (o *FirewallActionDetails) GetRlimitOk() (*bool, bool)`

GetRlimitOk returns a tuple with the Rlimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlimit

`func (o *FirewallActionDetails) SetRlimit(v bool)`

SetRlimit sets Rlimit field to given value.

### HasRlimit

`func (o *FirewallActionDetails) HasRlimit() bool`

HasRlimit returns a boolean if a field has been set.

### GetChallenge

`func (o *FirewallActionDetails) GetChallenge() bool`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *FirewallActionDetails) GetChallengeOk() (*bool, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *FirewallActionDetails) SetChallenge(v bool)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *FirewallActionDetails) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetWaf

`func (o *FirewallActionDetails) GetWaf() bool`

GetWaf returns the Waf field if non-nil, zero value otherwise.

### GetWafOk

`func (o *FirewallActionDetails) GetWafOk() (*bool, bool)`

GetWafOk returns a tuple with the Waf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaf

`func (o *FirewallActionDetails) SetWaf(v bool)`

SetWaf sets Waf field to given value.

### HasWaf

`func (o *FirewallActionDetails) HasWaf() bool`

HasWaf returns a boolean if a field has been set.

### GetMode

`func (o *FirewallActionDetails) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FirewallActionDetails) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FirewallActionDetails) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *FirewallActionDetails) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetTtl

`func (o *FirewallActionDetails) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *FirewallActionDetails) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *FirewallActionDetails) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *FirewallActionDetails) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetHttpsOnly

`func (o *FirewallActionDetails) GetHttpsOnly() bool`

GetHttpsOnly returns the HttpsOnly field if non-nil, zero value otherwise.

### GetHttpsOnlyOk

`func (o *FirewallActionDetails) GetHttpsOnlyOk() (*bool, bool)`

GetHttpsOnlyOk returns a tuple with the HttpsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsOnly

`func (o *FirewallActionDetails) SetHttpsOnly(v bool)`

SetHttpsOnly sets HttpsOnly field to given value.

### HasHttpsOnly

`func (o *FirewallActionDetails) HasHttpsOnly() bool`

HasHttpsOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


