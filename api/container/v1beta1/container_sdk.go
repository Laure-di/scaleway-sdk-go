// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package container provides methods and message types of the container v1beta1 API.
package container

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

type ContainerHTTPOption string

const (
	ContainerHTTPOptionUnknownHTTPOption = ContainerHTTPOption("unknown_http_option")
	ContainerHTTPOptionEnabled           = ContainerHTTPOption("enabled")
	ContainerHTTPOptionRedirected        = ContainerHTTPOption("redirected")
)

func (enum ContainerHTTPOption) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_http_option"
	}
	return string(enum)
}

func (enum ContainerHTTPOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContainerHTTPOption) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContainerHTTPOption(ContainerHTTPOption(tmp).String())
	return nil
}

type ContainerPrivacy string

const (
	ContainerPrivacyUnknownPrivacy = ContainerPrivacy("unknown_privacy")
	ContainerPrivacyPublic         = ContainerPrivacy("public")
	ContainerPrivacyPrivate        = ContainerPrivacy("private")
)

func (enum ContainerPrivacy) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_privacy"
	}
	return string(enum)
}

func (enum ContainerPrivacy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContainerPrivacy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContainerPrivacy(ContainerPrivacy(tmp).String())
	return nil
}

type ContainerProtocol string

const (
	ContainerProtocolUnknownProtocol = ContainerProtocol("unknown_protocol")
	ContainerProtocolHTTP1           = ContainerProtocol("http1")
	ContainerProtocolH2c             = ContainerProtocol("h2c")
)

func (enum ContainerProtocol) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_protocol"
	}
	return string(enum)
}

func (enum ContainerProtocol) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContainerProtocol) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContainerProtocol(ContainerProtocol(tmp).String())
	return nil
}

type ContainerStatus string

const (
	ContainerStatusUnknown  = ContainerStatus("unknown")
	ContainerStatusReady    = ContainerStatus("ready")
	ContainerStatusDeleting = ContainerStatus("deleting")
	ContainerStatusError    = ContainerStatus("error")
	ContainerStatusLocked   = ContainerStatus("locked")
	ContainerStatusCreating = ContainerStatus("creating")
	ContainerStatusPending  = ContainerStatus("pending")
	ContainerStatusCreated  = ContainerStatus("created")
)

func (enum ContainerStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum ContainerStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContainerStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContainerStatus(ContainerStatus(tmp).String())
	return nil
}

type CronStatus string

const (
	CronStatusUnknown  = CronStatus("unknown")
	CronStatusReady    = CronStatus("ready")
	CronStatusDeleting = CronStatus("deleting")
	CronStatusError    = CronStatus("error")
	CronStatusLocked   = CronStatus("locked")
	CronStatusCreating = CronStatus("creating")
	CronStatusPending  = CronStatus("pending")
)

func (enum CronStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum CronStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CronStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CronStatus(CronStatus(tmp).String())
	return nil
}

type DomainStatus string

const (
	DomainStatusUnknown  = DomainStatus("unknown")
	DomainStatusReady    = DomainStatus("ready")
	DomainStatusDeleting = DomainStatus("deleting")
	DomainStatusError    = DomainStatus("error")
	DomainStatusCreating = DomainStatus("creating")
	DomainStatusPending  = DomainStatus("pending")
)

func (enum DomainStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
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

type ListContainersRequestOrderBy string

const (
	ListContainersRequestOrderByCreatedAtAsc  = ListContainersRequestOrderBy("created_at_asc")
	ListContainersRequestOrderByCreatedAtDesc = ListContainersRequestOrderBy("created_at_desc")
	ListContainersRequestOrderByNameAsc       = ListContainersRequestOrderBy("name_asc")
	ListContainersRequestOrderByNameDesc      = ListContainersRequestOrderBy("name_desc")
)

func (enum ListContainersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListContainersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListContainersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListContainersRequestOrderBy(ListContainersRequestOrderBy(tmp).String())
	return nil
}

type ListCronsRequestOrderBy string

const (
	ListCronsRequestOrderByCreatedAtAsc  = ListCronsRequestOrderBy("created_at_asc")
	ListCronsRequestOrderByCreatedAtDesc = ListCronsRequestOrderBy("created_at_desc")
)

func (enum ListCronsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListCronsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListCronsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListCronsRequestOrderBy(ListCronsRequestOrderBy(tmp).String())
	return nil
}

type ListDomainsRequestOrderBy string

const (
	ListDomainsRequestOrderByCreatedAtAsc  = ListDomainsRequestOrderBy("created_at_asc")
	ListDomainsRequestOrderByCreatedAtDesc = ListDomainsRequestOrderBy("created_at_desc")
	ListDomainsRequestOrderByHostnameAsc   = ListDomainsRequestOrderBy("hostname_asc")
	ListDomainsRequestOrderByHostnameDesc  = ListDomainsRequestOrderBy("hostname_desc")
)

func (enum ListDomainsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
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

type ListLogsRequestOrderBy string

const (
	ListLogsRequestOrderByTimestampDesc = ListLogsRequestOrderBy("timestamp_desc")
	ListLogsRequestOrderByTimestampAsc  = ListLogsRequestOrderBy("timestamp_asc")
)

func (enum ListLogsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "timestamp_desc"
	}
	return string(enum)
}

func (enum ListLogsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListLogsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListLogsRequestOrderBy(ListLogsRequestOrderBy(tmp).String())
	return nil
}

type ListNamespacesRequestOrderBy string

const (
	ListNamespacesRequestOrderByCreatedAtAsc  = ListNamespacesRequestOrderBy("created_at_asc")
	ListNamespacesRequestOrderByCreatedAtDesc = ListNamespacesRequestOrderBy("created_at_desc")
	ListNamespacesRequestOrderByNameAsc       = ListNamespacesRequestOrderBy("name_asc")
	ListNamespacesRequestOrderByNameDesc      = ListNamespacesRequestOrderBy("name_desc")
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

type ListTokensRequestOrderBy string

const (
	ListTokensRequestOrderByCreatedAtAsc  = ListTokensRequestOrderBy("created_at_asc")
	ListTokensRequestOrderByCreatedAtDesc = ListTokensRequestOrderBy("created_at_desc")
)

func (enum ListTokensRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListTokensRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListTokensRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListTokensRequestOrderBy(ListTokensRequestOrderBy(tmp).String())
	return nil
}

type ListTriggersRequestOrderBy string

const (
	ListTriggersRequestOrderByCreatedAtAsc  = ListTriggersRequestOrderBy("created_at_asc")
	ListTriggersRequestOrderByCreatedAtDesc = ListTriggersRequestOrderBy("created_at_desc")
)

func (enum ListTriggersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListTriggersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListTriggersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListTriggersRequestOrderBy(ListTriggersRequestOrderBy(tmp).String())
	return nil
}

type LogStream string

const (
	LogStreamUnknown = LogStream("unknown")
	LogStreamStdout  = LogStream("stdout")
	LogStreamStderr  = LogStream("stderr")
)

func (enum LogStream) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum LogStream) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *LogStream) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = LogStream(LogStream(tmp).String())
	return nil
}

type NamespaceStatus string

const (
	NamespaceStatusUnknown  = NamespaceStatus("unknown")
	NamespaceStatusReady    = NamespaceStatus("ready")
	NamespaceStatusDeleting = NamespaceStatus("deleting")
	NamespaceStatusError    = NamespaceStatus("error")
	NamespaceStatusLocked   = NamespaceStatus("locked")
	NamespaceStatusCreating = NamespaceStatus("creating")
	NamespaceStatusPending  = NamespaceStatus("pending")
)

func (enum NamespaceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum NamespaceStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NamespaceStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NamespaceStatus(NamespaceStatus(tmp).String())
	return nil
}

type TokenStatus string

const (
	TokenStatusUnknown  = TokenStatus("unknown")
	TokenStatusReady    = TokenStatus("ready")
	TokenStatusDeleting = TokenStatus("deleting")
	TokenStatusError    = TokenStatus("error")
	TokenStatusCreating = TokenStatus("creating")
)

func (enum TokenStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum TokenStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TokenStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TokenStatus(TokenStatus(tmp).String())
	return nil
}

type TriggerInputType string

const (
	TriggerInputTypeUnknownInputType = TriggerInputType("unknown_input_type")
	TriggerInputTypeSqs              = TriggerInputType("sqs")
	TriggerInputTypeScwSqs           = TriggerInputType("scw_sqs")
	TriggerInputTypeNats             = TriggerInputType("nats")
	TriggerInputTypeScwNats          = TriggerInputType("scw_nats")
)

func (enum TriggerInputType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_input_type"
	}
	return string(enum)
}

func (enum TriggerInputType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TriggerInputType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TriggerInputType(TriggerInputType(tmp).String())
	return nil
}

type TriggerStatus string

const (
	TriggerStatusUnknownStatus = TriggerStatus("unknown_status")
	TriggerStatusReady         = TriggerStatus("ready")
	TriggerStatusDeleting      = TriggerStatus("deleting")
	TriggerStatusError         = TriggerStatus("error")
	TriggerStatusCreating      = TriggerStatus("creating")
	TriggerStatusPending       = TriggerStatus("pending")
)

func (enum TriggerStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum TriggerStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TriggerStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TriggerStatus(TriggerStatus(tmp).String())
	return nil
}

// SecretHashedValue:
type SecretHashedValue struct {
	// Key:
	Key string `json:"key"`
	// HashedValue:
	HashedValue string `json:"hashed_value"`
}

// TriggerMnqNatsClientConfig:
type TriggerMnqNatsClientConfig struct {
	// MnqNamespaceID:
	MnqNamespaceID string `json:"mnq_namespace_id"`
	// Subject:
	Subject string `json:"subject"`
	// MnqProjectID:
	MnqProjectID string `json:"mnq_project_id"`
	// MnqRegion:
	MnqRegion string `json:"mnq_region"`
	// MnqCredentialID:
	MnqCredentialID *string `json:"mnq_credential_id"`
}

// TriggerMnqSqsClientConfig:
type TriggerMnqSqsClientConfig struct {
	// MnqNamespaceID:
	MnqNamespaceID string `json:"mnq_namespace_id"`
	// Queue:
	Queue string `json:"queue"`
	// MnqProjectID:
	MnqProjectID string `json:"mnq_project_id"`
	// MnqRegion:
	MnqRegion string `json:"mnq_region"`
	// MnqCredentialID:
	MnqCredentialID *string `json:"mnq_credential_id"`
}

// TriggerSqsClientConfig:
type TriggerSqsClientConfig struct {
	// Endpoint:
	Endpoint string `json:"endpoint"`
	// QueueURL:
	QueueURL string `json:"queue_url"`
	// AccessKey:
	AccessKey string `json:"access_key"`
	// SecretKey:
	SecretKey string `json:"secret_key"`
}

// Secret:
type Secret struct {
	// Key:
	Key string `json:"key"`
	// Value:
	Value *string `json:"value"`
}

// CreateTriggerRequestMnqNatsClientConfig:
type CreateTriggerRequestMnqNatsClientConfig struct {
	// MnqNamespaceID:
	MnqNamespaceID string `json:"mnq_namespace_id"`
	// Subject:
	Subject string `json:"subject"`
	// MnqProjectID:
	MnqProjectID string `json:"mnq_project_id"`
	// MnqRegion:
	MnqRegion string `json:"mnq_region"`
}

// CreateTriggerRequestMnqSqsClientConfig:
type CreateTriggerRequestMnqSqsClientConfig struct {
	// MnqNamespaceID:
	MnqNamespaceID string `json:"mnq_namespace_id"`
	// Queue:
	Queue string `json:"queue"`
	// MnqProjectID:
	MnqProjectID string `json:"mnq_project_id"`
	// MnqRegion:
	MnqRegion string `json:"mnq_region"`
}

// CreateTriggerRequestSqsClientConfig:
type CreateTriggerRequestSqsClientConfig struct {
	// Endpoint:
	Endpoint string `json:"endpoint"`
	// QueueURL:
	QueueURL string `json:"queue_url"`
	// AccessKey:
	AccessKey string `json:"access_key"`
	// SecretKey:
	SecretKey string `json:"secret_key"`
}

// Container:
type Container struct {
	// ID: UUID of the container.
	ID string `json:"id"`
	// Name: Name of the container.
	Name string `json:"name"`
	// NamespaceID: UUID of the namespace the container belongs to.
	NamespaceID string `json:"namespace_id"`
	// Status: Status of the container.
	Status ContainerStatus `json:"status"`
	// EnvironmentVariables: Environment variables of the container.
	EnvironmentVariables map[string]string `json:"environment_variables"`
	// MinScale: Minimum number of instances to scale the container to.
	MinScale uint32 `json:"min_scale"`
	// MaxScale: Maximum number of instances to scale the container to.
	MaxScale uint32 `json:"max_scale"`
	// MemoryLimit: Memory limit of the container in MB.
	MemoryLimit uint32 `json:"memory_limit"`
	// CPULimit: CPU limit of the container in mvCPU.
	CPULimit uint32 `json:"cpu_limit"`
	// Timeout: Processing time limit for the container.
	Timeout *scw.Duration `json:"timeout"`
	// ErrorMessage: Last error message of the container.
	ErrorMessage *string `json:"error_message"`
	// Privacy: Privacy setting of the container.
	Privacy ContainerPrivacy `json:"privacy"`
	// Description: Description of the container.
	Description *string `json:"description"`
	// RegistryImage: Name of the registry image (e.g. "rg.fr-par.scw.cloud/something/image:tag").
	RegistryImage string `json:"registry_image"`
	// MaxConcurrency: Number of maximum concurrent executions of the container.
	MaxConcurrency uint32 `json:"max_concurrency"`
	// DomainName: Domain name attributed to the contaioner.
	DomainName string `json:"domain_name"`
	// Protocol: Protocol the container uses.
	Protocol ContainerProtocol `json:"protocol"`
	// Port: Port the container listens on.
	Port uint32 `json:"port"`
	// SecretEnvironmentVariables: Secret environment variables of the container.
	SecretEnvironmentVariables []*SecretHashedValue `json:"secret_environment_variables"`
	// HTTPOption: Possible values:
	//  - redirected: Responds to HTTP request with a 301 redirect to ask the clients to use HTTPS.
	//  - enabled: Serve both HTTP and HTTPS traffic.
	HTTPOption ContainerHTTPOption `json:"http_option"`
	// Region: Region in which the container will be deployed.
	Region scw.Region `json:"region"`
}

// Cron:
type Cron struct {
	// ID: UUID of the cron.
	ID string `json:"id"`
	// ContainerID: UUID of the container invoked by this cron.
	ContainerID string `json:"container_id"`
	// Schedule: UNIX cron shedule.
	Schedule string `json:"schedule"`
	// Args: Arguments to pass with the cron.
	Args *scw.JSONObject `json:"args"`
	// Status: Status of the cron.
	Status CronStatus `json:"status"`
	// Name: Name of the cron.
	Name string `json:"name"`
}

// Domain:
type Domain struct {
	// ID: UUID of the domain.
	ID string `json:"id"`
	// Hostname: Domain assigned to the container.
	Hostname string `json:"hostname"`
	// ContainerID: UUID of the container.
	ContainerID string `json:"container_id"`
	// URL: URL (TBD).
	URL string `json:"url"`
	// Status: Status of the domain.
	Status DomainStatus `json:"status"`
	// ErrorMessage: Last error message of the domain.
	ErrorMessage *string `json:"error_message"`
}

// Log:
type Log struct {
	// Message:
	Message string `json:"message"`
	// Timestamp:
	Timestamp *time.Time `json:"timestamp"`
	// ID:
	ID string `json:"id"`
	// Level: Contains the severity of the log (info, debug, error, ...).
	Level string `json:"level"`
	// Source: Source of the log (core runtime or user code).
	Source string `json:"source"`
	// Stream: Can be stdout or stderr.
	Stream LogStream `json:"stream"`
}

// Namespace:
type Namespace struct {
	// ID: UUID of the namespace.
	ID string `json:"id"`
	// Name: Name of the namespace.
	Name string `json:"name"`
	// EnvironmentVariables: Environment variables of the namespace.
	EnvironmentVariables map[string]string `json:"environment_variables"`
	// OrganizationID: UUID of the Organization the namespace belongs to.
	OrganizationID string `json:"organization_id"`
	// ProjectID: UUID of the Project the namespace belongs to.
	ProjectID string `json:"project_id"`
	// Status: Status of the namespace.
	Status NamespaceStatus `json:"status"`
	// RegistryNamespaceID: UUID of the registry namespace.
	RegistryNamespaceID string `json:"registry_namespace_id"`
	// ErrorMessage: Last error message of the namesace.
	ErrorMessage *string `json:"error_message"`
	// RegistryEndpoint: Registry endpoint of the namespace.
	RegistryEndpoint string `json:"registry_endpoint"`
	// Description: Description of the endpoint.
	Description *string `json:"description"`
	// SecretEnvironmentVariables: Secret environment variables of the namespace.
	SecretEnvironmentVariables []*SecretHashedValue `json:"secret_environment_variables"`
	// Region: Region in which the namespace will be created.
	Region scw.Region `json:"region"`
}

// Token:
type Token struct {
	// ID: UUID of the token.
	ID string `json:"id"`
	// Token: Identifier of the token.
	Token string `json:"token"`
	// ContainerID: UUID of the container the token belongs to.
	ContainerID *string `json:"container_id,omitempty"`
	// NamespaceID: UUID of the namespace the token belongs to.
	NamespaceID *string `json:"namespace_id,omitempty"`
	// Deprecated: PublicKey: Public key of the token.
	PublicKey *string `json:"public_key,omitempty"`
	// Status: Status of the token.
	Status TokenStatus `json:"status"`
	// Description: Description of the token.
	Description *string `json:"description"`
	// ExpiresAt: Expiry date of the token.
	ExpiresAt *time.Time `json:"expires_at"`
}

// Trigger:
type Trigger struct {
	// ID:
	ID string `json:"id"`
	// Name:
	Name string `json:"name"`
	// Description:
	Description string `json:"description"`
	// InputType:
	InputType TriggerInputType `json:"input_type"`
	// Status:
	Status TriggerStatus `json:"status"`
	// ErrorMessage:
	ErrorMessage *string `json:"error_message"`
	// ContainerID:
	ContainerID string `json:"container_id"`
	// ScwSqsConfig:
	ScwSqsConfig *TriggerMnqSqsClientConfig `json:"scw_sqs_config,omitempty"`
	// SqsConfig:
	SqsConfig *TriggerSqsClientConfig `json:"sqs_config,omitempty"`
	// ScwNatsConfig:
	ScwNatsConfig *TriggerMnqNatsClientConfig `json:"scw_nats_config,omitempty"`
}

// UpdateTriggerRequestSqsClientConfig:
type UpdateTriggerRequestSqsClientConfig struct {
	// AccessKey:
	AccessKey *string `json:"access_key"`
	// SecretKey:
	SecretKey *string `json:"secret_key"`
}

// CreateContainerRequest:
type CreateContainerRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NamespaceID: UUID of the namespace the container belongs to.
	NamespaceID string `json:"namespace_id"`
	// Name: Name of the container.
	Name string `json:"name"`
	// EnvironmentVariables: Environment variables of the container.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`
	// MinScale: Minimum number of instances to scale the container to.
	MinScale *uint32 `json:"min_scale,omitempty"`
	// MaxScale: Maximum number of instances to scale the container to.
	MaxScale *uint32 `json:"max_scale,omitempty"`
	// MemoryLimit: Memory limit of the container in MB.
	MemoryLimit *uint32 `json:"memory_limit,omitempty"`
	// CPULimit: CPU limit of the container in mvCPU.
	CPULimit *uint32 `json:"cpu_limit,omitempty"`
	// Timeout: Processing time limit for the container.
	Timeout *scw.Duration `json:"timeout,omitempty"`
	// Privacy: Privacy setting of the container.
	Privacy ContainerPrivacy `json:"privacy"`
	// Description: Description of the container.
	Description *string `json:"description,omitempty"`
	// RegistryImage: Name of the registry image (e.g. "rg.fr-par.scw.cloud/something/image:tag").
	RegistryImage *string `json:"registry_image,omitempty"`
	// MaxConcurrency: Number of maximum concurrent executions of the container.
	MaxConcurrency *uint32 `json:"max_concurrency,omitempty"`
	// Deprecated: DomainName: Domain name associated with the container.
	DomainName *string `json:"domain_name,omitempty"`
	// Protocol: Protocol the container uses.
	Protocol ContainerProtocol `json:"protocol"`
	// Port: Port the container listens on.
	Port *uint32 `json:"port,omitempty"`
	// SecretEnvironmentVariables: Secret environment variables of the container.
	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`
	// HTTPOption: Possible values:
	//  - redirected: Responds to HTTP request with a 301 redirect to ask the clients to use HTTPS.
	//  - enabled: Serve both HTTP and HTTPS traffic.
	HTTPOption ContainerHTTPOption `json:"http_option"`
}

// CreateCronRequest:
type CreateCronRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ContainerID: UUID of the container to invoke by the cron.
	ContainerID string `json:"container_id"`
	// Schedule: UNIX cron shedule.
	Schedule string `json:"schedule"`
	// Args: Arguments to pass with the cron.
	Args *scw.JSONObject `json:"args,omitempty"`
	// Name: Name of the cron to create.
	Name *string `json:"name,omitempty"`
}

// CreateDomainRequest:
type CreateDomainRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Hostname: Domain to assign.
	Hostname string `json:"hostname"`
	// ContainerID: UUID of the container to assign the domain to.
	ContainerID string `json:"container_id"`
}

// CreateNamespaceRequest:
type CreateNamespaceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Name: Name of the namespace to create.
	Name string `json:"name"`
	// EnvironmentVariables: Environment variables of the namespace to create.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`
	// ProjectID: UUID of the Project in which the namespace will be created.
	ProjectID string `json:"project_id"`
	// Description: Description of the namespace to create.
	Description *string `json:"description,omitempty"`
	// SecretEnvironmentVariables: Secret environment variables of the namespace to create.
	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`
}

// CreateTokenRequest:
type CreateTokenRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ContainerID: UUID of the container to create the token for.
	ContainerID *string `json:"container_id,omitempty"`
	// NamespaceID: UUID of the namespace to create the token for.
	NamespaceID *string `json:"namespace_id,omitempty"`
	// Description: Description of the token.
	Description *string `json:"description,omitempty"`
	// ExpiresAt: Expiry date of the token.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// CreateTriggerRequest:
type CreateTriggerRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Name:
	Name string `json:"name"`
	// Description:
	Description *string `json:"description,omitempty"`
	// ContainerID:
	ContainerID string `json:"container_id"`
	// ScwSqsConfig:
	ScwSqsConfig *CreateTriggerRequestMnqSqsClientConfig `json:"scw_sqs_config,omitempty"`
	// SqsConfig:
	SqsConfig *CreateTriggerRequestSqsClientConfig `json:"sqs_config,omitempty"`
	// ScwNatsConfig:
	ScwNatsConfig *CreateTriggerRequestMnqNatsClientConfig `json:"scw_nats_config,omitempty"`
}

// DeleteContainerRequest:
type DeleteContainerRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ContainerID: UUID of the container to delete.
	ContainerID string `json:"-"`
}

// DeleteCronRequest:
type DeleteCronRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// CronID: UUID of the cron to delete.
	CronID string `json:"-"`
}

// DeleteDomainRequest:
type DeleteDomainRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// DomainID: UUID of the domain to delete.
	DomainID string `json:"-"`
}

// DeleteNamespaceRequest:
type DeleteNamespaceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NamespaceID: UUID of the namespace to delete.
	NamespaceID string `json:"-"`
}

// DeleteTokenRequest:
type DeleteTokenRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// TokenID: UUID of the token to delete.
	TokenID string `json:"-"`
}

// DeleteTriggerRequest:
type DeleteTriggerRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// TriggerID:
	TriggerID string `json:"-"`
}

// DeployContainerRequest:
type DeployContainerRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ContainerID: UUID of the container to deploy.
	ContainerID string `json:"-"`
}

// GetContainerRequest:
type GetContainerRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ContainerID: UUID of the container to get.
	ContainerID string `json:"-"`
}

// GetCronRequest:
type GetCronRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// CronID: UUID of the cron to get.
	CronID string `json:"-"`
}

// GetDomainRequest:
type GetDomainRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// DomainID: UUID of the domain to get.
	DomainID string `json:"-"`
}

// GetNamespaceRequest:
type GetNamespaceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NamespaceID: UUID of the namespace to get.
	NamespaceID string `json:"-"`
}

// GetTokenRequest:
type GetTokenRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// TokenID: UUID of the token to get.
	TokenID string `json:"-"`
}

// GetTriggerRequest:
type GetTriggerRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// TriggerID:
	TriggerID string `json:"-"`
}

// IssueJWTRequest:
type IssueJWTRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ContainerID:
	ContainerID *string `json:"container_id,omitempty"`
	// NamespaceID:
	NamespaceID *string `json:"namespace_id,omitempty"`
	// ExpiresAt:
	ExpiresAt *time.Time `json:"-"`
}

// ListContainersRequest:
type ListContainersRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Page: Page number.
	Page *int32 `json:"-"`
	// PageSize: Number of containers per page.
	PageSize *uint32 `json:"-"`
	// OrderBy: Order of the containers.
	OrderBy ListContainersRequestOrderBy `json:"-"`
	// NamespaceID: UUID of the namespace the container belongs to.
	NamespaceID string `json:"-"`
	// Name: Name of the container.
	Name *string `json:"-"`
	// OrganizationID: UUID of the Organization the container belongs to.
	OrganizationID *string `json:"-"`
	// ProjectID: UUID of the Project the container belongs to.
	ProjectID *string `json:"-"`
}

// ListContainersResponse:
type ListContainersResponse struct {
	// Containers: Array of containers.
	Containers []*Container `json:"containers"`
	// TotalCount: Total number of containers.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListContainersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListContainersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListContainersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Containers = append(r.Containers, results.Containers...)
	r.TotalCount += uint32(len(results.Containers))
	return uint32(len(results.Containers)), nil
}

// ListCronsRequest:
type ListCronsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Page: Page number.
	Page *int32 `json:"-"`
	// PageSize: Number of crons per page.
	PageSize *uint32 `json:"-"`
	// OrderBy: Order of the crons.
	OrderBy ListCronsRequestOrderBy `json:"-"`
	// ContainerID: UUID of the container invoked by the cron.
	ContainerID string `json:"-"`
}

// ListCronsResponse:
type ListCronsResponse struct {
	// Crons: Array of crons.
	Crons []*Cron `json:"crons"`
	// TotalCount: Total number of crons.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCronsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCronsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListCronsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Crons = append(r.Crons, results.Crons...)
	r.TotalCount += uint32(len(results.Crons))
	return uint32(len(results.Crons)), nil
}

// ListDomainsRequest:
type ListDomainsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Page: Page number.
	Page *int32 `json:"-"`
	// PageSize: Number of domains per page.
	PageSize *uint32 `json:"-"`
	// OrderBy: Order of the domains.
	OrderBy ListDomainsRequestOrderBy `json:"-"`
	// ContainerID: UUID of the container the domain belongs to.
	ContainerID string `json:"-"`
}

// ListDomainsResponse:
type ListDomainsResponse struct {
	// Domains: Array of domains.
	Domains []*Domain `json:"domains"`
	// TotalCount: Total number of domains.
	TotalCount uint32 `json:"total_count"`
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

// ListLogsRequest:
type ListLogsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ContainerID: UUID of the container.
	ContainerID string `json:"-"`
	// Page: Page number.
	Page *int32 `json:"-"`
	// PageSize: Number of logs per page.
	PageSize *uint32 `json:"-"`
	// OrderBy: Order of the logs.
	OrderBy ListLogsRequestOrderBy `json:"-"`
}

// ListLogsResponse:
type ListLogsResponse struct {
	// Logs:
	Logs []*Log `json:"logs"`
	// TotalCount:
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListLogsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListLogsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListLogsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Logs = append(r.Logs, results.Logs...)
	r.TotalCount += uint32(len(results.Logs))
	return uint32(len(results.Logs)), nil
}

// ListNamespacesRequest:
type ListNamespacesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Page: Page number.
	Page *int32 `json:"-"`
	// PageSize: Number of namespaces per page.
	PageSize *uint32 `json:"-"`
	// OrderBy: Order of the namespaces.
	OrderBy ListNamespacesRequestOrderBy `json:"-"`
	// Name: Name of the namespaces.
	Name *string `json:"-"`
	// OrganizationID: UUID of the Organization the namespace belongs to.
	OrganizationID *string `json:"-"`
	// ProjectID: UUID of the Project the namespace belongs to.
	ProjectID *string `json:"-"`
}

// ListNamespacesResponse:
type ListNamespacesResponse struct {
	// Namespaces: Array of the namespaces.
	Namespaces []*Namespace `json:"namespaces"`
	// TotalCount: Total number of namespaces.
	TotalCount uint32 `json:"total_count"`
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

// ListTokensRequest:
type ListTokensRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Page: Page number.
	Page *int32 `json:"-"`
	// PageSize: Number of tokens per page.
	PageSize *uint32 `json:"-"`
	// OrderBy: Order of the tokens.
	OrderBy ListTokensRequestOrderBy `json:"-"`
	// ContainerID: UUID of the container the token belongs to.
	ContainerID *string `json:"-"`
	// NamespaceID: UUID of the namespace the token belongs to.
	NamespaceID *string `json:"-"`
}

// ListTokensResponse:
type ListTokensResponse struct {
	// Tokens:
	Tokens []*Token `json:"tokens"`
	// TotalCount:
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTokensResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTokensResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListTokensResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Tokens = append(r.Tokens, results.Tokens...)
	r.TotalCount += uint32(len(results.Tokens))
	return uint32(len(results.Tokens)), nil
}

// ListTriggersRequest:
type ListTriggersRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
	// OrderBy:
	OrderBy ListTriggersRequestOrderBy `json:"-"`
	// ContainerID:
	ContainerID *string `json:"container_id,omitempty"`
	// NamespaceID:
	NamespaceID *string `json:"namespace_id,omitempty"`
	// ProjectID:
	ProjectID *string `json:"project_id,omitempty"`
}

// ListTriggersResponse:
type ListTriggersResponse struct {
	// Triggers:
	Triggers []*Trigger `json:"triggers"`
	// TotalCount:
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTriggersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTriggersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListTriggersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Triggers = append(r.Triggers, results.Triggers...)
	r.TotalCount += uint32(len(results.Triggers))
	return uint32(len(results.Triggers)), nil
}

// UpdateContainerRequest:
type UpdateContainerRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ContainerID: UUID of the container to update.
	ContainerID string `json:"-"`
	// EnvironmentVariables: Environment variables of the container.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`
	// MinScale: Minimum number of instances to scale the container to.
	MinScale *uint32 `json:"min_scale,omitempty"`
	// MaxScale: Maximum number of instances to scale the container to.
	MaxScale *uint32 `json:"max_scale,omitempty"`
	// MemoryLimit: Memory limit of the container in MB.
	MemoryLimit *uint32 `json:"memory_limit,omitempty"`
	// CPULimit: CPU limit of the container in mvCPU.
	CPULimit *uint32 `json:"cpu_limit,omitempty"`
	// Timeout: Processing time limit for the container.
	Timeout *scw.Duration `json:"timeout,omitempty"`
	// Redeploy: Defines whether to redeploy failed containers.
	Redeploy *bool `json:"redeploy,omitempty"`
	// Privacy: Privacy settings of the container.
	Privacy ContainerPrivacy `json:"privacy"`
	// Description: Description of the container.
	Description *string `json:"description,omitempty"`
	// RegistryImage: Name of the registry image (e.g. "rg.fr-par.scw.cloud/something/image:tag").
	RegistryImage *string `json:"registry_image,omitempty"`
	// MaxConcurrency: Number of maximum concurrent executions of the container.
	MaxConcurrency *uint32 `json:"max_concurrency,omitempty"`
	// Deprecated: DomainName:
	DomainName *string `json:"domain_name,omitempty"`
	// Protocol:
	Protocol ContainerProtocol `json:"protocol"`
	// Port:
	Port *uint32 `json:"port,omitempty"`
	// SecretEnvironmentVariables:
	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`
	// HTTPOption: Possible values:
	//  - redirected: Responds to HTTP request with a 301 redirect to ask the clients to use HTTPS.
	//  - enabled: Serve both HTTP and HTTPS traffic.
	HTTPOption ContainerHTTPOption `json:"http_option"`
}

// UpdateCronRequest:
type UpdateCronRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// CronID: UUID of the cron to update.
	CronID string `json:"-"`
	// ContainerID: UUID of the container invoked by the cron.
	ContainerID *string `json:"container_id,omitempty"`
	// Schedule: UNIX cron schedule.
	Schedule *string `json:"schedule,omitempty"`
	// Args: Arguments to pass with the cron.
	Args *scw.JSONObject `json:"args,omitempty"`
	// Name: Name of the cron.
	Name *string `json:"name,omitempty"`
}

// UpdateNamespaceRequest:
type UpdateNamespaceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NamespaceID: UUID of the namespace to update.
	NamespaceID string `json:"-"`
	// EnvironmentVariables: Environment variables of the namespace to update.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`
	// Description: Description of the namespace to update.
	Description *string `json:"description,omitempty"`
	// SecretEnvironmentVariables: Secret environment variables of the namespace to update.
	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`
}

// UpdateTriggerRequest:
type UpdateTriggerRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// TriggerID:
	TriggerID string `json:"-"`
	// Name:
	Name *string `json:"name,omitempty"`
	// Description:
	Description *string `json:"description,omitempty"`
	// SqsConfig:
	SqsConfig *UpdateTriggerRequestSqsClientConfig `json:"sqs_config,omitempty"`
}

// Scaleway Serverless Containers is a «Container As A Service» product which gives users the ability to deploy atomic serverless workloads and only pay for resources used while containers are running.
//
// It provides many advantages, such as:
//
// - Containers are only executed when an event is triggered, which allows users to save money while code is not running
// - Auto-Scalability:
//   - Automated `Scaling up and down` based on user configuration (e.g. min: 0, max: 100 replicas of my container).
//   - Automated `Scaling to zero` when a container is not executed, which is cost-effective for the user and saves computing resources for the cloud provider.
//
// - Endpoint-only scaling
//
// ### Serverless Framework
//
// This page explains how to use the Scaleway Containers API, including a quickstart and the full API documentation. However, you may prefer to use the
// [Serverless Framework plugin](https://github.com/scaleway/serverless-scaleway-functions) enabling users to deploy their serverless workloads
// much more easily with a single `serverless deploy` command.
//
// If what you are looking for is an easy way to deploy your code, you may prefer Serverless Framework.
//
// Below, you will find a step-by-step guide on how to create a `namespace`, configure and deploy `containers`, and trigger your `containers` via HTTP and CRON.
//
// ## Concepts
//
// Refer to our [dedicated concepts page](https://www.scaleway.com/en/docs/serverless/containers/concepts/) to find definitions of the different terms referring to Serverless Containers.
// ## Quickstart
//
// (switchcolumn)
// (switchcolumn)
//
//  1. Configure your environment variables.
//     ```bash
//     export $SCW_SECRET_KEY="<Secret key of your token>"
//     export SCW_DEFAULT_REGION="<Choose your location (pl-waw/nl-ams/fr-par)>"
//     export SCW_PROJECT_ID="<Your Project ID>"
//     ```
//     <Message type="tip">
//     This is an optional step that seeks to simplify your usage of the Serverless Containers API. See [Regions](#availability-zones) below for help choosing a region. You can find your Project ID in the [Scaleway console](https://console.scaleway.com/Project/settings).
//     </Message>
//
//  2. Set the name for your namespace and configure your Project ID:
//     ```bash
//     curl -X POST "https://api.scaleway.com/containers/v1beta1/regions/$SCW_DEFAULT_REGION/namespaces" \
//     -H "accept: application/json" \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     -H "Content-Type: application/json" \
//     -d "{
//     \"name\": \"your-namespace-name\", \
//     \"project_id\": \"$SCW_PROJECT_ID\", \
//     \"environment_variables\": {\"YOUR_VARIABLE\": \"content\"} \
//     }"
//     ```
//
//  3. **Copy the response's `id` field**, you will need it for the next steps. For the sake of simplicity, we will save the ID to a variable, which will be used in the following examples:
//     ```bash
//     export NAMESPACE_ID="<your namespace id>"
//     ```
//     <Message type="note">
//     We suppose you already have a working image here. It can be anything which listens on a env variable \$PORT variable. Note that we run your container as user 1000, not root, so it must be runnable under these conditions.
//     For more information on how to push your image, refer to the [Container Registry documentation](https://www.scaleway.com/developers/api/registry/).
//     </Message>
//
//  4. **Edit the POST request payload** to use in the next step to create an Elastic Metal server. Modify the values in the example according to your requirements, using the information in the payload values section to help. For container resource criteria and default values, please refer to the [Containers Limitation documentation](https://www.scaleway.com/en/docs/serverless/containers/reference-content/containers-limitations/).
//     ```json
//     {
//     "namespace_id": "string",
//     "name": "string",
//     "environment_variables": {
//     "<key>": "string"
//     },
//     "min_scale": "integer",
//     "max_scale": "integer",
//     "memory_limit": "integer",
//     "cpu_limit": "integer",
//     "timeout": "integer",
//     "privacy": "unknown_privacy",
//     "description": "string",
//     "registry_image": "string",
//     "max_concurrency": "integer",
//     "protocol": "unknown_protocol",
//     "port": "integer",
//     "secret_environment_variables": [
//     {
//     "key": "string",
//     "value": "string"
//     }
//     ],
//     "http_option": "enabled"
//     }
//     ```
//
//     | Parameter        | Description                                                        |
//     | :--------------- | :----------------------------------------------------------------- |
//     | `region` | The region you want to target. Possible values are fr-par, nl-ams and pl-waw. |
//     | `namespace_id`| UUID of the namespace the container belongs to. |
//     | `name`| Name of the container. |
//     | `environment_variables` | **NULLABLE** Environment variables of the container. |
//     | `min_scale` | **NULLABLE** Minimum number of instances to scale the container to. |
//     | `max_scale` | **NULLABLE** Maximum number of instances to scale the container to. |
//     | `memory_limit`| **NULLABLE** Memory limit of the container in MiB. |
//     | `cpu_limit`| **NULLABLE** CPU limit of the container in mvCPU. |
//     | `timeout` | **NULLABLE** Request processing time limit for the container. (in seconds). |
//     | `privacy` | Privacy setting of the container. |
//     | `description` | **NULLABLE** Description of the container. |
//     | `registry_image`| **NULLABLE** Name of the registry image (e.g. "rg.fr-par.scw.cloud/something/image:tag"). |
//     | `max_concurrency` | **NULLABLE** Number of maximum concurrent executions of the container. |
//     | `protocol` | Protocol the container uses. Possible values are unknown_protocol, http1 and h2c. The default value is unknown_protocol. |
//     | `port` | **NULLABLE** Port the container listens on. |
//     | `secret_environment_variables` | Secret environment variables of the container. |
//     | `http_option` | Configure how HTTP and HTTPS requests are handled. Possible values:<br /> - `redirected`: Responds to HTTP request with a 301 redirect to ask the clients to use HTTPS.<br /> - `enabled`: Serve both HTTP and HTTPS traffic. |
//
//     <Message type="important">
//     All parameters are `required`, except for those marked as nullable.
//     </Message>
//
//  5. Run the following command to create your container. Make sure you include the payload you edited in the previous step.
//     ```bash
//     export REGISTRY_IMAGE="rg.fr-par.scw.cloud/myregistrynamespace/mycontainer:latest"
//     curl -X POST \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     "https://api.scaleway.com/containers/v1beta1/regions/$SCW_DEFAULT_REGION/containers" \
//     -d "{
//     "name": MyContainer, \
//     "registry_image": "$REGISTRY_IMAGE", \
//     "namespace_id": "$NAMESPACE_ID", \
//     "memory_limit": 300, \
//     "cpu_limit": 200, \
//     "min_scale": 0, \
//     "max_scale": 20 \
//     }"
//     ```
//
//  6. Run the following command to deploy your container:
//     ```bash
//     curl -X POST \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     "https://api.scaleway.com/containers/v1beta1/regions/$SCW_DEFAULT_REGION/containers/$CONTAINER_ID/deploy" \
//     -d "{}"
//     ```
//
//  7. Run the following command to trigger your container.
//     ```bash
//     curl -X GET \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     "https://api.scaleway.com/containers/v1beta1/regions/$SCW_DEFAULT_REGION/containers/$CONTAINER_ID"
//     export CONTAINER_ENDPOINT="<endpoint>"
//     curl -X GET "$CONTAINER_ENDPOINT"
//     ```
//
//  8. (optional) Use the following command to retrieve the containers output logs:
//     ```bash
//     curl -X GET \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     "https://api.scaleway.com/containers/v1beta1/regions/$SCW_DEFAULT_REGION/containers/$CONTAINER_ID/logs"
//     ```
//
//  9. (optional) Use the following call to destroy a namespace (along with all containers and crons):
//     ```bash
//     curl -s \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     -X DELETE "https://api.scaleway.com/containers/v1beta1/regions/$SCW_DEFAULT_REGION/namespaces/$NAMESPACE_ID"
//     ```
//
// (switchcolumn)
// <Message type="requirement">
//
//	To perform the following steps, you must first ensure that:
//	- You have a [Scaleway account](https://console.scaleway.com/)
//	- You have created an [API key](https://www.scaleway.com/en/docs/identity-and-access-management/iam/how-to/create-api-keys/) and that the API key has sufficient [IAM permissions](https://www.scaleway.com/en/docs/identity-and-access-management/iam/reference-content/permission-sets/) to perform the actions described on this page
//	- You have [installed `curl`](https://curl.se/download.html)
//	- You have [installed `jq`](https://stedolan.github.io/jq/) to improve readability of the API outputs
//
// </Message>
// (switchcolumn)
//
// ## Technical information
//
// A **Container** in Scaleway Containers consists of multiple components:
//
// - **Environment variables**: Users may configure specific environment variables (Database host/credentials for example) which are safely encrypted in our Database, and will be mounted inside your containers. **Note** that environment variables set at `namespace` level will also be mounted (in every container). Environment variables written at `container` level override the ones set at `namespace` level (if two of them bear the same name for example).
// - **Docker image**: A Docker image contains all the elements required to run your software: code, a runtime environment, tools, scripts, libraries, etc.
// - **Resources**: Users may decide how much computing resource to allocate to each container -> `Memory Limit` (in MB). We will then allocate the right amount of `CPU` based on Memory Limit choice. The right choice for your container's resources is very important, as you will be billed based on compute usage over time and the number of Containers executions.
//
// ### Product features
// - Fully isolated environments
// - Scaling to zero (saves money and computing resources while the code is not executed)
// - High Availability and scalability (automated and configurable, each container may scale automatically according to incoming workloads)
// - Multiple event sources:
//   - HTTP (request on our gateway will execute the container)
//   - CRON (time-based job, runs according to configurable cron schedule)
//
// - Integrated to the Scaleway Container Registry product:
//   - Deploy any docker image from one of your registry namespace
//
// - Flexible resources: you can choose values separately for your memory and CPU within a range in respect with maximum and minimum allowed values in [Containers Limitation documentation](https://www.scaleway.com/en/docs/serverless/containers/reference-content/containers-limitations/).
//
// ### Regions
//
// Serverless Containers is available in the Paris, Amsterdam and Warsaw regions, which are represented by the following path parameters:
//
// * `fr-par`
// * `nl-ams`
// * `pl-waw`
//
// ### CRON
//
// A `CRON` is a type of event which triggers a Scaleway Container: it is an `add-on` to your container.
//
// CRONs inside Scaleway Containers have the following properties:
//
// - `schedule`: UNIX Formatted CRON schedule. Your container will be executed based on this schedule. For example, `5 4 * * 0` means "execute my container at 04:05 AM every Sunday" (see this [page from Ubuntu's official documentation](https://doc.ubuntu-fr.org/cron)). The timezone is UTC+0.
// - `args`: JSON object passed to your container. You can use this property to define data that will be passed to your container's `event.body` object. For Containers, you might handle these arguments as the HTTP request's body.
//
// Under the hood, CRON Triggers are [Kubernetes JOBs](https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/) sending HTTP POST requests to your container.
//
// ### Authentication
//
// By default new containers are `public` meaning that no credentials are required to invoke them.
//
// A container can be `private` or `public`. This can be configured through the `privacy` parameter.
//
// Calling a `private` container without authentication will return HTTP code `403`.
//
// ## Going further
//
// For more help using Scaleway Serverless containers, check out the following resources:
//
// * Our [main documentation](https://www.scaleway.com/en/docs/serverless/containers/)
// * The #serverless-containers channel on our [Slack Community](https://www.scaleway.com/en/docs/tutorials/scaleway-slack-community/)
// * Our [support ticketing system](https://www.scaleway.com/en/docs/console/my-account/how-to/open-a-support-ticket).
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
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
}

// ListNamespaces: List all namespaces in a specified region.
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
		Query:  query,
	}

	var resp ListNamespacesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNamespace: Get the namespace associated with the specified ID.
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
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateNamespace: Create a new namespace in a specified region.
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
		req.Name = namegenerator.GetRandomName("cns")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
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

// UpdateNamespace: Update the space associated with the specified ID.
func (s *API) UpdateNamespace(req *UpdateNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
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
		Method: "PATCH",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
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

// DeleteNamespace: Delete the namespace associated with the specified ID.
func (s *API) DeleteNamespace(req *DeleteNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
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
		Method: "DELETE",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListContainers: List all containers for a specified region.
func (s *API) ListContainers(req *ListContainersRequest, opts ...scw.RequestOption) (*ListContainersResponse, error) {
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
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/containers",
		Query:  query,
	}

	var resp ListContainersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetContainer: Get the container associated with the specified ID.
func (s *API) GetContainer(req *GetContainerRequest, opts ...scw.RequestOption) (*Container, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ContainerID) == "" {
		return nil, errors.New("field ContainerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/containers/" + fmt.Sprint(req.ContainerID) + "",
	}

	var resp Container

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateContainer: Create a new container in the specified region.
func (s *API) CreateContainer(req *CreateContainerRequest, opts ...scw.RequestOption) (*Container, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/containers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Container

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateContainer: Update the container associated with the specified ID.
func (s *API) UpdateContainer(req *UpdateContainerRequest, opts ...scw.RequestOption) (*Container, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ContainerID) == "" {
		return nil, errors.New("field ContainerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/containers/" + fmt.Sprint(req.ContainerID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Container

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteContainer: Delete the container associated with the specified ID.
func (s *API) DeleteContainer(req *DeleteContainerRequest, opts ...scw.RequestOption) (*Container, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ContainerID) == "" {
		return nil, errors.New("field ContainerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/containers/" + fmt.Sprint(req.ContainerID) + "",
	}

	var resp Container

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeployContainer: Deploy a container associated with the specified ID.
func (s *API) DeployContainer(req *DeployContainerRequest, opts ...scw.RequestOption) (*Container, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ContainerID) == "" {
		return nil, errors.New("field ContainerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/containers/" + fmt.Sprint(req.ContainerID) + "/deploy",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Container

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListCrons: List all your crons.
func (s *API) ListCrons(req *ListCronsRequest, opts ...scw.RequestOption) (*ListCronsResponse, error) {
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
	parameter.AddToQuery(query, "container_id", req.ContainerID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons",
		Query:  query,
	}

	var resp ListCronsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCron: Get the cron associated with the specified ID.
func (s *API) GetCron(req *GetCronRequest, opts ...scw.RequestOption) (*Cron, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.CronID) == "" {
		return nil, errors.New("field CronID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons/" + fmt.Sprint(req.CronID) + "",
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateCron: Create a new cron.
func (s *API) CreateCron(req *CreateCronRequest, opts ...scw.RequestOption) (*Cron, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateCron: Update the cron associated with the specified ID.
func (s *API) UpdateCron(req *UpdateCronRequest, opts ...scw.RequestOption) (*Cron, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.CronID) == "" {
		return nil, errors.New("field CronID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons/" + fmt.Sprint(req.CronID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteCron: Delete the cron associated with the specified ID.
func (s *API) DeleteCron(req *DeleteCronRequest, opts ...scw.RequestOption) (*Cron, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.CronID) == "" {
		return nil, errors.New("field CronID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons/" + fmt.Sprint(req.CronID) + "",
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListLogs: List the logs of the container with the specified ID.
func (s *API) ListLogs(req *ListLogsRequest, opts ...scw.RequestOption) (*ListLogsResponse, error) {
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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ContainerID) == "" {
		return nil, errors.New("field ContainerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/containers/" + fmt.Sprint(req.ContainerID) + "/logs",
		Query:  query,
	}

	var resp ListLogsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDomains: List all domain name bindings in a specified region.
func (s *API) ListDomains(req *ListDomainsRequest, opts ...scw.RequestOption) (*ListDomainsResponse, error) {
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
	parameter.AddToQuery(query, "container_id", req.ContainerID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains",
		Query:  query,
	}

	var resp ListDomainsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDomain: Get a domain name binding for the container with the specified ID.
func (s *API) GetDomain(req *GetDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DomainID) == "" {
		return nil, errors.New("field DomainID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "",
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDomain: Create a domain name binding for the container with the specified ID.
func (s *API) CreateDomain(req *CreateDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains",
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

// DeleteDomain: Delete the domain name binding with the specific ID.
func (s *API) DeleteDomain(req *DeleteDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DomainID) == "" {
		return nil, errors.New("field DomainID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "",
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: IssueJWT:
func (s *API) IssueJWT(req *IssueJWTRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "expires_at", req.ExpiresAt)
	parameter.AddToQuery(query, "container_id", req.ContainerID)
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/issue-jwt",
		Query:  query,
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateToken: Create a new revocable token.
func (s *API) CreateToken(req *CreateTokenRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/tokens",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetToken: Get a token with a specified ID.
func (s *API) GetToken(req *GetTokenRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TokenID) == "" {
		return nil, errors.New("field TokenID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/tokens/" + fmt.Sprint(req.TokenID) + "",
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListTokens: List all tokens belonging to a specified Organization or Project.
func (s *API) ListTokens(req *ListTokensRequest, opts ...scw.RequestOption) (*ListTokensResponse, error) {
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
	parameter.AddToQuery(query, "container_id", req.ContainerID)
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/tokens",
		Query:  query,
	}

	var resp ListTokensResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteToken: Delete a token with a specified ID.
func (s *API) DeleteToken(req *DeleteTokenRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TokenID) == "" {
		return nil, errors.New("field TokenID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/tokens/" + fmt.Sprint(req.TokenID) + "",
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateTrigger: Create a new trigger for a specified container.
func (s *API) CreateTrigger(req *CreateTriggerRequest, opts ...scw.RequestOption) (*Trigger, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Trigger

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTrigger: Get a trigger with a specified ID.
func (s *API) GetTrigger(req *GetTriggerRequest, opts ...scw.RequestOption) (*Trigger, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TriggerID) == "" {
		return nil, errors.New("field TriggerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers/" + fmt.Sprint(req.TriggerID) + "",
	}

	var resp Trigger

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListTriggers: List all triggers belonging to a specified Organization or Project.
func (s *API) ListTriggers(req *ListTriggersRequest, opts ...scw.RequestOption) (*ListTriggersResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}
	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.ContainerID == nil && req.NamespaceID == nil && req.ProjectID == nil {
		req.ProjectID = &defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "container_id", req.ContainerID)
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers",
		Query:  query,
	}

	var resp ListTriggersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateTrigger: Update a trigger with a specified ID.
func (s *API) UpdateTrigger(req *UpdateTriggerRequest, opts ...scw.RequestOption) (*Trigger, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TriggerID) == "" {
		return nil, errors.New("field TriggerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers/" + fmt.Sprint(req.TriggerID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Trigger

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteTrigger: Delete a trigger with a specified ID.
func (s *API) DeleteTrigger(req *DeleteTriggerRequest, opts ...scw.RequestOption) (*Trigger, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TriggerID) == "" {
		return nil, errors.New("field TriggerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers/" + fmt.Sprint(req.TriggerID) + "",
	}

	var resp Trigger

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
