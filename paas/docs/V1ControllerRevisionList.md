# V1ControllerRevisionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**Items** | [**[]V1ControllerRevision**](V1ControllerRevision.md) | Items is the list of ControllerRevisions | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ListMeta**](V1ListMeta.md) |  | [optional] 

## Methods

### NewV1ControllerRevisionList

`func NewV1ControllerRevisionList(items []V1ControllerRevision, ) *V1ControllerRevisionList`

NewV1ControllerRevisionList instantiates a new V1ControllerRevisionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ControllerRevisionListWithDefaults

`func NewV1ControllerRevisionListWithDefaults() *V1ControllerRevisionList`

NewV1ControllerRevisionListWithDefaults instantiates a new V1ControllerRevisionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1ControllerRevisionList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1ControllerRevisionList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1ControllerRevisionList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1ControllerRevisionList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *V1ControllerRevisionList) GetItems() []V1ControllerRevision`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *V1ControllerRevisionList) GetItemsOk() (*[]V1ControllerRevision, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *V1ControllerRevisionList) SetItems(v []V1ControllerRevision)`

SetItems sets Items field to given value.


### GetKind

`func (o *V1ControllerRevisionList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1ControllerRevisionList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1ControllerRevisionList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1ControllerRevisionList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *V1ControllerRevisionList) GetMetadata() V1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1ControllerRevisionList) GetMetadataOk() (*V1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1ControllerRevisionList) SetMetadata(v V1ListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1ControllerRevisionList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


