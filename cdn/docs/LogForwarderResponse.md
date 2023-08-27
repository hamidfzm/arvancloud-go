# LogForwarderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**LogForwarderGeneric**](LogForwarderGeneric.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLogForwarderResponse

`func NewLogForwarderResponse() *LogForwarderResponse`

NewLogForwarderResponse instantiates a new LogForwarderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogForwarderResponseWithDefaults

`func NewLogForwarderResponseWithDefaults() *LogForwarderResponse`

NewLogForwarderResponseWithDefaults instantiates a new LogForwarderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *LogForwarderResponse) GetData() LogForwarderGeneric`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LogForwarderResponse) GetDataOk() (*LogForwarderGeneric, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LogForwarderResponse) SetData(v LogForwarderGeneric)`

SetData sets Data field to given value.

### HasData

`func (o *LogForwarderResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *LogForwarderResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LogForwarderResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LogForwarderResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LogForwarderResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *LogForwarderResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *LogForwarderResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


