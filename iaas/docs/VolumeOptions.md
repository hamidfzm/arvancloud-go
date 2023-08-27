# VolumeOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IopsPerGig** | Pointer to [**IOPSAndThroughput**](IOPSAndThroughput.md) |  | [optional] 
**IopsSecMax** | Pointer to [**MaxIOPSAndThroughput**](MaxIOPSAndThroughput.md) |  | [optional] 
**ThroughputKbPerGig** | Pointer to [**IOPSAndThroughput**](IOPSAndThroughput.md) |  | [optional] 
**ThroughputSecMax** | Pointer to [**MaxIOPSAndThroughput**](MaxIOPSAndThroughput.md) |  | [optional] 
**ThroughputSecMin** | Pointer to [**IOPSAndThroughput**](IOPSAndThroughput.md) |  | [optional] 

## Methods

### NewVolumeOptions

`func NewVolumeOptions() *VolumeOptions`

NewVolumeOptions instantiates a new VolumeOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeOptionsWithDefaults

`func NewVolumeOptionsWithDefaults() *VolumeOptions`

NewVolumeOptionsWithDefaults instantiates a new VolumeOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIopsPerGig

`func (o *VolumeOptions) GetIopsPerGig() IOPSAndThroughput`

GetIopsPerGig returns the IopsPerGig field if non-nil, zero value otherwise.

### GetIopsPerGigOk

`func (o *VolumeOptions) GetIopsPerGigOk() (*IOPSAndThroughput, bool)`

GetIopsPerGigOk returns a tuple with the IopsPerGig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsPerGig

`func (o *VolumeOptions) SetIopsPerGig(v IOPSAndThroughput)`

SetIopsPerGig sets IopsPerGig field to given value.

### HasIopsPerGig

`func (o *VolumeOptions) HasIopsPerGig() bool`

HasIopsPerGig returns a boolean if a field has been set.

### GetIopsSecMax

`func (o *VolumeOptions) GetIopsSecMax() MaxIOPSAndThroughput`

GetIopsSecMax returns the IopsSecMax field if non-nil, zero value otherwise.

### GetIopsSecMaxOk

`func (o *VolumeOptions) GetIopsSecMaxOk() (*MaxIOPSAndThroughput, bool)`

GetIopsSecMaxOk returns a tuple with the IopsSecMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsSecMax

`func (o *VolumeOptions) SetIopsSecMax(v MaxIOPSAndThroughput)`

SetIopsSecMax sets IopsSecMax field to given value.

### HasIopsSecMax

`func (o *VolumeOptions) HasIopsSecMax() bool`

HasIopsSecMax returns a boolean if a field has been set.

### GetThroughputKbPerGig

`func (o *VolumeOptions) GetThroughputKbPerGig() IOPSAndThroughput`

GetThroughputKbPerGig returns the ThroughputKbPerGig field if non-nil, zero value otherwise.

### GetThroughputKbPerGigOk

`func (o *VolumeOptions) GetThroughputKbPerGigOk() (*IOPSAndThroughput, bool)`

GetThroughputKbPerGigOk returns a tuple with the ThroughputKbPerGig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputKbPerGig

`func (o *VolumeOptions) SetThroughputKbPerGig(v IOPSAndThroughput)`

SetThroughputKbPerGig sets ThroughputKbPerGig field to given value.

### HasThroughputKbPerGig

`func (o *VolumeOptions) HasThroughputKbPerGig() bool`

HasThroughputKbPerGig returns a boolean if a field has been set.

### GetThroughputSecMax

`func (o *VolumeOptions) GetThroughputSecMax() MaxIOPSAndThroughput`

GetThroughputSecMax returns the ThroughputSecMax field if non-nil, zero value otherwise.

### GetThroughputSecMaxOk

`func (o *VolumeOptions) GetThroughputSecMaxOk() (*MaxIOPSAndThroughput, bool)`

GetThroughputSecMaxOk returns a tuple with the ThroughputSecMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputSecMax

`func (o *VolumeOptions) SetThroughputSecMax(v MaxIOPSAndThroughput)`

SetThroughputSecMax sets ThroughputSecMax field to given value.

### HasThroughputSecMax

`func (o *VolumeOptions) HasThroughputSecMax() bool`

HasThroughputSecMax returns a boolean if a field has been set.

### GetThroughputSecMin

`func (o *VolumeOptions) GetThroughputSecMin() IOPSAndThroughput`

GetThroughputSecMin returns the ThroughputSecMin field if non-nil, zero value otherwise.

### GetThroughputSecMinOk

`func (o *VolumeOptions) GetThroughputSecMinOk() (*IOPSAndThroughput, bool)`

GetThroughputSecMinOk returns a tuple with the ThroughputSecMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputSecMin

`func (o *VolumeOptions) SetThroughputSecMin(v IOPSAndThroughput)`

SetThroughputSecMin sets ThroughputSecMin field to given value.

### HasThroughputSecMin

`func (o *VolumeOptions) HasThroughputSecMin() bool`

HasThroughputSecMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


