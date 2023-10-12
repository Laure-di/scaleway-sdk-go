// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package mnq provides methods and message types of the mnq v1beta1 API.
package mnq

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

type ListNatsAccountsRequestOrderBy string

const (
	// Order by creation date (ascending chronological order).
	ListNatsAccountsRequestOrderByCreatedAtAsc = ListNatsAccountsRequestOrderBy("created_at_asc")
	// Order by creation date (descending chronological order).
	ListNatsAccountsRequestOrderByCreatedAtDesc = ListNatsAccountsRequestOrderBy("created_at_desc")
	// Order by last update date (ascending chronological order).
	ListNatsAccountsRequestOrderByUpdatedAtAsc = ListNatsAccountsRequestOrderBy("updated_at_asc")
	// Order by last update date (descending chronological order).
	ListNatsAccountsRequestOrderByUpdatedAtDesc = ListNatsAccountsRequestOrderBy("updated_at_desc")
	// Order by name (ascending alphabetical order).
	ListNatsAccountsRequestOrderByNameAsc = ListNatsAccountsRequestOrderBy("name_asc")
	// Order by name (descending alphabetical order).
	ListNatsAccountsRequestOrderByNameDesc = ListNatsAccountsRequestOrderBy("name_desc")
)

func (enum ListNatsAccountsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListNatsAccountsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListNatsAccountsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListNatsAccountsRequestOrderBy(ListNatsAccountsRequestOrderBy(tmp).String())
	return nil
}

type ListNatsCredentialsRequestOrderBy string

const (
	// Order by creation date (ascending chronological order).
	ListNatsCredentialsRequestOrderByCreatedAtAsc = ListNatsCredentialsRequestOrderBy("created_at_asc")
	// Order by creation date (descending chronological order).
	ListNatsCredentialsRequestOrderByCreatedAtDesc = ListNatsCredentialsRequestOrderBy("created_at_desc")
	// Order by last update date (ascending chronological order).
	ListNatsCredentialsRequestOrderByUpdatedAtAsc = ListNatsCredentialsRequestOrderBy("updated_at_asc")
	// Order by last update date (descending chronological order).
	ListNatsCredentialsRequestOrderByUpdatedAtDesc = ListNatsCredentialsRequestOrderBy("updated_at_desc")
	// Order by name (ascending alphabetical order).
	ListNatsCredentialsRequestOrderByNameAsc = ListNatsCredentialsRequestOrderBy("name_asc")
	// Order by name (descending alphabetical order).
	ListNatsCredentialsRequestOrderByNameDesc = ListNatsCredentialsRequestOrderBy("name_desc")
)

func (enum ListNatsCredentialsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListNatsCredentialsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListNatsCredentialsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListNatsCredentialsRequestOrderBy(ListNatsCredentialsRequestOrderBy(tmp).String())
	return nil
}

type ListSnsCredentialsRequestOrderBy string

const (
	// Order by creation date (ascending chronological order).
	ListSnsCredentialsRequestOrderByCreatedAtAsc = ListSnsCredentialsRequestOrderBy("created_at_asc")
	// Order by creation date (descending chronological order).
	ListSnsCredentialsRequestOrderByCreatedAtDesc = ListSnsCredentialsRequestOrderBy("created_at_desc")
	// Order by last update date (ascending chronological order).
	ListSnsCredentialsRequestOrderByUpdatedAtAsc = ListSnsCredentialsRequestOrderBy("updated_at_asc")
	// Order by last update date (descending chronological order).
	ListSnsCredentialsRequestOrderByUpdatedAtDesc = ListSnsCredentialsRequestOrderBy("updated_at_desc")
	// Order by name (ascending alphabetical order).
	ListSnsCredentialsRequestOrderByNameAsc = ListSnsCredentialsRequestOrderBy("name_asc")
	// Order by name (descending alphabetical order).
	ListSnsCredentialsRequestOrderByNameDesc = ListSnsCredentialsRequestOrderBy("name_desc")
)

func (enum ListSnsCredentialsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListSnsCredentialsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSnsCredentialsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSnsCredentialsRequestOrderBy(ListSnsCredentialsRequestOrderBy(tmp).String())
	return nil
}

type ListSqsCredentialsRequestOrderBy string

const (
	// Order by creation date (ascending chronological order).
	ListSqsCredentialsRequestOrderByCreatedAtAsc = ListSqsCredentialsRequestOrderBy("created_at_asc")
	// Order by creation date (descending chronological order).
	ListSqsCredentialsRequestOrderByCreatedAtDesc = ListSqsCredentialsRequestOrderBy("created_at_desc")
	// Order by last update date (ascending chronological order).
	ListSqsCredentialsRequestOrderByUpdatedAtAsc = ListSqsCredentialsRequestOrderBy("updated_at_asc")
	// Order by last update date (descending chronological order).
	ListSqsCredentialsRequestOrderByUpdatedAtDesc = ListSqsCredentialsRequestOrderBy("updated_at_desc")
	// Order by name (ascending alphabetical order).
	ListSqsCredentialsRequestOrderByNameAsc = ListSqsCredentialsRequestOrderBy("name_asc")
	// Order by name (descending alphabetical order).
	ListSqsCredentialsRequestOrderByNameDesc = ListSqsCredentialsRequestOrderBy("name_desc")
)

func (enum ListSqsCredentialsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListSqsCredentialsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSqsCredentialsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSqsCredentialsRequestOrderBy(ListSqsCredentialsRequestOrderBy(tmp).String())
	return nil
}

type SnsInfoStatus string

const (
	// Unknown status.
	SnsInfoStatusUnknownStatus = SnsInfoStatus("unknown_status")
	// Enabled status.
	SnsInfoStatusEnabled = SnsInfoStatus("enabled")
	// Disabled status.
	SnsInfoStatusDisabled = SnsInfoStatus("disabled")
)

func (enum SnsInfoStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum SnsInfoStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SnsInfoStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SnsInfoStatus(SnsInfoStatus(tmp).String())
	return nil
}

type SqsInfoStatus string

const (
	// Unknown status.
	SqsInfoStatusUnknownStatus = SqsInfoStatus("unknown_status")
	// Enabled status.
	SqsInfoStatusEnabled = SqsInfoStatus("enabled")
	// Disabled status.
	SqsInfoStatusDisabled = SqsInfoStatus("disabled")
)

func (enum SqsInfoStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum SqsInfoStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SqsInfoStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SqsInfoStatus(SqsInfoStatus(tmp).String())
	return nil
}

// File:
type File struct {
	// Name: File name.
	Name string `json:"name"`
	// Content: File content.
	Content string `json:"content"`
}

// SnsPermissions:
type SnsPermissions struct {
	// CanPublish: Defines whether the credentials bearer can publish messages to the service (publish to SNS topics).
	CanPublish *bool `json:"can_publish"`
	// CanReceive: Defines whether the credentials bearer can receive messages from the service (configure subscriptions).
	CanReceive *bool `json:"can_receive"`
	// CanManage: Defines whether the credentials bearer can manage the associated SNS topics or subscriptions.
	CanManage *bool `json:"can_manage"`
}

// SqsPermissions:
type SqsPermissions struct {
	// CanPublish: Defines whether the credentials bearer can publish messages to the service (send messages to SQS queues).
	CanPublish *bool `json:"can_publish"`
	// CanReceive: Defines whether the credentials bearer can receive messages from SQS queues.
	CanReceive *bool `json:"can_receive"`
	// CanManage: Defines whether the credentials bearer can manage the associated SQS queues.
	CanManage *bool `json:"can_manage"`
}

// NatsAccount:
type NatsAccount struct {
	// ID: NATS account ID.
	ID string `json:"id"`
	// Name: NATS account name.
	Name string `json:"name"`
	// Endpoint: Endpoint of the NATS service for this account.
	Endpoint string `json:"endpoint"`
	// ProjectID: Project ID of the Project containing the NATS account.
	ProjectID string `json:"project_id"`
	// Region: Region where the NATS account is deployed.
	Region scw.Region `json:"region"`
	// CreatedAt: NATS account creation date.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: NATS account last modification date.
	UpdatedAt *time.Time `json:"updated_at"`
}

// NatsCredentials:
type NatsCredentials struct {
	// ID: ID of the credentials.
	ID string `json:"id"`
	// Name: Name of the credentials.
	Name string `json:"name"`
	// NatsAccountID: NATS account containing the credentials.
	NatsAccountID string `json:"nats_account_id"`
	// CreatedAt: NATS credentials creation date.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: NATS credentials last modification date.
	UpdatedAt *time.Time `json:"updated_at"`
	// Credentials: Object containing the credentials file (Only returned by **Create Nats Credentials** call).
	Credentials *File `json:"credentials"`
	// Checksum: Checksum of the credentials file.
	Checksum string `json:"checksum"`
}

// SnsCredentials:
type SnsCredentials struct {
	// ID: ID of the credentials.
	ID string `json:"id"`
	// Name: Name of the credentials.
	Name string `json:"name"`
	// ProjectID: Project ID of the Project containing the credentials.
	ProjectID string `json:"project_id"`
	// Region: Region where the credentials exists.
	Region scw.Region `json:"region"`
	// CreatedAt: Credentials creation date.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Credentials last modification date.
	UpdatedAt *time.Time `json:"updated_at"`
	// AccessKey: Access key ID.
	AccessKey string `json:"access_key"`
	// SecretKey: Secret key ID (Only returned by **Create SNS Credentials** call).
	SecretKey string `json:"secret_key"`
	// SecretChecksum: Checksum of the Secret key.
	SecretChecksum string `json:"secret_checksum"`
	// Permissions: Permissions associated with these credentials.
	Permissions *SnsPermissions `json:"permissions"`
}

// SqsCredentials:
type SqsCredentials struct {
	// ID: ID of the credentials.
	ID string `json:"id"`
	// Name: Name of the credentials.
	Name string `json:"name"`
	// ProjectID: Project ID of the Project containing the credentials.
	ProjectID string `json:"project_id"`
	// Region: Region where the credentials exists.
	Region scw.Region `json:"region"`
	// CreatedAt: Credentials creation date.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Credentials last modification date.
	UpdatedAt *time.Time `json:"updated_at"`
	// AccessKey: Access key ID.
	AccessKey string `json:"access_key"`
	// SecretKey: Secret key ID (Only returned by **Create SQS Credentials** call).
	SecretKey string `json:"secret_key"`
	// SecretChecksum: Checksum of the Secret key.
	SecretChecksum string `json:"secret_checksum"`
	// Permissions: Permissions associated with these credentials.
	Permissions *SqsPermissions `json:"permissions"`
}

// ListNatsAccountsResponse:
type ListNatsAccountsResponse struct {
	// TotalCount: Total count of existing NATS accounts (matching any filters specified).
	TotalCount uint64 `json:"total_count"`
	// NatsAccounts: NATS accounts on this page.
	NatsAccounts []*NatsAccount `json:"nats_accounts"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNatsAccountsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNatsAccountsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListNatsAccountsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.NatsAccounts = append(r.NatsAccounts, results.NatsAccounts...)
	r.TotalCount += uint64(len(results.NatsAccounts))
	return uint64(len(results.NatsAccounts)), nil
}

// ListNatsCredentialsResponse:
type ListNatsCredentialsResponse struct {
	// TotalCount: Total count of existing credentials (matching any filters specified).
	TotalCount uint64 `json:"total_count"`
	// NatsCredentials: Credentials on this page.
	NatsCredentials []*NatsCredentials `json:"nats_credentials"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNatsCredentialsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNatsCredentialsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListNatsCredentialsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.NatsCredentials = append(r.NatsCredentials, results.NatsCredentials...)
	r.TotalCount += uint64(len(results.NatsCredentials))
	return uint64(len(results.NatsCredentials)), nil
}

// ListSnsCredentialsResponse:
type ListSnsCredentialsResponse struct {
	// TotalCount: Total count of existing credentials (matching any filters specified).
	TotalCount uint64 `json:"total_count"`
	// SnsCredentials: SNS credentials on this page.
	SnsCredentials []*SnsCredentials `json:"sns_credentials"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSnsCredentialsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSnsCredentialsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListSnsCredentialsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.SnsCredentials = append(r.SnsCredentials, results.SnsCredentials...)
	r.TotalCount += uint64(len(results.SnsCredentials))
	return uint64(len(results.SnsCredentials)), nil
}

// ListSqsCredentialsResponse:
type ListSqsCredentialsResponse struct {
	// TotalCount: Total count of existing credentials (matching any filters specified).
	TotalCount uint64 `json:"total_count"`
	// SqsCredentials: SQS credentials on this page.
	SqsCredentials []*SqsCredentials `json:"sqs_credentials"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSqsCredentialsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSqsCredentialsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListSqsCredentialsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.SqsCredentials = append(r.SqsCredentials, results.SqsCredentials...)
	r.TotalCount += uint64(len(results.SqsCredentials))
	return uint64(len(results.SqsCredentials)), nil
}

// NatsAPICreateNatsAccountRequest:
type NatsAPICreateNatsAccountRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Name: NATS account name.
	Name string `json:"name"`
	// ProjectID: Project containing the NATS account.
	ProjectID string `json:"project_id"`
}

// NatsAPICreateNatsCredentialsRequest:
type NatsAPICreateNatsCredentialsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NatsAccountID: NATS account containing the credentials.
	NatsAccountID string `json:"nats_account_id"`
	// Name: Name of the credentials.
	Name string `json:"name"`
}

// NatsAPIDeleteNatsAccountRequest:
type NatsAPIDeleteNatsAccountRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NatsAccountID: ID of the NATS account to delete.
	NatsAccountID string `json:"-"`
}

// NatsAPIDeleteNatsCredentialsRequest:
type NatsAPIDeleteNatsCredentialsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NatsCredentialsID: ID of the credentials to delete.
	NatsCredentialsID string `json:"-"`
}

// NatsAPIGetNatsAccountRequest:
type NatsAPIGetNatsAccountRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NatsAccountID: ID of the NATS account to get.
	NatsAccountID string `json:"-"`
}

// NatsAPIGetNatsCredentialsRequest:
type NatsAPIGetNatsCredentialsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NatsCredentialsID: ID of the credentials to get.
	NatsCredentialsID string `json:"-"`
}

// NatsAPIListNatsAccountsRequest:
type NatsAPIListNatsAccountsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID: Include only NATS accounts in this Project.
	ProjectID *string `json:"-"`
	// Page: Page number to return.
	Page *int32 `json:"-"`
	// PageSize: Maximum number of NATS accounts to return per page.
	PageSize *uint32 `json:"-"`
	// OrderBy: Order in which to return results.
	OrderBy ListNatsAccountsRequestOrderBy `json:"-"`
}

// NatsAPIListNatsCredentialsRequest:
type NatsAPIListNatsCredentialsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NatsAccountID: Include only credentials for this NATS account.
	NatsAccountID *string `json:"-"`
	// Page: Page number to return.
	Page *int32 `json:"-"`
	// PageSize: Maximum number of credentials to return per page.
	PageSize *uint32 `json:"-"`
	// OrderBy: Order in which to return results.
	OrderBy ListNatsCredentialsRequestOrderBy `json:"-"`
}

// NatsAPIUpdateNatsAccountRequest:
type NatsAPIUpdateNatsAccountRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NatsAccountID: ID of the NATS account to update.
	NatsAccountID string `json:"-"`
	// Name: NATS account name.
	Name *string `json:"name,omitempty"`
}

// SnsAPIActivateSnsRequest:
type SnsAPIActivateSnsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID: Project on which to activate the SNS service.
	ProjectID string `json:"project_id"`
}

// SnsAPICreateSnsCredentialsRequest:
type SnsAPICreateSnsCredentialsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID: Project containing the SNS credentials.
	ProjectID string `json:"project_id"`
	// Name: Name of the credentials.
	Name string `json:"name"`
	// Permissions: Permissions associated with these credentials.
	Permissions *SnsPermissions `json:"permissions,omitempty"`
}

// SnsAPIDeactivateSnsRequest:
type SnsAPIDeactivateSnsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID: Project on which to deactivate the SNS service.
	ProjectID string `json:"project_id"`
}

// SnsAPIDeleteSnsCredentialsRequest:
type SnsAPIDeleteSnsCredentialsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SnsCredentialsID: ID of the credentials to delete.
	SnsCredentialsID string `json:"-"`
}

// SnsAPIGetSnsCredentialsRequest:
type SnsAPIGetSnsCredentialsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SnsCredentialsID: ID of the SNS credentials to get.
	SnsCredentialsID string `json:"-"`
}

// SnsAPIGetSnsInfoRequest:
type SnsAPIGetSnsInfoRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID: Project to retrieve SNS info from.
	ProjectID string `json:"project_id"`
}

// SnsAPIListSnsCredentialsRequest:
type SnsAPIListSnsCredentialsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID: Include only SNS credentials in this Project.
	ProjectID *string `json:"-"`
	// Page: Page number to return.
	Page *int32 `json:"-"`
	// PageSize: Maximum number of credentials to return per page.
	PageSize *uint32 `json:"-"`
	// OrderBy: Order in which to return results.
	OrderBy ListSnsCredentialsRequestOrderBy `json:"-"`
}

// SnsAPIUpdateSnsCredentialsRequest:
type SnsAPIUpdateSnsCredentialsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SnsCredentialsID: ID of the SNS credentials to update.
	SnsCredentialsID string `json:"-"`
	// Name: Name of the credentials.
	Name *string `json:"name,omitempty"`
	// Permissions: Permissions associated with these credentials.
	Permissions *SnsPermissions `json:"permissions,omitempty"`
}

// SnsInfo:
type SnsInfo struct {
	// ProjectID: Project ID of the Project containing the service.
	ProjectID string `json:"project_id"`
	// Region: Region of the service.
	Region scw.Region `json:"region"`
	// CreatedAt: SNS creation date.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: SNS last modification date.
	UpdatedAt *time.Time `json:"updated_at"`
	// Status: SNS activation status.
	Status SnsInfoStatus `json:"status"`
	// SnsEndpointURL: Endpoint of the SNS service for this region and project.
	SnsEndpointURL string `json:"sns_endpoint_url"`
}

// SqsAPIActivateSqsRequest:
type SqsAPIActivateSqsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID: Project on which to activate the SQS service.
	ProjectID string `json:"project_id"`
}

// SqsAPICreateSqsCredentialsRequest:
type SqsAPICreateSqsCredentialsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID: Project containing the SQS credentials.
	ProjectID string `json:"project_id"`
	// Name: Name of the credentials.
	Name string `json:"name"`
	// Permissions: Permissions associated with these credentials.
	Permissions *SqsPermissions `json:"permissions,omitempty"`
}

// SqsAPIDeactivateSqsRequest:
type SqsAPIDeactivateSqsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID: Project on which to deactivate the SQS service.
	ProjectID string `json:"project_id"`
}

// SqsAPIDeleteSqsCredentialsRequest:
type SqsAPIDeleteSqsCredentialsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SqsCredentialsID: ID of the credentials to delete.
	SqsCredentialsID string `json:"-"`
}

// SqsAPIGetSqsCredentialsRequest:
type SqsAPIGetSqsCredentialsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SqsCredentialsID: ID of the SQS credentials to get.
	SqsCredentialsID string `json:"-"`
}

// SqsAPIGetSqsInfoRequest:
type SqsAPIGetSqsInfoRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID: Project to retrieve SQS info from.
	ProjectID string `json:"project_id"`
}

// SqsAPIListSqsCredentialsRequest:
type SqsAPIListSqsCredentialsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID: Include only SQS credentials in this Project.
	ProjectID *string `json:"-"`
	// Page: Page number to return.
	Page *int32 `json:"-"`
	// PageSize: Maximum number of credentials to return per page.
	PageSize *uint32 `json:"-"`
	// OrderBy: Order in which to return results.
	OrderBy ListSqsCredentialsRequestOrderBy `json:"-"`
}

// SqsAPIUpdateSqsCredentialsRequest:
type SqsAPIUpdateSqsCredentialsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SqsCredentialsID: ID of the SQS credentials to update.
	SqsCredentialsID string `json:"-"`
	// Name: Name of the credentials.
	Name *string `json:"name,omitempty"`
	// Permissions: Permissions associated with these credentials.
	Permissions *SqsPermissions `json:"permissions,omitempty"`
}

// SqsInfo:
type SqsInfo struct {
	// ProjectID: Project ID of the Project containing the service.
	ProjectID string `json:"project_id"`
	// Region: Region of the service.
	Region scw.Region `json:"region"`
	// CreatedAt: SQS creation date.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: SQS last modification date.
	UpdatedAt *time.Time `json:"updated_at"`
	// Status: SQS activation status.
	Status SqsInfoStatus `json:"status"`
	// SqsEndpointURL: Endpoint of the SQS service for this region and project.
	SqsEndpointURL string `json:"sqs_endpoint_url"`
}

// This API allows you to manage Scaleway Messaging and Queueing NATS accounts.
type NatsAPI struct {
	client *scw.Client
}

// NewNatsAPI returns a NatsAPI object from a Scaleway client.
func NewNatsAPI(client *scw.Client) *NatsAPI {
	return &NatsAPI{
		client: client,
	}
}
func (s *NatsAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar}
}

// CreateNatsAccount: Create a NATS account associated with a Project.
func (s *NatsAPI) CreateNatsAccount(req *NatsAPICreateNatsAccountRequest, opts ...scw.RequestOption) (*NatsAccount, error) {
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
		req.Name = namegenerator.GetRandomName("mnq")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/nats-accounts",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp NatsAccount

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteNatsAccount: Delete a NATS account, specified by its NATS account ID. Note that deleting a NATS account is irreversible, and any credentials, streams, consumer and stored messages belonging to this NATS account will also be deleted.
func (s *NatsAPI) DeleteNatsAccount(req *NatsAPIDeleteNatsAccountRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NatsAccountID) == "" {
		return errors.New("field NatsAccountID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/nats-accounts/" + fmt.Sprint(req.NatsAccountID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdateNatsAccount: Update the name of a NATS account, specified by its NATS account ID.
func (s *NatsAPI) UpdateNatsAccount(req *NatsAPIUpdateNatsAccountRequest, opts ...scw.RequestOption) (*NatsAccount, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NatsAccountID) == "" {
		return nil, errors.New("field NatsAccountID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/nats-accounts/" + fmt.Sprint(req.NatsAccountID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp NatsAccount

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNatsAccount: Retrieve information about an existing NATS account identified by its NATS account ID. Its full details, including name and endpoint, are returned in the response.
func (s *NatsAPI) GetNatsAccount(req *NatsAPIGetNatsAccountRequest, opts ...scw.RequestOption) (*NatsAccount, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NatsAccountID) == "" {
		return nil, errors.New("field NatsAccountID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/nats-accounts/" + fmt.Sprint(req.NatsAccountID) + "",
	}

	var resp NatsAccount

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNatsAccounts: List all NATS accounts in the specified region, for a Scaleway Organization or Project. By default, the NATS accounts returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *NatsAPI) ListNatsAccounts(req *NatsAPIListNatsAccountsRequest, opts ...scw.RequestOption) (*ListNatsAccountsResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/nats-accounts",
		Query:  query,
	}

	var resp ListNatsAccountsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateNatsCredentials: Create a set of credentials for a NATS account, specified by its NATS account ID.
func (s *NatsAPI) CreateNatsCredentials(req *NatsAPICreateNatsCredentialsRequest, opts ...scw.RequestOption) (*NatsCredentials, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("mnq")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/nats-credentials",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp NatsCredentials

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteNatsCredentials: Delete a set of credentials, specified by their credentials ID. Deleting credentials is irreversible and cannot be undone. The credentials can no longer be used to access the NATS account, and active connections using this credentials will be closed.
func (s *NatsAPI) DeleteNatsCredentials(req *NatsAPIDeleteNatsCredentialsRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NatsCredentialsID) == "" {
		return errors.New("field NatsCredentialsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/nats-credentials/" + fmt.Sprint(req.NatsCredentialsID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetNatsCredentials: Retrieve an existing set of credentials, identified by the `nats_credentials_id`. The credentials themselves are NOT returned, only their metadata (NATS account ID, credentials name, etc), are returned in the response.
func (s *NatsAPI) GetNatsCredentials(req *NatsAPIGetNatsCredentialsRequest, opts ...scw.RequestOption) (*NatsCredentials, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NatsCredentialsID) == "" {
		return nil, errors.New("field NatsCredentialsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/nats-credentials/" + fmt.Sprint(req.NatsCredentialsID) + "",
	}

	var resp NatsCredentials

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNatsCredentials: List existing credentials in the specified NATS account. The response contains only the metadata for the credentials, not the credentials themselves, which are only returned after a **Create Credentials** call.
func (s *NatsAPI) ListNatsCredentials(req *NatsAPIListNatsCredentialsRequest, opts ...scw.RequestOption) (*ListNatsCredentialsResponse, error) {
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
	parameter.AddToQuery(query, "nats_account_id", req.NatsAccountID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/nats-credentials",
		Query:  query,
	}

	var resp ListNatsCredentialsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This API allows you to manage Scaleway Messaging and Queueing SNS brokers.
type SnsAPI struct {
	client *scw.Client
}

// NewSnsAPI returns a SnsAPI object from a Scaleway client.
func NewSnsAPI(client *scw.Client) *SnsAPI {
	return &SnsAPI{
		client: client,
	}
}
func (s *SnsAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar}
}

// ActivateSns: Activate SNS for the specified Project ID. SNS must be activated before any usage. Activating SNS does not trigger any billing, and you can deactivate at any time.
func (s *SnsAPI) ActivateSns(req *SnsAPIActivateSnsRequest, opts ...scw.RequestOption) (*SnsInfo, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/activate-sns",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SnsInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSnsInfo: Retrieve the SNS information of the specified Project ID. Informations include the activation status and the SNS API endpoint URL.
func (s *SnsAPI) GetSnsInfo(req *SnsAPIGetSnsInfoRequest, opts ...scw.RequestOption) (*SnsInfo, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sns-info",
		Query:  query,
	}

	var resp SnsInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeactivateSns: Deactivate SNS for the specified Project ID.You must delete all topics and credentials before this call or you need to set the force_delete parameter.
func (s *SnsAPI) DeactivateSns(req *SnsAPIDeactivateSnsRequest, opts ...scw.RequestOption) (*SnsInfo, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deactivate-sns",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SnsInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSnsCredentials: Create a set of credentials for SNS, specified by a Project ID. Credentials give the bearer access to topics, and the level of permissions can be defined granularly.
func (s *SnsAPI) CreateSnsCredentials(req *SnsAPICreateSnsCredentialsRequest, opts ...scw.RequestOption) (*SnsCredentials, error) {
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
		req.Name = namegenerator.GetRandomName("mnq_sns")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sns-credentials",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SnsCredentials

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSnsCredentials: Delete a set of SNS credentials, specified by their credentials ID. Deleting credentials is irreversible and cannot be undone. The credentials can then no longer be used to access SNS.
func (s *SnsAPI) DeleteSnsCredentials(req *SnsAPIDeleteSnsCredentialsRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SnsCredentialsID) == "" {
		return errors.New("field SnsCredentialsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sns-credentials/" + fmt.Sprint(req.SnsCredentialsID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdateSnsCredentials: Update a set of SNS credentials. You can update the credentials' name, or their permissions.
func (s *SnsAPI) UpdateSnsCredentials(req *SnsAPIUpdateSnsCredentialsRequest, opts ...scw.RequestOption) (*SnsCredentials, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SnsCredentialsID) == "" {
		return nil, errors.New("field SnsCredentialsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sns-credentials/" + fmt.Sprint(req.SnsCredentialsID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SnsCredentials

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSnsCredentials: Retrieve an existing set of credentials, identified by the `credentials_id`. The credentials themselves, as well as their metadata (name, project ID etc), are returned in the response.
func (s *SnsAPI) GetSnsCredentials(req *SnsAPIGetSnsCredentialsRequest, opts ...scw.RequestOption) (*SnsCredentials, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SnsCredentialsID) == "" {
		return nil, errors.New("field SnsCredentialsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sns-credentials/" + fmt.Sprint(req.SnsCredentialsID) + "",
	}

	var resp SnsCredentials

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSnsCredentials: List existing SNS credentials in the specified region. The response contains only the metadata for the credentials, not the credentials themselves.
func (s *SnsAPI) ListSnsCredentials(req *SnsAPIListSnsCredentialsRequest, opts ...scw.RequestOption) (*ListSnsCredentialsResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sns-credentials",
		Query:  query,
	}

	var resp ListSnsCredentialsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This API allows you to manage Scaleway Messaging and Queueing SQS brokers.
type SqsAPI struct {
	client *scw.Client
}

// NewSqsAPI returns a SqsAPI object from a Scaleway client.
func NewSqsAPI(client *scw.Client) *SqsAPI {
	return &SqsAPI{
		client: client,
	}
}
func (s *SqsAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar}
}

// ActivateSqs: Activate SQS for the specified Project ID. SQS must be activated before any usage such as creating credentials and queues. Activating SQS does not trigger any billing, and you can deactivate at any time.
func (s *SqsAPI) ActivateSqs(req *SqsAPIActivateSqsRequest, opts ...scw.RequestOption) (*SqsInfo, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/activate-sqs",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SqsInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSqsInfo: Retrieve the SQS information of the specified Project ID. Informations include the activation status and the SQS API endpoint URL.
func (s *SqsAPI) GetSqsInfo(req *SqsAPIGetSqsInfoRequest, opts ...scw.RequestOption) (*SqsInfo, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sqs-info",
		Query:  query,
	}

	var resp SqsInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeactivateSqs: Deactivate SQS for the specified Project ID. You must delete all queues and credentials before this call or you need to set the force_delete parameter.
func (s *SqsAPI) DeactivateSqs(req *SqsAPIDeactivateSqsRequest, opts ...scw.RequestOption) (*SqsInfo, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deactivate-sqs",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SqsInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSqsCredentials: Create a set of credentials for SQS, specified by a Project ID. Credentials give the bearer access to queues, and the level of permissions can be defined granularly.
func (s *SqsAPI) CreateSqsCredentials(req *SqsAPICreateSqsCredentialsRequest, opts ...scw.RequestOption) (*SqsCredentials, error) {
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
		req.Name = namegenerator.GetRandomName("mnq_sqs")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sqs-credentials",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SqsCredentials

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSqsCredentials: Delete a set of SQS credentials, specified by their credentials ID. Deleting credentials is irreversible and cannot be undone. The credentials can then no longer be used to access SQS.
func (s *SqsAPI) DeleteSqsCredentials(req *SqsAPIDeleteSqsCredentialsRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SqsCredentialsID) == "" {
		return errors.New("field SqsCredentialsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sqs-credentials/" + fmt.Sprint(req.SqsCredentialsID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdateSqsCredentials: Update a set of SQS credentials. You can update the credentials' name, or their permissions.
func (s *SqsAPI) UpdateSqsCredentials(req *SqsAPIUpdateSqsCredentialsRequest, opts ...scw.RequestOption) (*SqsCredentials, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SqsCredentialsID) == "" {
		return nil, errors.New("field SqsCredentialsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sqs-credentials/" + fmt.Sprint(req.SqsCredentialsID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SqsCredentials

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSqsCredentials: Retrieve an existing set of credentials, identified by the `credentials_id`. The credentials themselves, as well as their metadata (name, project ID etc), are returned in the response.
func (s *SqsAPI) GetSqsCredentials(req *SqsAPIGetSqsCredentialsRequest, opts ...scw.RequestOption) (*SqsCredentials, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SqsCredentialsID) == "" {
		return nil, errors.New("field SqsCredentialsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sqs-credentials/" + fmt.Sprint(req.SqsCredentialsID) + "",
	}

	var resp SqsCredentials

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSqsCredentials: List existing SQS credentials in the specified region. The response contains only the metadata for the credentials, not the credentials themselves.
func (s *SqsAPI) ListSqsCredentials(req *SqsAPIListSqsCredentialsRequest, opts ...scw.RequestOption) (*ListSqsCredentialsResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sqs-credentials",
		Query:  query,
	}

	var resp ListSqsCredentialsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
