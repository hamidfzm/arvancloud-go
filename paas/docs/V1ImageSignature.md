# V1ImageSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**Conditions** | Pointer to [**[]V1SignatureCondition**](V1SignatureCondition.md) | Conditions represent the latest available observations of a signature&#39;s current state. | [optional] 
**Content** | **string** | Required: An opaque binary string which is an image&#39;s signature. | 
**Created** | Pointer to **string** | If specified, it is the time of signature&#39;s creation. | [optional] 
**ImageIdentity** | Pointer to **string** | A human readable string representing image&#39;s identity. It could be a product name and version, or an image pull spec (e.g. \&quot;registry.access.redhat.com/rhel7/rhel:7.2\&quot;). | [optional] 
**IssuedBy** | Pointer to [**V1SignatureIssuer**](V1SignatureIssuer.md) |  | [optional] 
**IssuedTo** | Pointer to [**V1SignatureSubject**](V1SignatureSubject.md) |  | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**SignedClaims** | Pointer to **map[string]interface{}** | Contains claims from the signature. | [optional] 
**Type** | **string** | Required: Describes a type of stored blob. | 

## Methods

### NewV1ImageSignature

`func NewV1ImageSignature(content string, type_ string, ) *V1ImageSignature`

NewV1ImageSignature instantiates a new V1ImageSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageSignatureWithDefaults

`func NewV1ImageSignatureWithDefaults() *V1ImageSignature`

NewV1ImageSignatureWithDefaults instantiates a new V1ImageSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1ImageSignature) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1ImageSignature) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1ImageSignature) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1ImageSignature) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetConditions

`func (o *V1ImageSignature) GetConditions() []V1SignatureCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *V1ImageSignature) GetConditionsOk() (*[]V1SignatureCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *V1ImageSignature) SetConditions(v []V1SignatureCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *V1ImageSignature) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetContent

`func (o *V1ImageSignature) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *V1ImageSignature) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *V1ImageSignature) SetContent(v string)`

SetContent sets Content field to given value.


### GetCreated

`func (o *V1ImageSignature) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *V1ImageSignature) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *V1ImageSignature) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *V1ImageSignature) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetImageIdentity

`func (o *V1ImageSignature) GetImageIdentity() string`

GetImageIdentity returns the ImageIdentity field if non-nil, zero value otherwise.

### GetImageIdentityOk

`func (o *V1ImageSignature) GetImageIdentityOk() (*string, bool)`

GetImageIdentityOk returns a tuple with the ImageIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageIdentity

`func (o *V1ImageSignature) SetImageIdentity(v string)`

SetImageIdentity sets ImageIdentity field to given value.

### HasImageIdentity

`func (o *V1ImageSignature) HasImageIdentity() bool`

HasImageIdentity returns a boolean if a field has been set.

### GetIssuedBy

`func (o *V1ImageSignature) GetIssuedBy() V1SignatureIssuer`

GetIssuedBy returns the IssuedBy field if non-nil, zero value otherwise.

### GetIssuedByOk

`func (o *V1ImageSignature) GetIssuedByOk() (*V1SignatureIssuer, bool)`

GetIssuedByOk returns a tuple with the IssuedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedBy

`func (o *V1ImageSignature) SetIssuedBy(v V1SignatureIssuer)`

SetIssuedBy sets IssuedBy field to given value.

### HasIssuedBy

`func (o *V1ImageSignature) HasIssuedBy() bool`

HasIssuedBy returns a boolean if a field has been set.

### GetIssuedTo

`func (o *V1ImageSignature) GetIssuedTo() V1SignatureSubject`

GetIssuedTo returns the IssuedTo field if non-nil, zero value otherwise.

### GetIssuedToOk

`func (o *V1ImageSignature) GetIssuedToOk() (*V1SignatureSubject, bool)`

GetIssuedToOk returns a tuple with the IssuedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedTo

`func (o *V1ImageSignature) SetIssuedTo(v V1SignatureSubject)`

SetIssuedTo sets IssuedTo field to given value.

### HasIssuedTo

`func (o *V1ImageSignature) HasIssuedTo() bool`

HasIssuedTo returns a boolean if a field has been set.

### GetKind

`func (o *V1ImageSignature) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1ImageSignature) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1ImageSignature) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1ImageSignature) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *V1ImageSignature) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1ImageSignature) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1ImageSignature) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1ImageSignature) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSignedClaims

`func (o *V1ImageSignature) GetSignedClaims() map[string]interface{}`

GetSignedClaims returns the SignedClaims field if non-nil, zero value otherwise.

### GetSignedClaimsOk

`func (o *V1ImageSignature) GetSignedClaimsOk() (*map[string]interface{}, bool)`

GetSignedClaimsOk returns a tuple with the SignedClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedClaims

`func (o *V1ImageSignature) SetSignedClaims(v map[string]interface{})`

SetSignedClaims sets SignedClaims field to given value.

### HasSignedClaims

`func (o *V1ImageSignature) HasSignedClaims() bool`

HasSignedClaims returns a boolean if a field has been set.

### GetType

`func (o *V1ImageSignature) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1ImageSignature) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1ImageSignature) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


