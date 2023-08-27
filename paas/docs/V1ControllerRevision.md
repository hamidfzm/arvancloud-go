# V1ControllerRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**Data** | Pointer to **string** | Data is the serialized representation of the state. | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Revision** | **int64** | Revision indicates the revision of the state represented by Data. | 

## Methods

### NewV1ControllerRevision

`func NewV1ControllerRevision(revision int64, ) *V1ControllerRevision`

NewV1ControllerRevision instantiates a new V1ControllerRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ControllerRevisionWithDefaults

`func NewV1ControllerRevisionWithDefaults() *V1ControllerRevision`

NewV1ControllerRevisionWithDefaults instantiates a new V1ControllerRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1ControllerRevision) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1ControllerRevision) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1ControllerRevision) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1ControllerRevision) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetData

`func (o *V1ControllerRevision) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V1ControllerRevision) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V1ControllerRevision) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *V1ControllerRevision) HasData() bool`

HasData returns a boolean if a field has been set.

### GetKind

`func (o *V1ControllerRevision) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1ControllerRevision) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1ControllerRevision) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1ControllerRevision) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *V1ControllerRevision) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1ControllerRevision) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1ControllerRevision) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1ControllerRevision) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRevision

`func (o *V1ControllerRevision) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *V1ControllerRevision) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *V1ControllerRevision) SetRevision(v int64)`

SetRevision sets Revision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


