// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package registry provides methods and message types of the registry v1 API.
package registry

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

type ImageStatus string

const (
	ImageStatusUnknown  = ImageStatus("unknown")
	ImageStatusReady    = ImageStatus("ready")
	ImageStatusDeleting = ImageStatus("deleting")
	ImageStatusError    = ImageStatus("error")
	ImageStatusLocked   = ImageStatus("locked")
)

func (enum ImageStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum ImageStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ImageStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ImageStatus(ImageStatus(tmp).String())
	return nil
}

type ImageVisibility string

const (
	ImageVisibilityVisibilityUnknown = ImageVisibility("visibility_unknown")
	ImageVisibilityInherit           = ImageVisibility("inherit")
	ImageVisibilityPublic            = ImageVisibility("public")
	ImageVisibilityPrivate           = ImageVisibility("private")
)

func (enum ImageVisibility) String() string {
	if enum == "" {
		// return default value if empty
		return "visibility_unknown"
	}
	return string(enum)
}

func (enum ImageVisibility) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ImageVisibility) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ImageVisibility(ImageVisibility(tmp).String())
	return nil
}

type ListImagesRequestOrderBy string

const (
	ListImagesRequestOrderByCreatedAtAsc  = ListImagesRequestOrderBy("created_at_asc")
	ListImagesRequestOrderByCreatedAtDesc = ListImagesRequestOrderBy("created_at_desc")
	ListImagesRequestOrderByNameAsc       = ListImagesRequestOrderBy("name_asc")
	ListImagesRequestOrderByNameDesc      = ListImagesRequestOrderBy("name_desc")
)

func (enum ListImagesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListImagesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListImagesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListImagesRequestOrderBy(ListImagesRequestOrderBy(tmp).String())
	return nil
}

type ListNamespacesRequestOrderBy string

const (
	ListNamespacesRequestOrderByCreatedAtAsc    = ListNamespacesRequestOrderBy("created_at_asc")
	ListNamespacesRequestOrderByCreatedAtDesc   = ListNamespacesRequestOrderBy("created_at_desc")
	ListNamespacesRequestOrderByDescriptionAsc  = ListNamespacesRequestOrderBy("description_asc")
	ListNamespacesRequestOrderByDescriptionDesc = ListNamespacesRequestOrderBy("description_desc")
	ListNamespacesRequestOrderByNameAsc         = ListNamespacesRequestOrderBy("name_asc")
	ListNamespacesRequestOrderByNameDesc        = ListNamespacesRequestOrderBy("name_desc")
)

func (enum ListNamespacesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListNamespacesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListNamespacesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListNamespacesRequestOrderBy(ListNamespacesRequestOrderBy(tmp).String())
	return nil
}

type ListTagsRequestOrderBy string

const (
	ListTagsRequestOrderByCreatedAtAsc  = ListTagsRequestOrderBy("created_at_asc")
	ListTagsRequestOrderByCreatedAtDesc = ListTagsRequestOrderBy("created_at_desc")
	ListTagsRequestOrderByNameAsc       = ListTagsRequestOrderBy("name_asc")
	ListTagsRequestOrderByNameDesc      = ListTagsRequestOrderBy("name_desc")
)

func (enum ListTagsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListTagsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListTagsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListTagsRequestOrderBy(ListTagsRequestOrderBy(tmp).String())
	return nil
}

type NamespaceStatus string

const (
	NamespaceStatusUnknown  = NamespaceStatus("unknown")
	NamespaceStatusReady    = NamespaceStatus("ready")
	NamespaceStatusDeleting = NamespaceStatus("deleting")
	NamespaceStatusError    = NamespaceStatus("error")
	NamespaceStatusLocked   = NamespaceStatus("locked")
)

func (enum NamespaceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum NamespaceStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NamespaceStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NamespaceStatus(NamespaceStatus(tmp).String())
	return nil
}

type TagStatus string

const (
	TagStatusUnknown  = TagStatus("unknown")
	TagStatusReady    = TagStatus("ready")
	TagStatusDeleting = TagStatus("deleting")
	TagStatusError    = TagStatus("error")
	TagStatusLocked   = TagStatus("locked")
)

func (enum TagStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum TagStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TagStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TagStatus(TagStatus(tmp).String())
	return nil
}

// Image:
type Image struct {
	// ID: UUID of the image.
	ID string `json:"id"`
	// Name: Name of the image, it must be unique within the namespace.
	Name string `json:"name"`
	// NamespaceID: UUID of the namespace the image belongs to.
	NamespaceID string `json:"namespace_id"`
	// Status: Status of the image.
	Status ImageStatus `json:"status"`
	// StatusMessage: Details of the image status.
	StatusMessage *string `json:"status_message"`
	// Visibility: Set to `public` to allow the image to be pulled without authentication. Else, set to  `private`. Set to `inherit` to keep the same visibility configuration as the namespace.
	Visibility ImageVisibility `json:"visibility"`
	// Size: Image size in bytes, calculated from the size of image layers. One layer used in two tags of the same image is counted once but one layer used in two images is counted twice.
	Size scw.Size `json:"size"`
	// CreatedAt: Date and time of image creation.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Date and time of last update.
	UpdatedAt *time.Time `json:"updated_at"`
	// Tags: List of docker tags of the image.
	Tags []string `json:"tags"`
}

// Namespace:
type Namespace struct {
	// ID: UUID of the namespace.
	ID string `json:"id"`
	// Name: Name of the namespace, unique in a region accross all organizations.
	Name string `json:"name"`
	// Description: Description of the namespace.
	Description string `json:"description"`
	// OrganizationID: Owner of the namespace.
	OrganizationID string `json:"organization_id"`
	// ProjectID: Project of the namespace.
	ProjectID string `json:"project_id"`
	// Status: Namespace status.
	Status NamespaceStatus `json:"status"`
	// StatusMessage: Namespace status details.
	StatusMessage string `json:"status_message"`
	// Endpoint: Endpoint reachable by docker.
	Endpoint string `json:"endpoint"`
	// IsPublic: Defines whether or not namespace is public.
	IsPublic bool `json:"is_public"`
	// Size: Total size of the namespace, calculated as the sum of the size of all images in the namespace.
	Size scw.Size `json:"size"`
	// CreatedAt: Date and time of creation.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Date and time of last update.
	UpdatedAt *time.Time `json:"updated_at"`
	// ImageCount: Number of images in the namespace.
	ImageCount uint32 `json:"image_count"`
	// Region: Region the namespace belongs to.
	Region scw.Region `json:"region"`
}

// Tag:
type Tag struct {
	// ID: UUID of the tag.
	ID string `json:"id"`
	// Name: Tag name, unique to an image.
	Name string `json:"name"`
	// ImageID: Image ID the of the image the tag belongs to.
	ImageID string `json:"image_id"`
	// Status: Tag status.
	Status TagStatus `json:"status"`
	// Digest: Hash of the tag content. Several tags of a same image may have the same digest.
	Digest string `json:"digest"`
	// CreatedAt: Date and time of creation.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Date and time of last update.
	UpdatedAt *time.Time `json:"updated_at"`
}

// CreateNamespaceRequest:
type CreateNamespaceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Name: Name of the namespace.
	Name string `json:"name"`
	// Description: Description of the namespace.
	Description string `json:"description"`
	// Deprecated: OrganizationID: Namespace owner (deprecated).
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: Project ID on which the namespace will be created.
	ProjectID *string `json:"project_id,omitempty"`
	// IsPublic: Defines whether or not namespace is public.
	IsPublic bool `json:"is_public"`
}

// DeleteImageRequest:
type DeleteImageRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ImageID: UUID of the image.
	ImageID string `json:"-"`
}

// DeleteNamespaceRequest:
type DeleteNamespaceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NamespaceID: UUID of the namespace.
	NamespaceID string `json:"-"`
}

// DeleteTagRequest:
type DeleteTagRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// TagID: UUID of the tag.
	TagID string `json:"-"`
	// Deprecated: Force: If two tags share the same digest the deletion will fail unless this parameter is set to true (deprecated).
	Force *bool `json:"-"`
}

// GetImageRequest:
type GetImageRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ImageID: UUID of the image.
	ImageID string `json:"-"`
}

// GetNamespaceRequest:
type GetNamespaceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NamespaceID: UUID of the namespace.
	NamespaceID string `json:"-"`
}

// GetTagRequest:
type GetTagRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// TagID: UUID of the tag.
	TagID string `json:"-"`
}

// ListImagesRequest:
type ListImagesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Page: A positive integer to choose the page to display.
	Page *int32 `json:"-"`
	// PageSize: A positive integer lower or equal to 100 to select the number of items to display.
	PageSize *uint32 `json:"-"`
	// OrderBy: Criteria to use when ordering image listings. Possible values are `created_at_asc`, `created_at_desc`, `name_asc`, `name_desc`, `region`, `status_asc` and `status_desc`. The default value is `created_at_asc`.
	OrderBy ListImagesRequestOrderBy `json:"-"`
	// NamespaceID: Filter by the namespace ID.
	NamespaceID *string `json:"-"`
	// Name: Filter by the image name (exact match).
	Name *string `json:"-"`
	// OrganizationID: Filter by Organization ID.
	OrganizationID *string `json:"-"`
	// ProjectID: Filter by Project ID.
	ProjectID *string `json:"-"`
}

// ListImagesResponse:
type ListImagesResponse struct {
	// Images: Paginated list of images that match the selected filters.
	Images []*Image `json:"images"`
	// TotalCount: Total number of images that match the selected filters.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListImagesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListImagesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListImagesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Images = append(r.Images, results.Images...)
	r.TotalCount += uint32(len(results.Images))
	return uint32(len(results.Images)), nil
}

// ListNamespacesRequest:
type ListNamespacesRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Page: A positive integer to choose the page to display.
	Page *int32 `json:"-"`
	// PageSize: A positive integer lower or equal to 100 to select the number of items to display.
	PageSize *uint32 `json:"-"`
	// OrderBy: Criteria to use when ordering namespace listings. Possible values are `created_at_asc`, `created_at_desc`, `name_asc`, `name_desc`, `region`, `status_asc` and `status_desc`. The default value is `created_at_asc`.
	OrderBy ListNamespacesRequestOrderBy `json:"-"`
	// OrganizationID: Filter by Organization ID.
	OrganizationID *string `json:"-"`
	// ProjectID: Filter by Project ID.
	ProjectID *string `json:"-"`
	// Name: Filter by the namespace name (exact match).
	Name *string `json:"-"`
}

// ListNamespacesResponse:
type ListNamespacesResponse struct {
	// Namespaces: Paginated list of namespaces that match the selected filters.
	Namespaces []*Namespace `json:"namespaces"`
	// TotalCount: Total number of namespaces that match the selected filters.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNamespacesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNamespacesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListNamespacesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Namespaces = append(r.Namespaces, results.Namespaces...)
	r.TotalCount += uint32(len(results.Namespaces))
	return uint32(len(results.Namespaces)), nil
}

// ListTagsRequest:
type ListTagsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ImageID: UUID of the image.
	ImageID string `json:"-"`
	// Page: A positive integer to choose the page to display.
	Page *int32 `json:"-"`
	// PageSize: A positive integer lower or equal to 100 to select the number of items to display.
	PageSize *uint32 `json:"-"`
	// OrderBy: Criteria to use when ordering tag listings. Possible values are `created_at_asc`, `created_at_desc`, `name_asc`, `name_desc`, `region`, `status_asc` and `status_desc`. The default value is `created_at_asc`.
	OrderBy ListTagsRequestOrderBy `json:"-"`
	// Name: Filter by the tag name (exact match).
	Name *string `json:"-"`
}

// ListTagsResponse:
type ListTagsResponse struct {
	// Tags: Paginated list of tags that match the selected filters.
	Tags []*Tag `json:"tags"`
	// TotalCount: Total number of tags that match the selected filters.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTagsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTagsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListTagsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Tags = append(r.Tags, results.Tags...)
	r.TotalCount += uint32(len(results.Tags))
	return uint32(len(results.Tags)), nil
}

// UpdateImageRequest:
type UpdateImageRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ImageID: ID of the image to update.
	ImageID string `json:"-"`
	// Visibility: Set to `public` to allow the image to be pulled without authentication. Else, set to  `private`. Set to `inherit` to keep the same visibility configuration as the namespace.
	Visibility ImageVisibility `json:"visibility"`
}

// UpdateNamespaceRequest:
type UpdateNamespaceRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// NamespaceID: ID of the namespace to update.
	NamespaceID string `json:"-"`
	// Description: Namespace description.
	Description *string `json:"description,omitempty"`
	// IsPublic: Defines whether or not the namespace is public.
	IsPublic *bool `json:"is_public,omitempty"`
}

// Container Registry API.
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

// ListNamespaces: List all namespaces in a specified region. By default, the namespaces listed are ordered by creation date in ascending order. This can be modified via the order_by field. You can also define additional parameters for your query, such as the `instance_id` and `project_id` parameters.
func (s *API) ListNamespaces(req *ListNamespacesRequest, opts ...scw.RequestOption) (*ListNamespacesResponse, error) {
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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/registry/v1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
		Query:  query,
	}

	var resp ListNamespacesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNamespace: Retrieve information about a given namespace, specified by its `namespace_id` and region. Full details about the namespace, such as `description`, `project_id`, `status`, `endpoint`, `is_public`, `size`, and `image_count` are returned in the response.
func (s *API) GetNamespace(req *GetNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NamespaceID) == "" {
		return nil, errors.New("field NamespaceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/registry/v1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateNamespace: Create a new Container Registry namespace. You must specify the namespace name and region in which you want it to be created. Optionally, you can specify the `project_id` and `is_public` in the request payload.
func (s *API) CreateNamespace(req *CreateNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
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

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("ns")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/registry/v1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateNamespace: Update the parameters of a given namespace, specified by its `namespace_id` and `region`. You can update the `description` and `is_public` parameters.
func (s *API) UpdateNamespace(req *UpdateNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NamespaceID) == "" {
		return nil, errors.New("field NamespaceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/registry/v1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteNamespace: Delete a given namespace. You must specify, in the endpoint, the `region` and `namespace_id` parameters of the namespace you want to delete.
func (s *API) DeleteNamespace(req *DeleteNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NamespaceID) == "" {
		return nil, errors.New("field NamespaceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/registry/v1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListImages: List all images in a specified region. By default, the images listed are ordered by creation date in ascending order. This can be modified via the order_by field. You can also define additional parameters for your query, such as the `namespace_id` and `project_id` parameters.
func (s *API) ListImages(req *ListImagesRequest, opts ...scw.RequestOption) (*ListImagesResponse, error) {
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
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/registry/v1/regions/" + fmt.Sprint(req.Region) + "/images",
		Query:  query,
	}

	var resp ListImagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetImage: Retrieve information about a given container image, specified by its `image_id` and region. Full details about the image, such as `name`, `namespace_id`, `status`, `visibility`, and `size` are returned in the response.
func (s *API) GetImage(req *GetImageRequest, opts ...scw.RequestOption) (*Image, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ImageID) == "" {
		return nil, errors.New("field ImageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/registry/v1/regions/" + fmt.Sprint(req.Region) + "/images/" + fmt.Sprint(req.ImageID) + "",
	}

	var resp Image

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateImage: Update the parameters of a given image, specified by its `image_id` and `region`. You can update the `visibility` parameter.
func (s *API) UpdateImage(req *UpdateImageRequest, opts ...scw.RequestOption) (*Image, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ImageID) == "" {
		return nil, errors.New("field ImageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/registry/v1/regions/" + fmt.Sprint(req.Region) + "/images/" + fmt.Sprint(req.ImageID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Image

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteImage: Delete a given image. You must specify, in the endpoint, the `region` and `image_id` parameters of the image you want to delete.
func (s *API) DeleteImage(req *DeleteImageRequest, opts ...scw.RequestOption) (*Image, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ImageID) == "" {
		return nil, errors.New("field ImageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/registry/v1/regions/" + fmt.Sprint(req.Region) + "/images/" + fmt.Sprint(req.ImageID) + "",
	}

	var resp Image

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListTags: List all tags for a given image, specified by region. By default, the tags listed are ordered by creation date in ascending order. This can be modified via the order_by field. You can also define additional parameters for your query, such as the `name`.
func (s *API) ListTags(req *ListTagsRequest, opts ...scw.RequestOption) (*ListTagsResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ImageID) == "" {
		return nil, errors.New("field ImageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/registry/v1/regions/" + fmt.Sprint(req.Region) + "/images/" + fmt.Sprint(req.ImageID) + "/tags",
		Query:  query,
	}

	var resp ListTagsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTag: Retrieve information about a given image tag, specified by its `tag_id` and region. Full details about the tag, such as `name`, `image_id`, `status`, and `digest` are returned in the response.
func (s *API) GetTag(req *GetTagRequest, opts ...scw.RequestOption) (*Tag, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TagID) == "" {
		return nil, errors.New("field TagID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/registry/v1/regions/" + fmt.Sprint(req.Region) + "/tags/" + fmt.Sprint(req.TagID) + "",
	}

	var resp Tag

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteTag: Delete a given image tag. You must specify, in the endpoint, the `region` and `tag_id` parameters of the tag you want to delete.
func (s *API) DeleteTag(req *DeleteTagRequest, opts ...scw.RequestOption) (*Tag, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "force", req.Force)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TagID) == "" {
		return nil, errors.New("field TagID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/registry/v1/regions/" + fmt.Sprint(req.Region) + "/tags/" + fmt.Sprint(req.TagID) + "",
		Query:  query,
	}

	var resp Tag

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
