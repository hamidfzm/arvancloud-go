# V1TemplateInstanceRequester

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extra** | Pointer to **map[string]interface{}** | extra holds additional information provided by the authenticator. | [optional] 
**Groups** | Pointer to **[]string** | groups represent the groups this user is a part of. | [optional] 
**Uid** | Pointer to **string** | uid is a unique value that identifies this user across time; if this user is deleted and another user by the same name is added, they will have different UIDs. | [optional] 
**Username** | Pointer to **string** | username uniquely identifies this user among all active users. | [optional] 

## Methods

### NewV1TemplateInstanceRequester

`func NewV1TemplateInstanceRequester() *V1TemplateInstanceRequester`

NewV1TemplateInstanceRequester instantiates a new V1TemplateInstanceRequester object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TemplateInstanceRequesterWithDefaults

`func NewV1TemplateInstanceRequesterWithDefaults() *V1TemplateInstanceRequester`

NewV1TemplateInstanceRequesterWithDefaults instantiates a new V1TemplateInstanceRequester object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtra

`func (o *V1TemplateInstanceRequester) GetExtra() map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *V1TemplateInstanceRequester) GetExtraOk() (*map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *V1TemplateInstanceRequester) SetExtra(v map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *V1TemplateInstanceRequester) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetGroups

`func (o *V1TemplateInstanceRequester) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *V1TemplateInstanceRequester) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *V1TemplateInstanceRequester) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *V1TemplateInstanceRequester) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUid

`func (o *V1TemplateInstanceRequester) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *V1TemplateInstanceRequester) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *V1TemplateInstanceRequester) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *V1TemplateInstanceRequester) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUsername

`func (o *V1TemplateInstanceRequester) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *V1TemplateInstanceRequester) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *V1TemplateInstanceRequester) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *V1TemplateInstanceRequester) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


