# DynamicFieldResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**DynamicField**](DynamicField.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDynamicFieldResponse

`func NewDynamicFieldResponse() *DynamicFieldResponse`

NewDynamicFieldResponse instantiates a new DynamicFieldResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicFieldResponseWithDefaults

`func NewDynamicFieldResponseWithDefaults() *DynamicFieldResponse`

NewDynamicFieldResponseWithDefaults instantiates a new DynamicFieldResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DynamicFieldResponse) GetData() DynamicField`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DynamicFieldResponse) GetDataOk() (*DynamicField, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DynamicFieldResponse) SetData(v DynamicField)`

SetData sets Data field to given value.

### HasData

`func (o *DynamicFieldResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *DynamicFieldResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DynamicFieldResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DynamicFieldResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DynamicFieldResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *DynamicFieldResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *DynamicFieldResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


