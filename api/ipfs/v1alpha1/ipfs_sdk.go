// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package ipfs provides methods and message types of the ipfs v1alpha1 API.
package ipfs

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

type APIListPinsRequestOrderBy string

const (
	APIListPinsRequestOrderByCreatedAtAsc  = APIListPinsRequestOrderBy("created_at_asc")
	APIListPinsRequestOrderByCreatedAtDesc = APIListPinsRequestOrderBy("created_at_desc")
)

func (enum APIListPinsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum APIListPinsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *APIListPinsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = APIListPinsRequestOrderBy(APIListPinsRequestOrderBy(tmp).String())
	return nil
}

type APIListVolumesRequestOrderBy string

const (
	APIListVolumesRequestOrderByCreatedAtAsc  = APIListVolumesRequestOrderBy("created_at_asc")
	APIListVolumesRequestOrderByCreatedAtDesc = APIListVolumesRequestOrderBy("created_at_desc")
)

func (enum APIListVolumesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum APIListVolumesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *APIListVolumesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = APIListVolumesRequestOrderBy(APIListVolumesRequestOrderBy(tmp).String())
	return nil
}

type PinDetails string

const (
	PinDetailsUnknownDetails                  = PinDetails("unknown_details")
	PinDetailsPinningLookingForProvider       = PinDetails("pinning_looking_for_provider")
	PinDetailsPinningInProgress               = PinDetails("pinning_in_progress")
	PinDetailsPinningBlocksFetched            = PinDetails("pinning_blocks_fetched")
	PinDetailsPinningFetchingURLData          = PinDetails("pinning_fetching_url_data")
	PinDetailsPinnedOk                        = PinDetails("pinned_ok")
	PinDetailsUnpinnedOk                      = PinDetails("unpinned_ok")
	PinDetailsUnpinningInProgress             = PinDetails("unpinning_in_progress")
	PinDetailsFailedContainsBannedCid         = PinDetails("failed_contains_banned_cid")
	PinDetailsFailedPinning                   = PinDetails("failed_pinning")
	PinDetailsFailedPinningNoProvider         = PinDetails("failed_pinning_no_provider")
	PinDetailsFailedPinningBadCidFormat       = PinDetails("failed_pinning_bad_cid_format")
	PinDetailsFailedPinningTimeout            = PinDetails("failed_pinning_timeout")
	PinDetailsFailedPinningTooBigContent      = PinDetails("failed_pinning_too_big_content")
	PinDetailsFailedPinningUnreachableURL     = PinDetails("failed_pinning_unreachable_url")
	PinDetailsFailedPinningBadURLFormat       = PinDetails("failed_pinning_bad_url_format")
	PinDetailsFailedPinningNoURLContentLength = PinDetails("failed_pinning_no_url_content_length")
	PinDetailsFailedPinningBadURLStatusCode   = PinDetails("failed_pinning_bad_url_status_code")
	PinDetailsFailedUnpinning                 = PinDetails("failed_unpinning")
	PinDetailsCheckingCoherence               = PinDetails("checking_coherence")
	PinDetailsRescheduled                     = PinDetails("rescheduled")
)

func (enum PinDetails) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_details"
	}
	return string(enum)
}

func (enum PinDetails) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PinDetails) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PinDetails(PinDetails(tmp).String())
	return nil
}

type PinStatus string

const (
	PinStatusUnknownStatus = PinStatus("unknown_status")
	PinStatusQueued        = PinStatus("queued")
	PinStatusPinning       = PinStatus("pinning")
	PinStatusFailed        = PinStatus("failed")
	PinStatusPinned        = PinStatus("pinned")
)

func (enum PinStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum PinStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PinStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PinStatus(PinStatus(tmp).String())
	return nil
}

// PinCIDMeta:
type PinCIDMeta struct {
	// ID:
	ID *string `json:"id"`
}

// PinCID:
type PinCID struct {
	// Cid:
	Cid *string `json:"cid"`
	// Name:
	Name *string `json:"name"`
	// Origins:
	Origins []string `json:"origins"`
	// Meta:
	Meta *PinCIDMeta `json:"meta"`
}

// PinInfo:
type PinInfo struct {
	// ID:
	ID *string `json:"id"`
	// URL:
	URL *string `json:"url"`
	// Size:
	Size *uint64 `json:"size"`
	// Progress:
	Progress *uint32 `json:"progress"`
	// StatusDetails:
	StatusDetails PinDetails `json:"status_details"`
}

// PinOptions:
type PinOptions struct {
	// RequiredZones:
	RequiredZones []string `json:"required_zones"`
	// ReplicationCount:
	ReplicationCount uint32 `json:"replication_count"`
}

// Pin:
type Pin struct {
	// PinID:
	PinID string `json:"pin_id"`
	// Status:
	Status PinStatus `json:"status"`
	// CreatedAt:
	CreatedAt *time.Time `json:"created_at"`
	// Cid:
	Cid *PinCID `json:"cid"`
	// Delegates:
	Delegates []string `json:"delegates"`
	// Info:
	Info *PinInfo `json:"info"`
}

// Volume:
type Volume struct {
	// ID:
	ID string `json:"id"`
	// ProjectID:
	ProjectID string `json:"project_id"`
	// Region:
	Region scw.Region `json:"region"`
	// CountPin:
	CountPin uint64 `json:"count_pin"`
	// CreatedAt:
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt:
	UpdatedAt *time.Time `json:"updated_at"`
	// Tags:
	Tags []string `json:"tags"`
	// Name:
	Name string `json:"name"`
	// Size:
	Size *uint64 `json:"size"`
}

// APICreatePinByCIDRequest:
type APICreatePinByCIDRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// VolumeID:
	VolumeID string `json:"volume_id"`
	// Cid:
	Cid string `json:"cid"`
	// Name:
	Name *string `json:"name,omitempty"`
	// Origins:
	Origins []string `json:"origins"`
	// PinOptions:
	PinOptions *PinOptions `json:"pin_options"`
}

// APICreatePinByURLRequest:
type APICreatePinByURLRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// VolumeID:
	VolumeID string `json:"volume_id"`
	// URL:
	URL string `json:"url"`
	// Name:
	Name *string `json:"name,omitempty"`
	// PinOptions:
	PinOptions *PinOptions `json:"pin_options"`
}

// APICreateVolumeRequest:
type APICreateVolumeRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID:
	ProjectID string `json:"project_id"`
	// Name:
	Name string `json:"name"`
}

// APIDeletePinRequest:
type APIDeletePinRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// PinID:
	PinID string `json:"-"`
	// VolumeID:
	VolumeID string `json:"-"`
}

// APIDeleteVolumeRequest:
type APIDeleteVolumeRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// VolumeID:
	VolumeID string `json:"-"`
}

// APIGetPinRequest:
type APIGetPinRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// PinID:
	PinID string `json:"-"`
	// VolumeID:
	VolumeID string `json:"-"`
}

// APIGetVolumeRequest:
type APIGetVolumeRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// VolumeID:
	VolumeID string `json:"-"`
}

// APIListPinsRequest:
type APIListPinsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// VolumeID:
	VolumeID string `json:"-"`
	// ProjectID:
	ProjectID *string `json:"-"`
	// OrganizationID:
	OrganizationID *string `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
	// OrderBy:
	OrderBy APIListPinsRequestOrderBy `json:"-"`
	// Status:
	Status PinStatus `json:"-"`
}

// APIListVolumesRequest:
type APIListVolumesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID:
	ProjectID string `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
	// OrderBy:
	OrderBy APIListVolumesRequestOrderBy `json:"-"`
}

// APIReplacePinRequest:
type APIReplacePinRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// PinID:
	PinID string `json:"-"`
	// VolumeID:
	VolumeID string `json:"volume_id"`
	// Cid:
	Cid string `json:"cid"`
	// Name:
	Name *string `json:"name,omitempty"`
	// Origins:
	Origins []string `json:"origins"`
	// PinOptions:
	PinOptions *PinOptions `json:"pin_options"`
}

// APIUpdateVolumeRequest:
type APIUpdateVolumeRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// VolumeID:
	VolumeID string `json:"-"`
	// Name:
	Name *string `json:"name,omitempty"`
	// Tags:
	Tags *[]string `json:"tags,omitempty"`
}

// ListPinsResponse:
type ListPinsResponse struct {
	// TotalCount:
	TotalCount uint64 `json:"total_count"`
	// Pins:
	Pins []*Pin `json:"pins"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPinsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPinsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListPinsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Pins = append(r.Pins, results.Pins...)
	r.TotalCount += uint64(len(results.Pins))
	return uint64(len(results.Pins)), nil
}

// ListVolumesResponse:
type ListVolumesResponse struct {
	// Volumes:
	Volumes []*Volume `json:"volumes"`
	// TotalCount:
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListVolumesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListVolumesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListVolumesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Volumes = append(r.Volumes, results.Volumes...)
	r.TotalCount += uint64(len(results.Volumes))
	return uint64(len(results.Volumes)), nil
}

// ReplacePinResponse:
type ReplacePinResponse struct {
	// Pin:
	Pin *Pin `json:"pin"`
}

// The InterPlanetary File System (IPFS) is a modular suite of open data storage and exchange protocols. Anybody can participate in this peer-to-peer network. In mid-2022 more than 300K nodes were a part of the network. Content stored on IPFS nodes is powered by a content-addressing design, making it location-agnostic, verifiable, and immutable. This design makes IPFS a good choice for versioning, archiving, NFT, or ample content storage and distribution.
//
// Scaleway pinning service enables you to keep (to 'pin') your data always available on the public IPFS network. Our pinning service is available in all of our regions to increase the accessibility of your data.
//
// ## Concepts
//
// Refer to our [dedicated concepts page](https://www.scaleway.com/en/docs/labs/ipfs/concepts/) to find definitions of all IPFS-related terminologies.
//
// ## Quickstart
// Refer to our dedicated [IPFS Pinning Quickstart page](https://www.scaleway.com/en/docs/labs/ipfs/quickstart/) to begin using the service.
//
// You can also find a tutorial on [how to manage volumes](https://www.scaleway.com/en/docs/labs/ipfs/api-cli/volumes-operations/) and [pins](https://www.scaleway.com/en/docs/labs/ipfs/api-cli/ipfs-operations/).
//
// You can find also detailed explainations on [how to get a pinned content](https://www.scaleway.com/en/docs/labs/ipfs/api-cli/ipfs-get-content/).
//
// ## Technical information
// The Scaleway IPFS pinning service provides features such as:
// - Unlimited storage via our Object Storage backend
// - Possibility to pin a remote content by an URL
//
// ### Data privacy
// - Data is **public**: our IPFS nodes are bootstrapped with public IPFS nodes. This implies that **any pinned content will be available on the public** IPFS network
// - Data is **shared**: public IPFS nodes can **fetch and host** your pinned content. Your data could be hosted **anywhere** (France, Europe, America, the Moon...). Scaleway is not enabled to delete once a content is delivered to external peers
// - Data is **not encrypted**: be aware that we do not apply any encryption algorithm over your data
//
// In consequence, consider to pin public-compatible data and/or encrypt by your own your content.
//
// ### 0% SLA
// We strive to provide you with the best possible experience on our IPFS pinning. However, this is an experimental service. This is why we are not able to contractually commit to a level of service, hence an SLA of 0%.
// The guarantees of these "Labs" offers are detailed in our special conditions for BETA services.
//
// ### Clients
// The Scaleway IPFS Pinning service is available on different clients:
//
// - [Kubo CLI](https://www.scaleway.com/en/docs/labs/ipfs/api-cli/ipfs-cli/)
// - [IPFS desktop](https://www.scaleway.com/en/docs/labs/ipfs/api-cli/ipfs-desktop/)
// - Scaleway CLI to [manage volumes](https://www.scaleway.com/en/docs/labs/ipfs/api-cli/volumes-operations/) and [pins](https://www.scaleway.com/en/docs/labs/ipfs/api-cli/ipfs-operations/)
//
// ### Regions
// Scaleway's infrastructure is spread across different [regions and Availability Zones](https://www.scaleway.com/en/docs/console/my-account/reference-content/products-availability/).
//
// Scaleway IPFS pinning service is available in the Paris, Amsterdam and Warsaw regions, which are represented by the following path parameters:
// - `fr-par`
// - `nl-ams`
// - `pl-waw`.
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

// CreateVolume: Create a new volume from a Project ID. Volume is identified by an ID and used to host pin references.
// Volume is personal (at least to your organization) even if IPFS blocks and CID are available to anyone.
// Should be the first command you made because every pin must be attached to a volume.
func (s *API) CreateVolume(req *APICreateVolumeRequest, opts ...scw.RequestOption) (*Volume, error) {
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
		Path:   "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/volumes",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Volume

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVolume: Retrieve information about a specific volume.
func (s *API) GetVolume(req *APIGetVolumeRequest, opts ...scw.RequestOption) (*Volume, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.VolumeID) == "" {
		return nil, errors.New("field VolumeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/volumes/" + fmt.Sprint(req.VolumeID) + "",
	}

	var resp Volume

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVolumes: Retrieve information about all volumes from a Project ID.
func (s *API) ListVolumes(req *APIListVolumesRequest, opts ...scw.RequestOption) (*ListVolumesResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}
	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/volumes",
		Query:  query,
	}

	var resp ListVolumesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVolume: Update volume information (tag, name...).
func (s *API) UpdateVolume(req *APIUpdateVolumeRequest, opts ...scw.RequestOption) (*Volume, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.VolumeID) == "" {
		return nil, errors.New("field VolumeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/volumes/" + fmt.Sprint(req.VolumeID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Volume

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVolume: Delete a volume by its ID and every pin attached to this volume. This process can take a while to conclude, depending on the size of your pinned content.
func (s *API) DeleteVolume(req *APIDeleteVolumeRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.VolumeID) == "" {
		return errors.New("field VolumeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/volumes/" + fmt.Sprint(req.VolumeID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CreatePinByURL: Will fetch and store the content pointed by the provided URL. The content must be available on the public IPFS network.
// The content (IPFS blocks) will be host by the pinning service until pin deletion.
// From that point, any other IPFS peer can fetch and host your content: Make sure to pin public or encrypted content.
// Many pin requests (from different users) can target the same CID.
// A pin is defined by its ID (UUID), its status (queued, pinning, pinned or failed) and target CID.
func (s *API) CreatePinByURL(req *APICreatePinByURLRequest, opts ...scw.RequestOption) (*Pin, error) {
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
		Path:   "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/pins/create-by-url",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Pin

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreatePinByCID: Will fetch and store the content pointed by the provided CID. The content must be available on the public IPFS network.
// The content (IPFS blocks) will be host by the pinning service until pin deletion.
// From that point, any other IPFS peer can fetch and host your content: Make sure to pin public or encrypted content.
// Many pin requests (from different users) can target the same CID.
// A pin is defined by its ID (UUID), its status (queued, pinning, pinned or failed) and target CID.
func (s *API) CreatePinByCID(req *APICreatePinByCIDRequest, opts ...scw.RequestOption) (*Pin, error) {
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
		Path:   "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/pins/create-by-cid",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Pin

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReplacePin:
func (s *API) ReplacePin(req *APIReplacePinRequest, opts ...scw.RequestOption) (*ReplacePinResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PinID) == "" {
		return nil, errors.New("field PinID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/pins/" + fmt.Sprint(req.PinID) + "/replace",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ReplacePinResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPin: Retrieve information about the provided **pin ID**, such as status, last modification, and CID.
func (s *API) GetPin(req *APIGetPinRequest, opts ...scw.RequestOption) (*Pin, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "volume_id", req.VolumeID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PinID) == "" {
		return nil, errors.New("field PinID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/pins/" + fmt.Sprint(req.PinID) + "",
		Query:  query,
	}

	var resp Pin

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPins: Retrieve information about all pins into a volume.
func (s *API) ListPins(req *APIListPinsRequest, opts ...scw.RequestOption) (*ListPinsResponse, error) {
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
	parameter.AddToQuery(query, "volume_id", req.VolumeID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "status", req.Status)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/pins",
		Query:  query,
	}

	var resp ListPinsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeletePin: An unpin request means that you no longer own the content.
// This content can therefore be removed and no longer provided on the IPFS network.
func (s *API) DeletePin(req *APIDeletePinRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "volume_id", req.VolumeID)

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PinID) == "" {
		return errors.New("field PinID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/pins/" + fmt.Sprint(req.PinID) + "",
		Query:  query,
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}
