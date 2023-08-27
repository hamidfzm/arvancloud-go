# V1TagImageHook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerName** | **string** | ContainerName is the name of a container in the deployment config whose image value will be used as the source of the tag. If there is only a single container this value will be defaulted to the name of that container. | 
**To** | [**V1ObjectReference**](V1ObjectReference.md) |  | 

## Methods

### NewV1TagImageHook

`func NewV1TagImageHook(containerName string, to V1ObjectReference, ) *V1TagImageHook`

NewV1TagImageHook instantiates a new V1TagImageHook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TagImageHookWithDefaults

`func NewV1TagImageHookWithDefaults() *V1TagImageHook`

NewV1TagImageHookWithDefaults instantiates a new V1TagImageHook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerName

`func (o *V1TagImageHook) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *V1TagImageHook) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *V1TagImageHook) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetTo

`func (o *V1TagImageHook) GetTo() V1ObjectReference`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *V1TagImageHook) GetToOk() (*V1ObjectReference, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *V1TagImageHook) SetTo(v V1ObjectReference)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


