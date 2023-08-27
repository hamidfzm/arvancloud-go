# V1SignatureIssuer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonName** | Pointer to **string** | Common name (e.g. openshift-signing-service). | [optional] 
**Organization** | Pointer to **string** | Organization name. | [optional] 

## Methods

### NewV1SignatureIssuer

`func NewV1SignatureIssuer() *V1SignatureIssuer`

NewV1SignatureIssuer instantiates a new V1SignatureIssuer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SignatureIssuerWithDefaults

`func NewV1SignatureIssuerWithDefaults() *V1SignatureIssuer`

NewV1SignatureIssuerWithDefaults instantiates a new V1SignatureIssuer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonName

`func (o *V1SignatureIssuer) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *V1SignatureIssuer) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *V1SignatureIssuer) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *V1SignatureIssuer) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetOrganization

`func (o *V1SignatureIssuer) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *V1SignatureIssuer) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *V1SignatureIssuer) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *V1SignatureIssuer) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


