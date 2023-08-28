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

// VPC API.
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
