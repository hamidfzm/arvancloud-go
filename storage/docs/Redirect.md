# Redirect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostName** | Pointer to **string** |  | [optional] 
**HttpRedirectCode** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to [**Protocol**](Protocol.md) |  | [optional] 
**ReplaceKeyPrefixWith** | Pointer to **string** |  | [optional] 
**ReplaceKeyWith** | Pointer to **string** |  | [optional] 

## Methods

### NewRedirect

`func NewRedirect() *Redirect`

NewRedirect instantiates a new Redirect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedirectWithDefaults

`func NewRedirectWithDefaults() *Redirect`

NewRedirectWithDefaults instantiates a new Redirect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostName

`func (o *Redirect) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *Redirect) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *Redirect) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *Redirect) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetHttpRedirectCode

`func (o *Redirect) GetHttpRedirectCode() string`

GetHttpRedirectCode returns the HttpRedirectCode field if non-nil, zero value otherwise.

### GetHttpRedirectCodeOk

`func (o *Redirect) GetHttpRedirectCodeOk() (*string, bool)`

GetHttpRedirectCodeOk returns a tuple with the HttpRedirectCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRedirectCode

`func (o *Redirect) SetHttpRedirectCode(v string)`

SetHttpRedirectCode sets HttpRedirectCode field to given value.

### HasHttpRedirectCode

`func (o *Redirect) HasHttpRedirectCode() bool`

HasHttpRedirectCode returns a boolean if a field has been set.

### GetProtocol

`func (o *Redirect) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Redirect) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Redirect) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *Redirect) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetReplaceKeyPrefixWith

`func (o *Redirect) GetReplaceKeyPrefixWith() string`

GetReplaceKeyPrefixWith returns the ReplaceKeyPrefixWith field if non-nil, zero value otherwise.

### GetReplaceKeyPrefixWithOk

`func (o *Redirect) GetReplaceKeyPrefixWithOk() (*string, bool)`

GetReplaceKeyPrefixWithOk returns a tuple with the ReplaceKeyPrefixWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceKeyPrefixWith

`func (o *Redirect) SetReplaceKeyPrefixWith(v string)`

SetReplaceKeyPrefixWith sets ReplaceKeyPrefixWith field to given value.

### HasReplaceKeyPrefixWith

`func (o *Redirect) HasReplaceKeyPrefixWith() bool`

HasReplaceKeyPrefixWith returns a boolean if a field has been set.

### GetReplaceKeyWith

`func (o *Redirect) GetReplaceKeyWith() string`

GetReplaceKeyWith returns the ReplaceKeyWith field if non-nil, zero value otherwise.

### GetReplaceKeyWithOk

`func (o *Redirect) GetReplaceKeyWithOk() (*string, bool)`

GetReplaceKeyWithOk returns a tuple with the ReplaceKeyWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceKeyWith

`func (o *Redirect) SetReplaceKeyWith(v string)`

SetReplaceKeyWith sets ReplaceKeyWith field to given value.

### HasReplaceKeyWith

`func (o *Redirect) HasReplaceKeyWith() bool`

HasReplaceKeyWith returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


