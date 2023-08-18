// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package applesilicon provides methods and message types of the applesilicon v1alpha1 API.
package applesilicon

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

type APIListServersRequestOrderBy string

const (
	APIListServersRequestOrderByCreatedAtAsc  = APIListServersRequestOrderBy("created_at_asc")
	APIListServersRequestOrderByCreatedAtDesc = APIListServersRequestOrderBy("created_at_desc")
)

func (enum APIListServersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum APIListServersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *APIListServersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = APIListServersRequestOrderBy(APIListServersRequestOrderBy(tmp).String())
	return nil
}

type ServerStatus string

const (
	ServerStatusUnknownStatus = ServerStatus("unknown_status")
	ServerStatusStarting      = ServerStatus("starting")
	ServerStatusReady         = ServerStatus("ready")
	ServerStatusError         = ServerStatus("error")
	ServerStatusRebooting     = ServerStatus("rebooting")
	ServerStatusUpdating      = ServerStatus("updating")
	ServerStatusLocking       = ServerStatus("locking")
	ServerStatusLocked        = ServerStatus("locked")
	ServerStatusUnlocking     = ServerStatus("unlocking")
	ServerStatusReinstalling  = ServerStatus("reinstalling")
)

func (enum ServerStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum ServerStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerStatus(ServerStatus(tmp).String())
	return nil
}

type ServerTypeStock string

const (
	ServerTypeStockUnknownStock = ServerTypeStock("unknown_stock")
	ServerTypeStockNoStock      = ServerTypeStock("no_stock")
	ServerTypeStockLowStock     = ServerTypeStock("low_stock")
	ServerTypeStockHighStock    = ServerTypeStock("high_stock")
)

func (enum ServerTypeStock) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_stock"
	}
	return string(enum)
}

func (enum ServerTypeStock) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerTypeStock) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerTypeStock(ServerTypeStock(tmp).String())
	return nil
}

// ServerTypeCPU:
type ServerTypeCPU struct {
	// Name:
	Name string `json:"name"`
	// CoreCount:
	CoreCount uint32 `json:"core_count"`
}

// ServerTypeDisk:
type ServerTypeDisk struct {
	// Capacity:
	Capacity scw.Size `json:"capacity"`
	// Type:
	Type string `json:"type"`
}

// ServerTypeMemory:
type ServerTypeMemory struct {
	// Capacity:
	Capacity scw.Size `json:"capacity"`
	// Type:
	Type string `json:"type"`
}

// OS:
type OS struct {
	// ID: Unique ID of the OS.
	ID string `json:"id"`
	// Name: OS name.
	Name string `json:"name"`
	// Label: OS name as it should be displayed.
	Label string `json:"label"`
	// ImageURL: URL of the image.
	ImageURL string `json:"image_url"`
	// CompatibleServerTypes: List of compatible server types.
	CompatibleServerTypes []string `json:"compatible_server_types"`
}

// ServerType:
type ServerType struct {
	// CPU: CPU description.
	CPU *ServerTypeCPU `json:"cpu"`
	// Disk: Size of the local disk of the server.
	Disk *ServerTypeDisk `json:"disk"`
	// Name: Name of the type.
	Name string `json:"name"`
	// Memory: Size of memory available.
	Memory *ServerTypeMemory `json:"memory"`
	// Stock: Current stock.
	Stock ServerTypeStock `json:"stock"`
	// MinimumLeaseDuration: Minimum duration of the lease in seconds (example. 3.4s).
	MinimumLeaseDuration *scw.Duration `json:"minimum_lease_duration"`
}

// Server:
type Server struct {
	// ID: UUID of the server.
	ID string `json:"id"`
	// Type: Type of the server.
	Type string `json:"type"`
	// Name: Name of the server.
	Name string `json:"name"`
	// ProjectID: Project this server is associated with.
	ProjectID string `json:"project_id"`
	// OrganizationID: Organization this server is associated with.
	OrganizationID string `json:"organization_id"`
	// IP: IPv4 address of the server.
	IP net.IP `json:"ip"`
	// VncURL: URL of the VNC.
	VncURL string `json:"vnc_url"`
	// Status: Current status of the server.
	Status ServerStatus `json:"status"`
	// CreatedAt: Date on which the server was created.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Date on which the server was last updated.
	UpdatedAt *time.Time `json:"updated_at"`
	// DeletableAt: Date on which the server was last deleted.
	DeletableAt *time.Time `json:"deletable_at"`
	// Zone: Zone of the server.
	Zone scw.Zone `json:"zone"`
}

// APICreateServerRequest:
type APICreateServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Name: Create a server with this given name.
	Name string `json:"name"`
	// ProjectID: Create a server in the given project ID.
	ProjectID string `json:"project_id"`
	// Type: Create a server of the given type.
	Type string `json:"type"`
}

// APIDeleteServerRequest:
type APIDeleteServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the server you want to delete.
	ServerID string `json:"-"`
}

// APIGetOSRequest:
type APIGetOSRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// OsID: UUID of the OS you want to get.
	OsID string `json:"-"`
}

// APIGetServerRequest:
type APIGetServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the server you want to get.
	ServerID string `json:"-"`
}

// APIGetServerTypeRequest:
type APIGetServerTypeRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerType: Server type identifier.
	ServerType string `json:"-"`
}

// APIListOSRequest:
type APIListOSRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Page: Positive integer to choose the page to return.
	Page *int32 `json:"-"`
	// PageSize: Positive integer lower or equal to 100 to select the number of items to return.
	PageSize *uint32 `json:"-"`
	// ServerType: List of compatible server types.
	ServerType *string `json:"-"`
	// Name: Filter OS by name (note that "11.1" will return "11.1.2" and "11.1" but not "12")).
	Name *string `json:"-"`
}

// APIListServerTypesRequest:
type APIListServerTypesRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
}

// APIListServersRequest:
type APIListServersRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// OrderBy: Sort order of the returned servers.
	OrderBy APIListServersRequestOrderBy `json:"-"`
	// ProjectID: Only list servers of this project ID.
	ProjectID *string `json:"-"`
	// OrganizationID: Only list servers of this Organization ID.
	OrganizationID *string `json:"-"`
	// Page: Positive integer to choose the page to return.
	Page *int32 `json:"-"`
	// PageSize: Positive integer lower or equal to 100 to select the number of items to return.
	PageSize *uint32 `json:"-"`
}

// APIRebootServerRequest:
type APIRebootServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the server you want to reboot.
	ServerID string `json:"-"`
}

// APIReinstallServerRequest:
type APIReinstallServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the server you want to reinstall.
	ServerID string `json:"-"`
}

// APIUpdateServerRequest:
type APIUpdateServerRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the server you want to update.
	ServerID string `json:"-"`
	// Name: Updated name for your server.
	Name *string `json:"name,omitempty"`
}

// ListOSResponse:
type ListOSResponse struct {
	// TotalCount: Total number of OS.
	TotalCount uint32 `json:"total_count"`
	// Os: List of OS.
	Os []*OS `json:"os"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListOSResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListOSResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListOSResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Os = append(r.Os, results.Os...)
	r.TotalCount += uint32(len(results.Os))
	return uint32(len(results.Os)), nil
}

// ListServerTypesResponse:
type ListServerTypesResponse struct {
	// ServerTypes: Available server types.
	ServerTypes []*ServerType `json:"server_types"`
}

// ListServersResponse:
type ListServersResponse struct {
	// TotalCount: Total number of servers.
	TotalCount uint32 `json:"total_count"`
	// Servers: Paginated returned servers.
	Servers []*Server `json:"servers"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListServersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Servers = append(r.Servers, results.Servers...)
	r.TotalCount += uint32(len(results.Servers))
	return uint32(len(results.Servers)), nil
}

// Scaleway Apple silicon is built using Apple's fifth generation Mac mini hardware. Powered by Apple’s M1 silicon and equipped with a 256 GB SSD and 8 GB of RAM, it's designed for towering performance and revolutionary power efficiency. The M1 chip has an Apple Silicon processor: a powerful 8-core CPU running at up to 3.2 GHz.
//
// This M1 processor is designed for developers to create and sign applications for a wide range of Apple devices: from iPhones and iPads to Mac computers and beyond. With its advanced neural engine, Mac mini M1 can process machine learning tasks up to 15 times faster than traditional computers.
//
// A dedicated Mac mini M1 thus allows you to enjoy dazzling performance and speed, enabling you to explore, learn, and build like never before.
//
// (switchcolumn)
// (switchcolumn)
//
// ## Concepts
//
// Refer to our [dedicated concepts page](https://www.scaleway.com/en/docs/compute/apple-silicon/concepts/) to find definitions of the different terms referring to Apple silicon.
//
// (switchcolumn)
// (switchcolumn)
//
// ## Quickstart
//
// 1. Configure your environment variables.
//
//	<Message type="note">
//	This is an optional step that seeks to simplify your usage of the APIs. Since there is only one Availability Zone for Apple silicon, the possible value for zone is `fr-par-3`.
//	</Message>
//
//	```bash
//	export SCW_SECRET_KEY="<API secret key>"
//	export SCW_DEFAULT_ZONE="<Scaleway Availability Zone>"
//	```
//
// 2. Edit the POST request payload that we will use in the next step to create an Apple silicon server.
//
//	Modify the values in the example according to your needs, using the information in the table to help.
//
//	```json
//	{
//	  "name": "My_AS_Server",
//	  "project_id": "7281793f-8474-727d-7688-72893f747g7e",
//	  "type": "M1-M"
//	}
//	```
//
//	| Parameter        | Description                                                                                              |
//	| :--------------- | :------------------------------------------------------------------------------------------------------- |
//	| `project_id`     | ID of the Project you want to create your server in. Your project name can only contain alphanumeric characters, spaces, dots and dashes. To find your Project ID, you can consult the **[Scaleway console](https://console.scaleway.com/project/settings)**                                                                                                       |
//	| `name`           | Create a server with this given name.                                                                    |
//	| `type`           | Create a server of the given type. As Apple silicon-as-a-service currently only offers one type, the possible value is `M1-M`.                                                                                                     |
//
// 3. Run the following command to create a server. Make sure you include the payload you edited in the previous step.
//
//	```bash
//	curl -X POST \
//	  -H "X-Auth-Token: $SECRET_KEY" \
//	  -H "Content-Type: application/json" \
//	  "https://api.scaleway.com/apple-silicon/v1alpha1/zones/$ZONE/servers" \
//	  -d '{
//	    "name": "My_Mac_Mini_M1",
//	    "project_id": "7281793f-8474-727d-7688-72893f747g7e",
//	    "type": "M1-M"
//	  }'
//	```
//
//	You should get an output similar to the following one, providing details about your Apple silicon server.
//
//	<Message type="note">
//	This is a response example, the UUIDs and IP address displayed are not real.
//	</Message>
//
//	```bash
//	{
//	  "id": "d6877c16-77af-40a2-8b13-514cee3e73f5",
//	  "type": "M1-M",
//	  "name": "My_Mac_Mini_M1",
//	  "project_id": "7281793f-8474-727d-7688-72893f747g7e",
//	  "organization_id": "b12d5c3g-c612-5674-c1e9-92627f36c5b9",
//	  "ip": "1.2.3.4",
//	  "vnc_url": "https://192.168.1.93",
//	  "status": "unknown_status",
//	  "created_at": "2022-03-22T12:34:56.123456Z",
//	  "updated_at": "2022-03-22T12:34:56.123456Z",
//	  "deletable_at": "2022-03-22T12:34:56.123456Z",
//	  "zone": "fr-par-3"
//	}
//	```
//
// 4. Retrieve your Mac mini M1 server IP from the response.
//
// 5. Connect to your Mac mini M1 server using SSH:
//
//	  ```bash
//	  ssh m1@<your_mac_mini_m1_ip>
//	  ```
//
//	(switchcolumn)
//	<Message type="requirement">
//	- You have a [Scaleway account](https://console.scaleway.com/)
//	- You have created an [API key](https://www.scaleway.com/en/docs/identity-and-access-management/iam/how-to/create-api-keys/) and that the API key has sufficient [IAM permissions](https://www.scaleway.com/en/docs/identity-and-access-management/iam/reference-content/permission-sets/) to perform the actions described on this page
//	- You have [installed `curl`](https://curl.se/download.html)
//	</Message>
//	(switchcolumn)
//
// ## Technical information
//
// Mac mini and macOS are trademarks of Apple Inc., registered in the U.S. and other countries and regions. IOS is a trademark or registered trademark of Cisco in the U.S. and other countries and is used by Apple under license. Scaleway is not affiliated with Apple Inc.
//
//	### Availability Zones
//
//	Mac mini M1 can be deployed in the following Availability Zone:
//
//	- Paris `fr-par-3`
//
// ## Technical limitations
//
// The following technical limiations apply when using Apple silicon:
//
// - Apple silicon-as-a-Service comes with a minimum allocation period of 24 hours.
// - There is a maximum of one Mac mini M1 allowed per Organization per Availability Zone. Learn more about account quotas with our dedicated [documentation](https://www.scaleway.com/en/docs/console/my-account/reference-content/account-quotas/#apple-silicon-m1).
//
// ## Going further
//
// For more help using Scaleway Apple silicon, check out the following resources:
// - Our [main documentation](https://www.scaleway.com/en/docs/compute/apple-silicon/)
// - Our [Slack Community](scaleway-community.slack.com)
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
	return []scw.Zone{scw.ZoneFrPar3}
}

// ListServerTypes: List all technical details about Apple silicon server types available in the specified zone. Since there is only one Availability Zone for Apple silicon servers, the targeted value is `fr-par-3`.
func (s *API) ListServerTypes(req *APIListServerTypesRequest, opts ...scw.RequestOption) (*ListServerTypesResponse, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/server-types",
	}

	var resp ListServerTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetServerType: Get technical details (CPU, disk size etc.) of a server type.
func (s *API) GetServerType(req *APIGetServerTypeRequest, opts ...scw.RequestOption) (*ServerType, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerType) == "" {
		return nil, errors.New("field ServerType cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/server-type/" + fmt.Sprint(req.ServerType) + "",
	}

	var resp ServerType

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateServer: Create a new server in the targeted zone, specifying its configuration including name and type.
func (s *API) CreateServer(req *APICreateServerRequest, opts ...scw.RequestOption) (*Server, error) {
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
		req.Name = namegenerator.GetRandomName("as")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListServers: List all servers in the specified zone. By default, returned servers in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListServers(req *APIListServersRequest, opts ...scw.RequestOption) (*ListServersResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers",
		Query:  query,
	}

	var resp ListServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListOS: List all Operating Systems (OS). The response will include the total number of OS as well as their associated IDs, names and labels.
func (s *API) ListOS(req *APIListOSRequest, opts ...scw.RequestOption) (*ListOSResponse, error) {
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "server_type", req.ServerType)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/os",
		Query:  query,
	}

	var resp ListOSResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOS: Get an Operating System (OS).  The response will include the OS's unique ID as well as its name and label.
func (s *API) GetOS(req *APIGetOSRequest, opts ...scw.RequestOption) (*OS, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.OsID) == "" {
		return nil, errors.New("field OsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/os/" + fmt.Sprint(req.OsID) + "",
	}

	var resp OS

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetServer: Retrieve information about an existing Apple silicon server, specified by its server ID. Its full details, including name, status and IP address, are returned in the response object.
func (s *API) GetServer(req *APIGetServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateServer: Update the parameters of an existing Apple silicon server, specified by its server ID.
func (s *API) UpdateServer(req *APIUpdateServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteServer: Delete an existing Apple silicon server, specified by its server ID. Deleting a server is permanent, and cannot be undone. Note that the minimum allocation period for Apple silicon-as-a-service is 24 hours, meaning you cannot delete your server prior to that.
func (s *API) DeleteServer(req *APIDeleteServerRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// RebootServer: Reboot an existing Apple silicon server, specified by its server ID.
func (s *API) RebootServer(req *APIRebootServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/reboot",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReinstallServer: Reinstall an existing Apple silicon server (specified by its server ID) from a new image (OS). All the data on the disk is deleted and all configuration is reset to the defailt configuration values of the image (OS).
func (s *API) ReinstallServer(req *APIReinstallServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/reinstall",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
