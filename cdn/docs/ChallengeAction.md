# ChallengeAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **int32** | The mode of mitigation (1: Cookie, 2: Javascript, 3: Captcha) | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 
**HttpsOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewChallengeAction

`func NewChallengeAction() *ChallengeAction`

NewChallengeAction instantiates a new ChallengeAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengeActionWithDefaults

`func NewChallengeActionWithDefaults() *ChallengeAction`

NewChallengeActionWithDefaults instantiates a new ChallengeAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *ChallengeAction) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ChallengeAction) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ChallengeAction) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ChallengeAction) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetTtl

`func (o *ChallengeAction) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ChallengeAction) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ChallengeAction) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ChallengeAction) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetHttpsOnly

`func (o *ChallengeAction) GetHttpsOnly() bool`

GetHttpsOnly returns the HttpsOnly field if non-nil, zero value otherwise.

### GetHttpsOnlyOk

`func (o *ChallengeAction) GetHttpsOnlyOk() (*bool, bool)`

GetHttpsOnlyOk returns a tuple with the HttpsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsOnly

`func (o *ChallengeAction) SetHttpsOnly(v bool)`

SetHttpsOnly sets HttpsOnly field to given value.

### HasHttpsOnly

`func (o *ChallengeAction) HasHttpsOnly() bool`

HasHttpsOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


