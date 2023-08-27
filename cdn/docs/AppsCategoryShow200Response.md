# AppsCategoryShow200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ApplicationCategory**](ApplicationCategory.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAppsCategoryShow200Response

`func NewAppsCategoryShow200Response() *AppsCategoryShow200Response`

NewAppsCategoryShow200Response instantiates a new AppsCategoryShow200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppsCategoryShow200ResponseWithDefaults

`func NewAppsCategoryShow200ResponseWithDefaults() *AppsCategoryShow200Response`

NewAppsCategoryShow200ResponseWithDefaults instantiates a new AppsCategoryShow200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppsCategoryShow200Response) GetData() ApplicationCategory`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppsCategoryShow200Response) GetDataOk() (*ApplicationCategory, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppsCategoryShow200Response) SetData(v ApplicationCategory)`

SetData sets Data field to given value.

### HasData

`func (o *AppsCategoryShow200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *AppsCategoryShow200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AppsCategoryShow200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AppsCategoryShow200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AppsCategoryShow200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *AppsCategoryShow200Response) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AppsCategoryShow200Response) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


