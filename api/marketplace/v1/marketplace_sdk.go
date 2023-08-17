// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package marketplace provides methods and message types of the marketplace v1 API.
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

type LocalImage struct {
	// ID: Version you will typically use to define an image in an API call.
	ID string `json:"id"`
	// CompatibleCommercialTypes: List of all commercial types that are compatible with this local image.
	CompatibleCommercialTypes []string `json:"compatible_commercial_types"`
	// Arch: Supported architecture for this local image.
	Arch string `json:"arch"`
	// Zone: Availability Zone where this local image is available.
	Zone scw.Zone `json:"zone"`
}

type Organization struct {
	// ID:
	ID string `json:"id"`
	// Name:
	Name string `json:"name"`
}

type Version struct {
	// ID: UUID of this version.
	ID string `json:"id"`
	// Name: Name of this version.
	Name string `json:"name"`
	// CreationDate: Creation date of this image version.
	CreationDate *time.Time `json:"creation_date,omitempty"`
	// ModificationDate: Date of the last modification of this version.
	ModificationDate *time.Time `json:"modification_date,omitempty"`
	// LocalImages: List of local images available in this version.
	LocalImages []*LocalImage `json:"local_images"`
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
	// CreationDate: Creation date of this image.
	CreationDate *time.Time `json:"creation_date,omitempty"`
	// ModificationDate: Date of the last modification of this image.
	ModificationDate *time.Time `json:"modification_date,omitempty"`
	// ValidUntil: Expiration date of this image.
	ValidUntil *time.Time `json:"valid_until,omitempty"`
	// Label: Typically an identifier for a distribution (ex. "ubuntu_focal").
	Label string `json:"label"`
	// Versions: List of versions of this image.
	Versions []*Version `json:"versions"`
	// Organization: Organization this image belongs to.
	Organization *Organization `json:"organization"`
	// CurrentPublicVersion:
	CurrentPublicVersion string `json:"current_public_version"`
}

type GetImageRequest struct {
	// ImageID: Display the image name.
	ImageID string `json:"-"`
}

type GetImageResponse struct {
	// Image:
	Image *Image `json:"image"`
}

type GetVersionRequest struct {
	// ImageID:
	ImageID string `json:"-"`
	// VersionID:
	VersionID string `json:"-"`
}

type GetVersionResponse struct {
	// Version:
	Version *Version `json:"version"`
}

type ListImagesRequest struct {
	// PerPage: A positive integer lower or equal to 100 to select the number of items to display.
	PerPage *uint32 `json:"per_page,omitempty"`
	// Page: A positive integer to choose the page to display.
	Page *int32 `json:"page,omitempty"`
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

type ListVersionsRequest struct {
	// ImageID:
	ImageID string `json:"-"`
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

// Marketplace API.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// ListImages: List marketplace images.
func (s *API) ListImages(req *ListImagesRequest, opts ...scw.RequestOption) (*ListImagesResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/marketplace/v1/images",
		Query:  query,
	}

	var resp ListImagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetImage: Get a specific marketplace image.
func (s *API) GetImage(req *GetImageRequest, opts ...scw.RequestOption) (*GetImageResponse, error) {
	var err error

	if fmt.Sprint(req.ImageID) == "" {
		return nil, errors.New("field ImageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/marketplace/v1/images/" + fmt.Sprint(req.ImageID) + "",
	}

	var resp GetImageResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVersions:
func (s *API) ListVersions(req *ListVersionsRequest, opts ...scw.RequestOption) (*ListVersionsResponse, error) {
	var err error

	if fmt.Sprint(req.ImageID) == "" {
		return nil, errors.New("field ImageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/marketplace/v1/images/" + fmt.Sprint(req.ImageID) + "/versions",
	}

	var resp ListVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVersion:
func (s *API) GetVersion(req *GetVersionRequest, opts ...scw.RequestOption) (*GetVersionResponse, error) {
	var err error

	if fmt.Sprint(req.ImageID) == "" {
		return nil, errors.New("field ImageID cannot be empty in request")
	}

	if fmt.Sprint(req.VersionID) == "" {
		return nil, errors.New("field VersionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/marketplace/v1/images/" + fmt.Sprint(req.ImageID) + "/versions/" + fmt.Sprint(req.VersionID) + "",
	}

	var resp GetVersionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
