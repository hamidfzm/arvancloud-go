# V1ResourceQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**V1ResourceQuotaSpec**](V1ResourceQuotaSpec.md) |  | [optional] 
**Status** | Pointer to [**V1ResourceQuotaStatus**](V1ResourceQuotaStatus.md) |  | [optional] 

## Methods

### NewV1ResourceQuota

`func NewV1ResourceQuota() *V1ResourceQuota`

NewV1ResourceQuota instantiates a new V1ResourceQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ResourceQuotaWithDefaults

`func NewV1ResourceQuotaWithDefaults() *V1ResourceQuota`

NewV1ResourceQuotaWithDefaults instantiates a new V1ResourceQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1ResourceQuota) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1ResourceQuota) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1ResourceQuota) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1ResourceQuota) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *V1ResourceQuota) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1ResourceQuota) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1ResourceQuota) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1ResourceQuota) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *V1ResourceQuota) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1ResourceQuota) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1ResourceQuota) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1ResourceQuota) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *V1ResourceQuota) GetSpec() V1ResourceQuotaSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *V1ResourceQuota) GetSpecOk() (*V1ResourceQuotaSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *V1ResourceQuota) SetSpec(v V1ResourceQuotaSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *V1ResourceQuota) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *V1ResourceQuota) GetStatus() V1ResourceQuotaStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1ResourceQuota) GetStatusOk() (*V1ResourceQuotaStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1ResourceQuota) SetStatus(v V1ResourceQuotaStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1ResourceQuota) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


