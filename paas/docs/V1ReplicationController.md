# V1ReplicationController

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**V1ReplicationControllerSpec**](V1ReplicationControllerSpec.md) |  | [optional] 
**Status** | Pointer to [**V1ReplicationControllerStatus**](V1ReplicationControllerStatus.md) |  | [optional] 

## Methods

### NewV1ReplicationController

`func NewV1ReplicationController() *V1ReplicationController`

NewV1ReplicationController instantiates a new V1ReplicationController object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ReplicationControllerWithDefaults

`func NewV1ReplicationControllerWithDefaults() *V1ReplicationController`

NewV1ReplicationControllerWithDefaults instantiates a new V1ReplicationController object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1ReplicationController) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1ReplicationController) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1ReplicationController) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1ReplicationController) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *V1ReplicationController) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1ReplicationController) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1ReplicationController) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1ReplicationController) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *V1ReplicationController) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1ReplicationController) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1ReplicationController) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1ReplicationController) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *V1ReplicationController) GetSpec() V1ReplicationControllerSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *V1ReplicationController) GetSpecOk() (*V1ReplicationControllerSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *V1ReplicationController) SetSpec(v V1ReplicationControllerSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *V1ReplicationController) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *V1ReplicationController) GetStatus() V1ReplicationControllerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1ReplicationController) GetStatusOk() (*V1ReplicationControllerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1ReplicationController) SetStatus(v V1ReplicationControllerStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1ReplicationController) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


