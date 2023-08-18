// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package instance provides methods and message types of the instance v1 API.
package instance

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

type Arch string

const (
	ArchX86_64 = Arch("x86_64")
	ArchArm    = Arch("arm")
	ArchArm64  = Arch("arm64")
)

func (enum Arch) String() string {
	if enum == "" {
		// return default value if empty
		return "x86_64"
	}
	return string(enum)
}

func (enum Arch) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *Arch) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = Arch(Arch(tmp).String())
	return nil
}

type BootType string

const (
	BootTypeLocal      = BootType("local")
	BootTypeBootscript = BootType("bootscript")
	BootTypeRescue     = BootType("rescue")
)

func (enum BootType) String() string {
	if enum == "" {
		// return default value if empty
		return "local"
	}
	return string(enum)
}

func (enum BootType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *BootType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = BootType(BootType(tmp).String())
	return nil
}

type IPState string

const (
	IPStateUnknownState = IPState("unknown_state")
	IPStateDetached     = IPState("detached")
	IPStateAttached     = IPState("attached")
	IPStatePending      = IPState("pending")
	IPStateError        = IPState("error")
)

func (enum IPState) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_state"
	}
	return string(enum)
}

func (enum IPState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *IPState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = IPState(IPState(tmp).String())
	return nil
}

type IPType string

const (
	IPTypeUnknownIptype = IPType("unknown_iptype")
	IPTypeNat           = IPType("nat")
	IPTypeRoutedIPv4    = IPType("routed_ipv4")
	IPTypeRoutedIPv6    = IPType("routed_ipv6")
)

func (enum IPType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_iptype"
	}
	return string(enum)
}

func (enum IPType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *IPType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = IPType(IPType(tmp).String())
	return nil
}

type ImageState string

const (
	ImageStateAvailable = ImageState("available")
	ImageStateCreating  = ImageState("creating")
	ImageStateError     = ImageState("error")
)

func (enum ImageState) String() string {
	if enum == "" {
		// return default value if empty
		return "available"
	}
	return string(enum)
}

func (enum ImageState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ImageState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ImageState(ImageState(tmp).String())
	return nil
}

type ListServersRequestOrder string

const (
	ListServersRequestOrderCreationDateDesc     = ListServersRequestOrder("creation_date_desc")
	ListServersRequestOrderCreationDateAsc      = ListServersRequestOrder("creation_date_asc")
	ListServersRequestOrderModificationDateDesc = ListServersRequestOrder("modification_date_desc")
	ListServersRequestOrderModificationDateAsc  = ListServersRequestOrder("modification_date_asc")
)

func (enum ListServersRequestOrder) String() string {
	if enum == "" {
		// return default value if empty
		return "creation_date_desc"
	}
	return string(enum)
}

func (enum ListServersRequestOrder) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListServersRequestOrder) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListServersRequestOrder(ListServersRequestOrder(tmp).String())
	return nil
}

type PlacementGroupPolicyMode string

const (
	PlacementGroupPolicyModeOptional = PlacementGroupPolicyMode("optional")
	PlacementGroupPolicyModeEnforced = PlacementGroupPolicyMode("enforced")
)

func (enum PlacementGroupPolicyMode) String() string {
	if enum == "" {
		// return default value if empty
		return "optional"
	}
	return string(enum)
}

func (enum PlacementGroupPolicyMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PlacementGroupPolicyMode) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PlacementGroupPolicyMode(PlacementGroupPolicyMode(tmp).String())
	return nil
}

type PlacementGroupPolicyType string

const (
	PlacementGroupPolicyTypeMaxAvailability = PlacementGroupPolicyType("max_availability")
	PlacementGroupPolicyTypeLowLatency      = PlacementGroupPolicyType("low_latency")
)

func (enum PlacementGroupPolicyType) String() string {
	if enum == "" {
		// return default value if empty
		return "max_availability"
	}
	return string(enum)
}

func (enum PlacementGroupPolicyType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PlacementGroupPolicyType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PlacementGroupPolicyType(PlacementGroupPolicyType(tmp).String())
	return nil
}

type PrivateNICState string

const (
	PrivateNICStateAvailable    = PrivateNICState("available")
	PrivateNICStateSyncing      = PrivateNICState("syncing")
	PrivateNICStateSyncingError = PrivateNICState("syncing_error")
)

func (enum PrivateNICState) String() string {
	if enum == "" {
		// return default value if empty
		return "available"
	}
	return string(enum)
}

func (enum PrivateNICState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PrivateNICState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PrivateNICState(PrivateNICState(tmp).String())
	return nil
}

type SecurityGroupPolicy string

const (
	SecurityGroupPolicyAccept = SecurityGroupPolicy("accept")
	SecurityGroupPolicyDrop   = SecurityGroupPolicy("drop")
)

func (enum SecurityGroupPolicy) String() string {
	if enum == "" {
		// return default value if empty
		return "accept"
	}
	return string(enum)
}

func (enum SecurityGroupPolicy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecurityGroupPolicy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecurityGroupPolicy(SecurityGroupPolicy(tmp).String())
	return nil
}

type SecurityGroupRuleAction string

const (
	SecurityGroupRuleActionAccept = SecurityGroupRuleAction("accept")
	SecurityGroupRuleActionDrop   = SecurityGroupRuleAction("drop")
)

func (enum SecurityGroupRuleAction) String() string {
	if enum == "" {
		// return default value if empty
		return "accept"
	}
	return string(enum)
}

func (enum SecurityGroupRuleAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecurityGroupRuleAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecurityGroupRuleAction(SecurityGroupRuleAction(tmp).String())
	return nil
}

type SecurityGroupRuleDirection string

const (
	SecurityGroupRuleDirectionInbound  = SecurityGroupRuleDirection("inbound")
	SecurityGroupRuleDirectionOutbound = SecurityGroupRuleDirection("outbound")
)

func (enum SecurityGroupRuleDirection) String() string {
	if enum == "" {
		// return default value if empty
		return "inbound"
	}
	return string(enum)
}

func (enum SecurityGroupRuleDirection) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecurityGroupRuleDirection) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecurityGroupRuleDirection(SecurityGroupRuleDirection(tmp).String())
	return nil
}

type SecurityGroupRuleProtocol string

const (
	SecurityGroupRuleProtocolTCP  = SecurityGroupRuleProtocol("TCP")
	SecurityGroupRuleProtocolUDP  = SecurityGroupRuleProtocol("UDP")
	SecurityGroupRuleProtocolICMP = SecurityGroupRuleProtocol("ICMP")
	SecurityGroupRuleProtocolANY  = SecurityGroupRuleProtocol("ANY")
)

func (enum SecurityGroupRuleProtocol) String() string {
	if enum == "" {
		// return default value if empty
		return "TCP"
	}
	return string(enum)
}

func (enum SecurityGroupRuleProtocol) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecurityGroupRuleProtocol) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecurityGroupRuleProtocol(SecurityGroupRuleProtocol(tmp).String())
	return nil
}

type SecurityGroupState string

const (
	SecurityGroupStateAvailable    = SecurityGroupState("available")
	SecurityGroupStateSyncing      = SecurityGroupState("syncing")
	SecurityGroupStateSyncingError = SecurityGroupState("syncing_error")
)

func (enum SecurityGroupState) String() string {
	if enum == "" {
		// return default value if empty
		return "available"
	}
	return string(enum)
}

func (enum SecurityGroupState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecurityGroupState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecurityGroupState(SecurityGroupState(tmp).String())
	return nil
}

type ServerAction string

const (
	ServerActionPoweron     = ServerAction("poweron")
	ServerActionBackup      = ServerAction("backup")
	ServerActionStopInPlace = ServerAction("stop_in_place")
	ServerActionPoweroff    = ServerAction("poweroff")
	ServerActionTerminate   = ServerAction("terminate")
	ServerActionReboot      = ServerAction("reboot")
)

func (enum ServerAction) String() string {
	if enum == "" {
		// return default value if empty
		return "poweron"
	}
	return string(enum)
}

func (enum ServerAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerAction(ServerAction(tmp).String())
	return nil
}

type ServerIPIPFamily string

const (
	ServerIPIPFamilyInet  = ServerIPIPFamily("inet")
	ServerIPIPFamilyInet6 = ServerIPIPFamily("inet6")
)

func (enum ServerIPIPFamily) String() string {
	if enum == "" {
		// return default value if empty
		return "inet"
	}
	return string(enum)
}

func (enum ServerIPIPFamily) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerIPIPFamily) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerIPIPFamily(ServerIPIPFamily(tmp).String())
	return nil
}

type ServerIPProvisioningMode string

const (
	ServerIPProvisioningModeManual = ServerIPProvisioningMode("manual")
	ServerIPProvisioningModeDHCP   = ServerIPProvisioningMode("dhcp")
	ServerIPProvisioningModeSlaac  = ServerIPProvisioningMode("slaac")
)

func (enum ServerIPProvisioningMode) String() string {
	if enum == "" {
		// return default value if empty
		return "manual"
	}
	return string(enum)
}

func (enum ServerIPProvisioningMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerIPProvisioningMode) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerIPProvisioningMode(ServerIPProvisioningMode(tmp).String())
	return nil
}

type ServerState string

const (
	ServerStateRunning        = ServerState("running")
	ServerStateStopped        = ServerState("stopped")
	ServerStateStoppedInPlace = ServerState("stopped in place")
	ServerStateStarting       = ServerState("starting")
	ServerStateStopping       = ServerState("stopping")
	ServerStateLocked         = ServerState("locked")
)

func (enum ServerState) String() string {
	if enum == "" {
		// return default value if empty
		return "running"
	}
	return string(enum)
}

func (enum ServerState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerState(ServerState(tmp).String())
	return nil
}

type ServerTypesAvailability string

const (
	ServerTypesAvailabilityAvailable = ServerTypesAvailability("available")
	ServerTypesAvailabilityScarce    = ServerTypesAvailability("scarce")
	ServerTypesAvailabilityShortage  = ServerTypesAvailability("shortage")
)

func (enum ServerTypesAvailability) String() string {
	if enum == "" {
		// return default value if empty
		return "available"
	}
	return string(enum)
}

func (enum ServerTypesAvailability) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerTypesAvailability) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerTypesAvailability(ServerTypesAvailability(tmp).String())
	return nil
}

type SnapshotState string

const (
	SnapshotStateAvailable    = SnapshotState("available")
	SnapshotStateSnapshotting = SnapshotState("snapshotting")
	SnapshotStateError        = SnapshotState("error")
	SnapshotStateInvalidData  = SnapshotState("invalid_data")
	SnapshotStateImporting    = SnapshotState("importing")
	SnapshotStateExporting    = SnapshotState("exporting")
)

func (enum SnapshotState) String() string {
	if enum == "" {
		// return default value if empty
		return "available"
	}
	return string(enum)
}

func (enum SnapshotState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SnapshotState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SnapshotState(SnapshotState(tmp).String())
	return nil
}

type SnapshotVolumeType string

const (
	SnapshotVolumeTypeUnknownVolumeType = SnapshotVolumeType("unknown_volume_type")
	SnapshotVolumeTypeLSSD              = SnapshotVolumeType("l_ssd")
	SnapshotVolumeTypeBSSD              = SnapshotVolumeType("b_ssd")
	SnapshotVolumeTypeUnified           = SnapshotVolumeType("unified")
)

func (enum SnapshotVolumeType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_volume_type"
	}
	return string(enum)
}

func (enum SnapshotVolumeType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SnapshotVolumeType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SnapshotVolumeType(SnapshotVolumeType(tmp).String())
	return nil
}

type TaskStatus string

const (
	TaskStatusPending = TaskStatus("pending")
	TaskStatusStarted = TaskStatus("started")
	TaskStatusSuccess = TaskStatus("success")
	TaskStatusFailure = TaskStatus("failure")
	TaskStatusRetry   = TaskStatus("retry")
)

func (enum TaskStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "pending"
	}
	return string(enum)
}

func (enum TaskStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TaskStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TaskStatus(TaskStatus(tmp).String())
	return nil
}

type VolumeServerState string

const (
	VolumeServerStateAvailable    = VolumeServerState("available")
	VolumeServerStateSnapshotting = VolumeServerState("snapshotting")
	VolumeServerStateError        = VolumeServerState("error")
	VolumeServerStateFetching     = VolumeServerState("fetching")
	VolumeServerStateResizing     = VolumeServerState("resizing")
	VolumeServerStateSaving       = VolumeServerState("saving")
	VolumeServerStateHotsyncing   = VolumeServerState("hotsyncing")
)

func (enum VolumeServerState) String() string {
	if enum == "" {
		// return default value if empty
		return "available"
	}
	return string(enum)
}

func (enum VolumeServerState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *VolumeServerState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = VolumeServerState(VolumeServerState(tmp).String())
	return nil
}

type VolumeServerVolumeType string

const (
	VolumeServerVolumeTypeLSSD = VolumeServerVolumeType("l_ssd")
	VolumeServerVolumeTypeBSSD = VolumeServerVolumeType("b_ssd")
)

func (enum VolumeServerVolumeType) String() string {
	if enum == "" {
		// return default value if empty
		return "l_ssd"
	}
	return string(enum)
}

func (enum VolumeServerVolumeType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *VolumeServerVolumeType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = VolumeServerVolumeType(VolumeServerVolumeType(tmp).String())
	return nil
}

type VolumeState string

const (
	VolumeStateAvailable    = VolumeState("available")
	VolumeStateSnapshotting = VolumeState("snapshotting")
	VolumeStateError        = VolumeState("error")
	VolumeStateFetching     = VolumeState("fetching")
	VolumeStateResizing     = VolumeState("resizing")
	VolumeStateSaving       = VolumeState("saving")
	VolumeStateHotsyncing   = VolumeState("hotsyncing")
)

func (enum VolumeState) String() string {
	if enum == "" {
		// return default value if empty
		return "available"
	}
	return string(enum)
}

func (enum VolumeState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *VolumeState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = VolumeState(VolumeState(tmp).String())
	return nil
}

type VolumeVolumeType string

const (
	VolumeVolumeTypeLSSD    = VolumeVolumeType("l_ssd")
	VolumeVolumeTypeBSSD    = VolumeVolumeType("b_ssd")
	VolumeVolumeTypeUnified = VolumeVolumeType("unified")
)

func (enum VolumeVolumeType) String() string {
	if enum == "" {
		// return default value if empty
		return "l_ssd"
	}
	return string(enum)
}

func (enum VolumeVolumeType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *VolumeVolumeType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = VolumeVolumeType(VolumeVolumeType(tmp).String())
	return nil
}

// ServerSummary:
type ServerSummary struct {
	// ID:
	ID string `json:"id"`
	// Name:
	Name string `json:"name"`
}

// Bootscript:
type Bootscript struct {
	// Bootcmdargs: Bootscript arguments.
	Bootcmdargs string `json:"bootcmdargs"`
	// Default: Display if the bootscript is the default bootscript (if no other boot option is configured).
	Default bool `json:"default"`
	// Dtb: Provide information regarding a Device Tree Binary (DTB) for use with C1 servers.
	Dtb string `json:"dtb"`
	// ID: Bootscript ID.
	ID string `json:"id"`
	// Initrd: Initrd (initial ramdisk) configuration.
	Initrd string `json:"initrd"`
	// Kernel: Instance kernel version.
	Kernel string `json:"kernel"`
	// Organization: Bootscript Organization ID.
	Organization string `json:"organization"`
	// Project: Bootscript Project ID.
	Project string `json:"project"`
	// Public: Provide information if the bootscript is public.
	Public bool `json:"public"`
	// Title: Bootscript title.
	Title string `json:"title"`
	// Arch: Bootscript architecture.
	Arch Arch `json:"arch"`
	// Zone: Zone in which the bootscript is located.
	Zone scw.Zone `json:"zone"`
}

// Volume:
type Volume struct {
	// ID: Volume unique ID.
	ID string `json:"id"`
	// Name: Volume name.
	Name string `json:"name"`
	// Deprecated: ExportURI: Show the volume NBD export URI.
	ExportURI *string `json:"export_uri,omitempty"`
	// Size: Volume disk size.
	Size scw.Size `json:"size"`
	// VolumeType: Volume type.
	VolumeType VolumeVolumeType `json:"volume_type"`
	// CreationDate: Volume creation date.
	CreationDate *time.Time `json:"creation_date"`
	// ModificationDate: Volume modification date.
	ModificationDate *time.Time `json:"modification_date"`
	// Organization: Volume Organization ID.
	Organization string `json:"organization"`
	// Project: Volume Project ID.
	Project string `json:"project"`
	// Tags: Volume tags.
	Tags []string `json:"tags"`
	// Server: Instance attached to the volume.
	Server *ServerSummary `json:"server"`
	// State: Volume state.
	State VolumeState `json:"state"`
	// Zone: Zone in which the volume is located.
	Zone scw.Zone `json:"zone"`
}

// VolumeSummary:
type VolumeSummary struct {
	// ID:
	ID string `json:"id"`
	// Name:
	Name string `json:"name"`
	// Size:
	Size scw.Size `json:"size"`
	// VolumeType:
	VolumeType VolumeVolumeType `json:"volume_type"`
}

// ServerTypeNetworkInterface:
type ServerTypeNetworkInterface struct {
	// InternalBandwidth: Maximum internal bandwidth in bits per seconds.
	InternalBandwidth *uint64 `json:"internal_bandwidth"`
	// InternetBandwidth: Maximum internet bandwidth in bits per seconds.
	InternetBandwidth *uint64 `json:"internet_bandwidth"`
}

// ServerTypeVolumeConstraintSizes:
type ServerTypeVolumeConstraintSizes struct {
	// MinSize: Minimum volume size in bytes.
	MinSize scw.Size `json:"min_size"`
	// MaxSize: Maximum volume size in bytes.
	MaxSize scw.Size `json:"max_size"`
}

// Image:
type Image struct {
	// ID:
	ID string `json:"id"`
	// Name:
	Name string `json:"name"`
	// Arch:
	Arch Arch `json:"arch"`
	// CreationDate:
	CreationDate *time.Time `json:"creation_date"`
	// ModificationDate:
	ModificationDate *time.Time `json:"modification_date"`
	// Deprecated: DefaultBootscript:
	DefaultBootscript *Bootscript `json:"default_bootscript,omitempty"`
	// ExtraVolumes:
	ExtraVolumes map[string]*Volume `json:"extra_volumes"`
	// FromServer:
	FromServer string `json:"from_server"`
	// Organization:
	Organization string `json:"organization"`
	// Public:
	Public bool `json:"public"`
	// RootVolume:
	RootVolume *VolumeSummary `json:"root_volume"`
	// State:
	State ImageState `json:"state"`
	// Project:
	Project string `json:"project"`
	// Tags:
	Tags []string `json:"tags"`
	// Zone:
	Zone scw.Zone `json:"zone"`
}

// PlacementGroup:
type PlacementGroup struct {
	// ID: Placement group unique ID.
	ID string `json:"id"`
	// Name: Placement group name.
	Name string `json:"name"`
	// Organization: Placement group Organization ID.
	Organization string `json:"organization"`
	// Project: Placement group Project ID.
	Project string `json:"project"`
	// Tags: Placement group tags.
	Tags []string `json:"tags"`
	// PolicyMode: Select the failure mode when the placement cannot be respected, either optional or enforced.
	PolicyMode PlacementGroupPolicyMode `json:"policy_mode"`
	// PolicyType: Select the behavior of the placement group, either low_latency (group) or max_availability (spread).
	PolicyType PlacementGroupPolicyType `json:"policy_type"`
	// PolicyRespected: Returns true if the policy is respected, false otherwise.
	PolicyRespected bool `json:"policy_respected"`
	// Zone: Zone in which the placement group is located.
	Zone scw.Zone `json:"zone"`
}

// PrivateNIC:
type PrivateNIC struct {
	// ID: Private NIC unique ID.
	ID string `json:"id"`
	// ServerID: Instance to which the private NIC is attached.
	ServerID string `json:"server_id"`
	// PrivateNetworkID: Private Network the private NIC is attached to.
	PrivateNetworkID string `json:"private_network_id"`
	// MacAddress: Private NIC MAC address.
	MacAddress string `json:"mac_address"`
	// State: Private NIC state.
	State PrivateNICState `json:"state"`
	// Tags: Private NIC tags.
	Tags []string `json:"tags"`
}

// SecurityGroupSummary:
type SecurityGroupSummary struct {
	// ID:
	ID string `json:"id"`
	// Name:
	Name string `json:"name"`
}

// ServerIP:
type ServerIP struct {
	// ID: Unique ID of the IP address.
	ID string `json:"id"`
	// Address: Instance's public IP-Address.
	Address net.IP `json:"address"`
	// Gateway: Gateway's IP address.
	Gateway net.IP `json:"gateway"`
	// Netmask: CIDR netmask.
	Netmask string `json:"netmask"`
	// Family: IP address family (inet or inet6).
	Family ServerIPIPFamily `json:"family"`
	// Dynamic: True if the IP address is dynamic.
	Dynamic bool `json:"dynamic"`
	// ProvisioningMode: Information about this address provisioning mode.
	ProvisioningMode ServerIPProvisioningMode `json:"provisioning_mode"`
}

// ServerIPv6:
type ServerIPv6 struct {
	// Address: Instance IPv6 IP-Address.
	Address net.IP `json:"address"`
	// Gateway: IPv6 IP-addresses gateway.
	Gateway net.IP `json:"gateway"`
	// Netmask: IPv6 IP-addresses CIDR netmask.
	Netmask string `json:"netmask"`
}

// ServerLocation:
type ServerLocation struct {
	// ClusterID:
	ClusterID string `json:"cluster_id"`
	// HypervisorID:
	HypervisorID string `json:"hypervisor_id"`
	// NodeID:
	NodeID string `json:"node_id"`
	// PlatformID:
	PlatformID string `json:"platform_id"`
	// ZoneID:
	ZoneID string `json:"zone_id"`
}

// ServerMaintenance:
type ServerMaintenance struct {
	// Reason:
	Reason string `json:"reason"`
}

// VolumeServer:
type VolumeServer struct {
	// ID:
	ID string `json:"id"`
	// Name:
	Name string `json:"name"`
	// ExportURI:
	ExportURI string `json:"export_uri"`
	// Organization:
	Organization string `json:"organization"`
	// Server:
	Server *ServerSummary `json:"server"`
	// Size:
	Size scw.Size `json:"size"`
	// VolumeType:
	VolumeType VolumeServerVolumeType `json:"volume_type"`
	// CreationDate:
	CreationDate *time.Time `json:"creation_date"`
	// ModificationDate:
	ModificationDate *time.Time `json:"modification_date"`
	// State:
	State VolumeServerState `json:"state"`
	// Project:
	Project string `json:"project"`
	// Boot:
	Boot bool `json:"boot"`
	// Zone:
	Zone scw.Zone `json:"zone"`
}

// SnapshotBaseVolume:
type SnapshotBaseVolume struct {
	// ID: Volume ID on which the snapshot is based.
	ID string `json:"id"`
	// Name: Volume name on which the snapshot is based on.
	Name string `json:"name"`
}

// ServerTypeCapabilities:
type ServerTypeCapabilities struct {
	// BlockStorage: Defines whether the Instance supports block storage.
	BlockStorage *bool `json:"block_storage"`
	// BootTypes: List of supported boot types.
	BootTypes []BootType `json:"boot_types"`
}

// ServerTypeNetwork:
type ServerTypeNetwork struct {
	// Interfaces: List of available network interfaces.
	Interfaces []*ServerTypeNetworkInterface `json:"interfaces"`
	// SumInternalBandwidth: Total maximum internal bandwidth in bits per seconds.
	SumInternalBandwidth *uint64 `json:"sum_internal_bandwidth"`
	// SumInternetBandwidth: Total maximum internet bandwidth in bits per seconds.
	SumInternetBandwidth *uint64 `json:"sum_internet_bandwidth"`
	// IPv6Support: True if IPv6 is enabled.
	IPv6Support bool `json:"ipv6_support"`
}

// ServerTypeVolumeConstraintsByType:
type ServerTypeVolumeConstraintsByType struct {
	// LSSD: Local SSD volumes.
	LSSD *ServerTypeVolumeConstraintSizes `json:"l_ssd"`
}

// VolumeTypeCapabilities:
type VolumeTypeCapabilities struct {
	// Snapshot:
	Snapshot bool `json:"snapshot"`
}

// VolumeTypeConstraints:
type VolumeTypeConstraints struct {
	// Min:
	Min scw.Size `json:"min"`
	// Max:
	Max scw.Size `json:"max"`
}

// IP:
type IP struct {
	// ID:
	ID string `json:"id"`
	// Address:
	Address net.IP `json:"address"`
	// Reverse:
	Reverse *string `json:"reverse"`
	// Server:
	Server *ServerSummary `json:"server"`
	// Organization:
	Organization string `json:"organization"`
	// Tags:
	Tags []string `json:"tags"`
	// Project:
	Project string `json:"project"`
	// Type:
	Type IPType `json:"type"`
	// State:
	State IPState `json:"state"`
	// Prefix:
	Prefix scw.IPNet `json:"prefix"`
	// Zone:
	Zone scw.Zone `json:"zone"`
}

// VolumeTemplate:
type VolumeTemplate struct {
	// ID: UUID of the volume.
	ID string `json:"id"`
	// Name: Name of the volume.
	Name string `json:"name"`
	// Size: Disk size of the volume, must be a multiple of 512.
	Size scw.Size `json:"size"`
	// VolumeType: Type of the volume.
	VolumeType VolumeVolumeType `json:"volume_type"`
	// Deprecated: Organization: Organization ID of the volume.
	Organization *string `json:"organization,omitempty"`
	// Project: Project ID of the volume.
	Project *string `json:"project,omitempty"`
}

// SecurityGroup:
type SecurityGroup struct {
	// ID: Security group unique ID.
	ID string `json:"id"`
	// Name: Security group name.
	Name string `json:"name"`
	// Description: Security group description.
	Description string `json:"description"`
	// EnableDefaultSecurity: True if SMTP is blocked on IPv4 and IPv6. This feature is read only, please open a support ticket if you need to make it configurable.
	EnableDefaultSecurity bool `json:"enable_default_security"`
	// InboundDefaultPolicy: Default inbound policy.
	InboundDefaultPolicy SecurityGroupPolicy `json:"inbound_default_policy"`
	// OutboundDefaultPolicy: Default outbound policy.
	OutboundDefaultPolicy SecurityGroupPolicy `json:"outbound_default_policy"`
	// Organization: Security group Organization ID.
	Organization string `json:"organization"`
	// Project: Security group Project ID.
	Project string `json:"project"`
	// Tags: Security group tags.
	Tags []string `json:"tags"`
	// Deprecated: OrganizationDefault: True if it is your default security group for this Organization ID.
	OrganizationDefault *bool `json:"organization_default,omitempty"`
	// ProjectDefault: True if it is your default security group for this Project ID.
	ProjectDefault bool `json:"project_default"`
	// CreationDate: Security group creation date.
	CreationDate *time.Time `json:"creation_date"`
	// ModificationDate: Security group modification date.
	ModificationDate *time.Time `json:"modification_date"`
	// Servers: List of Instances attached to this security group.
	Servers []*ServerSummary `json:"servers"`
	// Stateful: Defines whether the security group is stateful.
	Stateful bool `json:"stateful"`
	// State: Security group state.
	State SecurityGroupState `json:"state"`
	// Zone: Zone in which the security group is located.
	Zone scw.Zone `json:"zone"`
}

// SecurityGroupRule:
type SecurityGroupRule struct {
	// ID:
	ID string `json:"id"`
	// Protocol:
	Protocol SecurityGroupRuleProtocol `json:"protocol"`
	// Direction:
	Direction SecurityGroupRuleDirection `json:"direction"`
	// Action:
	Action SecurityGroupRuleAction `json:"action"`
	// IPRange:
	IPRange scw.IPNet `json:"ip_range"`
	// DestPortFrom:
	DestPortFrom *uint32 `json:"dest_port_from"`
	// DestPortTo:
	DestPortTo *uint32 `json:"dest_port_to"`
	// Position:
	Position uint32 `json:"position"`
	// Editable:
	Editable bool `json:"editable"`
	// Zone:
	Zone scw.Zone `json:"zone"`
}

// VolumeServerTemplate:
type VolumeServerTemplate struct {
	// ID: UUID of the volume.
	ID *string `json:"id"`
	// Boot: Force the Instance to boot on this volume.
	Boot *bool `json:"boot"`
	// Name: Name of the volume.
	Name *string `json:"name"`
	// Size: Disk size of the volume, must be a multiple of 512.
	Size *scw.Size `json:"size"`
	// VolumeType: Type of the volume.
	VolumeType VolumeVolumeType `json:"volume_type"`
	// BaseSnapshot: ID of the snapshot on which this volume will be based.
	BaseSnapshot *string `json:"base_snapshot"`
	// Organization: Organization ID of the volume.
	Organization *string `json:"organization"`
	// Project: Project ID of the volume.
	Project *string `json:"project"`
}

// Server:
type Server struct {
	// ID: Instance unique ID.
	ID string `json:"id"`
	// Name: Instance name.
	Name string `json:"name"`
	// Organization: Instance Organization ID.
	Organization string `json:"organization"`
	// Project: Instance Project ID.
	Project string `json:"project"`
	// AllowedActions: List of allowed actions on the Instance.
	AllowedActions []ServerAction `json:"allowed_actions"`
	// Tags: Tags associated with the Instance.
	Tags []string `json:"tags"`
	// CommercialType: Instance commercial type (eg. GP1-M).
	CommercialType string `json:"commercial_type"`
	// CreationDate: Instance creation date.
	CreationDate *time.Time `json:"creation_date"`
	// DynamicIPRequired: True if a dynamic IPv4 is required.
	DynamicIPRequired bool `json:"dynamic_ip_required"`
	// RoutedIPEnabled: True to configure the instance so it uses the new routed IP mode.
	RoutedIPEnabled bool `json:"routed_ip_enabled"`
	// EnableIPv6: True if IPv6 is enabled.
	EnableIPv6 bool `json:"enable_ipv6"`
	// Hostname: Instance host name.
	Hostname string `json:"hostname"`
	// Image: Information about the Instance image.
	Image *Image `json:"image"`
	// Protected: Defines whether the Instance protection option is activated.
	Protected bool `json:"protected"`
	// PrivateIP: Private IP address of the Instance.
	PrivateIP *string `json:"private_ip"`
	// PublicIP: Information about the public IP.
	PublicIP *ServerIP `json:"public_ip"`
	// PublicIPs: Information about all the public IPs attached to the server.
	PublicIPs []*ServerIP `json:"public_ips"`
	// MacAddress: The server's MAC address.
	MacAddress string `json:"mac_address"`
	// ModificationDate: Instance modification date.
	ModificationDate *time.Time `json:"modification_date"`
	// State: Instance state.
	State ServerState `json:"state"`
	// Location: Instance location.
	Location *ServerLocation `json:"location"`
	// IPv6: Instance IPv6 address.
	IPv6 *ServerIPv6 `json:"ipv6"`
	// Deprecated: Bootscript: Instance bootscript.
	Bootscript *Bootscript `json:"bootscript,omitempty"`
	// BootType: Instance boot type.
	BootType BootType `json:"boot_type"`
	// Volumes: Instance volumes.
	Volumes map[string]*VolumeServer `json:"volumes"`
	// SecurityGroup: Instance security group.
	SecurityGroup *SecurityGroupSummary `json:"security_group"`
	// Maintenances: Instance planned maintenance.
	Maintenances []*ServerMaintenance `json:"maintenances"`
	// StateDetail: Detailed information about the Instance state.
	StateDetail string `json:"state_detail"`
	// Arch: Instance architecture.
	Arch Arch `json:"arch"`
	// PlacementGroup: Instance placement group.
	PlacementGroup *PlacementGroup `json:"placement_group"`
	// PrivateNics: Instance private NICs.
	PrivateNics []*PrivateNIC `json:"private_nics"`
	// Zone: Zone in which the Instance is located.
	Zone scw.Zone `json:"zone"`
}

// Snapshot:
type Snapshot struct {
	// ID: Snapshot ID.
	ID string `json:"id"`
	// Name: Snapshot name.
	Name string `json:"name"`
	// Organization: Snapshot Organization ID.
	Organization string `json:"organization"`
	// Project: Snapshot Project ID.
	Project string `json:"project"`
	// Tags: Snapshot tags.
	Tags []string `json:"tags"`
	// VolumeType: Snapshot volume type.
	VolumeType VolumeVolumeType `json:"volume_type"`
	// Size: Snapshot size.
	Size scw.Size `json:"size"`
	// State: Snapshot state.
	State SnapshotState `json:"state"`
	// BaseVolume: Volume on which the snapshot is based on.
	BaseVolume *SnapshotBaseVolume `json:"base_volume"`
	// CreationDate: Snapshot creation date.
	CreationDate *time.Time `json:"creation_date"`
	// ModificationDate: Snapshot modification date.
	ModificationDate *time.Time `json:"modification_date"`
	// Zone: Snapshot zone.
	Zone scw.Zone `json:"zone"`
	// ErrorReason: Reason for the failed snapshot import.
	ErrorReason *string `json:"error_reason"`
}

// Task:
type Task struct {
	// ID: Unique ID of the task.
	ID string `json:"id"`
	// Description: Description of the task.
	Description string `json:"description"`
	// Progress: Progress of the task in percent.
	Progress int32 `json:"progress"`
	// StartedAt: Task start date.
	StartedAt *time.Time `json:"started_at"`
	// TerminatedAt: Task end date.
	TerminatedAt *time.Time `json:"terminated_at"`
	// Status: Task status.
	Status TaskStatus `json:"status"`
	// HrefFrom:
	HrefFrom string `json:"href_from"`
	// HrefResult:
	HrefResult string `json:"href_result"`
	// Zone: Zone in which the task is excecuted.
	Zone scw.Zone `json:"zone"`
}

// Dashboard:
type Dashboard struct {
	// VolumesCount:
	VolumesCount uint32 `json:"volumes_count"`
	// RunningServersCount:
	RunningServersCount uint32 `json:"running_servers_count"`
	// ServersByTypes:
	ServersByTypes map[string]uint32 `json:"servers_by_types"`
	// ImagesCount:
	ImagesCount uint32 `json:"images_count"`
	// SnapshotsCount:
	SnapshotsCount uint32 `json:"snapshots_count"`
	// ServersCount:
	ServersCount uint32 `json:"servers_count"`
	// IPsCount:
	IPsCount uint32 `json:"ips_count"`
	// SecurityGroupsCount:
	SecurityGroupsCount uint32 `json:"security_groups_count"`
	// IPsUnused:
	IPsUnused uint32 `json:"ips_unused"`
	// VolumesLSSDCount:
	VolumesLSSDCount uint32 `json:"volumes_l_ssd_count"`
	// VolumesBSSDCount:
	VolumesBSSDCount uint32 `json:"volumes_b_ssd_count"`
	// VolumesLSSDTotalSize:
	VolumesLSSDTotalSize scw.Size `json:"volumes_l_ssd_total_size"`
	// VolumesBSSDTotalSize:
	VolumesBSSDTotalSize scw.Size `json:"volumes_b_ssd_total_size"`
	// PrivateNicsCount:
	PrivateNicsCount uint32 `json:"private_nics_count"`
	// PlacementGroupsCount:
	PlacementGroupsCount uint32 `json:"placement_groups_count"`
}

// PlacementGroupServer:
type PlacementGroupServer struct {
	// ID: Instance UUID.
	ID string `json:"id"`
	// Name: Instance name.
	Name string `json:"name"`
	// PolicyRespected: Defines whether the placement group policy is respected (either 1 or 0).
	PolicyRespected bool `json:"policy_respected"`
}

// GetServerTypesAvailabilityResponseAvailability:
type GetServerTypesAvailabilityResponseAvailability struct {
	// Availability:
	Availability ServerTypesAvailability `json:"availability"`
}

// ServerType:
type ServerType struct {
	// Deprecated: MonthlyPrice: Estimated monthly price, for a 30 days month, in Euro.
	MonthlyPrice *float32 `json:"monthly_price,omitempty"`
	// HourlyPrice: Hourly price in Euro.
	HourlyPrice float32 `json:"hourly_price"`
	// AltNames: Alternative Instance name, if any.
	AltNames []string `json:"alt_names"`
	// PerVolumeConstraint: Additional volume constraints.
	PerVolumeConstraint *ServerTypeVolumeConstraintsByType `json:"per_volume_constraint"`
	// VolumesConstraint: Initial volume constraints.
	VolumesConstraint *ServerTypeVolumeConstraintSizes `json:"volumes_constraint"`
	// Ncpus: Number of CPU.
	Ncpus uint32 `json:"ncpus"`
	// Gpu: Number of GPU.
	Gpu *uint64 `json:"gpu"`
	// RAM: Available RAM in bytes.
	RAM uint64 `json:"ram"`
	// Arch: CPU architecture.
	Arch Arch `json:"arch"`
	// Baremetal: True if it is a baremetal Instance.
	Baremetal bool `json:"baremetal"`
	// Network: Network available for the Instance.
	Network *ServerTypeNetwork `json:"network"`
	// Capabilities: Capabilities.
	Capabilities *ServerTypeCapabilities `json:"capabilities"`
}

// VolumeType:
type VolumeType struct {
	// DisplayName:
	DisplayName string `json:"display_name"`
	// Capabilities:
	Capabilities *VolumeTypeCapabilities `json:"capabilities"`
	// Constraints:
	Constraints *VolumeTypeConstraints `json:"constraints"`
}

// ServerActionRequestVolumeBackupTemplate:
type ServerActionRequestVolumeBackupTemplate struct {
	// VolumeType: Overrides the `volume_type` of the snapshot for this volume.
	// If omitted, the volume type of the original volume will be used.
	VolumeType SnapshotVolumeType `json:"volume_type"`
}

// SetSecurityGroupRulesRequestRule:
type SetSecurityGroupRulesRequestRule struct {
	// ID: UUID of the security rule to update. If no value is provided, a new rule will be created.
	ID *string `json:"id"`
	// Action: Action to apply when the rule matches a packet.
	Action SecurityGroupRuleAction `json:"action"`
	// Protocol: Protocol family this rule applies to.
	Protocol SecurityGroupRuleProtocol `json:"protocol"`
	// Direction: Direction the rule applies to.
	Direction SecurityGroupRuleDirection `json:"direction"`
	// IPRange: Range of IP addresses these rules apply to.
	IPRange scw.IPNet `json:"ip_range"`
	// DestPortFrom: Beginning of the range of ports this rule applies to (inclusive). This value will be set to null if protocol is ICMP or ANY.
	DestPortFrom *uint32 `json:"dest_port_from"`
	// DestPortTo: End of the range of ports this rule applies to (inclusive). This value will be set to null if protocol is ICMP or ANY, or if it is equal to dest_port_from.
	DestPortTo *uint32 `json:"dest_port_to"`
	// Position: Position of this rule in the security group rules list. If several rules are passed with the same position, the resulting order is undefined.
	Position uint32 `json:"position"`
	// Editable: Indicates if this rule is editable. Rules with the value false will be ignored.
	Editable *bool `json:"editable"`
	// Zone: Zone of the rule. This field is ignored.
	Zone *scw.Zone `json:"zone"`
}

// NullableStringValue:
type NullableStringValue struct {
	// Null:
	Null bool `json:"null"`
	// Value:
	Value string `json:"value"`
}

// SecurityGroupTemplate:
type SecurityGroupTemplate struct {
	// ID:
	ID string `json:"id"`
	// Name:
	Name string `json:"name"`
}

// CreateIPRequest:
type CreateIPRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Deprecated: Organization: Organization ID in which the IP is reserved.
	Organization *string `json:"organization,omitempty"`
	// Project: Project ID in which the IP is reserved.
	Project *string `json:"project,omitempty"`
	// Tags: Tags of the IP.
	Tags []string `json:"tags,omitempty"`
	// Server: UUID of the Instance you want to attach the IP to.
	Server *string `json:"server,omitempty"`
	// Type: IP type to reserve (either 'nat', 'routed_ipv4' or 'routed_ipv6').
	Type IPType `json:"type,omitempty"`
}

// CreateIPResponse:
type CreateIPResponse struct {
	// IP:
	IP *IP `json:"ip"`
}

// CreateImageRequest:
type CreateImageRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Name: Name of the image.
	Name string `json:"name,omitempty"`
	// RootVolume: UUID of the snapshot.
	RootVolume string `json:"root_volume,omitempty"`
	// Arch: Architecture of the image.
	Arch Arch `json:"arch,omitempty"`
	// Deprecated: DefaultBootscript: Default bootscript of the image.
	DefaultBootscript *string `json:"default_bootscript,omitempty"`
	// ExtraVolumes: Additional volumes of the image.
	ExtraVolumes map[string]*VolumeTemplate `json:"extra_volumes,omitempty"`
	// Deprecated: Organization: Organization ID of the image.
	Organization *string `json:"organization,omitempty"`
	// Project: Project ID of the image.
	Project *string `json:"project,omitempty"`
	// Tags: Tags of the image.
	Tags []string `json:"tags,omitempty"`
	// Public: True to create a public image.
	Public *bool `json:"public,omitempty"`
}

// CreateImageResponse:
type CreateImageResponse struct {
	// Image:
	Image *Image `json:"image"`
}

// CreatePlacementGroupRequest:
type CreatePlacementGroupRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Name: Name of the placement group.
	Name string `json:"name,omitempty"`
	// Deprecated: Organization: Organization ID of the placement group.
	Organization *string `json:"organization,omitempty"`
	// Project: Project ID of the placement group.
	Project *string `json:"project,omitempty"`
	// Tags: Tags of the placement group.
	Tags []string `json:"tags,omitempty"`
	// PolicyMode: Operating mode of the placement group.
	PolicyMode PlacementGroupPolicyMode `json:"policy_mode,omitempty"`
	// PolicyType: Policy type of the placement group.
	PolicyType PlacementGroupPolicyType `json:"policy_type,omitempty"`
}

// CreatePlacementGroupResponse:
type CreatePlacementGroupResponse struct {
	// PlacementGroup:
	PlacementGroup *PlacementGroup `json:"placement_group"`
}

// CreatePrivateNICRequest:
type CreatePrivateNICRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the Instance the private NIC will be attached to.
	ServerID string `json:"-"`
	// PrivateNetworkID: UUID of the private network where the private NIC will be attached.
	PrivateNetworkID string `json:"private_network_id,omitempty"`
	// Tags: Private NIC tags.
	Tags []string `json:"tags,omitempty"`
	// IPIDs: Ip_ids defined from IPAM.
	IPIDs []string `json:"ip_ids,omitempty"`
}

// CreatePrivateNICResponse:
type CreatePrivateNICResponse struct {
	// PrivateNic:
	PrivateNic *PrivateNIC `json:"private_nic"`
}

// CreateSecurityGroupRequest:
type CreateSecurityGroupRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Name: Name of the security group.
	Name string `json:"name,omitempty"`
	// Description: Description of the security group.
	Description string `json:"description,omitempty"`
	// Deprecated: Organization: Organization ID the security group belongs to.
	Organization *string `json:"organization,omitempty"`
	// Project: Project ID the security group belong to.
	Project *string `json:"project,omitempty"`
	// Tags: Tags of the security group.
	Tags []string `json:"tags,omitempty"`
	// Deprecated: OrganizationDefault: Defines whether this security group becomes the default security group for new Instances.
	OrganizationDefault *bool `json:"organization_default,omitempty"`
	// ProjectDefault: Whether this security group becomes the default security group for new Instances.
	ProjectDefault *bool `json:"project_default,omitempty"`
	// Stateful: Whether the security group is stateful or not.
	Stateful bool `json:"stateful,omitempty"`
	// InboundDefaultPolicy: Default policy for inbound rules.
	InboundDefaultPolicy SecurityGroupPolicy `json:"inbound_default_policy,omitempty"`
	// OutboundDefaultPolicy: Default policy for outbound rules.
	OutboundDefaultPolicy SecurityGroupPolicy `json:"outbound_default_policy,omitempty"`
	// EnableDefaultSecurity: True to block SMTP on IPv4 and IPv6. This feature is read only, please open a support ticket if you need to make it configurable.
	EnableDefaultSecurity *bool `json:"enable_default_security,omitempty"`
}

// CreateSecurityGroupResponse:
type CreateSecurityGroupResponse struct {
	// SecurityGroup:
	SecurityGroup *SecurityGroup `json:"security_group"`
}

// CreateSecurityGroupRuleRequest:
type CreateSecurityGroupRuleRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// SecurityGroupID: UUID of the security group.
	SecurityGroupID string `json:"-"`
	// Protocol:
	Protocol SecurityGroupRuleProtocol `json:"protocol,omitempty"`
	// Direction:
	Direction SecurityGroupRuleDirection `json:"direction,omitempty"`
	// Action:
	Action SecurityGroupRuleAction `json:"action,omitempty"`
	// IPRange:
	IPRange scw.IPNet `json:"ip_range,omitempty"`
	// DestPortFrom: Beginning of the range of ports to apply this rule to (inclusive).
	DestPortFrom *uint32 `json:"dest_port_from,omitempty"`
	// DestPortTo: End of the range of ports to apply this rule to (inclusive).
	DestPortTo *uint32 `json:"dest_port_to,omitempty"`
	// Position: Position of this rule in the security group rules list.
	Position uint32 `json:"position,omitempty"`
	// Editable: Indicates if this rule is editable (will be ignored).
	Editable bool `json:"editable,omitempty"`
}

// CreateSecurityGroupRuleResponse:
type CreateSecurityGroupRuleResponse struct {
	// Rule:
	Rule *SecurityGroupRule `json:"rule"`
}

// CreateServerRequest:
type CreateServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Name: Instance name.
	Name string `json:"name,omitempty"`
	// DynamicIPRequired: Define if a dynamic IPv4 is required for the Instance.
	DynamicIPRequired *bool `json:"dynamic_ip_required,omitempty"`
	// RoutedIPEnabled: If true, configure the Instance so it uses the new routed IP mode.
	RoutedIPEnabled *bool `json:"routed_ip_enabled,omitempty"`
	// CommercialType: Define the Instance commercial type (i.e. GP1-S).
	CommercialType string `json:"commercial_type,omitempty"`
	// Image: Instance image ID or label.
	Image string `json:"image,omitempty"`
	// Volumes: Volumes attached to the server.
	Volumes map[string]*VolumeServerTemplate `json:"volumes,omitempty"`
	// EnableIPv6: True if IPv6 is enabled on the server.
	EnableIPv6 bool `json:"enable_ipv6,omitempty"`
	// PublicIP: ID of the reserved IP to attach to the Instance.
	PublicIP *string `json:"public_ip,omitempty"`
	// PublicIPs: A list of reserved IP IDs to attach to the Instance.
	PublicIPs *[]string `json:"public_ips,omitempty"`
	// BootType: Boot type to use.
	BootType *BootType `json:"boot_type,omitempty"`
	// Deprecated: Bootscript: Bootscript ID to use when `boot_type` is set to `bootscript`.
	Bootscript *string `json:"bootscript,omitempty"`
	// Deprecated: Organization: Instance Organization ID.
	Organization *string `json:"organization,omitempty"`
	// Project: Instance Project ID.
	Project *string `json:"project,omitempty"`
	// Tags: Instance tags.
	Tags []string `json:"tags,omitempty"`
	// SecurityGroup: Security group ID.
	SecurityGroup *string `json:"security_group,omitempty"`
	// PlacementGroup: Placement group ID if Instance must be part of a placement group.
	PlacementGroup *string `json:"placement_group,omitempty"`
}

// CreateServerResponse:
type CreateServerResponse struct {
	// Server:
	Server *Server `json:"server"`
}

// CreateSnapshotRequest:
type CreateSnapshotRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Name: Name of the snapshot.
	Name string `json:"name,omitempty"`
	// VolumeID: UUID of the volume.
	VolumeID *string `json:"volume_id,omitempty"`
	// Tags: Tags of the snapshot.
	Tags *[]string `json:"tags,omitempty"`
	// Deprecated: Organization: Organization ID of the snapshot.
	Organization *string `json:"organization,omitempty"`
	// Project: Project ID of the snapshot.
	Project *string `json:"project,omitempty"`
	// VolumeType: Overrides the volume_type of the snapshot.
	// If omitted, the volume type of the original volume will be used.
	VolumeType SnapshotVolumeType `json:"volume_type,omitempty"`
	// Bucket: Bucket name for snapshot imports.
	Bucket *string `json:"bucket,omitempty"`
	// Key: Object key for snapshot imports.
	Key *string `json:"key,omitempty"`
	// Size: Imported snapshot size, must be a multiple of 512.
	Size *scw.Size `json:"size,omitempty"`
}

// CreateSnapshotResponse:
type CreateSnapshotResponse struct {
	// Snapshot:
	Snapshot *Snapshot `json:"snapshot"`
	// Task:
	Task *Task `json:"task"`
}

// CreateVolumeRequest:
type CreateVolumeRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Name: Volume name.
	Name string `json:"name,omitempty"`
	// Deprecated: Organization: Volume Organization ID.
	Organization *string `json:"organization,omitempty"`
	// Project: Volume Project ID.
	Project *string `json:"project,omitempty"`
	// Tags: Volume tags.
	Tags []string `json:"tags,omitempty"`
	// VolumeType: Volume type.
	VolumeType VolumeVolumeType `json:"volume_type,omitempty"`
	// Size: Volume disk size, must be a multiple of 512.
	Size *scw.Size `json:"size,omitempty"`
	// BaseVolume: ID of the volume on which this volume will be based.
	BaseVolume *string `json:"base_volume,omitempty"`
	// BaseSnapshot: ID of the snapshot on which this volume will be based.
	BaseSnapshot *string `json:"base_snapshot,omitempty"`
}

// CreateVolumeResponse:
type CreateVolumeResponse struct {
	// Volume:
	Volume *Volume `json:"volume"`
}

// DeleteIPRequest:
type DeleteIPRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// IP: ID or address of the IP to delete.
	IP string `json:"-"`
}

// DeleteImageRequest:
type DeleteImageRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ImageID: UUID of the image you want to delete.
	ImageID string `json:"-"`
}

// DeletePlacementGroupRequest:
type DeletePlacementGroupRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// PlacementGroupID: UUID of the placement group you want to delete.
	PlacementGroupID string `json:"-"`
}

// DeletePrivateNICRequest:
type DeletePrivateNICRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: Instance to which the private NIC is attached.
	ServerID string `json:"-"`
	// PrivateNicID: Private NIC unique ID.
	PrivateNicID string `json:"-"`
}

// DeleteSecurityGroupRequest:
type DeleteSecurityGroupRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// SecurityGroupID: UUID of the security group you want to delete.
	SecurityGroupID string `json:"-"`
}

// DeleteSecurityGroupRuleRequest:
type DeleteSecurityGroupRuleRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// SecurityGroupID:
	SecurityGroupID string `json:"-"`
	// SecurityGroupRuleID:
	SecurityGroupRuleID string `json:"-"`
}

// DeleteServerRequest:
type DeleteServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID:
	ServerID string `json:"-"`
}

// DeleteServerUserDataRequest:
type DeleteServerUserDataRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the Instance.
	ServerID string `json:"-"`
	// Key: Key of the user data to delete.
	Key string `json:"-"`
}

// DeleteSnapshotRequest:
type DeleteSnapshotRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// SnapshotID: UUID of the snapshot you want to delete.
	SnapshotID string `json:"-"`
}

// DeleteVolumeRequest:
type DeleteVolumeRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// VolumeID: UUID of the volume you want to delete.
	VolumeID string `json:"-"`
}

// ExportSnapshotRequest:
type ExportSnapshotRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// SnapshotID: Snapshot ID.
	SnapshotID string `json:"-"`
	// Bucket: S3 bucket name.
	Bucket string `json:"bucket,omitempty"`
	// Key: S3 object key.
	Key string `json:"key,omitempty"`
}

// ExportSnapshotResponse:
type ExportSnapshotResponse struct {
	// Task:
	Task *Task `json:"task"`
}

// GetBootscriptRequest:
type GetBootscriptRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// BootscriptID:
	BootscriptID string `json:"-"`
}

// GetBootscriptResponse:
type GetBootscriptResponse struct {
	// Bootscript:
	Bootscript *Bootscript `json:"bootscript"`
}

// GetDashboardRequest:
type GetDashboardRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Organization:
	Organization *string `json:"-"`
	// Project:
	Project *string `json:"-"`
}

// GetDashboardResponse:
type GetDashboardResponse struct {
	// Dashboard:
	Dashboard *Dashboard `json:"dashboard"`
}

// GetIPRequest:
type GetIPRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// IP: IP ID or address to get.
	IP string `json:"-"`
}

// GetIPResponse:
type GetIPResponse struct {
	// IP:
	IP *IP `json:"ip"`
}

// GetImageRequest:
type GetImageRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ImageID: UUID of the image you want to get.
	ImageID string `json:"-"`
}

// GetImageResponse:
type GetImageResponse struct {
	// Image:
	Image *Image `json:"image"`
}

// GetPlacementGroupRequest:
type GetPlacementGroupRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// PlacementGroupID: UUID of the placement group you want to get.
	PlacementGroupID string `json:"-"`
}

// GetPlacementGroupResponse:
type GetPlacementGroupResponse struct {
	// PlacementGroup:
	PlacementGroup *PlacementGroup `json:"placement_group"`
}

// GetPlacementGroupServersRequest:
type GetPlacementGroupServersRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// PlacementGroupID: UUID of the placement group you want to get.
	PlacementGroupID string `json:"-"`
}

// GetPlacementGroupServersResponse:
type GetPlacementGroupServersResponse struct {
	// Servers: Instances attached to the placement group.
	Servers []*PlacementGroupServer `json:"servers"`
}

// GetPrivateNICRequest:
type GetPrivateNICRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: Instance to which the private NIC is attached.
	ServerID string `json:"-"`
	// PrivateNicID: Private NIC unique ID.
	PrivateNicID string `json:"-"`
}

// GetPrivateNICResponse:
type GetPrivateNICResponse struct {
	// PrivateNic:
	PrivateNic *PrivateNIC `json:"private_nic"`
}

// GetSecurityGroupRequest:
type GetSecurityGroupRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// SecurityGroupID: UUID of the security group you want to get.
	SecurityGroupID string `json:"-"`
}

// GetSecurityGroupResponse:
type GetSecurityGroupResponse struct {
	// SecurityGroup:
	SecurityGroup *SecurityGroup `json:"security_group"`
}

// GetSecurityGroupRuleRequest:
type GetSecurityGroupRuleRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// SecurityGroupID:
	SecurityGroupID string `json:"-"`
	// SecurityGroupRuleID:
	SecurityGroupRuleID string `json:"-"`
}

// GetSecurityGroupRuleResponse:
type GetSecurityGroupRuleResponse struct {
	// Rule:
	Rule *SecurityGroupRule `json:"rule"`
}

// GetServerRequest:
type GetServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the Instance you want to get.
	ServerID string `json:"-"`
}

// GetServerResponse:
type GetServerResponse struct {
	// Server:
	Server *Server `json:"server"`
}

// GetServerTypesAvailabilityRequest:
type GetServerTypesAvailabilityRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// PerPage: A positive integer lower or equal to 100 to select the number of items to return.
	PerPage *uint32 `json:"-"`
	// Page: A positive integer to choose the page to return.
	Page *int32 `json:"-"`
}

// GetServerTypesAvailabilityResponse:
type GetServerTypesAvailabilityResponse struct {
	// Servers: Map of server types.
	Servers map[string]*GetServerTypesAvailabilityResponseAvailability `json:"servers"`
	// TotalCount:
	TotalCount uint32 `json:"total_count"`
}

// GetSnapshotRequest:
type GetSnapshotRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// SnapshotID: UUID of the snapshot you want to get.
	SnapshotID string `json:"-"`
}

// GetSnapshotResponse:
type GetSnapshotResponse struct {
	// Snapshot:
	Snapshot *Snapshot `json:"snapshot"`
}

// GetVolumeRequest:
type GetVolumeRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// VolumeID: UUID of the volume you want to get.
	VolumeID string `json:"-"`
}

// GetVolumeResponse:
type GetVolumeResponse struct {
	// Volume:
	Volume *Volume `json:"volume"`
}

// ListBootscriptsRequest:
type ListBootscriptsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Arch:
	Arch *string `json:"-"`
	// Title:
	Title *string `json:"-"`
	// Default:
	Default *bool `json:"-"`
	// Public:
	Public *bool `json:"-"`
	// PerPage:
	PerPage *uint32 `json:"-"`
	// Page:
	Page *int32 `json:"-"`
}

// ListBootscriptsResponse:
type ListBootscriptsResponse struct {
	// TotalCount: Total number of bootscripts.
	TotalCount uint32 `json:"total_count"`
	// Bootscripts: List of bootscripts.
	Bootscripts []*Bootscript `json:"bootscripts"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListBootscriptsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListBootscriptsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListBootscriptsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Bootscripts = append(r.Bootscripts, results.Bootscripts...)
	r.TotalCount += uint32(len(results.Bootscripts))
	return uint32(len(results.Bootscripts)), nil
}

// ListDefaultSecurityGroupRulesRequest:
type ListDefaultSecurityGroupRulesRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
}

// ListIPsRequest:
type ListIPsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Project: Project ID in which the IPs are reserved.
	Project *string `json:"-"`
	// Organization: Organization ID in which the IPs are reserved.
	Organization *string `json:"-"`
	// Tags: Filter IPs with these exact tags (to filter with several tags, use commas to separate them).
	Tags []string `json:"-"`
	// Name: Filter on the IP address (Works as a LIKE operation on the IP address).
	Name *string `json:"-"`
	// PerPage: A positive integer lower or equal to 100 to select the number of items to return.
	PerPage *uint32 `json:"-"`
	// Page: A positive integer to choose the page to return.
	Page *int32 `json:"-"`
}

// ListIPsResponse:
type ListIPsResponse struct {
	// TotalCount: Total number of ips.
	TotalCount uint32 `json:"total_count"`
	// IPs: List of ips.
	IPs []*IP `json:"ips"`
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

// ListImagesRequest:
type ListImagesRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Organization:
	Organization *string `json:"-"`
	// PerPage:
	PerPage *uint32 `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// Name:
	Name *string `json:"-"`
	// Public:
	Public *bool `json:"-"`
	// Arch:
	Arch *string `json:"-"`
	// Project:
	Project *string `json:"-"`
	// Tags:
	Tags *string `json:"-"`
}

// ListImagesResponse:
type ListImagesResponse struct {
	// TotalCount: Total number of images.
	TotalCount uint32 `json:"total_count"`
	// Images: List of images.
	Images []*Image `json:"images"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListImagesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListImagesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListImagesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Images = append(r.Images, results.Images...)
	r.TotalCount += uint32(len(results.Images))
	return uint32(len(results.Images)), nil
}

// ListPlacementGroupsRequest:
type ListPlacementGroupsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// PerPage: A positive integer lower or equal to 100 to select the number of items to return.
	PerPage *int32 `json:"-"`
	// Page: A positive integer to choose the page to return.
	Page *int32 `json:"-"`
	// Organization: List only placement groups of this Organization ID.
	Organization *string `json:"-"`
	// Project: List only placement groups of this Project ID.
	Project *string `json:"-"`
	// Tags: List placement groups with these exact tags (to filter with several tags, use commas to separate them).
	Tags []string `json:"-"`
	// Name: Filter placement groups by name (for eg. "cluster1" will return "cluster100" and "cluster1" but not "foo").
	Name *string `json:"-"`
}

// ListPlacementGroupsResponse:
type ListPlacementGroupsResponse struct {
	// TotalCount: Total number of placement groups.
	TotalCount uint32 `json:"total_count"`
	// PlacementGroups: List of placement groups.
	PlacementGroups []*PlacementGroup `json:"placement_groups"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPlacementGroupsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPlacementGroupsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListPlacementGroupsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PlacementGroups = append(r.PlacementGroups, results.PlacementGroups...)
	r.TotalCount += uint32(len(results.PlacementGroups))
	return uint32(len(results.PlacementGroups)), nil
}

// ListPrivateNICsRequest:
type ListPrivateNICsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: Instance to which the private NIC is attached.
	ServerID string `json:"-"`
	// Tags: Private NIC tags.
	Tags []string `json:"-"`
	// PerPage: A positive integer lower or equal to 100 to select the number of items to return.
	PerPage *uint32 `json:"-"`
	// Page: A positive integer to choose the page to return.
	Page *int32 `json:"-"`
}

// ListPrivateNICsResponse:
type ListPrivateNICsResponse struct {
	// PrivateNics:
	PrivateNics []*PrivateNIC `json:"private_nics"`
	// TotalCount:
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPrivateNICsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPrivateNICsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListPrivateNICsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PrivateNics = append(r.PrivateNics, results.PrivateNics...)
	r.TotalCount += uint64(len(results.PrivateNics))
	return uint64(len(results.PrivateNics)), nil
}

// ListSecurityGroupRulesRequest:
type ListSecurityGroupRulesRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// SecurityGroupID: UUID of the security group.
	SecurityGroupID string `json:"-"`
	// PerPage: A positive integer lower or equal to 100 to select the number of items to return.
	PerPage *uint32 `json:"-"`
	// Page: A positive integer to choose the page to return.
	Page *int32 `json:"-"`
}

// ListSecurityGroupRulesResponse:
type ListSecurityGroupRulesResponse struct {
	// TotalCount: Total number of security groups.
	TotalCount uint32 `json:"total_count"`
	// Rules: List of security rules.
	Rules []*SecurityGroupRule `json:"rules"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSecurityGroupRulesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSecurityGroupRulesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSecurityGroupRulesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Rules = append(r.Rules, results.Rules...)
	r.TotalCount += uint32(len(results.Rules))
	return uint32(len(results.Rules)), nil
}

// ListSecurityGroupsRequest:
type ListSecurityGroupsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Name: Name of the security group.
	Name *string `json:"-"`
	// Organization: Security group Organization ID.
	Organization *string `json:"-"`
	// Project: Security group Project ID.
	Project *string `json:"-"`
	// Tags: List security groups with these exact tags (to filter with several tags, use commas to separate them).
	Tags []string `json:"-"`
	// ProjectDefault: Filter security groups with this value for project_default.
	ProjectDefault *bool `json:"-"`
	// PerPage: A positive integer lower or equal to 100 to select the number of items to return.
	PerPage *uint32 `json:"-"`
	// Page: A positive integer to choose the page to return.
	Page *int32 `json:"-"`
}

// ListSecurityGroupsResponse:
type ListSecurityGroupsResponse struct {
	// TotalCount: Total number of security groups.
	TotalCount uint32 `json:"total_count"`
	// SecurityGroups: List of security groups.
	SecurityGroups []*SecurityGroup `json:"security_groups"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSecurityGroupsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSecurityGroupsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSecurityGroupsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.SecurityGroups = append(r.SecurityGroups, results.SecurityGroups...)
	r.TotalCount += uint32(len(results.SecurityGroups))
	return uint32(len(results.SecurityGroups)), nil
}

// ListServerActionsRequest:
type ListServerActionsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID:
	ServerID string `json:"-"`
}

// ListServerActionsResponse:
type ListServerActionsResponse struct {
	// Actions:
	Actions []ServerAction `json:"actions"`
}

// ListServerUserDataRequest:
type ListServerUserDataRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the Instance.
	ServerID string `json:"-"`
}

// ListServerUserDataResponse:
type ListServerUserDataResponse struct {
	// UserData:
	UserData []string `json:"user_data"`
}

// ListServersRequest:
type ListServersRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// PerPage: A positive integer lower or equal to 100 to select the number of items to return.
	PerPage *uint32 `json:"-"`
	// Page: A positive integer to choose the page to return.
	Page *int32 `json:"-"`
	// Organization: List only Instances of this Organization ID.
	Organization *string `json:"-"`
	// Project: List only Instances of this Project ID.
	Project *string `json:"-"`
	// Name: Filter Instances by name (eg. "server1" will return "server100" and "server1" but not "foo").
	Name *string `json:"-"`
	// PrivateIP: List Instances by private_ip.
	PrivateIP *net.IP `json:"-"`
	// WithoutIP: List Instances that are not attached to a public IP.
	WithoutIP *bool `json:"-"`
	// CommercialType: List Instances of this commercial type.
	CommercialType *string `json:"-"`
	// State: List Instances in this state.
	State *ServerState `json:"-"`
	// Tags: List Instances with these exact tags (to filter with several tags, use commas to separate them).
	Tags []string `json:"-"`
	// PrivateNetwork: List Instances in this Private Network.
	PrivateNetwork *string `json:"-"`
	// Order: Define the order of the returned servers.
	Order ListServersRequestOrder `json:"-"`
	// PrivateNetworks: List Instances from the given Private Networks (use commas to separate them).
	PrivateNetworks []string `json:"-"`
}

// ListServersResponse:
type ListServersResponse struct {
	// TotalCount: Total number of Instances.
	TotalCount uint32 `json:"total_count"`
	// Servers: List of Instances.
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

// ListServersTypesRequest:
type ListServersTypesRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// PerPage:
	PerPage *uint32 `json:"-"`
	// Page:
	Page *int32 `json:"-"`
}

// ListServersTypesResponse:
type ListServersTypesResponse struct {
	// TotalCount: Total number of Instance types.
	TotalCount uint32 `json:"total_count"`
	// Servers: List of Instance types.
	Servers map[string]*ServerType `json:"servers"`
}

// ListSnapshotsRequest:
type ListSnapshotsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Organization:
	Organization *string `json:"-"`
	// PerPage:
	PerPage *uint32 `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// Name:
	Name *string `json:"-"`
	// Project:
	Project *string `json:"-"`
	// Tags:
	Tags *string `json:"-"`
}

// ListSnapshotsResponse:
type ListSnapshotsResponse struct {
	// TotalCount: Total number of snapshots.
	TotalCount uint32 `json:"total_count"`
	// Snapshots: List of snapshots.
	Snapshots []*Snapshot `json:"snapshots"`
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

// ListVolumesRequest:
type ListVolumesRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// VolumeType: Filter by volume type.
	VolumeType *VolumeVolumeType `json:"-"`
	// PerPage: A positive integer lower or equal to 100 to select the number of items to return.
	PerPage *uint32 `json:"-"`
	// Page: A positive integer to choose the page to return.
	Page *int32 `json:"-"`
	// Organization: Filter volume by Organization ID.
	Organization *string `json:"-"`
	// Project: Filter volume by Project ID.
	Project *string `json:"-"`
	// Tags: Filter volumes with these exact tags (to filter with several tags, use commas to separate them).
	Tags []string `json:"-"`
	// Name: Filter volume by name (for eg. "vol" will return "myvolume" but not "data").
	Name *string `json:"-"`
}

// ListVolumesResponse:
type ListVolumesResponse struct {
	// TotalCount: Total number of volumes.
	TotalCount uint32 `json:"total_count"`
	// Volumes: List of volumes.
	Volumes []*Volume `json:"volumes"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListVolumesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListVolumesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListVolumesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Volumes = append(r.Volumes, results.Volumes...)
	r.TotalCount += uint32(len(results.Volumes))
	return uint32(len(results.Volumes)), nil
}

// ListVolumesTypesRequest:
type ListVolumesTypesRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// PerPage:
	PerPage *uint32 `json:"-"`
	// Page:
	Page *int32 `json:"-"`
}

// ListVolumesTypesResponse:
type ListVolumesTypesResponse struct {
	// TotalCount: Total number of volume types.
	TotalCount uint32 `json:"total_count"`
	// Volumes: Map of volume types.
	Volumes map[string]*VolumeType `json:"volumes"`
}

// ServerActionRequest:
type ServerActionRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the Instance.
	ServerID string `json:"-"`
	// Action: Action to perform on the Instance.
	Action ServerAction `json:"action,omitempty"`
	// Name: Name of the backup you want to create.
	// This field should only be specified when performing a backup action.
	Name *string `json:"name,omitempty"`
	// Volumes: For each volume UUID, the snapshot parameters of the volume.
	// This field should only be specified when performing a backup action.
	Volumes map[string]*ServerActionRequestVolumeBackupTemplate `json:"volumes,omitempty"`
}

// ServerActionResponse:
type ServerActionResponse struct {
	// Task:
	Task *Task `json:"task"`
}

// SetImageRequest:
type SetImageRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ID:
	ID string `json:"-"`
	// Name:
	Name string `json:"name,omitempty"`
	// Arch:
	Arch Arch `json:"arch,omitempty"`
	// CreationDate:
	CreationDate *time.Time `json:"creation_date,omitempty"`
	// ModificationDate:
	ModificationDate *time.Time `json:"modification_date,omitempty"`
	// Deprecated: DefaultBootscript:
	DefaultBootscript *Bootscript `json:"default_bootscript,omitempty"`
	// ExtraVolumes:
	ExtraVolumes map[string]*Volume `json:"extra_volumes,omitempty"`
	// FromServer:
	FromServer string `json:"from_server,omitempty"`
	// Organization:
	Organization string `json:"organization,omitempty"`
	// Public:
	Public bool `json:"public,omitempty"`
	// RootVolume:
	RootVolume *VolumeSummary `json:"root_volume,omitempty"`
	// State:
	State ImageState `json:"state,omitempty"`
	// Project:
	Project string `json:"project,omitempty"`
	// Tags:
	Tags *[]string `json:"tags,omitempty"`
}

// SetPlacementGroupRequest:
type SetPlacementGroupRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// PlacementGroupID:
	PlacementGroupID string `json:"-"`
	// Name:
	Name string `json:"name,omitempty"`
	// Organization:
	Organization string `json:"organization,omitempty"`
	// PolicyMode:
	PolicyMode PlacementGroupPolicyMode `json:"policy_mode,omitempty"`
	// PolicyType:
	PolicyType PlacementGroupPolicyType `json:"policy_type,omitempty"`
	// Project:
	Project string `json:"project,omitempty"`
	// Tags:
	Tags *[]string `json:"tags,omitempty"`
}

// SetPlacementGroupResponse:
type SetPlacementGroupResponse struct {
	// PlacementGroup:
	PlacementGroup *PlacementGroup `json:"placement_group"`
}

// SetPlacementGroupServersRequest:
type SetPlacementGroupServersRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// PlacementGroupID: UUID of the placement group you want to set.
	PlacementGroupID string `json:"-"`
	// Servers: An array of the Instances' UUIDs you want to configure.
	Servers []string `json:"servers,omitempty"`
}

// SetPlacementGroupServersResponse:
type SetPlacementGroupServersResponse struct {
	// Servers: Instances attached to the placement group.
	Servers []*PlacementGroupServer `json:"servers"`
}

// SetSecurityGroupRulesRequest:
type SetSecurityGroupRulesRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// SecurityGroupID: UUID of the security group to update the rules on.
	SecurityGroupID string `json:"-"`
	// Rules: List of rules to update in the security group.
	Rules []*SetSecurityGroupRulesRequestRule `json:"rules,omitempty"`
}

// SetSecurityGroupRulesResponse:
type SetSecurityGroupRulesResponse struct {
	// Rules:
	Rules []*SecurityGroupRule `json:"rules"`
}

// UpdateIPRequest:
type UpdateIPRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// IP: IP ID or IP address.
	IP string `json:"-"`
	// Reverse: Reverse domain name.
	Reverse *NullableStringValue `json:"reverse,omitempty"`
	// Type: Convert a 'nat' IP to a 'routed_ipv4'.
	Type IPType `json:"type,omitempty"`
	// Tags: An array of keywords you want to tag this IP with.
	Tags *[]string `json:"tags,omitempty"`
	// Server:
	Server *NullableStringValue `json:"server,omitempty"`
}

// UpdateIPResponse:
type UpdateIPResponse struct {
	// IP:
	IP *IP `json:"ip"`
}

// UpdatePlacementGroupRequest:
type UpdatePlacementGroupRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// PlacementGroupID: UUID of the placement group.
	PlacementGroupID string `json:"-"`
	// Name: Name of the placement group.
	Name *string `json:"name,omitempty"`
	// Tags: Tags of the placement group.
	Tags *[]string `json:"tags,omitempty"`
	// PolicyMode: Operating mode of the placement group.
	PolicyMode *PlacementGroupPolicyMode `json:"policy_mode,omitempty"`
	// PolicyType: Policy type of the placement group.
	PolicyType *PlacementGroupPolicyType `json:"policy_type,omitempty"`
}

// UpdatePlacementGroupResponse:
type UpdatePlacementGroupResponse struct {
	// PlacementGroup:
	PlacementGroup *PlacementGroup `json:"placement_group"`
}

// UpdatePlacementGroupServersRequest:
type UpdatePlacementGroupServersRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// PlacementGroupID: UUID of the placement group you want to update.
	PlacementGroupID string `json:"-"`
	// Servers: An array of the Instances' UUIDs you want to configure.
	Servers []string `json:"servers,omitempty"`
}

// UpdatePlacementGroupServersResponse:
type UpdatePlacementGroupServersResponse struct {
	// Servers: Instances attached to the placement group.
	Servers []*PlacementGroupServer `json:"servers"`
}

// UpdatePrivateNICRequest:
type UpdatePrivateNICRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the Instance the private NIC will be attached to.
	ServerID string `json:"-"`
	// PrivateNicID: Private NIC unique ID.
	PrivateNicID string `json:"-"`
	// Tags: Tags used to select private NIC/s.
	Tags *[]string `json:"tags,omitempty"`
}

// UpdateServerRequest:
type UpdateServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the Instance.
	ServerID string `json:"-"`
	// Name: Name of the Instance.
	Name *string `json:"name,omitempty"`
	// BootType:
	BootType *BootType `json:"boot_type,omitempty"`
	// Tags: Tags of the Instance.
	Tags *[]string `json:"tags,omitempty"`
	// Volumes:
	Volumes *map[string]*VolumeServerTemplate `json:"volumes,omitempty"`
	// Deprecated: Bootscript:
	Bootscript *string `json:"bootscript,omitempty"`
	// DynamicIPRequired:
	DynamicIPRequired *bool `json:"dynamic_ip_required,omitempty"`
	// RoutedIPEnabled: True to configure the instance so it uses the new routed IP mode (once this is set to True you cannot set it back to False).
	RoutedIPEnabled *bool `json:"routed_ip_enabled,omitempty"`
	// PublicIPs:
	PublicIPs []*ServerIP `json:"public_ips,omitempty"`
	// EnableIPv6:
	EnableIPv6 *bool `json:"enable_ipv6,omitempty"`
	// Protected:
	Protected *bool `json:"protected,omitempty"`
	// SecurityGroup:
	SecurityGroup *SecurityGroupTemplate `json:"security_group,omitempty"`
	// PlacementGroup: Placement group ID if Instance must be part of a placement group.
	PlacementGroup *NullableStringValue `json:"placement_group,omitempty"`
	// PrivateNics: Instance private NICs.
	PrivateNics []*PrivateNIC `json:"private_nics,omitempty"`
	// CommercialType: Warning: This field has some restrictions:
	// - Cannot be changed if the Instance is not in `stopped` state.
	// - Cannot be changed if the Instance is in a placement group.
	// - Local storage requirements of the target commercial_types must be fulfilled (i.e. if an Instance has 80GB of local storage, it can be changed into a GP1-XS, which has a maximum of 150GB, but it cannot be changed into a DEV1-S, which has only 20GB).
	CommercialType *string `json:"commercial_type,omitempty"`
}

// UpdateServerResponse:
type UpdateServerResponse struct {
	// Server:
	Server *Server `json:"server"`
}

// UpdateVolumeRequest:
type UpdateVolumeRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// VolumeID: UUID of the volume.
	VolumeID string `json:"-"`
	// Name: Volume name.
	Name *string `json:"name,omitempty"`
	// Tags: Tags of the volume.
	Tags *[]string `json:"tags,omitempty"`
	// Size: Volume disk size, must be a multiple of 512.
	Size *scw.Size `json:"size,omitempty"`
}

// UpdateVolumeResponse:
type UpdateVolumeResponse struct {
	// Volume:
	Volume *Volume `json:"volume"`
}

// setImageResponse:
type setImageResponse struct {
	// Image:
	Image *Image `json:"image"`
}

// setSecurityGroupRequest:
type setSecurityGroupRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ID: ID of the security group (will be ignored).
	ID string `json:"-"`
	// Name: Name of the security group.
	Name string `json:"name,omitempty"`
	// Tags: Tags of the security group.
	Tags *[]string `json:"tags,omitempty"`
	// CreationDate: Creation date of the security group (will be ignored).
	CreationDate *time.Time `json:"creation_date,omitempty"`
	// ModificationDate: Modification date of the security group (will be ignored).
	ModificationDate *time.Time `json:"modification_date,omitempty"`
	// Description: Description of the security group.
	Description string `json:"description,omitempty"`
	// EnableDefaultSecurity: True to block SMTP on IPv4 and IPv6. This feature is read only, please open a support ticket if you need to make it configurable.
	EnableDefaultSecurity bool `json:"enable_default_security,omitempty"`
	// InboundDefaultPolicy: Default inbound policy.
	InboundDefaultPolicy SecurityGroupPolicy `json:"inbound_default_policy,omitempty"`
	// OutboundDefaultPolicy: Default outbound policy.
	OutboundDefaultPolicy SecurityGroupPolicy `json:"outbound_default_policy,omitempty"`
	// Organization: Security groups Organization ID.
	Organization string `json:"organization,omitempty"`
	// Project: Security group Project ID.
	Project string `json:"project,omitempty"`
	// Deprecated: OrganizationDefault: Please use project_default instead.
	OrganizationDefault *bool `json:"organization_default,omitempty"`
	// ProjectDefault: True use this security group for future Instances created in this project.
	ProjectDefault bool `json:"project_default,omitempty"`
	// Servers: Instances attached to this security group.
	Servers []*ServerSummary `json:"servers,omitempty"`
	// Stateful: True to set the security group as stateful.
	Stateful bool `json:"stateful,omitempty"`
}

// setSecurityGroupResponse:
type setSecurityGroupResponse struct {
	// SecurityGroup:
	SecurityGroup *SecurityGroup `json:"security_group"`
}

// setSecurityGroupRuleRequest:
type setSecurityGroupRuleRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// SecurityGroupID:
	SecurityGroupID string `json:"-"`
	// SecurityGroupRuleID:
	SecurityGroupRuleID string `json:"-"`
	// ID:
	ID string `json:"id,omitempty"`
	// Protocol:
	Protocol SecurityGroupRuleProtocol `json:"protocol,omitempty"`
	// Direction:
	Direction SecurityGroupRuleDirection `json:"direction,omitempty"`
	// Action:
	Action SecurityGroupRuleAction `json:"action,omitempty"`
	// IPRange:
	IPRange scw.IPNet `json:"ip_range,omitempty"`
	// DestPortFrom:
	DestPortFrom *uint32 `json:"dest_port_from,omitempty"`
	// DestPortTo:
	DestPortTo *uint32 `json:"dest_port_to,omitempty"`
	// Position:
	Position uint32 `json:"position,omitempty"`
	// Editable:
	Editable bool `json:"editable,omitempty"`
}

// setSecurityGroupRuleResponse:
type setSecurityGroupRuleResponse struct {
	// Rule:
	Rule *SecurityGroupRule `json:"rule"`
}

// setServerRequest:
type setServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ID: Instance unique ID.
	ID string `json:"-"`
	// Name: Instance name.
	Name string `json:"name,omitempty"`
	// Organization: Instance Organization ID.
	Organization string `json:"organization,omitempty"`
	// Project: Instance Project ID.
	Project string `json:"project,omitempty"`
	// AllowedActions: Provide a list of allowed actions on the server.
	AllowedActions []ServerAction `json:"allowed_actions,omitempty"`
	// Tags: Tags associated with the Instance.
	Tags *[]string `json:"tags,omitempty"`
	// CommercialType: Instance commercial type (eg. GP1-M).
	CommercialType string `json:"commercial_type,omitempty"`
	// CreationDate: Instance creation date.
	CreationDate *time.Time `json:"creation_date,omitempty"`
	// DynamicIPRequired: True if a dynamic IPv4 is required.
	DynamicIPRequired bool `json:"dynamic_ip_required,omitempty"`
	// RoutedIPEnabled: True to configure the instance so it uses the new routed IP mode (once this is set to True you cannot set it back to False).
	RoutedIPEnabled *bool `json:"routed_ip_enabled,omitempty"`
	// EnableIPv6: True if IPv6 is enabled.
	EnableIPv6 bool `json:"enable_ipv6,omitempty"`
	// Hostname: Instance host name.
	Hostname string `json:"hostname,omitempty"`
	// Image: Provide information on the Instance image.
	Image *Image `json:"image,omitempty"`
	// Protected: Instance protection option is activated.
	Protected bool `json:"protected,omitempty"`
	// PrivateIP: Instance private IP address.
	PrivateIP *string `json:"private_ip,omitempty"`
	// PublicIP: Information about the public IP.
	PublicIP *ServerIP `json:"public_ip,omitempty"`
	// PublicIPs: Information about all the public IPs attached to the server.
	PublicIPs []*ServerIP `json:"public_ips,omitempty"`
	// ModificationDate: Instance modification date.
	ModificationDate *time.Time `json:"modification_date,omitempty"`
	// State: Instance state.
	State ServerState `json:"state,omitempty"`
	// Location: Instance location.
	Location *ServerLocation `json:"location,omitempty"`
	// IPv6: Instance IPv6 address.
	IPv6 *ServerIPv6 `json:"ipv6,omitempty"`
	// Deprecated: Bootscript: Instance bootscript.
	Bootscript *Bootscript `json:"bootscript,omitempty"`
	// BootType: Instance boot type.
	BootType BootType `json:"boot_type,omitempty"`
	// Volumes: Instance volumes.
	Volumes map[string]*Volume `json:"volumes,omitempty"`
	// SecurityGroup: Instance security group.
	SecurityGroup *SecurityGroupSummary `json:"security_group,omitempty"`
	// Maintenances: Instance planned maintenances.
	Maintenances []*ServerMaintenance `json:"maintenances,omitempty"`
	// StateDetail: Instance state_detail.
	StateDetail string `json:"state_detail,omitempty"`
	// Arch: Instance architecture (refers to the CPU architecture used for the Instance, e.g. x86_64, arm64).
	Arch Arch `json:"arch,omitempty"`
	// PlacementGroup: Instance placement group.
	PlacementGroup *PlacementGroup `json:"placement_group,omitempty"`
	// PrivateNics: Instance private NICs.
	PrivateNics []*PrivateNIC `json:"private_nics,omitempty"`
}

// setServerResponse:
type setServerResponse struct {
	// Server:
	Server *Server `json:"server"`
}

// setSnapshotRequest:
type setSnapshotRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// SnapshotID:
	SnapshotID string `json:"-"`
	// ID:
	ID string `json:"id,omitempty"`
	// Name:
	Name string `json:"name,omitempty"`
	// Organization:
	Organization string `json:"organization,omitempty"`
	// VolumeType:
	VolumeType VolumeVolumeType `json:"volume_type,omitempty"`
	// Size:
	Size scw.Size `json:"size,omitempty"`
	// State:
	State SnapshotState `json:"state,omitempty"`
	// BaseVolume:
	BaseVolume *SnapshotBaseVolume `json:"base_volume,omitempty"`
	// CreationDate:
	CreationDate *time.Time `json:"creation_date,omitempty"`
	// ModificationDate:
	ModificationDate *time.Time `json:"modification_date,omitempty"`
	// Project:
	Project string `json:"project,omitempty"`
	// Tags:
	Tags *[]string `json:"tags,omitempty"`
}

// setSnapshotResponse:
type setSnapshotResponse struct {
	// Snapshot:
	Snapshot *Snapshot `json:"snapshot"`
}

// Scaleway Instances are virtual machines in the cloud. Different [Instance types](https://www.scaleway.com/en/docs/compute/instances/reference-content/choosing-instance-type/) offer different technical specifications in terms of vCPU, RAM, bandwidth and storage. Once you have created your Instance and installed your image of choice (e.g. an operating system), you can [connect to your Instance via SSH](https://www.scaleway.com/en/docs/compute/instances/how-to/connect-to-instance/) to use it as you wish. When you are done using the Instance, you can delete it from your account.
//
// (switchcolumn)
// <Message type="tip">
// To retrieve information about the different [images](#path-images) available to install on Scaleway Instances, check out our [Marketplace API](https://www.scaleway.com/en/developers/api/marketplace).
// </Message>
// (switchcolumn)
//
// ## Concepts
//
// Refer to our [dedicated concepts page](https://www.scaleway.com/en/docs/compute/instances/concepts/) to find definitions of all concepts and terminology related to Instances.
//
// (switchcolumn)
// (switchcolumn)
//
// ## Quickstart
//
// 1. Configure your environment variables
//
//	<Message type="note">
//	This is an optional step that seeks to simplify your usage of the Instances API. See [Availability Zones](#availability-zones) below for help choosing an Availability Zone. You can find your Project ID in the [Scaleway console](https://console.scaleway.com/project/settings).
//	</Message>
//
//	```bash
//	export SCW_SECRET_KEY="<API secret key>"
//	export SCW_DEFAULT_ZONE="<Scaleway Availability Zone>"
//	export SCW_PROJECT_ID="<Scaleway Project ID>"
//	```
//
// 2. **Create an Instance**: Run the following command to create an Instance. You can customize the details in the payload (name, description, type, tags etc) to your needs: use the information below to adjust the payload as necessary.
//
//	```bash
//	curl -X POST \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  -H "Content-Type: application/json" \
//	  "https://api.scaleway.com/instance/v1/zones/$SCW_DEFAULT_ZONE/servers" \
//	    -d '{
//	      "name": "my-new-instance",
//	      "project": "'"$SCW_PROJECT_ID"'",
//	      "commercial_type": "GP1-S",
//	      "image": "544f0add-626b-4e4f-8a96-79fa4414d99a",
//	      "enable_ipv6": true,
//	      "volumes": {
//	        "0":{
//	          "name": "my-volume",
//	          "size": 300000000000,
//	          "volume_type": "l_ssd"
//	        }
//	      }
//	    }'
//	```
//
//	    | Parameter       | Description                                                                                                                                              | Valid values                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
//	| --------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
//	| `name`            | A name of your choice for the Instance (string)                                                                                                          | Any string containing only alphanumeric characters, dots, spaces and dashes, e.g. `"my-new-instance"`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
//	| `project`         | The Project in which the Instance should be created (string)                                                                                             | Any valid Scaleway Project ID (see above), e.g. `"b4bd99e0-b389-11ed-afa1-0242ac120002"`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
//	| `commercial-type` | The commercial Instance type to create (string)                                                                                                          | Any valid ID of a Scaleway commercial Instance type, e.g. `"GP1-S"`, `"PRO2-M"`. Use the [List Instance Types](#path-instance-types-list-instance-types) endpoint to get a list of all valid Instance types and their IDs.                                                                                                                                                                                                                                                                               |
//	| `image`           | The image to install on the Instance, e.g. a particular OS (string)                                                                                      | Any valid Scaleway image ID, e.g. `"544f0add-626b-4e4f-8a96-79fa4414d99a"` which is the ID for the `Ubuntu 22.04 Jammy Jellyfish` image. Use the [List Instance Images](#path-images-list-instance-images) endpoint to get a list of all available images and their IDs, or check out the [Scaleway Marketplace API](https://www.scaleway.com/en/developers/api/marketplace).                                                                                                                                                                                                                                                                                                     |
//	| `enable_ipv6`     | Whether to enable IPv6 on the Instance (boolean)                                                                                                         | `true` or `false`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
//	| `volumes`         | An object that specifies the storage volumes to attach to the Instance. For more information, see **Creating an Instance: the volumes object** in the [Technical information](#technical-information) section of this quickstart. | A (dictionary) object with a minimum of one key (`"0"`) whose value is another object containing the parameters `"name"` (a name for the volume), `"size"` (the size for the volume, in bytes), and `"volume_type"` (`"l_ssd"`, `"b_ssd"` or `"unified"`). Additional keys for additional volumes should increment by 1 each time (the second volume would have a key of `1`.) Further parameters are available, and it is possible to attach existing volumes rather than creating a new one, or create a volume from a snapshot. |
//
// 3. **List your Instances**: run the following command to get a list of all the Instances in your account, with their details:
//
//	```bash
//	curl -X GET \
//	  -H "Content-Type: application/json" \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  "https://api.scaleway.com/instance/v1/zones/$SCW_DEFAULT_ZONE/servers/"
//	```
//
// 4. **Delete an Instance**: run the following command to delete an Instance, specified by its Instance ID:
//
//	```bash
//	curl -X DELETE \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  -H "Content-Type: application/json" \
//	  "https://api.scaleway.com/instance/v1/zones/$SCW_DEFAULT_ZONE/servers/<Instance-ID>"
//	```
//
//	The expected successful response is empty.
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
// Instances can be deployed in the following Availability Zones:
//
// | Name      | API ID                |
// |-----------|-----------------------|
// | Paris     | `fr-par-1` `fr-par-2` `fr-par-3` |
// | Amsterdam | `nl-ams-1` `nl-ams-2` |
// | Warsaw    | `pl-waw-1` `pl-waw-2` |
//
// (switchcolumn)
// (switchcolumn)
//
// ### Pagination
//
// Most listing requests receive a paginated response. Requests against paginated endpoints accept two `query` arguments:
//
// - `page`, a positive integer to choose which page to return.
// - `per_page`, an positive integer lower or equal to 100 to select the number of items to return per page. The default value is `50`.
//
// Paginated endpoints usually also accept filters to search and sort results.These filters are documented along each endpoint documentation.
//
// The `X-Total-Count` header contains the total number of items returned.
//
// (switchcolumn)
// (switchcolumn)
//
// ### Creating an Instance: the volumes object
//
// When [creating an Instance](#path-instances-create-an-instance), the `volumes` object is a required part of the payload. This is a dictionary with a minimum of one key (`"0"`) whose value is another object setting parameters for that volume. Additional keys for additional volumes should increment by 1 each time (the second volume would have a key of `1`.)
//
// Note that volume `size` must respect the volume constraints of the Instance's `commercial_type`: for each type of Instance, a minimum amount of storage is required, and there is also a maximum that cannot be exceeded. Some Instance types support only Block Storage (`b_ssd`), others also support local storage (`l_ssd`) ). Read more about these constraints in the [List Instance types](#path-instance-types-list-instance-types) documentation, specifically the `volume_constraints` parameter for each type listed in the response
//
// You can use the `volumes` object in different ways. The table below shows which parameters are required for each of the following use cases:
//
// | Use case                | Required params       | Optional params     | Notes                                  |
// |-------------------------|-----------------------|---------------------|----------------------------------------|
// | Create a volume from a snapshot of an image  |  | `volume_type`, `size`, `boot` | If the `size` parameter is not set, the size of the volume will equal the size of the corresponding snapshot of the image. |
// | Attach an existing volume   | `id`, `name` | `boot` |  |
// | Create an empty volume      | `name`, `volume_type`, `size` | `organization`, `project`, `boot` |  |
// | Create a volume from a snapshot     | `base_snapshot`, `name`, `volume_type` | `organization`, `project`, `boot` |  |
//
// (switchcolumn)
// <Message type="note">
// This information is designed to help you correctly configure the `volumes` object when using the [Create an Instance](#path-instances-create-an-instance) or [Update an Instance](#path-instances-update-an-instance) methods.
// </Message>
// (switchcolumn)
//
// ## Going further
//
// For more help using Scaleway Instances, check out the following resources:
// - Our [main documentation](https://www.scaleway.com/en/docs/compute/instances/)
// - The #instance channel on our [Slack Community](https://www.scaleway.com/en/docs/tutorials/scaleway-slack-community/)
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
func (s *API) Zones() []scw.Zone {
	return []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3}
}

// GetServerTypesAvailability: Get availability for all Instance types.
func (s *API) GetServerTypesAvailability(req *GetServerTypesAvailabilityRequest, opts ...scw.RequestOption) (*GetServerTypesAvailabilityResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/products/servers/availability",
		Query:  query,
	}

	var resp GetServerTypesAvailabilityResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListServersTypes: List available Instance types and their technical details.
func (s *API) ListServersTypes(req *ListServersTypesRequest, opts ...scw.RequestOption) (*ListServersTypesResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/products/servers",
		Query:  query,
	}

	var resp ListServersTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVolumesTypes: List all volume types and their technical details.
func (s *API) ListVolumesTypes(req *ListVolumesTypesRequest, opts ...scw.RequestOption) (*ListVolumesTypesResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/products/volumes",
		Query:  query,
	}

	var resp ListVolumesTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListServers: List all Instances in a specified Availability Zone, e.g. `fr-par-1`.
func (s *API) ListServers(req *ListServersRequest, opts ...scw.RequestOption) (*ListServersResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "project", req.Project)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "private_ip", req.PrivateIP)
	parameter.AddToQuery(query, "without_ip", req.WithoutIP)
	parameter.AddToQuery(query, "commercial_type", req.CommercialType)
	parameter.AddToQuery(query, "state", req.State)
	if len(req.Tags) != 0 {
		parameter.AddToQuery(query, "tags", strings.Join(req.Tags, ","))
	}
	parameter.AddToQuery(query, "private_network", req.PrivateNetwork)
	parameter.AddToQuery(query, "order", req.Order)
	if len(req.PrivateNetworks) != 0 {
		parameter.AddToQuery(query, "private_networks", strings.Join(req.PrivateNetworks, ","))
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers",
		Query:  query,
	}

	var resp ListServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// createServer: Create a new Instance of the specified commercial type in the specified zone. Pay attention to the volumes parameter, which takes an object which can be used in different ways to achieve different behaviors.
// Get more information in the [Technical Information](#technical-information) section of the introduction.
func (s *API) createServer(req *CreateServerRequest, opts ...scw.RequestOption) (*CreateServerResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultOrganization, exist := s.client.GetDefaultOrganizationID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Organization = &defaultOrganization
	}

	defaultProject, exist := s.client.GetDefaultProjectID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Project = &defaultProject
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("srv")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateServerResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteServer: Delete the Instance with the specified ID.
func (s *API) DeleteServer(req *DeleteServerRequest, opts ...scw.RequestOption) error {
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
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetServer: Get the details of a specified Instance.
func (s *API) GetServer(req *GetServerRequest, opts ...scw.RequestOption) (*GetServerResponse, error) {
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
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	var resp GetServerResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// setServer:
func (s *API) setServer(req *setServerRequest, opts ...scw.RequestOption) (*setServerResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}
	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}
	if req.Project == "" {
		defaultProject, _ := s.client.GetDefaultProjectID()
		req.Project = defaultProject
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ID) == "" {
		return nil, errors.New("field ID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp setServerResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// updateServer: Update the Instance information, such as name, boot mode, or tags.
func (s *API) updateServer(req *UpdateServerRequest, opts ...scw.RequestOption) (*UpdateServerResponse, error) {
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
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UpdateServerResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListServerActions: List all actions (e.g. power on, power off, reboot) that can currently be performed on an Instance.
func (s *API) ListServerActions(req *ListServerActionsRequest, opts ...scw.RequestOption) (*ListServerActionsResponse, error) {
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
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/action",
	}

	var resp ListServerActionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ServerAction: Perform an action on an Instance.
// Available actions are:
// * `poweron`: Start a stopped Instance.
// * `poweroff`: Fully stop the Instance and release the hypervisor slot.
// * `stop_in_place`: Stop the Instance, but keep the slot on the hypervisor.
// * `reboot`: Stop the instance and restart it.
// * `backup`:  Create an image with all the volumes of an Instance.
// * `terminate`: Delete the Instance along with all attached volumes.
// * `enable_routed_ip`: Migrate the Instance to the new network stack.
//
// Keep in mind that terminating an Instance will result in the deletion of all attached volumes, including local and block storage.
// If you want to preserve your local volumes, you should use the `archive` action instead of `terminate`. Similarly, if you want to keep your block storage volumes, you must first detach them before issuing the `terminate` command.
// For more information, read the [Volumes](#path-volumes-list-volumes) documentation.
func (s *API) ServerAction(req *ServerActionRequest, opts ...scw.RequestOption) (*ServerActionResponse, error) {
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
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/action",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ServerActionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListServerUserData: List all user data keys registered on a specified Instance.
func (s *API) ListServerUserData(req *ListServerUserDataRequest, opts ...scw.RequestOption) (*ListServerUserDataResponse, error) {
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
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/user_data",
	}

	var resp ListServerUserDataResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteServerUserData: Delete the specified key from an Instance's user data.
func (s *API) DeleteServerUserData(req *DeleteServerUserDataRequest, opts ...scw.RequestOption) error {
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

	if fmt.Sprint(req.Key) == "" {
		return errors.New("field Key cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/user_data/" + fmt.Sprint(req.Key) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListImages: List all existing Instance images.
func (s *API) ListImages(req *ListImagesRequest, opts ...scw.RequestOption) (*ListImagesResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "public", req.Public)
	parameter.AddToQuery(query, "arch", req.Arch)
	parameter.AddToQuery(query, "project", req.Project)
	parameter.AddToQuery(query, "tags", req.Tags)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/images",
		Query:  query,
	}

	var resp ListImagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetImage: Get details of an image with the specified ID.
func (s *API) GetImage(req *GetImageRequest, opts ...scw.RequestOption) (*GetImageResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ImageID) == "" {
		return nil, errors.New("field ImageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/images/" + fmt.Sprint(req.ImageID) + "",
	}

	var resp GetImageResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateImage: Create an Instance image from the specified snapshot ID.
func (s *API) CreateImage(req *CreateImageRequest, opts ...scw.RequestOption) (*CreateImageResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultOrganization, exist := s.client.GetDefaultOrganizationID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Organization = &defaultOrganization
	}

	defaultProject, exist := s.client.GetDefaultProjectID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Project = &defaultProject
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("img")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/images",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateImageResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// setImage: Replace all image properties with an image message.
func (s *API) setImage(req *SetImageRequest, opts ...scw.RequestOption) (*setImageResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}
	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}
	if req.Project == "" {
		defaultProject, _ := s.client.GetDefaultProjectID()
		req.Project = defaultProject
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ID) == "" {
		return nil, errors.New("field ID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/images/" + fmt.Sprint(req.ID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp setImageResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteImage: Delete the image with the specified ID.
func (s *API) DeleteImage(req *DeleteImageRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ImageID) == "" {
		return errors.New("field ImageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/images/" + fmt.Sprint(req.ImageID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListSnapshots: List all snapshots of an Organization in a specified Availability Zone.
func (s *API) ListSnapshots(req *ListSnapshotsRequest, opts ...scw.RequestOption) (*ListSnapshotsResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "project", req.Project)
	parameter.AddToQuery(query, "tags", req.Tags)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots",
		Query:  query,
	}

	var resp ListSnapshotsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSnapshot: Create a snapshot from a specified volume or from a QCOW2 file in a specified Availability Zone.
func (s *API) CreateSnapshot(req *CreateSnapshotRequest, opts ...scw.RequestOption) (*CreateSnapshotResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultOrganization, exist := s.client.GetDefaultOrganizationID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Organization = &defaultOrganization
	}

	defaultProject, exist := s.client.GetDefaultProjectID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Project = &defaultProject
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("snp")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateSnapshotResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSnapshot: Get details of a snapshot with the specified ID.
func (s *API) GetSnapshot(req *GetSnapshotRequest, opts ...scw.RequestOption) (*GetSnapshotResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SnapshotID) == "" {
		return nil, errors.New("field SnapshotID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
	}

	var resp GetSnapshotResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// setSnapshot: Replace all snapshot properties with a snapshot message.
func (s *API) setSnapshot(req *setSnapshotRequest, opts ...scw.RequestOption) (*setSnapshotResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}
	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}
	if req.Project == "" {
		defaultProject, _ := s.client.GetDefaultProjectID()
		req.Project = defaultProject
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SnapshotID) == "" {
		return nil, errors.New("field SnapshotID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp setSnapshotResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSnapshot: Delete the snapshot with the specified ID.
func (s *API) DeleteSnapshot(req *DeleteSnapshotRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SnapshotID) == "" {
		return errors.New("field SnapshotID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ExportSnapshot: Export a snapshot to a specified S3 bucket in the same region.
func (s *API) ExportSnapshot(req *ExportSnapshotRequest, opts ...scw.RequestOption) (*ExportSnapshotResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SnapshotID) == "" {
		return nil, errors.New("field SnapshotID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "/export",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ExportSnapshotResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVolumes: List volumes in the specified Availability Zone. You can filter the output by volume type.
func (s *API) ListVolumes(req *ListVolumesRequest, opts ...scw.RequestOption) (*ListVolumesResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "volume_type", req.VolumeType)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "project", req.Project)
	if len(req.Tags) != 0 {
		parameter.AddToQuery(query, "tags", strings.Join(req.Tags, ","))
	}
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/volumes",
		Query:  query,
	}

	var resp ListVolumesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVolume: Create a volume of a specified type in an Availability Zone.
func (s *API) CreateVolume(req *CreateVolumeRequest, opts ...scw.RequestOption) (*CreateVolumeResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultOrganization, exist := s.client.GetDefaultOrganizationID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Organization = &defaultOrganization
	}

	defaultProject, exist := s.client.GetDefaultProjectID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Project = &defaultProject
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("vol")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/volumes",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateVolumeResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVolume: Get details of a volume with the specified ID.
func (s *API) GetVolume(req *GetVolumeRequest, opts ...scw.RequestOption) (*GetVolumeResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.VolumeID) == "" {
		return nil, errors.New("field VolumeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/volumes/" + fmt.Sprint(req.VolumeID) + "",
	}

	var resp GetVolumeResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVolume: Replace the name and/or size properties of a volume specified by its ID, with the specified value(s). Any volume name can be changed, however only `b_ssd` volumes can currently be increased in size.
func (s *API) UpdateVolume(req *UpdateVolumeRequest, opts ...scw.RequestOption) (*UpdateVolumeResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.VolumeID) == "" {
		return nil, errors.New("field VolumeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/volumes/" + fmt.Sprint(req.VolumeID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UpdateVolumeResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVolume: Delete the volume with the specified ID.
func (s *API) DeleteVolume(req *DeleteVolumeRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.VolumeID) == "" {
		return errors.New("field VolumeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/volumes/" + fmt.Sprint(req.VolumeID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListSecurityGroups: List all existing security groups.
func (s *API) ListSecurityGroups(req *ListSecurityGroupsRequest, opts ...scw.RequestOption) (*ListSecurityGroupsResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "project", req.Project)
	if len(req.Tags) != 0 {
		parameter.AddToQuery(query, "tags", strings.Join(req.Tags, ","))
	}
	parameter.AddToQuery(query, "project_default", req.ProjectDefault)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups",
		Query:  query,
	}

	var resp ListSecurityGroupsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSecurityGroup: Create a security group with a specified name and description.
func (s *API) CreateSecurityGroup(req *CreateSecurityGroupRequest, opts ...scw.RequestOption) (*CreateSecurityGroupResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultOrganization, exist := s.client.GetDefaultOrganizationID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Organization = &defaultOrganization
	}

	defaultProject, exist := s.client.GetDefaultProjectID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Project = &defaultProject
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("sg")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateSecurityGroupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSecurityGroup: Get the details of a security group with the specified ID.
func (s *API) GetSecurityGroup(req *GetSecurityGroupRequest, opts ...scw.RequestOption) (*GetSecurityGroupResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return nil, errors.New("field SecurityGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "",
	}

	var resp GetSecurityGroupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSecurityGroup: Delete a security group with the specified ID.
func (s *API) DeleteSecurityGroup(req *DeleteSecurityGroupRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return errors.New("field SecurityGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// setSecurityGroup: Replace all security group properties with a security group message.
func (s *API) setSecurityGroup(req *setSecurityGroupRequest, opts ...scw.RequestOption) (*setSecurityGroupResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}
	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}
	if req.Project == "" {
		defaultProject, _ := s.client.GetDefaultProjectID()
		req.Project = defaultProject
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ID) == "" {
		return nil, errors.New("field ID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.ID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp setSecurityGroupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDefaultSecurityGroupRules: Lists the default rules applied to all the security groups.
func (s *API) ListDefaultSecurityGroupRules(req *ListDefaultSecurityGroupRulesRequest, opts ...scw.RequestOption) (*ListSecurityGroupRulesResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/default/rules",
	}

	var resp ListSecurityGroupRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSecurityGroupRules: List the rules of the a specified security group ID.
func (s *API) ListSecurityGroupRules(req *ListSecurityGroupRulesRequest, opts ...scw.RequestOption) (*ListSecurityGroupRulesResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return nil, errors.New("field SecurityGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "/rules",
		Query:  query,
	}

	var resp ListSecurityGroupRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSecurityGroupRule: Create a rule in the specified security group ID.
func (s *API) CreateSecurityGroupRule(req *CreateSecurityGroupRuleRequest, opts ...scw.RequestOption) (*CreateSecurityGroupRuleResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return nil, errors.New("field SecurityGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "/rules",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateSecurityGroupRuleResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetSecurityGroupRules: Replaces the existing rules of the security group with the rules provided. This endpoint supports the update of existing rules, creation of new rules and deletion of existing rules when they are not passed in the request.
func (s *API) SetSecurityGroupRules(req *SetSecurityGroupRulesRequest, opts ...scw.RequestOption) (*SetSecurityGroupRulesResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return nil, errors.New("field SecurityGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "/rules",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetSecurityGroupRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSecurityGroupRule: Delete a security group rule with the specified ID.
func (s *API) DeleteSecurityGroupRule(req *DeleteSecurityGroupRuleRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return errors.New("field SecurityGroupID cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupRuleID) == "" {
		return errors.New("field SecurityGroupRuleID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "/rules/" + fmt.Sprint(req.SecurityGroupRuleID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetSecurityGroupRule: Get details of a security group rule with the specified ID.
func (s *API) GetSecurityGroupRule(req *GetSecurityGroupRuleRequest, opts ...scw.RequestOption) (*GetSecurityGroupRuleResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return nil, errors.New("field SecurityGroupID cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupRuleID) == "" {
		return nil, errors.New("field SecurityGroupRuleID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "/rules/" + fmt.Sprint(req.SecurityGroupRuleID) + "",
	}

	var resp GetSecurityGroupRuleResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// setSecurityGroupRule: Update the rule of a specified security group ID.
func (s *API) setSecurityGroupRule(req *setSecurityGroupRuleRequest, opts ...scw.RequestOption) (*setSecurityGroupRuleResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return nil, errors.New("field SecurityGroupID cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupRuleID) == "" {
		return nil, errors.New("field SecurityGroupRuleID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "/rules/" + fmt.Sprint(req.SecurityGroupRuleID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp setSecurityGroupRuleResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPlacementGroups: List all placement groups in a specified Availability Zone.
func (s *API) ListPlacementGroups(req *ListPlacementGroupsRequest, opts ...scw.RequestOption) (*ListPlacementGroupsResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "project", req.Project)
	if len(req.Tags) != 0 {
		parameter.AddToQuery(query, "tags", strings.Join(req.Tags, ","))
	}
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups",
		Query:  query,
	}

	var resp ListPlacementGroupsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreatePlacementGroup: Create a new placement group in a specified Availability Zone.
func (s *API) CreatePlacementGroup(req *CreatePlacementGroupRequest, opts ...scw.RequestOption) (*CreatePlacementGroupResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultOrganization, exist := s.client.GetDefaultOrganizationID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Organization = &defaultOrganization
	}

	defaultProject, exist := s.client.GetDefaultProjectID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Project = &defaultProject
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("pg")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreatePlacementGroupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPlacementGroup: Get the specified placement group.
func (s *API) GetPlacementGroup(req *GetPlacementGroupRequest, opts ...scw.RequestOption) (*GetPlacementGroupResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return nil, errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "",
	}

	var resp GetPlacementGroupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetPlacementGroup: Set all parameters of the specified placement group.
func (s *API) SetPlacementGroup(req *SetPlacementGroupRequest, opts ...scw.RequestOption) (*SetPlacementGroupResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}
	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}
	if req.Project == "" {
		defaultProject, _ := s.client.GetDefaultProjectID()
		req.Project = defaultProject
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return nil, errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetPlacementGroupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePlacementGroup: Update one or more parameter of the specified placement group.
func (s *API) UpdatePlacementGroup(req *UpdatePlacementGroupRequest, opts ...scw.RequestOption) (*UpdatePlacementGroupResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return nil, errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UpdatePlacementGroupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeletePlacementGroup: Delete the specified placement group.
func (s *API) DeletePlacementGroup(req *DeletePlacementGroupRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetPlacementGroupServers: Get all Instances belonging to the specified placement group.
func (s *API) GetPlacementGroupServers(req *GetPlacementGroupServersRequest, opts ...scw.RequestOption) (*GetPlacementGroupServersResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return nil, errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "/servers",
	}

	var resp GetPlacementGroupServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetPlacementGroupServers: Set all Instances belonging to the specified placement group.
func (s *API) SetPlacementGroupServers(req *SetPlacementGroupServersRequest, opts ...scw.RequestOption) (*SetPlacementGroupServersResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return nil, errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "/servers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetPlacementGroupServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePlacementGroupServers: Update all Instances belonging to the specified placement group.
func (s *API) UpdatePlacementGroupServers(req *UpdatePlacementGroupServersRequest, opts ...scw.RequestOption) (*UpdatePlacementGroupServersResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return nil, errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "/servers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UpdatePlacementGroupServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListIPs: List all flexible IPs in a specified zone.
func (s *API) ListIPs(req *ListIPsRequest, opts ...scw.RequestOption) (*ListIPsResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project", req.Project)
	parameter.AddToQuery(query, "organization", req.Organization)
	if len(req.Tags) != 0 {
		parameter.AddToQuery(query, "tags", strings.Join(req.Tags, ","))
	}
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips",
		Query:  query,
	}

	var resp ListIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateIP: Reserve a flexible IP and attach it to the specified Instance.
func (s *API) CreateIP(req *CreateIPRequest, opts ...scw.RequestOption) (*CreateIPResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultOrganization, exist := s.client.GetDefaultOrganizationID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Organization = &defaultOrganization
	}

	defaultProject, exist := s.client.GetDefaultProjectID()
	if exist && req.Organization == nil && req.Project == nil {
		req.Project = &defaultProject
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateIPResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIP: Get details of an IP with the specified ID or address.
func (s *API) GetIP(req *GetIPRequest, opts ...scw.RequestOption) (*GetIPResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.IP) == "" {
		return nil, errors.New("field IP cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IP) + "",
	}

	var resp GetIPResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateIP: Update a flexible IP in the specified zone with the specified ID.
func (s *API) UpdateIP(req *UpdateIPRequest, opts ...scw.RequestOption) (*UpdateIPResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.IP) == "" {
		return nil, errors.New("field IP cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IP) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UpdateIPResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteIP: Delete the IP with the specified ID.
func (s *API) DeleteIP(req *DeleteIPRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.IP) == "" {
		return errors.New("field IP cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IP) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListPrivateNICs: List all private NICs of a specified Instance.
func (s *API) ListPrivateNICs(req *ListPrivateNICsRequest, opts ...scw.RequestOption) (*ListPrivateNICsResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	if len(req.Tags) != 0 {
		parameter.AddToQuery(query, "tags", strings.Join(req.Tags, ","))
	}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/private_nics",
		Query:  query,
	}

	var resp ListPrivateNICsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreatePrivateNIC: Create a private NIC connecting an Instance to a Private Network.
func (s *API) CreatePrivateNIC(req *CreatePrivateNICRequest, opts ...scw.RequestOption) (*CreatePrivateNICResponse, error) {
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
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/private_nics",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreatePrivateNICResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrivateNIC: Get private NIC properties.
func (s *API) GetPrivateNIC(req *GetPrivateNICRequest, opts ...scw.RequestOption) (*GetPrivateNICResponse, error) {
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

	if fmt.Sprint(req.PrivateNicID) == "" {
		return nil, errors.New("field PrivateNicID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/private_nics/" + fmt.Sprint(req.PrivateNicID) + "",
	}

	var resp GetPrivateNICResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePrivateNIC: Update one or more parameter(s) of a specified private NIC.
func (s *API) UpdatePrivateNIC(req *UpdatePrivateNICRequest, opts ...scw.RequestOption) (*PrivateNIC, error) {
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

	if fmt.Sprint(req.PrivateNicID) == "" {
		return nil, errors.New("field PrivateNicID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/private_nics/" + fmt.Sprint(req.PrivateNicID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PrivateNIC

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeletePrivateNIC: Delete a private NIC.
func (s *API) DeletePrivateNIC(req *DeletePrivateNICRequest, opts ...scw.RequestOption) error {
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

	if fmt.Sprint(req.PrivateNicID) == "" {
		return errors.New("field PrivateNicID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/private_nics/" + fmt.Sprint(req.PrivateNicID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// Deprecated: ListBootscripts: List bootscripts.
func (s *API) ListBootscripts(req *ListBootscriptsRequest, opts ...scw.RequestOption) (*ListBootscriptsResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "arch", req.Arch)
	parameter.AddToQuery(query, "title", req.Title)
	parameter.AddToQuery(query, "default", req.Default)
	parameter.AddToQuery(query, "public", req.Public)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/bootscripts",
		Query:  query,
	}

	var resp ListBootscriptsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: GetBootscript: Get details of a bootscript with the specified ID.
func (s *API) GetBootscript(req *GetBootscriptRequest, opts ...scw.RequestOption) (*GetBootscriptResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.BootscriptID) == "" {
		return nil, errors.New("field BootscriptID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/bootscripts/" + fmt.Sprint(req.BootscriptID) + "",
	}

	var resp GetBootscriptResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDashboard:
func (s *API) GetDashboard(req *GetDashboardRequest, opts ...scw.RequestOption) (*GetDashboardResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "project", req.Project)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/dashboard",
		Query:  query,
	}

	var resp GetDashboardResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
