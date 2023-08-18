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
	CanPublish *bool `json:"can_publish,omitempty"`
	// CanReceive: Defines whether the credentials bearer can receive messages from the service.
	CanReceive *bool `json:"can_receive,omitempty"`
	// CanManage: Defines whether the credentials bearer can manage the associated resources (SQS queues or SNS topics or subscriptions).
	CanManage *bool `json:"can_manage,omitempty"`
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
	SecretKey *string `json:"secret_key,omitempty"`
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
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Namespace last modification date.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
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
	Permissions *Permissions `json:"permissions"`
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
	Permissions *Permissions `json:"permissions"`
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

// Scaleway Messaging and Queuing is a message broker tool that allows you to transfer messages between different platforms and port your microservice applications to the cloud. Messaging and Queuing offers notifications, queues, FIFOs and streams, and includes all major features of a modern message broker (such as message persistence).
//
// Currently, two protocol options are available for your namespaces (brokers): **NATS** and **SQS/SNS**.
//
// (switchcolumn)
// <Message type="note">
// Messaging and Queuing is currently in **Public Beta**
// </Message>
// (switchcolumn)
//
// ## Concepts
//
// Refer to our [dedicated concepts page](https://www.scaleway.com/en/docs/serverless/messaging/concepts/) to find definitions of all terminology related to Messaging and Queuing.
//
// (switchcolumn)
// (switchcolumn)
//
// ## Quickstart
//
// ### Creating a namespace and credentials
//
// 1. Configure your environment variables.
//
//	<Message type="note">
//	This is an optional step that seeks to simplify your usage of the API.
//	</Message>
//
//	 ```bash
//	 export SCW_ACCESS_KEY="<API access key>"
//	 export SCW_SECRET_KEY="<API secret key>"
//	 export SCW_PROJECT_ID="<Scaleway Project ID>"
//	 ```
//
// 2. **Create a namespace.** Run the following command to create a namespace. You can customize the details in the payload to your needs (e.g. to create a NATS or SQS/SNS namespace): use the information in the table below to help you.
//
//	```bash
//	curl -X POST \
//	    -H 'X-Auth-Token: '$SCW_SECRET_KEY \
//	    -H 'Content-Type: application/json' \
//	    'https://api.scaleway.com/mnq/v1alpha1/regions/fr-par/namespaces' \
//	    -d '{
//	        "name": "my-broker",
//	        "project_id": "'$SCW_PROJECT_ID'",
//	        "protocol": "nats"
//	    }' | tee my-broker.json
//	```
//
//	|   Parameter   |   Required value for a NATS namespace    |   Required value for an SQS/SNS namespace |
//	|---------------|------------------------------------------|-------------------------------------------|
//	| protocol      | `nats`                                   | `sqs_sns`                                 |
//
//	<Message type="tip">
//	Adding `| tee my-broker.json` to the end of the command saves the JSON object returned by the API to a file, for use in later steps.
//	</Message>
//
// 3. **Create credentials**. These are necessary to connect to your namespace.  Modify the command shown below as necessary:
//
//   - For `namespace_id`, use the ID of the namespace you created in the previous step.
//
//   - If you are creating credentials for a NATS namespace, delete the permissions object from the payload, as it is only applicable for SQS/SNS namespaces.
//
//   - If you are creating credentials for a SQS/SNS namespace, you can customize the permissions values to your needs, using `true` or `false`.
//
//     Use the following command to create credentials:
//
//     ```bash
//     curl -X POST \
//     -H 'X-Auth-Token: '$SCW_SECRET_KEY \
//     -H 'Content-Type: application/json' \
//     'https://api.scaleway.com/mnq/v1alpha1/regions/fr-par/credentials' \
//     -d '{
//     "name": "my-credentials",
//     "namespace_id": "AARYMIB9CGRRTS7ZZ2JUFQNS9ROJYQVX7RMWWPBSPRUZEMVOPMQESFUZI",
//     "permissions": {
//     "can_publish": true,
//     "can_receive": true,
//     "can_manage": true
//     }
//     }' | tee my-credentials.json
//     ```
//
// ### Publishing and subscribing to messages with NATS
//
// Make sure you have [jq](https://stedolan.github.io/jq/download/) installed on your machine, and that you have [installed the NATS CLI](https://www.scaleway.com/en/docs/serverless/messaging/api-cli/nats-cli/#installing-the-nats-cli)
//
// If you created a NATS namespace and credentials, follow these steps to publish and subscribe to your first messages.
//
// 1. Execute the following command to save the credentials from step 3 above to the required file format:
//
//	```bash
//	`jq -r .nats_credentials.content my-credentials.json > nats.creds`
//	```
//
// 2. Open two terminals, one to subscribe to a subject and another to publish a message on this subject.
//
// 3. In the first terminal, execute the following command to subscribe to a subject:
//
//	```bash
//	nats sub -s $(jq -r .endpoint my-broker.json) --creds nats.creds my-subject
//	```
//
// 4. In the second terminal, execute the following command to publish a message:
//
//	```bash
//	nats pub -s $(jq -r .endpoint my-broker.json) --creds nats.creds my-subject 'Hi there!'
//	```
//
// ### Creating queues and sending/receiving messages with SQS/SNS
//
// 1. Follow the steps in our [dedicated documentation](https://www.scaleway.com/en/docs/serverless/messaging/api-cli/connect-aws-cli/) to connect your namespace to the AWS CLI.
//
// 2. Follow the steps in our [dedicated documentation](https://www.scaleway.com/en/docs/serverless/messaging/api-cli/sqs-sns-aws-cli/) to start creating queues and sending and receiving messages with SQS or SNS.
//
// ### Deleting a namespace
//
// Use the following command to delete your namespace when you're finished with it. Ensure you replace the `{namespace_id}` parameter in the path with your own namespace ID.
//
//	```bash
//	curl -X DELETE \
//	    -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	    -H "Content-Type: application/json" \
//	    "https://api.scaleway.com/mnq/v1alpha1/regions/fr-par/namespaces/{namespace_id}"
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
// ## Technical limitations
//
// The following limitations apply when using Messaging and Queuing:
//
// - Scaleway Messaging and Queuing does not fully support all AWS SQS/SNS actions. See the compatibility matrices for [SNS](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sns-support/) and [SQS](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sqs-support/) for more information.
// - SNS currently supports subscriptions with the following protocols:
//   - sqs
//   - http
//   - https (self-signed certificate are not supported)
//   - lambda (used to call a Scaleway Serverless function)
//
// - The maximum message size is currently 256 kB.
// - The maximum storage capacity is currently 100 MB per namespace, across all its streams.
// - The product is currently in Public Beta, we do not guarantee performance or stability. Use with care, as you may encounter data loss.
//
// ## Technical information
//
// ### Regional availability
//
// Scaleway Messaging and Queuing is currently available only in the Paris, France region (`fr-par`).
//
// ### Credentials
//
// Note the following when creating credentials:
//
// - **NATS** credentials give full access to the messaging namespace
// - **SQS/SNS** credentials use a simplified permissions system: each credential created by the  gives permissions to the whole namespace. There are 3 permissions:
//   - `can_publish` : Allows to send messages to a Queue/Topic.
//   - `can_receive` : Allows to receive and acknowledge (delete) messages and to subscribe to a Topic.
//   - `can_manage`  : Allows all other actions (Creating, Listing, Updating, Deleting Queues/Topics).
//
// ## Going further
//
// For more help using Scaleway Messaing and Queuing, check out the following resources:
// - Our [main documentation](https://www.scaleway.com/en/docs/serverless/messaging/)
// - The `#messaging-queuing-beta` channel on our [Slack community](http://slack.scaleway.com)
// - Our [support ticketing system](https://www.scaleway.com/en/docs/console/my-account/how-to/open-a-support-ticket/)
//
// ### Going further with NATS
//
// - [What is NATS](https://docs.nats.io/nats-concepts/what-is-nats)
// - [Installing & using the NATS CLI](https://docs.nats.io/using-nats/nats-tools/nats_cli)
// - [JetStream official documentation](https://docs.nats.io/nats-concepts/jetstream) (JetStream is enabled with quotas on our servers)
// - [Official NATS GitHub](https://github.com/nats-io/nats.go): official client library for many languages ([golang](https://github.com/nats-io/nats.go), [python](https://github.com/nats-io/nats.py), [rust](https://github.com/nats-io/nats.rs) etc.) and many other coding resources
// - [NATS community Slack channel](https://slack.nats.io)
//
// ### Going further with SQS/SNS
//
// - [Amazon Simple Queue Service Documentation](https://docs.aws.amazon.com/sqs).
// - [Amazon Simple Notification Service Documentation](https://docs.aws.amazon.com/sns).
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
