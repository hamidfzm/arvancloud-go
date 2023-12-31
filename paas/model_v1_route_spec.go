/*
Arvancloud PaaS REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paas

import (
	"encoding/json"
)

// checks if the V1RouteSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1RouteSpec{}

// V1RouteSpec RouteSpec describes the hostname or path the route exposes, any security information, and one to four backends (services) the route points to. Requests are distributed among the backends depending on the weights assigned to each backend. When using roundrobin scheduling the portion of requests that go to each backend is the backend weight divided by the sum of all of the backend weights. When the backend has more than one endpoint the requests that end up on the backend are roundrobin distributed among the endpoints. Weights are between 0 and 256 with default 1. Weight 0 causes no requests to the backend. If all weights are zero the route will be considered to have no backends and return a standard 503 response.  The `tls` field is optional and allows specific certificates or behavior for the route. Routers typically configure a default certificate on a wildcard domain to terminate routes without explicit certificates, but custom hostnames usually must choose passthrough (send traffic directly to the backend via the TLS Server-Name- Indication field) or provide a certificate.
type V1RouteSpec struct {
	// host is an alias/DNS that points to the service. Optional. If not specified a route name will typically be automatically chosen. Must follow DNS952 subdomain conventions.
	Host string `json:"host"`
	// alternateBackends allows up to 3 additional backends to be assigned to the route. Only the Service kind is allowed, and it will be defaulted to Service. Use the weight field in RouteTargetReference object to specify relative preference.
	AlternateBackends []V1RouteTargetReference `json:"alternateBackends,omitempty"`
	// Path that the router watches for, to route traffic for to the service. Optional
	Path *string `json:"path,omitempty"`
	Port *V1RoutePort `json:"port,omitempty"`
	Tls *V1TLSConfig `json:"tls,omitempty"`
	To V1RouteTargetReference `json:"to"`
	// Wildcard policy if any for the route. Currently only 'Subdomain' or 'None' is allowed.
	WildcardPolicy *string `json:"wildcardPolicy,omitempty"`
}

// NewV1RouteSpec instantiates a new V1RouteSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1RouteSpec(host string, to V1RouteTargetReference) *V1RouteSpec {
	this := V1RouteSpec{}
	this.Host = host
	this.To = to
	return &this
}

// NewV1RouteSpecWithDefaults instantiates a new V1RouteSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1RouteSpecWithDefaults() *V1RouteSpec {
	this := V1RouteSpec{}
	return &this
}

// GetHost returns the Host field value
func (o *V1RouteSpec) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *V1RouteSpec) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *V1RouteSpec) SetHost(v string) {
	o.Host = v
}

// GetAlternateBackends returns the AlternateBackends field value if set, zero value otherwise.
func (o *V1RouteSpec) GetAlternateBackends() []V1RouteTargetReference {
	if o == nil || IsNil(o.AlternateBackends) {
		var ret []V1RouteTargetReference
		return ret
	}
	return o.AlternateBackends
}

// GetAlternateBackendsOk returns a tuple with the AlternateBackends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RouteSpec) GetAlternateBackendsOk() ([]V1RouteTargetReference, bool) {
	if o == nil || IsNil(o.AlternateBackends) {
		return nil, false
	}
	return o.AlternateBackends, true
}

// HasAlternateBackends returns a boolean if a field has been set.
func (o *V1RouteSpec) HasAlternateBackends() bool {
	if o != nil && !IsNil(o.AlternateBackends) {
		return true
	}

	return false
}

// SetAlternateBackends gets a reference to the given []V1RouteTargetReference and assigns it to the AlternateBackends field.
func (o *V1RouteSpec) SetAlternateBackends(v []V1RouteTargetReference) {
	o.AlternateBackends = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *V1RouteSpec) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RouteSpec) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *V1RouteSpec) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *V1RouteSpec) SetPath(v string) {
	o.Path = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *V1RouteSpec) GetPort() V1RoutePort {
	if o == nil || IsNil(o.Port) {
		var ret V1RoutePort
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RouteSpec) GetPortOk() (*V1RoutePort, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *V1RouteSpec) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given V1RoutePort and assigns it to the Port field.
func (o *V1RouteSpec) SetPort(v V1RoutePort) {
	o.Port = &v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *V1RouteSpec) GetTls() V1TLSConfig {
	if o == nil || IsNil(o.Tls) {
		var ret V1TLSConfig
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RouteSpec) GetTlsOk() (*V1TLSConfig, bool) {
	if o == nil || IsNil(o.Tls) {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *V1RouteSpec) HasTls() bool {
	if o != nil && !IsNil(o.Tls) {
		return true
	}

	return false
}

// SetTls gets a reference to the given V1TLSConfig and assigns it to the Tls field.
func (o *V1RouteSpec) SetTls(v V1TLSConfig) {
	o.Tls = &v
}

// GetTo returns the To field value
func (o *V1RouteSpec) GetTo() V1RouteTargetReference {
	if o == nil {
		var ret V1RouteTargetReference
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *V1RouteSpec) GetToOk() (*V1RouteTargetReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *V1RouteSpec) SetTo(v V1RouteTargetReference) {
	o.To = v
}

// GetWildcardPolicy returns the WildcardPolicy field value if set, zero value otherwise.
func (o *V1RouteSpec) GetWildcardPolicy() string {
	if o == nil || IsNil(o.WildcardPolicy) {
		var ret string
		return ret
	}
	return *o.WildcardPolicy
}

// GetWildcardPolicyOk returns a tuple with the WildcardPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RouteSpec) GetWildcardPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.WildcardPolicy) {
		return nil, false
	}
	return o.WildcardPolicy, true
}

// HasWildcardPolicy returns a boolean if a field has been set.
func (o *V1RouteSpec) HasWildcardPolicy() bool {
	if o != nil && !IsNil(o.WildcardPolicy) {
		return true
	}

	return false
}

// SetWildcardPolicy gets a reference to the given string and assigns it to the WildcardPolicy field.
func (o *V1RouteSpec) SetWildcardPolicy(v string) {
	o.WildcardPolicy = &v
}

func (o V1RouteSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1RouteSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	if !IsNil(o.AlternateBackends) {
		toSerialize["alternateBackends"] = o.AlternateBackends
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Tls) {
		toSerialize["tls"] = o.Tls
	}
	toSerialize["to"] = o.To
	if !IsNil(o.WildcardPolicy) {
		toSerialize["wildcardPolicy"] = o.WildcardPolicy
	}
	return toSerialize, nil
}

type NullableV1RouteSpec struct {
	value *V1RouteSpec
	isSet bool
}

func (v NullableV1RouteSpec) Get() *V1RouteSpec {
	return v.value
}

func (v *NullableV1RouteSpec) Set(val *V1RouteSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableV1RouteSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableV1RouteSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1RouteSpec(val *V1RouteSpec) *NullableV1RouteSpec {
	return &NullableV1RouteSpec{value: val, isSet: true}
}

func (v NullableV1RouteSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1RouteSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


