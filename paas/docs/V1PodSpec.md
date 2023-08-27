# V1PodSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveDeadlineSeconds** | Pointer to **int64** | Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer. | [optional] 
**Affinity** | Pointer to [**V1Affinity**](V1Affinity.md) |  | [optional] 
**AutomountServiceAccountToken** | Pointer to **bool** | AutomountServiceAccountToken indicates whether a service account token should be automatically mounted. | [optional] 
**Containers** | [**[]V1Container**](V1Container.md) | List of containers belonging to the pod. Containers cannot currently be added or removed. There must be at least one container in a Pod. Cannot be updated. | 
**DnsConfig** | Pointer to [**V1PodDNSConfig**](V1PodDNSConfig.md) |  | [optional] 
**DnsPolicy** | Pointer to **string** | Set DNS policy for the pod. Defaults to \&quot;ClusterFirst\&quot;. Valid values are &#39;ClusterFirstWithHostNet&#39;, &#39;ClusterFirst&#39;, &#39;Default&#39; or &#39;None&#39;. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to &#39;ClusterFirstWithHostNet&#39;. | [optional] 
**HostAliases** | Pointer to [**[]V1HostAlias**](V1HostAlias.md) | HostAliases is an optional list of hosts and IPs that will be injected into the pod&#39;s hosts file if specified. This is only valid for non-hostNetwork pods. | [optional] 
**HostIPC** | Pointer to **bool** | Use the host&#39;s ipc namespace. Optional: Default to false. | [optional] 
**HostNetwork** | Pointer to **bool** | Host networking requested for this pod. Use the host&#39;s network namespace. If this option is set, the ports that will be used must be specified. Default to false. | [optional] 
**HostPID** | Pointer to **bool** | Use the host&#39;s pid namespace. Optional: Default to false. | [optional] 
**Hostname** | Pointer to **string** | Specifies the hostname of the Pod If not specified, the pod&#39;s hostname will be set to a system-defined value. | [optional] 
**ImagePullSecrets** | Pointer to [**[]V1LocalObjectReference**](V1LocalObjectReference.md) | ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod | [optional] 
**InitContainers** | Pointer to [**[]V1Container**](V1Container.md) | List of initialization containers belonging to the pod. Init containers are executed in order prior to containers being started. If any init container fails, the pod is considered to have failed and is handled according to its restartPolicy. The name for an init container or normal container must be unique among all containers. Init containers may not have Lifecycle actions, Readiness probes, or Liveness probes. The resourceRequirements of an init container are taken into account during scheduling by finding the highest request/limit for each resource type, and then using the max of of that value or the sum of the normal containers. Limits are applied to init containers in a similar fashion. Init containers cannot currently be added or removed. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/ | [optional] 
**NodeName** | Pointer to **string** | NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements. | [optional] 
**NodeSelector** | Pointer to **map[string]interface{}** | NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node&#39;s labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ | [optional] 
**Priority** | Pointer to **int32** | The priority value. Various system components use this field to find the priority of the pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority. | [optional] 
**PriorityClassName** | Pointer to **string** | If specified, indicates the pod&#39;s priority. \&quot;system-node-critical\&quot; and \&quot;system-cluster-critical\&quot; are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default. | [optional] 
**ReadinessGates** | Pointer to [**[]V1PodReadinessGate**](V1PodReadinessGate.md) | If specified, all readiness gates will be evaluated for pod readiness. A pod is ready when all its containers are ready AND all conditions specified in the readiness gates have status equal to \&quot;True\&quot; More info: https://github.com/kubernetes/community/blob/master/keps/sig-network/0007-pod-ready%2B%2B.md | [optional] 
**RestartPolicy** | Pointer to **string** | Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to Always. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy | [optional] 
**SchedulerName** | Pointer to **string** | If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler. | [optional] 
**SecurityContext** | Pointer to [**V1PodSecurityContext**](V1PodSecurityContext.md) |  | [optional] 
**ServiceAccount** | Pointer to **string** | DeprecatedServiceAccount is a depreciated alias for ServiceAccountName. Deprecated: Use serviceAccountName instead. | [optional] 
**ServiceAccountName** | Pointer to **string** | ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/ | [optional] 
**ShareProcessNamespace** | Pointer to **bool** | Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set. Optional: Default to false. This field is alpha-level and is honored only by servers that enable the PodShareProcessNamespace feature. | [optional] 
**Subdomain** | Pointer to **string** | If specified, the fully qualified Pod hostname will be \&quot;&lt;hostname&gt;.&lt;subdomain&gt;.&lt;pod namespace&gt;.svc.&lt;cluster domain&gt;\&quot;. If not specified, the pod will not have a domainname at all. | [optional] 
**TerminationGracePeriodSeconds** | Pointer to **int64** | Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds. | [optional] 
**Tolerations** | Pointer to [**[]V1Toleration**](V1Toleration.md) | If specified, the pod&#39;s tolerations. | [optional] 
**Volumes** | Pointer to [**[]V1Volume**](V1Volume.md) | List of volumes that can be mounted by containers belonging to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes | [optional] 

## Methods

### NewV1PodSpec

`func NewV1PodSpec(containers []V1Container, ) *V1PodSpec`

NewV1PodSpec instantiates a new V1PodSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PodSpecWithDefaults

`func NewV1PodSpecWithDefaults() *V1PodSpec`

NewV1PodSpecWithDefaults instantiates a new V1PodSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDeadlineSeconds

`func (o *V1PodSpec) GetActiveDeadlineSeconds() int64`

GetActiveDeadlineSeconds returns the ActiveDeadlineSeconds field if non-nil, zero value otherwise.

### GetActiveDeadlineSecondsOk

`func (o *V1PodSpec) GetActiveDeadlineSecondsOk() (*int64, bool)`

GetActiveDeadlineSecondsOk returns a tuple with the ActiveDeadlineSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDeadlineSeconds

`func (o *V1PodSpec) SetActiveDeadlineSeconds(v int64)`

SetActiveDeadlineSeconds sets ActiveDeadlineSeconds field to given value.

### HasActiveDeadlineSeconds

`func (o *V1PodSpec) HasActiveDeadlineSeconds() bool`

HasActiveDeadlineSeconds returns a boolean if a field has been set.

### GetAffinity

`func (o *V1PodSpec) GetAffinity() V1Affinity`

GetAffinity returns the Affinity field if non-nil, zero value otherwise.

### GetAffinityOk

`func (o *V1PodSpec) GetAffinityOk() (*V1Affinity, bool)`

GetAffinityOk returns a tuple with the Affinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinity

`func (o *V1PodSpec) SetAffinity(v V1Affinity)`

SetAffinity sets Affinity field to given value.

### HasAffinity

`func (o *V1PodSpec) HasAffinity() bool`

HasAffinity returns a boolean if a field has been set.

### GetAutomountServiceAccountToken

`func (o *V1PodSpec) GetAutomountServiceAccountToken() bool`

GetAutomountServiceAccountToken returns the AutomountServiceAccountToken field if non-nil, zero value otherwise.

### GetAutomountServiceAccountTokenOk

`func (o *V1PodSpec) GetAutomountServiceAccountTokenOk() (*bool, bool)`

GetAutomountServiceAccountTokenOk returns a tuple with the AutomountServiceAccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomountServiceAccountToken

`func (o *V1PodSpec) SetAutomountServiceAccountToken(v bool)`

SetAutomountServiceAccountToken sets AutomountServiceAccountToken field to given value.

### HasAutomountServiceAccountToken

`func (o *V1PodSpec) HasAutomountServiceAccountToken() bool`

HasAutomountServiceAccountToken returns a boolean if a field has been set.

### GetContainers

`func (o *V1PodSpec) GetContainers() []V1Container`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *V1PodSpec) GetContainersOk() (*[]V1Container, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *V1PodSpec) SetContainers(v []V1Container)`

SetContainers sets Containers field to given value.


### GetDnsConfig

`func (o *V1PodSpec) GetDnsConfig() V1PodDNSConfig`

GetDnsConfig returns the DnsConfig field if non-nil, zero value otherwise.

### GetDnsConfigOk

`func (o *V1PodSpec) GetDnsConfigOk() (*V1PodDNSConfig, bool)`

GetDnsConfigOk returns a tuple with the DnsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsConfig

`func (o *V1PodSpec) SetDnsConfig(v V1PodDNSConfig)`

SetDnsConfig sets DnsConfig field to given value.

### HasDnsConfig

`func (o *V1PodSpec) HasDnsConfig() bool`

HasDnsConfig returns a boolean if a field has been set.

### GetDnsPolicy

`func (o *V1PodSpec) GetDnsPolicy() string`

GetDnsPolicy returns the DnsPolicy field if non-nil, zero value otherwise.

### GetDnsPolicyOk

`func (o *V1PodSpec) GetDnsPolicyOk() (*string, bool)`

GetDnsPolicyOk returns a tuple with the DnsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPolicy

`func (o *V1PodSpec) SetDnsPolicy(v string)`

SetDnsPolicy sets DnsPolicy field to given value.

### HasDnsPolicy

`func (o *V1PodSpec) HasDnsPolicy() bool`

HasDnsPolicy returns a boolean if a field has been set.

### GetHostAliases

`func (o *V1PodSpec) GetHostAliases() []V1HostAlias`

GetHostAliases returns the HostAliases field if non-nil, zero value otherwise.

### GetHostAliasesOk

`func (o *V1PodSpec) GetHostAliasesOk() (*[]V1HostAlias, bool)`

GetHostAliasesOk returns a tuple with the HostAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostAliases

`func (o *V1PodSpec) SetHostAliases(v []V1HostAlias)`

SetHostAliases sets HostAliases field to given value.

### HasHostAliases

`func (o *V1PodSpec) HasHostAliases() bool`

HasHostAliases returns a boolean if a field has been set.

### GetHostIPC

`func (o *V1PodSpec) GetHostIPC() bool`

GetHostIPC returns the HostIPC field if non-nil, zero value otherwise.

### GetHostIPCOk

`func (o *V1PodSpec) GetHostIPCOk() (*bool, bool)`

GetHostIPCOk returns a tuple with the HostIPC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIPC

`func (o *V1PodSpec) SetHostIPC(v bool)`

SetHostIPC sets HostIPC field to given value.

### HasHostIPC

`func (o *V1PodSpec) HasHostIPC() bool`

HasHostIPC returns a boolean if a field has been set.

### GetHostNetwork

`func (o *V1PodSpec) GetHostNetwork() bool`

GetHostNetwork returns the HostNetwork field if non-nil, zero value otherwise.

### GetHostNetworkOk

`func (o *V1PodSpec) GetHostNetworkOk() (*bool, bool)`

GetHostNetworkOk returns a tuple with the HostNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNetwork

`func (o *V1PodSpec) SetHostNetwork(v bool)`

SetHostNetwork sets HostNetwork field to given value.

### HasHostNetwork

`func (o *V1PodSpec) HasHostNetwork() bool`

HasHostNetwork returns a boolean if a field has been set.

### GetHostPID

`func (o *V1PodSpec) GetHostPID() bool`

GetHostPID returns the HostPID field if non-nil, zero value otherwise.

### GetHostPIDOk

`func (o *V1PodSpec) GetHostPIDOk() (*bool, bool)`

GetHostPIDOk returns a tuple with the HostPID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPID

`func (o *V1PodSpec) SetHostPID(v bool)`

SetHostPID sets HostPID field to given value.

### HasHostPID

`func (o *V1PodSpec) HasHostPID() bool`

HasHostPID returns a boolean if a field has been set.

### GetHostname

`func (o *V1PodSpec) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V1PodSpec) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V1PodSpec) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *V1PodSpec) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetImagePullSecrets

`func (o *V1PodSpec) GetImagePullSecrets() []V1LocalObjectReference`

GetImagePullSecrets returns the ImagePullSecrets field if non-nil, zero value otherwise.

### GetImagePullSecretsOk

`func (o *V1PodSpec) GetImagePullSecretsOk() (*[]V1LocalObjectReference, bool)`

GetImagePullSecretsOk returns a tuple with the ImagePullSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePullSecrets

`func (o *V1PodSpec) SetImagePullSecrets(v []V1LocalObjectReference)`

SetImagePullSecrets sets ImagePullSecrets field to given value.

### HasImagePullSecrets

`func (o *V1PodSpec) HasImagePullSecrets() bool`

HasImagePullSecrets returns a boolean if a field has been set.

### GetInitContainers

`func (o *V1PodSpec) GetInitContainers() []V1Container`

GetInitContainers returns the InitContainers field if non-nil, zero value otherwise.

### GetInitContainersOk

`func (o *V1PodSpec) GetInitContainersOk() (*[]V1Container, bool)`

GetInitContainersOk returns a tuple with the InitContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitContainers

`func (o *V1PodSpec) SetInitContainers(v []V1Container)`

SetInitContainers sets InitContainers field to given value.

### HasInitContainers

`func (o *V1PodSpec) HasInitContainers() bool`

HasInitContainers returns a boolean if a field has been set.

### GetNodeName

`func (o *V1PodSpec) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *V1PodSpec) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *V1PodSpec) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *V1PodSpec) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetNodeSelector

`func (o *V1PodSpec) GetNodeSelector() map[string]interface{}`

GetNodeSelector returns the NodeSelector field if non-nil, zero value otherwise.

### GetNodeSelectorOk

`func (o *V1PodSpec) GetNodeSelectorOk() (*map[string]interface{}, bool)`

GetNodeSelectorOk returns a tuple with the NodeSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSelector

`func (o *V1PodSpec) SetNodeSelector(v map[string]interface{})`

SetNodeSelector sets NodeSelector field to given value.

### HasNodeSelector

`func (o *V1PodSpec) HasNodeSelector() bool`

HasNodeSelector returns a boolean if a field has been set.

### GetPriority

`func (o *V1PodSpec) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V1PodSpec) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V1PodSpec) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V1PodSpec) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPriorityClassName

`func (o *V1PodSpec) GetPriorityClassName() string`

GetPriorityClassName returns the PriorityClassName field if non-nil, zero value otherwise.

### GetPriorityClassNameOk

`func (o *V1PodSpec) GetPriorityClassNameOk() (*string, bool)`

GetPriorityClassNameOk returns a tuple with the PriorityClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityClassName

`func (o *V1PodSpec) SetPriorityClassName(v string)`

SetPriorityClassName sets PriorityClassName field to given value.

### HasPriorityClassName

`func (o *V1PodSpec) HasPriorityClassName() bool`

HasPriorityClassName returns a boolean if a field has been set.

### GetReadinessGates

`func (o *V1PodSpec) GetReadinessGates() []V1PodReadinessGate`

GetReadinessGates returns the ReadinessGates field if non-nil, zero value otherwise.

### GetReadinessGatesOk

`func (o *V1PodSpec) GetReadinessGatesOk() (*[]V1PodReadinessGate, bool)`

GetReadinessGatesOk returns a tuple with the ReadinessGates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessGates

`func (o *V1PodSpec) SetReadinessGates(v []V1PodReadinessGate)`

SetReadinessGates sets ReadinessGates field to given value.

### HasReadinessGates

`func (o *V1PodSpec) HasReadinessGates() bool`

HasReadinessGates returns a boolean if a field has been set.

### GetRestartPolicy

`func (o *V1PodSpec) GetRestartPolicy() string`

GetRestartPolicy returns the RestartPolicy field if non-nil, zero value otherwise.

### GetRestartPolicyOk

`func (o *V1PodSpec) GetRestartPolicyOk() (*string, bool)`

GetRestartPolicyOk returns a tuple with the RestartPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartPolicy

`func (o *V1PodSpec) SetRestartPolicy(v string)`

SetRestartPolicy sets RestartPolicy field to given value.

### HasRestartPolicy

`func (o *V1PodSpec) HasRestartPolicy() bool`

HasRestartPolicy returns a boolean if a field has been set.

### GetSchedulerName

`func (o *V1PodSpec) GetSchedulerName() string`

GetSchedulerName returns the SchedulerName field if non-nil, zero value otherwise.

### GetSchedulerNameOk

`func (o *V1PodSpec) GetSchedulerNameOk() (*string, bool)`

GetSchedulerNameOk returns a tuple with the SchedulerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerName

`func (o *V1PodSpec) SetSchedulerName(v string)`

SetSchedulerName sets SchedulerName field to given value.

### HasSchedulerName

`func (o *V1PodSpec) HasSchedulerName() bool`

HasSchedulerName returns a boolean if a field has been set.

### GetSecurityContext

`func (o *V1PodSpec) GetSecurityContext() V1PodSecurityContext`

GetSecurityContext returns the SecurityContext field if non-nil, zero value otherwise.

### GetSecurityContextOk

`func (o *V1PodSpec) GetSecurityContextOk() (*V1PodSecurityContext, bool)`

GetSecurityContextOk returns a tuple with the SecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityContext

`func (o *V1PodSpec) SetSecurityContext(v V1PodSecurityContext)`

SetSecurityContext sets SecurityContext field to given value.

### HasSecurityContext

`func (o *V1PodSpec) HasSecurityContext() bool`

HasSecurityContext returns a boolean if a field has been set.

### GetServiceAccount

`func (o *V1PodSpec) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *V1PodSpec) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *V1PodSpec) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *V1PodSpec) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetServiceAccountName

`func (o *V1PodSpec) GetServiceAccountName() string`

GetServiceAccountName returns the ServiceAccountName field if non-nil, zero value otherwise.

### GetServiceAccountNameOk

`func (o *V1PodSpec) GetServiceAccountNameOk() (*string, bool)`

GetServiceAccountNameOk returns a tuple with the ServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountName

`func (o *V1PodSpec) SetServiceAccountName(v string)`

SetServiceAccountName sets ServiceAccountName field to given value.

### HasServiceAccountName

`func (o *V1PodSpec) HasServiceAccountName() bool`

HasServiceAccountName returns a boolean if a field has been set.

### GetShareProcessNamespace

`func (o *V1PodSpec) GetShareProcessNamespace() bool`

GetShareProcessNamespace returns the ShareProcessNamespace field if non-nil, zero value otherwise.

### GetShareProcessNamespaceOk

`func (o *V1PodSpec) GetShareProcessNamespaceOk() (*bool, bool)`

GetShareProcessNamespaceOk returns a tuple with the ShareProcessNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareProcessNamespace

`func (o *V1PodSpec) SetShareProcessNamespace(v bool)`

SetShareProcessNamespace sets ShareProcessNamespace field to given value.

### HasShareProcessNamespace

`func (o *V1PodSpec) HasShareProcessNamespace() bool`

HasShareProcessNamespace returns a boolean if a field has been set.

### GetSubdomain

`func (o *V1PodSpec) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *V1PodSpec) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *V1PodSpec) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *V1PodSpec) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetTerminationGracePeriodSeconds

`func (o *V1PodSpec) GetTerminationGracePeriodSeconds() int64`

GetTerminationGracePeriodSeconds returns the TerminationGracePeriodSeconds field if non-nil, zero value otherwise.

### GetTerminationGracePeriodSecondsOk

`func (o *V1PodSpec) GetTerminationGracePeriodSecondsOk() (*int64, bool)`

GetTerminationGracePeriodSecondsOk returns a tuple with the TerminationGracePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationGracePeriodSeconds

`func (o *V1PodSpec) SetTerminationGracePeriodSeconds(v int64)`

SetTerminationGracePeriodSeconds sets TerminationGracePeriodSeconds field to given value.

### HasTerminationGracePeriodSeconds

`func (o *V1PodSpec) HasTerminationGracePeriodSeconds() bool`

HasTerminationGracePeriodSeconds returns a boolean if a field has been set.

### GetTolerations

`func (o *V1PodSpec) GetTolerations() []V1Toleration`

GetTolerations returns the Tolerations field if non-nil, zero value otherwise.

### GetTolerationsOk

`func (o *V1PodSpec) GetTolerationsOk() (*[]V1Toleration, bool)`

GetTolerationsOk returns a tuple with the Tolerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTolerations

`func (o *V1PodSpec) SetTolerations(v []V1Toleration)`

SetTolerations sets Tolerations field to given value.

### HasTolerations

`func (o *V1PodSpec) HasTolerations() bool`

HasTolerations returns a boolean if a field has been set.

### GetVolumes

`func (o *V1PodSpec) GetVolumes() []V1Volume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *V1PodSpec) GetVolumesOk() (*[]V1Volume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *V1PodSpec) SetVolumes(v []V1Volume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *V1PodSpec) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


