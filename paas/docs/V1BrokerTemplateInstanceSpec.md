# V1BrokerTemplateInstanceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindingIDs** | Pointer to **[]string** | bindingids is a list of &#39;binding_id&#39;s provided during successive bind calls to the template service broker. | [optional] 
**Secret** | [**V1ObjectReference**](V1ObjectReference.md) |  | 
**TemplateInstance** | [**V1ObjectReference**](V1ObjectReference.md) |  | 

## Methods

### NewV1BrokerTemplateInstanceSpec

`func NewV1BrokerTemplateInstanceSpec(secret V1ObjectReference, templateInstance V1ObjectReference, ) *V1BrokerTemplateInstanceSpec`

NewV1BrokerTemplateInstanceSpec instantiates a new V1BrokerTemplateInstanceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BrokerTemplateInstanceSpecWithDefaults

`func NewV1BrokerTemplateInstanceSpecWithDefaults() *V1BrokerTemplateInstanceSpec`

NewV1BrokerTemplateInstanceSpecWithDefaults instantiates a new V1BrokerTemplateInstanceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindingIDs

`func (o *V1BrokerTemplateInstanceSpec) GetBindingIDs() []string`

GetBindingIDs returns the BindingIDs field if non-nil, zero value otherwise.

### GetBindingIDsOk

`func (o *V1BrokerTemplateInstanceSpec) GetBindingIDsOk() (*[]string, bool)`

GetBindingIDsOk returns a tuple with the BindingIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingIDs

`func (o *V1BrokerTemplateInstanceSpec) SetBindingIDs(v []string)`

SetBindingIDs sets BindingIDs field to given value.

### HasBindingIDs

`func (o *V1BrokerTemplateInstanceSpec) HasBindingIDs() bool`

HasBindingIDs returns a boolean if a field has been set.

### GetSecret

`func (o *V1BrokerTemplateInstanceSpec) GetSecret() V1ObjectReference`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *V1BrokerTemplateInstanceSpec) GetSecretOk() (*V1ObjectReference, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *V1BrokerTemplateInstanceSpec) SetSecret(v V1ObjectReference)`

SetSecret sets Secret field to given value.


### GetTemplateInstance

`func (o *V1BrokerTemplateInstanceSpec) GetTemplateInstance() V1ObjectReference`

GetTemplateInstance returns the TemplateInstance field if non-nil, zero value otherwise.

### GetTemplateInstanceOk

`func (o *V1BrokerTemplateInstanceSpec) GetTemplateInstanceOk() (*V1ObjectReference, bool)`

GetTemplateInstanceOk returns a tuple with the TemplateInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateInstance

`func (o *V1BrokerTemplateInstanceSpec) SetTemplateInstance(v V1ObjectReference)`

SetTemplateInstance sets TemplateInstance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


