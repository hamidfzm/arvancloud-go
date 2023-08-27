# NsKeysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNsKeysResponse

`func NewNsKeysResponse() *NsKeysResponse`

NewNsKeysResponse instantiates a new NsKeysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNsKeysResponseWithDefaults

`func NewNsKeysResponseWithDefaults() *NsKeysResponse`

NewNsKeysResponseWithDefaults instantiates a new NsKeysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *NsKeysResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NsKeysResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NsKeysResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *NsKeysResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *NsKeysResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NsKeysResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NsKeysResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NsKeysResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *NsKeysResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *NsKeysResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


