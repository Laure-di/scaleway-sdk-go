// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package iot provides methods and message types of the iot v1 API.
package iot

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

type DeviceMessageFiltersRulePolicy string

const (
	DeviceMessageFiltersRulePolicyUnknown = DeviceMessageFiltersRulePolicy("unknown")
	DeviceMessageFiltersRulePolicyAccept  = DeviceMessageFiltersRulePolicy("accept")
	DeviceMessageFiltersRulePolicyReject  = DeviceMessageFiltersRulePolicy("reject")
)

func (enum DeviceMessageFiltersRulePolicy) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum DeviceMessageFiltersRulePolicy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DeviceMessageFiltersRulePolicy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DeviceMessageFiltersRulePolicy(DeviceMessageFiltersRulePolicy(tmp).String())
	return nil
}

type DeviceStatus string

const (
	DeviceStatusUnknown  = DeviceStatus("unknown")
	DeviceStatusError    = DeviceStatus("error")
	DeviceStatusEnabled  = DeviceStatus("enabled")
	DeviceStatusDisabled = DeviceStatus("disabled")
)

func (enum DeviceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum DeviceStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DeviceStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DeviceStatus(DeviceStatus(tmp).String())
	return nil
}

type HubProductPlan string

const (
	HubProductPlanPlanUnknown   = HubProductPlan("plan_unknown")
	HubProductPlanPlanShared    = HubProductPlan("plan_shared")
	HubProductPlanPlanDedicated = HubProductPlan("plan_dedicated")
	HubProductPlanPlanHa        = HubProductPlan("plan_ha")
)

func (enum HubProductPlan) String() string {
	if enum == "" {
		// return default value if empty
		return "plan_unknown"
	}
	return string(enum)
}

func (enum HubProductPlan) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *HubProductPlan) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = HubProductPlan(HubProductPlan(tmp).String())
	return nil
}

type HubStatus string

const (
	HubStatusUnknown   = HubStatus("unknown")
	HubStatusError     = HubStatus("error")
	HubStatusEnabling  = HubStatus("enabling")
	HubStatusReady     = HubStatus("ready")
	HubStatusDisabling = HubStatus("disabling")
	HubStatusDisabled  = HubStatus("disabled")
)

func (enum HubStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum HubStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *HubStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = HubStatus(HubStatus(tmp).String())
	return nil
}

type ListDevicesRequestOrderBy string

const (
	ListDevicesRequestOrderByNameAsc           = ListDevicesRequestOrderBy("name_asc")
	ListDevicesRequestOrderByNameDesc          = ListDevicesRequestOrderBy("name_desc")
	ListDevicesRequestOrderByStatusAsc         = ListDevicesRequestOrderBy("status_asc")
	ListDevicesRequestOrderByStatusDesc        = ListDevicesRequestOrderBy("status_desc")
	ListDevicesRequestOrderByHubIDAsc          = ListDevicesRequestOrderBy("hub_id_asc")
	ListDevicesRequestOrderByHubIDDesc         = ListDevicesRequestOrderBy("hub_id_desc")
	ListDevicesRequestOrderByCreatedAtAsc      = ListDevicesRequestOrderBy("created_at_asc")
	ListDevicesRequestOrderByCreatedAtDesc     = ListDevicesRequestOrderBy("created_at_desc")
	ListDevicesRequestOrderByUpdatedAtAsc      = ListDevicesRequestOrderBy("updated_at_asc")
	ListDevicesRequestOrderByUpdatedAtDesc     = ListDevicesRequestOrderBy("updated_at_desc")
	ListDevicesRequestOrderByAllowInsecureAsc  = ListDevicesRequestOrderBy("allow_insecure_asc")
	ListDevicesRequestOrderByAllowInsecureDesc = ListDevicesRequestOrderBy("allow_insecure_desc")
)

func (enum ListDevicesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListDevicesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDevicesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDevicesRequestOrderBy(ListDevicesRequestOrderBy(tmp).String())
	return nil
}

type ListHubsRequestOrderBy string

const (
	ListHubsRequestOrderByNameAsc         = ListHubsRequestOrderBy("name_asc")
	ListHubsRequestOrderByNameDesc        = ListHubsRequestOrderBy("name_desc")
	ListHubsRequestOrderByStatusAsc       = ListHubsRequestOrderBy("status_asc")
	ListHubsRequestOrderByStatusDesc      = ListHubsRequestOrderBy("status_desc")
	ListHubsRequestOrderByProductPlanAsc  = ListHubsRequestOrderBy("product_plan_asc")
	ListHubsRequestOrderByProductPlanDesc = ListHubsRequestOrderBy("product_plan_desc")
	ListHubsRequestOrderByCreatedAtAsc    = ListHubsRequestOrderBy("created_at_asc")
	ListHubsRequestOrderByCreatedAtDesc   = ListHubsRequestOrderBy("created_at_desc")
	ListHubsRequestOrderByUpdatedAtAsc    = ListHubsRequestOrderBy("updated_at_asc")
	ListHubsRequestOrderByUpdatedAtDesc   = ListHubsRequestOrderBy("updated_at_desc")
)

func (enum ListHubsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListHubsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListHubsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListHubsRequestOrderBy(ListHubsRequestOrderBy(tmp).String())
	return nil
}

type ListNetworksRequestOrderBy string

const (
	ListNetworksRequestOrderByNameAsc       = ListNetworksRequestOrderBy("name_asc")
	ListNetworksRequestOrderByNameDesc      = ListNetworksRequestOrderBy("name_desc")
	ListNetworksRequestOrderByTypeAsc       = ListNetworksRequestOrderBy("type_asc")
	ListNetworksRequestOrderByTypeDesc      = ListNetworksRequestOrderBy("type_desc")
	ListNetworksRequestOrderByCreatedAtAsc  = ListNetworksRequestOrderBy("created_at_asc")
	ListNetworksRequestOrderByCreatedAtDesc = ListNetworksRequestOrderBy("created_at_desc")
)

func (enum ListNetworksRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListNetworksRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListNetworksRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListNetworksRequestOrderBy(ListNetworksRequestOrderBy(tmp).String())
	return nil
}

type ListRoutesRequestOrderBy string

const (
	ListRoutesRequestOrderByNameAsc       = ListRoutesRequestOrderBy("name_asc")
	ListRoutesRequestOrderByNameDesc      = ListRoutesRequestOrderBy("name_desc")
	ListRoutesRequestOrderByHubIDAsc      = ListRoutesRequestOrderBy("hub_id_asc")
	ListRoutesRequestOrderByHubIDDesc     = ListRoutesRequestOrderBy("hub_id_desc")
	ListRoutesRequestOrderByTypeAsc       = ListRoutesRequestOrderBy("type_asc")
	ListRoutesRequestOrderByTypeDesc      = ListRoutesRequestOrderBy("type_desc")
	ListRoutesRequestOrderByCreatedAtAsc  = ListRoutesRequestOrderBy("created_at_asc")
	ListRoutesRequestOrderByCreatedAtDesc = ListRoutesRequestOrderBy("created_at_desc")
)

func (enum ListRoutesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
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

type NetworkNetworkType string

const (
	NetworkNetworkTypeUnknown = NetworkNetworkType("unknown")
	NetworkNetworkTypeSigfox  = NetworkNetworkType("sigfox")
	NetworkNetworkTypeRest    = NetworkNetworkType("rest")
)

func (enum NetworkNetworkType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum NetworkNetworkType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NetworkNetworkType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NetworkNetworkType(NetworkNetworkType(tmp).String())
	return nil
}

type RouteDatabaseConfigEngine string

const (
	RouteDatabaseConfigEngineUnknown    = RouteDatabaseConfigEngine("unknown")
	RouteDatabaseConfigEnginePostgresql = RouteDatabaseConfigEngine("postgresql")
	RouteDatabaseConfigEngineMysql      = RouteDatabaseConfigEngine("mysql")
)

func (enum RouteDatabaseConfigEngine) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum RouteDatabaseConfigEngine) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RouteDatabaseConfigEngine) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RouteDatabaseConfigEngine(RouteDatabaseConfigEngine(tmp).String())
	return nil
}

type RouteRestConfigHTTPVerb string

const (
	RouteRestConfigHTTPVerbUnknown = RouteRestConfigHTTPVerb("unknown")
	RouteRestConfigHTTPVerbGet     = RouteRestConfigHTTPVerb("get")
	RouteRestConfigHTTPVerbPost    = RouteRestConfigHTTPVerb("post")
	RouteRestConfigHTTPVerbPut     = RouteRestConfigHTTPVerb("put")
	RouteRestConfigHTTPVerbPatch   = RouteRestConfigHTTPVerb("patch")
	RouteRestConfigHTTPVerbDelete  = RouteRestConfigHTTPVerb("delete")
)

func (enum RouteRestConfigHTTPVerb) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum RouteRestConfigHTTPVerb) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RouteRestConfigHTTPVerb) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RouteRestConfigHTTPVerb(RouteRestConfigHTTPVerb(tmp).String())
	return nil
}

type RouteRouteType string

const (
	RouteRouteTypeUnknown  = RouteRouteType("unknown")
	RouteRouteTypeS3       = RouteRouteType("s3")
	RouteRouteTypeDatabase = RouteRouteType("database")
	RouteRouteTypeRest     = RouteRouteType("rest")
)

func (enum RouteRouteType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum RouteRouteType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RouteRouteType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RouteRouteType(RouteRouteType(tmp).String())
	return nil
}

type RouteS3ConfigS3Strategy string

const (
	RouteS3ConfigS3StrategyUnknown    = RouteS3ConfigS3Strategy("unknown")
	RouteS3ConfigS3StrategyPerTopic   = RouteS3ConfigS3Strategy("per_topic")
	RouteS3ConfigS3StrategyPerMessage = RouteS3ConfigS3Strategy("per_message")
)

func (enum RouteS3ConfigS3Strategy) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum RouteS3ConfigS3Strategy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RouteS3ConfigS3Strategy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RouteS3ConfigS3Strategy(RouteS3ConfigS3Strategy(tmp).String())
	return nil
}

type DeviceMessageFiltersRule struct {
	// Policy: If set to `accept`, all topics in the topics list will be allowed, with all other topics being denied.
	// If set to `reject`, all topics in the topics list will be denied, with all other topics being allowed.
	Policy DeviceMessageFiltersRulePolicy `json:"policy"`
	// Topics: List of topics to accept or reject. It must be valid MQTT topics and up to 65535 characters.
	Topics *[]string `json:"topics,omitempty"`
}

type DeviceMessageFilters struct {
	// Publish: Filtering rule to restrict topics the device can publish to.
	Publish *DeviceMessageFiltersRule `json:"publish"`
	// Subscribe: Filtering rule to restrict topics the device can subscribe to.
	Subscribe *DeviceMessageFiltersRule `json:"subscribe"`
}

type HubTwinsGraphiteConfig struct {
	// PushURI:
	PushURI string `json:"push_uri"`
}

type Certificate struct {
	// Crt:
	Crt string `json:"crt"`
	// Key:
	Key string `json:"key"`
}

type Device struct {
	// ID: Device ID, also used as MQTT Client ID or username.
	ID string `json:"id"`
	// Name: Device name.
	Name string `json:"name"`
	// Description: Device description.
	Description string `json:"description"`
	// Status: Device status.
	Status DeviceStatus `json:"status"`
	// HubID: Hub ID.
	HubID string `json:"hub_id"`
	// LastActivityAt: Last connection/activity date of a device.
	LastActivityAt *time.Time `json:"last_activity_at,omitempty"`
	// IsConnected: Defines whether the device is connected to the Hub.
	IsConnected bool `json:"is_connected"`
	// AllowInsecure: Defines whether to allow the device to connect to the Hub without TLS mutual authentication.
	AllowInsecure bool `json:"allow_insecure"`
	// AllowMultipleConnections: Defines whether to allow multiple physical devices to connect to the Hub with this device's credentials.
	AllowMultipleConnections bool `json:"allow_multiple_connections"`
	// MessageFilters: Filter-sets to restrict the topics the device can publish/subscribe to.
	MessageFilters *DeviceMessageFilters `json:"message_filters"`
	// HasCustomCertificate: Assigning a custom certificate allows a device to authenticate using that specific certificate without checking the Hub's CA certificate.
	HasCustomCertificate bool `json:"has_custom_certificate"`
	// CreatedAt: Date at which the device was added.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Date at which the device was last modified.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Network struct {
	// ID: Network ID.
	ID string `json:"id"`
	// Name: Network name.
	Name string `json:"name"`
	// Type: Type of network to connect with.
	Type NetworkNetworkType `json:"type"`
	// Endpoint: Endpoint to use for interacting with the network.
	Endpoint string `json:"endpoint"`
	// HubID: Hub ID to connect the Network to.
	HubID string `json:"hub_id"`
	// CreatedAt: Date at which the network was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// TopicPrefix: This prefix will be prepended to all topics for this Network.
	TopicPrefix string `json:"topic_prefix"`
}

type CreateRouteRequestDatabaseConfig struct {
	// Host:
	Host string `json:"host"`
	// Port:
	Port uint32 `json:"port"`
	// Dbname:
	Dbname string `json:"dbname"`
	// Username:
	Username string `json:"username"`
	// Password:
	Password string `json:"password"`
	// Query:
	Query string `json:"query"`
	// Engine:
	Engine RouteDatabaseConfigEngine `json:"engine"`
}

type CreateRouteRequestRestConfig struct {
	// Verb:
	Verb RouteRestConfigHTTPVerb `json:"verb"`
	// URI:
	URI string `json:"uri"`
	// Headers:
	Headers map[string]string `json:"headers"`
}

type CreateRouteRequestS3Config struct {
	// BucketRegion:
	BucketRegion string `json:"bucket_region"`
	// BucketName:
	BucketName string `json:"bucket_name"`
	// ObjectPrefix:
	ObjectPrefix string `json:"object_prefix"`
	// Strategy:
	Strategy RouteS3ConfigS3Strategy `json:"strategy"`
}

type Hub struct {
	// ID: Hub ID.
	ID string `json:"id"`
	// Name: Hub name.
	Name string `json:"name"`
	// Status: Current status of the Hub.
	Status HubStatus `json:"status"`
	// ProductPlan: Hub feature set.
	ProductPlan HubProductPlan `json:"product_plan"`
	// Enabled: Defines whether the hub has been enabled.
	Enabled bool `json:"enabled"`
	// DeviceCount: Number of registered devices.
	DeviceCount uint64 `json:"device_count"`
	// ConnectedDeviceCount: Number of currently connected devices.
	ConnectedDeviceCount uint64 `json:"connected_device_count"`
	// Endpoint: Devices should be connected to this host. Port may be 1883 (MQTT), 8883 (MQTT over TLS), 80 (MQTT over Websocket) or 443 (MQTT over Websocket over TLS).
	Endpoint string `json:"endpoint"`
	// DisableEvents: Defines whether to disable Hub events.
	DisableEvents bool `json:"disable_events"`
	// EventsTopicPrefix: Hub events topic prefix.
	EventsTopicPrefix string `json:"events_topic_prefix"`
	// Region: Region of the Hub.
	Region scw.Region `json:"region"`
	// CreatedAt: Hub creation date.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Hub last modification date.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// ProjectID: Project owning the resource.
	ProjectID string `json:"project_id"`
	// OrganizationID: Organization owning the resource.
	OrganizationID string `json:"organization_id"`
	// EnableDeviceAutoProvisioning: When an unknown device connects to your hub using a valid certificate chain, it will be automatically provisioned inside your Hub. The Hub uses the common name of the device certifcate to find out if a device with the same name already exists. This setting can only be enabled on a hub with a custom certificate authority.
	EnableDeviceAutoProvisioning bool `json:"enable_device_auto_provisioning"`
	// HasCustomCa: Flag is automatically set to `false` after Hub creation, as Hub certificates are managed by Scaleway. Once a custom certificate authority is set, the flag will be set to `true`.
	HasCustomCa bool `json:"has_custom_ca"`
	// TwinsGraphiteConfig: BETA - not implemented yet.
	TwinsGraphiteConfig *HubTwinsGraphiteConfig `json:"twins_graphite_config,omitempty"`
}

type RouteSummary struct {
	// ID: Route ID.
	ID string `json:"id"`
	// Name: Route name.
	Name string `json:"name"`
	// HubID: Hub ID of the route.
	HubID string `json:"hub_id"`
	// Topic: Topic the route subscribes to. It must be a valid MQTT topic and up to 65535 characters.
	Topic string `json:"topic"`
	// Type: Route type.
	Type RouteRouteType `json:"type"`
	// CreatedAt: Date at which the route was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Date at which the route was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type ListTwinDocumentsResponseDocumentSummary struct {
	// DocumentName: Name of the document.
	DocumentName string `json:"document_name"`
}

type RouteDatabaseConfig struct {
	// Engine: Database engine the route will connect to. If not specified, the default database will be 'PostgreSQL'.
	Engine RouteDatabaseConfigEngine `json:"engine"`
	// Host: Database host.
	Host string `json:"host"`
	// Port: Database port.
	Port uint32 `json:"port"`
	// Dbname: Database name.
	Dbname string `json:"dbname"`
	// Username: Database username. Make sure this account can execute the provided query.
	Username string `json:"username"`
	// Password: Database password.
	Password string `json:"password"`
	// Query: SQL query to be executed ($TOPIC and $PAYLOAD variables are available, see documentation).
	Query string `json:"query"`
}

type RouteRestConfig struct {
	// Verb: HTTP verb used to call REST URI.
	Verb RouteRestConfigHTTPVerb `json:"verb"`
	// URI: URI of the REST endpoint.
	URI string `json:"uri"`
	// Headers: HTTP call extra headers.
	Headers map[string]string `json:"headers"`
}

type RouteS3Config struct {
	// BucketRegion: Region of the S3 route's destination bucket (e.g., 'fr-par').
	BucketRegion string `json:"bucket_region"`
	// BucketName: Destination bucket name of the S3 route.
	BucketName string `json:"bucket_name"`
	// ObjectPrefix: Optional string to prefix object names with.
	ObjectPrefix string `json:"object_prefix"`
	// Strategy: How the S3 route's objects will be created: one per topic or one per message.
	Strategy RouteS3ConfigS3Strategy `json:"strategy"`
}

type UpdateRouteRequestDatabaseConfig struct {
	// Host:
	Host *string `json:"host,omitempty"`
	// Port:
	Port *uint32 `json:"port,omitempty"`
	// Dbname:
	Dbname *string `json:"dbname,omitempty"`
	// Username:
	Username *string `json:"username,omitempty"`
	// Password:
	Password *string `json:"password,omitempty"`
	// Query:
	Query *string `json:"query,omitempty"`
	// Engine:
	Engine RouteDatabaseConfigEngine `json:"engine"`
}

type UpdateRouteRequestRestConfig struct {
	// Verb:
	Verb RouteRestConfigHTTPVerb `json:"verb"`
	// URI:
	URI *string `json:"uri,omitempty"`
	// Headers:
	Headers *map[string]string `json:"headers,omitempty"`
}

type UpdateRouteRequestS3Config struct {
	// BucketRegion:
	BucketRegion *string `json:"bucket_region,omitempty"`
	// BucketName:
	BucketName *string `json:"bucket_name,omitempty"`
	// ObjectPrefix:
	ObjectPrefix *string `json:"object_prefix,omitempty"`
	// Strategy:
	Strategy RouteS3ConfigS3Strategy `json:"strategy"`
}

type CreateDeviceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Name: Device name.
	Name string `json:"name"`
	// HubID: Hub ID of the device.
	HubID string `json:"hub_id"`
	// AllowInsecure: Defines whether to allow plain and server-authenticated SSL connections in addition to mutually-authenticated ones.
	AllowInsecure bool `json:"allow_insecure"`
	// AllowMultipleConnections: Defines whether to allow multiple physical devices to connect with this device's credentials.
	AllowMultipleConnections bool `json:"allow_multiple_connections"`
	// MessageFilters: Filter-sets to authorize or deny the device to publish/subscribe to specific topics.
	MessageFilters *DeviceMessageFilters `json:"message_filters"`
	// Description: Device description.
	Description *string `json:"description,omitempty"`
}

type CreateDeviceResponse struct {
	// Device: Information related to the created device.
	Device *Device `json:"device"`
	// Certificate: Device certificate.
	Certificate *Certificate `json:"certificate"`
}

type CreateHubRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Name: Hub name (up to 255 characters).
	Name string `json:"name"`
	// ProjectID: Project/Organization ID to filter for, only Hubs from this Project/Organization will be returned.
	ProjectID string `json:"project_id"`
	// ProductPlan: Hub product plan.
	ProductPlan HubProductPlan `json:"product_plan"`
	// DisableEvents: Disable Hub events.
	DisableEvents *bool `json:"disable_events,omitempty"`
	// EventsTopicPrefix: Topic prefix (default '$SCW/events') of Hub events.
	EventsTopicPrefix *string `json:"events_topic_prefix,omitempty"`
	// TwinsGraphiteConfig: BETA - not implemented yet.
	TwinsGraphiteConfig *HubTwinsGraphiteConfig `json:"twins_graphite_config,omitempty"`
}

type CreateNetworkRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Name: Network name.
	Name string `json:"name"`
	// Type: Type of network to connect with.
	Type NetworkNetworkType `json:"type"`
	// HubID: Hub ID to connect the Network to.
	HubID string `json:"hub_id"`
	// TopicPrefix: Topic prefix for the Network.
	TopicPrefix string `json:"topic_prefix"`
}

type CreateNetworkResponse struct {
	// Network: Information related to the created network.
	Network *Network `json:"network"`
	// Secret: Endpoint Key to keep secret. This cannot be retrieved later.
	Secret string `json:"secret"`
}

type CreateRouteRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Name: Route name.
	Name string `json:"name"`
	// HubID: Hub ID of the route.
	HubID string `json:"hub_id"`
	// Topic: Topic the route subscribes to. It must be a valid MQTT topic and up to 65535 characters.
	Topic string `json:"topic"`
	// S3Config: If creating S3 Route, S3-specific configuration fields.
	S3Config *CreateRouteRequestS3Config `json:"s3_config,omitempty"`
	// DbConfig: If creating Database Route, DB-specific configuration fields.
	DbConfig *CreateRouteRequestDatabaseConfig `json:"db_config,omitempty"`
	// RestConfig: If creating Rest Route, Rest-specific configuration fields.
	RestConfig *CreateRouteRequestRestConfig `json:"rest_config,omitempty"`
}

type DeleteDeviceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// DeviceID: Device ID.
	DeviceID string `json:"-"`
}

type DeleteHubRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// HubID: Hub ID.
	HubID string `json:"-"`
	// DeleteDevices: Defines whether to force the deletion of devices added to this Hub or reject the operation.
	DeleteDevices *bool `json:"delete_devices,omitempty"`
}

type DeleteNetworkRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NetworkID: Network ID.
	NetworkID string `json:"-"`
}

type DeleteRouteRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// RouteID: Route ID.
	RouteID string `json:"-"`
}

type DeleteTwinDocumentRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// TwinID: Twin ID.
	TwinID string `json:"-"`
	// DocumentName: Name of the document.
	DocumentName string `json:"-"`
}

type DeleteTwinDocumentsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// TwinID: Twin ID.
	TwinID string `json:"-"`
}

type DisableDeviceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// DeviceID: Device ID.
	DeviceID string `json:"-"`
}

type DisableHubRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// HubID: Hub ID.
	HubID string `json:"-"`
}

type EnableDeviceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// DeviceID: Device ID.
	DeviceID string `json:"-"`
}

type EnableHubRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// HubID: Hub ID.
	HubID string `json:"-"`
}

type GetDeviceCertificateRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// DeviceID: Device ID.
	DeviceID string `json:"-"`
}

type GetDeviceCertificateResponse struct {
	// Device: Information related to the created device.
	Device *Device `json:"device"`
	// CertificatePem: Device certificate.
	CertificatePem string `json:"certificate_pem"`
}

type GetDeviceMetricsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// DeviceID: Device ID.
	DeviceID string `json:"-"`
	// StartDate: Start date used to compute the best scale for the returned metrics.
	StartDate *time.Time `json:"start_date,omitempty"`
}

type GetDeviceMetricsResponse struct {
	// Metrics: Metrics for a device over the requested period.
	Metrics []*scw.TimeSeries `json:"metrics"`
}

type GetDeviceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// DeviceID: Device ID.
	DeviceID string `json:"-"`
}

type GetHubCARequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// HubID:
	HubID string `json:"-"`
}

type GetHubCAResponse struct {
	// CaCertPem:
	CaCertPem string `json:"ca_cert_pem"`
}

type GetHubMetricsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// HubID: Hub ID.
	HubID string `json:"-"`
	// StartDate: Start date used to compute the best scale for returned metrics.
	StartDate *time.Time `json:"start_date,omitempty"`
}

type GetHubMetricsResponse struct {
	// Metrics: Metrics for a Hub over the requested period.
	Metrics []*scw.TimeSeries `json:"metrics"`
}

type GetHubRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// HubID: Hub ID.
	HubID string `json:"-"`
}

type GetNetworkRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NetworkID: Network ID.
	NetworkID string `json:"-"`
}

type GetRouteRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// RouteID: Route ID.
	RouteID string `json:"-"`
}

type GetTwinDocumentRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// TwinID: Twin ID.
	TwinID string `json:"-"`
	// DocumentName: Name of the document.
	DocumentName string `json:"-"`
}

type ListDevicesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Page: Page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Number of devices to return within a page. Maximum value is 100.
	PageSize *uint32 `json:"page_size,omitempty"`
	// OrderBy: Ordering of requested devices.
	OrderBy ListDevicesRequestOrderBy `json:"order_by"`
	// Name: Name to filter for, only devices with this name will be returned.
	Name *string `json:"name,omitempty"`
	// HubID: Hub ID to filter for, only devices attached to this Hub will be returned.
	HubID *string `json:"hub_id,omitempty"`
	// AllowInsecure: Defines wheter to filter the allow_insecure flag.
	AllowInsecure *bool `json:"allow_insecure,omitempty"`
	// Status: Device status (enabled, disabled, etc.).
	Status DeviceStatus `json:"status"`
}

type ListDevicesResponse struct {
	// TotalCount: Total number of devices.
	TotalCount uint32 `json:"total_count"`
	// Devices: Page of devices.
	Devices []*Device `json:"devices"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDevicesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDevicesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDevicesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Devices = append(r.Devices, results.Devices...)
	r.TotalCount += uint32(len(results.Devices))
	return uint32(len(results.Devices)), nil
}

type ListHubsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Page: Page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Number of Hubs to return within a page. Maximum value is 100.
	PageSize *uint32 `json:"page_size,omitempty"`
	// OrderBy: Sort order of Hubs in the response.
	OrderBy ListHubsRequestOrderBy `json:"order_by"`
	// ProjectID: Only list Hubs of this Project ID.
	ProjectID *string `json:"project_id,omitempty"`
	// OrganizationID: Only list Hubs of this Organization ID.
	OrganizationID *string `json:"organization_id,omitempty"`
	// Name: Hub name.
	Name *string `json:"name,omitempty"`
}

type ListHubsResponse struct {
	// TotalCount: Total number of Hubs.
	TotalCount uint32 `json:"total_count"`
	// Hubs: A page of hubs.
	Hubs []*Hub `json:"hubs"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListHubsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListHubsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListHubsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Hubs = append(r.Hubs, results.Hubs...)
	r.TotalCount += uint32(len(results.Hubs))
	return uint32(len(results.Hubs)), nil
}

type ListNetworksRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Page: Page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Number of networks to return. The maximum value is 100.
	PageSize *uint32 `json:"page_size,omitempty"`
	// OrderBy: Ordering of requested routes.
	OrderBy ListNetworksRequestOrderBy `json:"order_by"`
	// Name: Network name to filter for.
	Name *string `json:"name,omitempty"`
	// HubID: Hub ID to filter for.
	HubID *string `json:"hub_id,omitempty"`
	// TopicPrefix: Topic prefix to filter for.
	TopicPrefix *string `json:"topic_prefix,omitempty"`
}

type ListNetworksResponse struct {
	// TotalCount: Total number of Networks.
	TotalCount uint32 `json:"total_count"`
	// Networks: Page of networks.
	Networks []*Network `json:"networks"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNetworksResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNetworksResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListNetworksResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Networks = append(r.Networks, results.Networks...)
	r.TotalCount += uint32(len(results.Networks))
	return uint32(len(results.Networks)), nil
}

type ListRoutesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Page: Page number to return, from the paginated results.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Number of routes to return within a page. Maximum value is 100.
	PageSize *uint32 `json:"page_size,omitempty"`
	// OrderBy: Ordering of requested routes.
	OrderBy ListRoutesRequestOrderBy `json:"order_by"`
	// HubID: Hub ID to filter for.
	HubID *string `json:"hub_id,omitempty"`
	// Name: Route name to filter for.
	Name *string `json:"name,omitempty"`
}

type ListRoutesResponse struct {
	// TotalCount: Total number of routes.
	TotalCount uint32 `json:"total_count"`
	// Routes: Page of routes.
	Routes []*RouteSummary `json:"routes"`
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

type ListTwinDocumentsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// TwinID: Twin ID.
	TwinID string `json:"-"`
}

type ListTwinDocumentsResponse struct {
	// Documents: List of the twin document.
	Documents []*ListTwinDocumentsResponseDocumentSummary `json:"documents"`
}

type PatchTwinDocumentRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// TwinID: Twin ID.
	TwinID string `json:"-"`
	// DocumentName: Name of the document.
	DocumentName string `json:"-"`
	// Version: If set, ensures that the current version of the document matches before persisting the update.
	Version *uint32 `json:"version,omitempty"`
	// Data: A json data that will be applied on the document's current data.
	// Patching rules:
	// * The patch goes recursively through the patch objects.
	// * If the patch object property is null, it is removed from the final object.
	// * If the patch object property is a value (number, strings, bool, arrays), it is replaced.
	// * If the patch object property is an object, the previous rules will be applied recursively on it.
	Data *scw.JSONObject `json:"data,omitempty"`
}

type PutTwinDocumentRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// TwinID: Twin ID.
	TwinID string `json:"-"`
	// DocumentName: Name of the document.
	DocumentName string `json:"-"`
	// Version: If set, ensures that the current version of the document matches before persisting the update.
	Version *uint32 `json:"version,omitempty"`
	// Data: New data that will replace the contents of the document.
	Data *scw.JSONObject `json:"data,omitempty"`
}

type RenewDeviceCertificateRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// DeviceID: Device ID.
	DeviceID string `json:"-"`
}

type RenewDeviceCertificateResponse struct {
	// Device: Information related to the created device.
	Device *Device `json:"device"`
	// Certificate: Device certificate.
	Certificate *Certificate `json:"certificate"`
}

type Route struct {
	// ID: Route ID.
	ID string `json:"id"`
	// Name: Route name.
	Name string `json:"name"`
	// HubID: Hub ID of the route.
	HubID string `json:"hub_id"`
	// Topic: Topic the route subscribes to. It must be a valid MQTT topic and up to 65535 characters.
	Topic string `json:"topic"`
	// Type: Route type.
	Type RouteRouteType `json:"type"`
	// CreatedAt: Date at which the route was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// S3Config: When using S3 Route, S3-specific configuration fields.
	S3Config *RouteS3Config `json:"s3_config,omitempty"`
	// DbConfig: When using Database Route, DB-specific configuration fields.
	DbConfig *RouteDatabaseConfig `json:"db_config,omitempty"`
	// RestConfig: When using Rest Route, Rest-specific configuration fields.
	RestConfig *RouteRestConfig `json:"rest_config,omitempty"`
	// UpdatedAt: Date at which the route was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type SetDeviceCertificateRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// DeviceID: Device ID.
	DeviceID string `json:"-"`
	// CertificatePem: PEM-encoded custom certificate.
	CertificatePem string `json:"certificate_pem"`
}

type SetDeviceCertificateResponse struct {
	// Device:
	Device *Device `json:"device"`
	// CertificatePem:
	CertificatePem string `json:"certificate_pem"`
}

type SetHubCARequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// HubID: Hub ID.
	HubID string `json:"-"`
	// CaCertPem: CA's PEM-encoded certificate.
	CaCertPem string `json:"ca_cert_pem"`
	// ChallengeCertPem: Challenge is a PEM-encoded certificate that acts as proof of possession of the CA. It must be signed by the CA, and have a Common Name equal to the Hub ID.
	ChallengeCertPem string `json:"challenge_cert_pem"`
}

type TwinDocument struct {
	// TwinID: Parent twin ID of the document.
	TwinID string `json:"twin_id"`
	// DocumentName: Name of the document.
	DocumentName string `json:"document_name"`
	// Version: New version of the document.
	Version uint32 `json:"version"`
	// Data: New data related to the document.
	Data *scw.JSONObject `json:"data,omitempty"`
}

type UpdateDeviceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// DeviceID: Device ID.
	DeviceID string `json:"-"`
	// Description: Description for the device.
	Description *string `json:"description,omitempty"`
	// AllowInsecure: Defines whether to allow plain and server-authenticated SSL connections in addition to mutually-authenticated ones.
	AllowInsecure *bool `json:"allow_insecure,omitempty"`
	// AllowMultipleConnections: Defines whether to allow multiple physical devices to connect with this device's credentials.
	AllowMultipleConnections *bool `json:"allow_multiple_connections,omitempty"`
	// MessageFilters: Filter-sets to restrict the topics the device can publish/subscribe to.
	MessageFilters *DeviceMessageFilters `json:"message_filters"`
	// HubID: Change Hub for this device, additional fees may apply, see IoT Hub pricing.
	HubID *string `json:"hub_id,omitempty"`
}

type UpdateHubRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// HubID: ID of the Hub you want to update.
	HubID string `json:"-"`
	// Name: Hub name (up to 255 characters).
	Name *string `json:"name,omitempty"`
	// ProductPlan: Hub product plan.
	ProductPlan HubProductPlan `json:"product_plan"`
	// DisableEvents: Disable Hub events.
	DisableEvents *bool `json:"disable_events,omitempty"`
	// EventsTopicPrefix: Topic prefix of Hub events.
	EventsTopicPrefix *string `json:"events_topic_prefix,omitempty"`
	// EnableDeviceAutoProvisioning: Enable device auto provisioning.
	EnableDeviceAutoProvisioning *bool `json:"enable_device_auto_provisioning,omitempty"`
	// TwinsGraphiteConfig: BETA - not implemented yet.
	TwinsGraphiteConfig *HubTwinsGraphiteConfig `json:"twins_graphite_config,omitempty"`
}

type UpdateRouteRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// RouteID: Route id.
	RouteID string `json:"-"`
	// Name: Route name.
	Name *string `json:"name,omitempty"`
	// Topic: Topic the route subscribes to. It must be a valid MQTT topic and up to 65535 characters.
	Topic *string `json:"topic,omitempty"`
	// S3Config: When updating S3 Route, S3-specific configuration fields.
	S3Config *UpdateRouteRequestS3Config `json:"s3_config,omitempty"`
	// DbConfig: When updating Database Route, DB-specific configuration fields.
	DbConfig *UpdateRouteRequestDatabaseConfig `json:"db_config,omitempty"`
	// RestConfig: When updating Rest Route, Rest-specific configuration fields.
	RestConfig *UpdateRouteRequestRestConfig `json:"rest_config,omitempty"`
}

// Scaleway IoT Hub is a resilient message broker which supports multiple communication protocols and displays export capabilities. It is a PaaS (Platform-as-a-Service) enabling device-to-device and device-to-Scaleway communication.
//
// When connecting to a Hub, devices can exchange messages, granting them access to various cloud services. Hub routes also allow devices to send messages to other services that do not use MQTT (Message Queuing Telemetry Transport). MQTT is the most commonly used messaging protocol in IoT applications for its lightweight property and Publish/Subscribe model. Allowing interaction with services that do not use MQTT is an essential feature that increases flexibility and interoperability of the system.
//
// Scaleway IoT Hub is thus versatile and can easily adapt to the various use cases, requirements, and constraints of a given IoT project. IoT Hub features include :
//
// - Pub/Sub architecture
// - MQTT protocol, with or without TLS, with or without WebSockets
// - Mutual authentication available on TLS connections
// - High Availability and Scalability
// - Routes to Scaleway's ecosystem for data ingestion
// - Multiple other IoT Networks such as SigFox or LoRa
// - Usage metrics
//
// ## Concepts
//
// Refer to our [dedicated concepts page](https://www.scaleway.com/en/docs/managed-services/iot-hub/concepts/) to find definitions of the different terms referring to IoT Hub.
//
// ## Quickstart
//
// (switchcolumn)
// (switchcolumn)
//
// ### Set up your IoT Hub
//
//  1. Configure your environment variables.
//     <Message type="note">
//     This is an optional step that seeks to simplify your usage of the APIs. Since there is only one Availability Zone for IoT Hub, we have already set the value to the `fr-par` region for you.
//     </Message>
//
//     ```bash
//     export IOT_API="https://api.scaleway.com/iot/v1/regions/fr-par"
//     export SCW_ACCESS_KEY="<API access key>"
//     export SCW_SECRET_KEY="<API secret key>"
//     ```
//
//  2. Edit the POST request payload that we will use in the next step to create an IoT Hub.
//     ```
//     {
//     "name": "my_first_hub",
//     "product_plan": "plan_dedicated",
//     }
//     ```
//
//     | Parameter      | Description                                                                                                  |
//     | :------------- | :----------------------------------------------------------------------------------------------------------- |
//     | `project_id`   | ID of the Project to create your flexible IP in.                                                             |
//     | `name`         | Name of the Hub.                                                                                             |
//     | `product_plan` | Product plan for your IoT Hub. Three product plans are currently available for Hubs: `plan_shared`, `plan_dedicated` and `plan_ha`. To find which plan is better suited to your needs, you can consult the [IoT Hub Quickstart page](https://www.scaleway.com/en/docs/managed-services/iot-hub/quickstart/#how-to-create-an-iot-hub).                               |
//
//  3. Create a Hub. Run the following command to create the Hub. Make sure you include the payload you edited in the previous step.
//     <Message type="note">
//     We will save the output in a file called `myhub.json`.
//     </Message>
//
//     ```bash
//     curl -X POST \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     -H "Content-Type: application/json" \
//     -d '{
//     "project_id": "7281793f-8474-727d-7688-72893f747g7e",
//     "name": "my_first_hub",
//     "product_plan": "plan_dedicated"}' \
//     "$IOT_API/hubs" > myhub.json
//     ```
//
//  4. Get your Hub's information. Run the following command to pull your Hub's information from the `myhub.json` file.
//     ```bash
//     curl -X GET \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     "$IOT_API/hubs/$(jq -r '.id' myhub.json)"
//     ```
//
//     You should get an output similar to the following. The status should be indicated as `ready`.
//     <Message type="note">
//     This is a response example, the UUIDs displayed are not real.
//     </Message>
//
//     ```bash
//     {
//     "region":"fr-par",
//     "id":"6e38b739-2872-52fg-9gb7-e8gf8185be69",
//     "organization_id":"7281793f-8474-727d-7688-72893f747g7e",
//     "project_id":"7281793f-8474-727d-7688-72893f747g7e",
//     "name":"my_first_hub",
//     "status":"ready",
//     "product_plan":"plan_dedicated",
//     "endpoint":"iot.fr-par.scw.cloud",
//     "created_at":"2023-04-06T09:37:31.162Z",
//     "updated_at":"2023-04-06T09:37:31.162Z",
//     "enabled": true,
//     "device_count":0,
//     "connected_device_count":0,
//     "disable_events": false,
//     "events_topic_prefix":"$SCW/events",
//     "enable_device_auto_provisioning": false,
//     "has_custom_ca": false
//     }
//     ```
//
// ### Set up your IoT device
//
//  1. Edit the following POST request payload to create your first device.
//     <Message type="note">
//     For the sake of simplicity, the following payload allows the device to connect using insecure protocols (such as plain text MQTT or MQTTs without mutual authentication). In production, you should always deny insecure connections to ensure the highest level of security. This would be done by setting the field `allow_insecure` to `false`.
//     </Message>
//
//     ```
//     {
//     "name": "my-first-device",
//     "description": "This is a description for my first device that will act as subscriber"
//     "allow_insecure": "true",
//     }
//     ```
//     | Parameters       | Description                                                                |
//     | :--------------- | :------------------------------------------------------------------------- |
//     | `name`           | Device name                                                                |
//     | `description`    | A description of your device.                                              |
//     | `allow_insecure` | Whether or not you allow insecure connections. Values are either `true` or `false`.                                                                                        |
//
//  2. **Create your first device**. Run the following command to create your first device.
//     <Message type="note">
//     The response output will be saved in a file called `mydev1.json` and then piped to `jq < mydev1.json` to pretty-print the output in your terminal.
//     </Message>
//
//     ```bash
//     curl -X POST \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     -H "Content-Type: application/json" \
//     -d '{
//     "hub_id": "'$(jq -r '.id' myhub.json)'",
//     "name": "my_first_device",
//     "description": "This is a description for my first device that will act as subscriber",
//     "allow_insecure": true
//     }' \
//     "$IOT_API/devices" > mydev1.json
//     jq < mydev1.json
//     ```
//
//     You should get an output similar to the following, providing information about your device.
//     <Message type="note">
//     This is a response example, the UUIDs displayed are not real.
//     </Message>
//
//     ```bash
//     {
//     "device": {
//     "id": "98896561-9547-46a7-aa5b-8117a71a4b3c",
//     "name": "my_first_device",
//     "status": "enabled",
//     "hub_id": "6e38b739-2872-52fg-9gb7-e8gf8185be69",
//     "created_at": "2023-04-06T11:01:31.698Z",
//     "updated_at": "2023-04-06T11:01:31.698Z",
//     "allow_insecure": true,
//     "last_activity_at": "1970-01-01T00:00:00Z",
//     "is_connected": false,
//     "message_filters": {
//     "publish": {
//     "policy": "reject",
//     "topics": []
//     },
//     "subscribe": {
//     "policy": "reject",
//     "topics": []
//     }
//     },
//     "allow_multiple_connections": false,
//     "description": "This is a description for my first device that will act as subscriber",
//     "has_custom_certificate": false
//     },
//     "certificate": {
//     "crt": "<certificate here>",
//     "key": "<certificate key here>"
//     }
//     }
//     ```
//
//  3. Create your second device. Run the following command to create your second device.
//     ```bash
//     curl -X POST \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     -H "Content-Type: application/json" \
//     -d '{
//     "hub_id": "'$(jq -r '.id' myhub.json)'",
//     "name": "my_second_device",
//     "description": "This is a description for my second device that will act as publisher",
//     "allow_insecure": true
//     }' \
//     "$IOT_API/devices" > mydev2.json
//     jq < mydev2.json
//     ```
//
// ### Subscribe and publish
//
// In the next part, you will simulate two devices to indicate how to set them up as publishers and subscribers.
//
//  1. Set up the subscriber. Open a new terminal window and run the following command to subscribe to an MQTT topic.
//     ```bash
//     mosquitto_sub \
//     -h $(jq -r '.endpoint' myhub.json) \
//     -i $(jq -r '.device.id' mydev1.json) \
//     -t mytopic/mysubtopic
//     ```
//
//  2. Set up the publisher. Open another terminal window and run the following command to publish a `Hello, world!` message to an MQTT topic.
//     ```bash
//     mosquitto_pub \
//     -h $(jq -r '.endpoint' myhub.json) \
//     -i $(jq -r '.device.id' mydev2.json) \
//     -t mytopic/mysubtopic \
//     -m 'Hello, world!'
//     ```
//
//     When you open the terminal window you used to set up the subscriber, you should see that the `Hello, world!` message has successfuly been received.
//
// (switchcolumn)
// <Message type="requirement">
// To perform the following steps, you must first ensure that:
// - You have a [Scaleway account](https://console.scaleway.com/)
// - You have [created an API key](https://www.scaleway.com/en/docs/identity-and-access-management/iam/how-to/create-api-keys/) and that the API key has sufficient [IAM permissions](https://www.scaleway.com/en/docs/identity-and-access-management/iam/reference-content/permission-sets/) to perform the actions described on this page
// - You have [installed `curl`](https://curl.se/download.html)
// - You have [installed `mosquitto`](https://mosquitto.org/download/)
// - You have [installed `jq`](https://stedolan.github.io/jq/download/)
// </Message>
// (switchcolumn)
//
// ## Technical information
//
// ### Availability Zones
//
// Hubs can be created in the following Availability Zone:
//
// - Paris `fr-par`
//
// ## Going further
//
// For more help using Scaleway IoT Hub, check out the following resources:
// - Our [main documentation](https://www.scaleway.com/en/docs/managed-services/iot-hub/)
// - The #iot-hub channel on our [Slack Community](https://www.scaleway.com/en/docs/tutorials/scaleway-slack-community/)
// - Our [support ticketing system](https://www.scaleway.com/en/docs/console/my-account/how-to/open-a-support-ticket/).
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
	return []scw.Region{scw.RegionFrPar}
}

// ListHubs: List all Hubs in the specified zone. By default, returned Hubs are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListHubs(req *ListHubsRequest, opts ...scw.RequestOption) (*ListHubsResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs",
		Query:  query,
	}

	var resp ListHubsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateHub: Create a new Hub in the targeted region, specifying its configuration including name and product plan.
func (s *API) CreateHub(req *CreateHubRequest, opts ...scw.RequestOption) (*Hub, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("hub")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Hub

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHub: Retrieve information about an existing IoT Hub, specified by its Hub ID. Its full details, including name, status and endpoint, are returned in the response object.
func (s *API) GetHub(req *GetHubRequest, opts ...scw.RequestOption) (*Hub, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return nil, errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "",
	}

	var resp Hub

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHub: Update the parameters of an existing IoT Hub, specified by its Hub ID.
func (s *API) UpdateHub(req *UpdateHubRequest, opts ...scw.RequestOption) (*Hub, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return nil, errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Hub

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnableHub: Enable an existing IoT Hub, specified by its Hub ID.
func (s *API) EnableHub(req *EnableHubRequest, opts ...scw.RequestOption) (*Hub, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return nil, errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/enable",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Hub

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DisableHub: Disable an existing IoT Hub, specified by its Hub ID.
func (s *API) DisableHub(req *DisableHubRequest, opts ...scw.RequestOption) (*Hub, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return nil, errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/disable",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Hub

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteHub: Delete an existing IoT Hub, specified by its Hub ID. Deleting a Hub is permanent, and cannot be undone.
func (s *API) DeleteHub(req *DeleteHubRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "delete_devices", req.DeleteDevices)

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "",
		Query:  query,
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetHubMetrics: Get the metrics of an existing IoT Hub, specified by its Hub ID.
func (s *API) GetHubMetrics(req *GetHubMetricsRequest, opts ...scw.RequestOption) (*GetHubMetricsResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "start_date", req.StartDate)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return nil, errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/metrics",
		Query:  query,
	}

	var resp GetHubMetricsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetHubCA: Set a particular PEM-encoded certificate, specified by the Hub ID.
func (s *API) SetHubCA(req *SetHubCARequest, opts ...scw.RequestOption) (*Hub, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return nil, errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/ca",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Hub

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHubCA: Get information for a particular PEM-encoded certificate, specified by the Hub ID.
func (s *API) GetHubCA(req *GetHubCARequest, opts ...scw.RequestOption) (*GetHubCAResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return nil, errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/ca",
	}

	var resp GetHubCAResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDevices: List all devices in the specified region. By default, returned devices are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListDevices(req *ListDevicesRequest, opts ...scw.RequestOption) (*ListDevicesResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "hub_id", req.HubID)
	parameter.AddToQuery(query, "allow_insecure", req.AllowInsecure)
	parameter.AddToQuery(query, "status", req.Status)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices",
		Query:  query,
	}

	var resp ListDevicesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDevice: Attach a device to a given Hub.
func (s *API) CreateDevice(req *CreateDeviceRequest, opts ...scw.RequestOption) (*CreateDeviceResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("device")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateDeviceResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDevice: Retrieve information about an existing device, specified by its device ID. Its full details, including name, status and ID, are returned in the response object.
func (s *API) GetDevice(req *GetDeviceRequest, opts ...scw.RequestOption) (*Device, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeviceID) == "" {
		return nil, errors.New("field DeviceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "",
	}

	var resp Device

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateDevice: Update the parameters of an existing device, specified by its device ID.
func (s *API) UpdateDevice(req *UpdateDeviceRequest, opts ...scw.RequestOption) (*Device, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeviceID) == "" {
		return nil, errors.New("field DeviceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Device

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnableDevice: Enable a specific device, specified by its device ID.
func (s *API) EnableDevice(req *EnableDeviceRequest, opts ...scw.RequestOption) (*Device, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeviceID) == "" {
		return nil, errors.New("field DeviceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "/enable",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Device

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DisableDevice: Disable an existing device, specified by its device ID.
func (s *API) DisableDevice(req *DisableDeviceRequest, opts ...scw.RequestOption) (*Device, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeviceID) == "" {
		return nil, errors.New("field DeviceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "/disable",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Device

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RenewDeviceCertificate: Renew the certificate of an existing device, specified by its device ID.
func (s *API) RenewDeviceCertificate(req *RenewDeviceCertificateRequest, opts ...scw.RequestOption) (*RenewDeviceCertificateResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeviceID) == "" {
		return nil, errors.New("field DeviceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "/renew-certificate",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RenewDeviceCertificateResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetDeviceCertificate: Switch the existing certificate of a given device with an EM-encoded custom certificate.
func (s *API) SetDeviceCertificate(req *SetDeviceCertificateRequest, opts ...scw.RequestOption) (*SetDeviceCertificateResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeviceID) == "" {
		return nil, errors.New("field DeviceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "/certificate",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetDeviceCertificateResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDeviceCertificate: Get information for a particular PEM-encoded certificate, specified by the device ID. The response returns full details of the device, including its type of certificate.
func (s *API) GetDeviceCertificate(req *GetDeviceCertificateRequest, opts ...scw.RequestOption) (*GetDeviceCertificateResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeviceID) == "" {
		return nil, errors.New("field DeviceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "/certificate",
	}

	var resp GetDeviceCertificateResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDevice: Remove a specific device from the specific Hub it is attached to.
func (s *API) DeleteDevice(req *DeleteDeviceRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeviceID) == "" {
		return errors.New("field DeviceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetDeviceMetrics: Get the metrics of an existing device, specified by its device ID.
func (s *API) GetDeviceMetrics(req *GetDeviceMetricsRequest, opts ...scw.RequestOption) (*GetDeviceMetricsResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "start_date", req.StartDate)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeviceID) == "" {
		return nil, errors.New("field DeviceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "/metrics",
		Query:  query,
	}

	var resp GetDeviceMetricsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListRoutes: List all routes in the specified region. By default, returned routes are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "hub_id", req.HubID)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/routes",
		Query:  query,
	}

	var resp ListRoutesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRoute: Multiple kinds of routes can be created, such as:
//   - Database Route
//     Create a route that will record subscribed MQTT messages into your database.
//     <b>You need to manage the database by yourself</b>.
//   - REST Route.
//     Create a route that will call a REST API on received subscribed MQTT messages.
//   - S3 Routes.
//     Create a route that will put subscribed MQTT messages into an S3 bucket.
//     You need to create the bucket yourself and grant write access.
//     Granting can be done with s3cmd (`s3cmd setacl s3://<my-bucket> --acl-grant=write:555c69c3-87d0-4bf8-80f1-99a2f757d031:555c69c3-87d0-4bf8-80f1-99a2f757d031`).
func (s *API) CreateRoute(req *CreateRouteRequest, opts ...scw.RequestOption) (*Route, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("route")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/routes",
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

// UpdateRoute: Update the parameters of an existing route, specified by its route ID.
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
		Method: "PATCH",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/routes/" + fmt.Sprint(req.RouteID) + "",
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

// GetRoute: Get information for a particular route, specified by the route ID. The response returns full details of the route, including its type, the topic it subscribes to and its configuration.
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
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/routes/" + fmt.Sprint(req.RouteID) + "",
	}

	var resp Route

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteRoute: Delete an existing route, specified by its route ID. Deleting a route is permanent, and cannot be undone.
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
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/routes/" + fmt.Sprint(req.RouteID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListNetworks: List the networks.
func (s *API) ListNetworks(req *ListNetworksRequest, opts ...scw.RequestOption) (*ListNetworksResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "hub_id", req.HubID)
	parameter.AddToQuery(query, "topic_prefix", req.TopicPrefix)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/networks",
		Query:  query,
	}

	var resp ListNetworksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateNetwork: Create a new network for an existing hub.  Beside the default network, you can add networks for different data providers. Possible network types are Sigfox and REST.
func (s *API) CreateNetwork(req *CreateNetworkRequest, opts ...scw.RequestOption) (*CreateNetworkResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("network")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/networks",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateNetworkResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNetwork: Retrieve an existing network, specified by its network ID.  The response returns full details of the network, including its type, the topic prefix and its endpoint.
func (s *API) GetNetwork(req *GetNetworkRequest, opts ...scw.RequestOption) (*Network, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NetworkID) == "" {
		return nil, errors.New("field NetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/networks/" + fmt.Sprint(req.NetworkID) + "",
	}

	var resp Network

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteNetwork: Delete an existing network, specified by its network ID. Deleting a network is permanent, and cannot be undone.
func (s *API) DeleteNetwork(req *DeleteNetworkRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NetworkID) == "" {
		return errors.New("field NetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/networks/" + fmt.Sprint(req.NetworkID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetTwinDocument: BETA - Get a Cloud Twin Document.
func (s *API) GetTwinDocument(req *GetTwinDocumentRequest, opts ...scw.RequestOption) (*TwinDocument, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TwinID) == "" {
		return nil, errors.New("field TwinID cannot be empty in request")
	}

	if fmt.Sprint(req.DocumentName) == "" {
		return nil, errors.New("field DocumentName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/twins/" + fmt.Sprint(req.TwinID) + "/documents/" + fmt.Sprint(req.DocumentName) + "",
	}

	var resp TwinDocument

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PutTwinDocument: BETA - Update a Cloud Twin Document.
func (s *API) PutTwinDocument(req *PutTwinDocumentRequest, opts ...scw.RequestOption) (*TwinDocument, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TwinID) == "" {
		return nil, errors.New("field TwinID cannot be empty in request")
	}

	if fmt.Sprint(req.DocumentName) == "" {
		return nil, errors.New("field DocumentName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/twins/" + fmt.Sprint(req.TwinID) + "/documents/" + fmt.Sprint(req.DocumentName) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp TwinDocument

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PatchTwinDocument: BETA - Patch a Cloud Twin Document.
func (s *API) PatchTwinDocument(req *PatchTwinDocumentRequest, opts ...scw.RequestOption) (*TwinDocument, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TwinID) == "" {
		return nil, errors.New("field TwinID cannot be empty in request")
	}

	if fmt.Sprint(req.DocumentName) == "" {
		return nil, errors.New("field DocumentName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/twins/" + fmt.Sprint(req.TwinID) + "/documents/" + fmt.Sprint(req.DocumentName) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp TwinDocument

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteTwinDocument: BETA - Delete a Cloud Twin Document.
func (s *API) DeleteTwinDocument(req *DeleteTwinDocumentRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TwinID) == "" {
		return errors.New("field TwinID cannot be empty in request")
	}

	if fmt.Sprint(req.DocumentName) == "" {
		return errors.New("field DocumentName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/twins/" + fmt.Sprint(req.TwinID) + "/documents/" + fmt.Sprint(req.DocumentName) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListTwinDocuments: BETA - List the documents of a Cloud Twin.
func (s *API) ListTwinDocuments(req *ListTwinDocumentsRequest, opts ...scw.RequestOption) (*ListTwinDocumentsResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TwinID) == "" {
		return nil, errors.New("field TwinID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/twins/" + fmt.Sprint(req.TwinID) + "",
	}

	var resp ListTwinDocumentsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteTwinDocuments: BETA - Delete all the documents of a Cloud Twin.
func (s *API) DeleteTwinDocuments(req *DeleteTwinDocumentsRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TwinID) == "" {
		return errors.New("field TwinID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/twins/" + fmt.Sprint(req.TwinID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}
