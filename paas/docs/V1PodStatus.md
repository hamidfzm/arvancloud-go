# V1PodStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]V1PodCondition**](V1PodCondition.md) | Current service state of pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions | [optional] 
**ContainerStatuses** | Pointer to [**[]V1ContainerStatus**](V1ContainerStatus.md) | The list has one entry per container in the manifest. Each entry is currently the output of &#x60;docker inspect&#x60;. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status | [optional] 
**HostIP** | Pointer to **string** | IP address of the host to which the pod is assigned. Empty if not yet scheduled. | [optional] 
**InitContainerStatuses** | Pointer to [**[]V1ContainerStatus**](V1ContainerStatus.md) | The list has one entry per init container in the manifest. The most recent successful init container will have ready &#x3D; true, the most recently started container will have startTime set. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-and-container-status | [optional] 
**Message** | Pointer to **string** | A human readable message indicating details about why the pod is in this condition. | [optional] 
**NominatedNodeName** | Pointer to **string** | nominatedNodeName is set only when this pod preempts other pods on the node, but it cannot be scheduled right away as preemption victims receive their graceful termination periods. This field does not guarantee that the pod will be scheduled on this node. Scheduler may decide to place the pod elsewhere if other nodes become available sooner. Scheduler may also decide to give the resources on this node to a higher priority pod that is created after preemption. As a result, this field may be different than PodSpec.nodeName when the pod is scheduled. | [optional] 
**Phase** | Pointer to **string** | The phase of a Pod is a simple, high-level summary of where the Pod is in its lifecycle. The conditions array, the reason and message fields, and the individual container status arrays contain more detail about the pod&#39;s status. There are five possible phase values:  Pending: The pod has been accepted by the Kubernetes system, but one or more of the container images has not been created. This includes time before being scheduled as well as time spent downloading images over the network, which could take a while. Running: The pod has been bound to a node, and all of the containers have been created. At least one container is still running, or is in the process of starting or restarting. Succeeded: All containers in the pod have terminated in success, and will not be restarted. Failed: All containers in the pod have terminated, and at least one container has terminated in failure. The container either exited with non-zero status or was terminated by the system. Unknown: For some reason the state of the pod could not be obtained, typically due to an error in communicating with the host of the pod.  More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-phase | [optional] 
**PodIP** | Pointer to **string** | IP address allocated to the pod. Routable at least within the cluster. Empty if not yet allocated. | [optional] 
**QosClass** | Pointer to **string** | The Quality of Service (QOS) classification assigned to the pod based on resource requirements See PodQOSClass type for available QOS classes More info: https://git.k8s.io/community/contributors/design-proposals/node/resource-qos.md | [optional] 
**Reason** | Pointer to **string** | A brief CamelCase message indicating details about why the pod is in this state. e.g. &#39;Evicted&#39; | [optional] 
**StartTime** | Pointer to **string** | RFC 3339 date and time at which the object was acknowledged by the Kubelet. This is before the Kubelet pulled the container image(s) for the pod. | [optional] 

## Methods

### NewV1PodStatus

`func NewV1PodStatus() *V1PodStatus`

NewV1PodStatus instantiates a new V1PodStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PodStatusWithDefaults

`func NewV1PodStatusWithDefaults() *V1PodStatus`

NewV1PodStatusWithDefaults instantiates a new V1PodStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *V1PodStatus) GetConditions() []V1PodCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *V1PodStatus) GetConditionsOk() (*[]V1PodCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *V1PodStatus) SetConditions(v []V1PodCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *V1PodStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetContainerStatuses

`func (o *V1PodStatus) GetContainerStatuses() []V1ContainerStatus`

GetContainerStatuses returns the ContainerStatuses field if non-nil, zero value otherwise.

### GetContainerStatusesOk

`func (o *V1PodStatus) GetContainerStatusesOk() (*[]V1ContainerStatus, bool)`

GetContainerStatusesOk returns a tuple with the ContainerStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerStatuses

`func (o *V1PodStatus) SetContainerStatuses(v []V1ContainerStatus)`

SetContainerStatuses sets ContainerStatuses field to given value.

### HasContainerStatuses

`func (o *V1PodStatus) HasContainerStatuses() bool`

HasContainerStatuses returns a boolean if a field has been set.

### GetHostIP

`func (o *V1PodStatus) GetHostIP() string`

GetHostIP returns the HostIP field if non-nil, zero value otherwise.

### GetHostIPOk

`func (o *V1PodStatus) GetHostIPOk() (*string, bool)`

GetHostIPOk returns a tuple with the HostIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIP

`func (o *V1PodStatus) SetHostIP(v string)`

SetHostIP sets HostIP field to given value.

### HasHostIP

`func (o *V1PodStatus) HasHostIP() bool`

HasHostIP returns a boolean if a field has been set.

### GetInitContainerStatuses

`func (o *V1PodStatus) GetInitContainerStatuses() []V1ContainerStatus`

GetInitContainerStatuses returns the InitContainerStatuses field if non-nil, zero value otherwise.

### GetInitContainerStatusesOk

`func (o *V1PodStatus) GetInitContainerStatusesOk() (*[]V1ContainerStatus, bool)`

GetInitContainerStatusesOk returns a tuple with the InitContainerStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitContainerStatuses

`func (o *V1PodStatus) SetInitContainerStatuses(v []V1ContainerStatus)`

SetInitContainerStatuses sets InitContainerStatuses field to given value.

### HasInitContainerStatuses

`func (o *V1PodStatus) HasInitContainerStatuses() bool`

HasInitContainerStatuses returns a boolean if a field has been set.

### GetMessage

`func (o *V1PodStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1PodStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1PodStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1PodStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNominatedNodeName

`func (o *V1PodStatus) GetNominatedNodeName() string`

GetNominatedNodeName returns the NominatedNodeName field if non-nil, zero value otherwise.

### GetNominatedNodeNameOk

`func (o *V1PodStatus) GetNominatedNodeNameOk() (*string, bool)`

GetNominatedNodeNameOk returns a tuple with the NominatedNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNominatedNodeName

`func (o *V1PodStatus) SetNominatedNodeName(v string)`

SetNominatedNodeName sets NominatedNodeName field to given value.

### HasNominatedNodeName

`func (o *V1PodStatus) HasNominatedNodeName() bool`

HasNominatedNodeName returns a boolean if a field has been set.

### GetPhase

`func (o *V1PodStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *V1PodStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *V1PodStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *V1PodStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPodIP

`func (o *V1PodStatus) GetPodIP() string`

GetPodIP returns the PodIP field if non-nil, zero value otherwise.

### GetPodIPOk

`func (o *V1PodStatus) GetPodIPOk() (*string, bool)`

GetPodIPOk returns a tuple with the PodIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodIP

`func (o *V1PodStatus) SetPodIP(v string)`

SetPodIP sets PodIP field to given value.

### HasPodIP

`func (o *V1PodStatus) HasPodIP() bool`

HasPodIP returns a boolean if a field has been set.

### GetQosClass

`func (o *V1PodStatus) GetQosClass() string`

GetQosClass returns the QosClass field if non-nil, zero value otherwise.

### GetQosClassOk

`func (o *V1PodStatus) GetQosClassOk() (*string, bool)`

GetQosClassOk returns a tuple with the QosClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosClass

`func (o *V1PodStatus) SetQosClass(v string)`

SetQosClass sets QosClass field to given value.

### HasQosClass

`func (o *V1PodStatus) HasQosClass() bool`

HasQosClass returns a boolean if a field has been set.

### GetReason

`func (o *V1PodStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V1PodStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V1PodStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V1PodStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStartTime

`func (o *V1PodStatus) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *V1PodStatus) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *V1PodStatus) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *V1PodStatus) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


