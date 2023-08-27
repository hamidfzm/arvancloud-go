# V1NamespaceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Finalizers** | Pointer to [**[]V1FinalizerName**](V1FinalizerName.md) | Finalizers is an opaque list of values that must be empty to permanently remove object from storage. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/ | [optional] 

## Methods

### NewV1NamespaceSpec

`func NewV1NamespaceSpec() *V1NamespaceSpec`

NewV1NamespaceSpec instantiates a new V1NamespaceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NamespaceSpecWithDefaults

`func NewV1NamespaceSpecWithDefaults() *V1NamespaceSpec`

NewV1NamespaceSpecWithDefaults instantiates a new V1NamespaceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinalizers

`func (o *V1NamespaceSpec) GetFinalizers() []V1FinalizerName`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *V1NamespaceSpec) GetFinalizersOk() (*[]V1FinalizerName, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *V1NamespaceSpec) SetFinalizers(v []V1FinalizerName)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *V1NamespaceSpec) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


