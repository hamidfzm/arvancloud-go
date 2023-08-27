# V1ConfigMapNodeConfigSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KubeletConfigKey** | **string** | KubeletConfigKey declares which key of the referenced ConfigMap corresponds to the KubeletConfiguration structure This field is required in all cases. | 
**Name** | **string** | Name is the metadata.name of the referenced ConfigMap. This field is required in all cases. | 
**Namespace** | **string** | Namespace is the metadata.namespace of the referenced ConfigMap. This field is required in all cases. | 
**ResourceVersion** | Pointer to **string** | ResourceVersion is the metadata.ResourceVersion of the referenced ConfigMap. This field is forbidden in Node.Spec, and required in Node.Status. | [optional] 
**Uid** | Pointer to **string** | UID is the metadata.UID of the referenced ConfigMap. This field is forbidden in Node.Spec, and required in Node.Status. | [optional] 

## Methods

### NewV1ConfigMapNodeConfigSource

`func NewV1ConfigMapNodeConfigSource(kubeletConfigKey string, name string, namespace string, ) *V1ConfigMapNodeConfigSource`

NewV1ConfigMapNodeConfigSource instantiates a new V1ConfigMapNodeConfigSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ConfigMapNodeConfigSourceWithDefaults

`func NewV1ConfigMapNodeConfigSourceWithDefaults() *V1ConfigMapNodeConfigSource`

NewV1ConfigMapNodeConfigSourceWithDefaults instantiates a new V1ConfigMapNodeConfigSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubeletConfigKey

`func (o *V1ConfigMapNodeConfigSource) GetKubeletConfigKey() string`

GetKubeletConfigKey returns the KubeletConfigKey field if non-nil, zero value otherwise.

### GetKubeletConfigKeyOk

`func (o *V1ConfigMapNodeConfigSource) GetKubeletConfigKeyOk() (*string, bool)`

GetKubeletConfigKeyOk returns a tuple with the KubeletConfigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeletConfigKey

`func (o *V1ConfigMapNodeConfigSource) SetKubeletConfigKey(v string)`

SetKubeletConfigKey sets KubeletConfigKey field to given value.


### GetName

`func (o *V1ConfigMapNodeConfigSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1ConfigMapNodeConfigSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1ConfigMapNodeConfigSource) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *V1ConfigMapNodeConfigSource) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V1ConfigMapNodeConfigSource) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V1ConfigMapNodeConfigSource) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetResourceVersion

`func (o *V1ConfigMapNodeConfigSource) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *V1ConfigMapNodeConfigSource) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *V1ConfigMapNodeConfigSource) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *V1ConfigMapNodeConfigSource) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetUid

`func (o *V1ConfigMapNodeConfigSource) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *V1ConfigMapNodeConfigSource) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *V1ConfigMapNodeConfigSource) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *V1ConfigMapNodeConfigSource) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


