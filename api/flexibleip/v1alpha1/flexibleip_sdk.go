// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package flexibleip provides methods and message types of the flexibleip v1alpha1 API.
package flexibleip

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

type FlexibleIPStatus string

const (
	FlexibleIPStatusUnknown   = FlexibleIPStatus("unknown")
	FlexibleIPStatusReady     = FlexibleIPStatus("ready")
	FlexibleIPStatusUpdating  = FlexibleIPStatus("updating")
	FlexibleIPStatusAttached  = FlexibleIPStatus("attached")
	FlexibleIPStatusError     = FlexibleIPStatus("error")
	FlexibleIPStatusDetaching = FlexibleIPStatus("detaching")
	FlexibleIPStatusLocked    = FlexibleIPStatus("locked")
)

func (enum FlexibleIPStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum FlexibleIPStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FlexibleIPStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FlexibleIPStatus(FlexibleIPStatus(tmp).String())
	return nil
}

type ListFlexibleIPsRequestOrderBy string

const (
	ListFlexibleIPsRequestOrderByCreatedAtAsc  = ListFlexibleIPsRequestOrderBy("created_at_asc")
	ListFlexibleIPsRequestOrderByCreatedAtDesc = ListFlexibleIPsRequestOrderBy("created_at_desc")
)

func (enum ListFlexibleIPsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListFlexibleIPsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListFlexibleIPsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListFlexibleIPsRequestOrderBy(ListFlexibleIPsRequestOrderBy(tmp).String())
	return nil
}

type MACAddressStatus string

const (
	MACAddressStatusUnknown  = MACAddressStatus("unknown")
	MACAddressStatusReady    = MACAddressStatus("ready")
	MACAddressStatusUpdating = MACAddressStatus("updating")
	MACAddressStatusUsed     = MACAddressStatus("used")
	MACAddressStatusError    = MACAddressStatus("error")
	MACAddressStatusDeleting = MACAddressStatus("deleting")
)

func (enum MACAddressStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum MACAddressStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MACAddressStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MACAddressStatus(MACAddressStatus(tmp).String())
	return nil
}

type MACAddressType string

const (
	MACAddressTypeUnknownType = MACAddressType("unknown_type")
	MACAddressTypeVmware      = MACAddressType("vmware")
	MACAddressTypeXen         = MACAddressType("xen")
	MACAddressTypeKvm         = MACAddressType("kvm")
)

func (enum MACAddressType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum MACAddressType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MACAddressType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MACAddressType(MACAddressType(tmp).String())
	return nil
}

// MACAddress:
type MACAddress struct {
	// ID: ID of the flexible IP.
	ID string `json:"id"`
	// MacAddress: MAC address of the Virtual MAC.
	MacAddress string `json:"mac_address"`
	// MacType: Type of virtual MAC.
	MacType MACAddressType `json:"mac_type"`
	// Status: Status of virtual MAC.
	Status MACAddressStatus `json:"status"`
	// UpdatedAt: Date on which the virtual MAC was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// CreatedAt: Date on which the virtual MAC was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Zone: MAC address IP Availability Zone.
	Zone scw.Zone `json:"zone"`
}

// FlexibleIP:
type FlexibleIP struct {
	// ID: ID of the flexible IP.
	ID string `json:"id"`
	// OrganizationID: ID of the Organization the flexible IP is attached to.
	OrganizationID string `json:"organization_id"`
	// ProjectID: ID of the Project the flexible IP is attached to.
	ProjectID string `json:"project_id"`
	// Description: Flexible IP description.
	Description string `json:"description"`
	// Tags: Flexible IP tags.
	Tags []string `json:"tags"`
	// UpdatedAt: Date on which the flexible IP was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// CreatedAt: Date on which the flexible IP was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Status: - ready : flexible IP is created and ready to be attached to a server or to be associated with a virtual MAC.
	// - updating: flexible IP is being attached to a server or a virtual MAC operation is ongoing
	// - attached: flexible IP is attached to a server
	// - error: a flexible IP operation resulted in an error
	// - detaching: flexible IP is being detached from a server
	// - locked: the resource of the flexible IP is locked.
	Status FlexibleIPStatus `json:"status"`
	// IPAddress: IP of the flexible IP.
	IPAddress scw.IPNet `json:"ip_address"`
	// MacAddress: MAC address of the flexible IP.
	MacAddress *MACAddress `json:"mac_address"`
	// ServerID: ID of the server linked to the flexible IP.
	ServerID *string `json:"server_id,omitempty"`
	// Reverse: Reverse DNS value.
	Reverse string `json:"reverse"`
	// Zone: Availability Zone of the flexible IP.
	Zone scw.Zone `json:"zone"`
}

// AttachFlexibleIPRequest:
type AttachFlexibleIPRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// FipsIDs: Multiple IDs can be provided, but note that flexible IPs must belong to the same MAC group (see details about MAC groups).
	FipsIDs []string `json:"fips_ids"`
	// ServerID: ID of the server on which to attach the flexible IPs.
	ServerID string `json:"server_id"`
}

// AttachFlexibleIPsResponse:
type AttachFlexibleIPsResponse struct {
	// TotalCount: Total count of flexible IPs that are being updated.
	TotalCount uint32 `json:"total_count"`
	// FlexibleIPs: List of flexible IPs in an updating state.
	FlexibleIPs []*FlexibleIP `json:"flexible_ips"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *AttachFlexibleIPsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *AttachFlexibleIPsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*AttachFlexibleIPsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.FlexibleIPs = append(r.FlexibleIPs, results.FlexibleIPs...)
	r.TotalCount += uint32(len(results.FlexibleIPs))
	return uint32(len(results.FlexibleIPs)), nil
}

// CreateFlexibleIPRequest:
type CreateFlexibleIPRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ProjectID: ID of the project to associate with the Flexible IP.
	ProjectID string `json:"project_id"`
	// Description: Flexible IP description (max. of 255 characters).
	Description string `json:"description"`
	// Tags: Tags to associate to the flexible IP.
	Tags []string `json:"tags"`
	// ServerID: ID of the server to which the newly created flexible IP will be attached.
	ServerID *string `json:"server_id,omitempty"`
	// Reverse: Value of the reverse DNS.
	Reverse *string `json:"reverse,omitempty"`
	// IsIPv6: Defines whether the flexible IP has an IPv6 address.
	IsIPv6 bool `json:"is_ipv6"`
}

// DeleteFlexibleIPRequest:
type DeleteFlexibleIPRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// FipID: ID of the flexible IP to delete.
	FipID string `json:"-"`
}

// DeleteMACAddrRequest:
type DeleteMACAddrRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// FipID: If the flexible IP belongs to a MAC group, the MAC will be removed from both the MAC group and flexible IP.
	FipID string `json:"-"`
}

// DetachFlexibleIPRequest:
type DetachFlexibleIPRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// FipsIDs: List of flexible IP IDs to detach from a server. Multiple IDs can be provided. Note that flexible IPs must belong to the same MAC group.
	FipsIDs []string `json:"fips_ids"`
}

// DetachFlexibleIPsResponse:
type DetachFlexibleIPsResponse struct {
	// TotalCount: Total count of flexible IPs that are being detached.
	TotalCount uint32 `json:"total_count"`
	// FlexibleIPs: List of flexible IPs in a detaching state.
	FlexibleIPs []*FlexibleIP `json:"flexible_ips"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *DetachFlexibleIPsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *DetachFlexibleIPsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*DetachFlexibleIPsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.FlexibleIPs = append(r.FlexibleIPs, results.FlexibleIPs...)
	r.TotalCount += uint32(len(results.FlexibleIPs))
	return uint32(len(results.FlexibleIPs)), nil
}

// DuplicateMACAddrRequest:
type DuplicateMACAddrRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// FipID: Note that the flexible IPs need to be attached to the same server.
	FipID string `json:"-"`
	// DuplicateFromFipID: Note that flexible IPs need to be attached to the same server.
	DuplicateFromFipID string `json:"duplicate_from_fip_id"`
}

// GenerateMACAddrRequest:
type GenerateMACAddrRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// FipID: ID of the flexible IP for which to generate a virtual MAC.
	FipID string `json:"-"`
	// MacType: TODO.
	MacType MACAddressType `json:"mac_type"`
}

// GetFlexibleIPRequest:
type GetFlexibleIPRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// FipID: ID of the flexible IP.
	FipID string `json:"-"`
}

// ListFlexibleIPsRequest:
type ListFlexibleIPsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// OrderBy: Sort order of the returned flexible IPs.
	OrderBy ListFlexibleIPsRequestOrderBy `json:"order_by"`
	// Page: Page number.
	Page *int32 `json:"page,omitempty"`
	// PageSize: Maximum number of flexible IPs per page.
	PageSize *uint32 `json:"page_size,omitempty"`
	// Tags: Filter by tag, only flexible IPs with one or more matching tags will be returned.
	Tags []string `json:"tags"`
	// Status: Filter by status, only flexible IPs with this status will be returned.
	Status []FlexibleIPStatus `json:"status"`
	// ServerIDs: Filter by server IDs, only flexible IPs with these server IDs will be returned.
	ServerIDs []string `json:"server_ids"`
	// OrganizationID: Filter by Organization ID, only flexible IPs from this Organization will be returned.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: Filter by Project ID, only flexible IPs from this Project will be returned.
	ProjectID *string `json:"project_id,omitempty"`
}

// ListFlexibleIPsResponse:
type ListFlexibleIPsResponse struct {
	// TotalCount: Total count of matching flexible IPs.
	TotalCount uint32 `json:"total_count"`
	// FlexibleIPs: List of all flexible IPs.
	FlexibleIPs []*FlexibleIP `json:"flexible_ips"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListFlexibleIPsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListFlexibleIPsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListFlexibleIPsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.FlexibleIPs = append(r.FlexibleIPs, results.FlexibleIPs...)
	r.TotalCount += uint32(len(results.FlexibleIPs))
	return uint32(len(results.FlexibleIPs)), nil
}

// MoveMACAddrRequest:
type MoveMACAddrRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// FipID:
	FipID string `json:"-"`
	// DstFipID:
	DstFipID string `json:"dst_fip_id"`
}

// UpdateFlexibleIPRequest:
type UpdateFlexibleIPRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// FipID: ID of the flexible IP to update.
	FipID string `json:"-"`
	// Description: Flexible IP description (max. 255 characters).
	Description *string `json:"description,omitempty"`
	// Tags: Tags associated with the flexible IP.
	Tags *[]string `json:"tags,omitempty"`
	// Reverse: Value of the reverse DNS.
	Reverse *string `json:"reverse,omitempty"`
}

// Flexible IP addresses are additional public IP addresses that you can hold independently of any Elastic Metal server. They can either be IPv4 (single IP) or IPv6 (/64 IP block).
//
// Flexible IPs can be attached to and detached from any Elastic Metal server within the same zone. You can hold multiple flexible IPs in your account, and a given server can be linked to multiple flexible IPs. When you delete a flexible IP address, it is disassociated from your account.
//
// Flexible IPs can also be used to implement failovers. If any failure or maintenance issue occurs on a given Elastic Metal server, its flexible IP address can be transferred to another server.
//
// (switchcolumn)
// <Message type="important">
//
//	This documentation refers to flexible IPs for Elastic Metal servers. Refer to the corresponding product documentation if you are looking for information about flexible IPs for other products.
//
// </Message>
// (switchcolumn)
//
// ## Concepts
//
// Refer to our [dedicated concepts page](https://www.scaleway.com/en/docs/compute/elastic-metal/concepts/) to find definitions of the different terms referring to Elastic Metal and flexible IPs.
//
// ## Quickstart
//
// (switchcolumn)
// (switchcolumn)
//
//  1. **Configure your environment variables.**
//     ```bash
//     export SCW_ACCESS_KEY="<API access key>"
//     export SCW_SECRET_KEY="<API secret key>"
//     export SCW_DEFAULT_ZONE="<Scaleway default zone>"
//     ```
//     <Message type="note">
//     This is an optional step that seeks to simplify your usage of the APIs. Refer to the Availability Zones section to verify which zones are available for use.
//     </Message>
//
//  2. **Edit the POST request payload** that we will use in the next step to create a flexible IP.
//     ```
//     {
//     "project_id": "88f30nda-6768-9293-a89c-2b0b178628a6",
//     "description": "This is the description of my fIP",
//     "tags": [
//     "tag1"
//     ],
//     "server_id": "9dddd3sa-f13c-4351-9185-18f6b6d97t9w",
//     "reverse": "9dddd3se-f14c-4859-9185-18f6b6d78f8b.fr-par-1.baremetal.scw.cloud",
//     "is_ipv6": true
//     }
//     ```
//
//     | Parameter      | Description                                                        |
//     |----------------|--------------------------------------------------------------------|
//     | `project_id`   | **REQUIRED** ID of the Project to create your flexible IP in.      |
//     | `description`  | A description for your flexible IP (max. 255 characters).          |
//     | `tags`         | One or several tags for your flexible IP (optional)                |
//     | `server_id`    | ID of the server on which to attach your newly created flexible IP.|
//     | `reverse`      | Value of the server's reverse DNS.                                 |
//     | `is_ipv6`      | **BOOLEAN** Defines whether the flexible IP has an IPv6 address.   |
//
//     <Message type="note">
//     Except when specified, all values are nullable.
//     </Message>
//
//  3. **Create a flexible IP**: run the following command to create a flexible IP, including the payload you edited in the previous step.
//     ```bash
//     curl -X POST \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     -H "Content-Type: application/json" \
//     "https://api.scaleway.com/flexible-ip/v1alpha1/zones/$SCW_DEFAULT_ZONE/fips" \
//     -d '{
//     "project_id": "88f30nda-6768-9293-a89c-2b0b178628a6",
//     "description": "This is the description of my fIP",
//     "tags": [
//     "tag1"
//     ],
//     "server_id": "9dddd3sa-f13c-4351-9185-18f6b6d97t9w",
//     "reverse": "9dddd3se-f14c-4859-9185-18f6b6d78f8b.fr-par-1.baremetal.scw.cloud",
//     "is_ipv6": true
//     }'
//     ```
//
//     You should get an output similar to the following one, providing details about your flexible IP:
//
//     <Message type="note">
//     This is a response example, the UUIDs and IP address displayed are not real.
//     </Message>
//
//     ```bash
//     {
//     "id": "058d9f12-c33d-523d-b216-da4c9d0f3d66",
//     "organization_id": "88f30nda-6768-9293-a89c-2b0b178628a6",
//     "project_id": "88f30nda-6768-9293-a89c-2b0b178628a6",
//     "description": "This is the description of my fIP",
//     "updated_at": "2023-04-04T13:34:19.058178830Z",
//     "created_at": "2023-01-31T16:30:54.017824Z",
//     "status": "updating",
//     "tags": [
//     "tag1"
//     ],
//     "ip_address": "1998:cb9:813:24f3::/75",
//     "server_id": null,
//     "reverse": "9dddd3se-f14c-4859-9185-18f6b6d78f8b.fr-par-1.baremetal.scw.cloud",
//     "mac_address": null,
//     "zone": "fr-par-1"
//     }
//     ```
//
//  4. **Get a list of your flexible IPs**: run the following command to get a list of all the flexible IPs in your account, with their details:
//     ```bash
//     curl -X GET \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     -H "Content-Type: application/json" \
//     "https://api.scaleway.com/flexible-ip/v1alpha1/zones/$SCW_DEFAULT_ZONE/fips"
//     ```
//
//  5. **Generate a virtual MAC (Media Access Control) address**: run the following command to generate a virtual MAC address on a given flexible IP. Ensure that you replace `<FLEXIBLE-IP-ID>` in the URL with the ID of the flexible IP you want to create a virtual MAC address for.
//     ```bash
//     curl -X POST \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     -H "Content-Type: application/json" \
//     "https://api.scaleway.com/flexible-ip/v1alpha1/zones/$SCW_DEFAULT_ZONE/fips/<FLEXIBLE-IP-ID>/mac" \
//     -d '{
//     "mac_type": "<MAC_TYPE>"
//     }'
//     ```
//
//     **Payload value**
//
//     * **mac_type** (string): Choose the type of virtual MAC address you want to generate on your flexible IP: `vmware`, `xen` or `kvm` (with the default value being set to `unknown_type`). To get more information about the available virtual MAC addresses, refer to the "Technical information" part of this quickstart.
//
//  6. **Duplicate a virtual MAC (Media Access Control) address**: run the following command to duplicate a Virtual MAC from a given flexible IP onto another flexible IP attached to the same server.
//     ```bash
//     curl -X POST \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     -H "Content-Type: application/json" \
//     "https://api.scaleway.com/flexible-ip/v1alpha1/zones/$SCW_DEFAULT_ZONE/fips/<fip_id> \
//     -d '{
//     "duplicate_from_fip_id": "<ID_OF_THE_FIP_TO_DUPLICATE_MAC_FROM>"
//     }'
//     ```
//     **Payload values**
//
//     * **fip_id** (string): ID of the flexible IP on which to duplicate the Virtual MAC. Note that flexible IPs need to be attached to the same server for the operation to work.
//     * **duplicate_from_fip_id** (string): ID of the flexible IP to duplicate the Virtual MAC from.
//
// (switchcolumn)
// <Message type="requirement">
// To perform the following steps, you must first ensure that:
// - You have a [Scaleway account](https://console.scaleway.com/)
// - You have [created an API key](https://www.scaleway.com/en/docs/identity-and-access-management/iam/how-to/create-api-keys/) and that the API key has [sufficient IAM permissions](https://www.scaleway.com/en/docs/identity-and-access-management/iam/reference-content/permission-sets/) to perform the actions described on this page
// - You have [installed `curl`](https://curl.se/download.html)
// </Message>
// (switchcolumn)
//
// ## Technical information
//
// A virtual MAC (Media Access Control) address can be generated on a flexible IP. Virtual MAC addresses are unique identifiers assigned to a virtual machine's network interface. This is particularly useful for virtualization technologies enabling multiple virtual machines to run on a single host machine (now called a hypervisor).
//
// Flexible IPs can have virtual MAC addresses assigned to them, and it is possible to duplicate these virtual MAC addresses between flexible IPs on the same server. When a virtual MAC address is duplicated onto another flexible IP, the two become part of the same virtual MAC group.
//
// When flexible IPs belong to a given MAC group, they cannot be moved separately to another server. Both must be transferred to the new server, as a group. Subsequently, a MAC group can be moved by providing a list of flexible IP IDs in Attach/Detach requests.
//
// Note that if you detach a single flexible IP from a MAC group, the virtual MAC address will be removed from the detached flexible IP.
//
// ### Availability Zones
//
// Flexible IPs are available in the following Availability Zones:
//
// | Name      | API ID                           |
// |-----------|--------------------------------- |
// | Paris     | `fr-par-1` `fr-par-2`            |
// | Amsterdam | `nl-ams-1`                       |
//
// ## Technical limitations
//
// - Flexible IPs exist for many resources (Instances, Load Balancers, etc). Note, however, that each of these sets of flexible IPs is independent and usable only with that product. This API concerns flexible IPs for Elastic Metal servers only.
// - There is a limit of 64 flexible IPs per server
// ## Going further
//
// For more help using flexible IPs, check out the following resources:
// - Our [main documentation](https://www.scaleway.com/en/docs/compute/elastic-metal/)
// - Our [Slack Community](https://www.scaleway.com/en/docs/tutorials/scaleway-slack-community/)
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
func (s *API) Zones() []scw.Zone {
	return []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1}
}

// CreateFlexibleIP: Generate a new flexible IP within a given zone, specifying its configuration including Project ID and description.
func (s *API) CreateFlexibleIP(req *CreateFlexibleIPRequest, opts ...scw.RequestOption) (*FlexibleIP, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp FlexibleIP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFlexibleIP: Retrieve information about an existing flexible IP, specified by its ID and zone. Its full details, including Project ID, description and status, are returned in the response object.
func (s *API) GetFlexibleIP(req *GetFlexibleIPRequest, opts ...scw.RequestOption) (*FlexibleIP, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.FipID) == "" {
		return nil, errors.New("field FipID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "",
	}

	var resp FlexibleIP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListFlexibleIPs: List all flexible IPs within a given zone.
func (s *API) ListFlexibleIPs(req *ListFlexibleIPsRequest, opts ...scw.RequestOption) (*ListFlexibleIPsResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}
	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "server_ids", req.ServerIDs)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips",
		Query:  query,
	}

	var resp ListFlexibleIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateFlexibleIP: Update the parameters of an existing flexible IP, specified by its ID and zone. These parameters include tags and description.
func (s *API) UpdateFlexibleIP(req *UpdateFlexibleIPRequest, opts ...scw.RequestOption) (*FlexibleIP, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.FipID) == "" {
		return nil, errors.New("field FipID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp FlexibleIP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteFlexibleIP: Delete an existing flexible IP, specified by its ID and zone. Note that deleting a flexible IP is permanent and cannot be undone.
func (s *API) DeleteFlexibleIP(req *DeleteFlexibleIPRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.FipID) == "" {
		return errors.New("field FipID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// AttachFlexibleIP: Attach an existing flexible IP to a specified Elastic Metal server.
func (s *API) AttachFlexibleIP(req *AttachFlexibleIPRequest, opts ...scw.RequestOption) (*AttachFlexibleIPsResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/attach",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AttachFlexibleIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachFlexibleIP: Detach an existing flexible IP from a specified Elastic Metal server.
func (s *API) DetachFlexibleIP(req *DetachFlexibleIPRequest, opts ...scw.RequestOption) (*DetachFlexibleIPsResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/detach",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DetachFlexibleIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GenerateMACAddr: Generate a virtual MAC (Media Access Control) address on an existing flexible IP.
func (s *API) GenerateMACAddr(req *GenerateMACAddrRequest, opts ...scw.RequestOption) (*FlexibleIP, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.FipID) == "" {
		return nil, errors.New("field FipID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "/mac",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp FlexibleIP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DuplicateMACAddr: Duplicate a virtual MAC address from a given flexible IP to another flexible IP attached to the same server.
func (s *API) DuplicateMACAddr(req *DuplicateMACAddrRequest, opts ...scw.RequestOption) (*FlexibleIP, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.FipID) == "" {
		return nil, errors.New("field FipID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "/mac/duplicate",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp FlexibleIP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MoveMACAddr: Relocate a virtual MAC (Media Access Control) address from an existing flexible IP to a different flexible IP.
func (s *API) MoveMACAddr(req *MoveMACAddrRequest, opts ...scw.RequestOption) (*FlexibleIP, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.FipID) == "" {
		return nil, errors.New("field FipID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "/mac/move",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp FlexibleIP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteMACAddr: Detach a given MAC (Media Access Control) address from an existing flexible IP.
func (s *API) DeleteMACAddr(req *DeleteMACAddrRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.FipID) == "" {
		return errors.New("field FipID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "/mac",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}
