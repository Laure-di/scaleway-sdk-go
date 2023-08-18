// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package test provides methods and message types of the test v1 API.
package test

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

type EyeColors string

const (
	// Unknown color.
	EyeColorsUnknown = EyeColors("unknown")
	// Rare and striking shade that typically features a golden or yellowish-brown hue.
	EyeColorsAmber = EyeColors("amber")
	// Relatively rare, with the highest frequency found in eastern Europe.
	EyeColorsBlue = EyeColors("blue")
	// Most common eye color in the world caused by a high concentration of melanin in the iris.
	EyeColorsBrown = EyeColors("brown")
	// Relatively rare color which can change depending on the lighting conditions.
	EyeColorsGray = EyeColors("gray")
	// Rare and unique color characterized by a combination of yellow, brown, and blue pigments.
	EyeColorsGreen = EyeColors("green")
	// Brownish-yellow or greenish-brown with a hint of gold.
	EyeColorsHazel = EyeColors("hazel")
	// Rare mutation that results in a reddish-pink hue.
	EyeColorsRed = EyeColors("red")
	// Rare and striking shade that appears to be a mix of blue and purple.
	EyeColorsViolet = EyeColors("violet")
)

func (enum EyeColors) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum EyeColors) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *EyeColors) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = EyeColors(EyeColors(tmp).String())
	return nil
}

type HumanStatus string

const (
	// Unknown status.
	HumanStatusUnknown = HumanStatus("unknown")
	// The human is stopped.
	HumanStatusStopped = HumanStatus("stopped")
	// The human is running.
	HumanStatusRunning = HumanStatus("running")
)

func (enum HumanStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum HumanStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *HumanStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = HumanStatus(HumanStatus(tmp).String())
	return nil
}

type ListHumansRequestOrderBy string

const (
	// Ascending creation date.
	ListHumansRequestOrderByCreatedAtAsc = ListHumansRequestOrderBy("created_at_asc")
	// Descending creation date.
	ListHumansRequestOrderByCreatedAtDesc = ListHumansRequestOrderBy("created_at_desc")
	// Ascending update date.
	ListHumansRequestOrderByUpdatedAtAsc = ListHumansRequestOrderBy("updated_at_asc")
	// Descending update date.
	ListHumansRequestOrderByUpdatedAtDesc = ListHumansRequestOrderBy("updated_at_desc")
	// Ascending height.
	ListHumansRequestOrderByHeightAsc = ListHumansRequestOrderBy("height_asc")
	// Descending height.
	ListHumansRequestOrderByHeightDesc = ListHumansRequestOrderBy("height_desc")
)

func (enum ListHumansRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListHumansRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListHumansRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListHumansRequestOrderBy(ListHumansRequestOrderBy(tmp).String())
	return nil
}

// Human:
type Human struct {
	// ID:
	ID string `json:"id"`
	// OrganizationID:
	OrganizationID string `json:"organization_id"`
	// CreatedAt:
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt:
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Height:
	Height float64 `json:"height"`
	// ShoeSize:
	ShoeSize float32 `json:"shoe_size"`
	// AltitudeInMeter:
	AltitudeInMeter int32 `json:"altitude_in_meter"`
	// AltitudeInMillimeter:
	AltitudeInMillimeter int64 `json:"altitude_in_millimeter"`
	// FingersCount:
	FingersCount uint32 `json:"fingers_count"`
	// HairCount:
	HairCount uint64 `json:"hair_count"`
	// IsHappy:
	IsHappy bool `json:"is_happy"`
	// EyesColor:
	EyesColor EyeColors `json:"eyes_color"`
	// Status:
	Status HumanStatus `json:"status"`
	// Name:
	Name string `json:"name"`
	// ProjectID:
	ProjectID string `json:"project_id"`
}

// CreateHumanRequest:
type CreateHumanRequest struct {
	// Height:
	Height float64 `json:"height"`
	// ShoeSize:
	ShoeSize float32 `json:"shoe_size"`
	// AltitudeInMeter:
	AltitudeInMeter int32 `json:"altitude_in_meter"`
	// AltitudeInMillimeter:
	AltitudeInMillimeter int64 `json:"altitude_in_millimeter"`
	// FingersCount:
	FingersCount uint32 `json:"fingers_count"`
	// HairCount:
	HairCount uint64 `json:"hair_count"`
	// IsHappy:
	IsHappy bool `json:"is_happy"`
	// EyesColor:
	EyesColor EyeColors `json:"eyes_color"`
	// Deprecated: OrganizationID:
	OrganizationID *string `json:"organization_id,omitempty"`
	// Name:
	Name string `json:"name"`
	// ProjectID:
	ProjectID *string `json:"project_id,omitempty"`
}

// DeleteHumanRequest:
type DeleteHumanRequest struct {
	// HumanID: UUID of the human you want to delete.
	HumanID string `json:"-"`
}

// GetHumanRequest:
type GetHumanRequest struct {
	// HumanID: UUID of the human you want to get.
	HumanID string `json:"-"`
}

// ListHumansRequest:
type ListHumansRequest struct {
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
	// OrderBy:
	OrderBy ListHumansRequestOrderBy `json:"-"`
	// OrganizationID:
	OrganizationID *string `json:"-"`
	// ProjectID:
	ProjectID *string `json:"-"`
}

// ListHumansResponse:
type ListHumansResponse struct {
	// TotalCount:
	TotalCount uint32 `json:"total_count"`
	// Humans:
	Humans []*Human `json:"humans"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListHumansResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListHumansResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListHumansResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Humans = append(r.Humans, results.Humans...)
	r.TotalCount += uint32(len(results.Humans))
	return uint32(len(results.Humans)), nil
}

// RegisterRequest:
type RegisterRequest struct {
	// Username:
	Username string `json:"username"`
}

// RegisterResponse:
type RegisterResponse struct {
	// SecretKey:
	SecretKey string `json:"secret_key"`
	// AccessKey:
	AccessKey string `json:"access_key"`
}

// RunHumanRequest:
type RunHumanRequest struct {
	// HumanID: UUID of the human you want to make run.
	HumanID string `json:"-"`
}

// SmokeHumanRequest:
type SmokeHumanRequest struct {
	// Deprecated: HumanID: UUID of the human you want to make smoking.
	HumanID *string `json:"-"`
}

// UpdateHumanRequest:
type UpdateHumanRequest struct {
	// HumanID: UUID of the human you want to update.
	HumanID string `json:"-"`
	// Height:
	Height *float64 `json:"height,omitempty"`
	// ShoeSize:
	ShoeSize *float32 `json:"shoe_size,omitempty"`
	// AltitudeInMeter:
	AltitudeInMeter *int32 `json:"altitude_in_meter,omitempty"`
	// AltitudeInMillimeter:
	AltitudeInMillimeter *int64 `json:"altitude_in_millimeter,omitempty"`
	// FingersCount:
	FingersCount *uint32 `json:"fingers_count,omitempty"`
	// HairCount:
	HairCount *uint64 `json:"hair_count,omitempty"`
	// IsHappy:
	IsHappy *bool `json:"is_happy,omitempty"`
	// EyesColor:
	EyesColor EyeColors `json:"eyes_color"`
	// Name:
	Name *string `json:"name,omitempty"`
}

// Test is a fake service that aim to manage fake humans. It is used for internal and public end-to-end tests.
//
// This service don't use the Scaleway authentication service but a fake one.
// It allows to use this test service publicly without requiring a Scaleway account.
//
// First, you need to register a user with `scw test human register` to get an access-key.
// Then, you can use other test commands by setting the SCW_SECRET_KEY env variable.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// Register: Register a human and return a access-key and a secret-key that must be used in all other commands.
//
// Hint: you can use other test commands by setting the SCW_SECRET_KEY env variable.
func (s *API) Register(req *RegisterRequest, opts ...scw.RequestOption) (*RegisterResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/test/v1/register",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RegisterResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListHumans: List all your humans.
func (s *API) ListHumans(req *ListHumansRequest, opts ...scw.RequestOption) (*ListHumansResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/test/v1/humans",
		Query:  query,
	}

	var resp ListHumansResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHuman: Get the human details associated with the given id.
func (s *API) GetHuman(req *GetHumanRequest, opts ...scw.RequestOption) (*Human, error) {
	var err error

	if fmt.Sprint(req.HumanID) == "" {
		return nil, errors.New("field HumanID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/test/v1/humans/" + fmt.Sprint(req.HumanID) + "",
	}

	var resp Human

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateHuman: Create a new human.
func (s *API) CreateHuman(req *CreateHumanRequest, opts ...scw.RequestOption) (*Human, error) {
	var err error

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.OrganizationID == nil && req.ProjectID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.OrganizationID == nil && req.ProjectID == nil {
		req.ProjectID = &defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/test/v1/humans",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Human

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHuman: Update the human associated with the given id.
func (s *API) UpdateHuman(req *UpdateHumanRequest, opts ...scw.RequestOption) (*Human, error) {
	var err error

	if fmt.Sprint(req.HumanID) == "" {
		return nil, errors.New("field HumanID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/test/v1/humans/" + fmt.Sprint(req.HumanID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Human

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteHuman: Delete the human associated with the given id.
func (s *API) DeleteHuman(req *DeleteHumanRequest, opts ...scw.RequestOption) (*Human, error) {
	var err error

	if fmt.Sprint(req.HumanID) == "" {
		return nil, errors.New("field HumanID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/test/v1/humans/" + fmt.Sprint(req.HumanID) + "",
	}

	var resp Human

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RunHuman: Start a one hour running for the given human.
func (s *API) RunHuman(req *RunHumanRequest, opts ...scw.RequestOption) (*Human, error) {
	var err error

	if fmt.Sprint(req.HumanID) == "" {
		return nil, errors.New("field HumanID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/test/v1/humans/" + fmt.Sprint(req.HumanID) + "/run",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Human

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: SmokeHuman: Make a human smoke.
func (s *API) SmokeHuman(req *SmokeHumanRequest, opts ...scw.RequestOption) (*Human, error) {
	var err error

	if fmt.Sprint(req.HumanID) == "" {
		return nil, errors.New("field HumanID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/test/v1/humans/" + fmt.Sprint(req.HumanID) + "/smoke",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Human

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
