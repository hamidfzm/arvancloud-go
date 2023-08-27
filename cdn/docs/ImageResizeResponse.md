# ImageResizeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ImageResize**](ImageResize.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewImageResizeResponse

`func NewImageResizeResponse() *ImageResizeResponse`

NewImageResizeResponse instantiates a new ImageResizeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageResizeResponseWithDefaults

`func NewImageResizeResponseWithDefaults() *ImageResizeResponse`

NewImageResizeResponseWithDefaults instantiates a new ImageResizeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ImageResizeResponse) GetData() ImageResize`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ImageResizeResponse) GetDataOk() (*ImageResize, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ImageResizeResponse) SetData(v ImageResize)`

SetData sets Data field to given value.

### HasData

`func (o *ImageResizeResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *ImageResizeResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ImageResizeResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ImageResizeResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ImageResizeResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ImageResizeResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ImageResizeResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


