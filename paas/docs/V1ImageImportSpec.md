# V1ImageImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**V1ObjectReference**](V1ObjectReference.md) |  | 
**ImportPolicy** | Pointer to [**V1TagImportPolicy**](V1TagImportPolicy.md) |  | [optional] 
**IncludeManifest** | Pointer to **bool** | IncludeManifest determines if the manifest for each image is returned in the response | [optional] 
**ReferencePolicy** | Pointer to [**V1TagReferencePolicy**](V1TagReferencePolicy.md) |  | [optional] 
**To** | Pointer to [**V1LocalObjectReference**](V1LocalObjectReference.md) |  | [optional] 

## Methods

### NewV1ImageImportSpec

`func NewV1ImageImportSpec(from V1ObjectReference, ) *V1ImageImportSpec`

NewV1ImageImportSpec instantiates a new V1ImageImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageImportSpecWithDefaults

`func NewV1ImageImportSpecWithDefaults() *V1ImageImportSpec`

NewV1ImageImportSpecWithDefaults instantiates a new V1ImageImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *V1ImageImportSpec) GetFrom() V1ObjectReference`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *V1ImageImportSpec) GetFromOk() (*V1ObjectReference, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *V1ImageImportSpec) SetFrom(v V1ObjectReference)`

SetFrom sets From field to given value.


### GetImportPolicy

`func (o *V1ImageImportSpec) GetImportPolicy() V1TagImportPolicy`

GetImportPolicy returns the ImportPolicy field if non-nil, zero value otherwise.

### GetImportPolicyOk

`func (o *V1ImageImportSpec) GetImportPolicyOk() (*V1TagImportPolicy, bool)`

GetImportPolicyOk returns a tuple with the ImportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPolicy

`func (o *V1ImageImportSpec) SetImportPolicy(v V1TagImportPolicy)`

SetImportPolicy sets ImportPolicy field to given value.

### HasImportPolicy

`func (o *V1ImageImportSpec) HasImportPolicy() bool`

HasImportPolicy returns a boolean if a field has been set.

### GetIncludeManifest

`func (o *V1ImageImportSpec) GetIncludeManifest() bool`

GetIncludeManifest returns the IncludeManifest field if non-nil, zero value otherwise.

### GetIncludeManifestOk

`func (o *V1ImageImportSpec) GetIncludeManifestOk() (*bool, bool)`

GetIncludeManifestOk returns a tuple with the IncludeManifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeManifest

`func (o *V1ImageImportSpec) SetIncludeManifest(v bool)`

SetIncludeManifest sets IncludeManifest field to given value.

### HasIncludeManifest

`func (o *V1ImageImportSpec) HasIncludeManifest() bool`

HasIncludeManifest returns a boolean if a field has been set.

### GetReferencePolicy

`func (o *V1ImageImportSpec) GetReferencePolicy() V1TagReferencePolicy`

GetReferencePolicy returns the ReferencePolicy field if non-nil, zero value otherwise.

### GetReferencePolicyOk

`func (o *V1ImageImportSpec) GetReferencePolicyOk() (*V1TagReferencePolicy, bool)`

GetReferencePolicyOk returns a tuple with the ReferencePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencePolicy

`func (o *V1ImageImportSpec) SetReferencePolicy(v V1TagReferencePolicy)`

SetReferencePolicy sets ReferencePolicy field to given value.

### HasReferencePolicy

`func (o *V1ImageImportSpec) HasReferencePolicy() bool`

HasReferencePolicy returns a boolean if a field has been set.

### GetTo

`func (o *V1ImageImportSpec) GetTo() V1LocalObjectReference`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *V1ImageImportSpec) GetToOk() (*V1LocalObjectReference, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *V1ImageImportSpec) SetTo(v V1LocalObjectReference)`

SetTo sets To field to given value.

### HasTo

`func (o *V1ImageImportSpec) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


