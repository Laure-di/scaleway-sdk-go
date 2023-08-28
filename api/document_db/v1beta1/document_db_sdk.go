// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package document_db provides methods and message types of the document_db v1beta1 API.
package document_db

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

type ACLRuleAction string

const (
	ACLRuleActionAllow = ACLRuleAction("allow")
	ACLRuleActionDeny  = ACLRuleAction("deny")
)

func (enum ACLRuleAction) String() string {
	if enum == "" {
		// return default value if empty
		return "allow"
	}
	return string(enum)
}

func (enum ACLRuleAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ACLRuleAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ACLRuleAction(ACLRuleAction(tmp).String())
	return nil
}

type ACLRuleDirection string

const (
	ACLRuleDirectionInbound  = ACLRuleDirection("inbound")
	ACLRuleDirectionOutbound = ACLRuleDirection("outbound")
)

func (enum ACLRuleDirection) String() string {
	if enum == "" {
		// return default value if empty
		return "inbound"
	}
	return string(enum)
}

func (enum ACLRuleDirection) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ACLRuleDirection) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ACLRuleDirection(ACLRuleDirection(tmp).String())
	return nil
}

type ACLRuleProtocol string

const (
	ACLRuleProtocolTCP  = ACLRuleProtocol("tcp")
	ACLRuleProtocolUDP  = ACLRuleProtocol("udp")
	ACLRuleProtocolIcmp = ACLRuleProtocol("icmp")
)

func (enum ACLRuleProtocol) String() string {
	if enum == "" {
		// return default value if empty
		return "tcp"
	}
	return string(enum)
}

func (enum ACLRuleProtocol) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ACLRuleProtocol) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ACLRuleProtocol(ACLRuleProtocol(tmp).String())
	return nil
}

type EngineSettingPropertyType string

const (
	EngineSettingPropertyTypeBOOLEAN = EngineSettingPropertyType("BOOLEAN")
	EngineSettingPropertyTypeINT     = EngineSettingPropertyType("INT")
	EngineSettingPropertyTypeSTRING  = EngineSettingPropertyType("STRING")
	EngineSettingPropertyTypeFLOAT   = EngineSettingPropertyType("FLOAT")
)

func (enum EngineSettingPropertyType) String() string {
	if enum == "" {
		// return default value if empty
		return "BOOLEAN"
	}
	return string(enum)
}

func (enum EngineSettingPropertyType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *EngineSettingPropertyType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = EngineSettingPropertyType(EngineSettingPropertyType(tmp).String())
	return nil
}

type InstanceLogStatus string

const (
	InstanceLogStatusUnknown  = InstanceLogStatus("unknown")
	InstanceLogStatusReady    = InstanceLogStatus("ready")
	InstanceLogStatusCreating = InstanceLogStatus("creating")
	InstanceLogStatusError    = InstanceLogStatus("error")
)

func (enum InstanceLogStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum InstanceLogStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InstanceLogStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InstanceLogStatus(InstanceLogStatus(tmp).String())
	return nil
}

type InstanceStatus string

const (
	InstanceStatusUnknown      = InstanceStatus("unknown")
	InstanceStatusReady        = InstanceStatus("ready")
	InstanceStatusProvisioning = InstanceStatus("provisioning")
	InstanceStatusConfiguring  = InstanceStatus("configuring")
	InstanceStatusDeleting     = InstanceStatus("deleting")
	InstanceStatusError        = InstanceStatus("error")
	InstanceStatusAutohealing  = InstanceStatus("autohealing")
	InstanceStatusLocked       = InstanceStatus("locked")
	InstanceStatusInitializing = InstanceStatus("initializing")
	InstanceStatusDiskFull     = InstanceStatus("disk_full")
	InstanceStatusBackuping    = InstanceStatus("backuping")
	InstanceStatusSnapshotting = InstanceStatus("snapshotting")
	InstanceStatusRestarting   = InstanceStatus("restarting")
)

func (enum InstanceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum InstanceStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InstanceStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InstanceStatus(InstanceStatus(tmp).String())
	return nil
}

type ListDatabasesRequestOrderBy string

const (
	ListDatabasesRequestOrderByNameAsc  = ListDatabasesRequestOrderBy("name_asc")
	ListDatabasesRequestOrderByNameDesc = ListDatabasesRequestOrderBy("name_desc")
	ListDatabasesRequestOrderBySizeAsc  = ListDatabasesRequestOrderBy("size_asc")
	ListDatabasesRequestOrderBySizeDesc = ListDatabasesRequestOrderBy("size_desc")
)

func (enum ListDatabasesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListDatabasesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDatabasesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDatabasesRequestOrderBy(ListDatabasesRequestOrderBy(tmp).String())
	return nil
}

type ListInstanceLogsRequestOrderBy string

const (
	ListInstanceLogsRequestOrderByCreatedAtAsc  = ListInstanceLogsRequestOrderBy("created_at_asc")
	ListInstanceLogsRequestOrderByCreatedAtDesc = ListInstanceLogsRequestOrderBy("created_at_desc")
)

func (enum ListInstanceLogsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListInstanceLogsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListInstanceLogsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListInstanceLogsRequestOrderBy(ListInstanceLogsRequestOrderBy(tmp).String())
	return nil
}

type ListInstancesRequestOrderBy string

const (
	ListInstancesRequestOrderByCreatedAtAsc  = ListInstancesRequestOrderBy("created_at_asc")
	ListInstancesRequestOrderByCreatedAtDesc = ListInstancesRequestOrderBy("created_at_desc")
	ListInstancesRequestOrderByNameAsc       = ListInstancesRequestOrderBy("name_asc")
	ListInstancesRequestOrderByNameDesc      = ListInstancesRequestOrderBy("name_desc")
	ListInstancesRequestOrderByRegion        = ListInstancesRequestOrderBy("region")
	ListInstancesRequestOrderByStatusAsc     = ListInstancesRequestOrderBy("status_asc")
	ListInstancesRequestOrderByStatusDesc    = ListInstancesRequestOrderBy("status_desc")
)

func (enum ListInstancesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListInstancesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListInstancesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListInstancesRequestOrderBy(ListInstancesRequestOrderBy(tmp).String())
	return nil
}

type ListPrivilegesRequestOrderBy string

const (
	ListPrivilegesRequestOrderByUserNameAsc      = ListPrivilegesRequestOrderBy("user_name_asc")
	ListPrivilegesRequestOrderByUserNameDesc     = ListPrivilegesRequestOrderBy("user_name_desc")
	ListPrivilegesRequestOrderByDatabaseNameAsc  = ListPrivilegesRequestOrderBy("database_name_asc")
	ListPrivilegesRequestOrderByDatabaseNameDesc = ListPrivilegesRequestOrderBy("database_name_desc")
)

func (enum ListPrivilegesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "user_name_asc"
	}
	return string(enum)
}

func (enum ListPrivilegesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPrivilegesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPrivilegesRequestOrderBy(ListPrivilegesRequestOrderBy(tmp).String())
	return nil
}

type ListSnapshotsRequestOrderBy string

const (
	ListSnapshotsRequestOrderByCreatedAtAsc  = ListSnapshotsRequestOrderBy("created_at_asc")
	ListSnapshotsRequestOrderByCreatedAtDesc = ListSnapshotsRequestOrderBy("created_at_desc")
	ListSnapshotsRequestOrderByNameAsc       = ListSnapshotsRequestOrderBy("name_asc")
	ListSnapshotsRequestOrderByNameDesc      = ListSnapshotsRequestOrderBy("name_desc")
	ListSnapshotsRequestOrderByExpiresAtAsc  = ListSnapshotsRequestOrderBy("expires_at_asc")
	ListSnapshotsRequestOrderByExpiresAtDesc = ListSnapshotsRequestOrderBy("expires_at_desc")
)

func (enum ListSnapshotsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListSnapshotsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSnapshotsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSnapshotsRequestOrderBy(ListSnapshotsRequestOrderBy(tmp).String())
	return nil
}

type ListUsersRequestOrderBy string

const (
	ListUsersRequestOrderByNameAsc     = ListUsersRequestOrderBy("name_asc")
	ListUsersRequestOrderByNameDesc    = ListUsersRequestOrderBy("name_desc")
	ListUsersRequestOrderByIsAdminAsc  = ListUsersRequestOrderBy("is_admin_asc")
	ListUsersRequestOrderByIsAdminDesc = ListUsersRequestOrderBy("is_admin_desc")
)

func (enum ListUsersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListUsersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListUsersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListUsersRequestOrderBy(ListUsersRequestOrderBy(tmp).String())
	return nil
}

type MaintenanceStatus string

const (
	MaintenanceStatusUnknown  = MaintenanceStatus("unknown")
	MaintenanceStatusPending  = MaintenanceStatus("pending")
	MaintenanceStatusDone     = MaintenanceStatus("done")
	MaintenanceStatusCanceled = MaintenanceStatus("canceled")
)

func (enum MaintenanceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum MaintenanceStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MaintenanceStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MaintenanceStatus(MaintenanceStatus(tmp).String())
	return nil
}

type NodeTypeGeneration string

const (
	NodeTypeGenerationUnknownGeneration = NodeTypeGeneration("unknown_generation")
	NodeTypeGenerationGenerationV1      = NodeTypeGeneration("generation_v1")
	NodeTypeGenerationGenerationV2      = NodeTypeGeneration("generation_v2")
)

func (enum NodeTypeGeneration) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_generation"
	}
	return string(enum)
}

func (enum NodeTypeGeneration) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NodeTypeGeneration) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NodeTypeGeneration(NodeTypeGeneration(tmp).String())
	return nil
}

type NodeTypeStock string

const (
	NodeTypeStockUnknown    = NodeTypeStock("unknown")
	NodeTypeStockLowStock   = NodeTypeStock("low_stock")
	NodeTypeStockOutOfStock = NodeTypeStock("out_of_stock")
	NodeTypeStockAvailable  = NodeTypeStock("available")
)

func (enum NodeTypeStock) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum NodeTypeStock) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NodeTypeStock) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NodeTypeStock(NodeTypeStock(tmp).String())
	return nil
}

type Permission string

const (
	PermissionReadonly  = Permission("readonly")
	PermissionReadwrite = Permission("readwrite")
	PermissionAll       = Permission("all")
	PermissionCustom    = Permission("custom")
	PermissionNone      = Permission("none")
)

func (enum Permission) String() string {
	if enum == "" {
		// return default value if empty
		return "readonly"
	}
	return string(enum)
}

func (enum Permission) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *Permission) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = Permission(Permission(tmp).String())
	return nil
}

type ReadReplicaStatus string

const (
	ReadReplicaStatusUnknown      = ReadReplicaStatus("unknown")
	ReadReplicaStatusProvisioning = ReadReplicaStatus("provisioning")
	ReadReplicaStatusInitializing = ReadReplicaStatus("initializing")
	ReadReplicaStatusReady        = ReadReplicaStatus("ready")
	ReadReplicaStatusDeleting     = ReadReplicaStatus("deleting")
	ReadReplicaStatusError        = ReadReplicaStatus("error")
	ReadReplicaStatusLocked       = ReadReplicaStatus("locked")
	ReadReplicaStatusConfiguring  = ReadReplicaStatus("configuring")
	ReadReplicaStatusPromoting    = ReadReplicaStatus("promoting")
)

func (enum ReadReplicaStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum ReadReplicaStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ReadReplicaStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ReadReplicaStatus(ReadReplicaStatus(tmp).String())
	return nil
}

type SnapshotStatus string

const (
	SnapshotStatusUnknown   = SnapshotStatus("unknown")
	SnapshotStatusCreating  = SnapshotStatus("creating")
	SnapshotStatusReady     = SnapshotStatus("ready")
	SnapshotStatusRestoring = SnapshotStatus("restoring")
	SnapshotStatusDeleting  = SnapshotStatus("deleting")
	SnapshotStatusError     = SnapshotStatus("error")
	SnapshotStatusLocked    = SnapshotStatus("locked")
)

func (enum SnapshotStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum SnapshotStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SnapshotStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SnapshotStatus(SnapshotStatus(tmp).String())
	return nil
}

type VolumeType string

const (
	VolumeTypeLssd = VolumeType("lssd")
	VolumeTypeBssd = VolumeType("bssd")
)

func (enum VolumeType) String() string {
	if enum == "" {
		// return default value if empty
		return "lssd"
	}
	return string(enum)
}

func (enum VolumeType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *VolumeType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = VolumeType(VolumeType(tmp).String())
	return nil
}

// EndpointDirectAccessDetails:
type EndpointDirectAccessDetails struct {
}

// EndpointLoadBalancerDetails:
type EndpointLoadBalancerDetails struct {
}

// EndpointPrivateNetworkDetails:
type EndpointPrivateNetworkDetails struct {
	// PrivateNetworkID: UUID of the Private Network.
	PrivateNetworkID string `json:"private_network_id"`
	// ServiceIP: CIDR notation of the endpoint IPv4 address.
	ServiceIP scw.IPNet `json:"service_ip"`
	// Zone: Private network zone.
	Zone scw.Zone `json:"zone"`
}

// EndpointSpecPrivateNetworkIpamConfig:
type EndpointSpecPrivateNetworkIpamConfig struct {
}

// ReadReplicaEndpointSpecPrivateNetworkIpamConfig:
type ReadReplicaEndpointSpecPrivateNetworkIpamConfig struct {
}

// EngineSetting:
type EngineSetting struct {
	// Name: Setting name from the database engine.
	Name string `json:"name"`
	// DefaultValue: Value set when not specified.
	DefaultValue string `json:"default_value"`
	// HotConfigurable: Setting can be applied without restarting.
	HotConfigurable bool `json:"hot_configurable"`
	// Description: Setting description.
	Description string `json:"description"`
	// PropertyType: Setting type.
	PropertyType EngineSettingPropertyType `json:"property_type"`
	// Unit: Setting base unit.
	Unit *string `json:"unit"`
	// StringConstraint: Validation regex for string type settings.
	StringConstraint *string `json:"string_constraint"`
	// IntMin: Minimum value for int types.
	IntMin *int32 `json:"int_min"`
	// IntMax: Maximum value for int types.
	IntMax *int32 `json:"int_max"`
	// FloatMin: Minimum value for float types.
	FloatMin *float32 `json:"float_min"`
	// FloatMax: Maximum value for float types.
	FloatMax *float32 `json:"float_max"`
}

// Endpoint:
type Endpoint struct {
	// ID: UUID of the endpoint.
	ID string `json:"id"`
	// IP: IPv4 address of the endpoint.
	IP *net.IP `json:"ip,omitempty"`
	// Port: TCP port of the endpoint.
	Port uint32 `json:"port"`
	// Name: Name of the endpoint.
	Name *string `json:"name"`
	// PrivateNetwork: Private Network details. One maximum per Database Instance or Read Replica (a Database Instance and its Read Replica can have different Private Networks). Cannot be updated (has to be deleted and recreated).
	PrivateNetwork *EndpointPrivateNetworkDetails `json:"private_network,omitempty"`
	// LoadBalancer: Load Balancer details. Public endpoint for Database Instance which is systematically present. One per Database Instance.
	LoadBalancer *EndpointLoadBalancerDetails `json:"load_balancer,omitempty"`
	// DirectAccess: Direct access details. Public endpoint reserved for Read Replicas. One per Read Replica.
	DirectAccess *EndpointDirectAccessDetails `json:"direct_access,omitempty"`
	// Hostname: Hostname of the endpoint.
	Hostname *string `json:"hostname,omitempty"`
}

// EndpointSpecLoadBalancer:
type EndpointSpecLoadBalancer struct {
}

// EndpointSpecPrivateNetwork:
type EndpointSpecPrivateNetwork struct {
	// PrivateNetworkID: UUID of the Private Network to be connected to the Database Instance.
	PrivateNetworkID string `json:"private_network_id"`
	// ServiceIP: Endpoint IPv4 address with a CIDR notation. Refer to the official Scaleway documentation to learn more about IP and subnet limitations.
	ServiceIP *scw.IPNet `json:"service_ip,omitempty"`
	// IpamConfig: Automated configuration of your Private Network endpoint with Scaleway IPAM service. One at the most per Database Instance or Read Replica (a Database Instance and its Read Replica can have different Private Networks). Cannot be updated (has to be deleted and recreated).
	IpamConfig *EndpointSpecPrivateNetworkIpamConfig `json:"ipam_config,omitempty"`
}

// ReadReplicaEndpointSpecDirectAccess:
type ReadReplicaEndpointSpecDirectAccess struct {
}

// ReadReplicaEndpointSpecPrivateNetwork:
type ReadReplicaEndpointSpecPrivateNetwork struct {
	// PrivateNetworkID: UUID of the Private Network to be connected to the Read Replica.
	PrivateNetworkID string `json:"private_network_id"`
	// ServiceIP: Endpoint IPv4 address with a CIDR notation. Refer to the official Scaleway documentation to learn more about IP and subnet limitations.
	ServiceIP *scw.IPNet `json:"service_ip,omitempty"`
	// IpamConfig: Automated configuration of your Private Network endpoint with Scaleway IPAM service. One at the most per Database Instance or Read Replica (a Database Instance and its Read Replica can have different private networks). Cannot be updated (has to be deleted and recreated).
	IpamConfig *ReadReplicaEndpointSpecPrivateNetworkIpamConfig `json:"ipam_config,omitempty"`
}

// EngineVersion:
type EngineVersion struct {
	// Version: Database engine version.
	Version string `json:"version"`
	// Name: Database engine name.
	Name string `json:"name"`
	// EndOfLife: End of life date.
	EndOfLife *time.Time `json:"end_of_life"`
	// AvailableSettings: Engine settings available to be set.
	AvailableSettings []*EngineSetting `json:"available_settings"`
	// Disabled: Disabled versions cannot be created.
	Disabled bool `json:"disabled"`
	// Beta: Beta status of engine version.
	Beta bool `json:"beta"`
	// AvailableInitSettings: Engine settings available to be set at database initialization.
	AvailableInitSettings []*EngineSetting `json:"available_init_settings"`
}

// BackupSchedule:
type BackupSchedule struct {
	// Frequency: Frequency of the backup schedule (in hours).
	Frequency uint32 `json:"frequency"`
	// Retention: Default retention period of backups (in days).
	Retention uint32 `json:"retention"`
	// Disabled: Defines whether the backup schedule feature is disabled.
	Disabled bool `json:"disabled"`
	// NextRunAt: Next run of the backup schedule (accurate to 10 minutes).
	NextRunAt *time.Time `json:"next_run_at"`
}

// InstanceSetting:
type InstanceSetting struct {
	// Name:
	Name string `json:"name"`
	// Value:
	Value string `json:"value"`
}

// LogsPolicy:
type LogsPolicy struct {
	// MaxAgeRetention: Max age (in days) of remote logs to keep on the Database Instance.
	MaxAgeRetention *uint32 `json:"max_age_retention"`
	// TotalDiskRetention: Max disk size of remote logs to keep on the Database Instance.
	TotalDiskRetention *scw.Size `json:"total_disk_retention"`
}

// Maintenance:
type Maintenance struct {
	// StartsAt: Start date of the maintenance window.
	StartsAt *time.Time `json:"starts_at"`
	// StopsAt: End date of the maintenance window.
	StopsAt *time.Time `json:"stops_at"`
	// ClosedAt: Closed maintenance date.
	ClosedAt *time.Time `json:"closed_at"`
	// Reason: Maintenance information message.
	Reason string `json:"reason"`
	// Status: Status of the maintenance.
	Status MaintenanceStatus `json:"status"`
}

// ReadReplica:
type ReadReplica struct {
	// ID: UUID of the Read Replica.
	ID string `json:"id"`
	// Endpoints: Display Read Replica connection information.
	Endpoints []*Endpoint `json:"endpoints"`
	// Status: Read replica status.
	Status ReadReplicaStatus `json:"status"`
	// Region: Region the Read Replica is in.
	Region scw.Region `json:"region"`
	// SameZone: Whether the replica is in the same Availability Zone as the main Database Instance nodes or not.
	SameZone bool `json:"same_zone"`
}

// UpgradableVersion:
type UpgradableVersion struct {
	// ID:
	ID string `json:"id"`
	// Name:
	Name string `json:"name"`
	// Version:
	Version string `json:"version"`
	// MinorVersion:
	MinorVersion string `json:"minor_version"`
}

// Volume:
type Volume struct {
	// Type:
	Type VolumeType `json:"type"`
	// Size:
	Size scw.Size `json:"size"`
}

// NodeTypeVolumeConstraintSizes:
type NodeTypeVolumeConstraintSizes struct {
	// MinSize: [deprecated] Mimimum size required for the Volume.
	MinSize scw.Size `json:"min_size"`
	// MaxSize: [deprecated] Maximum size required for the Volume.
	MaxSize scw.Size `json:"max_size"`
}

// NodeTypeVolumeType:
type NodeTypeVolumeType struct {
	// Type: Volume Type.
	Type VolumeType `json:"type"`
	// Description: The description of the Volume.
	Description string `json:"description"`
	// MinSize: Mimimum size required for the Volume.
	MinSize scw.Size `json:"min_size"`
	// MaxSize: Maximum size required for the Volume.
	MaxSize scw.Size `json:"max_size"`
	// ChunkSize: Minimum increment level for a Block Storage volume size.
	ChunkSize scw.Size `json:"chunk_size"`
}

// ACLRuleRequest:
type ACLRuleRequest struct {
	// IP:
	IP scw.IPNet `json:"ip"`
	// Description:
	Description string `json:"description"`
}

// ACLRule:
type ACLRule struct {
	// IP:
	IP scw.IPNet `json:"ip"`
	// Deprecated: Port:
	Port *uint32 `json:"port,omitempty"`
	// Protocol:
	Protocol ACLRuleProtocol `json:"protocol"`
	// Direction:
	Direction ACLRuleDirection `json:"direction"`
	// Action:
	Action ACLRuleAction `json:"action"`
	// Description:
	Description string `json:"description"`
}

// EndpointSpec:
type EndpointSpec struct {
	// LoadBalancer: Load Balancer endpoint specifications. Public endpoint for Database Instance which is systematically present. One per Document Database Instance.
	LoadBalancer *EndpointSpecLoadBalancer `json:"load_balancer,omitempty"`
	// PrivateNetwork: Private Network endpoint specifications. One maximum per Database Instance or Read Replica (a Database Instance and its Read Replica can have different Private Networks). Cannot be updated (has to be deleted and recreated).
	PrivateNetwork *EndpointSpecPrivateNetwork `json:"private_network,omitempty"`
}

// ReadReplicaEndpointSpec:
type ReadReplicaEndpointSpec struct {
	// DirectAccess: Direct access endpoint specifications. Public endpoint reserved for Read Replicas. One per Read Replica.
	DirectAccess *ReadReplicaEndpointSpecDirectAccess `json:"direct_access,omitempty"`
	// PrivateNetwork: Private Network endpoint specifications. One at most, per Read Replica. Cannot be updated (has to be deleted and recreated).
	PrivateNetwork *ReadReplicaEndpointSpecPrivateNetwork `json:"private_network,omitempty"`
}

// DatabaseEngine:
type DatabaseEngine struct {
	// Name: Engine name.
	Name string `json:"name"`
	// LogoURL: Engine logo URL.
	LogoURL string `json:"logo_url"`
	// Versions: Available versions.
	Versions []*EngineVersion `json:"versions"`
	// Region: Region of this Database Instance.
	Region scw.Region `json:"region"`
}

// Database:
type Database struct {
	// Name: Name of the database.
	Name string `json:"name"`
	// Owner: Name of the database owner.
	Owner string `json:"owner"`
	// Managed: Defines whether the database is managed or not.
	Managed bool `json:"managed"`
	// Size: Size of the database.
	Size scw.Size `json:"size"`
}

// ListInstanceLogsDetailsResponseInstanceLogDetail:
type ListInstanceLogsDetailsResponseInstanceLogDetail struct {
	// LogName:
	LogName string `json:"log_name"`
	// Size:
	Size uint64 `json:"size"`
}

// InstanceLog:
type InstanceLog struct {
	// DownloadURL: Presigned S3 URL to download your log file.
	DownloadURL *string `json:"download_url"`
	// ID: UUID of the Database Instance log.
	ID string `json:"id"`
	// Status: Status of the logs in a Database Instance.
	Status InstanceLogStatus `json:"status"`
	// NodeName: Name of the underlying node.
	NodeName string `json:"node_name"`
	// ExpiresAt: Expiration date (must follow the ISO 8601 format).
	ExpiresAt *time.Time `json:"expires_at"`
	// CreatedAt: Creation date (must follow the ISO 8601 format).
	CreatedAt *time.Time `json:"created_at"`
	// Region: Region the Database Instance is in.
	Region scw.Region `json:"region"`
}

// Instance:
type Instance struct {
	// CreatedAt: Creation date (must follow the ISO 8601 format).
	CreatedAt *time.Time `json:"created_at"`
	// Volume: Volumes of the Database Instance.
	Volume *Volume `json:"volume"`
	// Region: Region the Database Instance is in.
	Region scw.Region `json:"region"`
	// ID: UUID of the Database Instance.
	ID string `json:"id"`
	// Name: Name of the Database Instance.
	Name string `json:"name"`
	// OrganizationID: Organization ID the Database Instance belongs to.
	OrganizationID string `json:"organization_id"`
	// ProjectID: Project ID the Database Instance belongs to.
	ProjectID string `json:"project_id"`
	// Status: Status of the Database Instance.
	Status InstanceStatus `json:"status"`
	// Engine: Database engine of the database.
	Engine string `json:"engine"`
	// UpgradableVersion: Available database engine versions for upgrade.
	UpgradableVersion []*UpgradableVersion `json:"upgradable_version"`
	// Deprecated: Endpoint: Endpoint of the Database Instance.
	Endpoint *Endpoint `json:"endpoint,omitempty"`
	// Tags: List of tags applied to the Database Instance.
	Tags []string `json:"tags"`
	// Settings: Advanced settings of the Database Instance.
	Settings []*InstanceSetting `json:"settings"`
	// BackupSchedule: Backup schedule of the Database Instance.
	BackupSchedule *BackupSchedule `json:"backup_schedule"`
	// IsHaCluster: Defines whether or not High-Availability is enabled.
	IsHaCluster bool `json:"is_ha_cluster"`
	// ReadReplicas: Read Replicas of the Database Instance.
	ReadReplicas []*ReadReplica `json:"read_replicas"`
	// NodeType: Node type of the Database Instance.
	NodeType string `json:"node_type"`
	// InitSettings: List of engine settings to be set at Database Instance initialization.
	InitSettings []*InstanceSetting `json:"init_settings"`
	// Endpoints: List of Database Instance endpoints.
	Endpoints []*Endpoint `json:"endpoints"`
	// LogsPolicy: Logs policy of the Database Instance.
	LogsPolicy *LogsPolicy `json:"logs_policy"`
	// BackupSameRegion: Store logical backups in the same region as the Database Instance.
	BackupSameRegion bool `json:"backup_same_region"`
	// Maintenances: List of Database Instance maintenance events.
	Maintenances []*Maintenance `json:"maintenances"`
}

// NodeType:
type NodeType struct {
	// Name: Node Type name identifier.
	Name string `json:"name"`
	// StockStatus: Current stock status for the Node Type.
	StockStatus NodeTypeStock `json:"stock_status"`
	// Description: Current specs of the offer.
	Description string `json:"description"`
	// Vcpus: Number of virtual CPUs.
	Vcpus uint32 `json:"vcpus"`
	// Memory: Quantity of RAM.
	Memory scw.Size `json:"memory"`
	// Deprecated: VolumeConstraint: [deprecated] Node Type volume constraints.
	VolumeConstraint *NodeTypeVolumeConstraintSizes `json:"volume_constraint,omitempty"`
	// Deprecated: IsBssdCompatible: The Node Type is compliant with Block Storage.
	IsBssdCompatible *bool `json:"is_bssd_compatible,omitempty"`
	// Disabled: The Node Type is currently disabled.
	Disabled bool `json:"disabled"`
	// Beta: The Node Type is currently in beta.
	Beta bool `json:"beta"`
	// AvailableVolumeTypes: Available storage options for the Node Type.
	AvailableVolumeTypes []*NodeTypeVolumeType `json:"available_volume_types"`
	// IsHaRequired: The Node Type can be used only with the High Availability option.
	IsHaRequired bool `json:"is_ha_required"`
	// Generation: Generation associated the NodeType offer.
	Generation NodeTypeGeneration `json:"generation"`
	// InstanceRange: Instance range associated with the NodeType offer.
	InstanceRange string `json:"instance_range"`
	// Region: Region the Node Type is in.
	Region scw.Region `json:"region"`
}

// Privilege:
type Privilege struct {
	// Permission: Permission (Read, Read/Write, All, Custom).
	Permission Permission `json:"permission"`
	// DatabaseName: Name of the database.
	DatabaseName string `json:"database_name"`
	// UserName: Name of the user.
	UserName string `json:"user_name"`
}

// Snapshot:
type Snapshot struct {
	// ID: UUID of the snapshot.
	ID string `json:"id"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"instance_id"`
	// Name: Name of the snapshot.
	Name string `json:"name"`
	// Status: Status of the snapshot.
	Status SnapshotStatus `json:"status"`
	// Size: Size of the snapshot.
	Size *scw.Size `json:"size"`
	// ExpiresAt: Expiration date (must follow the ISO 8601 format).
	ExpiresAt *time.Time `json:"expires_at"`
	// CreatedAt: Creation date (must follow the ISO 8601 format).
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Updated date (must follow the ISO 8601 format).
	UpdatedAt *time.Time `json:"updated_at"`
	// InstanceName: Name of the Database Instance of the snapshot.
	InstanceName string `json:"instance_name"`
	// NodeType: Source node type.
	NodeType string `json:"node_type"`
	// Region: Region of this snapshot.
	Region scw.Region `json:"region"`
}

// User:
type User struct {
	// Name: Name of the user (Length must be between 1 and 63 characters. First character must be an alphabet character (a-zA-Z). Your username cannot start with '_rdb' or 'pg_'. Only a-zA-Z0-9_$- characters are accepted).
	Name string `json:"name"`
	// IsAdmin: Defines whether or not a user got administrative privileges on the Database Instance.
	IsAdmin bool `json:"is_admin"`
}

// AddInstanceACLRulesRequest:
type AddInstanceACLRulesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want to add ACL rules to.
	InstanceID string `json:"-"`
	// Rules: ACL rules to add to the Database Instance.
	Rules []*ACLRuleRequest `json:"rules"`
}

// AddInstanceACLRulesResponse:
type AddInstanceACLRulesResponse struct {
	// Rules: ACL Rules enabled for the Database Instance.
	Rules []*ACLRule `json:"rules"`
}

// AddInstanceSettingsRequest:
type AddInstanceSettingsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want to add settings to.
	InstanceID string `json:"-"`
	// Settings: Settings to add to the Database Instance.
	Settings []*InstanceSetting `json:"settings"`
}

// AddInstanceSettingsResponse:
type AddInstanceSettingsResponse struct {
	// Settings: Settings available on the Database Instance.
	Settings []*InstanceSetting `json:"settings"`
}

// CloneInstanceRequest:
type CloneInstanceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want to clone.
	InstanceID string `json:"-"`
	// Name: Name of the Database Instance clone.
	Name string `json:"name"`
	// NodeType: Node type of the clone.
	NodeType *string `json:"node_type,omitempty"`
}

// CreateDatabaseRequest:
type CreateDatabaseRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance where to create the database.
	InstanceID string `json:"-"`
	// Name: Name of the database.
	Name string `json:"name"`
}

// CreateEndpointRequest:
type CreateEndpointRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you to which you want to add an endpoint.
	InstanceID string `json:"-"`
	// EndpointSpec: Specification of the endpoint you want to create.
	EndpointSpec *EndpointSpec `json:"endpoint_spec"`
}

// CreateInstanceFromSnapshotRequest:
type CreateInstanceFromSnapshotRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SnapshotID: Block snapshot of the Database Instance.
	SnapshotID string `json:"-"`
	// InstanceName: Name of the Database Instance created with the snapshot.
	InstanceName string `json:"instance_name"`
	// IsHaCluster: Defines whether or not High Availability is enabled on the new Database Instance.
	IsHaCluster *bool `json:"is_ha_cluster,omitempty"`
	// NodeType: The node type used to restore the snapshot.
	NodeType *string `json:"node_type,omitempty"`
}

// CreateInstanceRequest:
type CreateInstanceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Deprecated: OrganizationID: Please use project_id instead.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: The Project ID on which the Database Instance will be created.
	ProjectID *string `json:"project_id,omitempty"`
	// Name: Name of the Database Instance.
	Name string `json:"name"`
	// Engine: Database engine of the Database Instance.
	Engine string `json:"engine"`
	// UserName: Username created when the Database Instance is created.
	UserName string `json:"user_name"`
	// Password: Password of the user.
	Password string `json:"password"`
	// NodeType: Type of node to use for the Database Instance.
	NodeType string `json:"node_type"`
	// IsHaCluster: Defines whether or not High-Availability is enabled.
	IsHaCluster bool `json:"is_ha_cluster"`
	// DisableBackup: Defines whether or not backups are disabled.
	DisableBackup bool `json:"disable_backup"`
	// Tags: Tags to apply to the Database Instance.
	Tags []string `json:"tags"`
	// InitSettings: List of engine settings to be set upon Database Instance initialization.
	InitSettings []*InstanceSetting `json:"init_settings"`
	// VolumeType: Type of volume where data is stored (lssd, bssd, ...).
	VolumeType VolumeType `json:"volume_type"`
	// VolumeSize: Volume size when volume_type is not lssd.
	VolumeSize scw.Size `json:"volume_size"`
	// InitEndpoints: One or multiple EndpointSpec used to expose your Database Instance. A load_balancer public endpoint is systematically created.
	InitEndpoints []*EndpointSpec `json:"init_endpoints"`
	// BackupSameRegion: Defines whether to or not to store logical backups in the same region as the Database Instance.
	BackupSameRegion bool `json:"backup_same_region"`
}

// CreateReadReplicaEndpointRequest:
type CreateReadReplicaEndpointRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ReadReplicaID: UUID of the Read Replica.
	ReadReplicaID string `json:"-"`
	// EndpointSpec: Specification of the endpoint you want to create.
	EndpointSpec []*ReadReplicaEndpointSpec `json:"endpoint_spec"`
}

// CreateReadReplicaRequest:
type CreateReadReplicaRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want to create a Read Replica from.
	InstanceID string `json:"instance_id"`
	// EndpointSpec: Specification of the endpoint you want to create.
	EndpointSpec []*ReadReplicaEndpointSpec `json:"endpoint_spec"`
	// SameZone: Defines whether or not to create the replica in the same Availability Zone as the main Database Instance nodes.
	SameZone *bool `json:"same_zone,omitempty"`
}

// CreateSnapshotRequest:
type CreateSnapshotRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
	// Name: Name of the snapshot.
	Name string `json:"name"`
	// ExpiresAt: Expiration date (must follow the ISO 8601 format).
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// CreateUserRequest:
type CreateUserRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance in which you want to create a user.
	InstanceID string `json:"-"`
	// Name: Name of the user you want to create.
	Name string `json:"name"`
	// Password: Password of the user you want to create.
	Password string `json:"password"`
	// IsAdmin: Defines whether the user will have administrative privileges.
	IsAdmin bool `json:"is_admin"`
}

// DeleteDatabaseRequest:
type DeleteDatabaseRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance where to delete the database.
	InstanceID string `json:"-"`
	// Name: Name of the database to delete.
	Name string `json:"-"`
}

// DeleteEndpointRequest:
type DeleteEndpointRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// EndpointID: This endpoint can also be used to delete a Read Replica endpoint.
	EndpointID string `json:"-"`
}

// DeleteInstanceACLRulesRequest:
type DeleteInstanceACLRulesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want to delete an ACL rule from.
	InstanceID string `json:"-"`
	// ACLRuleIPs: IP addresses defined in the ACL rules of the Database Instance.
	ACLRuleIPs []string `json:"acl_rule_ips"`
}

// DeleteInstanceACLRulesResponse:
type DeleteInstanceACLRulesResponse struct {
	// Rules: IP addresses defined in the ACL rules of the Database Instance.
	Rules []*ACLRule `json:"rules"`
}

// DeleteInstanceRequest:
type DeleteInstanceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance to delete.
	InstanceID string `json:"-"`
}

// DeleteInstanceSettingsRequest:
type DeleteInstanceSettingsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance to delete settings from.
	InstanceID string `json:"-"`
	// SettingNames: Settings names to delete.
	SettingNames []string `json:"setting_names"`
}

// DeleteInstanceSettingsResponse:
type DeleteInstanceSettingsResponse struct {
	// Settings: Settings names to delete from the Database Instance.
	Settings []*InstanceSetting `json:"settings"`
}

// DeleteReadReplicaRequest:
type DeleteReadReplicaRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ReadReplicaID: UUID of the Read Replica.
	ReadReplicaID string `json:"-"`
}

// DeleteSnapshotRequest:
type DeleteSnapshotRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SnapshotID: UUID of the snapshot to delete.
	SnapshotID string `json:"-"`
}

// DeleteUserRequest:
type DeleteUserRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance to delete the user from.
	InstanceID string `json:"-"`
	// Name: Name of the user.
	Name string `json:"-"`
}

// GetEndpointRequest:
type GetEndpointRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// EndpointID: UUID of the endpoint you want to get.
	EndpointID string `json:"-"`
}

// GetInstanceCertificateRequest:
type GetInstanceCertificateRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
}

// GetInstanceLogRequest:
type GetInstanceLogRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceLogID: UUID of the instance_log you want.
	InstanceLogID string `json:"-"`
}

// GetInstanceMetricsRequest:
type GetInstanceMetricsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
	// StartDate: Start date to gather metrics from.
	StartDate *time.Time `json:"-"`
	// EndDate: End date to gather metrics from.
	EndDate *time.Time `json:"-"`
	// MetricName: Name of the metric to gather.
	MetricName *string `json:"-"`
}

// GetInstanceRequest:
type GetInstanceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
}

// GetReadReplicaRequest:
type GetReadReplicaRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ReadReplicaID: UUID of the Read Replica.
	ReadReplicaID string `json:"-"`
}

// GetSnapshotRequest:
type GetSnapshotRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SnapshotID: UUID of the snapshot.
	SnapshotID string `json:"-"`
}

// InstanceMetrics:
type InstanceMetrics struct {
	// Timeseries: Time series of metrics of a Database Instance.
	Timeseries []*scw.TimeSeries `json:"timeseries"`
}

// ListDatabaseEnginesRequest:
type ListDatabaseEnginesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Name: Name of the database engine.
	Name *string `json:"-"`
	// Version: Version of the database engine.
	Version *string `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
}

// ListDatabaseEnginesResponse:
type ListDatabaseEnginesResponse struct {
	// Engines: List of the available database engines.
	Engines []*DatabaseEngine `json:"engines"`
	// TotalCount: Total count of database engines available.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDatabaseEnginesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDatabaseEnginesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDatabaseEnginesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Engines = append(r.Engines, results.Engines...)
	r.TotalCount += uint32(len(results.Engines))
	return uint32(len(results.Engines)), nil
}

// ListDatabasesRequest:
type ListDatabasesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance to list the databases of.
	InstanceID string `json:"-"`
	// Name: Name of the database.
	Name *string `json:"-"`
	// Managed: Defines whether or not the database is managed.
	Managed *bool `json:"-"`
	// Owner: User that owns this database.
	Owner *string `json:"-"`
	// OrderBy: Criteria to use when ordering database listing.
	OrderBy ListDatabasesRequestOrderBy `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
}

// ListDatabasesResponse:
type ListDatabasesResponse struct {
	// Databases: List of the databases.
	Databases []*Database `json:"databases"`
	// TotalCount: Total count of databases present on a Database Instance.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDatabasesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDatabasesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDatabasesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Databases = append(r.Databases, results.Databases...)
	r.TotalCount += uint32(len(results.Databases))
	return uint32(len(results.Databases)), nil
}

// ListInstanceACLRulesRequest:
type ListInstanceACLRulesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
}

// ListInstanceACLRulesResponse:
type ListInstanceACLRulesResponse struct {
	// Rules: List of ACL rules present on a Database Instance.
	Rules []*ACLRule `json:"rules"`
	// TotalCount: Total count of ACL rules present on a Database Instance.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListInstanceACLRulesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListInstanceACLRulesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListInstanceACLRulesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Rules = append(r.Rules, results.Rules...)
	r.TotalCount += uint32(len(results.Rules))
	return uint32(len(results.Rules)), nil
}

// ListInstanceLogsDetailsRequest:
type ListInstanceLogsDetailsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want logs of.
	InstanceID string `json:"-"`
}

// ListInstanceLogsDetailsResponse:
type ListInstanceLogsDetailsResponse struct {
	// Details: Remote Database Instance logs details.
	Details []*ListInstanceLogsDetailsResponseInstanceLogDetail `json:"details"`
}

// ListInstanceLogsRequest:
type ListInstanceLogsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want logs of.
	InstanceID string `json:"-"`
	// OrderBy: Criteria to use when ordering Database Instance logs listing.
	OrderBy ListInstanceLogsRequestOrderBy `json:"-"`
}

// ListInstanceLogsResponse:
type ListInstanceLogsResponse struct {
	// InstanceLogs: Available logs in a Database Instance.
	InstanceLogs []*InstanceLog `json:"instance_logs"`
}

// ListInstancesRequest:
type ListInstancesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Tags: List Database Instances that have a given tag.
	Tags []string `json:"-"`
	// Name: Lists Database Instances that match a name pattern.
	Name *string `json:"-"`
	// OrderBy: Criteria to use when ordering Database Instance listings.
	OrderBy ListInstancesRequestOrderBy `json:"-"`
	// OrganizationID: Please use project_id instead.
	OrganizationID *string `json:"-"`
	// ProjectID: Project ID to list the Database Instance of.
	ProjectID *string `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
}

// ListInstancesResponse:
type ListInstancesResponse struct {
	// Instances: List of all Database Instances available in an Organization or Project.
	Instances []*Instance `json:"instances"`
	// TotalCount: Total count of Database Instances available in a Organization or Project.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListInstancesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListInstancesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListInstancesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Instances = append(r.Instances, results.Instances...)
	r.TotalCount += uint32(len(results.Instances))
	return uint32(len(results.Instances)), nil
}

// ListNodeTypesRequest:
type ListNodeTypesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// IncludeDisabledTypes: Defines whether or not to include disabled types.
	IncludeDisabledTypes bool `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
}

// ListNodeTypesResponse:
type ListNodeTypesResponse struct {
	// NodeTypes: Types of the node.
	NodeTypes []*NodeType `json:"node_types"`
	// TotalCount: Total count of node-types available.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNodeTypesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNodeTypesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListNodeTypesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.NodeTypes = append(r.NodeTypes, results.NodeTypes...)
	r.TotalCount += uint32(len(results.NodeTypes))
	return uint32(len(results.NodeTypes)), nil
}

// ListPrivilegesRequest:
type ListPrivilegesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
	// OrderBy: Criteria to use when ordering privileges listing.
	OrderBy ListPrivilegesRequestOrderBy `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
	// DatabaseName: Name of the database.
	DatabaseName *string `json:"-"`
	// UserName: Name of the user.
	UserName *string `json:"-"`
}

// ListPrivilegesResponse:
type ListPrivilegesResponse struct {
	// Privileges: Privileges of a user in a database in a Database Instance.
	Privileges []*Privilege `json:"privileges"`
	// TotalCount: Total count of privileges present on a database.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPrivilegesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPrivilegesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListPrivilegesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Privileges = append(r.Privileges, results.Privileges...)
	r.TotalCount += uint32(len(results.Privileges))
	return uint32(len(results.Privileges)), nil
}

// ListSnapshotsRequest:
type ListSnapshotsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Name: Name of the snapshot.
	Name *string `json:"-"`
	// OrderBy: Criteria to use when ordering snapshot listing.
	OrderBy ListSnapshotsRequestOrderBy `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID *string `json:"-"`
	// OrganizationID: Organization ID the snapshots belongs to.
	OrganizationID *string `json:"-"`
	// ProjectID: Project ID the snapshots belongs to.
	ProjectID *string `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
}

// ListSnapshotsResponse:
type ListSnapshotsResponse struct {
	// Snapshots: List of snapshots.
	Snapshots []*Snapshot `json:"snapshots"`
	// TotalCount: Total count of snapshots available.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSnapshotsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSnapshotsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSnapshotsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Snapshots = append(r.Snapshots, results.Snapshots...)
	r.TotalCount += uint32(len(results.Snapshots))
	return uint32(len(results.Snapshots)), nil
}

// ListUsersRequest:
type ListUsersRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
	// Name: Name of the user.
	Name *string `json:"-"`
	// OrderBy: Criteria to use when requesting user listing.
	OrderBy ListUsersRequestOrderBy `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
}

// ListUsersResponse:
type ListUsersResponse struct {
	// Users: List of users in a Database Instance.
	Users []*User `json:"users"`
	// TotalCount: Total count of users present on a Database Instance.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListUsersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListUsersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListUsersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Users = append(r.Users, results.Users...)
	r.TotalCount += uint32(len(results.Users))
	return uint32(len(results.Users)), nil
}

// MigrateEndpointRequest:
type MigrateEndpointRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// EndpointID: UUID of the endpoint you want to migrate.
	EndpointID string `json:"-"`
	// InstanceID: UUID of the instance you want to attach the endpoint to.
	InstanceID string `json:"instance_id"`
}

// PromoteReadReplicaRequest:
type PromoteReadReplicaRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ReadReplicaID: UUID of the Read Replica.
	ReadReplicaID string `json:"-"`
}

// PurgeInstanceLogsRequest:
type PurgeInstanceLogsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want logs of.
	InstanceID string `json:"-"`
	// LogName: Given log name to purge.
	LogName *string `json:"log_name,omitempty"`
}

// RenewInstanceCertificateRequest:
type RenewInstanceCertificateRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want logs of.
	InstanceID string `json:"-"`
}

// ResetReadReplicaRequest:
type ResetReadReplicaRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ReadReplicaID: UUID of the Read Replica.
	ReadReplicaID string `json:"-"`
}

// RestartInstanceRequest:
type RestartInstanceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want to restart.
	InstanceID string `json:"-"`
}

// SetInstanceACLRulesRequest:
type SetInstanceACLRulesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance where the ACL rules must be set.
	InstanceID string `json:"-"`
	// Rules: ACL rules to define for the Database Instance.
	Rules []*ACLRuleRequest `json:"rules"`
}

// SetInstanceACLRulesResponse:
type SetInstanceACLRulesResponse struct {
	// Rules: ACLs rules configured for a Database Instance.
	Rules []*ACLRule `json:"rules"`
}

// SetInstanceSettingsRequest:
type SetInstanceSettingsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance where the settings must be set.
	InstanceID string `json:"-"`
	// Settings: Settings to define for the Database Instance.
	Settings []*InstanceSetting `json:"settings"`
}

// SetInstanceSettingsResponse:
type SetInstanceSettingsResponse struct {
	// Settings: Settings configured for a Database Instance.
	Settings []*InstanceSetting `json:"settings"`
}

// SetPrivilegeRequest:
type SetPrivilegeRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
	// DatabaseName: Name of the database.
	DatabaseName string `json:"database_name"`
	// UserName: Name of the user.
	UserName string `json:"user_name"`
	// Permission: Permission to set (Read, Read/Write, All, Custom).
	Permission Permission `json:"permission"`
}

// UpdateInstanceRequest:
type UpdateInstanceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance to update.
	InstanceID string `json:"-"`
	// BackupScheduleFrequency: In hours.
	BackupScheduleFrequency *uint32 `json:"backup_schedule_frequency,omitempty"`
	// BackupScheduleRetention: In days.
	BackupScheduleRetention *uint32 `json:"backup_schedule_retention,omitempty"`
	// IsBackupScheduleDisabled: Defines whether or not the backup schedule is disabled.
	IsBackupScheduleDisabled *bool `json:"is_backup_schedule_disabled,omitempty"`
	// Name: Name of the Database Instance.
	Name *string `json:"name,omitempty"`
	// Tags: Tags of a Database Instance.
	Tags *[]string `json:"tags,omitempty"`
	// LogsPolicy: Logs policy of the Database Instance.
	LogsPolicy *LogsPolicy `json:"logs_policy"`
	// BackupSameRegion: Store logical backups in the same region as the Database Instance.
	BackupSameRegion *bool `json:"backup_same_region,omitempty"`
	// BackupScheduleStartHour: Defines the start time of the autobackup.
	BackupScheduleStartHour *uint32 `json:"backup_schedule_start_hour,omitempty"`
}

// UpdateSnapshotRequest:
type UpdateSnapshotRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SnapshotID: UUID of the snapshot to update.
	SnapshotID string `json:"-"`
	// Name: Name of the snapshot.
	Name *string `json:"name,omitempty"`
	// ExpiresAt: Expiration date (must follow the ISO 8601 format).
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// UpdateUserRequest:
type UpdateUserRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance the user belongs to.
	InstanceID string `json:"-"`
	// Name: Name of the database user.
	Name string `json:"-"`
	// Password: Password of the database user.
	Password *string `json:"password,omitempty"`
	// IsAdmin: Defines whether or not this user got administrative privileges.
	IsAdmin *bool `json:"is_admin,omitempty"`
}

// UpgradeInstanceRequest:
type UpgradeInstanceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want to upgrade.
	InstanceID string `json:"-"`
	// NodeType: Node type of the Database Instance you want to upgrade to.
	NodeType *string `json:"node_type,omitempty"`
	// EnableHa: Defines whether or not High Availability should be enabled on the Database Instance.
	EnableHa *bool `json:"enable_ha,omitempty"`
	// VolumeSize: Increase your Block volume size.
	VolumeSize *uint64 `json:"volume_size,omitempty"`
	// VolumeType: Change your Database Instance storage type.
	VolumeType *VolumeType `json:"volume_type,omitempty"`
	// UpgradableVersionID: This will create a new Database Instance with same specifications as the current one and perform a Database Engine upgrade.
	UpgradableVersionID *string `json:"upgradable_version_id,omitempty"`
}

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

// ListDatabaseEngines: List the FerretDB database engines available at Scaleway.
func (s *API) ListDatabaseEngines(req *ListDatabaseEnginesRequest, opts ...scw.RequestOption) (*ListDatabaseEnginesResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "version", req.Version)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/database-engines",
		Query:  query,
	}

	var resp ListDatabaseEnginesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNodeTypes: List all available node types. By default, the node types returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListNodeTypes(req *ListNodeTypesRequest, opts ...scw.RequestOption) (*ListNodeTypesResponse, error) {
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
	parameter.AddToQuery(query, "include_disabled_types", req.IncludeDisabledTypes)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/node-types",
		Query:  query,
	}

	var resp ListNodeTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpgradeInstance: Upgrade your current Database Instance specifications like node type, high availability, volume, or the database engine version. Note that upon upgrade the `enable_ha` parameter can only be set to `true`.
func (s *API) UpgradeInstance(req *UpgradeInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/upgrade",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListInstances: List all Database Instances in the specified region, for a given Scaleway Organization or Scaleway Project. By default, the Database Instances returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field. You can define additional parameters for your query, such as `tags` and `name`. For the `name` parameter, the value you include will be checked against the whole name string to see if it includes the string you put in the parameter.
func (s *API) ListInstances(req *ListInstancesRequest, opts ...scw.RequestOption) (*ListInstancesResponse, error) {
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
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances",
		Query:  query,
	}

	var resp ListInstancesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInstance: Retrieve information about a given Database Instance, specified by the `region` and `instance_id` parameters. Its full details, including name, status, IP address and port, are returned in the response object.
func (s *API) GetInstance(req *GetInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "",
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateInstance: Create a new Database Instance. You must set the `engine`, `user_name`, `password` and `node_type` parameters. Optionally, you can specify the volume type and size.
func (s *API) CreateInstance(req *CreateInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
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
		req.Name = namegenerator.GetRandomName("ins")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateInstance: Update the parameters of a Database Instance, including name, tags and backup schedule details.
func (s *API) UpdateInstance(req *UpdateInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteInstance: Delete a given Database Instance, specified by the `region` and `instance_id` parameters. Deleting a Database Instance is permanent, and cannot be undone. Note that upon deletion all your data will be lost.
func (s *API) DeleteInstance(req *DeleteInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "",
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CloneInstance: Clone a given Database Instance, specified by the `region` and `instance_id` parameters. The clone feature allows you to create a new Database Instance from an existing one. The clone includes all existing databases, users and permissions. You can create a clone on a Database Instance bigger than your current one.
func (s *API) CloneInstance(req *CloneInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/clone",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RestartInstance: Restart a given Database Instance, specified by the `region` and `instance_id` parameters. The status of the Database Instance returned in the response.
func (s *API) RestartInstance(req *RestartInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/restart",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInstanceCertificate: Retrieve information about the TLS certificate of a given Database Instance. Details like name and content are returned in the response.
func (s *API) GetInstanceCertificate(req *GetInstanceCertificateRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/certificate",
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RenewInstanceCertificate: Renew a TLS for a Database Instance. Renewing a certificate means that you will not be able to connect to your Database Instance using the previous certificate. You will also need to download and update the new certificate for all database clients.
func (s *API) RenewInstanceCertificate(req *RenewInstanceCertificateRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/renew-certificate",
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

// GetInstanceMetrics: Retrieve the time series metrics of a given Database Instance. You can define the period from which to retrieve metrics by specifying the `start_date` and `end_date`.
func (s *API) GetInstanceMetrics(req *GetInstanceMetricsRequest, opts ...scw.RequestOption) (*InstanceMetrics, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "start_date", req.StartDate)
	parameter.AddToQuery(query, "end_date", req.EndDate)
	parameter.AddToQuery(query, "metric_name", req.MetricName)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/metrics",
		Query:  query,
	}

	var resp InstanceMetrics

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateReadReplica: Create a new Read Replica of a Database Instance. You must specify the `region` and the `instance_id`. You can only create a maximum of 3 Read Replicas per Database Instance.
func (s *API) CreateReadReplica(req *CreateReadReplicaRequest, opts ...scw.RequestOption) (*ReadReplica, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/read-replicas",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ReadReplica

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetReadReplica: Retrieve information about a Database Instance Read Replica. Full details about the Read Replica, like `endpoints`, `status`  and `region` are returned in the response.
func (s *API) GetReadReplica(req *GetReadReplicaRequest, opts ...scw.RequestOption) (*ReadReplica, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ReadReplicaID) == "" {
		return nil, errors.New("field ReadReplicaID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/read-replicas/" + fmt.Sprint(req.ReadReplicaID) + "",
	}

	var resp ReadReplica

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteReadReplica: Delete a Read Replica of a Database Instance. You must specify the `region` and `read_replica_id` parameters of the Read Replica you want to delete.
func (s *API) DeleteReadReplica(req *DeleteReadReplicaRequest, opts ...scw.RequestOption) (*ReadReplica, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ReadReplicaID) == "" {
		return nil, errors.New("field ReadReplicaID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/read-replicas/" + fmt.Sprint(req.ReadReplicaID) + "",
	}

	var resp ReadReplica

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ResetReadReplica: When you resync a Read Replica, first it is reset, then its data is resynchronized from the primary node. Your Read Replica remains unavailable during the resync process. The duration of this process is proportional to the size of your Database Instance.
// The configured endpoints do not change.
func (s *API) ResetReadReplica(req *ResetReadReplicaRequest, opts ...scw.RequestOption) (*ReadReplica, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ReadReplicaID) == "" {
		return nil, errors.New("field ReadReplicaID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/read-replicas/" + fmt.Sprint(req.ReadReplicaID) + "/reset",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ReadReplica

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PromoteReadReplica: Promote a Read Replica to Database Instance automatically.
func (s *API) PromoteReadReplica(req *PromoteReadReplicaRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ReadReplicaID) == "" {
		return nil, errors.New("field ReadReplicaID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/read-replicas/" + fmt.Sprint(req.ReadReplicaID) + "/promote",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateReadReplicaEndpoint: Create a new endpoint for a Read Replica. Read Replicas can have at most one direct access and one Private Network endpoint.
func (s *API) CreateReadReplicaEndpoint(req *CreateReadReplicaEndpointRequest, opts ...scw.RequestOption) (*ReadReplica, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ReadReplicaID) == "" {
		return nil, errors.New("field ReadReplicaID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/read-replicas/" + fmt.Sprint(req.ReadReplicaID) + "/endpoints",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ReadReplica

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListInstanceLogs: List the available logs of a Database Instance. By default, the logs returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field.
func (s *API) ListInstanceLogs(req *ListInstanceLogsRequest, opts ...scw.RequestOption) (*ListInstanceLogsResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/logs",
		Query:  query,
	}

	var resp ListInstanceLogsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInstanceLog: Retrieve information about the logs of a Database Instance. Specify the `instance_log_id` and `region` in your request to get information such as `download_url`, `status`, `expires_at` and `created_at` about your logs in the response.
func (s *API) GetInstanceLog(req *GetInstanceLogRequest, opts ...scw.RequestOption) (*InstanceLog, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceLogID) == "" {
		return nil, errors.New("field InstanceLogID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/logs/" + fmt.Sprint(req.InstanceLogID) + "",
	}

	var resp InstanceLog

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PurgeInstanceLogs: Purge a given remote log from a Database Instance. You can specify the `log_name` of the log you wish to clean from your Database Instance.
func (s *API) PurgeInstanceLogs(req *PurgeInstanceLogsRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/purge-logs",
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

// ListInstanceLogsDetails: List remote log details. By default, the details returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field.
func (s *API) ListInstanceLogsDetails(req *ListInstanceLogsDetailsRequest, opts ...scw.RequestOption) (*ListInstanceLogsDetailsResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/logs-details",
	}

	var resp ListInstanceLogsDetailsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddInstanceSettings: Add an advanced setting to a Database Instance. You must set the `name` and the `value` of each setting.
func (s *API) AddInstanceSettings(req *AddInstanceSettingsRequest, opts ...scw.RequestOption) (*AddInstanceSettingsResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/settings",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AddInstanceSettingsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteInstanceSettings: Delete an advanced setting in a Database Instance. You must specify the names of the settings you want to delete in the request.
func (s *API) DeleteInstanceSettings(req *DeleteInstanceSettingsRequest, opts ...scw.RequestOption) (*DeleteInstanceSettingsResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/settings",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DeleteInstanceSettingsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetInstanceSettings: Update an advanced setting for a Database Instance. Settings added upon database engine initalization can only be defined once, and cannot, therefore, be updated.
func (s *API) SetInstanceSettings(req *SetInstanceSettingsRequest, opts ...scw.RequestOption) (*SetInstanceSettingsResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/settings",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetInstanceSettingsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListInstanceACLRules: List the ACL rules for a given Database Instance. The response is an array of ACL objects, each one representing an ACL that denies, allows or redirects traffic based on certain conditions.
func (s *API) ListInstanceACLRules(req *ListInstanceACLRulesRequest, opts ...scw.RequestOption) (*ListInstanceACLRulesResponse, error) {
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

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/acls",
		Query:  query,
	}

	var resp ListInstanceACLRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddInstanceACLRules: Add an additional ACL rule to a Database Instance.
func (s *API) AddInstanceACLRules(req *AddInstanceACLRulesRequest, opts ...scw.RequestOption) (*AddInstanceACLRulesResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/acls",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AddInstanceACLRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetInstanceACLRules: Replace all the ACL rules of a Database Instance.
func (s *API) SetInstanceACLRules(req *SetInstanceACLRulesRequest, opts ...scw.RequestOption) (*SetInstanceACLRulesResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/acls",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetInstanceACLRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteInstanceACLRules: Delete one or more ACL rules of a Database Instance.
func (s *API) DeleteInstanceACLRules(req *DeleteInstanceACLRulesRequest, opts ...scw.RequestOption) (*DeleteInstanceACLRulesResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/acls",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DeleteInstanceACLRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListUsers: List all users of a given Database Instance. By default, the users returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field.
func (s *API) ListUsers(req *ListUsersRequest, opts ...scw.RequestOption) (*ListUsersResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/users",
		Query:  query,
	}

	var resp ListUsersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateUser: Create a new user for a Database Instance. You must define the `name`, `password` and `is_admin` parameters.
func (s *API) CreateUser(req *CreateUserRequest, opts ...scw.RequestOption) (*User, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/users",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp User

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateUser: Update the parameters of a user on a Database Instance. You can update the `password` and `is_admin` parameters, but you cannot change the name of the user.
func (s *API) UpdateUser(req *UpdateUserRequest, opts ...scw.RequestOption) (*User, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	if fmt.Sprint(req.Name) == "" {
		return nil, errors.New("field Name cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/users/" + fmt.Sprint(req.Name) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp User

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteUser: Delete a given user on a Database Instance. You must specify, in the endpoint,  the `region`, `instance_id` and `name` parameters of the user you want to delete.
func (s *API) DeleteUser(req *DeleteUserRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return errors.New("field InstanceID cannot be empty in request")
	}

	if fmt.Sprint(req.Name) == "" {
		return errors.New("field Name cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/users/" + fmt.Sprint(req.Name) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListDatabases: List all databases of a given Database Instance. By default, the databases returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field. You can define additional parameters for your query, such as `name`, `managed` and `owner`.
func (s *API) ListDatabases(req *ListDatabasesRequest, opts ...scw.RequestOption) (*ListDatabasesResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "managed", req.Managed)
	parameter.AddToQuery(query, "owner", req.Owner)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/databases",
		Query:  query,
	}

	var resp ListDatabasesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDatabase: Create a new database. You must define the `name` parameter in the request.
func (s *API) CreateDatabase(req *CreateDatabaseRequest, opts ...scw.RequestOption) (*Database, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/databases",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Database

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDatabase: Delete a given database on a Database Instance. You must specify, in the endpoint, the `region`, `instance_id` and `name` parameters of the database you want to delete.
func (s *API) DeleteDatabase(req *DeleteDatabaseRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return errors.New("field InstanceID cannot be empty in request")
	}

	if fmt.Sprint(req.Name) == "" {
		return errors.New("field Name cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/databases/" + fmt.Sprint(req.Name) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListPrivileges: List privileges of a user on a database. By default, the details returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field. You can define additional parameters for your query, such as `database_name` and `user_name`.
func (s *API) ListPrivileges(req *ListPrivilegesRequest, opts ...scw.RequestOption) (*ListPrivilegesResponse, error) {
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
	parameter.AddToQuery(query, "database_name", req.DatabaseName)
	parameter.AddToQuery(query, "user_name", req.UserName)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/privileges",
		Query:  query,
	}

	var resp ListPrivilegesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetPrivilege: Set the privileges of a user on a database. You must define `database_name`, `user_name` and `permission` in the request body.
func (s *API) SetPrivilege(req *SetPrivilegeRequest, opts ...scw.RequestOption) (*Privilege, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/privileges",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Privilege

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSnapshots: List snapshots. You can include the `instance_id` or `project_id` in your query to get the list of snapshots for specific Database Instances and/or Projects. By default, the details returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListSnapshots(req *ListSnapshotsRequest, opts ...scw.RequestOption) (*ListSnapshotsResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "instance_id", req.InstanceID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/snapshots",
		Query:  query,
	}

	var resp ListSnapshotsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSnapshot: Retrieve information about a given snapshot, specified by its `snapshot_id` and `region`. Full details about the snapshot, like size and expiration date, are returned in the response.
func (s *API) GetSnapshot(req *GetSnapshotRequest, opts ...scw.RequestOption) (*Snapshot, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SnapshotID) == "" {
		return nil, errors.New("field SnapshotID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
	}

	var resp Snapshot

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSnapshot: Create a new snapshot of a Database Instance. You must define the `name` parameter in the request.
func (s *API) CreateSnapshot(req *CreateSnapshotRequest, opts ...scw.RequestOption) (*Snapshot, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("snp")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/snapshots",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Snapshot

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSnapshot: Update the parameters of a snapshot of a Database Instance. You can update the `name` and `expires_at` parameters.
func (s *API) UpdateSnapshot(req *UpdateSnapshotRequest, opts ...scw.RequestOption) (*Snapshot, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SnapshotID) == "" {
		return nil, errors.New("field SnapshotID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Snapshot

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSnapshot: Delete a given snapshot of a Database Instance. You must specify, in the endpoint,  the `region` and `snapshot_id` parameters of the snapshot you want to delete.
func (s *API) DeleteSnapshot(req *DeleteSnapshotRequest, opts ...scw.RequestOption) (*Snapshot, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SnapshotID) == "" {
		return nil, errors.New("field SnapshotID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
	}

	var resp Snapshot

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateInstanceFromSnapshot: Restore a snapshot. When you restore a snapshot, a new Instance is created and billed to your account. Note that is possible to select a larger node type for your new Database Instance. However, the Block volume size will be the same as the size of the restored snapshot. All Instance settings will be restored if you chose a node type with the same or more memory size than the initial Instance. Settings will be reset to the default if your node type has less memory.
func (s *API) CreateInstanceFromSnapshot(req *CreateInstanceFromSnapshotRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SnapshotID) == "" {
		return nil, errors.New("field SnapshotID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "/create-instance",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateEndpoint: Create a new endpoint for a Database Instance. You can add `load_balancer` and `private_network` specifications to the body of the request.
func (s *API) CreateEndpoint(req *CreateEndpointRequest, opts ...scw.RequestOption) (*Endpoint, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/endpoints",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Endpoint

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteEndpoint: Delete the endpoint of a Database Instance. You must specify the `region` and `endpoint_id` parameters of the endpoint you want to delete. Note that might need to update any environment configurations that point to the deleted endpoint.
func (s *API) DeleteEndpoint(req *DeleteEndpointRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.EndpointID) == "" {
		return errors.New("field EndpointID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetEndpoint: Retrieve information about a Database Instance endpoint. Full details about the endpoint, like `ip`, `port`, `private_network` and `load_balancer` specifications are returned in the response.
func (s *API) GetEndpoint(req *GetEndpointRequest, opts ...scw.RequestOption) (*Endpoint, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.EndpointID) == "" {
		return nil, errors.New("field EndpointID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
	}

	var resp Endpoint

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MigrateEndpoint: Migrate an existing Database Instance endpoint to another Database Instance.
func (s *API) MigrateEndpoint(req *MigrateEndpointRequest, opts ...scw.RequestOption) (*Endpoint, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.EndpointID) == "" {
		return nil, errors.New("field EndpointID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/document-db/v1beta1/regions/" + fmt.Sprint(req.Region) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "/migrate",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Endpoint

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
