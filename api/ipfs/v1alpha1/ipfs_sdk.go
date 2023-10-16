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

type ListNamesRequestOrderBy string

const (
	ListNamesRequestOrderByCreatedAtAsc  = ListNamesRequestOrderBy("created_at_asc")
	ListNamesRequestOrderByCreatedAtDesc = ListNamesRequestOrderBy("created_at_desc")
)

func (enum ListNamesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListNamesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListNamesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListNamesRequestOrderBy(ListNamesRequestOrderBy(tmp).String())
	return nil
}

type ListPinsRequestOrderBy string

const (
	ListPinsRequestOrderByCreatedAtAsc  = ListPinsRequestOrderBy("created_at_asc")
	ListPinsRequestOrderByCreatedAtDesc = ListPinsRequestOrderBy("created_at_desc")
)

func (enum ListPinsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListPinsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPinsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPinsRequestOrderBy(ListPinsRequestOrderBy(tmp).String())
	return nil
}

type ListVolumesRequestOrderBy string

const (
	ListVolumesRequestOrderByCreatedAtAsc  = ListVolumesRequestOrderBy("created_at_asc")
	ListVolumesRequestOrderByCreatedAtDesc = ListVolumesRequestOrderBy("created_at_desc")
)

func (enum ListVolumesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListVolumesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListVolumesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListVolumesRequestOrderBy(ListVolumesRequestOrderBy(tmp).String())
	return nil
}

type NameStatus string

const (
	NameStatusUnknownStatus = NameStatus("unknown_status")
	NameStatusQueued        = NameStatus("queued")
	NameStatusPublishing    = NameStatus("publishing")
	NameStatusFailed        = NameStatus("failed")
	NameStatusPublished     = NameStatus("published")
)

func (enum NameStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum NameStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NameStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NameStatus(NameStatus(tmp).String())
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

// Name:
type Name struct {
	// NameID:
	NameID string `json:"name_id"`
	// ProjectID:
	ProjectID string `json:"project_id"`
	// CreatedAt:
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt:
	UpdatedAt *time.Time `json:"updated_at"`
	// Tags:
	Tags []string `json:"tags"`
	// Name:
	Name string `json:"name"`
	// Key:
	Key string `json:"key"`
	// Status:
	Status NameStatus `json:"status"`
	// Value:
	Value string `json:"value"`
	// Region:
	Region scw.Region `json:"region"`
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

// CreatePinByCIDRequest:
type CreatePinByCIDRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// VolumeID: Volume ID on which you want to pin your content.
	VolumeID string `json:"volume_id"`
	// Cid: CID containing the content you want to pin.
	Cid string `json:"cid"`
	// Origins: Node containing the content you want to pin.
	Origins []string `json:"origins"`
	// Name: Pin name.
	Name *string `json:"name,omitempty"`
	// PinOptions: Pin options.
	PinOptions *PinOptions `json:"pin_options"`
}

// CreatePinByURLRequest:
type CreatePinByURLRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// VolumeID: Volume ID on which you want to pin your content.
	VolumeID string `json:"volume_id"`
	// URL: URL containing the content you want to pin.
	URL string `json:"url"`
	// Name: Pin name.
	Name *string `json:"name,omitempty"`
	// PinOptions: Pin options.
	PinOptions *PinOptions `json:"pin_options"`
}

// CreateVolumeRequest:
type CreateVolumeRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID: Project ID.
	ProjectID string `json:"project_id"`
	// Name: Volume name.
	Name string `json:"name"`
}

// DeletePinRequest:
type DeletePinRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// PinID: Pin ID you want to remove from the volume.
	PinID string `json:"-"`
	// VolumeID: Volume ID.
	VolumeID string `json:"volume_id"`
}

// DeleteVolumeRequest:
type DeleteVolumeRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// VolumeID: Volume ID.
	VolumeID string `json:"-"`
}

// ExportKeyNameResponse:
type ExportKeyNameResponse struct {
	// NameID:
	NameID string `json:"name_id"`
	// ProjectID:
	ProjectID string `json:"project_id"`
	// CreatedAt:
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt:
	UpdatedAt *time.Time `json:"updated_at"`
	// PublicKey:
	PublicKey string `json:"public_key"`
	// PrivateKey:
	PrivateKey string `json:"private_key"`
}

// GetPinRequest:
type GetPinRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// PinID: Pin ID of which you want to obtain information.
	PinID string `json:"-"`
	// VolumeID: Volume ID.
	VolumeID string `json:"volume_id"`
}

// GetVolumeRequest:
type GetVolumeRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// VolumeID: Volume ID.
	VolumeID string `json:"-"`
}

// IpnsAPICreateNameRequest:
type IpnsAPICreateNameRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID: Project ID.
	ProjectID string `json:"project_id"`
	// Name: Name for your records.
	Name string `json:"name"`
	// Value: Value you want to associate with your records, CID or IPNS key.
	Value string `json:"value"`
}

// IpnsAPIDeleteNameRequest:
type IpnsAPIDeleteNameRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NameID: Name ID you wish to delete.
	NameID string `json:"-"`
}

// IpnsAPIExportKeyNameRequest:
type IpnsAPIExportKeyNameRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NameID: Name ID whose keys you want to export.
	NameID string `json:"-"`
}

// IpnsAPIGetNameRequest:
type IpnsAPIGetNameRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NameID: Name ID whose information you want to retrieve.
	NameID string `json:"-"`
}

// IpnsAPIImportKeyNameRequest:
type IpnsAPIImportKeyNameRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID: Project ID.
	ProjectID string `json:"project_id"`
	// Name: Name for your records.
	Name string `json:"name"`
	// PrivateKey: Base64 private key.
	PrivateKey string `json:"private_key"`
	// Value: Value you want to associate with your records, CID or IPNS key.
	Value string `json:"value"`
}

// IpnsAPIListNamesRequest:
type IpnsAPIListNamesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID: Project ID.
	ProjectID *string `json:"-"`
	// OrganizationID: Organization ID.
	OrganizationID *string `json:"-"`
	// OrderBy: Sort the order of the returned names.
	OrderBy ListNamesRequestOrderBy `json:"-"`
	// Page: Page number.
	Page *int32 `json:"-"`
	// PageSize: Maximum number of names to return per page.
	PageSize *uint32 `json:"-"`
}

// IpnsAPIUpdateNameRequest:
type IpnsAPIUpdateNameRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NameID: Name ID you wish to update.
	NameID string `json:"-"`
	// Name: New name you want to associate with your record.
	Name *string `json:"name,omitempty"`
	// Tags: New tags you want to associate with your record.
	Tags *[]string `json:"tags,omitempty"`
	// Value: Value you want to associate with your records, CID or IPNS key.
	Value *string `json:"value,omitempty"`
}

// ListNamesResponse:
type ListNamesResponse struct {
	// Names:
	Names []*Name `json:"names"`
	// TotalCount:
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNamesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNamesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListNamesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Names = append(r.Names, results.Names...)
	r.TotalCount += uint64(len(results.Names))
	return uint64(len(results.Names)), nil
}

// ListPinsRequest:
type ListPinsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// VolumeID: Volume ID of which you want to list the pins.
	VolumeID string `json:"-"`
	// ProjectID: Project ID.
	ProjectID *string `json:"-"`
	// OrganizationID: Organization ID.
	OrganizationID *string `json:"-"`
	// OrderBy: Sort order of the returned Volume.
	OrderBy ListPinsRequestOrderBy `json:"-"`
	// Page: Page number.
	Page *int32 `json:"-"`
	// PageSize: Maximum number of volumes to return per page.
	PageSize *uint32 `json:"-"`
	// Status: List pins by status.
	Status PinStatus `json:"-"`
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

// ListVolumesRequest:
type ListVolumesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID: Project ID, only volumes belonging to this project will be listed.
	ProjectID string `json:"-"`
	// OrderBy: Sort the order of the returned volumes.
	OrderBy ListVolumesRequestOrderBy `json:"-"`
	// Page: Page number.
	Page *int32 `json:"-"`
	// PageSize: Maximum number of volumes to return per page.
	PageSize *uint32 `json:"-"`
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

// ReplacePinRequest:
type ReplacePinRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// PinID: Pin ID whose information you wish to replace.
	PinID string `json:"-"`
	// VolumeID: Volume ID.
	VolumeID string `json:"volume_id"`
	// Cid: New CID you want to pin in place of the old one.
	Cid string `json:"cid"`
	// Name: New name to replace.
	Name *string `json:"name,omitempty"`
	// Origins: Node containing the content you want to pin.
	Origins []string `json:"origins"`
	// PinOptions: Pin options.
	PinOptions *PinOptions `json:"pin_options"`
}

// ReplacePinResponse:
type ReplacePinResponse struct {
	// Pin:
	Pin *Pin `json:"pin"`
}

// UpdateVolumeRequest:
type UpdateVolumeRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// VolumeID: Volume ID.
	VolumeID string `json:"-"`
	// Name: Volume name.
	Name *string `json:"name,omitempty"`
	// Tags: Tags of the volume.
	Tags *[]string `json:"tags,omitempty"`
}

// IPFS Pinning service API.
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
func (s *API) CreateVolume(req *CreateVolumeRequest, opts ...scw.RequestOption) (*Volume, error) {
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
func (s *API) GetVolume(req *GetVolumeRequest, opts ...scw.RequestOption) (*Volume, error) {
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
func (s *API) ListVolumes(req *ListVolumesRequest, opts ...scw.RequestOption) (*ListVolumesResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

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
func (s *API) UpdateVolume(req *UpdateVolumeRequest, opts ...scw.RequestOption) (*Volume, error) {
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
func (s *API) DeleteVolume(req *DeleteVolumeRequest, opts ...scw.RequestOption) error {
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
func (s *API) CreatePinByURL(req *CreatePinByURLRequest, opts ...scw.RequestOption) (*Pin, error) {
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
func (s *API) CreatePinByCID(req *CreatePinByCIDRequest, opts ...scw.RequestOption) (*Pin, error) {
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
func (s *API) ReplacePin(req *ReplacePinRequest, opts ...scw.RequestOption) (*ReplacePinResponse, error) {
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
func (s *API) GetPin(req *GetPinRequest, opts ...scw.RequestOption) (*Pin, error) {
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

// ListPins: Retrieve information about all pins within a volume.
func (s *API) ListPins(req *ListPinsRequest, opts ...scw.RequestOption) (*ListPinsResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
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
func (s *API) DeletePin(req *DeletePinRequest, opts ...scw.RequestOption) error {
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

type IpnsAPI struct {
	client *scw.Client
}

// NewIpnsAPI returns a IpnsAPI object from a Scaleway client.
func NewIpnsAPI(client *scw.Client) *IpnsAPI {
	return &IpnsAPI{
		client: client,
	}
}
func (s *IpnsAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
}

// CreateName: You can use the `ipns key` command to list and generate more names and their respective keys.
func (s *IpnsAPI) CreateName(req *IpnsAPICreateNameRequest, opts ...scw.RequestOption) (*Name, error) {
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
		Path:   "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/names",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Name

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetName: Retrieve information about a specific name.
func (s *IpnsAPI) GetName(req *IpnsAPIGetNameRequest, opts ...scw.RequestOption) (*Name, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NameID) == "" {
		return nil, errors.New("field NameID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/names/" + fmt.Sprint(req.NameID) + "",
	}

	var resp Name

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteName: Delete a name by its ID.
func (s *IpnsAPI) DeleteName(req *IpnsAPIDeleteNameRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NameID) == "" {
		return errors.New("field NameID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/names/" + fmt.Sprint(req.NameID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListNames: Retrieve information about all names from a Project ID.
func (s *IpnsAPI) ListNames(req *IpnsAPIListNamesRequest, opts ...scw.RequestOption) (*ListNamesResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/names",
		Query:  query,
	}

	var resp ListNamesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateName: Update name information (CID, tag, name...).
func (s *IpnsAPI) UpdateName(req *IpnsAPIUpdateNameRequest, opts ...scw.RequestOption) (*Name, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NameID) == "" {
		return nil, errors.New("field NameID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/names/" + fmt.Sprint(req.NameID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Name

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExportKeyName: Export a private key by its ID.
func (s *IpnsAPI) ExportKeyName(req *IpnsAPIExportKeyNameRequest, opts ...scw.RequestOption) (*ExportKeyNameResponse, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NameID) == "" {
		return nil, errors.New("field NameID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/names/" + fmt.Sprint(req.NameID) + "/export-key",
	}

	var resp ExportKeyNameResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ImportKeyName: Import a private key.
func (s *IpnsAPI) ImportKeyName(req *IpnsAPIImportKeyNameRequest, opts ...scw.RequestOption) (*Name, error) {
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
		Path:   "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/names/import-key",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Name

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
