// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package vpc provides methods and message types of the vpc v1 API.
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

type APIListPrivateNetworksRequestOrderBy string

const (
	APIListPrivateNetworksRequestOrderByCreatedAtAsc  = APIListPrivateNetworksRequestOrderBy("created_at_asc")
	APIListPrivateNetworksRequestOrderByCreatedAtDesc = APIListPrivateNetworksRequestOrderBy("created_at_desc")
	APIListPrivateNetworksRequestOrderByNameAsc       = APIListPrivateNetworksRequestOrderBy("name_asc")
	APIListPrivateNetworksRequestOrderByNameDesc      = APIListPrivateNetworksRequestOrderBy("name_desc")
)

func (enum APIListPrivateNetworksRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum APIListPrivateNetworksRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *APIListPrivateNetworksRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = APIListPrivateNetworksRequestOrderBy(APIListPrivateNetworksRequestOrderBy(tmp).String())
	return nil
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
	// Zone: Availability Zone in which the Private Network is available.
	Zone scw.Zone `json:"zone"`
	// Tags: Tags of the Private Network.
	Tags []string `json:"tags"`
	// CreatedAt: Date the Private Network was created.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Date the Private Network was last modified.
	UpdatedAt *time.Time `json:"updated_at"`
	// Subnets: Private Network subnets CIDR.
	Subnets []scw.IPNet `json:"subnets"`
}

// APICreatePrivateNetworkRequest:
type APICreatePrivateNetworkRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Name: Name for the Private Network.
	Name string `json:"name"`
	// ProjectID: Scaleway Project in which to create the Private Network.
	ProjectID string `json:"project_id"`
	// Tags: Tags for the Private Network.
	Tags []string `json:"tags"`
	// Subnets: Private Network subnets CIDR.
	Subnets []scw.IPNet `json:"subnets"`
}

// APIDeletePrivateNetworkRequest:
type APIDeletePrivateNetworkRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// PrivateNetworkID: Private Network ID.
	PrivateNetworkID string `json:"-"`
}

// APIGetPrivateNetworkRequest:
type APIGetPrivateNetworkRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// PrivateNetworkID: Private Network ID.
	PrivateNetworkID string `json:"-"`
}

// APIListPrivateNetworksRequest:
type APIListPrivateNetworksRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// OrderBy: Sort order of the returned Private Networks.
	OrderBy APIListPrivateNetworksRequestOrderBy `json:"-"`
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
	// IncludeRegional: Defines whether to include regional Private Networks in the response.
	IncludeRegional *bool `json:"-"`
}

// APIUpdatePrivateNetworkRequest:
type APIUpdatePrivateNetworkRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// PrivateNetworkID: Private Network ID.
	PrivateNetworkID string `json:"-"`
	// Name: Name of the private network.
	Name *string `json:"name,omitempty"`
	// Tags: Tags for the Private Network.
	Tags *[]string `json:"tags,omitempty"`
	// Deprecated: Subnets: Private Network subnets CIDR (deprecated).
	Subnets *[]string `json:"subnets,omitempty"`
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

// This API concerns the zoned Private Networks service. Private Networks allows Scaleway resources (Instances, Load Balancers, Managed Databases etc.) within a single Availability Zone to be interconnected through a dedicated, private, and flexible [L2 network](https://en.wikipedia.org/wiki/Data_link_layer).
//
// You can add as many resources to your networks as you want, and add up to eight (8) different networks
// per resource. This allows you to run services isolated from the public internet and expose them to the rest of your infrastructure without worrying about public network filtering.
//
// (switchcolumn)
// (switchcolumn)
//
// ## Concepts
//
// Refer to the [Public Gateway concepts page](https://www.scaleway.com/en/docs/network/public-gateways/concepts/) to find definitions of concepts and terminology related to Private Networks.
//
// (switchcolumn)
// (switchcolumn)
//
// ## Quickstart
//
// 1. **Configure your environment variables**
//
//	<Message type="note">
//	This is an optional step that seeks to simplify your usage of the API. See [Availability Zones](#technical-information) below for help choosing an Availability Zone. You can find your Project ID in the [Scaleway console](https://console.scaleway.com/project/settings).
//	</Message>
//
//	```bash
//	export SCW_SECRET_KEY="<API secret key>"
//	export SCW_DEFAULT_ZONE="<Scaleway Availability Zone>"
//	export SCW_PROJECT_ID="<Scaleway Project ID>"
//	```
//
// 2. **Create a Private Network**. Run the following command to create a Private Network. You can customize the details in the payload (name, tags etc.) to your needs.
//
//	```bash
//	curl -X POST \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  -H "Content-Type: application/json" \
//	  "https://api.scaleway.com/vpc/v1/zones/$SCW_DEFAULT_ZONE/private-networks" \
//	  -d '{
//	    "name": "My new Private Network",
//	    "project_id": "'"$SCW_PROJECT_ID"'",
//	    "tags": ["test", "dev"]
//	  }'
//	```
//
//	<Message type="tip">
//	Keep the `id` field of the response: it is your Private Network ID, and is useable across all Scaleway products that support Private Networks. It may be useful to you to export the Private Network ID as a new environment variable `export PN_ID="<Your Private Network ID>`
//	</Message>
//
// 3. **Attach a resource to your Private Network**. Each Scaleway product has its own API to interact with Private Networks. To attach an Instance, Managed Database, Elastic Metal server, Load Balancer or Public Gateway to your Private Network, see instructions in the documentation of the relevant product API. Here, we take the example of an Instance.
//
//	Use the following call to attach an Instance to your Private Network. Ensure you replace `<Instance ID>` with the ID of your Instance, and `<Private Network ID>` with the ID of your Private Network.
//
//	```bash
//	curl -X POST \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  -H "Content-Type: application/json" \
//	  "https://api.scaleway.com/instance/v1/zones/$SCW_DEFAULT_ZONE/servers/<Instance ID>/private_nics" \
//	  -d '{"private_network_id": "<Private Network ID>"}'
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
// 6. **Configure the network interface on your Instance**. Refer to our [dedicated documentation](https://www.scaleway.com/en/docs/compute/instances/how-to/use-private-networks/#how-to-configure-the-private-network-interface-on-your-instances) for help with this step. Alternatively, use a [Public Gateway](https://www.scaleway.com/en/docs/network/public-gateways/quickstart/) to facilitate configuration of your resources on your Private Network.
//
// 7. **Delete your Private NIC**, which equates to unplugging your Instance from the Private Network. Use the following call. Ensure you replace `<Instance ID>` with the ID of your Instance, `<Private Network ID>` with the ID of your Private Network, and `<NIC ID>` with the ID of your Private NIC.
//
//	```bash
//	curl -X DELETE \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  -H "Content-Type: application/json" \
//	  "https://api.scaleway.com/instance/v1/zones/$SCW_DEFAULT_ZONE/servers/<Instance ID>/private_nics/<NIC ID>"
//	```
//
//	The network interface disappears from your Instance.
//
// 8. **Delete your Private Network**. Use the following call. Ensure you replace `<Private Network ID>` with the ID of your Private Network.
//
//	```bash
//	curl -X DELETE \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  -H "Content-Type: application/json" \
//	  "https://api.scaleway.com/vpc/v1/zones/$SCW_DEFAULT_ZONE/private-networks/<Private Network ID>"
//
//	```
//
//	<Message type="note">
//	Private Networks must be empty to be deleted. Ensure you have detached all resources from your network prior to deletion.
//	</Message>
//
// (switchcolumn)
// <Message type="requirement">
// - You have a [Scaleway account](https://console.scaleway.com/)
// - You have created an [API key](https://www.scaleway.com/en/docs/identity-and-access-management/iam/how-to/create-api-keys/) and that the API key has sufficient [IAM permissions](https://www.scaleway.com/en/docs/identity-and-access-management/iam/reference-content/permission-sets/) to perform the actions described on this page
// - You have [installed `curl`](https://curl.se/download.html)
// </Message>
// (switchcolumn)
//
// ## Technical information
//
// ### Availability Zones
//
// Scaleway Private Networks is a zoned product. It is available in the following Availability Zones:
//
// | Name      | API ID                |
// |-----------|-----------------------|
// | Paris     | `fr-par-1` `fr-par-2` `fr-par-3` |
// | Amsterdam | `nl-ams-1` `nl-ams-2` |
// | Warsaw    | `pl-waw-1` `pl-waw-2` |
//
// ## Technical limitations
//
// The following limitations apply to Scaleway Private Networks:
//
// - A maximum of eight (8) Private Networks can be plugged to any single Instance
// - The MAC address of an Instance in a Private Network cannot be changed
// - Broadcast and multicast traffic, while supported, are heavily rate-limited
//
// ## Going further
//
// For more help using Scaleway Private Networks, check out the following resources:
// - Our [main documentation](https://www.scaleway.com/en/docs/network/public-gateways/how-to/use-private-networks/)
// - The #private-networks channel on our [Slack Community](https://www.scaleway.com/en/docs/tutorials/scaleway-slack-community/)
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
	return []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2}
}

// ListPrivateNetworks: List existing Private Networks in a specified Availability Zone. By default, the Private Networks returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field.
func (s *API) ListPrivateNetworks(req *APIListPrivateNetworksRequest, opts ...scw.RequestOption) (*ListPrivateNetworksResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "private_network_ids", req.PrivateNetworkIDs)
	parameter.AddToQuery(query, "include_regional", req.IncludeRegional)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc/v1/zones/" + fmt.Sprint(req.Zone) + "/private-networks",
		Query:  query,
	}

	var resp ListPrivateNetworksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreatePrivateNetwork: Create a new Private Network. Once created, you can attach Scaleway resources in the same Availability Zone.
func (s *API) CreatePrivateNetwork(req *APICreatePrivateNetworkRequest, opts ...scw.RequestOption) (*PrivateNetwork, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("pn")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vpc/v1/zones/" + fmt.Sprint(req.Zone) + "/private-networks",
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
func (s *API) GetPrivateNetwork(req *APIGetPrivateNetworkRequest, opts ...scw.RequestOption) (*PrivateNetwork, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return nil, errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc/v1/zones/" + fmt.Sprint(req.Zone) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "",
	}

	var resp PrivateNetwork

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePrivateNetwork: Update parameters (such as name or tags) of an existing Private Network, specified by its Private Network ID.
func (s *API) UpdatePrivateNetwork(req *APIUpdatePrivateNetworkRequest, opts ...scw.RequestOption) (*PrivateNetwork, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return nil, errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/vpc/v1/zones/" + fmt.Sprint(req.Zone) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "",
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
func (s *API) DeletePrivateNetwork(req *APIDeletePrivateNetworkRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/vpc/v1/zones/" + fmt.Sprint(req.Zone) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}
