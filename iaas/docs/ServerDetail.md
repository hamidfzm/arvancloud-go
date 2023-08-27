# ServerDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to [**map[string][]ServerAddress**](array.md) |  | [optional] 
**ArNext** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Flavor** | Pointer to [**ServerFlavor**](ServerFlavor.md) |  | [optional] 
**HaEnabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Image** | Pointer to [**ServerImage**](ServerImage.md) |  | [optional] 
**KeyName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**SecurityGroups** | Pointer to [**[]SecurityGroup**](SecurityGroup.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**TaskState** | Pointer to **string** |  | [optional] 

## Methods

### NewServerDetail

`func NewServerDetail() *ServerDetail`

NewServerDetail instantiates a new ServerDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerDetailWithDefaults

`func NewServerDetailWithDefaults() *ServerDetail`

NewServerDetailWithDefaults instantiates a new ServerDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ServerDetail) GetAddresses() map[string][]ServerAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ServerDetail) GetAddressesOk() (*map[string][]ServerAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ServerDetail) SetAddresses(v map[string][]ServerAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *ServerDetail) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetArNext

`func (o *ServerDetail) GetArNext() string`

GetArNext returns the ArNext field if non-nil, zero value otherwise.

### GetArNextOk

`func (o *ServerDetail) GetArNextOk() (*string, bool)`

GetArNextOk returns a tuple with the ArNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArNext

`func (o *ServerDetail) SetArNext(v string)`

SetArNext sets ArNext field to given value.

### HasArNext

`func (o *ServerDetail) HasArNext() bool`

HasArNext returns a boolean if a field has been set.

### GetCreated

`func (o *ServerDetail) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ServerDetail) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ServerDetail) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ServerDetail) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetFlavor

`func (o *ServerDetail) GetFlavor() ServerFlavor`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *ServerDetail) GetFlavorOk() (*ServerFlavor, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *ServerDetail) SetFlavor(v ServerFlavor)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *ServerDetail) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### GetHaEnabled

`func (o *ServerDetail) GetHaEnabled() bool`

GetHaEnabled returns the HaEnabled field if non-nil, zero value otherwise.

### GetHaEnabledOk

`func (o *ServerDetail) GetHaEnabledOk() (*bool, bool)`

GetHaEnabledOk returns a tuple with the HaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaEnabled

`func (o *ServerDetail) SetHaEnabled(v bool)`

SetHaEnabled sets HaEnabled field to given value.

### HasHaEnabled

`func (o *ServerDetail) HasHaEnabled() bool`

HasHaEnabled returns a boolean if a field has been set.

### GetId

`func (o *ServerDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServerDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *ServerDetail) GetImage() ServerImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ServerDetail) GetImageOk() (*ServerImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ServerDetail) SetImage(v ServerImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *ServerDetail) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetKeyName

`func (o *ServerDetail) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *ServerDetail) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *ServerDetail) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *ServerDetail) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### GetName

`func (o *ServerDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *ServerDetail) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ServerDetail) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ServerDetail) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ServerDetail) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *ServerDetail) GetSecurityGroups() []SecurityGroup`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *ServerDetail) GetSecurityGroupsOk() (*[]SecurityGroup, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *ServerDetail) SetSecurityGroups(v []SecurityGroup)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *ServerDetail) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetStatus

`func (o *ServerDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServerDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *ServerDetail) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServerDetail) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServerDetail) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServerDetail) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTaskState

`func (o *ServerDetail) GetTaskState() string`

GetTaskState returns the TaskState field if non-nil, zero value otherwise.

### GetTaskStateOk

`func (o *ServerDetail) GetTaskStateOk() (*string, bool)`

GetTaskStateOk returns a tuple with the TaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskState

`func (o *ServerDetail) SetTaskState(v string)`

SetTaskState sets TaskState field to given value.

### HasTaskState

`func (o *ServerDetail) HasTaskState() bool`

HasTaskState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


