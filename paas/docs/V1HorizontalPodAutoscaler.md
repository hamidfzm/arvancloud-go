# V1HorizontalPodAutoscaler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**V1HorizontalPodAutoscalerSpec**](V1HorizontalPodAutoscalerSpec.md) |  | [optional] 
**Status** | Pointer to [**V1HorizontalPodAutoscalerStatus**](V1HorizontalPodAutoscalerStatus.md) |  | [optional] 

## Methods

### NewV1HorizontalPodAutoscaler

`func NewV1HorizontalPodAutoscaler() *V1HorizontalPodAutoscaler`

NewV1HorizontalPodAutoscaler instantiates a new V1HorizontalPodAutoscaler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1HorizontalPodAutoscalerWithDefaults

`func NewV1HorizontalPodAutoscalerWithDefaults() *V1HorizontalPodAutoscaler`

NewV1HorizontalPodAutoscalerWithDefaults instantiates a new V1HorizontalPodAutoscaler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1HorizontalPodAutoscaler) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1HorizontalPodAutoscaler) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1HorizontalPodAutoscaler) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1HorizontalPodAutoscaler) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *V1HorizontalPodAutoscaler) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1HorizontalPodAutoscaler) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1HorizontalPodAutoscaler) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1HorizontalPodAutoscaler) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *V1HorizontalPodAutoscaler) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1HorizontalPodAutoscaler) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1HorizontalPodAutoscaler) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1HorizontalPodAutoscaler) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *V1HorizontalPodAutoscaler) GetSpec() V1HorizontalPodAutoscalerSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *V1HorizontalPodAutoscaler) GetSpecOk() (*V1HorizontalPodAutoscalerSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *V1HorizontalPodAutoscaler) SetSpec(v V1HorizontalPodAutoscalerSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *V1HorizontalPodAutoscaler) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *V1HorizontalPodAutoscaler) GetStatus() V1HorizontalPodAutoscalerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1HorizontalPodAutoscaler) GetStatusOk() (*V1HorizontalPodAutoscalerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1HorizontalPodAutoscaler) SetStatus(v V1HorizontalPodAutoscalerStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1HorizontalPodAutoscaler) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


