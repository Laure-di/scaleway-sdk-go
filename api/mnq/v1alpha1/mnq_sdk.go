// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package mnq provides methods and message types of the mnq v1alpha1 API.
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

type ListCredentialsRequestOrderBy string

const (
	// Order by id (ascending alphabetical order).
	ListCredentialsRequestOrderByIDAsc = ListCredentialsRequestOrderBy("id_asc")
	// Order by id (descending alphabetical order).
	ListCredentialsRequestOrderByIDDesc = ListCredentialsRequestOrderBy("id_desc")
	// Order by name (ascending alphabetical order).
	ListCredentialsRequestOrderByNameAsc = ListCredentialsRequestOrderBy("name_asc")
	// Order by name (descending alphabetical order).
	ListCredentialsRequestOrderByNameDesc = ListCredentialsRequestOrderBy("name_desc")
)

func (enum ListCredentialsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "id_asc"
	}
	return string(enum)
}

func (enum ListCredentialsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListCredentialsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListCredentialsRequestOrderBy(ListCredentialsRequestOrderBy(tmp).String())
	return nil
}

type ListNamespacesRequestOrderBy string

const (
	// Order by creation date (ascending chronological order).
	ListNamespacesRequestOrderByCreatedAtAsc = ListNamespacesRequestOrderBy("created_at_asc")
	// Order by creation date (descending chronological order).
	ListNamespacesRequestOrderByCreatedAtDesc = ListNamespacesRequestOrderBy("created_at_desc")
	// Order by last update date (ascending chronological order).
	ListNamespacesRequestOrderByUpdatedAtAsc = ListNamespacesRequestOrderBy("updated_at_asc")
	// Order by last update date (descending chronological order).
	ListNamespacesRequestOrderByUpdatedAtDesc = ListNamespacesRequestOrderBy("updated_at_desc")
	// Order by id (ascending alphabetical order).
	ListNamespacesRequestOrderByIDAsc = ListNamespacesRequestOrderBy("id_asc")
	// Order by id (descending alphabetical order).
	ListNamespacesRequestOrderByIDDesc = ListNamespacesRequestOrderBy("id_desc")
	// Order by name (ascending alphabetical order).
	ListNamespacesRequestOrderByNameAsc = ListNamespacesRequestOrderBy("name_asc")
	// Order by name (descending alphabetical order).
	ListNamespacesRequestOrderByNameDesc = ListNamespacesRequestOrderBy("name_desc")
	// Order by project_id (ascending alphabetical order).
	ListNamespacesRequestOrderByProjectIDAsc = ListNamespacesRequestOrderBy("project_id_asc")
	// Order by project_id (descending alphabetical order).
	ListNamespacesRequestOrderByProjectIDDesc = ListNamespacesRequestOrderBy("project_id_desc")
)

func (enum ListNamespacesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListNamespacesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListNamespacesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListNamespacesRequestOrderBy(ListNamespacesRequestOrderBy(tmp).String())
	return nil
}

type NamespaceProtocol string

const (
	// Unknown protocol.
	NamespaceProtocolUnknown = NamespaceProtocol("unknown")
	// NATS protocol.
	NamespaceProtocolNats = NamespaceProtocol("nats")
	// SQS / SNS protocol.
	NamespaceProtocolSqsSns = NamespaceProtocol("sqs_sns")
)

func (enum NamespaceProtocol) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum NamespaceProtocol) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NamespaceProtocol) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NamespaceProtocol(NamespaceProtocol(tmp).String())
	return nil
}

// Permissions:
type Permissions struct {
	// CanPublish: Defines whether the credentials bearer can publish messages to the service (send messages to SQS queues or publish to SNS topics).
	CanPublish *bool `json:"can_publish"`
	// CanReceive: Defines whether the credentials bearer can receive messages from the service.
	CanReceive *bool `json:"can_receive"`
	// CanManage: Defines whether the credentials bearer can manage the associated resources (SQS queues or SNS topics or subscriptions).
	CanManage *bool `json:"can_manage"`
}

// CredentialSummarySQSSNSCreds:
type CredentialSummarySQSSNSCreds struct {
	// AccessKey: Access key ID.
	AccessKey string `json:"access_key"`
	// Permissions: Permissions associated with these credentials.
	Permissions *Permissions `json:"permissions"`
}

// CredentialNATSCredsFile:
type CredentialNATSCredsFile struct {
	// Content: Raw content of the NATS credentials file.
	Content string `json:"content"`
}

// CredentialSQSSNSCreds:
type CredentialSQSSNSCreds struct {
	// AccessKey: Access key ID.
	AccessKey string `json:"access_key"`
	// SecretKey: Secret key ID.
	SecretKey *string `json:"secret_key"`
	// Permissions: Permissions associated with these credentials.
	Permissions *Permissions `json:"permissions"`
}

// CredentialSummary:
type CredentialSummary struct {
	// ID: ID of the credentials.
	ID string `json:"id"`
	// Name: Name of the credentials.
	Name string `json:"name"`
	// NamespaceID: Namespace containing the credentials.
	NamespaceID string `json:"namespace_id"`
	// Protocol: Protocol associated with the credentials.
	Protocol NamespaceProtocol `json:"protocol"`
	// SqsSnsCredentials: Object containing the credentials and their metadata, if the credentials are for an SQS/SNS namespace.
	SqsSnsCredentials *CredentialSummarySQSSNSCreds `json:"sqs_sns_credentials,omitempty"`
}

// Namespace:
type Namespace struct {
	// ID: Namespace ID.
	ID string `json:"id"`
	// Name: Namespace name.
	Name string `json:"name"`
	// Endpoint: Endpoint of the service matching the namespace's protocol.
	Endpoint string `json:"endpoint"`
	// Protocol: Namespace protocol.
	Protocol NamespaceProtocol `json:"protocol"`
	// ProjectID: Project ID of the Project containing the namespace.
	ProjectID string `json:"project_id"`
	// Region: Region where the namespace is deployed.
	Region scw.Region `json:"region"`
	// CreatedAt: Namespace creation date.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Namespace last modification date.
	UpdatedAt *time.Time `json:"updated_at"`
}

// CreateCredentialRequest:
type CreateCredentialRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NamespaceID: Namespace containing the credentials.
	NamespaceID string `json:"namespace_id"`
	// Name: Name of the credentials.
	Name string `json:"name"`
	// Permissions: Permissions associated with these credentials.
	Permissions *Permissions `json:"permissions,omitempty"`
}

// CreateNamespaceRequest:
type CreateNamespaceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Name: Namespace name.
	Name string `json:"name"`
	// Protocol: Namespace protocol. You must specify a valid protocol (and not `unknown`) to avoid an error.
	Protocol NamespaceProtocol `json:"protocol"`
	// ProjectID: Project containing the Namespace.
	ProjectID string `json:"project_id"`
}

// Credential:
type Credential struct {
	// ID: ID of the credentials.
	ID string `json:"id"`
	// Name: Name of the credentials.
	Name string `json:"name"`
	// NamespaceID: Namespace containing the credentials.
	NamespaceID string `json:"namespace_id"`
	// Protocol: Protocol associated with the credentials.
	Protocol NamespaceProtocol `json:"protocol"`
	// NatsCredentials: Object containing the credentials, if the credentials are for a NATS namespace.
	NatsCredentials *CredentialNATSCredsFile `json:"nats_credentials,omitempty"`
	// SqsSnsCredentials: Object containing the credentials and their metadata, if the credentials are for an SQS/SNS namespace.
	SqsSnsCredentials *CredentialSQSSNSCreds `json:"sqs_sns_credentials,omitempty"`
}

// DeleteCredentialRequest:
type DeleteCredentialRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// CredentialID: ID of the credentials to delete.
	CredentialID string `json:"-"`
}

// DeleteNamespaceRequest:
type DeleteNamespaceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NamespaceID: ID of the namespace to delete.
	NamespaceID string `json:"-"`
}

// GetCredentialRequest:
type GetCredentialRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// CredentialID: ID of the credentials to get.
	CredentialID string `json:"-"`
}

// GetNamespaceRequest:
type GetNamespaceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NamespaceID: ID of the Namespace to get.
	NamespaceID string `json:"-"`
}

// ListCredentialsRequest:
type ListCredentialsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NamespaceID: Namespace containing the credentials.
	NamespaceID *string `json:"-"`
	// Page: Page number to return.
	Page *int32 `json:"-"`
	// PageSize: Maximum number of credentials to return per page.
	PageSize *uint32 `json:"-"`
	// OrderBy: Order in which to return results.
	OrderBy ListCredentialsRequestOrderBy `json:"-"`
}

// ListCredentialsResponse:
type ListCredentialsResponse struct {
	// TotalCount: Total count of existing credentials (matching any filters specified).
	TotalCount uint32 `json:"total_count"`
	// Credentials: Credentials on this page.
	Credentials []*CredentialSummary `json:"credentials"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCredentialsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCredentialsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListCredentialsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Credentials = append(r.Credentials, results.Credentials...)
	r.TotalCount += uint32(len(results.Credentials))
	return uint32(len(results.Credentials)), nil
}

// ListNamespacesRequest:
type ListNamespacesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// OrganizationID: Include only namespaces in this Organization.
	OrganizationID *string `json:"-"`
	// ProjectID: Include only namespaces in this Project.
	ProjectID *string `json:"-"`
	// Page: Page number to return.
	Page *int32 `json:"-"`
	// PageSize: Maximum number of namespaces to return per page.
	PageSize *uint32 `json:"-"`
	// OrderBy: Order in which to return results.
	OrderBy ListNamespacesRequestOrderBy `json:"-"`
}

// ListNamespacesResponse:
type ListNamespacesResponse struct {
	// TotalCount: Total count of existing namespaces (matching any filters specified).
	TotalCount uint32 `json:"total_count"`
	// Namespaces: Namespaces on this page.
	Namespaces []*Namespace `json:"namespaces"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNamespacesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNamespacesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListNamespacesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Namespaces = append(r.Namespaces, results.Namespaces...)
	r.TotalCount += uint32(len(results.Namespaces))
	return uint32(len(results.Namespaces)), nil
}

// UpdateCredentialRequest:
type UpdateCredentialRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// CredentialID: ID of the credentials to update.
	CredentialID string `json:"-"`
	// Name: Name of the credentials.
	Name *string `json:"name,omitempty"`
	// Permissions: Permissions associated with these credentials.
	Permissions *Permissions `json:"permissions,omitempty"`
}

// UpdateNamespaceRequest:
type UpdateNamespaceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NamespaceID: ID of the Namespace to update.
	NamespaceID string `json:"namespace_id"`
	// Name: Namespace name.
	Name *string `json:"name,omitempty"`
}

// This API allows you to manage Scaleway Messaging and Queueing brokers.
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

// ListNamespaces: List all Messaging and Queuing namespaces in the specified region, for a Scaleway Organization or Project. By default, the namespaces returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListNamespaces(req *ListNamespacesRequest, opts ...scw.RequestOption) (*ListNamespacesResponse, error) {
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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
		Query:  query,
	}

	var resp ListNamespacesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateNamespace: Create a Messaging and Queuing namespace, set to the desired protocol.
func (s *API) CreateNamespace(req *CreateNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
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
		Path:   "/mnq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateNamespace: Update the name of a Messaging and Queuing namespace, specified by its namespace ID.
func (s *API) UpdateNamespace(req *UpdateNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/mnq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNamespace: Retrieve information about an existing Messaging and Queuing namespace, identified by its namespace ID. Its full details, including name, endpoint and protocol, are returned in the response.
func (s *API) GetNamespace(req *GetNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NamespaceID) == "" {
		return nil, errors.New("field NamespaceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteNamespace: Delete a Messaging and Queuing namespace, specified by its namespace ID. Note that deleting a namespace is irreversible, and any URLs, credentials and queued messages belonging to this namespace will also be deleted.
func (s *API) DeleteNamespace(req *DeleteNamespaceRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NamespaceID) == "" {
		return errors.New("field NamespaceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/mnq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CreateCredential: Create a set of credentials for a Messaging and Queuing namespace, specified by its namespace ID. If creating credentials for a NATS namespace, the `permissions` object must not be included in the request. If creating credentials for an SQS/SNS namespace, the `permissions` object is required, with all three of its child attributes.
func (s *API) CreateCredential(req *CreateCredentialRequest, opts ...scw.RequestOption) (*Credential, error) {
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
		Path:   "/mnq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/credentials",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Credential

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteCredential: Delete a set of credentials, specified by their credential ID. Deleting credentials is irreversible and cannot be undone. The credentials can no longer be used to access the namespace.
func (s *API) DeleteCredential(req *DeleteCredentialRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.CredentialID) == "" {
		return errors.New("field CredentialID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/mnq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/credentials/" + fmt.Sprint(req.CredentialID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListCredentials: List existing credentials in the specified region. The response contains only the metadata for the credentials, not the credentials themselves (for this, use **Get Credentials**).
func (s *API) ListCredentials(req *ListCredentialsRequest, opts ...scw.RequestOption) (*ListCredentialsResponse, error) {
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
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/credentials",
		Query:  query,
	}

	var resp ListCredentialsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateCredential: Update a set of credentials. You can update the credentials' name, or (in the case of SQS/SNS credentials only) their permissions. To update the name of NATS credentials, do not include the `permissions` object in your request.
func (s *API) UpdateCredential(req *UpdateCredentialRequest, opts ...scw.RequestOption) (*Credential, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.CredentialID) == "" {
		return nil, errors.New("field CredentialID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/mnq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/credentials/" + fmt.Sprint(req.CredentialID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Credential

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCredential: Retrieve an existing set of credentials, identified by the `credential_id`. The credentials themselves, as well as their metadata (protocol, namespace ID etc), are returned in the response.
func (s *API) GetCredential(req *GetCredentialRequest, opts ...scw.RequestOption) (*Credential, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.CredentialID) == "" {
		return nil, errors.New("field CredentialID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/credentials/" + fmt.Sprint(req.CredentialID) + "",
	}

	var resp Credential

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
