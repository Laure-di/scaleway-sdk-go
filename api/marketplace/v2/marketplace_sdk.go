// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package marketplace provides methods and message types of the marketplace v2 API.
package marketplace

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

type ListImagesRequestOrderBy string

const (
	ListImagesRequestOrderByNameAsc       = ListImagesRequestOrderBy("name_asc")
	ListImagesRequestOrderByNameDesc      = ListImagesRequestOrderBy("name_desc")
	ListImagesRequestOrderByCreatedAtAsc  = ListImagesRequestOrderBy("created_at_asc")
	ListImagesRequestOrderByCreatedAtDesc = ListImagesRequestOrderBy("created_at_desc")
	ListImagesRequestOrderByUpdatedAtAsc  = ListImagesRequestOrderBy("updated_at_asc")
	ListImagesRequestOrderByUpdatedAtDesc = ListImagesRequestOrderBy("updated_at_desc")
)

func (enum ListImagesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
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

type ListLocalImagesRequestOrderBy string

const (
	ListLocalImagesRequestOrderByCreatedAtAsc  = ListLocalImagesRequestOrderBy("created_at_asc")
	ListLocalImagesRequestOrderByCreatedAtDesc = ListLocalImagesRequestOrderBy("created_at_desc")
)

func (enum ListLocalImagesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListLocalImagesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListLocalImagesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListLocalImagesRequestOrderBy(ListLocalImagesRequestOrderBy(tmp).String())
	return nil
}

type ListVersionsRequestOrderBy string

const (
	ListVersionsRequestOrderByCreatedAtAsc  = ListVersionsRequestOrderBy("created_at_asc")
	ListVersionsRequestOrderByCreatedAtDesc = ListVersionsRequestOrderBy("created_at_desc")
)

func (enum ListVersionsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListVersionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListVersionsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListVersionsRequestOrderBy(ListVersionsRequestOrderBy(tmp).String())
	return nil
}

type LocalImageType string

const (
	// Unspecified image type.
	LocalImageTypeUnknownType = LocalImageType("unknown_type")
	// An image type that can be used to create volumes which are managed via the Instance API.
	LocalImageTypeInstanceLocal = LocalImageType("instance_local")
	// An image type that can be used to create volumes which are managed via the Scaleway Block Storage (SBS) API.
	LocalImageTypeInstanceSbs = LocalImageType("instance_sbs")
)

func (enum LocalImageType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum LocalImageType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *LocalImageType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = LocalImageType(LocalImageType(tmp).String())
	return nil
}

type Category struct {
	// ID:
	ID string `json:"id"`
	// Name:
	Name string `json:"name"`
	// Description:
	Description string `json:"description"`
}

type Image struct {
	// ID: UUID of this image.
	ID string `json:"id"`
	// Name: Name of the image.
	Name string `json:"name"`
	// Description: Text description of this image.
	Description string `json:"description"`
	// Logo: URL of this image's logo.
	Logo string `json:"logo"`
	// Categories: List of categories this image belongs to.
	Categories []string `json:"categories"`
	// CreatedAt: Creation date of this image.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Date of the last modification of this image.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// ValidUntil: Expiration date of this image.
	ValidUntil *time.Time `json:"valid_until,omitempty"`
	// Label: Typically an identifier for a distribution (ex. "ubuntu_focal").
	Label string `json:"label"`
}

type LocalImage struct {
	// ID: Version you will typically use to define an image in an API call.
	ID string `json:"id"`
	// CompatibleCommercialTypes: List of all commercial types that are compatible with this local image.
	CompatibleCommercialTypes []string `json:"compatible_commercial_types"`
	// Arch: Supported architecture for this local image.
	Arch string `json:"arch"`
	// Zone: Availability Zone where this local image is available.
	Zone scw.Zone `json:"zone"`
	// Label: Image label this image belongs to.
	Label string `json:"label"`
	// Type: Type of this local image.
	Type LocalImageType `json:"type"`
}

type Version struct {
	// ID: UUID of this version.
	ID string `json:"id"`
	// Name: Name of this version.
	Name string `json:"name"`
	// CreatedAt: Creation date of this image version.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Date of the last modification of this version.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// PublishedAt: Date this version was officially published.
	PublishedAt *time.Time `json:"published_at,omitempty"`
}

type GetCategoryRequest struct {
	// CategoryID:
	CategoryID string `json:"-"`
}

type GetImageRequest struct {
	// ImageID: Display the image name.
	ImageID string `json:"-"`
}

type GetLocalImageRequest struct {
	// LocalImageID:
	LocalImageID string `json:"-"`
}

type GetVersionRequest struct {
	// VersionID:
	VersionID string `json:"-"`
}

type ListCategoriesRequest struct {
	// PageSize:
	PageSize *uint32 `json:"page_size,omitempty"`
	// Page:
	Page *int32 `json:"page,omitempty"`
}

type ListCategoriesResponse struct {
	// Categories:
	Categories []*Category `json:"categories"`
	// TotalCount:
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCategoriesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCategoriesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListCategoriesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Categories = append(r.Categories, results.Categories...)
	r.TotalCount += uint32(len(results.Categories))
	return uint32(len(results.Categories)), nil
}

type ListImagesRequest struct {
	// PageSize: A positive integer lower or equal to 100 to select the number of items to display.
	PageSize *uint32 `json:"page_size,omitempty"`
	// Page: A positive integer to choose the page to display.
	Page *int32 `json:"page,omitempty"`
	// OrderBy: Ordering to use.
	OrderBy ListImagesRequestOrderBy `json:"order_by"`
	// Arch: Choose for which machine architecture to return images.
	Arch *string `json:"arch,omitempty"`
	// Category: Choose the category of images to get.
	Category *string `json:"category,omitempty"`
	// IncludeEol: Choose to include end-of-life images.
	IncludeEol bool `json:"include_eol"`
}

type ListImagesResponse struct {
	// Images:
	Images []*Image `json:"images"`
	// TotalCount:
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

type ListLocalImagesRequest struct {
	// ImageID:
	ImageID *string `json:"image_id,omitempty"`
	// VersionID:
	VersionID *string `json:"version_id,omitempty"`
	// PageSize:
	PageSize *uint32 `json:"page_size,omitempty"`
	// Page:
	Page *int32 `json:"page,omitempty"`
	// OrderBy:
	OrderBy ListLocalImagesRequestOrderBy `json:"order_by"`
	// ImageLabel:
	ImageLabel *string `json:"image_label,omitempty"`
	// Zone:
	Zone *scw.Zone `json:"zone,omitempty"`
	// Type:
	Type LocalImageType `json:"type"`
}

type ListLocalImagesResponse struct {
	// LocalImages:
	LocalImages []*LocalImage `json:"local_images"`
	// TotalCount:
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListLocalImagesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListLocalImagesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListLocalImagesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.LocalImages = append(r.LocalImages, results.LocalImages...)
	r.TotalCount += uint32(len(results.LocalImages))
	return uint32(len(results.LocalImages)), nil
}

type ListVersionsRequest struct {
	// ImageID:
	ImageID string `json:"image_id"`
	// PageSize:
	PageSize *uint32 `json:"page_size,omitempty"`
	// Page:
	Page *int32 `json:"page,omitempty"`
	// OrderBy:
	OrderBy ListVersionsRequestOrderBy `json:"order_by"`
}

type ListVersionsResponse struct {
	// Versions:
	Versions []*Version `json:"versions"`
	// TotalCount:
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListVersionsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListVersionsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListVersionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Versions = append(r.Versions, results.Versions...)
	r.TotalCount += uint32(len(results.Versions))
	return uint32(len(results.Versions)), nil
}

// The Marketplace API serves as a catalog of available Instance images. Its goal is to help find the specific image ID to use when launching an Instance.
//
// Users can launch Instances through the Instances API. These Instances use a predefined disk state, called an image. To launch an Instance with a specific image, you must know the image's ID to pass to the Instance API.
//
// These IDs can change over time as the Scaleway team releases new versions of the images, so the Marketplace API can be used to find the most up-to-date image to use.
//
// (switchcolumn)
// <Message type="note">
// Check out the [Instance API](https://www.scaleway.com/en/developers/api/instance/)
// </Message>
// (switchcolumn)
//
// ## Concepts
//
// Refer to [the Instances concepts page](https://www.scaleway.com/en/docs/compute/instances/concepts/) to learn more about images and other concepts related to Instances.
//
// (switchcolumn)
// (switchcolumn)
//
// ## Quickstart
//
// This API uses calls to the `/local-images?image_label=<label>` endpoint to get the IDs of specific images.
//
// For example, make the following call to the endpoint `/local-images?image_label=ubuntu_jammy&zone=fr-par-1` to get the local images for **Ubuntu 22.04 LTS (Jammy Jellyfish)** in the `fr-par-1` Availability Zone:
//
//	```bash
//	$ curl -X GET https://api.scaleway.com/marketplace/v2/local-images\?image_label\=ubuntu_jammy\&zone\=fr-par-1 |jq
//	```
//
// The following list of images is returned:
//
//	```bash
//	{
//	  "local_images": [
//	    {
//	      "id": "350a06b2-ecd0-48b9-ba9e-495a51345a68",
//	      "arch": "x86_64",
//	      "zone": "fr-par-1",
//	      "compatible_commercial_types": [
//	        [...]
//	        "STARDUST1-S",
//	        "PLAY2-MICRO",
//	        "PLAY2-NANO",
//	        "PLAY2-PICO"
//	      ],
//	      "label": "ubuntu_jammy"
//	    },
//	    {
//	      "id": "94299452-12c4-4d79-9c79-ced52967d75f",
//	      "arch": "arm64",
//	      "zone": "fr-par-1",
//	      "compatible_commercial_types": [
//	        "AMP2-C1",
//	        "AMP2-C2",
//	        [...]
//	      ],
//	      "label": "ubuntu_jammy"
//	    }
//	  ],
//	  "total_count": 2
//	}
//	```
//
// (switchcolumn)
// <Message type="requirement">
// - You have [installed `curl`](https://curl.se/download.html)
// - You have [installed `jq`](https://stedolan.github.io/jq/download/) (this will make it significantly easier to read the information returned by the Marketplace API)
// </Message>
// (switchcolumn)
//
// ## Technical information
//
// ### Availability Zones
//
// Scaleway's infrastructure is spread across different [regions and Availability Zones](https://www.scaleway.com/en/docs/console/my-account/reference-content/products-availability/).
//
// The Marketplace API is a global API and provides information about Instance images in all Availability Zones.
//
// ### Authentication
//
// No authentication or API key is required to use the Marketplace API.
//
// ### Pagination
//
// Most listing requests receive a paginated response. Requests against paginated endpoints accept two `query` arguments:
//
// - `page`, a positive integer to choose which page to return. The default value is `1`.
// - `page_size`, an positive integer lower or equal to 100 to select the number of items to return per page. The default value is `20`.
//
// Paginated endpoints usually also accept filters to search and sort results. These filters are documented along each endpoint documentation.
//
// ## Going further
//
// For more help using the Marketplace API, check out the following resources:
//
// * Our [Instances API documentation](https://www.scaleway.com/en/developers/api/instance/).
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// ListImages: List all available images on the marketplace, their UUID, CPU architecture and description.
func (s *API) ListImages(req *ListImagesRequest, opts ...scw.RequestOption) (*ListImagesResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "arch", req.Arch)
	parameter.AddToQuery(query, "category", req.Category)
	parameter.AddToQuery(query, "include_eol", req.IncludeEol)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/marketplace/v2/images",
		Query:  query,
	}

	var resp ListImagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetImage: Get detailed information about a marketplace image, specified by its `image_id` (UUID format).
func (s *API) GetImage(req *GetImageRequest, opts ...scw.RequestOption) (*Image, error) {
	var err error

	if fmt.Sprint(req.ImageID) == "" {
		return nil, errors.New("field ImageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/marketplace/v2/images/" + fmt.Sprint(req.ImageID) + "",
	}

	var resp Image

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVersions: Get a list of all available version of an image, specified by its `image_id` (UUID format).
func (s *API) ListVersions(req *ListVersionsRequest, opts ...scw.RequestOption) (*ListVersionsResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "image_id", req.ImageID)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/marketplace/v2/versions",
		Query:  query,
	}

	var resp ListVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVersion: Get information such as the name, creation date, last update and published date for an image version specified by its `version_id` (UUID format).
func (s *API) GetVersion(req *GetVersionRequest, opts ...scw.RequestOption) (*Version, error) {
	var err error

	if fmt.Sprint(req.VersionID) == "" {
		return nil, errors.New("field VersionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/marketplace/v2/versions/" + fmt.Sprint(req.VersionID) + "",
	}

	var resp Version

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListLocalImages: List information about local images in a specific Availability Zone, specified by its `image_id` (UUID format), `version_id` (UUID format) or `image_label`. Only one of these three parameters may be set.
func (s *API) ListLocalImages(req *ListLocalImagesRequest, opts ...scw.RequestOption) (*ListLocalImagesResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "zone", req.Zone)
	parameter.AddToQuery(query, "type", req.Type)
	parameter.AddToQuery(query, "image_id", req.ImageID)
	parameter.AddToQuery(query, "version_id", req.VersionID)
	parameter.AddToQuery(query, "image_label", req.ImageLabel)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/marketplace/v2/local-images",
		Query:  query,
	}

	var resp ListLocalImagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLocalImage: Get detailed information about a local image, including compatible commercial types, supported architecture, labels and the Availability Zone of the image, specified by its `local_image_id` (UUID format).
func (s *API) GetLocalImage(req *GetLocalImageRequest, opts ...scw.RequestOption) (*LocalImage, error) {
	var err error

	if fmt.Sprint(req.LocalImageID) == "" {
		return nil, errors.New("field LocalImageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/marketplace/v2/local-images/" + fmt.Sprint(req.LocalImageID) + "",
	}

	var resp LocalImage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListCategories: Get a list of all existing categories. The output can be paginated.
func (s *API) ListCategories(req *ListCategoriesRequest, opts ...scw.RequestOption) (*ListCategoriesResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/marketplace/v2/categories",
		Query:  query,
	}

	var resp ListCategoriesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCategory: Get information about a specific category of the marketplace catalog, specified by its `category_id` (UUID format).
func (s *API) GetCategory(req *GetCategoryRequest, opts ...scw.RequestOption) (*Category, error) {
	var err error

	if fmt.Sprint(req.CategoryID) == "" {
		return nil, errors.New("field CategoryID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/marketplace/v2/categories/" + fmt.Sprint(req.CategoryID) + "",
	}

	var resp Category

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
