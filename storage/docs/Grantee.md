# Grantee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**EmailAddress** | Pointer to **string** |  | [optional] 
**ID** | Pointer to **string** |  | [optional] 
**Type** | [**Type**](Type.md) |  | 
**URI** | Pointer to **string** |  | [optional] 

## Methods

### NewGrantee

`func NewGrantee(type_ Type, ) *Grantee`

NewGrantee instantiates a new Grantee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGranteeWithDefaults

`func NewGranteeWithDefaults() *Grantee`

NewGranteeWithDefaults instantiates a new Grantee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *Grantee) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Grantee) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Grantee) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Grantee) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *Grantee) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *Grantee) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *Grantee) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *Grantee) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetID

`func (o *Grantee) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *Grantee) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *Grantee) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *Grantee) HasID() bool`

HasID returns a boolean if a field has been set.

### GetType

`func (o *Grantee) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Grantee) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Grantee) SetType(v Type)`

SetType sets Type field to given value.


### GetURI

`func (o *Grantee) GetURI() string`

GetURI returns the URI field if non-nil, zero value otherwise.

### GetURIOk

`func (o *Grantee) GetURIOk() (*string, bool)`

GetURIOk returns a tuple with the URI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURI

`func (o *Grantee) SetURI(v string)`

SetURI sets URI field to given value.

### HasURI

`func (o *Grantee) HasURI() bool`

HasURI returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


