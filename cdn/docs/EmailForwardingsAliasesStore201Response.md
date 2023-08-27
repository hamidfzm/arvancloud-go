# EmailForwardingsAliasesStore201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**EmailForwardingAlias**](EmailForwardingAlias.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEmailForwardingsAliasesStore201Response

`func NewEmailForwardingsAliasesStore201Response() *EmailForwardingsAliasesStore201Response`

NewEmailForwardingsAliasesStore201Response instantiates a new EmailForwardingsAliasesStore201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailForwardingsAliasesStore201ResponseWithDefaults

`func NewEmailForwardingsAliasesStore201ResponseWithDefaults() *EmailForwardingsAliasesStore201Response`

NewEmailForwardingsAliasesStore201ResponseWithDefaults instantiates a new EmailForwardingsAliasesStore201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EmailForwardingsAliasesStore201Response) GetData() EmailForwardingAlias`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EmailForwardingsAliasesStore201Response) GetDataOk() (*EmailForwardingAlias, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EmailForwardingsAliasesStore201Response) SetData(v EmailForwardingAlias)`

SetData sets Data field to given value.

### HasData

`func (o *EmailForwardingsAliasesStore201Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *EmailForwardingsAliasesStore201Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EmailForwardingsAliasesStore201Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EmailForwardingsAliasesStore201Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EmailForwardingsAliasesStore201Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *EmailForwardingsAliasesStore201Response) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *EmailForwardingsAliasesStore201Response) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


