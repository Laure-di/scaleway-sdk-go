// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package baremetal provides methods and message types of the baremetal v1 API.
package baremetal

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

type IPReverseStatus string

const (
	IPReverseStatusUnknown = IPReverseStatus("unknown")
	IPReverseStatusPending = IPReverseStatus("pending")
	IPReverseStatusActive  = IPReverseStatus("active")
	IPReverseStatusError   = IPReverseStatus("error")
)

func (enum IPReverseStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum IPReverseStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *IPReverseStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = IPReverseStatus(IPReverseStatus(tmp).String())
	return nil
}

type IPVersion string

const (
	IPVersionIPv4 = IPVersion("IPv4")
	IPVersionIPv6 = IPVersion("IPv6")
)

func (enum IPVersion) String() string {
	if enum == "" {
		// return default value if empty
		return "IPv4"
	}
	return string(enum)
}

func (enum IPVersion) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *IPVersion) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = IPVersion(IPVersion(tmp).String())
	return nil
}

type ListServerEventsRequestOrderBy string

const (
	ListServerEventsRequestOrderByCreatedAtAsc  = ListServerEventsRequestOrderBy("created_at_asc")
	ListServerEventsRequestOrderByCreatedAtDesc = ListServerEventsRequestOrderBy("created_at_desc")
)

func (enum ListServerEventsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListServerEventsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListServerEventsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListServerEventsRequestOrderBy(ListServerEventsRequestOrderBy(tmp).String())
	return nil
}

type ListServerPrivateNetworksRequestOrderBy string

const (
	ListServerPrivateNetworksRequestOrderByCreatedAtAsc  = ListServerPrivateNetworksRequestOrderBy("created_at_asc")
	ListServerPrivateNetworksRequestOrderByCreatedAtDesc = ListServerPrivateNetworksRequestOrderBy("created_at_desc")
	ListServerPrivateNetworksRequestOrderByUpdatedAtAsc  = ListServerPrivateNetworksRequestOrderBy("updated_at_asc")
	ListServerPrivateNetworksRequestOrderByUpdatedAtDesc = ListServerPrivateNetworksRequestOrderBy("updated_at_desc")
)

func (enum ListServerPrivateNetworksRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListServerPrivateNetworksRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListServerPrivateNetworksRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListServerPrivateNetworksRequestOrderBy(ListServerPrivateNetworksRequestOrderBy(tmp).String())
	return nil
}

type ListServersRequestOrderBy string

const (
	ListServersRequestOrderByCreatedAtAsc  = ListServersRequestOrderBy("created_at_asc")
	ListServersRequestOrderByCreatedAtDesc = ListServersRequestOrderBy("created_at_desc")
)

func (enum ListServersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListServersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListServersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListServersRequestOrderBy(ListServersRequestOrderBy(tmp).String())
	return nil
}

type ListSettingsRequestOrderBy string

const (
	ListSettingsRequestOrderByCreatedAtAsc  = ListSettingsRequestOrderBy("created_at_asc")
	ListSettingsRequestOrderByCreatedAtDesc = ListSettingsRequestOrderBy("created_at_desc")
)

func (enum ListSettingsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListSettingsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSettingsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSettingsRequestOrderBy(ListSettingsRequestOrderBy(tmp).String())
	return nil
}

type OfferStock string

const (
	OfferStockEmpty     = OfferStock("empty")
	OfferStockLow       = OfferStock("low")
	OfferStockAvailable = OfferStock("available")
)

func (enum OfferStock) String() string {
	if enum == "" {
		// return default value if empty
		return "empty"
	}
	return string(enum)
}

func (enum OfferStock) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OfferStock) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OfferStock(OfferStock(tmp).String())
	return nil
}

type OfferSubscriptionPeriod string

const (
	OfferSubscriptionPeriodUnknownSubscriptionPeriod = OfferSubscriptionPeriod("unknown_subscription_period")
	OfferSubscriptionPeriodHourly                    = OfferSubscriptionPeriod("hourly")
	OfferSubscriptionPeriodMonthly                   = OfferSubscriptionPeriod("monthly")
)

func (enum OfferSubscriptionPeriod) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_subscription_period"
	}
	return string(enum)
}

func (enum OfferSubscriptionPeriod) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OfferSubscriptionPeriod) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OfferSubscriptionPeriod(OfferSubscriptionPeriod(tmp).String())
	return nil
}

type ServerBootType string

const (
	ServerBootTypeUnknownBootType = ServerBootType("unknown_boot_type")
	ServerBootTypeNormal          = ServerBootType("normal")
	ServerBootTypeRescue          = ServerBootType("rescue")
)

func (enum ServerBootType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_boot_type"
	}
	return string(enum)
}

func (enum ServerBootType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerBootType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerBootType(ServerBootType(tmp).String())
	return nil
}

type ServerInstallStatus string

const (
	ServerInstallStatusUnknown    = ServerInstallStatus("unknown")
	ServerInstallStatusToInstall  = ServerInstallStatus("to_install")
	ServerInstallStatusInstalling = ServerInstallStatus("installing")
	ServerInstallStatusCompleted  = ServerInstallStatus("completed")
	ServerInstallStatusError      = ServerInstallStatus("error")
)

func (enum ServerInstallStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum ServerInstallStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerInstallStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerInstallStatus(ServerInstallStatus(tmp).String())
	return nil
}

type ServerOptionOptionStatus string

const (
	ServerOptionOptionStatusOptionStatusUnknown   = ServerOptionOptionStatus("option_status_unknown")
	ServerOptionOptionStatusOptionStatusEnable    = ServerOptionOptionStatus("option_status_enable")
	ServerOptionOptionStatusOptionStatusEnabling  = ServerOptionOptionStatus("option_status_enabling")
	ServerOptionOptionStatusOptionStatusDisabling = ServerOptionOptionStatus("option_status_disabling")
	ServerOptionOptionStatusOptionStatusError     = ServerOptionOptionStatus("option_status_error")
)

func (enum ServerOptionOptionStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "option_status_unknown"
	}
	return string(enum)
}

func (enum ServerOptionOptionStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerOptionOptionStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerOptionOptionStatus(ServerOptionOptionStatus(tmp).String())
	return nil
}

type ServerPingStatus string

const (
	ServerPingStatusPingStatusUnknown = ServerPingStatus("ping_status_unknown")
	ServerPingStatusPingStatusUp      = ServerPingStatus("ping_status_up")
	ServerPingStatusPingStatusDown    = ServerPingStatus("ping_status_down")
)

func (enum ServerPingStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "ping_status_unknown"
	}
	return string(enum)
}

func (enum ServerPingStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerPingStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerPingStatus(ServerPingStatus(tmp).String())
	return nil
}

type ServerPrivateNetworkStatus string

const (
	ServerPrivateNetworkStatusUnknown   = ServerPrivateNetworkStatus("unknown")
	ServerPrivateNetworkStatusAttaching = ServerPrivateNetworkStatus("attaching")
	ServerPrivateNetworkStatusAttached  = ServerPrivateNetworkStatus("attached")
	ServerPrivateNetworkStatusError     = ServerPrivateNetworkStatus("error")
	ServerPrivateNetworkStatusDetaching = ServerPrivateNetworkStatus("detaching")
	ServerPrivateNetworkStatusLocked    = ServerPrivateNetworkStatus("locked")
)

func (enum ServerPrivateNetworkStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum ServerPrivateNetworkStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerPrivateNetworkStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerPrivateNetworkStatus(ServerPrivateNetworkStatus(tmp).String())
	return nil
}

type ServerStatus string

const (
	ServerStatusUnknown    = ServerStatus("unknown")
	ServerStatusDelivering = ServerStatus("delivering")
	ServerStatusReady      = ServerStatus("ready")
	ServerStatusStopping   = ServerStatus("stopping")
	ServerStatusStopped    = ServerStatus("stopped")
	ServerStatusStarting   = ServerStatus("starting")
	ServerStatusError      = ServerStatus("error")
	ServerStatusDeleting   = ServerStatus("deleting")
	ServerStatusLocked     = ServerStatus("locked")
	ServerStatusOutOfStock = ServerStatus("out_of_stock")
	ServerStatusOrdered    = ServerStatus("ordered")
	ServerStatusResetting  = ServerStatus("resetting")
)

func (enum ServerStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum ServerStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerStatus(ServerStatus(tmp).String())
	return nil
}

type SettingType string

const (
	SettingTypeUnknown = SettingType("unknown")
	SettingTypeSMTP    = SettingType("smtp")
)

func (enum SettingType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum SettingType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SettingType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SettingType(SettingType(tmp).String())
	return nil
}

// OSOSField:
type OSOSField struct {
	// Editable:
	Editable bool `json:"editable"`
	// Required:
	Required bool `json:"required"`
	// DefaultValue:
	DefaultValue *string `json:"default_value,omitempty"`
}

// CPU:
type CPU struct {
	// Name: Name of the CPU.
	Name string `json:"name"`
	// CoreCount: Number of CPU cores.
	CoreCount uint32 `json:"core_count"`
	// ThreadCount: Number CPU threads.
	ThreadCount uint32 `json:"thread_count"`
	// Frequency: Frequency of the CPU in MHz.
	Frequency uint32 `json:"frequency"`
	// Benchmark: Benchmark of the CPU.
	Benchmark string `json:"benchmark"`
}

// Disk:
type Disk struct {
	// Capacity: Capacity of the disk in bytes.
	Capacity scw.Size `json:"capacity"`
	// Type: Type of the disk.
	Type string `json:"type"`
}

// Memory:
type Memory struct {
	// Capacity: Capacity of the memory in bytes.
	Capacity scw.Size `json:"capacity"`
	// Type: Type of the memory.
	Type string `json:"type"`
	// Frequency: Frequency of the memory in MHz.
	Frequency uint32 `json:"frequency"`
	// IsEcc: True if the memory is an error-correcting code memory.
	IsEcc bool `json:"is_ecc"`
}

// OfferOptionOffer:
type OfferOptionOffer struct {
	// ID: ID of the option.
	ID string `json:"id"`
	// Name: Name of the option.
	Name string `json:"name"`
	// Enabled: If true the option is enabled and included by default in the offer
	// If false the option is available for the offer but not included by default.
	Enabled bool `json:"enabled"`
	// SubscriptionPeriod: Period of subscription for the offer.
	SubscriptionPeriod OfferSubscriptionPeriod `json:"subscription_period"`
	// Price: Price of the option.
	Price *scw.Money `json:"price,omitempty"`
	// Manageable: Boolean to know if option could be managed.
	Manageable bool `json:"manageable"`
	// OsID: ID of the OS linked to the option.
	OsID *string `json:"os_id,omitempty"`
}

// PersistentMemory:
type PersistentMemory struct {
	// Capacity: Capacity of the memory in bytes.
	Capacity scw.Size `json:"capacity"`
	// Type: Type of the memory.
	Type string `json:"type"`
	// Frequency: Frequency of the memory in MHz.
	Frequency uint32 `json:"frequency"`
}

// RaidController:
type RaidController struct {
	// Model:
	Model string `json:"model"`
	// RaidLevel:
	RaidLevel []string `json:"raid_level"`
}

// IP:
type IP struct {
	// ID: ID of the IP.
	ID string `json:"id"`
	// Address: Address of the IP.
	Address net.IP `json:"address"`
	// Reverse: Reverse IP value.
	Reverse string `json:"reverse"`
	// Version: Version of IP (v4 or v6).
	Version IPVersion `json:"version"`
	// ReverseStatus: Status of the reverse.
	ReverseStatus IPReverseStatus `json:"reverse_status"`
	// ReverseStatusMessage: A message related to the reverse status, e.g. in case of an error.
	ReverseStatusMessage string `json:"reverse_status_message"`
}

// ServerInstall:
type ServerInstall struct {
	// OsID: ID of the OS.
	OsID string `json:"os_id"`
	// Hostname: Host defined during the server installation.
	Hostname string `json:"hostname"`
	// SSHKeyIDs: SSH public key IDs defined during server installation.
	SSHKeyIDs []string `json:"ssh_key_ids"`
	// Status: Status of the server installation.
	Status ServerInstallStatus `json:"status"`
	// User: User defined in the server installation, or the default user if none were specified.
	User string `json:"user"`
	// ServiceUser: Service user defined in the server installation, or the default user if none were specified.
	ServiceUser string `json:"service_user"`
	// ServiceURL: Address of the installed service.
	ServiceURL string `json:"service_url"`
}

// ServerOption:
type ServerOption struct {
	// ID: ID of the option.
	ID string `json:"id"`
	// Name: Name of the option.
	Name string `json:"name"`
	// Status: Status of the option on this server.
	Status ServerOptionOptionStatus `json:"status"`
	// Manageable: Defines whether the option can be managed (added or removed).
	Manageable bool `json:"manageable"`
	// ExpiresAt: Auto expiration date for compatible options.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// ServerRescueServer:
type ServerRescueServer struct {
	// User: Rescue user name.
	User string `json:"user"`
	// Password: Rescue password.
	Password string `json:"password"`
}

// CreateServerRequestInstall:
type CreateServerRequestInstall struct {
	// OsID: ID of the OS to installation on the server.
	OsID string `json:"os_id"`
	// Hostname: Hostname of the server.
	Hostname string `json:"hostname"`
	// SSHKeyIDs: SSH key IDs authorized on the server.
	SSHKeyIDs []string `json:"ssh_key_ids"`
	// User: User for the installation.
	User *string `json:"user,omitempty"`
	// Password: Password for the installation.
	Password *string `json:"password,omitempty"`
	// ServiceUser: Regular user that runs the service to be installed on the server.
	ServiceUser *string `json:"service_user,omitempty"`
	// ServicePassword: Password used for the service to install.
	ServicePassword *string `json:"service_password,omitempty"`
	// UserData: Cloud-init file.
	UserData *scw.File `json:"user_data,omitempty"`
}

// OS:
type OS struct {
	// ID: ID of the OS.
	ID string `json:"id"`
	// Name: Name of the OS.
	Name string `json:"name"`
	// Version: Version of the OS.
	Version string `json:"version"`
	// LogoURL: URL of this OS's logo.
	LogoURL string `json:"logo_url"`
	// SSH: Object defining the SSH requirements to install the OS.
	SSH *OSOSField `json:"ssh"`
	// User: Object defining the username requirements to install the OS.
	User *OSOSField `json:"user"`
	// Password: Object defining the password requirements to install the OS.
	Password *OSOSField `json:"password"`
	// ServiceUser: Object defining the username requirements to install the service.
	ServiceUser *OSOSField `json:"service_user"`
	// ServicePassword: Object defining the password requirements to install the service.
	ServicePassword *OSOSField `json:"service_password"`
	// Enabled: Defines if the operating system is enabled or not.
	Enabled bool `json:"enabled"`
	// LicenseRequired: License required (check server options for pricing details).
	LicenseRequired bool `json:"license_required"`
	// Allowed: Defines if a specific Organization is allowed to install this OS type.
	Allowed bool `json:"allowed"`
}

// Offer:
type Offer struct {
	// ID: ID of the offer.
	ID string `json:"id"`
	// Name: Name of the offer.
	Name string `json:"name"`
	// Stock: Stock level.
	Stock OfferStock `json:"stock"`
	// Bandwidth: Public bandwidth available (in bits/s) with the offer.
	Bandwidth uint64 `json:"bandwidth"`
	// CommercialRange: Commercial range of the offer.
	CommercialRange string `json:"commercial_range"`
	// PricePerHour: Price of the offer for the next 60 minutes (a server order at 11h32 will be payed until 12h32).
	PricePerHour *scw.Money `json:"price_per_hour,omitempty"`
	// PricePerMonth: Monthly price of the offer, if subscribing on a monthly basis.
	PricePerMonth *scw.Money `json:"price_per_month,omitempty"`
	// Disks: Disks specifications of the offer.
	Disks []*Disk `json:"disks"`
	// Enable: Defines whether the offer is currently available.
	Enable bool `json:"enable"`
	// CPUs: CPU specifications of the offer.
	CPUs []*CPU `json:"cpus"`
	// Memories: Memory specifications of the offer.
	Memories []*Memory `json:"memories"`
	// QuotaName: Name of the quota associated to the offer.
	QuotaName string `json:"quota_name"`
	// PersistentMemories: Persistent memory specifications of the offer.
	PersistentMemories []*PersistentMemory `json:"persistent_memories"`
	// RaidControllers: Raid controller specifications of the offer.
	RaidControllers []*RaidController `json:"raid_controllers"`
	// IncompatibleOsIDs: Array of OS images IDs incompatible with the server.
	IncompatibleOsIDs []string `json:"incompatible_os_ids"`
	// SubscriptionPeriod: Period of subscription for the offer.
	SubscriptionPeriod OfferSubscriptionPeriod `json:"subscription_period"`
	// OperationPath: Operation path of the service.
	OperationPath string `json:"operation_path"`
	// Fee: One time fee invoiced by Scaleway for the setup and activation of the server.
	Fee *scw.Money `json:"fee,omitempty"`
	// Options: Available options for customization of the server.
	Options []*OfferOptionOffer `json:"options"`
	// PrivateBandwidth: Private bandwidth available in bits/s with the offer.
	PrivateBandwidth uint64 `json:"private_bandwidth"`
	// SharedBandwidth: Defines whether the offer's bandwidth is shared or not.
	SharedBandwidth bool `json:"shared_bandwidth"`
	// Tags: Array of tags attached to the offer.
	Tags []string `json:"tags"`
}

// Option:
type Option struct {
	// ID: ID of the option.
	ID string `json:"id"`
	// Name: Name of the option.
	Name string `json:"name"`
	// Manageable: Defines whether the option is manageable (could be added or removed).
	Manageable bool `json:"manageable"`
}

// ServerEvent:
type ServerEvent struct {
	// ID: ID of the server to which the action will be applied.
	ID string `json:"id"`
	// Action: The action that will be applied to the server.
	Action string `json:"action"`
	// UpdatedAt: Date of last modification of the action.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// CreatedAt: Date of creation of the action.
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// ServerPrivateNetwork:
type ServerPrivateNetwork struct {
	// ID: The Private Network ID.
	ID string `json:"id"`
	// ProjectID: The Private Network Project ID.
	ProjectID string `json:"project_id"`
	// ServerID: The server ID.
	ServerID string `json:"server_id"`
	// PrivateNetworkID: The Private Network ID.
	PrivateNetworkID string `json:"private_network_id"`
	// Vlan: The VLAN ID associated to the Private Network.
	Vlan *uint32 `json:"vlan,omitempty"`
	// Status: The configuration status of the Private Network.
	Status ServerPrivateNetworkStatus `json:"status"`
	// CreatedAt: The Private Network creation date.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: The date the Private Network was last modified.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// Server:
type Server struct {
	// ID: ID of the server.
	ID string `json:"id"`
	// OrganizationID: Organization ID the server is attached to.
	OrganizationID string `json:"organization_id"`
	// ProjectID: Project ID the server is attached to.
	ProjectID string `json:"project_id"`
	// Name: Name of the server.
	Name string `json:"name"`
	// Description: Description of the server.
	Description string `json:"description"`
	// UpdatedAt: Last modification date of the server.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// CreatedAt: Creation date of the server.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Status: Status of the server.
	Status ServerStatus `json:"status"`
	// OfferID: Offer ID of the server.
	OfferID string `json:"offer_id"`
	// OfferName: Offer name of the server.
	OfferName string `json:"offer_name"`
	// Tags: Array of custom tags attached to the server.
	Tags []string `json:"tags"`
	// IPs: Array of IPs attached to the server.
	IPs []*IP `json:"ips"`
	// Domain: Domain of the server.
	Domain string `json:"domain"`
	// BootType: Boot type of the server.
	BootType ServerBootType `json:"boot_type"`
	// Zone: Zone in which is the server located.
	Zone scw.Zone `json:"zone"`
	// Install: Configuration of the installation.
	Install *ServerInstall `json:"install"`
	// PingStatus: Status of server ping.
	PingStatus ServerPingStatus `json:"ping_status"`
	// Options: Options enabled on the server.
	Options []*ServerOption `json:"options"`
	// RescueServer: Configuration of rescue boot.
	RescueServer *ServerRescueServer `json:"rescue_server"`
}

// Setting:
type Setting struct {
	// ID: ID of the setting.
	ID string `json:"id"`
	// Type: Type of the setting.
	Type SettingType `json:"type"`
	// ProjectID: ID of the Project ID.
	ProjectID string `json:"project_id"`
	// Enabled: Defines whether the setting is enabled.
	Enabled bool `json:"enabled"`
}

// AddOptionServerRequest:
type AddOptionServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server.
	ServerID string `json:"-"`
	// OptionID: ID of the option to add.
	OptionID string `json:"-"`
	// ExpiresAt: Auto expire the option after this date.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// BMCAccess:
type BMCAccess struct {
	// URL: URL to access to the server console.
	URL string `json:"url"`
	// Login: The login to use for the BMC (Baseboard Management Controller) access authentification.
	Login string `json:"login"`
	// Password: The password to use for the BMC (Baseboard Management Controller) access authentification.
	Password string `json:"password"`
	// ExpiresAt: The date after which the BMC (Baseboard Management Controller) access will be closed.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// CreateServerRequest:
type CreateServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// OfferID: Offer ID of the new server.
	OfferID string `json:"offer_id"`
	// Deprecated: OrganizationID: Organization ID with which the server will be created.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: Project ID with which the server will be created.
	ProjectID *string `json:"project_id,omitempty"`
	// Name: Name of the server (≠hostname).
	Name string `json:"name"`
	// Description: Description associated with the server, max 255 characters.
	Description string `json:"description"`
	// Tags: Tags to associate to the server.
	Tags []string `json:"tags"`
	// Install: Object describing the configuration details of the OS installation on the server.
	Install *CreateServerRequestInstall `json:"install"`
	// OptionIDs: IDs of options to enable on server.
	OptionIDs []string `json:"option_ids"`
}

// DeleteOptionServerRequest:
type DeleteOptionServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server.
	ServerID string `json:"-"`
	// OptionID: ID of the option to delete.
	OptionID string `json:"-"`
}

// DeleteServerRequest:
type DeleteServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server to delete.
	ServerID string `json:"-"`
}

// GetBMCAccessRequest:
type GetBMCAccessRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server.
	ServerID string `json:"-"`
}

// GetOSRequest:
type GetOSRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// OsID: ID of the OS.
	OsID string `json:"-"`
}

// GetOfferRequest:
type GetOfferRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// OfferID: ID of the researched Offer.
	OfferID string `json:"-"`
}

// GetOptionRequest:
type GetOptionRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// OptionID: ID of the option.
	OptionID string `json:"-"`
}

// GetServerMetricsRequest:
type GetServerMetricsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: Server ID to get the metrics.
	ServerID string `json:"-"`
}

// GetServerMetricsResponse:
type GetServerMetricsResponse struct {
	// Pings: Timeseries object representing pings on the server.
	Pings *scw.TimeSeries `json:"pings,omitempty"`
}

// GetServerRequest:
type GetServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server.
	ServerID string `json:"-"`
}

// InstallServerRequest:
type InstallServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: Server ID to install.
	ServerID string `json:"-"`
	// OsID: ID of the OS to installation on the server.
	OsID string `json:"os_id"`
	// Hostname: Hostname of the server.
	Hostname string `json:"hostname"`
	// SSHKeyIDs: SSH key IDs authorized on the server.
	SSHKeyIDs []string `json:"ssh_key_ids"`
	// User: User used for the installation.
	User *string `json:"user,omitempty"`
	// Password: Password used for the installation.
	Password *string `json:"password,omitempty"`
	// ServiceUser: User used for the service to install.
	ServiceUser *string `json:"service_user,omitempty"`
	// ServicePassword: Password used for the service to install.
	ServicePassword *string `json:"service_password,omitempty"`
	// UserData: Cloud-init file.
	UserData *scw.File `json:"user_data,omitempty"`
}

// ListOSRequest:
type ListOSRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Page: Page number.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Number of OS per page.
	PageSize *uint32 `json:"page_size,omitempty"`
	// OfferID: Offer IDs to filter OSes for.
	OfferID *string `json:"offer_id,omitempty"`
}

// ListOSResponse:
type ListOSResponse struct {
	// TotalCount: Total count of matching OS.
	TotalCount uint32 `json:"total_count"`
	// Os: OS that match filters.
	Os []*OS `json:"os"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListOSResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListOSResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListOSResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Os = append(r.Os, results.Os...)
	r.TotalCount += uint32(len(results.Os))
	return uint32(len(results.Os)), nil
}

// ListOffersRequest:
type ListOffersRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Page: Page number.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Number of offers per page.
	PageSize *uint32 `json:"page_size,omitempty"`
	// SubscriptionPeriod: Subscription period type to filter offers by.
	SubscriptionPeriod OfferSubscriptionPeriod `json:"subscription_period"`
}

// ListOffersResponse:
type ListOffersResponse struct {
	// TotalCount: Total count of matching offers.
	TotalCount uint32 `json:"total_count"`
	// Offers: Offers that match filters.
	Offers []*Offer `json:"offers"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListOffersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListOffersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListOffersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Offers = append(r.Offers, results.Offers...)
	r.TotalCount += uint32(len(results.Offers))
	return uint32(len(results.Offers)), nil
}

// ListOptionsRequest:
type ListOptionsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Page: Page number.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Number of options per page.
	PageSize *uint32 `json:"page_size,omitempty"`
	// OfferID: Offer ID to filter options for.
	OfferID *string `json:"offer_id,omitempty"`
	// Name: Name to filter options for.
	Name *string `json:"name,omitempty"`
}

// ListOptionsResponse:
type ListOptionsResponse struct {
	// TotalCount: Total count of matching options.
	TotalCount uint32 `json:"total_count"`
	// Options: Options that match filters.
	Options []*Option `json:"options"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListOptionsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListOptionsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListOptionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Options = append(r.Options, results.Options...)
	r.TotalCount += uint32(len(results.Options))
	return uint32(len(results.Options)), nil
}

// ListServerEventsRequest:
type ListServerEventsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server events searched.
	ServerID string `json:"-"`
	// Page: Page number.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Number of server events per page.
	PageSize *uint32 `json:"page_size,omitempty"`
	// OrderBy: Order of the server events.
	OrderBy ListServerEventsRequestOrderBy `json:"order_by"`
}

// ListServerEventsResponse:
type ListServerEventsResponse struct {
	// TotalCount: Total count of matching events.
	TotalCount uint32 `json:"total_count"`
	// Events: Server events that match filters.
	Events []*ServerEvent `json:"events"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServerEventsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServerEventsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListServerEventsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Events = append(r.Events, results.Events...)
	r.TotalCount += uint32(len(results.Events))
	return uint32(len(results.Events)), nil
}

// ListServerPrivateNetworksResponse:
type ListServerPrivateNetworksResponse struct {
	// ServerPrivateNetworks:
	ServerPrivateNetworks []*ServerPrivateNetwork `json:"server_private_networks"`
	// TotalCount:
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServerPrivateNetworksResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServerPrivateNetworksResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListServerPrivateNetworksResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ServerPrivateNetworks = append(r.ServerPrivateNetworks, results.ServerPrivateNetworks...)
	r.TotalCount += uint32(len(results.ServerPrivateNetworks))
	return uint32(len(results.ServerPrivateNetworks)), nil
}

// ListServersRequest:
type ListServersRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Page: Page number.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Number of servers per page.
	PageSize *uint32 `json:"page_size,omitempty"`
	// OrderBy: Order of the servers.
	OrderBy ListServersRequestOrderBy `json:"order_by"`
	// Tags: Tags to filter for.
	Tags []string `json:"tags"`
	// Status: Status to filter for.
	Status []string `json:"status"`
	// Name: Names to filter for.
	Name *string `json:"name,omitempty"`
	// OrganizationID: Organization ID to filter for.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: Project ID to filter for.
	ProjectID *string `json:"project_id,omitempty"`
	// OptionID: Option ID to filter for.
	OptionID *string `json:"option_id,omitempty"`
}

// ListServersResponse:
type ListServersResponse struct {
	// TotalCount: Total count of matching servers.
	TotalCount uint32 `json:"total_count"`
	// Servers: Array of Elastic Metal server objects matching the filters in the request.
	Servers []*Server `json:"servers"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListServersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Servers = append(r.Servers, results.Servers...)
	r.TotalCount += uint32(len(results.Servers))
	return uint32(len(results.Servers)), nil
}

// ListSettingsRequest:
type ListSettingsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Page: Page number.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Set the maximum list size.
	PageSize *uint32 `json:"page_size,omitempty"`
	// OrderBy: Sort order for items in the response.
	OrderBy ListSettingsRequestOrderBy `json:"order_by"`
	// ProjectID: ID of the Project.
	ProjectID *string `json:"project_id,omitempty"`
}

// ListSettingsResponse:
type ListSettingsResponse struct {
	// TotalCount: Total count of matching settings.
	TotalCount uint32 `json:"total_count"`
	// Settings: Settings that match filters.
	Settings []*Setting `json:"settings"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSettingsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSettingsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSettingsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Settings = append(r.Settings, results.Settings...)
	r.TotalCount += uint32(len(results.Settings))
	return uint32(len(results.Settings)), nil
}

// PrivateNetworkAPIAddServerPrivateNetworkRequest:
type PrivateNetworkAPIAddServerPrivateNetworkRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: The ID of the server.
	ServerID string `json:"-"`
	// PrivateNetworkID: The ID of the Private Network.
	PrivateNetworkID string `json:"private_network_id"`
}

// PrivateNetworkAPIDeleteServerPrivateNetworkRequest:
type PrivateNetworkAPIDeleteServerPrivateNetworkRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: The ID of the server.
	ServerID string `json:"-"`
	// PrivateNetworkID: The ID of the Private Network.
	PrivateNetworkID string `json:"-"`
}

// PrivateNetworkAPIListServerPrivateNetworksRequest:
type PrivateNetworkAPIListServerPrivateNetworksRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// OrderBy: The sort order for the returned Private Networks.
	OrderBy ListServerPrivateNetworksRequestOrderBy `json:"order_by"`
	// Page: The page number for the returned Private Networks.
	Page *int32 `json:"page,omitempty"`
	// PageSize: The maximum number of Private Networks per page.
	PageSize *uint32 `json:"page_size,omitempty"`
	// ServerID: Filter Private Networks by server ID.
	ServerID *string `json:"server_id,omitempty"`
	// PrivateNetworkID: Filter Private Networks by Private Network ID.
	PrivateNetworkID *string `json:"private_network_id,omitempty"`
	// OrganizationID: Filter Private Networks by Organization ID.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: Filter Private Networks by Project ID.
	ProjectID *string `json:"project_id,omitempty"`
}

// PrivateNetworkAPISetServerPrivateNetworksRequest:
type PrivateNetworkAPISetServerPrivateNetworksRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: The ID of the server.
	ServerID string `json:"-"`
	// PrivateNetworkIDs: The IDs of the Private Networks.
	PrivateNetworkIDs []string `json:"private_network_ids"`
}

// RebootServerRequest:
type RebootServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server to reboot.
	ServerID string `json:"-"`
	// BootType: The type of boot.
	BootType ServerBootType `json:"boot_type"`
}

// SetServerPrivateNetworksResponse:
type SetServerPrivateNetworksResponse struct {
	// ServerPrivateNetworks:
	ServerPrivateNetworks []*ServerPrivateNetwork `json:"server_private_networks"`
}

// StartBMCAccessRequest:
type StartBMCAccessRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server.
	ServerID string `json:"-"`
	// IP: The IP authorized to connect to the server.
	IP net.IP `json:"ip"`
}

// StartServerRequest:
type StartServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server to start.
	ServerID string `json:"-"`
	// BootType: The type of boot.
	BootType ServerBootType `json:"boot_type"`
}

// StopBMCAccessRequest:
type StopBMCAccessRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server.
	ServerID string `json:"-"`
}

// StopServerRequest:
type StopServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server to stop.
	ServerID string `json:"-"`
}

// UpdateIPRequest:
type UpdateIPRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server.
	ServerID string `json:"-"`
	// IPID: ID of the IP to update.
	IPID string `json:"-"`
	// Reverse: New reverse IP to update, not updated if null.
	Reverse *string `json:"reverse,omitempty"`
}

// UpdateServerRequest:
type UpdateServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server to update.
	ServerID string `json:"-"`
	// Name: Name of the server (≠hostname), not updated if null.
	Name *string `json:"name,omitempty"`
	// Description: Description associated with the server, max 255 characters, not updated if null.
	Description *string `json:"description,omitempty"`
	// Tags: Tags associated with the server, not updated if null.
	Tags *[]string `json:"tags,omitempty"`
}

// UpdateSettingRequest:
type UpdateSettingRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// SettingID: ID of the setting.
	SettingID string `json:"-"`
	// Enabled: Defines whether the setting is enabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// Scaleway Elastic Metal servers are dedicated physical servers that you can order on-demand, like Instances.
// You can install an OS or other images on your Elastic Metal server and connect to it via SSH to use it as you require.
// You can power off the server when you are not using or delete it from your account once you have finished using it.
// Elastic Metal servers are ideal for large workloads, big data, and applications that require increased security and dedicated resources.
//
// (switchcolumn)
// <Message type="tip">
//
//	Check out our dedicated APIs to manage [Private Networks](https://www.scaleway.com/en/developers/api/elastic-metal/private-network-api/) and [Flexible IPs](https://www.scaleway.com/en/developers/api/elastic-metal-flexible-ip) for Elastic Metal servers.
//
// </Message>
// (switchcolumn)
//
// ## Concepts
//
// Refer to our [dedicated concepts](https://www.scaleway.com/en/docs/compute/elastic-metal/concepts/) page to find definitions of the different terms referring to Elastic Metal servers.
//
// ## Quickstart
//
// (switchcolumn)
// (switchcolumn)
//
//  1. **Configure your environment variables.**
//     ```bash
//     export PROJECT_ID="<project-id>"
//     export ACCESS_KEY="<access-key>"
//     export SECRET_KEY="<secret-key>"
//     export ZONE="<availability-zone>"
//     ```
//     <Message type="note">
//     This is an optional step that seeks to simplify your usage of the Bare Metal API.
//     </Message>
//
//  2. **Edit the POST request payload** that we will use in the next step to create an Elastic Metal server. Modify the values in the example according to your needs, using the information in the table to help.
//     ```json
//     {
//     "offer_id": "string",
//     "project_id": "string",
//     "name": "string",
//     "description": "string",
//     "tags": [
//     "tag1", "tag2"
//     ],
//     "install": {
//     "os_id": "string",
//     "hostname": "string",
//     "ssh_key_ids": [
//     "string"
//     ],
//     "user": "string",
//     "password": "string",
//     "service_user": "string",
//     "service_password": "string"
//     },
//     "option_ids": [
//     "string"
//     ]
//     }
//     ```
//
//     | Parameter        | Description                                                        |
//     | :--------------- | :----------------------------------------------------------------- |
//     | `offer_id`           | **REQUIRED** UUID of the Elastic Metal offer                                         |
//     | `project_id`     | **REQUIRED** UUID of the project you want to create your Elastic Metal in.  |
//     | `name`           | **REQUIRED** Name of the Elastic Metal server (≠hostname)                                          |
//     | `description`     | **REQUIRED** A description of your server (max 255 characters)                             |
//     | `tags`  | **OPTIONAL** An array of tags associated with your server   |
//     | `os_id`  | The ID of the operating system image to install on the server   |
//     | `hostname`  | Hostname of the server   |
//     | `ssh_key_ids`  | SSH key IDs authorized on the server   |
//     | `user`  | **NULLABLE** A regular user to be configured on the server   |
//     | `password`  | **NULLABLE** The password for the user account   |
//     | `service_user`  | **NULLABLE** A service user for third party services (user to login in services such as BigBlueButton)  |
//     | `service password`  | **NULLABLE** Password for the service user   |
//     | `option_ids`  | IDs of options to enable on server  |
//
//     <Message type="tip">
//     To find your Project ID you can either use the [IAM API](https://www.scaleway.com/en/developers/api/account#path-projects-list-all-projects-of-an-organization) or the [Scaleway console](https://console.scaleway.com/project/settings):
//     </Message>
//
//  3. **Run the following command** to create an Elastic Metal server. Make sure you include the payload you edited in the previous step.
//     ```bash
//     curl -X POST \
//     -H "Content-Type: application/json" \
//     -H "X-Auth-Token: $SECRET_KEY" https://api.scaleway.com//baremetal/v1/zones/$ZONE/servers \
//     -d '{
//     "offer_id": "bd757ca3-a71b-4158-9ef5-39436b6db2a4",
//     "project_id": "cc6d123a-bc09-4e24-b5d9-3310f2104e13",
//     "name": "MyElasticMetal",
//     "description": "My_Elastic_Metal_Server",
//     "tags": [
//     "Ubuntu22", "Dedicated"
//     ],
//     "install": {
//     "os_id": "96e5f0f2-d216-4de2-8a15-68730d877885",
//     "hostname": "elasticmetal.exaple.com",
//     "ssh_key_ids": [
//     "fa05e77f-66b7-43b9-bc21-4dfe3c5bb624"
//     ],
//     "user": "ubuntu",
//     "password": "mySecretPa$$word"
//     "option_ids": [
//     "string"
//     ]
//     }"
//     ```
//  4. **List your Elastic Metal servers.**
//     ```bash
//     curl -X GET \
//     -H "Content-Type: application/json" \
//     -H "X-Auth-Token: $SECRET_KEY" https://api.scaleway.com/baremetal/v1/zones/$ZONE/servers
//     ```
//
// 5. **Retrieve your Elastic Metal server IP** from the response.
//
//  6. **Connect to your Elastic Metal server** using SSH
//     ```bash
//     ssh root@my-elastic-metal-server-ip
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
//
// ## Technical information
//
// ### Features
//
// - Installation (Server is installed by preseed [preseed: complete install from a virtual media], you must define at least one ssh key to install your server)
// - Start/Stop/Reboot
// - Rescue Reboot, a rescue image is an operating system image designed to help you diagnose and fix OS experiencing failures. When your server boot on rescue, you can mount your disks and start diagnosing/fixing your image.
// - Billed by the minute (billing starts when the server is delivered and stops when the server is deleted)
// - IPv6, all servers are available with a /128 IPv6 subnet
// - ReverseIP, You can configure your reverse IP (IPv4 and IPv6), you must register the server IP in your DNS records before calling the endpoint
// - Basic monitoring with ping status
// - Flexible IP is available ([documentation](https://www.scaleway.com/en/developers/api/elastic-metal-flexible-ip))
//
// ### Availability Zones
//
// Scaleway's infrastructure is spread across different [regions and Availability Zones](https://www.scaleway.com/en/docs/console/my-account/reference-content/products-availability/).
//
// Elastic Metal servers are available in the Paris and Amsterdam regions, with product availability in the following AZs:
// - `fr-par-1`
// - `fr-par-2`
// - `nl-ams-1`
//
// ## Technical limitations
//
// - Failover IPs are not available in API `v1`, use the API `v1alpha1`
// ## Going further
//
// For more help using Scaleway Elastic Metal servers, check out the following resources:
// - Our [main documentation](https://www.scaleway.com/en/docs/compute/elastic-metal/)
// - The #elastic-metal channel on our [Slack Community](https://www.scaleway.com/en/docs/tutorials/scaleway-slack-community/)
// - Our [support ticketing system](https://www.scaleway.com/en/docs/console/my-account/how-to/open-a-support-ticket)
// ### Troubleshooting
//
// #### How is the installation of Elastic Metal servers done?
//
// - The installation of Elastic Metal servers is done by preseed (± 10min) (preseed: complete install from a virtual media)
// #### How can I retrieve my secret key and Project ID?
//
// You can find your [secret key](https://console.scaleway.com/iam/api-keys) and your [Project ID](https://console.scaleway.com/project/credentials) in the Scaleway console.
//
// #### How can I add my server to a Private Network?
//
// See [our dedicated documentation](/en/developers/api/elastic-metal-flexible-ip).
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
	return []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1}
}

// ListServers: List Elastic Metal servers for a specific Organization.
func (s *API) ListServers(req *ListServersRequest, opts ...scw.RequestOption) (*ListServersResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "option_id", req.OptionID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers",
		Query:  query,
	}

	var resp ListServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetServer: Get full details of an existing Elastic Metal server associated with the ID.
func (s *API) GetServer(req *GetServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateServer: Create a new Elastic Metal server. Once the server is created, proceed with the [installation of an OS](#post-3e949e).
func (s *API) CreateServer(req *CreateServerRequest, opts ...scw.RequestOption) (*Server, error) {
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
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateServer: Update the server associated with the ID. You can update parameters such as the server's name, tags and description. Any parameters left null in the request body are not updated.
func (s *API) UpdateServer(req *UpdateServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// InstallServer: Install an Operating System (OS) on the Elastic Metal server with a specific ID.
func (s *API) InstallServer(req *InstallServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/install",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetServerMetrics: Get the ping status of the server associated with the ID.
func (s *API) GetServerMetrics(req *GetServerMetricsRequest, opts ...scw.RequestOption) (*GetServerMetricsResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/metrics",
	}

	var resp GetServerMetricsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteServer: Delete the server associated with the ID.
func (s *API) DeleteServer(req *DeleteServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RebootServer: Reboot the Elastic Metal server associated with the ID, use the `boot_type` `rescue` to reboot the server in rescue mode.
func (s *API) RebootServer(req *RebootServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/reboot",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StartServer: Start the server associated with the ID.
func (s *API) StartServer(req *StartServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/start",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StopServer: Stop the server associated with the ID. The server remains allocated to your account and all data remains on the local storage of the server.
func (s *API) StopServer(req *StopServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/stop",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListServerEvents: List event (i.e. start/stop/reboot) associated to the server ID.
func (s *API) ListServerEvents(req *ListServerEventsRequest, opts ...scw.RequestOption) (*ListServerEventsResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/events",
		Query:  query,
	}

	var resp ListServerEventsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StartBMCAccess: Start BMC (Baseboard Management Controller) access associated with the ID.
// The BMC (Baseboard Management Controller) access is available one hour after the installation of the server.
// You need first to create an option Remote Access. You will find the ID and the price with a call to listOffers (https://developers.scaleway.com/en/products/baremetal/api/#get-78db92). Then add the option https://developers.scaleway.com/en/products/baremetal/api/#post-b14abd.
// After adding the BMC option, you need to Get Remote Access to get the login/password https://developers.scaleway.com/en/products/baremetal/api/#get-cefc0f. Do not forget to delete the Option after use.
func (s *API) StartBMCAccess(req *StartBMCAccessRequest, opts ...scw.RequestOption) (*BMCAccess, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/bmc-access",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp BMCAccess

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBMCAccess: Get the BMC (Baseboard Management Controller) access associated with the ID, including the URL and login information needed to connect.
func (s *API) GetBMCAccess(req *GetBMCAccessRequest, opts ...scw.RequestOption) (*BMCAccess, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/bmc-access",
	}

	var resp BMCAccess

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StopBMCAccess: Stop BMC (Baseboard Management Controller) access associated with the ID.
func (s *API) StopBMCAccess(req *StopBMCAccessRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/bmc-access",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdateIP: Configure the IP address associated with the server ID and IP ID. You can use this method to set a reverse DNS for an IP address.
func (s *API) UpdateIP(req *UpdateIPRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/ips/" + fmt.Sprint(req.IPID) + "",
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

// AddOptionServer: Add an option, such as Private Networks, to a specific server.
func (s *API) AddOptionServer(req *AddOptionServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	if fmt.Sprint(req.OptionID) == "" {
		return nil, errors.New("field OptionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/options/" + fmt.Sprint(req.OptionID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteOptionServer: Delete an option from a specific server.
func (s *API) DeleteOptionServer(req *DeleteOptionServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	if fmt.Sprint(req.OptionID) == "" {
		return nil, errors.New("field OptionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/options/" + fmt.Sprint(req.OptionID) + "",
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListOffers: List all available Elastic Metal server configurations.
func (s *API) ListOffers(req *ListOffersRequest, opts ...scw.RequestOption) (*ListOffersResponse, error) {
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
	parameter.AddToQuery(query, "subscription_period", req.SubscriptionPeriod)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/offers",
		Query:  query,
	}

	var resp ListOffersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOffer: Get details of an offer identified by its offer ID.
func (s *API) GetOffer(req *GetOfferRequest, opts ...scw.RequestOption) (*Offer, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.OfferID) == "" {
		return nil, errors.New("field OfferID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/offers/" + fmt.Sprint(req.OfferID) + "",
	}

	var resp Offer

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOption: Return specific option for the ID.
func (s *API) GetOption(req *GetOptionRequest, opts ...scw.RequestOption) (*Option, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.OptionID) == "" {
		return nil, errors.New("field OptionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/options/" + fmt.Sprint(req.OptionID) + "",
	}

	var resp Option

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListOptions: List all options matching with filters.
func (s *API) ListOptions(req *ListOptionsRequest, opts ...scw.RequestOption) (*ListOptionsResponse, error) {
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
	parameter.AddToQuery(query, "offer_id", req.OfferID)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/options",
		Query:  query,
	}

	var resp ListOptionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSettings: Return all settings for a Project ID.
func (s *API) ListSettings(req *ListSettingsRequest, opts ...scw.RequestOption) (*ListSettingsResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/settings",
		Query:  query,
	}

	var resp ListSettingsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSetting: Update a setting for a Project ID (enable or disable).
func (s *API) UpdateSetting(req *UpdateSettingRequest, opts ...scw.RequestOption) (*Setting, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SettingID) == "" {
		return nil, errors.New("field SettingID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/settings/" + fmt.Sprint(req.SettingID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Setting

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListOS: List all OSes that are available for installation on Elastic Metal servers.
func (s *API) ListOS(req *ListOSRequest, opts ...scw.RequestOption) (*ListOSResponse, error) {
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
	parameter.AddToQuery(query, "offer_id", req.OfferID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/os",
		Query:  query,
	}

	var resp ListOSResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOS: Return the specific OS for the ID.
func (s *API) GetOS(req *GetOSRequest, opts ...scw.RequestOption) (*OS, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.OsID) == "" {
		return nil, errors.New("field OsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/os/" + fmt.Sprint(req.OsID) + "",
	}

	var resp OS

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// VPC is a set of products and features allowing you to build your own
// virtual private cloud on top of Scaleway's shared public cloud.
// It includes Private Networks, enabling resources to be interconnected
// through a dedicated, private, and flexible L2 network.
//
// You can add as many servers to your networks as you want and add
// up to eight (8) different networks per server, taking the form of additional
// network interfaces inside your server (VLANs). This service allows you to run
// services isolated from the public internet and expose them to the rest of your
// infrastructure without worrying about public network filtering.
// Servers can be plugged and unplugged from a network at will, even when the server
// is running: the network interface will be hot-plugged to the server,
// and software can be configured to set it up as soon as it appears automatically.
//
// ## Technical limitations
//
// - need to configure manually a VLAN on the BMaaS interface
// - Bandwidth is limited to 1Gbps inside the private network
// - Up to 8 private networks per server
// - Broadcast and Multicast traffic, while supported, are rate-limited
//
// # Quickstart
//
// ## Requirements
//
// You need to have an HTTP client such as [`curl`](https://curl.haxx.se) to use Scaleway's API. It is
// also a good idea to have [`jq`](https://stedolan.github.io/jq) which will help you read and parse
// JSON output. Make sure you have these two tools before you begin. Otherwise use your package manager
// to install them.
//
// To call Scaleway's API, you need an `X-Auth-Token`. If you don't have one yet, refer to our
// [doc about generating API keys](https://www.scaleway.com/en/docs/console/my-project/how-to/create-a-project/#-Generating-API-Keys).
//
// Next, you will need your Project ID to create VPC resources in. If you don't have it, refer to our
// [doc about creating a Project](https://www.scaleway.com/en/docs/console/my-project/how-to/create-a-project/#-Creating-a-Project).
//
// Finally, you will need to chose the Availability Zone in which to create your Private Network.
// Keep in mind that Private Networks are per zone and not per region, thus you will only be able to
// connect servers to networks from the same Availability Zone.
//
// ```sh
// export SECRET_KEY="<Your secret key>"
// export PROJECT_ID="<Your Project ID>"
// export ZONE="<Chosen zone (fr-par-1/nl-ams-1)>"
// ```
//
// ## Creating a Private Network
//
// See [our online documentation](https://developers.scaleway.com/en/products/vpc/api/)
//
// ## Adding a server to the Private Network
//
// ```sh
//
//	curl -s -H "Content-Type: application/json" -H "X-Auth-Token: $SECRET_KEY" \
//	    https://api.scaleway.com/baremetal/v1/zones/$ZONE/servers/$SRV_ID/private-networks \
//	    -d '{"private_network_id": "'$PN_ID'"}'
//
// ```
//
// Keep the `vlan` field from the response. It is your VLAN ID, and will be used
// to configure the server to handle traffic from and to the private network:
//
// ```sh
// sudo ip link add link eno1 name eno1.$VLAN type vlan id $VLAN
// sudo ip link set eno1.$VLAN up
// sudo ip addr add 192.168.0.10/24 dev eno1.$VLAN
// ```.
type PrivateNetworkAPI struct {
	client *scw.Client
}

// NewPrivateNetworkAPI returns a PrivateNetworkAPI object from a Scaleway client.
func NewPrivateNetworkAPI(client *scw.Client) *PrivateNetworkAPI {
	return &PrivateNetworkAPI{
		client: client,
	}
}
func (s *PrivateNetworkAPI) Zones() []scw.Zone {
	return []scw.Zone{scw.ZoneFrPar2}
}

// AddServerPrivateNetwork: Add a server to a Private Network.
func (s *PrivateNetworkAPI) AddServerPrivateNetwork(req *PrivateNetworkAPIAddServerPrivateNetworkRequest, opts ...scw.RequestOption) (*ServerPrivateNetwork, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/private-networks",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ServerPrivateNetwork

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetServerPrivateNetworks: Set multiple Private Networks on a server.
func (s *PrivateNetworkAPI) SetServerPrivateNetworks(req *PrivateNetworkAPISetServerPrivateNetworksRequest, opts ...scw.RequestOption) (*SetServerPrivateNetworksResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/private-networks",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetServerPrivateNetworksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListServerPrivateNetworks: List the Private Networks of a server.
func (s *PrivateNetworkAPI) ListServerPrivateNetworks(req *PrivateNetworkAPIListServerPrivateNetworksRequest, opts ...scw.RequestOption) (*ListServerPrivateNetworksResponse, error) {
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
	parameter.AddToQuery(query, "server_id", req.ServerID)
	parameter.AddToQuery(query, "private_network_id", req.PrivateNetworkID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/server-private-networks",
		Query:  query,
	}

	var resp ListServerPrivateNetworksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteServerPrivateNetwork: Delete a Private Network.
func (s *PrivateNetworkAPI) DeleteServerPrivateNetwork(req *PrivateNetworkAPIDeleteServerPrivateNetworkRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return errors.New("field ServerID cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/baremetal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}
