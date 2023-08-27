# V1TemplateInstanceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]V1TemplateInstanceCondition**](V1TemplateInstanceCondition.md) | conditions represent the latest available observations of a TemplateInstance&#39;s current state. | [optional] 
**Objects** | Pointer to [**[]V1TemplateInstanceObject**](V1TemplateInstanceObject.md) | Objects references the objects created by the TemplateInstance. | [optional] 

## Methods

### NewV1TemplateInstanceStatus

`func NewV1TemplateInstanceStatus() *V1TemplateInstanceStatus`

NewV1TemplateInstanceStatus instantiates a new V1TemplateInstanceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TemplateInstanceStatusWithDefaults

`func NewV1TemplateInstanceStatusWithDefaults() *V1TemplateInstanceStatus`

NewV1TemplateInstanceStatusWithDefaults instantiates a new V1TemplateInstanceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *V1TemplateInstanceStatus) GetConditions() []V1TemplateInstanceCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *V1TemplateInstanceStatus) GetConditionsOk() (*[]V1TemplateInstanceCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *V1TemplateInstanceStatus) SetConditions(v []V1TemplateInstanceCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *V1TemplateInstanceStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetObjects

`func (o *V1TemplateInstanceStatus) GetObjects() []V1TemplateInstanceObject`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *V1TemplateInstanceStatus) GetObjectsOk() (*[]V1TemplateInstanceObject, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *V1TemplateInstanceStatus) SetObjects(v []V1TemplateInstanceObject)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *V1TemplateInstanceStatus) HasObjects() bool`

HasObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


