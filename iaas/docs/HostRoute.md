# HostRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** |  | [optional] 
**Nexthop** | Pointer to **string** |  | [optional] 

## Methods

### NewHostRoute

`func NewHostRoute() *HostRoute`

NewHostRoute instantiates a new HostRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostRouteWithDefaults

`func NewHostRouteWithDefaults() *HostRoute`

NewHostRouteWithDefaults instantiates a new HostRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *HostRoute) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *HostRoute) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *HostRoute) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *HostRoute) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetNexthop

`func (o *HostRoute) GetNexthop() string`

GetNexthop returns the Nexthop field if non-nil, zero value otherwise.

### GetNexthopOk

`func (o *HostRoute) GetNexthopOk() (*string, bool)`

GetNexthopOk returns a tuple with the Nexthop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthop

`func (o *HostRoute) SetNexthop(v string)`

SetNexthop sets Nexthop field to given value.

### HasNexthop

`func (o *HostRoute) HasNexthop() bool`

HasNexthop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


