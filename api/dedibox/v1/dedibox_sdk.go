// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package dedibox provides methods and message types of the dedibox v1 API.
package dedibox

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

type AttachFailoverIPToMacAddressRequestMacType string

const (
	AttachFailoverIPToMacAddressRequestMacTypeMacTypeUnknown = AttachFailoverIPToMacAddressRequestMacType("mac_type_unknown")
	AttachFailoverIPToMacAddressRequestMacTypeVmware         = AttachFailoverIPToMacAddressRequestMacType("vmware")
	AttachFailoverIPToMacAddressRequestMacTypeKvm            = AttachFailoverIPToMacAddressRequestMacType("kvm")
	AttachFailoverIPToMacAddressRequestMacTypeXen            = AttachFailoverIPToMacAddressRequestMacType("xen")
)

func (enum AttachFailoverIPToMacAddressRequestMacType) String() string {
	if enum == "" {
		// return default value if empty
		return "mac_type_unknown"
	}
	return string(enum)
}

func (enum AttachFailoverIPToMacAddressRequestMacType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AttachFailoverIPToMacAddressRequestMacType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AttachFailoverIPToMacAddressRequestMacType(AttachFailoverIPToMacAddressRequestMacType(tmp).String())
	return nil
}

type BMCAccessStatus string

const (
	BMCAccessStatusUnknown  = BMCAccessStatus("unknown")
	BMCAccessStatusCreating = BMCAccessStatus("creating")
	BMCAccessStatusCreated  = BMCAccessStatus("created")
	BMCAccessStatusDeleting = BMCAccessStatus("deleting")
)

func (enum BMCAccessStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum BMCAccessStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *BMCAccessStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = BMCAccessStatus(BMCAccessStatus(tmp).String())
	return nil
}

type BackupStatus string

const (
	BackupStatusUnknownBackupStatus = BackupStatus("unknown_backup_status")
	BackupStatusUninitialized       = BackupStatus("uninitialized")
	BackupStatusInactive            = BackupStatus("inactive")
	BackupStatusReady               = BackupStatus("ready")
)

func (enum BackupStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_backup_status"
	}
	return string(enum)
}

func (enum BackupStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *BackupStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = BackupStatus(BackupStatus(tmp).String())
	return nil
}

type FailoverBlockVersion string

const (
	FailoverBlockVersionUnknownVersion = FailoverBlockVersion("unknown_version")
	FailoverBlockVersionIPv4           = FailoverBlockVersion("ipv4")
	FailoverBlockVersionIPv6           = FailoverBlockVersion("ipv6")
)

func (enum FailoverBlockVersion) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_version"
	}
	return string(enum)
}

func (enum FailoverBlockVersion) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FailoverBlockVersion) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FailoverBlockVersion(FailoverBlockVersion(tmp).String())
	return nil
}

type FailoverIPInterfaceType string

const (
	FailoverIPInterfaceTypeUnknown = FailoverIPInterfaceType("unknown")
	FailoverIPInterfaceTypeNormal  = FailoverIPInterfaceType("normal")
	FailoverIPInterfaceTypeIpmi    = FailoverIPInterfaceType("ipmi")
	FailoverIPInterfaceTypeVirtual = FailoverIPInterfaceType("virtual")
)

func (enum FailoverIPInterfaceType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum FailoverIPInterfaceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FailoverIPInterfaceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FailoverIPInterfaceType(FailoverIPInterfaceType(tmp).String())
	return nil
}

type FailoverIPStatus string

const (
	FailoverIPStatusUnknownStatus = FailoverIPStatus("unknown_status")
	FailoverIPStatusReady         = FailoverIPStatus("ready")
	FailoverIPStatusBusy          = FailoverIPStatus("busy")
	FailoverIPStatusLocked        = FailoverIPStatus("locked")
)

func (enum FailoverIPStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum FailoverIPStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FailoverIPStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FailoverIPStatus(FailoverIPStatus(tmp).String())
	return nil
}

type FailoverIPVersion string

const (
	FailoverIPVersionUnknownVersion = FailoverIPVersion("unknown_version")
	FailoverIPVersionIPv4           = FailoverIPVersion("ipv4")
	FailoverIPVersionIPv6           = FailoverIPVersion("ipv6")
)

func (enum FailoverIPVersion) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_version"
	}
	return string(enum)
}

func (enum FailoverIPVersion) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FailoverIPVersion) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FailoverIPVersion(FailoverIPVersion(tmp).String())
	return nil
}

type IPSemantic string

const (
	IPSemanticUnknown   = IPSemantic("unknown")
	IPSemanticProxad    = IPSemantic("proxad")
	IPSemanticExt       = IPSemantic("ext")
	IPSemanticPublic    = IPSemantic("public")
	IPSemanticPrivate   = IPSemantic("private")
	IPSemanticIpmi      = IPSemantic("ipmi")
	IPSemanticAdm       = IPSemantic("adm")
	IPSemanticRedirect  = IPSemantic("redirect")
	IPSemanticMigration = IPSemantic("migration")
)

func (enum IPSemantic) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum IPSemantic) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *IPSemantic) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = IPSemantic(IPSemantic(tmp).String())
	return nil
}

type IPStatus string

const (
	IPStatusUnknownStatus = IPStatus("unknown_status")
	IPStatusReady         = IPStatus("ready")
	IPStatusBusy          = IPStatus("busy")
	IPStatusLocked        = IPStatus("locked")
)

func (enum IPStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum IPStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *IPStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = IPStatus(IPStatus(tmp).String())
	return nil
}

type IPVersion string

const (
	IPVersionIPv4 = IPVersion("ipv4")
	IPVersionIPv6 = IPVersion("ipv6")
)

func (enum IPVersion) String() string {
	if enum == "" {
		// return default value if empty
		return "ipv4"
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

type ListFailoverIPsRequestOrderBy string

const (
	ListFailoverIPsRequestOrderByIPAsc  = ListFailoverIPsRequestOrderBy("ip_asc")
	ListFailoverIPsRequestOrderByIPDesc = ListFailoverIPsRequestOrderBy("ip_desc")
)

func (enum ListFailoverIPsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "ip_asc"
	}
	return string(enum)
}

func (enum ListFailoverIPsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListFailoverIPsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListFailoverIPsRequestOrderBy(ListFailoverIPsRequestOrderBy(tmp).String())
	return nil
}

type ListOSRequestOrderBy string

const (
	ListOSRequestOrderByCreatedAtAsc   = ListOSRequestOrderBy("created_at_asc")
	ListOSRequestOrderByCreatedAtDesc  = ListOSRequestOrderBy("created_at_desc")
	ListOSRequestOrderByReleasedAtAsc  = ListOSRequestOrderBy("released_at_asc")
	ListOSRequestOrderByReleasedAtDesc = ListOSRequestOrderBy("released_at_desc")
)

func (enum ListOSRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListOSRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListOSRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListOSRequestOrderBy(ListOSRequestOrderBy(tmp).String())
	return nil
}

type ListOffersRequestOrderBy string

const (
	ListOffersRequestOrderByCreatedAtAsc  = ListOffersRequestOrderBy("created_at_asc")
	ListOffersRequestOrderByCreatedAtDesc = ListOffersRequestOrderBy("created_at_desc")
	ListOffersRequestOrderByNameAsc       = ListOffersRequestOrderBy("name_asc")
	ListOffersRequestOrderByNameDesc      = ListOffersRequestOrderBy("name_desc")
	ListOffersRequestOrderByPriceAsc      = ListOffersRequestOrderBy("price_asc")
	ListOffersRequestOrderByPriceDesc     = ListOffersRequestOrderBy("price_desc")
)

func (enum ListOffersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListOffersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListOffersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListOffersRequestOrderBy(ListOffersRequestOrderBy(tmp).String())
	return nil
}

type ListServerDisksRequestOrderBy string

const (
	ListServerDisksRequestOrderByCreatedAtAsc  = ListServerDisksRequestOrderBy("created_at_asc")
	ListServerDisksRequestOrderByCreatedAtDesc = ListServerDisksRequestOrderBy("created_at_desc")
)

func (enum ListServerDisksRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListServerDisksRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListServerDisksRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListServerDisksRequestOrderBy(ListServerDisksRequestOrderBy(tmp).String())
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

type ListServicesRequestOrderBy string

const (
	ListServicesRequestOrderByCreatedAtAsc  = ListServicesRequestOrderBy("created_at_asc")
	ListServicesRequestOrderByCreatedAtDesc = ListServicesRequestOrderBy("created_at_desc")
)

func (enum ListServicesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListServicesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListServicesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListServicesRequestOrderBy(ListServicesRequestOrderBy(tmp).String())
	return nil
}

type MemoryType string

const (
	MemoryTypeDdr2 = MemoryType("ddr2")
	MemoryTypeDdr3 = MemoryType("ddr3")
	MemoryTypeDdr4 = MemoryType("ddr4")
)

func (enum MemoryType) String() string {
	if enum == "" {
		// return default value if empty
		return "ddr2"
	}
	return string(enum)
}

func (enum MemoryType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MemoryType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MemoryType(MemoryType(tmp).String())
	return nil
}

type NetworkInterfaceInterfaceType string

const (
	NetworkInterfaceInterfaceTypeUnknown = NetworkInterfaceInterfaceType("unknown")
	NetworkInterfaceInterfaceTypeNormal  = NetworkInterfaceInterfaceType("normal")
	NetworkInterfaceInterfaceTypeIpmi    = NetworkInterfaceInterfaceType("ipmi")
	NetworkInterfaceInterfaceTypeVirtual = NetworkInterfaceInterfaceType("virtual")
)

func (enum NetworkInterfaceInterfaceType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum NetworkInterfaceInterfaceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NetworkInterfaceInterfaceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NetworkInterfaceInterfaceType(NetworkInterfaceInterfaceType(tmp).String())
	return nil
}

type OSArch string

const (
	OSArchUnknownArch = OSArch("unknown_arch")
	OSArchAmd64       = OSArch("amd64")
	OSArchX86         = OSArch("x86")
	OSArchArm         = OSArch("arm")
	OSArchArm64       = OSArch("arm64")
)

func (enum OSArch) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_arch"
	}
	return string(enum)
}

func (enum OSArch) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OSArch) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OSArch(OSArch(tmp).String())
	return nil
}

type OSType string

const (
	OSTypeUnknownType = OSType("unknown_type")
	OSTypeServer      = OSType("server")
	OSTypeVirtu       = OSType("virtu")
	OSTypePanel       = OSType("panel")
	OSTypeDesktop     = OSType("desktop")
	OSTypeCustom      = OSType("custom")
	OSTypeRescue      = OSType("rescue")
)

func (enum OSType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum OSType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OSType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OSType(OSType(tmp).String())
	return nil
}

type OfferAntiDosInfoType string

const (
	OfferAntiDosInfoTypeMinimal    = OfferAntiDosInfoType("minimal")
	OfferAntiDosInfoTypePreventive = OfferAntiDosInfoType("preventive")
	OfferAntiDosInfoTypeCurative   = OfferAntiDosInfoType("curative")
)

func (enum OfferAntiDosInfoType) String() string {
	if enum == "" {
		// return default value if empty
		return "minimal"
	}
	return string(enum)
}

func (enum OfferAntiDosInfoType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OfferAntiDosInfoType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OfferAntiDosInfoType(OfferAntiDosInfoType(tmp).String())
	return nil
}

type OfferCatalog string

const (
	OfferCatalogAll      = OfferCatalog("all")
	OfferCatalogDefault  = OfferCatalog("default")
	OfferCatalogBeta     = OfferCatalog("beta")
	OfferCatalogReseller = OfferCatalog("reseller")
	OfferCatalogPremium  = OfferCatalog("premium")
	OfferCatalogVolume   = OfferCatalog("volume")
	OfferCatalogAdmin    = OfferCatalog("admin")
	OfferCatalogInactive = OfferCatalog("inactive")
)

func (enum OfferCatalog) String() string {
	if enum == "" {
		// return default value if empty
		return "all"
	}
	return string(enum)
}

func (enum OfferCatalog) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OfferCatalog) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OfferCatalog(OfferCatalog(tmp).String())
	return nil
}

type OfferPaymentFrequency string

const (
	OfferPaymentFrequencyMonthly = OfferPaymentFrequency("monthly")
	OfferPaymentFrequencyOneshot = OfferPaymentFrequency("oneshot")
)

func (enum OfferPaymentFrequency) String() string {
	if enum == "" {
		// return default value if empty
		return "monthly"
	}
	return string(enum)
}

func (enum OfferPaymentFrequency) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OfferPaymentFrequency) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OfferPaymentFrequency(OfferPaymentFrequency(tmp).String())
	return nil
}

type OfferSANInfoType string

const (
	OfferSANInfoTypeHdd = OfferSANInfoType("hdd")
	OfferSANInfoTypeSSD = OfferSANInfoType("ssd")
)

func (enum OfferSANInfoType) String() string {
	if enum == "" {
		// return default value if empty
		return "hdd"
	}
	return string(enum)
}

func (enum OfferSANInfoType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OfferSANInfoType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OfferSANInfoType(OfferSANInfoType(tmp).String())
	return nil
}

type OfferServerInfoStock string

const (
	OfferServerInfoStockEmpty     = OfferServerInfoStock("empty")
	OfferServerInfoStockLow       = OfferServerInfoStock("low")
	OfferServerInfoStockAvailable = OfferServerInfoStock("available")
)

func (enum OfferServerInfoStock) String() string {
	if enum == "" {
		// return default value if empty
		return "empty"
	}
	return string(enum)
}

func (enum OfferServerInfoStock) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OfferServerInfoStock) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OfferServerInfoStock(OfferServerInfoStock(tmp).String())
	return nil
}

type PartitionFileSystem string

const (
	PartitionFileSystemUnknown = PartitionFileSystem("unknown")
	PartitionFileSystemEfi     = PartitionFileSystem("efi")
	PartitionFileSystemSwap    = PartitionFileSystem("swap")
	PartitionFileSystemExt4    = PartitionFileSystem("ext4")
	PartitionFileSystemExt3    = PartitionFileSystem("ext3")
	PartitionFileSystemExt2    = PartitionFileSystem("ext2")
	PartitionFileSystemXfs     = PartitionFileSystem("xfs")
	PartitionFileSystemNtfs    = PartitionFileSystem("ntfs")
	PartitionFileSystemFat32   = PartitionFileSystem("fat32")
	PartitionFileSystemUfs     = PartitionFileSystem("ufs")
)

func (enum PartitionFileSystem) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum PartitionFileSystem) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PartitionFileSystem) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PartitionFileSystem(PartitionFileSystem(tmp).String())
	return nil
}

type PartitionType string

const (
	PartitionTypePrimary  = PartitionType("primary")
	PartitionTypeExtended = PartitionType("extended")
	PartitionTypeLogical  = PartitionType("logical")
)

func (enum PartitionType) String() string {
	if enum == "" {
		// return default value if empty
		return "primary"
	}
	return string(enum)
}

func (enum PartitionType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PartitionType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PartitionType(PartitionType(tmp).String())
	return nil
}

type RaidArrayRaidLevel string

const (
	RaidArrayRaidLevelNoRaid = RaidArrayRaidLevel("no_raid")
	RaidArrayRaidLevelRaid0  = RaidArrayRaidLevel("raid0")
	RaidArrayRaidLevelRaid1  = RaidArrayRaidLevel("raid1")
	RaidArrayRaidLevelRaid5  = RaidArrayRaidLevel("raid5")
	RaidArrayRaidLevelRaid6  = RaidArrayRaidLevel("raid6")
	RaidArrayRaidLevelRaid10 = RaidArrayRaidLevel("raid10")
)

func (enum RaidArrayRaidLevel) String() string {
	if enum == "" {
		// return default value if empty
		return "no_raid"
	}
	return string(enum)
}

func (enum RaidArrayRaidLevel) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RaidArrayRaidLevel) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RaidArrayRaidLevel(RaidArrayRaidLevel(tmp).String())
	return nil
}

type RescueProtocol string

const (
	RescueProtocolVnc = RescueProtocol("vnc")
	RescueProtocolSSH = RescueProtocol("ssh")
)

func (enum RescueProtocol) String() string {
	if enum == "" {
		// return default value if empty
		return "vnc"
	}
	return string(enum)
}

func (enum RescueProtocol) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RescueProtocol) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RescueProtocol(RescueProtocol(tmp).String())
	return nil
}

type ServerDiskType string

const (
	ServerDiskTypeSata = ServerDiskType("sata")
	ServerDiskTypeSSD  = ServerDiskType("ssd")
	ServerDiskTypeSas  = ServerDiskType("sas")
	ServerDiskTypeSshd = ServerDiskType("sshd")
	ServerDiskTypeUsb  = ServerDiskType("usb")
	ServerDiskTypeNvme = ServerDiskType("nvme")
)

func (enum ServerDiskType) String() string {
	if enum == "" {
		// return default value if empty
		return "sata"
	}
	return string(enum)
}

func (enum ServerDiskType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerDiskType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerDiskType(ServerDiskType(tmp).String())
	return nil
}

type ServerInstallStatus string

const (
	ServerInstallStatusUnknown               = ServerInstallStatus("unknown")
	ServerInstallStatusBooting               = ServerInstallStatus("booting")
	ServerInstallStatusSettingUpRaid         = ServerInstallStatus("setting_up_raid")
	ServerInstallStatusPartitioning          = ServerInstallStatus("partitioning")
	ServerInstallStatusFormatting            = ServerInstallStatus("formatting")
	ServerInstallStatusInstalling            = ServerInstallStatus("installing")
	ServerInstallStatusConfiguring           = ServerInstallStatus("configuring")
	ServerInstallStatusConfiguringBootloader = ServerInstallStatus("configuring_bootloader")
	ServerInstallStatusRebooting             = ServerInstallStatus("rebooting")
	ServerInstallStatusInstalled             = ServerInstallStatus("installed")
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

type ServerStatus string

const (
	ServerStatusUnknown    = ServerStatus("unknown")
	ServerStatusDelivering = ServerStatus("delivering")
	ServerStatusInstalling = ServerStatus("installing")
	ServerStatusReady      = ServerStatus("ready")
	ServerStatusStopped    = ServerStatus("stopped")
	ServerStatusError      = ServerStatus("error")
	ServerStatusLocked     = ServerStatus("locked")
	ServerStatusRescue     = ServerStatus("rescue")
	ServerStatusBusy       = ServerStatus("busy")
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

type ServiceLevelLevel string

const (
	ServiceLevelLevelUnknown  = ServiceLevelLevel("unknown")
	ServiceLevelLevelBasic    = ServiceLevelLevel("basic")
	ServiceLevelLevelBusiness = ServiceLevelLevel("business")
)

func (enum ServiceLevelLevel) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum ServiceLevelLevel) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServiceLevelLevel) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServiceLevelLevel(ServiceLevelLevel(tmp).String())
	return nil
}

type ServiceProvisioningStatus string

const (
	ServiceProvisioningStatusUnknown    = ServiceProvisioningStatus("unknown")
	ServiceProvisioningStatusDelivering = ServiceProvisioningStatus("delivering")
	ServiceProvisioningStatusReady      = ServiceProvisioningStatus("ready")
	ServiceProvisioningStatusError      = ServiceProvisioningStatus("error")
	ServiceProvisioningStatusExpiring   = ServiceProvisioningStatus("expiring")
	ServiceProvisioningStatusExpired    = ServiceProvisioningStatus("expired")
)

func (enum ServiceProvisioningStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum ServiceProvisioningStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServiceProvisioningStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServiceProvisioningStatus(ServiceProvisioningStatus(tmp).String())
	return nil
}

type ServiceType string

const (
	ServiceTypeUnknownType = ServiceType("unknown_type")
	ServiceTypeService     = ServiceType("service")
	ServiceTypeOrder       = ServiceType("order")
)

func (enum ServiceType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum ServiceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServiceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServiceType(ServiceType(tmp).String())
	return nil
}

// CPU:
type CPU struct {
	// Name: Name of CPU.
	Name string `json:"name"`
	// CoreCount: Number of cores of the CPU.
	CoreCount uint32 `json:"core_count"`
	// ThreadCount: Number of threads of the CPU.
	ThreadCount uint32 `json:"thread_count"`
	// Frequency: Frequency of the CPU.
	Frequency uint32 `json:"frequency"`
}

// Disk:
type Disk struct {
	// Capacity: Capacity of the disk.
	Capacity scw.Size `json:"capacity"`
	// Type: Type of the disk.
	Type ServerDiskType `json:"type"`
}

// Memory:
type Memory struct {
	// Capacity: Capacity of the memory.
	Capacity scw.Size `json:"capacity"`
	// Type: Type of the memory.
	Type MemoryType `json:"type"`
	// Frequency: Frequency of the memory.
	Frequency uint32 `json:"frequency"`
	// IsEcc: True if the memory is an error-correcting code memory.
	IsEcc bool `json:"is_ecc"`
}

// OfferAntiDosInfo:
type OfferAntiDosInfo struct {
	// Type:
	Type OfferAntiDosInfoType `json:"type"`
}

// OfferBackupInfo:
type OfferBackupInfo struct {
	// Size:
	Size scw.Size `json:"size"`
}

// OfferBandwidthInfo:
type OfferBandwidthInfo struct {
	// Speed:
	Speed uint32 `json:"speed"`
}

// OfferLicenseInfo:
type OfferLicenseInfo struct {
	// BoundToIP:
	BoundToIP bool `json:"bound_to_ip"`
}

// OfferRPNInfo:
type OfferRPNInfo struct {
	// Speed:
	Speed uint32 `json:"speed"`
}

// OfferSANInfo:
type OfferSANInfo struct {
	// Size: SAN size (in bytes).
	Size uint64 `json:"size"`
	// Ha: High availabilty offer.
	Ha bool `json:"ha"`
	// DeviceType: Type of SAN device (hdd / ssd).
	DeviceType OfferSANInfoType `json:"device_type"`
}

// OfferStorageInfo:
type OfferStorageInfo struct {
	// MaxQuota:
	MaxQuota uint32 `json:"max_quota"`
	// Size:
	Size scw.Size `json:"size"`
}

// PersistentMemory:
type PersistentMemory struct {
	// Capacity: Capacity of the persistent memory.
	Capacity scw.Size `json:"capacity"`
	// Frequency: Frequency of the persistent memory.
	Frequency uint32 `json:"frequency"`
	// Model: Model of the persistent memory.
	Model string `json:"model"`
}

// RaidController:
type RaidController struct {
	// Model: Model of the RAID controller.
	Model string `json:"model"`
	// RaidLevel: RAID level of the RAID controller.
	RaidLevel []string `json:"raid_level"`
}

// Offer:
type Offer struct {
	// ID: ID of the offer.
	ID uint64 `json:"id"`
	// Name: Name of the offer.
	Name string `json:"name"`
	// Catalog: Catalog of the offer.
	Catalog OfferCatalog `json:"catalog"`
	// PaymentFrequency: Payment frequency of the offer.
	PaymentFrequency OfferPaymentFrequency `json:"payment_frequency"`
	// Pricing: Price of the offer.
	Pricing *scw.Money `json:"pricing,omitempty"`
	// ServerInfo: Server info if it is a server offer.
	ServerInfo *OfferServerInfo `json:"server_info,omitempty"`
	// ServiceLevelInfo: Service level info if it is a service level offer.
	ServiceLevelInfo *OfferServiceLevelInfo `json:"service_level_info,omitempty"`
	// RpnInfo: RPN info if it is a RPN offer.
	RpnInfo *OfferRPNInfo `json:"rpn_info,omitempty"`
	// SanInfo: SAN info if it is a SAN offer.
	SanInfo *OfferSANInfo `json:"san_info,omitempty"`
	// AntidosInfo: AntiDOS info if it is a antiDOS offer.
	AntidosInfo *OfferAntiDosInfo `json:"antidos_info,omitempty"`
	// BackupInfo: Backup info if it is a backup offer.
	BackupInfo *OfferBackupInfo `json:"backup_info,omitempty"`
	// UsbStorageInfo: USB storage info if it is a USB storage offer.
	UsbStorageInfo *OfferStorageInfo `json:"usb_storage_info,omitempty"`
	// StorageInfo: Storage info if it is a storage offer.
	StorageInfo *OfferStorageInfo `json:"storage_info,omitempty"`
	// LicenseInfo: License info if it is a license offer.
	LicenseInfo *OfferLicenseInfo `json:"license_info,omitempty"`
	// FailoverIPInfo: Failover IP info if it is a failover IP offer.
	FailoverIPInfo *OfferFailoverIPInfo `json:"failover_ip_info,omitempty"`
	// FailoverBlockInfo: Failover block info if it is a failover block offer.
	FailoverBlockInfo *OfferFailoverBlockInfo `json:"failover_block_info,omitempty"`
	// BandwidthInfo: Bandwidth info if it is a bandwidth offer.
	BandwidthInfo *OfferBandwidthInfo `json:"bandwidth_info,omitempty"`
}

// OfferFailoverBlockInfo:
type OfferFailoverBlockInfo struct {
	// OnetimeFees:
	OnetimeFees *Offer `json:"onetime_fees"`
}

// OfferFailoverIPInfo:
type OfferFailoverIPInfo struct {
	// OnetimeFees:
	OnetimeFees *Offer `json:"onetime_fees"`
}

// OfferServiceLevelInfo:
type OfferServiceLevelInfo struct {
	// SupportTicket:
	SupportTicket bool `json:"support_ticket"`
	// SupportPhone:
	SupportPhone bool `json:"support_phone"`
	// SalesSupport:
	SalesSupport bool `json:"sales_support"`
	// Git:
	Git string `json:"git"`
	// SLA:
	SLA float32 `json:"sla"`
	// PrioritySupport:
	PrioritySupport bool `json:"priority_support"`
	// HighRpnBandwidth:
	HighRpnBandwidth bool `json:"high_rpn_bandwidth"`
	// Customization:
	Customization bool `json:"customization"`
	// Antidos:
	Antidos bool `json:"antidos"`
	// ExtraFailoverQuota:
	ExtraFailoverQuota uint32 `json:"extra_failover_quota"`
	// AvailableOptions:
	AvailableOptions []*Offer `json:"available_options"`
}

// IP:
type IP struct {
	// IPID: ID of the IP.
	IPID string `json:"ip_id"`
	// Address: Address of the IP.
	Address net.IP `json:"address"`
	// Reverse: Reverse IP value.
	Reverse string `json:"reverse"`
	// Version: Version of IP (v4 or v6).
	Version IPVersion `json:"version"`
	// Cidr: Classless InterDomain Routing notation of the IP.
	Cidr uint32 `json:"cidr"`
	// Netmask: Network mask of IP.
	Netmask net.IP `json:"netmask"`
	// Semantic: Semantic of IP.
	Semantic IPSemantic `json:"semantic"`
	// Gateway: Gateway of IP.
	Gateway net.IP `json:"gateway"`
	// Status: Status of the IP.
	Status IPStatus `json:"status"`
}

// OfferServerInfo:
type OfferServerInfo struct {
	// Bandwidth:
	Bandwidth uint64 `json:"bandwidth"`
	// Stock:
	Stock OfferServerInfoStock `json:"stock"`
	// CommercialRange:
	CommercialRange string `json:"commercial_range"`
	// Disks:
	Disks []*Disk `json:"disks"`
	// CPUs:
	CPUs []*CPU `json:"cpus"`
	// Memories:
	Memories []*Memory `json:"memories"`
	// PersistentMemories:
	PersistentMemories []*PersistentMemory `json:"persistent_memories"`
	// RaidControllers:
	RaidControllers []*RaidController `json:"raid_controllers"`
	// AvailableOptions:
	AvailableOptions []*Offer `json:"available_options"`
	// RpnVersion:
	RpnVersion *uint32 `json:"rpn_version,omitempty"`
	// Connectivity:
	Connectivity uint64 `json:"connectivity"`
	// OnetimeFees:
	OnetimeFees *Offer `json:"onetime_fees"`
	// StockByDatacenter:
	StockByDatacenter map[string]OfferServerInfoStock `json:"stock_by_datacenter"`
}

// NetworkInterface:
type NetworkInterface struct {
	// CardID: Card ID of the network interface.
	CardID uint64 `json:"card_id"`
	// DeviceID: Device ID of the network interface.
	DeviceID uint64 `json:"device_id"`
	// Mac: MAC address of the network interface.
	Mac string `json:"mac"`
	// Type: Network interface type.
	Type NetworkInterfaceInterfaceType `json:"type"`
	// IPs: IPs of the network interface.
	IPs []*IP `json:"ips"`
}

// OS:
type OS struct {
	// ID: ID of the OS.
	ID uint64 `json:"id"`
	// Name: Name of the OS.
	Name string `json:"name"`
	// Type: Type of the OS.
	Type OSType `json:"type"`
	// Version: Version of the OS.
	Version string `json:"version"`
	// Arch: Architecture of the OS.
	Arch OSArch `json:"arch"`
	// AllowCustomPartitioning: True if the OS allow custom partitioning.
	AllowCustomPartitioning bool `json:"allow_custom_partitioning"`
	// AllowSSHKeys: True if the OS allow SSH Keys.
	AllowSSHKeys bool `json:"allow_ssh_keys"`
	// RequiresUser: True if the OS requires user.
	RequiresUser bool `json:"requires_user"`
	// RequiresAdminPassword: True if the OS requires admin password.
	RequiresAdminPassword bool `json:"requires_admin_password"`
	// RequiresPanelPassword: True if the OS requires panel password.
	RequiresPanelPassword bool `json:"requires_panel_password"`
	// AllowedFilesystems: True if the OS allow file systems.
	AllowedFilesystems []PartitionFileSystem `json:"allowed_filesystems"`
	// RequiresLicense: True if the OS requires license.
	RequiresLicense bool `json:"requires_license"`
	// LicenseOffers: License offers available with the OS.
	LicenseOffers []*Offer `json:"license_offers"`
	// MaxPartitions: Maximum number of partitions which can be created.
	MaxPartitions *uint32 `json:"max_partitions,omitempty"`
	// DisplayName: Display name of the OS.
	DisplayName string `json:"display_name"`
	// PasswordRegex: Regex used to validate the installation passwords.
	PasswordRegex string `json:"password_regex"`
	// PanelPasswordRegex: Regex used to validate the panel installation password.
	PanelPasswordRegex *string `json:"panel_password_regex,omitempty"`
	// RequiresValidHostname: If both requires_valid_hostname & hostname_regex are set, it means that at least one of the criterias must be valid.
	RequiresValidHostname *bool `json:"requires_valid_hostname,omitempty"`
	// HostnameRegex: If both requires_valid_hostname & hostname_regex are set, it means that at least one of the criterias must be valid.
	HostnameRegex *string `json:"hostname_regex,omitempty"`
	// HostnameMaxLength: Hostname max length.
	HostnameMaxLength uint32 `json:"hostname_max_length"`
	// ReleasedAt: OS release date.
	ReleasedAt *time.Time `json:"released_at,omitempty"`
}

// ServerLocation:
type ServerLocation struct {
	// Rack:
	Rack string `json:"rack"`
	// Room:
	Room string `json:"room"`
	// DatacenterName:
	DatacenterName string `json:"datacenter_name"`
}

// ServerOption:
type ServerOption struct {
	// Offer:
	Offer *Offer `json:"offer"`
	// CreatedAt:
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt:
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// ExpiredAt:
	ExpiredAt *time.Time `json:"expired_at,omitempty"`
	// Options:
	Options []*ServerOption `json:"options"`
}

// ServiceLevel:
type ServiceLevel struct {
	// OfferID: Offer ID of service level.
	OfferID uint32 `json:"offer_id"`
	// Level: Level type of service level.
	Level ServiceLevelLevel `json:"level"`
}

// Server:
type Server struct {
	// ID: ID of the server.
	ID uint64 `json:"id"`
	// OrganizationID: Organization ID the server is attached to.
	OrganizationID string `json:"organization_id"`
	// ProjectID: Project ID the server is attached to.
	ProjectID string `json:"project_id"`
	// Hostname: Hostname of the server.
	Hostname string `json:"hostname"`
	// RebootedAt: Date of last reboot of the server.
	RebootedAt *time.Time `json:"rebooted_at,omitempty"`
	// CreatedAt: Date of creation of the server.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Date of last modification of the server.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// ExpiredAt: Date of release of the server.
	ExpiredAt *time.Time `json:"expired_at,omitempty"`
	// Offer: Offer of the server.
	Offer *Offer `json:"offer"`
	// Status: Status of the server.
	Status ServerStatus `json:"status"`
	// Location: Location of the server.
	Location *ServerLocation `json:"location"`
	// AbuseContact: Abuse contact of the server.
	AbuseContact string `json:"abuse_contact"`
	// Os: OS installed on the server.
	Os *OS `json:"os"`
	// Interfaces: Network interfaces of the server.
	Interfaces []*NetworkInterface `json:"interfaces"`
	// Zone: The zone in which is the server.
	Zone scw.Zone `json:"zone"`
	// Options: Options subscribe on the server.
	Options []*ServerOption `json:"options"`
	// Level: Service level of the server.
	Level *ServiceLevel `json:"level"`
	// HasBmc: Boolean if the server has a BMC.
	HasBmc bool `json:"has_bmc"`
	// RescueOs: Rescue OS of the server.
	RescueOs *OS `json:"rescue_os"`
	// Tags: Array of customs tags attached to the server.
	Tags []string `json:"tags"`
	// IsOutsourced: Whether the server is outsourced or not.
	IsOutsourced bool `json:"is_outsourced"`
	// IPv6Slaac: Whether or not you can enable/disable the IPv6.
	IPv6Slaac bool `json:"ipv6_slaac"`
	// Qinq: Whether the server is compatible with QinQ.
	Qinq bool `json:"qinq"`
	// IsRpnv2Member: Whether or not the server is already part of an rpnv2 group.
	IsRpnv2Member bool `json:"is_rpnv2_member"`
}

// FailoverBlock:
type FailoverBlock struct {
	// ID: ID of the failover block.
	ID uint64 `json:"id"`
	// Address: IP of the failover block.
	Address net.IP `json:"address"`
	// Nameservers: Name servers.
	Nameservers []string `json:"nameservers"`
	// IPVersion: IP version of the failover block.
	IPVersion FailoverBlockVersion `json:"ip_version"`
	// Cidr: Classless InterDomain Routing notation of the failover block.
	Cidr uint32 `json:"cidr"`
	// Netmask: Netmask of the failover block.
	Netmask net.IP `json:"netmask"`
	// GatewayIP: Gateway IP of the failover block.
	GatewayIP net.IP `json:"gateway_ip"`
}

// ServerDisk:
type ServerDisk struct {
	// ID:
	ID uint64 `json:"id"`
	// Connector:
	Connector string `json:"connector"`
	// Type:
	Type ServerDiskType `json:"type"`
	// Capacity:
	Capacity scw.Size `json:"capacity"`
	// IsAddon:
	IsAddon bool `json:"is_addon"`
}

// Service:
type Service struct {
	// ID: ID of the service.
	ID uint64 `json:"id"`
	// ResourceID: Resource ID of the service.
	ResourceID *uint64 `json:"resource_id,omitempty"`
	// ProvisioningStatus: Provisioning status of the service.
	ProvisioningStatus ServiceProvisioningStatus `json:"provisioning_status"`
	// Offer: Offer of the service.
	Offer *Offer `json:"offer"`
	// CreatedAt: Creation date of the service.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// DeliveredAt: Delivery date of the service.
	DeliveredAt *time.Time `json:"delivered_at,omitempty"`
	// TerminatedAt: Terminatation date of the service.
	TerminatedAt *time.Time `json:"terminated_at,omitempty"`
	// ExpiresAt: Expiration date of the service.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// Type: Service type, either order or service.
	Type ServiceType `json:"type"`
}

// InstallPartition:
type InstallPartition struct {
	// FileSystem: File system of the installation partition.
	FileSystem PartitionFileSystem `json:"file_system"`
	// MountPoint: Mount point of the installation partition.
	MountPoint *string `json:"mount_point,omitempty"`
	// RaidLevel: RAID level of the installation partition.
	RaidLevel RaidArrayRaidLevel `json:"raid_level"`
	// Capacity: Capacity of the installation partition.
	Capacity scw.Size `json:"capacity"`
	// Connectors: Connectors of the installation partition.
	Connectors []string `json:"connectors"`
}

// FailoverIP:
type FailoverIP struct {
	// ID: ID of the failover IP.
	ID uint64 `json:"id"`
	// Address: IP of the failover IP.
	Address net.IP `json:"address"`
	// Reverse: Reverse IP value.
	Reverse string `json:"reverse"`
	// IPVersion: IP version of the failover IP.
	IPVersion FailoverIPVersion `json:"ip_version"`
	// Cidr: Classless InterDomain Routing notation of the failover IP.
	Cidr uint32 `json:"cidr"`
	// Netmask: Netmask of the failover IP.
	Netmask net.IP `json:"netmask"`
	// GatewayIP: Gateway IP of the failover IP.
	GatewayIP net.IP `json:"gateway_ip"`
	// Mac: MAC address of the IP failover.
	Mac *string `json:"mac,omitempty"`
	// ServerID: Server ID linked to the IP failover.
	ServerID *uint64 `json:"server_id,omitempty"`
	// Status: Status of the IP failover.
	Status FailoverIPStatus `json:"status"`
	// Block: Block of the IP failover.
	Block *FailoverBlock `json:"block"`
	// Type: The interface type.
	Type FailoverIPInterfaceType `json:"type"`
	// ServerZone: The server zone (if assigned).
	ServerZone *string `json:"server_zone,omitempty"`
}

// ServerEvent:
type ServerEvent struct {
	// EventID: ID of the event.
	EventID uint64 `json:"event_id"`
	// Description: Descriptiion of the event.
	Description string `json:"description"`
	// Date: Date of the event.
	Date *time.Time `json:"date,omitempty"`
}

// ServerSummary:
type ServerSummary struct {
	// ID: ID of the server.
	ID uint64 `json:"id"`
	// DatacenterName: Datacenter of the server.
	DatacenterName string `json:"datacenter_name"`
	// OrganizationID: Organization ID the server is attached to.
	OrganizationID string `json:"organization_id"`
	// ProjectID: Project ID the server is attached to.
	ProjectID string `json:"project_id"`
	// Hostname: Hostname of the server.
	Hostname string `json:"hostname"`
	// CreatedAt: Date of creation of the server.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Date of last modification of the server.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// ExpiredAt: Date of release of the server.
	ExpiredAt *time.Time `json:"expired_at,omitempty"`
	// OfferID: Offer ID of the server.
	OfferID uint64 `json:"offer_id"`
	// OfferName: Offer name of the server.
	OfferName string `json:"offer_name"`
	// Status: Status of the server.
	Status ServerStatus `json:"status"`
	// OsID: OS ID installed on server.
	OsID *uint64 `json:"os_id,omitempty"`
	// Interfaces: Network interfaces of the server.
	Interfaces []*NetworkInterface `json:"interfaces"`
	// Zone: The zone in which is the server.
	Zone scw.Zone `json:"zone"`
	// Level: Service level of the server.
	Level *ServiceLevel `json:"level"`
	// IsOutsourced: Whether the server is outsourced or not.
	IsOutsourced bool `json:"is_outsourced"`
	// Qinq: Whether the server is compatible with QinQ.
	Qinq bool `json:"qinq"`
	// RpnVersion: Supported RPN version.
	RpnVersion *uint32 `json:"rpn_version,omitempty"`
}

// RaidArray:
type RaidArray struct {
	// RaidLevel: The RAID level.
	RaidLevel RaidArrayRaidLevel `json:"raid_level"`
	// Disks: Disks on the RAID controller.
	Disks []*ServerDisk `json:"disks"`
}

// Partition:
type Partition struct {
	// Type: Type of the partition.
	Type PartitionType `json:"type"`
	// FileSystem: File system of the partition.
	FileSystem PartitionFileSystem `json:"file_system"`
	// MountPoint: Mount point of the partition.
	MountPoint *string `json:"mount_point,omitempty"`
	// RaidLevel: Raid level of the partition.
	RaidLevel RaidArrayRaidLevel `json:"raid_level"`
	// Capacity: Capacity of the partition.
	Capacity scw.Size `json:"capacity"`
	// Connectors: Connectors of the partition.
	Connectors []string `json:"connectors"`
}

// UpdatableRaidArray:
type UpdatableRaidArray struct {
	// RaidLevel: The RAID level.
	RaidLevel RaidArrayRaidLevel `json:"raid_level"`
	// DiskIDs: The list of Disk ID of the updatable RAID.
	DiskIDs []uint64 `json:"disk_ids"`
}

// AttachFailoverIPToMacAddressRequest:
type AttachFailoverIPToMacAddressRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// IPID: ID of the failover IP.
	IPID uint64 `json:"-"`
	// Type: A mac type.
	Type AttachFailoverIPToMacAddressRequestMacType `json:"type"`
	// Mac: A valid mac address (existing or not).
	Mac *string `json:"mac,omitempty"`
}

// AttachFailoverIPsRequest:
type AttachFailoverIPsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server.
	ServerID uint64 `json:"server_id"`
	// FipsIDs: List of ID of failovers IP to attach.
	FipsIDs []uint64 `json:"fips_ids"`
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
	// Status: Status of the connection.
	Status BMCAccessStatus `json:"status"`
}

// Backup:
type Backup struct {
	// ID: ID of the backup.
	ID uint64 `json:"id"`
	// Login: Login of the backup.
	Login string `json:"login"`
	// Server: Server of the backup.
	Server string `json:"server"`
	// Status: Status of the backup.
	Status BackupStatus `json:"status"`
	// ACLEnabled: ACL enable boolean of the backup.
	ACLEnabled bool `json:"acl_enabled"`
	// Autologin: Autologin boolean of the backup.
	Autologin bool `json:"autologin"`
	// QuotaSpace: Total quota space of the backup.
	QuotaSpace uint64 `json:"quota_space"`
	// QuotaSpaceUsed: Quota space used of the backup.
	QuotaSpaceUsed uint64 `json:"quota_space_used"`
	// QuotaFiles: Total quota files of the backup.
	QuotaFiles uint64 `json:"quota_files"`
	// QuotaFilesUsed: Quota files used of the backup.
	QuotaFilesUsed uint64 `json:"quota_files_used"`
}

// CancelServerInstallRequest:
type CancelServerInstallRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: Server ID of the server to cancel install.
	ServerID uint64 `json:"-"`
}

// CreateFailoverIPsRequest:
type CreateFailoverIPsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// OfferID: Failover IP offer ID.
	OfferID uint64 `json:"offer_id"`
	// ProjectID: Project ID.
	ProjectID string `json:"project_id"`
	// Quantity: Quantity.
	Quantity uint32 `json:"quantity"`
}

// CreateFailoverIPsResponse:
type CreateFailoverIPsResponse struct {
	// TotalCount:
	TotalCount uint32 `json:"total_count"`
	// Services:
	Services []*Service `json:"services"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *CreateFailoverIPsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *CreateFailoverIPsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*CreateFailoverIPsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Services = append(r.Services, results.Services...)
	r.TotalCount += uint32(len(results.Services))
	return uint32(len(results.Services)), nil
}

// CreateServerRequest:
type CreateServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// OfferID: Offer ID of the new server.
	OfferID uint64 `json:"offer_id"`
	// ServerOptionIDs: Server option IDs of the new server.
	ServerOptionIDs []uint64 `json:"server_option_ids"`
	// ProjectID: Project ID of the new server.
	ProjectID string `json:"project_id"`
	// DatacenterName: Datacenter name of the new server.
	DatacenterName *string `json:"datacenter_name,omitempty"`
}

// DeleteFailoverIPRequest:
type DeleteFailoverIPRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// IPID: ID of the failover IP to delete.
	IPID uint64 `json:"-"`
}

// DeleteServerRequest:
type DeleteServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: Server ID to delete.
	ServerID uint64 `json:"-"`
}

// DeleteServiceRequest:
type DeleteServiceRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServiceID: ID of the service.
	ServiceID uint64 `json:"-"`
}

// DetachFailoverIPFromMacAddressRequest:
type DetachFailoverIPFromMacAddressRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// IPID: ID of the failover IP.
	IPID uint64 `json:"-"`
}

// DetachFailoverIPsRequest:
type DetachFailoverIPsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// FipsIDs: List of IDs of failovers IP to detach.
	FipsIDs []uint64 `json:"fips_ids"`
}

// GetBMCAccessRequest:
type GetBMCAccessRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server to get BMC access.
	ServerID uint64 `json:"-"`
}

// GetFailoverIPRequest:
type GetFailoverIPRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// IPID: ID of the failover IP.
	IPID uint64 `json:"-"`
}

// GetOSRequest:
type GetOSRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// OsID: ID of the OS.
	OsID uint64 `json:"-"`
	// ServerID: ID of the server.
	ServerID uint64 `json:"server_id"`
	// ProjectID: Project ID.
	ProjectID *string `json:"project_id,omitempty"`
}

// GetOfferRequest:
type GetOfferRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// OfferID: ID of offer.
	OfferID uint64 `json:"-"`
	// ProjectID: Project ID.
	ProjectID *string `json:"project_id,omitempty"`
}

// GetOrderedServiceRequest:
type GetOrderedServiceRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// OrderedServiceID:
	OrderedServiceID uint64 `json:"-"`
}

// GetRaidRequest:
type GetRaidRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server.
	ServerID uint64 `json:"-"`
}

// GetRemainingQuotaRequest:
type GetRemainingQuotaRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ProjectID: Project ID.
	ProjectID *string `json:"project_id,omitempty"`
}

// GetRemainingQuotaResponse:
type GetRemainingQuotaResponse struct {
	// FailoverIPQuota: Current failover IP quota.
	FailoverIPQuota uint32 `json:"failover_ip_quota"`
	// FailoverIPRemainingQuota: Remaining failover IP quota.
	FailoverIPRemainingQuota uint32 `json:"failover_ip_remaining_quota"`
	// FailoverBlockQuota: Current failover block quota.
	FailoverBlockQuota uint32 `json:"failover_block_quota"`
	// FailoverBlockRemainingQuota: Remaining failover block quota.
	FailoverBlockRemainingQuota uint32 `json:"failover_block_remaining_quota"`
}

// GetRescueRequest:
type GetRescueRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server to get rescue.
	ServerID uint64 `json:"-"`
}

// GetServerBackupRequest:
type GetServerBackupRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: Server ID of the backup.
	ServerID uint64 `json:"-"`
}

// GetServerDefaultPartitioningRequest:
type GetServerDefaultPartitioningRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server.
	ServerID uint64 `json:"-"`
	// OsID: OS ID of the default partitioning.
	OsID uint64 `json:"-"`
}

// GetServerInstallRequest:
type GetServerInstallRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: Server ID of the server to install.
	ServerID uint64 `json:"-"`
}

// GetServerRequest:
type GetServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server.
	ServerID uint64 `json:"-"`
}

// GetServiceRequest:
type GetServiceRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServiceID: ID of the service.
	ServiceID uint64 `json:"-"`
}

// InstallServerRequest:
type InstallServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: Server ID to install.
	ServerID uint64 `json:"-"`
	// OsID: OS ID to install on the server.
	OsID uint64 `json:"os_id"`
	// Hostname: Hostname of the server.
	Hostname string `json:"hostname"`
	// UserLogin: User to install on the server.
	UserLogin *string `json:"user_login,omitempty"`
	// UserPassword: User password to install on the server.
	UserPassword *string `json:"user_password,omitempty"`
	// PanelPassword: Panel password to install on the server.
	PanelPassword *string `json:"panel_password,omitempty"`
	// RootPassword: Root password to install on the server.
	RootPassword *string `json:"root_password,omitempty"`
	// Partitions: Partitions to install on the server.
	Partitions []*InstallPartition `json:"partitions"`
	// SSHKeyIDs: SSH key IDs authorized on the server.
	SSHKeyIDs []string `json:"ssh_key_ids"`
	// LicenseOfferID: Offer ID of license to install on server.
	LicenseOfferID *uint64 `json:"license_offer_id,omitempty"`
	// IPID: IP to link at the license to install on server.
	IPID *uint64 `json:"ip_id,omitempty"`
}

// ListFailoverIPsRequest:
type ListFailoverIPsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Page: Page number.
	Page *uint32 `json:"page,omitempty"`
	// PageSize: Number of failovers IP per page.
	PageSize *uint32 `json:"page_size,omitempty"`
	// OrderBy: Order of the failovers IP.
	OrderBy ListFailoverIPsRequestOrderBy `json:"order_by"`
	// ProjectID: Filter failovers IP by project ID.
	ProjectID *string `json:"project_id,omitempty"`
	// Search: Filter failovers IP which matching with this field.
	Search *string `json:"search,omitempty"`
	// OnlyAvailable: True: return all failovers IP not attached on server
	// false: return all failovers IP attached on server.
	OnlyAvailable *bool `json:"only_available,omitempty"`
}

// ListFailoverIPsResponse:
type ListFailoverIPsResponse struct {
	// TotalCount: Total count of matching failovers IP.
	TotalCount uint32 `json:"total_count"`
	// FailoverIPs: List of failover IPs that match filters.
	FailoverIPs []*FailoverIP `json:"failover_ips"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListFailoverIPsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListFailoverIPsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListFailoverIPsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.FailoverIPs = append(r.FailoverIPs, results.FailoverIPs...)
	r.TotalCount += uint32(len(results.FailoverIPs))
	return uint32(len(results.FailoverIPs)), nil
}

// ListOSRequest:
type ListOSRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Page: Page number.
	Page *uint32 `json:"page,omitempty"`
	// PageSize: Number of OS per page.
	PageSize *uint32 `json:"page_size,omitempty"`
	// OrderBy: Order of the OS.
	OrderBy ListOSRequestOrderBy `json:"order_by"`
	// Type: Type of the OS.
	Type OSType `json:"type"`
	// ServerID: Filter OS by compatible server ID.
	ServerID uint64 `json:"server_id"`
	// ProjectID: Project ID.
	ProjectID *string `json:"project_id,omitempty"`
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
	Page *uint32 `json:"page,omitempty"`
	// PageSize: Number of offer per page.
	PageSize *uint32 `json:"page_size,omitempty"`
	// OrderBy: Order of the offers.
	OrderBy ListOffersRequestOrderBy `json:"order_by"`
	// CommercialRange: Filter on commercial range.
	CommercialRange *string `json:"commercial_range,omitempty"`
	// Catalog: Filter on catalog.
	Catalog OfferCatalog `json:"catalog"`
	// ProjectID: Project ID.
	ProjectID *string `json:"project_id,omitempty"`
	// IsFailoverIP: Get the current failover IP offer.
	IsFailoverIP *bool `json:"is_failover_ip,omitempty"`
	// IsFailoverBlock: Get the current failover IP block offer.
	IsFailoverBlock *bool `json:"is_failover_block,omitempty"`
	// SoldIn: Filter offers depending on their datacenter.
	SoldIn []string `json:"sold_in,omitempty"`
	// AvailableOnly: Set this filter to true to only return available offers.
	AvailableOnly *bool `json:"available_only,omitempty"`
	// IsRpnSan: Get the RPN SAN offers.
	IsRpnSan *bool `json:"is_rpn_san,omitempty"`
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

// ListServerDisksRequest:
type ListServerDisksRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Page: Page number.
	Page *uint32 `json:"page,omitempty"`
	// PageSize: Number of server disk per page.
	PageSize *uint32 `json:"page_size,omitempty"`
	// OrderBy: Order of the server disks.
	OrderBy ListServerDisksRequestOrderBy `json:"order_by"`
	// ServerID: Server ID of the server disks.
	ServerID uint64 `json:"-"`
}

// ListServerDisksResponse:
type ListServerDisksResponse struct {
	// TotalCount: Total count of matching server disks.
	TotalCount uint32 `json:"total_count"`
	// Disks: Server disks that match filters.
	Disks []*ServerDisk `json:"disks"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServerDisksResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServerDisksResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListServerDisksResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Disks = append(r.Disks, results.Disks...)
	r.TotalCount += uint32(len(results.Disks))
	return uint32(len(results.Disks)), nil
}

// ListServerEventsRequest:
type ListServerEventsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Page: Page number.
	Page *uint32 `json:"page,omitempty"`
	// PageSize: Number of server event per page.
	PageSize *uint32 `json:"page_size,omitempty"`
	// OrderBy: Order of the server events.
	OrderBy ListServerEventsRequestOrderBy `json:"order_by"`
	// ServerID: Server ID of the server events.
	ServerID uint64 `json:"-"`
}

// ListServerEventsResponse:
type ListServerEventsResponse struct {
	// TotalCount: Total count of matching server events.
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

// ListServersRequest:
type ListServersRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Page: Page number.
	Page *uint32 `json:"page,omitempty"`
	// PageSize: Number of server per page.
	PageSize *uint32 `json:"page_size,omitempty"`
	// OrderBy: Order of the servers.
	OrderBy ListServersRequestOrderBy `json:"order_by"`
	// ProjectID: Filter servers by project ID.
	ProjectID *string `json:"project_id,omitempty"`
	// Search: Filter servers by hostname.
	Search *string `json:"search,omitempty"`
}

// ListServersResponse:
type ListServersResponse struct {
	// TotalCount: Total count of matching servers.
	TotalCount uint32 `json:"total_count"`
	// Servers: Servers that match filters.
	Servers []*ServerSummary `json:"servers"`
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

// ListServicesRequest:
type ListServicesRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Page: Page number.
	Page *uint32 `json:"page,omitempty"`
	// PageSize: Number of service per page.
	PageSize *uint32 `json:"page_size,omitempty"`
	// OrderBy: Order of the services.
	OrderBy ListServicesRequestOrderBy `json:"order_by"`
	// ProjectID: Project ID.
	ProjectID *string `json:"project_id,omitempty"`
}

// ListServicesResponse:
type ListServicesResponse struct {
	// TotalCount: Total count of matching services.
	TotalCount uint32 `json:"total_count"`
	// Services: Services that match filters.
	Services []*Service `json:"services"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServicesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServicesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListServicesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Services = append(r.Services, results.Services...)
	r.TotalCount += uint32(len(results.Services))
	return uint32(len(results.Services)), nil
}

// ListSubscribableServerOptionsRequest:
type ListSubscribableServerOptionsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Page: Page number.
	Page *uint32 `json:"page,omitempty"`
	// PageSize: Number of subscribable server option per page.
	PageSize *uint32 `json:"page_size,omitempty"`
	// ServerID: Server ID of the subscribable server options.
	ServerID uint64 `json:"-"`
}

// ListSubscribableServerOptionsResponse:
type ListSubscribableServerOptionsResponse struct {
	// TotalCount: Total count of matching subscribable server options.
	TotalCount uint32 `json:"total_count"`
	// ServerOptions: Server options that match filters.
	ServerOptions []*Offer `json:"server_options"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSubscribableServerOptionsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSubscribableServerOptionsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSubscribableServerOptionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ServerOptions = append(r.ServerOptions, results.ServerOptions...)
	r.TotalCount += uint32(len(results.ServerOptions))
	return uint32(len(results.ServerOptions)), nil
}

// Raid:
type Raid struct {
	// RaidArrays: Details about the RAID controller.
	RaidArrays []*RaidArray `json:"raid_arrays"`
}

// RebootServerRequest:
type RebootServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: Server ID to reboot.
	ServerID uint64 `json:"-"`
}

// Rescue:
type Rescue struct {
	// OsID: OS ID of the rescue.
	OsID uint64 `json:"os_id"`
	// Login: Login of the rescue.
	Login string `json:"login"`
	// Password: Password of the rescue.
	Password string `json:"password"`
	// Protocol: Protocol of the resuce.
	Protocol RescueProtocol `json:"protocol"`
}

// ServerDefaultPartitioning:
type ServerDefaultPartitioning struct {
	// Partitions: Default partitions.
	Partitions []*Partition `json:"partitions"`
}

// ServerInstall:
type ServerInstall struct {
	// OsID:
	OsID uint64 `json:"os_id"`
	// Hostname:
	Hostname string `json:"hostname"`
	// UserLogin:
	UserLogin *string `json:"user_login,omitempty"`
	// Partitions:
	Partitions []*Partition `json:"partitions"`
	// SSHKeyIDs:
	SSHKeyIDs []string `json:"ssh_key_ids"`
	// Status:
	Status ServerInstallStatus `json:"status"`
	// PanelURL:
	PanelURL *string `json:"panel_url,omitempty"`
}

// StartBMCAccessRequest:
type StartBMCAccessRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server to start the BMC access.
	ServerID uint64 `json:"-"`
	// IP: The IP authorized to connect to the given server.
	IP net.IP `json:"ip"`
}

// StartRescueRequest:
type StartRescueRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server to start rescue.
	ServerID uint64 `json:"-"`
	// OsID: OS ID to use to start rescue.
	OsID uint64 `json:"os_id"`
}

// StartServerRequest:
type StartServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: Server ID to start.
	ServerID uint64 `json:"-"`
}

// StopBMCAccessRequest:
type StopBMCAccessRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server to stop BMC access.
	ServerID uint64 `json:"-"`
}

// StopRescueRequest:
type StopRescueRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server to stop rescue.
	ServerID uint64 `json:"-"`
}

// StopServerRequest:
type StopServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: Server ID to stop.
	ServerID uint64 `json:"-"`
}

// SubscribeServerOptionRequest:
type SubscribeServerOptionRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: Server ID to subscribe server option.
	ServerID uint64 `json:"-"`
	// OptionID: Option ID to subscribe.
	OptionID uint64 `json:"option_id"`
}

// SubscribeStorageOptionsRequest:
type SubscribeStorageOptionsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: Server ID of the storage options to subscribe.
	ServerID uint64 `json:"-"`
	// OptionsIDs: Option IDs of the storage options to subscribe.
	OptionsIDs []uint64 `json:"options_ids"`
}

// SubscribeStorageOptionsResponse:
type SubscribeStorageOptionsResponse struct {
	// Services: Services subscribe storage options.
	Services []*Service `json:"services"`
}

// UpdateRaidRequest:
type UpdateRaidRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: ID of the server.
	ServerID uint64 `json:"-"`
	// RaidArrays: RAIDs to update.
	RaidArrays []*UpdatableRaidArray `json:"raid_arrays"`
}

// UpdateReverseRequest:
type UpdateReverseRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// IPID: ID of the IP.
	IPID uint64 `json:"-"`
	// Reverse: Reverse to apply on the IP.
	Reverse string `json:"reverse"`
}

// UpdateServerBackupRequest:
type UpdateServerBackupRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: Server ID to update backup.
	ServerID uint64 `json:"-"`
	// Password: Password of the server backup.
	Password *string `json:"password,omitempty"`
	// Autologin: Autologin of the server backup.
	Autologin *bool `json:"autologin,omitempty"`
	// ACLEnabled: Boolean to enable or disable ACL.
	ACLEnabled *bool `json:"acl_enabled,omitempty"`
}

// UpdateServerRequest:
type UpdateServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: Server ID to update.
	ServerID uint64 `json:"-"`
	// Hostname: Hostname of the server to update.
	Hostname *string `json:"hostname,omitempty"`
	// EnableIPv6: Flag to enable or not the IPv6 of server.
	EnableIPv6 *bool `json:"enable_ipv6,omitempty"`
}

// UpdateServerTagsRequest:
type UpdateServerTagsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: Server ID to update the tags.
	ServerID uint64 `json:"-"`
	// Tags: Tags of server to update.
	Tags []string `json:"tags"`
}

// Dedibox Phoenix API.
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

// ListServers: List baremetal servers for project.
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "search", req.Search)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers",
		Query:  query,
	}

	var resp ListServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetServer: Get the server associated with the given ID.
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetServerBackup:
func (s *API) GetServerBackup(req *GetServerBackupRequest, opts ...scw.RequestOption) (*Backup, error) {
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/backups",
	}

	var resp Backup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateServerBackup:
func (s *API) UpdateServerBackup(req *UpdateServerBackupRequest, opts ...scw.RequestOption) (*Backup, error) {
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/backups",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Backup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSubscribableServerOptions: List subscribable options associated to the given server ID.
func (s *API) ListSubscribableServerOptions(req *ListSubscribableServerOptionsRequest, opts ...scw.RequestOption) (*ListSubscribableServerOptionsResponse, error) {
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

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/subscribable-server-options",
		Query:  query,
	}

	var resp ListSubscribableServerOptionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SubscribeServerOption: Subscribe option for the given server ID.
func (s *API) SubscribeServerOption(req *SubscribeServerOptionRequest, opts ...scw.RequestOption) (*Service, error) {
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/subscribe-server-option",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Service

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateServer: Create a new baremetal server. The order return you a service ID to follow the provisionning status you could call GetService.
func (s *API) CreateServer(req *CreateServerRequest, opts ...scw.RequestOption) (*Service, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Service

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SubscribeStorageOptions: Subscribe storage option for the given server ID.
func (s *API) SubscribeStorageOptions(req *SubscribeStorageOptionsRequest, opts ...scw.RequestOption) (*SubscribeStorageOptionsResponse, error) {
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/subscribe-storage-options",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SubscribeStorageOptionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateServer: Update the server associated with the given ID.
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
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

// UpdateServerTags:
func (s *API) UpdateServerTags(req *UpdateServerTagsRequest, opts ...scw.RequestOption) (*Server, error) {
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/tags",
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

// RebootServer: Reboot the server associated with the given ID, use boot param to reboot in rescue.
func (s *API) RebootServer(req *RebootServerRequest, opts ...scw.RequestOption) error {
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
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/reboot",
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

// StartServer: Start the server associated with the given ID.
func (s *API) StartServer(req *StartServerRequest, opts ...scw.RequestOption) error {
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
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/start",
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

// StopServer: Stop the server associated with the given ID.
func (s *API) StopServer(req *StopServerRequest, opts ...scw.RequestOption) error {
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
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/stop",
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

// DeleteServer: Delete the server associated with the given ID.
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListServerEvents: List events associated to the given server ID.
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/events",
		Query:  query,
	}

	var resp ListServerEventsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListServerDisks: List disks associated to the given server ID.
func (s *API) ListServerDisks(req *ListServerDisksRequest, opts ...scw.RequestOption) (*ListServerDisksResponse, error) {
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/disks",
		Query:  query,
	}

	var resp ListServerDisksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOrderedService:
func (s *API) GetOrderedService(req *GetOrderedServiceRequest, opts ...scw.RequestOption) (*Service, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.OrderedServiceID) == "" {
		return nil, errors.New("field OrderedServiceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/ordered-services/" + fmt.Sprint(req.OrderedServiceID) + "",
	}

	var resp Service

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetService: Get the service associated with the given ID.
func (s *API) GetService(req *GetServiceRequest, opts ...scw.RequestOption) (*Service, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServiceID) == "" {
		return nil, errors.New("field ServiceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/services/" + fmt.Sprint(req.ServiceID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Service

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteService: Delete the service associated with the given ID.
func (s *API) DeleteService(req *DeleteServiceRequest, opts ...scw.RequestOption) (*Service, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServiceID) == "" {
		return nil, errors.New("field ServiceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/services/" + fmt.Sprint(req.ServiceID) + "",
	}

	var resp Service

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListServices: List services.
func (s *API) ListServices(req *ListServicesRequest, opts ...scw.RequestOption) (*ListServicesResponse, error) {
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/services",
		Query:  query,
	}

	var resp ListServicesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// InstallServer: Install an OS on the server associated with the given ID.
func (s *API) InstallServer(req *InstallServerRequest, opts ...scw.RequestOption) (*ServerInstall, error) {
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/install",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ServerInstall

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetServerInstall: Get the server installation status associated with the given server ID.
func (s *API) GetServerInstall(req *GetServerInstallRequest, opts ...scw.RequestOption) (*ServerInstall, error) {
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/install",
	}

	var resp ServerInstall

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CancelServerInstall: Cancels the current server installation associated with the given server ID.
func (s *API) CancelServerInstall(req *CancelServerInstallRequest, opts ...scw.RequestOption) error {
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
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/cancel-install",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetServerDefaultPartitioning: Get the server default partitioning schema associated with the given server ID and OS ID.
func (s *API) GetServerDefaultPartitioning(req *GetServerDefaultPartitioningRequest, opts ...scw.RequestOption) (*ServerDefaultPartitioning, error) {
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

	if fmt.Sprint(req.OsID) == "" {
		return nil, errors.New("field OsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/partitioning/" + fmt.Sprint(req.OsID) + "",
	}

	var resp ServerDefaultPartitioning

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StartBMCAccess: Start BMC (Baseboard Management Controller) access associated with the given ID.
// The BMC (Baseboard Management Controller) access is available one hour after the installation of the server.
func (s *API) StartBMCAccess(req *StartBMCAccessRequest, opts ...scw.RequestOption) error {
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
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/bmc-access",
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

// GetBMCAccess: Get the BMC (Baseboard Management Controller) access associated with the given ID.
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/bmc-access",
	}

	var resp BMCAccess

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StopBMCAccess: Stop BMC (Baseboard Management Controller) access associated with the given ID.
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/bmc-access",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListOffers: List all available server offers.
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "commercial_range", req.CommercialRange)
	parameter.AddToQuery(query, "catalog", req.Catalog)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "is_failover_ip", req.IsFailoverIP)
	parameter.AddToQuery(query, "is_failover_block", req.IsFailoverBlock)
	if len(req.SoldIn) != 0 {
		parameter.AddToQuery(query, "sold_in", strings.Join(req.SoldIn, ","))
	}
	parameter.AddToQuery(query, "available_only", req.AvailableOnly)
	parameter.AddToQuery(query, "is_rpn_san", req.IsRpnSan)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/offers",
		Query:  query,
	}

	var resp ListOffersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOffer: Return specific offer for the given ID.
func (s *API) GetOffer(req *GetOfferRequest, opts ...scw.RequestOption) (*Offer, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.OfferID) == "" {
		return nil, errors.New("field OfferID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/offers/" + fmt.Sprint(req.OfferID) + "",
		Query:  query,
	}

	var resp Offer

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListOS: List all available OS that can be install on a baremetal server.
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "type", req.Type)
	parameter.AddToQuery(query, "server_id", req.ServerID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/os",
		Query:  query,
	}

	var resp ListOSResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOS: Return specific OS for the given ID.
func (s *API) GetOS(req *GetOSRequest, opts ...scw.RequestOption) (*OS, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "server_id", req.ServerID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.OsID) == "" {
		return nil, errors.New("field OsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/os/" + fmt.Sprint(req.OsID) + "",
		Query:  query,
	}

	var resp OS

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateReverse: Update reverse of ip associated with the given ID.
func (s *API) UpdateReverse(req *UpdateReverseRequest, opts ...scw.RequestOption) (*IP, error) {
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/reverses/" + fmt.Sprint(req.IPID) + "",
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

// CreateFailoverIPs: Order X failover IPs.
func (s *API) CreateFailoverIPs(req *CreateFailoverIPsRequest, opts ...scw.RequestOption) (*CreateFailoverIPsResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/failover-ips",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateFailoverIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachFailoverIPs: Attach failovers on the server associated with the given ID.
func (s *API) AttachFailoverIPs(req *AttachFailoverIPsRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/failover-ips/attach",
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

// DetachFailoverIPs: Detach failovers on the server associated with the given ID.
func (s *API) DetachFailoverIPs(req *DetachFailoverIPsRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/failover-ips/detach",
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

// AttachFailoverIPToMacAddress: Attach a failover IP to a MAC address.
func (s *API) AttachFailoverIPToMacAddress(req *AttachFailoverIPToMacAddressRequest, opts ...scw.RequestOption) (*IP, error) {
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
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/failover-ips/" + fmt.Sprint(req.IPID) + "/attach-to-mac-address",
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

// DetachFailoverIPFromMacAddress: Detach a failover IP from a MAC address.
func (s *API) DetachFailoverIPFromMacAddress(req *DetachFailoverIPFromMacAddressRequest, opts ...scw.RequestOption) (*IP, error) {
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
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/failover-ips/" + fmt.Sprint(req.IPID) + "/detach-from-mac-address",
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

// DeleteFailoverIP: Delete the failover associated with the given ID.
func (s *API) DeleteFailoverIP(req *DeleteFailoverIPRequest, opts ...scw.RequestOption) error {
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/failover-ips/" + fmt.Sprint(req.IPID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListFailoverIPs: List failovers servers for project.
func (s *API) ListFailoverIPs(req *ListFailoverIPsRequest, opts ...scw.RequestOption) (*ListFailoverIPsResponse, error) {
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
	parameter.AddToQuery(query, "search", req.Search)
	parameter.AddToQuery(query, "only_available", req.OnlyAvailable)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/failover-ips",
		Query:  query,
	}

	var resp ListFailoverIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFailoverIP: Get the server associated with the given ID.
func (s *API) GetFailoverIP(req *GetFailoverIPRequest, opts ...scw.RequestOption) (*FailoverIP, error) {
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/failover-ips/" + fmt.Sprint(req.IPID) + "",
	}

	var resp FailoverIP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRemainingQuota: Get remaining quota.
func (s *API) GetRemainingQuota(req *GetRemainingQuotaRequest, opts ...scw.RequestOption) (*GetRemainingQuotaResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/remaining-quota",
		Query:  query,
	}

	var resp GetRemainingQuotaResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRaid: Return raid for the given server ID.
func (s *API) GetRaid(req *GetRaidRequest, opts ...scw.RequestOption) (*Raid, error) {
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/raid",
	}

	var resp Raid

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateRaid: Update RAID associated with the given server ID.
func (s *API) UpdateRaid(req *UpdateRaidRequest, opts ...scw.RequestOption) error {
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
		Method: "POST",
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/update-raid",
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

// StartRescue: Start in rescue the server associated with the given ID.
func (s *API) StartRescue(req *StartRescueRequest, opts ...scw.RequestOption) (*Rescue, error) {
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/rescue",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Rescue

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRescue: Return rescue information for the given server ID.
func (s *API) GetRescue(req *GetRescueRequest, opts ...scw.RequestOption) (*Rescue, error) {
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/rescue",
	}

	var resp Rescue

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StopRescue: Stop rescue on the server associated with the given ID.
func (s *API) StopRescue(req *StopRescueRequest, opts ...scw.RequestOption) error {
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
		Path:   "/dedibox/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/rescue",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}
