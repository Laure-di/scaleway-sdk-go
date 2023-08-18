// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package billing provides methods and message types of the billing v2alpha1 API.
package billing

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

type APIDownloadInvoiceRequestFileType string

const (
	APIDownloadInvoiceRequestFileTypePdf = APIDownloadInvoiceRequestFileType("pdf")
)

func (enum APIDownloadInvoiceRequestFileType) String() string {
	if enum == "" {
		// return default value if empty
		return "pdf"
	}
	return string(enum)
}

func (enum APIDownloadInvoiceRequestFileType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *APIDownloadInvoiceRequestFileType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = APIDownloadInvoiceRequestFileType(APIDownloadInvoiceRequestFileType(tmp).String())
	return nil
}

type APIListInvoicesRequestOrderBy string

const (
	APIListInvoicesRequestOrderByInvoiceNumberDesc = APIListInvoicesRequestOrderBy("invoice_number_desc")
	APIListInvoicesRequestOrderByInvoiceNumberAsc  = APIListInvoicesRequestOrderBy("invoice_number_asc")
	APIListInvoicesRequestOrderByStartDateDesc     = APIListInvoicesRequestOrderBy("start_date_desc")
	APIListInvoicesRequestOrderByStartDateAsc      = APIListInvoicesRequestOrderBy("start_date_asc")
	APIListInvoicesRequestOrderByIssuedDateDesc    = APIListInvoicesRequestOrderBy("issued_date_desc")
	APIListInvoicesRequestOrderByIssuedDateAsc     = APIListInvoicesRequestOrderBy("issued_date_asc")
	APIListInvoicesRequestOrderByDueDateDesc       = APIListInvoicesRequestOrderBy("due_date_desc")
	APIListInvoicesRequestOrderByDueDateAsc        = APIListInvoicesRequestOrderBy("due_date_asc")
	APIListInvoicesRequestOrderByTotalUntaxedDesc  = APIListInvoicesRequestOrderBy("total_untaxed_desc")
	APIListInvoicesRequestOrderByTotalUntaxedAsc   = APIListInvoicesRequestOrderBy("total_untaxed_asc")
	APIListInvoicesRequestOrderByTotalTaxedDesc    = APIListInvoicesRequestOrderBy("total_taxed_desc")
	APIListInvoicesRequestOrderByTotalTaxedAsc     = APIListInvoicesRequestOrderBy("total_taxed_asc")
	APIListInvoicesRequestOrderByInvoiceTypeDesc   = APIListInvoicesRequestOrderBy("invoice_type_desc")
	APIListInvoicesRequestOrderByInvoiceTypeAsc    = APIListInvoicesRequestOrderBy("invoice_type_asc")
)

func (enum APIListInvoicesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "invoice_number_desc"
	}
	return string(enum)
}

func (enum APIListInvoicesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *APIListInvoicesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = APIListInvoicesRequestOrderBy(APIListInvoicesRequestOrderBy(tmp).String())
	return nil
}

type InvoiceType string

const (
	InvoiceTypeUnknownType = InvoiceType("unknown_type")
	InvoiceTypePeriodic    = InvoiceType("periodic")
	InvoiceTypePurchase    = InvoiceType("purchase")
)

func (enum InvoiceType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum InvoiceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InvoiceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InvoiceType(InvoiceType(tmp).String())
	return nil
}

// GetConsumptionResponseConsumption:
type GetConsumptionResponseConsumption struct {
	// Value: Monetary value of the consumption.
	Value *scw.Money `json:"value"`
	// Description: Description of the consumption.
	Description string `json:"description"`
	// ProjectID: Project ID of the consumption.
	ProjectID string `json:"project_id"`
	// Category: Category of the consumption.
	Category string `json:"category"`
	// OperationPath: Unique identifier of the product.
	OperationPath string `json:"operation_path"`
}

// Invoice:
type Invoice struct {
	// ID: Invoice ID.
	ID string `json:"id"`
	// StartDate: Start date of the billing period.
	StartDate *time.Time `json:"start_date"`
	// IssuedDate: Date when the invoice was sent to the customer.
	IssuedDate *time.Time `json:"issued_date"`
	// DueDate: Payment time limit, set according to the Organization's payment conditions.
	DueDate *time.Time `json:"due_date"`
	// TotalUntaxed: Total amount, untaxed.
	TotalUntaxed *scw.Money `json:"total_untaxed"`
	// TotalTaxed: Total amount, taxed.
	TotalTaxed *scw.Money `json:"total_taxed"`
	// InvoiceType: Type of invoice.
	InvoiceType InvoiceType `json:"invoice_type"`
	// Number: Invoice number.
	Number int32 `json:"number"`
}

// APIDownloadInvoiceRequest:
type APIDownloadInvoiceRequest struct {
	// InvoiceID: Invoice ID.
	InvoiceID string `json:"-"`
	// FileType: Wanted file type.
	FileType APIDownloadInvoiceRequestFileType `json:"-"`
}

// APIGetConsumptionRequest:
type APIGetConsumptionRequest struct {
	// OrganizationID: Filter by organization ID.
	OrganizationID string `json:"-"`
}

// APIListInvoicesRequest:
type APIListInvoicesRequest struct {
	// OrganizationID: Organization ID to filter for, only invoices from this Organization will be returned.
	OrganizationID *string `json:"-"`
	// StartedAfter: Invoice's `start_date` is greater or equal to `started_after`.
	StartedAfter *time.Time `json:"-"`
	// StartedBefore: Invoice's `start_date` precedes `started_before`.
	StartedBefore *time.Time `json:"-"`
	// InvoiceType: Invoice type. It can either be `periodic` or `purchase`.
	InvoiceType InvoiceType `json:"-"`
	// Page: Positive integer to choose the page to return.
	Page *int32 `json:"-"`
	// PageSize: Positive integer lower or equal to 100 to select the number of items to return.
	PageSize *uint32 `json:"-"`
	// OrderBy: How invoices are ordered in the response.
	OrderBy APIListInvoicesRequestOrderBy `json:"-"`
}

// GetConsumptionResponse:
type GetConsumptionResponse struct {
	// Consumptions: Detailed consumption list.
	Consumptions []*GetConsumptionResponseConsumption `json:"consumptions"`
	// UpdatedAt: Last consumption update date.
	UpdatedAt *time.Time `json:"updated_at"`
}

// ListInvoicesResponse:
type ListInvoicesResponse struct {
	// TotalCount: Total number of invoices.
	TotalCount uint32 `json:"total_count"`
	// Invoices: Paginated returned invoices.
	Invoices []*Invoice `json:"invoices"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListInvoicesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListInvoicesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListInvoicesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Invoices = append(r.Invoices, results.Invoices...)
	r.TotalCount += uint32(len(results.Invoices))
	return uint32(len(results.Invoices)), nil
}

// Public cloud services are based on the “pay as you go” model, which means that you only pay for what you use. Your monthly invoice is calculated at the end of each month and based on your hourly resource usage during the month.
//
// With Scaleway’s Billing API, you can manage the billing of your Scaleway cloud services.
//
// (switchcolumn)
// <Message type="note">
// You may also be interested in the [Account API](https://www.scaleway.com/en/developers/api/account/) to manage your Projects, and the [IAM API](https://www.scaleway.com/en/developers/api/iam/) to manage users, permissions and API keys in your Organization.
// </Message>
// (switchcolumn)
//
// ## Concepts and pricing
//
// Refer to our [Account concepts page](https://www.scaleway.com/en/docs/console/my-account/concepts/) to find more information on features like billing alerts, and our [dedicated pricing page](https://www.scaleway.com/en/pricing/?tags=available) for full details about the rates applied on each Scaleway product.
//
// (switchcolumn)
// (switchcolumn)
//
// ## Quickstart
//
// 1. Configure your environment variables.
//
//	<Message type="note">
//	This is an optional step that seeks to simplify your usage of the API.
//	</Message>
//
//	```bash
//	export SCW_SECRET_KEY="<API secret key>"
//	export SCW_ORGANIZATION_ID="<Scaleway Organization ID>"
//	```
//
// 2. Run the following command to obtain your consumption over the current month.
//
//	```bash
//	curl -X GET \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  -H "Content-Type: application/json" \
//	  "https://api.scaleway.com/billing/v2alpha1/consumption?organization_id=$SCW_ORGANIZATION_ID"
//	```
//
// 3. Run the following command to list your invoices.
//
//	```bash
//	curl -X GET \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  -H "Content-Type: application/json" \
//	  "https://api.scaleway.com/billing/v2alpha1/invoices"
//	```
//
//	You should get an output similar to the following one, providing details about your invoices.
//
//	<Message type="note">
//	This is a response example, the UUIDs displayed are not real.
//	</Message>
//
//	```bash
//	{
//	  "total_count": "1",
//	  "invoices": [
//	    {
//	      "id": "66f588c7-91e9-5ce9-cedb-4733f791cf73",
//	      "start_date": "2022-03-22T12:34:56.123456Z",
//	      "issued_date": "2022-03-22T12:34:56.123456Z",
//	      "due_date": "2022-03-22T12:34:56.123456Z",
//	      "total_untaxed": {
//	        "currency_code": "EUR",
//	        "units": "9",
//	        "nanos": "360000000"
//	      },
//	      "total_taxed": {
//	        "currency_code": "EUR",
//	        "units": "9",
//	        "nanos": "360000000"
//	      },
//	      "invoice_type": "periodic",
//	      "number": "0"
//	    }
//	  ]
//	}
//	```
//
// 4. Run the following command to download an invoice based on its ID.
//
//	Make sure to replace the example ID in the URL with the actual invoice ID you want to download.
//
//	```bash
//	curl -X GET \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  -H "Content-Type: application/json" \
//	  "https://api.scaleway.com/billing/v2alpha1/invoices/66f588c7-91e9-5ce9-cedb-4733f791cf73/download"
//	```
//
// (switchcolumn)
// <Message type="requirement">
// - You have an account and are logged into the [Scaleway console](https://console.scaleway.com/organization)
// - You have created an [API key](https://www.scaleway.com/en/docs/identity-and-access-management/iam/how-to/create-api-keys/) and that the API key has sufficient [IAM permissions](https://www.scaleway.com/en/docs/identity-and-access-management/iam/reference-content/permission-sets/) to perform the actions described on this page
// - You have [installed `curl`](https://curl.se/download.html)
// </Message>
// (switchcolumn)
//
// ## Technical limitations
//
// The following limitations apply to use of the Billing API:
//
// - You must have appropriate IAM permissions to manage billing. If you are the Owner of the Organization, you will automatically have these permissions. Otherwise, you will need a policy giving you the BillingManager permission set. If you were previously a Billing Administrator, you will automatically have been migrated to the BillingAdministrator group when you activated IAM, which gives you the appropriate permissions.
// - You only need BillingReadOnly permissions to query consumption.
//
// ## Going further
//
// For more help using Scaleway’s Billing API, check out the following resources:
//
// - Our [main documentation](https://www.scaleway.com/en/docs/console/my-account/)
// - Our [Slack Community](https://www.scaleway-community.slack.com/)
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

// GetConsumption: The consumption reflects the amount of money you have spent for the products you have used.
// The consumption value is monetary and is not computed in real time.
func (s *API) GetConsumption(req *APIGetConsumptionRequest, opts ...scw.RequestOption) (*GetConsumptionResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing/v2alpha1/consumption",
		Query:  query,
	}

	var resp GetConsumptionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListInvoices: List all your invoices, filtering by `start_date` and `invoice_type`. Each invoice has its own ID.
func (s *API) ListInvoices(req *APIListInvoicesRequest, opts ...scw.RequestOption) (*ListInvoicesResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "started_after", req.StartedAfter)
	parameter.AddToQuery(query, "started_before", req.StartedBefore)
	parameter.AddToQuery(query, "invoice_type", req.InvoiceType)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing/v2alpha1/invoices",
		Query:  query,
	}

	var resp ListInvoicesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DownloadInvoice: Download a specific invoice, specified by its ID.
func (s *API) DownloadInvoice(req *APIDownloadInvoiceRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "file_type", req.FileType)

	if fmt.Sprint(req.InvoiceID) == "" {
		return nil, errors.New("field InvoiceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing/v2alpha1/invoices/" + fmt.Sprint(req.InvoiceID) + "/download",
		Query:  query,
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
