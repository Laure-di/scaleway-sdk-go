// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package k8s provides methods and message types of the k8s v1 API.
package k8s

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/marshaler"
	"github.com/scaleway/scaleway-sdk-go/internal/parameter"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ fmt.Stringer
	_ json.Unmarshaler
	_ url.URL
	_ net.IP
	_ http.Header
	_ bytes.Reader
	_ time.Time
	_ = strings.Join

	_ scw.ScalewayRequest
	_ marshaler.Duration
	_ scw.File
	_ = parameter.AddToQuery
	_ = namegenerator.GetRandomName
)

type AutoscalerEstimator string

const (
	AutoscalerEstimatorUnknownEstimator = AutoscalerEstimator("unknown_estimator")
	AutoscalerEstimatorBinpacking       = AutoscalerEstimator("binpacking")
)

func (enum AutoscalerEstimator) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_estimator"
	}
	return string(enum)
}

func (enum AutoscalerEstimator) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AutoscalerEstimator) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AutoscalerEstimator(AutoscalerEstimator(tmp).String())
	return nil
}

// Kubernetes autoscaler strategy to fit pods into cluster nodes (https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders).
type AutoscalerExpander string

const (
	AutoscalerExpanderUnknownExpander = AutoscalerExpander("unknown_expander")
	AutoscalerExpanderRandom          = AutoscalerExpander("random")
	AutoscalerExpanderMostPods        = AutoscalerExpander("most_pods")
	AutoscalerExpanderLeastWaste      = AutoscalerExpander("least_waste")
	AutoscalerExpanderPriority        = AutoscalerExpander("priority")
	AutoscalerExpanderPrice           = AutoscalerExpander("price")
)

func (enum AutoscalerExpander) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_expander"
	}
	return string(enum)
}

func (enum AutoscalerExpander) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AutoscalerExpander) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AutoscalerExpander(AutoscalerExpander(tmp).String())
	return nil
}

type CNI string

const (
	CNIUnknownCni = CNI("unknown_cni")
	// Cilium CNI will be configured (https://github.com/cilium/cilium).
	CNICilium = CNI("cilium")
	// Calico CNI will be configured (https://github.com/projectcalico/calico).
	CNICalico = CNI("calico")
	CNIWeave  = CNI("weave")
	// Kilo CNI will be configured (https://github.com/squat/kilo/). Note that this CNI is only available for Kosmos clusters.
	CNIKilo    = CNI("kilo")
	CNIFlannel = CNI("flannel")
)

func (enum CNI) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_cni"
	}
	return string(enum)
}

func (enum CNI) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CNI) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CNI(CNI(tmp).String())
	return nil
}

type ClusterStatus string

const (
	ClusterStatusUnknown = ClusterStatus("unknown")
	// Cluster is provisioning.
	ClusterStatusCreating = ClusterStatus("creating")
	// Cluster is ready to use.
	ClusterStatusReady = ClusterStatus("ready")
	// Cluster is waiting to be processed for deletion.
	ClusterStatusDeleting = ClusterStatus("deleting")
	ClusterStatusDeleted  = ClusterStatus("deleted")
	// Cluster is updating its own configuration, it can be a version upgrade too.
	ClusterStatusUpdating = ClusterStatus("updating")
	// Cluster is locked because an abuse has been detected or reported.
	ClusterStatusLocked = ClusterStatus("locked")
	// Cluster has no associated pool.
	ClusterStatusPoolRequired = ClusterStatus("pool_required")
)

func (enum ClusterStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum ClusterStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ClusterStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ClusterStatus(ClusterStatus(tmp).String())
	return nil
}

type ClusterTypeAvailability string

const (
	// Type is available in quantity.
	ClusterTypeAvailabilityAvailable = ClusterTypeAvailability("available")
	// Limited availability.
	ClusterTypeAvailabilityScarce = ClusterTypeAvailability("scarce")
	// Out of stock.
	ClusterTypeAvailabilityShortage = ClusterTypeAvailability("shortage")
)

func (enum ClusterTypeAvailability) String() string {
	if enum == "" {
		// return default value if empty
		return "available"
	}
	return string(enum)
}

func (enum ClusterTypeAvailability) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ClusterTypeAvailability) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ClusterTypeAvailability(ClusterTypeAvailability(tmp).String())
	return nil
}

type ClusterTypeResiliency string

const (
	ClusterTypeResiliencyUnknownResiliency = ClusterTypeResiliency("unknown_resiliency")
	// The control plane is rescheduled on other machines in case of failure of a lower layer.
	ClusterTypeResiliencyStandard = ClusterTypeResiliency("standard")
	// The control plane has replicas to ensure service continuity in case of failure of a lower layer.
	ClusterTypeResiliencyHighAvailability = ClusterTypeResiliency("high_availability")
)

func (enum ClusterTypeResiliency) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_resiliency"
	}
	return string(enum)
}

func (enum ClusterTypeResiliency) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ClusterTypeResiliency) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ClusterTypeResiliency(ClusterTypeResiliency(tmp).String())
	return nil
}

// Managed Ingress Controllers are deprecated. Use the Easy Deploy feature instead.
type Ingress string

const (
	IngressUnknownIngress = Ingress("unknown_ingress")
	IngressNone           = Ingress("none")
	IngressNginx          = Ingress("nginx")
	IngressTraefik        = Ingress("traefik")
	IngressTraefik2       = Ingress("traefik2")
)

func (enum Ingress) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_ingress"
	}
	return string(enum)
}

func (enum Ingress) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *Ingress) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = Ingress(Ingress(tmp).String())
	return nil
}

type ListClustersRequestOrderBy string

const (
	ListClustersRequestOrderByCreatedAtAsc  = ListClustersRequestOrderBy("created_at_asc")
	ListClustersRequestOrderByCreatedAtDesc = ListClustersRequestOrderBy("created_at_desc")
	ListClustersRequestOrderByUpdatedAtAsc  = ListClustersRequestOrderBy("updated_at_asc")
	ListClustersRequestOrderByUpdatedAtDesc = ListClustersRequestOrderBy("updated_at_desc")
	ListClustersRequestOrderByNameAsc       = ListClustersRequestOrderBy("name_asc")
	ListClustersRequestOrderByNameDesc      = ListClustersRequestOrderBy("name_desc")
	ListClustersRequestOrderByStatusAsc     = ListClustersRequestOrderBy("status_asc")
	ListClustersRequestOrderByStatusDesc    = ListClustersRequestOrderBy("status_desc")
	ListClustersRequestOrderByVersionAsc    = ListClustersRequestOrderBy("version_asc")
	ListClustersRequestOrderByVersionDesc   = ListClustersRequestOrderBy("version_desc")
)

func (enum ListClustersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListClustersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListClustersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListClustersRequestOrderBy(ListClustersRequestOrderBy(tmp).String())
	return nil
}

type ListNodesRequestOrderBy string

const (
	ListNodesRequestOrderByCreatedAtAsc  = ListNodesRequestOrderBy("created_at_asc")
	ListNodesRequestOrderByCreatedAtDesc = ListNodesRequestOrderBy("created_at_desc")
)

func (enum ListNodesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListNodesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListNodesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListNodesRequestOrderBy(ListNodesRequestOrderBy(tmp).String())
	return nil
}

type ListPoolsRequestOrderBy string

const (
	ListPoolsRequestOrderByCreatedAtAsc  = ListPoolsRequestOrderBy("created_at_asc")
	ListPoolsRequestOrderByCreatedAtDesc = ListPoolsRequestOrderBy("created_at_desc")
	ListPoolsRequestOrderByUpdatedAtAsc  = ListPoolsRequestOrderBy("updated_at_asc")
	ListPoolsRequestOrderByUpdatedAtDesc = ListPoolsRequestOrderBy("updated_at_desc")
	ListPoolsRequestOrderByNameAsc       = ListPoolsRequestOrderBy("name_asc")
	ListPoolsRequestOrderByNameDesc      = ListPoolsRequestOrderBy("name_desc")
	ListPoolsRequestOrderByStatusAsc     = ListPoolsRequestOrderBy("status_asc")
	ListPoolsRequestOrderByStatusDesc    = ListPoolsRequestOrderBy("status_desc")
	ListPoolsRequestOrderByVersionAsc    = ListPoolsRequestOrderBy("version_asc")
	ListPoolsRequestOrderByVersionDesc   = ListPoolsRequestOrderBy("version_desc")
)

func (enum ListPoolsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListPoolsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPoolsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPoolsRequestOrderBy(ListPoolsRequestOrderBy(tmp).String())
	return nil
}

type MaintenanceWindowDayOfTheWeek string

const (
	MaintenanceWindowDayOfTheWeekAny       = MaintenanceWindowDayOfTheWeek("any")
	MaintenanceWindowDayOfTheWeekMonday    = MaintenanceWindowDayOfTheWeek("monday")
	MaintenanceWindowDayOfTheWeekTuesday   = MaintenanceWindowDayOfTheWeek("tuesday")
	MaintenanceWindowDayOfTheWeekWednesday = MaintenanceWindowDayOfTheWeek("wednesday")
	MaintenanceWindowDayOfTheWeekThursday  = MaintenanceWindowDayOfTheWeek("thursday")
	MaintenanceWindowDayOfTheWeekFriday    = MaintenanceWindowDayOfTheWeek("friday")
	MaintenanceWindowDayOfTheWeekSaturday  = MaintenanceWindowDayOfTheWeek("saturday")
	MaintenanceWindowDayOfTheWeekSunday    = MaintenanceWindowDayOfTheWeek("sunday")
)

func (enum MaintenanceWindowDayOfTheWeek) String() string {
	if enum == "" {
		// return default value if empty
		return "any"
	}
	return string(enum)
}

func (enum MaintenanceWindowDayOfTheWeek) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MaintenanceWindowDayOfTheWeek) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MaintenanceWindowDayOfTheWeek(MaintenanceWindowDayOfTheWeek(tmp).String())
	return nil
}

type NodeStatus string

const (
	NodeStatusUnknown = NodeStatus("unknown")
	// Node is provisioning.
	NodeStatusCreating = NodeStatus("creating")
	// Node is unable to connect to apiserver.
	NodeStatusNotReady = NodeStatus("not_ready")
	// Node is ready to execute workload (marked schedulable by k8s scheduler).
	NodeStatusReady = NodeStatus("ready")
	// Node is waiting to be processed for deletion.
	NodeStatusDeleting = NodeStatus("deleting")
	NodeStatusDeleted  = NodeStatus("deleted")
	// Node is locked because an abuse has been detected or reported.
	NodeStatusLocked = NodeStatus("locked")
	// Node is rebooting.
	NodeStatusRebooting     = NodeStatus("rebooting")
	NodeStatusCreationError = NodeStatus("creation_error")
	NodeStatusUpgrading     = NodeStatus("upgrading")
	NodeStatusStarting      = NodeStatus("starting")
	NodeStatusRegistering   = NodeStatus("registering")
)

func (enum NodeStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum NodeStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NodeStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NodeStatus(NodeStatus(tmp).String())
	return nil
}

type PoolStatus string

const (
	PoolStatusUnknown = PoolStatus("unknown")
	// Pool has the right amount of nodes and is ready to process the workload.
	PoolStatusReady = PoolStatus("ready")
	// Pool is waiting to be processed for deletion.
	PoolStatusDeleting = PoolStatus("deleting")
	PoolStatusDeleted  = PoolStatus("deleted")
	// Pool is growing or shrinking.
	PoolStatusScaling = PoolStatus("scaling")
	// Pool is locked because an abuse has been detected or reported.
	PoolStatusLocked  = PoolStatus("locked")
	PoolStatusWarning = PoolStatus("warning")
	// Pool is upgrading its Kubernetes version.
	PoolStatusUpgrading = PoolStatus("upgrading")
)

func (enum PoolStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum PoolStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PoolStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PoolStatus(PoolStatus(tmp).String())
	return nil
}

type PoolVolumeType string

const (
	PoolVolumeTypeDefaultVolumeType = PoolVolumeType("default_volume_type")
	// Local Block Storage: your system is stored locally on your node hypervisor. Lower latency, no persistence across node replacements.
	PoolVolumeTypeLSSD = PoolVolumeType("l_ssd")
	// Remote Block Storage: your system is stored on a centralized and resilient cluster. Higher latency, persistence across node replacements.
	PoolVolumeTypeBSSD = PoolVolumeType("b_ssd")
)

func (enum PoolVolumeType) String() string {
	if enum == "" {
		// return default value if empty
		return "default_volume_type"
	}
	return string(enum)
}

func (enum PoolVolumeType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PoolVolumeType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PoolVolumeType(PoolVolumeType(tmp).String())
	return nil
}

type Runtime string

const (
	RuntimeUnknownRuntime = Runtime("unknown_runtime")
	RuntimeDocker         = Runtime("docker")
	// Containerd Runtime will be configured (https://github.com/containerd/containerd).
	RuntimeContainerd = Runtime("containerd")
	RuntimeCrio       = Runtime("crio")
)

func (enum Runtime) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_runtime"
	}
	return string(enum)
}

func (enum Runtime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *Runtime) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = Runtime(Runtime(tmp).String())
	return nil
}

type MaintenanceWindow struct {
	// StartHour: Start time of the two-hour maintenance window.
	StartHour uint32 `json:"start_hour"`
	// Day: Day of the week for the maintenance window.
	Day MaintenanceWindowDayOfTheWeek `json:"day"`
}

type PoolUpgradePolicy struct {
	// MaxUnavailable:
	MaxUnavailable uint32 `json:"max_unavailable"`
	// MaxSurge:
	MaxSurge uint32 `json:"max_surge"`
}

type CreateClusterRequestPoolConfigUpgradePolicy struct {
	// MaxUnavailable: The maximum number of nodes that can be not ready at the same time.
	MaxUnavailable *uint32 `json:"max_unavailable,omitempty"`
	// MaxSurge: The maximum number of nodes to be created during the upgrade.
	MaxSurge *uint32 `json:"max_surge,omitempty"`
}

type ClusterAutoUpgrade struct {
	// Enabled: Defines whether auto upgrade is enabled for the cluster.
	Enabled bool `json:"enabled"`
	// MaintenanceWindow: Maintenance window of the cluster auto upgrades.
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window"`
}

type ClusterAutoscalerConfig struct {
	// ScaleDownDisabled: Disable the cluster autoscaler.
	ScaleDownDisabled bool `json:"scale_down_disabled"`
	// ScaleDownDelayAfterAdd: How long after scale up that scale down evaluation resumes.
	ScaleDownDelayAfterAdd string `json:"scale_down_delay_after_add"`
	// Estimator: Type of resource estimator to be used in scale up.
	Estimator AutoscalerEstimator `json:"estimator"`
	// Expander: Type of node group expander to be used in scale up.
	Expander AutoscalerExpander `json:"expander"`
	// IgnoreDaemonsetsUtilization: Ignore DaemonSet pods when calculating resource utilization for scaling down.
	IgnoreDaemonsetsUtilization bool `json:"ignore_daemonsets_utilization"`
	// BalanceSimilarNodeGroups: Detect similar node groups and balance the number of nodes between them.
	BalanceSimilarNodeGroups bool `json:"balance_similar_node_groups"`
	// ExpendablePodsPriorityCutoff: Pods with priority below cutoff will be expendable. They can be killed without any consideration during scale down and they won't cause scale up. Pods with null priority (PodPriority disabled) are non expendable.
	ExpendablePodsPriorityCutoff int32 `json:"expendable_pods_priority_cutoff"`
	// ScaleDownUnneededTime: How long a node should be unneeded before it is eligible to be scaled down.
	ScaleDownUnneededTime string `json:"scale_down_unneeded_time"`
	// ScaleDownUtilizationThreshold: Node utilization level, defined as a sum of requested resources divided by capacity, below which a node can be considered for scale down.
	ScaleDownUtilizationThreshold float32 `json:"scale_down_utilization_threshold"`
	// MaxGracefulTerminationSec: Maximum number of seconds the cluster autoscaler waits for pod termination when trying to scale down a node.
	MaxGracefulTerminationSec uint32 `json:"max_graceful_termination_sec"`
}

type ClusterOpenIDConnectConfig struct {
	// IssuerURL: URL of the provider which allows the API server to discover public signing keys. Only URLs using the `https://` scheme are accepted. This is typically the provider's discovery URL without a path, for example "https://accounts.google.com" or "https://login.salesforce.com".
	IssuerURL string `json:"issuer_url"`
	// ClientID: A client ID that all tokens must be issued for.
	ClientID string `json:"client_id"`
	// UsernameClaim: JWT claim to use as the user name. The default is `sub`, which is expected to be the end user's unique identifier. Admins can choose other claims, such as `email` or `name`, depending on their provider. However, claims other than `email` will be prefixed with the issuer URL to prevent name collision.
	UsernameClaim string `json:"username_claim"`
	// UsernamePrefix: Prefix prepended to username claims to prevent name collision (such as `system:` users). For example, the value `oidc:` will create usernames like `oidc:jane.doe`. If this flag is not provided and `username_claim` is a value other than `email`, the prefix defaults to `( Issuer URL )#` where `( Issuer URL )` is the value of `issuer_url`. The value `-` can be used to disable all prefixing.
	UsernamePrefix string `json:"username_prefix"`
	// GroupsClaim: JWT claim to use as the user's group.
	GroupsClaim []string `json:"groups_claim"`
	// GroupsPrefix: Prefix prepended to group claims to prevent name collision (such as `system:` groups). For example, the value `oidc:` will create group names like `oidc:engineering` and `oidc:infra`.
	GroupsPrefix string `json:"groups_prefix"`
	// RequiredClaim: Multiple key=value pairs describing a required claim in the ID token. If set, the claims are verified to be present in the ID token with a matching value.
	RequiredClaim []string `json:"required_claim"`
}

type Pool struct {
	// ID: Pool ID.
	ID string `json:"id"`
	// ClusterID: Cluster ID of the pool.
	ClusterID string `json:"cluster_id"`
	// CreatedAt: Date on which the pool was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Date on which the pool was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Name: Pool name.
	Name string `json:"name"`
	// Status: Pool status.
	Status PoolStatus `json:"status"`
	// Version: Pool version.
	Version string `json:"version"`
	// NodeType: Node type is the type of Scaleway Instance wanted for the pool. Nodes with insufficient memory are not eligible (DEV1-S, PLAY2-PICO, STARDUST). 'external' is a special node type used to provision instances from other cloud providers in a Kosmos Cluster.
	NodeType string `json:"node_type"`
	// Autoscaling: Defines whether the autoscaling feature is enabled for the pool.
	Autoscaling bool `json:"autoscaling"`
	// Size: Size (number of nodes) of the pool.
	Size uint32 `json:"size"`
	// MinSize: Defines the minimum size of the pool. Note that this field is only used when autoscaling is enabled on the pool.
	MinSize uint32 `json:"min_size"`
	// MaxSize: Defines the maximum size of the pool. Note that this field is only used when autoscaling is enabled on the pool.
	MaxSize uint32 `json:"max_size"`
	// ContainerRuntime: Customization of the container runtime is available for each pool. Note that `docker` has been deprecated since version 1.20 and will be removed by version 1.24.
	ContainerRuntime Runtime `json:"container_runtime"`
	// Autohealing: Defines whether the autohealing feature is enabled for the pool.
	Autohealing bool `json:"autohealing"`
	// Tags: Tags associated with the pool.
	Tags []string `json:"tags"`
	// PlacementGroupID: Placement group ID in which all the nodes of the pool will be created.
	PlacementGroupID *string `json:"placement_group_id,omitempty"`
	// KubeletArgs: Kubelet arguments to be used by this pool. Note that this feature is experimental.
	KubeletArgs map[string]string `json:"kubelet_args"`
	// UpgradePolicy: Pool upgrade policy.
	UpgradePolicy *PoolUpgradePolicy `json:"upgrade_policy"`
	// Zone: Zone in which the pool's nodes will be spawned.
	Zone scw.Zone `json:"zone"`
	// RootVolumeType: Defines the system volume disk type. Two different types of volume (`volume_type`) are provided: `l_ssd` is a local block storage which means your system is stored locally on your node's hypervisor. `b_ssd` is a remote block storage which means your system is stored on a centralized and resilient cluster.
	RootVolumeType PoolVolumeType `json:"root_volume_type"`
	// RootVolumeSize: System volume disk size.
	RootVolumeSize *scw.Size `json:"root_volume_size,omitempty"`
	// Region: Cluster region of the pool.
	Region scw.Region `json:"region"`
}

type CreateClusterRequestAutoUpgrade struct {
	// Enable: Defines whether auto upgrade is enabled for the cluster.
	Enable bool `json:"enable"`
	// MaintenanceWindow: Maintenance window of the cluster auto upgrades.
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window"`
}

type CreateClusterRequestAutoscalerConfig struct {
	// ScaleDownDisabled: Disable the cluster autoscaler.
	ScaleDownDisabled *bool `json:"scale_down_disabled,omitempty"`
	// ScaleDownDelayAfterAdd: How long after scale up that scale down evaluation resumes.
	ScaleDownDelayAfterAdd *string `json:"scale_down_delay_after_add,omitempty"`
	// Estimator: Type of resource estimator to be used in scale up.
	Estimator AutoscalerEstimator `json:"estimator"`
	// Expander: Type of node group expander to be used in scale up.
	Expander AutoscalerExpander `json:"expander"`
	// IgnoreDaemonsetsUtilization: Ignore DaemonSet pods when calculating resource utilization for scaling down.
	IgnoreDaemonsetsUtilization *bool `json:"ignore_daemonsets_utilization,omitempty"`
	// BalanceSimilarNodeGroups: Detect similar node groups and balance the number of nodes between them.
	BalanceSimilarNodeGroups *bool `json:"balance_similar_node_groups,omitempty"`
	// ExpendablePodsPriorityCutoff: Pods with priority below cutoff will be expendable. They can be killed without any consideration during scale down and they won't cause scale up. Pods with null priority (PodPriority disabled) are non expendable.
	ExpendablePodsPriorityCutoff *int32 `json:"expendable_pods_priority_cutoff,omitempty"`
	// ScaleDownUnneededTime: How long a node should be unneeded before it is eligible to be scaled down.
	ScaleDownUnneededTime *string `json:"scale_down_unneeded_time,omitempty"`
	// ScaleDownUtilizationThreshold: Node utilization level, defined as a sum of requested resources divided by capacity, below which a node can be considered for scale down.
	ScaleDownUtilizationThreshold *float32 `json:"scale_down_utilization_threshold,omitempty"`
	// MaxGracefulTerminationSec: Maximum number of seconds the cluster autoscaler waits for pod termination when trying to scale down a node.
	MaxGracefulTerminationSec *uint32 `json:"max_graceful_termination_sec,omitempty"`
}

type CreateClusterRequestOpenIDConnectConfig struct {
	// IssuerURL: URL of the provider which allows the API server to discover public signing keys. Only URLs using the `https://` scheme are accepted. This is typically the provider's discovery URL without a path, for example "https://accounts.google.com" or "https://login.salesforce.com".
	IssuerURL string `json:"issuer_url"`
	// ClientID: A client ID that all tokens must be issued for.
	ClientID string `json:"client_id"`
	// UsernameClaim: JWT claim to use as the user name. The default is `sub`, which is expected to be the end user's unique identifier. Admins can choose other claims, such as `email` or `name`, depending on their provider. However, claims other than `email` will be prefixed with the issuer URL to prevent name collision.
	UsernameClaim *string `json:"username_claim,omitempty"`
	// UsernamePrefix: Prefix prepended to username claims to prevent name collision (such as `system:` users). For example, the value `oidc:` will create usernames like `oidc:jane.doe`. If this flag is not provided and `username_claim` is a value other than `email`, the prefix defaults to `( Issuer URL )#` where `( Issuer URL )` is the value of `issuer_url`. The value `-` can be used to disable all prefixing.
	UsernamePrefix *string `json:"username_prefix,omitempty"`
	// GroupsClaim: JWT claim to use as the user's group.
	GroupsClaim *[]string `json:"groups_claim,omitempty"`
	// GroupsPrefix: Prefix prepended to group claims to prevent name collision (such as `system:` groups). For example, the value `oidc:` will create group names like `oidc:engineering` and `oidc:infra`.
	GroupsPrefix *string `json:"groups_prefix,omitempty"`
	// RequiredClaim: Multiple key=value pairs describing a required claim in the ID token. If set, the claims are verified to be present in the ID token with a matching value.
	RequiredClaim *[]string `json:"required_claim,omitempty"`
}

type CreateClusterRequestPoolConfig struct {
	// Name: Name of the pool.
	Name string `json:"name"`
	// NodeType: Node type is the type of Scaleway Instance wanted for the pool. Nodes with insufficient memory are not eligible (DEV1-S, PLAY2-PICO, STARDUST). 'external' is a special node type used to provision instances from other cloud providers in a Kosmos Cluster.
	NodeType string `json:"node_type"`
	// PlacementGroupID: Placement group ID in which all the nodes of the pool will be created.
	PlacementGroupID *string `json:"placement_group_id,omitempty"`
	// Autoscaling: Defines whether the autoscaling feature is enabled for the pool.
	Autoscaling bool `json:"autoscaling"`
	// Size: Size (number of nodes) of the pool.
	Size uint32 `json:"size"`
	// MinSize: Defines the minimum size of the pool. Note that this field is only used when autoscaling is enabled on the pool.
	MinSize *uint32 `json:"min_size,omitempty"`
	// MaxSize: Defines the maximum size of the pool. Note that this field is only used when autoscaling is enabled on the pool.
	MaxSize *uint32 `json:"max_size,omitempty"`
	// ContainerRuntime: Customization of the container runtime is available for each pool. Note that `docker` has been deprecated since version 1.20 and will be removed by version 1.24.
	ContainerRuntime Runtime `json:"container_runtime"`
	// Autohealing: Defines whether the autohealing feature is enabled for the pool.
	Autohealing bool `json:"autohealing"`
	// Tags: Tags associated with the pool.
	Tags []string `json:"tags"`
	// KubeletArgs: Kubelet arguments to be used by this pool. Note that this feature is experimental.
	KubeletArgs map[string]string `json:"kubelet_args"`
	// UpgradePolicy: Pool upgrade policy.
	UpgradePolicy *CreateClusterRequestPoolConfigUpgradePolicy `json:"upgrade_policy"`
	// Zone: Zone in which the pool's nodes will be spawned.
	Zone scw.Zone `json:"zone"`
	// RootVolumeType: Defines the system volume disk type. Two different types of volume (`volume_type`) are provided: `l_ssd` is a local block storage which means your system is stored locally on your node's hypervisor. `b_ssd` is a remote block storage which means your system is stored on a centralized and resilient cluster.
	RootVolumeType PoolVolumeType `json:"root_volume_type"`
	// RootVolumeSize: System volume disk size.
	RootVolumeSize *scw.Size `json:"root_volume_size,omitempty"`
}

type CreatePoolRequestUpgradePolicy struct {
	// MaxUnavailable:
	MaxUnavailable *uint32 `json:"max_unavailable,omitempty"`
	// MaxSurge:
	MaxSurge *uint32 `json:"max_surge,omitempty"`
}

type ClusterType struct {
	// Name: Cluster type name.
	Name string `json:"name"`
	// Availability: Cluster type availability.
	Availability ClusterTypeAvailability `json:"availability"`
	// MaxNodes: Maximum number of nodes supported by the offer.
	MaxNodes uint32 `json:"max_nodes"`
	// CommitmentDelay: Time period during which you can no longer switch to a lower offer.
	CommitmentDelay *scw.Duration `json:"commitment_delay,omitempty"`
	// SLA: Value of the Service Level Agreement of the offer.
	SLA float32 `json:"sla"`
	// Resiliency: Resiliency offered by the offer.
	Resiliency ClusterTypeResiliency `json:"resiliency"`
	// Memory: Max RAM allowed for the control plane.
	Memory scw.Size `json:"memory"`
	// Dedicated: Returns information if this offer uses dedicated resources.
	Dedicated bool `json:"dedicated"`
}

type Version struct {
	// Name: Name of the Kubernetes version.
	Name string `json:"name"`
	// Label: Label of the Kubernetes version.
	Label string `json:"label"`
	// Region: Region in which this version is available.
	Region scw.Region `json:"region"`
	// AvailableCnis: Supported Container Network Interface (CNI) plugins for this version.
	AvailableCnis []CNI `json:"available_cnis"`
	// AvailableIngresses: Supported Ingress Controllers for this version.
	AvailableIngresses *[]Ingress `json:"available_ingresses,omitempty"`
	// AvailableContainerRuntimes: Supported container runtimes for this version.
	AvailableContainerRuntimes []Runtime `json:"available_container_runtimes"`
	// AvailableFeatureGates: Supported feature gates for this version.
	AvailableFeatureGates []string `json:"available_feature_gates"`
	// AvailableAdmissionPlugins: Supported admission plugins for this version.
	AvailableAdmissionPlugins []string `json:"available_admission_plugins"`
	// AvailableKubeletArgs: Supported kubelet arguments for this version.
	AvailableKubeletArgs map[string]string `json:"available_kubelet_args"`
}

type Cluster struct {
	// ID: Cluster ID.
	ID string `json:"id"`
	// Type: Cluster type.
	Type string `json:"type"`
	// Name: Cluster name.
	Name string `json:"name"`
	// Status: Status of the cluster.
	Status ClusterStatus `json:"status"`
	// Version: Kubernetes version of the cluster.
	Version string `json:"version"`
	// Region: Region in which the cluster is deployed.
	Region scw.Region `json:"region"`
	// OrganizationID: ID of the Organization owning the cluster.
	OrganizationID string `json:"organization_id"`
	// ProjectID: ID of the Project owning the cluster.
	ProjectID string `json:"project_id"`
	// Tags: Tags associated with the cluster.
	Tags []string `json:"tags"`
	// Cni: Container Network Interface (CNI) plugin running in the cluster.
	Cni CNI `json:"cni"`
	// Description: Cluster description.
	Description string `json:"description"`
	// ClusterURL: Kubernetes API server URL of the cluster.
	ClusterURL string `json:"cluster_url"`
	// DNSWildcard: Wildcard DNS resolving all the ready cluster nodes.
	DNSWildcard string `json:"dns_wildcard"`
	// CreatedAt: Date on which the cluster was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Date on which the cluster was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// AutoscalerConfig: Autoscaler config for the cluster.
	AutoscalerConfig *ClusterAutoscalerConfig `json:"autoscaler_config"`
	// DashboardEnabled: Defines whether the Kubernetes dashboard is enabled for the cluster.
	DashboardEnabled *bool `json:"dashboard_enabled,omitempty"`
	// Ingress: Managed Ingress controller used in the cluster (deprecated feature).
	Ingress *Ingress `json:"ingress,omitempty"`
	// AutoUpgrade: Auto upgrade configuration of the cluster.
	AutoUpgrade *ClusterAutoUpgrade `json:"auto_upgrade"`
	// UpgradeAvailable: Defines whether a new Kubernetes version is available.
	UpgradeAvailable bool `json:"upgrade_available"`
	// FeatureGates: List of enabled feature gates.
	FeatureGates []string `json:"feature_gates"`
	// AdmissionPlugins: List of enabled admission plugins.
	AdmissionPlugins []string `json:"admission_plugins"`
	// OpenIDConnectConfig: This configuration enables to update the OpenID Connect configuration of the Kubernetes API server.
	OpenIDConnectConfig *ClusterOpenIDConnectConfig `json:"open_id_connect_config"`
	// ApiserverCertSans: Additional Subject Alternative Names for the Kubernetes API server certificate.
	ApiserverCertSans []string `json:"apiserver_cert_sans"`
	// PrivateNetworkID: Private network ID for internal cluster communication.
	PrivateNetworkID *string `json:"private_network_id,omitempty"`
	// CommitmentEndsAt: Date on which it will be possible to switch to a smaller offer.
	CommitmentEndsAt *time.Time `json:"commitment_ends_at,omitempty"`
	// AuditLog: Enables the Kubernetes audit logging feature. Audit logs are sent to Cockpit. Please note that this feature is only available for clusters with a Dedicated Control Plane (https://www.scaleway.com/en/kubernetes-dedicated-control-plane/).
	AuditLog bool `json:"audit_log"`
}

type Node struct {
	// ID: Node ID.
	ID string `json:"id"`
	// PoolID: Pool ID of the node.
	PoolID string `json:"pool_id"`
	// ClusterID: Cluster ID of the node.
	ClusterID string `json:"cluster_id"`
	// ProviderID: Underlying instance ID. It is prefixed by instance type and location information (see https://pkg.go.dev/k8s.io/api/core/v1#NodeSpec.ProviderID).
	ProviderID string `json:"provider_id"`
	// Region: Cluster region of the node.
	Region scw.Region `json:"region"`
	// Name: Name of the node.
	Name string `json:"name"`
	// PublicIPV4: Public IPv4 address of the node.
	PublicIPV4 *net.IP `json:"public_ip_v4,omitempty"`
	// PublicIPV6: Public IPv6 address of the node.
	PublicIPV6 *net.IP `json:"public_ip_v6,omitempty"`
	// Conditions: Conditions of the node. These conditions contain the Node Problem Detector conditions, as well as some in house conditions.
	Conditions *map[string]string `json:"conditions,omitempty"`
	// Status: Status of the node.
	Status NodeStatus `json:"status"`
	// ErrorMessage: Details of the error, if any occurred when managing the node.
	ErrorMessage *string `json:"error_message,omitempty"`
	// CreatedAt: Date on which the node was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Date on which the node was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type UpdateClusterRequestAutoUpgrade struct {
	// Enable: Defines whether auto upgrade is enabled for the cluster.
	Enable *bool `json:"enable,omitempty"`
	// MaintenanceWindow: Maintenance window of the cluster auto upgrades.
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window"`
}

type UpdateClusterRequestAutoscalerConfig struct {
	// ScaleDownDisabled: Disable the cluster autoscaler.
	ScaleDownDisabled *bool `json:"scale_down_disabled,omitempty"`
	// ScaleDownDelayAfterAdd: How long after scale up that scale down evaluation resumes.
	ScaleDownDelayAfterAdd *string `json:"scale_down_delay_after_add,omitempty"`
	// Estimator: Type of resource estimator to be used in scale up.
	Estimator AutoscalerEstimator `json:"estimator"`
	// Expander: Type of node group expander to be used in scale up.
	Expander AutoscalerExpander `json:"expander"`
	// IgnoreDaemonsetsUtilization: Ignore DaemonSet pods when calculating resource utilization for scaling down.
	IgnoreDaemonsetsUtilization *bool `json:"ignore_daemonsets_utilization,omitempty"`
	// BalanceSimilarNodeGroups: Detect similar node groups and balance the number of nodes between them.
	BalanceSimilarNodeGroups *bool `json:"balance_similar_node_groups,omitempty"`
	// ExpendablePodsPriorityCutoff: Pods with priority below cutoff will be expendable. They can be killed without any consideration during scale down and they won't cause scale up. Pods with null priority (PodPriority disabled) are non expendable.
	ExpendablePodsPriorityCutoff *int32 `json:"expendable_pods_priority_cutoff,omitempty"`
	// ScaleDownUnneededTime: How long a node should be unneeded before it is eligible to be scaled down.
	ScaleDownUnneededTime *string `json:"scale_down_unneeded_time,omitempty"`
	// ScaleDownUtilizationThreshold: Node utilization level, defined as a sum of requested resources divided by capacity, below which a node can be considered for scale down.
	ScaleDownUtilizationThreshold *float32 `json:"scale_down_utilization_threshold,omitempty"`
	// MaxGracefulTerminationSec: Maximum number of seconds the cluster autoscaler waits for pod termination when trying to scale down a node.
	MaxGracefulTerminationSec *uint32 `json:"max_graceful_termination_sec,omitempty"`
}

type UpdateClusterRequestOpenIDConnectConfig struct {
	// IssuerURL: URL of the provider which allows the API server to discover public signing keys. Only URLs using the `https://` scheme are accepted. This is typically the provider's discovery URL without a path, for example "https://accounts.google.com" or "https://login.salesforce.com".
	IssuerURL *string `json:"issuer_url,omitempty"`
	// ClientID: A client ID that all tokens must be issued for.
	ClientID *string `json:"client_id,omitempty"`
	// UsernameClaim: JWT claim to use as the user name. The default is `sub`, which is expected to be the end user's unique identifier. Admins can choose other claims, such as `email` or `name`, depending on their provider. However, claims other than `email` will be prefixed with the issuer URL to prevent name collision.
	UsernameClaim *string `json:"username_claim,omitempty"`
	// UsernamePrefix: Prefix prepended to username claims to prevent name collision (such as `system:` users). For example, the value `oidc:` will create usernames like `oidc:jane.doe`. If this flag is not provided and `username_claim` is a value other than `email`, the prefix defaults to `( Issuer URL )#` where `( Issuer URL )` is the value of `issuer_url`. The value `-` can be used to disable all prefixing.
	UsernamePrefix *string `json:"username_prefix,omitempty"`
	// GroupsClaim: JWT claim to use as the user's group.
	GroupsClaim *[]string `json:"groups_claim,omitempty"`
	// GroupsPrefix: Prefix prepended to group claims to prevent name collision (such as `system:` groups). For example, the value `oidc:` will create group names like `oidc:engineering` and `oidc:infra`.
	GroupsPrefix *string `json:"groups_prefix,omitempty"`
	// RequiredClaim: Multiple key=value pairs describing a required claim in the ID token. If set, the claims are verified to be present in the ID token with a matching value.
	RequiredClaim *[]string `json:"required_claim,omitempty"`
}

type UpdatePoolRequestUpgradePolicy struct {
	// MaxUnavailable:
	MaxUnavailable *uint32 `json:"max_unavailable,omitempty"`
	// MaxSurge:
	MaxSurge *uint32 `json:"max_surge,omitempty"`
}

type CreateClusterRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// OrganizationID: Organization ID in which the cluster will be created.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: Project ID in which the cluster will be created.
	ProjectID *string `json:"project_id,omitempty"`
	// Type: Type of the cluster (possible values are kapsule, multicloud, kapsule-dedicated-8, kapsule-dedicated-16).
	Type string `json:"type"`
	// Name: Cluster name.
	Name string `json:"name"`
	// Description: Cluster description.
	Description string `json:"description"`
	// Tags: Tags associated with the cluster.
	Tags []string `json:"tags"`
	// Version: Kubernetes version of the cluster.
	Version string `json:"version"`
	// Cni: Container Network Interface (CNI) plugin running in the cluster.
	Cni CNI `json:"cni"`
	// EnableDashboard: Defines whether the Kubernetes Dashboard is enabled in the cluster.
	EnableDashboard *bool `json:"enable_dashboard,omitempty"`
	// Ingress: Ingress Controller running in the cluster (deprecated feature).
	Ingress *Ingress `json:"ingress,omitempty"`
	// Pools: Pools created along with the cluster.
	Pools []*CreateClusterRequestPoolConfig `json:"pools"`
	// AutoscalerConfig: Autoscaler configuration for the cluster. It allows you to set (to an extent) your preferred autoscaler configuration, which is an implementation of the cluster-autoscaler (https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler/).
	AutoscalerConfig *CreateClusterRequestAutoscalerConfig `json:"autoscaler_config"`
	// AutoUpgrade: Auto upgrade configuration of the cluster. This configuration enables to set a specific 2-hour time window in which the cluster can be automatically updated to the latest patch version.
	AutoUpgrade *CreateClusterRequestAutoUpgrade `json:"auto_upgrade"`
	// FeatureGates: List of feature gates to enable.
	FeatureGates []string `json:"feature_gates"`
	// AdmissionPlugins: List of admission plugins to enable.
	AdmissionPlugins []string `json:"admission_plugins"`
	// OpenIDConnectConfig: OpenID Connect configuration of the cluster. This configuration enables to update the OpenID Connect configuration of the Kubernetes API server.
	OpenIDConnectConfig *CreateClusterRequestOpenIDConnectConfig `json:"open_id_connect_config"`
	// ApiserverCertSans: Additional Subject Alternative Names for the Kubernetes API server certificate.
	ApiserverCertSans []string `json:"apiserver_cert_sans"`
	// PrivateNetworkID: Private network ID for internal cluster communication (cannot be changed later).
	PrivateNetworkID *string `json:"private_network_id,omitempty"`
	// AuditLog: Enables the Kubernetes audit logging feature. Audit logs are sent to Cockpit. Please note that this feature is only available for clusters with a Dedicated Control Plane (https://www.scaleway.com/en/kubernetes-dedicated-control-plane/).
	AuditLog bool `json:"audit_log"`
}

type CreateExternalNodeRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// PoolID:
	PoolID string `json:"-"`
}

type CreatePoolRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ClusterID: Cluster ID to which the pool will be attached.
	ClusterID string `json:"-"`
	// Name: Pool name.
	Name string `json:"name"`
	// NodeType: Node type is the type of Scaleway Instance wanted for the pool. Nodes with insufficient memory are not eligible (DEV1-S, PLAY2-PICO, STARDUST). 'external' is a special node type used to provision instances from other cloud providers in a Kosmos Cluster.
	NodeType string `json:"node_type"`
	// PlacementGroupID: Placement group ID in which all the nodes of the pool will be created.
	PlacementGroupID *string `json:"placement_group_id,omitempty"`
	// Autoscaling: Defines whether the autoscaling feature is enabled for the pool.
	Autoscaling bool `json:"autoscaling"`
	// Size: Size (number of nodes) of the pool.
	Size uint32 `json:"size"`
	// MinSize: Defines the minimum size of the pool. Note that this field is only used when autoscaling is enabled on the pool.
	MinSize *uint32 `json:"min_size,omitempty"`
	// MaxSize: Defines the maximum size of the pool. Note that this field is only used when autoscaling is enabled on the pool.
	MaxSize *uint32 `json:"max_size,omitempty"`
	// ContainerRuntime: Customization of the container runtime is available for each pool. Note that `docker` has been deprecated since version 1.20 and will be removed by version 1.24.
	ContainerRuntime Runtime `json:"container_runtime"`
	// Autohealing: Defines whether the autohealing feature is enabled for the pool.
	Autohealing bool `json:"autohealing"`
	// Tags: Tags associated with the pool.
	Tags []string `json:"tags"`
	// KubeletArgs: Kubelet arguments to be used by this pool. Note that this feature is experimental.
	KubeletArgs map[string]string `json:"kubelet_args"`
	// UpgradePolicy: Pool upgrade policy.
	UpgradePolicy *CreatePoolRequestUpgradePolicy `json:"upgrade_policy"`
	// Zone: Zone in which the pool's nodes will be spawned.
	Zone scw.Zone `json:"zone"`
	// RootVolumeType: Defines the system volume disk type. Two different types of volume (`volume_type`) are provided: `l_ssd` is a local block storage which means your system is stored locally on your node's hypervisor. `b_ssd` is a remote block storage which means your system is stored on a centralized and resilient cluster.
	RootVolumeType PoolVolumeType `json:"root_volume_type"`
	// RootVolumeSize: System volume disk size.
	RootVolumeSize *scw.Size `json:"root_volume_size,omitempty"`
}

type DeleteClusterRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ClusterID: ID of the cluster to delete.
	ClusterID string `json:"-"`
	// WithAdditionalResources: Defines whether all volumes (including retain volume type), empty Private Networks and Load Balancers with a name starting with the cluster ID will also be deleted.
	WithAdditionalResources bool `json:"with_additional_resources"`
}

type DeleteNodeRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NodeID: ID of the node to replace.
	NodeID string `json:"-"`
	// SkipDrain: Skip draining node from its workload.
	SkipDrain bool `json:"skip_drain"`
	// Replace: Add a new node after the deletion of this node.
	Replace bool `json:"replace"`
}

type DeletePoolRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// PoolID: ID of the pool to delete.
	PoolID string `json:"-"`
}

type ExternalNode struct {
	// ID:
	ID string `json:"id"`
	// Name:
	Name string `json:"name"`
	// ClusterURL:
	ClusterURL string `json:"cluster_url"`
	// PoolVersion:
	PoolVersion string `json:"pool_version"`
	// ClusterCa:
	ClusterCa string `json:"cluster_ca"`
	// KubeToken:
	KubeToken string `json:"kube_token"`
	// KubeletConfig:
	KubeletConfig string `json:"kubelet_config"`
	// ExternalIP:
	ExternalIP string `json:"external_ip"`
}

type GetClusterKubeConfigRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ClusterID: Cluster ID for which to download the kubeconfig.
	ClusterID string `json:"-"`
}

type GetClusterRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ClusterID: ID of the requested cluster.
	ClusterID string `json:"-"`
}

type GetNodeRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NodeID: ID of the requested node.
	NodeID string `json:"-"`
}

type GetPoolRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// PoolID: ID of the requested pool.
	PoolID string `json:"-"`
}

type GetVersionRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// VersionName: Requested version name.
	VersionName string `json:"-"`
}

type ListClusterAvailableTypesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ClusterID: Cluster ID for which the available Kubernetes types will be listed.
	ClusterID string `json:"-"`
}

type ListClusterAvailableTypesResponse struct {
	// ClusterTypes: Available cluster types for the cluster.
	ClusterTypes []*ClusterType `json:"cluster_types"`
	// TotalCount: Total number of types.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListClusterAvailableTypesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListClusterAvailableTypesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListClusterAvailableTypesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ClusterTypes = append(r.ClusterTypes, results.ClusterTypes...)
	r.TotalCount += uint32(len(results.ClusterTypes))
	return uint32(len(results.ClusterTypes)), nil
}

type ListClusterAvailableVersionsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ClusterID: Cluster ID for which the available Kubernetes versions will be listed.
	ClusterID string `json:"-"`
}

type ListClusterAvailableVersionsResponse struct {
	// Versions: Available Kubernetes versions for the cluster.
	Versions []*Version `json:"versions"`
}

type ListClusterTypesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Page: Page number, from the paginated results, to return for cluster-types.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Maximum number of clusters per page.
	PageSize *uint32 `json:"page_size,omitempty"`
}

type ListClusterTypesResponse struct {
	// TotalCount: Total number of cluster-types.
	TotalCount uint32 `json:"total_count"`
	// ClusterTypes: Paginated returned cluster-types.
	ClusterTypes []*ClusterType `json:"cluster_types"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListClusterTypesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListClusterTypesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListClusterTypesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ClusterTypes = append(r.ClusterTypes, results.ClusterTypes...)
	r.TotalCount += uint32(len(results.ClusterTypes))
	return uint32(len(results.ClusterTypes)), nil
}

type ListClustersRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// OrganizationID: Organization ID on which to filter the returned clusters.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: Project ID on which to filter the returned clusters.
	ProjectID *string `json:"project_id,omitempty"`
	// OrderBy: Sort order of returned clusters.
	OrderBy ListClustersRequestOrderBy `json:"order_by"`
	// Page: Page number to return for clusters, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Maximum number of clusters per page.
	PageSize *uint32 `json:"page_size,omitempty"`
	// Name: Name to filter on, only clusters containing this substring in their name will be returned.
	Name *string `json:"name,omitempty"`
	// Status: Status to filter on, only clusters with this status will be returned.
	Status ClusterStatus `json:"status"`
	// Type: Type to filter on, only clusters with this type will be returned.
	Type *string `json:"type,omitempty"`
}

type ListClustersResponse struct {
	// TotalCount: Total number of clusters.
	TotalCount uint32 `json:"total_count"`
	// Clusters: Paginated returned clusters.
	Clusters []*Cluster `json:"clusters"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListClustersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListClustersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListClustersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Clusters = append(r.Clusters, results.Clusters...)
	r.TotalCount += uint32(len(results.Clusters))
	return uint32(len(results.Clusters)), nil
}

type ListNodesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ClusterID: Cluster ID from which the nodes will be listed from.
	ClusterID string `json:"-"`
	// PoolID: Pool ID on which to filter the returned nodes.
	PoolID *string `json:"pool_id,omitempty"`
	// OrderBy: Sort order of the returned nodes.
	OrderBy ListNodesRequestOrderBy `json:"order_by"`
	// Page: Page number for the returned nodes.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Maximum number of nodes per page.
	PageSize *uint32 `json:"page_size,omitempty"`
	// Name: Name to filter on, only nodes containing this substring in their name will be returned.
	Name *string `json:"name,omitempty"`
	// Status: Status to filter on, only nodes with this status will be returned.
	Status NodeStatus `json:"status"`
}

type ListNodesResponse struct {
	// TotalCount: Total number of nodes.
	TotalCount uint32 `json:"total_count"`
	// Nodes: Paginated returned nodes.
	Nodes []*Node `json:"nodes"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNodesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNodesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListNodesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Nodes = append(r.Nodes, results.Nodes...)
	r.TotalCount += uint32(len(results.Nodes))
	return uint32(len(results.Nodes)), nil
}

type ListPoolsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ClusterID: ID of the cluster whose pools will be listed.
	ClusterID string `json:"-"`
	// OrderBy: Sort order of returned pools.
	OrderBy ListPoolsRequestOrderBy `json:"order_by"`
	// Page: Page number for the returned pools.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Maximum number of pools per page.
	PageSize *uint32 `json:"page_size,omitempty"`
	// Name: Name to filter on, only pools containing this substring in their name will be returned.
	Name *string `json:"name,omitempty"`
	// Status: Status to filter on, only pools with this status will be returned.
	Status PoolStatus `json:"status"`
}

type ListPoolsResponse struct {
	// TotalCount: Total number of pools that exists for the cluster.
	TotalCount uint32 `json:"total_count"`
	// Pools: Paginated returned pools.
	Pools []*Pool `json:"pools"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPoolsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPoolsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListPoolsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Pools = append(r.Pools, results.Pools...)
	r.TotalCount += uint32(len(results.Pools))
	return uint32(len(results.Pools)), nil
}

type ListVersionsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
}

type ListVersionsResponse struct {
	// Versions: Available Kubernetes versions.
	Versions []*Version `json:"versions"`
}

type MigrateToPrivateNetworkClusterRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ClusterID: ID of the cluster to migrate.
	ClusterID string `json:"-"`
	// PrivateNetworkID: ID of the Private Network to link to the cluster.
	PrivateNetworkID string `json:"private_network_id"`
}

type RebootNodeRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NodeID: ID of the node to reboot.
	NodeID string `json:"-"`
}

type ReplaceNodeRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NodeID: ID of the node to replace.
	NodeID string `json:"-"`
}

type ResetClusterAdminTokenRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ClusterID: Cluster ID on which the admin token will be renewed.
	ClusterID string `json:"-"`
}

type SetClusterTypeRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ClusterID: ID of the cluster to migrate from one type to another.
	ClusterID string `json:"-"`
	// Type: Type of the cluster. Note that some migrations are not possible (please refer to product documentation).
	Type string `json:"type"`
}

type UpdateClusterRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ClusterID: ID of the cluster to update.
	ClusterID string `json:"-"`
	// Name: New external name for the cluster.
	Name *string `json:"name,omitempty"`
	// Description: New description for the cluster.
	Description *string `json:"description,omitempty"`
	// Tags: New tags associated with the cluster.
	Tags *[]string `json:"tags,omitempty"`
	// AutoscalerConfig: New autoscaler config for the cluster.
	AutoscalerConfig *UpdateClusterRequestAutoscalerConfig `json:"autoscaler_config"`
	// EnableDashboard: New value for the Kubernetes Dashboard enablement.
	EnableDashboard *bool `json:"enable_dashboard,omitempty"`
	// Ingress: New Ingress Controller for the cluster (deprecated feature).
	Ingress *Ingress `json:"ingress,omitempty"`
	// AutoUpgrade: New auto upgrade configuration for the cluster. Note that all fields need to be set.
	AutoUpgrade *UpdateClusterRequestAutoUpgrade `json:"auto_upgrade"`
	// FeatureGates: List of feature gates to enable.
	FeatureGates *[]string `json:"feature_gates,omitempty"`
	// AdmissionPlugins: List of admission plugins to enable.
	AdmissionPlugins *[]string `json:"admission_plugins,omitempty"`
	// OpenIDConnectConfig: OpenID Connect configuration of the cluster. This configuration enables to update the OpenID Connect configuration of the Kubernetes API server.
	OpenIDConnectConfig *UpdateClusterRequestOpenIDConnectConfig `json:"open_id_connect_config"`
	// ApiserverCertSans: Additional Subject Alternative Names for the Kubernetes API server certificate.
	ApiserverCertSans *[]string `json:"apiserver_cert_sans,omitempty"`
	// AuditLog: Enables the Kubernetes audit logging feature. Audit logs are sent to Cockpit. Please note that this feature is only available for clusters with a Dedicated Control Plane (https://www.scaleway.com/en/kubernetes-dedicated-control-plane/).
	AuditLog *bool `json:"audit_log,omitempty"`
}

type UpdatePoolRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// PoolID: ID of the pool to update.
	PoolID string `json:"-"`
	// Autoscaling: New value for the pool autoscaling enablement.
	Autoscaling *bool `json:"autoscaling,omitempty"`
	// Size: New desired pool size.
	Size *uint32 `json:"size,omitempty"`
	// MinSize: New minimum size for the pool.
	MinSize *uint32 `json:"min_size,omitempty"`
	// MaxSize: New maximum size for the pool.
	MaxSize *uint32 `json:"max_size,omitempty"`
	// Autohealing: New value for the pool autohealing enablement.
	Autohealing *bool `json:"autohealing,omitempty"`
	// Tags: New tags associated with the pool.
	Tags *[]string `json:"tags,omitempty"`
	// KubeletArgs: New Kubelet arguments to be used by this pool. Note that this feature is experimental.
	KubeletArgs *map[string]string `json:"kubelet_args,omitempty"`
	// UpgradePolicy: New upgrade policy for the pool.
	UpgradePolicy *UpdatePoolRequestUpgradePolicy `json:"upgrade_policy"`
}

type UpgradeClusterRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ClusterID: ID of the cluster to upgrade.
	ClusterID string `json:"-"`
	// Version: New Kubernetes version of the cluster. Note that the version should either be a higher patch version of the same minor version or the direct minor version after the current one.
	Version string `json:"version"`
	// UpgradePools: Defines whether pools will also be upgraded once the control plane is upgraded.
	UpgradePools bool `json:"upgrade_pools"`
}

type UpgradePoolRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// PoolID: ID of the pool to upgrade.
	PoolID string `json:"-"`
	// Version: New Kubernetes version for the pool.
	Version string `json:"version"`
}

// Kubernetes is an open-source platform that enables developers to manage their containerized applications. Scaleway Kubernetes Kapsule and Kosmos are powerful tools to help you manage your containerized workloads and services.
//
// They both provide a managed environment for creating, configuring, and running clusters of pre-configured machines.
//
// The primary difference between Kapsule and Kosmos is that Kapsule clusters are composed solely of Scaleway Instances. In contrast, Kosmos is a managed Multi-Cloud Kubernetes Engine that allows you to connect Instances and virtual or dedicated servers from any cloud provider to a single managed Control-Plane.
//
// ## Concepts
//
// Refer to our [dedicated concepts page](https://www.scaleway.com/en/docs/containers/kubernetes/concepts/) to find definitions of all Kubetnetes-related terminology.
//
// ## Quickstart
//
// (switchcolumn)
// (switchcolumn)
//
//  1. Configure your environment variables. Note: This is an optional step that seeks to simplify your usage of the Kapsule and Kosmos API.
//     ```bash
//     export SCW_ACCESS_KEY="<API access key>"
//     export SCW_SECRET_KEY="<API secret key>"
//     export SCW_REGION="<Scaleway service region>"
//     export SCW_PROJECT_ID="<Scaleway Project ID>"
//     ```
//
//  2. Edit the POST request payload you will use to create your Kubernetes cluster. Replace the parameters in the following example:
//     ```json
//     {
//     "project_id": "$SCW_PROJECT_ID",
//     "type": "string",
//     "name": "string",
//     "description": "string",
//     "tags": [
//     "string"
//     ],
//     "version": "string",
//     "cni": "unknown_cni",
//     "pools": [
//     {
//     "name": "string",
//     "node_type": "string",
//     "size": "integer",
//     "container_runtime": "unknown_runtime",
//     "tags": [
//     "string"
//     ],
//
//     "zone": "string",
//     "root_volume_type": "default_volume_type",
//     "root_volume_size": "integer"
//     }
//     ]
//     }
//     ```
//
//     | Parameter        | Description                                                        |
//     | :--------------- | :----------------------------------------------------------------- |
//     | `project_id`     | The ID of the Project you want to create your Kubernetes cluster in. To find your Project ID you can [list the projects](https://www.scaleway.com/en/docs/identity-and-access-management/iam/api-cli/managing-projects#listing-all-your-projects) or consult the [Scaleway console](https://console.scaleway.com/project/settings). |
//     | `type`         | The type of the cluster (possible values are `kapsule`, `multicloud`). |
//     | `name`           | **REQUIRED** Name of the cmister. |
//     | `description`  | Description of the cluster.
//     | `tags`         | Tags associated with the cluster. |
//     | `version`      | **REQUIRED** Kubernetes version of the cluster. |
//     | `cni`          | **REQUIRED** Container Network Interface (CNI) plugin that will run on the cluster. The default value is `unknown_cni`. |
//     | `pools`        | Pools to be created along with the cluster. |
//     | `name`         | **REQUIRED** Name of the pool.
//     | `node-type`    | **REQUIRED** The node type of the Scaleway Instance wanted for the pool.
//     | `size`         | **REQUIRED** The number of nodes in the pool. |
//     | `tags`         | Tags associated with the pool.
//     | `zone`         | Availability zone in which the pools will be deployed in.
//     | `root_volume_type` | The root volume type. The default value is `default_volume_type`. |
//     | `root_volume_size` | The system volume disk size in bytes. |
//
//  3. Create a Kapsule cluster and node pool by running the following command. Make sure you include the payload you edited in the previous step.
//     ```bash
//     curl -X POST \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     "Content-Type: application/json" \
//     "https://api.scaleway.com/k8s/v1/regions/$SCW_REGION/clusters" \
//     -d '{
//     "project_id": "$SCW_REGION",
//     "type": "kapsule",
//     "name": "MyFirstKapsuleCluster",
//     "description": "My first Kapsule Cluster",
//     "tags": [
//     "kapsule",
//     "kubernetes"
//     ],
//     "version": "1.27.0",
//     "cni": "unknown_cni",
//     "pools": [
//     {
//     "name": "MyFirstKapsulePool",
//     "node_type": "PLAY2-MICRO",
//     "size": "2",
//     "container_runtime": "unknown_runtime",
//     "tags": [
//     "pool"
//     ],
//     "zone": "fr-par-1",
//     "root_volume_type": "default_volume_type",
//     "root_volume_size": "20000000000"
//     }
//     ]
//     }`
//     ```
//
//  4. List your Kapsule clusters.
//     ```bash
//     curl -X GET -H "Content-Type: application/json" \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" https://api.scaleway.com/k8s/v1/regions/$SCW_REGION/clusters
//     ```
//
//     You will see detailed information about your clusters.
//
//  5. Download the kubeconfig file for your cluster.
//     ```bash
//     curl -X GET -H "Content-Type: application/json" \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" https://api.scaleway.com/k8s/v1/regions/$SCW_REGION/clusters/$SCW_CLUSTER_ID/kubeconfig
//     ```
//     <Message type="tip">
//     Add `?dl=1` at the end of the URL to directly get the `base64`-decoded kubeconfig. If not, the kubeconfig will be `base64`-encoded.
//     </Message>
//
// 6. [Connect to your cluster](https://www.scaleway.com/en/docs/containers/kubernetes/how-to/connect-cluster-kubectl) using kubectl.
//
//  7. Delete your cluster.
//     ```bash
//     curl -X GET -H "Content-Type: application/json" \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" https://api.scaleway.com/k8s/v1/regions/$SCW_REGION/clusters/$SCW_CLUSTER_ID
//     ```
//
// (switchcolumn)
// <Message type="requirement">
// To perform the following steps, you must first ensure that:
//   - you have an account and are logged into the [Scaleway console](https://console.scaleway.com/organization)
//   - you have created an [API key](https://www.scaleway.com/en/docs/identity-and-access-management/iam/how-to/create-api-keys/) and that the API key has sufficient [IAM permissions](https://www.scaleway.com/en/docs/identity-and-access-management/iam/reference-content/permission-sets/) to perform the actions described on this page.
//   - you have [installed `curl`](https://curl.se/download.html)
//
// </Message>
// (switchcolumn)
// ## Technical information
//
// Kubernetes Kapsule provides features such as:
// - Persistent Volume Claims (PVC) are available through Scaleway Block Volumes.
// - Pool autoscaling and autohealing is available.
// - Kubernetes auto upgrades features is available.
// ### Regions
//
// Scaleway's infrastructure is spread across different [regions and Availability Zones](https://www.scaleway.com/en/docs/console/my-account/reference-content/products-availability/).
//
// Kubernetes Kapsule and Kosmos are available in the Paris, Amsterdam and Warsaw regions, which are represented by the following path parameters:
//
// - `fr-par`
// - `nl-ams`
// - `pl-waw`
//
// ### Versions
//
// Kubernetes Kosmos and Kapsule supports at least the latest version of the last 3 major Kubernetes releases.
// ## Technical limitations
//
// # The following limitations should be taken into account when using the Kubernetes API
//
// - The maximum number of pods per node is 110.
// ## Going further
//
// For more help using Kubenetes Kapsule and Kosmos, check out the following resources:
// - Our [main documentation](https://www.scaleway.com/en/docs/containers/kubernetes/)
// - The #k8s channel on our [Slack Community](https://www.scaleway.com/en/docs/tutorials/scaleway-slack-community/)
// - Our [support ticketing system](https://www.scaleway.com/en/docs/console/my-account/how-to/open-a-support-ticket).
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}
func (s *API) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
}

// ListClusters: List all existing Kubernetes clusters in a specific region.
func (s *API) ListClusters(req *ListClustersRequest, opts ...scw.RequestOption) (*ListClustersResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}
	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "type", req.Type)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters",
		Query:  query,
	}

	var resp ListClustersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateCluster: Create a new Kubernetes cluster in a Scaleway region.
func (s *API) CreateCluster(req *CreateClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.OrganizationID == nil && req.ProjectID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.OrganizationID == nil && req.ProjectID == nil {
		req.ProjectID = &defaultProjectID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("k8s")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCluster: Retrieve information about a specific Kubernetes cluster.
func (s *API) GetCluster(req *GetClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateCluster: Update information on a specific Kubernetes cluster. You can update details such as its name, description, tags and configuration. To upgrade a cluster, you will need to use the dedicated endpoint.
func (s *API) UpdateCluster(req *UpdateClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteCluster: Delete a specific Kubernetes cluster and all its associated pools and nodes. Note that this method will not delete any Load Balancer or Block Volume that are associated with the cluster.
func (s *API) DeleteCluster(req *DeleteClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "with_additional_resources", req.WithAdditionalResources)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
		Query:  query,
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpgradeCluster: Upgrade a specific Kubernetes cluster and possibly its associated pools to a specific and supported Kubernetes version.
func (s *API) UpgradeCluster(req *UpgradeClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/upgrade",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetClusterType: Change the type of a specific Kubernetes cluster.
func (s *API) SetClusterType(req *SetClusterTypeRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/set-type",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListClusterAvailableVersions: List the versions that a specific Kubernetes cluster is allowed to upgrade to. Results will include every patch version greater than the current patch, as well as one minor version ahead of the current version. Any upgrade skipping a minor version will not work.
func (s *API) ListClusterAvailableVersions(req *ListClusterAvailableVersionsRequest, opts ...scw.RequestOption) (*ListClusterAvailableVersionsResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/available-versions",
	}

	var resp ListClusterAvailableVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListClusterAvailableTypes: List the cluster types that a specific Kubernetes cluster is allowed to switch to.
func (s *API) ListClusterAvailableTypes(req *ListClusterAvailableTypesRequest, opts ...scw.RequestOption) (*ListClusterAvailableTypesResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/available-types",
	}

	var resp ListClusterAvailableTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// getClusterKubeConfig: Download the Kubernetes cluster config file (also known as `kubeconfig`) for a specific cluster to use it with `kubectl`.
// Tip: add `?dl=1` at the end of the URL to directly retrieve the base64 decoded kubeconfig. If you choose not to, the kubeconfig will be base64 encoded.
func (s *API) getClusterKubeConfig(req *GetClusterKubeConfigRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/kubeconfig",
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ResetClusterAdminToken: Reset the admin token for a specific Kubernetes cluster. This will revoke the old admin token (which will not be usable afterwards) and create a new one. Note that you will need to download kubeconfig again to keep interacting with the cluster.
func (s *API) ResetClusterAdminToken(req *ResetClusterAdminTokenRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/reset-admin-token",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return err
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// MigrateToPrivateNetworkCluster: Migrate a cluster that was created before the release of Private Network clusters to a new one with a Private Network.
func (s *API) MigrateToPrivateNetworkCluster(req *MigrateToPrivateNetworkClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/migrate-to-private-network",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPools: List all the existing pools for a specific Kubernetes cluster.
func (s *API) ListPools(req *ListPoolsRequest, opts ...scw.RequestOption) (*ListPoolsResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}
	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "status", req.Status)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/pools",
		Query:  query,
	}

	var resp ListPoolsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreatePool: Create a new pool in a specific Kubernetes cluster.
func (s *API) CreatePool(req *CreatePoolRequest, opts ...scw.RequestOption) (*Pool, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("pool")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/pools",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Pool

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPool: Retrieve details about a specific pool in a Kubernetes cluster.
func (s *API) GetPool(req *GetPoolRequest, opts ...scw.RequestOption) (*Pool, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PoolID) == "" {
		return nil, errors.New("field PoolID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/pools/" + fmt.Sprint(req.PoolID) + "",
	}

	var resp Pool

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpgradePool: Upgrade the Kubernetes version of a specific pool. Note that it only works if the targeted version matches the cluster's version.
func (s *API) UpgradePool(req *UpgradePoolRequest, opts ...scw.RequestOption) (*Pool, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PoolID) == "" {
		return nil, errors.New("field PoolID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/pools/" + fmt.Sprint(req.PoolID) + "/upgrade",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Pool

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePool: Update the attributes of a specific pool, such as its desired size, autoscaling settings, and tags.
func (s *API) UpdatePool(req *UpdatePoolRequest, opts ...scw.RequestOption) (*Pool, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PoolID) == "" {
		return nil, errors.New("field PoolID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/pools/" + fmt.Sprint(req.PoolID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Pool

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeletePool: Delete a specific pool from a cluster. Note that all the pool's nodes will also be deleted.
func (s *API) DeletePool(req *DeletePoolRequest, opts ...scw.RequestOption) (*Pool, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PoolID) == "" {
		return nil, errors.New("field PoolID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/pools/" + fmt.Sprint(req.PoolID) + "",
	}

	var resp Pool

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateExternalNode: Retrieve metadata for a Kosmos node. This method is not intended to be called by end users but rather programmatically by the kapsule-node-agent.
func (s *API) CreateExternalNode(req *CreateExternalNodeRequest, opts ...scw.RequestOption) (*ExternalNode, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PoolID) == "" {
		return nil, errors.New("field PoolID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/pools/" + fmt.Sprint(req.PoolID) + "/external-nodes",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ExternalNode

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNodes: List all the existing nodes for a specific Kubernetes cluster.
func (s *API) ListNodes(req *ListNodesRequest, opts ...scw.RequestOption) (*ListNodesResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}
	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "pool_id", req.PoolID)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "status", req.Status)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/nodes",
		Query:  query,
	}

	var resp ListNodesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNode: Retrieve details about a specific Kubernetes Node.
func (s *API) GetNode(req *GetNodeRequest, opts ...scw.RequestOption) (*Node, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NodeID) == "" {
		return nil, errors.New("field NodeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/nodes/" + fmt.Sprint(req.NodeID) + "",
	}

	var resp Node

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReplaceNode: Replace a specific Node. The node will first be cordoned (scheduling will be disabled on it). The existing pods on the node will then be drained and rescheduled onto another schedulable node. Note that when there is not enough space to reschedule all the pods (such as in a one-node cluster), disruption of your applications can be expected.
func (s *API) ReplaceNode(req *ReplaceNodeRequest, opts ...scw.RequestOption) (*Node, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NodeID) == "" {
		return nil, errors.New("field NodeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/nodes/" + fmt.Sprint(req.NodeID) + "/replace",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Node

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RebootNode: Reboot a specific Node. The node will first be cordoned (scheduling will be disabled on it). The existing pods on the node will then be drained and rescheduled onto another schedulable node. Note that when there is not enough space to reschedule all the pods (such as in a one-node cluster), disruption of your applications can be expected.
func (s *API) RebootNode(req *RebootNodeRequest, opts ...scw.RequestOption) (*Node, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NodeID) == "" {
		return nil, errors.New("field NodeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/nodes/" + fmt.Sprint(req.NodeID) + "/reboot",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Node

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteNode: Delete a specific Node. Note that when there is not enough space to reschedule all the pods (such as in a one-node cluster), disruption of your applications can be expected.
func (s *API) DeleteNode(req *DeleteNodeRequest, opts ...scw.RequestOption) (*Node, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "skip_drain", req.SkipDrain)
	parameter.AddToQuery(query, "replace", req.Replace)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NodeID) == "" {
		return nil, errors.New("field NodeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/nodes/" + fmt.Sprint(req.NodeID) + "",
		Query:  query,
	}

	var resp Node

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVersions: List all available versions for the creation of a new Kubernetes cluster.
func (s *API) ListVersions(req *ListVersionsRequest, opts ...scw.RequestOption) (*ListVersionsResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/versions",
	}

	var resp ListVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVersion: Retrieve a specific Kubernetes version and its details.
func (s *API) GetVersion(req *GetVersionRequest, opts ...scw.RequestOption) (*Version, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.VersionName) == "" {
		return nil, errors.New("field VersionName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/versions/" + fmt.Sprint(req.VersionName) + "",
	}

	var resp Version

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListClusterTypes: List available cluster types and their technical details.
func (s *API) ListClusterTypes(req *ListClusterTypesRequest, opts ...scw.RequestOption) (*ListClusterTypesResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}
	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/cluster-types",
		Query:  query,
	}

	var resp ListClusterTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
