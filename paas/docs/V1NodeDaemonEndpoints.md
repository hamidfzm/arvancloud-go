# V1NodeDaemonEndpoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KubeletEndpoint** | Pointer to [**V1DaemonEndpoint**](V1DaemonEndpoint.md) |  | [optional] 

## Methods

### NewV1NodeDaemonEndpoints

`func NewV1NodeDaemonEndpoints() *V1NodeDaemonEndpoints`

NewV1NodeDaemonEndpoints instantiates a new V1NodeDaemonEndpoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NodeDaemonEndpointsWithDefaults

`func NewV1NodeDaemonEndpointsWithDefaults() *V1NodeDaemonEndpoints`

NewV1NodeDaemonEndpointsWithDefaults instantiates a new V1NodeDaemonEndpoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubeletEndpoint

`func (o *V1NodeDaemonEndpoints) GetKubeletEndpoint() V1DaemonEndpoint`

GetKubeletEndpoint returns the KubeletEndpoint field if non-nil, zero value otherwise.

### GetKubeletEndpointOk

`func (o *V1NodeDaemonEndpoints) GetKubeletEndpointOk() (*V1DaemonEndpoint, bool)`

GetKubeletEndpointOk returns a tuple with the KubeletEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeletEndpoint

`func (o *V1NodeDaemonEndpoints) SetKubeletEndpoint(v V1DaemonEndpoint)`

SetKubeletEndpoint sets KubeletEndpoint field to given value.

### HasKubeletEndpoint

`func (o *V1NodeDaemonEndpoints) HasKubeletEndpoint() bool`

HasKubeletEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


