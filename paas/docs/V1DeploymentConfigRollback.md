# V1DeploymentConfigRollback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**Name** | **string** | Name of the deployment config that will be rolled back. | 
**Spec** | [**V1DeploymentConfigRollbackSpec**](V1DeploymentConfigRollbackSpec.md) |  | 
**UpdatedAnnotations** | Pointer to **map[string]interface{}** | UpdatedAnnotations is a set of new annotations that will be added in the deployment config. | [optional] 

## Methods

### NewV1DeploymentConfigRollback

`func NewV1DeploymentConfigRollback(name string, spec V1DeploymentConfigRollbackSpec, ) *V1DeploymentConfigRollback`

NewV1DeploymentConfigRollback instantiates a new V1DeploymentConfigRollback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeploymentConfigRollbackWithDefaults

`func NewV1DeploymentConfigRollbackWithDefaults() *V1DeploymentConfigRollback`

NewV1DeploymentConfigRollbackWithDefaults instantiates a new V1DeploymentConfigRollback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1DeploymentConfigRollback) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1DeploymentConfigRollback) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1DeploymentConfigRollback) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1DeploymentConfigRollback) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *V1DeploymentConfigRollback) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1DeploymentConfigRollback) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1DeploymentConfigRollback) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1DeploymentConfigRollback) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *V1DeploymentConfigRollback) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1DeploymentConfigRollback) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1DeploymentConfigRollback) SetName(v string)`

SetName sets Name field to given value.


### GetSpec

`func (o *V1DeploymentConfigRollback) GetSpec() V1DeploymentConfigRollbackSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *V1DeploymentConfigRollback) GetSpecOk() (*V1DeploymentConfigRollbackSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *V1DeploymentConfigRollback) SetSpec(v V1DeploymentConfigRollbackSpec)`

SetSpec sets Spec field to given value.


### GetUpdatedAnnotations

`func (o *V1DeploymentConfigRollback) GetUpdatedAnnotations() map[string]interface{}`

GetUpdatedAnnotations returns the UpdatedAnnotations field if non-nil, zero value otherwise.

### GetUpdatedAnnotationsOk

`func (o *V1DeploymentConfigRollback) GetUpdatedAnnotationsOk() (*map[string]interface{}, bool)`

GetUpdatedAnnotationsOk returns a tuple with the UpdatedAnnotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAnnotations

`func (o *V1DeploymentConfigRollback) SetUpdatedAnnotations(v map[string]interface{})`

SetUpdatedAnnotations sets UpdatedAnnotations field to given value.

### HasUpdatedAnnotations

`func (o *V1DeploymentConfigRollback) HasUpdatedAnnotations() bool`

HasUpdatedAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


