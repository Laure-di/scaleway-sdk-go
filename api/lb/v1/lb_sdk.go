// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package lb provides methods and message types of the lb v1 API.
package lb

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

type ACLActionRedirectRedirectType string

const (
	ACLActionRedirectRedirectTypeLocation = ACLActionRedirectRedirectType("location")
	ACLActionRedirectRedirectTypeScheme   = ACLActionRedirectRedirectType("scheme")
)

func (enum ACLActionRedirectRedirectType) String() string {
	if enum == "" {
		// return default value if empty
		return "location"
	}
	return string(enum)
}

func (enum ACLActionRedirectRedirectType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ACLActionRedirectRedirectType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ACLActionRedirectRedirectType(ACLActionRedirectRedirectType(tmp).String())
	return nil
}

type ACLActionType string

const (
	ACLActionTypeAllow    = ACLActionType("allow")
	ACLActionTypeDeny     = ACLActionType("deny")
	ACLActionTypeRedirect = ACLActionType("redirect")
)

func (enum ACLActionType) String() string {
	if enum == "" {
		// return default value if empty
		return "allow"
	}
	return string(enum)
}

func (enum ACLActionType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ACLActionType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ACLActionType(ACLActionType(tmp).String())
	return nil
}

type ACLHTTPFilter string

const (
	ACLHTTPFilterACLHTTPFilterNone = ACLHTTPFilter("acl_http_filter_none")
	ACLHTTPFilterPathBegin         = ACLHTTPFilter("path_begin")
	ACLHTTPFilterPathEnd           = ACLHTTPFilter("path_end")
	ACLHTTPFilterRegex             = ACLHTTPFilter("regex")
	ACLHTTPFilterHTTPHeaderMatch   = ACLHTTPFilter("http_header_match")
)

func (enum ACLHTTPFilter) String() string {
	if enum == "" {
		// return default value if empty
		return "acl_http_filter_none"
	}
	return string(enum)
}

func (enum ACLHTTPFilter) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ACLHTTPFilter) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ACLHTTPFilter(ACLHTTPFilter(tmp).String())
	return nil
}

type BackendServerStatsHealthCheckStatus string

const (
	BackendServerStatsHealthCheckStatusUnknown  = BackendServerStatsHealthCheckStatus("unknown")
	BackendServerStatsHealthCheckStatusNeutral  = BackendServerStatsHealthCheckStatus("neutral")
	BackendServerStatsHealthCheckStatusFailed   = BackendServerStatsHealthCheckStatus("failed")
	BackendServerStatsHealthCheckStatusPassed   = BackendServerStatsHealthCheckStatus("passed")
	BackendServerStatsHealthCheckStatusCondpass = BackendServerStatsHealthCheckStatus("condpass")
)

func (enum BackendServerStatsHealthCheckStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum BackendServerStatsHealthCheckStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *BackendServerStatsHealthCheckStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = BackendServerStatsHealthCheckStatus(BackendServerStatsHealthCheckStatus(tmp).String())
	return nil
}

type BackendServerStatsServerState string

const (
	BackendServerStatsServerStateStopped  = BackendServerStatsServerState("stopped")
	BackendServerStatsServerStateStarting = BackendServerStatsServerState("starting")
	BackendServerStatsServerStateRunning  = BackendServerStatsServerState("running")
	BackendServerStatsServerStateStopping = BackendServerStatsServerState("stopping")
)

func (enum BackendServerStatsServerState) String() string {
	if enum == "" {
		// return default value if empty
		return "stopped"
	}
	return string(enum)
}

func (enum BackendServerStatsServerState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *BackendServerStatsServerState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = BackendServerStatsServerState(BackendServerStatsServerState(tmp).String())
	return nil
}

type CertificateStatus string

const (
	CertificateStatusPending = CertificateStatus("pending")
	CertificateStatusReady   = CertificateStatus("ready")
	CertificateStatusError   = CertificateStatus("error")
)

func (enum CertificateStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "pending"
	}
	return string(enum)
}

func (enum CertificateStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CertificateStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CertificateStatus(CertificateStatus(tmp).String())
	return nil
}

type CertificateType string

const (
	CertificateTypeLetsencryt = CertificateType("letsencryt")
	CertificateTypeCustom     = CertificateType("custom")
)

func (enum CertificateType) String() string {
	if enum == "" {
		// return default value if empty
		return "letsencryt"
	}
	return string(enum)
}

func (enum CertificateType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CertificateType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CertificateType(CertificateType(tmp).String())
	return nil
}

type ForwardPortAlgorithm string

const (
	ForwardPortAlgorithmRoundrobin = ForwardPortAlgorithm("roundrobin")
	ForwardPortAlgorithmLeastconn  = ForwardPortAlgorithm("leastconn")
	ForwardPortAlgorithmFirst      = ForwardPortAlgorithm("first")
)

func (enum ForwardPortAlgorithm) String() string {
	if enum == "" {
		// return default value if empty
		return "roundrobin"
	}
	return string(enum)
}

func (enum ForwardPortAlgorithm) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ForwardPortAlgorithm) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ForwardPortAlgorithm(ForwardPortAlgorithm(tmp).String())
	return nil
}

type InstanceStatus string

const (
	InstanceStatusUnknown   = InstanceStatus("unknown")
	InstanceStatusReady     = InstanceStatus("ready")
	InstanceStatusPending   = InstanceStatus("pending")
	InstanceStatusStopped   = InstanceStatus("stopped")
	InstanceStatusError     = InstanceStatus("error")
	InstanceStatusLocked    = InstanceStatus("locked")
	InstanceStatusMigrating = InstanceStatus("migrating")
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

type LBStatus string

const (
	LBStatusUnknown   = LBStatus("unknown")
	LBStatusReady     = LBStatus("ready")
	LBStatusPending   = LBStatus("pending")
	LBStatusStopped   = LBStatus("stopped")
	LBStatusError     = LBStatus("error")
	LBStatusLocked    = LBStatus("locked")
	LBStatusMigrating = LBStatus("migrating")
	LBStatusToCreate  = LBStatus("to_create")
	LBStatusCreating  = LBStatus("creating")
	LBStatusToDelete  = LBStatus("to_delete")
	LBStatusDeleting  = LBStatus("deleting")
)

func (enum LBStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum LBStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *LBStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = LBStatus(LBStatus(tmp).String())
	return nil
}

type LBTypeStock string

const (
	LBTypeStockUnknown    = LBTypeStock("unknown")
	LBTypeStockLowStock   = LBTypeStock("low_stock")
	LBTypeStockOutOfStock = LBTypeStock("out_of_stock")
	LBTypeStockAvailable  = LBTypeStock("available")
)

func (enum LBTypeStock) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum LBTypeStock) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *LBTypeStock) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = LBTypeStock(LBTypeStock(tmp).String())
	return nil
}

type ListACLRequestOrderBy string

const (
	ListACLRequestOrderByCreatedAtAsc  = ListACLRequestOrderBy("created_at_asc")
	ListACLRequestOrderByCreatedAtDesc = ListACLRequestOrderBy("created_at_desc")
	ListACLRequestOrderByNameAsc       = ListACLRequestOrderBy("name_asc")
	ListACLRequestOrderByNameDesc      = ListACLRequestOrderBy("name_desc")
)

func (enum ListACLRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListACLRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListACLRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListACLRequestOrderBy(ListACLRequestOrderBy(tmp).String())
	return nil
}

type ListBackendsRequestOrderBy string

const (
	ListBackendsRequestOrderByCreatedAtAsc  = ListBackendsRequestOrderBy("created_at_asc")
	ListBackendsRequestOrderByCreatedAtDesc = ListBackendsRequestOrderBy("created_at_desc")
	ListBackendsRequestOrderByNameAsc       = ListBackendsRequestOrderBy("name_asc")
	ListBackendsRequestOrderByNameDesc      = ListBackendsRequestOrderBy("name_desc")
)

func (enum ListBackendsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListBackendsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListBackendsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListBackendsRequestOrderBy(ListBackendsRequestOrderBy(tmp).String())
	return nil
}

type ListCertificatesRequestOrderBy string

const (
	ListCertificatesRequestOrderByCreatedAtAsc  = ListCertificatesRequestOrderBy("created_at_asc")
	ListCertificatesRequestOrderByCreatedAtDesc = ListCertificatesRequestOrderBy("created_at_desc")
	ListCertificatesRequestOrderByNameAsc       = ListCertificatesRequestOrderBy("name_asc")
	ListCertificatesRequestOrderByNameDesc      = ListCertificatesRequestOrderBy("name_desc")
)

func (enum ListCertificatesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListCertificatesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListCertificatesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListCertificatesRequestOrderBy(ListCertificatesRequestOrderBy(tmp).String())
	return nil
}

type ListFrontendsRequestOrderBy string

const (
	ListFrontendsRequestOrderByCreatedAtAsc  = ListFrontendsRequestOrderBy("created_at_asc")
	ListFrontendsRequestOrderByCreatedAtDesc = ListFrontendsRequestOrderBy("created_at_desc")
	ListFrontendsRequestOrderByNameAsc       = ListFrontendsRequestOrderBy("name_asc")
	ListFrontendsRequestOrderByNameDesc      = ListFrontendsRequestOrderBy("name_desc")
)

func (enum ListFrontendsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListFrontendsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListFrontendsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListFrontendsRequestOrderBy(ListFrontendsRequestOrderBy(tmp).String())
	return nil
}

type ListLBsRequestOrderBy string

const (
	ListLBsRequestOrderByCreatedAtAsc  = ListLBsRequestOrderBy("created_at_asc")
	ListLBsRequestOrderByCreatedAtDesc = ListLBsRequestOrderBy("created_at_desc")
	ListLBsRequestOrderByNameAsc       = ListLBsRequestOrderBy("name_asc")
	ListLBsRequestOrderByNameDesc      = ListLBsRequestOrderBy("name_desc")
)

func (enum ListLBsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListLBsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListLBsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListLBsRequestOrderBy(ListLBsRequestOrderBy(tmp).String())
	return nil
}

type ListPrivateNetworksRequestOrderBy string

const (
	ListPrivateNetworksRequestOrderByCreatedAtAsc  = ListPrivateNetworksRequestOrderBy("created_at_asc")
	ListPrivateNetworksRequestOrderByCreatedAtDesc = ListPrivateNetworksRequestOrderBy("created_at_desc")
)

func (enum ListPrivateNetworksRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListPrivateNetworksRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPrivateNetworksRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPrivateNetworksRequestOrderBy(ListPrivateNetworksRequestOrderBy(tmp).String())
	return nil
}

type ListRoutesRequestOrderBy string

const (
	ListRoutesRequestOrderByCreatedAtAsc  = ListRoutesRequestOrderBy("created_at_asc")
	ListRoutesRequestOrderByCreatedAtDesc = ListRoutesRequestOrderBy("created_at_desc")
)

func (enum ListRoutesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListRoutesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListRoutesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListRoutesRequestOrderBy(ListRoutesRequestOrderBy(tmp).String())
	return nil
}

type ListSubscriberRequestOrderBy string

const (
	ListSubscriberRequestOrderByCreatedAtAsc  = ListSubscriberRequestOrderBy("created_at_asc")
	ListSubscriberRequestOrderByCreatedAtDesc = ListSubscriberRequestOrderBy("created_at_desc")
	ListSubscriberRequestOrderByNameAsc       = ListSubscriberRequestOrderBy("name_asc")
	ListSubscriberRequestOrderByNameDesc      = ListSubscriberRequestOrderBy("name_desc")
)

func (enum ListSubscriberRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListSubscriberRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSubscriberRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSubscriberRequestOrderBy(ListSubscriberRequestOrderBy(tmp).String())
	return nil
}

type OnMarkedDownAction string

const (
	OnMarkedDownActionOnMarkedDownActionNone = OnMarkedDownAction("on_marked_down_action_none")
	OnMarkedDownActionShutdownSessions       = OnMarkedDownAction("shutdown_sessions")
)

func (enum OnMarkedDownAction) String() string {
	if enum == "" {
		// return default value if empty
		return "on_marked_down_action_none"
	}
	return string(enum)
}

func (enum OnMarkedDownAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OnMarkedDownAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OnMarkedDownAction(OnMarkedDownAction(tmp).String())
	return nil
}

type PrivateNetworkStatus string

const (
	PrivateNetworkStatusUnknown = PrivateNetworkStatus("unknown")
	PrivateNetworkStatusReady   = PrivateNetworkStatus("ready")
	PrivateNetworkStatusPending = PrivateNetworkStatus("pending")
	PrivateNetworkStatusError   = PrivateNetworkStatus("error")
)

func (enum PrivateNetworkStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum PrivateNetworkStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PrivateNetworkStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PrivateNetworkStatus(PrivateNetworkStatus(tmp).String())
	return nil
}

type Protocol string

const (
	ProtocolTCP  = Protocol("tcp")
	ProtocolHTTP = Protocol("http")
)

func (enum Protocol) String() string {
	if enum == "" {
		// return default value if empty
		return "tcp"
	}
	return string(enum)
}

func (enum Protocol) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *Protocol) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = Protocol(Protocol(tmp).String())
	return nil
}

type ProxyProtocol string

const (
	ProxyProtocolProxyProtocolUnknown = ProxyProtocol("proxy_protocol_unknown")
	ProxyProtocolProxyProtocolNone    = ProxyProtocol("proxy_protocol_none")
	ProxyProtocolProxyProtocolV1      = ProxyProtocol("proxy_protocol_v1")
	ProxyProtocolProxyProtocolV2      = ProxyProtocol("proxy_protocol_v2")
	ProxyProtocolProxyProtocolV2Ssl   = ProxyProtocol("proxy_protocol_v2_ssl")
	ProxyProtocolProxyProtocolV2SslCn = ProxyProtocol("proxy_protocol_v2_ssl_cn")
)

func (enum ProxyProtocol) String() string {
	if enum == "" {
		// return default value if empty
		return "proxy_protocol_unknown"
	}
	return string(enum)
}

func (enum ProxyProtocol) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ProxyProtocol) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ProxyProtocol(ProxyProtocol(tmp).String())
	return nil
}

type SSLCompatibilityLevel string

const (
	SSLCompatibilityLevelSslCompatibilityLevelUnknown      = SSLCompatibilityLevel("ssl_compatibility_level_unknown")
	SSLCompatibilityLevelSslCompatibilityLevelIntermediate = SSLCompatibilityLevel("ssl_compatibility_level_intermediate")
	SSLCompatibilityLevelSslCompatibilityLevelModern       = SSLCompatibilityLevel("ssl_compatibility_level_modern")
	SSLCompatibilityLevelSslCompatibilityLevelOld          = SSLCompatibilityLevel("ssl_compatibility_level_old")
)

func (enum SSLCompatibilityLevel) String() string {
	if enum == "" {
		// return default value if empty
		return "ssl_compatibility_level_unknown"
	}
	return string(enum)
}

func (enum SSLCompatibilityLevel) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SSLCompatibilityLevel) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SSLCompatibilityLevel(SSLCompatibilityLevel(tmp).String())
	return nil
}

type StickySessionsType string

const (
	StickySessionsTypeNone   = StickySessionsType("none")
	StickySessionsTypeCookie = StickySessionsType("cookie")
	StickySessionsTypeTable  = StickySessionsType("table")
)

func (enum StickySessionsType) String() string {
	if enum == "" {
		// return default value if empty
		return "none"
	}
	return string(enum)
}

func (enum StickySessionsType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *StickySessionsType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = StickySessionsType(StickySessionsType(tmp).String())
	return nil
}

type SubscriberEmailConfig struct {
	// Email: Email address to send alerts to.
	Email string `json:"email"`
}

// Webhook alert of subscriber.
type SubscriberWebhookConfig struct {
	// URI: URI to receive POST requests.
	URI string `json:"uri"`
}

type HealthCheckHTTPConfig struct {
	// URI: The HTTP URI to use when performing a health check on backend servers.
	URI string `json:"uri"`
	// Method: The HTTP method used when performing a health check on backend servers.
	Method string `json:"method"`
	// Code: The HTTP response code that should be returned for a health check to be considered successful.
	Code *int32 `json:"code,omitempty"`
	// HostHeader: The HTTP host header used when performing a health check on backend servers.
	HostHeader string `json:"host_header"`
}

type HealthCheckHTTPSConfig struct {
	// URI: The HTTP URI to use when performing a health check on backend servers.
	URI string `json:"uri"`
	// Method: The HTTP method used when performing a health check on backend servers.
	Method string `json:"method"`
	// Code: The HTTP response code that should be returned for a health check to be considered successful.
	Code *int32 `json:"code,omitempty"`
	// HostHeader: The HTTP host header used when performing a health check on backend servers.
	HostHeader string `json:"host_header"`
	// Sni: The SNI value used when performing a health check on backend servers over SSL.
	Sni string `json:"sni"`
}

type HealthCheckLdapConfig struct {
}

type HealthCheckMysqlConfig struct {
	// User: MySQL user to use for the health check.
	User string `json:"user"`
}

type HealthCheckPgsqlConfig struct {
	// User: PostgreSQL user to use for the health check.
	User string `json:"user"`
}

type HealthCheckRedisConfig struct {
}

type HealthCheckTCPConfig struct {
}

type IP struct {
	// ID: IP address ID.
	ID string `json:"id"`
	// IPAddress: IP address.
	IPAddress string `json:"ip_address"`
	// OrganizationID: Organization ID of the Scaleway Organization the IP address is in.
	OrganizationID string `json:"organization_id"`
	// ProjectID: Project ID of the Scaleway Project the IP address is in.
	ProjectID string `json:"project_id"`
	// LBID: Load Balancer ID.
	LBID *string `json:"lb_id,omitempty"`
	// Reverse: Reverse DNS (domain name) of the IP address.
	Reverse string `json:"reverse"`
	// Region: The region the IP address is in.
	Region *scw.Region `json:"region,omitempty"`
	// Zone: The zone the IP address is in.
	Zone scw.Zone `json:"zone"`
}

type Instance struct {
	// ID: Underlying Instance ID.
	ID string `json:"id"`
	// Status: Instance status.
	Status InstanceStatus `json:"status"`
	// IPAddress: Instance IP address.
	IPAddress string `json:"ip_address"`
	// CreatedAt: Date on which the Instance was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Date on which the Instance was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Region: The region the Instance is in.
	Region *scw.Region `json:"region,omitempty"`
	// Zone: The zone the Instance is in.
	Zone scw.Zone `json:"zone"`
}

// Subscriber.
type Subscriber struct {
	// ID: Subscriber ID.
	ID string `json:"id"`
	// Name: Subscriber name.
	Name string `json:"name"`
	// EmailConfig: Email address of subscriber.
	EmailConfig *SubscriberEmailConfig `json:"email_config,omitempty"`
	// WebhookConfig: Webhook URI of subscriber.
	WebhookConfig *SubscriberWebhookConfig `json:"webhook_config,omitempty"`
}

type HealthCheck struct {
	// Port: Port to use for the backend server health check.
	Port int32 `json:"port"`
	// CheckDelay: Time to wait between two consecutive health checks.
	CheckDelay *time.Duration `json:"check_delay,omitempty"`
	// CheckTimeout: Maximum time a backend server has to reply to the health check.
	CheckTimeout *time.Duration `json:"check_timeout,omitempty"`
	// CheckMaxRetries: Number of consecutive unsuccessful health checks after which the server will be considered dead.
	CheckMaxRetries int32 `json:"check_max_retries"`
	// TCPConfig: Object to configure a basic TCP health check.
	TCPConfig *HealthCheckTCPConfig `json:"tcp_config,omitempty"`
	// MysqlConfig: Object to configure a MySQL health check. The check requires MySQL >=3.22, for older versions, use a TCP health check.
	MysqlConfig *HealthCheckMysqlConfig `json:"mysql_config,omitempty"`
	// PgsqlConfig: Object to configure a PostgreSQL health check.
	PgsqlConfig *HealthCheckPgsqlConfig `json:"pgsql_config,omitempty"`
	// LdapConfig: Object to configure an LDAP health check. The response is analyzed to find the LDAPv3 response message.
	LdapConfig *HealthCheckLdapConfig `json:"ldap_config,omitempty"`
	// RedisConfig: Object to configure a Redis health check. The response is analyzed to find the +PONG response message.
	RedisConfig *HealthCheckRedisConfig `json:"redis_config,omitempty"`
	// HTTPConfig: Object to configure an HTTP health check.
	HTTPConfig *HealthCheckHTTPConfig `json:"http_config,omitempty"`
	// HTTPSConfig: Object to configure an HTTPS health check.
	HTTPSConfig *HealthCheckHTTPSConfig `json:"https_config,omitempty"`
	// CheckSendProxy: Defines whether proxy protocol should be activated for the health check.
	CheckSendProxy bool `json:"check_send_proxy"`
	// TransientCheckDelay: Time to wait between two consecutive health checks when a backend server is in a transient state (going UP or DOWN).
	TransientCheckDelay *scw.Duration `json:"transient_check_delay,omitempty"`
}

func (m *HealthCheck) UnmarshalJSON(b []byte) error {
	type tmpType HealthCheck
	tmp := struct {
		tmpType
		TmpCheckDelay   *marshaler.Duration `json:"check_delay,omitempty"`
		TmpCheckTimeout *marshaler.Duration `json:"check_timeout,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = HealthCheck(tmp.tmpType)
	m.CheckDelay = tmp.TmpCheckDelay.Standard()
	m.CheckTimeout = tmp.TmpCheckTimeout.Standard()
	return nil
}

func (m HealthCheck) MarshalJSON() ([]byte, error) {
	type tmpType HealthCheck
	tmp := struct {
		tmpType
		TmpCheckDelay   *marshaler.Duration `json:"check_delay,omitempty"`
		TmpCheckTimeout *marshaler.Duration `json:"check_timeout,omitempty"`
	}{
		tmpType:         tmpType(m),
		TmpCheckDelay:   marshaler.NewDuration(m.CheckDelay),
		TmpCheckTimeout: marshaler.NewDuration(m.CheckTimeout),
	}
	return json.Marshal(tmp)
}

type LB struct {
	// ID: Underlying Instance ID.
	ID string `json:"id"`
	// Name: Load Balancer name.
	Name string `json:"name"`
	// Description: Load Balancer description.
	Description string `json:"description"`
	// Status: Load Balancer status.
	Status LBStatus `json:"status"`
	// Instances: List of underlying Instances.
	Instances []*Instance `json:"instances"`
	// OrganizationID: Scaleway Organization ID.
	OrganizationID string `json:"organization_id"`
	// ProjectID: Scaleway Project ID.
	ProjectID string `json:"project_id"`
	// IP: List of IP addresses attached to the Load Balancer.
	IP []*IP `json:"ip"`
	// Tags: Load Balancer tags.
	Tags []string `json:"tags"`
	// FrontendCount: Number of frontends the Load Balancer has.
	FrontendCount int32 `json:"frontend_count"`
	// BackendCount: Number of backends the Load Balancer has.
	BackendCount int32 `json:"backend_count"`
	// Type: Load Balancer offer type.
	Type string `json:"type"`
	// Subscriber: Subscriber information.
	Subscriber *Subscriber `json:"subscriber"`
	// SslCompatibilityLevel: Determines the minimal SSL version which needs to be supported on client side.
	SslCompatibilityLevel SSLCompatibilityLevel `json:"ssl_compatibility_level"`
	// CreatedAt: Date on which the Load Balancer was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Date on which the Load Balancer was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// PrivateNetworkCount: Number of Private Networks attached to the Load Balancer.
	PrivateNetworkCount int32 `json:"private_network_count"`
	// RouteCount: Number of routes configured on the Load Balancer.
	RouteCount int32 `json:"route_count"`
	// Region: The region the Load Balancer is in.
	Region *scw.Region `json:"region,omitempty"`
	// Zone: The zone the Load Balancer is in.
	Zone scw.Zone `json:"zone"`
}

type ACLActionRedirect struct {
	// Type: Redirect type.
	Type ACLActionRedirectRedirectType `json:"type"`
	// Target: Redirect target. For a location redirect, you can use a URL e.g. `https://scaleway.com`. Using a scheme name (e.g. `https`, `http`, `ftp`, `git`) will replace the request's original scheme. This can be useful to implement HTTP to HTTPS redirects. Valid placeholders that can be used in a `location` redirect to preserve parts of the original request in the redirection URL are \{\{host\}\}, \{\{query\}\}, \{\{path\}\} and \{\{scheme\}\}.
	Target string `json:"target"`
	// Code: HTTP redirect code to use. Valid values are 301, 302, 303, 307 and 308. Default value is 302.
	Code *int32 `json:"code,omitempty"`
}

type Backend struct {
	// ID: Backend ID.
	ID string `json:"id"`
	// Name: Name of the backend.
	Name string `json:"name"`
	// ForwardProtocol: Protocol used by the backend when forwarding traffic to backend servers.
	ForwardProtocol Protocol `json:"forward_protocol"`
	// ForwardPort: Port used by the backend when forwarding traffic to backend servers.
	ForwardPort int32 `json:"forward_port"`
	// ForwardPortAlgorithm: Load balancing algorithm to use when determining which backend server to forward new traffic to.
	ForwardPortAlgorithm ForwardPortAlgorithm `json:"forward_port_algorithm"`
	// StickySessions: Defines whether sticky sessions (binding a particular session to a particular backend server) are activated and the method to use if so. None disables sticky sessions. Cookie-based uses an HTTP cookie to stick a session to a backend server. Table-based uses the source (client) IP address to stick a session to a backend server.
	StickySessions StickySessionsType `json:"sticky_sessions"`
	// StickySessionsCookieName: Cookie name for cookie-based sticky sessions.
	StickySessionsCookieName string `json:"sticky_sessions_cookie_name"`
	// HealthCheck: Object defining the health check to be carried out by the backend when checking the status and health of backend servers.
	HealthCheck *HealthCheck `json:"health_check"`
	// Pool: List of IP addresses of backend servers attached to this backend.
	Pool []string `json:"pool"`
	// LB: Load Balancer the backend is attached to.
	LB *LB `json:"lb"`
	// SendProxyV2: Deprecated in favor of proxy_protocol field.
	SendProxyV2 *bool `json:"send_proxy_v2,omitempty"`
	// TimeoutServer: Maximum allowed time for a backend server to process a request.
	TimeoutServer *time.Duration `json:"timeout_server,omitempty"`
	// TimeoutConnect: Maximum allowed time for establishing a connection to a backend server.
	TimeoutConnect *time.Duration `json:"timeout_connect,omitempty"`
	// TimeoutTunnel: Maximum allowed tunnel inactivity time after Websocket is established (takes precedence over client and server timeout).
	TimeoutTunnel *time.Duration `json:"timeout_tunnel,omitempty"`
	// OnMarkedDownAction: Action to take when a backend server is marked as down.
	OnMarkedDownAction OnMarkedDownAction `json:"on_marked_down_action"`
	// ProxyProtocol: Protocol to use between the Load Balancer and backend servers. Allows the backend servers to be informed of the client's real IP address. The PROXY protocol must be supported by the backend servers' software.
	ProxyProtocol ProxyProtocol `json:"proxy_protocol"`
	// CreatedAt: Date at which the backend was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Date at which the backend was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// FailoverHost: Scaleway S3 bucket website to be served as failover if all backend servers are down, e.g. failover-website.s3-website.fr-par.scw.cloud.
	FailoverHost *string `json:"failover_host,omitempty"`
	// SslBridging: Defines whether to enable SSL bridging between the Load Balancer and backend servers.
	SslBridging *bool `json:"ssl_bridging,omitempty"`
	// IgnoreSslServerVerify: Defines whether the server certificate verification should be ignored.
	IgnoreSslServerVerify *bool `json:"ignore_ssl_server_verify,omitempty"`
	// RedispatchAttemptCount: Whether to use another backend server on each attempt.
	RedispatchAttemptCount *int32 `json:"redispatch_attempt_count,omitempty"`
	// MaxRetries: Number of retries when a backend server connection failed.
	MaxRetries *int32 `json:"max_retries,omitempty"`
	// MaxConnections: Maximum number of connections allowed per backend server.
	MaxConnections *int32 `json:"max_connections,omitempty"`
	// TimeoutQueue: Maximum time for a request to be left pending in queue when `max_connections` is reached.
	TimeoutQueue *scw.Duration `json:"timeout_queue,omitempty"`
}

func (m *Backend) UnmarshalJSON(b []byte) error {
	type tmpType Backend
	tmp := struct {
		tmpType
		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = Backend(tmp.tmpType)
	m.TimeoutServer = tmp.TmpTimeoutServer.Standard()
	m.TimeoutConnect = tmp.TmpTimeoutConnect.Standard()
	m.TimeoutTunnel = tmp.TmpTimeoutTunnel.Standard()
	return nil
}

func (m Backend) MarshalJSON() ([]byte, error) {
	type tmpType Backend
	tmp := struct {
		tmpType
		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
	}{
		tmpType:           tmpType(m),
		TmpTimeoutServer:  marshaler.NewDuration(m.TimeoutServer),
		TmpTimeoutConnect: marshaler.NewDuration(m.TimeoutConnect),
		TmpTimeoutTunnel:  marshaler.NewDuration(m.TimeoutTunnel),
	}
	return json.Marshal(tmp)
}

type Certificate struct {
	// Type: Certificate type (Let's Encrypt or custom).
	Type CertificateType `json:"type"`
	// ID: Certificate ID.
	ID string `json:"id"`
	// CommonName: Main domain name of certificate.
	CommonName string `json:"common_name"`
	// SubjectAlternativeName: Alternative domain names.
	SubjectAlternativeName []string `json:"subject_alternative_name"`
	// Fingerprint: Identifier (SHA-1) of the certificate.
	Fingerprint string `json:"fingerprint"`
	// NotValidBefore: Lower validity bound.
	NotValidBefore *time.Time `json:"not_valid_before,omitempty"`
	// NotValidAfter: Upper validity bound.
	NotValidAfter *time.Time `json:"not_valid_after,omitempty"`
	// Status: Certificate status.
	Status CertificateStatus `json:"status"`
	// LB: Load Balancer object the certificate is attached to.
	LB *LB `json:"lb"`
	// Name: Certificate name.
	Name string `json:"name"`
	// CreatedAt: Date on which the certificate was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Date on which the certificate was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// StatusDetails: Additional information about the certificate status (useful in case of certificate generation failure, for example).
	StatusDetails *string `json:"status_details,omitempty"`
}

type ACLAction struct {
	// Type: Action to take when incoming traffic matches an ACL filter.
	Type ACLActionType `json:"type"`
	// Redirect: Redirection parameters when using an ACL with a `redirect` action.
	Redirect *ACLActionRedirect `json:"redirect"`
}

type ACLMatch struct {
	// IPSubnet: List of IPs or CIDR v4/v6 addresses to filter for from the client side.
	IPSubnet []*string `json:"ip_subnet"`
	// HTTPFilter: Type of HTTP filter to match. Extracts the request's URL path, which starts at the first slash and ends before the question mark (without the host part). Defines where to filter for the http_filter_value. Only supported for HTTP backends.
	HTTPFilter ACLHTTPFilter `json:"http_filter"`
	// HTTPFilterValue: List of values to filter for.
	HTTPFilterValue []*string `json:"http_filter_value"`
	// HTTPFilterOption: Name of the HTTP header to filter on if `http_header_match` was selected in `http_filter`.
	HTTPFilterOption *string `json:"http_filter_option,omitempty"`
	// Invert: Defines whether to invert the match condition. If set to `true`, the ACL carries out its action when the condition DOES NOT match.
	Invert bool `json:"invert"`
}

type Frontend struct {
	// ID: Frontend ID.
	ID string `json:"id"`
	// Name: Name of the frontend.
	Name string `json:"name"`
	// InboundPort: Port the frontend listens on.
	InboundPort int32 `json:"inbound_port"`
	// Backend: Backend object the frontend is attached to.
	Backend *Backend `json:"backend"`
	// LB: Load Balancer object the frontend is attached to.
	LB *LB `json:"lb"`
	// TimeoutClient: Maximum allowed inactivity time on the client side.
	TimeoutClient *time.Duration `json:"timeout_client,omitempty"`
	// Certificate: Certificate, deprecated in favor of certificate_ids array.
	Certificate *Certificate `json:"certificate,omitempty"`
	// CertificateIDs: List of SSL/TLS certificate IDs to bind to the frontend.
	CertificateIDs []string `json:"certificate_ids"`
	// CreatedAt: Date on which the frontend was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Date on which the frontend was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// EnableHTTP3: Defines whether to enable HTTP/3 protocol on the frontend.
	EnableHTTP3 bool `json:"enable_http3"`
}

func (m *Frontend) UnmarshalJSON(b []byte) error {
	type tmpType Frontend
	tmp := struct {
		tmpType
		TmpTimeoutClient *marshaler.Duration `json:"timeout_client,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = Frontend(tmp.tmpType)
	m.TimeoutClient = tmp.TmpTimeoutClient.Standard()
	return nil
}

func (m Frontend) MarshalJSON() ([]byte, error) {
	type tmpType Frontend
	tmp := struct {
		tmpType
		TmpTimeoutClient *marshaler.Duration `json:"timeout_client,omitempty"`
	}{
		tmpType:          tmpType(m),
		TmpTimeoutClient: marshaler.NewDuration(m.TimeoutClient),
	}
	return json.Marshal(tmp)
}

type PrivateNetworkDHCPConfig struct {
}

type PrivateNetworkIpamConfig struct {
}

type PrivateNetworkStaticConfig struct {
	// IPAddress: Array of a local IP address for the Load Balancer on this Private Network.
	IPAddress []string `json:"ip_address"`
}

type RouteMatch struct {
	// Sni: Value to match in the Server Name Indication TLS extension (SNI) field from an incoming connection made via an SSL/TLS transport layer. This field should be set for routes on TCP Load Balancers.
	Sni *string `json:"sni,omitempty"`
	// HostHeader: Value to match in the HTTP Host request header from an incoming connection. This field should be set for routes on HTTP Load Balancers.
	HostHeader *string `json:"host_header,omitempty"`
}

type CreateCertificateRequestCustomCertificate struct {
	// CertificateChain: Full PEM-formatted certificate, consisting of the entire certificate chain including public key, private key, and (optionally) Certificate Authorities.
	CertificateChain string `json:"certificate_chain"`
}

type CreateCertificateRequestLetsencryptConfig struct {
	// CommonName: Main domain name of certificate (this domain must exist and resolve to your Load Balancer IP address).
	CommonName string `json:"common_name"`
	// SubjectAlternativeName: Alternative domain names (all domain names must exist and resolve to your Load Balancer IP address).
	SubjectAlternativeName []string `json:"subject_alternative_name"`
}

type BackendServerStats struct {
	// InstanceID: ID of your Load Balancer's underlying Instance.
	InstanceID string `json:"instance_id"`
	// BackendID: Backend ID.
	BackendID string `json:"backend_id"`
	// IP: IPv4 or IPv6 address of the backend server.
	IP string `json:"ip"`
	// ServerState: Server operational state (stopped/starting/running/stopping).
	ServerState BackendServerStatsServerState `json:"server_state"`
	// ServerStateChangedAt: Time since last operational change.
	ServerStateChangedAt *time.Time `json:"server_state_changed_at,omitempty"`
	// LastHealthCheckStatus: Last health check status (unknown/neutral/failed/passed/condpass).
	LastHealthCheckStatus BackendServerStatsHealthCheckStatus `json:"last_health_check_status"`
}

type ACL struct {
	// ID: ACL ID.
	ID string `json:"id"`
	// Name: ACL name.
	Name string `json:"name"`
	// Match: ACL match filter object. One of `ip_subnet` or `http_filter` & `http_filter_value` are required.
	Match *ACLMatch `json:"match"`
	// Action: Action to take when incoming traffic matches an ACL filter.
	Action *ACLAction `json:"action"`
	// Frontend: ACL is attached to this frontend object.
	Frontend *Frontend `json:"frontend"`
	// Index: Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
	Index int32 `json:"index"`
	// CreatedAt: Date on which the ACL was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Date on which the ACL was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Description: ACL description.
	Description string `json:"description"`
}

type PrivateNetwork struct {
	// LB: Load Balancer object which is attached to the Private Network.
	LB *LB `json:"lb"`
	// StaticConfig: Object containing an array of a local IP address for the Load Balancer on this Private Network.
	StaticConfig *PrivateNetworkStaticConfig `json:"static_config,omitempty"`
	// DHCPConfig: Defines whether to let DHCP assign IP addresses.
	DHCPConfig *PrivateNetworkDHCPConfig `json:"dhcp_config,omitempty"`
	// IpamConfig: For internal use only.
	IpamConfig *PrivateNetworkIpamConfig `json:"ipam_config,omitempty"`
	// PrivateNetworkID: Private Network ID.
	PrivateNetworkID string `json:"private_network_id"`
	// Status: Status of Private Network connection.
	Status PrivateNetworkStatus `json:"status"`
	// CreatedAt: Date on which the Private Network was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Date on which the PN was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type LBType struct {
	// Name: Load Balancer commercial offer type name.
	Name string `json:"name"`
	// StockStatus: Current stock status for a given Load Balancer type.
	StockStatus LBTypeStock `json:"stock_status"`
	// Description: Load Balancer commercial offer type description.
	Description string `json:"description"`
	// Region: The region the Load Balancer stock is in.
	Region *scw.Region `json:"region,omitempty"`
	// Zone: The zone the Load Balancer stock is in.
	Zone scw.Zone `json:"zone"`
}

type Route struct {
	// ID: Route ID.
	ID string `json:"id"`
	// FrontendID: ID of the source frontend.
	FrontendID string `json:"frontend_id"`
	// BackendID: ID of the target backend.
	BackendID string `json:"backend_id"`
	// Match: Object defining the match condition for a route to be applied. If an incoming client session matches the specified condition (i.e. it has a matching SNI value or HTTP Host header value), it will be passed to the target backend.
	Match *RouteMatch `json:"match"`
	// CreatedAt: Date on which the route was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Date on which the route was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type ACLSpec struct {
	// Name: ACL name.
	Name string `json:"name"`
	// Action: Action to take when incoming traffic matches an ACL filter.
	Action *ACLAction `json:"action"`
	// Match: ACL match filter object. One of `ip_subnet` or `http_filter` and `http_filter_value` are required.
	Match *ACLMatch `json:"match"`
	// Index: Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
	Index int32 `json:"index"`
	// Description: ACL description.
	Description string `json:"description"`
}

type AddBackendServersRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// BackendID: Backend ID.
	BackendID string `json:"-"`
	// ServerIP: List of IP addresses to add to backend servers.
	ServerIP []string `json:"server_ip"`
}

type AttachPrivateNetworkRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// PrivateNetworkID: Private Network ID.
	PrivateNetworkID string `json:"-"`
	// StaticConfig: Object containing an array of a local IP address for the Load Balancer on this Private Network.
	StaticConfig *PrivateNetworkStaticConfig `json:"static_config,omitempty"`
	// DHCPConfig: Defines whether to let DHCP assign IP addresses.
	DHCPConfig *PrivateNetworkDHCPConfig `json:"dhcp_config,omitempty"`
	// IpamConfig: For internal use only.
	IpamConfig *PrivateNetworkIpamConfig `json:"ipam_config,omitempty"`
}

// Add an ACL to a Load Balancer frontend.
type CreateACLRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// FrontendID: Frontend ID to attach the ACL to.
	FrontendID string `json:"-"`
	// Name: ACL name.
	Name string `json:"name"`
	// Action: Action to take when incoming traffic matches an ACL filter.
	Action *ACLAction `json:"action"`
	// Match: ACL match filter object. One of `ip_subnet` or `http_filter` & `http_filter_value` are required.
	Match *ACLMatch `json:"match"`
	// Index: Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
	Index int32 `json:"index"`
	// Description: ACL description.
	Description string `json:"description"`
}

type CreateBackendRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Name: Name for the backend.
	Name string `json:"name"`
	// ForwardProtocol: Protocol to be used by the backend when forwarding traffic to backend servers.
	ForwardProtocol Protocol `json:"forward_protocol"`
	// ForwardPort: Port to be used by the backend when forwarding traffic to backend servers.
	ForwardPort int32 `json:"forward_port"`
	// ForwardPortAlgorithm: Load balancing algorithm to be used when determining which backend server to forward new traffic to.
	ForwardPortAlgorithm ForwardPortAlgorithm `json:"forward_port_algorithm"`
	// StickySessions: Defines whether to activate sticky sessions (binding a particular session to a particular backend server) and the method to use if so. None disables sticky sessions. Cookie-based uses an HTTP cookie TO stick a session to a backend server. Table-based uses the source (client) IP address to stick a session to a backend server.
	StickySessions StickySessionsType `json:"sticky_sessions"`
	// StickySessionsCookieName: Cookie name for cookie-based sticky sessions.
	StickySessionsCookieName string `json:"sticky_sessions_cookie_name"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// HealthCheck: Object defining the health check to be carried out by the backend when checking the status and health of backend servers.
	HealthCheck *HealthCheck `json:"health_check"`
	// ServerIP: List of backend server IP addresses (IPv4 or IPv6) the backend should forward traffic to.
	ServerIP []string `json:"server_ip"`
	// SendProxyV2: Deprecated in favor of proxy_protocol field.
	SendProxyV2 *bool `json:"send_proxy_v2,omitempty"`
	// TimeoutServer: Maximum allowed time for a backend server to process a request.
	TimeoutServer *time.Duration `json:"timeout_server,omitempty"`
	// TimeoutConnect: Maximum allowed time for establishing a connection to a backend server.
	TimeoutConnect *time.Duration `json:"timeout_connect,omitempty"`
	// TimeoutTunnel: Maximum allowed tunnel inactivity time after Websocket is established (takes precedence over client and server timeout).
	TimeoutTunnel *time.Duration `json:"timeout_tunnel,omitempty"`
	// OnMarkedDownAction: Action to take when a backend server is marked as down.
	OnMarkedDownAction OnMarkedDownAction `json:"on_marked_down_action"`
	// ProxyProtocol: Protocol to use between the Load Balancer and backend servers. Allows the backend servers to be informed of the client's real IP address. The PROXY protocol must be supported by the backend servers' software.
	ProxyProtocol ProxyProtocol `json:"proxy_protocol"`
	// FailoverHost: Scaleway S3 bucket website to be served as failover if all backend servers are down, e.g. failover-website.s3-website.fr-par.scw.cloud.
	FailoverHost *string `json:"failover_host,omitempty"`
	// SslBridging: Defines whether to enable SSL bridging between the Load Balancer and backend servers.
	SslBridging *bool `json:"ssl_bridging,omitempty"`
	// IgnoreSslServerVerify: Defines whether the server certificate verification should be ignored.
	IgnoreSslServerVerify *bool `json:"ignore_ssl_server_verify,omitempty"`
	// RedispatchAttemptCount: Whether to use another backend server on each attempt.
	RedispatchAttemptCount *int32 `json:"redispatch_attempt_count,omitempty"`
	// MaxRetries: Number of retries when a backend server connection failed.
	MaxRetries *int32 `json:"max_retries,omitempty"`
	// MaxConnections: Maximum number of connections allowed per backend server.
	MaxConnections *int32 `json:"max_connections,omitempty"`
	// TimeoutQueue: Maximum time for a request to be left pending in queue when `max_connections` is reached.
	TimeoutQueue *scw.Duration `json:"timeout_queue,omitempty"`
}

func (m *CreateBackendRequest) UnmarshalJSON(b []byte) error {
	type tmpType CreateBackendRequest
	tmp := struct {
		tmpType
		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = CreateBackendRequest(tmp.tmpType)
	m.TimeoutServer = tmp.TmpTimeoutServer.Standard()
	m.TimeoutConnect = tmp.TmpTimeoutConnect.Standard()
	m.TimeoutTunnel = tmp.TmpTimeoutTunnel.Standard()
	return nil
}

func (m CreateBackendRequest) MarshalJSON() ([]byte, error) {
	type tmpType CreateBackendRequest
	tmp := struct {
		tmpType
		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
	}{
		tmpType:           tmpType(m),
		TmpTimeoutServer:  marshaler.NewDuration(m.TimeoutServer),
		TmpTimeoutConnect: marshaler.NewDuration(m.TimeoutConnect),
		TmpTimeoutTunnel:  marshaler.NewDuration(m.TimeoutTunnel),
	}
	return json.Marshal(tmp)
}

type CreateCertificateRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// Name: Name for the certificate.
	Name string `json:"name"`
	// Letsencrypt: Object to define a new Let's Encrypt certificate to be generated.
	Letsencrypt *CreateCertificateRequestLetsencryptConfig `json:"letsencrypt,omitempty"`
	// CustomCertificate: Object to define an existing custom certificate to be imported.
	CustomCertificate *CreateCertificateRequestCustomCertificate `json:"custom_certificate,omitempty"`
}

type CreateFrontendRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Name: Name for the frontend.
	Name string `json:"name"`
	// InboundPort: Port the frontend should listen on.
	InboundPort int32 `json:"inbound_port"`
	// LBID: Load Balancer ID (ID of the Load Balancer to attach the frontend to).
	LBID string `json:"-"`
	// BackendID: Backend ID (ID of the backend the frontend should pass traffic to).
	BackendID string `json:"backend_id"`
	// TimeoutClient: Maximum allowed inactivity time on the client side.
	TimeoutClient *time.Duration `json:"timeout_client,omitempty"`
	// CertificateID: Certificate ID, deprecated in favor of certificate_ids array.
	CertificateID *string `json:"certificate_id,omitempty"`
	// CertificateIDs: List of SSL/TLS certificate IDs to bind to the frontend.
	CertificateIDs *[]string `json:"certificate_ids,omitempty"`
	// EnableHTTP3: Defines whether to enable HTTP/3 protocol on the frontend.
	EnableHTTP3 bool `json:"enable_http3"`
}

func (m *CreateFrontendRequest) UnmarshalJSON(b []byte) error {
	type tmpType CreateFrontendRequest
	tmp := struct {
		tmpType
		TmpTimeoutClient *marshaler.Duration `json:"timeout_client,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = CreateFrontendRequest(tmp.tmpType)
	m.TimeoutClient = tmp.TmpTimeoutClient.Standard()
	return nil
}

func (m CreateFrontendRequest) MarshalJSON() ([]byte, error) {
	type tmpType CreateFrontendRequest
	tmp := struct {
		tmpType
		TmpTimeoutClient *marshaler.Duration `json:"timeout_client,omitempty"`
	}{
		tmpType:          tmpType(m),
		TmpTimeoutClient: marshaler.NewDuration(m.TimeoutClient),
	}
	return json.Marshal(tmp)
}

type CreateIPRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// OrganizationID: Organization ID of the Organization where the IP address should be created.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: Project ID of the Project where the IP address should be created.
	ProjectID *string `json:"project_id,omitempty"`
	// Reverse: Reverse DNS (domain name) for the IP address.
	Reverse *string `json:"reverse,omitempty"`
	// IsIPv6: If true, creates a Flexible IP with an ipv6 address.
	IsIPv6 bool `json:"is_ipv6"`
}

type CreateLBRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// OrganizationID: Scaleway Organization to create the Load Balancer in.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: Scaleway Project to create the Load Balancer in.
	ProjectID *string `json:"project_id,omitempty"`
	// Name: Name for the Load Balancer.
	Name string `json:"name"`
	// Description: Description for the Load Balancer.
	Description string `json:"description"`
	// IPID: ID of an existing flexible IP address to attach to the Load Balancer.
	IPID *string `json:"ip_id,omitempty"`
	// AssignFlexibleIP: Defines whether to automatically assign a flexible public IP to lb. Default value is `false` (do not assign).
	AssignFlexibleIP *bool `json:"assign_flexible_ip,omitempty"`
	// IPIDs: List of IP IDs to attach to the Load Balancer.
	IPIDs []string `json:"ip_ids"`
	// Tags: List of tags for the Load Balancer.
	Tags []string `json:"tags"`
	// Type: Load Balancer commercial offer type. Use the Load Balancer types endpoint to retrieve a list of available offer types.
	Type string `json:"type"`
	// SslCompatibilityLevel: Determines the minimal SSL version which needs to be supported on the client side, in an SSL/TLS offloading context. Intermediate is suitable for general-purpose servers with a variety of clients, recommended for almost all systems. Modern is suitable for services with clients that support TLS 1.3 and do not need backward compatibility. Old is compatible with a small number of very old clients and should be used only as a last resort.
	SslCompatibilityLevel SSLCompatibilityLevel `json:"ssl_compatibility_level"`
}

type CreateRouteRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// FrontendID: ID of the source frontend to create the route on.
	FrontendID string `json:"frontend_id"`
	// BackendID: ID of the target backend for the route.
	BackendID string `json:"backend_id"`
	// Match: Object defining the match condition for a route to be applied. If an incoming client session matches the specified condition (i.e. it has a matching SNI value or HTTP Host header value), it will be passed to the target backend.
	Match *RouteMatch `json:"match"`
}

// Create a new alert subscriber (webhook or email).
type CreateSubscriberRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Name: Subscriber name.
	Name string `json:"name"`
	// EmailConfig: Email address configuration.
	EmailConfig *SubscriberEmailConfig `json:"email_config,omitempty"`
	// WebhookConfig: WebHook URI configuration.
	WebhookConfig *SubscriberWebhookConfig `json:"webhook_config,omitempty"`
	// OrganizationID: Organization ID to create the subscriber in.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: Project ID to create the subscriber in.
	ProjectID *string `json:"project_id,omitempty"`
}

type DeleteACLRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ACLID: ACL ID.
	ACLID string `json:"-"`
}

type DeleteBackendRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// BackendID: ID of the backend to delete.
	BackendID string `json:"-"`
}

type DeleteCertificateRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// CertificateID: Certificate ID.
	CertificateID string `json:"-"`
}

type DeleteFrontendRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// FrontendID: ID of the frontend to delete.
	FrontendID string `json:"-"`
}

type DeleteLBRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// LBID: ID of the Load Balancer to delete.
	LBID string `json:"-"`
	// ReleaseIP: Defines whether the Load Balancer's flexible IP should be deleted. Set to true to release the flexible IP, or false to keep it available in your account for future Load Balancers.
	ReleaseIP bool `json:"release_ip"`
}

type DeleteRouteRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// RouteID: Route ID.
	RouteID string `json:"-"`
}

type DeleteSubscriberRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SubscriberID: Subscriber ID.
	SubscriberID string `json:"-"`
}

type DetachPrivateNetworkRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// LBID: Load balancer ID.
	LBID string `json:"-"`
	// PrivateNetworkID: Set your instance private network id.
	PrivateNetworkID string `json:"-"`
}

type GetACLRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ACLID: ACL ID.
	ACLID string `json:"-"`
}

type GetBackendRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// BackendID: Backend ID.
	BackendID string `json:"-"`
}

type GetCertificateRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// CertificateID: Certificate ID.
	CertificateID string `json:"-"`
}

type GetFrontendRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// FrontendID: Frontend ID.
	FrontendID string `json:"-"`
}

type GetIPRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// IPID: IP address ID.
	IPID string `json:"-"`
}

type GetLBRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
}

// Get Load Balancer stats.
type GetLBStatsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// BackendID: ID of the backend.
	BackendID *string `json:"backend_id,omitempty"`
}

type GetRouteRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// RouteID: Route ID.
	RouteID string `json:"-"`
}

type GetSubscriberRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SubscriberID: Subscriber ID.
	SubscriberID string `json:"-"`
}

type LBStats struct {
	// BackendServersStats: List of objects containing Load Balancer statistics.
	BackendServersStats []*BackendServerStats `json:"backend_servers_stats"`
}

type ListACLResponse struct {
	// ACLs: List of ACL objects.
	ACLs []*ACL `json:"acls"`
	// TotalCount: The total number of objects.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListACLResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListACLResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListACLResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ACLs = append(r.ACLs, results.ACLs...)
	r.TotalCount += uint32(len(results.ACLs))
	return uint32(len(results.ACLs)), nil
}

type ListACLsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// FrontendID: Frontend ID (ACLs attached to this frontend will be returned in the response).
	FrontendID string `json:"-"`
	// OrderBy: Sort order of ACLs in the response.
	OrderBy ListACLRequestOrderBy `json:"order_by"`
	// Page: The page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// PageSize: The number of ACLs to return.
	PageSize *uint32 `json:"page_size,omitempty"`
	// Name: ACL name to filter for.
	Name *string `json:"name,omitempty"`
}

type ListBackendStatsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// Page: The page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Number of items to return.
	PageSize *uint32 `json:"page_size,omitempty"`
	// BackendID: ID of the backend.
	BackendID *string `json:"backend_id,omitempty"`
}

type ListBackendStatsResponse struct {
	// BackendServersStats: List of objects containing backend server statistics.
	BackendServersStats []*BackendServerStats `json:"backend_servers_stats"`
	// TotalCount: The total number of objects.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListBackendStatsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListBackendStatsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListBackendStatsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.BackendServersStats = append(r.BackendServersStats, results.BackendServersStats...)
	r.TotalCount += uint32(len(results.BackendServersStats))
	return uint32(len(results.BackendServersStats)), nil
}

type ListBackendsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// Name: Name of the backend to filter for.
	Name *string `json:"name,omitempty"`
	// OrderBy: Sort order of backends in the response.
	OrderBy ListBackendsRequestOrderBy `json:"order_by"`
	// Page: The page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Number of backends to return.
	PageSize *uint32 `json:"page_size,omitempty"`
}

type ListBackendsResponse struct {
	// Backends: List of backend objects of a given Load Balancer.
	Backends []*Backend `json:"backends"`
	// TotalCount: Total count of backend objects, without pagination.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListBackendsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListBackendsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListBackendsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Backends = append(r.Backends, results.Backends...)
	r.TotalCount += uint32(len(results.Backends))
	return uint32(len(results.Backends)), nil
}

type ListCertificatesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// OrderBy: Sort order of certificates in the response.
	OrderBy ListCertificatesRequestOrderBy `json:"order_by"`
	// Page: The page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Number of certificates to return.
	PageSize *uint32 `json:"page_size,omitempty"`
	// Name: Certificate name to filter for, only certificates of this name will be returned.
	Name *string `json:"name,omitempty"`
}

type ListCertificatesResponse struct {
	// Certificates: List of certificate objects.
	Certificates []*Certificate `json:"certificates"`
	// TotalCount: The total number of objects.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCertificatesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCertificatesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListCertificatesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Certificates = append(r.Certificates, results.Certificates...)
	r.TotalCount += uint32(len(results.Certificates))
	return uint32(len(results.Certificates)), nil
}

type ListFrontendsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// Name: Name of the frontend to filter for.
	Name *string `json:"name,omitempty"`
	// OrderBy: Sort order of frontends in the response.
	OrderBy ListFrontendsRequestOrderBy `json:"order_by"`
	// Page: The page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Number of frontends to return.
	PageSize *uint32 `json:"page_size,omitempty"`
}

type ListFrontendsResponse struct {
	// Frontends: List of frontend objects of a given Load Balancer.
	Frontends []*Frontend `json:"frontends"`
	// TotalCount: Total count of frontend objects, without pagination.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListFrontendsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListFrontendsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListFrontendsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Frontends = append(r.Frontends, results.Frontends...)
	r.TotalCount += uint32(len(results.Frontends))
	return uint32(len(results.Frontends)), nil
}

type ListIPsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Page: The page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Number of IP addresses to return.
	PageSize *uint32 `json:"page_size,omitempty"`
	// IPAddress: IP address to filter for.
	IPAddress *string `json:"ip_address,omitempty"`
	// OrganizationID: Organization ID to filter for, only Load Balancer IP addresses from this Organization will be returned.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: Project ID to filter for, only Load Balancer IP addresses from this Project will be returned.
	ProjectID *string `json:"project_id,omitempty"`
}

type ListIPsResponse struct {
	// IPs: List of IP address objects.
	IPs []*IP `json:"ips"`
	// TotalCount: Total count of IP address objects, without pagination.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListIPsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListIPsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListIPsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.IPs = append(r.IPs, results.IPs...)
	r.TotalCount += uint32(len(results.IPs))
	return uint32(len(results.IPs)), nil
}

type ListLBPrivateNetworksRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// OrderBy: Sort order of Private Network objects in the response.
	OrderBy ListPrivateNetworksRequestOrderBy `json:"order_by"`
	// PageSize: Number of objects to return.
	PageSize *uint32 `json:"page_size,omitempty"`
	// Page: The page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
}

type ListLBPrivateNetworksResponse struct {
	// PrivateNetwork: List of Private Network objects attached to the Load Balancer.
	PrivateNetwork []*PrivateNetwork `json:"private_network"`
	// TotalCount: Total number of objects in the response.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListLBPrivateNetworksResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListLBPrivateNetworksResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListLBPrivateNetworksResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PrivateNetwork = append(r.PrivateNetwork, results.PrivateNetwork...)
	r.TotalCount += uint32(len(results.PrivateNetwork))
	return uint32(len(results.PrivateNetwork)), nil
}

type ListLBTypesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Page: The page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// PageSize: The number of items to return.
	PageSize *uint32 `json:"page_size,omitempty"`
}

type ListLBTypesResponse struct {
	// LBTypes: List of Load Balancer commercial offer type objects.
	LBTypes []*LBType `json:"lb_types"`
	// TotalCount: Total number of Load Balancer offer type objects.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListLBTypesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListLBTypesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListLBTypesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.LBTypes = append(r.LBTypes, results.LBTypes...)
	r.TotalCount += uint32(len(results.LBTypes))
	return uint32(len(results.LBTypes)), nil
}

type ListLBsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Name: Load Balancer name to filter for.
	Name *string `json:"name,omitempty"`
	// OrderBy: Sort order of Load Balancers in the response.
	OrderBy ListLBsRequestOrderBy `json:"order_by"`
	// PageSize: Number of Load Balancers to return.
	PageSize *uint32 `json:"page_size,omitempty"`
	// Page: Page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// OrganizationID: Organization ID to filter for, only Load Balancers from this Organization will be returned.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: Project ID to filter for, only Load Balancers from this Project will be returned.
	ProjectID *string `json:"project_id,omitempty"`
}

type ListLBsResponse struct {
	// LBs: List of Load Balancer objects.
	LBs []*LB `json:"lbs"`
	// TotalCount: The total number of Load Balancer objects.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListLBsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListLBsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListLBsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.LBs = append(r.LBs, results.LBs...)
	r.TotalCount += uint32(len(results.LBs))
	return uint32(len(results.LBs)), nil
}

type ListRoutesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// OrderBy: Sort order of routes in the response.
	OrderBy ListRoutesRequestOrderBy `json:"order_by"`
	// PageSize: The number of route objects to return.
	PageSize *uint32 `json:"page_size,omitempty"`
	// Page: The page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// FrontendID: Frontend ID to filter for, only Routes from this Frontend will be returned.
	FrontendID *string `json:"frontend_id,omitempty"`
}

type ListRoutesResponse struct {
	// Routes: List of route objects.
	Routes []*Route `json:"routes"`
	// TotalCount: The total number of route objects.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRoutesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRoutesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListRoutesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Routes = append(r.Routes, results.Routes...)
	r.TotalCount += uint32(len(results.Routes))
	return uint32(len(results.Routes)), nil
}

type ListSubscriberRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// OrderBy: Sort order of subscribers in the response.
	OrderBy ListSubscriberRequestOrderBy `json:"order_by"`
	// Page: The page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// PageSize: The number of items to return.
	PageSize *uint32 `json:"page_size,omitempty"`
	// Name: Subscriber name to search for.
	Name *string `json:"name,omitempty"`
	// OrganizationID: Filter subscribers by Organization ID.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: Filter subscribers by Project ID.
	ProjectID *string `json:"project_id,omitempty"`
}

type ListSubscriberResponse struct {
	// Subscribers: List of subscriber objects.
	Subscribers []*Subscriber `json:"subscribers"`
	// TotalCount: The total number of objects.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSubscriberResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSubscriberResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSubscriberResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Subscribers = append(r.Subscribers, results.Subscribers...)
	r.TotalCount += uint32(len(results.Subscribers))
	return uint32(len(results.Subscribers)), nil
}

type MigrateLBRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// Type: Load Balancer type to migrate to (use the List all Load Balancer offer types endpoint to get a list of available offer types).
	Type string `json:"type"`
}

type ReleaseIPRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// IPID: IP address ID.
	IPID string `json:"-"`
}

type RemoveBackendServersRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// BackendID: Backend ID.
	BackendID string `json:"-"`
	// ServerIP: List of IP addresses to remove from backend servers.
	ServerIP []string `json:"server_ip"`
}

type SetACLsResponse struct {
	// ACLs: List of ACL objects.
	ACLs []*ACL `json:"acls"`
	// TotalCount: The total number of ACL objects.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *SetACLsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *SetACLsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*SetACLsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ACLs = append(r.ACLs, results.ACLs...)
	r.TotalCount += uint32(len(results.ACLs))
	return uint32(len(results.ACLs)), nil
}

type SetBackendServersRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// BackendID: Backend ID.
	BackendID string `json:"-"`
	// ServerIP: List of IP addresses for backend servers. Any other existing backend servers will be removed.
	ServerIP []string `json:"server_ip"`
}

type SubscribeToLBRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// SubscriberID: Subscriber ID.
	SubscriberID string `json:"subscriber_id"`
}

type UnsubscribeFromLBRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
}

type UpdateACLRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ACLID: ACL ID.
	ACLID string `json:"-"`
	// Name: ACL name.
	Name string `json:"name"`
	// Action: Action to take when incoming traffic matches an ACL filter.
	Action *ACLAction `json:"action"`
	// Match: ACL match filter object. One of `ip_subnet` or `http_filter` & `http_filter_value` are required.
	Match *ACLMatch `json:"match"`
	// Index: Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
	Index int32 `json:"index"`
	// Description: ACL description.
	Description *string `json:"description,omitempty"`
}

type UpdateBackendRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// BackendID: Backend ID.
	BackendID string `json:"-"`
	// Name: Backend name.
	Name string `json:"name"`
	// ForwardProtocol: Protocol to be used by the backend when forwarding traffic to backend servers.
	ForwardProtocol Protocol `json:"forward_protocol"`
	// ForwardPort: Port to be used by the backend when forwarding traffic to backend servers.
	ForwardPort int32 `json:"forward_port"`
	// ForwardPortAlgorithm: Load balancing algorithm to be used when determining which backend server to forward new traffic to.
	ForwardPortAlgorithm ForwardPortAlgorithm `json:"forward_port_algorithm"`
	// StickySessions: Defines whether to activate sticky sessions (binding a particular session to a particular backend server) and the method to use if so. None disables sticky sessions. Cookie-based uses an HTTP cookie to stick a session to a backend server. Table-based uses the source (client) IP address to stick a session to a backend server.
	StickySessions StickySessionsType `json:"sticky_sessions"`
	// StickySessionsCookieName: Cookie name for cookie-based sticky sessions.
	StickySessionsCookieName string `json:"sticky_sessions_cookie_name"`
	// SendProxyV2: Deprecated in favor of proxy_protocol field.
	SendProxyV2 *bool `json:"send_proxy_v2,omitempty"`
	// TimeoutServer: Maximum allowed time for a backend server to process a request.
	TimeoutServer *time.Duration `json:"timeout_server,omitempty"`
	// TimeoutConnect: Maximum allowed time for establishing a connection to a backend server.
	TimeoutConnect *time.Duration `json:"timeout_connect,omitempty"`
	// TimeoutTunnel: Maximum allowed tunnel inactivity time after Websocket is established (takes precedence over client and server timeout).
	TimeoutTunnel *time.Duration `json:"timeout_tunnel,omitempty"`
	// OnMarkedDownAction: Action to take when a backend server is marked as down.
	OnMarkedDownAction OnMarkedDownAction `json:"on_marked_down_action"`
	// ProxyProtocol: Protocol to use between the Load Balancer and backend servers. Allows the backend servers to be informed of the client's real IP address. The PROXY protocol must be supported by the backend servers' software.
	ProxyProtocol ProxyProtocol `json:"proxy_protocol"`
	// FailoverHost: Scaleway S3 bucket website to be served as failover if all backend servers are down, e.g. failover-website.s3-website.fr-par.scw.cloud.
	FailoverHost *string `json:"failover_host,omitempty"`
	// SslBridging: Defines whether to enable SSL bridging between the Load Balancer and backend servers.
	SslBridging *bool `json:"ssl_bridging,omitempty"`
	// IgnoreSslServerVerify: Defines whether the server certificate verification should be ignored.
	IgnoreSslServerVerify *bool `json:"ignore_ssl_server_verify,omitempty"`
	// RedispatchAttemptCount: Whether to use another backend server on each attempt.
	RedispatchAttemptCount *int32 `json:"redispatch_attempt_count,omitempty"`
	// MaxRetries: Number of retries when a backend server connection failed.
	MaxRetries *int32 `json:"max_retries,omitempty"`
	// MaxConnections: Maximum number of connections allowed per backend server.
	MaxConnections *int32 `json:"max_connections,omitempty"`
	// TimeoutQueue: Maximum time for a request to be left pending in queue when `max_connections` is reached.
	TimeoutQueue *scw.Duration `json:"timeout_queue,omitempty"`
}

func (m *UpdateBackendRequest) UnmarshalJSON(b []byte) error {
	type tmpType UpdateBackendRequest
	tmp := struct {
		tmpType
		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = UpdateBackendRequest(tmp.tmpType)
	m.TimeoutServer = tmp.TmpTimeoutServer.Standard()
	m.TimeoutConnect = tmp.TmpTimeoutConnect.Standard()
	m.TimeoutTunnel = tmp.TmpTimeoutTunnel.Standard()
	return nil
}

func (m UpdateBackendRequest) MarshalJSON() ([]byte, error) {
	type tmpType UpdateBackendRequest
	tmp := struct {
		tmpType
		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
	}{
		tmpType:           tmpType(m),
		TmpTimeoutServer:  marshaler.NewDuration(m.TimeoutServer),
		TmpTimeoutConnect: marshaler.NewDuration(m.TimeoutConnect),
		TmpTimeoutTunnel:  marshaler.NewDuration(m.TimeoutTunnel),
	}
	return json.Marshal(tmp)
}

type UpdateCertificateRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// CertificateID: Certificate ID.
	CertificateID string `json:"-"`
	// Name: Certificate name.
	Name string `json:"name"`
}

type UpdateFrontendRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// FrontendID: Frontend ID.
	FrontendID string `json:"-"`
	// Name: Frontend name.
	Name string `json:"name"`
	// InboundPort: Port the frontend should listen on.
	InboundPort int32 `json:"inbound_port"`
	// BackendID: Backend ID (ID of the backend the frontend should pass traffic to).
	BackendID string `json:"backend_id"`
	// TimeoutClient: Maximum allowed inactivity time on the client side.
	TimeoutClient *time.Duration `json:"timeout_client,omitempty"`
	// CertificateID: Certificate ID, deprecated in favor of certificate_ids array.
	CertificateID *string `json:"certificate_id,omitempty"`
	// CertificateIDs: List of SSL/TLS certificate IDs to bind to the frontend.
	CertificateIDs *[]string `json:"certificate_ids,omitempty"`
	// EnableHTTP3: Defines whether to enable HTTP/3 protocol on the frontend.
	EnableHTTP3 bool `json:"enable_http3"`
}

func (m *UpdateFrontendRequest) UnmarshalJSON(b []byte) error {
	type tmpType UpdateFrontendRequest
	tmp := struct {
		tmpType
		TmpTimeoutClient *marshaler.Duration `json:"timeout_client,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = UpdateFrontendRequest(tmp.tmpType)
	m.TimeoutClient = tmp.TmpTimeoutClient.Standard()
	return nil
}

func (m UpdateFrontendRequest) MarshalJSON() ([]byte, error) {
	type tmpType UpdateFrontendRequest
	tmp := struct {
		tmpType
		TmpTimeoutClient *marshaler.Duration `json:"timeout_client,omitempty"`
	}{
		tmpType:          tmpType(m),
		TmpTimeoutClient: marshaler.NewDuration(m.TimeoutClient),
	}
	return json.Marshal(tmp)
}

type UpdateHealthCheckRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Port: Port to use for the backend server health check.
	Port int32 `json:"port"`
	// CheckDelay: Time to wait between two consecutive health checks.
	CheckDelay *time.Duration `json:"check_delay,omitempty"`
	// CheckTimeout: Maximum time a backend server has to reply to the health check.
	CheckTimeout *time.Duration `json:"check_timeout,omitempty"`
	// CheckMaxRetries: Number of consecutive unsuccessful health checks after which the server will be considered dead.
	CheckMaxRetries int32 `json:"check_max_retries"`
	// BackendID: Backend ID.
	BackendID string `json:"-"`
	// CheckSendProxy: Defines whether proxy protocol should be activated for the health check.
	CheckSendProxy bool `json:"check_send_proxy"`
	// TCPConfig: Object to configure a basic TCP health check.
	TCPConfig *HealthCheckTCPConfig `json:"tcp_config,omitempty"`
	// MysqlConfig: Object to configure a MySQL health check. The check requires MySQL >=3.22, for older versions, use a TCP health check.
	MysqlConfig *HealthCheckMysqlConfig `json:"mysql_config,omitempty"`
	// PgsqlConfig: Object to configure a PostgreSQL health check.
	PgsqlConfig *HealthCheckPgsqlConfig `json:"pgsql_config,omitempty"`
	// LdapConfig: Object to configure an LDAP health check. The response is analyzed to find the LDAPv3 response message.
	LdapConfig *HealthCheckLdapConfig `json:"ldap_config,omitempty"`
	// RedisConfig: Object to configure a Redis health check. The response is analyzed to find the +PONG response message.
	RedisConfig *HealthCheckRedisConfig `json:"redis_config,omitempty"`
	// HTTPConfig: Object to configure an HTTP health check.
	HTTPConfig *HealthCheckHTTPConfig `json:"http_config,omitempty"`
	// HTTPSConfig: Object to configure an HTTPS health check.
	HTTPSConfig *HealthCheckHTTPSConfig `json:"https_config,omitempty"`
	// TransientCheckDelay: Time to wait between two consecutive health checks when a backend server is in a transient state (going UP or DOWN).
	TransientCheckDelay *scw.Duration `json:"transient_check_delay,omitempty"`
}

func (m *UpdateHealthCheckRequest) UnmarshalJSON(b []byte) error {
	type tmpType UpdateHealthCheckRequest
	tmp := struct {
		tmpType
		TmpCheckDelay   *marshaler.Duration `json:"check_delay,omitempty"`
		TmpCheckTimeout *marshaler.Duration `json:"check_timeout,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = UpdateHealthCheckRequest(tmp.tmpType)
	m.CheckDelay = tmp.TmpCheckDelay.Standard()
	m.CheckTimeout = tmp.TmpCheckTimeout.Standard()
	return nil
}

func (m UpdateHealthCheckRequest) MarshalJSON() ([]byte, error) {
	type tmpType UpdateHealthCheckRequest
	tmp := struct {
		tmpType
		TmpCheckDelay   *marshaler.Duration `json:"check_delay,omitempty"`
		TmpCheckTimeout *marshaler.Duration `json:"check_timeout,omitempty"`
	}{
		tmpType:         tmpType(m),
		TmpCheckDelay:   marshaler.NewDuration(m.CheckDelay),
		TmpCheckTimeout: marshaler.NewDuration(m.CheckTimeout),
	}
	return json.Marshal(tmp)
}

type UpdateIPRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// IPID: IP address ID.
	IPID string `json:"-"`
	// Reverse: Reverse DNS (domain name) for the IP address.
	Reverse *string `json:"reverse,omitempty"`
}

type UpdateLBRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// Name: Load Balancer name.
	Name string `json:"name"`
	// Description: Load Balancer description.
	Description string `json:"description"`
	// Tags: List of tags for the Load Balancer.
	Tags []string `json:"tags"`
	// SslCompatibilityLevel: Determines the minimal SSL version which needs to be supported on the client side, in an SSL/TLS offloading context. Intermediate is suitable for general-purpose servers with a variety of clients, recommended for almost all systems. Modern is suitable for services with clients that support TLS 1.3 and don't need backward compatibility. Old is compatible with a small number of very old clients and should be used only as a last resort.
	SslCompatibilityLevel SSLCompatibilityLevel `json:"ssl_compatibility_level"`
}

type UpdateRouteRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// RouteID: Route ID.
	RouteID string `json:"-"`
	// BackendID: ID of the target backend for the route.
	BackendID string `json:"backend_id"`
	// Match: Object defining the match condition for a route to be applied. If an incoming client session matches the specified condition (i.e. it has a matching SNI value or HTTP Host header value), it will be passed to the target backend.
	Match *RouteMatch `json:"match"`
}

type UpdateSubscriberRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SubscriberID: Subscriber ID.
	SubscriberID string `json:"-"`
	// Name: Subscriber name.
	Name string `json:"name"`
	// EmailConfig: Email address configuration.
	EmailConfig *SubscriberEmailConfig `json:"email_config,omitempty"`
	// WebhookConfig: Webhook URI configuration.
	WebhookConfig *SubscriberWebhookConfig `json:"webhook_config,omitempty"`
}

type ZonedAPIAddBackendServersRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// BackendID: Backend ID.
	BackendID string `json:"-"`
	// ServerIP: List of IP addresses to add to backend servers.
	ServerIP []string `json:"server_ip"`
}

type ZonedAPIAttachPrivateNetworkRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// PrivateNetworkID: Private Network ID.
	PrivateNetworkID string `json:"-"`
	// StaticConfig: Object containing an array of a local IP address for the Load Balancer on this Private Network.
	StaticConfig *PrivateNetworkStaticConfig `json:"static_config,omitempty"`
	// DHCPConfig: Defines whether to let DHCP assign IP addresses.
	DHCPConfig *PrivateNetworkDHCPConfig `json:"dhcp_config,omitempty"`
	// IpamConfig: For internal use only.
	IpamConfig *PrivateNetworkIpamConfig `json:"ipam_config,omitempty"`
}

// Add an ACL to a Load Balancer frontend.
type ZonedAPICreateACLRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// FrontendID: Frontend ID to attach the ACL to.
	FrontendID string `json:"-"`
	// Name: ACL name.
	Name string `json:"name"`
	// Action: Action to take when incoming traffic matches an ACL filter.
	Action *ACLAction `json:"action"`
	// Match: ACL match filter object. One of `ip_subnet` or `http_filter` & `http_filter_value` are required.
	Match *ACLMatch `json:"match"`
	// Index: Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
	Index int32 `json:"index"`
	// Description: ACL description.
	Description string `json:"description"`
}

type ZonedAPICreateBackendRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Name: Name for the backend.
	Name string `json:"name"`
	// ForwardProtocol: Protocol to be used by the backend when forwarding traffic to backend servers.
	ForwardProtocol Protocol `json:"forward_protocol"`
	// ForwardPort: Port to be used by the backend when forwarding traffic to backend servers.
	ForwardPort int32 `json:"forward_port"`
	// ForwardPortAlgorithm: Load balancing algorithm to be used when determining which backend server to forward new traffic to.
	ForwardPortAlgorithm ForwardPortAlgorithm `json:"forward_port_algorithm"`
	// StickySessions: Defines whether to activate sticky sessions (binding a particular session to a particular backend server) and the method to use if so. None disables sticky sessions. Cookie-based uses an HTTP cookie TO stick a session to a backend server. Table-based uses the source (client) IP address to stick a session to a backend server.
	StickySessions StickySessionsType `json:"sticky_sessions"`
	// StickySessionsCookieName: Cookie name for cookie-based sticky sessions.
	StickySessionsCookieName string `json:"sticky_sessions_cookie_name"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// HealthCheck: Object defining the health check to be carried out by the backend when checking the status and health of backend servers.
	HealthCheck *HealthCheck `json:"health_check"`
	// ServerIP: List of backend server IP addresses (IPv4 or IPv6) the backend should forward traffic to.
	ServerIP []string `json:"server_ip"`
	// SendProxyV2: Deprecated in favor of proxy_protocol field.
	SendProxyV2 *bool `json:"send_proxy_v2,omitempty"`
	// TimeoutServer: Maximum allowed time for a backend server to process a request.
	TimeoutServer *time.Duration `json:"timeout_server,omitempty"`
	// TimeoutConnect: Maximum allowed time for establishing a connection to a backend server.
	TimeoutConnect *time.Duration `json:"timeout_connect,omitempty"`
	// TimeoutTunnel: Maximum allowed tunnel inactivity time after Websocket is established (takes precedence over client and server timeout).
	TimeoutTunnel *time.Duration `json:"timeout_tunnel,omitempty"`
	// OnMarkedDownAction: Action to take when a backend server is marked as down.
	OnMarkedDownAction OnMarkedDownAction `json:"on_marked_down_action"`
	// ProxyProtocol: Protocol to use between the Load Balancer and backend servers. Allows the backend servers to be informed of the client's real IP address. The PROXY protocol must be supported by the backend servers' software.
	ProxyProtocol ProxyProtocol `json:"proxy_protocol"`
	// FailoverHost: Scaleway S3 bucket website to be served as failover if all backend servers are down, e.g. failover-website.s3-website.fr-par.scw.cloud.
	FailoverHost *string `json:"failover_host,omitempty"`
	// SslBridging: Defines whether to enable SSL bridging between the Load Balancer and backend servers.
	SslBridging *bool `json:"ssl_bridging,omitempty"`
	// IgnoreSslServerVerify: Defines whether the server certificate verification should be ignored.
	IgnoreSslServerVerify *bool `json:"ignore_ssl_server_verify,omitempty"`
	// RedispatchAttemptCount: Whether to use another backend server on each attempt.
	RedispatchAttemptCount *int32 `json:"redispatch_attempt_count,omitempty"`
	// MaxRetries: Number of retries when a backend server connection failed.
	MaxRetries *int32 `json:"max_retries,omitempty"`
	// MaxConnections: Maximum number of connections allowed per backend server.
	MaxConnections *int32 `json:"max_connections,omitempty"`
	// TimeoutQueue: Maximum time for a request to be left pending in queue when `max_connections` is reached.
	TimeoutQueue *scw.Duration `json:"timeout_queue,omitempty"`
}

func (m *ZonedAPICreateBackendRequest) UnmarshalJSON(b []byte) error {
	type tmpType ZonedAPICreateBackendRequest
	tmp := struct {
		tmpType
		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = ZonedAPICreateBackendRequest(tmp.tmpType)
	m.TimeoutServer = tmp.TmpTimeoutServer.Standard()
	m.TimeoutConnect = tmp.TmpTimeoutConnect.Standard()
	m.TimeoutTunnel = tmp.TmpTimeoutTunnel.Standard()
	return nil
}

func (m ZonedAPICreateBackendRequest) MarshalJSON() ([]byte, error) {
	type tmpType ZonedAPICreateBackendRequest
	tmp := struct {
		tmpType
		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
	}{
		tmpType:           tmpType(m),
		TmpTimeoutServer:  marshaler.NewDuration(m.TimeoutServer),
		TmpTimeoutConnect: marshaler.NewDuration(m.TimeoutConnect),
		TmpTimeoutTunnel:  marshaler.NewDuration(m.TimeoutTunnel),
	}
	return json.Marshal(tmp)
}

type ZonedAPICreateCertificateRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// Name: Name for the certificate.
	Name string `json:"name"`
	// Letsencrypt: Object to define a new Let's Encrypt certificate to be generated.
	Letsencrypt *CreateCertificateRequestLetsencryptConfig `json:"letsencrypt,omitempty"`
	// CustomCertificate: Object to define an existing custom certificate to be imported.
	CustomCertificate *CreateCertificateRequestCustomCertificate `json:"custom_certificate,omitempty"`
}

type ZonedAPICreateFrontendRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Name: Name for the frontend.
	Name string `json:"name"`
	// InboundPort: Port the frontend should listen on.
	InboundPort int32 `json:"inbound_port"`
	// LBID: Load Balancer ID (ID of the Load Balancer to attach the frontend to).
	LBID string `json:"-"`
	// BackendID: Backend ID (ID of the backend the frontend should pass traffic to).
	BackendID string `json:"backend_id"`
	// TimeoutClient: Maximum allowed inactivity time on the client side.
	TimeoutClient *time.Duration `json:"timeout_client,omitempty"`
	// CertificateID: Certificate ID, deprecated in favor of certificate_ids array.
	CertificateID *string `json:"certificate_id,omitempty"`
	// CertificateIDs: List of SSL/TLS certificate IDs to bind to the frontend.
	CertificateIDs *[]string `json:"certificate_ids,omitempty"`
	// EnableHTTP3: Defines whether to enable HTTP/3 protocol on the frontend.
	EnableHTTP3 bool `json:"enable_http3"`
}

func (m *ZonedAPICreateFrontendRequest) UnmarshalJSON(b []byte) error {
	type tmpType ZonedAPICreateFrontendRequest
	tmp := struct {
		tmpType
		TmpTimeoutClient *marshaler.Duration `json:"timeout_client,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = ZonedAPICreateFrontendRequest(tmp.tmpType)
	m.TimeoutClient = tmp.TmpTimeoutClient.Standard()
	return nil
}

func (m ZonedAPICreateFrontendRequest) MarshalJSON() ([]byte, error) {
	type tmpType ZonedAPICreateFrontendRequest
	tmp := struct {
		tmpType
		TmpTimeoutClient *marshaler.Duration `json:"timeout_client,omitempty"`
	}{
		tmpType:          tmpType(m),
		TmpTimeoutClient: marshaler.NewDuration(m.TimeoutClient),
	}
	return json.Marshal(tmp)
}

type ZonedAPICreateIPRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// OrganizationID: Organization ID of the Organization where the IP address should be created.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: Project ID of the Project where the IP address should be created.
	ProjectID *string `json:"project_id,omitempty"`
	// Reverse: Reverse DNS (domain name) for the IP address.
	Reverse *string `json:"reverse,omitempty"`
	// IsIPv6: If true, creates a Flexible IP with an ipv6 address.
	IsIPv6 bool `json:"is_ipv6"`
}

type ZonedAPICreateLBRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// OrganizationID: Scaleway Organization to create the Load Balancer in.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: Scaleway Project to create the Load Balancer in.
	ProjectID *string `json:"project_id,omitempty"`
	// Name: Name for the Load Balancer.
	Name string `json:"name"`
	// Description: Description for the Load Balancer.
	Description string `json:"description"`
	// IPID: ID of an existing flexible IP address to attach to the Load Balancer.
	IPID *string `json:"ip_id,omitempty"`
	// AssignFlexibleIP: Defines whether to automatically assign a flexible public IP to lb. Default value is `false` (do not assign).
	AssignFlexibleIP *bool `json:"assign_flexible_ip,omitempty"`
	// IPIDs: List of IP IDs to attach to the Load Balancer.
	IPIDs []string `json:"ip_ids"`
	// Tags: List of tags for the Load Balancer.
	Tags []string `json:"tags"`
	// Type: Load Balancer commercial offer type. Use the Load Balancer types endpoint to retrieve a list of available offer types.
	Type string `json:"type"`
	// SslCompatibilityLevel: Determines the minimal SSL version which needs to be supported on the client side, in an SSL/TLS offloading context. Intermediate is suitable for general-purpose servers with a variety of clients, recommended for almost all systems. Modern is suitable for services with clients that support TLS 1.3 and do not need backward compatibility. Old is compatible with a small number of very old clients and should be used only as a last resort.
	SslCompatibilityLevel SSLCompatibilityLevel `json:"ssl_compatibility_level"`
}

type ZonedAPICreateRouteRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// FrontendID: ID of the source frontend to create the route on.
	FrontendID string `json:"frontend_id"`
	// BackendID: ID of the target backend for the route.
	BackendID string `json:"backend_id"`
	// Match: Object defining the match condition for a route to be applied. If an incoming client session matches the specified condition (i.e. it has a matching SNI value or HTTP Host header value), it will be passed to the target backend.
	Match *RouteMatch `json:"match"`
}

// Create a new alert subscriber (webhook or email).
type ZonedAPICreateSubscriberRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Name: Subscriber name.
	Name string `json:"name"`
	// EmailConfig: Email address configuration.
	EmailConfig *SubscriberEmailConfig `json:"email_config,omitempty"`
	// WebhookConfig: WebHook URI configuration.
	WebhookConfig *SubscriberWebhookConfig `json:"webhook_config,omitempty"`
	// OrganizationID: Organization ID to create the subscriber in.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: Project ID to create the subscriber in.
	ProjectID *string `json:"project_id,omitempty"`
}

type ZonedAPIDeleteACLRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ACLID: ACL ID.
	ACLID string `json:"-"`
}

type ZonedAPIDeleteBackendRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// BackendID: ID of the backend to delete.
	BackendID string `json:"-"`
}

type ZonedAPIDeleteCertificateRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// CertificateID: Certificate ID.
	CertificateID string `json:"-"`
}

type ZonedAPIDeleteFrontendRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// FrontendID: ID of the frontend to delete.
	FrontendID string `json:"-"`
}

type ZonedAPIDeleteLBRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// LBID: ID of the Load Balancer to delete.
	LBID string `json:"-"`
	// ReleaseIP: Defines whether the Load Balancer's flexible IP should be deleted. Set to true to release the flexible IP, or false to keep it available in your account for future Load Balancers.
	ReleaseIP bool `json:"release_ip"`
}

type ZonedAPIDeleteRouteRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// RouteID: Route ID.
	RouteID string `json:"-"`
}

type ZonedAPIDeleteSubscriberRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// SubscriberID: Subscriber ID.
	SubscriberID string `json:"-"`
}

type ZonedAPIDetachPrivateNetworkRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// LBID: Load balancer ID.
	LBID string `json:"-"`
	// PrivateNetworkID: Set your instance private network id.
	PrivateNetworkID string `json:"-"`
}

type ZonedAPIGetACLRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ACLID: ACL ID.
	ACLID string `json:"-"`
}

type ZonedAPIGetBackendRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// BackendID: Backend ID.
	BackendID string `json:"-"`
}

type ZonedAPIGetCertificateRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// CertificateID: Certificate ID.
	CertificateID string `json:"-"`
}

type ZonedAPIGetFrontendRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// FrontendID: Frontend ID.
	FrontendID string `json:"-"`
}

type ZonedAPIGetIPRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// IPID: IP address ID.
	IPID string `json:"-"`
}

type ZonedAPIGetLBRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
}

// Get Load Balancer stats.
type ZonedAPIGetLBStatsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// BackendID: ID of the backend.
	BackendID *string `json:"backend_id,omitempty"`
}

type ZonedAPIGetRouteRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// RouteID: Route ID.
	RouteID string `json:"-"`
}

type ZonedAPIGetSubscriberRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// SubscriberID: Subscriber ID.
	SubscriberID string `json:"-"`
}

type ZonedAPIListACLsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// FrontendID: Frontend ID (ACLs attached to this frontend will be returned in the response).
	FrontendID string `json:"-"`
	// OrderBy: Sort order of ACLs in the response.
	OrderBy ListACLRequestOrderBy `json:"order_by"`
	// Page: The page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// PageSize: The number of ACLs to return.
	PageSize *uint32 `json:"page_size,omitempty"`
	// Name: ACL name to filter for.
	Name *string `json:"name,omitempty"`
}

type ZonedAPIListBackendStatsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// Page: The page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Number of items to return.
	PageSize *uint32 `json:"page_size,omitempty"`
	// BackendID: ID of the backend.
	BackendID *string `json:"backend_id,omitempty"`
}

type ZonedAPIListBackendsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// Name: Name of the backend to filter for.
	Name *string `json:"name,omitempty"`
	// OrderBy: Sort order of backends in the response.
	OrderBy ListBackendsRequestOrderBy `json:"order_by"`
	// Page: The page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Number of backends to return.
	PageSize *uint32 `json:"page_size,omitempty"`
}

type ZonedAPIListCertificatesRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// OrderBy: Sort order of certificates in the response.
	OrderBy ListCertificatesRequestOrderBy `json:"order_by"`
	// Page: The page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Number of certificates to return.
	PageSize *uint32 `json:"page_size,omitempty"`
	// Name: Certificate name to filter for, only certificates of this name will be returned.
	Name *string `json:"name,omitempty"`
}

type ZonedAPIListFrontendsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// Name: Name of the frontend to filter for.
	Name *string `json:"name,omitempty"`
	// OrderBy: Sort order of frontends in the response.
	OrderBy ListFrontendsRequestOrderBy `json:"order_by"`
	// Page: The page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Number of frontends to return.
	PageSize *uint32 `json:"page_size,omitempty"`
}

type ZonedAPIListIPsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Page: The page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Number of IP addresses to return.
	PageSize *uint32 `json:"page_size,omitempty"`
	// IPAddress: IP address to filter for.
	IPAddress *string `json:"ip_address,omitempty"`
	// OrganizationID: Organization ID to filter for, only Load Balancer IP addresses from this Organization will be returned.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: Project ID to filter for, only Load Balancer IP addresses from this Project will be returned.
	ProjectID *string `json:"project_id,omitempty"`
}

type ZonedAPIListLBPrivateNetworksRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// OrderBy: Sort order of Private Network objects in the response.
	OrderBy ListPrivateNetworksRequestOrderBy `json:"order_by"`
	// PageSize: Number of objects to return.
	PageSize *uint32 `json:"page_size,omitempty"`
	// Page: The page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
}

type ZonedAPIListLBTypesRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Page: The page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// PageSize: The number of items to return.
	PageSize *uint32 `json:"page_size,omitempty"`
}

type ZonedAPIListLBsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Name: Load Balancer name to filter for.
	Name *string `json:"name,omitempty"`
	// OrderBy: Sort order of Load Balancers in the response.
	OrderBy ListLBsRequestOrderBy `json:"order_by"`
	// PageSize: Number of Load Balancers to return.
	PageSize *uint32 `json:"page_size,omitempty"`
	// Page: Page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// OrganizationID: Organization ID to filter for, only Load Balancers from this Organization will be returned.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: Project ID to filter for, only Load Balancers from this Project will be returned.
	ProjectID *string `json:"project_id,omitempty"`
}

type ZonedAPIListRoutesRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// OrderBy: Sort order of routes in the response.
	OrderBy ListRoutesRequestOrderBy `json:"order_by"`
	// PageSize: The number of route objects to return.
	PageSize *uint32 `json:"page_size,omitempty"`
	// Page: The page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// FrontendID: Frontend ID to filter for, only Routes from this Frontend will be returned.
	FrontendID *string `json:"frontend_id,omitempty"`
}

type ZonedAPIListSubscriberRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// OrderBy: Sort order of subscribers in the response.
	OrderBy ListSubscriberRequestOrderBy `json:"order_by"`
	// Page: The page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// PageSize: The number of items to return.
	PageSize *uint32 `json:"page_size,omitempty"`
	// Name: Subscriber name to search for.
	Name *string `json:"name,omitempty"`
	// OrganizationID: Filter subscribers by Organization ID.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: Filter subscribers by Project ID.
	ProjectID *string `json:"project_id,omitempty"`
}

type ZonedAPIMigrateLBRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// Type: Load Balancer type to migrate to (use the List all Load Balancer offer types endpoint to get a list of available offer types).
	Type string `json:"type"`
}

type ZonedAPIReleaseIPRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// IPID: IP address ID.
	IPID string `json:"-"`
}

type ZonedAPIRemoveBackendServersRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// BackendID: Backend ID.
	BackendID string `json:"-"`
	// ServerIP: List of IP addresses to remove from backend servers.
	ServerIP []string `json:"server_ip"`
}

type ZonedAPISetACLsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ACLs: List of ACLs for this frontend. Any other existing ACLs on this frontend will be removed.
	ACLs []*ACLSpec `json:"acls"`
	// FrontendID: Frontend ID.
	FrontendID string `json:"-"`
}

type ZonedAPISetBackendServersRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// BackendID: Backend ID.
	BackendID string `json:"-"`
	// ServerIP: List of IP addresses for backend servers. Any other existing backend servers will be removed.
	ServerIP []string `json:"server_ip"`
}

type ZonedAPISubscribeToLBRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// SubscriberID: Subscriber ID.
	SubscriberID string `json:"subscriber_id"`
}

type ZonedAPIUnsubscribeFromLBRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
}

type ZonedAPIUpdateACLRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ACLID: ACL ID.
	ACLID string `json:"-"`
	// Name: ACL name.
	Name string `json:"name"`
	// Action: Action to take when incoming traffic matches an ACL filter.
	Action *ACLAction `json:"action"`
	// Match: ACL match filter object. One of `ip_subnet` or `http_filter` & `http_filter_value` are required.
	Match *ACLMatch `json:"match"`
	// Index: Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
	Index int32 `json:"index"`
	// Description: ACL description.
	Description *string `json:"description,omitempty"`
}

type ZonedAPIUpdateBackendRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// BackendID: Backend ID.
	BackendID string `json:"-"`
	// Name: Backend name.
	Name string `json:"name"`
	// ForwardProtocol: Protocol to be used by the backend when forwarding traffic to backend servers.
	ForwardProtocol Protocol `json:"forward_protocol"`
	// ForwardPort: Port to be used by the backend when forwarding traffic to backend servers.
	ForwardPort int32 `json:"forward_port"`
	// ForwardPortAlgorithm: Load balancing algorithm to be used when determining which backend server to forward new traffic to.
	ForwardPortAlgorithm ForwardPortAlgorithm `json:"forward_port_algorithm"`
	// StickySessions: Defines whether to activate sticky sessions (binding a particular session to a particular backend server) and the method to use if so. None disables sticky sessions. Cookie-based uses an HTTP cookie to stick a session to a backend server. Table-based uses the source (client) IP address to stick a session to a backend server.
	StickySessions StickySessionsType `json:"sticky_sessions"`
	// StickySessionsCookieName: Cookie name for cookie-based sticky sessions.
	StickySessionsCookieName string `json:"sticky_sessions_cookie_name"`
	// SendProxyV2: Deprecated in favor of proxy_protocol field.
	SendProxyV2 *bool `json:"send_proxy_v2,omitempty"`
	// TimeoutServer: Maximum allowed time for a backend server to process a request.
	TimeoutServer *time.Duration `json:"timeout_server,omitempty"`
	// TimeoutConnect: Maximum allowed time for establishing a connection to a backend server.
	TimeoutConnect *time.Duration `json:"timeout_connect,omitempty"`
	// TimeoutTunnel: Maximum allowed tunnel inactivity time after Websocket is established (takes precedence over client and server timeout).
	TimeoutTunnel *time.Duration `json:"timeout_tunnel,omitempty"`
	// OnMarkedDownAction: Action to take when a backend server is marked as down.
	OnMarkedDownAction OnMarkedDownAction `json:"on_marked_down_action"`
	// ProxyProtocol: Protocol to use between the Load Balancer and backend servers. Allows the backend servers to be informed of the client's real IP address. The PROXY protocol must be supported by the backend servers' software.
	ProxyProtocol ProxyProtocol `json:"proxy_protocol"`
	// FailoverHost: Scaleway S3 bucket website to be served as failover if all backend servers are down, e.g. failover-website.s3-website.fr-par.scw.cloud.
	FailoverHost *string `json:"failover_host,omitempty"`
	// SslBridging: Defines whether to enable SSL bridging between the Load Balancer and backend servers.
	SslBridging *bool `json:"ssl_bridging,omitempty"`
	// IgnoreSslServerVerify: Defines whether the server certificate verification should be ignored.
	IgnoreSslServerVerify *bool `json:"ignore_ssl_server_verify,omitempty"`
	// RedispatchAttemptCount: Whether to use another backend server on each attempt.
	RedispatchAttemptCount *int32 `json:"redispatch_attempt_count,omitempty"`
	// MaxRetries: Number of retries when a backend server connection failed.
	MaxRetries *int32 `json:"max_retries,omitempty"`
	// MaxConnections: Maximum number of connections allowed per backend server.
	MaxConnections *int32 `json:"max_connections,omitempty"`
	// TimeoutQueue: Maximum time for a request to be left pending in queue when `max_connections` is reached.
	TimeoutQueue *scw.Duration `json:"timeout_queue,omitempty"`
}

func (m *ZonedAPIUpdateBackendRequest) UnmarshalJSON(b []byte) error {
	type tmpType ZonedAPIUpdateBackendRequest
	tmp := struct {
		tmpType
		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = ZonedAPIUpdateBackendRequest(tmp.tmpType)
	m.TimeoutServer = tmp.TmpTimeoutServer.Standard()
	m.TimeoutConnect = tmp.TmpTimeoutConnect.Standard()
	m.TimeoutTunnel = tmp.TmpTimeoutTunnel.Standard()
	return nil
}

func (m ZonedAPIUpdateBackendRequest) MarshalJSON() ([]byte, error) {
	type tmpType ZonedAPIUpdateBackendRequest
	tmp := struct {
		tmpType
		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
	}{
		tmpType:           tmpType(m),
		TmpTimeoutServer:  marshaler.NewDuration(m.TimeoutServer),
		TmpTimeoutConnect: marshaler.NewDuration(m.TimeoutConnect),
		TmpTimeoutTunnel:  marshaler.NewDuration(m.TimeoutTunnel),
	}
	return json.Marshal(tmp)
}

type ZonedAPIUpdateCertificateRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// CertificateID: Certificate ID.
	CertificateID string `json:"-"`
	// Name: Certificate name.
	Name string `json:"name"`
}

type ZonedAPIUpdateFrontendRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// FrontendID: Frontend ID.
	FrontendID string `json:"-"`
	// Name: Frontend name.
	Name string `json:"name"`
	// InboundPort: Port the frontend should listen on.
	InboundPort int32 `json:"inbound_port"`
	// BackendID: Backend ID (ID of the backend the frontend should pass traffic to).
	BackendID string `json:"backend_id"`
	// TimeoutClient: Maximum allowed inactivity time on the client side.
	TimeoutClient *time.Duration `json:"timeout_client,omitempty"`
	// CertificateID: Certificate ID, deprecated in favor of certificate_ids array.
	CertificateID *string `json:"certificate_id,omitempty"`
	// CertificateIDs: List of SSL/TLS certificate IDs to bind to the frontend.
	CertificateIDs *[]string `json:"certificate_ids,omitempty"`
	// EnableHTTP3: Defines whether to enable HTTP/3 protocol on the frontend.
	EnableHTTP3 bool `json:"enable_http3"`
}

func (m *ZonedAPIUpdateFrontendRequest) UnmarshalJSON(b []byte) error {
	type tmpType ZonedAPIUpdateFrontendRequest
	tmp := struct {
		tmpType
		TmpTimeoutClient *marshaler.Duration `json:"timeout_client,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = ZonedAPIUpdateFrontendRequest(tmp.tmpType)
	m.TimeoutClient = tmp.TmpTimeoutClient.Standard()
	return nil
}

func (m ZonedAPIUpdateFrontendRequest) MarshalJSON() ([]byte, error) {
	type tmpType ZonedAPIUpdateFrontendRequest
	tmp := struct {
		tmpType
		TmpTimeoutClient *marshaler.Duration `json:"timeout_client,omitempty"`
	}{
		tmpType:          tmpType(m),
		TmpTimeoutClient: marshaler.NewDuration(m.TimeoutClient),
	}
	return json.Marshal(tmp)
}

type ZonedAPIUpdateHealthCheckRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Port: Port to use for the backend server health check.
	Port int32 `json:"port"`
	// CheckDelay: Time to wait between two consecutive health checks.
	CheckDelay *time.Duration `json:"check_delay,omitempty"`
	// CheckTimeout: Maximum time a backend server has to reply to the health check.
	CheckTimeout *time.Duration `json:"check_timeout,omitempty"`
	// CheckMaxRetries: Number of consecutive unsuccessful health checks after which the server will be considered dead.
	CheckMaxRetries int32 `json:"check_max_retries"`
	// BackendID: Backend ID.
	BackendID string `json:"-"`
	// CheckSendProxy: Defines whether proxy protocol should be activated for the health check.
	CheckSendProxy bool `json:"check_send_proxy"`
	// TCPConfig: Object to configure a basic TCP health check.
	TCPConfig *HealthCheckTCPConfig `json:"tcp_config,omitempty"`
	// MysqlConfig: Object to configure a MySQL health check. The check requires MySQL >=3.22, for older versions, use a TCP health check.
	MysqlConfig *HealthCheckMysqlConfig `json:"mysql_config,omitempty"`
	// PgsqlConfig: Object to configure a PostgreSQL health check.
	PgsqlConfig *HealthCheckPgsqlConfig `json:"pgsql_config,omitempty"`
	// LdapConfig: Object to configure an LDAP health check. The response is analyzed to find the LDAPv3 response message.
	LdapConfig *HealthCheckLdapConfig `json:"ldap_config,omitempty"`
	// RedisConfig: Object to configure a Redis health check. The response is analyzed to find the +PONG response message.
	RedisConfig *HealthCheckRedisConfig `json:"redis_config,omitempty"`
	// HTTPConfig: Object to configure an HTTP health check.
	HTTPConfig *HealthCheckHTTPConfig `json:"http_config,omitempty"`
	// HTTPSConfig: Object to configure an HTTPS health check.
	HTTPSConfig *HealthCheckHTTPSConfig `json:"https_config,omitempty"`
	// TransientCheckDelay: Time to wait between two consecutive health checks when a backend server is in a transient state (going UP or DOWN).
	TransientCheckDelay *scw.Duration `json:"transient_check_delay,omitempty"`
}

func (m *ZonedAPIUpdateHealthCheckRequest) UnmarshalJSON(b []byte) error {
	type tmpType ZonedAPIUpdateHealthCheckRequest
	tmp := struct {
		tmpType
		TmpCheckDelay   *marshaler.Duration `json:"check_delay,omitempty"`
		TmpCheckTimeout *marshaler.Duration `json:"check_timeout,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = ZonedAPIUpdateHealthCheckRequest(tmp.tmpType)
	m.CheckDelay = tmp.TmpCheckDelay.Standard()
	m.CheckTimeout = tmp.TmpCheckTimeout.Standard()
	return nil
}

func (m ZonedAPIUpdateHealthCheckRequest) MarshalJSON() ([]byte, error) {
	type tmpType ZonedAPIUpdateHealthCheckRequest
	tmp := struct {
		tmpType
		TmpCheckDelay   *marshaler.Duration `json:"check_delay,omitempty"`
		TmpCheckTimeout *marshaler.Duration `json:"check_timeout,omitempty"`
	}{
		tmpType:         tmpType(m),
		TmpCheckDelay:   marshaler.NewDuration(m.CheckDelay),
		TmpCheckTimeout: marshaler.NewDuration(m.CheckTimeout),
	}
	return json.Marshal(tmp)
}

type ZonedAPIUpdateIPRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// IPID: IP address ID.
	IPID string `json:"-"`
	// Reverse: Reverse DNS (domain name) for the IP address.
	Reverse *string `json:"reverse,omitempty"`
}

type ZonedAPIUpdateLBRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// LBID: Load Balancer ID.
	LBID string `json:"-"`
	// Name: Load Balancer name.
	Name string `json:"name"`
	// Description: Load Balancer description.
	Description string `json:"description"`
	// Tags: List of tags for the Load Balancer.
	Tags []string `json:"tags"`
	// SslCompatibilityLevel: Determines the minimal SSL version which needs to be supported on the client side, in an SSL/TLS offloading context. Intermediate is suitable for general-purpose servers with a variety of clients, recommended for almost all systems. Modern is suitable for services with clients that support TLS 1.3 and don't need backward compatibility. Old is compatible with a small number of very old clients and should be used only as a last resort.
	SslCompatibilityLevel SSLCompatibilityLevel `json:"ssl_compatibility_level"`
}

type ZonedAPIUpdateRouteRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// RouteID: Route ID.
	RouteID string `json:"-"`
	// BackendID: ID of the target backend for the route.
	BackendID string `json:"backend_id"`
	// Match: Object defining the match condition for a route to be applied. If an incoming client session matches the specified condition (i.e. it has a matching SNI value or HTTP Host header value), it will be passed to the target backend.
	Match *RouteMatch `json:"match"`
}

type ZonedAPIUpdateSubscriberRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// SubscriberID: Subscriber ID.
	SubscriberID string `json:"-"`
	// Name: Subscriber name.
	Name string `json:"name"`
	// EmailConfig: Email address configuration.
	EmailConfig *SubscriberEmailConfig `json:"email_config,omitempty"`
	// WebhookConfig: Webhook URI configuration.
	WebhookConfig *SubscriberWebhookConfig `json:"webhook_config,omitempty"`
}

// Load Balancers are highly available and fully managed [Instances](https://www.scaleway.com/en/docs/compute/instances/concepts/#instance) that facilitate the distribution of incoming traffic across multiple servers. Load Balancers allow you to scale your applications while ensuring their continuous availability, even in the event of heavy traffic. They are commonly used to improve the performance and reliability of websites, applications, databases and other services.
//
// A Scaleway Load Balancer has a **public IP address** which is always reachable: in the event of hardware failure it is rerouted to a backup Instance. The Load Balancer receives incoming traffic at this IP address, monitors the health and availability of its **backend servers** via **health checks**, and balances traffic load between all healthy and available backend servers.
//
// You can create as many **frontends** and **backends** for each Load Balancer as you wish, with frontends configured to listen on defined ports and forward traffic to specific backends. You can also add **certificates** to frontends to enable secure, encrypted connections and facilitate SSL bridging or offloading. Backends can be configured to use your **protocol** of choice (TCP or HTTP) to connect to their backend servers. Additional features such as Access Control Lists (ACLs) and routes allow you to further configure the flow of traffic through your Scaleway Load Balancer.
//
// (switchcolumn)
// (switchcolumn)
//
// ## Concepts
//
// Refer to our [dedicated concepts page](https://www.scaleway.com/en/docs/network/load-balancer/concepts/) to find definitions of all Load Balancer-related terminology.
//
// (switchcolumn)
// (switchcolumn)
//
// ## Quickstart
//
// **Requirements**:
// To perform the following steps, you must first ensure that:
// - You have a [Scaleway account](https://console.scaleway.com/)
// - You have created an [API key](https://www.scaleway.com/en/docs/identity-and-access-management/iam/how-to/create-api-keys/) and that the API key has sufficient [IAM permissions](https://www.scaleway.com/en/docs/identity-and-access-management/iam/reference-content/permission-sets/) to perform the actions described on this page
// - You have [installed `curl`](https://curl.se/download.html)
//
// 1. Configure your environment variables.
//
//	<Message type="note">
//	This is an optional step that seeks to simplify your usage of the Load Balancer API.
//	</Message>
//
//	```bash
//	export SCW_SECRET_KEY="<API secret key>"
//	export SCW_DEFAULT_ZONE="<Scaleway default Availability Zone>"
//	export SCW_PROJECT_ID="<Scaleway Project ID>"
//	```
//
// 2. Create a Load Balancer: run the following command to create a Load Balancer. You can customize the details in the payload (name, description, tags, etc) as you wish.
//
//	```bash
//	curl -X POST \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  -H "Content-Type: application/json" \
//	  "https://api.scaleway.com/lb/v1/zones/$SCW_DEFAULT_ZONE/lbs" \
//	  -d '{
//	    "name":"API Test LB",
//	    "description": "my new Load Balancer",
//	    "project_id":"'"$SCW_PROJECT_ID"'",
//	    "tags":["test","another tag"]
//	  }'
//	```
//
// 3. **Get a list of your Load Balancers**: run the following command to get a list of all the Load Balancers in your account, with their details:
//
//	```bash
//	curl -X GET \
//	  -H "Content-Type: application/json" \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  "https://api.scaleway.com/lb/v1/zones/$SCW_DEFAULT_ZONE/lbs"
//	```
//
// 4. **Create a backend for your Load Balancer**: run the following command to create a backend for a specified Load Balancer. Ensure that you replace `<LOAD-BALANCER-ID>` in the URL with the ID of the Load Balancer you want to create a backend for. You can customize the configuration of the backend according to your needs: use the information below to adjust the payload as necessary.
//
//	```bash
//	curl -X POST \
//	  -H "Content-Type: application/json" \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  "https://api.scaleway.com/lb/v1/zones/$SCW_DEFAULT_ZONE/lbs/<LOAD-BALANCER-ID>/backends" \
//	  -d '{
//	    "name":"main backend",
//	    "forward_port": 80,
//	    "forward_port_algorithm": "roundrobin",
//	    "forward_protocol": "tcp",
//	    "health_check":{
//	        "check_delay": 2000,
//	        "check_max_retries": 3,
//	        "check_timeout": 1000,
//	        "port": 80,
//	        "tcp_config":{}
//	    },
//	    "server_ip": ["184.224.68.187", "139.7.243.56"]
//	  }'
//	```
//
//	  **Payload values**:
//
//	  * **name** (string): A name for the backend, e.g. `"main backend"`
//	  * **forward_port** (number): The port to use when connecting to a backend server to forward a user session, e.g. `8080`
//	  * **forward_port_algorithm** (string): The forwarding algorithm to use when determining which backend server to forward a user session to. Must be one of the following:
//	      * `"roundrobin"`: New sessions are forwarded to each backend server in turn.
//	      * `"leastconn"`: New sessions are forwarded to the backend server with the fewest current active sessions.
//	      * `"first"`: New sessions are forwarded to the first backend server found .
//	  * **forward_protocol** (string): The protocol to use when connecting to a backend server to forward a user session. Must be one of the following:
//	      * `"tcp"`: Transmission Control Protocol
//	      * `"http"`: Hypertext Transfer Protocol
//	  * **health_check** (object): The definition of the health check to use to check that backend servers are available and able to receive forwarded user sessions. Must contain the following parameters:
//	      * `"check_delay"`: The time between two consecutive health checks (in milliseconds)
//	      * `"check_max_retries"`: The number of consecutive unsuccessful health checks, after which the server will be considered dead
//	      * `"check_timeout:"` The maximum time a backend server has to reply to the health check (in milliseconds)
//	      * `"port"`: The port to use when connecting to a backend server for a health check
//	      * `"<type>_config"`: The health check type to use. This parameter name must be one of `"mysqlconfig"`, `"ldap_config"`, `"redis_config"`, `"tcp_config"`, `"pgsql_config"`, `"http_config"` or `"https_config"`. The parameter value may be an empty object, or an object containing further parameters, depending on the health check configuration type selected. See the full documentation below on health checks for further details.
//	          * Example of valid `health_check` value `{"check_delay":2000,"check_max_retries":3,"check_timeout":1000,"port":80,"tcp_config":{}}`
//	  * **server_ip** (array): The list of IPv4 or IPv6 addresses of the backend servers to forward user sessions to, e.g. `["184.224.68.187", "139.7.243.56"]`
//
// 5. **Create a frontend for your Load Balancer**: run the following command to create a frontend for a specified Load Balancer. Ensure that you replace `<LOAD-BALANCER-ID>` in the URL with the ID of the Load Balancer you want to create a frontend for. You can customize the configuration according to your needs: use the information below to adjust the payload as necessary.
//
//	```bash
//	curl -X POST \
//	  -H "Content-Type: application/json" \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  "https://api.scaleway.com/lb/v1/zones/$SCW_DEFAULT_ZONE/lbs/<LOAD-BALANCER-ID>/frontends" \
//	  -d '{
//	    "name": "main frontend",
//	    "backend_id": "6920166d-8665-426e-85a6-f137e7e71e7f",
//	    "inbound_port": 80,
//	    "timeout_client": 5000
//	  }'
//	```
//
//	**Payload values**:
//
//	* **name** (string): A name for the frontend, e.g. `"main frontend"`
//	* **backend_id** (number): The ID of the backend to attach to the frontend, e.g. `6920166d-8665-426e-85a6-f137e7e71e7f`
//	* **inbound_port** (number): The port that the frontend should listen on for incoming connections, e.g. `80`
//	* **timeout_client** (number): The maximum amount of inactivity time, in milliseconds, the frontend should allow before closing the connection, e.g. `5000`
//
// 6. **Delete your Load Balancer**: run the following command to delete a Load Balancer. Ensure that you replace `<LOAD-BALANCER-ID>` in the URL with the ID of the Load Balancer you want to delete. You can customize the payload to specify whether you want to release (delete) the [Flexible IP](https://www.scaleway.com/en/docs/network/load-balancer/concepts/#flexible-ip-address) associated with the Load Balancer or not.
//
//	```bash
//	curl -X DELETE \
//	  -H "Content-Type: application/json" \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  "https://api.scaleway.com/lb/v1/zones/$SCW_DEFAULT_ZONE/lbs/<LOAD-BALANCER-ID>" \
//	  -d '{
//	    "release_ip": false
//	  }'
//	```
//
// (switchcolumn)
// <Message type="requirement">
// - You have a [Scaleway account](https://console.scaleway.com/)
// - You have created an [API key](https://www.scaleway.com/en/docs/identity-and-access-management/iam/how-to/create-api-keys/) and that the API key has sufficient [IAM permissions](https://www.scaleway.com/en/docs/identity-and-access-management/iam/reference-content/permission-sets/) to perform the actions described on this page
// - You have [installed `curl`](https://curl.se/download.html)
// </Message>
// (switchcolumn)
//
// ## Technical information
//
// ### Availability Zones
//
// Load Balancers can be deployed in the following Availability Zones:
//
// | Name      | API ID                |
// |-----------|-----------------------|
// | Paris     | `fr-par-1` `fr-par-2` |
// | Amsterdam | `nl-ams-1` `nl-ams-2` |
// | Warsaw    | `pl-waw-1` `pl-waw-2` |
//
// The Scaleway Load Balancer API is a **zoned** API, meaning that each call must specify in its path parameters the Availability Zone for the resources concerned by the call.
//
// ## Going further
//
// For more help using Scaleway Load Balancers, check out the following resources:
// - Our [main documentation](https://www.scaleway.com/en/docs/network/load-balancer/)
// - The #load-balancer channel on our [Slack Community](https://www.scaleway.com/en/docs/tutorials/scaleway-slack-community/)
// - Our [support ticketing system](https://www.scaleway.com/en/docs/console/my-account/how-to/open-a-support-ticket/).
type ZonedAPI struct {
	client *scw.Client
}

// NewZonedAPI returns a ZonedAPI object from a Scaleway client.
func NewZonedAPI(client *scw.Client) *ZonedAPI {
	return &ZonedAPI{
		client: client,
	}
}
func (s *ZonedAPI) Zones() []scw.Zone {
	return []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2}
}

// ListLBs: List all Load Balancers in the specified zone, for a Scaleway Organization or Scaleway Project. By default, the Load Balancers returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *ZonedAPI) ListLBs(req *ZonedAPIListLBsRequest, opts ...scw.RequestOption) (*ListLBsResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/lbs",
		Query:  query,
	}

	var resp ListLBsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateLB: Create a new Load Balancer. Note that the Load Balancer will be created without frontends or backends; these must be created separately via the dedicated endpoints.
func (s *ZonedAPI) CreateLB(req *ZonedAPICreateLBRequest, opts ...scw.RequestOption) (*LB, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
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
		req.Name = namegenerator.GetRandomName("lb")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/lbs",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp LB

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLB: Retrieve information about an existing Load Balancer, specified by its Load Balancer ID. Its full details, including name, status and IP address, are returned in the response object.
func (s *ZonedAPI) GetLB(req *ZonedAPIGetLBRequest, opts ...scw.RequestOption) (*LB, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/lbs/" + fmt.Sprint(req.LBID) + "",
	}

	var resp LB

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateLB: Update the parameters of an existing Load Balancer, specified by its Load Balancer ID. Note that the request type is PUT and not PATCH. You must set all parameters.
func (s *ZonedAPI) UpdateLB(req *ZonedAPIUpdateLBRequest, opts ...scw.RequestOption) (*LB, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/lbs/" + fmt.Sprint(req.LBID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp LB

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteLB: Delete an existing Load Balancer, specified by its Load Balancer ID. Deleting a Load Balancer is permanent, and cannot be undone. The Load Balancer's flexible IP address can either be deleted with the Load Balancer, or kept in your account for future use.
func (s *ZonedAPI) DeleteLB(req *ZonedAPIDeleteLBRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "release_ip", req.ReleaseIP)

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/lbs/" + fmt.Sprint(req.LBID) + "",
		Query:  query,
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// MigrateLB: Migrate an existing Load Balancer from one commercial type to another. Allows you to scale your Load Balancer up or down in terms of bandwidth or multi-cloud provision.
func (s *ZonedAPI) MigrateLB(req *ZonedAPIMigrateLBRequest, opts ...scw.RequestOption) (*LB, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/lbs/" + fmt.Sprint(req.LBID) + "/migrate",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp LB

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListIPs: List the Load Balancer flexible IP addresses held in the account (filtered by Organization ID or Project ID). It is also possible to search for a specific IP address.
func (s *ZonedAPI) ListIPs(req *ZonedAPIListIPsRequest, opts ...scw.RequestOption) (*ListIPsResponse, error) {
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "ip_address", req.IPAddress)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/ips",
		Query:  query,
	}

	var resp ListIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateIP: Create a new Load Balancer flexible IP address, in the specified Scaleway Project. This can be attached to new Load Balancers created in the future.
func (s *ZonedAPI) CreateIP(req *ZonedAPICreateIPRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.OrganizationID == nil && req.ProjectID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.OrganizationID == nil && req.ProjectID == nil {
		req.ProjectID = &defaultProjectID
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/ips",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIP: Retrieve the full details of a Load Balancer flexible IP address.
func (s *ZonedAPI) GetIP(req *ZonedAPIGetIPRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IPID) + "",
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReleaseIP: Delete a Load Balancer flexible IP address. This action is irreversible, and cannot be undone.
func (s *ZonedAPI) ReleaseIP(req *ZonedAPIReleaseIPRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IPID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdateIP: Update the reverse DNS of a Load Balancer flexible IP address.
func (s *ZonedAPI) UpdateIP(req *ZonedAPIUpdateIPRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IPID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListBackends: List all the backends of a Load Balancer, specified by its Load Balancer ID. By default, results are returned in ascending order by the creation date of each backend. The response is an array of backend objects, containing full details of each one including their configuration parameters such as protocol, port and forwarding algorithm.
func (s *ZonedAPI) ListBackends(req *ZonedAPIListBackendsRequest, opts ...scw.RequestOption) (*ListBackendsResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/lbs/" + fmt.Sprint(req.LBID) + "/backends",
		Query:  query,
	}

	var resp ListBackendsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateBackend: Create a new backend for a given Load Balancer, specifying its full configuration including protocol, port and forwarding algorithm.
func (s *ZonedAPI) CreateBackend(req *ZonedAPICreateBackendRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("lbb")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/lbs/" + fmt.Sprint(req.LBID) + "/backends",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Backend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBackend: Get the full details of a given backend, specified by its backend ID. The response contains the backend's full configuration parameters including protocol, port and forwarding algorithm.
func (s *ZonedAPI) GetBackend(req *ZonedAPIGetBackendRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.BackendID) == "" {
		return nil, errors.New("field BackendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/backends/" + fmt.Sprint(req.BackendID) + "",
	}

	var resp Backend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateBackend: Update a backend of a given Load Balancer, specified by its backend ID. Note that the request type is PUT and not PATCH. You must set all parameters.
func (s *ZonedAPI) UpdateBackend(req *ZonedAPIUpdateBackendRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.BackendID) == "" {
		return nil, errors.New("field BackendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/backends/" + fmt.Sprint(req.BackendID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Backend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteBackend: Delete a backend of a given Load Balancer, specified by its backend ID. This action is irreversible and cannot be undone.
func (s *ZonedAPI) DeleteBackend(req *ZonedAPIDeleteBackendRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.BackendID) == "" {
		return errors.New("field BackendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/backends/" + fmt.Sprint(req.BackendID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// AddBackendServers: For a given backend specified by its backend ID, add a set of backend servers (identified by their IP addresses) it should forward traffic to. These will be appended to any existing set of backend servers for this backend.
func (s *ZonedAPI) AddBackendServers(req *ZonedAPIAddBackendServersRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.BackendID) == "" {
		return nil, errors.New("field BackendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/backends/" + fmt.Sprint(req.BackendID) + "/servers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Backend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveBackendServers: For a given backend specified by its backend ID, remove the specified backend servers (identified by their IP addresses) so that it no longer forwards traffic to them.
func (s *ZonedAPI) RemoveBackendServers(req *ZonedAPIRemoveBackendServersRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.BackendID) == "" {
		return nil, errors.New("field BackendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/backends/" + fmt.Sprint(req.BackendID) + "/servers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Backend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetBackendServers: For a given backend specified by its backend ID, define the set of backend servers (identified by their IP addresses) that it should forward traffic to. Any existing backend servers configured for this backend will be removed.
func (s *ZonedAPI) SetBackendServers(req *ZonedAPISetBackendServersRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.BackendID) == "" {
		return nil, errors.New("field BackendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/backends/" + fmt.Sprint(req.BackendID) + "/servers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Backend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHealthCheck: Update the configuration of the health check performed by a given backend to verify the health of its backend servers, identified by its backend ID. Note that the request type is PUT and not PATCH. You must set all parameters.
func (s *ZonedAPI) UpdateHealthCheck(req *ZonedAPIUpdateHealthCheckRequest, opts ...scw.RequestOption) (*HealthCheck, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.BackendID) == "" {
		return nil, errors.New("field BackendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/backends/" + fmt.Sprint(req.BackendID) + "/healthcheck",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp HealthCheck

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListFrontends: List all the frontends of a Load Balancer, specified by its Load Balancer ID. By default, results are returned in ascending order by the creation date of each frontend. The response is an array of frontend objects, containing full details of each one including the port they listen on and the backend they are attached to.
func (s *ZonedAPI) ListFrontends(req *ZonedAPIListFrontendsRequest, opts ...scw.RequestOption) (*ListFrontendsResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/lbs/" + fmt.Sprint(req.LBID) + "/frontends",
		Query:  query,
	}

	var resp ListFrontendsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFrontend: Create a new frontend for a given Load Balancer, specifying its configuration including the port it should listen on and the backend to attach it to.
func (s *ZonedAPI) CreateFrontend(req *ZonedAPICreateFrontendRequest, opts ...scw.RequestOption) (*Frontend, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("lbf")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/lbs/" + fmt.Sprint(req.LBID) + "/frontends",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Frontend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFrontend: Get the full details of a given frontend, specified by its frontend ID. The response contains the frontend's full configuration parameters including the backend it is attached to, the port it listens on, and any certificates it has.
func (s *ZonedAPI) GetFrontend(req *ZonedAPIGetFrontendRequest, opts ...scw.RequestOption) (*Frontend, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.FrontendID) == "" {
		return nil, errors.New("field FrontendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/frontends/" + fmt.Sprint(req.FrontendID) + "",
	}

	var resp Frontend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateFrontend: Update a given frontend, specified by its frontend ID. You can update configuration parameters including its name and the port it listens on. Note that the request type is PUT and not PATCH. You must set all parameters.
func (s *ZonedAPI) UpdateFrontend(req *ZonedAPIUpdateFrontendRequest, opts ...scw.RequestOption) (*Frontend, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.FrontendID) == "" {
		return nil, errors.New("field FrontendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/frontends/" + fmt.Sprint(req.FrontendID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Frontend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteFrontend: Delete a given frontend, specified by its frontend ID. This action is irreversible and cannot be undone.
func (s *ZonedAPI) DeleteFrontend(req *ZonedAPIDeleteFrontendRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.FrontendID) == "" {
		return errors.New("field FrontendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/frontends/" + fmt.Sprint(req.FrontendID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListRoutes: List all routes for a given frontend. The response is an array of routes, each one  with a specified backend to direct to if a certain condition is matched (based on the value of the SNI field or HTTP Host header).
func (s *ZonedAPI) ListRoutes(req *ZonedAPIListRoutesRequest, opts ...scw.RequestOption) (*ListRoutesResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "frontend_id", req.FrontendID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/routes",
		Query:  query,
	}

	var resp ListRoutesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRoute: Create a new route on a given frontend. To configure a route, specify the backend to direct to if a certain condition is matched (based on the value of the SNI field or HTTP Host header).
func (s *ZonedAPI) CreateRoute(req *ZonedAPICreateRouteRequest, opts ...scw.RequestOption) (*Route, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/routes",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Route

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRoute: Retrieve information about an existing route, specified by its route ID. Its full details, origin frontend, target backend and match condition, are returned in the response object.
func (s *ZonedAPI) GetRoute(req *ZonedAPIGetRouteRequest, opts ...scw.RequestOption) (*Route, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.RouteID) == "" {
		return nil, errors.New("field RouteID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/routes/" + fmt.Sprint(req.RouteID) + "",
	}

	var resp Route

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateRoute: Update the configuration of an existing route, specified by its route ID.
func (s *ZonedAPI) UpdateRoute(req *ZonedAPIUpdateRouteRequest, opts ...scw.RequestOption) (*Route, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.RouteID) == "" {
		return nil, errors.New("field RouteID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/routes/" + fmt.Sprint(req.RouteID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Route

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteRoute: Delete an existing route, specified by its route ID. Deleting a route is permanent, and cannot be undone.
func (s *ZonedAPI) DeleteRoute(req *ZonedAPIDeleteRouteRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.RouteID) == "" {
		return errors.New("field RouteID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/routes/" + fmt.Sprint(req.RouteID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetLBStats: Get usage statistics of a given Load Balancer.
func (s *ZonedAPI) GetLBStats(req *ZonedAPIGetLBStatsRequest, opts ...scw.RequestOption) (*LBStats, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "backend_id", req.BackendID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/lbs/" + fmt.Sprint(req.LBID) + "/stats",
		Query:  query,
	}

	var resp LBStats

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListBackendStats: List information about your backend servers, including their state and the result of their last health check.
func (s *ZonedAPI) ListBackendStats(req *ZonedAPIListBackendStatsRequest, opts ...scw.RequestOption) (*ListBackendStatsResponse, error) {
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "backend_id", req.BackendID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/lbs/" + fmt.Sprint(req.LBID) + "/backend-stats",
		Query:  query,
	}

	var resp ListBackendStatsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListACLs: List the ACLs for a given frontend, specified by its frontend ID. The response is an array of ACL objects, each one representing an ACL that denies or allows traffic based on certain conditions.
func (s *ZonedAPI) ListACLs(req *ZonedAPIListACLsRequest, opts ...scw.RequestOption) (*ListACLResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.FrontendID) == "" {
		return nil, errors.New("field FrontendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/frontends/" + fmt.Sprint(req.FrontendID) + "/acls",
		Query:  query,
	}

	var resp ListACLResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateACL: Create a new ACL for a given frontend. Each ACL must have a name, an action to perform (allow or deny), and a match rule (the action is carried out when the incoming traffic matches the rule).
func (s *ZonedAPI) CreateACL(req *ZonedAPICreateACLRequest, opts ...scw.RequestOption) (*ACL, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("acl")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.FrontendID) == "" {
		return nil, errors.New("field FrontendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/frontends/" + fmt.Sprint(req.FrontendID) + "/acls",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ACL

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetACL: Get information for a particular ACL, specified by its ACL ID. The response returns full details of the ACL, including its name, action, match rule and frontend.
func (s *ZonedAPI) GetACL(req *ZonedAPIGetACLRequest, opts ...scw.RequestOption) (*ACL, error) {
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
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/acls/" + fmt.Sprint(req.ACLID) + "",
	}

	var resp ACL

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateACL: Update a particular ACL, specified by its ACL ID. You can update details including its name, action and match rule.
func (s *ZonedAPI) UpdateACL(req *ZonedAPIUpdateACLRequest, opts ...scw.RequestOption) (*ACL, error) {
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
		Method: "PUT",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/acls/" + fmt.Sprint(req.ACLID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ACL

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteACL: Delete an ACL, specified by its ACL ID. Deleting an ACL is irreversible and cannot be undone.
func (s *ZonedAPI) DeleteACL(req *ZonedAPIDeleteACLRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ACLID) == "" {
		return errors.New("field ACLID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/acls/" + fmt.Sprint(req.ACLID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// SetACLs: For a given frontend specified by its frontend ID, define and add the complete set of ACLS for that frontend. Any existing ACLs on this frontend will be removed.
func (s *ZonedAPI) SetACLs(req *ZonedAPISetACLsRequest, opts ...scw.RequestOption) (*SetACLsResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.FrontendID) == "" {
		return nil, errors.New("field FrontendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/frontends/" + fmt.Sprint(req.FrontendID) + "/acls",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetACLsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateCertificate: Generate a new SSL/TLS certificate for a given Load Balancer. You can choose to create a Let's Encrypt certificate, or import a custom certificate.
func (s *ZonedAPI) CreateCertificate(req *ZonedAPICreateCertificateRequest, opts ...scw.RequestOption) (*Certificate, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("certificate")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/lbs/" + fmt.Sprint(req.LBID) + "/certificates",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Certificate

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListCertificates: List all the SSL/TLS certificates on a given Load Balancer. The response is an array of certificate objects, which are by default listed in ascending order of creation date.
func (s *ZonedAPI) ListCertificates(req *ZonedAPIListCertificatesRequest, opts ...scw.RequestOption) (*ListCertificatesResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/lbs/" + fmt.Sprint(req.LBID) + "/certificates",
		Query:  query,
	}

	var resp ListCertificatesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCertificate: Get information for a particular SSL/TLS certificate, specified by its certificate ID. The response returns full details of the certificate, including its type, main domain name, and alternative domain names.
func (s *ZonedAPI) GetCertificate(req *ZonedAPIGetCertificateRequest, opts ...scw.RequestOption) (*Certificate, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.CertificateID) == "" {
		return nil, errors.New("field CertificateID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/certificates/" + fmt.Sprint(req.CertificateID) + "",
	}

	var resp Certificate

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateCertificate: Update the name of a particular SSL/TLS certificate, specified by its certificate ID.
func (s *ZonedAPI) UpdateCertificate(req *ZonedAPIUpdateCertificateRequest, opts ...scw.RequestOption) (*Certificate, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.CertificateID) == "" {
		return nil, errors.New("field CertificateID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/certificates/" + fmt.Sprint(req.CertificateID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Certificate

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteCertificate: Delete an SSL/TLS certificate, specified by its certificate ID. Deleting a certificate is irreversible and cannot be undone.
func (s *ZonedAPI) DeleteCertificate(req *ZonedAPIDeleteCertificateRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.CertificateID) == "" {
		return errors.New("field CertificateID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/certificates/" + fmt.Sprint(req.CertificateID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListLBTypes: List all the different commercial Load Balancer types. The response includes an array of offer types, each with a name, description, and information about its stock availability.
func (s *ZonedAPI) ListLBTypes(req *ZonedAPIListLBTypesRequest, opts ...scw.RequestOption) (*ListLBTypesResponse, error) {
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/lb-types",
		Query:  query,
	}

	var resp ListLBTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSubscriber: Create a new subscriber, either with an email configuration or a webhook configuration, for a specified Scaleway Project.
func (s *ZonedAPI) CreateSubscriber(req *ZonedAPICreateSubscriberRequest, opts ...scw.RequestOption) (*Subscriber, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.OrganizationID == nil && req.ProjectID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.OrganizationID == nil && req.ProjectID == nil {
		req.ProjectID = &defaultProjectID
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/subscribers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Subscriber

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSubscriber: Retrieve information about an existing subscriber, specified by its subscriber ID. Its full details, including name and email/webhook configuration, are returned in the response object.
func (s *ZonedAPI) GetSubscriber(req *ZonedAPIGetSubscriberRequest, opts ...scw.RequestOption) (*Subscriber, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SubscriberID) == "" {
		return nil, errors.New("field SubscriberID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/subscribers/" + fmt.Sprint(req.SubscriberID) + "",
	}

	var resp Subscriber

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSubscriber: List all subscribers to Load Balancer alerts. By default, returns all subscribers to Load Balancer alerts for the Organization associated with the authentication token used for the request.
func (s *ZonedAPI) ListSubscriber(req *ZonedAPIListSubscriberRequest, opts ...scw.RequestOption) (*ListSubscriberResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/subscribers",
		Query:  query,
	}

	var resp ListSubscriberResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSubscriber: Update the parameters of a given subscriber (e.g. name, webhook configuration, email configuration), specified by its subscriber ID.
func (s *ZonedAPI) UpdateSubscriber(req *ZonedAPIUpdateSubscriberRequest, opts ...scw.RequestOption) (*Subscriber, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SubscriberID) == "" {
		return nil, errors.New("field SubscriberID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/subscribers/" + fmt.Sprint(req.SubscriberID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Subscriber

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSubscriber: Delete an existing subscriber, specified by its subscriber ID. Deleting a subscriber is permanent, and cannot be undone.
func (s *ZonedAPI) DeleteSubscriber(req *ZonedAPIDeleteSubscriberRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SubscriberID) == "" {
		return errors.New("field SubscriberID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/lb/subscription/" + fmt.Sprint(req.SubscriberID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// SubscribeToLB: Subscribe an existing subscriber to alerts for a given Load Balancer.
func (s *ZonedAPI) SubscribeToLB(req *ZonedAPISubscribeToLBRequest, opts ...scw.RequestOption) (*LB, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/lb/" + fmt.Sprint(req.LBID) + "/subscribe",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp LB

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsubscribeFromLB: Unsubscribe a subscriber from alerts for a given Load Balancer. The subscriber is not deleted, and can be resubscribed in the future if necessary.
func (s *ZonedAPI) UnsubscribeFromLB(req *ZonedAPIUnsubscribeFromLBRequest, opts ...scw.RequestOption) (*LB, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/lb/" + fmt.Sprint(req.LBID) + "/unsubscribe",
	}

	var resp LB

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListLBPrivateNetworks: List the Private Networks attached to a given Load Balancer, specified by its Load Balancer ID. The response is an array of Private Network objects, giving information including the status, configuration, name and creation date of each Private Network.
func (s *ZonedAPI) ListLBPrivateNetworks(req *ZonedAPIListLBPrivateNetworksRequest, opts ...scw.RequestOption) (*ListLBPrivateNetworksResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/lbs/" + fmt.Sprint(req.LBID) + "/private-networks",
		Query:  query,
	}

	var resp ListLBPrivateNetworksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachPrivateNetwork: Attach a specified Load Balancer to a specified Private Network, defining a static or DHCP configuration for the Load Balancer on the network.
func (s *ZonedAPI) AttachPrivateNetwork(req *ZonedAPIAttachPrivateNetworkRequest, opts ...scw.RequestOption) (*PrivateNetwork, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return nil, errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/lbs/" + fmt.Sprint(req.LBID) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "/attach",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PrivateNetwork

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachPrivateNetwork: Detach a specified Load Balancer from a specified Private Network.
func (s *ZonedAPI) DetachPrivateNetwork(req *ZonedAPIDetachPrivateNetworkRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return errors.New("field LBID cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/zones/" + fmt.Sprint(req.Zone) + "/lbs/" + fmt.Sprint(req.LBID) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "/detach",
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

// This API allows you to manage your load balancer service.
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

// ListLBs: List load balancers.
func (s *API) ListLBs(req *ListLBsRequest, opts ...scw.RequestOption) (*ListLBsResponse, error) {
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
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs",
		Query:  query,
	}

	var resp ListLBsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateLB: Create a load balancer.
func (s *API) CreateLB(req *CreateLBRequest, opts ...scw.RequestOption) (*LB, error) {
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
		req.Name = namegenerator.GetRandomName("lb")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp LB

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLB: Get a load balancer.
func (s *API) GetLB(req *GetLBRequest, opts ...scw.RequestOption) (*LB, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LBID) + "",
	}

	var resp LB

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateLB: Update a load balancer.
func (s *API) UpdateLB(req *UpdateLBRequest, opts ...scw.RequestOption) (*LB, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LBID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp LB

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteLB: Delete a load balancer.
func (s *API) DeleteLB(req *DeleteLBRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "release_ip", req.ReleaseIP)

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LBID) + "",
		Query:  query,
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// MigrateLB: Migrate a load balancer.
func (s *API) MigrateLB(req *MigrateLBRequest, opts ...scw.RequestOption) (*LB, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LBID) + "/migrate",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp LB

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListIPs: List IPs.
func (s *API) ListIPs(req *ListIPsRequest, opts ...scw.RequestOption) (*ListIPsResponse, error) {
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
	parameter.AddToQuery(query, "ip_address", req.IPAddress)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/ips",
		Query:  query,
	}

	var resp ListIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateIP: Create an IP.
func (s *API) CreateIP(req *CreateIPRequest, opts ...scw.RequestOption) (*IP, error) {
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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/ips",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIP: Get an IP.
func (s *API) GetIP(req *GetIPRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IPID) + "",
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReleaseIP: Delete an IP.
func (s *API) ReleaseIP(req *ReleaseIPRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IPID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdateIP: Update an IP.
func (s *API) UpdateIP(req *UpdateIPRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IPID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListBackends: List backends in a given load balancer.
func (s *API) ListBackends(req *ListBackendsRequest, opts ...scw.RequestOption) (*ListBackendsResponse, error) {
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

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LBID) + "/backends",
		Query:  query,
	}

	var resp ListBackendsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateBackend: Create a backend in a given load balancer.
func (s *API) CreateBackend(req *CreateBackendRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("lbb")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LBID) + "/backends",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Backend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBackend: Get a backend in a given load balancer.
func (s *API) GetBackend(req *GetBackendRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BackendID) == "" {
		return nil, errors.New("field BackendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/backends/" + fmt.Sprint(req.BackendID) + "",
	}

	var resp Backend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateBackend: Update a backend in a given load balancer.
func (s *API) UpdateBackend(req *UpdateBackendRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BackendID) == "" {
		return nil, errors.New("field BackendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/backends/" + fmt.Sprint(req.BackendID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Backend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteBackend: Delete a backend in a given load balancer.
func (s *API) DeleteBackend(req *DeleteBackendRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BackendID) == "" {
		return errors.New("field BackendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/backends/" + fmt.Sprint(req.BackendID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// AddBackendServers: Add a set of servers in a given backend.
func (s *API) AddBackendServers(req *AddBackendServersRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BackendID) == "" {
		return nil, errors.New("field BackendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/backends/" + fmt.Sprint(req.BackendID) + "/servers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Backend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveBackendServers: Remove a set of servers for a given backend.
func (s *API) RemoveBackendServers(req *RemoveBackendServersRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BackendID) == "" {
		return nil, errors.New("field BackendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/backends/" + fmt.Sprint(req.BackendID) + "/servers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Backend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetBackendServers: Define all servers in a given backend.
func (s *API) SetBackendServers(req *SetBackendServersRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BackendID) == "" {
		return nil, errors.New("field BackendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/backends/" + fmt.Sprint(req.BackendID) + "/servers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Backend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHealthCheck: Update an health check for a given backend.
func (s *API) UpdateHealthCheck(req *UpdateHealthCheckRequest, opts ...scw.RequestOption) (*HealthCheck, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BackendID) == "" {
		return nil, errors.New("field BackendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/backends/" + fmt.Sprint(req.BackendID) + "/healthcheck",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp HealthCheck

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListFrontends: List frontends in a given load balancer.
func (s *API) ListFrontends(req *ListFrontendsRequest, opts ...scw.RequestOption) (*ListFrontendsResponse, error) {
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

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LBID) + "/frontends",
		Query:  query,
	}

	var resp ListFrontendsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFrontend: Create a frontend in a given load balancer.
func (s *API) CreateFrontend(req *CreateFrontendRequest, opts ...scw.RequestOption) (*Frontend, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("lbf")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LBID) + "/frontends",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Frontend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFrontend: Get a frontend.
func (s *API) GetFrontend(req *GetFrontendRequest, opts ...scw.RequestOption) (*Frontend, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FrontendID) == "" {
		return nil, errors.New("field FrontendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/frontends/" + fmt.Sprint(req.FrontendID) + "",
	}

	var resp Frontend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateFrontend: Update a frontend.
func (s *API) UpdateFrontend(req *UpdateFrontendRequest, opts ...scw.RequestOption) (*Frontend, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FrontendID) == "" {
		return nil, errors.New("field FrontendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/frontends/" + fmt.Sprint(req.FrontendID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Frontend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteFrontend: Delete a frontend.
func (s *API) DeleteFrontend(req *DeleteFrontendRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FrontendID) == "" {
		return errors.New("field FrontendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/frontends/" + fmt.Sprint(req.FrontendID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListRoutes: List all backend redirections.
func (s *API) ListRoutes(req *ListRoutesRequest, opts ...scw.RequestOption) (*ListRoutesResponse, error) {
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
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "frontend_id", req.FrontendID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/routes",
		Query:  query,
	}

	var resp ListRoutesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRoute: Create a backend redirection.
func (s *API) CreateRoute(req *CreateRouteRequest, opts ...scw.RequestOption) (*Route, error) {
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
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/routes",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Route

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRoute: Get single backend redirection.
func (s *API) GetRoute(req *GetRouteRequest, opts ...scw.RequestOption) (*Route, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.RouteID) == "" {
		return nil, errors.New("field RouteID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/routes/" + fmt.Sprint(req.RouteID) + "",
	}

	var resp Route

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateRoute: Edit a backend redirection.
func (s *API) UpdateRoute(req *UpdateRouteRequest, opts ...scw.RequestOption) (*Route, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.RouteID) == "" {
		return nil, errors.New("field RouteID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/routes/" + fmt.Sprint(req.RouteID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Route

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteRoute: Delete a backend redirection.
func (s *API) DeleteRoute(req *DeleteRouteRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.RouteID) == "" {
		return errors.New("field RouteID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/routes/" + fmt.Sprint(req.RouteID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetLBStats: Get usage statistics of a given load balancer.
func (s *API) GetLBStats(req *GetLBStatsRequest, opts ...scw.RequestOption) (*LBStats, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "backend_id", req.BackendID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LBID) + "/stats",
		Query:  query,
	}

	var resp LBStats

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListBackendStats: List backend server statistics.
func (s *API) ListBackendStats(req *ListBackendStatsRequest, opts ...scw.RequestOption) (*ListBackendStatsResponse, error) {
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
	parameter.AddToQuery(query, "backend_id", req.BackendID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LBID) + "/backend-stats",
		Query:  query,
	}

	var resp ListBackendStatsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListACLs: List ACL for a given frontend.
func (s *API) ListACLs(req *ListACLsRequest, opts ...scw.RequestOption) (*ListACLResponse, error) {
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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FrontendID) == "" {
		return nil, errors.New("field FrontendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/frontends/" + fmt.Sprint(req.FrontendID) + "/acls",
		Query:  query,
	}

	var resp ListACLResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateACL: Create an ACL for a given frontend.
func (s *API) CreateACL(req *CreateACLRequest, opts ...scw.RequestOption) (*ACL, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("acl")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FrontendID) == "" {
		return nil, errors.New("field FrontendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/frontends/" + fmt.Sprint(req.FrontendID) + "/acls",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ACL

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetACL: Get an ACL.
func (s *API) GetACL(req *GetACLRequest, opts ...scw.RequestOption) (*ACL, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ACLID) == "" {
		return nil, errors.New("field ACLID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/acls/" + fmt.Sprint(req.ACLID) + "",
	}

	var resp ACL

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateACL: Update an ACL.
func (s *API) UpdateACL(req *UpdateACLRequest, opts ...scw.RequestOption) (*ACL, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ACLID) == "" {
		return nil, errors.New("field ACLID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/acls/" + fmt.Sprint(req.ACLID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ACL

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteACL: Delete an ACL.
func (s *API) DeleteACL(req *DeleteACLRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ACLID) == "" {
		return errors.New("field ACLID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/acls/" + fmt.Sprint(req.ACLID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CreateCertificate: Generate a new TLS certificate using Let's Encrypt or import your certificate.
func (s *API) CreateCertificate(req *CreateCertificateRequest, opts ...scw.RequestOption) (*Certificate, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("certificate")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LBID) + "/certificates",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Certificate

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListCertificates: List all TLS certificates on a given load balancer.
func (s *API) ListCertificates(req *ListCertificatesRequest, opts ...scw.RequestOption) (*ListCertificatesResponse, error) {
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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LBID) + "/certificates",
		Query:  query,
	}

	var resp ListCertificatesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCertificate: Get a TLS certificate.
func (s *API) GetCertificate(req *GetCertificateRequest, opts ...scw.RequestOption) (*Certificate, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.CertificateID) == "" {
		return nil, errors.New("field CertificateID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/certificates/" + fmt.Sprint(req.CertificateID) + "",
	}

	var resp Certificate

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateCertificate: Update a TLS certificate.
func (s *API) UpdateCertificate(req *UpdateCertificateRequest, opts ...scw.RequestOption) (*Certificate, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.CertificateID) == "" {
		return nil, errors.New("field CertificateID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/certificates/" + fmt.Sprint(req.CertificateID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Certificate

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteCertificate: Delete a TLS certificate.
func (s *API) DeleteCertificate(req *DeleteCertificateRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.CertificateID) == "" {
		return errors.New("field CertificateID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/certificates/" + fmt.Sprint(req.CertificateID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListLBTypes: List all load balancer offer type.
func (s *API) ListLBTypes(req *ListLBTypesRequest, opts ...scw.RequestOption) (*ListLBTypesResponse, error) {
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
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lb-types",
		Query:  query,
	}

	var resp ListLBTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSubscriber: Create a subscriber, webhook or email.
func (s *API) CreateSubscriber(req *CreateSubscriberRequest, opts ...scw.RequestOption) (*Subscriber, error) {
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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/subscribers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Subscriber

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSubscriber: Get a subscriber.
func (s *API) GetSubscriber(req *GetSubscriberRequest, opts ...scw.RequestOption) (*Subscriber, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SubscriberID) == "" {
		return nil, errors.New("field SubscriberID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/subscribers/" + fmt.Sprint(req.SubscriberID) + "",
	}

	var resp Subscriber

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSubscriber: List all subscriber.
func (s *API) ListSubscriber(req *ListSubscriberRequest, opts ...scw.RequestOption) (*ListSubscriberResponse, error) {
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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/subscribers",
		Query:  query,
	}

	var resp ListSubscriberResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSubscriber: Update a subscriber.
func (s *API) UpdateSubscriber(req *UpdateSubscriberRequest, opts ...scw.RequestOption) (*Subscriber, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SubscriberID) == "" {
		return nil, errors.New("field SubscriberID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/subscribers/" + fmt.Sprint(req.SubscriberID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Subscriber

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSubscriber: Delete a subscriber.
func (s *API) DeleteSubscriber(req *DeleteSubscriberRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SubscriberID) == "" {
		return errors.New("field SubscriberID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lb/subscriber/" + fmt.Sprint(req.SubscriberID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// SubscribeToLB: Subscribe a subscriber to a given load balancer.
func (s *API) SubscribeToLB(req *SubscribeToLBRequest, opts ...scw.RequestOption) (*LB, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lb/" + fmt.Sprint(req.LBID) + "/subscribe",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp LB

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsubscribeFromLB: Unsubscribe a subscriber from a given load balancer.
func (s *API) UnsubscribeFromLB(req *UnsubscribeFromLBRequest, opts ...scw.RequestOption) (*LB, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lb/" + fmt.Sprint(req.LBID) + "/unsubscribe",
	}

	var resp LB

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListLBPrivateNetworks: List attached private network of load balancer.
func (s *API) ListLBPrivateNetworks(req *ListLBPrivateNetworksRequest, opts ...scw.RequestOption) (*ListLBPrivateNetworksResponse, error) {
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
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LBID) + "/private-networks",
		Query:  query,
	}

	var resp ListLBPrivateNetworksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachPrivateNetwork: Add load balancer on instance private network.
func (s *API) AttachPrivateNetwork(req *AttachPrivateNetworkRequest, opts ...scw.RequestOption) (*PrivateNetwork, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return nil, errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LBID) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "/attach",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PrivateNetwork

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachPrivateNetwork: Remove load balancer of private network.
func (s *API) DetachPrivateNetwork(req *DetachPrivateNetworkRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return errors.New("field LBID cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LBID) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "/detach",
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
