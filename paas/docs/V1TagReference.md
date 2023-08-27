# V1TagReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | **map[string]interface{}** | Optional; if specified, annotations that are applied to images retrieved via ImageStreamTags. | 
**From** | Pointer to [**V1ObjectReference**](V1ObjectReference.md) |  | [optional] 
**Generation** | **int64** | Generation is a counter that tracks mutations to the spec tag (user intent). When a tag reference is changed the generation is set to match the current stream generation (which is incremented every time spec is changed). Other processes in the system like the image importer observe that the generation of spec tag is newer than the generation recorded in the status and use that as a trigger to import the newest remote tag. To trigger a new import, clients may set this value to zero which will reset the generation to the latest stream generation. Legacy clients will send this value as nil which will be merged with the current tag generation. | 
**ImportPolicy** | Pointer to [**V1TagImportPolicy**](V1TagImportPolicy.md) |  | [optional] 
**Name** | **string** | Name of the tag | 
**Reference** | Pointer to **bool** | Reference states if the tag will be imported. Default value is false, which means the tag will be imported. | [optional] 
**ReferencePolicy** | Pointer to [**V1TagReferencePolicy**](V1TagReferencePolicy.md) |  | [optional] 

## Methods

### NewV1TagReference

`func NewV1TagReference(annotations map[string]interface{}, generation int64, name string, ) *V1TagReference`

NewV1TagReference instantiates a new V1TagReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TagReferenceWithDefaults

`func NewV1TagReferenceWithDefaults() *V1TagReference`

NewV1TagReferenceWithDefaults instantiates a new V1TagReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *V1TagReference) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *V1TagReference) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *V1TagReference) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.


### GetFrom

`func (o *V1TagReference) GetFrom() V1ObjectReference`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *V1TagReference) GetFromOk() (*V1ObjectReference, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *V1TagReference) SetFrom(v V1ObjectReference)`

SetFrom sets From field to given value.

### HasFrom

`func (o *V1TagReference) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetGeneration

`func (o *V1TagReference) GetGeneration() int64`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *V1TagReference) GetGenerationOk() (*int64, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *V1TagReference) SetGeneration(v int64)`

SetGeneration sets Generation field to given value.


### GetImportPolicy

`func (o *V1TagReference) GetImportPolicy() V1TagImportPolicy`

GetImportPolicy returns the ImportPolicy field if non-nil, zero value otherwise.

### GetImportPolicyOk

`func (o *V1TagReference) GetImportPolicyOk() (*V1TagImportPolicy, bool)`

GetImportPolicyOk returns a tuple with the ImportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPolicy

`func (o *V1TagReference) SetImportPolicy(v V1TagImportPolicy)`

SetImportPolicy sets ImportPolicy field to given value.

### HasImportPolicy

`func (o *V1TagReference) HasImportPolicy() bool`

HasImportPolicy returns a boolean if a field has been set.

### GetName

`func (o *V1TagReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1TagReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1TagReference) SetName(v string)`

SetName sets Name field to given value.


### GetReference

`func (o *V1TagReference) GetReference() bool`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *V1TagReference) GetReferenceOk() (*bool, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *V1TagReference) SetReference(v bool)`

SetReference sets Reference field to given value.

### HasReference

`func (o *V1TagReference) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferencePolicy

`func (o *V1TagReference) GetReferencePolicy() V1TagReferencePolicy`

GetReferencePolicy returns the ReferencePolicy field if non-nil, zero value otherwise.

### GetReferencePolicyOk

`func (o *V1TagReference) GetReferencePolicyOk() (*V1TagReferencePolicy, bool)`

GetReferencePolicyOk returns a tuple with the ReferencePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencePolicy

`func (o *V1TagReference) SetReferencePolicy(v V1TagReferencePolicy)`

SetReferencePolicy sets ReferencePolicy field to given value.

### HasReferencePolicy

`func (o *V1TagReference) HasReferencePolicy() bool`

HasReferencePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


