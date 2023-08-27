# V1RepositoryImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**V1ObjectReference**](V1ObjectReference.md) |  | 
**ImportPolicy** | Pointer to [**V1TagImportPolicy**](V1TagImportPolicy.md) |  | [optional] 
**IncludeManifest** | Pointer to **bool** | IncludeManifest determines if the manifest for each image is returned in the response | [optional] 
**ReferencePolicy** | Pointer to [**V1TagReferencePolicy**](V1TagReferencePolicy.md) |  | [optional] 

## Methods

### NewV1RepositoryImportSpec

`func NewV1RepositoryImportSpec(from V1ObjectReference, ) *V1RepositoryImportSpec`

NewV1RepositoryImportSpec instantiates a new V1RepositoryImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RepositoryImportSpecWithDefaults

`func NewV1RepositoryImportSpecWithDefaults() *V1RepositoryImportSpec`

NewV1RepositoryImportSpecWithDefaults instantiates a new V1RepositoryImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *V1RepositoryImportSpec) GetFrom() V1ObjectReference`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *V1RepositoryImportSpec) GetFromOk() (*V1ObjectReference, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *V1RepositoryImportSpec) SetFrom(v V1ObjectReference)`

SetFrom sets From field to given value.


### GetImportPolicy

`func (o *V1RepositoryImportSpec) GetImportPolicy() V1TagImportPolicy`

GetImportPolicy returns the ImportPolicy field if non-nil, zero value otherwise.

### GetImportPolicyOk

`func (o *V1RepositoryImportSpec) GetImportPolicyOk() (*V1TagImportPolicy, bool)`

GetImportPolicyOk returns a tuple with the ImportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPolicy

`func (o *V1RepositoryImportSpec) SetImportPolicy(v V1TagImportPolicy)`

SetImportPolicy sets ImportPolicy field to given value.

### HasImportPolicy

`func (o *V1RepositoryImportSpec) HasImportPolicy() bool`

HasImportPolicy returns a boolean if a field has been set.

### GetIncludeManifest

`func (o *V1RepositoryImportSpec) GetIncludeManifest() bool`

GetIncludeManifest returns the IncludeManifest field if non-nil, zero value otherwise.

### GetIncludeManifestOk

`func (o *V1RepositoryImportSpec) GetIncludeManifestOk() (*bool, bool)`

GetIncludeManifestOk returns a tuple with the IncludeManifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeManifest

`func (o *V1RepositoryImportSpec) SetIncludeManifest(v bool)`

SetIncludeManifest sets IncludeManifest field to given value.

### HasIncludeManifest

`func (o *V1RepositoryImportSpec) HasIncludeManifest() bool`

HasIncludeManifest returns a boolean if a field has been set.

### GetReferencePolicy

`func (o *V1RepositoryImportSpec) GetReferencePolicy() V1TagReferencePolicy`

GetReferencePolicy returns the ReferencePolicy field if non-nil, zero value otherwise.

### GetReferencePolicyOk

`func (o *V1RepositoryImportSpec) GetReferencePolicyOk() (*V1TagReferencePolicy, bool)`

GetReferencePolicyOk returns a tuple with the ReferencePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencePolicy

`func (o *V1RepositoryImportSpec) SetReferencePolicy(v V1TagReferencePolicy)`

SetReferencePolicy sets ReferencePolicy field to given value.

### HasReferencePolicy

`func (o *V1RepositoryImportSpec) HasReferencePolicy() bool`

HasReferencePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


