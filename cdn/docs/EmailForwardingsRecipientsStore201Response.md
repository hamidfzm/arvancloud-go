# EmailForwardingsRecipientsStore201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**EmailForwardingRecipient**](EmailForwardingRecipient.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEmailForwardingsRecipientsStore201Response

`func NewEmailForwardingsRecipientsStore201Response() *EmailForwardingsRecipientsStore201Response`

NewEmailForwardingsRecipientsStore201Response instantiates a new EmailForwardingsRecipientsStore201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailForwardingsRecipientsStore201ResponseWithDefaults

`func NewEmailForwardingsRecipientsStore201ResponseWithDefaults() *EmailForwardingsRecipientsStore201Response`

NewEmailForwardingsRecipientsStore201ResponseWithDefaults instantiates a new EmailForwardingsRecipientsStore201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EmailForwardingsRecipientsStore201Response) GetData() EmailForwardingRecipient`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EmailForwardingsRecipientsStore201Response) GetDataOk() (*EmailForwardingRecipient, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EmailForwardingsRecipientsStore201Response) SetData(v EmailForwardingRecipient)`

SetData sets Data field to given value.

### HasData

`func (o *EmailForwardingsRecipientsStore201Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *EmailForwardingsRecipientsStore201Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EmailForwardingsRecipientsStore201Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EmailForwardingsRecipientsStore201Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EmailForwardingsRecipientsStore201Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *EmailForwardingsRecipientsStore201Response) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *EmailForwardingsRecipientsStore201Response) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


