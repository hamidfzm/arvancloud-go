# V1ServiceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterIP** | Pointer to **string** | clusterIP is the IP address of the service and is usually assigned randomly by the master. If an address is specified manually and is not in use by others, it will be allocated to the service; otherwise, creation of the service will fail. This field can not be changed through updates. Valid values are \&quot;None\&quot;, empty string (\&quot;\&quot;), or a valid IP address. \&quot;None\&quot; can be specified for headless services when proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies | [optional] 
**ExternalIPs** | Pointer to **[]string** | externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service.  These IPs are not managed by Kubernetes.  The user is responsible for ensuring that traffic arrives at a node with this IP.  A common example is external load-balancers that are not part of the Kubernetes system. | [optional] 
**ExternalName** | Pointer to **string** | externalName is the external reference that kubedns or equivalent will return as a CNAME record for this service. No proxying will be involved. Must be a valid RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires Type to be ExternalName. | [optional] 
**ExternalTrafficPolicy** | Pointer to **string** | externalTrafficPolicy denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints. \&quot;Local\&quot; preserves the client source IP and avoids a second hop for LoadBalancer and Nodeport type services, but risks potentially imbalanced traffic spreading. \&quot;Cluster\&quot; obscures the client source IP and may cause a second hop to another node, but should have good overall load-spreading. | [optional] 
**HealthCheckNodePort** | Pointer to **int32** | healthCheckNodePort specifies the healthcheck nodePort for the service. If not specified, HealthCheckNodePort is created by the service api backend with the allocated nodePort. Will use user-specified nodePort value if specified by the client. Only effects when Type is set to LoadBalancer and ExternalTrafficPolicy is set to Local. | [optional] 
**LoadBalancerIP** | Pointer to **string** | Only applies to Service Type: LoadBalancer LoadBalancer will get created with the IP specified in this field. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature. | [optional] 
**LoadBalancerSourceRanges** | Pointer to **[]string** | If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature.\&quot; More info: https://kubernetes.io/docs/tasks/access-application-cluster/configure-cloud-provider-firewall/ | [optional] 
**Ports** | Pointer to [**[]V1ServicePort**](V1ServicePort.md) | The list of ports that are exposed by this service. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies | [optional] 
**PublishNotReadyAddresses** | Pointer to **bool** | publishNotReadyAddresses, when set to true, indicates that DNS implementations must publish the notReadyAddresses of subsets for the Endpoints associated with the Service. The default value is false. The primary use case for setting this field is to use a StatefulSet&#39;s Headless Service to propagate SRV records for its Pods without respect to their readiness for purpose of peer discovery. | [optional] 
**Selector** | Pointer to **map[string]interface{}** | Route service traffic to pods with label keys and values matching this selector. If empty or not present, the service is assumed to have an external process managing its endpoints, which Kubernetes will not modify. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/ | [optional] 
**SessionAffinity** | Pointer to **string** | Supports \&quot;ClientIP\&quot; and \&quot;None\&quot;. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies | [optional] 
**SessionAffinityConfig** | Pointer to [**V1SessionAffinityConfig**](V1SessionAffinityConfig.md) |  | [optional] 
**Type** | Pointer to **string** | type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. \&quot;ExternalName\&quot; maps to the specified externalName. \&quot;ClusterIP\&quot; allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object. If clusterIP is \&quot;None\&quot;, no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a stable IP. \&quot;NodePort\&quot; builds on ClusterIP and allocates a port on every node which routes to the clusterIP. \&quot;LoadBalancer\&quot; builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the clusterIP. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services  | [optional] 

## Methods

### NewV1ServiceSpec

`func NewV1ServiceSpec() *V1ServiceSpec`

NewV1ServiceSpec instantiates a new V1ServiceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ServiceSpecWithDefaults

`func NewV1ServiceSpecWithDefaults() *V1ServiceSpec`

NewV1ServiceSpecWithDefaults instantiates a new V1ServiceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIP

`func (o *V1ServiceSpec) GetClusterIP() string`

GetClusterIP returns the ClusterIP field if non-nil, zero value otherwise.

### GetClusterIPOk

`func (o *V1ServiceSpec) GetClusterIPOk() (*string, bool)`

GetClusterIPOk returns a tuple with the ClusterIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIP

`func (o *V1ServiceSpec) SetClusterIP(v string)`

SetClusterIP sets ClusterIP field to given value.

### HasClusterIP

`func (o *V1ServiceSpec) HasClusterIP() bool`

HasClusterIP returns a boolean if a field has been set.

### GetExternalIPs

`func (o *V1ServiceSpec) GetExternalIPs() []string`

GetExternalIPs returns the ExternalIPs field if non-nil, zero value otherwise.

### GetExternalIPsOk

`func (o *V1ServiceSpec) GetExternalIPsOk() (*[]string, bool)`

GetExternalIPsOk returns a tuple with the ExternalIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIPs

`func (o *V1ServiceSpec) SetExternalIPs(v []string)`

SetExternalIPs sets ExternalIPs field to given value.

### HasExternalIPs

`func (o *V1ServiceSpec) HasExternalIPs() bool`

HasExternalIPs returns a boolean if a field has been set.

### GetExternalName

`func (o *V1ServiceSpec) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *V1ServiceSpec) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *V1ServiceSpec) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *V1ServiceSpec) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### GetExternalTrafficPolicy

`func (o *V1ServiceSpec) GetExternalTrafficPolicy() string`

GetExternalTrafficPolicy returns the ExternalTrafficPolicy field if non-nil, zero value otherwise.

### GetExternalTrafficPolicyOk

`func (o *V1ServiceSpec) GetExternalTrafficPolicyOk() (*string, bool)`

GetExternalTrafficPolicyOk returns a tuple with the ExternalTrafficPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTrafficPolicy

`func (o *V1ServiceSpec) SetExternalTrafficPolicy(v string)`

SetExternalTrafficPolicy sets ExternalTrafficPolicy field to given value.

### HasExternalTrafficPolicy

`func (o *V1ServiceSpec) HasExternalTrafficPolicy() bool`

HasExternalTrafficPolicy returns a boolean if a field has been set.

### GetHealthCheckNodePort

`func (o *V1ServiceSpec) GetHealthCheckNodePort() int32`

GetHealthCheckNodePort returns the HealthCheckNodePort field if non-nil, zero value otherwise.

### GetHealthCheckNodePortOk

`func (o *V1ServiceSpec) GetHealthCheckNodePortOk() (*int32, bool)`

GetHealthCheckNodePortOk returns a tuple with the HealthCheckNodePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckNodePort

`func (o *V1ServiceSpec) SetHealthCheckNodePort(v int32)`

SetHealthCheckNodePort sets HealthCheckNodePort field to given value.

### HasHealthCheckNodePort

`func (o *V1ServiceSpec) HasHealthCheckNodePort() bool`

HasHealthCheckNodePort returns a boolean if a field has been set.

### GetLoadBalancerIP

`func (o *V1ServiceSpec) GetLoadBalancerIP() string`

GetLoadBalancerIP returns the LoadBalancerIP field if non-nil, zero value otherwise.

### GetLoadBalancerIPOk

`func (o *V1ServiceSpec) GetLoadBalancerIPOk() (*string, bool)`

GetLoadBalancerIPOk returns a tuple with the LoadBalancerIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerIP

`func (o *V1ServiceSpec) SetLoadBalancerIP(v string)`

SetLoadBalancerIP sets LoadBalancerIP field to given value.

### HasLoadBalancerIP

`func (o *V1ServiceSpec) HasLoadBalancerIP() bool`

HasLoadBalancerIP returns a boolean if a field has been set.

### GetLoadBalancerSourceRanges

`func (o *V1ServiceSpec) GetLoadBalancerSourceRanges() []string`

GetLoadBalancerSourceRanges returns the LoadBalancerSourceRanges field if non-nil, zero value otherwise.

### GetLoadBalancerSourceRangesOk

`func (o *V1ServiceSpec) GetLoadBalancerSourceRangesOk() (*[]string, bool)`

GetLoadBalancerSourceRangesOk returns a tuple with the LoadBalancerSourceRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerSourceRanges

`func (o *V1ServiceSpec) SetLoadBalancerSourceRanges(v []string)`

SetLoadBalancerSourceRanges sets LoadBalancerSourceRanges field to given value.

### HasLoadBalancerSourceRanges

`func (o *V1ServiceSpec) HasLoadBalancerSourceRanges() bool`

HasLoadBalancerSourceRanges returns a boolean if a field has been set.

### GetPorts

`func (o *V1ServiceSpec) GetPorts() []V1ServicePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *V1ServiceSpec) GetPortsOk() (*[]V1ServicePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *V1ServiceSpec) SetPorts(v []V1ServicePort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *V1ServiceSpec) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPublishNotReadyAddresses

`func (o *V1ServiceSpec) GetPublishNotReadyAddresses() bool`

GetPublishNotReadyAddresses returns the PublishNotReadyAddresses field if non-nil, zero value otherwise.

### GetPublishNotReadyAddressesOk

`func (o *V1ServiceSpec) GetPublishNotReadyAddressesOk() (*bool, bool)`

GetPublishNotReadyAddressesOk returns a tuple with the PublishNotReadyAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishNotReadyAddresses

`func (o *V1ServiceSpec) SetPublishNotReadyAddresses(v bool)`

SetPublishNotReadyAddresses sets PublishNotReadyAddresses field to given value.

### HasPublishNotReadyAddresses

`func (o *V1ServiceSpec) HasPublishNotReadyAddresses() bool`

HasPublishNotReadyAddresses returns a boolean if a field has been set.

### GetSelector

`func (o *V1ServiceSpec) GetSelector() map[string]interface{}`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *V1ServiceSpec) GetSelectorOk() (*map[string]interface{}, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *V1ServiceSpec) SetSelector(v map[string]interface{})`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *V1ServiceSpec) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetSessionAffinity

`func (o *V1ServiceSpec) GetSessionAffinity() string`

GetSessionAffinity returns the SessionAffinity field if non-nil, zero value otherwise.

### GetSessionAffinityOk

`func (o *V1ServiceSpec) GetSessionAffinityOk() (*string, bool)`

GetSessionAffinityOk returns a tuple with the SessionAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionAffinity

`func (o *V1ServiceSpec) SetSessionAffinity(v string)`

SetSessionAffinity sets SessionAffinity field to given value.

### HasSessionAffinity

`func (o *V1ServiceSpec) HasSessionAffinity() bool`

HasSessionAffinity returns a boolean if a field has been set.

### GetSessionAffinityConfig

`func (o *V1ServiceSpec) GetSessionAffinityConfig() V1SessionAffinityConfig`

GetSessionAffinityConfig returns the SessionAffinityConfig field if non-nil, zero value otherwise.

### GetSessionAffinityConfigOk

`func (o *V1ServiceSpec) GetSessionAffinityConfigOk() (*V1SessionAffinityConfig, bool)`

GetSessionAffinityConfigOk returns a tuple with the SessionAffinityConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionAffinityConfig

`func (o *V1ServiceSpec) SetSessionAffinityConfig(v V1SessionAffinityConfig)`

SetSessionAffinityConfig sets SessionAffinityConfig field to given value.

### HasSessionAffinityConfig

`func (o *V1ServiceSpec) HasSessionAffinityConfig() bool`

HasSessionAffinityConfig returns a boolean if a field has been set.

### GetType

`func (o *V1ServiceSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1ServiceSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1ServiceSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1ServiceSpec) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


