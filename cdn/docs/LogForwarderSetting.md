# LogForwarderSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SampleRate** | Pointer to **int32** |  | [optional] 
**S3Endpoint** | Pointer to **string** |  | [optional] 
**AccessKey** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**BucketName** | Pointer to **string** |  | [optional] 
**ObjectSize** | Pointer to **int32** |  | [optional] 
**FlushInterval** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 
**AppKey** | Pointer to **string** |  | [optional] 
**BufferSize** | Pointer to **int32** |  | [optional] 
**KafkaVersion** | Pointer to **string** |  | [optional] 
**KafkaBrokers** | Pointer to **[]string** |  | [optional] 
**KafkaTopicToWrite** | Pointer to **string** |  | [optional] 
**KafkaProducerBatchSize** | Pointer to **int32** |  | [optional] 
**KafkaProducerFlushFrequencyMs** | Pointer to **int32** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Logtype** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Tls** | Pointer to **bool** |  | [optional] 
**Cert** | Pointer to **string** |  | [optional] 
**RetryTime** | Pointer to **int32** |  | [optional] 

## Methods

### NewLogForwarderSetting

`func NewLogForwarderSetting() *LogForwarderSetting`

NewLogForwarderSetting instantiates a new LogForwarderSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogForwarderSettingWithDefaults

`func NewLogForwarderSettingWithDefaults() *LogForwarderSetting`

NewLogForwarderSettingWithDefaults instantiates a new LogForwarderSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSampleRate

`func (o *LogForwarderSetting) GetSampleRate() int32`

GetSampleRate returns the SampleRate field if non-nil, zero value otherwise.

### GetSampleRateOk

`func (o *LogForwarderSetting) GetSampleRateOk() (*int32, bool)`

GetSampleRateOk returns a tuple with the SampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleRate

`func (o *LogForwarderSetting) SetSampleRate(v int32)`

SetSampleRate sets SampleRate field to given value.

### HasSampleRate

`func (o *LogForwarderSetting) HasSampleRate() bool`

HasSampleRate returns a boolean if a field has been set.

### GetS3Endpoint

`func (o *LogForwarderSetting) GetS3Endpoint() string`

GetS3Endpoint returns the S3Endpoint field if non-nil, zero value otherwise.

### GetS3EndpointOk

`func (o *LogForwarderSetting) GetS3EndpointOk() (*string, bool)`

GetS3EndpointOk returns a tuple with the S3Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Endpoint

`func (o *LogForwarderSetting) SetS3Endpoint(v string)`

SetS3Endpoint sets S3Endpoint field to given value.

### HasS3Endpoint

`func (o *LogForwarderSetting) HasS3Endpoint() bool`

HasS3Endpoint returns a boolean if a field has been set.

### GetAccessKey

`func (o *LogForwarderSetting) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *LogForwarderSetting) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *LogForwarderSetting) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *LogForwarderSetting) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *LogForwarderSetting) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *LogForwarderSetting) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *LogForwarderSetting) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *LogForwarderSetting) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetBucketName

`func (o *LogForwarderSetting) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *LogForwarderSetting) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *LogForwarderSetting) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *LogForwarderSetting) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetObjectSize

`func (o *LogForwarderSetting) GetObjectSize() int32`

GetObjectSize returns the ObjectSize field if non-nil, zero value otherwise.

### GetObjectSizeOk

`func (o *LogForwarderSetting) GetObjectSizeOk() (*int32, bool)`

GetObjectSizeOk returns a tuple with the ObjectSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSize

`func (o *LogForwarderSetting) SetObjectSize(v int32)`

SetObjectSize sets ObjectSize field to given value.

### HasObjectSize

`func (o *LogForwarderSetting) HasObjectSize() bool`

HasObjectSize returns a boolean if a field has been set.

### GetFlushInterval

`func (o *LogForwarderSetting) GetFlushInterval() int32`

GetFlushInterval returns the FlushInterval field if non-nil, zero value otherwise.

### GetFlushIntervalOk

`func (o *LogForwarderSetting) GetFlushIntervalOk() (*int32, bool)`

GetFlushIntervalOk returns a tuple with the FlushInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushInterval

`func (o *LogForwarderSetting) SetFlushInterval(v int32)`

SetFlushInterval sets FlushInterval field to given value.

### HasFlushInterval

`func (o *LogForwarderSetting) HasFlushInterval() bool`

HasFlushInterval returns a boolean if a field has been set.

### GetUrl

`func (o *LogForwarderSetting) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LogForwarderSetting) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LogForwarderSetting) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LogForwarderSetting) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetApiKey

`func (o *LogForwarderSetting) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *LogForwarderSetting) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *LogForwarderSetting) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *LogForwarderSetting) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetAppKey

`func (o *LogForwarderSetting) GetAppKey() string`

GetAppKey returns the AppKey field if non-nil, zero value otherwise.

### GetAppKeyOk

`func (o *LogForwarderSetting) GetAppKeyOk() (*string, bool)`

GetAppKeyOk returns a tuple with the AppKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppKey

`func (o *LogForwarderSetting) SetAppKey(v string)`

SetAppKey sets AppKey field to given value.

### HasAppKey

`func (o *LogForwarderSetting) HasAppKey() bool`

HasAppKey returns a boolean if a field has been set.

### GetBufferSize

`func (o *LogForwarderSetting) GetBufferSize() int32`

GetBufferSize returns the BufferSize field if non-nil, zero value otherwise.

### GetBufferSizeOk

`func (o *LogForwarderSetting) GetBufferSizeOk() (*int32, bool)`

GetBufferSizeOk returns a tuple with the BufferSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferSize

`func (o *LogForwarderSetting) SetBufferSize(v int32)`

SetBufferSize sets BufferSize field to given value.

### HasBufferSize

`func (o *LogForwarderSetting) HasBufferSize() bool`

HasBufferSize returns a boolean if a field has been set.

### GetKafkaVersion

`func (o *LogForwarderSetting) GetKafkaVersion() string`

GetKafkaVersion returns the KafkaVersion field if non-nil, zero value otherwise.

### GetKafkaVersionOk

`func (o *LogForwarderSetting) GetKafkaVersionOk() (*string, bool)`

GetKafkaVersionOk returns a tuple with the KafkaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaVersion

`func (o *LogForwarderSetting) SetKafkaVersion(v string)`

SetKafkaVersion sets KafkaVersion field to given value.

### HasKafkaVersion

`func (o *LogForwarderSetting) HasKafkaVersion() bool`

HasKafkaVersion returns a boolean if a field has been set.

### GetKafkaBrokers

`func (o *LogForwarderSetting) GetKafkaBrokers() []string`

GetKafkaBrokers returns the KafkaBrokers field if non-nil, zero value otherwise.

### GetKafkaBrokersOk

`func (o *LogForwarderSetting) GetKafkaBrokersOk() (*[]string, bool)`

GetKafkaBrokersOk returns a tuple with the KafkaBrokers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaBrokers

`func (o *LogForwarderSetting) SetKafkaBrokers(v []string)`

SetKafkaBrokers sets KafkaBrokers field to given value.

### HasKafkaBrokers

`func (o *LogForwarderSetting) HasKafkaBrokers() bool`

HasKafkaBrokers returns a boolean if a field has been set.

### GetKafkaTopicToWrite

`func (o *LogForwarderSetting) GetKafkaTopicToWrite() string`

GetKafkaTopicToWrite returns the KafkaTopicToWrite field if non-nil, zero value otherwise.

### GetKafkaTopicToWriteOk

`func (o *LogForwarderSetting) GetKafkaTopicToWriteOk() (*string, bool)`

GetKafkaTopicToWriteOk returns a tuple with the KafkaTopicToWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaTopicToWrite

`func (o *LogForwarderSetting) SetKafkaTopicToWrite(v string)`

SetKafkaTopicToWrite sets KafkaTopicToWrite field to given value.

### HasKafkaTopicToWrite

`func (o *LogForwarderSetting) HasKafkaTopicToWrite() bool`

HasKafkaTopicToWrite returns a boolean if a field has been set.

### GetKafkaProducerBatchSize

`func (o *LogForwarderSetting) GetKafkaProducerBatchSize() int32`

GetKafkaProducerBatchSize returns the KafkaProducerBatchSize field if non-nil, zero value otherwise.

### GetKafkaProducerBatchSizeOk

`func (o *LogForwarderSetting) GetKafkaProducerBatchSizeOk() (*int32, bool)`

GetKafkaProducerBatchSizeOk returns a tuple with the KafkaProducerBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaProducerBatchSize

`func (o *LogForwarderSetting) SetKafkaProducerBatchSize(v int32)`

SetKafkaProducerBatchSize sets KafkaProducerBatchSize field to given value.

### HasKafkaProducerBatchSize

`func (o *LogForwarderSetting) HasKafkaProducerBatchSize() bool`

HasKafkaProducerBatchSize returns a boolean if a field has been set.

### GetKafkaProducerFlushFrequencyMs

`func (o *LogForwarderSetting) GetKafkaProducerFlushFrequencyMs() int32`

GetKafkaProducerFlushFrequencyMs returns the KafkaProducerFlushFrequencyMs field if non-nil, zero value otherwise.

### GetKafkaProducerFlushFrequencyMsOk

`func (o *LogForwarderSetting) GetKafkaProducerFlushFrequencyMsOk() (*int32, bool)`

GetKafkaProducerFlushFrequencyMsOk returns a tuple with the KafkaProducerFlushFrequencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaProducerFlushFrequencyMs

`func (o *LogForwarderSetting) SetKafkaProducerFlushFrequencyMs(v int32)`

SetKafkaProducerFlushFrequencyMs sets KafkaProducerFlushFrequencyMs field to given value.

### HasKafkaProducerFlushFrequencyMs

`func (o *LogForwarderSetting) HasKafkaProducerFlushFrequencyMs() bool`

HasKafkaProducerFlushFrequencyMs returns a boolean if a field has been set.

### GetToken

`func (o *LogForwarderSetting) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LogForwarderSetting) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LogForwarderSetting) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LogForwarderSetting) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetLogtype

`func (o *LogForwarderSetting) GetLogtype() string`

GetLogtype returns the Logtype field if non-nil, zero value otherwise.

### GetLogtypeOk

`func (o *LogForwarderSetting) GetLogtypeOk() (*string, bool)`

GetLogtypeOk returns a tuple with the Logtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogtype

`func (o *LogForwarderSetting) SetLogtype(v string)`

SetLogtype sets Logtype field to given value.

### HasLogtype

`func (o *LogForwarderSetting) HasLogtype() bool`

HasLogtype returns a boolean if a field has been set.

### GetHost

`func (o *LogForwarderSetting) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *LogForwarderSetting) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *LogForwarderSetting) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *LogForwarderSetting) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *LogForwarderSetting) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LogForwarderSetting) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LogForwarderSetting) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *LogForwarderSetting) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetTls

`func (o *LogForwarderSetting) GetTls() bool`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *LogForwarderSetting) GetTlsOk() (*bool, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *LogForwarderSetting) SetTls(v bool)`

SetTls sets Tls field to given value.

### HasTls

`func (o *LogForwarderSetting) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetCert

`func (o *LogForwarderSetting) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *LogForwarderSetting) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *LogForwarderSetting) SetCert(v string)`

SetCert sets Cert field to given value.

### HasCert

`func (o *LogForwarderSetting) HasCert() bool`

HasCert returns a boolean if a field has been set.

### GetRetryTime

`func (o *LogForwarderSetting) GetRetryTime() int32`

GetRetryTime returns the RetryTime field if non-nil, zero value otherwise.

### GetRetryTimeOk

`func (o *LogForwarderSetting) GetRetryTimeOk() (*int32, bool)`

GetRetryTimeOk returns a tuple with the RetryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryTime

`func (o *LogForwarderSetting) SetRetryTime(v int32)`

SetRetryTime sets RetryTime field to given value.

### HasRetryTime

`func (o *LogForwarderSetting) HasRetryTime() bool`

HasRetryTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


