# LogForwarderDataFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **bool** |  | [optional] 
**Scheme** | Pointer to **bool** |  | [optional] 
**Domain** | Pointer to **bool** |  | [optional] 
**Uri** | Pointer to **bool** |  | [optional] 
**QueryString** | Pointer to **bool** |  | [optional] 
**Referer** | Pointer to **bool** |  | [optional] 
**Ip** | Pointer to **bool** |  | [optional] 
**Ua** | Pointer to **bool** |  | [optional] 
**Country** | Pointer to **bool** |  | [optional] 
**Asn** | Pointer to **bool** |  | [optional] 
**ContentType** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**TlsFingerprint** | Pointer to **bool** |  | [optional] 
**ServerPort** | Pointer to **bool** |  | [optional] 
**BytesSent** | Pointer to **bool** |  | [optional] 
**BytesReceived** | Pointer to **bool** |  | [optional] 
**UpstreamTime** | Pointer to **bool** |  | [optional] 
**Cache** | Pointer to **bool** |  | [optional] 
**RequestId** | Pointer to **bool** |  | [optional] 
**Timestamp** | Pointer to **bool** |  | [optional] 
**Product** | Pointer to **bool** |  | [optional] 
**RemoteAddress** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to **bool** |  | [optional] 
**Uuid** | Pointer to **bool** |  | [optional] 
**Record** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **bool** |  | [optional] 
**ResponseCode** | Pointer to **bool** |  | [optional] 
**ProcessTime** | Pointer to **bool** |  | [optional] 
**ClientIp** | Pointer to **bool** |  | [optional] 
**UpstreamProto** | Pointer to **bool** |  | [optional] 
**UpstreamUri** | Pointer to **bool** |  | [optional] 
**UpstreamPort** | Pointer to **bool** |  | [optional] 
**UpstreamIp** | Pointer to **bool** |  | [optional] 
**DomainName** | Pointer to **bool** |  | [optional] 
**HttpVersion** | Pointer to **bool** |  | [optional] 
**RequestMethod** | Pointer to **bool** |  | [optional] 
**RequestUri** | Pointer to **bool** |  | [optional] 
**RealTimestamp** | Pointer to **bool** |  | [optional] 
**ErrorMessage** | Pointer to **bool** |  | [optional] 
**PopSite** | Pointer to **bool** |  | [optional] 
**ServerIp** | Pointer to **bool** |  | [optional] 
**Firewall** | Pointer to **bool** |  | [optional] 
**Proxy** | Pointer to **bool** |  | [optional] 
**DnsResolver** | Pointer to **bool** |  | [optional] 
**Ddos** | Pointer to **bool** |  | [optional] 
**Ratelimit** | Pointer to **bool** |  | [optional] 
**Waf** | Pointer to **bool** |  | [optional] 

## Methods

### NewLogForwarderDataFormat

`func NewLogForwarderDataFormat() *LogForwarderDataFormat`

NewLogForwarderDataFormat instantiates a new LogForwarderDataFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogForwarderDataFormatWithDefaults

`func NewLogForwarderDataFormatWithDefaults() *LogForwarderDataFormat`

NewLogForwarderDataFormatWithDefaults instantiates a new LogForwarderDataFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *LogForwarderDataFormat) GetMethod() bool`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LogForwarderDataFormat) GetMethodOk() (*bool, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LogForwarderDataFormat) SetMethod(v bool)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LogForwarderDataFormat) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetScheme

`func (o *LogForwarderDataFormat) GetScheme() bool`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *LogForwarderDataFormat) GetSchemeOk() (*bool, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *LogForwarderDataFormat) SetScheme(v bool)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *LogForwarderDataFormat) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetDomain

`func (o *LogForwarderDataFormat) GetDomain() bool`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *LogForwarderDataFormat) GetDomainOk() (*bool, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *LogForwarderDataFormat) SetDomain(v bool)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *LogForwarderDataFormat) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetUri

`func (o *LogForwarderDataFormat) GetUri() bool`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *LogForwarderDataFormat) GetUriOk() (*bool, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *LogForwarderDataFormat) SetUri(v bool)`

SetUri sets Uri field to given value.

### HasUri

`func (o *LogForwarderDataFormat) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetQueryString

`func (o *LogForwarderDataFormat) GetQueryString() bool`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *LogForwarderDataFormat) GetQueryStringOk() (*bool, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *LogForwarderDataFormat) SetQueryString(v bool)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *LogForwarderDataFormat) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### GetReferer

`func (o *LogForwarderDataFormat) GetReferer() bool`

GetReferer returns the Referer field if non-nil, zero value otherwise.

### GetRefererOk

`func (o *LogForwarderDataFormat) GetRefererOk() (*bool, bool)`

GetRefererOk returns a tuple with the Referer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferer

`func (o *LogForwarderDataFormat) SetReferer(v bool)`

SetReferer sets Referer field to given value.

### HasReferer

`func (o *LogForwarderDataFormat) HasReferer() bool`

HasReferer returns a boolean if a field has been set.

### GetIp

`func (o *LogForwarderDataFormat) GetIp() bool`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *LogForwarderDataFormat) GetIpOk() (*bool, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *LogForwarderDataFormat) SetIp(v bool)`

SetIp sets Ip field to given value.

### HasIp

`func (o *LogForwarderDataFormat) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetUa

`func (o *LogForwarderDataFormat) GetUa() bool`

GetUa returns the Ua field if non-nil, zero value otherwise.

### GetUaOk

`func (o *LogForwarderDataFormat) GetUaOk() (*bool, bool)`

GetUaOk returns a tuple with the Ua field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUa

`func (o *LogForwarderDataFormat) SetUa(v bool)`

SetUa sets Ua field to given value.

### HasUa

`func (o *LogForwarderDataFormat) HasUa() bool`

HasUa returns a boolean if a field has been set.

### GetCountry

`func (o *LogForwarderDataFormat) GetCountry() bool`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *LogForwarderDataFormat) GetCountryOk() (*bool, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *LogForwarderDataFormat) SetCountry(v bool)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *LogForwarderDataFormat) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetAsn

`func (o *LogForwarderDataFormat) GetAsn() bool`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *LogForwarderDataFormat) GetAsnOk() (*bool, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *LogForwarderDataFormat) SetAsn(v bool)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *LogForwarderDataFormat) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetContentType

`func (o *LogForwarderDataFormat) GetContentType() bool`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *LogForwarderDataFormat) GetContentTypeOk() (*bool, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *LogForwarderDataFormat) SetContentType(v bool)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *LogForwarderDataFormat) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetStatus

`func (o *LogForwarderDataFormat) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LogForwarderDataFormat) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LogForwarderDataFormat) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LogForwarderDataFormat) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTlsFingerprint

`func (o *LogForwarderDataFormat) GetTlsFingerprint() bool`

GetTlsFingerprint returns the TlsFingerprint field if non-nil, zero value otherwise.

### GetTlsFingerprintOk

`func (o *LogForwarderDataFormat) GetTlsFingerprintOk() (*bool, bool)`

GetTlsFingerprintOk returns a tuple with the TlsFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsFingerprint

`func (o *LogForwarderDataFormat) SetTlsFingerprint(v bool)`

SetTlsFingerprint sets TlsFingerprint field to given value.

### HasTlsFingerprint

`func (o *LogForwarderDataFormat) HasTlsFingerprint() bool`

HasTlsFingerprint returns a boolean if a field has been set.

### GetServerPort

`func (o *LogForwarderDataFormat) GetServerPort() bool`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *LogForwarderDataFormat) GetServerPortOk() (*bool, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *LogForwarderDataFormat) SetServerPort(v bool)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *LogForwarderDataFormat) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetBytesSent

`func (o *LogForwarderDataFormat) GetBytesSent() bool`

GetBytesSent returns the BytesSent field if non-nil, zero value otherwise.

### GetBytesSentOk

`func (o *LogForwarderDataFormat) GetBytesSentOk() (*bool, bool)`

GetBytesSentOk returns a tuple with the BytesSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesSent

`func (o *LogForwarderDataFormat) SetBytesSent(v bool)`

SetBytesSent sets BytesSent field to given value.

### HasBytesSent

`func (o *LogForwarderDataFormat) HasBytesSent() bool`

HasBytesSent returns a boolean if a field has been set.

### GetBytesReceived

`func (o *LogForwarderDataFormat) GetBytesReceived() bool`

GetBytesReceived returns the BytesReceived field if non-nil, zero value otherwise.

### GetBytesReceivedOk

`func (o *LogForwarderDataFormat) GetBytesReceivedOk() (*bool, bool)`

GetBytesReceivedOk returns a tuple with the BytesReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesReceived

`func (o *LogForwarderDataFormat) SetBytesReceived(v bool)`

SetBytesReceived sets BytesReceived field to given value.

### HasBytesReceived

`func (o *LogForwarderDataFormat) HasBytesReceived() bool`

HasBytesReceived returns a boolean if a field has been set.

### GetUpstreamTime

`func (o *LogForwarderDataFormat) GetUpstreamTime() bool`

GetUpstreamTime returns the UpstreamTime field if non-nil, zero value otherwise.

### GetUpstreamTimeOk

`func (o *LogForwarderDataFormat) GetUpstreamTimeOk() (*bool, bool)`

GetUpstreamTimeOk returns a tuple with the UpstreamTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamTime

`func (o *LogForwarderDataFormat) SetUpstreamTime(v bool)`

SetUpstreamTime sets UpstreamTime field to given value.

### HasUpstreamTime

`func (o *LogForwarderDataFormat) HasUpstreamTime() bool`

HasUpstreamTime returns a boolean if a field has been set.

### GetCache

`func (o *LogForwarderDataFormat) GetCache() bool`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *LogForwarderDataFormat) GetCacheOk() (*bool, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *LogForwarderDataFormat) SetCache(v bool)`

SetCache sets Cache field to given value.

### HasCache

`func (o *LogForwarderDataFormat) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetRequestId

`func (o *LogForwarderDataFormat) GetRequestId() bool`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *LogForwarderDataFormat) GetRequestIdOk() (*bool, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *LogForwarderDataFormat) SetRequestId(v bool)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *LogForwarderDataFormat) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetTimestamp

`func (o *LogForwarderDataFormat) GetTimestamp() bool`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LogForwarderDataFormat) GetTimestampOk() (*bool, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LogForwarderDataFormat) SetTimestamp(v bool)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LogForwarderDataFormat) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetProduct

`func (o *LogForwarderDataFormat) GetProduct() bool`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *LogForwarderDataFormat) GetProductOk() (*bool, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *LogForwarderDataFormat) SetProduct(v bool)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *LogForwarderDataFormat) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *LogForwarderDataFormat) GetRemoteAddress() bool`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *LogForwarderDataFormat) GetRemoteAddressOk() (*bool, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *LogForwarderDataFormat) SetRemoteAddress(v bool)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *LogForwarderDataFormat) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetData

`func (o *LogForwarderDataFormat) GetData() bool`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LogForwarderDataFormat) GetDataOk() (*bool, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LogForwarderDataFormat) SetData(v bool)`

SetData sets Data field to given value.

### HasData

`func (o *LogForwarderDataFormat) HasData() bool`

HasData returns a boolean if a field has been set.

### GetUuid

`func (o *LogForwarderDataFormat) GetUuid() bool`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *LogForwarderDataFormat) GetUuidOk() (*bool, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *LogForwarderDataFormat) SetUuid(v bool)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *LogForwarderDataFormat) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetRecord

`func (o *LogForwarderDataFormat) GetRecord() bool`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *LogForwarderDataFormat) GetRecordOk() (*bool, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *LogForwarderDataFormat) SetRecord(v bool)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *LogForwarderDataFormat) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetType

`func (o *LogForwarderDataFormat) GetType() bool`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogForwarderDataFormat) GetTypeOk() (*bool, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogForwarderDataFormat) SetType(v bool)`

SetType sets Type field to given value.

### HasType

`func (o *LogForwarderDataFormat) HasType() bool`

HasType returns a boolean if a field has been set.

### GetResponseCode

`func (o *LogForwarderDataFormat) GetResponseCode() bool`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *LogForwarderDataFormat) GetResponseCodeOk() (*bool, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *LogForwarderDataFormat) SetResponseCode(v bool)`

SetResponseCode sets ResponseCode field to given value.

### HasResponseCode

`func (o *LogForwarderDataFormat) HasResponseCode() bool`

HasResponseCode returns a boolean if a field has been set.

### GetProcessTime

`func (o *LogForwarderDataFormat) GetProcessTime() bool`

GetProcessTime returns the ProcessTime field if non-nil, zero value otherwise.

### GetProcessTimeOk

`func (o *LogForwarderDataFormat) GetProcessTimeOk() (*bool, bool)`

GetProcessTimeOk returns a tuple with the ProcessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessTime

`func (o *LogForwarderDataFormat) SetProcessTime(v bool)`

SetProcessTime sets ProcessTime field to given value.

### HasProcessTime

`func (o *LogForwarderDataFormat) HasProcessTime() bool`

HasProcessTime returns a boolean if a field has been set.

### GetClientIp

`func (o *LogForwarderDataFormat) GetClientIp() bool`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *LogForwarderDataFormat) GetClientIpOk() (*bool, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *LogForwarderDataFormat) SetClientIp(v bool)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *LogForwarderDataFormat) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetUpstreamProto

`func (o *LogForwarderDataFormat) GetUpstreamProto() bool`

GetUpstreamProto returns the UpstreamProto field if non-nil, zero value otherwise.

### GetUpstreamProtoOk

`func (o *LogForwarderDataFormat) GetUpstreamProtoOk() (*bool, bool)`

GetUpstreamProtoOk returns a tuple with the UpstreamProto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamProto

`func (o *LogForwarderDataFormat) SetUpstreamProto(v bool)`

SetUpstreamProto sets UpstreamProto field to given value.

### HasUpstreamProto

`func (o *LogForwarderDataFormat) HasUpstreamProto() bool`

HasUpstreamProto returns a boolean if a field has been set.

### GetUpstreamUri

`func (o *LogForwarderDataFormat) GetUpstreamUri() bool`

GetUpstreamUri returns the UpstreamUri field if non-nil, zero value otherwise.

### GetUpstreamUriOk

`func (o *LogForwarderDataFormat) GetUpstreamUriOk() (*bool, bool)`

GetUpstreamUriOk returns a tuple with the UpstreamUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamUri

`func (o *LogForwarderDataFormat) SetUpstreamUri(v bool)`

SetUpstreamUri sets UpstreamUri field to given value.

### HasUpstreamUri

`func (o *LogForwarderDataFormat) HasUpstreamUri() bool`

HasUpstreamUri returns a boolean if a field has been set.

### GetUpstreamPort

`func (o *LogForwarderDataFormat) GetUpstreamPort() bool`

GetUpstreamPort returns the UpstreamPort field if non-nil, zero value otherwise.

### GetUpstreamPortOk

`func (o *LogForwarderDataFormat) GetUpstreamPortOk() (*bool, bool)`

GetUpstreamPortOk returns a tuple with the UpstreamPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamPort

`func (o *LogForwarderDataFormat) SetUpstreamPort(v bool)`

SetUpstreamPort sets UpstreamPort field to given value.

### HasUpstreamPort

`func (o *LogForwarderDataFormat) HasUpstreamPort() bool`

HasUpstreamPort returns a boolean if a field has been set.

### GetUpstreamIp

`func (o *LogForwarderDataFormat) GetUpstreamIp() bool`

GetUpstreamIp returns the UpstreamIp field if non-nil, zero value otherwise.

### GetUpstreamIpOk

`func (o *LogForwarderDataFormat) GetUpstreamIpOk() (*bool, bool)`

GetUpstreamIpOk returns a tuple with the UpstreamIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamIp

`func (o *LogForwarderDataFormat) SetUpstreamIp(v bool)`

SetUpstreamIp sets UpstreamIp field to given value.

### HasUpstreamIp

`func (o *LogForwarderDataFormat) HasUpstreamIp() bool`

HasUpstreamIp returns a boolean if a field has been set.

### GetDomainName

`func (o *LogForwarderDataFormat) GetDomainName() bool`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *LogForwarderDataFormat) GetDomainNameOk() (*bool, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *LogForwarderDataFormat) SetDomainName(v bool)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *LogForwarderDataFormat) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetHttpVersion

`func (o *LogForwarderDataFormat) GetHttpVersion() bool`

GetHttpVersion returns the HttpVersion field if non-nil, zero value otherwise.

### GetHttpVersionOk

`func (o *LogForwarderDataFormat) GetHttpVersionOk() (*bool, bool)`

GetHttpVersionOk returns a tuple with the HttpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpVersion

`func (o *LogForwarderDataFormat) SetHttpVersion(v bool)`

SetHttpVersion sets HttpVersion field to given value.

### HasHttpVersion

`func (o *LogForwarderDataFormat) HasHttpVersion() bool`

HasHttpVersion returns a boolean if a field has been set.

### GetRequestMethod

`func (o *LogForwarderDataFormat) GetRequestMethod() bool`

GetRequestMethod returns the RequestMethod field if non-nil, zero value otherwise.

### GetRequestMethodOk

`func (o *LogForwarderDataFormat) GetRequestMethodOk() (*bool, bool)`

GetRequestMethodOk returns a tuple with the RequestMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMethod

`func (o *LogForwarderDataFormat) SetRequestMethod(v bool)`

SetRequestMethod sets RequestMethod field to given value.

### HasRequestMethod

`func (o *LogForwarderDataFormat) HasRequestMethod() bool`

HasRequestMethod returns a boolean if a field has been set.

### GetRequestUri

`func (o *LogForwarderDataFormat) GetRequestUri() bool`

GetRequestUri returns the RequestUri field if non-nil, zero value otherwise.

### GetRequestUriOk

`func (o *LogForwarderDataFormat) GetRequestUriOk() (*bool, bool)`

GetRequestUriOk returns a tuple with the RequestUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUri

`func (o *LogForwarderDataFormat) SetRequestUri(v bool)`

SetRequestUri sets RequestUri field to given value.

### HasRequestUri

`func (o *LogForwarderDataFormat) HasRequestUri() bool`

HasRequestUri returns a boolean if a field has been set.

### GetRealTimestamp

`func (o *LogForwarderDataFormat) GetRealTimestamp() bool`

GetRealTimestamp returns the RealTimestamp field if non-nil, zero value otherwise.

### GetRealTimestampOk

`func (o *LogForwarderDataFormat) GetRealTimestampOk() (*bool, bool)`

GetRealTimestampOk returns a tuple with the RealTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealTimestamp

`func (o *LogForwarderDataFormat) SetRealTimestamp(v bool)`

SetRealTimestamp sets RealTimestamp field to given value.

### HasRealTimestamp

`func (o *LogForwarderDataFormat) HasRealTimestamp() bool`

HasRealTimestamp returns a boolean if a field has been set.

### GetErrorMessage

`func (o *LogForwarderDataFormat) GetErrorMessage() bool`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *LogForwarderDataFormat) GetErrorMessageOk() (*bool, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *LogForwarderDataFormat) SetErrorMessage(v bool)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *LogForwarderDataFormat) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetPopSite

`func (o *LogForwarderDataFormat) GetPopSite() bool`

GetPopSite returns the PopSite field if non-nil, zero value otherwise.

### GetPopSiteOk

`func (o *LogForwarderDataFormat) GetPopSiteOk() (*bool, bool)`

GetPopSiteOk returns a tuple with the PopSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopSite

`func (o *LogForwarderDataFormat) SetPopSite(v bool)`

SetPopSite sets PopSite field to given value.

### HasPopSite

`func (o *LogForwarderDataFormat) HasPopSite() bool`

HasPopSite returns a boolean if a field has been set.

### GetServerIp

`func (o *LogForwarderDataFormat) GetServerIp() bool`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *LogForwarderDataFormat) GetServerIpOk() (*bool, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *LogForwarderDataFormat) SetServerIp(v bool)`

SetServerIp sets ServerIp field to given value.

### HasServerIp

`func (o *LogForwarderDataFormat) HasServerIp() bool`

HasServerIp returns a boolean if a field has been set.

### GetFirewall

`func (o *LogForwarderDataFormat) GetFirewall() bool`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *LogForwarderDataFormat) GetFirewallOk() (*bool, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *LogForwarderDataFormat) SetFirewall(v bool)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *LogForwarderDataFormat) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.

### GetProxy

`func (o *LogForwarderDataFormat) GetProxy() bool`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *LogForwarderDataFormat) GetProxyOk() (*bool, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *LogForwarderDataFormat) SetProxy(v bool)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *LogForwarderDataFormat) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetDnsResolver

`func (o *LogForwarderDataFormat) GetDnsResolver() bool`

GetDnsResolver returns the DnsResolver field if non-nil, zero value otherwise.

### GetDnsResolverOk

`func (o *LogForwarderDataFormat) GetDnsResolverOk() (*bool, bool)`

GetDnsResolverOk returns a tuple with the DnsResolver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsResolver

`func (o *LogForwarderDataFormat) SetDnsResolver(v bool)`

SetDnsResolver sets DnsResolver field to given value.

### HasDnsResolver

`func (o *LogForwarderDataFormat) HasDnsResolver() bool`

HasDnsResolver returns a boolean if a field has been set.

### GetDdos

`func (o *LogForwarderDataFormat) GetDdos() bool`

GetDdos returns the Ddos field if non-nil, zero value otherwise.

### GetDdosOk

`func (o *LogForwarderDataFormat) GetDdosOk() (*bool, bool)`

GetDdosOk returns a tuple with the Ddos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdos

`func (o *LogForwarderDataFormat) SetDdos(v bool)`

SetDdos sets Ddos field to given value.

### HasDdos

`func (o *LogForwarderDataFormat) HasDdos() bool`

HasDdos returns a boolean if a field has been set.

### GetRatelimit

`func (o *LogForwarderDataFormat) GetRatelimit() bool`

GetRatelimit returns the Ratelimit field if non-nil, zero value otherwise.

### GetRatelimitOk

`func (o *LogForwarderDataFormat) GetRatelimitOk() (*bool, bool)`

GetRatelimitOk returns a tuple with the Ratelimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatelimit

`func (o *LogForwarderDataFormat) SetRatelimit(v bool)`

SetRatelimit sets Ratelimit field to given value.

### HasRatelimit

`func (o *LogForwarderDataFormat) HasRatelimit() bool`

HasRatelimit returns a boolean if a field has been set.

### GetWaf

`func (o *LogForwarderDataFormat) GetWaf() bool`

GetWaf returns the Waf field if non-nil, zero value otherwise.

### GetWafOk

`func (o *LogForwarderDataFormat) GetWafOk() (*bool, bool)`

GetWafOk returns a tuple with the Waf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaf

`func (o *LogForwarderDataFormat) SetWaf(v bool)`

SetWaf sets Waf field to given value.

### HasWaf

`func (o *LogForwarderDataFormat) HasWaf() bool`

HasWaf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


