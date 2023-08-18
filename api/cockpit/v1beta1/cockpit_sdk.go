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
	// SetupLogsRules: Permission to setup logs rules.
	SetupLogsRules bool `json:"setup_logs_rules"`
	// SetupAlerts: Permission to setup alerts.
	SetupAlerts bool `json:"setup_alerts"`
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
	ContactPoint *ContactPoint `json:"contact_point"`
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
	ContactPoint *ContactPoint `json:"contact_point"`
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
	// OrderBy:
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

// Scaleway's Cockpit allows you to monitor your applications and their infrastructure by giving you insights and context into their behavior. Cockpit also enables you to visualize your metrics and logs through a Grafana dashboard. With Cockpit, you can also push your own data as Scaleway products' data is included by default.
//
// The Observability Cockpit provides you with two Prometheus Remote Write endpoints for pushing metrics and logs. You can push metrics with any `Prometheus Remote Write` compatible agent such as [Prometheus](https://prometheus.io/docs/introduction/overview/), [Grafana](https://grafana.com/docs/agent/latest/) or [OpenTelemetry Collector](https://opentelemetry.io/docs/collector/).
//
// You can push logs with any Loki compatible agent such as [Promtail](https://grafana.com/docs/loki/latest/clients/promtail/), [Fluentd](https://docs.fluentd.org/), [Fluent Bit](https://docs.fluentbit.io/manual/) or [Logstash](https://www.elastic.co/guide/en/logstash/current/introduction.html).
//
// (switchcolumn)
// <Message type="note">
// This API concerns the Cockpit service which is currently in **Public Beta**.
// </Message>
// (switchcolumn)
//
// ## Concepts
//
// Refer to our [dedicated concepts page](https://www.scaleway.com/en/docs/observability/cockpit/concepts/) to find definitions of the different terms referring to the Observability Cockpit.
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
//	export SCW_SECRET_KEY="<API secret key>"
//	export SCW_PROJECT_ID="<Scaleway Project ID>
//	```
//
// 2. **Activate your Cockpit.**
//
//	Run the following command to activate your Cockpit:
//
//	```bash
//	curl -X POST \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  https://api.scaleway.com/cockpit/v1beta1/activate \
//	  -H "Content-Type: application/json" \
//	  -d '{"project_id": "'"$SCW_PROJECT_ID"'"}'
//	```
//
// 3.  **Build push URLs.**
//
//	A Cockpit has 4 [endpoints](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#endpoints) which are given when creating or getting a Cockpit. The endpoints look like the following:
//
//	```json
//	{
//	    [...],
//	    "endpoints": {
//	        "metrics_url": "https://metrics.cockpit.fr-par.scw.cloud",
//	        "logs_url": "https://logs.cockpit.fr-par.scw.cloud",
//	        "alertmanager_url": "https://alertmanager.cockpit.fr-par.scw.cloud"
//	        "grafana_url": "<project_id>.dashboard.obs.fr-par.scw.cloud"
//	    },
//	    [...]
//	}
//	```
//
//	The `metrics_url` is a domain that exposes a Prometheus-like API to manage metrics, and the `logs_url` exposes a Loki API to manage logs.
//
//	<Message type="important">
//	To be able to send metrics and logs, you need the **exact URL** to use with your **pushers**.
//	</Message>
//
//	- The Prometheus Remote Write endpoint to push your metrics is the following: `https://metrics.cockpit.fr-par.scw.cloud/api/v1/push`
//	- The Remote Write endpoint to push your logs is the following: `https://logs.cockpit.fr-par.scw.cloud/loki/api/v1/push`
//	- The Prometheus Alertmanager endpoint is the following: `https://alertmanager.cockpit.fr-par.scw.cloud/alertmanager`
//
// 4. **Create your token to push metrics and logs.**
//
//	Find out [how to create your token via the console](https://www.scaleway.com/en/docs/observability/cockpit/how-to/create-token/) or run the following command to create your token via API.
//
//	```bash
//	curl -X POST \
//	"https://api.scaleway.com/cockpit/v1beta1/tokens" \
//	-H "Content-Type: application/json" \
//	-H "X-Auth-Token: $SCW_SECRET_KEY" \
//	-d '{
//	  "project_id": "00000000-0000-0000-0000-000000000000",
//	  "name": "token-name",
//	  "scopes": {
//	    "query_metrics": false,
//	    "write_metrics": true,
//	    "setup_metrics_rules": false,
//	    "query_logs": false,
//	    "write_logs": true,
//	    "setup_logs_rules": false,
//	    "setup_alerts": false
//	  }
//	}'
//	```
//
//	<Message type="important">
//	  Your token's `secret_key` only displays once. Make sure you save it. We strongly recommend that you only give the minimal amount of permissions to your token. Metric pushers should only have the `write_metrics` scope, and log pushers the `write_logs` one.
//	</Message>
//
// 5. **Configure the Grafana agent.**
//
//	Find out [how to configure, start the Grafana agent and see your metrics and logs](https://www.scaleway.com/en/docs/observability/cockpit/api-cli/configuring-grafana-agent/) in our documentation.
//
//	<Message type="note">
//	The promtail configuration for the logs does not support custom headers. The `tenant_id` field corresponds to the header `X-Scope-OrgID` which
//	is one of the supported headers. For more details on the different supported headers, see our [troubleshooting documentation for when your pusher does not support custom HTTP headers](https://www.scaleway.com/en/docs/observability/cockpit/troubleshooting/pusher-does-not-support-custom-http-headers/).
//	</Message>
//
// 6. **Configure rules and alerts.**
//
//	Upon its creation, a Grafana instance is already configured by default with logs and metrics data sources. The alert manager data source is also configured. This means that you will be able to configure your alerting rules, for metrics and logs, configure your contact points and notification policies and manage your alert silences from Grafana.
//
//	To do so, on your Grafana instance, click the **Bell** icon on the left of your screen. The `Alerting` page displays with tabs. Refer to the **External links** section of this page if you are not familiar with the alert manager or alert rules.
//
//	<Message type="important">
//	For the alerts and rules to be created in your Cockpit, you must use the data sources provided by Scaleway.
//	</Message>
//
//	In the `Alert rules` tab, you must select `Mimir or Loki alert` when creating your alert rules in the `Rule type` field. Then, on the `Select data source` field that appears on the right, select the data source you wish to create the rule on.
//
//	For all the other tabs, from `Contact points` to `Alert groups`, you will be configuring the **alert manager**. The `Choose Alertmananger` field displays at the top of the screen.
//
//	<Message type="important">
//	Make sure that you select `Scaleway Alerting` as your alert manager as the alerts are managed by the platform. If you use `Grafana` you will work with the local alert manager which is deactivated.
//	</Message>
//
// (switchcolumn)
// <Message type="requirement">
//
//   - You have a [Scaleway account](https://console.scaleway.com/)
//   - You have [created an API key](https://www.scaleway.com/en/docs/identity-and-access-management/iam/how-to/create-api-keys/) and that the API key has sufficient [IAM permissions](https://www.scaleway.com/en/docs/identity-and-access-management/iam/reference-content/permission-sets/) to perform the actions described on this page
//   - You have activated your Cockpit via the [console](https://www.scaleway.com/en/docs/observability/cockpit/how-to/activate-cockpit/) or [API](#path-cockpits-activate-a-cockpit)
//   - You have [installed `curl`](https://curl.se/download.html)
//   - You have [installed docker](https://www.docker.com/)
//   - You have [installed docker compose](https://docs.docker.com/compose/install/)
//   - You have a Linux machine to run the container on. You can still run the container on macOS or Windows, but your logs will not be available in your Cockpit.
//
// </Message>
// (switchcolumn)
//
// ## Technical information
//
// ### Regional availability
//
// Scaleway's Cockpit is available globally. Find out about our [product availability in our dedicated documentation](https://grafana.com/docs/mimir/latest/references/http-api/).
//
// ### Current supported endpoints
//
//   - The current supported endpoints for metrics are the following. For more information about the endpoints, refer to the [Mimir documentation](https://grafana.com/docs/mimir/latest/references/http-api/)
//     ```
//     # Remote write endpoints
//     path /api/v1/push
//
//     # Query endpoints
//     path /prometheus/api/v1/query
//     path /prometheus/api/v1/query_range
//     path /prometheus/api/v1/query_exemplars
//     path /prometheus/api/v1/series
//     path /prometheus/api/v1/labels
//     path /prometheus/api/v1/label/*
//     path /prometheus/api/v1/metadata
//     path /prometheus/api/v1/read
//     path /prometheus/api/v1/status/buildinfo
//
//     # Ruler endpoints
//     path /prometheus/api/v1/rules
//     path /prometheus/api/v1/alerts
//     path /prometheus/config/v1/rules
//     path /prometheus/config/v1/rules/*
//     ```
//
//   - The current supported endpoints for logs are the following. For more information about the endpoints, refer to the [Loki documentation](https://grafana.com/docs/loki/latest/api/)
//     ```
//     # Remote write endpoints
//     path /loki/api/v1/push
//
//     # Query endpoints
//     path /loki/api/v1/query
//     path /loki/api/v1/query_range
//     path /loki/api/v1/labels
//     path /loki/api/v1/label
//     path /loki/api/v1/label/*
//     path /loki/api/v1/tail
//     path /loki/api/v1/series
//
//     # Ruler endpoints
//     path /loki/api/v1/rules
//     path /loki/api/v1/rules/*
//     path /api/prom/rules
//     path /api/prom/rules/*
//     path /prometheus/api/v1/rules
//     path /prometheus/api/v1/alerts
//     ```
//
//   - The current supported endpoints for the alert manager are the following. For more information on the endpoints, refer to the [Prometheus documentation](https://prometheus.io/docs/prometheus/latest/querying/api/#alerts)
//     ```
//     # Prometheus alertmanager
//     path /alertmanager/*
//
//     # Alertmanager configuration (see Mimir doc)
//     path /api/v1/alerts
//     ```
//
// ### Troubleshooting
//
// Refer to our troubleshooting documentation if [your pusher does not support HTTP headers](https://www.scaleway.com/en/docs/observability/cockpit/troubleshooting/pusher-does-not-support-custom-http-headers/) or if you want to [reset your Grafana password.](https://www.scaleway.com/en/docs/observability/cockpit/troubleshooting/resetting-grafana-password-via-the-api/)
//
// ## Technical limitations
//
// - Metrics and logs data is retained **31 days**. After this period, data older than 31 days is deleted.
// - The number of **active metrics time series** is limited to **250,000 per Cockpit** by default.
// - The number of **active log streams** is limited to **5000 per Cockpit**.
// - It is not yet possible to downsample metrics.
//
// ## Going further
//
// For more information about Cockpit, you can check out the following pages:
//
// - [Cockpit Documentation](https://www.scaleway.com/en/docs/observability/cockpit/)
// - [Scaleway Slack Community](http://slack.scaleway.com) join the #observability-beta and the #observability channels
// - [Contact our support team](https://console.scaleway.com/support/tickets)
//
// ### External links
//
// If you are interested in learning more, you can check out the following pages:
//
// - [Introduction to Prometheus](https://prometheus.io/docs/introduction/overview/). This page gives a general introduction to what Prometheus is.
// - [Introduction to Promtail](https://grafana.com/docs/loki/latest/clients/promtail/). Promtail is an agent that collects and forwards logs to a Grafana Loki instance.
// - [Introduction to Grafana Agent](https://grafana.com/docs/grafana-cloud/agent/). Grafana Agent is an agent that embeds multiple other agents like Prometheus or Promtail.
// - [Pushing logs using Fluentbit](https://docs.fluentbit.io/manual/pipeline/outputs/loki). Fluentbit is another commonly used agent that collects and forwards logs.
// - [Pushing metrics using Prometheus's agent mode](https://prometheus.io/blog/2021/11/16/agent/). Prometheus can be used in an "agent-mode" that can scrape local targets and forward them using remote-write. This is useful for scenarios where targets are exposed in a private network (like a VPN or a single host) but need to be forwarded outside of it.
// - [Prometheus remote write configuration](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#remote_write). This page gives the exact configuration parameters that can be used to configure Prometheus's remote-write.
// - [Prometheus data model](https://prometheus.io/docs/concepts/data_model/). This page explains the Prometheus data model to help understand the concepts of metrics, time series, labels and samples, and how they work together.
// - [Article on Prometheus label cardinality and performance](https://www.robustperception.io/cardinality-is-key). When working with metrics, the cardinality, which is the number time series for a given metric, is extremely important. This article explains why.
// - [Integrations supported by Grafana Agent](https://grafana.com/docs/agent/latest/configuration/integrations/). Grafana Agent has many other embedded agents that can be used to monitor various software, like PostgreSQL or MongoDB. All of them can be found on the left menu of this page.
// - [Introduction to alerting using Prometheus](https://prometheus.io/docs/alerting/latest/overview/)
// - [How to create a Mimir or Loki managed alerting rule using Grafana](https://grafana.com/docs/grafana/latest/alerting/alerting-rules/create-mimir-loki-managed-rule/)
// - [How to manage alerting rules using Grafana](https://grafana.com/docs/grafana/latest/alerting/unified-alerting/alerting-rules/rule-list/)
// - [How to manage contact points using Grafana](https://grafana.com/docs/grafana/latest/alerting/unified-alerting/contact-points/)
// - [How to manage notification policies using Grafana](https://grafana.com/docs/grafana/latest/alerting/unified-alerting/notifications/)
// - [How to manage silences using Grafana](https://grafana.com/docs/grafana/latest/alerting/unified-alerting/silences/).
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
