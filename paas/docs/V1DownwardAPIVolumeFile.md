# V1DownwardAPIVolumeFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldRef** | Pointer to [**V1ObjectFieldSelector**](V1ObjectFieldSelector.md) |  | [optional] 
**Mode** | Pointer to **int32** | Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set. | [optional] 
**Path** | **string** | Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the &#39;..&#39; path. Must be utf-8 encoded. The first item of the relative path must not start with &#39;..&#39; | 
**ResourceFieldRef** | Pointer to [**V1ResourceFieldSelector**](V1ResourceFieldSelector.md) |  | [optional] 

## Methods

### NewV1DownwardAPIVolumeFile

`func NewV1DownwardAPIVolumeFile(path string, ) *V1DownwardAPIVolumeFile`

NewV1DownwardAPIVolumeFile instantiates a new V1DownwardAPIVolumeFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DownwardAPIVolumeFileWithDefaults

`func NewV1DownwardAPIVolumeFileWithDefaults() *V1DownwardAPIVolumeFile`

NewV1DownwardAPIVolumeFileWithDefaults instantiates a new V1DownwardAPIVolumeFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldRef

`func (o *V1DownwardAPIVolumeFile) GetFieldRef() V1ObjectFieldSelector`

GetFieldRef returns the FieldRef field if non-nil, zero value otherwise.

### GetFieldRefOk

`func (o *V1DownwardAPIVolumeFile) GetFieldRefOk() (*V1ObjectFieldSelector, bool)`

GetFieldRefOk returns a tuple with the FieldRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldRef

`func (o *V1DownwardAPIVolumeFile) SetFieldRef(v V1ObjectFieldSelector)`

SetFieldRef sets FieldRef field to given value.

### HasFieldRef

`func (o *V1DownwardAPIVolumeFile) HasFieldRef() bool`

HasFieldRef returns a boolean if a field has been set.

### GetMode

`func (o *V1DownwardAPIVolumeFile) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *V1DownwardAPIVolumeFile) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *V1DownwardAPIVolumeFile) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *V1DownwardAPIVolumeFile) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPath

`func (o *V1DownwardAPIVolumeFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *V1DownwardAPIVolumeFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *V1DownwardAPIVolumeFile) SetPath(v string)`

SetPath sets Path field to given value.


### GetResourceFieldRef

`func (o *V1DownwardAPIVolumeFile) GetResourceFieldRef() V1ResourceFieldSelector`

GetResourceFieldRef returns the ResourceFieldRef field if non-nil, zero value otherwise.

### GetResourceFieldRefOk

`func (o *V1DownwardAPIVolumeFile) GetResourceFieldRefOk() (*V1ResourceFieldSelector, bool)`

GetResourceFieldRefOk returns a tuple with the ResourceFieldRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceFieldRef

`func (o *V1DownwardAPIVolumeFile) SetResourceFieldRef(v V1ResourceFieldSelector)`

SetResourceFieldRef sets ResourceFieldRef field to given value.

### HasResourceFieldRef

`func (o *V1DownwardAPIVolumeFile) HasResourceFieldRef() bool`

HasResourceFieldRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


