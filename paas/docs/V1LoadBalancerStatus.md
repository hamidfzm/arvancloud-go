# V1LoadBalancerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ingress** | Pointer to [**[]V1LoadBalancerIngress**](V1LoadBalancerIngress.md) | Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points. | [optional] 

## Methods

### NewV1LoadBalancerStatus

`func NewV1LoadBalancerStatus() *V1LoadBalancerStatus`

NewV1LoadBalancerStatus instantiates a new V1LoadBalancerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1LoadBalancerStatusWithDefaults

`func NewV1LoadBalancerStatusWithDefaults() *V1LoadBalancerStatus`

NewV1LoadBalancerStatusWithDefaults instantiates a new V1LoadBalancerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngress

`func (o *V1LoadBalancerStatus) GetIngress() []V1LoadBalancerIngress`

GetIngress returns the Ingress field if non-nil, zero value otherwise.

### GetIngressOk

`func (o *V1LoadBalancerStatus) GetIngressOk() (*[]V1LoadBalancerIngress, bool)`

GetIngressOk returns a tuple with the Ingress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngress

`func (o *V1LoadBalancerStatus) SetIngress(v []V1LoadBalancerIngress)`

SetIngress sets Ingress field to given value.

### HasIngress

`func (o *V1LoadBalancerStatus) HasIngress() bool`

HasIngress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


