// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package domain provides methods and message types of the domain v2beta1 API.
package domain

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
	std "github.com/scaleway/scaleway-sdk-go/api/std"
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

type ContactEmailStatus string

const (
	// If unspecified, the status is unknown by default.
	ContactEmailStatusEmailStatusUnknown = ContactEmailStatus("email_status_unknown")
	// The contact email has been validated.
	ContactEmailStatusValidated = ContactEmailStatus("validated")
	// The contact email has not been validated.
	ContactEmailStatusNotValidated = ContactEmailStatus("not_validated")
	// The contact email is invalid.
	ContactEmailStatusInvalidEmail = ContactEmailStatus("invalid_email")
)

func (enum ContactEmailStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "email_status_unknown"
	}
	return string(enum)
}

func (enum ContactEmailStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContactEmailStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContactEmailStatus(ContactEmailStatus(tmp).String())
	return nil
}

type ContactExtensionFRMode string

const (
	// If unspecified, the status is unknown by default.
	ContactExtensionFRModeModeUnknown = ContactExtensionFRMode("mode_unknown")
	// The contact is a physical person (only for .fr domains).
	ContactExtensionFRModeIndividual = ContactExtensionFRMode("individual")
	// The contact is a company with a SIRET/SIREN code (only for .fr domains).
	ContactExtensionFRModeCompanyIdentificationCode = ContactExtensionFRMode("company_identification_code")
	// The contact has a Data Universal Numbering System code (only for .fr domains).
	ContactExtensionFRModeDuns = ContactExtensionFRMode("duns")
	// The contact has a local or a country ID (only for .fr domains).
	ContactExtensionFRModeLocal = ContactExtensionFRMode("local")
	// The contact is an association (only for .fr domains).
	ContactExtensionFRModeAssociation = ContactExtensionFRMode("association")
	// The contact is a brand (only for .fr domains).
	ContactExtensionFRModeTrademark = ContactExtensionFRMode("trademark")
	// The contact has an intervention code (DSIA) from AFNIC (only for .fr domains).
	ContactExtensionFRModeCodeAuthAfnic = ContactExtensionFRMode("code_auth_afnic")
)

func (enum ContactExtensionFRMode) String() string {
	if enum == "" {
		// return default value if empty
		return "mode_unknown"
	}
	return string(enum)
}

func (enum ContactExtensionFRMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContactExtensionFRMode) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContactExtensionFRMode(ContactExtensionFRMode(tmp).String())
	return nil
}

type ContactExtensionNLLegalForm string

const (
	// If unspecified, the status is unknown by default.
	ContactExtensionNLLegalFormLegalFormUnknown = ContactExtensionNLLegalForm("legal_form_unknown")
	// The contact's legal form is not listed below (only for .nl domains).
	ContactExtensionNLLegalFormOther = ContactExtensionNLLegalForm("other")
	// The contact is a non-Dutch EC company (only for .nl domains).
	ContactExtensionNLLegalFormNonDutchEuCompany = ContactExtensionNLLegalForm("non_dutch_eu_company")
	// The contact is a non-Dutch legal form/enterprise/subsidiary (only for .nl domains).
	ContactExtensionNLLegalFormNonDutchLegalFormEnterpriseSubsidiary = ContactExtensionNLLegalForm("non_dutch_legal_form_enterprise_subsidiary")
	// The contact is a limited company (only for .nl domains).
	ContactExtensionNLLegalFormLimitedCompany = ContactExtensionNLLegalForm("limited_company")
	// The contact is a limited company in formation (only for .nl domains).
	ContactExtensionNLLegalFormLimitedCompanyInFormation = ContactExtensionNLLegalForm("limited_company_in_formation")
	// The contact is a cooperative (only for .nl domains).
	ContactExtensionNLLegalFormCooperative = ContactExtensionNLLegalForm("cooperative")
	// The contact is a limited Partnership (only for .nl domains).
	ContactExtensionNLLegalFormLimitedPartnership = ContactExtensionNLLegalForm("limited_partnership")
	// The contact is a sole trader (only for .nl domains).
	ContactExtensionNLLegalFormSoleCompany = ContactExtensionNLLegalForm("sole_company")
	// The contact is a European Economic Interest Group (only for .nl domains).
	ContactExtensionNLLegalFormEuropeanEconomicInterestGroup = ContactExtensionNLLegalForm("european_economic_interest_group")
	// The contact is a religious society (only for .nl domains).
	ContactExtensionNLLegalFormReligiousEntity = ContactExtensionNLLegalForm("religious_entity")
	// The contact is a partnership (only for .nl domains).
	ContactExtensionNLLegalFormPartnership = ContactExtensionNLLegalForm("partnership")
	// The contact is a public Company (only for .nl domains).
	ContactExtensionNLLegalFormPublicCompany = ContactExtensionNLLegalForm("public_company")
	// The contact is a mutual benefit company (only for .nl domains).
	ContactExtensionNLLegalFormMutualBenefitCompany = ContactExtensionNLLegalForm("mutual_benefit_company")
	// The contact is a natural person (only for .nl domains).
	ContactExtensionNLLegalFormResidential = ContactExtensionNLLegalForm("residential")
	// The contact is a shipping company (only for .nl domains).
	ContactExtensionNLLegalFormShippingCompany = ContactExtensionNLLegalForm("shipping_company")
	// The contact is a foundation (only for .nl domains).
	ContactExtensionNLLegalFormFoundation = ContactExtensionNLLegalForm("foundation")
	// The contact is a association (only for .nl domains).
	ContactExtensionNLLegalFormAssociation = ContactExtensionNLLegalForm("association")
	// The contact is a trading partnership (only for .nl domains).
	ContactExtensionNLLegalFormTradingPartnership = ContactExtensionNLLegalForm("trading_partnership")
	// The contact is a physical person (only for .nl domains).
	ContactExtensionNLLegalFormNaturalPerson = ContactExtensionNLLegalForm("natural_person")
)

func (enum ContactExtensionNLLegalForm) String() string {
	if enum == "" {
		// return default value if empty
		return "legal_form_unknown"
	}
	return string(enum)
}

func (enum ContactExtensionNLLegalForm) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContactExtensionNLLegalForm) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContactExtensionNLLegalForm(ContactExtensionNLLegalForm(tmp).String())
	return nil
}

type ContactLegalForm string

const (
	// If unspecified, the status is unknown by default.
	ContactLegalFormLegalFormUnknown = ContactLegalForm("legal_form_unknown")
	// The contact is a physical person.
	ContactLegalFormIndividual = ContactLegalForm("individual")
	// The contact is a corporate or a society.
	ContactLegalFormCorporate = ContactLegalForm("corporate")
	// The contact is an association.
	ContactLegalFormAssociation = ContactLegalForm("association")
	// The contact is not represented by a physical person, a corporate or an association.
	ContactLegalFormOther = ContactLegalForm("other")
)

func (enum ContactLegalForm) String() string {
	if enum == "" {
		// return default value if empty
		return "legal_form_unknown"
	}
	return string(enum)
}

func (enum ContactLegalForm) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContactLegalForm) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContactLegalForm(ContactLegalForm(tmp).String())
	return nil
}

type DNSZoneStatus string

const (
	// If unspecified, the DNS zone's status is unknown by default.
	DNSZoneStatusUnknown = DNSZoneStatus("unknown")
	// The DNS zone is active and healthy.
	DNSZoneStatusActive = DNSZoneStatus("active")
	// The DNS zone is updating.
	DNSZoneStatusPending = DNSZoneStatus("pending")
	// An error occurred after updating the DNS zone.
	DNSZoneStatusError = DNSZoneStatus("error")
	// The DNS zone is locked and cannot be updated anymore.
	DNSZoneStatusLocked = DNSZoneStatus("locked")
)

func (enum DNSZoneStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum DNSZoneStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DNSZoneStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DNSZoneStatus(DNSZoneStatus(tmp).String())
	return nil
}

type DSRecordAlgorithm string

const (
	// Code 1, algorithm: 'RSAMD5'.
	DSRecordAlgorithmRsamd5 = DSRecordAlgorithm("rsamd5")
	// Code 2, algorithm: 'DIFFIE_HELLMAN'.
	DSRecordAlgorithmDh = DSRecordAlgorithm("dh")
	// Code 3, algorithm: 'DSA_SHA1'.
	DSRecordAlgorithmDsa = DSRecordAlgorithm("dsa")
	// Code 5, algorithm: 'RSA_SHA1'.
	DSRecordAlgorithmRsasha1 = DSRecordAlgorithm("rsasha1")
	// Code 6, algorithm: 'DSA_NSEC3_SHA1'.
	DSRecordAlgorithmDsaNsec3Sha1 = DSRecordAlgorithm("dsa_nsec3_sha1")
	// Code 7, algorithm: 'RSASHA1_NSEC3_SHA1'.
	DSRecordAlgorithmRsasha1Nsec3Sha1 = DSRecordAlgorithm("rsasha1_nsec3_sha1")
	// Code 8, algorithm: 'RSASHA256'.
	DSRecordAlgorithmRsasha256 = DSRecordAlgorithm("rsasha256")
	// Code 10, algorithm: 'RSASHA512'.
	DSRecordAlgorithmRsasha512 = DSRecordAlgorithm("rsasha512")
	// Code 12, algorithm: 'ECC_GOST'.
	DSRecordAlgorithmEccGost = DSRecordAlgorithm("ecc_gost")
	// Code 13, algorithm: 'ECDSAP256SHA256'.
	DSRecordAlgorithmEcdsap256sha256 = DSRecordAlgorithm("ecdsap256sha256")
	// Code 14, algorithm: 'ECDSAP384SHA384'.
	DSRecordAlgorithmEcdsap384sha384 = DSRecordAlgorithm("ecdsap384sha384")
	// Code 15, algorithm: 'ED25519'.
	DSRecordAlgorithmEd25519 = DSRecordAlgorithm("ed25519")
	// Code 16, algorithm: 'ED448'.
	DSRecordAlgorithmEd448 = DSRecordAlgorithm("ed448")
)

func (enum DSRecordAlgorithm) String() string {
	if enum == "" {
		// return default value if empty
		return "rsamd5"
	}
	return string(enum)
}

func (enum DSRecordAlgorithm) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DSRecordAlgorithm) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DSRecordAlgorithm(DSRecordAlgorithm(tmp).String())
	return nil
}

type DSRecordDigestType string

const (
	// Code 1, digest type: 'SHA_1'.
	DSRecordDigestTypeSha1 = DSRecordDigestType("sha_1")
	// Code 2, digest type: 'SHA_256'.
	DSRecordDigestTypeSha256 = DSRecordDigestType("sha_256")
	// Code 3, digest type: 'GOST_R_34_11_94'.
	DSRecordDigestTypeGostR34_11_94 = DSRecordDigestType("gost_r_34_11_94")
	// Code 4, digest type: 'SHA_384'.
	DSRecordDigestTypeSha384 = DSRecordDigestType("sha_384")
)

func (enum DSRecordDigestType) String() string {
	if enum == "" {
		// return default value if empty
		return "sha_1"
	}
	return string(enum)
}

func (enum DSRecordDigestType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DSRecordDigestType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DSRecordDigestType(DSRecordDigestType(tmp).String())
	return nil
}

type DomainFeatureStatus string

const (
	// Default unknown status.
	DomainFeatureStatusFeatureStatusUnknown = DomainFeatureStatus("feature_status_unknown")
	// A feature (auto renew, DNSSEC) is being enabled.
	DomainFeatureStatusEnabling = DomainFeatureStatus("enabling")
	// A feature (auto renew, DNSSEC) has been enabled.
	DomainFeatureStatusEnabled = DomainFeatureStatus("enabled")
	// A feature (auto renew, DNSSEC) is being disabled.
	DomainFeatureStatusDisabling = DomainFeatureStatus("disabling")
	// A feature (auto renew, DNSSEC) has been disabled.
	DomainFeatureStatusDisabled = DomainFeatureStatus("disabled")
)

func (enum DomainFeatureStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "feature_status_unknown"
	}
	return string(enum)
}

func (enum DomainFeatureStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainFeatureStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainFeatureStatus(DomainFeatureStatus(tmp).String())
	return nil
}

type DomainRegistrationStatusTransferStatus string

const (
	// If unspecified, the status is unknown by default.
	DomainRegistrationStatusTransferStatusStatusUnknown = DomainRegistrationStatusTransferStatus("status_unknown")
	// The domain transfer is being initialized.
	DomainRegistrationStatusTransferStatusPending = DomainRegistrationStatusTransferStatus("pending")
	// The domain transfer has started. The process can be accelerated if you accept the vote.
	DomainRegistrationStatusTransferStatusWaitingVote = DomainRegistrationStatusTransferStatus("waiting_vote")
	// The domain transfer has been rejected.
	DomainRegistrationStatusTransferStatusRejected = DomainRegistrationStatusTransferStatus("rejected")
	// The domain transfer has been accepted. Your resources are being created.
	DomainRegistrationStatusTransferStatusProcessing = DomainRegistrationStatusTransferStatus("processing")
	// The domain transfer is complete.
	DomainRegistrationStatusTransferStatusDone = DomainRegistrationStatusTransferStatus("done")
)

func (enum DomainRegistrationStatusTransferStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "status_unknown"
	}
	return string(enum)
}

func (enum DomainRegistrationStatusTransferStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainRegistrationStatusTransferStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainRegistrationStatusTransferStatus(DomainRegistrationStatusTransferStatus(tmp).String())
	return nil
}

type DomainStatus string

const (
	// If unspecified, the status is unknown by default.
	DomainStatusStatusUnknown = DomainStatus("status_unknown")
	// The domain is active.
	DomainStatusActive = DomainStatus("active")
	// The domain is in the process of being created.
	DomainStatusCreating = DomainStatus("creating")
	// An error occurred during the domain's creation process.
	DomainStatusCreateError = DomainStatus("create_error")
	// The domain is being renewed.
	DomainStatusRenewing = DomainStatus("renewing")
	// An error occurred during the domain's renewal process.
	DomainStatusRenewError = DomainStatus("renew_error")
	// The domain is being transferred to Scaleway Domains and DNS.
	DomainStatusXfering = DomainStatus("xfering")
	// An error occurred during the domain's transfer process.
	DomainStatusXferError = DomainStatus("xfer_error")
	// The domain is expired but it can be renewed.
	DomainStatusExpired = DomainStatus("expired")
	// The domain is expiring but it is still renewable.
	DomainStatusExpiring = DomainStatus("expiring")
	// The domain's information is updating.
	DomainStatusUpdating = DomainStatus("updating")
	// The external domain has not yet been validated. It will be automatically removed after 48 hours if it still has not been validated by then.
	DomainStatusChecking = DomainStatus("checking")
	// The domain is locked. Contact Scaleway's support team for more information.
	DomainStatusLocked = DomainStatus("locked")
	// The domain will be deleted soon. This process cannot be canceled.
	DomainStatusDeleting = DomainStatus("deleting")
)

func (enum DomainStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "status_unknown"
	}
	return string(enum)
}

func (enum DomainStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainStatus(DomainStatus(tmp).String())
	return nil
}

type HostStatus string

const (
	// If unspecified, the status is unknown by default.
	HostStatusUnknownStatus = HostStatus("unknown_status")
	// The host is active.
	HostStatusActive = HostStatus("active")
	// The host is being updated.
	HostStatusUpdating = HostStatus("updating")
	// The host is being deleted.
	HostStatusDeleting = HostStatus("deleting")
)

func (enum HostStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum HostStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *HostStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = HostStatus(HostStatus(tmp).String())
	return nil
}

type ListContactsRequestRole string

const (
	ListContactsRequestRoleUnknownRole = ListContactsRequestRole("unknown_role")
	// The contact is a domain's owner.
	ListContactsRequestRoleOwner = ListContactsRequestRole("owner")
	// The contact is a domain's administrative contact.
	ListContactsRequestRoleAdministrative = ListContactsRequestRole("administrative")
	// The contact is a domain's technical contact.
	ListContactsRequestRoleTechnical = ListContactsRequestRole("technical")
)

func (enum ListContactsRequestRole) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_role"
	}
	return string(enum)
}

func (enum ListContactsRequestRole) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListContactsRequestRole) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListContactsRequestRole(ListContactsRequestRole(tmp).String())
	return nil
}

type ListDNSZoneRecordsRequestOrderBy string

const (
	// Order by record name (ascending).
	ListDNSZoneRecordsRequestOrderByNameAsc = ListDNSZoneRecordsRequestOrderBy("name_asc")
	// Order by record name (descending).
	ListDNSZoneRecordsRequestOrderByNameDesc = ListDNSZoneRecordsRequestOrderBy("name_desc")
)

func (enum ListDNSZoneRecordsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListDNSZoneRecordsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDNSZoneRecordsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDNSZoneRecordsRequestOrderBy(ListDNSZoneRecordsRequestOrderBy(tmp).String())
	return nil
}

type ListDNSZonesRequestOrderBy string

const (
	// Order by domain name (ascending).
	ListDNSZonesRequestOrderByDomainAsc = ListDNSZonesRequestOrderBy("domain_asc")
	// Order by domain name (descending).
	ListDNSZonesRequestOrderByDomainDesc = ListDNSZonesRequestOrderBy("domain_desc")
	// Order by subdomain name (ascending).
	ListDNSZonesRequestOrderBySubdomainAsc = ListDNSZonesRequestOrderBy("subdomain_asc")
	// Order by subdomain name (descending).
	ListDNSZonesRequestOrderBySubdomainDesc = ListDNSZonesRequestOrderBy("subdomain_desc")
)

func (enum ListDNSZonesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "domain_asc"
	}
	return string(enum)
}

func (enum ListDNSZonesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDNSZonesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDNSZonesRequestOrderBy(ListDNSZonesRequestOrderBy(tmp).String())
	return nil
}

type ListDomainsRequestOrderBy string

const (
	// Order by domain name (ascending).
	ListDomainsRequestOrderByDomainAsc = ListDomainsRequestOrderBy("domain_asc")
	// Order by domain name (descending).
	ListDomainsRequestOrderByDomainDesc = ListDomainsRequestOrderBy("domain_desc")
)

func (enum ListDomainsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "domain_asc"
	}
	return string(enum)
}

func (enum ListDomainsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDomainsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDomainsRequestOrderBy(ListDomainsRequestOrderBy(tmp).String())
	return nil
}

type ListRenewableDomainsRequestOrderBy string

const (
	// Order by domain name (ascending).
	ListRenewableDomainsRequestOrderByDomainAsc = ListRenewableDomainsRequestOrderBy("domain_asc")
	// Order by domain name (descending).
	ListRenewableDomainsRequestOrderByDomainDesc = ListRenewableDomainsRequestOrderBy("domain_desc")
)

func (enum ListRenewableDomainsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "domain_asc"
	}
	return string(enum)
}

func (enum ListRenewableDomainsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListRenewableDomainsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListRenewableDomainsRequestOrderBy(ListRenewableDomainsRequestOrderBy(tmp).String())
	return nil
}

type ListTasksRequestOrderBy string

const (
	// Order by domain name (descending).
	ListTasksRequestOrderByDomainDesc = ListTasksRequestOrderBy("domain_desc")
	// Order by domain name (ascending).
	ListTasksRequestOrderByDomainAsc = ListTasksRequestOrderBy("domain_asc")
	// Order by type (ascending).
	ListTasksRequestOrderByTypeAsc = ListTasksRequestOrderBy("type_asc")
	// Order by type (descending).
	ListTasksRequestOrderByTypeDesc = ListTasksRequestOrderBy("type_desc")
	// Order by status (ascending).
	ListTasksRequestOrderByStatusAsc = ListTasksRequestOrderBy("status_asc")
	// Order by status (descending).
	ListTasksRequestOrderByStatusDesc = ListTasksRequestOrderBy("status_desc")
	// Order by updated date (ascending).
	ListTasksRequestOrderByUpdatedAtAsc = ListTasksRequestOrderBy("updated_at_asc")
	// Order by updated date (descending).
	ListTasksRequestOrderByUpdatedAtDesc = ListTasksRequestOrderBy("updated_at_desc")
)

func (enum ListTasksRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "domain_desc"
	}
	return string(enum)
}

func (enum ListTasksRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListTasksRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListTasksRequestOrderBy(ListTasksRequestOrderBy(tmp).String())
	return nil
}

type RawFormat string

const (
	// If unspecified, the format is unknown by default.
	RawFormatUnknownRawFormat = RawFormat("unknown_raw_format")
	// Export the DNS zone in text bind format.
	RawFormatBind = RawFormat("bind")
)

func (enum RawFormat) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_raw_format"
	}
	return string(enum)
}

func (enum RawFormat) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RawFormat) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RawFormat(RawFormat(tmp).String())
	return nil
}

type RecordHTTPServiceConfigStrategy string

const (
	// Returns a random IP based of the list of IPs available.
	RecordHTTPServiceConfigStrategyRandom = RecordHTTPServiceConfigStrategy("random")
	// Based on the hash of bestwho, returns a random functioning IP out of the best IPs available.
	RecordHTTPServiceConfigStrategyHashed = RecordHTTPServiceConfigStrategy("hashed")
	// Return all functioning IPs available.
	RecordHTTPServiceConfigStrategyAll = RecordHTTPServiceConfigStrategy("all")
)

func (enum RecordHTTPServiceConfigStrategy) String() string {
	if enum == "" {
		// return default value if empty
		return "random"
	}
	return string(enum)
}

func (enum RecordHTTPServiceConfigStrategy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RecordHTTPServiceConfigStrategy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RecordHTTPServiceConfigStrategy(RecordHTTPServiceConfigStrategy(tmp).String())
	return nil
}

type RecordType string

const (
	// If unspecified, the record's type is unknown by default.
	RecordTypeUnknown = RecordType("unknown")
	// An A record contains an IP address. Example: '203.0.113.210'.
	RecordTypeA = RecordType("A")
	// An AAAA record contains an IPv6 address. Example: '2001:DB8:2000:bf0::1'.
	RecordTypeAAAA = RecordType("AAAA")
	// A CNAME record specifies the canonical name of a record. Example 'webserver-01.yourcompany.com'.
	RecordTypeCNAME = RecordType("CNAME")
	// A TXT record can be used to attach textual data to a domain. Example 'v=spf1 include:_spf.tem.scw.cloud -all'.
	RecordTypeTXT = RecordType("TXT")
	// SRV records can be used to encode the location and port of services on a domain name. Example : '20 443 sipdir.scaleway.example.com'.
	RecordTypeSRV = RecordType("SRV")
	// TLSA records are used to bind SSL/TLS certificates to named hosts and ports.
	RecordTypeTLSA = RecordType("TLSA")
	// An MX record specifies a mail exchanger host for a domain. Example '10 mx.example.net.'.
	RecordTypeMX = RecordType("MX")
	// Specifies nameservers for a domain. Example: 'ns1.yourcompany.com'.
	RecordTypeNS = RecordType("NS")
	// A reverse pointer is used to specify the hostname that belongs to an IP or an IPv6 address. Example: 'www.yourcompany.com.'.
	RecordTypePTR = RecordType("PTR")
	// A 'Certification Authority Authorization' record is used to specify certificate authorities that may issue certificates for a domain. Example: '0 issue ca.yourcompany.com'.
	RecordTypeCAA = RecordType("CAA")
	// The ALIAS pseudo-record type is supported to provide CNAME-like mechanisms on a zone's apex.
	RecordTypeALIAS = RecordType("ALIAS")
	// A LOC record is a way of expressing geographic location information for a domain name. It contains WGS84 latitude, longitude and altitude. Example: '51 56 0.123 N 5 54 0.000 E 4.00m 1.00m 10000.00m 10.00m'.
	RecordTypeLOC = RecordType("LOC")
	// An SSHFP record type is used for storing Secure Shell (SSH) fingerprints. Example: '2 1 123456789abcdef67890123456789abcdef67890'.
	RecordTypeSSHFP = RecordType("SSHFP")
	// A Hardware Info record is used to specify the CPU and operating system you are using. Example: 'i386 Linux'.
	RecordTypeHINFO = RecordType("HINFO")
	// A Responsible Person record stores the mailbox name and the more-information pointer. Example: 'michel.yourcompany.com michel.people.yourcompany.com', to indicate that michel@yourcompany.com is responsible and that more information about Michel is available by querying the `TXT` record of 'michel.people.yourcompany.com'.
	RecordTypeRP = RecordType("RP")
	// A URI record, is used to publish mappings from hostnames to URIs. Example: '10 1 'ftp://ftp.yourcompany.com/public'.
	RecordTypeURI = RecordType("URI")
	// DS records (Delegation Signer) are used to secure delegations (DNSSEC). Example: '2371 13 2 1F987CC6583E92DF0890718C42'.
	RecordTypeDS = RecordType("DS")
	// A Naming Authority Pointer record is used to set rules for how websites process requests. Example: '100 50 "s" "z3950+I2L+I2C" "" _z3950._tcp.yourcompany.com'.
	RecordTypeNAPTR = RecordType("NAPTR")
	// A DNAME record provides redirection from one part of the DNS name tree to another part of the DNS name tree. DNAME and CNAME records both cause a lookup to (potentially) return data corresponding to a different domain name from the queried domain name. Example: 'yourcompany.com'.
	RecordTypeDNAME = RecordType("DNAME")
)

func (enum RecordType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum RecordType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RecordType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RecordType(RecordType(tmp).String())
	return nil
}

type RenewableDomainStatus string

const (
	// If unspecified, the status is unknown by default.
	RenewableDomainStatusUnknown = RenewableDomainStatus("unknown")
	// The domain can be renewed.
	RenewableDomainStatusRenewable = RenewableDomainStatus("renewable")
	// The domain is expired, but it still can be late renewed.
	RenewableDomainStatusLateReneweable = RenewableDomainStatus("late_reneweable")
	// The domain cannot be renewed.
	RenewableDomainStatusNotRenewable = RenewableDomainStatus("not_renewable")
)

func (enum RenewableDomainStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum RenewableDomainStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RenewableDomainStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RenewableDomainStatus(RenewableDomainStatus(tmp).String())
	return nil
}

type SSLCertificateStatus string

const (
	// If unspecified, the SSL certificate's status is unknown by default.
	SSLCertificateStatusUnknown = SSLCertificateStatus("unknown")
	// The SSL certificate has been created but it has not been processed yet.
	SSLCertificateStatusNew = SSLCertificateStatus("new")
	// The SSL certificate's status is pending.
	SSLCertificateStatusPending = SSLCertificateStatus("pending")
	// The SSL certificate has been created and processed.
	SSLCertificateStatusSuccess = SSLCertificateStatus("success")
	// An error occurred during the SSL certificate's creation.
	SSLCertificateStatusError = SSLCertificateStatus("error")
)

func (enum SSLCertificateStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum SSLCertificateStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SSLCertificateStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SSLCertificateStatus(SSLCertificateStatus(tmp).String())
	return nil
}

type TaskStatus string

const (
	// If unspecified, the status is unavailable by default.
	TaskStatusUnavailable = TaskStatus("unavailable")
	// The task has been created but it has not yet started.
	TaskStatusNew = TaskStatus("new")
	// The task is waiting for a payment to be validated.
	TaskStatusWaitingPayment = TaskStatus("waiting_payment")
	// The task is pending.
	TaskStatusPending = TaskStatus("pending")
	// The task has been completed.
	TaskStatusSuccess = TaskStatus("success")
	// The task is in an error state.
	TaskStatusError = TaskStatus("error")
)

func (enum TaskStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unavailable"
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

type TaskType string

const (
	// If unspecified, the status is unknown by default.
	TaskTypeUnknown = TaskType("unknown")
	// Create a new internal domain.
	TaskTypeCreateDomain = TaskType("create_domain")
	// Create a new external domain.
	TaskTypeCreateExternalDomain = TaskType("create_external_domain")
	// Renew a domain.
	TaskTypeRenewDomain = TaskType("renew_domain")
	// Transfer a domain to Scaleway Domains and DNS.
	TaskTypeTransferDomain = TaskType("transfer_domain")
	// Trade a domain to a new owner.
	TaskTypeTradeDomain = TaskType("trade_domain")
	// Lock the transfer of a domain for protection.
	TaskTypeLockDomainTransfer = TaskType("lock_domain_transfer")
	// Unlock the transfer of a domain.
	TaskTypeUnlockDomainTransfer = TaskType("unlock_domain_transfer")
	// Enable DNSSEC for a domain.
	TaskTypeEnableDnssec = TaskType("enable_dnssec")
	// Disable DNSSEC for a domain.
	TaskTypeDisableDnssec = TaskType("disable_dnssec")
	// Update the domain's information.
	TaskTypeUpdateDomain = TaskType("update_domain")
	// Change the technical or administrative contact.
	TaskTypeUpdateContact = TaskType("update_contact")
	// Delete a domain and destroy its zone versions, zones, and SSL certificates.
	TaskTypeDeleteDomain = TaskType("delete_domain")
	// Cancel a task that has not yet started.
	TaskTypeCancelTask = TaskType("cancel_task")
	// Generate a new SSL certificate.
	TaskTypeGenerateSslCertificate = TaskType("generate_ssl_certificate")
	// Renew an SSL certificate.
	TaskTypeRenewSslCertificate = TaskType("renew_ssl_certificate")
	// Send a message. For most cases, it will be followed by an email.
	TaskTypeSendMessage = TaskType("send_message")
	// Delete a domain that has expired and not been restored for at least 3 months.
	TaskTypeDeleteDomainExpired = TaskType("delete_domain_expired")
	// Delete a newly registered external domain that has not been validated after 48 hours or when the external domain fails to point to our name servers for more than 14 days.
	TaskTypeDeleteExternalDomain = TaskType("delete_external_domain")
	// Create domain's hostname with glue IPs.
	TaskTypeCreateHost = TaskType("create_host")
	// Update domain's hostname with glue IPs.
	TaskTypeUpdateHost = TaskType("update_host")
	// Delete domain's hostname.
	TaskTypeDeleteHost = TaskType("delete_host")
)

func (enum TaskType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum TaskType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TaskType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TaskType(TaskType(tmp).String())
	return nil
}

// RecordGeoIPConfigMatch:
type RecordGeoIPConfigMatch struct {
	// Countries:
	Countries []string `json:"countries"`
	// Continents:
	Continents []string `json:"continents"`
	// Data:
	Data string `json:"data"`
}

// RecordViewConfigView:
type RecordViewConfigView struct {
	// Subnet:
	Subnet string `json:"subnet"`
	// Data:
	Data string `json:"data"`
}

// RecordWeightedConfigWeightedIP:
type RecordWeightedConfigWeightedIP struct {
	// IP:
	IP net.IP `json:"ip"`
	// Weight:
	Weight uint32 `json:"weight"`
}

// DSRecordPublicKey:
type DSRecordPublicKey struct {
	// Key:
	Key string `json:"key"`
}

// RecordGeoIPConfig:
type RecordGeoIPConfig struct {
	// Matches:
	Matches []*RecordGeoIPConfigMatch `json:"matches"`
	// Default:
	Default string `json:"default"`
}

// RecordHTTPServiceConfig:
type RecordHTTPServiceConfig struct {
	// IPs:
	IPs []net.IP `json:"ips"`
	// MustContain:
	MustContain *string `json:"must_contain,omitempty"`
	// URL:
	URL string `json:"url"`
	// UserAgent:
	UserAgent *string `json:"user_agent,omitempty"`
	// Strategy:
	Strategy RecordHTTPServiceConfigStrategy `json:"strategy"`
}

// RecordViewConfig:
type RecordViewConfig struct {
	// Views:
	Views []*RecordViewConfigView `json:"views"`
}

// RecordWeightedConfig:
type RecordWeightedConfig struct {
	// WeightedIPs:
	WeightedIPs []*RecordWeightedConfigWeightedIP `json:"weighted_ips"`
}

// ContactExtensionFRAssociationInfo:
type ContactExtensionFRAssociationInfo struct {
	// PublicationJo:
	PublicationJo *time.Time `json:"publication_jo,omitempty"`
	// PublicationJoPage:
	PublicationJoPage uint32 `json:"publication_jo_page"`
}

// ContactExtensionFRCodeAuthAfnicInfo:
type ContactExtensionFRCodeAuthAfnicInfo struct {
	// CodeAuthAfnic:
	CodeAuthAfnic string `json:"code_auth_afnic"`
}

// ContactExtensionFRDunsInfo:
type ContactExtensionFRDunsInfo struct {
	// DunsID:
	DunsID string `json:"duns_id"`
	// LocalID:
	LocalID string `json:"local_id"`
}

// ContactExtensionFRIndividualInfo:
type ContactExtensionFRIndividualInfo struct {
	// WhoisOptIn:
	WhoisOptIn bool `json:"whois_opt_in"`
}

// ContactExtensionFRTrademarkInfo:
type ContactExtensionFRTrademarkInfo struct {
	// TrademarkInpi:
	TrademarkInpi string `json:"trademark_inpi"`
}

// DSRecordDigest:
type DSRecordDigest struct {
	// Type:
	Type DSRecordDigestType `json:"type"`
	// Digest:
	Digest string `json:"digest"`
	// PublicKey:
	PublicKey *DSRecordPublicKey `json:"public_key"`
}

// Record:
type Record struct {
	// Data:
	Data string `json:"data"`
	// Name:
	Name string `json:"name"`
	// Priority:
	Priority uint32 `json:"priority"`
	// TTL:
	TTL uint32 `json:"ttl"`
	// Type:
	Type RecordType `json:"type"`
	// Comment:
	Comment *string `json:"comment,omitempty"`
	// GeoIPConfig:
	GeoIPConfig *RecordGeoIPConfig `json:"geo_ip_config,omitempty"`
	// HTTPServiceConfig:
	HTTPServiceConfig *RecordHTTPServiceConfig `json:"http_service_config,omitempty"`
	// WeightedConfig:
	WeightedConfig *RecordWeightedConfig `json:"weighted_config,omitempty"`
	// ViewConfig:
	ViewConfig *RecordViewConfig `json:"view_config,omitempty"`
	// ID:
	ID string `json:"id"`
}

// RecordIdentifier:
type RecordIdentifier struct {
	// Name:
	Name string `json:"name"`
	// Type:
	Type RecordType `json:"type"`
	// Data:
	Data *string `json:"data,omitempty"`
	// TTL:
	TTL *uint32 `json:"ttl,omitempty"`
}

// ContactExtensionEU:
type ContactExtensionEU struct {
	// EuropeanCitizenship:
	EuropeanCitizenship string `json:"european_citizenship"`
}

// ContactExtensionFR:
type ContactExtensionFR struct {
	// Mode:
	Mode ContactExtensionFRMode `json:"mode"`
	// IndividualInfo:
	IndividualInfo *ContactExtensionFRIndividualInfo `json:"individual_info,omitempty"`
	// DunsInfo:
	DunsInfo *ContactExtensionFRDunsInfo `json:"duns_info,omitempty"`
	// AssociationInfo:
	AssociationInfo *ContactExtensionFRAssociationInfo `json:"association_info,omitempty"`
	// TrademarkInfo:
	TrademarkInfo *ContactExtensionFRTrademarkInfo `json:"trademark_info,omitempty"`
	// CodeAuthAfnicInfo:
	CodeAuthAfnicInfo *ContactExtensionFRCodeAuthAfnicInfo `json:"code_auth_afnic_info,omitempty"`
}

// ContactExtensionNL:
type ContactExtensionNL struct {
	// LegalForm:
	LegalForm ContactExtensionNLLegalForm `json:"legal_form"`
	// LegalFormRegistrationNumber:
	LegalFormRegistrationNumber string `json:"legal_form_registration_number"`
}

// ContactQuestion:
type ContactQuestion struct {
	// Question:
	Question string `json:"question"`
	// Answer:
	Answer string `json:"answer"`
}

// TldOffer:
type TldOffer struct {
	// Action:
	Action string `json:"action"`
	// OperationPath:
	OperationPath string `json:"operation_path"`
	// Price:
	Price *scw.Money `json:"price,omitempty"`
}

// DSRecord:
type DSRecord struct {
	// KeyID:
	KeyID uint32 `json:"key_id"`
	// Algorithm:
	Algorithm DSRecordAlgorithm `json:"algorithm"`
	// Digest:
	Digest *DSRecordDigest `json:"digest,omitempty"`
	// PublicKey:
	PublicKey *DSRecordPublicKey `json:"public_key,omitempty"`
}

// RecordChangeAdd:
type RecordChangeAdd struct {
	// Records:
	Records []*Record `json:"records"`
}

// RecordChangeClear:
type RecordChangeClear struct {
}

// RecordChangeDelete:
type RecordChangeDelete struct {
	// ID:
	ID *string `json:"id,omitempty"`
	// IDFields:
	IDFields *RecordIdentifier `json:"id_fields,omitempty"`
}

// RecordChangeSet:
type RecordChangeSet struct {
	// ID:
	ID *string `json:"id,omitempty"`
	// IDFields:
	IDFields *RecordIdentifier `json:"id_fields,omitempty"`
	// Records:
	Records []*Record `json:"records"`
}

// ImportRawDNSZoneRequestTsigKey:
type ImportRawDNSZoneRequestTsigKey struct {
	// Name:
	Name string `json:"name"`
	// Key:
	Key string `json:"key"`
	// Algorithm:
	Algorithm string `json:"algorithm"`
}

// Contact:
type Contact struct {
	// ID:
	ID string `json:"id"`
	// LegalForm:
	LegalForm ContactLegalForm `json:"legal_form"`
	// Firstname:
	Firstname string `json:"firstname"`
	// Lastname:
	Lastname string `json:"lastname"`
	// CompanyName:
	CompanyName string `json:"company_name"`
	// Email:
	Email string `json:"email"`
	// EmailAlt:
	EmailAlt string `json:"email_alt"`
	// PhoneNumber:
	PhoneNumber string `json:"phone_number"`
	// FaxNumber:
	FaxNumber string `json:"fax_number"`
	// AddressLine1:
	AddressLine1 string `json:"address_line_1"`
	// AddressLine2:
	AddressLine2 string `json:"address_line_2"`
	// Zip:
	Zip string `json:"zip"`
	// City:
	City string `json:"city"`
	// Country:
	Country string `json:"country"`
	// VatIDentificationCode:
	VatIDentificationCode string `json:"vat_identification_code"`
	// CompanyIDentificationCode:
	CompanyIDentificationCode string `json:"company_identification_code"`
	// Lang:
	Lang std.LanguageCode `json:"lang"`
	// Resale:
	Resale bool `json:"resale"`
	// Deprecated: Questions:
	Questions *[]*ContactQuestion `json:"questions,omitempty"`
	// ExtensionFr:
	ExtensionFr *ContactExtensionFR `json:"extension_fr"`
	// ExtensionEu:
	ExtensionEu *ContactExtensionEU `json:"extension_eu"`
	// WhoisOptIn:
	WhoisOptIn bool `json:"whois_opt_in"`
	// EmailStatus:
	EmailStatus ContactEmailStatus `json:"email_status"`
	// State:
	State string `json:"state"`
	// ExtensionNl:
	ExtensionNl *ContactExtensionNL `json:"extension_nl"`
}

// ContactRolesRoles:
type ContactRolesRoles struct {
	// IsOwner:
	IsOwner bool `json:"is_owner"`
	// IsAdministrative:
	IsAdministrative bool `json:"is_administrative"`
	// IsTechnical:
	IsTechnical bool `json:"is_technical"`
}

// DomainRegistrationStatusExternalDomain:
type DomainRegistrationStatusExternalDomain struct {
	// ValidationToken:
	ValidationToken string `json:"validation_token"`
}

// DomainRegistrationStatusTransfer:
type DomainRegistrationStatusTransfer struct {
	// Status:
	Status DomainRegistrationStatusTransferStatus `json:"status"`
	// VoteCurrentOwner:
	VoteCurrentOwner bool `json:"vote_current_owner"`
	// VoteNewOwner:
	VoteNewOwner bool `json:"vote_new_owner"`
}

// Tld:
type Tld struct {
	// Name:
	Name string `json:"name"`
	// DnssecSupport:
	DnssecSupport bool `json:"dnssec_support"`
	// DurationInYearsMin:
	DurationInYearsMin uint32 `json:"duration_in_years_min"`
	// DurationInYearsMax:
	DurationInYearsMax uint32 `json:"duration_in_years_max"`
	// IDnSupport:
	IDnSupport bool `json:"idn_support"`
	// Offers:
	Offers map[string]*TldOffer `json:"offers"`
	// Specifications:
	Specifications map[string]string `json:"specifications"`
}

// NewContact:
type NewContact struct {
	// LegalForm:
	LegalForm ContactLegalForm `json:"legal_form"`
	// Firstname:
	Firstname string `json:"firstname"`
	// Lastname:
	Lastname string `json:"lastname"`
	// CompanyName:
	CompanyName *string `json:"company_name,omitempty"`
	// Email:
	Email string `json:"email"`
	// EmailAlt:
	EmailAlt *string `json:"email_alt,omitempty"`
	// PhoneNumber:
	PhoneNumber string `json:"phone_number"`
	// FaxNumber:
	FaxNumber *string `json:"fax_number,omitempty"`
	// AddressLine1:
	AddressLine1 string `json:"address_line_1"`
	// AddressLine2:
	AddressLine2 *string `json:"address_line_2,omitempty"`
	// Zip:
	Zip string `json:"zip"`
	// City:
	City string `json:"city"`
	// Country:
	Country string `json:"country"`
	// VatIDentificationCode:
	VatIDentificationCode *string `json:"vat_identification_code,omitempty"`
	// CompanyIDentificationCode:
	CompanyIDentificationCode *string `json:"company_identification_code,omitempty"`
	// Lang:
	Lang std.LanguageCode `json:"lang"`
	// Resale:
	Resale bool `json:"resale"`
	// Deprecated: Questions:
	Questions *[]*ContactQuestion `json:"questions,omitempty"`
	// ExtensionFr:
	ExtensionFr *ContactExtensionFR `json:"extension_fr"`
	// ExtensionEu:
	ExtensionEu *ContactExtensionEU `json:"extension_eu"`
	// WhoisOptIn:
	WhoisOptIn bool `json:"whois_opt_in"`
	// State:
	State *string `json:"state,omitempty"`
	// ExtensionNl:
	ExtensionNl *ContactExtensionNL `json:"extension_nl"`
}

// CheckContactsCompatibilityResponseContactCheckResult:
type CheckContactsCompatibilityResponseContactCheckResult struct {
	// Compatible:
	Compatible bool `json:"compatible"`
	// ErrorMessage:
	ErrorMessage *string `json:"error_message,omitempty"`
}

// DNSZone:
type DNSZone struct {
	// Domain:
	Domain string `json:"domain"`
	// Subdomain:
	Subdomain string `json:"subdomain"`
	// Ns:
	Ns []string `json:"ns"`
	// NsDefault:
	NsDefault []string `json:"ns_default"`
	// NsMaster:
	NsMaster []string `json:"ns_master"`
	// Status:
	Status DNSZoneStatus `json:"status"`
	// Message:
	Message *string `json:"message,omitempty"`
	// UpdatedAt:
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// ProjectID:
	ProjectID string `json:"project_id"`
}

// DomainDNSSEC:
type DomainDNSSEC struct {
	// Status:
	Status DomainFeatureStatus `json:"status"`
	// DsRecords:
	DsRecords []*DSRecord `json:"ds_records"`
}

// RecordChange:
type RecordChange struct {
	// Add:
	Add *RecordChangeAdd `json:"add,omitempty"`
	// Set:
	Set *RecordChangeSet `json:"set,omitempty"`
	// Delete:
	Delete *RecordChangeDelete `json:"delete,omitempty"`
	// Clear:
	Clear *RecordChangeClear `json:"clear,omitempty"`
}

// ImportProviderDNSZoneRequestOnlineV1:
type ImportProviderDNSZoneRequestOnlineV1 struct {
	// Token:
	Token string `json:"token"`
}

// ImportRawDNSZoneRequestAXFRSource:
type ImportRawDNSZoneRequestAXFRSource struct {
	// NameServer:
	NameServer string `json:"name_server"`
	// TsigKey:
	TsigKey *ImportRawDNSZoneRequestTsigKey `json:"tsig_key"`
}

// ImportRawDNSZoneRequestBindSource:
type ImportRawDNSZoneRequestBindSource struct {
	// Content:
	Content string `json:"content"`
}

// ContactRoles:
type ContactRoles struct {
	// Contact:
	Contact *Contact `json:"contact"`
	// Roles:
	Roles map[string]*ContactRolesRoles `json:"roles"`
}

// Nameserver:
type Nameserver struct {
	// Name:
	Name string `json:"name"`
	// IP:
	IP []string `json:"ip"`
}

// DNSZoneVersion:
type DNSZoneVersion struct {
	// ID:
	ID string `json:"id"`
	// CreatedAt:
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// Host:
type Host struct {
	// Domain:
	Domain string `json:"domain"`
	// Name:
	Name string `json:"name"`
	// IPs:
	IPs []net.IP `json:"ips"`
	// Status:
	Status HostStatus `json:"status"`
}

// DomainSummary:
type DomainSummary struct {
	// Domain:
	Domain string `json:"domain"`
	// ProjectID:
	ProjectID string `json:"project_id"`
	// AutoRenewStatus:
	AutoRenewStatus DomainFeatureStatus `json:"auto_renew_status"`
	// DnssecStatus:
	DnssecStatus DomainFeatureStatus `json:"dnssec_status"`
	// EppCode:
	EppCode []string `json:"epp_code"`
	// ExpiredAt:
	ExpiredAt *time.Time `json:"expired_at,omitempty"`
	// UpdatedAt:
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Registrar:
	Registrar string `json:"registrar"`
	// IsExternal:
	IsExternal bool `json:"is_external"`
	// Status:
	Status DomainStatus `json:"status"`
	// ExternalDomainRegistrationStatus:
	ExternalDomainRegistrationStatus *DomainRegistrationStatusExternalDomain `json:"external_domain_registration_status,omitempty"`
	// TransferRegistrationStatus:
	TransferRegistrationStatus *DomainRegistrationStatusTransfer `json:"transfer_registration_status,omitempty"`
	// OrganizationID:
	OrganizationID string `json:"organization_id"`
}

// RenewableDomain:
type RenewableDomain struct {
	// Domain:
	Domain string `json:"domain"`
	// ProjectID:
	ProjectID string `json:"project_id"`
	// OrganizationID:
	OrganizationID string `json:"organization_id"`
	// Status:
	Status RenewableDomainStatus `json:"status"`
	// RenewableDurationInYears:
	RenewableDurationInYears *int32 `json:"renewable_duration_in_years,omitempty"`
	// ExpiredAt:
	ExpiredAt *time.Time `json:"expired_at,omitempty"`
	// LimitRenewAt:
	LimitRenewAt *time.Time `json:"limit_renew_at,omitempty"`
	// LimitRedemptionAt:
	LimitRedemptionAt *time.Time `json:"limit_redemption_at,omitempty"`
	// EstimatedDeleteAt:
	EstimatedDeleteAt *time.Time `json:"estimated_delete_at,omitempty"`
	// Tld:
	Tld *Tld `json:"tld"`
}

// SSLCertificate:
type SSLCertificate struct {
	// DNSZone:
	DNSZone string `json:"dns_zone"`
	// AlternativeDNSZones:
	AlternativeDNSZones []string `json:"alternative_dns_zones"`
	// Status:
	Status SSLCertificateStatus `json:"status"`
	// PrivateKey:
	PrivateKey string `json:"private_key"`
	// CertificateChain:
	CertificateChain string `json:"certificate_chain"`
	// CreatedAt:
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// ExpiredAt:
	ExpiredAt *time.Time `json:"expired_at,omitempty"`
}

// Task:
type Task struct {
	// ID:
	ID string `json:"id"`
	// ProjectID:
	ProjectID string `json:"project_id"`
	// OrganizationID:
	OrganizationID string `json:"organization_id"`
	// Domain:
	Domain *string `json:"domain,omitempty"`
	// Type:
	Type TaskType `json:"type"`
	// Status:
	Status TaskStatus `json:"status"`
	// StartedAt:
	StartedAt *time.Time `json:"started_at,omitempty"`
	// UpdatedAt:
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Message:
	Message *string `json:"message,omitempty"`
}

// TransferInDomainRequestTransferRequest:
type TransferInDomainRequestTransferRequest struct {
	// Domain:
	Domain string `json:"domain"`
	// AuthCode:
	AuthCode string `json:"auth_code"`
}

// UpdateContactRequestQuestion:
type UpdateContactRequestQuestion struct {
	// Question:
	Question *string `json:"question,omitempty"`
	// Answer:
	Answer *string `json:"answer,omitempty"`
}

// AvailableDomain:
type AvailableDomain struct {
	// Domain:
	Domain string `json:"domain"`
	// Available:
	Available bool `json:"available"`
	// Tld:
	Tld *Tld `json:"tld"`
}

// CheckContactsCompatibilityResponse:
type CheckContactsCompatibilityResponse struct {
	// Compatible:
	Compatible bool `json:"compatible"`
	// OwnerCheckResult:
	OwnerCheckResult *CheckContactsCompatibilityResponseContactCheckResult `json:"owner_check_result"`
	// AdministrativeCheckResult:
	AdministrativeCheckResult *CheckContactsCompatibilityResponseContactCheckResult `json:"administrative_check_result"`
	// TechnicalCheckResult:
	TechnicalCheckResult *CheckContactsCompatibilityResponseContactCheckResult `json:"technical_check_result"`
}

// ClearDNSZoneRecordsRequest:
type ClearDNSZoneRecordsRequest struct {
	// DNSZone: DNS zone to clear.
	DNSZone string `json:"-"`
}

// ClearDNSZoneRecordsResponse:
type ClearDNSZoneRecordsResponse struct {
}

// CloneDNSZoneRequest:
type CloneDNSZoneRequest struct {
	// DNSZone: DNS zone to clone.
	DNSZone string `json:"-"`
	// DestDNSZone: Destination DNS zone in which to clone the chosen DNS zone.
	DestDNSZone string `json:"dest_dns_zone"`
	// Overwrite: Specifies whether or not the destination DNS zone will be overwritten.
	Overwrite bool `json:"overwrite"`
	// ProjectID: Project ID of the destination DNS zone.
	ProjectID *string `json:"project_id,omitempty"`
}

// CreateDNSZoneRequest:
type CreateDNSZoneRequest struct {
	// Domain: Domain in which to crreate the DNS zone.
	Domain string `json:"domain"`
	// Subdomain: Subdomain of the DNS zone to create.
	Subdomain string `json:"subdomain"`
	// ProjectID: Project ID in which to create the DNS zone.
	ProjectID string `json:"project_id"`
}

// CreateSSLCertificateRequest:
type CreateSSLCertificateRequest struct {
	// DNSZone:
	DNSZone string `json:"dns_zone"`
	// AlternativeDNSZones:
	AlternativeDNSZones []string `json:"alternative_dns_zones"`
}

// DeleteDNSZoneRequest:
type DeleteDNSZoneRequest struct {
	// DNSZone: DNS zone to delete.
	DNSZone string `json:"-"`
	// ProjectID: Project ID of the DNS zone to delete.
	ProjectID string `json:"-"`
}

// DeleteDNSZoneResponse:
type DeleteDNSZoneResponse struct {
}

// DeleteDNSZoneTsigKeyRequest:
type DeleteDNSZoneTsigKeyRequest struct {
	// DNSZone:
	DNSZone string `json:"-"`
}

// DeleteExternalDomainResponse:
type DeleteExternalDomainResponse struct {
}

// DeleteSSLCertificateRequest:
type DeleteSSLCertificateRequest struct {
	// DNSZone:
	DNSZone string `json:"-"`
}

// DeleteSSLCertificateResponse:
type DeleteSSLCertificateResponse struct {
}

// Domain:
type Domain struct {
	// Domain:
	Domain string `json:"domain"`
	// OrganizationID:
	OrganizationID string `json:"organization_id"`
	// ProjectID:
	ProjectID string `json:"project_id"`
	// AutoRenewStatus:
	AutoRenewStatus DomainFeatureStatus `json:"auto_renew_status"`
	// Dnssec:
	Dnssec *DomainDNSSEC `json:"dnssec"`
	// EppCode:
	EppCode []string `json:"epp_code"`
	// ExpiredAt:
	ExpiredAt *time.Time `json:"expired_at,omitempty"`
	// UpdatedAt:
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Registrar:
	Registrar string `json:"registrar"`
	// IsExternal:
	IsExternal bool `json:"is_external"`
	// Status:
	Status DomainStatus `json:"status"`
	// DNSZones:
	DNSZones []*DNSZone `json:"dns_zones"`
	// OwnerContact:
	OwnerContact *Contact `json:"owner_contact"`
	// TechnicalContact:
	TechnicalContact *Contact `json:"technical_contact"`
	// AdministrativeContact:
	AdministrativeContact *Contact `json:"administrative_contact"`
	// ExternalDomainRegistrationStatus:
	ExternalDomainRegistrationStatus *DomainRegistrationStatusExternalDomain `json:"external_domain_registration_status,omitempty"`
	// TransferRegistrationStatus:
	TransferRegistrationStatus *DomainRegistrationStatusTransfer `json:"transfer_registration_status,omitempty"`
	// Tld:
	Tld *Tld `json:"tld"`
}

// ExportRawDNSZoneRequest:
type ExportRawDNSZoneRequest struct {
	// DNSZone: DNS zone to export.
	DNSZone string `json:"-"`
	// Format: DNS zone format.
	Format RawFormat `json:"-"`
}

// GetDNSZoneTsigKeyRequest:
type GetDNSZoneTsigKeyRequest struct {
	// DNSZone:
	DNSZone string `json:"-"`
}

// GetDNSZoneTsigKeyResponse:
type GetDNSZoneTsigKeyResponse struct {
	// Name:
	Name string `json:"name"`
	// Key:
	Key string `json:"key"`
	// Algorithm:
	Algorithm string `json:"algorithm"`
}

// GetDNSZoneVersionDiffRequest:
type GetDNSZoneVersionDiffRequest struct {
	// DNSZoneVersionID:
	DNSZoneVersionID string `json:"-"`
}

// GetDNSZoneVersionDiffResponse:
type GetDNSZoneVersionDiffResponse struct {
	// Changes:
	Changes []*RecordChange `json:"changes"`
}

// GetDomainAuthCodeResponse:
type GetDomainAuthCodeResponse struct {
	// AuthCode:
	AuthCode string `json:"auth_code"`
}

// GetSSLCertificateRequest:
type GetSSLCertificateRequest struct {
	// DNSZone:
	DNSZone string `json:"-"`
}

// ImportProviderDNSZoneRequest:
type ImportProviderDNSZoneRequest struct {
	// DNSZone:
	DNSZone string `json:"-"`
	// OnlineV1:
	OnlineV1 *ImportProviderDNSZoneRequestOnlineV1 `json:"online_v1,omitempty"`
}

// ImportProviderDNSZoneResponse:
type ImportProviderDNSZoneResponse struct {
	// Records:
	Records []*Record `json:"records"`
}

// ImportRawDNSZoneRequest:
type ImportRawDNSZoneRequest struct {
	// DNSZone: DNS zone to import.
	DNSZone string `json:"-"`
	// Deprecated: Content:
	Content *string `json:"content,omitempty"`
	// ProjectID:
	ProjectID string `json:"project_id"`
	// Deprecated: Format:
	Format *RawFormat `json:"format,omitempty"`
	// BindSource: Import a bind file format.
	BindSource *ImportRawDNSZoneRequestBindSource `json:"bind_source,omitempty"`
	// AxfrSource: Import from the name server given with TSIG, to use or not.
	AxfrSource *ImportRawDNSZoneRequestAXFRSource `json:"axfr_source,omitempty"`
}

// ImportRawDNSZoneResponse:
type ImportRawDNSZoneResponse struct {
	// Records:
	Records []*Record `json:"records"`
}

// ListContactsResponse:
type ListContactsResponse struct {
	// TotalCount:
	TotalCount uint32 `json:"total_count"`
	// Contacts:
	Contacts []*ContactRoles `json:"contacts"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListContactsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListContactsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListContactsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Contacts = append(r.Contacts, results.Contacts...)
	r.TotalCount += uint32(len(results.Contacts))
	return uint32(len(results.Contacts)), nil
}

// ListDNSZoneNameserversRequest:
type ListDNSZoneNameserversRequest struct {
	// DNSZone: DNS zone on which to filter the returned DNS zone name servers.
	DNSZone string `json:"-"`
	// ProjectID: Project ID on which to filter the returned DNS zone name servers.
	ProjectID *string `json:"-"`
}

// ListDNSZoneNameserversResponse:
type ListDNSZoneNameserversResponse struct {
	// Ns: DNS zone name servers returned.
	Ns []*Nameserver `json:"ns"`
}

// ListDNSZoneRecordsRequest:
type ListDNSZoneRecordsRequest struct {
	// DNSZone: DNS zone on which to filter the returned DNS zone records.
	DNSZone string `json:"-"`
	// ProjectID: Project ID on which to filter the returned DNS zone records.
	ProjectID *string `json:"-"`
	// OrderBy: Sort order of the returned DNS zone records.
	OrderBy ListDNSZoneRecordsRequestOrderBy `json:"-"`
	// Page: Page number to return, from the paginated results.
	Page *int32 `json:"-"`
	// PageSize: Maximum number of DNS zone records per page.
	PageSize *uint32 `json:"-"`
	// Name: Name on which to filter the returned DNS zone records.
	Name string `json:"-"`
	// Type: Record type on which to filter the returned DNS zone records.
	Type RecordType `json:"-"`
	// ID: Record ID on which to filter the returned DNS zone records.
	ID *string `json:"-"`
}

// ListDNSZoneRecordsResponse:
type ListDNSZoneRecordsResponse struct {
	// TotalCount: Total number of DNS zone records.
	TotalCount uint32 `json:"total_count"`
	// Records: Paginated returned DNS zone records.
	Records []*Record `json:"records"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDNSZoneRecordsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDNSZoneRecordsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDNSZoneRecordsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Records = append(r.Records, results.Records...)
	r.TotalCount += uint32(len(results.Records))
	return uint32(len(results.Records)), nil
}

// ListDNSZoneVersionRecordsRequest:
type ListDNSZoneVersionRecordsRequest struct {
	// DNSZoneVersionID:
	DNSZoneVersionID string `json:"-"`
	// Page: Page number to return, from the paginated results.
	Page *int32 `json:"-"`
	// PageSize: Maximum number of DNS zones versions records per page.
	PageSize *uint32 `json:"-"`
}

// ListDNSZoneVersionRecordsResponse:
type ListDNSZoneVersionRecordsResponse struct {
	// TotalCount: Total number of DNS zones versions records.
	TotalCount uint32 `json:"total_count"`
	// Records:
	Records []*Record `json:"records"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDNSZoneVersionRecordsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDNSZoneVersionRecordsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDNSZoneVersionRecordsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Records = append(r.Records, results.Records...)
	r.TotalCount += uint32(len(results.Records))
	return uint32(len(results.Records)), nil
}

// ListDNSZoneVersionsRequest:
type ListDNSZoneVersionsRequest struct {
	// DNSZone:
	DNSZone string `json:"-"`
	// Page: Page number to return, from the paginated results.
	Page *int32 `json:"-"`
	// PageSize: Maximum number of DNS zones versions per page.
	PageSize *uint32 `json:"-"`
}

// ListDNSZoneVersionsResponse:
type ListDNSZoneVersionsResponse struct {
	// TotalCount: Total number of DNS zones versions.
	TotalCount uint32 `json:"total_count"`
	// Versions:
	Versions []*DNSZoneVersion `json:"versions"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDNSZoneVersionsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDNSZoneVersionsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDNSZoneVersionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Versions = append(r.Versions, results.Versions...)
	r.TotalCount += uint32(len(results.Versions))
	return uint32(len(results.Versions)), nil
}

// ListDNSZonesRequest:
type ListDNSZonesRequest struct {
	// OrganizationID: Organization ID on which to filter the returned DNS zones.
	OrganizationID *string `json:"-"`
	// ProjectID: Project ID on which to filter the returned DNS zones.
	ProjectID *string `json:"-"`
	// OrderBy: Sort order of the returned DNS zones.
	OrderBy ListDNSZonesRequestOrderBy `json:"-"`
	// Page: Page number to return, from the paginated results.
	Page *int32 `json:"-"`
	// PageSize: Maximum number of DNS zones to return per page.
	PageSize *uint32 `json:"-"`
	// Domain: Domain on which to filter the returned DNS zones.
	Domain string `json:"-"`
	// DNSZone: DNS zone on which to filter the returned DNS zones.
	DNSZone string `json:"-"`
}

// ListDNSZonesResponse:
type ListDNSZonesResponse struct {
	// TotalCount: Total number of DNS zones matching the requested criteria.
	TotalCount uint32 `json:"total_count"`
	// DNSZones: Paginated returned DNS zones.
	DNSZones []*DNSZone `json:"dns_zones"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDNSZonesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDNSZonesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDNSZonesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.DNSZones = append(r.DNSZones, results.DNSZones...)
	r.TotalCount += uint32(len(results.DNSZones))
	return uint32(len(results.DNSZones)), nil
}

// ListDomainHostsResponse:
type ListDomainHostsResponse struct {
	// TotalCount:
	TotalCount uint32 `json:"total_count"`
	// Hosts:
	Hosts []*Host `json:"hosts"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDomainHostsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDomainHostsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDomainHostsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Hosts = append(r.Hosts, results.Hosts...)
	r.TotalCount += uint32(len(results.Hosts))
	return uint32(len(results.Hosts)), nil
}

// ListDomainsResponse:
type ListDomainsResponse struct {
	// TotalCount:
	TotalCount uint32 `json:"total_count"`
	// Domains:
	Domains []*DomainSummary `json:"domains"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDomainsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDomainsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDomainsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Domains = append(r.Domains, results.Domains...)
	r.TotalCount += uint32(len(results.Domains))
	return uint32(len(results.Domains)), nil
}

// ListRenewableDomainsResponse:
type ListRenewableDomainsResponse struct {
	// TotalCount:
	TotalCount uint32 `json:"total_count"`
	// Domains:
	Domains []*RenewableDomain `json:"domains"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRenewableDomainsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRenewableDomainsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListRenewableDomainsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Domains = append(r.Domains, results.Domains...)
	r.TotalCount += uint32(len(results.Domains))
	return uint32(len(results.Domains)), nil
}

// ListSSLCertificatesRequest:
type ListSSLCertificatesRequest struct {
	// DNSZone:
	DNSZone string `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
	// ProjectID:
	ProjectID *string `json:"-"`
}

// ListSSLCertificatesResponse:
type ListSSLCertificatesResponse struct {
	// TotalCount:
	TotalCount uint32 `json:"total_count"`
	// Certificates:
	Certificates []*SSLCertificate `json:"certificates"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSSLCertificatesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSSLCertificatesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSSLCertificatesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Certificates = append(r.Certificates, results.Certificates...)
	r.TotalCount += uint32(len(results.Certificates))
	return uint32(len(results.Certificates)), nil
}

// ListTasksResponse:
type ListTasksResponse struct {
	// TotalCount:
	TotalCount uint32 `json:"total_count"`
	// Tasks:
	Tasks []*Task `json:"tasks"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTasksResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTasksResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListTasksResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Tasks = append(r.Tasks, results.Tasks...)
	r.TotalCount += uint32(len(results.Tasks))
	return uint32(len(results.Tasks)), nil
}

// OrderResponse:
type OrderResponse struct {
	// Domains:
	Domains []string `json:"domains"`
	// OrganizationID:
	OrganizationID string `json:"organization_id"`
	// ProjectID:
	ProjectID string `json:"project_id"`
	// TaskID:
	TaskID string `json:"task_id"`
	// CreatedAt:
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// RefreshDNSZoneRequest:
type RefreshDNSZoneRequest struct {
	// DNSZone: DNS zone to refresh.
	DNSZone string `json:"-"`
	// RecreateDNSZone: Specifies whether or not to recreate the DNS zone.
	RecreateDNSZone bool `json:"recreate_dns_zone"`
	// RecreateSubDNSZone: Specifies whether or not to recreate the sub DNS zone.
	RecreateSubDNSZone bool `json:"recreate_sub_dns_zone"`
}

// RefreshDNSZoneResponse:
type RefreshDNSZoneResponse struct {
	// DNSZones: DNS zones returned.
	DNSZones []*DNSZone `json:"dns_zones"`
}

// RegisterExternalDomainResponse:
type RegisterExternalDomainResponse struct {
	// Domain:
	Domain string `json:"domain"`
	// OrganizationID:
	OrganizationID string `json:"organization_id"`
	// ValidationToken:
	ValidationToken string `json:"validation_token"`
	// CreatedAt:
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// ProjectID:
	ProjectID string `json:"project_id"`
}

// RegistrarAPIBuyDomainsRequest:
type RegistrarAPIBuyDomainsRequest struct {
	// Domains:
	Domains []string `json:"domains"`
	// DurationInYears:
	DurationInYears uint32 `json:"duration_in_years"`
	// ProjectID:
	ProjectID string `json:"project_id"`
	// OwnerContactID:
	OwnerContactID *string `json:"owner_contact_id,omitempty"`
	// OwnerContact:
	OwnerContact *NewContact `json:"owner_contact,omitempty"`
	// AdministrativeContactID:
	AdministrativeContactID *string `json:"administrative_contact_id,omitempty"`
	// AdministrativeContact:
	AdministrativeContact *NewContact `json:"administrative_contact,omitempty"`
	// TechnicalContactID:
	TechnicalContactID *string `json:"technical_contact_id,omitempty"`
	// TechnicalContact:
	TechnicalContact *NewContact `json:"technical_contact,omitempty"`
}

// RegistrarAPICheckContactsCompatibilityRequest:
type RegistrarAPICheckContactsCompatibilityRequest struct {
	// Domains:
	Domains []string `json:"domains"`
	// Tlds:
	Tlds []string `json:"tlds"`
	// OwnerContactID:
	OwnerContactID *string `json:"owner_contact_id,omitempty"`
	// OwnerContact:
	OwnerContact *NewContact `json:"owner_contact,omitempty"`
	// AdministrativeContactID:
	AdministrativeContactID *string `json:"administrative_contact_id,omitempty"`
	// AdministrativeContact:
	AdministrativeContact *NewContact `json:"administrative_contact,omitempty"`
	// TechnicalContactID:
	TechnicalContactID *string `json:"technical_contact_id,omitempty"`
	// TechnicalContact:
	TechnicalContact *NewContact `json:"technical_contact,omitempty"`
}

// RegistrarAPICreateDomainHostRequest:
type RegistrarAPICreateDomainHostRequest struct {
	// Domain:
	Domain string `json:"-"`
	// Name:
	Name string `json:"name"`
	// IPs:
	IPs []net.IP `json:"ips"`
}

// RegistrarAPIDeleteDomainHostRequest:
type RegistrarAPIDeleteDomainHostRequest struct {
	// Domain:
	Domain string `json:"-"`
	// Name:
	Name string `json:"-"`
}

// RegistrarAPIDeleteExternalDomainRequest:
type RegistrarAPIDeleteExternalDomainRequest struct {
	// Domain:
	Domain string `json:"-"`
}

// RegistrarAPIDisableDomainAutoRenewRequest:
type RegistrarAPIDisableDomainAutoRenewRequest struct {
	// Domain:
	Domain string `json:"-"`
}

// RegistrarAPIDisableDomainDNSSECRequest:
type RegistrarAPIDisableDomainDNSSECRequest struct {
	// Domain:
	Domain string `json:"-"`
}

// RegistrarAPIEnableDomainAutoRenewRequest:
type RegistrarAPIEnableDomainAutoRenewRequest struct {
	// Domain:
	Domain string `json:"-"`
}

// RegistrarAPIEnableDomainDNSSECRequest:
type RegistrarAPIEnableDomainDNSSECRequest struct {
	// Domain:
	Domain string `json:"-"`
	// DsRecord:
	DsRecord *DSRecord `json:"ds_record"`
}

// RegistrarAPIGetContactRequest:
type RegistrarAPIGetContactRequest struct {
	// ContactID:
	ContactID string `json:"-"`
}

// RegistrarAPIGetDomainAuthCodeRequest:
type RegistrarAPIGetDomainAuthCodeRequest struct {
	// Domain:
	Domain string `json:"-"`
}

// RegistrarAPIGetDomainRequest:
type RegistrarAPIGetDomainRequest struct {
	// Domain:
	Domain string `json:"-"`
}

// RegistrarAPIListContactsRequest:
type RegistrarAPIListContactsRequest struct {
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
	// Domain:
	Domain *string `json:"-"`
	// ProjectID:
	ProjectID *string `json:"-"`
	// OrganizationID:
	OrganizationID *string `json:"-"`
	// Role:
	Role ListContactsRequestRole `json:"-"`
	// EmailStatus:
	EmailStatus ContactEmailStatus `json:"-"`
}

// RegistrarAPIListDomainHostsRequest:
type RegistrarAPIListDomainHostsRequest struct {
	// Domain:
	Domain string `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
}

// RegistrarAPIListDomainsRequest:
type RegistrarAPIListDomainsRequest struct {
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
	// OrderBy:
	OrderBy ListDomainsRequestOrderBy `json:"-"`
	// Registrar:
	Registrar *string `json:"-"`
	// Status:
	Status DomainStatus `json:"-"`
	// ProjectID:
	ProjectID *string `json:"-"`
	// OrganizationID:
	OrganizationID *string `json:"-"`
	// IsExternal:
	IsExternal *bool `json:"-"`
	// Domain:
	Domain *string `json:"-"`
}

// RegistrarAPIListRenewableDomainsRequest:
type RegistrarAPIListRenewableDomainsRequest struct {
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
	// OrderBy:
	OrderBy ListRenewableDomainsRequestOrderBy `json:"-"`
	// ProjectID:
	ProjectID *string `json:"-"`
	// OrganizationID:
	OrganizationID *string `json:"-"`
}

// RegistrarAPIListTasksRequest:
type RegistrarAPIListTasksRequest struct {
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
	// ProjectID:
	ProjectID *string `json:"-"`
	// OrganizationID:
	OrganizationID *string `json:"-"`
	// Domain:
	Domain *string `json:"-"`
	// Types:
	Types []TaskType `json:"-"`
	// Statuses:
	Statuses []TaskStatus `json:"-"`
	// OrderBy:
	OrderBy ListTasksRequestOrderBy `json:"-"`
}

// RegistrarAPILockDomainTransferRequest:
type RegistrarAPILockDomainTransferRequest struct {
	// Domain:
	Domain string `json:"-"`
}

// RegistrarAPIRegisterExternalDomainRequest:
type RegistrarAPIRegisterExternalDomainRequest struct {
	// Domain:
	Domain string `json:"domain"`
	// ProjectID:
	ProjectID string `json:"project_id"`
}

// RegistrarAPIRenewDomainsRequest:
type RegistrarAPIRenewDomainsRequest struct {
	// Domains:
	Domains []string `json:"domains"`
	// DurationInYears:
	DurationInYears uint32 `json:"duration_in_years"`
	// ForceLateRenewal:
	ForceLateRenewal *bool `json:"force_late_renewal,omitempty"`
}

// RegistrarAPISearchAvailableDomainsRequest:
type RegistrarAPISearchAvailableDomainsRequest struct {
	// Domains: A list of domain to search, TLD is optional.
	Domains []string `json:"-"`
	// Tlds: Array of tlds to search on.
	Tlds []string `json:"-"`
	// StrictSearch: Search exact match.
	StrictSearch bool `json:"-"`
}

// RegistrarAPITradeDomainRequest:
type RegistrarAPITradeDomainRequest struct {
	// Domain:
	Domain string `json:"-"`
	// ProjectID:
	ProjectID *string `json:"project_id,omitempty"`
	// NewOwnerContactID:
	NewOwnerContactID *string `json:"new_owner_contact_id,omitempty"`
	// NewOwnerContact:
	NewOwnerContact *NewContact `json:"new_owner_contact,omitempty"`
}

// RegistrarAPITransferInDomainRequest:
type RegistrarAPITransferInDomainRequest struct {
	// Domains:
	Domains []*TransferInDomainRequestTransferRequest `json:"domains"`
	// ProjectID:
	ProjectID string `json:"project_id"`
	// OwnerContactID:
	OwnerContactID *string `json:"owner_contact_id,omitempty"`
	// OwnerContact:
	OwnerContact *NewContact `json:"owner_contact,omitempty"`
	// AdministrativeContactID:
	AdministrativeContactID *string `json:"administrative_contact_id,omitempty"`
	// AdministrativeContact:
	AdministrativeContact *NewContact `json:"administrative_contact,omitempty"`
	// TechnicalContactID:
	TechnicalContactID *string `json:"technical_contact_id,omitempty"`
	// TechnicalContact:
	TechnicalContact *NewContact `json:"technical_contact,omitempty"`
}

// RegistrarAPIUnlockDomainTransferRequest:
type RegistrarAPIUnlockDomainTransferRequest struct {
	// Domain:
	Domain string `json:"-"`
}

// RegistrarAPIUpdateContactRequest:
type RegistrarAPIUpdateContactRequest struct {
	// ContactID:
	ContactID string `json:"-"`
	// Email:
	Email *string `json:"email,omitempty"`
	// EmailAlt:
	EmailAlt *string `json:"email_alt,omitempty"`
	// PhoneNumber:
	PhoneNumber *string `json:"phone_number,omitempty"`
	// FaxNumber:
	FaxNumber *string `json:"fax_number,omitempty"`
	// AddressLine1:
	AddressLine1 *string `json:"address_line_1,omitempty"`
	// AddressLine2:
	AddressLine2 *string `json:"address_line_2,omitempty"`
	// Zip:
	Zip *string `json:"zip,omitempty"`
	// City:
	City *string `json:"city,omitempty"`
	// Country:
	Country *string `json:"country,omitempty"`
	// VatIDentificationCode:
	VatIDentificationCode *string `json:"vat_identification_code,omitempty"`
	// CompanyIDentificationCode:
	CompanyIDentificationCode *string `json:"company_identification_code,omitempty"`
	// Lang:
	Lang std.LanguageCode `json:"lang"`
	// Resale:
	Resale *bool `json:"resale,omitempty"`
	// Deprecated: Questions:
	Questions *[]*UpdateContactRequestQuestion `json:"questions,omitempty"`
	// ExtensionFr:
	ExtensionFr *ContactExtensionFR `json:"extension_fr"`
	// ExtensionEu:
	ExtensionEu *ContactExtensionEU `json:"extension_eu"`
	// WhoisOptIn:
	WhoisOptIn *bool `json:"whois_opt_in,omitempty"`
	// State:
	State *string `json:"state,omitempty"`
	// ExtensionNl:
	ExtensionNl *ContactExtensionNL `json:"extension_nl"`
}

// RegistrarAPIUpdateDomainHostRequest:
type RegistrarAPIUpdateDomainHostRequest struct {
	// Domain:
	Domain string `json:"-"`
	// Name:
	Name string `json:"-"`
	// IPs:
	IPs *[]string `json:"ips,omitempty"`
}

// RegistrarAPIUpdateDomainRequest:
type RegistrarAPIUpdateDomainRequest struct {
	// Domain:
	Domain string `json:"-"`
	// TechnicalContactID:
	TechnicalContactID *string `json:"technical_contact_id,omitempty"`
	// TechnicalContact:
	TechnicalContact *NewContact `json:"technical_contact,omitempty"`
	// OwnerContactID:
	OwnerContactID *string `json:"owner_contact_id,omitempty"`
	// OwnerContact:
	OwnerContact *NewContact `json:"owner_contact,omitempty"`
	// AdministrativeContactID:
	AdministrativeContactID *string `json:"administrative_contact_id,omitempty"`
	// AdministrativeContact:
	AdministrativeContact *NewContact `json:"administrative_contact,omitempty"`
}

// RestoreDNSZoneVersionRequest:
type RestoreDNSZoneVersionRequest struct {
	// DNSZoneVersionID:
	DNSZoneVersionID string `json:"-"`
}

// RestoreDNSZoneVersionResponse:
type RestoreDNSZoneVersionResponse struct {
}

// SearchAvailableDomainsResponse:
type SearchAvailableDomainsResponse struct {
	// AvailableDomains: Array of available domains.
	AvailableDomains []*AvailableDomain `json:"available_domains"`
}

// UpdateDNSZoneNameserversRequest:
type UpdateDNSZoneNameserversRequest struct {
	// DNSZone: DNS zone in which to update the DNS zone name servers.
	DNSZone string `json:"-"`
	// Ns: New DNS zone name servers.
	Ns []*Nameserver `json:"ns"`
}

// UpdateDNSZoneNameserversResponse:
type UpdateDNSZoneNameserversResponse struct {
	// Ns: DNS zone name servers returned.
	Ns []*Nameserver `json:"ns"`
}

// UpdateDNSZoneRecordsRequest:
type UpdateDNSZoneRecordsRequest struct {
	// DNSZone: DNS zone in which to update the DNS zone records.
	DNSZone string `json:"-"`
	// Changes: Changes made to the records.
	Changes []*RecordChange `json:"changes"`
	// ReturnAllRecords: Specifies whether or not to return all the records.
	ReturnAllRecords *bool `json:"return_all_records,omitempty"`
	// DisallowNewZoneCreation: Disable the creation of the target zone if it does not exist. Target zone creation is disabled by default.
	DisallowNewZoneCreation bool `json:"disallow_new_zone_creation"`
	// Serial: Use the provided serial (0) instead of the auto-increment serial.
	Serial *uint64 `json:"serial,omitempty"`
}

// UpdateDNSZoneRecordsResponse:
type UpdateDNSZoneRecordsResponse struct {
	// Records: DNS zone records returned.
	Records []*Record `json:"records"`
}

// UpdateDNSZoneRequest:
type UpdateDNSZoneRequest struct {
	// DNSZone: DNS zone to update.
	DNSZone string `json:"-"`
	// NewDNSZone: Name of the new DNS zone to create.
	NewDNSZone *string `json:"new_dns_zone,omitempty"`
	// ProjectID: Project ID in which to create the new DNS zone.
	ProjectID string `json:"project_id"`
}

// The Domains and DNS API documentation allows you to configure and manage your domains' DNS zones and records. You can also use dynamic records to optimize and easily use your infrastructure as code.
// Read our [reference content documentation](https://www.scaleway.com/en/docs/network/domains-and-dns/reference-content/) for more information about domains.
//
// Refer to the [Domains and DNS Registrar API documentation](https://www.scaleway.com/en/developers/api/domains-and-dns/registrar-api/) to find out how to buy,transfer and manage your domains and contacts (DNSSEC included).
//
// (switchcolumn)
// (switchcolumn)
// ## Concepts
//
// Refer to our [dedicated concepts page](https://www.scaleway.com/en/docs/network/domains-and-dns/concepts/) to find definitions of the different terms referring to Scaleway Domains and DNS.
//
// (switchcolumn)
// (switchcolumn)
// ## Quickstart
//
// 1. **Configure your environment variables.**
//
//	<Message type="note">
//	This is an optional step that seeks to simplify your usage of the API.
//	</Message>
//
//	```bash
//	export SCW_ACCESS_KEY="<API access key>"
//	export SCW_SECRET_KEY="<API secret key>"
//	export SCW_PROJECT_ID="<Scaleway Project ID>"
//	```
//
// 2. **Register an external domain.**
//
//	Run the following command to register your domain:
//
//	```bash
//	curl https://api.scaleway.com/domain/v2beta1/external-domains \
//	  -X POST \
//	  -H "Content-Type: application/json" \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  -d '{
//	    "domain": "my-external-domain.test",
//	    "project_id": "<Scaleway Project ID>"
//	}'
//	```
//
// 3. **Set up a TXT record.**
//
//	<Message type="important">
//	  After you register your external domain, you must set up a TXT record **within 48 hours**. This step is optional if you have an internal domain.
//	</Message>
//
//	Set up your TXT record as explained in **step 4** of our [documentation on how to add an external domain](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/add-external-domain/) to confirm you are the owner of the domain registered.
//
//	<Message type="note">
//	  You will receive a confirmation email once your domain is validated. Your domain will then be available on the API.
//	</Message>
//
// 4. **Update your DNS name servers.**
//
//	[Update your DNS name servers](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-nameservers-internal-domain) to Scaleway's DNS name servers: `ns0.dom.scw.cloud` and `ns1.dom.scw.cloud`.
//
//	- For your domain: update the name servers at your registrar to point your domain to Scaleway's DNS servers
//	- For your subdomain: update the name servers at your parent DNS server to point your subdomain to Scaleway's DNS servers
//
// 5. **Set up your domain's DNSSEC.**
//
//	You need to set up your domain's DNSSEC on the registry by your registrar and on the DNS server with the same information.
//
//	To set up the DNSSEC:
//	* use the [Registrar API](https://www.scaleway.com/en/developers/api/domains-and-dns/registrar-api/#path-domains-update-domain-dnssec) `EnableDomainDNSSEC` to generate the setup on the DNS servers
//	* if your domain is not registered on Scaleway but on another registrar, use the DS or public key information to configure the DNSSEC on your registrar
//
//	<Message type="tip">
//	  You can remove your DNSSEC settings on your registrar before you migrate to our service. You will be able to set it up again afterwards.
//	</Message>
//
// 6. **List your domains.**
//
//	Run the following command to retrieve the list of your domains:
//
//	```bash
//	curl https://api.scaleway.com/domain/v2beta1/domains \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY"
//	```
//
//	The following output displays:
//
//	```bash
//	{
//	  "total_count": 1,
//	  "domains": [
//	    {
//	      "domain": "domain-external.test",
//	      "organization_id": "<Scaleway Organization ID>",
//	      "project_id": "<Scaleway Project ID>",
//	      "auto_renew_status": "disabled",
//	      "dnssec_status": "disabled",
//	      "epp_code": [],
//	      "expired_at": null,
//	      "updated_at": "1970-01-01T00:00:00Z",
//	      "registrar": "EXTERNAL",
//	      "is_external": true,
//	      "status": "active"
//	    }
//	  ]
//	}
//	```
//
// 7. **List the zones you can manage.**
//
//	Run the following command to list the DNS zones you can manage:
//
//	```bash
//	curl https://api.scaleway.com/domain/v2beta1/dns-zones/ \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY"
//	```
//	An output similar to the following displays:
//
//	```bash
//	{
//	  "total_count": 3,
//	  "dns_zones": [
//	    {
//	      "domain": "domain-external.test",
//	      "subdomain": "",
//	      "ns": [
//	        "ns0.dom.scw.cloud",
//	        "ns1.dom.scw.cloud"
//	      ],
//	      "ns_default": [
//	        "ns0.dom.scw.cloud",
//	        "ns1.dom.scw.cloud"
//	      ],
//	      "ns_master": [],
//	      "status": "active",
//	      "message": null,
//	      "updated_at": "2022-06-21T01:34:55Z",
//	      "project_id": "<Scaleway Project ID>"
//	    },
//	    {
//	      "domain": "domain-external.test",
//	      "subdomain": "test-zone",
//	      "ns": [
//	        "ns0.dom.scw.cloud",
//	        "ns1.dom.scw.cloud"
//	      ],
//	      "ns_default": [
//	        "ns0.dom.scw.cloud",
//	        "ns1.dom.scw.cloud"
//	      ],
//	      "ns_master": [],
//	      "status": "active",
//	      "message": null,
//	      "updated_at": "2022-06-21T01:34:55Z",
//	      "project_id": "<Scaleway Project ID>"
//	    },
//	    {
//	      "domain": "other-domain-external.test",
//	      "subdomain": "",
//	      "ns": [
//	        "ns0.dom.scw.cloud",
//	        "ns1.dom.scw.cloud"
//	      ],
//	      "ns_default": [
//	        "ns0.dom.scw.cloud",
//	        "ns1.dom.scw.cloud"
//	      ],
//	      "ns_master": [],
//	      "status": "active",
//	      "message": null,
//	      "updated_at": "2022-06-21T01:34:55Z",
//	      "project_id": "<Scaleway Project ID>"
//	    }
//	  ]
//	}
//
//	```
//
// 8. **Update a zone's record(s).**
//
//	<Message type="important">
//	  - The `clear` command clears all your records. You will have to setup your records again
//	  - If the zone you want to update does not exist but is included in a managed zone, the API will create the new (sub) zone
//	</Message>
//
//	Run the following command to update a zone's record:
//
//	```bash
//	curl https://api.scaleway.com/domain/v2beta1/dns-zones/subdomain.domain-external.test/records \
//	  -X PATCH \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  -d '{
//	    "changes": [
//	      {
//	        "clear": {}
//	      },
//	      {
//	        "add": {
//	          "records": [
//	            {
//	              "name": "test1",
//	              "data": "127.0.0.1",
//	              "type": "A",
//	              "ttl": 3600
//	            },
//	            {
//	              "name": "test2",
//	              "data": "127.0.0.1",
//	              "type": "A",
//	              "ttl": 3600
//	            }
//	          ]
//	        }
//	      },
//	      {
//	        "set": {
//	          "id_fields": {
//	            "type": "A",
//	            "name": "test2"
//	          },
//	          "records": [
//	            {
//	              "name": "test2",
//	              "data": "127.0.0.2",
//	              "type": "A",
//	              "ttl": 3600
//	            }
//	          ]
//	        }
//	      },
//	      {
//	        "add": {
//	          "records": [
//	            {
//	              "name": "test3",
//	              "data": "127.0.0.3",
//	              "type": "A",
//	              "ttl": 3600
//	            },
//	            {
//	              "name": "test3",
//	              "data": "127.0.0.4",
//	              "type": "A",
//	              "ttl": 3600
//	            }
//	          ]
//	        }
//	      },
//	      {
//	        "delete": {
//	          "id_fields": {
//	            "name": "test3",
//	            "data": "127.0.0.4",
//	            "type": "A",
//	            "ttl": 3600
//	          }
//	        }
//	      }
//	    ]
//	  }'
//	```
//
//	An output similar to the following displays:
//
//	```bash
//	{
//	  "records": [
//	    {
//	      "id": "<UUID>",
//	      "data": "127.0.0.1",
//	      "name": "test1",
//	      "priority": 0,
//	      "ttl": 3600,
//	      "type": "A"
//	    },
//	    {
//	      "id": "<UUID>",
//	      "data": "127.0.0.1",
//	      "name": "test2",
//	      "priority": 0,
//	      "ttl": 3600,
//	      "type": "A"
//	    },
//	    {
//	      "id": "<UUID>",
//	      "data": "127.0.0.2",
//	      "name": "test2",
//	      "priority": 0,
//	      "ttl": 3600,
//	      "type": "A"
//	    },
//	    {
//	      "id": "<UUID>",
//	      "data": "127.0.0.3",
//	      "name": "test3",
//	      "priority": 0,
//	      "ttl": 3600,
//	      "type": "A"
//	    }
//	  ]
//	}
//	```
//
// 9. **Enable DNSSEC for external domains.**
//
//   - Run the following command if you have not set up DNSSEC at your registrar:
//
//     ```bash
//     curl https://api.scaleway.com/domain/v2beta1/domains/example.com/enable-dnssec \
//     -X POST \
//     -H "Content-Type: application/json" \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     -d '{}'
//     ```
//     An output similar to the following displays:
//
//     ```json
//     {
//     "domain":{
//     "domain":"example.fr",
//     "auto_renew_status":"enabled",
//     "dnssec_status":"disabled",
//     "ds_records":[
//     {
//     "key_id":49071,
//     "algorithm":"ecdsap256sha256",
//     "digest":{
//     "type":"sha_384",
//     "digest":"c14421f739e35b62e3383d5972263f93610dea31da48a1990453752cf1b4f4f0314f75ecd0bc9061c51ee41ec7692204"
//     }
//     }
//     ],
//     "epp":[
//     "clientTransferProhibited"
//     ],
//     "expired_at":"2021-04-04T12:46:29Z",
//     "updated_at":"2020-07-07T13:24:27Z",
//     "registrar":"EXTERNAL",
//     "is_external":true,
//     "status":"active",
//     "......"
//     }
//     ```
//
//     <Message type="important">
//     Set up the information returned under `ds_records` in the output, in your registrar to configure DNSSEC for your external domain.
//     </Message>
//
//   - Run the following command if you have already set up DNSSEC at your registrar:
//
//     ```bash
//     curl https://api.scaleway.com/domain/v2beta1/domains/example.com/enable-dnssec \
//     -X POST \
//     -H "Content-Type: application/json" \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     -d '{
//     "ds_record": {
//     "algorithm": "rsamd5",
//     "digest":
//     {
//     "type": "sha_1",
//     "digest": "ABCD"
//     }
//     }
//     }'
//     ```
//
// 10. **Enable DNSSEC for internal domains.**
//
//	Find out how to enable DNSSEC for your internal domains in our [dedicated documentation](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/configure-dnssec-internal-domain/).
//
// (switchcolumn)
// <Message type="requirement">
// - You have a [Scaleway account](https://console.scaleway.com/)
// - You have [created an API key](https://www.scaleway.com/en/docs/identity-and-access-management/iam/how-to/create-api-keys/) and that the API key has sufficient [IAM permissions](https://www.scaleway.com/en/docs/identity-and-access-management/iam/reference-content/permission-sets/) to perform the actions described on this page
// - You have your [Organization ID](https://console.scaleway.com/organization/settings)
// - You have [installed `curl`](https://curl.se/download.html)
// </Message>
// (switchcolumn)
//
// ## Technical information
//
// ### Regional availability
//
// Scaleway Domains and DNS is currently available globally. Find out about [product availability in our dedicated documentation](https://www.scaleway.com/en/docs/console/my-account/reference-content/products-availability).
//
// ### Format
//
// - The format of record names is always in "short" format (i.e."test" instead of the FQDN "test.mydomain.com")
// - If a data value is not in RFC format, quotes will be added for `TXT`records. If `CNAME`records are not in FQDDN, the zone will be added
// ## Technical limitations
//
// - Only the owner of a domain can use and configure it with Domains and DNS
// - Each zone supports the creation of sub-zones, which inherit the parent's Project ID
// - We provide the following record types: `A`, `AAAA`, `CNAME`, `CAA`, `DNAME`, `MX`, `NAPTR`, `NS`, `SRV`, `TLSA` and `TXT`
// - We support the following dynamic records: `http(s)` `health check` and `views`
//
// ## Going futher
//
// For more information about Domains and DNS, you can check out the following pages:
//
// - [Domains and DNS Documentation](https://www.scaleway.com/en/docs/network/domains-and-dns/)
// - [Scaleway Slack Community](http://slack.scaleway.com) join the #domains channel
// - [Contact our support team](https://console.scaleway.com/support/tickets)
//
// ### Test if the service is up
//
// Run the following command to test if an `http` or `https` service is up on a provided list of IPs:
//
// ```bash
//
//	curl https://api.scaleway.com/domain/v2beta1/dns-zones/dom.example.com/records \
//	  -X PATCH \
//	  -H "Content-Type: application/json" \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  -d '{
//	    "return_all_records": false,
//	    "changes": [
//	      {
//	        "add": {
//	          "records": [
//	            {
//	              "name": "service-www",
//	              "data": "10.51.2.8",
//	              "type": "A",
//	              "ttl": 600,
//	              "http_service_config": {
//	                "ips": [
//	                  "10.51.2.8",
//	                  "10.51.3.7",
//	                  "10.51.2.9"
//	                ],
//	                "must_contain": "I am fine",
//	                "url": "https://www.domain-test.test"
//	              }
//	            }
//	          ]
//	        }
//	      }
//	    ]
//	  }'
//
// ```
//
// <Message type="note">
//
//	An IP is considered healthy if the query response contains the text specified in `must_contain`. The record type can be one of `A` and `AAAA`.
//
// </Message>
//
// ### View
//
// Run the following command to view responses depending on the requester's or resolver's IP:
//
// ```bash
//
//	curl https://api.scaleway.com/domain/v2beta1/dns-zones/dom.example.com/records \
//	  -X PATCH \
//	  -H "Content-Type: application/json" \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  -d '{
//	    "return_all_records": false,
//	    "changes": [
//	      {
//	        "set": {
//	          "id_fields": {
//	            "name": "test",
//	            "type": "CNAME"
//	          },
//	          "records": [
//	            {
//	              "name": "test",
//	              "data": "default-domain.com.",
//	              "type": "CNAME",
//	              "ttl": 600,
//	              "view_config": {
//	                "views": [
//	                  {
//	                    "subnet": "10.51.0.0/16",
//	                    "data": "filtered-domain.com."
//	                  }
//	                ]
//	              }
//	            }
//	          ]
//	        }
//	      }
//	    ]
//	  }'
//
// ```
// <Message type="note">
//
//	Record types include `A`, `AAAA`, `CNAME` and `TXT`.
//
// </Message>
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// ListDNSZones: Retrieve the list of DNS zones you can manage and filter DNS zones associated with specific domain names.
func (s *API) ListDNSZones(req *ListDNSZonesRequest, opts ...scw.RequestOption) (*ListDNSZonesResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "domain", req.Domain)
	parameter.AddToQuery(query, "dns_zone", req.DNSZone)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain/v2beta1/dns-zones",
		Query:  query,
	}

	var resp ListDNSZonesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDNSZone: Create a new DNS zone specified by the domain name, the subdomain and the Project ID.
func (s *API) CreateDNSZone(req *CreateDNSZoneRequest, opts ...scw.RequestOption) (*DNSZone, error) {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain/v2beta1/dns-zones",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DNSZone

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateDNSZone: Update the name and/or the Organizations for a DNS zone.
func (s *API) UpdateDNSZone(req *UpdateDNSZoneRequest, opts ...scw.RequestOption) (*DNSZone, error) {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/domain/v2beta1/dns-zones/" + fmt.Sprint(req.DNSZone) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DNSZone

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CloneDNSZone: Clone an existing DNS zone with all its records into a new DNS zone.
func (s *API) CloneDNSZone(req *CloneDNSZoneRequest, opts ...scw.RequestOption) (*DNSZone, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain/v2beta1/dns-zones/" + fmt.Sprint(req.DNSZone) + "/clone",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DNSZone

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDNSZone: Delete a DNS zone and all its records.
func (s *API) DeleteDNSZone(req *DeleteDNSZoneRequest, opts ...scw.RequestOption) (*DeleteDNSZoneResponse, error) {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/domain/v2beta1/dns-zones/" + fmt.Sprint(req.DNSZone) + "",
		Query:  query,
	}

	var resp DeleteDNSZoneResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDNSZoneRecords: Retrieve a list of DNS records within a DNS zone that has default name servers.
// You can filter records by type and name.
func (s *API) ListDNSZoneRecords(req *ListDNSZoneRecordsRequest, opts ...scw.RequestOption) (*ListDNSZoneRecordsResponse, error) {
	var err error
	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "type", req.Type)
	parameter.AddToQuery(query, "id", req.ID)

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain/v2beta1/dns-zones/" + fmt.Sprint(req.DNSZone) + "/records",
		Query:  query,
	}

	var resp ListDNSZoneRecordsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateDNSZoneRecords: Update records within a DNS zone that has default name servers and perform several actions on your records.
//
// Actions include:
//   - add: allows you to add a new record or add a new IP to an existing A record, for example
//   - set: allows you to edit a record or edit an IP from an existing A record, for example
//   - delete: allows you to delete a record or delete an IP from an existing A record, for example
//   - clear: allows you to delete all records from a DNS zone
//
// All edits will be versioned.
func (s *API) UpdateDNSZoneRecords(req *UpdateDNSZoneRecordsRequest, opts ...scw.RequestOption) (*UpdateDNSZoneRecordsResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/domain/v2beta1/dns-zones/" + fmt.Sprint(req.DNSZone) + "/records",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UpdateDNSZoneRecordsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDNSZoneNameservers: Retrieve a list of name servers within a DNS zone and their optional glue records.
func (s *API) ListDNSZoneNameservers(req *ListDNSZoneNameserversRequest, opts ...scw.RequestOption) (*ListDNSZoneNameserversResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain/v2beta1/dns-zones/" + fmt.Sprint(req.DNSZone) + "/nameservers",
		Query:  query,
	}

	var resp ListDNSZoneNameserversResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateDNSZoneNameservers: Update name servers within a DNS zone and set optional glue records.
func (s *API) UpdateDNSZoneNameservers(req *UpdateDNSZoneNameserversRequest, opts ...scw.RequestOption) (*UpdateDNSZoneNameserversResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/domain/v2beta1/dns-zones/" + fmt.Sprint(req.DNSZone) + "/nameservers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UpdateDNSZoneNameserversResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ClearDNSZoneRecords: Delete all records within a DNS zone that has default name servers.<br/>
// All edits will be versioned.
func (s *API) ClearDNSZoneRecords(req *ClearDNSZoneRecordsRequest, opts ...scw.RequestOption) (*ClearDNSZoneRecordsResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/domain/v2beta1/dns-zones/" + fmt.Sprint(req.DNSZone) + "/records",
	}

	var resp ClearDNSZoneRecordsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExportRawDNSZone: Export a DNS zone with default name servers, in a specific format.
func (s *API) ExportRawDNSZone(req *ExportRawDNSZoneRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "format", req.Format)

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain/v2beta1/dns-zones/" + fmt.Sprint(req.DNSZone) + "/raw",
		Query:  query,
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ImportRawDNSZone: Import and replace the format of records from a given provider, with default name servers.
func (s *API) ImportRawDNSZone(req *ImportRawDNSZoneRequest, opts ...scw.RequestOption) (*ImportRawDNSZoneResponse, error) {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain/v2beta1/dns-zones/" + fmt.Sprint(req.DNSZone) + "/raw",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ImportRawDNSZoneResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ImportProviderDNSZone: Import and replace the format of records from a given provider, with default name servers.
func (s *API) ImportProviderDNSZone(req *ImportProviderDNSZoneRequest, opts ...scw.RequestOption) (*ImportProviderDNSZoneResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain/v2beta1/dns-zones/" + fmt.Sprint(req.DNSZone) + "/import-provider",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ImportProviderDNSZoneResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RefreshDNSZone: Refresh an SOA DNS zone to reload the records in the DNS zone and update the SOA serial.
// You can recreate the given DNS zone and its sub DNS zone if needed.
func (s *API) RefreshDNSZone(req *RefreshDNSZoneRequest, opts ...scw.RequestOption) (*RefreshDNSZoneResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain/v2beta1/dns-zones/" + fmt.Sprint(req.DNSZone) + "/refresh",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RefreshDNSZoneResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDNSZoneVersions: Retrieve a list of a DNS zone's versions.<br/>
// The maximum version count is 100. If the count reaches this limit, the oldest version will be deleted after each new modification.
func (s *API) ListDNSZoneVersions(req *ListDNSZoneVersionsRequest, opts ...scw.RequestOption) (*ListDNSZoneVersionsResponse, error) {
	var err error
	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain/v2beta1/dns-zones/" + fmt.Sprint(req.DNSZone) + "/versions",
		Query:  query,
	}

	var resp ListDNSZoneVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDNSZoneVersionRecords: Retrieve a list of records from a specific DNS zone version.
func (s *API) ListDNSZoneVersionRecords(req *ListDNSZoneVersionRecordsRequest, opts ...scw.RequestOption) (*ListDNSZoneVersionRecordsResponse, error) {
	var err error
	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.DNSZoneVersionID) == "" {
		return nil, errors.New("field DNSZoneVersionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain/v2beta1/dns-zones/version/" + fmt.Sprint(req.DNSZoneVersionID) + "",
		Query:  query,
	}

	var resp ListDNSZoneVersionRecordsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDNSZoneVersionDiff: Access a previous DNS zone version to see the differences from another specific version.
func (s *API) GetDNSZoneVersionDiff(req *GetDNSZoneVersionDiffRequest, opts ...scw.RequestOption) (*GetDNSZoneVersionDiffResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZoneVersionID) == "" {
		return nil, errors.New("field DNSZoneVersionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain/v2beta1/dns-zones/version/" + fmt.Sprint(req.DNSZoneVersionID) + "/diff",
	}

	var resp GetDNSZoneVersionDiffResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RestoreDNSZoneVersion: Restore and activate a version of a specific DNS zone.
func (s *API) RestoreDNSZoneVersion(req *RestoreDNSZoneVersionRequest, opts ...scw.RequestOption) (*RestoreDNSZoneVersionResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZoneVersionID) == "" {
		return nil, errors.New("field DNSZoneVersionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain/v2beta1/dns-zones/version/" + fmt.Sprint(req.DNSZoneVersionID) + "/restore",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RestoreDNSZoneVersionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSSLCertificate: Get the DNS zone's TLS certificate. If you do not have a certificate, the ouptut returns `no certificate found`.
func (s *API) GetSSLCertificate(req *GetSSLCertificateRequest, opts ...scw.RequestOption) (*SSLCertificate, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain/v2beta1/ssl-certificates/" + fmt.Sprint(req.DNSZone) + "",
	}

	var resp SSLCertificate

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSSLCertificate: Create a new TLS certificate or retrieve information about an existing TLS certificate.
func (s *API) CreateSSLCertificate(req *CreateSSLCertificateRequest, opts ...scw.RequestOption) (*SSLCertificate, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain/v2beta1/ssl-certificates",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SSLCertificate

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSSLCertificates: List all the TLS certificates a user has created, specified by the user's Project ID and the DNS zone.
func (s *API) ListSSLCertificates(req *ListSSLCertificatesRequest, opts ...scw.RequestOption) (*ListSSLCertificatesResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "dns_zone", req.DNSZone)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain/v2beta1/ssl-certificates",
		Query:  query,
	}

	var resp ListSSLCertificatesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSSLCertificate: Delete an existing TLS certificate specified by its DNS zone. Deleting a TLS certificate is permanent and cannot be undone.
func (s *API) DeleteSSLCertificate(req *DeleteSSLCertificateRequest, opts ...scw.RequestOption) (*DeleteSSLCertificateResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/domain/v2beta1/ssl-certificates/" + fmt.Sprint(req.DNSZone) + "",
	}

	var resp DeleteSSLCertificateResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDNSZoneTsigKey: Retrieve information about the TSIG key of a given DNS zone to allow AXFR requests.
func (s *API) GetDNSZoneTsigKey(req *GetDNSZoneTsigKeyRequest, opts ...scw.RequestOption) (*GetDNSZoneTsigKeyResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain/v2beta1/dns-zones/" + fmt.Sprint(req.DNSZone) + "/tsig-key",
	}

	var resp GetDNSZoneTsigKeyResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDNSZoneTsigKey: Delete an existing TSIG key specified by its DNS zone. Deleting a TSIG key is permanent and cannot be undone.
func (s *API) DeleteDNSZoneTsigKey(req *DeleteDNSZoneTsigKeyRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/domain/v2beta1/dns-zones/" + fmt.Sprint(req.DNSZone) + "/tsig-key",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// The Domains and DNS registrar API documentation allows you to buy, transfer and manage your domains and contacts (DNSSEC included) as a simple client or a reseller.
//
// Read our [reference content documentation for more information about domains](https://www.scaleway.com/en/docs/network/domains-and-dns/reference-content/).
//
// (switchcolumn)
// (switchcolumn)
// ## Concepts
//
// Refer to our [dedicated concepts page](https://www.scaleway.com/en/docs/network/domains-and-dns/concepts/) to find definitions of the different terms referring to Scaleway Domains and DNS.
//
// (switchcolumn)
// (switchcolumn)
// ## Quickstart
//
// 1. **Configure your environment variables.**
//
//	<Message type="note">
//	This is an optional step that seeks to simplify your usage of the API.
//	</Message>
//
//	```bash
//	export SCW_ACCESS_KEY="<API access key>"
//	export SCW_SECRET_KEY="<API secret key>"
//	export SCW_PROJECT_ID="<Scaleway Project ID>"
//	```
//
// 2. **Get information about your domains.**
//
//		<Message type="important">
//	      You must specify the domain you want to get information for if you want to get information about one specific domain.
//	    </Message>
//
//		Run the following command to get information about your domains:
//
//		```bash
//		curl https://api.scaleway.com/domain/v2beta1/search-domains?domains=domain-example.fr \
//		-H "X-Auth-Token: $secret_key"
//		```
//
//		An output similar to the following displays:
//
//		```bash
//		{
//			"available_domains": [
//				{
//					"domain": "domain-example.fr",
//					"available": true,
//					"tld": {
//						"name": "fr",
//						"dnssec_support": true,
//						"duration_in_years_min": 1,
//						"duration_in_years_max": 10,
//						"idn_support": true,
//						"offers": {
//							"create": {
//								"action": "create",
//								"operation_path": "/network/domain/domains/fr/register",
//								"price": {
//									"currency_code": "EUR",
//									"units": 123,
//									"nanos": 456000000
//								}
//							},
//							"late_renew": {
//								"action": "late_renew",
//								"operation_path": "/network/domain/domains/fr/late_renew",
//								"price": {
//									"currency_code": "EUR",
//									"units": 123,
//									"nanos": 456000000
//								}
//							},
//							"renew": {
//								"action": "renew",
//								"operation_path": "/network/domain/domains/fr/renew",
//								"price": {
//									"currency_code": "EUR",
//									"units": 123,
//									"nanos": 456000000
//								}
//							},
//							"trade": {
//								"action": "trade",
//								"operation_path": "/network/domain/domains/fr/trade",
//								"price": {
//									"currency_code": "EUR",
//									"units": 123,
//									"nanos": 456000000
//								}
//							},
//							"transfer": {
//								"action": "transfer",
//								"operation_path": "/network/domain/domains/fr/transfer",
//								"price": {
//									"currency_code": "EUR",
//									"units": 123,
//									"nanos": 456000000
//								}
//							}
//						},
//						"specifications": {
//							"afnic": "1",
//							"grace_period": "1",
//							"transfer_without_contact": "1"
//						}
//					}
//				}
//			]
//		}
//		```
//		<Message type="note">
//
//		 - `available` indicates wether the domain is available to be ordered and created at the registry
//		 - `tld` indicates wether TLD is supported, the different actions and prices available. If TLD is not supported, `null` is returned in the output
//		 - `offers` indicates the different actions and prices available
//		 - `specifications` includes tags to indicate TLD specifications
//		 - `currency_code` indicates the available currency. Currently only euro is available
//		 - `units` and `nanos` compose the price
//	    </Message>
//
// 3. **Manage an external domain.**
//
//	Refer to the Domains and DNS API documentation to find out [how to register and manage an external domain](https://www.scaleway.com/en/developers/api/domains-and-dns/#quickstart).
//
// 4. **Purchase a domain.**
//
//	Run the following command to purchase a domain:
//
//	```bash
//	curl https://api.scaleway.com/domain/v2beta1/buy-domains \
//	-X POST \
//	-H "X-Auth-Token: $SCW_SECRET_KEY" \
//	-d '{
//		"domains": [
//			"my_domain.fr"
//		],
//		"project_id": "$project_id",
//		"duration_in_years": 1,
//		"owner_contact": {
//			"legal_form": "individual",
//			"country": "FR",
//			"lastname": "DOE",
//			"firstname": "John",
//			"city": "Lille",
//			"address_line_1": "37 rue de la barre",
//			"zip": "59800",
//			"email": "email@mail.fr",
//			"country": "FR",
//			"phone_number": "0123456789",
//			"resale": true,
//			"extension_fr": {
//				"mode": "individual",
//				"individual_info": {
//					"whois_opt_in": true
//				}
//			}
//		}
//	}
//	```
//	<Message type="note">
//
//	`extension_fr.individual_info.whois_opt_in` returns `false` if the information of the domain's owner is anonymous. It returns `true` if the information of the domain's owner is published in whois.
//	</Message>
//
// 4. **List your domains.**
//
//	Run the following command to list your domains:
//
//	```bash
//	curl https://api.scaleway.com/domain/v2beta1/domains \
//	-H "X-Auth-Token: $SCW_SECRET_KEY"
//	```
//	An output similar to the following displays:
//
//	```bash
//	{
//	"total_count": 1,
//	"domains": [
//		{
//		"domain": "my_domain.fr",
//		"project_id": "xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxxxx",
//		"organization_id": "xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxxxx",
//		"auto_renew_status": "disabled",
//		"dnssec_status": "disabled",
//		"epp_code": [
//			"clientTransferProhibited"
//		],
//		"expired_at": "1970-01-01T00:00:00Z",
//		"updated_at": "1970-01-01T00:00:00Z",
//		"registrar": "SCALEWAY",
//		"is_external": false,
//		"status": "active"
//		}
//	]
//	}
//	```
//
// 5. **Enable DNSSEC for external domains.**
//
//   - Run the following command if you have not set up DNSSEC at your registrar:
//
//     ```bash
//     curl https://api.scaleway.com/domain/v2beta1/domains/example.com/enable-dnssec \
//     -X POST \
//     -H "Content-Type: application/json" \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     -d '{}'
//     ```
//     An output similar to the following displays:
//
//     ```json
//     {
//     "domain":{
//     "domain":"example.fr",
//     "auto_renew_status":"enabled",
//     "dnssec_status":"disabled",
//     "ds_records":[
//     {
//     "key_id":49071,
//     "algorithm":"ecdsap256sha256",
//     "digest":{
//     "type":"sha_384",
//     "digest":"c14421f739e35b62e3383d5972263f93610dea31da48a1990453752cf1b4f4f0314f75ecd0bc9061c51ee41ec7692204"
//     }
//     }
//     ],
//     "epp":[
//     "clientTransferProhibited"
//     ],
//     "expired_at":"2021-04-04T12:46:29Z",
//     "updated_at":"2020-07-07T13:24:27Z",
//     "registrar":"EXTERNAL",
//     "is_external":true,
//     "status":"active",
//     "......"
//     }
//     ```
//
//     <Message type="important">
//     Set up the information returned under `ds_records` in the output, in your registrar to configure DNSSEC for your external domain.
//     </Message>
//
//   - Run the following command if you have already set up DNSSEC at your registrar:
//
//     ```bash
//     curl https://api.scaleway.com/domain/v2beta1/domains/example.com/enable-dnssec \
//     -X POST \
//     -H "Content-Type: application/json" \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     -d '{
//     "ds_record": {
//     "algorithm": "rsamd5",
//     "digest":
//     {
//     "type": "sha_1",
//     "digest": "ABCD"
//     }
//     }
//     }'
//     ```
//
// 6. **Enable DNSSEC for internal domains.**
//
//	    Find out how to enable DNSSEC for your internal domains in our [dedicated documentation](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/configure-dnssec-internal-domain/).
//
//		<Message type="important">
//		 If your DNS service is managed by Scaleway, no other information is required from you you will need to provide the DS or public key information.
//		</Message>
//
// (switchcolumn)
// <Message type="requirement">
// - You have a [Scaleway account](https://console.scaleway.com/)
// - You have [created an API key](https://www.scaleway.com/en/docs/identity-and-access-management/iam/how-to/create-api-keys/) and that the API key has sufficient [IAM permissions](https://www.scaleway.com/en/docs/identity-and-access-management/iam/reference-content/permission-sets/) to perform the actions described on this page
// - You have your [Organization ID](https://console.scaleway.com/organization/settings)
// - You are the owner of your domain
// - You have [installed `curl`](https://curl.se/download.html)
// </Message>
// (switchcolumn)
//
// ## Technical information
//
// ### Regional availability
//
// Scaleway Domains and DNS is currently available globally. Find out about [product availability in our dedicated documentation](https://www.scaleway.com/en/docs/console/my-account/reference-content/products-availability).
//
// ## Technical limitations
//
// - Only the owner of a domain can use and configure it with Domains and DNS
// - Each zone supports the creation of sub-zones, which inherit the parent's Project ID
//
// ## Going futher
//
// For more information about Domains and DNS, you can check out the following pages:
//
// - [Domains and DNS Documentation](https://www.scaleway.com/en/docs/network/domains-and-dns/)
// - [Scaleway Slack Community](http://slack.scaleway.com) join the #domains channel
// - [Contact our support team](https://console.scaleway.com/support/tickets).
type RegistrarAPI struct {
	client *scw.Client
}

// NewRegistrarAPI returns a RegistrarAPI object from a Scaleway client.
func NewRegistrarAPI(client *scw.Client) *RegistrarAPI {
	return &RegistrarAPI{
		client: client,
	}
}

// ListTasks: List all operations performed on the account.
// You can filter the list of tasks by domain name.
func (s *RegistrarAPI) ListTasks(req *RegistrarAPIListTasksRequest, opts ...scw.RequestOption) (*ListTasksResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "domain", req.Domain)
	parameter.AddToQuery(query, "types", req.Types)
	parameter.AddToQuery(query, "statuses", req.Statuses)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain/v2beta1/tasks",
		Query:  query,
	}

	var resp ListTasksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// BuyDomains: Request the registration of domain names.
// You can provide a domain's already existing contact or a new contact.
func (s *RegistrarAPI) BuyDomains(req *RegistrarAPIBuyDomainsRequest, opts ...scw.RequestOption) (*OrderResponse, error) {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain/v2beta1/buy-domains",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp OrderResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RenewDomains: Request the renewal of one or more domain names.
func (s *RegistrarAPI) RenewDomains(req *RegistrarAPIRenewDomainsRequest, opts ...scw.RequestOption) (*OrderResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain/v2beta1/renew-domains",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp OrderResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// TransferInDomain: Request the transfer of a domain from another registrar to Scaleway Domains and DNS.
func (s *RegistrarAPI) TransferInDomain(req *RegistrarAPITransferInDomainRequest, opts ...scw.RequestOption) (*OrderResponse, error) {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain/v2beta1/domains/transfer-domains",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp OrderResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// TradeDomain: Request to change a domain's contact owner.<br/>
// If you specify the `organization_id` of the domain's new owner, the contact will change from the current owner's Scaleway account to the new owner's Scaleway account.<br/>
// If the new owner's current contact information is not available, the first ever contact they have created for previous domains is taken into account to operate the change.<br/>
// If the new owner has never created a contact to register domains before, an error message displays.
func (s *RegistrarAPI) TradeDomain(req *RegistrarAPITradeDomainRequest, opts ...scw.RequestOption) (*OrderResponse, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain/v2beta1/domains/" + fmt.Sprint(req.Domain) + "/trade",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp OrderResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RegisterExternalDomain: Request the registration of an external domain name.
func (s *RegistrarAPI) RegisterExternalDomain(req *RegistrarAPIRegisterExternalDomainRequest, opts ...scw.RequestOption) (*RegisterExternalDomainResponse, error) {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain/v2beta1/external-domains",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RegisterExternalDomainResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteExternalDomain: Delete an external domain name.
func (s *RegistrarAPI) DeleteExternalDomain(req *RegistrarAPIDeleteExternalDomainRequest, opts ...scw.RequestOption) (*DeleteExternalDomainResponse, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/domain/v2beta1/external-domains/" + fmt.Sprint(req.Domain) + "",
	}

	var resp DeleteExternalDomainResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckContactsCompatibility: Check whether contacts are compatible with a domain or a TLD.
// If contacts are not compatible with either the domain or the TLD, the information that needs to be corrected is returned.
func (s *RegistrarAPI) CheckContactsCompatibility(req *RegistrarAPICheckContactsCompatibilityRequest, opts ...scw.RequestOption) (*CheckContactsCompatibilityResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain/v2beta1/check-contacts-compatibility",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CheckContactsCompatibilityResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListContacts: Retrieve the list of contacts and their associated domains and roles.
// You can filter the list by domain name.
func (s *RegistrarAPI) ListContacts(req *RegistrarAPIListContactsRequest, opts ...scw.RequestOption) (*ListContactsResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "domain", req.Domain)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "role", req.Role)
	parameter.AddToQuery(query, "email_status", req.EmailStatus)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain/v2beta1/contacts",
		Query:  query,
	}

	var resp ListContactsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetContact: Retrieve a contact's details from the registrar using the given contact's ID.
func (s *RegistrarAPI) GetContact(req *RegistrarAPIGetContactRequest, opts ...scw.RequestOption) (*Contact, error) {
	var err error

	if fmt.Sprint(req.ContactID) == "" {
		return nil, errors.New("field ContactID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain/v2beta1/contacts/" + fmt.Sprint(req.ContactID) + "",
	}

	var resp Contact

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateContact: Edit the contact's information.
func (s *RegistrarAPI) UpdateContact(req *RegistrarAPIUpdateContactRequest, opts ...scw.RequestOption) (*Contact, error) {
	var err error

	if fmt.Sprint(req.ContactID) == "" {
		return nil, errors.New("field ContactID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/domain/v2beta1/contacts/" + fmt.Sprint(req.ContactID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Contact

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDomains: Retrieve the list of domains you own.
func (s *RegistrarAPI) ListDomains(req *RegistrarAPIListDomainsRequest, opts ...scw.RequestOption) (*ListDomainsResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "registrar", req.Registrar)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "is_external", req.IsExternal)
	parameter.AddToQuery(query, "domain", req.Domain)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain/v2beta1/domains",
		Query:  query,
	}

	var resp ListDomainsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListRenewableDomains: Retrieve the list of domains you own that can be renewed. You can also see the maximum renewal duration in years for your domains that are renewable.
func (s *RegistrarAPI) ListRenewableDomains(req *RegistrarAPIListRenewableDomainsRequest, opts ...scw.RequestOption) (*ListRenewableDomainsResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain/v2beta1/renewable-domains",
		Query:  query,
	}

	var resp ListRenewableDomainsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDomain: Retrieve a specific domain and display the domain's information.
func (s *RegistrarAPI) GetDomain(req *RegistrarAPIGetDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain/v2beta1/domains/" + fmt.Sprint(req.Domain) + "",
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateDomain: Update contacts for a specific domain or create a new contact.<br/>
// If you add the same contact for multiple roles (owner, administrative, technical), only one ID will be created and used for all of the roles.
func (s *RegistrarAPI) UpdateDomain(req *RegistrarAPIUpdateDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/domain/v2beta1/domains/" + fmt.Sprint(req.Domain) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// LockDomainTransfer: Lock the transfer of a domain. This means that the domain cannot be transferred and the authorization code cannot be requested to your current registrar.
func (s *RegistrarAPI) LockDomainTransfer(req *RegistrarAPILockDomainTransferRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain/v2beta1/domains/" + fmt.Sprint(req.Domain) + "/lock-transfer",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnlockDomainTransfer: Unlock the transfer of a domain. This means that the domain can be transferred and the authorization code can be requested to your current registrar.
func (s *RegistrarAPI) UnlockDomainTransfer(req *RegistrarAPIUnlockDomainTransferRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain/v2beta1/domains/" + fmt.Sprint(req.Domain) + "/unlock-transfer",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnableDomainAutoRenew: Enable the `auto renew` feature for a domain. This means the domain will be automatically renewed before its expiry date.
func (s *RegistrarAPI) EnableDomainAutoRenew(req *RegistrarAPIEnableDomainAutoRenewRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain/v2beta1/domains/" + fmt.Sprint(req.Domain) + "/enable-auto-renew",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DisableDomainAutoRenew: Disable the `auto renew` feature for a domain. This means the domain will not be renewed before its expiry date.
func (s *RegistrarAPI) DisableDomainAutoRenew(req *RegistrarAPIDisableDomainAutoRenewRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain/v2beta1/domains/" + fmt.Sprint(req.Domain) + "/disable-auto-renew",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDomainAuthCode: Retrieve the authorization code to tranfer an unlocked domain. The output returns an error if the domain is locked.
// Some TLDs may have a different procedure to retrieve the authorization code. In that case, the information displays in the message field.
func (s *RegistrarAPI) GetDomainAuthCode(req *RegistrarAPIGetDomainAuthCodeRequest, opts ...scw.RequestOption) (*GetDomainAuthCodeResponse, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain/v2beta1/domains/" + fmt.Sprint(req.Domain) + "/auth-code",
	}

	var resp GetDomainAuthCodeResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnableDomainDNSSEC: If your domain has the default Scaleway NS and uses another registrar, you have to update the DS record manually.
func (s *RegistrarAPI) EnableDomainDNSSEC(req *RegistrarAPIEnableDomainDNSSECRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain/v2beta1/domains/" + fmt.Sprint(req.Domain) + "/enable-dnssec",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DisableDomainDNSSEC: Disable DNSSEC for a domain.
func (s *RegistrarAPI) DisableDomainDNSSEC(req *RegistrarAPIDisableDomainDNSSECRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain/v2beta1/domains/" + fmt.Sprint(req.Domain) + "/disable-dnssec",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SearchAvailableDomains: Search a domain or a maximum of 10 domains that are available.
//
// If the TLD list is empty or not set, the search returns the results from the most popular TLDs.
func (s *RegistrarAPI) SearchAvailableDomains(req *RegistrarAPISearchAvailableDomainsRequest, opts ...scw.RequestOption) (*SearchAvailableDomainsResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "domains", req.Domains)
	parameter.AddToQuery(query, "tlds", req.Tlds)
	parameter.AddToQuery(query, "strict_search", req.StrictSearch)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain/v2beta1/search-domains",
		Query:  query,
	}

	var resp SearchAvailableDomainsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDomainHost: Create a hostname for a domain with glue IPs.
func (s *RegistrarAPI) CreateDomainHost(req *RegistrarAPICreateDomainHostRequest, opts ...scw.RequestOption) (*Host, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain/v2beta1/domains/" + fmt.Sprint(req.Domain) + "/hosts",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Host

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDomainHosts: List a domain's hostnames using their glue IPs.
func (s *RegistrarAPI) ListDomainHosts(req *RegistrarAPIListDomainHostsRequest, opts ...scw.RequestOption) (*ListDomainHostsResponse, error) {
	var err error
	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain/v2beta1/domains/" + fmt.Sprint(req.Domain) + "/hosts",
		Query:  query,
	}

	var resp ListDomainHostsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateDomainHost: Update a domain's hostname with glue IPs.
func (s *RegistrarAPI) UpdateDomainHost(req *RegistrarAPIUpdateDomainHostRequest, opts ...scw.RequestOption) (*Host, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	if fmt.Sprint(req.Name) == "" {
		return nil, errors.New("field Name cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/domain/v2beta1/domains/" + fmt.Sprint(req.Domain) + "/hosts/" + fmt.Sprint(req.Name) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Host

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDomainHost: Delete a domain's hostname.
func (s *RegistrarAPI) DeleteDomainHost(req *RegistrarAPIDeleteDomainHostRequest, opts ...scw.RequestOption) (*Host, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	if fmt.Sprint(req.Name) == "" {
		return nil, errors.New("field Name cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/domain/v2beta1/domains/" + fmt.Sprint(req.Domain) + "/hosts/" + fmt.Sprint(req.Name) + "",
	}

	var resp Host

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
