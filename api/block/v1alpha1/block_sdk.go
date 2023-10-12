// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package block provides methods and message types of the block v1alpha1 API.
package block

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

type ListSnapshotsRequestOrderBy string

const (
	// Order by creation date (ascending chronological order).
	ListSnapshotsRequestOrderByCreatedAtAsc = ListSnapshotsRequestOrderBy("created_at_asc")
	// Order by creation date (descending chronological order).
	ListSnapshotsRequestOrderByCreatedAtDesc = ListSnapshotsRequestOrderBy("created_at_desc")
	// Order by name (ascending order).
	ListSnapshotsRequestOrderByNameAsc = ListSnapshotsRequestOrderBy("name_asc")
	// Order by name (descending order).
	ListSnapshotsRequestOrderByNameDesc = ListSnapshotsRequestOrderBy("name_desc")
)

func (enum ListSnapshotsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListSnapshotsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSnapshotsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSnapshotsRequestOrderBy(ListSnapshotsRequestOrderBy(tmp).String())
	return nil
}

type ListVolumesRequestOrderBy string

const (
	// Order by creation date (ascending chronological order).
	ListVolumesRequestOrderByCreatedAtAsc = ListVolumesRequestOrderBy("created_at_asc")
	// Order by creation date (descending chronological order).
	ListVolumesRequestOrderByCreatedAtDesc = ListVolumesRequestOrderBy("created_at_desc")
	// Order by name (ascending order).
	ListVolumesRequestOrderByNameAsc = ListVolumesRequestOrderBy("name_asc")
	// Order by name (descending order).
	ListVolumesRequestOrderByNameDesc = ListVolumesRequestOrderBy("name_desc")
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

type ReferenceStatus string

const (
	// If unspecified, the status of the reference is unknown by default.
	ReferenceStatusUnknownStatus = ReferenceStatus("unknown_status")
	// When the reference is being attached (transient).
	ReferenceStatusAttaching = ReferenceStatus("attaching")
	// When the reference attached to a volume.
	ReferenceStatusAttached = ReferenceStatus("attached")
	// When the reference is being detached (transient).
	ReferenceStatusDetaching = ReferenceStatus("detaching")
	// When the reference is detached from a volume - the reference ceases to exist.
	ReferenceStatusDetached = ReferenceStatus("detached")
	// Reference undergoing snapshotting operation (transient).
	ReferenceStatusSnapshotting = ReferenceStatus("snapshotting")
	// Error status.
	ReferenceStatusError = ReferenceStatus("error")
)

func (enum ReferenceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum ReferenceStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ReferenceStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ReferenceStatus(ReferenceStatus(tmp).String())
	return nil
}

type ReferenceType string

const (
	// If unspecified, the reference type is unknown by default.
	ReferenceTypeUnknownType = ReferenceType("unknown_type")
	// Reference linked to a snapshot (for snapshots only).
	ReferenceTypeLink = ReferenceType("link")
	// Exclusive reference that can be associated to a volume (for volumes only).
	ReferenceTypeExclusive = ReferenceType("exclusive")
	// Access to the volume or snapshot in a read-only mode, without storage write access to the resource.
	ReferenceTypeReadOnly = ReferenceType("read_only")
)

func (enum ReferenceType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum ReferenceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ReferenceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ReferenceType(ReferenceType(tmp).String())
	return nil
}

type SnapshotStatus string

const (
	// If unspecified, the snapshot status is unknown by default.
	SnapshotStatusUnknownStatus = SnapshotStatus("unknown_status")
	// The snapshot is under creation (transient).
	SnapshotStatusCreating = SnapshotStatus("creating")
	// Snapshot exists and is not attached to any reference.
	SnapshotStatusAvailable = SnapshotStatus("available")
	// Snapshot in an error status.
	SnapshotStatusError = SnapshotStatus("error")
	// Snapshot is being deleted (transient).
	SnapshotStatusDeleting = SnapshotStatus("deleting")
	// Snapshot was deleted.
	SnapshotStatusDeleted = SnapshotStatus("deleted")
	// Snapshot attached to one or more references.
	SnapshotStatusInUse  = SnapshotStatus("in_use")
	SnapshotStatusLocked = SnapshotStatus("locked")
)

func (enum SnapshotStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum SnapshotStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SnapshotStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SnapshotStatus(SnapshotStatus(tmp).String())
	return nil
}

type StorageClass string

const (
	// If unspecified, the Storage Class is unknown by default.
	StorageClassUnknownStorageClass = StorageClass("unknown_storage_class")
	// No specific Storage Class selected.
	StorageClassUnspecified = StorageClass("unspecified")
	// Classic storage.
	StorageClassBssd = StorageClass("bssd")
	// Performance storage with lower latency.
	StorageClassSbs = StorageClass("sbs")
)

func (enum StorageClass) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_storage_class"
	}
	return string(enum)
}

func (enum StorageClass) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *StorageClass) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = StorageClass(StorageClass(tmp).String())
	return nil
}

type VolumeStatus string

const (
	// If unspecified, the volume status is unknown by default.
	VolumeStatusUnknownStatus = VolumeStatus("unknown_status")
	// The volume is under creation (transient).
	VolumeStatusCreating = VolumeStatus("creating")
	// The volume exists and is not attached to any reference.
	VolumeStatusAvailable = VolumeStatus("available")
	// The volume exists and is already attached to a reference.
	VolumeStatusInUse = VolumeStatus("in_use")
	// The volume undergoing deletion (transient).
	VolumeStatusDeleting = VolumeStatus("deleting")
	VolumeStatusDeleted  = VolumeStatus("deleted")
	// The volume is being increased (transient).
	VolumeStatusResizing = VolumeStatus("resizing")
	// The volume is an error status.
	VolumeStatusError = VolumeStatus("error")
	// The volume is undergoing snapshotting operation (transient).
	VolumeStatusSnapshotting = VolumeStatus("snapshotting")
	VolumeStatusLocked       = VolumeStatus("locked")
)

func (enum VolumeStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum VolumeStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *VolumeStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = VolumeStatus(VolumeStatus(tmp).String())
	return nil
}

// SnapshotParentVolume:
type SnapshotParentVolume struct {
	// ID: Parent volume UUID (volume from which the snapshot originates).
	ID string `json:"id"`
	// Name: Name of the parent volume.
	Name string `json:"name"`
	// Type: Volume type of the parent volume.
	Type string `json:"type"`
	// Status: Current status the parent volume.
	Status VolumeStatus `json:"status"`
}

// VolumeSpecifications:
type VolumeSpecifications struct {
	// PerfIops: The maximum IO/s expected, according to the different options available in stock (`5000 | 15000`).
	PerfIops *uint32 `json:"perf_iops"`
	// Class: The storage class of the volume.
	Class StorageClass `json:"class"`
}

// Reference:
type Reference struct {
	// ID: UUID of the reference.
	ID string `json:"id"`
	// ProductResourceType: Type of resoruce to which the reference is associated (snapshot or volume).
	ProductResourceType string `json:"product_resource_type"`
	// ProductResourceID: UUID of the volume or the snapshot it refers to (according to the product_resource_type).
	ProductResourceID string `json:"product_resource_id"`
	// CreatedAt: Creation date of the reference.
	CreatedAt *time.Time `json:"created_at"`
	// Type: Type of reference (link, exclusive, read_only).
	Type ReferenceType `json:"type"`
	// Status: Status of reference (attaching, attached, detaching).
	Status ReferenceStatus `json:"status"`
}

// CreateVolumeRequestFromEmpty:
type CreateVolumeRequestFromEmpty struct {
	// Size: Must be compliant with the minimum (1 GB) and maximum (10 TB) allowed size.
	Size scw.Size `json:"size"`
}

// CreateVolumeRequestFromSnapshot:
type CreateVolumeRequestFromSnapshot struct {
	// Size: Must be compliant with the minimum (1 GB) and maximum (10 TB) allowed size.
	// Size is optional and is used only if a resize of the volume is requested, otherwise original snapshot size will be used.
	Size *scw.Size `json:"size"`
	// SnapshotID: Source snapshot from which volume will be created.
	SnapshotID string `json:"snapshot_id"`
}

// SnapshotSummary:
type SnapshotSummary struct {
	// ID: UUID of the snapshot.
	ID string `json:"id"`
	// Name: Name of the snapshot.
	Name string `json:"name"`
	// ParentVolume: If the parent volume has been deleted, value is null.
	ParentVolume *SnapshotParentVolume `json:"parent_volume"`
	// Size: Size of the snapshot in bytes.
	Size scw.Size `json:"size"`
	// ProjectID: UUID of the project the snapshot belongs to.
	ProjectID string `json:"project_id"`
	// CreatedAt: Creation date of the snapshot.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Last modification date of the properties of a snapshot.
	UpdatedAt *time.Time `json:"updated_at"`
	// Status: Current status of the snapshot (available, in_use, ...).
	Status SnapshotStatus `json:"status"`
	// Tags: List of tags assigned to the volume.
	Tags []string `json:"tags"`
	// Zone: Snapshot Availability Zone.
	Zone scw.Zone `json:"zone"`
	// Class: Storage class of the snapshot.
	Class StorageClass `json:"class"`
}

// VolumeType:
type VolumeType struct {
	// Type: Volume type.
	Type string `json:"type"`
	// Pricing: Price of the volume billed in GB/hour.
	Pricing *scw.Money `json:"pricing"`
	// SnapshotPricing: Price of the snapshot billed in GB/hour.
	SnapshotPricing *scw.Money `json:"snapshot_pricing"`
	// Specs: Volume specifications of the volume type.
	Specs *VolumeSpecifications `json:"specs"`
}

// Volume:
type Volume struct {
	// ID: UUID of the volume.
	ID string `json:"id"`
	// Name: Name of the volume.
	Name string `json:"name"`
	// Type: Volume type.
	Type string `json:"type"`
	// Size: Volume size in bytes.
	Size scw.Size `json:"size"`
	// ProjectID: UUID of the project to which the volume belongs.
	ProjectID string `json:"project_id"`
	// CreatedAt: Creation date of the volume.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Last update of the properties of a volume.
	UpdatedAt *time.Time `json:"updated_at"`
	// References: List of the references to the volume.
	References []*Reference `json:"references"`
	// ParentSnapshotID: When a volume is created from a snapshot, is the UUID of the snapshot from which the volume has been created.
	ParentSnapshotID *string `json:"parent_snapshot_id"`
	// Status: Current status of the volume (available, in_use, ...).
	Status VolumeStatus `json:"status"`
	// Tags: List of tags assigned to the volume.
	Tags []string `json:"tags"`
	// Zone: Volume zone.
	Zone scw.Zone `json:"zone"`
	// Specs: Specifications of the volume.
	Specs *VolumeSpecifications `json:"specs"`
}

// CreateSnapshotRequest:
type CreateSnapshotRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// VolumeID: UUID of the volume to snapshot.
	VolumeID string `json:"volume_id"`
	// Name: Name of the snapshot.
	Name string `json:"name"`
	// ProjectID: UUID of the project to which the volume and the snapshot belong.
	ProjectID string `json:"project_id"`
	// Tags: List of tags assigned to the snapshot.
	Tags []string `json:"tags"`
}

// CreateVolumeRequest:
type CreateVolumeRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Name: Name of the volume.
	Name string `json:"name"`
	// PerfIops: The maximum IO/s expected, according to the different options available in stock (`5000 | 15000`).
	PerfIops *uint32 `json:"perf_iops,omitempty"`
	// ProjectID: UUID of the project the volume belongs to.
	ProjectID string `json:"project_id"`
	// FromEmpty: Specify the size of the new volume if creating a new one from scratch.
	FromEmpty *CreateVolumeRequestFromEmpty `json:"from_empty,omitempty"`
	// FromSnapshot: Specify the snapshot ID of the original snapshot.
	FromSnapshot *CreateVolumeRequestFromSnapshot `json:"from_snapshot,omitempty"`
	// Tags: List of tags assigned to the volume.
	Tags []string `json:"tags"`
}

// DeleteSnapshotRequest:
type DeleteSnapshotRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// SnapshotID: UUID of the snapshot.
	SnapshotID string `json:"-"`
}

// DeleteVolumeRequest:
type DeleteVolumeRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// VolumeID: UUID of the volume.
	VolumeID string `json:"-"`
}

// GetSnapshotRequest:
type GetSnapshotRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// SnapshotID: UUID of the snapshot.
	SnapshotID string `json:"-"`
}

// GetVolumeRequest:
type GetVolumeRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// VolumeID: UUID of the volume.
	VolumeID string `json:"-"`
}

// ImportSnapshotFromS3Request:
type ImportSnapshotFromS3Request struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Bucket:
	Bucket string `json:"bucket"`
	// Key:
	Key string `json:"key"`
	// Name:
	Name string `json:"name"`
	// ProjectID:
	ProjectID string `json:"project_id"`
	// Tags:
	Tags []string `json:"tags"`
}

// ListSnapshotsRequest:
type ListSnapshotsRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// OrderBy: Criteria to use when ordering the list.
	OrderBy ListSnapshotsRequestOrderBy `json:"-"`
	// ProjectID: Filter by Project ID.
	ProjectID *string `json:"-"`
	// Page: Page number.
	Page *int32 `json:"-"`
	// PageSize: Page size, defines how many entries are returned in one page, must be lower or equal to 100.
	PageSize *uint32 `json:"-"`
	// VolumeID: Filter snapshots by the ID of the original volume.
	VolumeID *string `json:"-"`
	// Name: Filter snapshots by their names.
	Name *string `json:"-"`
}

// ListSnapshotsResponse:
type ListSnapshotsResponse struct {
	// Snapshots: Paginated returned list of snapshots.
	Snapshots []*SnapshotSummary `json:"snapshots"`
	// TotalCount: Total number of snpashots in the project.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSnapshotsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSnapshotsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListSnapshotsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Snapshots = append(r.Snapshots, results.Snapshots...)
	r.TotalCount += uint64(len(results.Snapshots))
	return uint64(len(results.Snapshots)), nil
}

// ListVolumeTypesRequest:
type ListVolumeTypesRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// Page: Page number.
	Page *int32 `json:"-"`
	// PageSize: Page size, defines how many entries are returned in one page, must be lower or equal to 100.
	PageSize *uint32 `json:"-"`
}

// ListVolumeTypesResponse:
type ListVolumeTypesResponse struct {
	// VolumeTypes: Returns paginated list of volume-types.
	VolumeTypes []*VolumeType `json:"volume_types"`
	// TotalCount: Total number of volume-types currently available in stock.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListVolumeTypesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListVolumeTypesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListVolumeTypesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.VolumeTypes = append(r.VolumeTypes, results.VolumeTypes...)
	r.TotalCount += uint64(len(results.VolumeTypes))
	return uint64(len(results.VolumeTypes)), nil
}

// ListVolumesRequest:
type ListVolumesRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// OrderBy: Criteria to use when ordering the list.
	OrderBy ListVolumesRequestOrderBy `json:"-"`
	// ProjectID: Filter by Project ID.
	ProjectID *string `json:"-"`
	// Page: Page number.
	Page *int32 `json:"-"`
	// PageSize: Page size, defines how many entries are returned in one page, must be lower or equal to 100.
	PageSize *uint32 `json:"-"`
	// Name: Filter the return volumes by their names.
	Name *string `json:"-"`
	// ProductResourceID: Filter by a product resource ID linked to this volume (such as an Instance ID).
	ProductResourceID *string `json:"-"`
}

// ListVolumesResponse:
type ListVolumesResponse struct {
	// Volumes: Paginated returned list of volumes.
	Volumes []*Volume `json:"volumes"`
	// TotalCount: Total number of volumes in the project.
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

// Snapshot:
type Snapshot struct {
	// ID: UUID of the snapshot.
	ID string `json:"id"`
	// Name: Name of the snapshot.
	Name string `json:"name"`
	// ParentVolume: If the parent volume was deleted, value is null.
	ParentVolume *SnapshotParentVolume `json:"parent_volume"`
	// Size: Size in bytes of the snapshot.
	Size scw.Size `json:"size"`
	// ProjectID: UUID of the project the snapshot belongs to.
	ProjectID string `json:"project_id"`
	// CreatedAt: Creation date of the snapshot.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Last modification date of the properties of a snapshot.
	UpdatedAt *time.Time `json:"updated_at"`
	// References: List of the references to the snapshot.
	References []*Reference `json:"references"`
	// Status: Current status of the snapshot (available, in_use, ...).
	Status SnapshotStatus `json:"status"`
	// Tags: List of tags assigned to the volume.
	Tags []string `json:"tags"`
	// Zone: Snapshot zone.
	Zone scw.Zone `json:"zone"`
	// Class: Storage class of the snapshot.
	Class StorageClass `json:"class"`
}

// UpdateSnapshotRequest:
type UpdateSnapshotRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// SnapshotID: UUID of the snapshot.
	SnapshotID string `json:"-"`
	// Name: When defined, is the name of the snapshot.
	Name *string `json:"name,omitempty"`
	// Tags: List of tags assigned to the snapshot.
	Tags *[]string `json:"tags,omitempty"`
}

// UpdateVolumeRequest:
type UpdateVolumeRequest struct {
	// Zone:
	Zone scw.Zone `json:"-"`
	// VolumeID: UUID of the volume.
	VolumeID string `json:"-"`
	// Name: When defined, is the new name of the volume.
	Name *string `json:"name,omitempty"`
	// Size: Size in bytes of the volume, with a granularity of 1 GB (10^9 bytes).
	// Must be compliant with the minimum (1GB) and maximum (10TB) allowed size.
	Size *scw.Size `json:"size,omitempty"`
	// Tags: List of tags assigned to the volume.
	Tags *[]string `json:"tags,omitempty"`
	// PerfIops: The selected value must be available for the volume's current storage class.
	PerfIops *uint32 `json:"perf_iops,omitempty"`
}

// This API allows you to use and manage your Block Storage volumes.
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
	return []scw.Zone{scw.ZoneFrPar1, scw.ZonePlWaw3}
}

// ListVolumeTypes: List all available volume types in a specified zone. The volume types listed are ordered by name in ascending order.
func (s *API) ListVolumeTypes(req *ListVolumeTypesRequest, opts ...scw.RequestOption) (*ListVolumeTypesResponse, error) {
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

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/block/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/volume-types",
		Query:  query,
	}

	var resp ListVolumeTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVolumes: List all existing volumes in a specified zone. By default, the volumes listed are ordered by creation date in ascending order. This can be modified via the `order_by` field.
func (s *API) ListVolumes(req *ListVolumesRequest, opts ...scw.RequestOption) (*ListVolumesResponse, error) {
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "product_resource_id", req.ProductResourceID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/block/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/volumes",
		Query:  query,
	}

	var resp ListVolumesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVolume: To create a new volume from scratch, you must specify `from_empty` and the `size`.
// To create a volume from an existing snapshot, specify `from_snapshot` and the `snapshot_id` in the request payload instead, size is optional and can be specified if you need to extend the original size. The volume will take on the same volume class and underlying IOPS limitations as the original snapshot.
func (s *API) CreateVolume(req *CreateVolumeRequest, opts ...scw.RequestOption) (*Volume, error) {
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
		Path:   "/block/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/volumes",
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

// GetVolume: Retrieve technical information about a specific volume. Details such as size, type, and status are returned in the response.
func (s *API) GetVolume(req *GetVolumeRequest, opts ...scw.RequestOption) (*Volume, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.VolumeID) == "" {
		return nil, errors.New("field VolumeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/block/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/volumes/" + fmt.Sprint(req.VolumeID) + "",
	}

	var resp Volume

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVolume: You must specify the `volume_id` of the volume you want to delete. The volume must not be in the `in_use` status.
func (s *API) DeleteVolume(req *DeleteVolumeRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.VolumeID) == "" {
		return errors.New("field VolumeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/block/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/volumes/" + fmt.Sprint(req.VolumeID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdateVolume: Update the technical details of a volume, such as its name, tags, or its new size and `volume_type` (within the same Block Storage class).
// You can only resize a volume to a larger size. It is currently not possible to change your Block Storage Class.
func (s *API) UpdateVolume(req *UpdateVolumeRequest, opts ...scw.RequestOption) (*Volume, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.VolumeID) == "" {
		return nil, errors.New("field VolumeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/block/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/volumes/" + fmt.Sprint(req.VolumeID) + "",
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

// ListSnapshots: List all available snapshots in a specified zone. By default, the snapshots listed are ordered by creation date in ascending order. This can be modified via the `order_by` field.
func (s *API) ListSnapshots(req *ListSnapshotsRequest, opts ...scw.RequestOption) (*ListSnapshotsResponse, error) {
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "volume_id", req.VolumeID)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/block/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/snapshots",
		Query:  query,
	}

	var resp ListSnapshotsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSnapshot: Retrieve technical information about a specific snapshot. Details such as size, volume type, and status are returned in the response.
func (s *API) GetSnapshot(req *GetSnapshotRequest, opts ...scw.RequestOption) (*Snapshot, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SnapshotID) == "" {
		return nil, errors.New("field SnapshotID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/block/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
	}

	var resp Snapshot

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSnapshot: To create a snapshot, the volume must be in the `in_use` or the `available` status.
// If your volume is in a transient state, you need to wait until the end of the current operation.
func (s *API) CreateSnapshot(req *CreateSnapshotRequest, opts ...scw.RequestOption) (*Snapshot, error) {
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
		Path:   "/block/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/snapshots",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Snapshot

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ImportSnapshotFromS3:
func (s *API) ImportSnapshotFromS3(req *ImportSnapshotFromS3Request, opts ...scw.RequestOption) (*Snapshot, error) {
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
		Path:   "/block/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/snapshots/import-from-s3",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Snapshot

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSnapshot: You must specify the `snapshot_id` of the snapshot you want to delete. The snapshot must not be in use.
func (s *API) DeleteSnapshot(req *DeleteSnapshotRequest, opts ...scw.RequestOption) error {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SnapshotID) == "" {
		return errors.New("field SnapshotID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/block/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdateSnapshot: Update the name or tags of the snapshot.
func (s *API) UpdateSnapshot(req *UpdateSnapshotRequest, opts ...scw.RequestOption) (*Snapshot, error) {
	var err error
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SnapshotID) == "" {
		return nil, errors.New("field SnapshotID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/block/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Snapshot

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
