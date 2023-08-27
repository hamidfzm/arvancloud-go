# LifecycleRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiration** | Pointer to [**LifecycleRuleExpiration**](LifecycleRuleExpiration.md) |  | [optional] 
**ID** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**Filter** | Pointer to [**LifecycleRuleFilter**](LifecycleRuleFilter.md) |  | [optional] 
**Status** | [**ExpirationStatus**](ExpirationStatus.md) |  | 
**Transitions** | Pointer to [**Array**](array.md) |  | [optional] 
**NoncurrentVersionTransitions** | Pointer to [**Array**](array.md) |  | [optional] 
**NoncurrentVersionExpiration** | Pointer to [**NoncurrentVersionExpiration**](NoncurrentVersionExpiration.md) |  | [optional] 
**AbortIncompleteMultipartUpload** | Pointer to [**AbortIncompleteMultipartUpload**](AbortIncompleteMultipartUpload.md) |  | [optional] 

## Methods

### NewLifecycleRule

`func NewLifecycleRule(status ExpirationStatus, ) *LifecycleRule`

NewLifecycleRule instantiates a new LifecycleRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifecycleRuleWithDefaults

`func NewLifecycleRuleWithDefaults() *LifecycleRule`

NewLifecycleRuleWithDefaults instantiates a new LifecycleRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiration

`func (o *LifecycleRule) GetExpiration() LifecycleRuleExpiration`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *LifecycleRule) GetExpirationOk() (*LifecycleRuleExpiration, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *LifecycleRule) SetExpiration(v LifecycleRuleExpiration)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *LifecycleRule) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetID

`func (o *LifecycleRule) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *LifecycleRule) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *LifecycleRule) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *LifecycleRule) HasID() bool`

HasID returns a boolean if a field has been set.

### GetPrefix

`func (o *LifecycleRule) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *LifecycleRule) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *LifecycleRule) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *LifecycleRule) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetFilter

`func (o *LifecycleRule) GetFilter() LifecycleRuleFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LifecycleRule) GetFilterOk() (*LifecycleRuleFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LifecycleRule) SetFilter(v LifecycleRuleFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *LifecycleRule) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetStatus

`func (o *LifecycleRule) GetStatus() ExpirationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LifecycleRule) GetStatusOk() (*ExpirationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LifecycleRule) SetStatus(v ExpirationStatus)`

SetStatus sets Status field to given value.


### GetTransitions

`func (o *LifecycleRule) GetTransitions() Array`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *LifecycleRule) GetTransitionsOk() (*Array, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *LifecycleRule) SetTransitions(v Array)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *LifecycleRule) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.

### GetNoncurrentVersionTransitions

`func (o *LifecycleRule) GetNoncurrentVersionTransitions() Array`

GetNoncurrentVersionTransitions returns the NoncurrentVersionTransitions field if non-nil, zero value otherwise.

### GetNoncurrentVersionTransitionsOk

`func (o *LifecycleRule) GetNoncurrentVersionTransitionsOk() (*Array, bool)`

GetNoncurrentVersionTransitionsOk returns a tuple with the NoncurrentVersionTransitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoncurrentVersionTransitions

`func (o *LifecycleRule) SetNoncurrentVersionTransitions(v Array)`

SetNoncurrentVersionTransitions sets NoncurrentVersionTransitions field to given value.

### HasNoncurrentVersionTransitions

`func (o *LifecycleRule) HasNoncurrentVersionTransitions() bool`

HasNoncurrentVersionTransitions returns a boolean if a field has been set.

### GetNoncurrentVersionExpiration

`func (o *LifecycleRule) GetNoncurrentVersionExpiration() NoncurrentVersionExpiration`

GetNoncurrentVersionExpiration returns the NoncurrentVersionExpiration field if non-nil, zero value otherwise.

### GetNoncurrentVersionExpirationOk

`func (o *LifecycleRule) GetNoncurrentVersionExpirationOk() (*NoncurrentVersionExpiration, bool)`

GetNoncurrentVersionExpirationOk returns a tuple with the NoncurrentVersionExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoncurrentVersionExpiration

`func (o *LifecycleRule) SetNoncurrentVersionExpiration(v NoncurrentVersionExpiration)`

SetNoncurrentVersionExpiration sets NoncurrentVersionExpiration field to given value.

### HasNoncurrentVersionExpiration

`func (o *LifecycleRule) HasNoncurrentVersionExpiration() bool`

HasNoncurrentVersionExpiration returns a boolean if a field has been set.

### GetAbortIncompleteMultipartUpload

`func (o *LifecycleRule) GetAbortIncompleteMultipartUpload() AbortIncompleteMultipartUpload`

GetAbortIncompleteMultipartUpload returns the AbortIncompleteMultipartUpload field if non-nil, zero value otherwise.

### GetAbortIncompleteMultipartUploadOk

`func (o *LifecycleRule) GetAbortIncompleteMultipartUploadOk() (*AbortIncompleteMultipartUpload, bool)`

GetAbortIncompleteMultipartUploadOk returns a tuple with the AbortIncompleteMultipartUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortIncompleteMultipartUpload

`func (o *LifecycleRule) SetAbortIncompleteMultipartUpload(v AbortIncompleteMultipartUpload)`

SetAbortIncompleteMultipartUpload sets AbortIncompleteMultipartUpload field to given value.

### HasAbortIncompleteMultipartUpload

`func (o *LifecycleRule) HasAbortIncompleteMultipartUpload() bool`

HasAbortIncompleteMultipartUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


