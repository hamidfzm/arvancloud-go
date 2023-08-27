# V1RouteStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ingress** | [**[]V1RouteIngress**](V1RouteIngress.md) | ingress describes the places where the route may be exposed. The list of ingress points may contain duplicate Host or RouterName values. Routes are considered live once they are &#x60;Ready&#x60; | 

## Methods

### NewV1RouteStatus

`func NewV1RouteStatus(ingress []V1RouteIngress, ) *V1RouteStatus`

NewV1RouteStatus instantiates a new V1RouteStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RouteStatusWithDefaults

`func NewV1RouteStatusWithDefaults() *V1RouteStatus`

NewV1RouteStatusWithDefaults instantiates a new V1RouteStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngress

`func (o *V1RouteStatus) GetIngress() []V1RouteIngress`

GetIngress returns the Ingress field if non-nil, zero value otherwise.

### GetIngressOk

`func (o *V1RouteStatus) GetIngressOk() (*[]V1RouteIngress, bool)`

GetIngressOk returns a tuple with the Ingress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngress

`func (o *V1RouteStatus) SetIngress(v []V1RouteIngress)`

SetIngress sets Ingress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


