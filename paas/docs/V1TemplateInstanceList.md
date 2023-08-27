# V1TemplateInstanceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**Items** | [**[]V1TemplateInstance**](V1TemplateInstance.md) | items is a list of Templateinstances | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ListMeta**](V1ListMeta.md) |  | [optional] 

## Methods

### NewV1TemplateInstanceList

`func NewV1TemplateInstanceList(items []V1TemplateInstance, ) *V1TemplateInstanceList`

NewV1TemplateInstanceList instantiates a new V1TemplateInstanceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TemplateInstanceListWithDefaults

`func NewV1TemplateInstanceListWithDefaults() *V1TemplateInstanceList`

NewV1TemplateInstanceListWithDefaults instantiates a new V1TemplateInstanceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1TemplateInstanceList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1TemplateInstanceList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1TemplateInstanceList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1TemplateInstanceList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *V1TemplateInstanceList) GetItems() []V1TemplateInstance`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *V1TemplateInstanceList) GetItemsOk() (*[]V1TemplateInstance, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *V1TemplateInstanceList) SetItems(v []V1TemplateInstance)`

SetItems sets Items field to given value.


### GetKind

`func (o *V1TemplateInstanceList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1TemplateInstanceList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1TemplateInstanceList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1TemplateInstanceList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *V1TemplateInstanceList) GetMetadata() V1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1TemplateInstanceList) GetMetadataOk() (*V1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1TemplateInstanceList) SetMetadata(v V1ListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1TemplateInstanceList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


