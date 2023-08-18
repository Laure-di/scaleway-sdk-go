// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package account provides methods and message types of the account v3 API.
package account

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

type ListProjectsRequestOrderBy string

const (
	// Creation date ascending.
	ListProjectsRequestOrderByCreatedAtAsc = ListProjectsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListProjectsRequestOrderByCreatedAtDesc = ListProjectsRequestOrderBy("created_at_desc")
	// Name ascending.
	ListProjectsRequestOrderByNameAsc = ListProjectsRequestOrderBy("name_asc")
	// Name descending.
	ListProjectsRequestOrderByNameDesc = ListProjectsRequestOrderBy("name_desc")
)

func (enum ListProjectsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListProjectsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListProjectsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListProjectsRequestOrderBy(ListProjectsRequestOrderBy(tmp).String())
	return nil
}

// Project:
type Project struct {
	// ID: ID of the Project.
	ID string `json:"id"`
	// Name: Name of the Project.
	Name string `json:"name"`
	// OrganizationID: Organization ID of the Project.
	OrganizationID string `json:"organization_id"`
	// CreatedAt: Creation date of the Project.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: Update date of the Project.
	UpdatedAt *time.Time `json:"updated_at"`
	// Description: Description of the Project.
	Description string `json:"description"`
}

// ListProjectsResponse:
type ListProjectsResponse struct {
	// TotalCount: Total number of Projects.
	TotalCount uint64 `json:"total_count"`
	// Projects: Paginated returned Projects.
	Projects []*Project `json:"projects"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListProjectsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListProjectsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListProjectsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Projects = append(r.Projects, results.Projects...)
	r.TotalCount += uint64(len(results.Projects))
	return uint64(len(results.Projects)), nil
}

// ProjectAPICreateProjectRequest:
type ProjectAPICreateProjectRequest struct {
	// Name: Name of the Project.
	Name string `json:"name"`
	// OrganizationID: Organization ID of the Project.
	OrganizationID string `json:"organization_id"`
	// Description: Description of the Project.
	Description string `json:"description"`
}

// ProjectAPIDeleteProjectRequest:
type ProjectAPIDeleteProjectRequest struct {
	// ProjectID: Project ID of the Project.
	ProjectID string `json:"-"`
}

// ProjectAPIGetProjectRequest:
type ProjectAPIGetProjectRequest struct {
	// ProjectID: Project ID of the Project.
	ProjectID string `json:"-"`
}

// ProjectAPIListProjectsRequest:
type ProjectAPIListProjectsRequest struct {
	// OrganizationID: Organization ID of the Project.
	OrganizationID string `json:"-"`
	// Name: Name of the Project.
	Name *string `json:"-"`
	// Page: Page number for the returned Projects.
	Page *int32 `json:"-"`
	// PageSize: Maximum number of Project per page.
	PageSize *uint32 `json:"-"`
	// OrderBy: Sort order of the returned Projects.
	OrderBy ListProjectsRequestOrderBy `json:"-"`
	// ProjectIDs: Project IDs to filter for. The results will be limited to any Projects with an ID in this array.
	ProjectIDs []string `json:"-"`
}

// ProjectAPIUpdateProjectRequest:
type ProjectAPIUpdateProjectRequest struct {
	// ProjectID: Project ID of the Project.
	ProjectID string `json:"-"`
	// Name: Name of the Project.
	Name *string `json:"name,omitempty"`
	// Description: Description of the Project.
	Description *string `json:"description,omitempty"`
}

// The Account API currently allows you to manage **Projects**. Projects are Scaleway's resource management feature. Designed to help you manage your infrastructure and cloud services, you can create multiple Projects within a single Organization. This allows you to group resources into different Projects, providing better resource isolation and organization, which, in turn, leads to improved management efficiency.
//
// It also increases transparency for users since resources in an invoice are grouped into Projects. Additionally, access for each Scaleway product is managed at Project level: this means a user can be granted Project-specific rights without being given full access to all Projects, allowing for more targeted and controlled access. Read our [dedicated documentation](https://www.scaleway.com/en/docs/identity-and-access-management/iam/reference-content/overview/#background:-organizations-projects-and-resources) to learn more.
//
// (switchcolumn)
// <Message type="tip">
// You may be interested in the following related APIs:
// - [Billing API](https://www.scaleway.com/en/developers/api/billing/)
// - [Identity and Access Management (IAM) API](https://www.scaleway.com/en/developers/api/iam/)
// </Message>
// (switchcolumn)
//
// ## Concepts
//
// Refer to our [dedicated concepts page](https://scaleway.com/en/docs/console/my-account/concepts/) to find definitions of the different terms related to accounts.
//
// (switchcolumn)
// (switchcolumn)
//
// ## Quickstart
//
// 1. **Configure your environment variables.**
//
//	<Message type="note">
//	This is an optional step that seeks to simplify your usage of the APIs.
//	</Message>
//
//	```bash
//	export SCW_SECRET_KEY="<API secret key>"
//	```
//
// 2. Edit the **POST** request payload you will use to create a Project.
//
//	Replace the parameters in the following example:
//
//	```json
//	    '{
//	    "name": "Dis-Iz-My-Projekt",
//	    "organization_id": "b12d5c3g-c612-5674-c1e9-92627f36c5b9",
//	    "description": "This is the description of my Project",
//	    }'
//	```
//
//	| Parameter         | Description                                        |
//	| :---------------- | :------------------------------------------------- |
//	| `name`            | Name for the Project you want to create.        |
//	| `organization_id` | ID of the Organization in which to create the Project. It must be in UUID format.   |
//	| `description`     | Description for the Project you want to create.  |
//
// 3. **Create a Project**: Run the following command to create a Project.
//
//	Make sure you include the payload you edited in the previous step.
//
//	```bash
//	curl -X POST \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  -H "Content-Type: application/json" \
//	  "https://api.scaleway.com/account/v3/projects" \
//	  -d '{
//	    "name": "Dis-Iz-My-Projekt",
//	    "organization_id": "b12d5c3g-c612-5674-c1e9-92627f36c5b9",
//	    "type": "This is the description of my Project"
//	  }'
//	```
//
//	You should get a response like the following:
//
//	<Message type="note">
//	This is a response example, the UUIDs displayed are not real.
//	</Message>
//
//	```bash
//	{
//	  "id": "6170692e-7363-616c-6577-61792e636f6d",
//	  "name": "Dis-Iz-My-Projekt",
//	  "organization_id": "b12d5c3g-c612-5674-c1e9-92627f36c5b9",
//	  "created_at": "2023-03-02T22:00:28.888380Z",
//	  "updated_at": "2023-03-02T22:00:28.888380Z",
//	  "description": "This is the description of my project"
//	}
//	```
//
// 4. **List your Projects**: Run the following command to list your Projects. Replace the Organization ID in the endpoint with your own Organization ID.
//
//	```bash
//	curl -X GET \
//	  -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	  -H "Content-Type: application/json" \
//	  "https://api.scaleway.com/account/v3/projects?organization_id=b12d5c3g-c612-5674-c1e9-92627f36c5b9"
//	```
//
//	You should get a response like the following:
//
//	```bash
//	{
//	  "total_count": "2",
//	  "projects": [
//	    {
//	      "id": "0fe11800-ddbb-4b62-8906-ed08780cdddb",
//	      "name": "default",
//	      "organization_id": "b12d5c3g-c612-5674-c1e9-92627f36c5b9",
//	      "created_at": "2022-08-31T07:45:53.657735Z",
//	      "updated_at": "2022-08-31T07:45:53.657735Z",
//	      "description": "cannot_be_deleted"
//	    },
//	    {
//	        "id":"6170692e-7363-616c-6577-61792e636f6d",
//	        "name":"Dis-Iz-My-Projekt",
//	        "organization_id":"b12d5c3g-c612-5674-c1e9-92627f36c5b9",
//	        "created_at":"2023-03-02T22:00:28.888380Z",
//	        "updated_at":"2023-03-02T22:00:28.888380Z",
//	        "description":"This is the description of my Project"
//	    },
//	]
//	}
//	```
//
//  5. **Update a Project**: Run the following command to update a Project's name and description.
//     <Message type="note">
//     Do not forget to replace the Project ID in the endpoint, and the Organization ID in the payload, with your own. You can retrieve the Project ID from the "List Project" response above.
//     </Message>
//
//     ```bash
//     curl -X PATCH \
//     -H "X-Auth-Token: $SCW_SECRET_KEY" \
//     -H "Content-Type: application/json" \
//     "https://api.scaleway.com/account/v3/projects/6170692e-7363-616c-6577-61792e636f6d" \
//     -d '{
//     "name": "Dis-Iz-My-Updated-Projekt",
//     "organization_id": "b12d5c3g-c612-5674-c1e9-92627f36c5b9",
//     "description": "This is the updated description of my Project"
//     }'
//     ```
//
// (switchcolumn)
// <Message type="requirement">
// - You have a [Scaleway account](https://console.scaleway.com/) and you know your Organization ID
// - You have created an [API key](https://www.scaleway.com/en/docs/identity-and-access-management/iam/how-to/create-api-keys/) and that the API key has sufficient [IAM permissions](https://www.scaleway.com/en/docs/identity-and-access-management/iam/reference-content/permission-sets/) to perform the actions described on this page
// - You know your [Organization ID](https://console.scaleway.com/organization/settings)
// - You have [installed `curl`](https://curl.se/download.html)
// </Message>
// (switchcolumn)
//
// ## Technical limitations
//
// On the Scaleway Console, the Account scope covers account and Organization creation as well as personal data configuration and management of Projects. Currently, the public Account API only allows you to manage Projects.
//
// ## Going further
//
// For more help using Scaleway Account, check out the following resources:
// - Our [main documentation](https://www.scaleway.com/en/docs/console/my-account/)
// - Our [Slack Community](https://www.scaleway-community.slack.com)
// - Our [support ticketing system](https://www.scaleway.com/en/docs/console/my-account/how-to/open-a-support-ticket/).
type ProjectAPI struct {
	client *scw.Client
}

// NewProjectAPI returns a ProjectAPI object from a Scaleway client.
func NewProjectAPI(client *scw.Client) *ProjectAPI {
	return &ProjectAPI{
		client: client,
	}
}

// CreateProject: Generate a new Project for an Organization, specifying its configuration including name and description.
func (s *ProjectAPI) CreateProject(req *ProjectAPICreateProjectRequest, opts ...scw.RequestOption) (*Project, error) {
	var err error
	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("proj")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v3/projects",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Project

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListProjects: List all Projects of an Organization. The response will include the total number of Projects as well as their associated Organizations, names, and IDs. Other information includes the creation and update date of the Project.
func (s *ProjectAPI) ListProjects(req *ProjectAPIListProjectsRequest, opts ...scw.RequestOption) (*ListProjectsResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_ids", req.ProjectIDs)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v3/projects",
		Query:  query,
	}

	var resp ListProjectsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetProject: Retrieve information about an existing Project, specified by its Project ID. Its full details, including ID, name and description, are returned in the response object.
func (s *ProjectAPI) GetProject(req *ProjectAPIGetProjectRequest, opts ...scw.RequestOption) (*Project, error) {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.ProjectID) == "" {
		return nil, errors.New("field ProjectID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v3/projects/" + fmt.Sprint(req.ProjectID) + "",
	}

	var resp Project

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteProject: Delete an existing Project, specified by its Project ID. The Project needs to be empty (meaning there are no resources left in it) to be deleted effectively. Note that deleting a Project is permanent, and cannot be undone.
func (s *ProjectAPI) DeleteProject(req *ProjectAPIDeleteProjectRequest, opts ...scw.RequestOption) error {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.ProjectID) == "" {
		return errors.New("field ProjectID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/account/v3/projects/" + fmt.Sprint(req.ProjectID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdateProject: Update the parameters of an existing Project, specified by its Project ID. These parameters include the name and description.
func (s *ProjectAPI) UpdateProject(req *ProjectAPIUpdateProjectRequest, opts ...scw.RequestOption) (*Project, error) {
	var err error
	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.ProjectID) == "" {
		return nil, errors.New("field ProjectID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/account/v3/projects/" + fmt.Sprint(req.ProjectID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Project

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
