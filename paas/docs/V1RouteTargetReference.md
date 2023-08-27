# V1RouteTargetReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The kind of target that the route is referring to. Currently, only &#39;Service&#39; is allowed | 
**Name** | **string** | name of the service/target that is being referred to. e.g. name of the service | 
**Weight** | **int32** | weight as an integer between 0 and 256, default 1, that specifies the target&#39;s relative weight against other target reference objects. 0 suppresses requests to this backend. | 

## Methods

### NewV1RouteTargetReference

`func NewV1RouteTargetReference(kind string, name string, weight int32, ) *V1RouteTargetReference`

NewV1RouteTargetReference instantiates a new V1RouteTargetReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RouteTargetReferenceWithDefaults

`func NewV1RouteTargetReferenceWithDefaults() *V1RouteTargetReference`

NewV1RouteTargetReferenceWithDefaults instantiates a new V1RouteTargetReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *V1RouteTargetReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1RouteTargetReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1RouteTargetReference) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *V1RouteTargetReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1RouteTargetReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1RouteTargetReference) SetName(v string)`

SetName sets Name field to given value.


### GetWeight

`func (o *V1RouteTargetReference) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V1RouteTargetReference) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V1RouteTargetReference) SetWeight(v int32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


