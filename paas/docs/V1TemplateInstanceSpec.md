# V1TemplateInstanceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requester** | [**V1TemplateInstanceRequester**](V1TemplateInstanceRequester.md) |  | 
**Secret** | Pointer to [**V1LocalObjectReference**](V1LocalObjectReference.md) |  | [optional] 
**Template** | [**V1Template**](V1Template.md) |  | 

## Methods

### NewV1TemplateInstanceSpec

`func NewV1TemplateInstanceSpec(requester V1TemplateInstanceRequester, template V1Template, ) *V1TemplateInstanceSpec`

NewV1TemplateInstanceSpec instantiates a new V1TemplateInstanceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TemplateInstanceSpecWithDefaults

`func NewV1TemplateInstanceSpecWithDefaults() *V1TemplateInstanceSpec`

NewV1TemplateInstanceSpecWithDefaults instantiates a new V1TemplateInstanceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequester

`func (o *V1TemplateInstanceSpec) GetRequester() V1TemplateInstanceRequester`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *V1TemplateInstanceSpec) GetRequesterOk() (*V1TemplateInstanceRequester, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *V1TemplateInstanceSpec) SetRequester(v V1TemplateInstanceRequester)`

SetRequester sets Requester field to given value.


### GetSecret

`func (o *V1TemplateInstanceSpec) GetSecret() V1LocalObjectReference`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *V1TemplateInstanceSpec) GetSecretOk() (*V1LocalObjectReference, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *V1TemplateInstanceSpec) SetSecret(v V1LocalObjectReference)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *V1TemplateInstanceSpec) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetTemplate

`func (o *V1TemplateInstanceSpec) GetTemplate() V1Template`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *V1TemplateInstanceSpec) GetTemplateOk() (*V1Template, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *V1TemplateInstanceSpec) SetTemplate(v V1Template)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


