# TransportLayerProxyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**TransportLayerProxy**](TransportLayerProxy.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTransportLayerProxyResponse

`func NewTransportLayerProxyResponse() *TransportLayerProxyResponse`

NewTransportLayerProxyResponse instantiates a new TransportLayerProxyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportLayerProxyResponseWithDefaults

`func NewTransportLayerProxyResponseWithDefaults() *TransportLayerProxyResponse`

NewTransportLayerProxyResponseWithDefaults instantiates a new TransportLayerProxyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TransportLayerProxyResponse) GetData() TransportLayerProxy`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransportLayerProxyResponse) GetDataOk() (*TransportLayerProxy, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransportLayerProxyResponse) SetData(v TransportLayerProxy)`

SetData sets Data field to given value.

### HasData

`func (o *TransportLayerProxyResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *TransportLayerProxyResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TransportLayerProxyResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TransportLayerProxyResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TransportLayerProxyResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *TransportLayerProxyResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *TransportLayerProxyResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


