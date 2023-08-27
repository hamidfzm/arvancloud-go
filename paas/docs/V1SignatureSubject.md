# V1SignatureSubject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonName** | Pointer to **string** | Common name (e.g. openshift-signing-service). | [optional] 
**Organization** | Pointer to **string** | Organization name. | [optional] 
**PublicKeyID** | **string** | If present, it is a human readable key id of public key belonging to the subject used to verify image signature. It should contain at least 64 lowest bits of public key&#39;s fingerprint (e.g. 0x685ebe62bf278440). | 

## Methods

### NewV1SignatureSubject

`func NewV1SignatureSubject(publicKeyID string, ) *V1SignatureSubject`

NewV1SignatureSubject instantiates a new V1SignatureSubject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SignatureSubjectWithDefaults

`func NewV1SignatureSubjectWithDefaults() *V1SignatureSubject`

NewV1SignatureSubjectWithDefaults instantiates a new V1SignatureSubject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonName

`func (o *V1SignatureSubject) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *V1SignatureSubject) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *V1SignatureSubject) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *V1SignatureSubject) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetOrganization

`func (o *V1SignatureSubject) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *V1SignatureSubject) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *V1SignatureSubject) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *V1SignatureSubject) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPublicKeyID

`func (o *V1SignatureSubject) GetPublicKeyID() string`

GetPublicKeyID returns the PublicKeyID field if non-nil, zero value otherwise.

### GetPublicKeyIDOk

`func (o *V1SignatureSubject) GetPublicKeyIDOk() (*string, bool)`

GetPublicKeyIDOk returns a tuple with the PublicKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyID

`func (o *V1SignatureSubject) SetPublicKeyID(v string)`

SetPublicKeyID sets PublicKeyID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


