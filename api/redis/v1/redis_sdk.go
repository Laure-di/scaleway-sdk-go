// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package redis provides methods and message types of the redis v1 API.
package redis

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

type AvailableClusterSettingPropertyType string

const (
	AvailableClusterSettingPropertyTypeUNKNOWN = AvailableClusterSettingPropertyType("UNKNOWN")
	AvailableClusterSettingPropertyTypeBOOLEAN = AvailableClusterSettingPropertyType("BOOLEAN")
	AvailableClusterSettingPropertyTypeINT     = AvailableClusterSettingPropertyType("INT")
	AvailableClusterSettingPropertyTypeSTRING  = AvailableClusterSettingPropertyType("STRING")
)

func (enum AvailableClusterSettingPropertyType) String() string {
	if enum == "" {
		// return default value if empty
		return "UNKNOWN"
	}
	return string(enum)
}

func (enum AvailableClusterSettingPropertyType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AvailableClusterSettingPropertyType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AvailableClusterSettingPropertyType(AvailableClusterSettingPropertyType(tmp).String())
	return nil
}

type ClusterStatus string

const (
	ClusterStatusUnknown      = ClusterStatus("unknown")
	ClusterStatusReady        = ClusterStatus("ready")
	ClusterStatusProvisioning = ClusterStatus("provisioning")
	ClusterStatusConfiguring  = ClusterStatus("configuring")
	ClusterStatusDeleting     = ClusterStatus("deleting")
	ClusterStatusError        = ClusterStatus("error")
	ClusterStatusAutohealing  = ClusterStatus("autohealing")
	ClusterStatusLocked       = ClusterStatus("locked")
	ClusterStatusSuspended    = ClusterStatus("suspended")
	ClusterStatusInitializing = ClusterStatus("initializing")
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

type ListClustersRequestOrderBy string

const (
	ListClustersRequestOrderByCreatedAtAsc  = ListClustersRequestOrderBy("created_at_asc")
	ListClustersRequestOrderByCreatedAtDesc = ListClustersRequestOrderBy("created_at_desc")
	ListClustersRequestOrderByNameAsc       = ListClustersRequestOrderBy("name_asc")
	ListClustersRequestOrderByNameDesc      = ListClustersRequestOrderBy("name_desc")
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

type PrivateNetworkProvisioningMode string

const (
	PrivateNetworkProvisioningModeStatic = PrivateNetworkProvisioningMode("static")
	PrivateNetworkProvisioningModeIpam   = PrivateNetworkProvisioningMode("ipam")
)

func (enum PrivateNetworkProvisioningMode) String() string {
	if enum == "" {
		// return default value if empty
		return "static"
	}
	return string(enum)
}

func (enum PrivateNetworkProvisioningMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PrivateNetworkProvisioningMode) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PrivateNetworkProvisioningMode(PrivateNetworkProvisioningMode(tmp).String())
	return nil
}

// EndpointSpecPrivateNetworkSpecIpamConfig:
type EndpointSpecPrivateNetworkSpecIpamConfig struct {
}

// PrivateNetwork:
type PrivateNetwork struct {
	// ID: UUID of the Private Network.
	ID string `json:"id"`
	// ServiceIPs: List of IPv4 CIDR notation addresses of the endpoint.
	ServiceIPs []scw.IPNet `json:"service_ips"`
	// Zone: Zone of the Private Network.
	Zone scw.Zone `json:"zone"`
	// ProvisioningMode: How your endpoint ips are provisioned.
	ProvisioningMode PrivateNetworkProvisioningMode `json:"provisioning_mode"`
}

// PublicNetwork:
type PublicNetwork struct {
}

// EndpointSpecPrivateNetworkSpec:
type EndpointSpecPrivateNetworkSpec struct {
	// ID: UUID of the Private Network to connect to the Database Instance.
	ID string `json:"id"`
	// ServiceIPs: Endpoint IPv4 address with a CIDR notation. You must provide at least one IPv4 per node.
	ServiceIPs []scw.IPNet `json:"service_ips"`
	// IpamConfig: Automated configuration of your Private Network endpoint with Scaleway IPAM service.
	IpamConfig *EndpointSpecPrivateNetworkSpecIpamConfig `json:"ipam_config"`
}

// EndpointSpecPublicNetworkSpec:
type EndpointSpecPublicNetworkSpec struct {
}

// AvailableClusterSetting:
type AvailableClusterSetting struct {
	// Name: Name of the setting.
	Name string `json:"name"`
	// DefaultValue: Default value of the setting.
	DefaultValue *string `json:"default_value"`
	// Type: Type of setting.
	Type AvailableClusterSettingPropertyType `json:"type"`
	// Description: Description of the setting.
	Description string `json:"description"`
	// MaxValue: Optional maximum value of the setting.
	MaxValue *int64 `json:"max_value"`
	// MinValue: Optional minimum value of the setting.
	MinValue *int64 `json:"min_value"`
	// Regex: Optional validation rule of the setting.
	Regex *string `json:"regex"`
	// Deprecated: Defines whether or not the setting is deprecated.
	Deprecated bool `json:"deprecated"`
}

// ACLRule:
type ACLRule struct {
	// ID: ID of the rule.
	ID string `json:"id"`
	// IPCidr: IPv4 network address of the rule.
	IPCidr *scw.IPNet `json:"ip_cidr"`
	// Description: Description of the rule.
	Description *string `json:"description"`
}

// ClusterSetting:
type ClusterSetting struct {
	// Value: Value of the setting.
	Value string `json:"value"`
	// Name: Name of the setting.
	Name string `json:"name"`
}

// Endpoint:
type Endpoint struct {
	// Port: TCP port of the endpoint.
	Port uint32 `json:"port"`
	// PrivateNetwork: Private Network details.
	PrivateNetwork *PrivateNetwork `json:"private_network,omitempty"`
	// PublicNetwork: Public network details.
	PublicNetwork *PublicNetwork `json:"public_network,omitempty"`
	// IPs: List of IPv4 addresses of the endpoint.
	IPs []net.IP `json:"ips"`
	// ID: UUID of the endpoint.
	ID string `json:"id"`
}

// ACLRuleSpec:
type ACLRuleSpec struct {
	// IPCidr: IPv4 network address of the rule.
	IPCidr scw.IPNet `json:"ip_cidr"`
	// Description: Description of the rule.
	Description string `json:"description"`
}

// EndpointSpec:
type EndpointSpec struct {
	// PrivateNetwork: Private Network specification details.
	PrivateNetwork *EndpointSpecPrivateNetworkSpec `json:"private_network,omitempty"`
	// PublicNetwork: Public network specification details.
	PublicNetwork *EndpointSpecPublicNetworkSpec `json:"public_network,omitempty"`
}

// ClusterVersion:
type ClusterVersion struct {
	// Version: Redis™ engine version.
	Version string `json:"version"`
	// EndOfLifeAt: Date of End of Life.
	EndOfLifeAt *time.Time `json:"end_of_life_at"`
	// AvailableSettings: Cluster settings available to be updated.
	AvailableSettings []*AvailableClusterSetting `json:"available_settings"`
	// LogoURL: Redis™ logo url.
	LogoURL string `json:"logo_url"`
	// Zone: Zone of the Redis™ Database Instance.
	Zone scw.Zone `json:"zone"`
}

// Cluster:
type Cluster struct {
	// ID: UUID of the Database Instance.
	ID string `json:"id"`
	// Name: Name of the Database Instance.
	Name string `json:"name"`
	// ProjectID: Project ID the Database Instance belongs to.
	ProjectID string `json:"project_id"`
	// Status: Status of the Database Instance.
	Status ClusterStatus `json:"status"`
	// Version: Redis™ engine version of the Database Instance.
	Version string `json:"version"`
	// Endpoints: List of Database Instance endpoints.
	Endpoints []*Endpoint `json:"endpoints"`
	// Tags: List of tags applied to the Database Instance.
	Tags []string `json:"tags"`
	// NodeType: Node type of the Database Instance.
	NodeType string `json:"node_type"`
	// CreatedAt: Creation date (Format ISO 8601).
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Update date (Format ISO 8601).
	UpdatedAt *time.Time `json:"updated_at"`
	// TLSEnabled: Defines whether or not TLS is enabled.
	TLSEnabled bool `json:"tls_enabled"`
	// ClusterSettings: List of Database Instance settings.
	ClusterSettings []*ClusterSetting `json:"cluster_settings"`
	// ACLRules: List of ACL rules.
	ACLRules []*ACLRule `json:"acl_rules"`
	// ClusterSize: Number of nodes of the Database Instance cluster.
	ClusterSize uint32 `json:"cluster_size"`
	// Zone: Zone of the Database Instance.
	Zone scw.Zone `json:"zone"`
	// UserName: Name of the user associated to the cluster.
	UserName string `json:"user_name"`
	// UpgradableVersions: List of engine versions the Database Instance can upgrade to.
	UpgradableVersions []string `json:"upgradable_versions"`
}

// NodeType:
type NodeType struct {
	// Name: Node type name.
	Name string `json:"name"`
	// StockStatus: Current stock status of the node type.
	StockStatus NodeTypeStock `json:"stock_status"`
	// Description: Current specifications of the offer.
	Description string `json:"description"`
	// Vcpus: Number of virtual CPUs.
	Vcpus uint32 `json:"vcpus"`
	// Memory: Quantity of RAM.
	Memory scw.Size `json:"memory"`
	// Disabled: Defines whether node type is currently disabled or not.
	Disabled bool `json:"disabled"`
	// Beta: Defines whether node type is currently in beta.
	Beta bool `json:"beta"`
	// Zone: Zone of the node type.
	Zone scw.Zone `json:"zone"`
}

// AddACLRulesRequest:
type AddACLRulesRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance you want to add ACL rules to.
	ClusterID string `json:"-"`
	// ACLRules: ACLs rules to add to the cluster.
	ACLRules []*ACLRuleSpec `json:"acl_rules"`
}

// AddACLRulesResponse:
type AddACLRulesResponse struct {
	// ACLRules: ACL Rules enabled for the Database Instance.
	ACLRules []*ACLRule `json:"acl_rules"`
	// TotalCount: Total count of ACL rules of the Database Instance.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *AddACLRulesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *AddACLRulesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*AddACLRulesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ACLRules = append(r.ACLRules, results.ACLRules...)
	r.TotalCount += uint32(len(results.ACLRules))
	return uint32(len(results.ACLRules)), nil
}

// AddClusterSettingsRequest:
type AddClusterSettingsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance you want to add settings to.
	ClusterID string `json:"-"`
	// Settings: Settings to add to the cluster.
	Settings []*ClusterSetting `json:"settings"`
}

// AddEndpointsRequest:
type AddEndpointsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance you want to add endpoints to.
	ClusterID string `json:"-"`
	// Endpoints: Endpoints to add to the Database Instance.
	Endpoints []*EndpointSpec `json:"endpoints"`
}

// AddEndpointsResponse:
type AddEndpointsResponse struct {
	// Endpoints: Endpoints defined on the Database Instance.
	Endpoints []*Endpoint `json:"endpoints"`
	// TotalCount: Total count of endpoints of the Database Instance.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *AddEndpointsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *AddEndpointsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*AddEndpointsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Endpoints = append(r.Endpoints, results.Endpoints...)
	r.TotalCount += uint32(len(results.Endpoints))
	return uint32(len(results.Endpoints)), nil
}

// ClusterMetricsResponse:
type ClusterMetricsResponse struct {
	// Timeseries: Time series of metrics of a given cluster.
	Timeseries []*scw.TimeSeries `json:"timeseries"`
}

// ClusterSettingsResponse:
type ClusterSettingsResponse struct {
	// Settings: Settings configured for a given Database Instance.
	Settings []*ClusterSetting `json:"settings"`
}

// CreateClusterRequest:
type CreateClusterRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ProjectID: Project ID in which to create the Database Instance.
	ProjectID string `json:"project_id"`
	// Name: Name of the Database Instance.
	Name string `json:"name"`
	// Version: Redis™ engine version of the Database Instance.
	Version string `json:"version"`
	// Tags: Tags to apply to the Database Instance.
	Tags []string `json:"tags"`
	// NodeType: Type of node to use for the Database Instance.
	NodeType string `json:"node_type"`
	// UserName: Name of the user created upon Database Instance creation.
	UserName string `json:"user_name"`
	// Password: Password of the user.
	Password string `json:"password"`
	// ClusterSize: Number of nodes in the Redis™ cluster.
	ClusterSize *int32 `json:"cluster_size,omitempty"`
	// ACLRules: List of ACLRuleSpec used to secure your publicly exposed cluster.
	ACLRules []*ACLRuleSpec `json:"acl_rules"`
	// Endpoints: Zero or multiple EndpointSpec used to expose your cluster publicly and inside private networks. If no EndpoindSpec is given the cluster will be publicly exposed by default.
	Endpoints []*EndpointSpec `json:"endpoints"`
	// TLSEnabled: Defines whether or not TLS is enabled.
	TLSEnabled bool `json:"tls_enabled"`
	// ClusterSettings: List of advanced settings to be set upon Database Instance initialization.
	ClusterSettings []*ClusterSetting `json:"cluster_settings"`
}

// DeleteACLRuleRequest:
type DeleteACLRuleRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ACLID: UUID of the ACL rule you want to delete.
	ACLID string `json:"-"`
}

// DeleteClusterRequest:
type DeleteClusterRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance to delete.
	ClusterID string `json:"-"`
}

// DeleteClusterSettingRequest:
type DeleteClusterSettingRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance where the settings must be set.
	ClusterID string `json:"-"`
	// SettingName: Setting name to delete.
	SettingName string `json:"-"`
}

// DeleteEndpointRequest:
type DeleteEndpointRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// EndpointID: UUID of the endpoint you want to delete.
	EndpointID string `json:"-"`
}

// GetACLRuleRequest:
type GetACLRuleRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ACLID: UUID of the ACL rule you want to get.
	ACLID string `json:"-"`
}

// GetClusterCertificateRequest:
type GetClusterCertificateRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster.
	ClusterID string `json:"-"`
}

// GetClusterMetricsRequest:
type GetClusterMetricsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster.
	ClusterID string `json:"-"`
	// StartAt: Start date.
	StartAt *time.Time `json:"-"`
	// EndAt: End date.
	EndAt *time.Time `json:"-"`
	// MetricName: Name of the metric to gather.
	MetricName *string `json:"-"`
}

// GetClusterRequest:
type GetClusterRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster.
	ClusterID string `json:"-"`
}

// GetEndpointRequest:
type GetEndpointRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// EndpointID: UUID of the endpoint you want to get.
	EndpointID string `json:"-"`
}

// ListClusterVersionsRequest:
type ListClusterVersionsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// IncludeDisabled: Defines whether or not to include disabled Redis™ engine versions.
	IncludeDisabled bool `json:"-"`
	// IncludeBeta: Defines whether or not to include beta Redis™ engine versions.
	IncludeBeta bool `json:"-"`
	// IncludeDeprecated: Defines whether or not to include deprecated Redis™ engine versions.
	IncludeDeprecated bool `json:"-"`
	// Version: List Redis™ engine versions that match a given name pattern.
	Version *string `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
}

// ListClusterVersionsResponse:
type ListClusterVersionsResponse struct {
	// Versions: List of available Redis™ engine versions.
	Versions []*ClusterVersion `json:"versions"`
	// TotalCount: Total count of available Redis™ engine versions.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListClusterVersionsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListClusterVersionsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListClusterVersionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Versions = append(r.Versions, results.Versions...)
	r.TotalCount += uint32(len(results.Versions))
	return uint32(len(results.Versions)), nil
}

// ListClustersRequest:
type ListClustersRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Tags: Filter by Database Instance tags.
	Tags []string `json:"-"`
	// Name: Filter by Database Instance names.
	Name *string `json:"-"`
	// OrderBy: Criteria to use when ordering the list.
	OrderBy ListClustersRequestOrderBy `json:"-"`
	// ProjectID: Filter by Project ID.
	ProjectID *string `json:"-"`
	// OrganizationID: Filter by Organization ID.
	OrganizationID *string `json:"-"`
	// Version: Filter by Redis™ engine version.
	Version *string `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
}

// ListClustersResponse:
type ListClustersResponse struct {
	// Clusters: List all Database Instances.
	Clusters []*Cluster `json:"clusters"`
	// TotalCount: Total count of Database Instances.
	TotalCount uint32 `json:"total_count"`
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

// ListNodeTypesRequest:
type ListNodeTypesRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// IncludeDisabledTypes: Defines whether or not to include disabled types.
	IncludeDisabledTypes bool `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
}

// ListNodeTypesResponse:
type ListNodeTypesResponse struct {
	// NodeTypes: Types of node.
	NodeTypes []*NodeType `json:"node_types"`
	// TotalCount: Total count of node types available.
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

// MigrateClusterRequest:
type MigrateClusterRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance to update.
	ClusterID string `json:"-"`
	// Version: Redis™ engine version of the Database Instance.
	Version *string `json:"version,omitempty"`
	// NodeType: Type of node to use for the Database Instance.
	NodeType *string `json:"node_type,omitempty"`
	// ClusterSize: Number of nodes for the Database Instance.
	ClusterSize *uint32 `json:"cluster_size,omitempty"`
}

// RenewClusterCertificateRequest:
type RenewClusterCertificateRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster.
	ClusterID string `json:"-"`
}

// SetACLRulesRequest:
type SetACLRulesRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance where the ACL rules have to be set.
	ClusterID string `json:"-"`
	// ACLRules: ACLs rules to define for the cluster.
	ACLRules []*ACLRuleSpec `json:"acl_rules"`
}

// SetACLRulesResponse:
type SetACLRulesResponse struct {
	// ACLRules: ACL Rules enabled for the Database Instance.
	ACLRules []*ACLRule `json:"acl_rules"`
}

// SetClusterSettingsRequest:
type SetClusterSettingsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance where the settings must be set.
	ClusterID string `json:"-"`
	// Settings: Settings to define for the Database Instance.
	Settings []*ClusterSetting `json:"settings"`
}

// SetEndpointsRequest:
type SetEndpointsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance where the endpoints have to be set.
	ClusterID string `json:"-"`
	// Endpoints: Endpoints to define for the Database Instance.
	Endpoints []*EndpointSpec `json:"endpoints"`
}

// SetEndpointsResponse:
type SetEndpointsResponse struct {
	// Endpoints: Endpoints defined on the Database Instance.
	Endpoints []*Endpoint `json:"endpoints"`
}

// UpdateClusterRequest:
type UpdateClusterRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance to update.
	ClusterID string `json:"-"`
	// Name: Name of the Database Instance.
	Name *string `json:"name,omitempty"`
	// Tags: Database Instance tags.
	Tags *[]string `json:"tags,omitempty"`
	// UserName: Name of the Database Instance user.
	UserName *string `json:"user_name,omitempty"`
	// Password: Password of the Database Instance user.
	Password *string `json:"password,omitempty"`
}

// UpdateEndpointRequest:
type UpdateEndpointRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// EndpointID:
	EndpointID string `json:"-"`
	// PrivateNetwork:
	PrivateNetwork *EndpointSpecPrivateNetworkSpec `json:"private_network,omitempty"`
	// PublicNetwork:
	PublicNetwork *EndpointSpecPublicNetworkSpec `json:"public_network,omitempty"`
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
func (s *API) Zones() []scw.Zone {
	return []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw1, scw.ZonePlWaw2}
}

// CreateCluster: Create a new Redis™ Database Instance (Redis™ cluster). You must set the `zone`, `project_id`, `version`, `node_type`, `user_name` and `password` parameters. Optionally you can define `acl_rules`, `endpoints`, `tls_enabled` and `cluster_settings`.
func (s *API) CreateCluster(req *CreateClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("ins")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters",
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

// UpdateCluster: Update the parameters of a Redis™ Database Instance (Redis™ cluster), including `name`, `tags`, `user_name` and `password`.
func (s *API) UpdateCluster(req *UpdateClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
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

// GetCluster: Retrieve information about a Redis™ Database Instance (Redis™ cluster). Specify the `cluster_id` and `region` in your request to get information such as `id`, `status`, `version`, `tls_enabled`, `cluster_settings`, `upgradable_versions` and `endpoints` about your cluster in the response.
func (s *API) GetCluster(req *GetClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListClusters: List all Redis™ Database Instances (Redis™ cluster) in the specified zone. By default, the Database Instances returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field. You can define additional parameters for your query, such as `tags`, `name`, `organization_id` and `version`.
func (s *API) ListClusters(req *ListClustersRequest, opts ...scw.RequestOption) (*ListClustersResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}
	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "version", req.Version)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters",
		Query:  query,
	}

	var resp ListClustersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MigrateCluster: Upgrade your standalone Redis™ Database Instance node, either by upgrading to a bigger node type (vertical scaling) or by adding more nodes to your Database Instance to increase your number of endpoints and distribute cache (horizontal scaling). Note that scaling horizontally your Redis™ Database Instance will not renew its TLS certificate. In order to refresh the TLS certificate, you must use the Renew TLS certificate endpoint.
func (s *API) MigrateCluster(req *MigrateClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/migrate",
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

// DeleteCluster: Delete a Redis™ Database Instance (Redis™ cluster), specified by the `region` and `cluster_id` parameters. Deleting a Database Instance is permanent, and cannot be undone. Note that upon deletion all your data will be lost.
func (s *API) DeleteCluster(req *DeleteClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetClusterMetrics: Retrieve the metrics of a Redis™ Database Instance (Redis™ cluster). You can define the period from which to retrieve metrics by specifying the `start_date` and `end_date`.
func (s *API) GetClusterMetrics(req *GetClusterMetricsRequest, opts ...scw.RequestOption) (*ClusterMetricsResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "start_at", req.StartAt)
	parameter.AddToQuery(query, "end_at", req.EndAt)
	parameter.AddToQuery(query, "metric_name", req.MetricName)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/metrics",
		Query:  query,
	}

	var resp ClusterMetricsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNodeTypes: List all available node types. By default, the node types returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListNodeTypes(req *ListNodeTypesRequest, opts ...scw.RequestOption) (*ListNodeTypesResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}
	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "include_disabled_types", req.IncludeDisabledTypes)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/node-types",
		Query:  query,
	}

	var resp ListNodeTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListClusterVersions: List the Redis™ database engine versions available. You can define additional parameters for your query, such as `include_disabled`, `include_beta`, `include_deprecated` and `version`.
func (s *API) ListClusterVersions(req *ListClusterVersionsRequest, opts ...scw.RequestOption) (*ListClusterVersionsResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}
	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "include_disabled", req.IncludeDisabled)
	parameter.AddToQuery(query, "include_beta", req.IncludeBeta)
	parameter.AddToQuery(query, "include_deprecated", req.IncludeDeprecated)
	parameter.AddToQuery(query, "version", req.Version)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/cluster-versions",
		Query:  query,
	}

	var resp ListClusterVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetClusterCertificate: Retrieve information about the TLS certificate of a Redis™ Database Instance (Redis™ cluster). Details like name and content are returned in the response.
func (s *API) GetClusterCertificate(req *GetClusterCertificateRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/certificate",
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RenewClusterCertificate: Renew a TLS certificate for a Redis™ Database Instance (Redis™ cluster). Renewing a certificate means that you will not be able to connect to your Database Instance using the previous certificate. You will also need to download and update the new certificate for all database clients.
func (s *API) RenewClusterCertificate(req *RenewClusterCertificateRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/renew-certificate",
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

// AddClusterSettings: Add an advanced setting to a Redis™ Database Instance (Redis™ cluster). You must set the `name` and the `value` of each setting.
func (s *API) AddClusterSettings(req *AddClusterSettingsRequest, opts ...scw.RequestOption) (*ClusterSettingsResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/settings",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ClusterSettingsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteClusterSetting: Delete an advanced setting in a Redis™ Database Instance (Redis™ cluster). You must specify the names of the settings you want to delete in the request body.
func (s *API) DeleteClusterSetting(req *DeleteClusterSettingRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	if fmt.Sprint(req.SettingName) == "" {
		return nil, errors.New("field SettingName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/settings/" + fmt.Sprint(req.SettingName) + "",
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetClusterSettings: Update an advanced setting for a Redis™ Database Instance (Redis™ cluster). Settings added upon database engine initalization can only be defined once, and cannot, therefore, be updated.
func (s *API) SetClusterSettings(req *SetClusterSettingsRequest, opts ...scw.RequestOption) (*ClusterSettingsResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/settings",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ClusterSettingsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetACLRules: Replace all the ACL rules of a Redis™ Database Instance (Redis™ cluster).
func (s *API) SetACLRules(req *SetACLRulesRequest, opts ...scw.RequestOption) (*SetACLRulesResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/acls",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetACLRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddACLRules: Add an additional ACL rule to a Redis™ Database Instance (Redis™ cluster).
func (s *API) AddACLRules(req *AddACLRulesRequest, opts ...scw.RequestOption) (*AddACLRulesResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/acls",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AddACLRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteACLRule: Delete an ACL rule of a Redis™ Database Instance (Redis™ cluster). You must specify the `acl_id` of the rule you want to delete in your request.
func (s *API) DeleteACLRule(req *DeleteACLRuleRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ACLID) == "" {
		return nil, errors.New("field ACLID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/acls/" + fmt.Sprint(req.ACLID) + "",
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetACLRule: Retrieve information about an ACL rule of a Redis™ Database Instance (Redis™ cluster). You must specify the `acl_id` of the rule in your request.
func (s *API) GetACLRule(req *GetACLRuleRequest, opts ...scw.RequestOption) (*ACLRule, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ACLID) == "" {
		return nil, errors.New("field ACLID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/acls/" + fmt.Sprint(req.ACLID) + "",
	}

	var resp ACLRule

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetEndpoints: Update an endpoint for a Redis™ Database Instance (Redis™ cluster). You must specify the `cluster_id` and the `endpoints` parameters in your request.
func (s *API) SetEndpoints(req *SetEndpointsRequest, opts ...scw.RequestOption) (*SetEndpointsResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/endpoints",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetEndpointsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddEndpoints: Add a new endpoint for a Redis™ Database Instance (Redis™ cluster). You can add `private_network` or `public_network` specifications to the body of the request.
func (s *API) AddEndpoints(req *AddEndpointsRequest, opts ...scw.RequestOption) (*AddEndpointsResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/endpoints",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AddEndpointsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteEndpoint: Delete the endpoint of a Redis™ Database Instance (Redis™ cluster). You must specify the `region` and `endpoint_id` parameters of the endpoint you want to delete. Note that might need to update any environment configurations that point to the deleted endpoint.
func (s *API) DeleteEndpoint(req *DeleteEndpointRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.EndpointID) == "" {
		return nil, errors.New("field EndpointID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetEndpoint: Retrieve information about a Redis™ Database Instance (Redis™ cluster) endpoint. Full details about the endpoint, like `ips`, `port`, `private_network` and `public_network` specifications are returned in the response.
func (s *API) GetEndpoint(req *GetEndpointRequest, opts ...scw.RequestOption) (*Endpoint, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.EndpointID) == "" {
		return nil, errors.New("field EndpointID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
	}

	var resp Endpoint

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateEndpoint: Update information about a Redis™ Database Instance (Redis™ cluster) endpoint. Full details about the endpoint, like `ips`, `port`, `private_network` and `public_network` specifications are returned in the response.
func (s *API) UpdateEndpoint(req *UpdateEndpointRequest, opts ...scw.RequestOption) (*Endpoint, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.EndpointID) == "" {
		return nil, errors.New("field EndpointID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/redis/v1/zones/" + fmt.Sprint(req.Zone) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
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
