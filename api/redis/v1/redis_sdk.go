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

type APIListClustersRequestOrderBy string

const (
	APIListClustersRequestOrderByCreatedAtAsc  = APIListClustersRequestOrderBy("created_at_asc")
	APIListClustersRequestOrderByCreatedAtDesc = APIListClustersRequestOrderBy("created_at_desc")
	APIListClustersRequestOrderByNameAsc       = APIListClustersRequestOrderBy("name_asc")
	APIListClustersRequestOrderByNameDesc      = APIListClustersRequestOrderBy("name_desc")
)

func (enum APIListClustersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum APIListClustersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *APIListClustersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = APIListClustersRequestOrderBy(APIListClustersRequestOrderBy(tmp).String())
	return nil
}

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

// APIAddACLRulesRequest:
type APIAddACLRulesRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance you want to add ACL rules to.
	ClusterID string `json:"-"`
	// ACLRules: ACLs rules to add to the cluster.
	ACLRules []*ACLRuleSpec `json:"acl_rules"`
}

// APIAddClusterSettingsRequest:
type APIAddClusterSettingsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance you want to add settings to.
	ClusterID string `json:"-"`
	// Settings: Settings to add to the cluster.
	Settings []*ClusterSetting `json:"settings"`
}

// APIAddEndpointsRequest:
type APIAddEndpointsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance you want to add endpoints to.
	ClusterID string `json:"-"`
	// Endpoints: Endpoints to add to the Database Instance.
	Endpoints []*EndpointSpec `json:"endpoints"`
}

// APICreateClusterRequest:
type APICreateClusterRequest struct {
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

// APIDeleteACLRuleRequest:
type APIDeleteACLRuleRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ACLID: UUID of the ACL rule you want to delete.
	ACLID string `json:"-"`
}

// APIDeleteClusterRequest:
type APIDeleteClusterRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance to delete.
	ClusterID string `json:"-"`
}

// APIDeleteClusterSettingRequest:
type APIDeleteClusterSettingRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance where the settings must be set.
	ClusterID string `json:"-"`
	// SettingName: Setting name to delete.
	SettingName string `json:"-"`
}

// APIDeleteEndpointRequest:
type APIDeleteEndpointRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// EndpointID: UUID of the endpoint you want to delete.
	EndpointID string `json:"-"`
}

// APIGetACLRuleRequest:
type APIGetACLRuleRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ACLID: UUID of the ACL rule you want to get.
	ACLID string `json:"-"`
}

// APIGetClusterCertificateRequest:
type APIGetClusterCertificateRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster.
	ClusterID string `json:"-"`
}

// APIGetClusterMetricsRequest:
type APIGetClusterMetricsRequest struct {
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

// APIGetClusterRequest:
type APIGetClusterRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster.
	ClusterID string `json:"-"`
}

// APIGetEndpointRequest:
type APIGetEndpointRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// EndpointID: UUID of the endpoint you want to get.
	EndpointID string `json:"-"`
}

// APIListClusterVersionsRequest:
type APIListClusterVersionsRequest struct {
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

// APIListClustersRequest:
type APIListClustersRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Tags: Filter by Database Instance tags.
	Tags []string `json:"-"`
	// Name: Filter by Database Instance names.
	Name *string `json:"-"`
	// OrderBy: Criteria to use when ordering the list.
	OrderBy APIListClustersRequestOrderBy `json:"-"`
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

// APIListNodeTypesRequest:
type APIListNodeTypesRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// IncludeDisabledTypes: Defines whether or not to include disabled types.
	IncludeDisabledTypes bool `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
}

// APIMigrateClusterRequest:
type APIMigrateClusterRequest struct {
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

// APIRenewClusterCertificateRequest:
type APIRenewClusterCertificateRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster.
	ClusterID string `json:"-"`
}

// APISetACLRulesRequest:
type APISetACLRulesRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance where the ACL rules have to be set.
	ClusterID string `json:"-"`
	// ACLRules: ACLs rules to define for the cluster.
	ACLRules []*ACLRuleSpec `json:"acl_rules"`
}

// APISetClusterSettingsRequest:
type APISetClusterSettingsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance where the settings must be set.
	ClusterID string `json:"-"`
	// Settings: Settings to define for the Database Instance.
	Settings []*ClusterSetting `json:"settings"`
}

// APISetEndpointsRequest:
type APISetEndpointsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the Database Instance where the endpoints have to be set.
	ClusterID string `json:"-"`
	// Endpoints: Endpoints to define for the Database Instance.
	Endpoints []*EndpointSpec `json:"endpoints"`
}

// APIUpdateClusterRequest:
type APIUpdateClusterRequest struct {
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

// APIUpdateEndpointRequest:
type APIUpdateEndpointRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// EndpointID:
	EndpointID string `json:"-"`
	// PrivateNetwork:
	PrivateNetwork *EndpointSpecPrivateNetworkSpec `json:"private_network,omitempty"`
	// PublicNetwork:
	PublicNetwork *EndpointSpecPublicNetworkSpec `json:"public_network,omitempty"`
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

// SetACLRulesResponse:
type SetACLRulesResponse struct {
	// ACLRules: ACL Rules enabled for the Database Instance.
	ACLRules []*ACLRule `json:"acl_rules"`
}

// SetEndpointsResponse:
type SetEndpointsResponse struct {
	// Endpoints: Endpoints defined on the Database Instance.
	Endpoints []*Endpoint `json:"endpoints"`
}

// Managed Database for Redis™ is a low-latency caching solution. It allows you to easily set up a secure cache and lighten the load on your main database. Based on the in-memory data storage, Managed Database for Redis™ improves your application response time and helps you provide a better experience to your users.
//
// Using Managed Database for Redis™ as a cache optimizes the speed of your requests as copies of the most frequently used data are stored in memory, making them accessible in milliseconds.
//
// (switchcolumn)
// (switchcolumn)
//
// ## Concepts
//
// Refer to our [dedicated concepts page](https://www.scaleway.com/en/docs/managed-databases/redis/concepts/) to find definitions of the different terms referring to Managed Database for Redis™.
//
// (switchcolumn)
// (switchcolumn)
//
// ## Quickstart
//
//  1. Configure your environment variables.
//     <Message type="note">
//     This is an optional step that seeks to simplify your usage of the APIs
//     </Message>
//
//     ```bash
//     export ACCESS_KEY="<access-key>"
//     export SECRET_KEY="<secret-key>"
//     export SCW_ZONE="<zone>"
//     ```
//
//  2. Edit the POST request payload you will use to create your Redis™ Database Instance cluster. Replace the parameters in the following example:
//     ```json
//     {
//     "project_id":"50e8d5d3-c623-4df8-a5ef-1972a5432001",
//     "name":"cluster1",
//     "version":"7.0.5",
//     "tags":[
//     "tag1"
//     ],
//     "node_type":"RED1-micro",
//     "user_name":"redis-user",
//     "password":"lxbemloZiPT1l37*",
//     "cluster_size": 3,
//     "endpoints": [
//     {
//     "private_network": {
//     "id": "d8e65f2b-cce9-40b7-80fc-6a2902db6826",
//     "service_ips": [
//     "192.0.2.1/24",
//     "192.0.2.2/24",
//     "192.0.2.3/24"
//     ]
//     }
//     }
//     ],
//     "tls_enabled": true
//     }
//     ```
//
//     | Parameter        | Description                                                        |
//     | :--------------- | :----------------------------------------------------------------- |
//     | `project_id`     | **REQUIRED** The ID of the Project you want to create your Database Instance in. To find your Project ID you can **[list the projects](/api/account#path-projects-list-all-projects-of-an-organization)** or consult the **[Scaleway console](https://console.scaleway.com/project/settings)**. |
//     | `name`           | Name of the Redis™ Database Instance                                          |
//     | `version`        | **REQUIRED** Version of the Redis™ engine. To check the list of available versions you can use the folowing endpoint: `https://api.scaleway.com/redis/v1/zones/$SCW_ZONE/cluster-versions`                                           |
//     | `tags`           | The list of tags `["tag1", "tag2", ...]` that will be associated with the Redis™ Database Instance. Tags can be appended to the query of the **[List Database Instances](#path-redistm-database-instance-list-redistm-database-instances)** call to show a list of the Database Instances using a specific tag. You can also combine tags to list Database Instances that possess **all** of the appended tags. |
//     | `node_type`      | **REQUIRED** The node type. To check the list of available node types you can use the folowing endpoint: `https://api.scaleway.com/redis/v1/zones/$SCW_ZONE/node-types` |
//     | `user_name`      | **REQUIRED** Identifier of the default user, which is created concurrently with the Redis™ Database Instance |
//     | `password`       | **REQUIRED** Password for the default user |
//     | `cluster_size`   | **INTEGER** The number of nodes in the Redis™ Database Instance cluster. You can either set it to 1 for a standalone Database Instance, or set it to 3 to 6, for Database Instances in cluster mode |
//     | `endpoints`      | The network configuration of your Redis™ Database Instance. You can either set `private_network` or `public_network`. If you set up a Private Network, you must define the ID of the Private Network connected to the cluster, and one IPv4 address endpoint per node. The address must be in the [RFC1918 subnet](https://datatracker.ietf.org/doc/html/rfc1918) range and follow the [CIDR notation method](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing). |
//     | `tls_enabled`    | **BOOLEAN** Whether or not to enable TLS certificates |
//
//  3. Create a Redis™ Database Instance by running the following command. Make sure you include the payload you edited in the previous step.
//     ```bash
//     curl -X POST \
//     -H "Content-Type: application/json" \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" https://api.scaleway.com/redis/v1/zones/$SCW_REGION/clusters \
//     -d  '{
//     "project_id":"50e8d5d3-c623-4df8-a5ef-1972a5432001",
//     "name":"cluster1",
//     "version":"7.0.5",
//     "tags":[
//     "tag1"
//     ],
//     "node_type":"RED1-micro",
//     "user_name":"redis-user",
//     "password":"lxbemloZiPT1l37*",
//     "cluster_size": 3,
//     "endpoints": [
//     {
//     "private_network": {
//     "id": "d8e65f2b-cce9-40b7-80fc-6a2902db6826",
//     "service_ips": [
//     "192.0.2.1/24",
//     "192.0.2.2/24",
//     "192.0.2.3/24"
//     ]
//     }
//     }
//     ],
//     "tls_enabled": true
//     }'
//     ```
//
//  4. List your Redis™ Database Instances.
//     ```bash
//     curl -X GET
//     -H "Content-Type: application/json" \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" https://api.scaleway.com/redis/v1/zones/$SCW_ZONE/clusters
//     ```
//
//  5. Retrieve your Redis™ Database Instance IP and port from the response.
//     <Message type="note">
//     In this tutorial, we will use `192.0.2.1` and `6379` as the IP and port, respectively
//     </Message>
//
//  6. Connect to your Database Instance with the Redis™ client.
//     <Message type="note">
//     You can use only one of your node IP addresses at a time to connect to your Redis™ Database Instance, as the Redis™ CLI does not support cluster mode.
//     </Message>
//
//     ```bash
//     redis-cli -h 192.0.2.1 -p 6379 --user redis-user --askpass --tls --cacert SSL_redis-cluster1.pem
//     ```
//
//     <Message type="note">
//     The command above uses TLS to add an extra layer of security to your connection. The TLS certificate is generated automatically if you set tls_enabled to true. The certificates take on the following name structure: `SSL_redis-<name-of-your-redis-database-instance>.pem`. When using connectors other than redis-cli, you might need to specify the path to your certificate.
//     </Message>
//
//     <Message type="important">
//     Scaleway supports TLS1.2 and TLS1.3. If you use older versions of `libssl`, you might encounter connexion issues when using `redis-cli`. If this is the case, we recommend you check the libssl versions installed on your local machine and update if necessary.
//     </Message>
//
// 7. Enter the database password that you defined upon creation.
//
// You are now connected to your Managed Database for Redis™.
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
// ## Going Further
//
// For more information about Managed Database for Redis™, you can check out the following pages:
//
// * [Managed Database for Redis™ Documentation](https://www.scaleway.com/en/docs/managed-databases/redis/quickstart/)
// * [Managed Database for Redis™ FAQ](https://www.scaleway.com/en/docs/faq/databases-for-redis/)
// * [Scaleway Slack Community](https://scaleway-community.slack.com/) join the #database channel
// * [Contact our support team](https://console.scaleway.com/support/tickets/).
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
func (s *API) CreateCluster(req *APICreateClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
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
func (s *API) UpdateCluster(req *APIUpdateClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
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
func (s *API) GetCluster(req *APIGetClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
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
func (s *API) ListClusters(req *APIListClustersRequest, opts ...scw.RequestOption) (*ListClustersResponse, error) {
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
func (s *API) MigrateCluster(req *APIMigrateClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
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
func (s *API) DeleteCluster(req *APIDeleteClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
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
func (s *API) GetClusterMetrics(req *APIGetClusterMetricsRequest, opts ...scw.RequestOption) (*ClusterMetricsResponse, error) {
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
func (s *API) ListNodeTypes(req *APIListNodeTypesRequest, opts ...scw.RequestOption) (*ListNodeTypesResponse, error) {
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
func (s *API) ListClusterVersions(req *APIListClusterVersionsRequest, opts ...scw.RequestOption) (*ListClusterVersionsResponse, error) {
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
func (s *API) GetClusterCertificate(req *APIGetClusterCertificateRequest, opts ...scw.RequestOption) (*scw.File, error) {
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
func (s *API) RenewClusterCertificate(req *APIRenewClusterCertificateRequest, opts ...scw.RequestOption) (*Cluster, error) {
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
func (s *API) AddClusterSettings(req *APIAddClusterSettingsRequest, opts ...scw.RequestOption) (*ClusterSettingsResponse, error) {
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
func (s *API) DeleteClusterSetting(req *APIDeleteClusterSettingRequest, opts ...scw.RequestOption) (*Cluster, error) {
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
func (s *API) SetClusterSettings(req *APISetClusterSettingsRequest, opts ...scw.RequestOption) (*ClusterSettingsResponse, error) {
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
func (s *API) SetACLRules(req *APISetACLRulesRequest, opts ...scw.RequestOption) (*SetACLRulesResponse, error) {
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
func (s *API) AddACLRules(req *APIAddACLRulesRequest, opts ...scw.RequestOption) (*AddACLRulesResponse, error) {
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
func (s *API) DeleteACLRule(req *APIDeleteACLRuleRequest, opts ...scw.RequestOption) (*Cluster, error) {
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
func (s *API) GetACLRule(req *APIGetACLRuleRequest, opts ...scw.RequestOption) (*ACLRule, error) {
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
func (s *API) SetEndpoints(req *APISetEndpointsRequest, opts ...scw.RequestOption) (*SetEndpointsResponse, error) {
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
func (s *API) AddEndpoints(req *APIAddEndpointsRequest, opts ...scw.RequestOption) (*AddEndpointsResponse, error) {
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
func (s *API) DeleteEndpoint(req *APIDeleteEndpointRequest, opts ...scw.RequestOption) (*Cluster, error) {
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
func (s *API) GetEndpoint(req *APIGetEndpointRequest, opts ...scw.RequestOption) (*Endpoint, error) {
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
func (s *API) UpdateEndpoint(req *APIUpdateEndpointRequest, opts ...scw.RequestOption) (*Endpoint, error) {
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
