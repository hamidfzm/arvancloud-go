# V1beta1Eviction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**DeleteOptions** | Pointer to [**V1DeleteOptions**](V1DeleteOptions.md) |  | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 

## Methods

### NewV1beta1Eviction

`func NewV1beta1Eviction() *V1beta1Eviction`

NewV1beta1Eviction instantiates a new V1beta1Eviction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1beta1EvictionWithDefaults

`func NewV1beta1EvictionWithDefaults() *V1beta1Eviction`

NewV1beta1EvictionWithDefaults instantiates a new V1beta1Eviction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1beta1Eviction) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1beta1Eviction) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1beta1Eviction) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1beta1Eviction) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetDeleteOptions

`func (o *V1beta1Eviction) GetDeleteOptions() V1DeleteOptions`

GetDeleteOptions returns the DeleteOptions field if non-nil, zero value otherwise.

### GetDeleteOptionsOk

`func (o *V1beta1Eviction) GetDeleteOptionsOk() (*V1DeleteOptions, bool)`

GetDeleteOptionsOk returns a tuple with the DeleteOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOptions

`func (o *V1beta1Eviction) SetDeleteOptions(v V1DeleteOptions)`

SetDeleteOptions sets DeleteOptions field to given value.

### HasDeleteOptions

`func (o *V1beta1Eviction) HasDeleteOptions() bool`

HasDeleteOptions returns a boolean if a field has been set.

### GetKind

`func (o *V1beta1Eviction) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1beta1Eviction) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1beta1Eviction) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1beta1Eviction) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *V1beta1Eviction) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1beta1Eviction) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1beta1Eviction) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1beta1Eviction) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


