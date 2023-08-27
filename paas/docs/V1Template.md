# V1Template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | Pointer to [**[]V1Parameter**](V1Parameter.md) | parameters is an optional array of Parameters used during the Template to Config transformation. | [optional] 
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**Labels** | Pointer to **map[string]interface{}** | labels is a optional set of labels that are applied to every object during the Template to Config transformation. | [optional] 
**Message** | Pointer to **string** | message is an optional instructional message that will be displayed when this template is instantiated. This field should inform the user how to utilize the newly created resources. Parameter substitution will be performed on the message before being displayed so that generated credentials and other parameters can be included in the output. | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Objects** | [**[]RuntimeRawExtension**](RuntimeRawExtension.md) | objects is an array of resources to include in this template. If a namespace value is hardcoded in the object, it will be removed during template instantiation, however if the namespace value is, or contains, a ${PARAMETER_REFERENCE}, the resolved value after parameter substitution will be respected and the object will be created in that namespace. | 

## Methods

### NewV1Template

`func NewV1Template(objects []RuntimeRawExtension, ) *V1Template`

NewV1Template instantiates a new V1Template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TemplateWithDefaults

`func NewV1TemplateWithDefaults() *V1Template`

NewV1TemplateWithDefaults instantiates a new V1Template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *V1Template) GetParameters() []V1Parameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *V1Template) GetParametersOk() (*[]V1Parameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *V1Template) SetParameters(v []V1Parameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *V1Template) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetApiVersion

`func (o *V1Template) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1Template) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1Template) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1Template) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *V1Template) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1Template) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1Template) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1Template) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabels

`func (o *V1Template) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *V1Template) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *V1Template) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *V1Template) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMessage

`func (o *V1Template) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1Template) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1Template) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1Template) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMetadata

`func (o *V1Template) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1Template) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1Template) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1Template) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetObjects

`func (o *V1Template) GetObjects() []RuntimeRawExtension`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *V1Template) GetObjectsOk() (*[]RuntimeRawExtension, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *V1Template) SetObjects(v []RuntimeRawExtension)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


