# V1RouteIngress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Host is the host string under which the route is exposed; this value is required | [optional] 
**Conditions** | Pointer to [**[]V1RouteIngressCondition**](V1RouteIngressCondition.md) | Conditions is the state of the route, may be empty. | [optional] 
**RouterCanonicalHostname** | Pointer to **string** | CanonicalHostname is the external host name for the router that can be used as a CNAME for the host requested for this route. This value is optional and may not be set in all cases. | [optional] 
**RouterName** | Pointer to **string** | Name is a name chosen by the router to identify itself; this value is required | [optional] 
**WildcardPolicy** | Pointer to **string** | Wildcard policy is the wildcard policy that was allowed where this route is exposed. | [optional] 

## Methods

### NewV1RouteIngress

`func NewV1RouteIngress() *V1RouteIngress`

NewV1RouteIngress instantiates a new V1RouteIngress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RouteIngressWithDefaults

`func NewV1RouteIngressWithDefaults() *V1RouteIngress`

NewV1RouteIngressWithDefaults instantiates a new V1RouteIngress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *V1RouteIngress) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *V1RouteIngress) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *V1RouteIngress) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *V1RouteIngress) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetConditions

`func (o *V1RouteIngress) GetConditions() []V1RouteIngressCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *V1RouteIngress) GetConditionsOk() (*[]V1RouteIngressCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *V1RouteIngress) SetConditions(v []V1RouteIngressCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *V1RouteIngress) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetRouterCanonicalHostname

`func (o *V1RouteIngress) GetRouterCanonicalHostname() string`

GetRouterCanonicalHostname returns the RouterCanonicalHostname field if non-nil, zero value otherwise.

### GetRouterCanonicalHostnameOk

`func (o *V1RouteIngress) GetRouterCanonicalHostnameOk() (*string, bool)`

GetRouterCanonicalHostnameOk returns a tuple with the RouterCanonicalHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterCanonicalHostname

`func (o *V1RouteIngress) SetRouterCanonicalHostname(v string)`

SetRouterCanonicalHostname sets RouterCanonicalHostname field to given value.

### HasRouterCanonicalHostname

`func (o *V1RouteIngress) HasRouterCanonicalHostname() bool`

HasRouterCanonicalHostname returns a boolean if a field has been set.

### GetRouterName

`func (o *V1RouteIngress) GetRouterName() string`

GetRouterName returns the RouterName field if non-nil, zero value otherwise.

### GetRouterNameOk

`func (o *V1RouteIngress) GetRouterNameOk() (*string, bool)`

GetRouterNameOk returns a tuple with the RouterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterName

`func (o *V1RouteIngress) SetRouterName(v string)`

SetRouterName sets RouterName field to given value.

### HasRouterName

`func (o *V1RouteIngress) HasRouterName() bool`

HasRouterName returns a boolean if a field has been set.

### GetWildcardPolicy

`func (o *V1RouteIngress) GetWildcardPolicy() string`

GetWildcardPolicy returns the WildcardPolicy field if non-nil, zero value otherwise.

### GetWildcardPolicyOk

`func (o *V1RouteIngress) GetWildcardPolicyOk() (*string, bool)`

GetWildcardPolicyOk returns a tuple with the WildcardPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcardPolicy

`func (o *V1RouteIngress) SetWildcardPolicy(v string)`

SetWildcardPolicy sets WildcardPolicy field to given value.

### HasWildcardPolicy

`func (o *V1RouteIngress) HasWildcardPolicy() bool`

HasWildcardPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


