# Plan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandwidthInBytes** | Pointer to **int64** |  | [optional] 
**BasePackage** | Pointer to **string** |  | [optional] 
**Canary** | Pointer to **bool** |  | [optional] 
**CpuCount** | Pointer to **int64** |  | [optional] 
**CpuShare** | Pointer to **string** |  | [optional] 
**CreateType** | Pointer to **string** |  | [optional] 
**Disk** | Pointer to **int64** |  | [optional] 
**DiskInBytes** | Pointer to **int64** |  | [optional] 
**Generation** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IopsMaxHdd** | Pointer to **int64** |  | [optional] 
**IopsMaxSsd** | Pointer to **int64** |  | [optional] 
**Memory** | Pointer to **float64** |  | [optional] 
**MemoryInBytes** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Off** | Pointer to **string** |  | [optional] 
**OffPercent** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**Outbound** | Pointer to **int64** |  | [optional] 
**PortSpeed** | Pointer to **float64** |  | [optional] 
**Pps** | Pointer to **[]int64** |  | [optional] 
**PricePerHour** | Pointer to **int64** |  | [optional] 
**PricePerMonth** | Pointer to **int64** |  | [optional] 
**Subtype** | Pointer to **string** |  | [optional] 
**Throughput** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewPlan

`func NewPlan() *Plan`

NewPlan instantiates a new Plan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanWithDefaults

`func NewPlanWithDefaults() *Plan`

NewPlanWithDefaults instantiates a new Plan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidthInBytes

`func (o *Plan) GetBandwidthInBytes() int64`

GetBandwidthInBytes returns the BandwidthInBytes field if non-nil, zero value otherwise.

### GetBandwidthInBytesOk

`func (o *Plan) GetBandwidthInBytesOk() (*int64, bool)`

GetBandwidthInBytesOk returns a tuple with the BandwidthInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthInBytes

`func (o *Plan) SetBandwidthInBytes(v int64)`

SetBandwidthInBytes sets BandwidthInBytes field to given value.

### HasBandwidthInBytes

`func (o *Plan) HasBandwidthInBytes() bool`

HasBandwidthInBytes returns a boolean if a field has been set.

### GetBasePackage

`func (o *Plan) GetBasePackage() string`

GetBasePackage returns the BasePackage field if non-nil, zero value otherwise.

### GetBasePackageOk

`func (o *Plan) GetBasePackageOk() (*string, bool)`

GetBasePackageOk returns a tuple with the BasePackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePackage

`func (o *Plan) SetBasePackage(v string)`

SetBasePackage sets BasePackage field to given value.

### HasBasePackage

`func (o *Plan) HasBasePackage() bool`

HasBasePackage returns a boolean if a field has been set.

### GetCanary

`func (o *Plan) GetCanary() bool`

GetCanary returns the Canary field if non-nil, zero value otherwise.

### GetCanaryOk

`func (o *Plan) GetCanaryOk() (*bool, bool)`

GetCanaryOk returns a tuple with the Canary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanary

`func (o *Plan) SetCanary(v bool)`

SetCanary sets Canary field to given value.

### HasCanary

`func (o *Plan) HasCanary() bool`

HasCanary returns a boolean if a field has been set.

### GetCpuCount

`func (o *Plan) GetCpuCount() int64`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *Plan) GetCpuCountOk() (*int64, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *Plan) SetCpuCount(v int64)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *Plan) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetCpuShare

`func (o *Plan) GetCpuShare() string`

GetCpuShare returns the CpuShare field if non-nil, zero value otherwise.

### GetCpuShareOk

`func (o *Plan) GetCpuShareOk() (*string, bool)`

GetCpuShareOk returns a tuple with the CpuShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuShare

`func (o *Plan) SetCpuShare(v string)`

SetCpuShare sets CpuShare field to given value.

### HasCpuShare

`func (o *Plan) HasCpuShare() bool`

HasCpuShare returns a boolean if a field has been set.

### GetCreateType

`func (o *Plan) GetCreateType() string`

GetCreateType returns the CreateType field if non-nil, zero value otherwise.

### GetCreateTypeOk

`func (o *Plan) GetCreateTypeOk() (*string, bool)`

GetCreateTypeOk returns a tuple with the CreateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateType

`func (o *Plan) SetCreateType(v string)`

SetCreateType sets CreateType field to given value.

### HasCreateType

`func (o *Plan) HasCreateType() bool`

HasCreateType returns a boolean if a field has been set.

### GetDisk

`func (o *Plan) GetDisk() int64`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *Plan) GetDiskOk() (*int64, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *Plan) SetDisk(v int64)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *Plan) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetDiskInBytes

`func (o *Plan) GetDiskInBytes() int64`

GetDiskInBytes returns the DiskInBytes field if non-nil, zero value otherwise.

### GetDiskInBytesOk

`func (o *Plan) GetDiskInBytesOk() (*int64, bool)`

GetDiskInBytesOk returns a tuple with the DiskInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskInBytes

`func (o *Plan) SetDiskInBytes(v int64)`

SetDiskInBytes sets DiskInBytes field to given value.

### HasDiskInBytes

`func (o *Plan) HasDiskInBytes() bool`

HasDiskInBytes returns a boolean if a field has been set.

### GetGeneration

`func (o *Plan) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *Plan) GetGenerationOk() (*string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *Plan) SetGeneration(v string)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *Plan) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetId

`func (o *Plan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Plan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Plan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Plan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIopsMaxHdd

`func (o *Plan) GetIopsMaxHdd() int64`

GetIopsMaxHdd returns the IopsMaxHdd field if non-nil, zero value otherwise.

### GetIopsMaxHddOk

`func (o *Plan) GetIopsMaxHddOk() (*int64, bool)`

GetIopsMaxHddOk returns a tuple with the IopsMaxHdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsMaxHdd

`func (o *Plan) SetIopsMaxHdd(v int64)`

SetIopsMaxHdd sets IopsMaxHdd field to given value.

### HasIopsMaxHdd

`func (o *Plan) HasIopsMaxHdd() bool`

HasIopsMaxHdd returns a boolean if a field has been set.

### GetIopsMaxSsd

`func (o *Plan) GetIopsMaxSsd() int64`

GetIopsMaxSsd returns the IopsMaxSsd field if non-nil, zero value otherwise.

### GetIopsMaxSsdOk

`func (o *Plan) GetIopsMaxSsdOk() (*int64, bool)`

GetIopsMaxSsdOk returns a tuple with the IopsMaxSsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsMaxSsd

`func (o *Plan) SetIopsMaxSsd(v int64)`

SetIopsMaxSsd sets IopsMaxSsd field to given value.

### HasIopsMaxSsd

`func (o *Plan) HasIopsMaxSsd() bool`

HasIopsMaxSsd returns a boolean if a field has been set.

### GetMemory

`func (o *Plan) GetMemory() float64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Plan) GetMemoryOk() (*float64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Plan) SetMemory(v float64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *Plan) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMemoryInBytes

`func (o *Plan) GetMemoryInBytes() int64`

GetMemoryInBytes returns the MemoryInBytes field if non-nil, zero value otherwise.

### GetMemoryInBytesOk

`func (o *Plan) GetMemoryInBytesOk() (*int64, bool)`

GetMemoryInBytesOk returns a tuple with the MemoryInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryInBytes

`func (o *Plan) SetMemoryInBytes(v int64)`

SetMemoryInBytes sets MemoryInBytes field to given value.

### HasMemoryInBytes

`func (o *Plan) HasMemoryInBytes() bool`

HasMemoryInBytes returns a boolean if a field has been set.

### GetName

`func (o *Plan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Plan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Plan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Plan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOff

`func (o *Plan) GetOff() string`

GetOff returns the Off field if non-nil, zero value otherwise.

### GetOffOk

`func (o *Plan) GetOffOk() (*string, bool)`

GetOffOk returns a tuple with the Off field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOff

`func (o *Plan) SetOff(v string)`

SetOff sets Off field to given value.

### HasOff

`func (o *Plan) HasOff() bool`

HasOff returns a boolean if a field has been set.

### GetOffPercent

`func (o *Plan) GetOffPercent() string`

GetOffPercent returns the OffPercent field if non-nil, zero value otherwise.

### GetOffPercentOk

`func (o *Plan) GetOffPercentOk() (*string, bool)`

GetOffPercentOk returns a tuple with the OffPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPercent

`func (o *Plan) SetOffPercent(v string)`

SetOffPercent sets OffPercent field to given value.

### HasOffPercent

`func (o *Plan) HasOffPercent() bool`

HasOffPercent returns a boolean if a field has been set.

### GetOrder

`func (o *Plan) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Plan) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Plan) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Plan) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOutbound

`func (o *Plan) GetOutbound() int64`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *Plan) GetOutboundOk() (*int64, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *Plan) SetOutbound(v int64)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *Plan) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.

### GetPortSpeed

`func (o *Plan) GetPortSpeed() float64`

GetPortSpeed returns the PortSpeed field if non-nil, zero value otherwise.

### GetPortSpeedOk

`func (o *Plan) GetPortSpeedOk() (*float64, bool)`

GetPortSpeedOk returns a tuple with the PortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSpeed

`func (o *Plan) SetPortSpeed(v float64)`

SetPortSpeed sets PortSpeed field to given value.

### HasPortSpeed

`func (o *Plan) HasPortSpeed() bool`

HasPortSpeed returns a boolean if a field has been set.

### GetPps

`func (o *Plan) GetPps() []int64`

GetPps returns the Pps field if non-nil, zero value otherwise.

### GetPpsOk

`func (o *Plan) GetPpsOk() (*[]int64, bool)`

GetPpsOk returns a tuple with the Pps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPps

`func (o *Plan) SetPps(v []int64)`

SetPps sets Pps field to given value.

### HasPps

`func (o *Plan) HasPps() bool`

HasPps returns a boolean if a field has been set.

### GetPricePerHour

`func (o *Plan) GetPricePerHour() int64`

GetPricePerHour returns the PricePerHour field if non-nil, zero value otherwise.

### GetPricePerHourOk

`func (o *Plan) GetPricePerHourOk() (*int64, bool)`

GetPricePerHourOk returns a tuple with the PricePerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerHour

`func (o *Plan) SetPricePerHour(v int64)`

SetPricePerHour sets PricePerHour field to given value.

### HasPricePerHour

`func (o *Plan) HasPricePerHour() bool`

HasPricePerHour returns a boolean if a field has been set.

### GetPricePerMonth

`func (o *Plan) GetPricePerMonth() int64`

GetPricePerMonth returns the PricePerMonth field if non-nil, zero value otherwise.

### GetPricePerMonthOk

`func (o *Plan) GetPricePerMonthOk() (*int64, bool)`

GetPricePerMonthOk returns a tuple with the PricePerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerMonth

`func (o *Plan) SetPricePerMonth(v int64)`

SetPricePerMonth sets PricePerMonth field to given value.

### HasPricePerMonth

`func (o *Plan) HasPricePerMonth() bool`

HasPricePerMonth returns a boolean if a field has been set.

### GetSubtype

`func (o *Plan) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *Plan) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *Plan) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *Plan) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetThroughput

`func (o *Plan) GetThroughput() int64`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *Plan) GetThroughputOk() (*int64, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *Plan) SetThroughput(v int64)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *Plan) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.

### GetType

`func (o *Plan) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Plan) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Plan) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Plan) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


