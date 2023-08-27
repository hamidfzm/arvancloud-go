# V1SecurityContextConstraintsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**Items** | [**[]V1SecurityContextConstraints**](V1SecurityContextConstraints.md) | List of security context constraints. | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ListMeta**](V1ListMeta.md) |  | [optional] 

## Methods

### NewV1SecurityContextConstraintsList

`func NewV1SecurityContextConstraintsList(items []V1SecurityContextConstraints, ) *V1SecurityContextConstraintsList`

NewV1SecurityContextConstraintsList instantiates a new V1SecurityContextConstraintsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SecurityContextConstraintsListWithDefaults

`func NewV1SecurityContextConstraintsListWithDefaults() *V1SecurityContextConstraintsList`

NewV1SecurityContextConstraintsListWithDefaults instantiates a new V1SecurityContextConstraintsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1SecurityContextConstraintsList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1SecurityContextConstraintsList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1SecurityContextConstraintsList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1SecurityContextConstraintsList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *V1SecurityContextConstraintsList) GetItems() []V1SecurityContextConstraints`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *V1SecurityContextConstraintsList) GetItemsOk() (*[]V1SecurityContextConstraints, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *V1SecurityContextConstraintsList) SetItems(v []V1SecurityContextConstraints)`

SetItems sets Items field to given value.


### GetKind

`func (o *V1SecurityContextConstraintsList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1SecurityContextConstraintsList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1SecurityContextConstraintsList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1SecurityContextConstraintsList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *V1SecurityContextConstraintsList) GetMetadata() V1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1SecurityContextConstraintsList) GetMetadataOk() (*V1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1SecurityContextConstraintsList) SetMetadata(v V1ListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1SecurityContextConstraintsList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


