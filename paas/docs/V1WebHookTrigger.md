# V1WebHookTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowEnv** | Pointer to **bool** | allowEnv determines whether the webhook can set environment variables; can only be set to true for GenericWebHook. | [optional] 
**Secret** | Pointer to **string** | secret used to validate requests. Deprecated: use SecretReference instead. | [optional] 
**SecretReference** | Pointer to [**V1SecretLocalReference**](V1SecretLocalReference.md) |  | [optional] 

## Methods

### NewV1WebHookTrigger

`func NewV1WebHookTrigger() *V1WebHookTrigger`

NewV1WebHookTrigger instantiates a new V1WebHookTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1WebHookTriggerWithDefaults

`func NewV1WebHookTriggerWithDefaults() *V1WebHookTrigger`

NewV1WebHookTriggerWithDefaults instantiates a new V1WebHookTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowEnv

`func (o *V1WebHookTrigger) GetAllowEnv() bool`

GetAllowEnv returns the AllowEnv field if non-nil, zero value otherwise.

### GetAllowEnvOk

`func (o *V1WebHookTrigger) GetAllowEnvOk() (*bool, bool)`

GetAllowEnvOk returns a tuple with the AllowEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEnv

`func (o *V1WebHookTrigger) SetAllowEnv(v bool)`

SetAllowEnv sets AllowEnv field to given value.

### HasAllowEnv

`func (o *V1WebHookTrigger) HasAllowEnv() bool`

HasAllowEnv returns a boolean if a field has been set.

### GetSecret

`func (o *V1WebHookTrigger) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *V1WebHookTrigger) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *V1WebHookTrigger) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *V1WebHookTrigger) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetSecretReference

`func (o *V1WebHookTrigger) GetSecretReference() V1SecretLocalReference`

GetSecretReference returns the SecretReference field if non-nil, zero value otherwise.

### GetSecretReferenceOk

`func (o *V1WebHookTrigger) GetSecretReferenceOk() (*V1SecretLocalReference, bool)`

GetSecretReferenceOk returns a tuple with the SecretReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretReference

`func (o *V1WebHookTrigger) SetSecretReference(v V1SecretLocalReference)`

SetSecretReference sets SecretReference field to given value.

### HasSecretReference

`func (o *V1WebHookTrigger) HasSecretReference() bool`

HasSecretReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


