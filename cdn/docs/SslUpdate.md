# SslUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | a user uploaded certificate&#39;s uuid or &#x60;managed&#x60; | [optional] 
**CertificateKeyType** | Pointer to **string** |  | [optional] 
**FingerprintStatus** | Pointer to **bool** | Whether Domain is using fingerprint or not | [optional] 
**SslStatus** | Pointer to **bool** | Whether Domain is using ssl module or not | [optional] 
**CertificateMode** | Pointer to **string** | Indicates certificate is managed by arvan, or its up to the user | [optional] [readonly] 
**TlsVersion** | Pointer to **string** | Minimum version of TLS. Empty (&#39;&#39;) means default. | [optional] 
**HstsStatus** | Pointer to **bool** | Whether HSTS is enabled | [optional] 
**HstsMaxAge** | Pointer to **string** | HSTS max age directive | [optional] 
**HstsSubdomain** | Pointer to **bool** |  | [optional] 
**HstsPreload** | Pointer to **bool** |  | [optional] 
**HttpsRedirect** | Pointer to **bool** |  | [optional] 
**ReplaceHttp** | Pointer to **bool** | Replace HTTP with HTTPs in HTML and JS sources | [optional] 
**Certificates** | Pointer to [**[]Certificate**](Certificate.md) |  | [optional] [readonly] 
**Orders** | Pointer to [**[]CertificateOrder**](CertificateOrder.md) | returns all \&quot;certificate orders\&quot; since the last invalid or canceled order | [optional] [readonly] 

## Methods

### NewSslUpdate

`func NewSslUpdate() *SslUpdate`

NewSslUpdate instantiates a new SslUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslUpdateWithDefaults

`func NewSslUpdateWithDefaults() *SslUpdate`

NewSslUpdateWithDefaults instantiates a new SslUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *SslUpdate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SslUpdate) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SslUpdate) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *SslUpdate) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCertificateKeyType

`func (o *SslUpdate) GetCertificateKeyType() string`

GetCertificateKeyType returns the CertificateKeyType field if non-nil, zero value otherwise.

### GetCertificateKeyTypeOk

`func (o *SslUpdate) GetCertificateKeyTypeOk() (*string, bool)`

GetCertificateKeyTypeOk returns a tuple with the CertificateKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateKeyType

`func (o *SslUpdate) SetCertificateKeyType(v string)`

SetCertificateKeyType sets CertificateKeyType field to given value.

### HasCertificateKeyType

`func (o *SslUpdate) HasCertificateKeyType() bool`

HasCertificateKeyType returns a boolean if a field has been set.

### GetFingerprintStatus

`func (o *SslUpdate) GetFingerprintStatus() bool`

GetFingerprintStatus returns the FingerprintStatus field if non-nil, zero value otherwise.

### GetFingerprintStatusOk

`func (o *SslUpdate) GetFingerprintStatusOk() (*bool, bool)`

GetFingerprintStatusOk returns a tuple with the FingerprintStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintStatus

`func (o *SslUpdate) SetFingerprintStatus(v bool)`

SetFingerprintStatus sets FingerprintStatus field to given value.

### HasFingerprintStatus

`func (o *SslUpdate) HasFingerprintStatus() bool`

HasFingerprintStatus returns a boolean if a field has been set.

### GetSslStatus

`func (o *SslUpdate) GetSslStatus() bool`

GetSslStatus returns the SslStatus field if non-nil, zero value otherwise.

### GetSslStatusOk

`func (o *SslUpdate) GetSslStatusOk() (*bool, bool)`

GetSslStatusOk returns a tuple with the SslStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslStatus

`func (o *SslUpdate) SetSslStatus(v bool)`

SetSslStatus sets SslStatus field to given value.

### HasSslStatus

`func (o *SslUpdate) HasSslStatus() bool`

HasSslStatus returns a boolean if a field has been set.

### GetCertificateMode

`func (o *SslUpdate) GetCertificateMode() string`

GetCertificateMode returns the CertificateMode field if non-nil, zero value otherwise.

### GetCertificateModeOk

`func (o *SslUpdate) GetCertificateModeOk() (*string, bool)`

GetCertificateModeOk returns a tuple with the CertificateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateMode

`func (o *SslUpdate) SetCertificateMode(v string)`

SetCertificateMode sets CertificateMode field to given value.

### HasCertificateMode

`func (o *SslUpdate) HasCertificateMode() bool`

HasCertificateMode returns a boolean if a field has been set.

### GetTlsVersion

`func (o *SslUpdate) GetTlsVersion() string`

GetTlsVersion returns the TlsVersion field if non-nil, zero value otherwise.

### GetTlsVersionOk

`func (o *SslUpdate) GetTlsVersionOk() (*string, bool)`

GetTlsVersionOk returns a tuple with the TlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsVersion

`func (o *SslUpdate) SetTlsVersion(v string)`

SetTlsVersion sets TlsVersion field to given value.

### HasTlsVersion

`func (o *SslUpdate) HasTlsVersion() bool`

HasTlsVersion returns a boolean if a field has been set.

### GetHstsStatus

`func (o *SslUpdate) GetHstsStatus() bool`

GetHstsStatus returns the HstsStatus field if non-nil, zero value otherwise.

### GetHstsStatusOk

`func (o *SslUpdate) GetHstsStatusOk() (*bool, bool)`

GetHstsStatusOk returns a tuple with the HstsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsStatus

`func (o *SslUpdate) SetHstsStatus(v bool)`

SetHstsStatus sets HstsStatus field to given value.

### HasHstsStatus

`func (o *SslUpdate) HasHstsStatus() bool`

HasHstsStatus returns a boolean if a field has been set.

### GetHstsMaxAge

`func (o *SslUpdate) GetHstsMaxAge() string`

GetHstsMaxAge returns the HstsMaxAge field if non-nil, zero value otherwise.

### GetHstsMaxAgeOk

`func (o *SslUpdate) GetHstsMaxAgeOk() (*string, bool)`

GetHstsMaxAgeOk returns a tuple with the HstsMaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsMaxAge

`func (o *SslUpdate) SetHstsMaxAge(v string)`

SetHstsMaxAge sets HstsMaxAge field to given value.

### HasHstsMaxAge

`func (o *SslUpdate) HasHstsMaxAge() bool`

HasHstsMaxAge returns a boolean if a field has been set.

### GetHstsSubdomain

`func (o *SslUpdate) GetHstsSubdomain() bool`

GetHstsSubdomain returns the HstsSubdomain field if non-nil, zero value otherwise.

### GetHstsSubdomainOk

`func (o *SslUpdate) GetHstsSubdomainOk() (*bool, bool)`

GetHstsSubdomainOk returns a tuple with the HstsSubdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsSubdomain

`func (o *SslUpdate) SetHstsSubdomain(v bool)`

SetHstsSubdomain sets HstsSubdomain field to given value.

### HasHstsSubdomain

`func (o *SslUpdate) HasHstsSubdomain() bool`

HasHstsSubdomain returns a boolean if a field has been set.

### GetHstsPreload

`func (o *SslUpdate) GetHstsPreload() bool`

GetHstsPreload returns the HstsPreload field if non-nil, zero value otherwise.

### GetHstsPreloadOk

`func (o *SslUpdate) GetHstsPreloadOk() (*bool, bool)`

GetHstsPreloadOk returns a tuple with the HstsPreload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsPreload

`func (o *SslUpdate) SetHstsPreload(v bool)`

SetHstsPreload sets HstsPreload field to given value.

### HasHstsPreload

`func (o *SslUpdate) HasHstsPreload() bool`

HasHstsPreload returns a boolean if a field has been set.

### GetHttpsRedirect

`func (o *SslUpdate) GetHttpsRedirect() bool`

GetHttpsRedirect returns the HttpsRedirect field if non-nil, zero value otherwise.

### GetHttpsRedirectOk

`func (o *SslUpdate) GetHttpsRedirectOk() (*bool, bool)`

GetHttpsRedirectOk returns a tuple with the HttpsRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsRedirect

`func (o *SslUpdate) SetHttpsRedirect(v bool)`

SetHttpsRedirect sets HttpsRedirect field to given value.

### HasHttpsRedirect

`func (o *SslUpdate) HasHttpsRedirect() bool`

HasHttpsRedirect returns a boolean if a field has been set.

### GetReplaceHttp

`func (o *SslUpdate) GetReplaceHttp() bool`

GetReplaceHttp returns the ReplaceHttp field if non-nil, zero value otherwise.

### GetReplaceHttpOk

`func (o *SslUpdate) GetReplaceHttpOk() (*bool, bool)`

GetReplaceHttpOk returns a tuple with the ReplaceHttp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceHttp

`func (o *SslUpdate) SetReplaceHttp(v bool)`

SetReplaceHttp sets ReplaceHttp field to given value.

### HasReplaceHttp

`func (o *SslUpdate) HasReplaceHttp() bool`

HasReplaceHttp returns a boolean if a field has been set.

### GetCertificates

`func (o *SslUpdate) GetCertificates() []Certificate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *SslUpdate) GetCertificatesOk() (*[]Certificate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *SslUpdate) SetCertificates(v []Certificate)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *SslUpdate) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetOrders

`func (o *SslUpdate) GetOrders() []CertificateOrder`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *SslUpdate) GetOrdersOk() (*[]CertificateOrder, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *SslUpdate) SetOrders(v []CertificateOrder)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *SslUpdate) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### SetOrdersNil

`func (o *SslUpdate) SetOrdersNil(b bool)`

 SetOrdersNil sets the value for Orders to be an explicit nil

### UnsetOrders
`func (o *SslUpdate) UnsetOrders()`

UnsetOrders ensures that no value is present for Orders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


