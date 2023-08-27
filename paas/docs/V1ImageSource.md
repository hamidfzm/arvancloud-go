# V1ImageSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | [**[]V1ImageSourcePath**](V1ImageSourcePath.md) | paths is a list of source and destination paths to copy from the image. This content will be copied into the build context prior to starting the build. If no paths are set, the build context will not be altered. | 
**As** | **[]string** | A list of image names that this source will be used in place of during a multi-stage Docker image build. For instance, a Dockerfile that uses \&quot;COPY --from&#x3D;nginx:latest\&quot; will first check for an image source that has \&quot;nginx:latest\&quot; in this field before attempting to pull directly. If the Dockerfile does not reference an image source it is ignored. This field and paths may both be set, in which case the contents will be used twice. | 
**From** | [**V1ObjectReference**](V1ObjectReference.md) |  | 
**PullSecret** | Pointer to [**V1LocalObjectReference**](V1LocalObjectReference.md) |  | [optional] 

## Methods

### NewV1ImageSource

`func NewV1ImageSource(paths []V1ImageSourcePath, as []string, from V1ObjectReference, ) *V1ImageSource`

NewV1ImageSource instantiates a new V1ImageSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageSourceWithDefaults

`func NewV1ImageSourceWithDefaults() *V1ImageSource`

NewV1ImageSourceWithDefaults instantiates a new V1ImageSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *V1ImageSource) GetPaths() []V1ImageSourcePath`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *V1ImageSource) GetPathsOk() (*[]V1ImageSourcePath, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *V1ImageSource) SetPaths(v []V1ImageSourcePath)`

SetPaths sets Paths field to given value.


### GetAs

`func (o *V1ImageSource) GetAs() []string`

GetAs returns the As field if non-nil, zero value otherwise.

### GetAsOk

`func (o *V1ImageSource) GetAsOk() (*[]string, bool)`

GetAsOk returns a tuple with the As field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAs

`func (o *V1ImageSource) SetAs(v []string)`

SetAs sets As field to given value.


### GetFrom

`func (o *V1ImageSource) GetFrom() V1ObjectReference`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *V1ImageSource) GetFromOk() (*V1ObjectReference, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *V1ImageSource) SetFrom(v V1ObjectReference)`

SetFrom sets From field to given value.


### GetPullSecret

`func (o *V1ImageSource) GetPullSecret() V1LocalObjectReference`

GetPullSecret returns the PullSecret field if non-nil, zero value otherwise.

### GetPullSecretOk

`func (o *V1ImageSource) GetPullSecretOk() (*V1LocalObjectReference, bool)`

GetPullSecretOk returns a tuple with the PullSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullSecret

`func (o *V1ImageSource) SetPullSecret(v V1LocalObjectReference)`

SetPullSecret sets PullSecret field to given value.

### HasPullSecret

`func (o *V1ImageSource) HasPullSecret() bool`

HasPullSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


