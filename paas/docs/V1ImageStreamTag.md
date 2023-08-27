# V1ImageStreamTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**Conditions** | Pointer to [**[]V1TagEventCondition**](V1TagEventCondition.md) | conditions is an array of conditions that apply to the image stream tag. | [optional] 
**Generation** | **int64** | generation is the current generation of the tagged image - if tag is provided and this value is not equal to the tag generation, a user has requested an import that has not completed, or conditions will be filled out indicating any error. | 
**Image** | [**V1Image**](V1Image.md) |  | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**LookupPolicy** | [**V1ImageLookupPolicy**](V1ImageLookupPolicy.md) |  | 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Tag** | [**V1TagReference**](V1TagReference.md) |  | 

## Methods

### NewV1ImageStreamTag

`func NewV1ImageStreamTag(generation int64, image V1Image, lookupPolicy V1ImageLookupPolicy, tag V1TagReference, ) *V1ImageStreamTag`

NewV1ImageStreamTag instantiates a new V1ImageStreamTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageStreamTagWithDefaults

`func NewV1ImageStreamTagWithDefaults() *V1ImageStreamTag`

NewV1ImageStreamTagWithDefaults instantiates a new V1ImageStreamTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1ImageStreamTag) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1ImageStreamTag) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1ImageStreamTag) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1ImageStreamTag) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetConditions

`func (o *V1ImageStreamTag) GetConditions() []V1TagEventCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *V1ImageStreamTag) GetConditionsOk() (*[]V1TagEventCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *V1ImageStreamTag) SetConditions(v []V1TagEventCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *V1ImageStreamTag) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetGeneration

`func (o *V1ImageStreamTag) GetGeneration() int64`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *V1ImageStreamTag) GetGenerationOk() (*int64, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *V1ImageStreamTag) SetGeneration(v int64)`

SetGeneration sets Generation field to given value.


### GetImage

`func (o *V1ImageStreamTag) GetImage() V1Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *V1ImageStreamTag) GetImageOk() (*V1Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *V1ImageStreamTag) SetImage(v V1Image)`

SetImage sets Image field to given value.


### GetKind

`func (o *V1ImageStreamTag) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1ImageStreamTag) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1ImageStreamTag) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1ImageStreamTag) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLookupPolicy

`func (o *V1ImageStreamTag) GetLookupPolicy() V1ImageLookupPolicy`

GetLookupPolicy returns the LookupPolicy field if non-nil, zero value otherwise.

### GetLookupPolicyOk

`func (o *V1ImageStreamTag) GetLookupPolicyOk() (*V1ImageLookupPolicy, bool)`

GetLookupPolicyOk returns a tuple with the LookupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupPolicy

`func (o *V1ImageStreamTag) SetLookupPolicy(v V1ImageLookupPolicy)`

SetLookupPolicy sets LookupPolicy field to given value.


### GetMetadata

`func (o *V1ImageStreamTag) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1ImageStreamTag) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1ImageStreamTag) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1ImageStreamTag) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTag

`func (o *V1ImageStreamTag) GetTag() V1TagReference`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *V1ImageStreamTag) GetTagOk() (*V1TagReference, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *V1ImageStreamTag) SetTag(v V1TagReference)`

SetTag sets Tag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


