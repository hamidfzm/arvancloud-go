# V1ListMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Continue** | Pointer to **string** | continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response. | [optional] 
**ResourceVersion** | Pointer to **string** | String that identifies the server&#39;s internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency | [optional] 
**SelfLink** | Pointer to **string** | selfLink is a URL representing this object. Populated by the system. Read-only. | [optional] 

## Methods

### NewV1ListMeta

`func NewV1ListMeta() *V1ListMeta`

NewV1ListMeta instantiates a new V1ListMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListMetaWithDefaults

`func NewV1ListMetaWithDefaults() *V1ListMeta`

NewV1ListMetaWithDefaults instantiates a new V1ListMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinue

`func (o *V1ListMeta) GetContinue() string`

GetContinue returns the Continue field if non-nil, zero value otherwise.

### GetContinueOk

`func (o *V1ListMeta) GetContinueOk() (*string, bool)`

GetContinueOk returns a tuple with the Continue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinue

`func (o *V1ListMeta) SetContinue(v string)`

SetContinue sets Continue field to given value.

### HasContinue

`func (o *V1ListMeta) HasContinue() bool`

HasContinue returns a boolean if a field has been set.

### GetResourceVersion

`func (o *V1ListMeta) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *V1ListMeta) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *V1ListMeta) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *V1ListMeta) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetSelfLink

`func (o *V1ListMeta) GetSelfLink() string`

GetSelfLink returns the SelfLink field if non-nil, zero value otherwise.

### GetSelfLinkOk

`func (o *V1ListMeta) GetSelfLinkOk() (*string, bool)`

GetSelfLinkOk returns a tuple with the SelfLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfLink

`func (o *V1ListMeta) SetSelfLink(v string)`

SetSelfLink sets SelfLink field to given value.

### HasSelfLink

`func (o *V1ListMeta) HasSelfLink() bool`

HasSelfLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


