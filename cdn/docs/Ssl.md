# Ssl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewSsl

`func NewSsl() *Ssl`

NewSsl instantiates a new Ssl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslWithDefaults

`func NewSslWithDefaults() *Ssl`

NewSslWithDefaults instantiates a new Ssl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFingerprintStatus

`func (o *Ssl) GetFingerprintStatus() bool`

GetFingerprintStatus returns the FingerprintStatus field if non-nil, zero value otherwise.

### GetFingerprintStatusOk

`func (o *Ssl) GetFingerprintStatusOk() (*bool, bool)`

GetFingerprintStatusOk returns a tuple with the FingerprintStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintStatus

`func (o *Ssl) SetFingerprintStatus(v bool)`

SetFingerprintStatus sets FingerprintStatus field to given value.

### HasFingerprintStatus

`func (o *Ssl) HasFingerprintStatus() bool`

HasFingerprintStatus returns a boolean if a field has been set.

### GetSslStatus

`func (o *Ssl) GetSslStatus() bool`

GetSslStatus returns the SslStatus field if non-nil, zero value otherwise.

### GetSslStatusOk

`func (o *Ssl) GetSslStatusOk() (*bool, bool)`

GetSslStatusOk returns a tuple with the SslStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslStatus

`func (o *Ssl) SetSslStatus(v bool)`

SetSslStatus sets SslStatus field to given value.

### HasSslStatus

`func (o *Ssl) HasSslStatus() bool`

HasSslStatus returns a boolean if a field has been set.

### GetCertificateMode

`func (o *Ssl) GetCertificateMode() string`

GetCertificateMode returns the CertificateMode field if non-nil, zero value otherwise.

### GetCertificateModeOk

`func (o *Ssl) GetCertificateModeOk() (*string, bool)`

GetCertificateModeOk returns a tuple with the CertificateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateMode

`func (o *Ssl) SetCertificateMode(v string)`

SetCertificateMode sets CertificateMode field to given value.

### HasCertificateMode

`func (o *Ssl) HasCertificateMode() bool`

HasCertificateMode returns a boolean if a field has been set.

### GetTlsVersion

`func (o *Ssl) GetTlsVersion() string`

GetTlsVersion returns the TlsVersion field if non-nil, zero value otherwise.

### GetTlsVersionOk

`func (o *Ssl) GetTlsVersionOk() (*string, bool)`

GetTlsVersionOk returns a tuple with the TlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsVersion

`func (o *Ssl) SetTlsVersion(v string)`

SetTlsVersion sets TlsVersion field to given value.

### HasTlsVersion

`func (o *Ssl) HasTlsVersion() bool`

HasTlsVersion returns a boolean if a field has been set.

### GetHstsStatus

`func (o *Ssl) GetHstsStatus() bool`

GetHstsStatus returns the HstsStatus field if non-nil, zero value otherwise.

### GetHstsStatusOk

`func (o *Ssl) GetHstsStatusOk() (*bool, bool)`

GetHstsStatusOk returns a tuple with the HstsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsStatus

`func (o *Ssl) SetHstsStatus(v bool)`

SetHstsStatus sets HstsStatus field to given value.

### HasHstsStatus

`func (o *Ssl) HasHstsStatus() bool`

HasHstsStatus returns a boolean if a field has been set.

### GetHstsMaxAge

`func (o *Ssl) GetHstsMaxAge() string`

GetHstsMaxAge returns the HstsMaxAge field if non-nil, zero value otherwise.

### GetHstsMaxAgeOk

`func (o *Ssl) GetHstsMaxAgeOk() (*string, bool)`

GetHstsMaxAgeOk returns a tuple with the HstsMaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsMaxAge

`func (o *Ssl) SetHstsMaxAge(v string)`

SetHstsMaxAge sets HstsMaxAge field to given value.

### HasHstsMaxAge

`func (o *Ssl) HasHstsMaxAge() bool`

HasHstsMaxAge returns a boolean if a field has been set.

### GetHstsSubdomain

`func (o *Ssl) GetHstsSubdomain() bool`

GetHstsSubdomain returns the HstsSubdomain field if non-nil, zero value otherwise.

### GetHstsSubdomainOk

`func (o *Ssl) GetHstsSubdomainOk() (*bool, bool)`

GetHstsSubdomainOk returns a tuple with the HstsSubdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsSubdomain

`func (o *Ssl) SetHstsSubdomain(v bool)`

SetHstsSubdomain sets HstsSubdomain field to given value.

### HasHstsSubdomain

`func (o *Ssl) HasHstsSubdomain() bool`

HasHstsSubdomain returns a boolean if a field has been set.

### GetHstsPreload

`func (o *Ssl) GetHstsPreload() bool`

GetHstsPreload returns the HstsPreload field if non-nil, zero value otherwise.

### GetHstsPreloadOk

`func (o *Ssl) GetHstsPreloadOk() (*bool, bool)`

GetHstsPreloadOk returns a tuple with the HstsPreload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHstsPreload

`func (o *Ssl) SetHstsPreload(v bool)`

SetHstsPreload sets HstsPreload field to given value.

### HasHstsPreload

`func (o *Ssl) HasHstsPreload() bool`

HasHstsPreload returns a boolean if a field has been set.

### GetHttpsRedirect

`func (o *Ssl) GetHttpsRedirect() bool`

GetHttpsRedirect returns the HttpsRedirect field if non-nil, zero value otherwise.

### GetHttpsRedirectOk

`func (o *Ssl) GetHttpsRedirectOk() (*bool, bool)`

GetHttpsRedirectOk returns a tuple with the HttpsRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsRedirect

`func (o *Ssl) SetHttpsRedirect(v bool)`

SetHttpsRedirect sets HttpsRedirect field to given value.

### HasHttpsRedirect

`func (o *Ssl) HasHttpsRedirect() bool`

HasHttpsRedirect returns a boolean if a field has been set.

### GetReplaceHttp

`func (o *Ssl) GetReplaceHttp() bool`

GetReplaceHttp returns the ReplaceHttp field if non-nil, zero value otherwise.

### GetReplaceHttpOk

`func (o *Ssl) GetReplaceHttpOk() (*bool, bool)`

GetReplaceHttpOk returns a tuple with the ReplaceHttp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceHttp

`func (o *Ssl) SetReplaceHttp(v bool)`

SetReplaceHttp sets ReplaceHttp field to given value.

### HasReplaceHttp

`func (o *Ssl) HasReplaceHttp() bool`

HasReplaceHttp returns a boolean if a field has been set.

### GetCertificates

`func (o *Ssl) GetCertificates() []Certificate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *Ssl) GetCertificatesOk() (*[]Certificate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *Ssl) SetCertificates(v []Certificate)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *Ssl) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetOrders

`func (o *Ssl) GetOrders() []CertificateOrder`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *Ssl) GetOrdersOk() (*[]CertificateOrder, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *Ssl) SetOrders(v []CertificateOrder)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *Ssl) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### SetOrdersNil

`func (o *Ssl) SetOrdersNil(b bool)`

 SetOrdersNil sets the value for Orders to be an explicit nil

### UnsetOrders
`func (o *Ssl) UnsetOrders()`

UnsetOrders ensures that no value is present for Orders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


