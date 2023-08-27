# LoadBalancerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**LoadBalancer**](LoadBalancer.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLoadBalancerResponse

`func NewLoadBalancerResponse() *LoadBalancerResponse`

NewLoadBalancerResponse instantiates a new LoadBalancerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerResponseWithDefaults

`func NewLoadBalancerResponseWithDefaults() *LoadBalancerResponse`

NewLoadBalancerResponseWithDefaults instantiates a new LoadBalancerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *LoadBalancerResponse) GetData() LoadBalancer`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LoadBalancerResponse) GetDataOk() (*LoadBalancer, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LoadBalancerResponse) SetData(v LoadBalancer)`

SetData sets Data field to given value.

### HasData

`func (o *LoadBalancerResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *LoadBalancerResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LoadBalancerResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LoadBalancerResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LoadBalancerResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *LoadBalancerResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *LoadBalancerResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


