# V1RouteSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | host is an alias/DNS that points to the service. Optional. If not specified a route name will typically be automatically chosen. Must follow DNS952 subdomain conventions. | 
**AlternateBackends** | Pointer to [**[]V1RouteTargetReference**](V1RouteTargetReference.md) | alternateBackends allows up to 3 additional backends to be assigned to the route. Only the Service kind is allowed, and it will be defaulted to Service. Use the weight field in RouteTargetReference object to specify relative preference. | [optional] 
**Path** | Pointer to **string** | Path that the router watches for, to route traffic for to the service. Optional | [optional] 
**Port** | Pointer to [**V1RoutePort**](V1RoutePort.md) |  | [optional] 
**Tls** | Pointer to [**V1TLSConfig**](V1TLSConfig.md) |  | [optional] 
**To** | [**V1RouteTargetReference**](V1RouteTargetReference.md) |  | 
**WildcardPolicy** | Pointer to **string** | Wildcard policy if any for the route. Currently only &#39;Subdomain&#39; or &#39;None&#39; is allowed. | [optional] 

## Methods

### NewV1RouteSpec

`func NewV1RouteSpec(host string, to V1RouteTargetReference, ) *V1RouteSpec`

NewV1RouteSpec instantiates a new V1RouteSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RouteSpecWithDefaults

`func NewV1RouteSpecWithDefaults() *V1RouteSpec`

NewV1RouteSpecWithDefaults instantiates a new V1RouteSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *V1RouteSpec) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *V1RouteSpec) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *V1RouteSpec) SetHost(v string)`

SetHost sets Host field to given value.


### GetAlternateBackends

`func (o *V1RouteSpec) GetAlternateBackends() []V1RouteTargetReference`

GetAlternateBackends returns the AlternateBackends field if non-nil, zero value otherwise.

### GetAlternateBackendsOk

`func (o *V1RouteSpec) GetAlternateBackendsOk() (*[]V1RouteTargetReference, bool)`

GetAlternateBackendsOk returns a tuple with the AlternateBackends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateBackends

`func (o *V1RouteSpec) SetAlternateBackends(v []V1RouteTargetReference)`

SetAlternateBackends sets AlternateBackends field to given value.

### HasAlternateBackends

`func (o *V1RouteSpec) HasAlternateBackends() bool`

HasAlternateBackends returns a boolean if a field has been set.

### GetPath

`func (o *V1RouteSpec) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *V1RouteSpec) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *V1RouteSpec) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *V1RouteSpec) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPort

`func (o *V1RouteSpec) GetPort() V1RoutePort`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *V1RouteSpec) GetPortOk() (*V1RoutePort, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *V1RouteSpec) SetPort(v V1RoutePort)`

SetPort sets Port field to given value.

### HasPort

`func (o *V1RouteSpec) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetTls

`func (o *V1RouteSpec) GetTls() V1TLSConfig`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *V1RouteSpec) GetTlsOk() (*V1TLSConfig, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *V1RouteSpec) SetTls(v V1TLSConfig)`

SetTls sets Tls field to given value.

### HasTls

`func (o *V1RouteSpec) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetTo

`func (o *V1RouteSpec) GetTo() V1RouteTargetReference`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *V1RouteSpec) GetToOk() (*V1RouteTargetReference, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *V1RouteSpec) SetTo(v V1RouteTargetReference)`

SetTo sets To field to given value.


### GetWildcardPolicy

`func (o *V1RouteSpec) GetWildcardPolicy() string`

GetWildcardPolicy returns the WildcardPolicy field if non-nil, zero value otherwise.

### GetWildcardPolicyOk

`func (o *V1RouteSpec) GetWildcardPolicyOk() (*string, bool)`

GetWildcardPolicyOk returns a tuple with the WildcardPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcardPolicy

`func (o *V1RouteSpec) SetWildcardPolicy(v string)`

SetWildcardPolicy sets WildcardPolicy field to given value.

### HasWildcardPolicy

`func (o *V1RouteSpec) HasWildcardPolicy() bool`

HasWildcardPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


