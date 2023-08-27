# V1ComponentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**Conditions** | Pointer to [**[]V1ComponentCondition**](V1ComponentCondition.md) | List of component conditions observed | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 

## Methods

### NewV1ComponentStatus

`func NewV1ComponentStatus() *V1ComponentStatus`

NewV1ComponentStatus instantiates a new V1ComponentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ComponentStatusWithDefaults

`func NewV1ComponentStatusWithDefaults() *V1ComponentStatus`

NewV1ComponentStatusWithDefaults instantiates a new V1ComponentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1ComponentStatus) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1ComponentStatus) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1ComponentStatus) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1ComponentStatus) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetConditions

`func (o *V1ComponentStatus) GetConditions() []V1ComponentCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *V1ComponentStatus) GetConditionsOk() (*[]V1ComponentCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *V1ComponentStatus) SetConditions(v []V1ComponentCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *V1ComponentStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetKind

`func (o *V1ComponentStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1ComponentStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1ComponentStatus) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1ComponentStatus) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *V1ComponentStatus) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1ComponentStatus) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1ComponentStatus) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1ComponentStatus) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


