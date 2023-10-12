// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package cockpit provides methods and message types of the cockpit v1beta1 API.
package cockpit

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

type CockpitStatus string

const (
	CockpitStatusUnknownStatus = CockpitStatus("unknown_status")
	CockpitStatusCreating      = CockpitStatus("creating")
	CockpitStatusReady         = CockpitStatus("ready")
	CockpitStatusDeleting      = CockpitStatus("deleting")
	CockpitStatusUpdating      = CockpitStatus("updating")
	CockpitStatusError         = CockpitStatus("error")
)

func (enum CockpitStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum CockpitStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CockpitStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CockpitStatus(CockpitStatus(tmp).String())
	return nil
}

type DatasourceType string

const (
	DatasourceTypeUnknownDatasourceType = DatasourceType("unknown_datasource_type")
	DatasourceTypeMetrics               = DatasourceType("metrics")
	DatasourceTypeLogs                  = DatasourceType("logs")
	DatasourceTypeTraces                = DatasourceType("traces")
	DatasourceTypeAlerts                = DatasourceType("alerts")
)

func (enum DatasourceType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_datasource_type"
	}
	return string(enum)
}

func (enum DatasourceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DatasourceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DatasourceType(DatasourceType(tmp).String())
	return nil
}

type GrafanaUserRole string

const (
	GrafanaUserRoleUnknownRole = GrafanaUserRole("unknown_role")
	GrafanaUserRoleEditor      = GrafanaUserRole("editor")
	GrafanaUserRoleViewer      = GrafanaUserRole("viewer")
)

func (enum GrafanaUserRole) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_role"
	}
	return string(enum)
}

func (enum GrafanaUserRole) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *GrafanaUserRole) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = GrafanaUserRole(GrafanaUserRole(tmp).String())
	return nil
}

type ListDatasourcesRequestOrderBy string

const (
	ListDatasourcesRequestOrderByCreatedAtAsc  = ListDatasourcesRequestOrderBy("created_at_asc")
	ListDatasourcesRequestOrderByCreatedAtDesc = ListDatasourcesRequestOrderBy("created_at_desc")
	ListDatasourcesRequestOrderByNameAsc       = ListDatasourcesRequestOrderBy("name_asc")
	ListDatasourcesRequestOrderByNameDesc      = ListDatasourcesRequestOrderBy("name_desc")
)

func (enum ListDatasourcesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListDatasourcesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDatasourcesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDatasourcesRequestOrderBy(ListDatasourcesRequestOrderBy(tmp).String())
	return nil
}

type ListGrafanaUsersRequestOrderBy string

const (
	ListGrafanaUsersRequestOrderByLoginAsc  = ListGrafanaUsersRequestOrderBy("login_asc")
	ListGrafanaUsersRequestOrderByLoginDesc = ListGrafanaUsersRequestOrderBy("login_desc")
)

func (enum ListGrafanaUsersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "login_asc"
	}
	return string(enum)
}

func (enum ListGrafanaUsersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListGrafanaUsersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListGrafanaUsersRequestOrderBy(ListGrafanaUsersRequestOrderBy(tmp).String())
	return nil
}

type ListPlansRequestOrderBy string

const (
	ListPlansRequestOrderByNameAsc  = ListPlansRequestOrderBy("name_asc")
	ListPlansRequestOrderByNameDesc = ListPlansRequestOrderBy("name_desc")
)

func (enum ListPlansRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListPlansRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPlansRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPlansRequestOrderBy(ListPlansRequestOrderBy(tmp).String())
	return nil
}

type ListTokensRequestOrderBy string

const (
	ListTokensRequestOrderByCreatedAtAsc  = ListTokensRequestOrderBy("created_at_asc")
	ListTokensRequestOrderByCreatedAtDesc = ListTokensRequestOrderBy("created_at_desc")
	ListTokensRequestOrderByNameAsc       = ListTokensRequestOrderBy("name_asc")
	ListTokensRequestOrderByNameDesc      = ListTokensRequestOrderBy("name_desc")
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

type PlanName string

const (
	PlanNameUnknownName = PlanName("unknown_name")
	PlanNameFree        = PlanName("free")
	PlanNamePremium     = PlanName("premium")
	PlanNameCustom      = PlanName("custom")
)

func (enum PlanName) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_name"
	}
	return string(enum)
}

func (enum PlanName) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PlanName) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PlanName(PlanName(tmp).String())
	return nil
}

// ContactPointEmail:
type ContactPointEmail struct {
	// To:
	To string `json:"to"`
}

// TokenScopes:
type TokenScopes struct {
	// QueryMetrics: Permission to fetch metrics.
	QueryMetrics bool `json:"query_metrics"`
	// WriteMetrics: Permission to write metrics.
	WriteMetrics bool `json:"write_metrics"`
	// SetupMetricsRules: Permission to setup metrics rules.
	SetupMetricsRules bool `json:"setup_metrics_rules"`
	// QueryLogs: Permission to fetch logs.
	QueryLogs bool `json:"query_logs"`
	// WriteLogs: Permission to write logs.
	WriteLogs bool `json:"write_logs"`
	// SetupLogsRules: Permission to set up logs rules.
	SetupLogsRules bool `json:"setup_logs_rules"`
	// SetupAlerts: Permission to set up alerts.
	SetupAlerts bool `json:"setup_alerts"`
	// QueryTraces: Permission to fetch traces.
	QueryTraces bool `json:"query_traces"`
	// WriteTraces: Permission to write traces.
	WriteTraces bool `json:"write_traces"`
}

// CockpitEndpoints:
type CockpitEndpoints struct {
	// MetricsURL: URL for metrics.
	MetricsURL string `json:"metrics_url"`
	// LogsURL: URL for logs.
	LogsURL string `json:"logs_url"`
	// AlertmanagerURL: URL for the alert manager.
	AlertmanagerURL string `json:"alertmanager_url"`
	// GrafanaURL: URL for the Grafana dashboard.
	GrafanaURL string `json:"grafana_url"`
}

// Plan: Pricing plan.
type Plan struct {
	// ID: ID of a given pricing plan.
	ID string `json:"id"`
	// Name: Name of a given pricing plan.
	Name PlanName `json:"name"`
	// RetentionMetricsInterval: Retention for metrics.
	RetentionMetricsInterval *scw.Duration `json:"retention_metrics_interval"`
	// RetentionLogsInterval: Retention for logs.
	RetentionLogsInterval *scw.Duration `json:"retention_logs_interval"`
	// SampleIngestionPrice: Ingestion price for 1 million samples in cents.
	SampleIngestionPrice uint32 `json:"sample_ingestion_price"`
	// LogsIngestionPrice: Ingestion price for 1 GB of logs in cents.
	LogsIngestionPrice uint32 `json:"logs_ingestion_price"`
	// RetentionPrice: Retention price in euros per month.
	RetentionPrice uint32 `json:"retention_price"`
}

// ContactPoint: Contact point.
type ContactPoint struct {
	// Email: Contact point configuration.
	Email *ContactPointEmail `json:"email,omitempty"`
}

// Datasource: Datasource.
type Datasource struct {
	// ID: ID of the datasource.
	ID string `json:"id"`
	// ProjectID: ID of the Project the Cockpit belongs to.
	ProjectID string `json:"project_id"`
	// Name: Datasource name.
	Name string `json:"name"`
	// URL: Datasource URL.
	URL string `json:"url"`
	// Type: Datasource type.
	Type DatasourceType `json:"type"`
}

// GrafanaProductDashboard: Grafana dashboard.
type GrafanaProductDashboard struct {
	// DashboardName: Name of the dashboard.
	DashboardName string `json:"dashboard_name"`
	// Title: Title of the dashboard.
	Title string `json:"title"`
	// URL: URL of the dashboard.
	URL string `json:"url"`
	// Tags: Tags of the dashboard.
	Tags []string `json:"tags"`
	// Variables: Variables of the dashboard.
	Variables []string `json:"variables"`
}

// GrafanaUser: Grafana user.
type GrafanaUser struct {
	// ID: ID of the Grafana user.
	ID uint32 `json:"id"`
	// Login: Username of the Grafana user.
	Login string `json:"login"`
	// Role: Role assigned to the Grafana user.
	Role GrafanaUserRole `json:"role"`
	// Password: The Grafana user's password.
	Password *string `json:"password"`
}

// Token:
type Token struct {
	// ID: ID of the token.
	ID string `json:"id"`
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
	// Name: Name of the token.
	Name string `json:"name"`
	// CreatedAt: Date and time of the token's creation.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Date and time of the token's last update.
	UpdatedAt *time.Time `json:"updated_at"`
	// Scopes: Token's permissions.
	Scopes *TokenScopes `json:"scopes"`
	// SecretKey: Token's secret key.
	SecretKey *string `json:"secret_key"`
}

// ActivateCockpitRequest:
type ActivateCockpitRequest struct {
	// ProjectID: ID of the Project the Cockpit belongs to.
	ProjectID string `json:"project_id"`
}

// Cockpit: Cockpit.
type Cockpit struct {
	// ProjectID: ID of the Project the Cockpit belongs to.
	ProjectID string `json:"project_id"`
	// CreatedAt: Date and time of the Cockpit's creation.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Date and time of the Cockpit's last update.
	UpdatedAt *time.Time `json:"updated_at"`
	// Endpoints: Endpoints of the Cockpit.
	Endpoints *CockpitEndpoints `json:"endpoints"`
	// Status: Status of the Cockpit.
	Status CockpitStatus `json:"status"`
	// ManagedAlertsEnabled: Specifies whether managed alerts are enabled or disabled.
	ManagedAlertsEnabled bool `json:"managed_alerts_enabled"`
	// Plan: Pricing plan information.
	Plan *Plan `json:"plan"`
}

// CockpitMetrics: Metrics for a given Cockpit.
type CockpitMetrics struct {
	// Timeseries: Time series array.
	Timeseries []*scw.TimeSeries `json:"timeseries"`
}

// CreateContactPointRequest: Request to create a contact point.
type CreateContactPointRequest struct {
	// ProjectID: ID of the Project in which to create the contact point.
	ProjectID string `json:"project_id"`
	// ContactPoint: Contact point to create.
	ContactPoint *ContactPoint `json:"contact_point,omitempty"`
}

// CreateDatasourceRequest: Request to create a datasource.
type CreateDatasourceRequest struct {
	// ProjectID: ID of the Project the Cockpit belongs to.
	ProjectID string `json:"project_id"`
	// Name: Datasource name.
	Name string `json:"name"`
	// Type: Datasource type.
	Type DatasourceType `json:"type"`
}

// CreateGrafanaUserRequest: Request to create a Grafana user.
type CreateGrafanaUserRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
	// Login: Username of the Grafana user.
	Login string `json:"login"`
	// Role: Role assigned to the Grafana user.
	Role GrafanaUserRole `json:"role"`
}

// CreateTokenRequest:
type CreateTokenRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
	// Name: Name of the token.
	Name string `json:"name"`
	// Scopes: Token's permissions.
	Scopes *TokenScopes `json:"scopes"`
}

// DeactivateCockpitRequest:
type DeactivateCockpitRequest struct {
	// ProjectID: ID of the Project the Cockpit belongs to.
	ProjectID string `json:"project_id"`
}

// DeleteContactPointRequest: Request to delete a contact point.
type DeleteContactPointRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
	// ContactPoint: Contact point to delete.
	ContactPoint *ContactPoint `json:"contact_point,omitempty"`
}

// DeleteGrafanaUserRequest: Request to delete a Grafana user.
type DeleteGrafanaUserRequest struct {
	// GrafanaUserID: ID of the Grafana user.
	GrafanaUserID uint32 `json:"-"`
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// DeleteTokenRequest:
type DeleteTokenRequest struct {
	// TokenID: ID of the token.
	TokenID string `json:"-"`
}

// DisableManagedAlertsRequest: Request to disable the sending of managed alerts.
type DisableManagedAlertsRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// EnableManagedAlertsRequest: Request to enable the sending of managed alerts.
type EnableManagedAlertsRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// GetCockpitMetricsRequest: Request to get a given Cockpit's metrics.
type GetCockpitMetricsRequest struct {
	// ProjectID: ID of the Project the Cockpit belongs to.
	ProjectID string `json:"-"`
	// StartDate: Desired time range's start date for the metrics.
	StartDate *time.Time `json:"-"`
	// EndDate: Desired time range's end date for the metrics.
	EndDate *time.Time `json:"-"`
	// MetricName: Name of the metric requested.
	MetricName *string `json:"-"`
}

// GetCockpitRequest:
type GetCockpitRequest struct {
	// ProjectID: ID of the Project the Cockpit belongs to.
	ProjectID string `json:"-"`
}

// GetGrafanaProductDashboardRequest: Request to get a dashboard.
type GetGrafanaProductDashboardRequest struct {
	// DashboardName: Name of the dashboard.
	DashboardName string `json:"-"`
	// ProjectID: ID of the Project.
	ProjectID string `json:"-"`
}

// GetTokenRequest:
type GetTokenRequest struct {
	// TokenID: ID of the token.
	TokenID string `json:"-"`
}

// ListContactPointsRequest: Request to list all contact points.
type ListContactPointsRequest struct {
	// Page: Page number.
	Page *int32 `json:"-"`
	// PageSize: Page size.
	PageSize *uint32 `json:"-"`
	// ProjectID: ID of the Project from which to list the contact points.
	ProjectID string `json:"-"`
}

// ListContactPointsResponse: Response returned when listing contact points.
type ListContactPointsResponse struct {
	// TotalCount: Count of all contact points created.
	TotalCount uint32 `json:"total_count"`
	// ContactPoints: Array of contact points.
	ContactPoints []*ContactPoint `json:"contact_points"`
	// HasAdditionalReceivers: Specifies whether the contact point has other receivers than the default receiver.
	HasAdditionalReceivers bool `json:"has_additional_receivers"`
	// HasAdditionalContactPoints: Specifies whether there are unmanaged contact points.
	HasAdditionalContactPoints bool `json:"has_additional_contact_points"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListContactPointsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListContactPointsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListContactPointsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ContactPoints = append(r.ContactPoints, results.ContactPoints...)
	r.TotalCount += uint32(len(results.ContactPoints))
	return uint32(len(results.ContactPoints)), nil
}

// ListDatasourcesRequest:
type ListDatasourcesRequest struct {
	// Page: Page number.
	Page *int32 `json:"-"`
	// PageSize: Page size.
	PageSize *uint32 `json:"-"`
	// OrderBy: How the response is ordered.
	OrderBy ListDatasourcesRequestOrderBy `json:"-"`
	// ProjectID: ID of the Project.
	ProjectID string `json:"-"`
	// Types: Filter by datasource types.
	Types []DatasourceType `json:"-"`
}

// ListDatasourcesResponse:
type ListDatasourcesResponse struct {
	// TotalCount: Count of all datasources corresponding to the request.
	TotalCount uint32 `json:"total_count"`
	// Datasources: List of the datasources within the pagination.
	Datasources []*Datasource `json:"datasources"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDatasourcesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDatasourcesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDatasourcesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Datasources = append(r.Datasources, results.Datasources...)
	r.TotalCount += uint32(len(results.Datasources))
	return uint32(len(results.Datasources)), nil
}

// ListGrafanaProductDashboardsRequest: Request to get a list of dashboards.
type ListGrafanaProductDashboardsRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"-"`
	// Page: Page number.
	Page *int32 `json:"-"`
	// PageSize: Page size.
	PageSize *uint32 `json:"-"`
	// Tags: Tags to filter the dashboards.
	Tags []string `json:"-"`
}

// ListGrafanaProductDashboardsResponse: Response returned when getting a list of dashboards.
type ListGrafanaProductDashboardsResponse struct {
	// TotalCount: Count of grafana dasboards.
	TotalCount uint64 `json:"total_count"`
	// Dashboards: Information on grafana dashboards.
	Dashboards []*GrafanaProductDashboard `json:"dashboards"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListGrafanaProductDashboardsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListGrafanaProductDashboardsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListGrafanaProductDashboardsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Dashboards = append(r.Dashboards, results.Dashboards...)
	r.TotalCount += uint64(len(results.Dashboards))
	return uint64(len(results.Dashboards)), nil
}

// ListGrafanaUsersRequest: Request to list all Grafana users.
type ListGrafanaUsersRequest struct {
	// Page: Page number.
	Page *int32 `json:"-"`
	// PageSize: Page size.
	PageSize *uint32 `json:"-"`
	// OrderBy:
	OrderBy ListGrafanaUsersRequestOrderBy `json:"-"`
	// ProjectID: ID of the Project.
	ProjectID string `json:"-"`
}

// ListGrafanaUsersResponse: Response returned when listing Grafana users.
type ListGrafanaUsersResponse struct {
	// TotalCount: Count of all Grafana users.
	TotalCount uint32 `json:"total_count"`
	// GrafanaUsers: Information on all Grafana users.
	GrafanaUsers []*GrafanaUser `json:"grafana_users"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListGrafanaUsersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListGrafanaUsersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListGrafanaUsersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.GrafanaUsers = append(r.GrafanaUsers, results.GrafanaUsers...)
	r.TotalCount += uint32(len(results.GrafanaUsers))
	return uint32(len(results.GrafanaUsers)), nil
}

// ListPlansRequest: Request to list all pricing plans.
type ListPlansRequest struct {
	// Page: Page number.
	Page *int32 `json:"-"`
	// PageSize: Page size.
	PageSize *uint32 `json:"-"`
	// OrderBy:
	OrderBy ListPlansRequestOrderBy `json:"-"`
}

// ListPlansResponse: Response returned when listing all pricing plans.
type ListPlansResponse struct {
	// TotalCount: Count of all pricing plans.
	TotalCount uint64 `json:"total_count"`
	// Plans: Information on plans.
	Plans []*Plan `json:"plans"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPlansResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPlansResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListPlansResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Plans = append(r.Plans, results.Plans...)
	r.TotalCount += uint64(len(results.Plans))
	return uint64(len(results.Plans)), nil
}

// ListTokensRequest:
type ListTokensRequest struct {
	// Page: Page number.
	Page *int32 `json:"-"`
	// PageSize: Page size.
	PageSize *uint32 `json:"-"`
	// OrderBy: How the response is ordered.
	OrderBy ListTokensRequestOrderBy `json:"-"`
	// ProjectID: ID of the Project.
	ProjectID string `json:"-"`
}

// ListTokensResponse:
type ListTokensResponse struct {
	// TotalCount: Count of all tokens created.
	TotalCount uint32 `json:"total_count"`
	// Tokens: List of all tokens created.
	Tokens []*Token `json:"tokens"`
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

// ResetCockpitGrafanaRequest:
type ResetCockpitGrafanaRequest struct {
	// ProjectID: ID of the Project the Cockpit belongs to.
	ProjectID string `json:"project_id"`
}

// ResetGrafanaUserPasswordRequest: Request to reset a Grafana user's password.
type ResetGrafanaUserPasswordRequest struct {
	// GrafanaUserID: ID of the Grafana user.
	GrafanaUserID uint32 `json:"-"`
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// SelectPlanRequest: Request to select a specific pricing plan.
type SelectPlanRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
	// PlanID: ID of the pricing plan.
	PlanID string `json:"plan_id"`
}

// SelectPlanResponse: Response returned when selecting a pricing plan.
type SelectPlanResponse struct {
}

// TriggerTestAlertRequest:
type TriggerTestAlertRequest struct {
	// ProjectID:
	ProjectID string `json:"project_id"`
}

// Cockpit's API allows you to activate your Cockpit on your Projects. Scaleway's Cockpit stores metrics and logs and provides a dedicated Grafana for dashboarding to visualize them.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// ActivateCockpit: Activate the Cockpit of the specified Project ID.
func (s *API) ActivateCockpit(req *ActivateCockpitRequest, opts ...scw.RequestOption) (*Cockpit, error) {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/activate",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cockpit

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCockpit: Retrieve the Cockpit of the specified Project ID.
func (s *API) GetCockpit(req *GetCockpitRequest, opts ...scw.RequestOption) (*Cockpit, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1beta1/cockpit",
		Query:  query,
	}

	var resp Cockpit

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCockpitMetrics: Get metrics from your Cockpit with the specified Project ID.
func (s *API) GetCockpitMetrics(req *GetCockpitMetricsRequest, opts ...scw.RequestOption) (*CockpitMetrics, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "start_date", req.StartDate)
	parameter.AddToQuery(query, "end_date", req.EndDate)
	parameter.AddToQuery(query, "metric_name", req.MetricName)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1beta1/cockpit/metrics",
		Query:  query,
	}

	var resp CockpitMetrics

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeactivateCockpit: Deactivate the Cockpit of the specified Project ID.
func (s *API) DeactivateCockpit(req *DeactivateCockpitRequest, opts ...scw.RequestOption) (*Cockpit, error) {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/deactivate",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cockpit

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ResetCockpitGrafana: Reset your Cockpit's Grafana associated with the specified Project ID.
func (s *API) ResetCockpitGrafana(req *ResetCockpitGrafanaRequest, opts ...scw.RequestOption) (*Cockpit, error) {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/reset-grafana",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cockpit

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDatasource: Create a datasource for the specified Project ID and the given type.
func (s *API) CreateDatasource(req *CreateDatasourceRequest, opts ...scw.RequestOption) (*Datasource, error) {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/datasources",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Datasource

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDatasources: Get a list of datasources for the specified Project ID.
func (s *API) ListDatasources(req *ListDatasourcesRequest, opts ...scw.RequestOption) (*ListDatasourcesResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "types", req.Types)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1beta1/datasources",
		Query:  query,
	}

	var resp ListDatasourcesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateToken: Create a token associated with the specified Project ID.
func (s *API) CreateToken(req *CreateTokenRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("token")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/tokens",
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

// ListTokens: Get a list of tokens associated with the specified Project ID.
func (s *API) ListTokens(req *ListTokensRequest, opts ...scw.RequestOption) (*ListTokensResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1beta1/tokens",
		Query:  query,
	}

	var resp ListTokensResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetToken: Retrieve the token associated with the specified token ID.
func (s *API) GetToken(req *GetTokenRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error

	if fmt.Sprint(req.TokenID) == "" {
		return nil, errors.New("field TokenID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1beta1/tokens/" + fmt.Sprint(req.TokenID) + "",
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteToken: Delete the token associated with the specified token ID.
func (s *API) DeleteToken(req *DeleteTokenRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.TokenID) == "" {
		return errors.New("field TokenID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/cockpit/v1beta1/tokens/" + fmt.Sprint(req.TokenID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CreateContactPoint: Create a contact point to receive alerts for the default receiver.
func (s *API) CreateContactPoint(req *CreateContactPointRequest, opts ...scw.RequestOption) (*ContactPoint, error) {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/contact-points",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ContactPoint

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListContactPoints: Get a list of contact points for the Cockpit associated with the specified Project ID.
func (s *API) ListContactPoints(req *ListContactPointsRequest, opts ...scw.RequestOption) (*ListContactPointsResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1beta1/contact-points",
		Query:  query,
	}

	var resp ListContactPointsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteContactPoint: Delete a contact point for the default receiver.
func (s *API) DeleteContactPoint(req *DeleteContactPointRequest, opts ...scw.RequestOption) error {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/delete-contact-point",
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

// EnableManagedAlerts: Enable the sending of managed alerts for the specified Project's Cockpit.
func (s *API) EnableManagedAlerts(req *EnableManagedAlertsRequest, opts ...scw.RequestOption) error {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/enable-managed-alerts",
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

// DisableManagedAlerts: Disable the sending of managed alerts for the specified Project's Cockpit.
func (s *API) DisableManagedAlerts(req *DisableManagedAlertsRequest, opts ...scw.RequestOption) error {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/disable-managed-alerts",
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

// TriggerTestAlert: Trigger a test alert to all of the Cockpit's receivers.
func (s *API) TriggerTestAlert(req *TriggerTestAlertRequest, opts ...scw.RequestOption) error {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/trigger-test-alert",
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

// CreateGrafanaUser: Create a Grafana user for your Cockpit's Grafana instance. Make sure you save the automatically-generated password and the Grafana user ID.
func (s *API) CreateGrafanaUser(req *CreateGrafanaUserRequest, opts ...scw.RequestOption) (*GrafanaUser, error) {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/grafana-users",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp GrafanaUser

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListGrafanaUsers: Get a list of Grafana users who are able to connect to the Cockpit's Grafana instance.
func (s *API) ListGrafanaUsers(req *ListGrafanaUsersRequest, opts ...scw.RequestOption) (*ListGrafanaUsersResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1beta1/grafana-users",
		Query:  query,
	}

	var resp ListGrafanaUsersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteGrafanaUser: Delete a Grafana user from a Grafana instance, specified by the Cockpit's Project ID and the Grafana user ID.
func (s *API) DeleteGrafanaUser(req *DeleteGrafanaUserRequest, opts ...scw.RequestOption) error {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.GrafanaUserID) == "" {
		return errors.New("field GrafanaUserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/grafana-users/" + fmt.Sprint(req.GrafanaUserID) + "/delete",
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

// ResetGrafanaUserPassword: Reset a Grafana user's password specified by the Cockpit's Project ID and the Grafana user ID.
func (s *API) ResetGrafanaUserPassword(req *ResetGrafanaUserPasswordRequest, opts ...scw.RequestOption) (*GrafanaUser, error) {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.GrafanaUserID) == "" {
		return nil, errors.New("field GrafanaUserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/grafana-users/" + fmt.Sprint(req.GrafanaUserID) + "/reset-password",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp GrafanaUser

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPlans: Get a list of all pricing plans available.
func (s *API) ListPlans(req *ListPlansRequest, opts ...scw.RequestOption) (*ListPlansResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1beta1/plans",
		Query:  query,
	}

	var resp ListPlansResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SelectPlan: Select your chosen pricing plan for your Cockpit, specifying the Cockpit's Project ID and the pricing plan's ID in the request.
func (s *API) SelectPlan(req *SelectPlanRequest, opts ...scw.RequestOption) (*SelectPlanResponse, error) {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/select-plan",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SelectPlanResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListGrafanaProductDashboards: Get a list of available product dashboards.
func (s *API) ListGrafanaProductDashboards(req *ListGrafanaProductDashboardsRequest, opts ...scw.RequestOption) (*ListGrafanaProductDashboardsResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "tags", req.Tags)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1beta1/grafana-product-dashboards",
		Query:  query,
	}

	var resp ListGrafanaProductDashboardsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGrafanaProductDashboard: Get a product dashboard specified by the dashboard ID.
func (s *API) GetGrafanaProductDashboard(req *GetGrafanaProductDashboardRequest, opts ...scw.RequestOption) (*GrafanaProductDashboard, error) {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.DashboardName) == "" {
		return nil, errors.New("field DashboardName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1beta1/grafana-product-dashboards/" + fmt.Sprint(req.DashboardName) + "",
		Query:  query,
	}

	var resp GrafanaProductDashboard

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
