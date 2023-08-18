// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package function provides methods and message types of the function v1beta1 API.
package function

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

type FunctionHTTPOption string

const (
	FunctionHTTPOptionUnknownHTTPOption = FunctionHTTPOption("unknown_http_option")
	FunctionHTTPOptionEnabled           = FunctionHTTPOption("enabled")
	FunctionHTTPOptionRedirected        = FunctionHTTPOption("redirected")
)

func (enum FunctionHTTPOption) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_http_option"
	}
	return string(enum)
}

func (enum FunctionHTTPOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FunctionHTTPOption) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FunctionHTTPOption(FunctionHTTPOption(tmp).String())
	return nil
}

type FunctionPrivacy string

const (
	FunctionPrivacyUnknownPrivacy = FunctionPrivacy("unknown_privacy")
	FunctionPrivacyPublic         = FunctionPrivacy("public")
	FunctionPrivacyPrivate        = FunctionPrivacy("private")
)

func (enum FunctionPrivacy) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_privacy"
	}
	return string(enum)
}

func (enum FunctionPrivacy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FunctionPrivacy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FunctionPrivacy(FunctionPrivacy(tmp).String())
	return nil
}

type FunctionRuntime string

const (
	FunctionRuntimeUnknownRuntime = FunctionRuntime("unknown_runtime")
	FunctionRuntimeGolang         = FunctionRuntime("golang")
	FunctionRuntimePython         = FunctionRuntime("python")
	FunctionRuntimePython3        = FunctionRuntime("python3")
	FunctionRuntimeNode8          = FunctionRuntime("node8")
	FunctionRuntimeNode10         = FunctionRuntime("node10")
	FunctionRuntimeNode14         = FunctionRuntime("node14")
	FunctionRuntimeNode16         = FunctionRuntime("node16")
	FunctionRuntimeNode17         = FunctionRuntime("node17")
	FunctionRuntimePython37       = FunctionRuntime("python37")
	FunctionRuntimePython38       = FunctionRuntime("python38")
	FunctionRuntimePython39       = FunctionRuntime("python39")
	FunctionRuntimePython310      = FunctionRuntime("python310")
	FunctionRuntimeGo113          = FunctionRuntime("go113")
	FunctionRuntimeGo117          = FunctionRuntime("go117")
	FunctionRuntimeGo118          = FunctionRuntime("go118")
	FunctionRuntimeNode18         = FunctionRuntime("node18")
	FunctionRuntimeRust165        = FunctionRuntime("rust165")
	FunctionRuntimeGo119          = FunctionRuntime("go119")
	FunctionRuntimePython311      = FunctionRuntime("python311")
	FunctionRuntimePhp82          = FunctionRuntime("php82")
	FunctionRuntimeNode19         = FunctionRuntime("node19")
	FunctionRuntimeGo120          = FunctionRuntime("go120")
	FunctionRuntimeNode20         = FunctionRuntime("node20")
)

func (enum FunctionRuntime) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_runtime"
	}
	return string(enum)
}

func (enum FunctionRuntime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FunctionRuntime) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FunctionRuntime(FunctionRuntime(tmp).String())
	return nil
}

type FunctionStatus string

const (
	FunctionStatusUnknown  = FunctionStatus("unknown")
	FunctionStatusReady    = FunctionStatus("ready")
	FunctionStatusDeleting = FunctionStatus("deleting")
	FunctionStatusError    = FunctionStatus("error")
	FunctionStatusLocked   = FunctionStatus("locked")
	FunctionStatusCreating = FunctionStatus("creating")
	FunctionStatusPending  = FunctionStatus("pending")
	FunctionStatusCreated  = FunctionStatus("created")
)

func (enum FunctionStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum FunctionStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FunctionStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FunctionStatus(FunctionStatus(tmp).String())
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

type ListFunctionsRequestOrderBy string

const (
	ListFunctionsRequestOrderByCreatedAtAsc  = ListFunctionsRequestOrderBy("created_at_asc")
	ListFunctionsRequestOrderByCreatedAtDesc = ListFunctionsRequestOrderBy("created_at_desc")
	ListFunctionsRequestOrderByNameAsc       = ListFunctionsRequestOrderBy("name_asc")
	ListFunctionsRequestOrderByNameDesc      = ListFunctionsRequestOrderBy("name_desc")
)

func (enum ListFunctionsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListFunctionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListFunctionsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListFunctionsRequestOrderBy(ListFunctionsRequestOrderBy(tmp).String())
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

type RuntimeStatus string

const (
	RuntimeStatusUnknownStatus = RuntimeStatus("unknown_status")
	RuntimeStatusBeta          = RuntimeStatus("beta")
	RuntimeStatusAvailable     = RuntimeStatus("available")
	RuntimeStatusDeprecated    = RuntimeStatus("deprecated")
	RuntimeStatusEndOfSupport  = RuntimeStatus("end_of_support")
	RuntimeStatusEndOfLife     = RuntimeStatus("end_of_life")
)

func (enum RuntimeStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum RuntimeStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RuntimeStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RuntimeStatus(RuntimeStatus(tmp).String())
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

// Cron:
type Cron struct {
	// ID: UUID of the cron.
	ID string `json:"id"`
	// FunctionID: UUID of the function the cron applies to.
	FunctionID string `json:"function_id"`
	// Schedule: Schedule of the cron.
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
	// Hostname: Hostname associated with the function.
	Hostname string `json:"hostname"`
	// FunctionID: UUID of the function the domain is associated with.
	FunctionID string `json:"function_id"`
	// URL: URL of the function.
	URL string `json:"url"`
	// Status: State of the doamin.
	Status DomainStatus `json:"status"`
	// ErrorMessage: Error message if the domain is in "error" state.
	ErrorMessage *string `json:"error_message"`
}

// Runtime:
type Runtime struct {
	// Name:
	Name string `json:"name"`
	// Language:
	Language string `json:"language"`
	// Version:
	Version string `json:"version"`
	// DefaultHandler:
	DefaultHandler string `json:"default_handler"`
	// CodeSample:
	CodeSample string `json:"code_sample"`
	// Status:
	Status RuntimeStatus `json:"status"`
	// StatusMessage:
	StatusMessage string `json:"status_message"`
	// Extension:
	Extension string `json:"extension"`
	// Implementation:
	Implementation string `json:"implementation"`
	// LogoURL:
	LogoURL string `json:"logo_url"`
}

// Function:
type Function struct {
	// ID: UUID of the function.
	ID string `json:"id"`
	// Name: Name of the function.
	Name string `json:"name"`
	// NamespaceID: UUID of the namespace the function belongs to.
	NamespaceID string `json:"namespace_id"`
	// Status: Status of the function.
	Status FunctionStatus `json:"status"`
	// EnvironmentVariables: Environment variables of the function.
	EnvironmentVariables map[string]string `json:"environment_variables"`
	// MinScale: Minimum number of instances to scale the function to.
	MinScale uint32 `json:"min_scale"`
	// MaxScale: Maximum number of instances to scale the function to.
	MaxScale uint32 `json:"max_scale"`
	// Runtime: Runtime of the function.
	Runtime FunctionRuntime `json:"runtime"`
	// MemoryLimit: Memory limit of the function in MB.
	MemoryLimit uint32 `json:"memory_limit"`
	// CPULimit: CPU limit of the function.
	CPULimit uint32 `json:"cpu_limit"`
	// Timeout: Request processing time limit for the function.
	Timeout *scw.Duration `json:"timeout"`
	// Handler: Handler to use for the function.
	Handler string `json:"handler"`
	// ErrorMessage: Error message if the function is in "error" state.
	ErrorMessage *string `json:"error_message"`
	// BuildMessage: Description of the current build step.
	BuildMessage *string `json:"build_message"`
	// Privacy: Privacy setting of the function.
	Privacy FunctionPrivacy `json:"privacy"`
	// Description: Description of the function.
	Description *string `json:"description"`
	// DomainName: Domain name associated with the function.
	DomainName string `json:"domain_name"`
	// SecretEnvironmentVariables: Secret environment variables of the function.
	SecretEnvironmentVariables []*SecretHashedValue `json:"secret_environment_variables"`
	// Region: Region in which the function is deployed.
	Region scw.Region `json:"region"`
	// HTTPOption: Possible values:
	//  - redirected: Responds to HTTP request with a 301 redirect to ask the clients to use HTTPS.
	//  - enabled: Serve both HTTP and HTTPS traffic.
	HTTPOption FunctionHTTPOption `json:"http_option"`
	// RuntimeMessage:
	RuntimeMessage string `json:"runtime_message"`
}

// Log:
type Log struct {
	// Message: Message of the log.
	Message string `json:"message"`
	// Timestamp: Timestamp of the log.
	Timestamp *time.Time `json:"timestamp"`
	// ID: UUID of the log.
	ID string `json:"id"`
	// Level: Severity of the log (info, debug, error etc.).
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
	// ErrorMessage: Error message if the namespace is in "error" state.
	ErrorMessage *string `json:"error_message"`
	// RegistryEndpoint: Registry endpoint of the namespace.
	RegistryEndpoint string `json:"registry_endpoint"`
	// Description: Description of the namespace.
	Description *string `json:"description"`
	// SecretEnvironmentVariables: Secret environment variables of the namespace.
	SecretEnvironmentVariables []*SecretHashedValue `json:"secret_environment_variables"`
	// Region: Region in which the namespace is located.
	Region scw.Region `json:"region"`
}

// Token:
type Token struct {
	// ID: UUID of the token.
	ID string `json:"id"`
	// Token: String of the token.
	Token string `json:"token"`
	// FunctionID: UUID of the function the token is associated with.
	FunctionID *string `json:"function_id,omitempty"`
	// NamespaceID: UUID of the namespace the token is assoicated with.
	NamespaceID *string `json:"namespace_id,omitempty"`
	// Deprecated: PublicKey: Public key of the token.
	PublicKey *string `json:"public_key,omitempty"`
	// Status: Status of the token.
	Status TokenStatus `json:"status"`
	// Description: Description of the token.
	Description *string `json:"description"`
	// ExpiresAt: Date on which the token expires.
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
	// FunctionID:
	FunctionID string `json:"function_id"`
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

// CreateCronRequest:
type CreateCronRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// FunctionID: UUID of the function to use the cron with.
	FunctionID string `json:"function_id"`
	// Schedule: Schedule of the cron in UNIX cron format.
	Schedule string `json:"schedule"`
	// Args: Arguments to use with the cron.
	Args *scw.JSONObject `json:"args,omitempty"`
	// Name: Name of the cron.
	Name *string `json:"name,omitempty"`
}

// CreateDomainRequest:
type CreateDomainRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Hostname: Hostame to create.
	Hostname string `json:"hostname"`
	// FunctionID: UUID of the function to associate the domain with.
	FunctionID string `json:"function_id"`
}

// CreateFunctionRequest:
type CreateFunctionRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Name: Name of the function to create.
	Name string `json:"name"`
	// NamespaceID: UUID of the namespace the function will be created in.
	NamespaceID string `json:"namespace_id"`
	// EnvironmentVariables: Environment variables of the function.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`
	// MinScale: Minumum number of instances to scale the function to.
	MinScale *uint32 `json:"min_scale,omitempty"`
	// MaxScale: Maximum number of instances to scale the function to.
	MaxScale *uint32 `json:"max_scale,omitempty"`
	// Runtime: Runtime to use with the function.
	Runtime FunctionRuntime `json:"runtime"`
	// MemoryLimit: Memory limit of the function in MB.
	MemoryLimit *uint32 `json:"memory_limit,omitempty"`
	// Timeout: Request processing time limit for the function.
	Timeout *scw.Duration `json:"timeout,omitempty"`
	// Handler: Handler to use with the function.
	Handler *string `json:"handler,omitempty"`
	// Privacy: Privacy setting of the function.
	Privacy FunctionPrivacy `json:"privacy"`
	// Description: Description of the function.
	Description *string `json:"description,omitempty"`
	// Deprecated: DomainName: Domain name associated with the function.
	DomainName *string `json:"domain_name,omitempty"`
	// SecretEnvironmentVariables:
	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`
	// HTTPOption: Possible values:
	//  - redirected: Responds to HTTP request with a 301 redirect to ask the clients to use HTTPS.
	//  - enabled: Serve both HTTP and HTTPS traffic.
	HTTPOption FunctionHTTPOption `json:"http_option"`
}

// CreateNamespaceRequest:
type CreateNamespaceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Name:
	Name string `json:"name"`
	// EnvironmentVariables: Environment variables of the namespace.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`
	// ProjectID: UUID of the project in which the namespace will be created.
	ProjectID string `json:"project_id"`
	// Description: Description of the namespace.
	Description *string `json:"description,omitempty"`
	// SecretEnvironmentVariables: Secret environment variables of the namespace.
	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`
}

// CreateTokenRequest:
type CreateTokenRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// FunctionID: UUID of the function to associate the token with.
	FunctionID *string `json:"function_id,omitempty"`
	// NamespaceID: UUID of the namespace to associate the token with.
	NamespaceID *string `json:"namespace_id,omitempty"`
	// Description: Description of the token.
	Description *string `json:"description,omitempty"`
	// ExpiresAt: Date on which the token expires.
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
	// FunctionID:
	FunctionID string `json:"function_id"`
	// ScwSqsConfig:
	ScwSqsConfig *CreateTriggerRequestMnqSqsClientConfig `json:"scw_sqs_config,omitempty"`
	// SqsConfig:
	SqsConfig *CreateTriggerRequestSqsClientConfig `json:"sqs_config,omitempty"`
	// ScwNatsConfig:
	ScwNatsConfig *CreateTriggerRequestMnqNatsClientConfig `json:"scw_nats_config,omitempty"`
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

// DeleteFunctionRequest:
type DeleteFunctionRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// FunctionID: UUID of the function to delete.
	FunctionID string `json:"-"`
}

// DeleteNamespaceRequest:
type DeleteNamespaceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NamespaceID: UUID of the namespace.
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

// DeployFunctionRequest:
type DeployFunctionRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// FunctionID: UUID of the function to deploy.
	FunctionID string `json:"-"`
}

// DownloadURL:
type DownloadURL struct {
	// URL:
	URL string `json:"url"`
	// Headers:
	Headers map[string]*[]string `json:"headers"`
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

// GetFunctionDownloadURLRequest:
type GetFunctionDownloadURLRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// FunctionID: UUID of the function to get the the download URL for.
	FunctionID string `json:"-"`
}

// GetFunctionRequest:
type GetFunctionRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// FunctionID: UUID of the function.
	FunctionID string `json:"-"`
}

// GetFunctionUploadURLRequest:
type GetFunctionUploadURLRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// FunctionID: UUID of the function to get the upload URL for.
	FunctionID string `json:"-"`
	// ContentLength: Size of the archive to upload in bytes.
	ContentLength uint64 `json:"-"`
}

// GetNamespaceRequest:
type GetNamespaceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NamespaceID: UUID of the namespace.
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
	// FunctionID:
	FunctionID *string `json:"function_id,omitempty"`
	// NamespaceID:
	NamespaceID *string `json:"namespace_id,omitempty"`
	// ExpiresAt:
	ExpiresAt *time.Time `json:"-"`
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
	// FunctionID: UUID of the function.
	FunctionID string `json:"-"`
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
	// FunctionID: UUID of the function the domain is assoicated with.
	FunctionID string `json:"-"`
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

// ListFunctionRuntimesRequest:
type ListFunctionRuntimesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
}

// ListFunctionRuntimesResponse:
type ListFunctionRuntimesResponse struct {
	// Runtimes: Array of runtimes available.
	Runtimes []*Runtime `json:"runtimes"`
	// TotalCount: Total number of runtimes.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListFunctionRuntimesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListFunctionRuntimesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListFunctionRuntimesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Runtimes = append(r.Runtimes, results.Runtimes...)
	r.TotalCount += uint32(len(results.Runtimes))
	return uint32(len(results.Runtimes)), nil
}

// ListFunctionsRequest:
type ListFunctionsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Page: Page number.
	Page *int32 `json:"-"`
	// PageSize: Number of functions per page.
	PageSize *uint32 `json:"-"`
	// OrderBy: Order of the functions.
	OrderBy ListFunctionsRequestOrderBy `json:"-"`
	// NamespaceID: UUID of the namespace the function belongs to.
	NamespaceID string `json:"-"`
	// Name: Name of the function.
	Name *string `json:"-"`
	// OrganizationID: UUID of the Organziation the function belongs to.
	OrganizationID *string `json:"-"`
	// ProjectID: UUID of the Project the function belongs to.
	ProjectID *string `json:"-"`
}

// ListFunctionsResponse:
type ListFunctionsResponse struct {
	// Functions: Array of functions.
	Functions []*Function `json:"functions"`
	// TotalCount: Total number of functions.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListFunctionsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListFunctionsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListFunctionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Functions = append(r.Functions, results.Functions...)
	r.TotalCount += uint32(len(results.Functions))
	return uint32(len(results.Functions)), nil
}

// ListLogsRequest:
type ListLogsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// FunctionID: UUID of the function to get the logs for.
	FunctionID string `json:"-"`
	// Page: Page number.
	Page *int32 `json:"-"`
	// PageSize: Number of logs per page.
	PageSize *uint32 `json:"-"`
	// OrderBy: Order of the logs.
	OrderBy ListLogsRequestOrderBy `json:"-"`
}

// ListLogsResponse:
type ListLogsResponse struct {
	// Logs: Array of logs.
	Logs []*Log `json:"logs"`
	// TotalCount: Total number of logs.
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
	// Name: Name of the namespace.
	Name *string `json:"-"`
	// OrganizationID: UUID of the Organization the namespace belongs to.
	OrganizationID *string `json:"-"`
	// ProjectID: UUID of the Project the namespace belongs to.
	ProjectID *string `json:"-"`
}

// ListNamespacesResponse:
type ListNamespacesResponse struct {
	// Namespaces:
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
	// OrderBy: Sort order for the tokens.
	OrderBy ListTokensRequestOrderBy `json:"-"`
	// FunctionID: UUID of the function the token is assoicated with.
	FunctionID *string `json:"-"`
	// NamespaceID: UUID of the namespace the token is associated with.
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
	// FunctionID:
	FunctionID *string `json:"function_id,omitempty"`
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

// UpdateCronRequest:
type UpdateCronRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// CronID: UUID of the cron to update.
	CronID string `json:"-"`
	// FunctionID: UUID of the function to use the cron with.
	FunctionID *string `json:"function_id,omitempty"`
	// Schedule: Schedule of the cron in UNIX cron format.
	Schedule *string `json:"schedule,omitempty"`
	// Args: Arguments to use with the cron.
	Args *scw.JSONObject `json:"args,omitempty"`
	// Name: Name of the cron.
	Name *string `json:"name,omitempty"`
}

// UpdateFunctionRequest:
type UpdateFunctionRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// FunctionID: UUID of the function to update.
	FunctionID string `json:"-"`
	// EnvironmentVariables: Environment variables of the function to update.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`
	// MinScale: Minumum number of instances to scale the function to.
	MinScale *uint32 `json:"min_scale,omitempty"`
	// MaxScale: Maximum number of instances to scale the function to.
	MaxScale *uint32 `json:"max_scale,omitempty"`
	// Runtime: Runtime to use with the function.
	Runtime FunctionRuntime `json:"runtime"`
	// MemoryLimit: Memory limit of the function in MB.
	MemoryLimit *uint32 `json:"memory_limit,omitempty"`
	// Timeout: Processing time limit for the function.
	Timeout *scw.Duration `json:"timeout,omitempty"`
	// Redeploy: Redeploy failed function.
	Redeploy *bool `json:"redeploy,omitempty"`
	// Handler: Handler to use with the function.
	Handler *string `json:"handler,omitempty"`
	// Privacy: Privacy setting of the function.
	Privacy FunctionPrivacy `json:"privacy"`
	// Description: Description of the function.
	Description *string `json:"description,omitempty"`
	// Deprecated: DomainName: Domain name associated with the function.
	DomainName *string `json:"domain_name,omitempty"`
	// SecretEnvironmentVariables: Secret environment variables of the function.
	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`
	// HTTPOption: Possible values:
	//  - redirected: Responds to HTTP request with a 301 redirect to ask the clients to use HTTPS.
	//  - enabled: Serve both HTTP and HTTPS traffic.
	HTTPOption FunctionHTTPOption `json:"http_option"`
}

// UpdateNamespaceRequest:
type UpdateNamespaceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NamespaceID: UUID of the namespapce.
	NamespaceID string `json:"-"`
	// EnvironmentVariables: Environment variables of the namespace.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`
	// Description: Description of the namespace.
	Description *string `json:"description,omitempty"`
	// SecretEnvironmentVariables: Secret environment variables of the namespace.
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

// UploadURL:
type UploadURL struct {
	// URL: Upload URL to upload the function to.
	URL string `json:"url"`
	// Headers: HTTP headers.
	Headers map[string]*[]string `json:"headers"`
}

// Deploying applications to the cloud becomes simpler with Serverless Functions. Instead of provisioning and maintaining servers, you just need to install a "function" of your business logic on the Scaleway cloud platform. The platform executes this function on demand, managing resource allocation for you. If the system needs to accommodate a hundred simultaneous requests, it creates a hundred (or more) copies of your service. Conversely, if demand drops to two concurrent requests, it destroys the unneeded ones. You only pay for the resources your functions use when they need them.
//
// The benefits of Scaleway Serverless Functions include:
//
// * Saving money while code is not running, as functions are only executed when an event is triggered
// * Auto-scalability:
//   - Scaling up and down automatically based on user configuration (e.g., min: 0, max: 100 replicas of my function).
//   - Scaling to zero when the function is not executed, which saves money for the user and computing resources for the cloud provider.
//
// * Endpoint-only scaling.
//
// ### Deploying functions
//
// This page explains how to use the Scaleway Serverless Functions API, including a quickstart and the full API documentation. However we provide several
// ways to deploy functions: [Function Deployment Methods](https://www.scaleway.com/en/docs/serverless/functions/reference-content/deploy-function/).
//
// ## Concepts
//
// Refer to our [dedicated concepts](https://www.scaleway.com/en/docs/serverless/functions/concepts/) page to find definitions of all terminology related to Scaleway Serverless Functions.
//
// ## Quickstart
//
// <Message type="note">
//
//	For functions concepts and advanced documentation, please refer to [Scaleway Quickstart for Functions](https://www.scaleway.com/en/docs/serverless/functions/quickstart/).
//
// </Message>
//
// (switchcolumn)
// (switchcolumn)
//
//  1. Configure your environment variables.
//     ```bash
//     export SCW_SECRET_KEY="<your secret key>"
//     export SCW_DEFAULT_REGION="<choose your location (fr-par/nl-ams/pl-waw)>"
//     export SCW_DEFAULT_PROJECT_ID="<your project ID>"
//     ```
//     <Message type="note">
//     A namespace is a project that allows you to group your functions. Functions in the same namespace can share environment variables and access tokens, defined at the namespace level.
//     </Message>
//
//  2. Set the name for your namespce and set your Project ID.
//     ```bash
//     curl -X POST \
//     "https://api.scaleway.com/functions/v1beta1/regions/$SCW_DEFAULT_REGION/namespaces" \
//     -H "accept: application/json" \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     -H "Content-Type: application/json" \
//     -d '{
//     name: your-namespace-name, \
//     project_id: "$SCW_DEFAULT_PROJECT_ID", \
//     environment_variables: \
//     {
//     YOUR_VARIABLE: content
//     } \
//     }'
//     ```
//  3. Copy the `id` field of the response to use in the next steps. For the sake of simplicity, we will save the ID to a variable, which we will use in the following examples.
//     ```bash
//     export SCW_NAMESPACE_ID="<your namespace id>"
//     ```
//
//  4. Edit the POST request payload that we will use in the next step to create a function. Modify the values in the example according to your needs, using the information in the table to help.
//     ```
//     {
//     "name": "string",
//     "namespace_id": "string",
//     "environment_variables": {
//     "<key>": "string"
//     },
//     "min_scale": "integer",
//     "max_scale": "integer",
//     "runtime": "unknown_runtime",
//     "memory_limit": "integer",
//     "timeout": "2.5s",
//     "handler": "string",
//     "privacy": "unknown_privacy",
//     "description": "string",
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
//     | `name`           | The name of your function |
//     | `namespace_id`   | ID of the namespace in which you want to create your function |
//     | `runtime`        | Your function's runtime, check the supported runtimes above |
//     | `environment_variables` | **NULLABLE** Environment variables of the function. |
//     | `memory_limit`   |**NULLABLE** Memory (in MB) allocated to your function, see the table of memory/CPU allocation above (increasing the memory limit will increase the cost of your function executions as we allocate more resources to your functions). |
//     | `min_scale`      | **NULLABLE** Minimum replicas for your function, defaults to `0`, **Note** that a function is `billed` when it gets executed, and using a `min_scale` greater than 0 will cause your function to run all the time. |
//     | `max_scale`      | **NULLABLE** Maximum replicas for your function (defaults to `20`), our system will scale your functions automatically based on incoming workload, but will never scale the number of replicas above the configured `max_scale`. |
//     | `timeout`        | **NULLABLE** Holds the max duration (in seconds) the function is allowed for responding to a request |
//     | `description`    | **NULLABLE** Description of the function. |
//     | `handler`        | **NULLABLE** More details with examples in each language/runtime section below |
//
// Note that the meaning of the value set in `handler` will change depending on the `runtime` parameter:
// *  `Node`, `PHP` and `Python` runtimes: `handler` is the path to the function file, followed by the name of the function to call. For example, `src/handler.handle` specifies a function file at `src/handler.js`, `src/handler.php` and `src/handler.py` respectively, with a function called `handle`.
// * `Go` and `Rust` runtimes: `handler` is the path to the package containing the handler. For example, `my_handler` specifies that the function file is located in a `my_handler` directory. The functions must then be part of a `main` package, exposing a `main` function.
//
//	<Message type="note">
//	  All parameters, except the ones marked with `nullable` are **required**.
//	</Message>
//
//  5. Run the following command to create your function. Make sure you include the payload you edited in the previous step.
//     ```bash
//     curl -X POST \
//     -H "X-Auth-Token: $SCW_SECRET_KEY"\
//     "https://api.scaleway.com/functions/v1beta1/regions/$SCW_DEFAULT_REGION/functions"\
//     -d '{
//     name: function-name, \
//     namespace_id: "$SCW_NAMESPACE_ID", \
//     memory_limit: 128, \
//     min_scale: 0, \
//     max_scale: 20, \
//     runtime: node16", \
//     handler: handler.myHandler \
//     }'
//     export SCW_FUNCTION_ID = "<your-function-id>"
//     ```
//
// 6. Upload your code from the console or directly via the API. Refer to our documentation [How to package your function dependencies and upload it as a zip-file](https://www.scaleway.com/en/docs/serverless/functions/how-to/package-function-dependencies-in-zip/) for more information.
//
//  7. Run the following command to deploy your function.
//     ```bash
//     curl -X POST \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" "https://api.scaleway.com/functions/v1beta1/regions/$SCW_DEFAULT_REGION/functions/$SCW_FUNCTION_ID/deploy" \
//     -d "{}"
//     ```
//     <Message type="note">
//     The process may take some minutes, as we have to build your source code into an executable function (wrapped by our runtimes), and deploy it to the Scaleway cloud platform.
//     </Message>
//
//  8. Retrieve your function's HTTP(s) endpoint with the following command once your function has been properly deployed.
//     ```bash
//     curl -X GET \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" "https://api.scaleway.com/functions/v1beta1/regions/$SCW_DEFAULT_REGION/functions/$SCW_FUNCTION_ID"
//     export FUNCTION_ENDPOINT="<endpoint>"
//     ```
//
//  9. Call your function via its endpoint.
//     ```bash
//     curl -X GET "$FUNCTION_ENDPOINT"
//     ```
//
//  10. (optional) Destroy the namespace (along with all functions and crons) by using the following call.
//     ```bash
//     curl -s \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" -X DELETE "https://api.scaleway.com/functions/v1beta1/regions/$SCW_DEFAULT_REGION/namespaces/$SCW_NAMESPACE_ID"
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
// - Functions are fully isolated environments
// - Scaling to zero (save money and computing resources while code is not executed)
// - High Availability and Scalability (Automated and configurable, each function may scale automatically according to incoming workloads)
// - Different programming languages supported
// - Multiple event sources:
//   - HTTP (request on our Gateway will execute the function)
//   - CRON (time-based job, runs according to configurable cron schedule)
//
// - Integrated with the Scaleway Container Registry product
//   - Each of your functions namespace has an associated registry namespace
//   - All your functions are available as docker image in this registry namespace
//   - Each version of your function matches a tag of this image
//
// ### Token management
//
// To get all tokens associated with a function:
//
// ```bash
// export SCW_SECRET_KEY=<your_secret_key>
// export SCW_DEFAULT_REGION="<Scaleway default region>"
// export SCW_FUNCTION_ID=<funcion_id_you_want_to_inspect>
//
//	curl -X GET \
//	  "https://api.scaleway.com/functions/v1beta1/regions/$SCW_DEFAULT_REGION/tokens?function_id=$SCW_FUNCTION_ID" \
//	  -H "accept: application/json" \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  -H "Content-Type: application/json"
//
// ```
//
// <Message type="note">
//
//	Tokens can only be read at creation time, and are not stored, hence they can not be retrieved if lost.
//
// </Message>
//
// To generate a token for a function:
//
// ```bash
// export SCW_SECRET_KEY="<API secret key>"
// export SCW_DEFAULT_REGION="<Scaleway default region>"
// export SCW_FUNCTION_ID=<funcion_id_you_want_to_add_token>
// export EXPIRE_AT=$(date -d "+90 days" +%Y-%m-%dT%H:%M:%S.000Z)
//
//	curl -X POST \
//	  "https://api.scaleway.com/functions/v1beta1/regions/$SCW_DEFAULT_REGION/tokens" \
//	  -H "accept: application/json" \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  -H "Content-Type: application/json" \
//	  -d '{
//	      descrption:desc, \
//	      expire_at:"$EXPIRE_AT", \
//	      function_id:"$SCW_FUNCTION_ID"
//	      }'
//
// ```
//
// <Message type="tip">
//
//	`expire_at` is optional, and in this example is set to 90 days from today (see `EXPIRE_AT`). If you don't set `expire_at`, your token will never expire.
//
// </Message>
//
// To revoke a token you need to delete it:
//
// ```bash
// export SCW_ACCESS_KEY="<API access key>"
// export SCW_SECRET_KEY="<API secret key>"
// export TOKEN_ID=<token_you_want_to_delete>
//
//	curl -X DELETE \
//	  "https://api.scaleway.com/functions/v1beta1/regions/$SCW_DEFAULT_REGION/tokens/$SCW_ACCESS_KEY" \
//	  -H "accept: application/json" \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  -H "Content-Type: application/json"
//
// ```
//
// ### Authentication
//
// By default, new functions are `public` meaning that no credentials are required to invoke functions.
//
// A function can be `private` or `public`. This can be configured through the `privacy` parameter.
//
// Calling a `private` function without authentication will return HTTP code `403`.
//
// ### Authentication for private functions
//
// A `private` function will:
//
// - Reject a call without an `X-Auth-Token` header, returning HTTP status Code `403`
// - Validate the token _before_ invoking your code if it is provided via the `X-Auth-Token` header
//
// For example, to execute a private function by providing a token using `curl`, you may run the following command:
//
// ```bash
//
//	curl \
//	  -H "X-Auth-Token: <generated-token>" <your-function-host>
//
// ```
//
// ### Logs
//
// Functions and containers output logs can be retrieved from the endpoint GET `/logs`. You need to pass its ID as a `function_id` parameter.
// ### Pagination
//
// Most listing requests receive a paginated response. Requests against paginated endpoints accept two `query` arguments:
//
// - `page`, a positive integer to choose which page to return.
// - `page_size`, an positive integer lower or equal to 100 to select the number of items to return per page. The default value is `20`.
//
// Paginated endpoints usually also accept filters to search and sort results. These filters are documented along each endpoint documentation.
//
// ### Regions
//
// Serverless Functions is available in the Paris, Amsterdam and Warsaw regions, which are represented by the following path parameters:
//
// * `fr-par`
// * `nl-ams`
// * `pl-waw`
//
// ## Technical limitations
// ### Supported runtimes and function lifecyle
//
// You can find all information about runtimes and functions in the [Functions' lifecycle reference page](https://www.scaleway.com/en/docs/compute/functions/reference-content/functions-lifecycle/).
// ## Going further
//
// For more help using Scaleway Serverless functions, check out the following resources:
//
// - Our [main documentation](https://www.scaleway.com/en/docs/serverless/functions/)
// - The #serverless-functions channel on our [[Slack Community](https://www.scaleway.com/en/docs/tutorials/scaleway-slack-community/)
// - Our [support ticketing system](https://www.scaleway.com/en/docs/console/my-account/how-to/open-a-support-ticket)
//
// ### Local testing
//
// Scaleway provides a number of open-source tools for offline testing. The details can be found in the [local testing docs](https://www.scaleway.com/en/docs/serverless/functions/reference-content/local-testing/).
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

// ListNamespaces: List all existing namespaces in the specified region.
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
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
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
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateNamespace: Create a new namespace in a specified Organization or Project.
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
		req.Name = namegenerator.GetRandomName("ns")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
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

// UpdateNamespace: Update the namespace associated with the specified ID.
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
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
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
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListFunctions: List all your functions.
func (s *API) ListFunctions(req *ListFunctionsRequest, opts ...scw.RequestOption) (*ListFunctionsResponse, error) {
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
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions",
		Query:  query,
	}

	var resp ListFunctionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFunction: Get the function associated with the specified ID.
func (s *API) GetFunction(req *GetFunctionRequest, opts ...scw.RequestOption) (*Function, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "",
	}

	var resp Function

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFunction: Create a new function in the specified region for a specified Organization or Project.
func (s *API) CreateFunction(req *CreateFunctionRequest, opts ...scw.RequestOption) (*Function, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("fn")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Function

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateFunction: Update the function associated with the specified ID.
func (s *API) UpdateFunction(req *UpdateFunctionRequest, opts ...scw.RequestOption) (*Function, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Function

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteFunction: Delete the function associated with the specified ID.
func (s *API) DeleteFunction(req *DeleteFunctionRequest, opts ...scw.RequestOption) (*Function, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "",
	}

	var resp Function

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeployFunction: Deploy a function associated with the specified ID.
func (s *API) DeployFunction(req *DeployFunctionRequest, opts ...scw.RequestOption) (*Function, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "/deploy",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Function

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListFunctionRuntimes: List available function runtimes.
func (s *API) ListFunctionRuntimes(req *ListFunctionRuntimesRequest, opts ...scw.RequestOption) (*ListFunctionRuntimesResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/runtimes",
	}

	var resp ListFunctionRuntimesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFunctionUploadURL: Get an upload URL of a function associated with the specified ID.
func (s *API) GetFunctionUploadURL(req *GetFunctionUploadURLRequest, opts ...scw.RequestOption) (*UploadURL, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "content_length", req.ContentLength)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "/upload-url",
		Query:  query,
	}

	var resp UploadURL

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFunctionDownloadURL: Get a download URL for a function associated with the specified ID.
func (s *API) GetFunctionDownloadURL(req *GetFunctionDownloadURLRequest, opts ...scw.RequestOption) (*DownloadURL, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "/download-url",
	}

	var resp DownloadURL

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListCrons: List all the cronjobs in a specified region.
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
	parameter.AddToQuery(query, "function_id", req.FunctionID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons",
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
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons/" + fmt.Sprint(req.CronID) + "",
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateCron: Create a new cronjob for a function with the specified ID.
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
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons",
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
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons/" + fmt.Sprint(req.CronID) + "",
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
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons/" + fmt.Sprint(req.CronID) + "",
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListLogs: List the application logs of the function with the specified ID.
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

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "/logs",
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
	parameter.AddToQuery(query, "function_id", req.FunctionID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains",
		Query:  query,
	}

	var resp ListDomainsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDomain: Get a domain name binding for the function with the specified ID.
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
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "",
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDomain: Create a domain name binding for the function with the specified ID.
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
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains",
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

// DeleteDomain: Delete a domain name binding for the function with the specified ID.
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
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "",
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
	parameter.AddToQuery(query, "function_id", req.FunctionID)
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/issue-jwt",
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
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/tokens",
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

// GetToken: Get a token.
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
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/tokens/" + fmt.Sprint(req.TokenID) + "",
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListTokens: List all tokens.
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
	parameter.AddToQuery(query, "function_id", req.FunctionID)
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/tokens",
		Query:  query,
	}

	var resp ListTokensResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteToken: Delete a token.
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
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/tokens/" + fmt.Sprint(req.TokenID) + "",
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateTrigger:
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
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers",
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

// GetTrigger:
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
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers/" + fmt.Sprint(req.TriggerID) + "",
	}

	var resp Trigger

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListTriggers:
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
	if exist && req.FunctionID == nil && req.NamespaceID == nil && req.ProjectID == nil {
		req.ProjectID = &defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "function_id", req.FunctionID)
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers",
		Query:  query,
	}

	var resp ListTriggersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateTrigger:
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
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers/" + fmt.Sprint(req.TriggerID) + "",
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

// DeleteTrigger:
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
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers/" + fmt.Sprint(req.TriggerID) + "",
	}

	var resp Trigger

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
