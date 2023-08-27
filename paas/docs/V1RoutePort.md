# V1RoutePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetPort** | **string** | The target port on pods selected by the service this route points to. If this is a string, it will be looked up as a named port in the target endpoints port list. Required | 

## Methods

### NewV1RoutePort

`func NewV1RoutePort(targetPort string, ) *V1RoutePort`

NewV1RoutePort instantiates a new V1RoutePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RoutePortWithDefaults

`func NewV1RoutePortWithDefaults() *V1RoutePort`

NewV1RoutePortWithDefaults instantiates a new V1RoutePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetPort

`func (o *V1RoutePort) GetTargetPort() string`

GetTargetPort returns the TargetPort field if non-nil, zero value otherwise.

### GetTargetPortOk

`func (o *V1RoutePort) GetTargetPortOk() (*string, bool)`

GetTargetPortOk returns a tuple with the TargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPort

`func (o *V1RoutePort) SetTargetPort(v string)`

SetTargetPort sets TargetPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


