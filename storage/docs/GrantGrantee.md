# GrantGrantee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**EmailAddress** | Pointer to **string** |  | [optional] 
**ID** | Pointer to **string** |  | [optional] 
**Type** | [**Type**](Type.md) |  | 
**URI** | Pointer to **string** |  | [optional] 

## Methods

### NewGrantGrantee

`func NewGrantGrantee(type_ Type, ) *GrantGrantee`

NewGrantGrantee instantiates a new GrantGrantee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantGranteeWithDefaults

`func NewGrantGranteeWithDefaults() *GrantGrantee`

NewGrantGranteeWithDefaults instantiates a new GrantGrantee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *GrantGrantee) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GrantGrantee) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GrantGrantee) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GrantGrantee) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *GrantGrantee) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *GrantGrantee) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *GrantGrantee) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *GrantGrantee) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetID

`func (o *GrantGrantee) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *GrantGrantee) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *GrantGrantee) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *GrantGrantee) HasID() bool`

HasID returns a boolean if a field has been set.

### GetType

`func (o *GrantGrantee) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GrantGrantee) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GrantGrantee) SetType(v Type)`

SetType sets Type field to given value.


### GetURI

`func (o *GrantGrantee) GetURI() string`

GetURI returns the URI field if non-nil, zero value otherwise.

### GetURIOk

`func (o *GrantGrantee) GetURIOk() (*string, bool)`

GetURIOk returns a tuple with the URI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURI

`func (o *GrantGrantee) SetURI(v string)`

SetURI sets URI field to given value.

### HasURI

`func (o *GrantGrantee) HasURI() bool`

HasURI returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


