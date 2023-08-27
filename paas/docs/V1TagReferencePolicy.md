# V1TagReferencePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type determines how the image pull spec should be transformed when the image stream tag is used in deployment config triggers or new builds. The default value is &#x60;Source&#x60;, indicating the original location of the image should be used (if imported). The user may also specify &#x60;Local&#x60;, indicating that the pull spec should point to the integrated Docker registry and leverage the registry&#39;s ability to proxy the pull to an upstream registry. &#x60;Local&#x60; allows the credentials used to pull this image to be managed from the image stream&#39;s namespace, so others on the platform can access a remote image but have no access to the remote secret. It also allows the image layers to be mirrored into the local registry which the images can still be pulled even if the upstream registry is unavailable. | 

## Methods

### NewV1TagReferencePolicy

`func NewV1TagReferencePolicy(type_ string, ) *V1TagReferencePolicy`

NewV1TagReferencePolicy instantiates a new V1TagReferencePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TagReferencePolicyWithDefaults

`func NewV1TagReferencePolicyWithDefaults() *V1TagReferencePolicy`

NewV1TagReferencePolicyWithDefaults instantiates a new V1TagReferencePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *V1TagReferencePolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1TagReferencePolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1TagReferencePolicy) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


