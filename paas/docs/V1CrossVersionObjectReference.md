# V1CrossVersionObjectReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | API version of the referent | [optional] 
**Kind** | **string** | Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds\&quot; | 
**Name** | **string** | Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names | 

## Methods

### NewV1CrossVersionObjectReference

`func NewV1CrossVersionObjectReference(kind string, name string, ) *V1CrossVersionObjectReference`

NewV1CrossVersionObjectReference instantiates a new V1CrossVersionObjectReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1CrossVersionObjectReferenceWithDefaults

`func NewV1CrossVersionObjectReferenceWithDefaults() *V1CrossVersionObjectReference`

NewV1CrossVersionObjectReferenceWithDefaults instantiates a new V1CrossVersionObjectReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1CrossVersionObjectReference) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1CrossVersionObjectReference) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1CrossVersionObjectReference) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1CrossVersionObjectReference) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *V1CrossVersionObjectReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1CrossVersionObjectReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1CrossVersionObjectReference) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *V1CrossVersionObjectReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1CrossVersionObjectReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1CrossVersionObjectReference) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


