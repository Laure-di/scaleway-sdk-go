// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package rdb provides methods and message types of the rdb v1 API.
package rdb

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

type APIListDatabaseBackupsRequestOrderBy string

const (
	APIListDatabaseBackupsRequestOrderByCreatedAtAsc  = APIListDatabaseBackupsRequestOrderBy("created_at_asc")
	APIListDatabaseBackupsRequestOrderByCreatedAtDesc = APIListDatabaseBackupsRequestOrderBy("created_at_desc")
	APIListDatabaseBackupsRequestOrderByNameAsc       = APIListDatabaseBackupsRequestOrderBy("name_asc")
	APIListDatabaseBackupsRequestOrderByNameDesc      = APIListDatabaseBackupsRequestOrderBy("name_desc")
	APIListDatabaseBackupsRequestOrderByStatusAsc     = APIListDatabaseBackupsRequestOrderBy("status_asc")
	APIListDatabaseBackupsRequestOrderByStatusDesc    = APIListDatabaseBackupsRequestOrderBy("status_desc")
)

func (enum APIListDatabaseBackupsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum APIListDatabaseBackupsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *APIListDatabaseBackupsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = APIListDatabaseBackupsRequestOrderBy(APIListDatabaseBackupsRequestOrderBy(tmp).String())
	return nil
}

type APIListDatabasesRequestOrderBy string

const (
	APIListDatabasesRequestOrderByNameAsc  = APIListDatabasesRequestOrderBy("name_asc")
	APIListDatabasesRequestOrderByNameDesc = APIListDatabasesRequestOrderBy("name_desc")
	APIListDatabasesRequestOrderBySizeAsc  = APIListDatabasesRequestOrderBy("size_asc")
	APIListDatabasesRequestOrderBySizeDesc = APIListDatabasesRequestOrderBy("size_desc")
)

func (enum APIListDatabasesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum APIListDatabasesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *APIListDatabasesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = APIListDatabasesRequestOrderBy(APIListDatabasesRequestOrderBy(tmp).String())
	return nil
}

type APIListInstanceLogsRequestOrderBy string

const (
	APIListInstanceLogsRequestOrderByCreatedAtAsc  = APIListInstanceLogsRequestOrderBy("created_at_asc")
	APIListInstanceLogsRequestOrderByCreatedAtDesc = APIListInstanceLogsRequestOrderBy("created_at_desc")
)

func (enum APIListInstanceLogsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum APIListInstanceLogsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *APIListInstanceLogsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = APIListInstanceLogsRequestOrderBy(APIListInstanceLogsRequestOrderBy(tmp).String())
	return nil
}

type APIListInstancesRequestOrderBy string

const (
	APIListInstancesRequestOrderByCreatedAtAsc  = APIListInstancesRequestOrderBy("created_at_asc")
	APIListInstancesRequestOrderByCreatedAtDesc = APIListInstancesRequestOrderBy("created_at_desc")
	APIListInstancesRequestOrderByNameAsc       = APIListInstancesRequestOrderBy("name_asc")
	APIListInstancesRequestOrderByNameDesc      = APIListInstancesRequestOrderBy("name_desc")
	APIListInstancesRequestOrderByRegion        = APIListInstancesRequestOrderBy("region")
	APIListInstancesRequestOrderByStatusAsc     = APIListInstancesRequestOrderBy("status_asc")
	APIListInstancesRequestOrderByStatusDesc    = APIListInstancesRequestOrderBy("status_desc")
)

func (enum APIListInstancesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum APIListInstancesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *APIListInstancesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = APIListInstancesRequestOrderBy(APIListInstancesRequestOrderBy(tmp).String())
	return nil
}

type APIListPrivilegesRequestOrderBy string

const (
	APIListPrivilegesRequestOrderByUserNameAsc      = APIListPrivilegesRequestOrderBy("user_name_asc")
	APIListPrivilegesRequestOrderByUserNameDesc     = APIListPrivilegesRequestOrderBy("user_name_desc")
	APIListPrivilegesRequestOrderByDatabaseNameAsc  = APIListPrivilegesRequestOrderBy("database_name_asc")
	APIListPrivilegesRequestOrderByDatabaseNameDesc = APIListPrivilegesRequestOrderBy("database_name_desc")
)

func (enum APIListPrivilegesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "user_name_asc"
	}
	return string(enum)
}

func (enum APIListPrivilegesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *APIListPrivilegesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = APIListPrivilegesRequestOrderBy(APIListPrivilegesRequestOrderBy(tmp).String())
	return nil
}

type APIListSnapshotsRequestOrderBy string

const (
	APIListSnapshotsRequestOrderByCreatedAtAsc  = APIListSnapshotsRequestOrderBy("created_at_asc")
	APIListSnapshotsRequestOrderByCreatedAtDesc = APIListSnapshotsRequestOrderBy("created_at_desc")
	APIListSnapshotsRequestOrderByNameAsc       = APIListSnapshotsRequestOrderBy("name_asc")
	APIListSnapshotsRequestOrderByNameDesc      = APIListSnapshotsRequestOrderBy("name_desc")
	APIListSnapshotsRequestOrderByExpiresAtAsc  = APIListSnapshotsRequestOrderBy("expires_at_asc")
	APIListSnapshotsRequestOrderByExpiresAtDesc = APIListSnapshotsRequestOrderBy("expires_at_desc")
)

func (enum APIListSnapshotsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum APIListSnapshotsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *APIListSnapshotsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = APIListSnapshotsRequestOrderBy(APIListSnapshotsRequestOrderBy(tmp).String())
	return nil
}

type APIListUsersRequestOrderBy string

const (
	APIListUsersRequestOrderByNameAsc     = APIListUsersRequestOrderBy("name_asc")
	APIListUsersRequestOrderByNameDesc    = APIListUsersRequestOrderBy("name_desc")
	APIListUsersRequestOrderByIsAdminAsc  = APIListUsersRequestOrderBy("is_admin_asc")
	APIListUsersRequestOrderByIsAdminDesc = APIListUsersRequestOrderBy("is_admin_desc")
)

func (enum APIListUsersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum APIListUsersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *APIListUsersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = APIListUsersRequestOrderBy(APIListUsersRequestOrderBy(tmp).String())
	return nil
}

type DatabaseBackupStatus string

const (
	DatabaseBackupStatusUnknown   = DatabaseBackupStatus("unknown")
	DatabaseBackupStatusCreating  = DatabaseBackupStatus("creating")
	DatabaseBackupStatusReady     = DatabaseBackupStatus("ready")
	DatabaseBackupStatusRestoring = DatabaseBackupStatus("restoring")
	DatabaseBackupStatusDeleting  = DatabaseBackupStatus("deleting")
	DatabaseBackupStatusError     = DatabaseBackupStatus("error")
	DatabaseBackupStatusExporting = DatabaseBackupStatus("exporting")
	DatabaseBackupStatusLocked    = DatabaseBackupStatus("locked")
)

func (enum DatabaseBackupStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum DatabaseBackupStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DatabaseBackupStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DatabaseBackupStatus(DatabaseBackupStatus(tmp).String())
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
	// PrivateNetworkID: UUID of the private network.
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
	// LoadBalancer: Load balancer details. Public endpoint for Database Instance which is systematically present. One per Database Instance.
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
	// SameZone: Whether the replica is in the same availability zone as the main instance nodes or not.
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

// EndpointSpec:
type EndpointSpec struct {
	// LoadBalancer: Load balancer endpoint specifications. Public endpoint for Database Instance which is systematically present. One per RDB instance.
	LoadBalancer *EndpointSpecLoadBalancer `json:"load_balancer,omitempty"`
	// PrivateNetwork: Private Network endpoint specifications. One maximum per Database Instance or Read Replica (a Database Instance and its Read Replica can have different Private Networks). Cannot be updated (has to be deleted and recreated).
	PrivateNetwork *EndpointSpecPrivateNetwork `json:"private_network,omitempty"`
}

// ReadReplicaEndpointSpec:
type ReadReplicaEndpointSpec struct {
	// DirectAccess: Direct access endpoint specifications. Public endpoint reserved for Read Replicas. One per Read Replica.
	DirectAccess *ReadReplicaEndpointSpecDirectAccess `json:"direct_access,omitempty"`
	// PrivateNetwork: Private Network endpoint specifications. One at the most per Read Replica. Cannot be updated (has to be deleted and recreated).
	PrivateNetwork *ReadReplicaEndpointSpecPrivateNetwork `json:"private_network,omitempty"`
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

// DatabaseBackup:
type DatabaseBackup struct {
	// ID: UUID of the database backup.
	ID string `json:"id"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"instance_id"`
	// DatabaseName: Name of backed up database.
	DatabaseName string `json:"database_name"`
	// Name: Name of the backup.
	Name string `json:"name"`
	// Status: Status of the backup.
	Status DatabaseBackupStatus `json:"status"`
	// Size: Size of the database backup.
	Size *scw.Size `json:"size"`
	// ExpiresAt: Expiration date (must follow the ISO 8601 format).
	ExpiresAt *time.Time `json:"expires_at"`
	// CreatedAt: Creation date (must follow the ISO 8601 format).
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Updated date (must follow the ISO 8601 format).
	UpdatedAt *time.Time `json:"updated_at"`
	// InstanceName: Name of the Database Instance of the backup.
	InstanceName string `json:"instance_name"`
	// DownloadURL: URL you can download the backup from.
	DownloadURL *string `json:"download_url"`
	// DownloadURLExpiresAt: Expiration date of the download link.
	DownloadURLExpiresAt *time.Time `json:"download_url_expires_at"`
	// Region: Region of the database backup.
	Region scw.Region `json:"region"`
	// SameRegion: Store logical backups in the same region as the source Database Instance.
	SameRegion bool `json:"same_region"`
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
	// Engine: Database engine of the database (PostgreSQL, MySQL, ...).
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
	// InitSettings: List of engine settings to be set at database initialization.
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
	// IsHaRequired: The Node Type can be used only with high availability option.
	IsHaRequired bool `json:"is_ha_required"`
	// Generation: Generation associated with the NodeType offer.
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
	// Name: Name of the user (Length must be between 1 and 63 characters for PostgreSQL and between 1 and 32 characters for MySQL. First character must be an alphabet character (a-zA-Z). Your username cannot start with '_rdb' or in PostgreSQL, 'pg_'. Only a-zA-Z0-9_$- characters are accepted).
	Name string `json:"name"`
	// IsAdmin: Defines whether or not a user got administrative privileges on the Database Instance.
	IsAdmin bool `json:"is_admin"`
}

// APIAddInstanceACLRulesRequest:
type APIAddInstanceACLRulesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want to add ACL rules to.
	InstanceID string `json:"-"`
	// Rules: ACL rules to add to the Database Instance.
	Rules []*ACLRuleRequest `json:"rules"`
}

// APIAddInstanceSettingsRequest:
type APIAddInstanceSettingsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want to add settings to.
	InstanceID string `json:"-"`
	// Settings: Settings to add to the Database Instance.
	Settings []*InstanceSetting `json:"settings"`
}

// APICloneInstanceRequest:
type APICloneInstanceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want to clone.
	InstanceID string `json:"-"`
	// Name: Name of the Database Instance clone.
	Name string `json:"name"`
	// NodeType: Node type of the clone.
	NodeType *string `json:"node_type,omitempty"`
}

// APICreateDatabaseBackupRequest:
type APICreateDatabaseBackupRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"instance_id"`
	// DatabaseName: Name of the database you want to back up.
	DatabaseName string `json:"database_name"`
	// Name: Name of the backup.
	Name string `json:"name"`
	// ExpiresAt: Expiration date (must follow the ISO 8601 format).
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// APICreateDatabaseRequest:
type APICreateDatabaseRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance where to create the database.
	InstanceID string `json:"-"`
	// Name: Name of the database.
	Name string `json:"name"`
}

// APICreateEndpointRequest:
type APICreateEndpointRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you to which you want to add an endpoint.
	InstanceID string `json:"-"`
	// EndpointSpec: Specification of the endpoint you want to create.
	EndpointSpec *EndpointSpec `json:"endpoint_spec"`
}

// APICreateInstanceFromSnapshotRequest:
type APICreateInstanceFromSnapshotRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SnapshotID: Block snapshot of the Database Instance.
	SnapshotID string `json:"-"`
	// InstanceName: Name of the Database Instance created with the snapshot.
	InstanceName string `json:"instance_name"`
	// IsHaCluster: Defines whether or not High-Availability is enabled on the new Database Instance.
	IsHaCluster *bool `json:"is_ha_cluster,omitempty"`
	// NodeType: The node type used to restore the snapshot.
	NodeType *string `json:"node_type,omitempty"`
}

// APICreateInstanceRequest:
type APICreateInstanceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Deprecated: OrganizationID: Please use project_id instead.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: The Project ID on which the Database Instance will be created.
	ProjectID *string `json:"project_id,omitempty"`
	// Name: Name of the Database Instance.
	Name string `json:"name"`
	// Engine: Database engine of the Database Instance (PostgreSQL, MySQL, ...).
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

// APICreateReadReplicaEndpointRequest:
type APICreateReadReplicaEndpointRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ReadReplicaID: UUID of the Read Replica.
	ReadReplicaID string `json:"-"`
	// EndpointSpec: Specification of the endpoint you want to create.
	EndpointSpec []*ReadReplicaEndpointSpec `json:"endpoint_spec"`
}

// APICreateReadReplicaRequest:
type APICreateReadReplicaRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want to create a Read Replica from.
	InstanceID string `json:"instance_id"`
	// EndpointSpec: Specification of the endpoint you want to create.
	EndpointSpec []*ReadReplicaEndpointSpec `json:"endpoint_spec"`
	// SameZone: Defines whether to create the replica in the same availability zone as the main instance nodes or not.
	SameZone *bool `json:"same_zone,omitempty"`
}

// APICreateSnapshotRequest:
type APICreateSnapshotRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
	// Name: Name of the snapshot.
	Name string `json:"name"`
	// ExpiresAt: Expiration date (must follow the ISO 8601 format).
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// APICreateUserRequest:
type APICreateUserRequest struct {
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

// APIDeleteDatabaseBackupRequest:
type APIDeleteDatabaseBackupRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// DatabaseBackupID: UUID of the database backup to delete.
	DatabaseBackupID string `json:"-"`
}

// APIDeleteDatabaseRequest:
type APIDeleteDatabaseRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance where to delete the database.
	InstanceID string `json:"-"`
	// Name: Name of the database to delete.
	Name string `json:"-"`
}

// APIDeleteEndpointRequest:
type APIDeleteEndpointRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// EndpointID: This endpoint can also be used to delete a Read Replica endpoint.
	EndpointID string `json:"-"`
}

// APIDeleteInstanceACLRulesRequest:
type APIDeleteInstanceACLRulesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want to delete an ACL rule from.
	InstanceID string `json:"-"`
	// ACLRuleIPs: IP addresses defined in the ACL rules of the Database Instance.
	ACLRuleIPs []string `json:"acl_rule_ips"`
}

// APIDeleteInstanceRequest:
type APIDeleteInstanceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance to delete.
	InstanceID string `json:"-"`
}

// APIDeleteInstanceSettingsRequest:
type APIDeleteInstanceSettingsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance to delete settings from.
	InstanceID string `json:"-"`
	// SettingNames: Settings names to delete.
	SettingNames []string `json:"setting_names"`
}

// APIDeleteReadReplicaRequest:
type APIDeleteReadReplicaRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ReadReplicaID: UUID of the Read Replica.
	ReadReplicaID string `json:"-"`
}

// APIDeleteSnapshotRequest:
type APIDeleteSnapshotRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SnapshotID: UUID of the snapshot to delete.
	SnapshotID string `json:"-"`
}

// APIDeleteUserRequest:
type APIDeleteUserRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance to delete the user from.
	InstanceID string `json:"-"`
	// Name: Name of the user.
	Name string `json:"-"`
}

// APIExportDatabaseBackupRequest:
type APIExportDatabaseBackupRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// DatabaseBackupID: UUID of the database backup you want to export.
	DatabaseBackupID string `json:"-"`
}

// APIGetDatabaseBackupRequest:
type APIGetDatabaseBackupRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// DatabaseBackupID: UUID of the database backup.
	DatabaseBackupID string `json:"-"`
}

// APIGetEndpointRequest:
type APIGetEndpointRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// EndpointID: UUID of the endpoint you want to get.
	EndpointID string `json:"-"`
}

// APIGetInstanceCertificateRequest:
type APIGetInstanceCertificateRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
}

// APIGetInstanceLogRequest:
type APIGetInstanceLogRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceLogID: UUID of the instance_log you want.
	InstanceLogID string `json:"-"`
}

// APIGetInstanceMetricsRequest:
type APIGetInstanceMetricsRequest struct {
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

// APIGetInstanceRequest:
type APIGetInstanceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
}

// APIGetReadReplicaRequest:
type APIGetReadReplicaRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ReadReplicaID: UUID of the Read Replica.
	ReadReplicaID string `json:"-"`
}

// APIGetSnapshotRequest:
type APIGetSnapshotRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SnapshotID: UUID of the snapshot.
	SnapshotID string `json:"-"`
}

// APIListDatabaseBackupsRequest:
type APIListDatabaseBackupsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Name: Name of the database backups.
	Name *string `json:"-"`
	// OrderBy: Criteria to use when ordering database backups listing.
	OrderBy APIListDatabaseBackupsRequestOrderBy `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID *string `json:"-"`
	// OrganizationID: Organization ID of the Organization the database backups belong to.
	OrganizationID *string `json:"-"`
	// ProjectID: Project ID of the Project the database backups belong to.
	ProjectID *string `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
}

// APIListDatabaseEnginesRequest:
type APIListDatabaseEnginesRequest struct {
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

// APIListDatabasesRequest:
type APIListDatabasesRequest struct {
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
	OrderBy APIListDatabasesRequestOrderBy `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
}

// APIListInstanceACLRulesRequest:
type APIListInstanceACLRulesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
}

// APIListInstanceLogsDetailsRequest:
type APIListInstanceLogsDetailsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want logs of.
	InstanceID string `json:"-"`
}

// APIListInstanceLogsRequest:
type APIListInstanceLogsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want logs of.
	InstanceID string `json:"-"`
	// OrderBy: Criteria to use when ordering Database Instance logs listing.
	OrderBy APIListInstanceLogsRequestOrderBy `json:"-"`
}

// APIListInstancesRequest:
type APIListInstancesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Tags: List Database Instances that have a given tag.
	Tags []string `json:"-"`
	// Name: Lists Database Instances that match a name pattern.
	Name *string `json:"-"`
	// OrderBy: Criteria to use when ordering Database Instance listings.
	OrderBy APIListInstancesRequestOrderBy `json:"-"`
	// OrganizationID: Please use project_id instead.
	OrganizationID *string `json:"-"`
	// ProjectID: Project ID to list the Database Instance of.
	ProjectID *string `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
}

// APIListNodeTypesRequest:
type APIListNodeTypesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// IncludeDisabledTypes: Defines whether or not to include disabled types.
	IncludeDisabledTypes bool `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
}

// APIListPrivilegesRequest:
type APIListPrivilegesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
	// OrderBy: Criteria to use when ordering privileges listing.
	OrderBy APIListPrivilegesRequestOrderBy `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
	// DatabaseName: Name of the database.
	DatabaseName *string `json:"-"`
	// UserName: Name of the user.
	UserName *string `json:"-"`
}

// APIListSnapshotsRequest:
type APIListSnapshotsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Name: Name of the snapshot.
	Name *string `json:"-"`
	// OrderBy: Criteria to use when ordering snapshot listing.
	OrderBy APIListSnapshotsRequestOrderBy `json:"-"`
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

// APIListUsersRequest:
type APIListUsersRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
	// Name: Name of the user.
	Name *string `json:"-"`
	// OrderBy: Criteria to use when requesting user listing.
	OrderBy APIListUsersRequestOrderBy `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
}

// APIMigrateEndpointRequest:
type APIMigrateEndpointRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// EndpointID: UUID of the endpoint you want to migrate.
	EndpointID string `json:"-"`
	// InstanceID: UUID of the instance you want to attach the endpoint to.
	InstanceID string `json:"instance_id"`
}

// APIPrepareInstanceLogsRequest:
type APIPrepareInstanceLogsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want logs of.
	InstanceID string `json:"-"`
	// StartDate: Start datetime of your log. (RFC 3339 format).
	StartDate *time.Time `json:"start_date,omitempty"`
	// EndDate: End datetime of your log. (RFC 3339 format).
	EndDate *time.Time `json:"end_date,omitempty"`
}

// APIPromoteReadReplicaRequest:
type APIPromoteReadReplicaRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ReadReplicaID: UUID of the Read Replica.
	ReadReplicaID string `json:"-"`
}

// APIPurgeInstanceLogsRequest:
type APIPurgeInstanceLogsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want logs of.
	InstanceID string `json:"-"`
	// LogName: Given log name to purge.
	LogName *string `json:"log_name,omitempty"`
}

// APIRenewInstanceCertificateRequest:
type APIRenewInstanceCertificateRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want logs of.
	InstanceID string `json:"-"`
}

// APIResetReadReplicaRequest:
type APIResetReadReplicaRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ReadReplicaID: UUID of the Read Replica.
	ReadReplicaID string `json:"-"`
}

// APIRestartInstanceRequest:
type APIRestartInstanceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want to restart.
	InstanceID string `json:"-"`
}

// APIRestoreDatabaseBackupRequest:
type APIRestoreDatabaseBackupRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// DatabaseBackupID: Backup of a logical database.
	DatabaseBackupID string `json:"-"`
	// DatabaseName: Defines the destination database to restore into a specified database (the default destination is set to the origin database of the backup).
	DatabaseName *string `json:"database_name,omitempty"`
	// InstanceID: Defines the Database Instance where the backup has to be restored.
	InstanceID string `json:"instance_id"`
}

// APISetInstanceACLRulesRequest:
type APISetInstanceACLRulesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance where the ACL rules must be set.
	InstanceID string `json:"-"`
	// Rules: ACL rules to define for the Database Instance.
	Rules []*ACLRuleRequest `json:"rules"`
}

// APISetInstanceSettingsRequest:
type APISetInstanceSettingsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance where the settings must be set.
	InstanceID string `json:"-"`
	// Settings: Settings to define for the Database Instance.
	Settings []*InstanceSetting `json:"settings"`
}

// APISetPrivilegeRequest:
type APISetPrivilegeRequest struct {
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

// APIUpdateDatabaseBackupRequest:
type APIUpdateDatabaseBackupRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// DatabaseBackupID: UUID of the database backup to update.
	DatabaseBackupID string `json:"-"`
	// Name: Name of the Database Backup.
	Name *string `json:"name,omitempty"`
	// ExpiresAt: Expiration date (must follow the ISO 8601 format).
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// APIUpdateInstanceRequest:
type APIUpdateInstanceRequest struct {
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

// APIUpdateSnapshotRequest:
type APIUpdateSnapshotRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SnapshotID: UUID of the snapshot to update.
	SnapshotID string `json:"-"`
	// Name: Name of the snapshot.
	Name *string `json:"name,omitempty"`
	// ExpiresAt: Expiration date (must follow the ISO 8601 format).
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// APIUpdateUserRequest:
type APIUpdateUserRequest struct {
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

// APIUpgradeInstanceRequest:
type APIUpgradeInstanceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want to upgrade.
	InstanceID string `json:"-"`
	// NodeType: Node type of the Database Instance you want to upgrade to.
	NodeType *string `json:"node_type,omitempty"`
	// EnableHa: Defines whether or not high availability should be enabled on the Database Instance.
	EnableHa *bool `json:"enable_ha,omitempty"`
	// VolumeSize: Increase your block storage volume size.
	VolumeSize *uint64 `json:"volume_size,omitempty"`
	// VolumeType: Change your Database Instance storage type.
	VolumeType *VolumeType `json:"volume_type,omitempty"`
	// UpgradableVersionID: This will create a new Database Instance with same specifications as the current one and perform a Database Engine upgrade.
	UpgradableVersionID *string `json:"upgradable_version_id,omitempty"`
}

// AddInstanceACLRulesResponse:
type AddInstanceACLRulesResponse struct {
	// Rules: ACL Rules enabled for the Database Instance.
	Rules []*ACLRule `json:"rules"`
}

// AddInstanceSettingsResponse:
type AddInstanceSettingsResponse struct {
	// Settings: Settings available on the Database Instance.
	Settings []*InstanceSetting `json:"settings"`
}

// DeleteInstanceACLRulesResponse:
type DeleteInstanceACLRulesResponse struct {
	// Rules: IP addresses defined in the ACL rules of the Database Instance.
	Rules []*ACLRule `json:"rules"`
}

// DeleteInstanceSettingsResponse:
type DeleteInstanceSettingsResponse struct {
	// Settings: Settings names to delete from the Database Instance.
	Settings []*InstanceSetting `json:"settings"`
}

// InstanceMetrics:
type InstanceMetrics struct {
	// Timeseries: Time series of metrics of a Database Instance.
	Timeseries []*scw.TimeSeries `json:"timeseries"`
}

// ListDatabaseBackupsResponse:
type ListDatabaseBackupsResponse struct {
	// DatabaseBackups: List of database backups.
	DatabaseBackups []*DatabaseBackup `json:"database_backups"`
	// TotalCount: Total count of database backups available.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDatabaseBackupsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDatabaseBackupsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDatabaseBackupsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.DatabaseBackups = append(r.DatabaseBackups, results.DatabaseBackups...)
	r.TotalCount += uint32(len(results.DatabaseBackups))
	return uint32(len(results.DatabaseBackups)), nil
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

// ListInstanceLogsDetailsResponse:
type ListInstanceLogsDetailsResponse struct {
	// Details: Remote Database Instance logs details.
	Details []*ListInstanceLogsDetailsResponseInstanceLogDetail `json:"details"`
}

// ListInstanceLogsResponse:
type ListInstanceLogsResponse struct {
	// InstanceLogs: Available logs in a Database Instance.
	InstanceLogs []*InstanceLog `json:"instance_logs"`
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

// PrepareInstanceLogsResponse:
type PrepareInstanceLogsResponse struct {
	// InstanceLogs: Instance logs for a Database Instance between a start and an end date.
	InstanceLogs []*InstanceLog `json:"instance_logs"`
}

// SetInstanceACLRulesResponse:
type SetInstanceACLRulesResponse struct {
	// Rules: ACLs rules configured for a Database Instance.
	Rules []*ACLRule `json:"rules"`
}

// SetInstanceSettingsResponse:
type SetInstanceSettingsResponse struct {
	// Settings: Settings configured for a Database Instance.
	Settings []*InstanceSetting `json:"settings"`
}

// Managed Database for PostgreSQL and MySQL provides fully-managed relational Database Instances, with MySQL or PostgreSQL as database engines. The resource allows you to focus on development rather than administration or configuration. It comes with a high-availability mode, data replication, and automatic backups.
//
// Compared to traditional database management, which requires customers to provide their infrastructure and resources to manage their databases, Managed Database for PostgreSQL and MySQL Instance offers the user access to Database Instances without setting up the hardware or configuring the software. Scaleway handles the provisioning, manages the configuration, and provides useful features as high availability, automated backup, user management, and more.
//
// (switchcolumn)
// (switchcolumn)
//
// # Concepts
//
// Refer to our [dedicated concepts page](https://www.scaleway.com/en/docs/managed-databases/postgresql-and-mysql/concepts/) to find definitions of the different terms referring to Managed Database for PostgreSQL and MySQL.
//
// (switchcolumn)
// (switchcolumn)
//
// # Quickstart
//
//  1. Configure your environment variables.
//     <Message type="note">
//     This is an optional step that seeks to simplify your usage of the APIs.
//     </Message>
//
//     ```bash
//     export SCW_ACCESS_KEY="<API access key>"
//     export SCW_SECRET_KEY="<API secret key>"
//     export SCW_REGION="<Scaleway region>"
//     ```
//
//  2. Edit the POST request payload you will use to create your Database Instance. Replace the parameters in the following example:
//     ```json
//     '{
//     "project_id": "d8e65f2b-cce9-40b7-80fc-6a2902db6826",
//     "name": "myDB",
//     "engine": "PostgreSQL-15",
//     "tags": ["donnerstag"],
//     "is_ha_cluster": true,
//     "node_type": "db-pro2-xxs",
//     "disable_backup": false,
//     "user_name": "my_initial_user",
//     "password": "thiZ_is_v0ry_s3cret",
//     "volume_type": "bssd",
//     "volume_size": "30000000000"
//     }'
//     ```
//
//     | Parameter        | Description                                                        |
//     | :--------------- | :----------------------------------------------------------------- |
//     | `project_id`     | The ID of the Project you want to create your Database Instance in. To find your Project ID you can **[list the projects](/api/account#path-projects-list-all-projects-of-an-organization)** or consult the **[Scaleway console](https://console.scaleway.com/project/settings)**. |
//     | `engine`         | **REQUIRED** Version ID of the database engine. To check the list of available engines you can use the folowing endpoint: `https://api.scaleway.com/rdb/v1/regions/$SCW_REGION/database-engines`                                           |
//     | `name`           | Name of the Database Instance                                          |
//     | `node_type`      | **REQUIRED** The node type. To check the list of available node types you can use the folowing endpoint: `https://api.scaleway.com/rdb/v1/regions/$SCW_REGION/node-types`                             |
//     | `is_ha_cluster`  | **BOOLEAN** Defines whether High Availability is enabled for the Database Instance   |
//     | `disable_backup` | **BOOLEAN** Defines whether automated backups are disabled for the Database Instance                 |
//     | `tags`           | The list of tags `["tag1", "tag2", ...]` that will be associated with the Database Instance. Tags can be appended to the query of the [List Database Instances](#path-database-instances-list-database-instances) call to show results for only the Database Instances using a specific tag. You can also combine tags to list Database Instances that posess all the appended tags. |
//     | `user_name`      | **REQUIRED** Identifier of the default user, which is created concurrently with the Database Instance |
//     | `password`       | **REQUIRED** Password for the default user |
//     | `volume_type`    | Type of volume where data is stored. You can specify either local volume (`lssd`) or block volume (`bssd`). The default value is `lssd` |
//     | `volume_size`    | Volume size when volume_type is `bssd`. The value should be expressed in bytes. For example 30GB is expressed as 30000000000 |
//
//  3. Create a Database Instance by running the following command. Make sure you include the payload you edited in the previous step.
//     ```bash
//     curl -X POST \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     "Content-Type: application/json" \
//     https://api.scaleway.com/rdb/v1/regions/$SCW_REGION/instances \
//     -d '{
//     "project_id": "d8e65f2b-cce9-40b7-80fc-6a2902db6826",
//     "name": "myDB",
//     "engine": "PostgreSQL-15",
//     "tags": ["donnerstag"],
//     "is_ha_cluster": true,
//     "node_type": "db-pro2-xxs",
//     "disable_backup": false,
//     "user_name": "my_initial_user",
//     "password": "thiZ_is_v0ry_s3cret",
//     "volume_type": "bssd",
//     "volume_size": "30000000000"
//     }'
//     ```
//
//  4. List your Database Instances.
//     ```bash
//     curl -X GET \
//     -H "Content-Type: application/json" \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" https://api.scaleway.com/rdb/v1/regions/$SCW_REGION/instances
//     ```
//
//     You should get a response like the following:
//
//     <Message type="note">
//     This is a response example, the UUIDs and IP address displayed are not real.
//     </Message>
//
//     ```json
//     {
//     "id": "f5122f66-fb50-4cef-aa02-487ef4fc1af0",
//     "name": "myDB",
//     "organization_id": "895693aa-3915-4896-8761-c2923b008be7",
//     "project_id": "d8e65f2b-cce9-40b7-80fc-6a2902db6826",
//     "status": "ready",
//     "engine": "PostgreSQL-15",
//     "endpoint": {
//     "ip": "198.51.100.0",
//     "port": 22245,
//     "name": null
//     },
//     "tags": [
//     "donnerstag"
//     ],
//     "settings": [],
//     "backup_schedule": {
//     "frequency": 24,
//     "retention": 7,
//     "disabled": true
//     },
//     "is_ha_cluster": true,
//     "read_replicas": [],
//     "node_type": "db-pro2-xxs",
//     "volume": {
//     "type": "bssd",
//     "size": 30000000000
//     }
//     "created_at": "2019-04-19T16:24:52.591417Z",
//     "region": "fr-par"
//     }
//     ```
//
//  5. Retrieve your Database Instance IP and port from the response.
//     <Message type="note">
//     In the example above, the IP and port are `198.51.100.0` and `22245`, respectively.
//     </Message>
//
//  6. Connect to your Database Instance with the database client of the engine you selected.
//     For MySQL, run the following command:
//     ```bash
//     mysql -h <ip-address> --port <port> -p -u <user_name>
//     ```
//
//     For PostgreSQL, run:
//     ```bash
//     psql -h <ip-address> -p <port> -U <username> -d rdb
//     ```
//
//     For the recurring example, the command would look like:
//
//     ```bash
//     psql -h 198.51.100.0 -p 22245 -U my_initial_user -d rdb
//     ```
//
// 7. Enter the database password that you defined upon creation.
//
// You are now connected to your Managed Database.
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
//
// # Technical Information
//
// ## Regions
//
// Scaleway's infrastructure is spread across different [regions and Availability Zones](https://www.scaleway.com/en/docs/console/my-account/reference-content/products-availability/).
//
// Managed Database for PostgreSQL and MySQL is available in the Paris, Amsterdam and Warsaw regions, which are represented by the following path parameters:
//
// - `fr-par`
// - `nl-ams`
// - `pl-waw`
//
// ## PostgreSQL specifications
//
// ### Versions
//
// Scaleway Database for PostgreSQL supports PostgreSQL versions 11, 12, 13, 14 and 15.
//
// ### System
//
// Different modules are available for installation, including TimescaleDB and PostGIS. Refer to the [Managed Database for PostgreSQL and MySQL FAQ page](https://www.scaleway.com/en/docs/faq/databases-for-postgresql-and-mysql/#which-postgresql-extensions-are-available) for an extensive list of PostgreSQL extensions.
//
// ### Database Management
//
// You can create logical databases through the Scaleway console, the Scaleway APIs or SQL.
//
// - databases created using the Scaleway console or the API are owned by an internal system user. These are called "managed databases".
// - databases created using SQL will be owned by the creator. These are called "unmanaged databases".
//
// ## MySQL specifications
//
// ### Versions
//
// Scaleway Database for MySQL supports MySQL 8.
//
// ### System
//
// - only the [InnoDB engine](https://dev.mysql.com/doc/refman/8.0/en/innodb-storage-engine.html) is supported
// - the [Global Transaction Identifier (GTID)](https://dev.mysql.com/doc/refman/8.0/en/replication-gtids-concepts.html) is enabled.
// - [`mysql_native_password`](https://dev.mysql.com/doc/refman/8.0/en/native-pluggable-authentication.html) (default) and [`caching_sha2_password`](https://dev.mysql.com/doc/refman/8.0/en/caching-sha2-pluggable-authentication.html) authentication are supported.
//
// ### User Management
//
// - users with an `admin` role have access to all logical databases and can create new ones.
// - users created via the API are authenticated using the default authentication plugin, which can be changed in the settings.
//
// # Technical Limitations
//
// ## PostgreSQL
//
// ### User Management
//
// - users with an `admin` role have `CREATEROLE` and `CREATEDB` privileges.
// - users do NOT have `SUPERUSER` nor `REPLICATION` privileges.
// - permission management through the Scaleway console or API is only possible for the "managed databases".
//
// ### Backup and restoration
//
// Databases that have been backed up and then restored retain the user permission settings in use at the time of backup. If you delete users after backup and then restore your backup in the same database, or if you restore a backup to a different database with different or no users, the permissions configured for them continue to exist, but with no associated owner. This error will put a stop to the restoration process.
//
// To avoid this issue, we recommmend you re-create the users you deleted. In the occasion you restore the backup to a new database, you must create new users with the same names.
//
// # Going Further
//
// For more information about Managed Database for PostgreSQL and MySQL, you can check out the following pages:
//
// * [Managed Database for PostgreSQL and MySQL Documentation](https://www.scaleway.com/en/docs/managed-databases/postgresql-and-mysql/)
// * [Managed Database for PostgreSQL and MySQL FAQ](https://www.scaleway.com/en/docs/faq/databases-for-postgresql-and-mysql/)
// * [Scaleway Slack Community](https://scaleway-community.slack.com/) join the #database channel
// * [Contact our support team](https://console.scaleway.com/support/tickets)
//
// ## How to migrate a database
//
// If you wish to migrate existing databases to a Managed Database for PostgreSQL or MySQL, you can refer to the [Migrating existing databases to a Database Instance](https://www.scaleway.com/en/docs/tutorials/migrate-databases-instance/) tutorial page.
//
// ## Troubleshoooting
//
// ### Disk full status
//
// If your Database Instance uses local storage, your local volume might eventually approach full capacity and shift to `disk_full` mode. This mode grants you enough space to either [upgrade your node type](https://www.scaleway.com/en/docs/managed-databases/postgresql-and-mysql/how-to/upgrade-a-database/#how-to-change-the-node-type) or [clear out space in your volume](https://www.scaleway.com/en/docs/managed-databases/postgresql-and-mysql/troubleshooting/disk-full/).
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

// ListDatabaseEngines: List the PostgreSQL and MySQL database engines available at Scaleway.
func (s *API) ListDatabaseEngines(req *APIListDatabaseEnginesRequest, opts ...scw.RequestOption) (*ListDatabaseEnginesResponse, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/database-engines",
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
func (s *API) ListNodeTypes(req *APIListNodeTypesRequest, opts ...scw.RequestOption) (*ListNodeTypesResponse, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/node-types",
		Query:  query,
	}

	var resp ListNodeTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDatabaseBackups: List all backups in a specified region, for a given Scaleway Organization or Scaleway Project. By default, the backups listed are ordered by creation date in ascending order. This can be modified via the `order_by` field.
func (s *API) ListDatabaseBackups(req *APIListDatabaseBackupsRequest, opts ...scw.RequestOption) (*ListDatabaseBackupsResponse, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/backups",
		Query:  query,
	}

	var resp ListDatabaseBackupsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDatabaseBackup: Create a new backup. You must set the `instance_id`, `database_name`, `name` and `expires_at` parameters.
func (s *API) CreateDatabaseBackup(req *APICreateDatabaseBackupRequest, opts ...scw.RequestOption) (*DatabaseBackup, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("bkp")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/backups",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DatabaseBackup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDatabaseBackup: Retrieve information about a given backup, specified by its database backup ID and region. Full details about the backup, like size, URL and expiration date, are returned in the response.
func (s *API) GetDatabaseBackup(req *APIGetDatabaseBackupRequest, opts ...scw.RequestOption) (*DatabaseBackup, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DatabaseBackupID) == "" {
		return nil, errors.New("field DatabaseBackupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/backups/" + fmt.Sprint(req.DatabaseBackupID) + "",
	}

	var resp DatabaseBackup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateDatabaseBackup: Update the parameters of a backup, including name and expiration date.
func (s *API) UpdateDatabaseBackup(req *APIUpdateDatabaseBackupRequest, opts ...scw.RequestOption) (*DatabaseBackup, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DatabaseBackupID) == "" {
		return nil, errors.New("field DatabaseBackupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/backups/" + fmt.Sprint(req.DatabaseBackupID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DatabaseBackup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDatabaseBackup: Delete a backup, specified by its database backup ID and region. Deleting a backup is permanent, and cannot be undone.
func (s *API) DeleteDatabaseBackup(req *APIDeleteDatabaseBackupRequest, opts ...scw.RequestOption) (*DatabaseBackup, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DatabaseBackupID) == "" {
		return nil, errors.New("field DatabaseBackupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/backups/" + fmt.Sprint(req.DatabaseBackupID) + "",
	}

	var resp DatabaseBackup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RestoreDatabaseBackup: Launch the process of restoring database backup. You must specify the `instance_id` of the Database Instance of destination, where the backup will be restored. Note that large database backups can take up to several hours to restore.
func (s *API) RestoreDatabaseBackup(req *APIRestoreDatabaseBackupRequest, opts ...scw.RequestOption) (*DatabaseBackup, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DatabaseBackupID) == "" {
		return nil, errors.New("field DatabaseBackupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/backups/" + fmt.Sprint(req.DatabaseBackupID) + "/restore",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DatabaseBackup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExportDatabaseBackup: Export a backup, specified by the `database_backup_id` and the `region` parameters. The download URL is returned in the response.
func (s *API) ExportDatabaseBackup(req *APIExportDatabaseBackupRequest, opts ...scw.RequestOption) (*DatabaseBackup, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DatabaseBackupID) == "" {
		return nil, errors.New("field DatabaseBackupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/backups/" + fmt.Sprint(req.DatabaseBackupID) + "/export",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DatabaseBackup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpgradeInstance: Upgrade your current Database Instance specifications like node type, high availability, volume, or the database engine version. Note that upon upgrade the `enable_ha` parameter can only be set to `true`.
func (s *API) UpgradeInstance(req *APIUpgradeInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/upgrade",
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
func (s *API) ListInstances(req *APIListInstancesRequest, opts ...scw.RequestOption) (*ListInstancesResponse, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances",
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
func (s *API) GetInstance(req *APIGetInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "",
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateInstance: Create a new Database Instance. You must set the `engine`, `user_name`, `password` and `node_type` parameters. Optionally, you can specify the volume type and size.
func (s *API) CreateInstance(req *APICreateInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances",
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
func (s *API) UpdateInstance(req *APIUpdateInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "",
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
func (s *API) DeleteInstance(req *APIDeleteInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "",
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CloneInstance: Clone a given Database Instance, specified by the `region` and `instance_id` parameters. The clone feature allows you to create a new Database Instance from an existing one. The clone includes all existing databases, users and permissions. You can create a clone on a Database Instance bigger than your current one.
func (s *API) CloneInstance(req *APICloneInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/clone",
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
func (s *API) RestartInstance(req *APIRestartInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/restart",
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
func (s *API) GetInstanceCertificate(req *APIGetInstanceCertificateRequest, opts ...scw.RequestOption) (*scw.File, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/certificate",
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RenewInstanceCertificate: Renew a TLS for a Database Instance. Renewing a certificate means that you will not be able to connect to your Database Instance using the previous certificate. You will also need to download and update the new certificate for all database clients.
func (s *API) RenewInstanceCertificate(req *APIRenewInstanceCertificateRequest, opts ...scw.RequestOption) error {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/renew-certificate",
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
func (s *API) GetInstanceMetrics(req *APIGetInstanceMetricsRequest, opts ...scw.RequestOption) (*InstanceMetrics, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/metrics",
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
func (s *API) CreateReadReplica(req *APICreateReadReplicaRequest, opts ...scw.RequestOption) (*ReadReplica, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/read-replicas",
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
func (s *API) GetReadReplica(req *APIGetReadReplicaRequest, opts ...scw.RequestOption) (*ReadReplica, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/read-replicas/" + fmt.Sprint(req.ReadReplicaID) + "",
	}

	var resp ReadReplica

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteReadReplica: Delete a Read Replica of a Database Instance. You must specify the `region` and `read_replica_id` parameters of the Read Replica you want to delete.
func (s *API) DeleteReadReplica(req *APIDeleteReadReplicaRequest, opts ...scw.RequestOption) (*ReadReplica, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/read-replicas/" + fmt.Sprint(req.ReadReplicaID) + "",
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
func (s *API) ResetReadReplica(req *APIResetReadReplicaRequest, opts ...scw.RequestOption) (*ReadReplica, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/read-replicas/" + fmt.Sprint(req.ReadReplicaID) + "/reset",
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
func (s *API) PromoteReadReplica(req *APIPromoteReadReplicaRequest, opts ...scw.RequestOption) (*Instance, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/read-replicas/" + fmt.Sprint(req.ReadReplicaID) + "/promote",
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
func (s *API) CreateReadReplicaEndpoint(req *APICreateReadReplicaEndpointRequest, opts ...scw.RequestOption) (*ReadReplica, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/read-replicas/" + fmt.Sprint(req.ReadReplicaID) + "/endpoints",
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

// PrepareInstanceLogs: Prepare your Database Instance logs. You can define the `start_date` and `end_date` parameters for your query. The download URL is returned in the response. Logs are recorded from 00h00 to 23h59 and then aggregated in a `.log` file once a day. Therefore, even if you specify a timeframe from which you want to get the logs, you will receive logs from the full 24 hours.
func (s *API) PrepareInstanceLogs(req *APIPrepareInstanceLogsRequest, opts ...scw.RequestOption) (*PrepareInstanceLogsResponse, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/prepare-logs",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PrepareInstanceLogsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListInstanceLogs: List the available logs of a Database Instance. By default, the logs returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field.
func (s *API) ListInstanceLogs(req *APIListInstanceLogsRequest, opts ...scw.RequestOption) (*ListInstanceLogsResponse, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/logs",
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
func (s *API) GetInstanceLog(req *APIGetInstanceLogRequest, opts ...scw.RequestOption) (*InstanceLog, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/logs/" + fmt.Sprint(req.InstanceLogID) + "",
	}

	var resp InstanceLog

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PurgeInstanceLogs: Purge a given remote log from a Database Instance. You can specify the `log_name` of the log you wish to clean from your Database Instance.
func (s *API) PurgeInstanceLogs(req *APIPurgeInstanceLogsRequest, opts ...scw.RequestOption) error {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/purge-logs",
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
func (s *API) ListInstanceLogsDetails(req *APIListInstanceLogsDetailsRequest, opts ...scw.RequestOption) (*ListInstanceLogsDetailsResponse, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/logs-details",
	}

	var resp ListInstanceLogsDetailsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddInstanceSettings: Add an advanced setting to a Database Instance. You must set the `name` and the `value` of each setting.
func (s *API) AddInstanceSettings(req *APIAddInstanceSettingsRequest, opts ...scw.RequestOption) (*AddInstanceSettingsResponse, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/settings",
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
func (s *API) DeleteInstanceSettings(req *APIDeleteInstanceSettingsRequest, opts ...scw.RequestOption) (*DeleteInstanceSettingsResponse, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/settings",
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
func (s *API) SetInstanceSettings(req *APISetInstanceSettingsRequest, opts ...scw.RequestOption) (*SetInstanceSettingsResponse, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/settings",
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
func (s *API) ListInstanceACLRules(req *APIListInstanceACLRulesRequest, opts ...scw.RequestOption) (*ListInstanceACLRulesResponse, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/acls",
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
func (s *API) AddInstanceACLRules(req *APIAddInstanceACLRulesRequest, opts ...scw.RequestOption) (*AddInstanceACLRulesResponse, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/acls",
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
func (s *API) SetInstanceACLRules(req *APISetInstanceACLRulesRequest, opts ...scw.RequestOption) (*SetInstanceACLRulesResponse, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/acls",
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
func (s *API) DeleteInstanceACLRules(req *APIDeleteInstanceACLRulesRequest, opts ...scw.RequestOption) (*DeleteInstanceACLRulesResponse, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/acls",
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
func (s *API) ListUsers(req *APIListUsersRequest, opts ...scw.RequestOption) (*ListUsersResponse, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/users",
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
func (s *API) CreateUser(req *APICreateUserRequest, opts ...scw.RequestOption) (*User, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/users",
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
func (s *API) UpdateUser(req *APIUpdateUserRequest, opts ...scw.RequestOption) (*User, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/users/" + fmt.Sprint(req.Name) + "",
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
func (s *API) DeleteUser(req *APIDeleteUserRequest, opts ...scw.RequestOption) error {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/users/" + fmt.Sprint(req.Name) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListDatabases: List all databases of a given Database Instance. By default, the databases returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field. You can define additional parameters for your query, such as `name`, `managed` and `owner`.
func (s *API) ListDatabases(req *APIListDatabasesRequest, opts ...scw.RequestOption) (*ListDatabasesResponse, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/databases",
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
func (s *API) CreateDatabase(req *APICreateDatabaseRequest, opts ...scw.RequestOption) (*Database, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/databases",
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
func (s *API) DeleteDatabase(req *APIDeleteDatabaseRequest, opts ...scw.RequestOption) error {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/databases/" + fmt.Sprint(req.Name) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListPrivileges: List privileges of a user on a database. By default, the details returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field. You can define additional parameters for your query, such as `database_name` and `user_name`.
func (s *API) ListPrivileges(req *APIListPrivilegesRequest, opts ...scw.RequestOption) (*ListPrivilegesResponse, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/privileges",
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
func (s *API) SetPrivilege(req *APISetPrivilegeRequest, opts ...scw.RequestOption) (*Privilege, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/privileges",
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
func (s *API) ListSnapshots(req *APIListSnapshotsRequest, opts ...scw.RequestOption) (*ListSnapshotsResponse, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/snapshots",
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
func (s *API) GetSnapshot(req *APIGetSnapshotRequest, opts ...scw.RequestOption) (*Snapshot, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
	}

	var resp Snapshot

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSnapshot: Create a new snapshot of a Database Instance. You must define the `name` parameter in the request.
func (s *API) CreateSnapshot(req *APICreateSnapshotRequest, opts ...scw.RequestOption) (*Snapshot, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/snapshots",
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
func (s *API) UpdateSnapshot(req *APIUpdateSnapshotRequest, opts ...scw.RequestOption) (*Snapshot, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
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
func (s *API) DeleteSnapshot(req *APIDeleteSnapshotRequest, opts ...scw.RequestOption) (*Snapshot, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
	}

	var resp Snapshot

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateInstanceFromSnapshot: Restore a snapshot. When you restore a snapshot, a new Instance is created and billed to your account. Note that is possible to select a larger node type for your new Database Instance. However, the Block volume size will be the same as the size of the restored snapshot. All Instance settings will be restored if you chose a node type with the same or more memory size than the initial Instance. Settings will be reset to the default if your node type has less memory.
func (s *API) CreateInstanceFromSnapshot(req *APICreateInstanceFromSnapshotRequest, opts ...scw.RequestOption) (*Instance, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "/create-instance",
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
func (s *API) CreateEndpoint(req *APICreateEndpointRequest, opts ...scw.RequestOption) (*Endpoint, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/endpoints",
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
func (s *API) DeleteEndpoint(req *APIDeleteEndpointRequest, opts ...scw.RequestOption) error {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetEndpoint: Retrieve information about a Database Instance endpoint. Full details about the endpoint, like `ip`, `port`, `private_network` and `load_balancer` specifications are returned in the response.
func (s *API) GetEndpoint(req *APIGetEndpointRequest, opts ...scw.RequestOption) (*Endpoint, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
	}

	var resp Endpoint

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MigrateEndpoint: Migrate an existing instance endpoint to another instance.
func (s *API) MigrateEndpoint(req *APIMigrateEndpointRequest, opts ...scw.RequestOption) (*Endpoint, error) {
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
		Path:   "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "/migrate",
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
