# V1TLSConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaCertificate** | Pointer to **string** | caCertificate provides the cert authority certificate contents | [optional] 
**Certificate** | Pointer to **string** | certificate provides certificate contents | [optional] 
**DestinationCACertificate** | Pointer to **string** | destinationCACertificate provides the contents of the ca certificate of the final destination.  When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify. | [optional] 
**InsecureEdgeTerminationPolicy** | Pointer to **string** | insecureEdgeTerminationPolicy indicates the desired behavior for insecure connections to a route. While each router may make its own decisions on which ports to expose, this is normally port 80.  * Allow - traffic is sent to the server on the insecure port (default) * Disable - no traffic is allowed on the insecure port. * Redirect - clients are redirected to the secure port. | [optional] 
**Key** | Pointer to **string** | key provides key file contents | [optional] 
**Termination** | **string** | termination indicates termination type. | 

## Methods

### NewV1TLSConfig

`func NewV1TLSConfig(termination string, ) *V1TLSConfig`

NewV1TLSConfig instantiates a new V1TLSConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TLSConfigWithDefaults

`func NewV1TLSConfigWithDefaults() *V1TLSConfig`

NewV1TLSConfigWithDefaults instantiates a new V1TLSConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaCertificate

`func (o *V1TLSConfig) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *V1TLSConfig) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *V1TLSConfig) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *V1TLSConfig) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetCertificate

`func (o *V1TLSConfig) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *V1TLSConfig) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *V1TLSConfig) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *V1TLSConfig) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetDestinationCACertificate

`func (o *V1TLSConfig) GetDestinationCACertificate() string`

GetDestinationCACertificate returns the DestinationCACertificate field if non-nil, zero value otherwise.

### GetDestinationCACertificateOk

`func (o *V1TLSConfig) GetDestinationCACertificateOk() (*string, bool)`

GetDestinationCACertificateOk returns a tuple with the DestinationCACertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCACertificate

`func (o *V1TLSConfig) SetDestinationCACertificate(v string)`

SetDestinationCACertificate sets DestinationCACertificate field to given value.

### HasDestinationCACertificate

`func (o *V1TLSConfig) HasDestinationCACertificate() bool`

HasDestinationCACertificate returns a boolean if a field has been set.

### GetInsecureEdgeTerminationPolicy

`func (o *V1TLSConfig) GetInsecureEdgeTerminationPolicy() string`

GetInsecureEdgeTerminationPolicy returns the InsecureEdgeTerminationPolicy field if non-nil, zero value otherwise.

### GetInsecureEdgeTerminationPolicyOk

`func (o *V1TLSConfig) GetInsecureEdgeTerminationPolicyOk() (*string, bool)`

GetInsecureEdgeTerminationPolicyOk returns a tuple with the InsecureEdgeTerminationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureEdgeTerminationPolicy

`func (o *V1TLSConfig) SetInsecureEdgeTerminationPolicy(v string)`

SetInsecureEdgeTerminationPolicy sets InsecureEdgeTerminationPolicy field to given value.

### HasInsecureEdgeTerminationPolicy

`func (o *V1TLSConfig) HasInsecureEdgeTerminationPolicy() bool`

HasInsecureEdgeTerminationPolicy returns a boolean if a field has been set.

### GetKey

`func (o *V1TLSConfig) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V1TLSConfig) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V1TLSConfig) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *V1TLSConfig) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetTermination

`func (o *V1TLSConfig) GetTermination() string`

GetTermination returns the Termination field if non-nil, zero value otherwise.

### GetTerminationOk

`func (o *V1TLSConfig) GetTerminationOk() (*string, bool)`

GetTerminationOk returns a tuple with the Termination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermination

`func (o *V1TLSConfig) SetTermination(v string)`

SetTermination sets Termination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


