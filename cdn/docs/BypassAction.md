# BypassAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rlimit** | Pointer to **bool** |  | [optional] 
**Challenge** | Pointer to **bool** |  | [optional] 
**Waf** | Pointer to **bool** |  | [optional] 

## Methods

### NewBypassAction

`func NewBypassAction() *BypassAction`

NewBypassAction instantiates a new BypassAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBypassActionWithDefaults

`func NewBypassActionWithDefaults() *BypassAction`

NewBypassActionWithDefaults instantiates a new BypassAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRlimit

`func (o *BypassAction) GetRlimit() bool`

GetRlimit returns the Rlimit field if non-nil, zero value otherwise.

### GetRlimitOk

`func (o *BypassAction) GetRlimitOk() (*bool, bool)`

GetRlimitOk returns a tuple with the Rlimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlimit

`func (o *BypassAction) SetRlimit(v bool)`

SetRlimit sets Rlimit field to given value.

### HasRlimit

`func (o *BypassAction) HasRlimit() bool`

HasRlimit returns a boolean if a field has been set.

### GetChallenge

`func (o *BypassAction) GetChallenge() bool`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *BypassAction) GetChallengeOk() (*bool, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *BypassAction) SetChallenge(v bool)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *BypassAction) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetWaf

`func (o *BypassAction) GetWaf() bool`

GetWaf returns the Waf field if non-nil, zero value otherwise.

### GetWafOk

`func (o *BypassAction) GetWafOk() (*bool, bool)`

GetWafOk returns a tuple with the Waf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaf

`func (o *BypassAction) SetWaf(v bool)`

SetWaf sets Waf field to given value.

### HasWaf

`func (o *BypassAction) HasWaf() bool`

HasWaf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


