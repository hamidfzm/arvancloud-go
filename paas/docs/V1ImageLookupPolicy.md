# V1ImageLookupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Local** | **bool** | local will change the docker short image references (like \&quot;mysql\&quot; or \&quot;php:latest\&quot;) on objects in this namespace to the image ID whenever they match this image stream, instead of reaching out to a remote registry. The name will be fully qualified to an image ID if found. The tag&#39;s referencePolicy is taken into account on the replaced value. Only works within the current namespace. | 

## Methods

### NewV1ImageLookupPolicy

`func NewV1ImageLookupPolicy(local bool, ) *V1ImageLookupPolicy`

NewV1ImageLookupPolicy instantiates a new V1ImageLookupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageLookupPolicyWithDefaults

`func NewV1ImageLookupPolicyWithDefaults() *V1ImageLookupPolicy`

NewV1ImageLookupPolicyWithDefaults instantiates a new V1ImageLookupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocal

`func (o *V1ImageLookupPolicy) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *V1ImageLookupPolicy) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *V1ImageLookupPolicy) SetLocal(v bool)`

SetLocal sets Local field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


