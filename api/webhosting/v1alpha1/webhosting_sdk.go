// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package webhosting provides methods and message types of the webhosting v1alpha1 API.
package webhosting

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

type APIListHostingsRequestOrderBy string

const (
	APIListHostingsRequestOrderByCreatedAtAsc  = APIListHostingsRequestOrderBy("created_at_asc")
	APIListHostingsRequestOrderByCreatedAtDesc = APIListHostingsRequestOrderBy("created_at_desc")
)

func (enum APIListHostingsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum APIListHostingsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *APIListHostingsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = APIListHostingsRequestOrderBy(APIListHostingsRequestOrderBy(tmp).String())
	return nil
}

type APIListOffersRequestOrderBy string

const (
	APIListOffersRequestOrderByPriceAsc = APIListOffersRequestOrderBy("price_asc")
)

func (enum APIListOffersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "price_asc"
	}
	return string(enum)
}

func (enum APIListOffersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *APIListOffersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = APIListOffersRequestOrderBy(APIListOffersRequestOrderBy(tmp).String())
	return nil
}

type DNSRecordStatus string

const (
	DNSRecordStatusUnknownStatus = DNSRecordStatus("unknown_status")
	DNSRecordStatusValid         = DNSRecordStatus("valid")
	DNSRecordStatusInvalid       = DNSRecordStatus("invalid")
)

func (enum DNSRecordStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum DNSRecordStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DNSRecordStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DNSRecordStatus(DNSRecordStatus(tmp).String())
	return nil
}

type DNSRecordType string

const (
	DNSRecordTypeUnknownType = DNSRecordType("unknown_type")
	DNSRecordTypeA           = DNSRecordType("a")
	DNSRecordTypeCname       = DNSRecordType("cname")
	DNSRecordTypeMx          = DNSRecordType("mx")
	DNSRecordTypeTxt         = DNSRecordType("txt")
	DNSRecordTypeNs          = DNSRecordType("ns")
	DNSRecordTypeAaaa        = DNSRecordType("aaaa")
)

func (enum DNSRecordType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum DNSRecordType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DNSRecordType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DNSRecordType(DNSRecordType(tmp).String())
	return nil
}

type DNSRecordsStatus string

const (
	DNSRecordsStatusUnknown = DNSRecordsStatus("unknown")
	DNSRecordsStatusValid   = DNSRecordsStatus("valid")
	DNSRecordsStatusInvalid = DNSRecordsStatus("invalid")
)

func (enum DNSRecordsStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum DNSRecordsStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DNSRecordsStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DNSRecordsStatus(DNSRecordsStatus(tmp).String())
	return nil
}

type HostingDNSStatus string

const (
	HostingDNSStatusUnknownDNSStatus = HostingDNSStatus("unknown_dns_status")
	HostingDNSStatusValid            = HostingDNSStatus("valid")
	HostingDNSStatusInvalid          = HostingDNSStatus("invalid")
)

func (enum HostingDNSStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_dns_status"
	}
	return string(enum)
}

func (enum HostingDNSStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *HostingDNSStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = HostingDNSStatus(HostingDNSStatus(tmp).String())
	return nil
}

type HostingStatus string

const (
	HostingStatusUnknownStatus = HostingStatus("unknown_status")
	HostingStatusDelivering    = HostingStatus("delivering")
	HostingStatusReady         = HostingStatus("ready")
	HostingStatusDeleting      = HostingStatus("deleting")
	HostingStatusError         = HostingStatus("error")
	HostingStatusLocked        = HostingStatus("locked")
	HostingStatusMigrating     = HostingStatus("migrating")
)

func (enum HostingStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum HostingStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *HostingStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = HostingStatus(HostingStatus(tmp).String())
	return nil
}

type NameserverStatus string

const (
	NameserverStatusUnknownStatus = NameserverStatus("unknown_status")
	NameserverStatusValid         = NameserverStatus("valid")
	NameserverStatusInvalid       = NameserverStatus("invalid")
)

func (enum NameserverStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum NameserverStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NameserverStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NameserverStatus(NameserverStatus(tmp).String())
	return nil
}

type OfferQuotaWarning string

const (
	OfferQuotaWarningUnknownQuotaWarning   = OfferQuotaWarning("unknown_quota_warning")
	OfferQuotaWarningEmailCountExceeded    = OfferQuotaWarning("email_count_exceeded")
	OfferQuotaWarningDatabaseCountExceeded = OfferQuotaWarning("database_count_exceeded")
	OfferQuotaWarningDiskUsageExceeded     = OfferQuotaWarning("disk_usage_exceeded")
)

func (enum OfferQuotaWarning) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_quota_warning"
	}
	return string(enum)
}

func (enum OfferQuotaWarning) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OfferQuotaWarning) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OfferQuotaWarning(OfferQuotaWarning(tmp).String())
	return nil
}

// HostingCpanelURLs:
type HostingCpanelURLs struct {
	// Dashboard:
	Dashboard string `json:"dashboard"`
	// Webmail:
	Webmail string `json:"webmail"`
}

// HostingOption:
type HostingOption struct {
	// ID: Option ID.
	ID string `json:"id"`
	// Name: Option name.
	Name string `json:"name"`
}

// OfferProduct:
type OfferProduct struct {
	// Name: Product name.
	Name string `json:"name"`
	// Option: Product option.
	Option bool `json:"option"`
	// EmailAccountsQuota:
	EmailAccountsQuota int32 `json:"email_accounts_quota"`
	// EmailStorageQuota:
	EmailStorageQuota int32 `json:"email_storage_quota"`
	// DatabasesQuota:
	DatabasesQuota int32 `json:"databases_quota"`
	// HostingStorageQuota:
	HostingStorageQuota uint32 `json:"hosting_storage_quota"`
	// SupportIncluded:
	SupportIncluded bool `json:"support_included"`
	// VCPU:
	VCPU uint32 `json:"v_cpu"`
	// RAM:
	RAM uint32 `json:"ram"`
}

// DNSRecord:
type DNSRecord struct {
	// Name: Record name.
	Name string `json:"name"`
	// Type: Record type.
	Type DNSRecordType `json:"type"`
	// TTL: Record time-to-live.
	TTL uint32 `json:"ttl"`
	// Value: Record value.
	Value string `json:"value"`
	// Priority: Record priority level.
	Priority *uint32 `json:"priority"`
	// Status: Record status.
	Status DNSRecordStatus `json:"status"`
}

// Nameserver:
type Nameserver struct {
	// Hostname: Hostname of the nameserver.
	Hostname string `json:"hostname"`
	// Status: Status of the nameserver.
	Status NameserverStatus `json:"status"`
	// IsDefault: Defines whether the nameserver is the default one.
	IsDefault bool `json:"is_default"`
}

// Hosting:
type Hosting struct {
	// ID: ID of the Web Hosting plan.
	ID string `json:"id"`
	// OrganizationID: ID of the Scaleway Organization the Web Hosting plan belongs to.
	OrganizationID string `json:"organization_id"`
	// ProjectID: ID of the Scaleway Project the Web Hosting plan belongs to.
	ProjectID string `json:"project_id"`
	// UpdatedAt: Date on which the Web Hosting plan was last updated.
	UpdatedAt *time.Time `json:"updated_at"`
	// CreatedAt: Date on which the Web Hosting plan was created.
	CreatedAt *time.Time `json:"created_at"`
	// Status: Status of the Web Hosting plan.
	Status HostingStatus `json:"status"`
	// PlatformHostname: Hostname of the host platform.
	PlatformHostname string `json:"platform_hostname"`
	// PlatformNumber: Number of the host platform.
	PlatformNumber *int32 `json:"platform_number"`
	// OfferID: ID of the active offer for the Web Hosting plan.
	OfferID string `json:"offer_id"`
	// OfferName: Name of the active offer for the Web Hosting plan.
	OfferName string `json:"offer_name"`
	// Domain: Main domain associated with the Web Hosting plan.
	Domain string `json:"domain"`
	// Tags: List of tags associated with the Web Hosting plan.
	Tags []string `json:"tags"`
	// Options: Array of any options activated for the Web Hosting plan.
	Options []*HostingOption `json:"options"`
	// DNSStatus: DNS status of the Web Hosting plan.
	DNSStatus HostingDNSStatus `json:"dns_status"`
	// CpanelURLs: URL to connect to cPanel dashboard and to Webmail interface.
	CpanelURLs *HostingCpanelURLs `json:"cpanel_urls"`
	// Username: Main Web Hosting cPanel username.
	Username string `json:"username"`
	// OfferEndOfLife: Indicates if the hosting offer has reached its end of life.
	OfferEndOfLife bool `json:"offer_end_of_life"`
	// Region: Region where the Web Hosting plan is hosted.
	Region scw.Region `json:"region"`
}

// Offer:
type Offer struct {
	// ID: Offer ID.
	ID string `json:"id"`
	// BillingOperationPath: Unique identifier used for billing.
	BillingOperationPath string `json:"billing_operation_path"`
	// Product: Product constituting this offer.
	Product *OfferProduct `json:"product"`
	// Price: Price of this offer.
	Price *scw.Money `json:"price"`
	// Available: If a hosting_id was specified in the call, defines whether this offer is available for that Web Hosting plan to migrate (update) to.
	Available bool `json:"available"`
	// QuotaWarnings: Quota warnings, if the offer is not available for the specified hosting_id.
	QuotaWarnings []OfferQuotaWarning `json:"quota_warnings"`
	// EndOfLife: Indicates if the offer has reached its end of life.
	EndOfLife bool `json:"end_of_life"`
}

// APICreateHostingRequest:
type APICreateHostingRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// OfferID: ID of the selected offer for the Web Hosting plan.
	OfferID string `json:"offer_id"`
	// ProjectID: ID of the Scaleway Project in which to create the Web Hosting plan.
	ProjectID string `json:"project_id"`
	// Email: Contact email for the Web Hosting client.
	Email *string `json:"email,omitempty"`
	// Tags: List of tags for the Web Hosting plan.
	Tags []string `json:"tags"`
	// Domain: Domain name to link to the Web Hosting plan. You must already own this domain name, and have completed the DNS validation process beforehand.
	Domain string `json:"domain"`
	// OptionIDs: IDs of any selected additional options for the Web Hosting plan.
	OptionIDs []string `json:"option_ids"`
}

// APIDeleteHostingRequest:
type APIDeleteHostingRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// HostingID: Hosting ID.
	HostingID string `json:"-"`
}

// APIGetDomainDNSRecordsRequest:
type APIGetDomainDNSRecordsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Domain: Domain associated with the DNS records.
	Domain string `json:"-"`
}

// APIGetHostingRequest:
type APIGetHostingRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// HostingID: Hosting ID.
	HostingID string `json:"-"`
}

// APIListHostingsRequest:
type APIListHostingsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Page: Page number to return, from the paginated results (must be a positive integer).
	Page *int32 `json:"-"`
	// PageSize: Number of Web Hosting plans to return (must be a positive integer lower or equal to 100).
	PageSize *uint32 `json:"-"`
	// OrderBy: Sort order for Web Hosting plans in the response.
	OrderBy APIListHostingsRequestOrderBy `json:"-"`
	// Tags: Tags to filter for, only Web Hosting plans with matching tags will be returned.
	Tags *[]string `json:"-"`
	// Statuses: Statuses to filter for, only Web Hosting plans with matching statuses will be returned.
	Statuses []HostingStatus `json:"-"`
	// Domain: Domain to filter for, only Web Hosting plans associated with this domain will be returned.
	Domain *string `json:"-"`
	// ProjectID: Project ID to filter for, only Web Hosting plans from this Project will be returned.
	ProjectID *string `json:"-"`
	// OrganizationID: Organization ID to filter for, only Web Hosting plans from this Organization will be returned.
	OrganizationID *string `json:"-"`
}

// APIListOffersRequest:
type APIListOffersRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// OrderBy: Sort order of offers in the response.
	OrderBy APIListOffersRequestOrderBy `json:"-"`
	// WithoutOptions: Defines whether the response should consist of offers only, without options.
	WithoutOptions bool `json:"-"`
	// OnlyOptions: Defines whether the response should consist of options only, without offers.
	OnlyOptions bool `json:"-"`
	// HostingID: ID of a Web Hosting plan, to check compatibility with returned offers (in case of wanting to update the plan).
	HostingID *string `json:"-"`
}

// APIRestoreHostingRequest:
type APIRestoreHostingRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// HostingID: Hosting ID.
	HostingID string `json:"-"`
}

// APIUpdateHostingRequest:
type APIUpdateHostingRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// HostingID: Hosting ID.
	HostingID string `json:"-"`
	// Email: New contact email for the Web Hosting plan.
	Email *string `json:"email,omitempty"`
	// Tags: New tags for the Web Hosting plan.
	Tags *[]string `json:"tags,omitempty"`
	// OptionIDs: IDs of the new options for the Web Hosting plan.
	OptionIDs *[]string `json:"option_ids,omitempty"`
	// OfferID: ID of the new offer for the Web Hosting plan.
	OfferID *string `json:"offer_id,omitempty"`
}

// DNSRecords:
type DNSRecords struct {
	// Records: List of DNS records.
	Records []*DNSRecord `json:"records"`
	// NameServers: List of nameservers.
	NameServers []*Nameserver `json:"name_servers"`
	// Status: Status of the records.
	Status DNSRecordsStatus `json:"status"`
}

// ListHostingsResponse:
type ListHostingsResponse struct {
	// TotalCount: Number of Web Hosting plans returned.
	TotalCount uint32 `json:"total_count"`
	// Hostings: List of Web Hosting plans.
	Hostings []*Hosting `json:"hostings"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListHostingsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListHostingsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListHostingsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Hostings = append(r.Hostings, results.Hostings...)
	r.TotalCount += uint32(len(results.Hostings))
	return uint32(len(results.Hostings)), nil
}

// ListOffersResponse:
type ListOffersResponse struct {
	// Offers: List of offers.
	Offers []*Offer `json:"offers"`
}

// Scaleway provides several Web Hosting plans for individuals, professionals, and everyone in between. Our Web Hosting plans include:
//
// - A domain name
// - A configurable webhosting service
// - The management of your emails, including anti-spam, antivirus and filter systems
// - Unlimited sub-domains
// - An FTP account to upload your website
// - At least one database
//
// (switchcolumn)
// (switchcolumn)
//
// ## Concepts
//
// Refer to our [dedicated concepts page](https://www.scaleway.com/en/docs/managed-services/webhosting/concepts/) to find definitions of all Web Hosting-related terminology.
//
// (switchcolumn)
// (switchcolumn)
//
// ## Quickstart
//
// 1. Configure your environment variables.
//
//	<Message type="note">
//	This is an optional step that seeks to simplify your usage of the Web Hosting API.
//	</Message>
//
//	```bash
//	export SCW_SECRET_KEY="<API secret key>"
//	export SCW_PROJECT_ID="<Scaleway Project ID>"
//	```
//
// 2. **Choose a Web Hosting offer**: run the following command to list all Web Hosting offers. The `| jq` appendage at the end of the command makes the output easier to read.
//
//	```bash
//	curl -X GET \
//	    -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	    -H "Content-Type: application/json" \
//	    "https://api.scaleway.com/webhosting/v1alpha1/regions/fr-par/offers?without_options=true" | jq
//	```
//
//	<Message type="note">
//	In the above example, we choose to get offers only. Adjust the query parameter if you would also like to get information about the different options you can add to offers.
//	</Message>
//
// 3. **Order a Web Hosting plan**: run the following command to create a Web Hosting plan. You can customize the details in the payload to your needs, using the table below to help.
//
//	```bash
//	curl -X POST \
//	    -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	    -H "Content-Type: application/json" \
//	    "https://api.scaleway.com/webhosting/v1alpha1/regions/fr-par/hostings" \
//	    -d '{
//	    "domain": "mydomain.net",
//	    "offer_id": "f5c2ae8f-7625-4bca-b711-b44bb3d08694",
//	    "project_id": "'"$SCW_PROJECT_ID"'",
//	    "tags": ["my-tag"]
//	}'
//	```
//
//	| Parameter        | Description                                                                            | Valid values OR Example     |
//	|------------------|----------------------------------------------------------------------------------------|-----------------------------|
//	| `domain`         | The domain name you would like to link to your Web Hosting account. You must already own this domain name and have completed the DNS validation process beforehand.                                                            | `mydomain.net`               |
//	| `offer_id`       | The offer ID of the desired offer. Use one of the offer IDs returned in step 2         |`f5c2ae8f-7625-4bca-b711-b44bb3d08694` |
//	| `project_id`     | The ID of the Scaleway Project to create the Web Hosting plan in.                      | `277ed74e-ea73-11ed-a05b-0242ac120003`                   |
//	| `tags`           | Any tags you would like to associate with this Web Hosting plan                         | `my-tag`                   |
//
//
//	<Message type="tip">
//	Once you've ordered your Web Hosting plan, check out our [dedicated documentation](https://www.scaleway.com/en/docs/managed-services/webhosting/quickstart/) to get started with your configuration.
//	</Message>
//
// 4. **Delete your Web Hosting plan**: run the following command to delete your Web Hosting plan. Ensure that you replace `{hosting-id}` in the URL with the ID of the Web Hosting plan you want to delete.
//
//	```bash
//	curl -X DELETE \
//	    -H "Content-Type: application/json" \
//	    -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	    "https://api.scaleway.com/rest-of-endpoint/{hosting-id}"
//	```
//
// (switchcolumn)
// <Message type="requirement">
// - You have a [Scaleway account](https://console.scaleway.com/)
// - You have created an [API key](https://www.scaleway.com/en/docs/identity-and-access-management/iam/how-to/create-api-keys/) and that the API key has sufficient [IAM permissions](https://www.scaleway.com/en/docs/identity-and-access-management/iam/reference-content/permission-sets/) to perform the actions described on this page
// - You have [installed `curl`](https://curl.se/download.html)
// - You have [installaed `jq`](https://stedolan.github.io/jq/download/)
// </Message>
// (switchcolumn)
//
// ## Technical information
//
// All Scaleway Web Hosting plans use the cPanel website management control panel tool. It provides a graphical interface and fast access icons which allow the configuration and monitoring of your hosting solutions. You can access cPanel via the [Scaleway console](https://www.scaleway.com/en/docs/managed-services/webhosting/quickstart/#how-to-access-cpanel-via-the-console).
//
// ### Regions
//
// The Scaleway Web Hosting API is a **regional** API, meaning that each call must specify in its path parameters the region for the resources concerned by the call.
//
// Scaleway Web Hosting is available only in the `fr-par` region.
//
// ## Going further
//
// For more help using Scaleway Web Hosting, check out the following resources:
// - Our [main documentation](https://www.scaleway.com/en/docs/managed-services/webhosting/)
// - The #webhosting-early-access channel on our [Slack Community](https://www.scaleway.com/en/docs/tutorials/scaleway-slack-community/)
// - Our [support ticketing system](https://www.scaleway.com/en/docs/console/my-account/how-to/open-a-support-ticket/).
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

// CreateHosting: Order a Web Hosting plan, specifying the offer type required via the `offer_id` parameter.
func (s *API) CreateHosting(req *APICreateHostingRequest, opts ...scw.RequestOption) (*Hosting, error) {
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
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/hostings",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Hosting

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListHostings: List all of your existing Web Hosting plans. Various filters are available to limit the results, including filtering by domain, status, tag and Project ID.
func (s *API) ListHostings(req *APIListHostingsRequest, opts ...scw.RequestOption) (*ListHostingsResponse, error) {
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
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "statuses", req.Statuses)
	parameter.AddToQuery(query, "domain", req.Domain)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/hostings",
		Query:  query,
	}

	var resp ListHostingsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHosting: Get the details of one of your existing Web Hosting plans, specified by its `hosting_id`.
func (s *API) GetHosting(req *APIGetHostingRequest, opts ...scw.RequestOption) (*Hosting, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HostingID) == "" {
		return nil, errors.New("field HostingID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "",
	}

	var resp Hosting

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHosting: Update the details of one of your existing Web Hosting plans, specified by its `hosting_id`. You can update parameters including the contact email address, tags, options and offer.
func (s *API) UpdateHosting(req *APIUpdateHostingRequest, opts ...scw.RequestOption) (*Hosting, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HostingID) == "" {
		return nil, errors.New("field HostingID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Hosting

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteHosting: Delete a Web Hosting plan, specified by its `hosting_id`. Note that deletion is not immediate: it will take place at the end of the calendar month, after which time your Web Hosting plan and all its data (files and emails) will be irreversibly lost.
func (s *API) DeleteHosting(req *APIDeleteHostingRequest, opts ...scw.RequestOption) (*Hosting, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HostingID) == "" {
		return nil, errors.New("field HostingID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "",
	}

	var resp Hosting

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RestoreHosting: When you [delete a Web Hosting plan](#path-hostings-delete-a-hosting), definitive deletion does not take place until the end of the calendar month. In the time between initiating the deletion, and definitive deletion at the end of the month, you can choose to **restore** the Web Hosting plan, using this endpoint and specifying its `hosting_id`.
func (s *API) RestoreHosting(req *APIRestoreHostingRequest, opts ...scw.RequestOption) (*Hosting, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HostingID) == "" {
		return nil, errors.New("field HostingID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/restore",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Hosting

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDomainDNSRecords: Get the set of DNS records of a specified domain associated with a Web Hosting plan.
func (s *API) GetDomainDNSRecords(req *APIGetDomainDNSRecordsRequest, opts ...scw.RequestOption) (*DNSRecords, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.Domain) + "/dns-records",
	}

	var resp DNSRecords

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListOffers: List the different Web Hosting offers, and their options, available to order from Scaleway.
func (s *API) ListOffers(req *APIListOffersRequest, opts ...scw.RequestOption) (*ListOffersResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "without_options", req.WithoutOptions)
	parameter.AddToQuery(query, "only_options", req.OnlyOptions)
	parameter.AddToQuery(query, "hosting_id", req.HostingID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/offers",
		Query:  query,
	}

	var resp ListOffersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
