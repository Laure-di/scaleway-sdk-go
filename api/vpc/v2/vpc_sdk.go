// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package vpc provides methods and message types of the vpc v2 API.
package vpc

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

type ListPrivateNetworksRequestOrderBy string

const (
	ListPrivateNetworksRequestOrderByCreatedAtAsc  = ListPrivateNetworksRequestOrderBy("created_at_asc")
	ListPrivateNetworksRequestOrderByCreatedAtDesc = ListPrivateNetworksRequestOrderBy("created_at_desc")
	ListPrivateNetworksRequestOrderByNameAsc       = ListPrivateNetworksRequestOrderBy("name_asc")
	ListPrivateNetworksRequestOrderByNameDesc      = ListPrivateNetworksRequestOrderBy("name_desc")
)

func (enum ListPrivateNetworksRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListPrivateNetworksRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPrivateNetworksRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPrivateNetworksRequestOrderBy(ListPrivateNetworksRequestOrderBy(tmp).String())
	return nil
}

type ListVPCsRequestOrderBy string

const (
	ListVPCsRequestOrderByCreatedAtAsc  = ListVPCsRequestOrderBy("created_at_asc")
	ListVPCsRequestOrderByCreatedAtDesc = ListVPCsRequestOrderBy("created_at_desc")
	ListVPCsRequestOrderByNameAsc       = ListVPCsRequestOrderBy("name_asc")
	ListVPCsRequestOrderByNameDesc      = ListVPCsRequestOrderBy("name_desc")
)

func (enum ListVPCsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListVPCsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListVPCsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListVPCsRequestOrderBy(ListVPCsRequestOrderBy(tmp).String())
	return nil
}

// Subnet:
type Subnet struct {
	// ID: ID of the subnet.
	ID string `json:"id"`
	// CreatedAt: Subnet creation date.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Subnet last modification date.
	UpdatedAt *time.Time `json:"updated_at"`
	// Subnet: Subnet CIDR.
	Subnet scw.IPNet `json:"subnet"`
}

// PrivateNetwork:
type PrivateNetwork struct {
	// ID: Private Network ID.
	ID string `json:"id"`
	// Name: Private Network name.
	Name string `json:"name"`
	// OrganizationID: Scaleway Organization the Private Network belongs to.
	OrganizationID string `json:"organization_id"`
	// ProjectID: Scaleway Project the Private Network belongs to.
	ProjectID string `json:"project_id"`
	// Region: Region in which the Private Network is available.
	Region scw.Region `json:"region"`
	// Tags: Tags of the Private Network.
	Tags []string `json:"tags"`
	// CreatedAt: Date the Private Network was created.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Date the Private Network was last modified.
	UpdatedAt *time.Time `json:"updated_at"`
	// Subnets: Private Network subnets.
	Subnets []*Subnet `json:"subnets"`
	// VpcID: VPC the Private Network belongs to.
	VpcID string `json:"vpc_id"`
	// DHCPEnabled: Defines whether managed DHCP is enabled for this Private Network.
	DHCPEnabled bool `json:"dhcp_enabled"`
}

// VPC:
type VPC struct {
	// ID: VPC ID.
	ID string `json:"id"`
	// Name: VPC name.
	Name string `json:"name"`
	// OrganizationID: Scaleway Organization the VPC belongs to.
	OrganizationID string `json:"organization_id"`
	// ProjectID: Scaleway Project the VPC belongs to.
	ProjectID string `json:"project_id"`
	// Region: Region of the VPC.
	Region scw.Region `json:"region"`
	// Tags: Tags for the VPC.
	Tags []string `json:"tags"`
	// IsDefault: Defines whether the VPC is the default one for its Project.
	IsDefault bool `json:"is_default"`
	// CreatedAt: Date the VPC was created.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Date the VPC was last modified.
	UpdatedAt *time.Time `json:"updated_at"`
	// PrivateNetworkCount: Number of Private Networks within this VPC.
	PrivateNetworkCount uint32 `json:"private_network_count"`
}

// AddSubnetsRequest:
type AddSubnetsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// PrivateNetworkID: Private Network ID.
	PrivateNetworkID string `json:"-"`
	// Subnets: Private Network subnets CIDR.
	Subnets []scw.IPNet `json:"subnets"`
}

// AddSubnetsResponse:
type AddSubnetsResponse struct {
	// Subnets:
	Subnets []scw.IPNet `json:"subnets"`
}

// CreatePrivateNetworkRequest:
type CreatePrivateNetworkRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Name: Name for the Private Network.
	Name string `json:"name"`
	// ProjectID: Scaleway Project in which to create the Private Network.
	ProjectID string `json:"project_id"`
	// Tags: Tags for the Private Network.
	Tags []string `json:"tags"`
	// Subnets: Private Network subnets CIDR.
	Subnets []scw.IPNet `json:"subnets"`
	// VpcID: VPC in which to create the Private Network.
	VpcID *string `json:"vpc_id,omitempty"`
}

// CreateVPCRequest:
type CreateVPCRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Name: Name for the VPC.
	Name string `json:"name"`
	// ProjectID: Scaleway Project in which to create the VPC.
	ProjectID string `json:"project_id"`
	// Tags: Tags for the VPC.
	Tags []string `json:"tags"`
}

// DeletePrivateNetworkRequest:
type DeletePrivateNetworkRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// PrivateNetworkID: Private Network ID.
	PrivateNetworkID string `json:"-"`
}

// DeleteSubnetsRequest:
type DeleteSubnetsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// PrivateNetworkID: Private Network ID.
	PrivateNetworkID string `json:"-"`
	// Subnets: Private Network subnets CIDR.
	Subnets []scw.IPNet `json:"subnets"`
}

// DeleteSubnetsResponse:
type DeleteSubnetsResponse struct {
	// Subnets:
	Subnets []scw.IPNet `json:"subnets"`
}

// DeleteVPCRequest:
type DeleteVPCRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// VpcID: VPC ID.
	VpcID string `json:"-"`
}

// EnableDHCPRequest:
type EnableDHCPRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// PrivateNetworkID: Private Network ID.
	PrivateNetworkID string `json:"-"`
}

// GetPrivateNetworkRequest:
type GetPrivateNetworkRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// PrivateNetworkID: Private Network ID.
	PrivateNetworkID string `json:"-"`
}

// GetVPCRequest:
type GetVPCRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// VpcID: VPC ID.
	VpcID string `json:"-"`
}

// ListPrivateNetworksRequest:
type ListPrivateNetworksRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// OrderBy: Sort order of the returned Private Networks.
	OrderBy ListPrivateNetworksRequestOrderBy `json:"-"`
	// Page: Page number to return, from the paginated results.
	Page *int32 `json:"-"`
	// PageSize: Maximum number of Private Networks to return per page.
	PageSize *uint32 `json:"-"`
	// Name: Name to filter for. Only Private Networks with names containing this string will be returned.
	Name *string `json:"-"`
	// Tags: Tags to filter for. Only Private Networks with one or more matching tags will be returned.
	Tags []string `json:"-"`
	// OrganizationID: Organization ID to filter for. Only Private Networks belonging to this Organization will be returned.
	OrganizationID *string `json:"-"`
	// ProjectID: Project ID to filter for. Only Private Networks belonging to this Project will be returned.
	ProjectID *string `json:"-"`
	// PrivateNetworkIDs: Private Network IDs to filter for. Only Private Networks with one of these IDs will be returned.
	PrivateNetworkIDs []string `json:"-"`
	// VpcID: VPC ID to filter for. Only Private Networks belonging to this VPC will be returned.
	VpcID *string `json:"-"`
	// DHCPEnabled: DHCP status to filter for. When true, only Private Networks with managed DHCP enabled will be returned.
	DHCPEnabled *bool `json:"-"`
}

// ListPrivateNetworksResponse:
type ListPrivateNetworksResponse struct {
	// PrivateNetworks:
	PrivateNetworks []*PrivateNetwork `json:"private_networks"`
	// TotalCount:
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPrivateNetworksResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPrivateNetworksResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListPrivateNetworksResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PrivateNetworks = append(r.PrivateNetworks, results.PrivateNetworks...)
	r.TotalCount += uint32(len(results.PrivateNetworks))
	return uint32(len(results.PrivateNetworks)), nil
}

// ListVPCsRequest:
type ListVPCsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// OrderBy: Sort order of the returned VPCs.
	OrderBy ListVPCsRequestOrderBy `json:"-"`
	// Page: Page number to return, from the paginated results.
	Page *int32 `json:"-"`
	// PageSize: Maximum number of VPCs to return per page.
	PageSize *uint32 `json:"-"`
	// Name: Name to filter for. Only VPCs with names containing this string will be returned.
	Name *string `json:"-"`
	// Tags: Tags to filter for. Only VPCs with one more more matching tags will be returned.
	Tags []string `json:"-"`
	// OrganizationID: Organization ID to filter for. Only VPCs belonging to this Organization will be returned.
	OrganizationID *string `json:"-"`
	// ProjectID: Project ID to filter for. Only VPCs belonging to this Project will be returned.
	ProjectID *string `json:"-"`
	// IsDefault: Defines whether to filter only for VPCs which are the default one for their Project.
	IsDefault *bool `json:"-"`
}

// ListVPCsResponse:
type ListVPCsResponse struct {
	// Vpcs:
	Vpcs []*VPC `json:"vpcs"`
	// TotalCount:
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListVPCsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListVPCsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListVPCsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Vpcs = append(r.Vpcs, results.Vpcs...)
	r.TotalCount += uint32(len(results.Vpcs))
	return uint32(len(results.Vpcs)), nil
}

// MigrateZonalPrivateNetworksRequest:
type MigrateZonalPrivateNetworksRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// OrganizationID: Organization ID to target. The specified zoned Private Networks within this Organization will be migrated to regional.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: Project to target. The specified zoned Private Networks within this Project will be migrated to regional.
	ProjectID *string `json:"project_id,omitempty"`
	// PrivateNetworkIDs: IDs of the Private Networks to migrate.
	PrivateNetworkIDs []string `json:"private_network_ids"`
}

// SetSubnetsRequest:
type SetSubnetsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// PrivateNetworkID: Private Network ID.
	PrivateNetworkID string `json:"-"`
	// Subnets: Private Network subnets CIDR.
	Subnets []scw.IPNet `json:"subnets"`
}

// SetSubnetsResponse:
type SetSubnetsResponse struct {
	// Subnets:
	Subnets []scw.IPNet `json:"subnets"`
}

// UpdatePrivateNetworkRequest:
type UpdatePrivateNetworkRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// PrivateNetworkID: Private Network ID.
	PrivateNetworkID string `json:"-"`
	// Name: Name for the Private Network.
	Name *string `json:"name,omitempty"`
	// Tags: Tags for the Private Network.
	Tags *[]string `json:"tags,omitempty"`
}

// UpdateVPCRequest:
type UpdateVPCRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// VpcID: VPC ID.
	VpcID string `json:"-"`
	// Name: Name for the VPC.
	Name *string `json:"name,omitempty"`
	// Tags: Tags for the VPC.
	Tags *[]string `json:"tags,omitempty"`
}

// VPC allows you to build your own **V**irtual **P**rivate **C**loud on top of Scaleway’s shared public cloud.
//
// One default VPC for every available region is automatically created in each Scaleway Project.
//
// VPC currently comprises the regional Private Networks product. Layer 2 regional Private Networks sit inside the layer 3 VPC. Private Networks allows Scaleway resources (Instances, Load Balancers, Managed Databases etc.) within a single region to be interconnected through a dedicated, private, and flexible [L2 network](https://en.wikipedia.org/wiki/Data_link_layer).
//
// You can add as many resources to your networks as you want, and add up to eight (8) different networks
// per resource. This allows you to run services isolated from the public internet and expose them to the rest of your infrastructure without worrying about public network filtering.
//
// (switchcolumn)
// <Message type="note">
// VPC v2 is now in **General Availability**.
// </Message>
// (switchcolumn)
//
// ## Concepts
//
// Refer to our [dedicated concepts page](https://www.scaleway.com/en/docs/network/vpc/concepts/) to find definitions of all concepts and terminology related to VPC.
//
// (switchcolumn)
// (switchcolumn)
//
// ## Quickstart
//
// 1. **Configure your environment variables**
//
//	<Message type="note">
//	This is an optional step that seeks to simplify your usage of the API. See the [Technical information](#technical-information) section below for help choosing an Availability Zone and Region. You can find your Project ID in the [Scaleway console](https://console.scaleway.com/project/settings).
//	</Message>
//
//	```bash
//	export SCW_SECRET_KEY="<API secret key>"
//	export SCW_DEFAULT_REGION="<Scaleway region>"
//	export SCW_DEFAULT_ZONE="<Scaleway Availability Zone>"
//	export SCW_PROJECT_ID="<Scaleway Project ID>"
//	```
//
// 2. **Create a Private Network**. Run the following command to create a Private Network. You can customize the details in the payload (name, tags etc.) to your needs.
//
//	```bash
//	curl -X POST \
//	    -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	    -H "Content-Type: application/json" \
//	    "https://api.scaleway.com/vpc/v2/regions/$SCW_DEFAULT_REGION/private-networks" \
//	    -d '{
//	        "name": "My new Private Network",
//	        "project_id": "'"$SCW_PROJECT_ID"'",
//	        "tags": ["test", "dev"]
//	    }'
//	```
//
//	<Message type="tip">
//	Keep the `id` field of the response: it is your Private Network ID, and is useable across all Scaleway products that support Private Networks. It may be useful to you to export the Private Network ID as a new environment variable `export PN_ID="<Your Private Network ID>`
//	</Message>
//
// 3. **Attach a resource to your Private Network**. Each Scaleway product has its own API to interact with Private Networks. To attach an Instance, Managed Database, Elastic Metal server, Load Balancer or Public Gateway to your Private Network, see instructions in the documentation of the relevant product API. Here, we take the example of an Instance.
//
//	Use the following call to attach an Instance to your Private Network. Ensure you replace `<Instance ID>` with the ID of your Instance, and `<Private Network ID>` with the ID of your Private Network. Note that the Instance must be in an Availability Zone that is part of the region of your Private Network.
//
//	```bash
//	curl -X POST \
//	    -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	    -H "Content-Type: application/json" \
//	    "https://api.scaleway.com/instance/v1/zones/$SCW_DEFAULT_ZONE/servers/<Instance ID>/private_nics" \
//	    -d '{"private_network_id": "<Private Network ID>"}'
//	```
//
//	<Message type="tip">
//	Keep the `id` field of the response: it is your Private NIC ID. It may be useful to you to export
//	the Private NIC ID as a new environment variable `export NIC_ID="<Your Private NIC ID>`.
//	</Message>
//
//	<Message type="tip">
//	Keep the `mac_address` field of the response, as it will allow you to identify the Private NIC inside your Instance. If successful, a new network interface will appear inside your Instance, ready to be configured to transmit traffic to other Instances of the same network, with the MAC address returned by the API call.
//	</Message>
//
// 4. **Confirm that the network interface has been plugged in**. To do this, connect to your Instance and run `dmseg`. You should see an output similar to the following:
//
//	```bash
//	[1579004.592869] pci 0000:00:05.0: [1af4:1000] type 00 class 0x020000
//	[1579004.594835] pci 0000:00:05.0: reg 0x10: [io  0x0000-0x003f]
//	[1579004.596715] pci 0000:00:05.0: reg 0x14: [mem 0x00000000-0x00000fff]
//	[1579004.598732] pci 0000:00:05.0: reg 0x20: [mem 0x00000000-0x00003fff 64bit pref]
//	[1579004.600765] pci 0000:00:05.0: reg 0x30: [mem 0x00000000-0x0007ffff pref]
//	[1579004.603819] pci 0000:00:05.0: BAR 6: assigned [mem 0xc0100000-0xc017ffff pref]
//	[1579004.604582] pci 0000:00:05.0: BAR 4: assigned [mem 0x100000c000-0x100000ffff 64bit pref]
//	[1579004.605555] pci 0000:00:05.0: BAR 1: assigned [mem 0xc0003000-0xc0003fff]
//	[1579004.606383] pci 0000:00:05.0: BAR 0: assigned [io  0x1000-0x103f]
//	[1579004.607212] virtio-pci 0000:00:05.0: enabling device (0000 -> 0003)
//	[1579004.625149] PCI Interrupt Link [LNKA] enabled at IRQ 11
//	[1579004.644930] virtio_net virtio3 ens5: renamed from eth0
//	```
//
// 5. **Confirm the presence of the network interface, and confirm its name if several networks are plugged into your Instance**. To do this, run `ip -br link`. You should see an output similar to the following:
//
//	```bash
//	lo               UNKNOWN        00:00:00:00:00:00 <LOOPBACK,UP,LOWER_UP>
//	ens2             UP             de:1c:94:44:d0:04 <BROADCAST,MULTICAST,UP,LOWER_UP>
//	ens5             DOWN           02:00:00:00:00:31 <BROADCAST,MULTICAST>
//	ens6             DOWN           02:00:00:00:01:5b <BROADCAST,MULTICAST>
//	ens7             DOWN           02:00:00:00:01:5e <BROADCAST,MULTICAST>
//	```
//
// 6. **Configure the Instance's IP address**. DHCP is activated by default on new Private Networks, and automatically assigns IP addresses to resources on the network. If you have an older Private Network, [check whether DHCP is activated](https://www.scaleway.com/en/docs/network/vpc/reference-content/vpc-migration/) and either activate DHCP for automatic IP configuration, or [manually configure](https://www.scaleway.com/en/docs/compute/instances/reference-content/manual-configuration-private-ips/) the network interface on your Instance if necessary.
//
// 7. **Delete your Private NIC**, which equates to unplugging your Instance from the Private Network. Use the following call. Ensure you replace `<Instance ID>` with the ID of your Instance, `<Private Network ID>` with the ID of your Private Network, and `<NIC ID>` with the ID of your Private NIC.
//
//	```bash
//	curl -X DELETE \
//	    -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	    -H "Content-Type: application/json" \
//	    "https://api.scaleway.com/instance/v1/zones/$SCW_DEFAULT_ZONE/servers/<Instance ID>/private_nics/<NIC ID>"
//	```
//
//	The network interface disappears from your Instance.
//
// 8. **Delete your Private Network**. Use the following call. Ensure you replace `<Private Network ID>` with the ID of your Private Network.
//
//	 ```bash
//	 curl -X DELETE \
//	     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	     -H "Content-Type: application/json" \
//	     "https://api.scaleway.com/vpc/v2/regions/$SCW_DEFAULT_REGION/private-networks/<Private Network ID>"
//	 ```
//
//	 <Message type="note">
//	 Private Networks must be empty to be deleted. Ensure you have detached all resources from your network prior to deletion.
//	 </Message>
//
//	(switchcolumn)
//	<Message type="requirement">
//	 - You have a [Scaleway account](https://console.scaleway.com/)
//	 - You have created an [API key](https://www.scaleway.com/en/docs/identity-and-access-management/iam/how-to/create-api-keys/) and that the API key has sufficient [IAM permissions](https://www.scaleway.com/en/docs/identity-and-access-management/iam/reference-content/permission-sets/) to perform the actions described on this page
//	 - You have [installed `curl`](https://curl.se/download.html)
//	 </Message>
//	 (switchcolumn)
//
// ## Technical information
//
// VPC and Private Networks are available in the Paris, Amsterdam and Warsaw regions, which are represented by the following path parameters:
//
// * `fr-par`
// * `nl-ams`
// * `pl-waw`
//
// ## Technical limitations
//
// The following limitations apply to Scaleway VPC:
//
// - The only product currently available within the regional VPC is Private Networks.
// - Up to 250 resources can be attached to a Private Network.
// - A resource can be attached to up to 8 Private Networks.
// - The following resource types can be attached to a Private Network:
//   - Instances
//   - Elastic Metal servers
//   - Load Balancers
//   - Public Gateways
//   - Managed Databases for PostgreSQL and MySQL
//   - Managed Databases for Redis (only at the time of resource creation)
//   - Kubernetes Kapsule (only at the time of resource creation)
//
// - The MAC address of an Instance in a Private Network cannot be changed.
// - Broadcast and multicast traffic, while supported, are heavily rate-limited.
//
// ## Going further
//
// For more help using Scaleway VPC and Private Networks, check out the following resources:
// - Our [main documentation](https://www.scaleway.com/en/docs/network/vpc/)
// - The #regional-vpc-beta channel on our [Slack Community](https://www.scaleway.com/en/docs/tutorials/scaleway-slack-community/)
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
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
}

// ListVPCs: List existing VPCs in the specified region.
func (s *API) ListVPCs(req *ListVPCsRequest, opts ...scw.RequestOption) (*ListVPCsResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "is_default", req.IsDefault)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/vpcs",
		Query:  query,
	}

	var resp ListVPCsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVPC: Create a new VPC in the specified region.
func (s *API) CreateVPC(req *CreateVPCRequest, opts ...scw.RequestOption) (*VPC, error) {
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
		req.Name = namegenerator.GetRandomName("vpc")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/vpcs",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp VPC

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVPC: Retrieve details of an existing VPC, specified by its VPC ID.
func (s *API) GetVPC(req *GetVPCRequest, opts ...scw.RequestOption) (*VPC, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.VpcID) == "" {
		return nil, errors.New("field VpcID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/vpcs/" + fmt.Sprint(req.VpcID) + "",
	}

	var resp VPC

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVPC: Update parameters including name and tags of the specified VPC.
func (s *API) UpdateVPC(req *UpdateVPCRequest, opts ...scw.RequestOption) (*VPC, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.VpcID) == "" {
		return nil, errors.New("field VpcID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/vpcs/" + fmt.Sprint(req.VpcID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp VPC

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVPC: Delete a VPC specified by its VPC ID.
func (s *API) DeleteVPC(req *DeleteVPCRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.VpcID) == "" {
		return errors.New("field VpcID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/vpcs/" + fmt.Sprint(req.VpcID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListPrivateNetworks: List existing Private Networks in the specified region. By default, the Private Networks returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field.
func (s *API) ListPrivateNetworks(req *ListPrivateNetworksRequest, opts ...scw.RequestOption) (*ListPrivateNetworksResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "private_network_ids", req.PrivateNetworkIDs)
	parameter.AddToQuery(query, "vpc_id", req.VpcID)
	parameter.AddToQuery(query, "dhcp_enabled", req.DHCPEnabled)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/private-networks",
		Query:  query,
	}

	var resp ListPrivateNetworksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreatePrivateNetwork: Create a new Private Network. Once created, you can attach Scaleway resources which are in the same region.
func (s *API) CreatePrivateNetwork(req *CreatePrivateNetworkRequest, opts ...scw.RequestOption) (*PrivateNetwork, error) {
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
		req.Name = namegenerator.GetRandomName("pn")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/private-networks",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PrivateNetwork

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrivateNetwork: Retrieve information about an existing Private Network, specified by its Private Network ID. Its full details are returned in the response object.
func (s *API) GetPrivateNetwork(req *GetPrivateNetworkRequest, opts ...scw.RequestOption) (*PrivateNetwork, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return nil, errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "",
	}

	var resp PrivateNetwork

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePrivateNetwork: Update parameters (such as name or tags) of an existing Private Network, specified by its Private Network ID.
func (s *API) UpdatePrivateNetwork(req *UpdatePrivateNetworkRequest, opts ...scw.RequestOption) (*PrivateNetwork, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return nil, errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PrivateNetwork

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeletePrivateNetwork: Delete an existing Private Network. Note that you must first detach all resources from the network, in order to delete it.
func (s *API) DeletePrivateNetwork(req *DeletePrivateNetworkRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// MigrateZonalPrivateNetworks: Transform multiple existing zoned Private Networks (scoped to a single Availability Zone) into regional Private Networks, scoped to an entire region. You can transform one or many Private Networks (specified by their Private Network IDs) within a single Scaleway Organization or Project, with the same call.
func (s *API) MigrateZonalPrivateNetworks(req *MigrateZonalPrivateNetworksRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.OrganizationID == nil && req.ProjectID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.OrganizationID == nil && req.ProjectID == nil {
		req.ProjectID = &defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/private-networks/migrate-zonal",
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

// EnableDHCP: Enable DHCP managed on an existing Private Network. Note that you will not be able to deactivate it afterwards.
func (s *API) EnableDHCP(req *EnableDHCPRequest, opts ...scw.RequestOption) (*PrivateNetwork, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return nil, errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "/enable-dhcp",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PrivateNetwork

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetSubnets: Set subnets for an existing Private Network. Note that the method is PUT and not PATCH. Any existing subnets will be removed in favor of the new specified set of subnets.
func (s *API) SetSubnets(req *SetSubnetsRequest, opts ...scw.RequestOption) (*SetSubnetsResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return nil, errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "/subnets",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetSubnetsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddSubnets: Add new subnets to an existing Private Network.
func (s *API) AddSubnets(req *AddSubnetsRequest, opts ...scw.RequestOption) (*AddSubnetsResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return nil, errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "/subnets",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AddSubnetsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSubnets: Delete the specified subnets from a Private Network.
func (s *API) DeleteSubnets(req *DeleteSubnetsRequest, opts ...scw.RequestOption) (*DeleteSubnetsResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return nil, errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "/subnets",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DeleteSubnetsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
