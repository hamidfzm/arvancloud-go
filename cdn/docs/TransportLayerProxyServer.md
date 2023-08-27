# TransportLayerProxyServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Port** | **int32** |  | 
**Weight** | **int32** |  | 
**Check** | **string** |  | 
**Fall** | Pointer to **int32** |  | [optional] 
**Inter** | Pointer to **int32** |  | [optional] 
**Rise** | Pointer to **int32** |  | [optional] 

## Methods

### NewTransportLayerProxyServer

`func NewTransportLayerProxyServer(address string, port int32, weight int32, check string, ) *TransportLayerProxyServer`

NewTransportLayerProxyServer instantiates a new TransportLayerProxyServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportLayerProxyServerWithDefaults

`func NewTransportLayerProxyServerWithDefaults() *TransportLayerProxyServer`

NewTransportLayerProxyServerWithDefaults instantiates a new TransportLayerProxyServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TransportLayerProxyServer) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransportLayerProxyServer) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransportLayerProxyServer) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPort

`func (o *TransportLayerProxyServer) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TransportLayerProxyServer) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TransportLayerProxyServer) SetPort(v int32)`

SetPort sets Port field to given value.


### GetWeight

`func (o *TransportLayerProxyServer) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *TransportLayerProxyServer) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *TransportLayerProxyServer) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetCheck

`func (o *TransportLayerProxyServer) GetCheck() string`

GetCheck returns the Check field if non-nil, zero value otherwise.

### GetCheckOk

`func (o *TransportLayerProxyServer) GetCheckOk() (*string, bool)`

GetCheckOk returns a tuple with the Check field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheck

`func (o *TransportLayerProxyServer) SetCheck(v string)`

SetCheck sets Check field to given value.


### GetFall

`func (o *TransportLayerProxyServer) GetFall() int32`

GetFall returns the Fall field if non-nil, zero value otherwise.

### GetFallOk

`func (o *TransportLayerProxyServer) GetFallOk() (*int32, bool)`

GetFallOk returns a tuple with the Fall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFall

`func (o *TransportLayerProxyServer) SetFall(v int32)`

SetFall sets Fall field to given value.

### HasFall

`func (o *TransportLayerProxyServer) HasFall() bool`

HasFall returns a boolean if a field has been set.

### GetInter

`func (o *TransportLayerProxyServer) GetInter() int32`

GetInter returns the Inter field if non-nil, zero value otherwise.

### GetInterOk

`func (o *TransportLayerProxyServer) GetInterOk() (*int32, bool)`

GetInterOk returns a tuple with the Inter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInter

`func (o *TransportLayerProxyServer) SetInter(v int32)`

SetInter sets Inter field to given value.

### HasInter

`func (o *TransportLayerProxyServer) HasInter() bool`

HasInter returns a boolean if a field has been set.

### GetRise

`func (o *TransportLayerProxyServer) GetRise() int32`

GetRise returns the Rise field if non-nil, zero value otherwise.

### GetRiseOk

`func (o *TransportLayerProxyServer) GetRiseOk() (*int32, bool)`

GetRiseOk returns a tuple with the Rise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRise

`func (o *TransportLayerProxyServer) SetRise(v int32)`

SetRise sets Rise field to given value.

### HasRise

`func (o *TransportLayerProxyServer) HasRise() bool`

HasRise returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


