# SinglePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandwidthInBytes** | Pointer to **int64** |  | [optional] 
**BasePackage** | Pointer to **string** |  | [optional] 
**CpuCount** | Pointer to **int64** |  | [optional] 
**CpuShare** | Pointer to **string** |  | [optional] 
**CreateType** | Pointer to **string** |  | [optional] 
**Disk** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IopsMaxHdd** | Pointer to **int64** |  | [optional] 
**IopsMaxSsd** | Pointer to **int64** |  | [optional] 
**Memory** | Pointer to **float64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Off** | Pointer to **string** |  | [optional] 
**OffPercent** | Pointer to **string** |  | [optional] 
**PortSpeed** | Pointer to **float64** |  | [optional] 
**Pps** | Pointer to **[]int64** |  | [optional] 
**PricePerHour** | Pointer to **int64** |  | [optional] 
**PricePerMonth** | Pointer to **int64** |  | [optional] 
**Subtype** | Pointer to **string** |  | [optional] 
**Swap** | Pointer to **string** |  | [optional] 
**Throughput** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewSinglePlan

`func NewSinglePlan() *SinglePlan`

NewSinglePlan instantiates a new SinglePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSinglePlanWithDefaults

`func NewSinglePlanWithDefaults() *SinglePlan`

NewSinglePlanWithDefaults instantiates a new SinglePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidthInBytes

`func (o *SinglePlan) GetBandwidthInBytes() int64`

GetBandwidthInBytes returns the BandwidthInBytes field if non-nil, zero value otherwise.

### GetBandwidthInBytesOk

`func (o *SinglePlan) GetBandwidthInBytesOk() (*int64, bool)`

GetBandwidthInBytesOk returns a tuple with the BandwidthInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthInBytes

`func (o *SinglePlan) SetBandwidthInBytes(v int64)`

SetBandwidthInBytes sets BandwidthInBytes field to given value.

### HasBandwidthInBytes

`func (o *SinglePlan) HasBandwidthInBytes() bool`

HasBandwidthInBytes returns a boolean if a field has been set.

### GetBasePackage

`func (o *SinglePlan) GetBasePackage() string`

GetBasePackage returns the BasePackage field if non-nil, zero value otherwise.

### GetBasePackageOk

`func (o *SinglePlan) GetBasePackageOk() (*string, bool)`

GetBasePackageOk returns a tuple with the BasePackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePackage

`func (o *SinglePlan) SetBasePackage(v string)`

SetBasePackage sets BasePackage field to given value.

### HasBasePackage

`func (o *SinglePlan) HasBasePackage() bool`

HasBasePackage returns a boolean if a field has been set.

### GetCpuCount

`func (o *SinglePlan) GetCpuCount() int64`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *SinglePlan) GetCpuCountOk() (*int64, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *SinglePlan) SetCpuCount(v int64)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *SinglePlan) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetCpuShare

`func (o *SinglePlan) GetCpuShare() string`

GetCpuShare returns the CpuShare field if non-nil, zero value otherwise.

### GetCpuShareOk

`func (o *SinglePlan) GetCpuShareOk() (*string, bool)`

GetCpuShareOk returns a tuple with the CpuShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuShare

`func (o *SinglePlan) SetCpuShare(v string)`

SetCpuShare sets CpuShare field to given value.

### HasCpuShare

`func (o *SinglePlan) HasCpuShare() bool`

HasCpuShare returns a boolean if a field has been set.

### GetCreateType

`func (o *SinglePlan) GetCreateType() string`

GetCreateType returns the CreateType field if non-nil, zero value otherwise.

### GetCreateTypeOk

`func (o *SinglePlan) GetCreateTypeOk() (*string, bool)`

GetCreateTypeOk returns a tuple with the CreateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateType

`func (o *SinglePlan) SetCreateType(v string)`

SetCreateType sets CreateType field to given value.

### HasCreateType

`func (o *SinglePlan) HasCreateType() bool`

HasCreateType returns a boolean if a field has been set.

### GetDisk

`func (o *SinglePlan) GetDisk() int64`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *SinglePlan) GetDiskOk() (*int64, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *SinglePlan) SetDisk(v int64)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *SinglePlan) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetId

`func (o *SinglePlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SinglePlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SinglePlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SinglePlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIopsMaxHdd

`func (o *SinglePlan) GetIopsMaxHdd() int64`

GetIopsMaxHdd returns the IopsMaxHdd field if non-nil, zero value otherwise.

### GetIopsMaxHddOk

`func (o *SinglePlan) GetIopsMaxHddOk() (*int64, bool)`

GetIopsMaxHddOk returns a tuple with the IopsMaxHdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsMaxHdd

`func (o *SinglePlan) SetIopsMaxHdd(v int64)`

SetIopsMaxHdd sets IopsMaxHdd field to given value.

### HasIopsMaxHdd

`func (o *SinglePlan) HasIopsMaxHdd() bool`

HasIopsMaxHdd returns a boolean if a field has been set.

### GetIopsMaxSsd

`func (o *SinglePlan) GetIopsMaxSsd() int64`

GetIopsMaxSsd returns the IopsMaxSsd field if non-nil, zero value otherwise.

### GetIopsMaxSsdOk

`func (o *SinglePlan) GetIopsMaxSsdOk() (*int64, bool)`

GetIopsMaxSsdOk returns a tuple with the IopsMaxSsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsMaxSsd

`func (o *SinglePlan) SetIopsMaxSsd(v int64)`

SetIopsMaxSsd sets IopsMaxSsd field to given value.

### HasIopsMaxSsd

`func (o *SinglePlan) HasIopsMaxSsd() bool`

HasIopsMaxSsd returns a boolean if a field has been set.

### GetMemory

`func (o *SinglePlan) GetMemory() float64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *SinglePlan) GetMemoryOk() (*float64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *SinglePlan) SetMemory(v float64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *SinglePlan) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetName

`func (o *SinglePlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SinglePlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SinglePlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SinglePlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOff

`func (o *SinglePlan) GetOff() string`

GetOff returns the Off field if non-nil, zero value otherwise.

### GetOffOk

`func (o *SinglePlan) GetOffOk() (*string, bool)`

GetOffOk returns a tuple with the Off field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOff

`func (o *SinglePlan) SetOff(v string)`

SetOff sets Off field to given value.

### HasOff

`func (o *SinglePlan) HasOff() bool`

HasOff returns a boolean if a field has been set.

### GetOffPercent

`func (o *SinglePlan) GetOffPercent() string`

GetOffPercent returns the OffPercent field if non-nil, zero value otherwise.

### GetOffPercentOk

`func (o *SinglePlan) GetOffPercentOk() (*string, bool)`

GetOffPercentOk returns a tuple with the OffPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPercent

`func (o *SinglePlan) SetOffPercent(v string)`

SetOffPercent sets OffPercent field to given value.

### HasOffPercent

`func (o *SinglePlan) HasOffPercent() bool`

HasOffPercent returns a boolean if a field has been set.

### GetPortSpeed

`func (o *SinglePlan) GetPortSpeed() float64`

GetPortSpeed returns the PortSpeed field if non-nil, zero value otherwise.

### GetPortSpeedOk

`func (o *SinglePlan) GetPortSpeedOk() (*float64, bool)`

GetPortSpeedOk returns a tuple with the PortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSpeed

`func (o *SinglePlan) SetPortSpeed(v float64)`

SetPortSpeed sets PortSpeed field to given value.

### HasPortSpeed

`func (o *SinglePlan) HasPortSpeed() bool`

HasPortSpeed returns a boolean if a field has been set.

### GetPps

`func (o *SinglePlan) GetPps() []int64`

GetPps returns the Pps field if non-nil, zero value otherwise.

### GetPpsOk

`func (o *SinglePlan) GetPpsOk() (*[]int64, bool)`

GetPpsOk returns a tuple with the Pps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPps

`func (o *SinglePlan) SetPps(v []int64)`

SetPps sets Pps field to given value.

### HasPps

`func (o *SinglePlan) HasPps() bool`

HasPps returns a boolean if a field has been set.

### GetPricePerHour

`func (o *SinglePlan) GetPricePerHour() int64`

GetPricePerHour returns the PricePerHour field if non-nil, zero value otherwise.

### GetPricePerHourOk

`func (o *SinglePlan) GetPricePerHourOk() (*int64, bool)`

GetPricePerHourOk returns a tuple with the PricePerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerHour

`func (o *SinglePlan) SetPricePerHour(v int64)`

SetPricePerHour sets PricePerHour field to given value.

### HasPricePerHour

`func (o *SinglePlan) HasPricePerHour() bool`

HasPricePerHour returns a boolean if a field has been set.

### GetPricePerMonth

`func (o *SinglePlan) GetPricePerMonth() int64`

GetPricePerMonth returns the PricePerMonth field if non-nil, zero value otherwise.

### GetPricePerMonthOk

`func (o *SinglePlan) GetPricePerMonthOk() (*int64, bool)`

GetPricePerMonthOk returns a tuple with the PricePerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerMonth

`func (o *SinglePlan) SetPricePerMonth(v int64)`

SetPricePerMonth sets PricePerMonth field to given value.

### HasPricePerMonth

`func (o *SinglePlan) HasPricePerMonth() bool`

HasPricePerMonth returns a boolean if a field has been set.

### GetSubtype

`func (o *SinglePlan) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *SinglePlan) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *SinglePlan) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *SinglePlan) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetSwap

`func (o *SinglePlan) GetSwap() string`

GetSwap returns the Swap field if non-nil, zero value otherwise.

### GetSwapOk

`func (o *SinglePlan) GetSwapOk() (*string, bool)`

GetSwapOk returns a tuple with the Swap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwap

`func (o *SinglePlan) SetSwap(v string)`

SetSwap sets Swap field to given value.

### HasSwap

`func (o *SinglePlan) HasSwap() bool`

HasSwap returns a boolean if a field has been set.

### GetThroughput

`func (o *SinglePlan) GetThroughput() int64`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *SinglePlan) GetThroughputOk() (*int64, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *SinglePlan) SetThroughput(v int64)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *SinglePlan) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.

### GetType

`func (o *SinglePlan) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SinglePlan) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SinglePlan) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SinglePlan) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


