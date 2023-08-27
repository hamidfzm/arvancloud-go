# V1ReplicaSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**V1ReplicaSetSpec**](V1ReplicaSetSpec.md) |  | [optional] 
**Status** | Pointer to [**V1ReplicaSetStatus**](V1ReplicaSetStatus.md) |  | [optional] 

## Methods

### NewV1ReplicaSet

`func NewV1ReplicaSet() *V1ReplicaSet`

NewV1ReplicaSet instantiates a new V1ReplicaSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ReplicaSetWithDefaults

`func NewV1ReplicaSetWithDefaults() *V1ReplicaSet`

NewV1ReplicaSetWithDefaults instantiates a new V1ReplicaSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1ReplicaSet) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1ReplicaSet) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1ReplicaSet) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1ReplicaSet) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *V1ReplicaSet) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1ReplicaSet) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1ReplicaSet) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1ReplicaSet) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *V1ReplicaSet) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1ReplicaSet) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1ReplicaSet) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1ReplicaSet) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *V1ReplicaSet) GetSpec() V1ReplicaSetSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *V1ReplicaSet) GetSpecOk() (*V1ReplicaSetSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *V1ReplicaSet) SetSpec(v V1ReplicaSetSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *V1ReplicaSet) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *V1ReplicaSet) GetStatus() V1ReplicaSetStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1ReplicaSet) GetStatusOk() (*V1ReplicaSetStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1ReplicaSet) SetStatus(v V1ReplicaSetStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1ReplicaSet) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


