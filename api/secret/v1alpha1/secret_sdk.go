// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package secret provides methods and message types of the secret v1alpha1 API.
package secret

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

type ListSecretsRequestOrderBy string

const (
	ListSecretsRequestOrderByNameAsc       = ListSecretsRequestOrderBy("name_asc")
	ListSecretsRequestOrderByNameDesc      = ListSecretsRequestOrderBy("name_desc")
	ListSecretsRequestOrderByCreatedAtAsc  = ListSecretsRequestOrderBy("created_at_asc")
	ListSecretsRequestOrderByCreatedAtDesc = ListSecretsRequestOrderBy("created_at_desc")
	ListSecretsRequestOrderByUpdatedAtAsc  = ListSecretsRequestOrderBy("updated_at_asc")
	ListSecretsRequestOrderByUpdatedAtDesc = ListSecretsRequestOrderBy("updated_at_desc")
)

func (enum ListSecretsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListSecretsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSecretsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSecretsRequestOrderBy(ListSecretsRequestOrderBy(tmp).String())
	return nil
}

type Product string

const (
	ProductUnknown = Product("unknown")
)

func (enum Product) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum Product) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *Product) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = Product(Product(tmp).String())
	return nil
}

type SecretStatus string

const (
	SecretStatusReady  = SecretStatus("ready")
	SecretStatusLocked = SecretStatus("locked")
)

func (enum SecretStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "ready"
	}
	return string(enum)
}

func (enum SecretStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecretStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecretStatus(SecretStatus(tmp).String())
	return nil
}

type SecretType string

const (
	SecretTypeUnknownSecretType = SecretType("unknown_secret_type")
	// Default type.
	SecretTypeOpaque = SecretType("opaque")
	// Certificates used by load balancers, the format must be PEM concatenated and contains exactly one PKCS8 private key and the certificate full chain containing all intermediates CA.
	SecretTypeCertificate = SecretType("certificate")
)

func (enum SecretType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_secret_type"
	}
	return string(enum)
}

func (enum SecretType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecretType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecretType(SecretType(tmp).String())
	return nil
}

type SecretVersionStatus string

const (
	SecretVersionStatusUnknown   = SecretVersionStatus("unknown")
	SecretVersionStatusEnabled   = SecretVersionStatus("enabled")
	SecretVersionStatusDisabled  = SecretVersionStatus("disabled")
	SecretVersionStatusDestroyed = SecretVersionStatus("destroyed")
)

func (enum SecretVersionStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum SecretVersionStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecretVersionStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecretVersionStatus(SecretVersionStatus(tmp).String())
	return nil
}

type PasswordGenerationParams struct {
	// Length: Length of the password to generate (between 1 and 1024).
	Length uint32 `json:"length"`
	// NoLowercaseLetters: Do not include lower case letters by default in the alphabet.
	NoLowercaseLetters bool `json:"no_lowercase_letters"`
	// NoUppercaseLetters: Do not include upper case letters by default in the alphabet.
	NoUppercaseLetters bool `json:"no_uppercase_letters"`
	// NoDigits: Do not include digits by default in the alphabet.
	NoDigits bool `json:"no_digits"`
	// AdditionalChars: Additional ascii characters to be included in the alphabet.
	AdditionalChars string `json:"additional_chars"`
}

type SecretVersion struct {
	// Revision: The first version of the secret is numbered 1, and all subsequent revisions augment by 1.
	Revision uint32 `json:"revision"`
	// SecretID: ID of the secret.
	SecretID string `json:"secret_id"`
	// Status: * `unknown`: the version is in an invalid state.
	// * `enabled`: the version is accessible.
	// * `disabled`: the version is not accessible but can be enabled.
	// * `destroyed`: the version is permanently deleted. It is not possible to recover it.
	Status SecretVersionStatus `json:"status"`
	// CreatedAt: Date and time of the version's creation.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Last update of the version.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Description: Description of the version.
	Description *string `json:"description,omitempty"`
	// IsLatest: Returns `true` if the version is the latest.
	IsLatest bool `json:"is_latest"`
}

type Secret struct {
	// ID: ID of the secret.
	ID string `json:"id"`
	// ProjectID: ID of the Project containing the secret.
	ProjectID string `json:"project_id"`
	// Name: Name of the secret.
	Name string `json:"name"`
	// Status: * `ready`: the secret can be read, modified and deleted.
	// * `locked`: no action can be performed on the secret. This status can only be applied and removed by Scaleway.
	Status SecretStatus `json:"status"`
	// CreatedAt: Date and time of the secret's creation.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Last update of the secret.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Tags: List of the secret's tags.
	Tags []string `json:"tags"`
	// VersionCount: Number of versions for this secret.
	VersionCount uint32 `json:"version_count"`
	// Description: Updated description of the secret.
	Description *string `json:"description,omitempty"`
	// IsManaged: Returns `true` for secrets that are managed by another product.
	IsManaged bool `json:"is_managed"`
	// IsProtected: Returns `true` for protected secrets that cannot be deleted.
	IsProtected bool `json:"is_protected"`
	// Type: See `Secret.Type` enum for description of values.
	Type SecretType `json:"type"`
	// Region: Region of the secret.
	Region scw.Region `json:"region"`
}

type AccessSecretVersionByNameRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SecretName: Name of the secret.
	SecretName string `json:"-"`
	// Revision: The first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be a number or "latest".
	Revision string `json:"-"`
	// ProjectID: (Optional.) If not specified, Secret Manager will look for the secret version in all Projects.
	ProjectID *string `json:"project_id,omitempty"`
}

type AccessSecretVersionRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
	// Revision: The first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be a number or "latest".
	Revision string `json:"-"`
}

type AccessSecretVersionResponse struct {
	// SecretID: ID of the secret.
	SecretID string `json:"secret_id"`
	// Revision: The first version of the secret is numbered 1, and all subsequent revisions augment by 1.
	Revision uint32 `json:"revision"`
	// Data: The base64-encoded secret payload of the version.
	Data []byte `json:"data"`
	// DataCrc32: This field is only available if a CRC32 was supplied during the creation of the version.
	DataCrc32 *uint32 `json:"data_crc32,omitempty"`
}

type AddSecretOwnerRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
	// ProductName: (Deprecated: use `product` field) Name of the product to add.
	ProductName *string `json:"product_name,omitempty"`
	// Product: See `Product` enum for description of values.
	Product Product `json:"product"`
}

type CreateSecretRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID: ID of the Project containing the secret.
	ProjectID string `json:"project_id"`
	// Name: Name of the secret.
	Name string `json:"name"`
	// Tags: List of the secret's tags.
	Tags []string `json:"tags"`
	// Description: Description of the secret.
	Description *string `json:"description,omitempty"`
	// Type: (Optional.) See `Secret.Type` enum for description of values. If not specified, the type is `Opaque`.
	Type SecretType `json:"type"`
}

type CreateSecretVersionRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
	// Data: The base64-encoded secret payload of the version.
	Data []byte `json:"data"`
	// Description: Description of the version.
	Description *string `json:"description,omitempty"`
	// DisablePrevious: (Optional.) If there is no previous version or if the previous version was already disabled, does nothing.
	DisablePrevious *bool `json:"disable_previous,omitempty"`
	// PasswordGeneration: (Optional.) If specified, a random password will be generated. The `data` and `data_crc32` fields must be empty. By default, the generator will use upper and lower case letters, and digits. This behavior can be tuned using the generation parameters.
	PasswordGeneration *PasswordGenerationParams `json:"password_generation,omitempty"`
	// DataCrc32: If specified, Secret Manager will verify the integrity of the data received against the given CRC32 checksum. An error is returned if the CRC32 does not match. If, however, the CRC32 matches, it will be stored and returned along with the SecretVersion on future access requests.
	DataCrc32 *uint32 `json:"data_crc32,omitempty"`
}

type DeleteSecretRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
}

type DestroySecretVersionRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
	// Revision: The first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be a number or "latest".
	Revision string `json:"-"`
}

type DisableSecretVersionRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
	// Revision: The first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be a number or "latest".
	Revision string `json:"-"`
}

type EnableSecretVersionRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
	// Revision: The first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be a number or "latest".
	Revision string `json:"-"`
}

type GeneratePasswordRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
	// Description: Description of the version.
	Description *string `json:"description,omitempty"`
	// DisablePrevious: This has no effect if there is no previous version or if the previous version was already disabled.
	DisablePrevious *bool `json:"disable_previous,omitempty"`
	// Length: Length of the password to generate (between 1 and 1024 characters).
	Length uint32 `json:"length"`
	// NoLowercaseLetters: (Optional.) Exclude lower case letters by default in the password character set.
	NoLowercaseLetters *bool `json:"no_lowercase_letters,omitempty"`
	// NoUppercaseLetters: (Optional.) Exclude upper case letters by default in the password character set.
	NoUppercaseLetters *bool `json:"no_uppercase_letters,omitempty"`
	// NoDigits: (Optional.) Exclude digits by default in the password character set.
	NoDigits *bool `json:"no_digits,omitempty"`
	// AdditionalChars: (Optional.) Additional ASCII characters to be included in the password character set.
	AdditionalChars *string `json:"additional_chars,omitempty"`
}

type GetSecretByNameRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SecretName: Name of the secret.
	SecretName string `json:"-"`
	// ProjectID: (Optional.) If not specified, Secret Manager will look for the secret in all Projects.
	ProjectID *string `json:"project_id,omitempty"`
}

type GetSecretRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
}

type GetSecretVersionByNameRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SecretName: Name of the secret.
	SecretName string `json:"-"`
	// Revision: The first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be a number or "latest".
	Revision string `json:"-"`
	// ProjectID: (Optional.) If not specified, Secret Manager will look for the secret version in all Projects.
	ProjectID *string `json:"project_id,omitempty"`
}

type GetSecretVersionRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
	// Revision: The first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be a number or "latest".
	Revision string `json:"-"`
}

type ListSecretVersionsByNameRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SecretName: Name of the secret.
	SecretName string `json:"-"`
	// Page:
	Page *int32 `json:"page,omitempty"`
	// PageSize:
	PageSize *uint32 `json:"page_size,omitempty"`
	// Status: Filter results by status.
	Status []SecretVersionStatus `json:"status"`
	// ProjectID: (Optional.) If not specified, Secret Manager will look for the secret in all Projects.
	ProjectID *string `json:"project_id,omitempty"`
}

type ListSecretVersionsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
	// Page:
	Page *int32 `json:"page,omitempty"`
	// PageSize:
	PageSize *uint32 `json:"page_size,omitempty"`
	// Status: Filter results by status.
	Status []SecretVersionStatus `json:"status"`
}

type ListSecretVersionsResponse struct {
	// Versions: Single page of versions.
	Versions []*SecretVersion `json:"versions"`
	// TotalCount: Number of versions.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSecretVersionsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSecretVersionsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSecretVersionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Versions = append(r.Versions, results.Versions...)
	r.TotalCount += uint32(len(results.Versions))
	return uint32(len(results.Versions)), nil
}

type ListSecretsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// OrganizationID: Filter by Organization ID (optional).
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: Filter by Project ID (optional).
	ProjectID *string `json:"project_id,omitempty"`
	// OrderBy:
	OrderBy ListSecretsRequestOrderBy `json:"order_by"`
	// Page:
	Page *int32 `json:"page,omitempty"`
	// PageSize:
	PageSize *uint32 `json:"page_size,omitempty"`
	// Tags: List of tags to filter on (optional).
	Tags []string `json:"tags"`
	// Name: Filter by secret name (optional).
	Name *string `json:"name,omitempty"`
	// IsManaged: Filter by managed / not managed (optional).
	IsManaged *bool `json:"is_managed,omitempty"`
}

type ListSecretsResponse struct {
	// Secrets: Single page of secrets matching the requested criteria.
	Secrets []*Secret `json:"secrets"`
	// TotalCount: Count of all secrets matching the requested criteria.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSecretsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSecretsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSecretsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Secrets = append(r.Secrets, results.Secrets...)
	r.TotalCount += uint32(len(results.Secrets))
	return uint32(len(results.Secrets)), nil
}

type ListTagsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID: (Optional.) If not specified, Secret Manager will look for tags in all Projects.
	ProjectID *string `json:"project_id,omitempty"`
	// Page:
	Page *int32 `json:"page,omitempty"`
	// PageSize:
	PageSize *uint32 `json:"page_size,omitempty"`
}

type ListTagsResponse struct {
	// Tags: List of tags.
	Tags []string `json:"tags"`
	// TotalCount: Count of all tags matching the requested criteria.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTagsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTagsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListTagsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Tags = append(r.Tags, results.Tags...)
	r.TotalCount += uint32(len(results.Tags))
	return uint32(len(results.Tags)), nil
}

type ProtectSecretRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret to protect.
	SecretID string `json:"-"`
}

type UnprotectSecretRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret to unprotect.
	SecretID string `json:"-"`
}

type UpdateSecretRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
	// Name: Secret's updated name (optional).
	Name *string `json:"name,omitempty"`
	// Tags: Secret's updated list of tags (optional).
	Tags *[]string `json:"tags,omitempty"`
	// Description: Description of the secret.
	Description *string `json:"description,omitempty"`
}

type UpdateSecretVersionRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
	// Revision: The first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be a number or "latest".
	Revision string `json:"-"`
	// Description: Description of the version.
	Description *string `json:"description,omitempty"`
}

// (switchcolumn)
// (switchcolumn)
// ## Concepts
//
// Refer to our [dedicated concepts page](https://www.scaleway.com/en/docs/identity-and-access-management/secret-manager/concepts/) to find definitions of the different terms referring to Secret Manager.
//
// (switchcolumn)
// (switchcolumn)
// ## Quickstart
//
// 1. **Configure your environment variables.**
//
//	<Message type="note">
//	  This is an optional step that seeks to simplify your usage of the API.
//	</Message>
//
//	```bash
//	export SCW_ACCESS_KEY="<API access key>"
//	export SCW_SECRET_KEY="<API secret key>"
//	export SCW_PROJECT_ID="<Scaleway Project ID>
//	```
//
// 2. **Create a secret**. Run the following command to create a secret:
//
//	```bash
//	curl "https://api.scaleway.com/secret-manager/v1alpha1/regions/$REGION/secrets" \
//	  -H "Content-Type: application/json" \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  -d '{
//	    "name": "my-secret",
//	    "project_id": "'"$PROJECT_ID"'"
//	  }'
//	```
//
// 3. **Create a secret version**. Run the following command to create a version of your secret:
//
//	```bash
//	curl "https://api.scaleway.com/secret-manager/v1alpha1/regions/$REGION/secrets/<SECRET_ID>/versions" \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  -d "{\"data\":\"$(echo -n "p@sSw0Rd_" | base64)\"}"
//	```
//
// 4. **Access data from your latest secret version**. Run the following command to access the     data of your most recent secret version:
//
//	```bash
//	curl "https://api.scaleway.com/secret-manager/v1alpha1/regions/$REGION/secrets/<SECRET_ID>/versions/latest/access" \
//	  -H "Content-Type: application/json" \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY"
//	```
//
//	<Message type="note">
//	  Requests can either target a specific version or the latest.
//	</Message>
//
// (switchcolumn)
//
// <Message type="requirement">
//   - You have your [Organization and your Project ID](https://console.scaleway.com/project/settings)
//   - You have [created an API key](https://www.scaleway.com/en/docs/identity-and-access-management/iam/how-to/create-api-keys/)
//   - You have [installed `curl`](https://curl.se/download.html)
//   - You have created an [API key](https://www.scaleway.com/en/docs/identity-and-access-management/iam/how-to/create-api-keys/) and that the API key has sufficient [IAM permissions](https://www.scaleway.com/en/docs/identity-and-access-management/iam/reference-content/permission-sets/) to perform the actions described on this page
//
// </Message>
//
// (switchcolumn)
// ## Technical information
//
// ### Regions
//
// Scaleway's infrastructure spans different [regions and Availability Zones](https://www.scaleway.com/en/docs/console/my-account/reference-content/products-availability/).
//
// Secret Manager is available in the Paris, Amsterdam and Warsaw regions, which are represented by the following path parameters:
//
// - fr-par
// - nl-ams
// - pl-waw
//
// ## Technical limitations
//
// - Operations on secrets and versions are limited to CRUDL
// - A secret's payload size is limited to 64KiB
//
// ## Going further
//
// For more information about Secret Manager, you can check out the following pages:
//
// * [Secret Manager Documentation](https://www.scaleway.com/en/docs/identity-and-access-management/secret-manager/)
// * [Scaleway Slack Community](https://scaleway-community.slack.com/) join the #secret-manager channel
// * [Contact our support team](https://console.scaleway.com/support/tickets).
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

// CreateSecret: You must sepcify the `region` to create a secret.
func (s *API) CreateSecret(req *CreateSecretRequest, opts ...scw.RequestOption) (*Secret, error) {
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
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Secret

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSecret: Retrieve the metadata of a secret specified by the `region` and `secret_id` parameters.
func (s *API) GetSecret(req *GetSecretRequest, opts ...scw.RequestOption) (*Secret, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return nil, errors.New("field SecretID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "",
	}

	var resp Secret

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSecretByName: Retrieve the metadata of a secret specified by the `region` and `secret_name` parameters.
func (s *API) GetSecretByName(req *GetSecretByNameRequest, opts ...scw.RequestOption) (*Secret, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretName) == "" {
		return nil, errors.New("field SecretName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets-by-name/" + fmt.Sprint(req.SecretName) + "",
		Query:  query,
	}

	var resp Secret

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSecret: Edit a secret's metadata such as name, tag(s) and description. The secret to update is specified by the `secret_id` and `region` parameters.
func (s *API) UpdateSecret(req *UpdateSecretRequest, opts ...scw.RequestOption) (*Secret, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return nil, errors.New("field SecretID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Secret

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSecrets: Retrieve the list of secrets created within an Organization and/or Project. You must specify either the `organization_id` or the `project_id` and the `region`.
func (s *API) ListSecrets(req *ListSecretsRequest, opts ...scw.RequestOption) (*ListSecretsResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "is_managed", req.IsManaged)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets",
		Query:  query,
	}

	var resp ListSecretsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSecret: Delete a given secret specified by the `region` and `secret_id` parameters.
func (s *API) DeleteSecret(req *DeleteSecretRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return errors.New("field SecretID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ProtectSecret: Protect a given secret specified by the `secret_id` parameter. A protected secret can be read and modified but cannot be deleted.
func (s *API) ProtectSecret(req *ProtectSecretRequest, opts ...scw.RequestOption) (*Secret, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return nil, errors.New("field SecretID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/protect",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Secret

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnprotectSecret: Unprotect a given secret specified by the `secret_id` parameter. An unprotected secret can be read, modified and deleted.
func (s *API) UnprotectSecret(req *UnprotectSecretRequest, opts ...scw.RequestOption) (*Secret, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return nil, errors.New("field SecretID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/unprotect",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Secret

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddSecretOwner: Allow a product to use the secret.
func (s *API) AddSecretOwner(req *AddSecretOwnerRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return errors.New("field SecretID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/add-owner",
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

// CreateSecretVersion: Create a version of a given secret specified by the `region` and `secret_id` parameters.
func (s *API) CreateSecretVersion(req *CreateSecretVersionRequest, opts ...scw.RequestOption) (*SecretVersion, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return nil, errors.New("field SecretID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SecretVersion

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GeneratePassword: Generate a password for the given secret specified by the `region` and `secret_id` parameters. This will also create a new version of the secret that will store the password.
func (s *API) GeneratePassword(req *GeneratePasswordRequest, opts ...scw.RequestOption) (*SecretVersion, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return nil, errors.New("field SecretID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/generate-password",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SecretVersion

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSecretVersion: Retrieve the metadata of a secret's given version specified by the `region`, `secret_id` and `revision` parameters.
func (s *API) GetSecretVersion(req *GetSecretVersionRequest, opts ...scw.RequestOption) (*SecretVersion, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return nil, errors.New("field SecretID cannot be empty in request")
	}

	if fmt.Sprint(req.Revision) == "" {
		return nil, errors.New("field Revision cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "",
	}

	var resp SecretVersion

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSecretVersionByName: Retrieve the metadata of a secret's given version specified by the `region`, `secret_name`, `revision` and `project_id` parameters.
func (s *API) GetSecretVersionByName(req *GetSecretVersionByNameRequest, opts ...scw.RequestOption) (*SecretVersion, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretName) == "" {
		return nil, errors.New("field SecretName cannot be empty in request")
	}

	if fmt.Sprint(req.Revision) == "" {
		return nil, errors.New("field Revision cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets-by-name/" + fmt.Sprint(req.SecretName) + "/versions/" + fmt.Sprint(req.Revision) + "",
		Query:  query,
	}

	var resp SecretVersion

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSecretVersion: Edit the metadata of a secret's given version, specified by the `region`, `secret_id` and `revision` parameters.
func (s *API) UpdateSecretVersion(req *UpdateSecretVersionRequest, opts ...scw.RequestOption) (*SecretVersion, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return nil, errors.New("field SecretID cannot be empty in request")
	}

	if fmt.Sprint(req.Revision) == "" {
		return nil, errors.New("field Revision cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SecretVersion

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSecretVersions: Retrieve the list of a given secret's versions specified by the `secret_id` and `region` parameters.
func (s *API) ListSecretVersions(req *ListSecretVersionsRequest, opts ...scw.RequestOption) (*ListSecretVersionsResponse, error) {
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
	parameter.AddToQuery(query, "status", req.Status)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return nil, errors.New("field SecretID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions",
		Query:  query,
	}

	var resp ListSecretVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSecretVersionsByName: Retrieve the list of a given secret's versions specified by the `secret_name`,`region` and `project_id` parameters.
func (s *API) ListSecretVersionsByName(req *ListSecretVersionsByNameRequest, opts ...scw.RequestOption) (*ListSecretVersionsResponse, error) {
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
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretName) == "" {
		return nil, errors.New("field SecretName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets-by-name/" + fmt.Sprint(req.SecretName) + "/versions",
		Query:  query,
	}

	var resp ListSecretVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnableSecretVersion: Make a specific version accessible. You must specify the `region`, `secret_id` and `revision` parameters.
func (s *API) EnableSecretVersion(req *EnableSecretVersionRequest, opts ...scw.RequestOption) (*SecretVersion, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return nil, errors.New("field SecretID cannot be empty in request")
	}

	if fmt.Sprint(req.Revision) == "" {
		return nil, errors.New("field Revision cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "/enable",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SecretVersion

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DisableSecretVersion: Make a specific version inaccessible. You must specify the `region`, `secret_id` and `revision` parameters.
func (s *API) DisableSecretVersion(req *DisableSecretVersionRequest, opts ...scw.RequestOption) (*SecretVersion, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return nil, errors.New("field SecretID cannot be empty in request")
	}

	if fmt.Sprint(req.Revision) == "" {
		return nil, errors.New("field Revision cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "/disable",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SecretVersion

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccessSecretVersion: Access sensitive data in a secret's version specified by the `region`, `secret_id` and `revision` parameters.
func (s *API) AccessSecretVersion(req *AccessSecretVersionRequest, opts ...scw.RequestOption) (*AccessSecretVersionResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return nil, errors.New("field SecretID cannot be empty in request")
	}

	if fmt.Sprint(req.Revision) == "" {
		return nil, errors.New("field Revision cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "/access",
	}

	var resp AccessSecretVersionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccessSecretVersionByName: Access sensitive data in a secret's version specified by the `region`, `secret_name`, `revision` and `project_id` parameters.
func (s *API) AccessSecretVersionByName(req *AccessSecretVersionByNameRequest, opts ...scw.RequestOption) (*AccessSecretVersionResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretName) == "" {
		return nil, errors.New("field SecretName cannot be empty in request")
	}

	if fmt.Sprint(req.Revision) == "" {
		return nil, errors.New("field Revision cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets-by-name/" + fmt.Sprint(req.SecretName) + "/versions/" + fmt.Sprint(req.Revision) + "/access",
		Query:  query,
	}

	var resp AccessSecretVersionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DestroySecretVersion: Delete a secret's version and the sensitive data contained in it. Deleting a version is permanent and cannot be undone.
func (s *API) DestroySecretVersion(req *DestroySecretVersionRequest, opts ...scw.RequestOption) (*SecretVersion, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return nil, errors.New("field SecretID cannot be empty in request")
	}

	if fmt.Sprint(req.Revision) == "" {
		return nil, errors.New("field Revision cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "/destroy",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SecretVersion

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListTags: List all tags associated with secrets within a given Project.
func (s *API) ListTags(req *ListTagsRequest, opts ...scw.RequestOption) (*ListTagsResponse, error) {
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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/tags",
		Query:  query,
	}

	var resp ListTagsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
