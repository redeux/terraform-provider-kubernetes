---
page_title: "kubernetes_deployment Resource - terraform-provider-kubernetes"
subcategory: ""
description: |-
  
---

# Resource `kubernetes_deployment`





## Schema

### Required

- **metadata** (Block List, Min: 1, Max: 1) Standard deployment's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata (see [below for nested schema](#nestedblock--metadata))
- **spec** (Block List, Min: 1, Max: 1) Spec defines the specification of the desired behavior of the deployment. More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.9/#deployment-v1-apps (see [below for nested schema](#nestedblock--spec))

### Optional

- **id** (String, Optional) The ID of this resource.
- **timeouts** (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- **wait_for_rollout** (Boolean, Optional) Wait for the rollout of the deployment to complete. Defaults to true.

<a id="nestedblock--metadata"></a>
### Nested Schema for `metadata`

Optional:

- **annotations** (Map of String, Optional) An unstructured key value map stored with the deployment that may be used to store arbitrary metadata. More info: http://kubernetes.io/docs/user-guide/annotations
- **generate_name** (String, Optional) Prefix, used by the server, to generate a unique name ONLY IF the `name` field has not been provided. This value will also be combined with a unique suffix. Read more: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency
- **labels** (Map of String, Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the deployment. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels
- **name** (String, Optional) Name of the deployment, must be unique. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names
- **namespace** (String, Optional) Namespace defines the space within which name of the deployment must be unique.

Read-only:

- **generation** (Number, Read-only) A sequence number representing a specific generation of the desired state.
- **resource_version** (String, Read-only) An opaque value that represents the internal version of this deployment that can be used by clients to determine when deployment has changed. Read more: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
- **self_link** (String, Read-only) A URL representing this deployment.
- **uid** (String, Read-only) The unique in time and space value for this deployment. More info: http://kubernetes.io/docs/user-guide/identifiers#uids


<a id="nestedblock--spec"></a>
### Nested Schema for `spec`

Required:

- **template** (Block List, Min: 1, Max: 1) Template describes the pods that will be created. (see [below for nested schema](#nestedblock--spec--template))

Optional:

- **min_ready_seconds** (Number, Optional) Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
- **paused** (Boolean, Optional) Indicates that the deployment is paused.
- **progress_deadline_seconds** (Number, Optional) The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. Defaults to 600s.
- **replicas** (String, Optional) Number of desired pods. This is a string to be able to distinguish between explicit zero and not specified.
- **revision_history_limit** (Number, Optional) The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.
- **selector** (Block List, Max: 1) A label query over pods that should match the Replicas count. (see [below for nested schema](#nestedblock--spec--selector))
- **strategy** (Block List, Max: 1) The deployment strategy to use to replace existing pods with new ones. (see [below for nested schema](#nestedblock--spec--strategy))

<a id="nestedblock--spec--template"></a>
### Nested Schema for `spec.template`

Required:

- **metadata** (Block List, Min: 1, Max: 1) Standard pod's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata (see [below for nested schema](#nestedblock--spec--template--metadata))
- **spec** (Block List, Min: 1, Max: 1) Spec defines the specification of the desired behavior of the deployment. More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.9/#deployment-v1-apps (see [below for nested schema](#nestedblock--spec--template--spec))

<a id="nestedblock--spec--template--metadata"></a>
### Nested Schema for `spec.template.metadata`

Optional:

- **annotations** (Map of String, Optional) An unstructured key value map stored with the pod that may be used to store arbitrary metadata. More info: http://kubernetes.io/docs/user-guide/annotations
- **generate_name** (String, Optional) Prefix, used by the server, to generate a unique name ONLY IF the `name` field has not been provided. This value will also be combined with a unique suffix. Read more: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency
- **labels** (Map of String, Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the pod. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels
- **name** (String, Optional) Name of the pod, must be unique. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names
- **namespace** (String, Optional) Namespace defines the space within which name of the pod must be unique.

Read-only:

- **generation** (Number, Read-only) A sequence number representing a specific generation of the desired state.
- **resource_version** (String, Read-only) An opaque value that represents the internal version of this pod that can be used by clients to determine when pod has changed. Read more: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
- **self_link** (String, Read-only) A URL representing this pod.
- **uid** (String, Read-only) The unique in time and space value for this pod. More info: http://kubernetes.io/docs/user-guide/identifiers#uids


<a id="nestedblock--spec--template--spec"></a>
### Nested Schema for `spec.template.spec`

Optional:

- **active_deadline_seconds** (Number, Optional) Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.
- **affinity** (Block List, Max: 1) Optional pod scheduling constraints. (see [below for nested schema](#nestedblock--spec--template--spec--affinity))
- **automount_service_account_token** (Boolean, Optional) AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.
- **container** (Block List) List of containers belonging to the pod. Containers cannot currently be added or removed. There must be at least one container in a Pod. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/containers (see [below for nested schema](#nestedblock--spec--template--spec--container))
- **dns_config** (Block List, Max: 1) Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy. Optional: Defaults to empty (see [below for nested schema](#nestedblock--spec--template--spec--dns_config))
- **dns_policy** (String, Optional) Set DNS policy for containers within the pod. Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'. Optional: Defaults to 'ClusterFirst', see [Kubernetes reference](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-policy).
- **enable_service_links** (Boolean, Optional) Enables generating environment variables for service discovery. Optional: service links are enabled by default.
- **host_aliases** (Block List) List of hosts and IPs that will be injected into the pod's hosts file if specified. Optional: Defaults to empty. (see [below for nested schema](#nestedblock--spec--template--spec--host_aliases))
- **host_ipc** (Boolean, Optional) Use the host's ipc namespace. Optional: Defaults to false.
- **host_network** (Boolean, Optional) Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified.
- **host_pid** (Boolean, Optional) Use the host's pid namespace.
- **hostname** (String, Optional) Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a system-defined value.
- **image_pull_secrets** (Block List) ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. More info: http://kubernetes.io/docs/user-guide/images#specifying-imagepullsecrets-on-a-pod (see [below for nested schema](#nestedblock--spec--template--spec--image_pull_secrets))
- **init_container** (Block List) List of init containers belonging to the pod. Init containers always run to completion and each must complete successfully before the next is started. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/ (see [below for nested schema](#nestedblock--spec--template--spec--init_container))
- **node_name** (String, Optional) NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.
- **node_selector** (Map of String, Optional) NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: http://kubernetes.io/docs/user-guide/node-selection.
- **priority_class_name** (String, Optional) If specified, indicates the pod's priority. "system-node-critical" and "system-cluster-critical" are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.
- **readiness_gate** (Block List) If specified, all readiness gates will be evaluated for pod readiness. A pod is ready when all its containers are ready AND all conditions specified in the readiness gates have status equal to "True" More info: https://git.k8s.io/enhancements/keps/sig-network/0007-pod-ready%2B%2B.md (see [below for nested schema](#nestedblock--spec--template--spec--readiness_gate))
- **restart_policy** (String, Optional) Restart policy for all containers within the pod. One of Always, OnFailure, Never. More info: http://kubernetes.io/docs/user-guide/pod-states#restartpolicy.
- **security_context** (Block List, Max: 1) SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty (see [below for nested schema](#nestedblock--spec--template--spec--security_context))
- **service_account_name** (String, Optional) ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: http://releases.k8s.io/HEAD/docs/design/service_accounts.md.
- **share_process_namespace** (Boolean, Optional) Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set. Optional: Defaults to false.
- **subdomain** (String, Optional) If specified, the fully qualified Pod hostname will be "...svc.". If not specified, the pod will not have a domainname at all..
- **termination_grace_period_seconds** (Number, Optional) Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process.
- **toleration** (Block List) If specified, the pod's toleration. Optional: Defaults to empty (see [below for nested schema](#nestedblock--spec--template--spec--toleration))
- **volume** (Block List) List of volumes that can be mounted by containers belonging to the pod. More info: http://kubernetes.io/docs/user-guide/volumes (see [below for nested schema](#nestedblock--spec--template--spec--volume))

<a id="nestedblock--spec--template--spec--affinity"></a>
### Nested Schema for `spec.template.spec.volume`

Optional:

- **node_affinity** (Block List, Max: 1) Node affinity scheduling rules for the pod. (see [below for nested schema](#nestedblock--spec--template--spec--volume--node_affinity))
- **pod_affinity** (Block List, Max: 1) Inter-pod topological affinity. rules that specify that certain pods should be placed in the same topological domain (e.g. same node, same rack, same zone, same power domain, etc.) (see [below for nested schema](#nestedblock--spec--template--spec--volume--pod_affinity))
- **pod_anti_affinity** (Block List, Max: 1) Inter-pod topological affinity. rules that specify that certain pods should be placed in the same topological domain (e.g. same node, same rack, same zone, same power domain, etc.) (see [below for nested schema](#nestedblock--spec--template--spec--volume--pod_anti_affinity))

<a id="nestedblock--spec--template--spec--volume--node_affinity"></a>
### Nested Schema for `spec.template.spec.volume.node_affinity`

Optional:

- **preferred_during_scheduling_ignored_during_execution** (Block List) The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, RequiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding MatchExpressions; the node(s) with the highest sum are the most preferred. (see [below for nested schema](#nestedblock--spec--template--spec--volume--node_affinity--preferred_during_scheduling_ignored_during_execution))
- **required_during_scheduling_ignored_during_execution** (Block List, Max: 1) If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a node label update), the system may or may not try to eventually evict the pod from its node. (see [below for nested schema](#nestedblock--spec--template--spec--volume--node_affinity--required_during_scheduling_ignored_during_execution))

<a id="nestedblock--spec--template--spec--volume--node_affinity--preferred_during_scheduling_ignored_during_execution"></a>
### Nested Schema for `spec.template.spec.volume.node_affinity.required_during_scheduling_ignored_during_execution`

Required:

- **preference** (Block List, Min: 1, Max: 1) A node selector term, associated with the corresponding weight. (see [below for nested schema](#nestedblock--spec--template--spec--volume--node_affinity--required_during_scheduling_ignored_during_execution--preference))
- **weight** (Number, Required) weight is in the range 1-100

<a id="nestedblock--spec--template--spec--volume--node_affinity--required_during_scheduling_ignored_during_execution--preference"></a>
### Nested Schema for `spec.template.spec.volume.node_affinity.required_during_scheduling_ignored_during_execution.weight`

Optional:

- **match_expressions** (Block List) List of node selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedblock--spec--template--spec--volume--node_affinity--required_during_scheduling_ignored_during_execution--weight--match_expressions))

<a id="nestedblock--spec--template--spec--volume--node_affinity--required_during_scheduling_ignored_during_execution--weight--match_expressions"></a>
### Nested Schema for `spec.template.spec.volume.node_affinity.required_during_scheduling_ignored_during_execution.weight.match_expressions`

Optional:

- **key** (String, Optional) The label key that the selector applies to.
- **operator** (String, Optional) Operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
- **values** (Set of String, Optional) Values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.




<a id="nestedblock--spec--template--spec--volume--node_affinity--required_during_scheduling_ignored_during_execution"></a>
### Nested Schema for `spec.template.spec.volume.node_affinity.required_during_scheduling_ignored_during_execution`

Optional:

- **node_selector_term** (Block List) List of node selector terms. The terms are ORed. (see [below for nested schema](#nestedblock--spec--template--spec--volume--node_affinity--required_during_scheduling_ignored_during_execution--node_selector_term))

<a id="nestedblock--spec--template--spec--volume--node_affinity--required_during_scheduling_ignored_during_execution--node_selector_term"></a>
### Nested Schema for `spec.template.spec.volume.node_affinity.required_during_scheduling_ignored_during_execution.node_selector_term`

Optional:

- **match_expressions** (Block List) List of node selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedblock--spec--template--spec--volume--node_affinity--required_during_scheduling_ignored_during_execution--node_selector_term--match_expressions))

<a id="nestedblock--spec--template--spec--volume--node_affinity--required_during_scheduling_ignored_during_execution--node_selector_term--match_expressions"></a>
### Nested Schema for `spec.template.spec.volume.node_affinity.required_during_scheduling_ignored_during_execution.node_selector_term.match_expressions`

Optional:

- **key** (String, Optional) The label key that the selector applies to.
- **operator** (String, Optional) Operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
- **values** (Set of String, Optional) Values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.





<a id="nestedblock--spec--template--spec--volume--pod_affinity"></a>
### Nested Schema for `spec.template.spec.volume.pod_affinity`

Optional:

- **preferred_during_scheduling_ignored_during_execution** (Block List) The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, RequiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding MatchExpressions; the node(s) with the highest sum are the most preferred. (see [below for nested schema](#nestedblock--spec--template--spec--volume--pod_affinity--preferred_during_scheduling_ignored_during_execution))
- **required_during_scheduling_ignored_during_execution** (Block List) If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each PodAffinityTerm are intersected, i.e. all terms must be satisfied. (see [below for nested schema](#nestedblock--spec--template--spec--volume--pod_affinity--required_during_scheduling_ignored_during_execution))

<a id="nestedblock--spec--template--spec--volume--pod_affinity--preferred_during_scheduling_ignored_during_execution"></a>
### Nested Schema for `spec.template.spec.volume.pod_affinity.required_during_scheduling_ignored_during_execution`

Required:

- **pod_affinity_term** (Block List, Min: 1, Max: 1) A pod affinity term, associated with the corresponding weight (see [below for nested schema](#nestedblock--spec--template--spec--volume--pod_affinity--required_during_scheduling_ignored_during_execution--pod_affinity_term))
- **weight** (Number, Required) weight associated with matching the corresponding podAffinityTerm, in the range 1-100

<a id="nestedblock--spec--template--spec--volume--pod_affinity--required_during_scheduling_ignored_during_execution--pod_affinity_term"></a>
### Nested Schema for `spec.template.spec.volume.pod_affinity.required_during_scheduling_ignored_during_execution.weight`

Optional:

- **label_selector** (Block List) A label query over a set of resources, in this case pods. (see [below for nested schema](#nestedblock--spec--template--spec--volume--pod_affinity--required_during_scheduling_ignored_during_execution--weight--label_selector))
- **namespaces** (Set of String, Optional) namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'
- **topology_key** (String, Optional) empty topology key is interpreted by the scheduler as 'all topologies'

<a id="nestedblock--spec--template--spec--volume--pod_affinity--required_during_scheduling_ignored_during_execution--weight--label_selector"></a>
### Nested Schema for `spec.template.spec.volume.pod_affinity.required_during_scheduling_ignored_during_execution.weight.topology_key`

Optional:

- **match_expressions** (Block List) A list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedblock--spec--template--spec--volume--pod_affinity--required_during_scheduling_ignored_during_execution--weight--topology_key--match_expressions))
- **match_labels** (Map of String, Optional) A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of `match_expressions`, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.

<a id="nestedblock--spec--template--spec--volume--pod_affinity--required_during_scheduling_ignored_during_execution--weight--topology_key--match_expressions"></a>
### Nested Schema for `spec.template.spec.volume.pod_affinity.required_during_scheduling_ignored_during_execution.weight.topology_key.match_expressions`

Optional:

- **key** (String, Optional) The label key that the selector applies to.
- **operator** (String, Optional) A key's relationship to a set of values. Valid operators ard `In`, `NotIn`, `Exists` and `DoesNotExist`.
- **values** (Set of String, Optional) An array of string values. If the operator is `In` or `NotIn`, the values array must be non-empty. If the operator is `Exists` or `DoesNotExist`, the values array must be empty. This array is replaced during a strategic merge patch.





<a id="nestedblock--spec--template--spec--volume--pod_affinity--required_during_scheduling_ignored_during_execution"></a>
### Nested Schema for `spec.template.spec.volume.pod_affinity.required_during_scheduling_ignored_during_execution`

Optional:

- **label_selector** (Block List) A label query over a set of resources, in this case pods. (see [below for nested schema](#nestedblock--spec--template--spec--volume--pod_affinity--required_during_scheduling_ignored_during_execution--label_selector))
- **namespaces** (Set of String, Optional) namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'
- **topology_key** (String, Optional) empty topology key is interpreted by the scheduler as 'all topologies'

<a id="nestedblock--spec--template--spec--volume--pod_affinity--required_during_scheduling_ignored_during_execution--label_selector"></a>
### Nested Schema for `spec.template.spec.volume.pod_affinity.required_during_scheduling_ignored_during_execution.topology_key`

Optional:

- **match_expressions** (Block List) A list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedblock--spec--template--spec--volume--pod_affinity--required_during_scheduling_ignored_during_execution--topology_key--match_expressions))
- **match_labels** (Map of String, Optional) A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of `match_expressions`, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.

<a id="nestedblock--spec--template--spec--volume--pod_affinity--required_during_scheduling_ignored_during_execution--topology_key--match_expressions"></a>
### Nested Schema for `spec.template.spec.volume.pod_affinity.required_during_scheduling_ignored_during_execution.topology_key.match_labels`

Optional:

- **key** (String, Optional) The label key that the selector applies to.
- **operator** (String, Optional) A key's relationship to a set of values. Valid operators ard `In`, `NotIn`, `Exists` and `DoesNotExist`.
- **values** (Set of String, Optional) An array of string values. If the operator is `In` or `NotIn`, the values array must be non-empty. If the operator is `Exists` or `DoesNotExist`, the values array must be empty. This array is replaced during a strategic merge patch.





<a id="nestedblock--spec--template--spec--volume--pod_anti_affinity"></a>
### Nested Schema for `spec.template.spec.volume.pod_anti_affinity`

Optional:

- **preferred_during_scheduling_ignored_during_execution** (Block List) The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, RequiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding MatchExpressions; the node(s) with the highest sum are the most preferred. (see [below for nested schema](#nestedblock--spec--template--spec--volume--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution))
- **required_during_scheduling_ignored_during_execution** (Block List) If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each PodAffinityTerm are intersected, i.e. all terms must be satisfied. (see [below for nested schema](#nestedblock--spec--template--spec--volume--pod_anti_affinity--required_during_scheduling_ignored_during_execution))

<a id="nestedblock--spec--template--spec--volume--pod_anti_affinity--preferred_during_scheduling_ignored_during_execution"></a>
### Nested Schema for `spec.template.spec.volume.pod_anti_affinity.required_during_scheduling_ignored_during_execution`

Required:

- **pod_affinity_term** (Block List, Min: 1, Max: 1) A pod affinity term, associated with the corresponding weight (see [below for nested schema](#nestedblock--spec--template--spec--volume--pod_anti_affinity--required_during_scheduling_ignored_during_execution--pod_affinity_term))
- **weight** (Number, Required) weight associated with matching the corresponding podAffinityTerm, in the range 1-100

<a id="nestedblock--spec--template--spec--volume--pod_anti_affinity--required_during_scheduling_ignored_during_execution--pod_affinity_term"></a>
### Nested Schema for `spec.template.spec.volume.pod_anti_affinity.required_during_scheduling_ignored_during_execution.weight`

Optional:

- **label_selector** (Block List) A label query over a set of resources, in this case pods. (see [below for nested schema](#nestedblock--spec--template--spec--volume--pod_anti_affinity--required_during_scheduling_ignored_during_execution--weight--label_selector))
- **namespaces** (Set of String, Optional) namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'
- **topology_key** (String, Optional) empty topology key is interpreted by the scheduler as 'all topologies'

<a id="nestedblock--spec--template--spec--volume--pod_anti_affinity--required_during_scheduling_ignored_during_execution--weight--label_selector"></a>
### Nested Schema for `spec.template.spec.volume.pod_anti_affinity.required_during_scheduling_ignored_during_execution.weight.topology_key`

Optional:

- **match_expressions** (Block List) A list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedblock--spec--template--spec--volume--pod_anti_affinity--required_during_scheduling_ignored_during_execution--weight--topology_key--match_expressions))
- **match_labels** (Map of String, Optional) A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of `match_expressions`, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.

<a id="nestedblock--spec--template--spec--volume--pod_anti_affinity--required_during_scheduling_ignored_during_execution--weight--topology_key--match_expressions"></a>
### Nested Schema for `spec.template.spec.volume.pod_anti_affinity.required_during_scheduling_ignored_during_execution.weight.topology_key.match_expressions`

Optional:

- **key** (String, Optional) The label key that the selector applies to.
- **operator** (String, Optional) A key's relationship to a set of values. Valid operators ard `In`, `NotIn`, `Exists` and `DoesNotExist`.
- **values** (Set of String, Optional) An array of string values. If the operator is `In` or `NotIn`, the values array must be non-empty. If the operator is `Exists` or `DoesNotExist`, the values array must be empty. This array is replaced during a strategic merge patch.





<a id="nestedblock--spec--template--spec--volume--pod_anti_affinity--required_during_scheduling_ignored_during_execution"></a>
### Nested Schema for `spec.template.spec.volume.pod_anti_affinity.required_during_scheduling_ignored_during_execution`

Optional:

- **label_selector** (Block List) A label query over a set of resources, in this case pods. (see [below for nested schema](#nestedblock--spec--template--spec--volume--pod_anti_affinity--required_during_scheduling_ignored_during_execution--label_selector))
- **namespaces** (Set of String, Optional) namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'
- **topology_key** (String, Optional) empty topology key is interpreted by the scheduler as 'all topologies'

<a id="nestedblock--spec--template--spec--volume--pod_anti_affinity--required_during_scheduling_ignored_during_execution--label_selector"></a>
### Nested Schema for `spec.template.spec.volume.pod_anti_affinity.required_during_scheduling_ignored_during_execution.topology_key`

Optional:

- **match_expressions** (Block List) A list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedblock--spec--template--spec--volume--pod_anti_affinity--required_during_scheduling_ignored_during_execution--topology_key--match_expressions))
- **match_labels** (Map of String, Optional) A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of `match_expressions`, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.

<a id="nestedblock--spec--template--spec--volume--pod_anti_affinity--required_during_scheduling_ignored_during_execution--topology_key--match_expressions"></a>
### Nested Schema for `spec.template.spec.volume.pod_anti_affinity.required_during_scheduling_ignored_during_execution.topology_key.match_labels`

Optional:

- **key** (String, Optional) The label key that the selector applies to.
- **operator** (String, Optional) A key's relationship to a set of values. Valid operators ard `In`, `NotIn`, `Exists` and `DoesNotExist`.
- **values** (Set of String, Optional) An array of string values. If the operator is `In` or `NotIn`, the values array must be non-empty. If the operator is `Exists` or `DoesNotExist`, the values array must be empty. This array is replaced during a strategic merge patch.






<a id="nestedblock--spec--template--spec--container"></a>
### Nested Schema for `spec.template.spec.volume`

Required:

- **name** (String, Required) Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.

Optional:

- **args** (List of String, Optional) Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/containers#containers-and-commands
- **command** (List of String, Optional) Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/containers#containers-and-commands
- **env** (Block List) List of environment variables to set in the container. Cannot be updated. (see [below for nested schema](#nestedblock--spec--template--spec--volume--env))
- **env_from** (Block List) List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated. (see [below for nested schema](#nestedblock--spec--template--spec--volume--env_from))
- **image** (String, Optional) Docker image name. More info: http://kubernetes.io/docs/user-guide/images
- **image_pull_policy** (String, Optional) Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/images#updating-images
- **lifecycle** (Block List, Max: 1) Actions that the management system should take in response to container lifecycle events (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle))
- **liveness_probe** (Block List, Max: 1) Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes (see [below for nested schema](#nestedblock--spec--template--spec--volume--liveness_probe))
- **port** (Block List) List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default "0.0.0.0" address inside a container will be accessible from the network. Cannot be updated. (see [below for nested schema](#nestedblock--spec--template--spec--volume--port))
- **readiness_probe** (Block List, Max: 1) Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes (see [below for nested schema](#nestedblock--spec--template--spec--volume--readiness_probe))
- **resources** (Block List, Max: 1) Compute Resources required by this container. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#resources (see [below for nested schema](#nestedblock--spec--template--spec--volume--resources))
- **security_context** (Block List, Max: 1) Security options the pod should run with. More info: http://releases.k8s.io/HEAD/docs/design/security_context.md (see [below for nested schema](#nestedblock--spec--template--spec--volume--security_context))
- **startup_probe** (Block List, Max: 1) StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. This is an alpha feature enabled by the StartupProbe feature flag. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes (see [below for nested schema](#nestedblock--spec--template--spec--volume--startup_probe))
- **stdin** (Boolean, Optional) Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF.
- **stdin_once** (Boolean, Optional) Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF.
- **termination_message_path** (String, Optional) Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Defaults to /dev/termination-log. Cannot be updated.
- **termination_message_policy** (String, Optional) Optional: Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.
- **tty** (Boolean, Optional) Whether this container should allocate a TTY for itself
- **volume_mount** (Block List) Pod volumes to mount into the container's filesystem. Cannot be updated. (see [below for nested schema](#nestedblock--spec--template--spec--volume--volume_mount))
- **working_dir** (String, Optional) Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.

<a id="nestedblock--spec--template--spec--volume--env"></a>
### Nested Schema for `spec.template.spec.volume.env`

Required:

- **name** (String, Required) Name of the environment variable. Must be a C_IDENTIFIER

Optional:

- **value** (String, Optional) Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "".
- **value_from** (Block List, Max: 1) Source for the environment variable's value (see [below for nested schema](#nestedblock--spec--template--spec--volume--env--value_from))

<a id="nestedblock--spec--template--spec--volume--env--value_from"></a>
### Nested Schema for `spec.template.spec.volume.env.value_from`

Optional:

- **config_map_key_ref** (Block List, Max: 1) Selects a key of a ConfigMap. (see [below for nested schema](#nestedblock--spec--template--spec--volume--env--value_from--config_map_key_ref))
- **field_ref** (Block List, Max: 1) Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.podIP. (see [below for nested schema](#nestedblock--spec--template--spec--volume--env--value_from--field_ref))
- **resource_field_ref** (Block List, Max: 1) Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported. (see [below for nested schema](#nestedblock--spec--template--spec--volume--env--value_from--resource_field_ref))
- **secret_key_ref** (Block List, Max: 1) Selects a key of a secret in the pod's namespace. (see [below for nested schema](#nestedblock--spec--template--spec--volume--env--value_from--secret_key_ref))

<a id="nestedblock--spec--template--spec--volume--env--value_from--config_map_key_ref"></a>
### Nested Schema for `spec.template.spec.volume.env.value_from.secret_key_ref`

Optional:

- **key** (String, Optional) The key to select.
- **name** (String, Optional) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
- **optional** (Boolean, Optional) Specify whether the ConfigMap or its key must be defined.


<a id="nestedblock--spec--template--spec--volume--env--value_from--field_ref"></a>
### Nested Schema for `spec.template.spec.volume.env.value_from.secret_key_ref`

Optional:

- **api_version** (String, Optional) Version of the schema the FieldPath is written in terms of, defaults to "v1".
- **field_path** (String, Optional) Path of the field to select in the specified API version


<a id="nestedblock--spec--template--spec--volume--env--value_from--resource_field_ref"></a>
### Nested Schema for `spec.template.spec.volume.env.value_from.secret_key_ref`

Required:

- **resource** (String, Required) Resource to select

Optional:

- **container_name** (String, Optional)
- **divisor** (String, Optional)


<a id="nestedblock--spec--template--spec--volume--env--value_from--secret_key_ref"></a>
### Nested Schema for `spec.template.spec.volume.env.value_from.secret_key_ref`

Optional:

- **key** (String, Optional) The key of the secret to select from. Must be a valid secret key.
- **name** (String, Optional) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
- **optional** (Boolean, Optional) Specify whether the Secret or its key must be defined.




<a id="nestedblock--spec--template--spec--volume--env_from"></a>
### Nested Schema for `spec.template.spec.volume.env_from`

Optional:

- **config_map_ref** (Block List, Max: 1) The ConfigMap to select from (see [below for nested schema](#nestedblock--spec--template--spec--volume--env_from--config_map_ref))
- **prefix** (String, Optional) An optional identifer to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
- **secret_ref** (Block List, Max: 1) The Secret to select from (see [below for nested schema](#nestedblock--spec--template--spec--volume--env_from--secret_ref))

<a id="nestedblock--spec--template--spec--volume--env_from--config_map_ref"></a>
### Nested Schema for `spec.template.spec.volume.env_from.secret_ref`

Required:

- **name** (String, Required) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names

Optional:

- **optional** (Boolean, Optional) Specify whether the ConfigMap must be defined


<a id="nestedblock--spec--template--spec--volume--env_from--secret_ref"></a>
### Nested Schema for `spec.template.spec.volume.env_from.secret_ref`

Required:

- **name** (String, Required) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names

Optional:

- **optional** (Boolean, Optional) Specify whether the Secret must be defined



<a id="nestedblock--spec--template--spec--volume--lifecycle"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle`

Optional:

- **post_start** (Block List) post_start is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: http://kubernetes.io/docs/user-guide/container-environment#hook-details (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle--post_start))
- **pre_stop** (Block List) pre_stop is called immediately before a container is terminated. The container is terminated after the handler completes. The reason for termination is passed to the handler. Regardless of the outcome of the handler, the container is eventually terminated. Other management of the container blocks until the hook completes. More info: http://kubernetes.io/docs/user-guide/container-environment#hook-details (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle--pre_stop))

<a id="nestedblock--spec--template--spec--volume--lifecycle--post_start"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle.pre_stop`

Optional:

- **exec** (Block List, Max: 1) exec specifies the action to take. (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle--pre_stop--exec))
- **http_get** (Block List, Max: 1) Specifies the http request to perform. (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle--pre_stop--http_get))
- **tcp_socket** (Block List) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle--pre_stop--tcp_socket))

<a id="nestedblock--spec--template--spec--volume--lifecycle--pre_stop--exec"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle.pre_stop.tcp_socket`

Optional:

- **command** (List of String, Optional) Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.


<a id="nestedblock--spec--template--spec--volume--lifecycle--pre_stop--http_get"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle.pre_stop.tcp_socket`

Optional:

- **host** (String, Optional) Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
- **http_header** (Block List) Scheme to use for connecting to the host. (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle--pre_stop--tcp_socket--http_header))
- **path** (String, Optional) Path to access on the HTTP server.
- **port** (String, Optional) Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
- **scheme** (String, Optional) Scheme to use for connecting to the host.

<a id="nestedblock--spec--template--spec--volume--lifecycle--pre_stop--tcp_socket--http_header"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle.pre_stop.tcp_socket.scheme`

Optional:

- **name** (String, Optional) The header field name
- **value** (String, Optional) The header field value



<a id="nestedblock--spec--template--spec--volume--lifecycle--pre_stop--tcp_socket"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle.pre_stop.tcp_socket`

Required:

- **port** (String, Required) Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.



<a id="nestedblock--spec--template--spec--volume--lifecycle--pre_stop"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle.pre_stop`

Optional:

- **exec** (Block List, Max: 1) exec specifies the action to take. (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle--pre_stop--exec))
- **http_get** (Block List, Max: 1) Specifies the http request to perform. (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle--pre_stop--http_get))
- **tcp_socket** (Block List) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle--pre_stop--tcp_socket))

<a id="nestedblock--spec--template--spec--volume--lifecycle--pre_stop--exec"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle.pre_stop.tcp_socket`

Optional:

- **command** (List of String, Optional) Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.


<a id="nestedblock--spec--template--spec--volume--lifecycle--pre_stop--http_get"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle.pre_stop.tcp_socket`

Optional:

- **host** (String, Optional) Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
- **http_header** (Block List) Scheme to use for connecting to the host. (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle--pre_stop--tcp_socket--http_header))
- **path** (String, Optional) Path to access on the HTTP server.
- **port** (String, Optional) Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
- **scheme** (String, Optional) Scheme to use for connecting to the host.

<a id="nestedblock--spec--template--spec--volume--lifecycle--pre_stop--tcp_socket--http_header"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle.pre_stop.tcp_socket.scheme`

Optional:

- **name** (String, Optional) The header field name
- **value** (String, Optional) The header field value



<a id="nestedblock--spec--template--spec--volume--lifecycle--pre_stop--tcp_socket"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle.pre_stop.tcp_socket`

Required:

- **port** (String, Required) Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.




<a id="nestedblock--spec--template--spec--volume--liveness_probe"></a>
### Nested Schema for `spec.template.spec.volume.liveness_probe`

Optional:

- **exec** (Block List, Max: 1) exec specifies the action to take. (see [below for nested schema](#nestedblock--spec--template--spec--volume--liveness_probe--exec))
- **failure_threshold** (Number, Optional) Minimum consecutive failures for the probe to be considered failed after having succeeded.
- **http_get** (Block List, Max: 1) Specifies the http request to perform. (see [below for nested schema](#nestedblock--spec--template--spec--volume--liveness_probe--http_get))
- **initial_delay_seconds** (Number, Optional) Number of seconds after the container has started before liveness probes are initiated. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes
- **period_seconds** (Number, Optional) How often (in seconds) to perform the probe
- **success_threshold** (Number, Optional) Minimum consecutive successes for the probe to be considered successful after having failed.
- **tcp_socket** (Block List) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported (see [below for nested schema](#nestedblock--spec--template--spec--volume--liveness_probe--tcp_socket))
- **timeout_seconds** (Number, Optional) Number of seconds after which the probe times out. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes

<a id="nestedblock--spec--template--spec--volume--liveness_probe--exec"></a>
### Nested Schema for `spec.template.spec.volume.liveness_probe.timeout_seconds`

Optional:

- **command** (List of String, Optional) Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.


<a id="nestedblock--spec--template--spec--volume--liveness_probe--http_get"></a>
### Nested Schema for `spec.template.spec.volume.liveness_probe.timeout_seconds`

Optional:

- **host** (String, Optional) Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
- **http_header** (Block List) Scheme to use for connecting to the host. (see [below for nested schema](#nestedblock--spec--template--spec--volume--liveness_probe--timeout_seconds--http_header))
- **path** (String, Optional) Path to access on the HTTP server.
- **port** (String, Optional) Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
- **scheme** (String, Optional) Scheme to use for connecting to the host.

<a id="nestedblock--spec--template--spec--volume--liveness_probe--timeout_seconds--http_header"></a>
### Nested Schema for `spec.template.spec.volume.liveness_probe.timeout_seconds.scheme`

Optional:

- **name** (String, Optional) The header field name
- **value** (String, Optional) The header field value



<a id="nestedblock--spec--template--spec--volume--liveness_probe--tcp_socket"></a>
### Nested Schema for `spec.template.spec.volume.liveness_probe.timeout_seconds`

Required:

- **port** (String, Required) Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.



<a id="nestedblock--spec--template--spec--volume--port"></a>
### Nested Schema for `spec.template.spec.volume.port`

Required:

- **container_port** (Number, Required) Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.

Optional:

- **host_ip** (String, Optional) What host IP to bind the external port to.
- **host_port** (Number, Optional) Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.
- **name** (String, Optional) If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services
- **protocol** (String, Optional) Protocol for port. Must be UDP or TCP. Defaults to "TCP".


<a id="nestedblock--spec--template--spec--volume--readiness_probe"></a>
### Nested Schema for `spec.template.spec.volume.readiness_probe`

Optional:

- **exec** (Block List, Max: 1) exec specifies the action to take. (see [below for nested schema](#nestedblock--spec--template--spec--volume--readiness_probe--exec))
- **failure_threshold** (Number, Optional) Minimum consecutive failures for the probe to be considered failed after having succeeded.
- **http_get** (Block List, Max: 1) Specifies the http request to perform. (see [below for nested schema](#nestedblock--spec--template--spec--volume--readiness_probe--http_get))
- **initial_delay_seconds** (Number, Optional) Number of seconds after the container has started before liveness probes are initiated. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes
- **period_seconds** (Number, Optional) How often (in seconds) to perform the probe
- **success_threshold** (Number, Optional) Minimum consecutive successes for the probe to be considered successful after having failed.
- **tcp_socket** (Block List) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported (see [below for nested schema](#nestedblock--spec--template--spec--volume--readiness_probe--tcp_socket))
- **timeout_seconds** (Number, Optional) Number of seconds after which the probe times out. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes

<a id="nestedblock--spec--template--spec--volume--readiness_probe--exec"></a>
### Nested Schema for `spec.template.spec.volume.readiness_probe.timeout_seconds`

Optional:

- **command** (List of String, Optional) Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.


<a id="nestedblock--spec--template--spec--volume--readiness_probe--http_get"></a>
### Nested Schema for `spec.template.spec.volume.readiness_probe.timeout_seconds`

Optional:

- **host** (String, Optional) Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
- **http_header** (Block List) Scheme to use for connecting to the host. (see [below for nested schema](#nestedblock--spec--template--spec--volume--readiness_probe--timeout_seconds--http_header))
- **path** (String, Optional) Path to access on the HTTP server.
- **port** (String, Optional) Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
- **scheme** (String, Optional) Scheme to use for connecting to the host.

<a id="nestedblock--spec--template--spec--volume--readiness_probe--timeout_seconds--http_header"></a>
### Nested Schema for `spec.template.spec.volume.readiness_probe.timeout_seconds.scheme`

Optional:

- **name** (String, Optional) The header field name
- **value** (String, Optional) The header field value



<a id="nestedblock--spec--template--spec--volume--readiness_probe--tcp_socket"></a>
### Nested Schema for `spec.template.spec.volume.readiness_probe.timeout_seconds`

Required:

- **port** (String, Required) Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.



<a id="nestedblock--spec--template--spec--volume--resources"></a>
### Nested Schema for `spec.template.spec.volume.resources`

Optional:

- **limits** (Map of String, Optional) Describes the maximum amount of compute resources allowed. More info: http://kubernetes.io/docs/user-guide/compute-resources/
- **requests** (Map of String, Optional) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/


<a id="nestedblock--spec--template--spec--volume--security_context"></a>
### Nested Schema for `spec.template.spec.volume.security_context`

Optional:

- **allow_privilege_escalation** (Boolean, Optional) AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN
- **capabilities** (Block List, Max: 1) The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime. (see [below for nested schema](#nestedblock--spec--template--spec--volume--security_context--capabilities))
- **privileged** (Boolean, Optional) Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.
- **read_only_root_filesystem** (Boolean, Optional) Whether this container has a read-only root filesystem. Default is false.
- **run_as_group** (Number, Optional) The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
- **run_as_non_root** (Boolean, Optional) Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
- **run_as_user** (Number, Optional) The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
- **se_linux_options** (Block List, Max: 1) The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. (see [below for nested schema](#nestedblock--spec--template--spec--volume--security_context--se_linux_options))

<a id="nestedblock--spec--template--spec--volume--security_context--capabilities"></a>
### Nested Schema for `spec.template.spec.volume.security_context.se_linux_options`

Optional:

- **add** (List of String, Optional) Added capabilities
- **drop** (List of String, Optional) Removed capabilities


<a id="nestedblock--spec--template--spec--volume--security_context--se_linux_options"></a>
### Nested Schema for `spec.template.spec.volume.security_context.se_linux_options`

Optional:

- **level** (String, Optional) Level is SELinux level label that applies to the container.
- **role** (String, Optional) Role is a SELinux role label that applies to the container.
- **type** (String, Optional) Type is a SELinux type label that applies to the container.
- **user** (String, Optional) User is a SELinux user label that applies to the container.



<a id="nestedblock--spec--template--spec--volume--startup_probe"></a>
### Nested Schema for `spec.template.spec.volume.startup_probe`

Optional:

- **exec** (Block List, Max: 1) exec specifies the action to take. (see [below for nested schema](#nestedblock--spec--template--spec--volume--startup_probe--exec))
- **failure_threshold** (Number, Optional) Minimum consecutive failures for the probe to be considered failed after having succeeded.
- **http_get** (Block List, Max: 1) Specifies the http request to perform. (see [below for nested schema](#nestedblock--spec--template--spec--volume--startup_probe--http_get))
- **initial_delay_seconds** (Number, Optional) Number of seconds after the container has started before liveness probes are initiated. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes
- **period_seconds** (Number, Optional) How often (in seconds) to perform the probe
- **success_threshold** (Number, Optional) Minimum consecutive successes for the probe to be considered successful after having failed.
- **tcp_socket** (Block List) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported (see [below for nested schema](#nestedblock--spec--template--spec--volume--startup_probe--tcp_socket))
- **timeout_seconds** (Number, Optional) Number of seconds after which the probe times out. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes

<a id="nestedblock--spec--template--spec--volume--startup_probe--exec"></a>
### Nested Schema for `spec.template.spec.volume.startup_probe.timeout_seconds`

Optional:

- **command** (List of String, Optional) Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.


<a id="nestedblock--spec--template--spec--volume--startup_probe--http_get"></a>
### Nested Schema for `spec.template.spec.volume.startup_probe.timeout_seconds`

Optional:

- **host** (String, Optional) Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
- **http_header** (Block List) Scheme to use for connecting to the host. (see [below for nested schema](#nestedblock--spec--template--spec--volume--startup_probe--timeout_seconds--http_header))
- **path** (String, Optional) Path to access on the HTTP server.
- **port** (String, Optional) Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
- **scheme** (String, Optional) Scheme to use for connecting to the host.

<a id="nestedblock--spec--template--spec--volume--startup_probe--timeout_seconds--http_header"></a>
### Nested Schema for `spec.template.spec.volume.startup_probe.timeout_seconds.scheme`

Optional:

- **name** (String, Optional) The header field name
- **value** (String, Optional) The header field value



<a id="nestedblock--spec--template--spec--volume--startup_probe--tcp_socket"></a>
### Nested Schema for `spec.template.spec.volume.startup_probe.timeout_seconds`

Required:

- **port** (String, Required) Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.



<a id="nestedblock--spec--template--spec--volume--volume_mount"></a>
### Nested Schema for `spec.template.spec.volume.volume_mount`

Required:

- **mount_path** (String, Required) Path within the container at which the volume should be mounted. Must not contain ':'.
- **name** (String, Required) This must match the Name of a Volume.

Optional:

- **mount_propagation** (String, Optional) Mount propagation mode. mount_propagation determines how mounts are propagated from the host to container and the other way around. Valid values are None (default), HostToContainer and Bidirectional.
- **read_only** (Boolean, Optional) Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.
- **sub_path** (String, Optional) Path within the volume from which the container's volume should be mounted. Defaults to "" (volume's root).



<a id="nestedblock--spec--template--spec--dns_config"></a>
### Nested Schema for `spec.template.spec.volume`

Optional:

- **nameservers** (List of String, Optional) A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.
- **option** (Block List) A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy. (see [below for nested schema](#nestedblock--spec--template--spec--volume--option))
- **searches** (List of String, Optional) A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.

<a id="nestedblock--spec--template--spec--volume--option"></a>
### Nested Schema for `spec.template.spec.volume.option`

Required:

- **name** (String, Required) Name of the option.

Optional:

- **value** (String, Optional) Value of the option. Optional: Defaults to empty.



<a id="nestedblock--spec--template--spec--host_aliases"></a>
### Nested Schema for `spec.template.spec.volume`

Required:

- **hostnames** (List of String, Required) Hostnames for the IP address.
- **ip** (String, Required) IP address of the host file entry.


<a id="nestedblock--spec--template--spec--image_pull_secrets"></a>
### Nested Schema for `spec.template.spec.volume`

Required:

- **name** (String, Required) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names


<a id="nestedblock--spec--template--spec--init_container"></a>
### Nested Schema for `spec.template.spec.volume`

Required:

- **name** (String, Required) Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.

Optional:

- **args** (List of String, Optional) Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/containers#containers-and-commands
- **command** (List of String, Optional) Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/containers#containers-and-commands
- **env** (Block List) List of environment variables to set in the container. Cannot be updated. (see [below for nested schema](#nestedblock--spec--template--spec--volume--env))
- **env_from** (Block List) List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated. (see [below for nested schema](#nestedblock--spec--template--spec--volume--env_from))
- **image** (String, Optional) Docker image name. More info: http://kubernetes.io/docs/user-guide/images
- **image_pull_policy** (String, Optional) Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/images#updating-images
- **lifecycle** (Block List, Max: 1) Actions that the management system should take in response to container lifecycle events (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle))
- **liveness_probe** (Block List, Max: 1) Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes (see [below for nested schema](#nestedblock--spec--template--spec--volume--liveness_probe))
- **port** (Block List) List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default "0.0.0.0" address inside a container will be accessible from the network. Cannot be updated. (see [below for nested schema](#nestedblock--spec--template--spec--volume--port))
- **readiness_probe** (Block List, Max: 1) Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes (see [below for nested schema](#nestedblock--spec--template--spec--volume--readiness_probe))
- **resources** (Block List, Max: 1) Compute Resources required by this container. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#resources (see [below for nested schema](#nestedblock--spec--template--spec--volume--resources))
- **security_context** (Block List, Max: 1) Security options the pod should run with. More info: http://releases.k8s.io/HEAD/docs/design/security_context.md (see [below for nested schema](#nestedblock--spec--template--spec--volume--security_context))
- **startup_probe** (Block List, Max: 1) StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. This is an alpha feature enabled by the StartupProbe feature flag. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes (see [below for nested schema](#nestedblock--spec--template--spec--volume--startup_probe))
- **stdin** (Boolean, Optional) Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF.
- **stdin_once** (Boolean, Optional) Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF.
- **termination_message_path** (String, Optional) Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Defaults to /dev/termination-log. Cannot be updated.
- **termination_message_policy** (String, Optional) Optional: Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.
- **tty** (Boolean, Optional) Whether this container should allocate a TTY for itself
- **volume_mount** (Block List) Pod volumes to mount into the container's filesystem. Cannot be updated. (see [below for nested schema](#nestedblock--spec--template--spec--volume--volume_mount))
- **working_dir** (String, Optional) Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.

<a id="nestedblock--spec--template--spec--volume--env"></a>
### Nested Schema for `spec.template.spec.volume.env`

Required:

- **name** (String, Required) Name of the environment variable. Must be a C_IDENTIFIER

Optional:

- **value** (String, Optional) Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "".
- **value_from** (Block List, Max: 1) Source for the environment variable's value (see [below for nested schema](#nestedblock--spec--template--spec--volume--env--value_from))

<a id="nestedblock--spec--template--spec--volume--env--value_from"></a>
### Nested Schema for `spec.template.spec.volume.env.value_from`

Optional:

- **config_map_key_ref** (Block List, Max: 1) Selects a key of a ConfigMap. (see [below for nested schema](#nestedblock--spec--template--spec--volume--env--value_from--config_map_key_ref))
- **field_ref** (Block List, Max: 1) Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.podIP. (see [below for nested schema](#nestedblock--spec--template--spec--volume--env--value_from--field_ref))
- **resource_field_ref** (Block List, Max: 1) Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported. (see [below for nested schema](#nestedblock--spec--template--spec--volume--env--value_from--resource_field_ref))
- **secret_key_ref** (Block List, Max: 1) Selects a key of a secret in the pod's namespace. (see [below for nested schema](#nestedblock--spec--template--spec--volume--env--value_from--secret_key_ref))

<a id="nestedblock--spec--template--spec--volume--env--value_from--config_map_key_ref"></a>
### Nested Schema for `spec.template.spec.volume.env.value_from.secret_key_ref`

Optional:

- **key** (String, Optional) The key to select.
- **name** (String, Optional) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
- **optional** (Boolean, Optional) Specify whether the ConfigMap or its key must be defined.


<a id="nestedblock--spec--template--spec--volume--env--value_from--field_ref"></a>
### Nested Schema for `spec.template.spec.volume.env.value_from.secret_key_ref`

Optional:

- **api_version** (String, Optional) Version of the schema the FieldPath is written in terms of, defaults to "v1".
- **field_path** (String, Optional) Path of the field to select in the specified API version


<a id="nestedblock--spec--template--spec--volume--env--value_from--resource_field_ref"></a>
### Nested Schema for `spec.template.spec.volume.env.value_from.secret_key_ref`

Required:

- **resource** (String, Required) Resource to select

Optional:

- **container_name** (String, Optional)
- **divisor** (String, Optional)


<a id="nestedblock--spec--template--spec--volume--env--value_from--secret_key_ref"></a>
### Nested Schema for `spec.template.spec.volume.env.value_from.secret_key_ref`

Optional:

- **key** (String, Optional) The key of the secret to select from. Must be a valid secret key.
- **name** (String, Optional) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
- **optional** (Boolean, Optional) Specify whether the Secret or its key must be defined.




<a id="nestedblock--spec--template--spec--volume--env_from"></a>
### Nested Schema for `spec.template.spec.volume.env_from`

Optional:

- **config_map_ref** (Block List, Max: 1) The ConfigMap to select from (see [below for nested schema](#nestedblock--spec--template--spec--volume--env_from--config_map_ref))
- **prefix** (String, Optional) An optional identifer to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
- **secret_ref** (Block List, Max: 1) The Secret to select from (see [below for nested schema](#nestedblock--spec--template--spec--volume--env_from--secret_ref))

<a id="nestedblock--spec--template--spec--volume--env_from--config_map_ref"></a>
### Nested Schema for `spec.template.spec.volume.env_from.secret_ref`

Required:

- **name** (String, Required) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names

Optional:

- **optional** (Boolean, Optional) Specify whether the ConfigMap must be defined


<a id="nestedblock--spec--template--spec--volume--env_from--secret_ref"></a>
### Nested Schema for `spec.template.spec.volume.env_from.secret_ref`

Required:

- **name** (String, Required) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names

Optional:

- **optional** (Boolean, Optional) Specify whether the Secret must be defined



<a id="nestedblock--spec--template--spec--volume--lifecycle"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle`

Optional:

- **post_start** (Block List) post_start is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: http://kubernetes.io/docs/user-guide/container-environment#hook-details (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle--post_start))
- **pre_stop** (Block List) pre_stop is called immediately before a container is terminated. The container is terminated after the handler completes. The reason for termination is passed to the handler. Regardless of the outcome of the handler, the container is eventually terminated. Other management of the container blocks until the hook completes. More info: http://kubernetes.io/docs/user-guide/container-environment#hook-details (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle--pre_stop))

<a id="nestedblock--spec--template--spec--volume--lifecycle--post_start"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle.pre_stop`

Optional:

- **exec** (Block List, Max: 1) exec specifies the action to take. (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle--pre_stop--exec))
- **http_get** (Block List, Max: 1) Specifies the http request to perform. (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle--pre_stop--http_get))
- **tcp_socket** (Block List) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle--pre_stop--tcp_socket))

<a id="nestedblock--spec--template--spec--volume--lifecycle--pre_stop--exec"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle.pre_stop.tcp_socket`

Optional:

- **command** (List of String, Optional) Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.


<a id="nestedblock--spec--template--spec--volume--lifecycle--pre_stop--http_get"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle.pre_stop.tcp_socket`

Optional:

- **host** (String, Optional) Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
- **http_header** (Block List) Scheme to use for connecting to the host. (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle--pre_stop--tcp_socket--http_header))
- **path** (String, Optional) Path to access on the HTTP server.
- **port** (String, Optional) Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
- **scheme** (String, Optional) Scheme to use for connecting to the host.

<a id="nestedblock--spec--template--spec--volume--lifecycle--pre_stop--tcp_socket--http_header"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle.pre_stop.tcp_socket.scheme`

Optional:

- **name** (String, Optional) The header field name
- **value** (String, Optional) The header field value



<a id="nestedblock--spec--template--spec--volume--lifecycle--pre_stop--tcp_socket"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle.pre_stop.tcp_socket`

Required:

- **port** (String, Required) Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.



<a id="nestedblock--spec--template--spec--volume--lifecycle--pre_stop"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle.pre_stop`

Optional:

- **exec** (Block List, Max: 1) exec specifies the action to take. (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle--pre_stop--exec))
- **http_get** (Block List, Max: 1) Specifies the http request to perform. (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle--pre_stop--http_get))
- **tcp_socket** (Block List) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle--pre_stop--tcp_socket))

<a id="nestedblock--spec--template--spec--volume--lifecycle--pre_stop--exec"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle.pre_stop.tcp_socket`

Optional:

- **command** (List of String, Optional) Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.


<a id="nestedblock--spec--template--spec--volume--lifecycle--pre_stop--http_get"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle.pre_stop.tcp_socket`

Optional:

- **host** (String, Optional) Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
- **http_header** (Block List) Scheme to use for connecting to the host. (see [below for nested schema](#nestedblock--spec--template--spec--volume--lifecycle--pre_stop--tcp_socket--http_header))
- **path** (String, Optional) Path to access on the HTTP server.
- **port** (String, Optional) Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
- **scheme** (String, Optional) Scheme to use for connecting to the host.

<a id="nestedblock--spec--template--spec--volume--lifecycle--pre_stop--tcp_socket--http_header"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle.pre_stop.tcp_socket.scheme`

Optional:

- **name** (String, Optional) The header field name
- **value** (String, Optional) The header field value



<a id="nestedblock--spec--template--spec--volume--lifecycle--pre_stop--tcp_socket"></a>
### Nested Schema for `spec.template.spec.volume.lifecycle.pre_stop.tcp_socket`

Required:

- **port** (String, Required) Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.




<a id="nestedblock--spec--template--spec--volume--liveness_probe"></a>
### Nested Schema for `spec.template.spec.volume.liveness_probe`

Optional:

- **exec** (Block List, Max: 1) exec specifies the action to take. (see [below for nested schema](#nestedblock--spec--template--spec--volume--liveness_probe--exec))
- **failure_threshold** (Number, Optional) Minimum consecutive failures for the probe to be considered failed after having succeeded.
- **http_get** (Block List, Max: 1) Specifies the http request to perform. (see [below for nested schema](#nestedblock--spec--template--spec--volume--liveness_probe--http_get))
- **initial_delay_seconds** (Number, Optional) Number of seconds after the container has started before liveness probes are initiated. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes
- **period_seconds** (Number, Optional) How often (in seconds) to perform the probe
- **success_threshold** (Number, Optional) Minimum consecutive successes for the probe to be considered successful after having failed.
- **tcp_socket** (Block List) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported (see [below for nested schema](#nestedblock--spec--template--spec--volume--liveness_probe--tcp_socket))
- **timeout_seconds** (Number, Optional) Number of seconds after which the probe times out. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes

<a id="nestedblock--spec--template--spec--volume--liveness_probe--exec"></a>
### Nested Schema for `spec.template.spec.volume.liveness_probe.timeout_seconds`

Optional:

- **command** (List of String, Optional) Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.


<a id="nestedblock--spec--template--spec--volume--liveness_probe--http_get"></a>
### Nested Schema for `spec.template.spec.volume.liveness_probe.timeout_seconds`

Optional:

- **host** (String, Optional) Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
- **http_header** (Block List) Scheme to use for connecting to the host. (see [below for nested schema](#nestedblock--spec--template--spec--volume--liveness_probe--timeout_seconds--http_header))
- **path** (String, Optional) Path to access on the HTTP server.
- **port** (String, Optional) Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
- **scheme** (String, Optional) Scheme to use for connecting to the host.

<a id="nestedblock--spec--template--spec--volume--liveness_probe--timeout_seconds--http_header"></a>
### Nested Schema for `spec.template.spec.volume.liveness_probe.timeout_seconds.scheme`

Optional:

- **name** (String, Optional) The header field name
- **value** (String, Optional) The header field value



<a id="nestedblock--spec--template--spec--volume--liveness_probe--tcp_socket"></a>
### Nested Schema for `spec.template.spec.volume.liveness_probe.timeout_seconds`

Required:

- **port** (String, Required) Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.



<a id="nestedblock--spec--template--spec--volume--port"></a>
### Nested Schema for `spec.template.spec.volume.port`

Required:

- **container_port** (Number, Required) Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.

Optional:

- **host_ip** (String, Optional) What host IP to bind the external port to.
- **host_port** (Number, Optional) Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.
- **name** (String, Optional) If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services
- **protocol** (String, Optional) Protocol for port. Must be UDP or TCP. Defaults to "TCP".


<a id="nestedblock--spec--template--spec--volume--readiness_probe"></a>
### Nested Schema for `spec.template.spec.volume.readiness_probe`

Optional:

- **exec** (Block List, Max: 1) exec specifies the action to take. (see [below for nested schema](#nestedblock--spec--template--spec--volume--readiness_probe--exec))
- **failure_threshold** (Number, Optional) Minimum consecutive failures for the probe to be considered failed after having succeeded.
- **http_get** (Block List, Max: 1) Specifies the http request to perform. (see [below for nested schema](#nestedblock--spec--template--spec--volume--readiness_probe--http_get))
- **initial_delay_seconds** (Number, Optional) Number of seconds after the container has started before liveness probes are initiated. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes
- **period_seconds** (Number, Optional) How often (in seconds) to perform the probe
- **success_threshold** (Number, Optional) Minimum consecutive successes for the probe to be considered successful after having failed.
- **tcp_socket** (Block List) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported (see [below for nested schema](#nestedblock--spec--template--spec--volume--readiness_probe--tcp_socket))
- **timeout_seconds** (Number, Optional) Number of seconds after which the probe times out. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes

<a id="nestedblock--spec--template--spec--volume--readiness_probe--exec"></a>
### Nested Schema for `spec.template.spec.volume.readiness_probe.timeout_seconds`

Optional:

- **command** (List of String, Optional) Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.


<a id="nestedblock--spec--template--spec--volume--readiness_probe--http_get"></a>
### Nested Schema for `spec.template.spec.volume.readiness_probe.timeout_seconds`

Optional:

- **host** (String, Optional) Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
- **http_header** (Block List) Scheme to use for connecting to the host. (see [below for nested schema](#nestedblock--spec--template--spec--volume--readiness_probe--timeout_seconds--http_header))
- **path** (String, Optional) Path to access on the HTTP server.
- **port** (String, Optional) Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
- **scheme** (String, Optional) Scheme to use for connecting to the host.

<a id="nestedblock--spec--template--spec--volume--readiness_probe--timeout_seconds--http_header"></a>
### Nested Schema for `spec.template.spec.volume.readiness_probe.timeout_seconds.scheme`

Optional:

- **name** (String, Optional) The header field name
- **value** (String, Optional) The header field value



<a id="nestedblock--spec--template--spec--volume--readiness_probe--tcp_socket"></a>
### Nested Schema for `spec.template.spec.volume.readiness_probe.timeout_seconds`

Required:

- **port** (String, Required) Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.



<a id="nestedblock--spec--template--spec--volume--resources"></a>
### Nested Schema for `spec.template.spec.volume.resources`

Optional:

- **limits** (Map of String, Optional) Describes the maximum amount of compute resources allowed. More info: http://kubernetes.io/docs/user-guide/compute-resources/
- **requests** (Map of String, Optional) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/


<a id="nestedblock--spec--template--spec--volume--security_context"></a>
### Nested Schema for `spec.template.spec.volume.security_context`

Optional:

- **allow_privilege_escalation** (Boolean, Optional) AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN
- **capabilities** (Block List, Max: 1) The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime. (see [below for nested schema](#nestedblock--spec--template--spec--volume--security_context--capabilities))
- **privileged** (Boolean, Optional) Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.
- **read_only_root_filesystem** (Boolean, Optional) Whether this container has a read-only root filesystem. Default is false.
- **run_as_group** (Number, Optional) The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
- **run_as_non_root** (Boolean, Optional) Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
- **run_as_user** (Number, Optional) The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
- **se_linux_options** (Block List, Max: 1) The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. (see [below for nested schema](#nestedblock--spec--template--spec--volume--security_context--se_linux_options))

<a id="nestedblock--spec--template--spec--volume--security_context--capabilities"></a>
### Nested Schema for `spec.template.spec.volume.security_context.se_linux_options`

Optional:

- **add** (List of String, Optional) Added capabilities
- **drop** (List of String, Optional) Removed capabilities


<a id="nestedblock--spec--template--spec--volume--security_context--se_linux_options"></a>
### Nested Schema for `spec.template.spec.volume.security_context.se_linux_options`

Optional:

- **level** (String, Optional) Level is SELinux level label that applies to the container.
- **role** (String, Optional) Role is a SELinux role label that applies to the container.
- **type** (String, Optional) Type is a SELinux type label that applies to the container.
- **user** (String, Optional) User is a SELinux user label that applies to the container.



<a id="nestedblock--spec--template--spec--volume--startup_probe"></a>
### Nested Schema for `spec.template.spec.volume.startup_probe`

Optional:

- **exec** (Block List, Max: 1) exec specifies the action to take. (see [below for nested schema](#nestedblock--spec--template--spec--volume--startup_probe--exec))
- **failure_threshold** (Number, Optional) Minimum consecutive failures for the probe to be considered failed after having succeeded.
- **http_get** (Block List, Max: 1) Specifies the http request to perform. (see [below for nested schema](#nestedblock--spec--template--spec--volume--startup_probe--http_get))
- **initial_delay_seconds** (Number, Optional) Number of seconds after the container has started before liveness probes are initiated. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes
- **period_seconds** (Number, Optional) How often (in seconds) to perform the probe
- **success_threshold** (Number, Optional) Minimum consecutive successes for the probe to be considered successful after having failed.
- **tcp_socket** (Block List) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported (see [below for nested schema](#nestedblock--spec--template--spec--volume--startup_probe--tcp_socket))
- **timeout_seconds** (Number, Optional) Number of seconds after which the probe times out. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes

<a id="nestedblock--spec--template--spec--volume--startup_probe--exec"></a>
### Nested Schema for `spec.template.spec.volume.startup_probe.timeout_seconds`

Optional:

- **command** (List of String, Optional) Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.


<a id="nestedblock--spec--template--spec--volume--startup_probe--http_get"></a>
### Nested Schema for `spec.template.spec.volume.startup_probe.timeout_seconds`

Optional:

- **host** (String, Optional) Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
- **http_header** (Block List) Scheme to use for connecting to the host. (see [below for nested schema](#nestedblock--spec--template--spec--volume--startup_probe--timeout_seconds--http_header))
- **path** (String, Optional) Path to access on the HTTP server.
- **port** (String, Optional) Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
- **scheme** (String, Optional) Scheme to use for connecting to the host.

<a id="nestedblock--spec--template--spec--volume--startup_probe--timeout_seconds--http_header"></a>
### Nested Schema for `spec.template.spec.volume.startup_probe.timeout_seconds.scheme`

Optional:

- **name** (String, Optional) The header field name
- **value** (String, Optional) The header field value



<a id="nestedblock--spec--template--spec--volume--startup_probe--tcp_socket"></a>
### Nested Schema for `spec.template.spec.volume.startup_probe.timeout_seconds`

Required:

- **port** (String, Required) Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.



<a id="nestedblock--spec--template--spec--volume--volume_mount"></a>
### Nested Schema for `spec.template.spec.volume.volume_mount`

Required:

- **mount_path** (String, Required) Path within the container at which the volume should be mounted. Must not contain ':'.
- **name** (String, Required) This must match the Name of a Volume.

Optional:

- **mount_propagation** (String, Optional) Mount propagation mode. mount_propagation determines how mounts are propagated from the host to container and the other way around. Valid values are None (default), HostToContainer and Bidirectional.
- **read_only** (Boolean, Optional) Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.
- **sub_path** (String, Optional) Path within the volume from which the container's volume should be mounted. Defaults to "" (volume's root).



<a id="nestedblock--spec--template--spec--readiness_gate"></a>
### Nested Schema for `spec.template.spec.volume`

Required:

- **condition_type** (String, Required) refers to a condition in the pod's condition list with matching type.


<a id="nestedblock--spec--template--spec--security_context"></a>
### Nested Schema for `spec.template.spec.volume`

Optional:

- **fs_group** (Number, Optional) A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod: 1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw---- If unset, the Kubelet will not modify the ownership and permissions of any volume.
- **run_as_group** (Number, Optional) The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.
- **run_as_non_root** (Boolean, Optional) Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
- **run_as_user** (Number, Optional) The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.
- **se_linux_options** (Block List, Max: 1) The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. (see [below for nested schema](#nestedblock--spec--template--spec--volume--se_linux_options))
- **supplemental_groups** (Set of Number, Optional) A list of groups applied to the first process run in each container, in addition to the container's primary GID. If unspecified, no groups will be added to any container.
- **sysctl** (Block List) holds a list of namespaced sysctls used for the pod. (see [below for nested schema](#nestedblock--spec--template--spec--volume--sysctl))

<a id="nestedblock--spec--template--spec--volume--se_linux_options"></a>
### Nested Schema for `spec.template.spec.volume.se_linux_options`

Optional:

- **level** (String, Optional) Level is SELinux level label that applies to the container.
- **role** (String, Optional) Role is a SELinux role label that applies to the container.
- **type** (String, Optional) Type is a SELinux type label that applies to the container.
- **user** (String, Optional) User is a SELinux user label that applies to the container.


<a id="nestedblock--spec--template--spec--volume--sysctl"></a>
### Nested Schema for `spec.template.spec.volume.sysctl`

Required:

- **name** (String, Required) Name of a property to set.
- **value** (String, Required) Value of a property to set.



<a id="nestedblock--spec--template--spec--toleration"></a>
### Nested Schema for `spec.template.spec.volume`

Optional:

- **effect** (String, Optional) Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.
- **key** (String, Optional) Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.
- **operator** (String, Optional) Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.
- **toleration_seconds** (String, Optional) TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.
- **value** (String, Optional) Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.


<a id="nestedblock--spec--template--spec--volume"></a>
### Nested Schema for `spec.template.spec.volume`

Optional:

- **aws_elastic_block_store** (Block List, Max: 1) Represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore (see [below for nested schema](#nestedblock--spec--template--spec--volume--aws_elastic_block_store))
- **azure_disk** (Block List, Max: 1) Represents an Azure Data Disk mount on the host and bind mount to the pod. (see [below for nested schema](#nestedblock--spec--template--spec--volume--azure_disk))
- **azure_file** (Block List, Max: 1) Represents an Azure File Service mount on the host and bind mount to the pod. (see [below for nested schema](#nestedblock--spec--template--spec--volume--azure_file))
- **ceph_fs** (Block List, Max: 1) Represents a Ceph FS mount on the host that shares a pod's lifetime (see [below for nested schema](#nestedblock--spec--template--spec--volume--ceph_fs))
- **cinder** (Block List, Max: 1) Represents a cinder volume attached and mounted on kubelets host machine. More info: http://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md (see [below for nested schema](#nestedblock--spec--template--spec--volume--cinder))
- **config_map** (Block List, Max: 1) ConfigMap represents a configMap that should populate this volume (see [below for nested schema](#nestedblock--spec--template--spec--volume--config_map))
- **csi** (Block List, Max: 1) Represents a CSI Volume. More info: http://kubernetes.io/docs/user-guide/volumes#csi (see [below for nested schema](#nestedblock--spec--template--spec--volume--csi))
- **downward_api** (Block List, Max: 1) DownwardAPI represents downward API about the pod that should populate this volume (see [below for nested schema](#nestedblock--spec--template--spec--volume--downward_api))
- **empty_dir** (Block List, Max: 1) EmptyDir represents a temporary directory that shares a pod's lifetime. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir (see [below for nested schema](#nestedblock--spec--template--spec--volume--empty_dir))
- **fc** (Block List, Max: 1) Represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod. (see [below for nested schema](#nestedblock--spec--template--spec--volume--fc))
- **flex_volume** (Block List, Max: 1) Represents a generic volume resource that is provisioned/attached using an exec based plugin. This is an alpha feature and may change in future. (see [below for nested schema](#nestedblock--spec--template--spec--volume--flex_volume))
- **flocker** (Block List, Max: 1) Represents a Flocker volume attached to a kubelet's host machine and exposed to the pod for its usage. This depends on the Flocker control service being running (see [below for nested schema](#nestedblock--spec--template--spec--volume--flocker))
- **gce_persistent_disk** (Block List, Max: 1) Represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin. More info: http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk (see [below for nested schema](#nestedblock--spec--template--spec--volume--gce_persistent_disk))
- **git_repo** (Block List, Max: 1) GitRepo represents a git repository at a particular revision. (see [below for nested schema](#nestedblock--spec--template--spec--volume--git_repo))
- **glusterfs** (Block List, Max: 1) Represents a Glusterfs volume that is attached to a host and exposed to the pod. Provisioned by an admin. More info: http://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md (see [below for nested schema](#nestedblock--spec--template--spec--volume--glusterfs))
- **host_path** (Block List, Max: 1) Represents a directory on the host. Provisioned by a developer or tester. This is useful for single-node development and testing only! On-host storage is not supported in any way and WILL NOT WORK in a multi-node cluster. More info: http://kubernetes.io/docs/user-guide/volumes#hostpath (see [below for nested schema](#nestedblock--spec--template--spec--volume--host_path))
- **iscsi** (Block List, Max: 1) Represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin. (see [below for nested schema](#nestedblock--spec--template--spec--volume--iscsi))
- **local** (Block List, Max: 1) Represents a mounted local storage device such as a disk, partition or directory. Local volumes can only be used as a statically created PersistentVolume. Dynamic provisioning is not supported yet. More info: http://kubernetes.io/docs/user-guide/volumes#local (see [below for nested schema](#nestedblock--spec--template--spec--volume--local))
- **name** (String, Optional) Volume's name. Must be a DNS_LABEL and unique within the pod. More info: http://kubernetes.io/docs/user-guide/identifiers#names
- **nfs** (Block List, Max: 1) Represents an NFS mount on the host. Provisioned by an admin. More info: http://kubernetes.io/docs/user-guide/volumes#nfs (see [below for nested schema](#nestedblock--spec--template--spec--volume--nfs))
- **persistent_volume_claim** (Block List, Max: 1) The specification of a persistent volume. (see [below for nested schema](#nestedblock--spec--template--spec--volume--persistent_volume_claim))
- **photon_persistent_disk** (Block List, Max: 1) Represents a PhotonController persistent disk attached and mounted on kubelets host machine (see [below for nested schema](#nestedblock--spec--template--spec--volume--photon_persistent_disk))
- **projected** (Block List) Projected represents a single volume that projects several volume sources into the same directory. More info: https://kubernetes.io/docs/concepts/storage/volumes/#projected (see [below for nested schema](#nestedblock--spec--template--spec--volume--projected))
- **quobyte** (Block List, Max: 1) Quobyte represents a Quobyte mount on the host that shares a pod's lifetime (see [below for nested schema](#nestedblock--spec--template--spec--volume--quobyte))
- **rbd** (Block List, Max: 1) Represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: http://releases.k8s.io/HEAD/examples/volumes/rbd/README.md (see [below for nested schema](#nestedblock--spec--template--spec--volume--rbd))
- **secret** (Block List, Max: 1) Secret represents a secret that should populate this volume. More info: http://kubernetes.io/docs/user-guide/volumes#secrets (see [below for nested schema](#nestedblock--spec--template--spec--volume--secret))
- **vsphere_volume** (Block List, Max: 1) Represents a vSphere volume attached and mounted on kubelets host machine (see [below for nested schema](#nestedblock--spec--template--spec--volume--vsphere_volume))

<a id="nestedblock--spec--template--spec--volume--aws_elastic_block_store"></a>
### Nested Schema for `spec.template.spec.volume.aws_elastic_block_store`

Required:

- **volume_id** (String, Required) Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore

Optional:

- **fs_type** (String, Optional) Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore
- **partition** (Number, Optional) The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).
- **read_only** (Boolean, Optional) Whether to set the read-only property in VolumeMounts to "true". If omitted, the default is "false". More info: http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore


<a id="nestedblock--spec--template--spec--volume--azure_disk"></a>
### Nested Schema for `spec.template.spec.volume.azure_disk`

Required:

- **caching_mode** (String, Required) Host Caching mode: None, Read Only, Read Write.
- **data_disk_uri** (String, Required) The URI the data disk in the blob storage
- **disk_name** (String, Required) The Name of the data disk in the blob storage

Optional:

- **fs_type** (String, Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
- **kind** (String, Optional) The type for the data disk. Expected values: Shared, Dedicated, Managed. Defaults to Shared
- **read_only** (Boolean, Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).


<a id="nestedblock--spec--template--spec--volume--azure_file"></a>
### Nested Schema for `spec.template.spec.volume.azure_file`

Required:

- **secret_name** (String, Required) The name of secret that contains Azure Storage Account Name and Key
- **share_name** (String, Required) Share Name

Optional:

- **read_only** (Boolean, Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).


<a id="nestedblock--spec--template--spec--volume--ceph_fs"></a>
### Nested Schema for `spec.template.spec.volume.ceph_fs`

Required:

- **monitors** (Set of String, Required) Monitors is a collection of Ceph monitors More info: http://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it

Optional:

- **path** (String, Optional) Used as the mounted root, rather than the full Ceph tree, default is /
- **read_only** (Boolean, Optional) Whether to force the read-only setting in VolumeMounts. Defaults to `false` (read/write). More info: http://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
- **secret_file** (String, Optional) The path to key ring for User, default is /etc/ceph/user.secret More info: http://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
- **secret_ref** (Block List, Max: 1) Reference to the authentication secret for User, default is empty. More info: http://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it (see [below for nested schema](#nestedblock--spec--template--spec--volume--ceph_fs--secret_ref))
- **user** (String, Optional) User is the rados user name, default is admin. More info: http://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it

<a id="nestedblock--spec--template--spec--volume--ceph_fs--secret_ref"></a>
### Nested Schema for `spec.template.spec.volume.ceph_fs.user`

Optional:

- **name** (String, Optional) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
- **namespace** (String, Optional) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names



<a id="nestedblock--spec--template--spec--volume--cinder"></a>
### Nested Schema for `spec.template.spec.volume.cinder`

Required:

- **volume_id** (String, Required) Volume ID used to identify the volume in Cinder. More info: http://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md

Optional:

- **fs_type** (String, Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: http://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
- **read_only** (Boolean, Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write). More info: http://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md


<a id="nestedblock--spec--template--spec--volume--config_map"></a>
### Nested Schema for `spec.template.spec.volume.config_map`

Optional:

- **default_mode** (String, Optional) Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
- **items** (Block List) If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'. (see [below for nested schema](#nestedblock--spec--template--spec--volume--config_map--items))
- **name** (String, Optional) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
- **optional** (Boolean, Optional) Optional: Specify whether the ConfigMap or its keys must be defined.

<a id="nestedblock--spec--template--spec--volume--config_map--items"></a>
### Nested Schema for `spec.template.spec.volume.config_map.optional`

Optional:

- **key** (String, Optional) The key to project.
- **mode** (String, Optional) Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
- **path** (String, Optional) The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.



<a id="nestedblock--spec--template--spec--volume--csi"></a>
### Nested Schema for `spec.template.spec.volume.csi`

Required:

- **driver** (String, Required) the name of the volume driver to use. More info: https://kubernetes.io/docs/concepts/storage/volumes/#csi
- **volume_handle** (String, Required) A string value that uniquely identifies the volume. More info: https://kubernetes.io/docs/concepts/storage/volumes/#csi

Optional:

- **controller_expand_secret_ref** (Block List, Max: 1) A reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI ControllerExpandVolume call. (see [below for nested schema](#nestedblock--spec--template--spec--volume--csi--controller_expand_secret_ref))
- **controller_publish_secret_ref** (Block List, Max: 1) A reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI ControllerPublishVolume and ControllerUnpublishVolume calls. (see [below for nested schema](#nestedblock--spec--template--spec--volume--csi--controller_publish_secret_ref))
- **fs_type** (String, Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
- **node_publish_secret_ref** (Block List, Max: 1) A reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. (see [below for nested schema](#nestedblock--spec--template--spec--volume--csi--node_publish_secret_ref))
- **node_stage_secret_ref** (Block List, Max: 1) A reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodeStageVolume and NodeStageVolume and NodeUnstageVolume calls. (see [below for nested schema](#nestedblock--spec--template--spec--volume--csi--node_stage_secret_ref))
- **read_only** (Boolean, Optional) Whether to set the read-only property in VolumeMounts to "true". If omitted, the default is "false". More info: http://kubernetes.io/docs/user-guide/volumes#csi
- **volume_attributes** (Map of String, Optional) Attributes of the volume to publish.

<a id="nestedblock--spec--template--spec--volume--csi--controller_expand_secret_ref"></a>
### Nested Schema for `spec.template.spec.volume.csi.volume_attributes`

Optional:

- **name** (String, Optional) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
- **namespace** (String, Optional) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names


<a id="nestedblock--spec--template--spec--volume--csi--controller_publish_secret_ref"></a>
### Nested Schema for `spec.template.spec.volume.csi.volume_attributes`

Optional:

- **name** (String, Optional) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
- **namespace** (String, Optional) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names


<a id="nestedblock--spec--template--spec--volume--csi--node_publish_secret_ref"></a>
### Nested Schema for `spec.template.spec.volume.csi.volume_attributes`

Optional:

- **name** (String, Optional) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
- **namespace** (String, Optional) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names


<a id="nestedblock--spec--template--spec--volume--csi--node_stage_secret_ref"></a>
### Nested Schema for `spec.template.spec.volume.csi.volume_attributes`

Optional:

- **name** (String, Optional) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
- **namespace** (String, Optional) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names



<a id="nestedblock--spec--template--spec--volume--downward_api"></a>
### Nested Schema for `spec.template.spec.volume.downward_api`

Optional:

- **default_mode** (String, Optional) Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
- **items** (Block List) If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error. Paths must be relative and may not contain the '..' path or start with '..'. (see [below for nested schema](#nestedblock--spec--template--spec--volume--downward_api--items))

<a id="nestedblock--spec--template--spec--volume--downward_api--items"></a>
### Nested Schema for `spec.template.spec.volume.downward_api.items`

Required:

- **field_ref** (Block List, Min: 1, Max: 1) Required: Selects a field of the pod: only annotations, labels, name and namespace are supported. (see [below for nested schema](#nestedblock--spec--template--spec--volume--downward_api--items--field_ref))
- **path** (String, Required) Path is the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'

Optional:

- **mode** (String, Optional) Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
- **resource_field_ref** (Block List, Max: 1) Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported. (see [below for nested schema](#nestedblock--spec--template--spec--volume--downward_api--items--resource_field_ref))

<a id="nestedblock--spec--template--spec--volume--downward_api--items--field_ref"></a>
### Nested Schema for `spec.template.spec.volume.downward_api.items.resource_field_ref`

Optional:

- **api_version** (String, Optional) Version of the schema the FieldPath is written in terms of, defaults to "v1".
- **field_path** (String, Optional) Path of the field to select in the specified API version


<a id="nestedblock--spec--template--spec--volume--downward_api--items--resource_field_ref"></a>
### Nested Schema for `spec.template.spec.volume.downward_api.items.resource_field_ref`

Required:

- **container_name** (String, Required)
- **resource** (String, Required) Resource to select

Optional:

- **divisor** (String, Optional)




<a id="nestedblock--spec--template--spec--volume--empty_dir"></a>
### Nested Schema for `spec.template.spec.volume.empty_dir`

Optional:

- **medium** (String, Optional) What type of storage medium should back this directory. The default is "" which means to use the node's default medium. Must be an empty string (default) or Memory. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir
- **size_limit** (String, Optional) Total amount of local storage required for this EmptyDir volume.


<a id="nestedblock--spec--template--spec--volume--fc"></a>
### Nested Schema for `spec.template.spec.volume.fc`

Required:

- **lun** (Number, Required) FC target lun number
- **target_ww_ns** (Set of String, Required) FC target worldwide names (WWNs)

Optional:

- **fs_type** (String, Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
- **read_only** (Boolean, Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).


<a id="nestedblock--spec--template--spec--volume--flex_volume"></a>
### Nested Schema for `spec.template.spec.volume.flex_volume`

Required:

- **driver** (String, Required) Driver is the name of the driver to use for this volume.

Optional:

- **fs_type** (String, Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
- **options** (Map of String, Optional) Extra command options if any.
- **read_only** (Boolean, Optional) Whether to force the ReadOnly setting in VolumeMounts. Defaults to false (read/write).
- **secret_ref** (Block List, Max: 1) Reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts. (see [below for nested schema](#nestedblock--spec--template--spec--volume--flex_volume--secret_ref))

<a id="nestedblock--spec--template--spec--volume--flex_volume--secret_ref"></a>
### Nested Schema for `spec.template.spec.volume.flex_volume.secret_ref`

Optional:

- **name** (String, Optional) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
- **namespace** (String, Optional) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names



<a id="nestedblock--spec--template--spec--volume--flocker"></a>
### Nested Schema for `spec.template.spec.volume.flocker`

Optional:

- **dataset_name** (String, Optional) Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated
- **dataset_uuid** (String, Optional) UUID of the dataset. This is unique identifier of a Flocker dataset


<a id="nestedblock--spec--template--spec--volume--gce_persistent_disk"></a>
### Nested Schema for `spec.template.spec.volume.gce_persistent_disk`

Required:

- **pd_name** (String, Required) Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk

Optional:

- **fs_type** (String, Optional) Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk
- **partition** (Number, Optional) The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty). More info: http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk
- **read_only** (Boolean, Optional) Whether to force the ReadOnly setting in VolumeMounts. Defaults to false. More info: http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk


<a id="nestedblock--spec--template--spec--volume--git_repo"></a>
### Nested Schema for `spec.template.spec.volume.git_repo`

Optional:

- **directory** (String, Optional) Target directory name. Must not contain or start with '..'. If '.' is supplied, the volume directory will be the git repository. Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.
- **repository** (String, Optional) Repository URL
- **revision** (String, Optional) Commit hash for the specified revision.


<a id="nestedblock--spec--template--spec--volume--glusterfs"></a>
### Nested Schema for `spec.template.spec.volume.glusterfs`

Required:

- **endpoints_name** (String, Required) The endpoint name that details Glusterfs topology. More info: http://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod
- **path** (String, Required) The Glusterfs volume path. More info: http://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod

Optional:

- **read_only** (Boolean, Optional) Whether to force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: http://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod


<a id="nestedblock--spec--template--spec--volume--host_path"></a>
### Nested Schema for `spec.template.spec.volume.host_path`

Optional:

- **path** (String, Optional) Path of the directory on the host. More info: http://kubernetes.io/docs/user-guide/volumes#hostpath
- **type** (String, Optional) Type for HostPath volume. Allowed values are "" (default), DirectoryOrCreate, Directory, FileOrCreate, File, Socket, CharDevice and BlockDevice


<a id="nestedblock--spec--template--spec--volume--iscsi"></a>
### Nested Schema for `spec.template.spec.volume.iscsi`

Required:

- **iqn** (String, Required) Target iSCSI Qualified Name.
- **target_portal** (String, Required) iSCSI target portal. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).

Optional:

- **fs_type** (String, Optional) Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: http://kubernetes.io/docs/user-guide/volumes#iscsi
- **iscsi_interface** (String, Optional) iSCSI interface name that uses an iSCSI transport. Defaults to 'default' (tcp).
- **lun** (Number, Optional) iSCSI target lun number.
- **read_only** (Boolean, Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false.


<a id="nestedblock--spec--template--spec--volume--local"></a>
### Nested Schema for `spec.template.spec.volume.local`

Optional:

- **path** (String, Optional) Path of the directory on the host. More info: http://kubernetes.io/docs/user-guide/volumes#local


<a id="nestedblock--spec--template--spec--volume--nfs"></a>
### Nested Schema for `spec.template.spec.volume.nfs`

Required:

- **path** (String, Required) Path that is exported by the NFS server. More info: http://kubernetes.io/docs/user-guide/volumes#nfs
- **server** (String, Required) Server is the hostname or IP address of the NFS server. More info: http://kubernetes.io/docs/user-guide/volumes#nfs

Optional:

- **read_only** (Boolean, Optional) Whether to force the NFS export to be mounted with read-only permissions. Defaults to false. More info: http://kubernetes.io/docs/user-guide/volumes#nfs


<a id="nestedblock--spec--template--spec--volume--persistent_volume_claim"></a>
### Nested Schema for `spec.template.spec.volume.persistent_volume_claim`

Optional:

- **claim_name** (String, Optional) ClaimName is the name of a PersistentVolumeClaim in the same
- **read_only** (Boolean, Optional) Will force the ReadOnly setting in VolumeMounts.


<a id="nestedblock--spec--template--spec--volume--photon_persistent_disk"></a>
### Nested Schema for `spec.template.spec.volume.photon_persistent_disk`

Required:

- **pd_id** (String, Required) ID that identifies Photon Controller persistent disk

Optional:

- **fs_type** (String, Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.


<a id="nestedblock--spec--template--spec--volume--projected"></a>
### Nested Schema for `spec.template.spec.volume.projected`

Required:

- **sources** (Block List, Min: 1) Source of the volume to project in the directory. (see [below for nested schema](#nestedblock--spec--template--spec--volume--projected--sources))

Optional:

- **default_mode** (String, Optional) Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.

<a id="nestedblock--spec--template--spec--volume--projected--sources"></a>
### Nested Schema for `spec.template.spec.volume.projected.default_mode`

Optional:

- **config_map** (Block List) ConfigMap represents a configMap that should populate this volume (see [below for nested schema](#nestedblock--spec--template--spec--volume--projected--default_mode--config_map))
- **downward_api** (Block List, Max: 1) DownwardAPI represents downward API about the pod that should populate this volume (see [below for nested schema](#nestedblock--spec--template--spec--volume--projected--default_mode--downward_api))
- **secret** (Block List) Secret represents a secret that should populate this volume. More info: http://kubernetes.io/docs/user-guide/volumes#secrets (see [below for nested schema](#nestedblock--spec--template--spec--volume--projected--default_mode--secret))
- **service_account_token** (Block List, Max: 1) A projected service account token volume (see [below for nested schema](#nestedblock--spec--template--spec--volume--projected--default_mode--service_account_token))

<a id="nestedblock--spec--template--spec--volume--projected--default_mode--config_map"></a>
### Nested Schema for `spec.template.spec.volume.projected.default_mode.service_account_token`

Optional:

- **items** (Block List) If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error. Paths must be relative and may not contain the '..' path or start with '..'. (see [below for nested schema](#nestedblock--spec--template--spec--volume--projected--default_mode--service_account_token--items))
- **name** (String, Optional) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
- **optional** (Boolean, Optional) Optional: Specify whether the ConfigMap or it's keys must be defined.

<a id="nestedblock--spec--template--spec--volume--projected--default_mode--service_account_token--items"></a>
### Nested Schema for `spec.template.spec.volume.projected.default_mode.service_account_token.optional`

Optional:

- **key** (String, Optional) The key to project.
- **mode** (String, Optional) Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
- **path** (String, Optional) The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.



<a id="nestedblock--spec--template--spec--volume--projected--default_mode--downward_api"></a>
### Nested Schema for `spec.template.spec.volume.projected.default_mode.service_account_token`

Optional:

- **items** (Block List) Represents a volume containing downward API info. Downward API volumes support ownership management and SELinux relabeling. (see [below for nested schema](#nestedblock--spec--template--spec--volume--projected--default_mode--service_account_token--items))

<a id="nestedblock--spec--template--spec--volume--projected--default_mode--service_account_token--items"></a>
### Nested Schema for `spec.template.spec.volume.projected.default_mode.service_account_token.items`

Required:

- **path** (String, Required) Path is the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'

Optional:

- **field_ref** (Block List, Max: 1) Selects a field of the pod: only annotations, labels, name and namespace are supported. (see [below for nested schema](#nestedblock--spec--template--spec--volume--projected--default_mode--service_account_token--items--field_ref))
- **mode** (String, Optional) Mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
- **resource_field_ref** (Block List, Max: 1) Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported. (see [below for nested schema](#nestedblock--spec--template--spec--volume--projected--default_mode--service_account_token--items--resource_field_ref))

<a id="nestedblock--spec--template--spec--volume--projected--default_mode--service_account_token--items--field_ref"></a>
### Nested Schema for `spec.template.spec.volume.projected.default_mode.service_account_token.items.field_ref`

Optional:

- **api_version** (String, Optional) Version of the schema the FieldPath is written in terms of, defaults to 'v1'.
- **field_path** (String, Optional) Path of the field to select in the specified API version


<a id="nestedblock--spec--template--spec--volume--projected--default_mode--service_account_token--items--resource_field_ref"></a>
### Nested Schema for `spec.template.spec.volume.projected.default_mode.service_account_token.items.resource_field_ref`

Required:

- **container_name** (String, Required)
- **resource** (String, Required) Resource to select

Optional:

- **quantity** (String, Optional)




<a id="nestedblock--spec--template--spec--volume--projected--default_mode--secret"></a>
### Nested Schema for `spec.template.spec.volume.projected.default_mode.service_account_token`

Optional:

- **items** (Block List) If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'. (see [below for nested schema](#nestedblock--spec--template--spec--volume--projected--default_mode--service_account_token--items))
- **name** (String, Optional) Name of the secret in the pod's namespace to use. More info: http://kubernetes.io/docs/user-guide/volumes#secrets
- **optional** (Boolean, Optional) Optional: Specify whether the Secret or it's keys must be defined.

<a id="nestedblock--spec--template--spec--volume--projected--default_mode--service_account_token--items"></a>
### Nested Schema for `spec.template.spec.volume.projected.default_mode.service_account_token.optional`

Optional:

- **key** (String, Optional) The key to project.
- **mode** (String, Optional) Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
- **path** (String, Optional) The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.



<a id="nestedblock--spec--template--spec--volume--projected--default_mode--service_account_token"></a>
### Nested Schema for `spec.template.spec.volume.projected.default_mode.service_account_token`

Required:

- **path** (String, Required) Path specifies a relative path to the mount point of the projected volume.

Optional:

- **audience** (String, Optional) Audience is the intended audience of the token
- **expiration_seconds** (Number, Optional) ExpirationSeconds is the expected duration of validity of the service account token. It defaults to 1 hour and must be at least 10 minutes (600 seconds).




<a id="nestedblock--spec--template--spec--volume--quobyte"></a>
### Nested Schema for `spec.template.spec.volume.quobyte`

Required:

- **registry** (String, Required) Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes
- **volume** (String, Required) Volume is a string that references an already created Quobyte volume by name.

Optional:

- **group** (String, Optional) Group to map volume access to Default is no group
- **read_only** (Boolean, Optional) Whether to force the Quobyte volume to be mounted with read-only permissions. Defaults to false.
- **user** (String, Optional) User to map volume access to Defaults to serivceaccount user


<a id="nestedblock--spec--template--spec--volume--rbd"></a>
### Nested Schema for `spec.template.spec.volume.rbd`

Required:

- **ceph_monitors** (Set of String, Required) A collection of Ceph monitors. More info: http://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
- **rbd_image** (String, Required) The rados image name. More info: http://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it

Optional:

- **fs_type** (String, Optional) Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: http://kubernetes.io/docs/user-guide/volumes#rbd
- **keyring** (String, Optional) Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: http://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
- **rados_user** (String, Optional) The rados user name. Default is admin. More info: http://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
- **rbd_pool** (String, Optional) The rados pool name. Default is rbd. More info: http://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it.
- **read_only** (Boolean, Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false. More info: http://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
- **secret_ref** (Block List, Max: 1) Name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: http://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it (see [below for nested schema](#nestedblock--spec--template--spec--volume--rbd--secret_ref))

<a id="nestedblock--spec--template--spec--volume--rbd--secret_ref"></a>
### Nested Schema for `spec.template.spec.volume.rbd.secret_ref`

Optional:

- **name** (String, Optional) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
- **namespace** (String, Optional) Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names



<a id="nestedblock--spec--template--spec--volume--secret"></a>
### Nested Schema for `spec.template.spec.volume.secret`

Optional:

- **default_mode** (String, Optional) Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
- **items** (Block List) If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'. (see [below for nested schema](#nestedblock--spec--template--spec--volume--secret--items))
- **optional** (Boolean, Optional) Optional: Specify whether the Secret or its keys must be defined.
- **secret_name** (String, Optional) Name of the secret in the pod's namespace to use. More info: http://kubernetes.io/docs/user-guide/volumes#secrets

<a id="nestedblock--spec--template--spec--volume--secret--items"></a>
### Nested Schema for `spec.template.spec.volume.secret.secret_name`

Optional:

- **key** (String, Optional) The key to project.
- **mode** (String, Optional) Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
- **path** (String, Optional) The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.



<a id="nestedblock--spec--template--spec--volume--vsphere_volume"></a>
### Nested Schema for `spec.template.spec.volume.vsphere_volume`

Required:

- **volume_path** (String, Required) Path that identifies vSphere volume vmdk

Optional:

- **fs_type** (String, Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.





<a id="nestedblock--spec--selector"></a>
### Nested Schema for `spec.selector`

Optional:

- **match_expressions** (Block List) A list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedblock--spec--selector--match_expressions))
- **match_labels** (Map of String, Optional) A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of `match_expressions`, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.

<a id="nestedblock--spec--selector--match_expressions"></a>
### Nested Schema for `spec.selector.match_expressions`

Optional:

- **key** (String, Optional) The label key that the selector applies to.
- **operator** (String, Optional) A key's relationship to a set of values. Valid operators ard `In`, `NotIn`, `Exists` and `DoesNotExist`.
- **values** (Set of String, Optional) An array of string values. If the operator is `In` or `NotIn`, the values array must be non-empty. If the operator is `Exists` or `DoesNotExist`, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedblock--spec--strategy"></a>
### Nested Schema for `spec.strategy`

Optional:

- **rolling_update** (Block List, Max: 1) Rolling update config params. Present only if DeploymentStrategyType = RollingUpdate. (see [below for nested schema](#nestedblock--spec--strategy--rolling_update))
- **type** (String, Optional) Type of deployment. Can be 'Recreate' or 'RollingUpdate'. Default is RollingUpdate.

<a id="nestedblock--spec--strategy--rolling_update"></a>
### Nested Schema for `spec.strategy.rolling_update`

Optional:

- **max_surge** (String, Optional) The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%. Example: when this is set to 30%, the new RC can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new RC can be scaled up further, ensuring that total number of pods running at any time during the update is atmost 130% of desired pods.
- **max_unavailable** (String, Optional) The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%. Example: when this is set to 30%, the old RC can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old RC can be scaled down further, followed by scaling up the new RC, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods.




<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- **create** (String, Optional)
- **delete** (String, Optional)
- **update** (String, Optional)


