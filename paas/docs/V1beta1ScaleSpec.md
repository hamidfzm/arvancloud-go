# V1beta1ScaleSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Replicas** | Pointer to **int32** | desired number of instances for the scaled object. | [optional] 

## Methods

### NewV1beta1ScaleSpec

`func NewV1beta1ScaleSpec() *V1beta1ScaleSpec`

NewV1beta1ScaleSpec instantiates a new V1beta1ScaleSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1beta1ScaleSpecWithDefaults

`func NewV1beta1ScaleSpecWithDefaults() *V1beta1ScaleSpec`

NewV1beta1ScaleSpecWithDefaults instantiates a new V1beta1ScaleSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicas

`func (o *V1beta1ScaleSpec) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *V1beta1ScaleSpec) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *V1beta1ScaleSpec) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *V1beta1ScaleSpec) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


