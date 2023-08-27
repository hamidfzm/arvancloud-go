# LogForwarderKafkaConnectionType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SampleRate** | Pointer to **int32** |  | [optional] 
**KafkaVersion** | Pointer to **string** |  | [optional] 
**KafkaBrokers** | Pointer to **[]string** |  | [optional] 
**KafkaTopicToWrite** | Pointer to **string** |  | [optional] 
**KafkaProducerBatchSize** | Pointer to **int32** |  | [optional] 
**KafkaProducerFlushFrequencyMs** | Pointer to **int32** |  | [optional] 

## Methods

### NewLogForwarderKafkaConnectionType

`func NewLogForwarderKafkaConnectionType() *LogForwarderKafkaConnectionType`

NewLogForwarderKafkaConnectionType instantiates a new LogForwarderKafkaConnectionType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogForwarderKafkaConnectionTypeWithDefaults

`func NewLogForwarderKafkaConnectionTypeWithDefaults() *LogForwarderKafkaConnectionType`

NewLogForwarderKafkaConnectionTypeWithDefaults instantiates a new LogForwarderKafkaConnectionType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSampleRate

`func (o *LogForwarderKafkaConnectionType) GetSampleRate() int32`

GetSampleRate returns the SampleRate field if non-nil, zero value otherwise.

### GetSampleRateOk

`func (o *LogForwarderKafkaConnectionType) GetSampleRateOk() (*int32, bool)`

GetSampleRateOk returns a tuple with the SampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleRate

`func (o *LogForwarderKafkaConnectionType) SetSampleRate(v int32)`

SetSampleRate sets SampleRate field to given value.

### HasSampleRate

`func (o *LogForwarderKafkaConnectionType) HasSampleRate() bool`

HasSampleRate returns a boolean if a field has been set.

### GetKafkaVersion

`func (o *LogForwarderKafkaConnectionType) GetKafkaVersion() string`

GetKafkaVersion returns the KafkaVersion field if non-nil, zero value otherwise.

### GetKafkaVersionOk

`func (o *LogForwarderKafkaConnectionType) GetKafkaVersionOk() (*string, bool)`

GetKafkaVersionOk returns a tuple with the KafkaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaVersion

`func (o *LogForwarderKafkaConnectionType) SetKafkaVersion(v string)`

SetKafkaVersion sets KafkaVersion field to given value.

### HasKafkaVersion

`func (o *LogForwarderKafkaConnectionType) HasKafkaVersion() bool`

HasKafkaVersion returns a boolean if a field has been set.

### GetKafkaBrokers

`func (o *LogForwarderKafkaConnectionType) GetKafkaBrokers() []string`

GetKafkaBrokers returns the KafkaBrokers field if non-nil, zero value otherwise.

### GetKafkaBrokersOk

`func (o *LogForwarderKafkaConnectionType) GetKafkaBrokersOk() (*[]string, bool)`

GetKafkaBrokersOk returns a tuple with the KafkaBrokers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaBrokers

`func (o *LogForwarderKafkaConnectionType) SetKafkaBrokers(v []string)`

SetKafkaBrokers sets KafkaBrokers field to given value.

### HasKafkaBrokers

`func (o *LogForwarderKafkaConnectionType) HasKafkaBrokers() bool`

HasKafkaBrokers returns a boolean if a field has been set.

### GetKafkaTopicToWrite

`func (o *LogForwarderKafkaConnectionType) GetKafkaTopicToWrite() string`

GetKafkaTopicToWrite returns the KafkaTopicToWrite field if non-nil, zero value otherwise.

### GetKafkaTopicToWriteOk

`func (o *LogForwarderKafkaConnectionType) GetKafkaTopicToWriteOk() (*string, bool)`

GetKafkaTopicToWriteOk returns a tuple with the KafkaTopicToWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaTopicToWrite

`func (o *LogForwarderKafkaConnectionType) SetKafkaTopicToWrite(v string)`

SetKafkaTopicToWrite sets KafkaTopicToWrite field to given value.

### HasKafkaTopicToWrite

`func (o *LogForwarderKafkaConnectionType) HasKafkaTopicToWrite() bool`

HasKafkaTopicToWrite returns a boolean if a field has been set.

### GetKafkaProducerBatchSize

`func (o *LogForwarderKafkaConnectionType) GetKafkaProducerBatchSize() int32`

GetKafkaProducerBatchSize returns the KafkaProducerBatchSize field if non-nil, zero value otherwise.

### GetKafkaProducerBatchSizeOk

`func (o *LogForwarderKafkaConnectionType) GetKafkaProducerBatchSizeOk() (*int32, bool)`

GetKafkaProducerBatchSizeOk returns a tuple with the KafkaProducerBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaProducerBatchSize

`func (o *LogForwarderKafkaConnectionType) SetKafkaProducerBatchSize(v int32)`

SetKafkaProducerBatchSize sets KafkaProducerBatchSize field to given value.

### HasKafkaProducerBatchSize

`func (o *LogForwarderKafkaConnectionType) HasKafkaProducerBatchSize() bool`

HasKafkaProducerBatchSize returns a boolean if a field has been set.

### GetKafkaProducerFlushFrequencyMs

`func (o *LogForwarderKafkaConnectionType) GetKafkaProducerFlushFrequencyMs() int32`

GetKafkaProducerFlushFrequencyMs returns the KafkaProducerFlushFrequencyMs field if non-nil, zero value otherwise.

### GetKafkaProducerFlushFrequencyMsOk

`func (o *LogForwarderKafkaConnectionType) GetKafkaProducerFlushFrequencyMsOk() (*int32, bool)`

GetKafkaProducerFlushFrequencyMsOk returns a tuple with the KafkaProducerFlushFrequencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaProducerFlushFrequencyMs

`func (o *LogForwarderKafkaConnectionType) SetKafkaProducerFlushFrequencyMs(v int32)`

SetKafkaProducerFlushFrequencyMs sets KafkaProducerFlushFrequencyMs field to given value.

### HasKafkaProducerFlushFrequencyMs

`func (o *LogForwarderKafkaConnectionType) HasKafkaProducerFlushFrequencyMs() bool`

HasKafkaProducerFlushFrequencyMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


