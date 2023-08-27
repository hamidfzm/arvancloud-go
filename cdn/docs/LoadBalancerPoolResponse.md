# LoadBalancerPoolResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]LoadBalancerPool**](LoadBalancerPool.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLoadBalancerPoolResponse

`func NewLoadBalancerPoolResponse() *LoadBalancerPoolResponse`

NewLoadBalancerPoolResponse instantiates a new LoadBalancerPoolResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerPoolResponseWithDefaults

`func NewLoadBalancerPoolResponseWithDefaults() *LoadBalancerPoolResponse`

NewLoadBalancerPoolResponseWithDefaults instantiates a new LoadBalancerPoolResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *LoadBalancerPoolResponse) GetData() []LoadBalancerPool`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LoadBalancerPoolResponse) GetDataOk() (*[]LoadBalancerPool, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LoadBalancerPoolResponse) SetData(v []LoadBalancerPool)`

SetData sets Data field to given value.

### HasData

`func (o *LoadBalancerPoolResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *LoadBalancerPoolResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LoadBalancerPoolResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LoadBalancerPoolResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LoadBalancerPoolResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *LoadBalancerPoolResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *LoadBalancerPoolResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


