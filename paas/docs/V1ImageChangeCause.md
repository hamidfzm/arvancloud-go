# V1ImageChangeCause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromRef** | Pointer to [**V1ObjectReference**](V1ObjectReference.md) |  | [optional] 
**ImageID** | Pointer to **string** | imageID is the ID of the image that triggered a a new build. | [optional] 

## Methods

### NewV1ImageChangeCause

`func NewV1ImageChangeCause() *V1ImageChangeCause`

NewV1ImageChangeCause instantiates a new V1ImageChangeCause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageChangeCauseWithDefaults

`func NewV1ImageChangeCauseWithDefaults() *V1ImageChangeCause`

NewV1ImageChangeCauseWithDefaults instantiates a new V1ImageChangeCause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromRef

`func (o *V1ImageChangeCause) GetFromRef() V1ObjectReference`

GetFromRef returns the FromRef field if non-nil, zero value otherwise.

### GetFromRefOk

`func (o *V1ImageChangeCause) GetFromRefOk() (*V1ObjectReference, bool)`

GetFromRefOk returns a tuple with the FromRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromRef

`func (o *V1ImageChangeCause) SetFromRef(v V1ObjectReference)`

SetFromRef sets FromRef field to given value.

### HasFromRef

`func (o *V1ImageChangeCause) HasFromRef() bool`

HasFromRef returns a boolean if a field has been set.

### GetImageID

`func (o *V1ImageChangeCause) GetImageID() string`

GetImageID returns the ImageID field if non-nil, zero value otherwise.

### GetImageIDOk

`func (o *V1ImageChangeCause) GetImageIDOk() (*string, bool)`

GetImageIDOk returns a tuple with the ImageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageID

`func (o *V1ImageChangeCause) SetImageID(v string)`

SetImageID sets ImageID field to given value.

### HasImageID

`func (o *V1ImageChangeCause) HasImageID() bool`

HasImageID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


