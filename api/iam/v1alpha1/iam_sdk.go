// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package iam provides methods and message types of the iam v1alpha1 API.
package iam

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

type BearerType string

const (
	// Unknown bearer type.
	BearerTypeUnknownBearerType = BearerType("unknown_bearer_type")
	// User.
	BearerTypeUser = BearerType("user")
	// Application.
	BearerTypeApplication = BearerType("application")
)

func (enum BearerType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_bearer_type"
	}
	return string(enum)
}

func (enum BearerType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *BearerType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = BearerType(BearerType(tmp).String())
	return nil
}

type ListAPIKeysRequestOrderBy string

const (
	// Creation date ascending.
	ListAPIKeysRequestOrderByCreatedAtAsc = ListAPIKeysRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListAPIKeysRequestOrderByCreatedAtDesc = ListAPIKeysRequestOrderBy("created_at_desc")
	// Update date ascending.
	ListAPIKeysRequestOrderByUpdatedAtAsc = ListAPIKeysRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListAPIKeysRequestOrderByUpdatedAtDesc = ListAPIKeysRequestOrderBy("updated_at_desc")
	// Expiration date ascending.
	ListAPIKeysRequestOrderByExpiresAtAsc = ListAPIKeysRequestOrderBy("expires_at_asc")
	// Expiration date descending.
	ListAPIKeysRequestOrderByExpiresAtDesc = ListAPIKeysRequestOrderBy("expires_at_desc")
	// Access key ascending.
	ListAPIKeysRequestOrderByAccessKeyAsc = ListAPIKeysRequestOrderBy("access_key_asc")
	// Access key descending.
	ListAPIKeysRequestOrderByAccessKeyDesc = ListAPIKeysRequestOrderBy("access_key_desc")
)

func (enum ListAPIKeysRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListAPIKeysRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListAPIKeysRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListAPIKeysRequestOrderBy(ListAPIKeysRequestOrderBy(tmp).String())
	return nil
}

type ListApplicationsRequestOrderBy string

const (
	// Creation date ascending.
	ListApplicationsRequestOrderByCreatedAtAsc = ListApplicationsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListApplicationsRequestOrderByCreatedAtDesc = ListApplicationsRequestOrderBy("created_at_desc")
	// Update date ascending.
	ListApplicationsRequestOrderByUpdatedAtAsc = ListApplicationsRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListApplicationsRequestOrderByUpdatedAtDesc = ListApplicationsRequestOrderBy("updated_at_desc")
	// Name ascending.
	ListApplicationsRequestOrderByNameAsc = ListApplicationsRequestOrderBy("name_asc")
	// Name descending.
	ListApplicationsRequestOrderByNameDesc = ListApplicationsRequestOrderBy("name_desc")
)

func (enum ListApplicationsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListApplicationsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListApplicationsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListApplicationsRequestOrderBy(ListApplicationsRequestOrderBy(tmp).String())
	return nil
}

type ListGroupsRequestOrderBy string

const (
	// Creation date ascending.
	ListGroupsRequestOrderByCreatedAtAsc = ListGroupsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListGroupsRequestOrderByCreatedAtDesc = ListGroupsRequestOrderBy("created_at_desc")
	// Update date ascending.
	ListGroupsRequestOrderByUpdatedAtAsc = ListGroupsRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListGroupsRequestOrderByUpdatedAtDesc = ListGroupsRequestOrderBy("updated_at_desc")
	// Name ascending.
	ListGroupsRequestOrderByNameAsc = ListGroupsRequestOrderBy("name_asc")
	// Name descending.
	ListGroupsRequestOrderByNameDesc = ListGroupsRequestOrderBy("name_desc")
)

func (enum ListGroupsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListGroupsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListGroupsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListGroupsRequestOrderBy(ListGroupsRequestOrderBy(tmp).String())
	return nil
}

type ListJWTsRequestOrderBy string

const (
	// Creation date ascending.
	ListJWTsRequestOrderByCreatedAtAsc = ListJWTsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListJWTsRequestOrderByCreatedAtDesc = ListJWTsRequestOrderBy("created_at_desc")
	// Update date ascending.
	ListJWTsRequestOrderByUpdatedAtAsc = ListJWTsRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListJWTsRequestOrderByUpdatedAtDesc = ListJWTsRequestOrderBy("updated_at_desc")
)

func (enum ListJWTsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListJWTsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListJWTsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListJWTsRequestOrderBy(ListJWTsRequestOrderBy(tmp).String())
	return nil
}

type ListPermissionSetsRequestOrderBy string

const (
	// Name ascending.
	ListPermissionSetsRequestOrderByNameAsc = ListPermissionSetsRequestOrderBy("name_asc")
	// Name descending.
	ListPermissionSetsRequestOrderByNameDesc = ListPermissionSetsRequestOrderBy("name_desc")
	// Creation date ascending.
	ListPermissionSetsRequestOrderByCreatedAtAsc = ListPermissionSetsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListPermissionSetsRequestOrderByCreatedAtDesc = ListPermissionSetsRequestOrderBy("created_at_desc")
)

func (enum ListPermissionSetsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListPermissionSetsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPermissionSetsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPermissionSetsRequestOrderBy(ListPermissionSetsRequestOrderBy(tmp).String())
	return nil
}

type ListPoliciesRequestOrderBy string

const (
	// Policy name ascending.
	ListPoliciesRequestOrderByPolicyNameAsc = ListPoliciesRequestOrderBy("policy_name_asc")
	// Policy name descending.
	ListPoliciesRequestOrderByPolicyNameDesc = ListPoliciesRequestOrderBy("policy_name_desc")
	// Creation date ascending.
	ListPoliciesRequestOrderByCreatedAtAsc = ListPoliciesRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListPoliciesRequestOrderByCreatedAtDesc = ListPoliciesRequestOrderBy("created_at_desc")
)

func (enum ListPoliciesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "policy_name_asc"
	}
	return string(enum)
}

func (enum ListPoliciesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPoliciesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPoliciesRequestOrderBy(ListPoliciesRequestOrderBy(tmp).String())
	return nil
}

type ListQuotaRequestOrderBy string

const (
	// Name ascending.
	ListQuotaRequestOrderByNameAsc = ListQuotaRequestOrderBy("name_asc")
	// Name descending.
	ListQuotaRequestOrderByNameDesc = ListQuotaRequestOrderBy("name_desc")
)

func (enum ListQuotaRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListQuotaRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListQuotaRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListQuotaRequestOrderBy(ListQuotaRequestOrderBy(tmp).String())
	return nil
}

type ListSSHKeysRequestOrderBy string

const (
	// Creation date ascending.
	ListSSHKeysRequestOrderByCreatedAtAsc = ListSSHKeysRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListSSHKeysRequestOrderByCreatedAtDesc = ListSSHKeysRequestOrderBy("created_at_desc")
	// Update date ascending.
	ListSSHKeysRequestOrderByUpdatedAtAsc = ListSSHKeysRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListSSHKeysRequestOrderByUpdatedAtDesc = ListSSHKeysRequestOrderBy("updated_at_desc")
	// Name ascending.
	ListSSHKeysRequestOrderByNameAsc = ListSSHKeysRequestOrderBy("name_asc")
	// Name descending.
	ListSSHKeysRequestOrderByNameDesc = ListSSHKeysRequestOrderBy("name_desc")
)

func (enum ListSSHKeysRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListSSHKeysRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSSHKeysRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSSHKeysRequestOrderBy(ListSSHKeysRequestOrderBy(tmp).String())
	return nil
}

type ListUsersRequestOrderBy string

const (
	// Creation date ascending.
	ListUsersRequestOrderByCreatedAtAsc = ListUsersRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListUsersRequestOrderByCreatedAtDesc = ListUsersRequestOrderBy("created_at_desc")
	// Update date ascending.
	ListUsersRequestOrderByUpdatedAtAsc = ListUsersRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListUsersRequestOrderByUpdatedAtDesc = ListUsersRequestOrderBy("updated_at_desc")
	// Email ascending.
	ListUsersRequestOrderByEmailAsc = ListUsersRequestOrderBy("email_asc")
	// Email descending.
	ListUsersRequestOrderByEmailDesc = ListUsersRequestOrderBy("email_desc")
	// Last login ascending.
	ListUsersRequestOrderByLastLoginAsc = ListUsersRequestOrderBy("last_login_asc")
	// Last login descending.
	ListUsersRequestOrderByLastLoginDesc = ListUsersRequestOrderBy("last_login_desc")
)

func (enum ListUsersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListUsersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListUsersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListUsersRequestOrderBy(ListUsersRequestOrderBy(tmp).String())
	return nil
}

type PermissionSetScopeType string

const (
	// Unknown scope type.
	PermissionSetScopeTypeUnknownScopeType = PermissionSetScopeType("unknown_scope_type")
	// Projects.
	PermissionSetScopeTypeProjects = PermissionSetScopeType("projects")
	// Organization.
	PermissionSetScopeTypeOrganization = PermissionSetScopeType("organization")
	// Account root user.
	PermissionSetScopeTypeAccountRootUser = PermissionSetScopeType("account_root_user")
)

func (enum PermissionSetScopeType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_scope_type"
	}
	return string(enum)
}

func (enum PermissionSetScopeType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PermissionSetScopeType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PermissionSetScopeType(PermissionSetScopeType(tmp).String())
	return nil
}

type UserStatus string

const (
	// Unknown status.
	UserStatusUnknownStatus = UserStatus("unknown_status")
	// Invitation pending.
	UserStatusInvitationPending = UserStatus("invitation_pending")
	// Activated.
	UserStatusActivated = UserStatus("activated")
)

func (enum UserStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum UserStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *UserStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = UserStatus(UserStatus(tmp).String())
	return nil
}

type UserType string

const (
	// Unknown type.
	UserTypeUnknownType = UserType("unknown_type")
	// Guest.
	UserTypeGuest = UserType("guest")
	// Owner.
	UserTypeOwner = UserType("owner")
)

func (enum UserType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum UserType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *UserType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = UserType(UserType(tmp).String())
	return nil
}

// RuleSpecs:
type RuleSpecs struct {
	// PermissionSetNames: Names of permission sets bound to the rule.
	PermissionSetNames *[]string `json:"permission_set_names"`
	// ProjectIDs: List of Project IDs the rule is scoped to.
	ProjectIDs *[]string `json:"project_ids,omitempty"`
	// OrganizationID: ID of Organization the rule is scoped to.
	OrganizationID *string `json:"organization_id,omitempty"`
}

// JWT:
type JWT struct {
	// Jti: JWT ID.
	Jti string `json:"jti"`
	// IssuerID: ID of the user who issued the JWT.
	IssuerID string `json:"issuer_id"`
	// AudienceID: ID of the user targeted by the JWT.
	AudienceID string `json:"audience_id"`
	// CreatedAt: Creation date of the JWT.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Last update date of the JWT.
	UpdatedAt *time.Time `json:"updated_at"`
	// ExpiresAt: Expiration date of the JWT.
	ExpiresAt *time.Time `json:"expires_at"`
	// IP: IP address used during the creation of the JWT.
	IP net.IP `json:"ip"`
	// UserAgent: User-agent used during the creation of the JWT.
	UserAgent string `json:"user_agent"`
}

// APIKey:
type APIKey struct {
	// AccessKey: Access key of the API key.
	AccessKey string `json:"access_key"`
	// SecretKey: Secret key of the API Key.
	SecretKey *string `json:"secret_key"`
	// ApplicationID: ID of application that bears the API key.
	ApplicationID *string `json:"application_id,omitempty"`
	// UserID: ID of user that bears the API key.
	UserID *string `json:"user_id,omitempty"`
	// Description: Description of API key.
	Description string `json:"description"`
	// CreatedAt: Date and time of API key creation.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Date and time of last API key update.
	UpdatedAt *time.Time `json:"updated_at"`
	// ExpiresAt: Date and time of API key expiration.
	ExpiresAt *time.Time `json:"expires_at"`
	// DefaultProjectID: Default Project ID specified for this API key.
	DefaultProjectID string `json:"default_project_id"`
	// Editable: Defines whether or not the API key is editable.
	Editable bool `json:"editable"`
	// CreationIP: IP address of the device that created the API key.
	CreationIP string `json:"creation_ip"`
}

// Application:
type Application struct {
	// ID: ID of the application.
	ID string `json:"id"`
	// Name: Name of the application.
	Name string `json:"name"`
	// Description: Description of the application.
	Description string `json:"description"`
	// CreatedAt: Date and time application was created.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Date and time of last application update.
	UpdatedAt *time.Time `json:"updated_at"`
	// OrganizationID: ID of the Organization.
	OrganizationID string `json:"organization_id"`
	// Editable: Defines whether or not the application is editable.
	Editable bool `json:"editable"`
	// NbAPIKeys: Number of API keys attributed to the application.
	NbAPIKeys uint32 `json:"nb_api_keys"`
}

// Group:
type Group struct {
	// ID: ID of the group.
	ID string `json:"id"`
	// CreatedAt: Date and time of group creation.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Date and time of last group update.
	UpdatedAt *time.Time `json:"updated_at"`
	// OrganizationID: ID of Organization linked to the group.
	OrganizationID string `json:"organization_id"`
	// Name: Name of the group.
	Name string `json:"name"`
	// Description: Description of the group.
	Description string `json:"description"`
	// UserIDs: IDs of users attached to this group.
	UserIDs []string `json:"user_ids"`
	// ApplicationIDs: IDs of applications attached to this group.
	ApplicationIDs []string `json:"application_ids"`
}

// PermissionSet:
type PermissionSet struct {
	// ID: Id of the permission set.
	ID string `json:"id"`
	// Name: Name of the permission set.
	Name string `json:"name"`
	// ScopeType: Scope of the permission set.
	ScopeType PermissionSetScopeType `json:"scope_type"`
	// Description: Description of the permission set.
	Description string `json:"description"`
	// Categories: Categories of the permission set.
	Categories *[]string `json:"categories"`
}

// Policy:
type Policy struct {
	// ID: Id of the policy.
	ID string `json:"id"`
	// Name: Name of the policy.
	Name string `json:"name"`
	// Description: Description of the policy.
	Description string `json:"description"`
	// OrganizationID: Organization ID of the policy.
	OrganizationID string `json:"organization_id"`
	// CreatedAt: Date and time of policy creation.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Date and time of last policy update.
	UpdatedAt *time.Time `json:"updated_at"`
	// Editable: Defines whether or not a policy is editable.
	Editable bool `json:"editable"`
	// NbRules: Number of rules of the policy.
	NbRules uint32 `json:"nb_rules"`
	// NbScopes: Number of policy scopes.
	NbScopes uint32 `json:"nb_scopes"`
	// NbPermissionSets: Number of permission sets of the policy.
	NbPermissionSets uint32 `json:"nb_permission_sets"`
	// UserID: ID of the user attributed to the policy.
	UserID *string `json:"user_id,omitempty"`
	// GroupID: ID of the group attributed to the policy.
	GroupID *string `json:"group_id,omitempty"`
	// ApplicationID: ID of the application attributed to the policy.
	ApplicationID *string `json:"application_id,omitempty"`
	// NoPrincipal: Defines whether or not a policy is attributed to a principal.
	NoPrincipal *bool `json:"no_principal,omitempty"`
}

// Quotum:
type Quotum struct {
	// Name: Name of the quota.
	Name string `json:"name"`
	// Limit: Maximum limit of the quota.
	Limit *uint64 `json:"limit,omitempty"`
	// Unlimited: Defines whether or not the quota is unlimited.
	Unlimited *bool `json:"unlimited,omitempty"`
}

// Rule:
type Rule struct {
	// ID: Id of rule.
	ID string `json:"id"`
	// PermissionSetNames: Names of permission sets bound to the rule.
	PermissionSetNames *[]string `json:"permission_set_names"`
	// PermissionSetsScopeType: Permission_set_names have the same scope_type.
	PermissionSetsScopeType PermissionSetScopeType `json:"permission_sets_scope_type"`
	// ProjectIDs: List of Project IDs the rule is scoped to.
	ProjectIDs *[]string `json:"project_ids,omitempty"`
	// OrganizationID: ID of Organization the rule is scoped to.
	OrganizationID *string `json:"organization_id,omitempty"`
	// AccountRootUserID: ID of account root user the rule is scoped to.
	AccountRootUserID *string `json:"account_root_user_id,omitempty"`
}

// SSHKey:
type SSHKey struct {
	// ID: ID of SSH key.
	ID string `json:"id"`
	// Name: Name of SSH key.
	Name string `json:"name"`
	// PublicKey: Public key of SSH key.
	PublicKey string `json:"public_key"`
	// Fingerprint: Fingerprint of the SSH key.
	Fingerprint string `json:"fingerprint"`
	// CreatedAt: Creation date of SSH key.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Last update date of SSH key.
	UpdatedAt *time.Time `json:"updated_at"`
	// OrganizationID: ID of Organization linked to the SSH key.
	OrganizationID string `json:"organization_id"`
	// ProjectID: ID of Project linked to the SSH key.
	ProjectID string `json:"project_id"`
	// Disabled: SSH key status.
	Disabled bool `json:"disabled"`
}

// User:
type User struct {
	// ID: ID of user.
	ID string `json:"id"`
	// Email: Email of user.
	Email string `json:"email"`
	// CreatedAt: Date user was created.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Date of last user update.
	UpdatedAt *time.Time `json:"updated_at"`
	// OrganizationID: ID of the Organization.
	OrganizationID string `json:"organization_id"`
	// Deletable: Deletion status of user. Owners cannot be deleted.
	Deletable bool `json:"deletable"`
	// LastLoginAt: Date of the last login.
	LastLoginAt *time.Time `json:"last_login_at"`
	// Type: Type of user.
	Type UserType `json:"type"`
	// Deprecated: TwoFactorEnabled: Deprecated, use "mfa" instead.
	TwoFactorEnabled *bool `json:"two_factor_enabled,omitempty"`
	// Status: Status of user invitation.
	Status UserStatus `json:"status"`
	// Mfa: Defines whether MFA is enabled.
	Mfa bool `json:"mfa"`
	// AccountRootUserID: ID of the account root user associated with the user.
	AccountRootUserID string `json:"account_root_user_id"`
}

// AddGroupMemberRequest:
type AddGroupMemberRequest struct {
	// GroupID: ID of the group.
	GroupID string `json:"-"`
	// UserID: ID of the user to add.
	UserID *string `json:"user_id,omitempty"`
	// ApplicationID: ID of the application to add.
	ApplicationID *string `json:"application_id,omitempty"`
}

// AddGroupMembersRequest:
type AddGroupMembersRequest struct {
	// GroupID: ID of the group.
	GroupID string `json:"-"`
	// UserIDs: IDs of the users to add.
	UserIDs []string `json:"user_ids"`
	// ApplicationIDs: IDs of the applications to add.
	ApplicationIDs []string `json:"application_ids"`
}

// ClonePolicyRequest:
type ClonePolicyRequest struct {
	// PolicyID:
	PolicyID string `json:"-"`
}

// CreateAPIKeyRequest:
type CreateAPIKeyRequest struct {
	// ApplicationID: ID of the application.
	ApplicationID *string `json:"application_id,omitempty"`
	// UserID: ID of the user.
	UserID *string `json:"user_id,omitempty"`
	// ExpiresAt: Expiration date of the API key.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// DefaultProjectID: Default Project ID to use with Object Storage.
	DefaultProjectID *string `json:"default_project_id,omitempty"`
	// Description: Description of the API key (max length is 200 characters).
	Description string `json:"description"`
}

// CreateApplicationRequest:
type CreateApplicationRequest struct {
	// Name: Name of the application to create (max length is 64 characters).
	Name string `json:"name"`
	// OrganizationID: ID of the Organization.
	OrganizationID string `json:"organization_id"`
	// Description: Description of the application (max length is 200 characters).
	Description string `json:"description"`
}

// CreateGroupRequest:
type CreateGroupRequest struct {
	// OrganizationID: ID of Organization linked to the group.
	OrganizationID string `json:"organization_id"`
	// Name: Name of the group to create (max length is 64 chars). MUST be unique inside an Organization.
	Name string `json:"name"`
	// Description: Description of the group to create (max length is 200 chars).
	Description string `json:"description"`
}

// CreatePolicyRequest:
type CreatePolicyRequest struct {
	// Name: Name of the policy to create (max length is 64 characters).
	Name string `json:"name"`
	// Description: Description of the policy to create (max length is 200 characters).
	Description string `json:"description"`
	// OrganizationID: ID of the Organization.
	OrganizationID string `json:"organization_id"`
	// Rules: Rules of the policy to create.
	Rules []*RuleSpecs `json:"rules"`
	// UserID: ID of user attributed to the policy.
	UserID *string `json:"user_id,omitempty"`
	// GroupID: ID of group attributed to the policy.
	GroupID *string `json:"group_id,omitempty"`
	// ApplicationID: ID of application attributed to the policy.
	ApplicationID *string `json:"application_id,omitempty"`
	// NoPrincipal: Defines whether or not a policy is attributed to a principal.
	NoPrincipal *bool `json:"no_principal,omitempty"`
}

// CreateSSHKeyRequest:
type CreateSSHKeyRequest struct {
	// Name: Name of the SSH key. Max length is 1000.
	Name string `json:"name"`
	// PublicKey: SSH public key. Currently only the ssh-rsa, ssh-dss (DSA), ssh-ed25519 and ecdsa keys with NIST curves are supported. Max length is 65000.
	PublicKey string `json:"public_key"`
	// ProjectID: Project the resource is attributed to.
	ProjectID string `json:"project_id"`
}

// CreateUserRequest:
type CreateUserRequest struct {
	// OrganizationID: ID of the Organization.
	OrganizationID string `json:"organization_id"`
	// Email: Email of the user.
	Email string `json:"email"`
}

// DeleteAPIKeyRequest:
type DeleteAPIKeyRequest struct {
	// AccessKey: Access key to delete.
	AccessKey string `json:"-"`
}

// DeleteApplicationRequest:
type DeleteApplicationRequest struct {
	// ApplicationID: ID of the application to delete.
	ApplicationID string `json:"-"`
}

// DeleteGroupRequest:
type DeleteGroupRequest struct {
	// GroupID: ID of the group to delete.
	GroupID string `json:"-"`
}

// DeleteJWTRequest:
type DeleteJWTRequest struct {
	// Jti: JWT ID of the JWT to delete.
	Jti string `json:"-"`
}

// DeletePolicyRequest:
type DeletePolicyRequest struct {
	// PolicyID: Id of policy to delete.
	PolicyID string `json:"-"`
}

// DeleteSSHKeyRequest:
type DeleteSSHKeyRequest struct {
	// SSHKeyID:
	SSHKeyID string `json:"-"`
}

// DeleteUserRequest:
type DeleteUserRequest struct {
	// UserID: ID of the user to delete.
	UserID string `json:"-"`
}

// GetAPIKeyRequest:
type GetAPIKeyRequest struct {
	// AccessKey: Access key to search for.
	AccessKey string `json:"-"`
}

// GetApplicationRequest:
type GetApplicationRequest struct {
	// ApplicationID: ID of the application to find.
	ApplicationID string `json:"-"`
}

// GetGroupRequest:
type GetGroupRequest struct {
	// GroupID: ID of the group.
	GroupID string `json:"-"`
}

// GetJWTRequest:
type GetJWTRequest struct {
	// Jti: JWT ID of the JWT to get.
	Jti string `json:"-"`
}

// GetPolicyRequest:
type GetPolicyRequest struct {
	// PolicyID: Id of policy to search.
	PolicyID string `json:"-"`
}

// GetQuotumRequest:
type GetQuotumRequest struct {
	// QuotumName: Name of the quota to get.
	QuotumName string `json:"-"`
	// OrganizationID: ID of the Organization.
	OrganizationID string `json:"-"`
}

// GetSSHKeyRequest:
type GetSSHKeyRequest struct {
	// SSHKeyID: ID of the SSH key.
	SSHKeyID string `json:"-"`
}

// GetUserRequest:
type GetUserRequest struct {
	// UserID: ID of the user to find.
	UserID string `json:"-"`
}

// ListAPIKeysRequest:
type ListAPIKeysRequest struct {
	// OrderBy: Criteria for sorting results.
	OrderBy ListAPIKeysRequestOrderBy `json:"-"`
	// Page: Page number. Value must be greater or equal to 1.
	Page *int32 `json:"-"`
	// PageSize: Number of results per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`
	// OrganizationID: ID of Organization.
	OrganizationID *string `json:"-"`
	// Deprecated: ApplicationID: ID of application that bears the API key.
	ApplicationID *string `json:"application_id,omitempty"`
	// Deprecated: UserID: ID of user that bears the API key.
	UserID *string `json:"user_id,omitempty"`
	// Editable: Defines whether to filter out editable API keys or not.
	Editable *bool `json:"-"`
	// Expired: Defines whether to filter out expired API keys or not.
	Expired *bool `json:"-"`
	// AccessKey: Filter by access key.
	AccessKey *string `json:"-"`
	// Description: Filter by description.
	Description *string `json:"-"`
	// BearerID: Filter by bearer ID.
	BearerID *string `json:"-"`
	// BearerType: Filter by type of bearer.
	BearerType BearerType `json:"-"`
}

// ListAPIKeysResponse:
type ListAPIKeysResponse struct {
	// APIKeys: List of API keys.
	APIKeys []*APIKey `json:"api_keys"`
	// TotalCount: Total count of API Keys.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListAPIKeysResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListAPIKeysResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListAPIKeysResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.APIKeys = append(r.APIKeys, results.APIKeys...)
	r.TotalCount += uint32(len(results.APIKeys))
	return uint32(len(results.APIKeys)), nil
}

// ListApplicationsRequest:
type ListApplicationsRequest struct {
	// OrderBy: Criteria for sorting results.
	OrderBy ListApplicationsRequestOrderBy `json:"-"`
	// PageSize: Number of results per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`
	// Page: Page number. Value must be greater than 1.
	Page *int32 `json:"-"`
	// Name: Name of the application to filter.
	Name *string `json:"-"`
	// OrganizationID: ID of the Organization to filter.
	OrganizationID string `json:"-"`
	// Editable: Defines whether to filter out editable applications or not.
	Editable *bool `json:"-"`
	// ApplicationIDs: Filter by list of IDs.
	ApplicationIDs []string `json:"-"`
}

// ListApplicationsResponse:
type ListApplicationsResponse struct {
	// Applications: List of applications.
	Applications []*Application `json:"applications"`
	// TotalCount: Total count of applications.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListApplicationsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListApplicationsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListApplicationsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Applications = append(r.Applications, results.Applications...)
	r.TotalCount += uint32(len(results.Applications))
	return uint32(len(results.Applications)), nil
}

// ListGroupsRequest:
type ListGroupsRequest struct {
	// OrderBy: Sort order of groups.
	OrderBy ListGroupsRequestOrderBy `json:"-"`
	// Page: Requested page number. Value must be greater or equal to 1.
	Page *int32 `json:"-"`
	// PageSize: Number of items per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`
	// OrganizationID: Filter by Organization ID.
	OrganizationID string `json:"-"`
	// Name: Name of group to find.
	Name *string `json:"-"`
	// ApplicationIDs: Filter by a list of application IDs.
	ApplicationIDs []string `json:"-"`
	// UserIDs: Filter by a list of user IDs.
	UserIDs []string `json:"-"`
	// GroupIDs: Filter by a list of group IDs.
	GroupIDs []string `json:"-"`
}

// ListGroupsResponse:
type ListGroupsResponse struct {
	// Groups: List of groups.
	Groups []*Group `json:"groups"`
	// TotalCount: Total count of groups.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListGroupsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListGroupsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListGroupsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Groups = append(r.Groups, results.Groups...)
	r.TotalCount += uint32(len(results.Groups))
	return uint32(len(results.Groups)), nil
}

// ListJWTsRequest:
type ListJWTsRequest struct {
	// OrderBy: Criteria for sorting results.
	OrderBy ListJWTsRequestOrderBy `json:"-"`
	// AudienceID: ID of the user to search.
	AudienceID *string `json:"-"`
	// PageSize: Number of results per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`
	// Page: Page number. Value must be greater to 1.
	Page *int32 `json:"-"`
	// Expired: Filter out expired JWTs or not.
	Expired *bool `json:"-"`
}

// ListJWTsResponse:
type ListJWTsResponse struct {
	// Jwts:
	Jwts []*JWT `json:"jwts"`
	// TotalCount:
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListJWTsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListJWTsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListJWTsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Jwts = append(r.Jwts, results.Jwts...)
	r.TotalCount += uint64(len(results.Jwts))
	return uint64(len(results.Jwts)), nil
}

// ListPermissionSetsRequest:
type ListPermissionSetsRequest struct {
	// OrderBy: Criteria for sorting results.
	OrderBy ListPermissionSetsRequestOrderBy `json:"-"`
	// PageSize: Number of results per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`
	// Page: Page number. Value must be greater than 1.
	Page *int32 `json:"-"`
	// OrganizationID: Filter by Organization ID.
	OrganizationID string `json:"-"`
}

// ListPermissionSetsResponse:
type ListPermissionSetsResponse struct {
	// PermissionSets: List of permission sets.
	PermissionSets []*PermissionSet `json:"permission_sets"`
	// TotalCount: Total count of permission sets.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPermissionSetsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPermissionSetsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListPermissionSetsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PermissionSets = append(r.PermissionSets, results.PermissionSets...)
	r.TotalCount += uint32(len(results.PermissionSets))
	return uint32(len(results.PermissionSets)), nil
}

// ListPoliciesRequest:
type ListPoliciesRequest struct {
	// OrderBy: Criteria for sorting results.
	OrderBy ListPoliciesRequestOrderBy `json:"-"`
	// PageSize: Number of results per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`
	// Page: Page number. Value must be greater than 1.
	Page *int32 `json:"-"`
	// OrganizationID: ID of the Organization to filter.
	OrganizationID string `json:"-"`
	// Editable: Defines whether or not filter out editable policies.
	Editable *bool `json:"-"`
	// UserIDs: Defines whether or not to filter by list of user IDs.
	UserIDs []string `json:"-"`
	// GroupIDs: Defines whether or not to filter by list of group IDs.
	GroupIDs []string `json:"-"`
	// ApplicationIDs: Filter by a list of application IDs.
	ApplicationIDs []string `json:"-"`
	// NoPrincipal: Defines whether or not the policy is attributed to a principal.
	NoPrincipal *bool `json:"-"`
	// PolicyName: Name of the policy to fetch.
	PolicyName *string `json:"-"`
}

// ListPoliciesResponse:
type ListPoliciesResponse struct {
	// Policies: List of policies.
	Policies []*Policy `json:"policies"`
	// TotalCount: Total count of policies.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPoliciesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPoliciesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListPoliciesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Policies = append(r.Policies, results.Policies...)
	r.TotalCount += uint32(len(results.Policies))
	return uint32(len(results.Policies)), nil
}

// ListQuotaRequest:
type ListQuotaRequest struct {
	// OrderBy: Criteria for sorting results.
	OrderBy ListQuotaRequestOrderBy `json:"-"`
	// PageSize: Number of results per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`
	// Page: Page number. Value must be greater than 1.
	Page *int32 `json:"-"`
	// OrganizationID: Filter by Organization ID.
	OrganizationID string `json:"-"`
}

// ListQuotaResponse:
type ListQuotaResponse struct {
	// Quota: List of quota.
	Quota []*Quotum `json:"quota"`
	// TotalCount: Total count of quota.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListQuotaResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListQuotaResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListQuotaResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Quota = append(r.Quota, results.Quota...)
	r.TotalCount += uint64(len(results.Quota))
	return uint64(len(results.Quota)), nil
}

// ListRulesRequest:
type ListRulesRequest struct {
	// PolicyID: Id of policy to search.
	PolicyID string `json:"-"`
	// PageSize: Number of results per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`
	// Page: Page number. Value must be greater than 1.
	Page *int32 `json:"-"`
}

// ListRulesResponse:
type ListRulesResponse struct {
	// Rules: Rules of the policy.
	Rules []*Rule `json:"rules"`
	// TotalCount: Total count of rules.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRulesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRulesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListRulesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Rules = append(r.Rules, results.Rules...)
	r.TotalCount += uint32(len(results.Rules))
	return uint32(len(results.Rules)), nil
}

// ListSSHKeysRequest:
type ListSSHKeysRequest struct {
	// OrderBy: Sort order of the SSH keys.
	OrderBy ListSSHKeysRequestOrderBy `json:"-"`
	// Page: Requested page number. Value must be greater or equal to 1.
	Page *int32 `json:"-"`
	// PageSize: Number of items per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`
	// OrganizationID: Filter by Organization ID.
	OrganizationID *string `json:"-"`
	// Name: Name of group to find.
	Name *string `json:"-"`
	// ProjectID: Filter by Project ID.
	ProjectID *string `json:"-"`
	// Disabled: Defines whether to include disabled SSH keys or not.
	Disabled *bool `json:"-"`
}

// ListSSHKeysResponse:
type ListSSHKeysResponse struct {
	// SSHKeys: List of SSH keys.
	SSHKeys []*SSHKey `json:"ssh_keys"`
	// TotalCount: Total count of SSH keys.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSSHKeysResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSSHKeysResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSSHKeysResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.SSHKeys = append(r.SSHKeys, results.SSHKeys...)
	r.TotalCount += uint32(len(results.SSHKeys))
	return uint32(len(results.SSHKeys)), nil
}

// ListUsersRequest:
type ListUsersRequest struct {
	// OrderBy: Criteria for sorting results.
	OrderBy ListUsersRequestOrderBy `json:"-"`
	// PageSize: Number of results per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`
	// Page: Page number. Value must be greater or equal to 1.
	Page *int32 `json:"-"`
	// OrganizationID: ID of the Organization to filter.
	OrganizationID *string `json:"-"`
	// UserIDs: Filter by list of IDs.
	UserIDs []string `json:"-"`
	// Mfa: Filter by MFA status.
	Mfa *bool `json:"-"`
}

// ListUsersResponse:
type ListUsersResponse struct {
	// Users: List of users.
	Users []*User `json:"users"`
	// TotalCount: Total count of users.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListUsersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListUsersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListUsersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Users = append(r.Users, results.Users...)
	r.TotalCount += uint32(len(results.Users))
	return uint32(len(results.Users)), nil
}

// RemoveGroupMemberRequest:
type RemoveGroupMemberRequest struct {
	// GroupID: ID of the group.
	GroupID string `json:"-"`
	// UserID: ID of the user to remove.
	UserID *string `json:"user_id,omitempty"`
	// ApplicationID: ID of the application to remove.
	ApplicationID *string `json:"application_id,omitempty"`
}

// SetGroupMembersRequest:
type SetGroupMembersRequest struct {
	// GroupID:
	GroupID string `json:"-"`
	// UserIDs:
	UserIDs []string `json:"user_ids"`
	// ApplicationIDs:
	ApplicationIDs []string `json:"application_ids"`
}

// SetRulesRequest:
type SetRulesRequest struct {
	// PolicyID: Id of policy to update.
	PolicyID string `json:"policy_id"`
	// Rules: Rules of the policy to set.
	Rules []*RuleSpecs `json:"rules"`
}

// SetRulesResponse:
type SetRulesResponse struct {
	// Rules: Rules of the policy.
	Rules []*Rule `json:"rules"`
}

// UpdateAPIKeyRequest:
type UpdateAPIKeyRequest struct {
	// AccessKey: Access key to update.
	AccessKey string `json:"-"`
	// DefaultProjectID: New default Project ID to set.
	DefaultProjectID *string `json:"default_project_id,omitempty"`
	// Description: New description to update.
	Description *string `json:"description,omitempty"`
}

// UpdateApplicationRequest:
type UpdateApplicationRequest struct {
	// ApplicationID: ID of the application to update.
	ApplicationID string `json:"-"`
	// Name: New name for the application (max length is 64 chars).
	Name *string `json:"name,omitempty"`
	// Description: New description for the application (max length is 200 chars).
	Description *string `json:"description,omitempty"`
}

// UpdateGroupRequest:
type UpdateGroupRequest struct {
	// GroupID: ID of the group to update.
	GroupID string `json:"-"`
	// Name: New name for the group (max length is 64 chars). MUST be unique inside an Organization.
	Name *string `json:"name,omitempty"`
	// Description: New description for the group (max length is 200 chars).
	Description *string `json:"description,omitempty"`
}

// UpdatePolicyRequest:
type UpdatePolicyRequest struct {
	// PolicyID: Id of policy to update.
	PolicyID string `json:"-"`
	// Name: New name for the policy (max length is 64 characters).
	Name *string `json:"name,omitempty"`
	// Description: New description of policy (max length is 200 characters).
	Description *string `json:"description,omitempty"`
	// UserID: New ID of user attributed to the policy.
	UserID *string `json:"user_id,omitempty"`
	// GroupID: New ID of group attributed to the policy.
	GroupID *string `json:"group_id,omitempty"`
	// ApplicationID: New ID of application attributed to the policy.
	ApplicationID *string `json:"application_id,omitempty"`
	// NoPrincipal: Defines whether or not the policy is attributed to a principal.
	NoPrincipal *bool `json:"no_principal,omitempty"`
}

// UpdateSSHKeyRequest:
type UpdateSSHKeyRequest struct {
	// SSHKeyID:
	SSHKeyID string `json:"-"`
	// Name: Name of the SSH key. Max length is 1000.
	Name *string `json:"name,omitempty"`
	// Disabled: Enable or disable the SSH key.
	Disabled *bool `json:"disabled,omitempty"`
}

// IAM API.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// ListSSHKeys: List SSH keys. By default, the SSH keys listed are ordered by creation date in ascending order. This can be modified via the `order_by` field. You can define additional parameters for your query such as `organization_id`, `name`, `project_id` and `disabled`.
func (s *API) ListSSHKeys(req *ListSSHKeysRequest, opts ...scw.RequestOption) (*ListSSHKeysResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "disabled", req.Disabled)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam/v1alpha1/ssh-keys",
		Query:  query,
	}

	var resp ListSSHKeysResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSSHKey: Add a new SSH key to a Scaleway Project. You must specify the `name`, `public_key` and `project_id`.
func (s *API) CreateSSHKey(req *CreateSSHKeyRequest, opts ...scw.RequestOption) (*SSHKey, error) {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("key")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam/v1alpha1/ssh-keys",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SSHKey

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSSHKey: Retrieve information about a given SSH key, specified by the `ssh_key_id` parameter. The SSH key's full details, including `id`, `name`, `public_key`, and `project_id` are returned in the response.
func (s *API) GetSSHKey(req *GetSSHKeyRequest, opts ...scw.RequestOption) (*SSHKey, error) {
	var err error

	if fmt.Sprint(req.SSHKeyID) == "" {
		return nil, errors.New("field SSHKeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam/v1alpha1/ssh-keys/" + fmt.Sprint(req.SSHKeyID) + "",
	}

	var resp SSHKey

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSSHKey: Update the parameters of an SSH key, including `name` and `disable`.
func (s *API) UpdateSSHKey(req *UpdateSSHKeyRequest, opts ...scw.RequestOption) (*SSHKey, error) {
	var err error

	if fmt.Sprint(req.SSHKeyID) == "" {
		return nil, errors.New("field SSHKeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/iam/v1alpha1/ssh-keys/" + fmt.Sprint(req.SSHKeyID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SSHKey

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSSHKey: Delete a given SSH key, specified by the `ssh_key_id`. Deleting an SSH is permanent, and cannot be undone. Note that you might need to update any configurations that used the SSH key.
func (s *API) DeleteSSHKey(req *DeleteSSHKeyRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.SSHKeyID) == "" {
		return errors.New("field SSHKeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iam/v1alpha1/ssh-keys/" + fmt.Sprint(req.SSHKeyID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListUsers: List the users of an Organization. By default, the users listed are ordered by creation date in ascending order. This can be modified via the `order_by` field. You must define the `organization_id` in the query path of your request. You can also define additional parameters for your query such as `user_ids`.
func (s *API) ListUsers(req *ListUsersRequest, opts ...scw.RequestOption) (*ListUsersResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "user_ids", req.UserIDs)
	parameter.AddToQuery(query, "mfa", req.Mfa)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam/v1alpha1/users",
		Query:  query,
	}

	var resp ListUsersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetUser: Retrieve information about a user, specified by the `user_id` parameter. The user's full details, including `id`, `email`, `organization_id`, `status` and `mfa` are returned in the response.
func (s *API) GetUser(req *GetUserRequest, opts ...scw.RequestOption) (*User, error) {
	var err error

	if fmt.Sprint(req.UserID) == "" {
		return nil, errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam/v1alpha1/users/" + fmt.Sprint(req.UserID) + "",
	}

	var resp User

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteUser: Remove a user from an Organization in which they are a guest. You must define the `user_id` in your request. Note that removing a user from an Organization automatically deletes their API keys, and any policies directly attached to them become orphaned.
func (s *API) DeleteUser(req *DeleteUserRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.UserID) == "" {
		return errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iam/v1alpha1/users/" + fmt.Sprint(req.UserID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CreateUser: Create a new user. You must define the `organization_id` and the `email` in your request.
func (s *API) CreateUser(req *CreateUserRequest, opts ...scw.RequestOption) (*User, error) {
	var err error
	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam/v1alpha1/users",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp User

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListApplications: List the applications of an Organization. By default, the applications listed are ordered by creation date in ascending order. This can be modified via the `order_by` field. You must define the `organization_id` in the query path of your request. You can also define additional parameters for your query such as `application_ids`.
func (s *API) ListApplications(req *ListApplicationsRequest, opts ...scw.RequestOption) (*ListApplicationsResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "editable", req.Editable)
	parameter.AddToQuery(query, "application_ids", req.ApplicationIDs)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam/v1alpha1/applications",
		Query:  query,
	}

	var resp ListApplicationsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateApplication: Create a new application. You must define the `name` parameter in the request.
func (s *API) CreateApplication(req *CreateApplicationRequest, opts ...scw.RequestOption) (*Application, error) {
	var err error
	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("app")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam/v1alpha1/applications",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Application

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetApplication: Retrieve information about an application, specified by the `application_id` parameter. The application's full details, including `id`, `email`, `organization_id`, `status` and `two_factor_enabled` are returned in the response.
func (s *API) GetApplication(req *GetApplicationRequest, opts ...scw.RequestOption) (*Application, error) {
	var err error

	if fmt.Sprint(req.ApplicationID) == "" {
		return nil, errors.New("field ApplicationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam/v1alpha1/applications/" + fmt.Sprint(req.ApplicationID) + "",
	}

	var resp Application

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateApplication: Update the parameters of an application, including `name` and `description`.
func (s *API) UpdateApplication(req *UpdateApplicationRequest, opts ...scw.RequestOption) (*Application, error) {
	var err error

	if fmt.Sprint(req.ApplicationID) == "" {
		return nil, errors.New("field ApplicationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/iam/v1alpha1/applications/" + fmt.Sprint(req.ApplicationID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Application

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteApplication: Delete an application. Note that this action is irreversible and will automatically delete the application's API keys. Policies attached to users and applications via this group will no longer apply.
func (s *API) DeleteApplication(req *DeleteApplicationRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.ApplicationID) == "" {
		return errors.New("field ApplicationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iam/v1alpha1/applications/" + fmt.Sprint(req.ApplicationID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListGroups: List groups. By default, the groups listed are ordered by creation date in ascending order. This can be modified via the `order_by` field. You can define additional parameters to filter your query. Use `user_ids` or `application_ids` to list all groups certain users or applications belong to.
func (s *API) ListGroups(req *ListGroupsRequest, opts ...scw.RequestOption) (*ListGroupsResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "application_ids", req.ApplicationIDs)
	parameter.AddToQuery(query, "user_ids", req.UserIDs)
	parameter.AddToQuery(query, "group_ids", req.GroupIDs)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam/v1alpha1/groups",
		Query:  query,
	}

	var resp ListGroupsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateGroup: Create a new group. You must define the `name` and `organization_id` parameters in the request.
func (s *API) CreateGroup(req *CreateGroupRequest, opts ...scw.RequestOption) (*Group, error) {
	var err error
	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("grp")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam/v1alpha1/groups",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Group

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGroup: Retrive information about a given group, specified by the `group_id` parameter. The group's full details, including `user_ids` and `application_ids` are returned in the response.
func (s *API) GetGroup(req *GetGroupRequest, opts ...scw.RequestOption) (*Group, error) {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return nil, errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam/v1alpha1/groups/" + fmt.Sprint(req.GroupID) + "",
	}

	var resp Group

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateGroup: Update the parameters of group, including `name` and `description`.
func (s *API) UpdateGroup(req *UpdateGroupRequest, opts ...scw.RequestOption) (*Group, error) {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return nil, errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/iam/v1alpha1/groups/" + fmt.Sprint(req.GroupID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Group

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetGroupMembers: Overwrite users and applications configuration in a group. Any information that you add using this command will overwrite the previous configuration.
func (s *API) SetGroupMembers(req *SetGroupMembersRequest, opts ...scw.RequestOption) (*Group, error) {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return nil, errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/iam/v1alpha1/groups/" + fmt.Sprint(req.GroupID) + "/members",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Group

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddGroupMember: Add a user or an application to a group. You can specify a `user_id` and and `application_id` in the body of your request. Note that you can only add one of each per request.
func (s *API) AddGroupMember(req *AddGroupMemberRequest, opts ...scw.RequestOption) (*Group, error) {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return nil, errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam/v1alpha1/groups/" + fmt.Sprint(req.GroupID) + "/add-member",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Group

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddGroupMembers: Add multiple users and applications to a group in a single call. You can specify an array of `user_id`s and `application_id`s. Note that any existing users and applications in the group will remain. To add new users/applications and delete pre-existing ones, use the [Overwrite users and applications of a group](#path-groups-overwrite-users-and-applications-of-a-group) method.
func (s *API) AddGroupMembers(req *AddGroupMembersRequest, opts ...scw.RequestOption) (*Group, error) {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return nil, errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam/v1alpha1/groups/" + fmt.Sprint(req.GroupID) + "/add-members",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Group

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveGroupMember: Remove a user or an application from a group. You can specify a `user_id` and and `application_id` in the body of your request. Note that you can only remove one of each per request. Removing a user from a group means that any permissions given to them via the group (i.e. from an attached policy) will no longer apply. Be sure you want to remove these permissions from the user before proceeding.
func (s *API) RemoveGroupMember(req *RemoveGroupMemberRequest, opts ...scw.RequestOption) (*Group, error) {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return nil, errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam/v1alpha1/groups/" + fmt.Sprint(req.GroupID) + "/remove-member",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Group

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteGroup: Delete a group. Note that this action is irreversible and could delete permissions for group members. Policies attached to users and applications via this group will no longer apply.
func (s *API) DeleteGroup(req *DeleteGroupRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iam/v1alpha1/groups/" + fmt.Sprint(req.GroupID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListPolicies: List the policies of an Organization. By default, the policies listed are ordered by creation date in ascending order. This can be modified via the `order_by` field. You must define the `organization_id` in the query path of your request. You can also define additional parameters to filter your query, such as `user_ids`, `groups_ids`, `application_ids`, and `policy_name`.
func (s *API) ListPolicies(req *ListPoliciesRequest, opts ...scw.RequestOption) (*ListPoliciesResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "editable", req.Editable)
	parameter.AddToQuery(query, "user_ids", req.UserIDs)
	parameter.AddToQuery(query, "group_ids", req.GroupIDs)
	parameter.AddToQuery(query, "application_ids", req.ApplicationIDs)
	parameter.AddToQuery(query, "no_principal", req.NoPrincipal)
	parameter.AddToQuery(query, "policy_name", req.PolicyName)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam/v1alpha1/policies",
		Query:  query,
	}

	var resp ListPoliciesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreatePolicy: Create a new application. You must define the `name` parameter in the request. You can specify parameters such as `user_id`, `groups_id`, `application_id`, `no_principal`, `rules` and its child attributes.
func (s *API) CreatePolicy(req *CreatePolicyRequest, opts ...scw.RequestOption) (*Policy, error) {
	var err error
	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("pol")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam/v1alpha1/policies",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Policy

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPolicy: Retrieve information about a policy, speficified by the `policy_id` parameter. The policy's full details, including `id`, `name`, `organization_id`, `nb_rules` and `nb_scopes`, `nb_permission_sets` are returned in the response.
func (s *API) GetPolicy(req *GetPolicyRequest, opts ...scw.RequestOption) (*Policy, error) {
	var err error

	if fmt.Sprint(req.PolicyID) == "" {
		return nil, errors.New("field PolicyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam/v1alpha1/policies/" + fmt.Sprint(req.PolicyID) + "",
	}

	var resp Policy

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePolicy: Update the parameters of a policy, including `name`, `description`, `user_id`, `group_id`, `application_id` and `no_principal`.
func (s *API) UpdatePolicy(req *UpdatePolicyRequest, opts ...scw.RequestOption) (*Policy, error) {
	var err error

	if fmt.Sprint(req.PolicyID) == "" {
		return nil, errors.New("field PolicyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/iam/v1alpha1/policies/" + fmt.Sprint(req.PolicyID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Policy

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeletePolicy: Delete a policy. You must define specify the `policy_id` parameter in your request. Note that when deleting a policy, all permissions it gives to its principal (user, group or application) will be revoked.
func (s *API) DeletePolicy(req *DeletePolicyRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.PolicyID) == "" {
		return errors.New("field PolicyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iam/v1alpha1/policies/" + fmt.Sprint(req.PolicyID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ClonePolicy: Clone a policy. You must define specify the `policy_id` parameter in your request.
func (s *API) ClonePolicy(req *ClonePolicyRequest, opts ...scw.RequestOption) (*Policy, error) {
	var err error

	if fmt.Sprint(req.PolicyID) == "" {
		return nil, errors.New("field PolicyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam/v1alpha1/policies/" + fmt.Sprint(req.PolicyID) + "/clone",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Policy

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetRules: Overwrite the rules of a given policy. Any information that you add using this command will overwrite the previous configuration. If you include some of the rules you already had in your previous configuration in your new one, but you change their order, the new order of display will apply. While policy rules are ordered, they have no impact on the access logic of IAM because rules are allow-only.
func (s *API) SetRules(req *SetRulesRequest, opts ...scw.RequestOption) (*SetRulesResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/iam/v1alpha1/rules",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListRules: List the rules of a given policy. By default, the rules listed are ordered by creation date in ascending order. This can be modified via the `order_by` field. You must define the `policy_id` in the query path of your request.
func (s *API) ListRules(req *ListRulesRequest, opts ...scw.RequestOption) (*ListRulesResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "policy_id", req.PolicyID)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam/v1alpha1/rules",
		Query:  query,
	}

	var resp ListRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPermissionSets: List permission sets available for given Organization. You must define the `organization_id` in the query path of your request.
func (s *API) ListPermissionSets(req *ListPermissionSetsRequest, opts ...scw.RequestOption) (*ListPermissionSetsResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam/v1alpha1/permission-sets",
		Query:  query,
	}

	var resp ListPermissionSetsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListAPIKeys: List API keys. By default, the API keys listed are ordered by creation date in ascending order. This can be modified via the `order_by` field. You can define additional parameters for your query such as `editable`, `expired`, `access_key` and `bearer_id`.
func (s *API) ListAPIKeys(req *ListAPIKeysRequest, opts ...scw.RequestOption) (*ListAPIKeysResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "editable", req.Editable)
	parameter.AddToQuery(query, "expired", req.Expired)
	parameter.AddToQuery(query, "access_key", req.AccessKey)
	parameter.AddToQuery(query, "description", req.Description)
	parameter.AddToQuery(query, "bearer_id", req.BearerID)
	parameter.AddToQuery(query, "bearer_type", req.BearerType)
	parameter.AddToQuery(query, "application_id", req.ApplicationID)
	parameter.AddToQuery(query, "user_id", req.UserID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam/v1alpha1/api-keys",
		Query:  query,
	}

	var resp ListAPIKeysResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateAPIKey: Create an API key. You must specify the `application_id` or the `user_id` and the description. You can also specify the `default_project_id` which is the Project ID of your preferred Project, to use with Object Storage. The `access_key` and `secret_key` values are returned in the response. Note that he secret key is only showed once. Make sure that you copy and store both keys somewhere safe.
func (s *API) CreateAPIKey(req *CreateAPIKeyRequest, opts ...scw.RequestOption) (*APIKey, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam/v1alpha1/api-keys",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp APIKey

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAPIKey: Retrive information about an API key, specified by the `access_key` parameter. The API key's details, including either the `user_id` or `application_id` of its bearer are returned in the response. Note that the string value for the `secret_key` is nullable, and therefore is not displayed in the response. The `secret_key` value is only displayed upon API key creation.
func (s *API) GetAPIKey(req *GetAPIKeyRequest, opts ...scw.RequestOption) (*APIKey, error) {
	var err error

	if fmt.Sprint(req.AccessKey) == "" {
		return nil, errors.New("field AccessKey cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam/v1alpha1/api-keys/" + fmt.Sprint(req.AccessKey) + "",
	}

	var resp APIKey

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAPIKey: Update the parameters of an API key, including `default_project_id` and `description`.
func (s *API) UpdateAPIKey(req *UpdateAPIKeyRequest, opts ...scw.RequestOption) (*APIKey, error) {
	var err error

	if fmt.Sprint(req.AccessKey) == "" {
		return nil, errors.New("field AccessKey cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/iam/v1alpha1/api-keys/" + fmt.Sprint(req.AccessKey) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp APIKey

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteAPIKey: Delete an API key. Note that this action is irreversible and cannot be undone. Make sure you update any configurations using the API keys you delete.
func (s *API) DeleteAPIKey(req *DeleteAPIKeyRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.AccessKey) == "" {
		return errors.New("field AccessKey cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iam/v1alpha1/api-keys/" + fmt.Sprint(req.AccessKey) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListQuota: List all product and features quota for an Organization, with their associated limits. By default, the quota listed are ordered by creation date in ascending order. This can be modified via the `order_by` field. You must define the `organization_id` in the query path of your request.
func (s *API) ListQuota(req *ListQuotaRequest, opts ...scw.RequestOption) (*ListQuotaResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam/v1alpha1/quota",
		Query:  query,
	}

	var resp ListQuotaResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetQuotum: Retrieve information about a resource quota, speficified by the `quotum_name` parameter. The quota's `limit`, or whether it is unlimited, is returned in the response.
func (s *API) GetQuotum(req *GetQuotumRequest, opts ...scw.RequestOption) (*Quotum, error) {
	var err error
	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	if fmt.Sprint(req.QuotumName) == "" {
		return nil, errors.New("field QuotumName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam/v1alpha1/quota/" + fmt.Sprint(req.QuotumName) + "",
		Query:  query,
	}

	var resp Quotum

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListJWTs: List JWTs.
func (s *API) ListJWTs(req *ListJWTsRequest, opts ...scw.RequestOption) (*ListJWTsResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "audience_id", req.AudienceID)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "expired", req.Expired)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam/v1alpha1/jwts",
		Query:  query,
	}

	var resp ListJWTsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetJWT: Get a JWT.
func (s *API) GetJWT(req *GetJWTRequest, opts ...scw.RequestOption) (*JWT, error) {
	var err error

	if fmt.Sprint(req.Jti) == "" {
		return nil, errors.New("field Jti cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam/v1alpha1/jwts/" + fmt.Sprint(req.Jti) + "",
	}

	var resp JWT

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteJWT: Delete a JWT.
func (s *API) DeleteJWT(req *DeleteJWTRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.Jti) == "" {
		return errors.New("field Jti cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iam/v1alpha1/jwts/" + fmt.Sprint(req.Jti) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}
