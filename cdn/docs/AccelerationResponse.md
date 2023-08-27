# AccelerationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Acceleration**](Acceleration.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAccelerationResponse

`func NewAccelerationResponse() *AccelerationResponse`

NewAccelerationResponse instantiates a new AccelerationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccelerationResponseWithDefaults

`func NewAccelerationResponseWithDefaults() *AccelerationResponse`

NewAccelerationResponseWithDefaults instantiates a new AccelerationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AccelerationResponse) GetData() Acceleration`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccelerationResponse) GetDataOk() (*Acceleration, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccelerationResponse) SetData(v Acceleration)`

SetData sets Data field to given value.

### HasData

`func (o *AccelerationResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *AccelerationResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AccelerationResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AccelerationResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AccelerationResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *AccelerationResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AccelerationResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


