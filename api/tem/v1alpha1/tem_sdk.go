// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package tem provides methods and message types of the tem v1alpha1 API.
package tem

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

type DomainLastStatusRecordStatus string

const (
	// If unspecified, the status of the domain's record is unknown by default.
	DomainLastStatusRecordStatusUnknownRecordStatus = DomainLastStatusRecordStatus("unknown_record_status")
	// The record is valid.
	DomainLastStatusRecordStatusValid = DomainLastStatusRecordStatus("valid")
	// The record is invalid.
	DomainLastStatusRecordStatusInvalid = DomainLastStatusRecordStatus("invalid")
	// The record was not found.
	DomainLastStatusRecordStatusNotFound = DomainLastStatusRecordStatus("not_found")
)

func (enum DomainLastStatusRecordStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_record_status"
	}
	return string(enum)
}

func (enum DomainLastStatusRecordStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainLastStatusRecordStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainLastStatusRecordStatus(DomainLastStatusRecordStatus(tmp).String())
	return nil
}

type DomainStatus string

const (
	// If unspecified, the status of the domain is unknown by default.
	DomainStatusUnknown = DomainStatus("unknown")
	// The domain is checked.
	DomainStatusChecked = DomainStatus("checked")
	// The domain is unchecked.
	DomainStatusUnchecked = DomainStatus("unchecked")
	// The domain is invalid.
	DomainStatusInvalid = DomainStatus("invalid")
	// The domain is locked.
	DomainStatusLocked = DomainStatus("locked")
	// The domain is revoked.
	DomainStatusRevoked = DomainStatus("revoked")
	// The domain is pending for check.
	DomainStatusPending = DomainStatus("pending")
)

func (enum DomainStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum DomainStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainStatus(DomainStatus(tmp).String())
	return nil
}

type EmailFlag string

const (
	// If unspecified, the flag type is unknown by default.
	EmailFlagUnknownFlag = EmailFlag("unknown_flag")
	// Refers to a non critical error received while sending the email(s). Soft bounced emails are retried.
	EmailFlagSoftBounce = EmailFlag("soft_bounce")
	// Refers to a critical error that happened while sending the email(s).
	EmailFlagHardBounce = EmailFlag("hard_bounce")
	// Refers to an email considered as spam.
	EmailFlagSpam = EmailFlag("spam")
	// Refers to an email not delivered because the recipient mailbox is full.
	EmailFlagMailboxFull = EmailFlag("mailbox_full")
)

func (enum EmailFlag) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_flag"
	}
	return string(enum)
}

func (enum EmailFlag) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *EmailFlag) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = EmailFlag(EmailFlag(tmp).String())
	return nil
}

type EmailRcptType string

const (
	// If unspecified, the recipient type is unknown by default.
	EmailRcptTypeUnknownRcptType = EmailRcptType("unknown_rcpt_type")
	// Primary recipient.
	EmailRcptTypeTo = EmailRcptType("to")
	// Carbon copy recipient.
	EmailRcptTypeCc = EmailRcptType("cc")
	// Blind carbon copy recipient.
	EmailRcptTypeBcc = EmailRcptType("bcc")
)

func (enum EmailRcptType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_rcpt_type"
	}
	return string(enum)
}

func (enum EmailRcptType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *EmailRcptType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = EmailRcptType(EmailRcptType(tmp).String())
	return nil
}

type EmailStatus string

const (
	// If unspecified, the status of the email is unknown by default.
	EmailStatusUnknown = EmailStatus("unknown")
	// The email is new.
	EmailStatusNew = EmailStatus("new")
	// The email is in the process of being sent.
	EmailStatusSending = EmailStatus("sending")
	// The email was sent.
	EmailStatusSent = EmailStatus("sent")
	// The sending of the email failed.
	EmailStatusFailed = EmailStatus("failed")
	// The sending of the email was canceled.
	EmailStatusCanceled = EmailStatus("canceled")
)

func (enum EmailStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum EmailStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *EmailStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = EmailStatus(EmailStatus(tmp).String())
	return nil
}

type ListEmailsRequestOrderBy string

const (
	// Order by creation date (descending chronological order).
	ListEmailsRequestOrderByCreatedAtDesc = ListEmailsRequestOrderBy("created_at_desc")
	// Order by creation date (ascending chronological order).
	ListEmailsRequestOrderByCreatedAtAsc = ListEmailsRequestOrderBy("created_at_asc")
	// Order by last update date (descending chronological order).
	ListEmailsRequestOrderByUpdatedAtDesc = ListEmailsRequestOrderBy("updated_at_desc")
	// Order by last update date (ascending chronological order).
	ListEmailsRequestOrderByUpdatedAtAsc = ListEmailsRequestOrderBy("updated_at_asc")
	// Order by status (descending alphabetical order).
	ListEmailsRequestOrderByStatusDesc = ListEmailsRequestOrderBy("status_desc")
	// Order by status (ascending alphabetical order).
	ListEmailsRequestOrderByStatusAsc = ListEmailsRequestOrderBy("status_asc")
	// Order by mail_from (descending alphabetical order).
	ListEmailsRequestOrderByMailFromDesc = ListEmailsRequestOrderBy("mail_from_desc")
	// Order by mail_from (ascending alphabetical order).
	ListEmailsRequestOrderByMailFromAsc = ListEmailsRequestOrderBy("mail_from_asc")
	// Order by mail recipient (descending alphabetical order).
	ListEmailsRequestOrderByMailRcptDesc = ListEmailsRequestOrderBy("mail_rcpt_desc")
	// Order by mail recipient (ascending alphabetical order).
	ListEmailsRequestOrderByMailRcptAsc = ListEmailsRequestOrderBy("mail_rcpt_asc")
	// Order by subject (descending alphabetical order).
	ListEmailsRequestOrderBySubjectDesc = ListEmailsRequestOrderBy("subject_desc")
	// Order by subject (ascending alphabetical order).
	ListEmailsRequestOrderBySubjectAsc = ListEmailsRequestOrderBy("subject_asc")
)

func (enum ListEmailsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_desc"
	}
	return string(enum)
}

func (enum ListEmailsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListEmailsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListEmailsRequestOrderBy(ListEmailsRequestOrderBy(tmp).String())
	return nil
}

// EmailTry:
type EmailTry struct {
	// Rank: Rank number of this attempt to send the email.
	Rank uint32 `json:"rank"`
	// TriedAt: Date of the attempt to send the email.
	TriedAt *time.Time `json:"tried_at,omitempty"`
	// Code: The SMTP status code received after the attempt. 0 if the attempt did not reach an SMTP server.
	Code int32 `json:"code"`
	// Message: The SMTP message received. If the attempt did not reach an SMTP server, the message returned explains what happened.
	Message string `json:"message"`
}

// DomainStatistics:
type DomainStatistics struct {
	// TotalCount:
	TotalCount uint32 `json:"total_count"`
	// SentCount:
	SentCount uint32 `json:"sent_count"`
	// FailedCount:
	FailedCount uint32 `json:"failed_count"`
	// CanceledCount:
	CanceledCount uint32 `json:"canceled_count"`
}

// CreateEmailRequestAddress:
type CreateEmailRequestAddress struct {
	// Email: Email address.
	Email string `json:"email"`
	// Name: (Optional) Name displayed.
	Name *string `json:"name,omitempty"`
}

// CreateEmailRequestAttachment:
type CreateEmailRequestAttachment struct {
	// Name: Filename of the attachment.
	Name string `json:"name"`
	// Type: MIME type of the attachment.
	Type string `json:"type"`
	// Content: Content of the attachment encoded in base64.
	Content []byte `json:"content"`
}

// Email:
type Email struct {
	// ID: Technical ID of the email.
	ID string `json:"id"`
	// MessageID: Message ID of the email.
	MessageID string `json:"message_id"`
	// ProjectID: ID of the Project to which the email belongs.
	ProjectID string `json:"project_id"`
	// MailFrom: Email address of the sender.
	MailFrom string `json:"mail_from"`
	// Deprecated: RcptTo: Email address of the recipient.
	RcptTo *string `json:"rcpt_to,omitempty"`
	// MailRcpt: Email address of the recipient.
	MailRcpt string `json:"mail_rcpt"`
	// RcptType: Type of recipient.
	RcptType EmailRcptType `json:"rcpt_type"`
	// Subject: Subject of the email.
	Subject string `json:"subject"`
	// CreatedAt: Creation date of the email object.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt: Last update of the email object.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Status: Status of the email.
	Status EmailStatus `json:"status"`
	// StatusDetails: Additional status information.
	StatusDetails *string `json:"status_details,omitempty"`
	// TryCount: Number of attempts to send the email.
	TryCount uint32 `json:"try_count"`
	// LastTries: Information about the last three attempts to send the email.
	LastTries []*EmailTry `json:"last_tries"`
	// Flags: Flags categorize emails. They allow you to obtain more information about recurring errors, for example.
	Flags []EmailFlag `json:"flags"`
}

// DomainLastStatusDkimRecord:
type DomainLastStatusDkimRecord struct {
	// Status: Status of the DKIM record's configurartion.
	Status DomainLastStatusRecordStatus `json:"status"`
	// LastValidAt: Time and date the DKIM record was last valid.
	LastValidAt *time.Time `json:"last_valid_at,omitempty"`
	// Error: An error text displays in case the record is not valid.
	Error *string `json:"error,omitempty"`
}

// DomainLastStatusSpfRecord:
type DomainLastStatusSpfRecord struct {
	// Status: Status of the SPF record's configurartion.
	Status DomainLastStatusRecordStatus `json:"status"`
	// LastValidAt: Time and date the SPF record was last valid.
	LastValidAt *time.Time `json:"last_valid_at,omitempty"`
	// Error: An error text displays in case the record is not valid.
	Error *string `json:"error,omitempty"`
}

// Domain:
type Domain struct {
	// ID: ID of the domain.
	ID string `json:"id"`
	// OrganizationID: ID of the domain's Organization.
	OrganizationID string `json:"organization_id"`
	// ProjectID: ID of the domain's Project.
	ProjectID string `json:"project_id"`
	// Name: Domain name (example.com).
	Name string `json:"name"`
	// Status: Status of the domain.
	Status DomainStatus `json:"status"`
	// CreatedAt: Date and time of domain creation.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// NextCheckAt: Date and time of the next scheduled check.
	NextCheckAt *time.Time `json:"next_check_at,omitempty"`
	// LastValidAt: Date and time the domain was last valid.
	LastValidAt *time.Time `json:"last_valid_at,omitempty"`
	// RevokedAt: Date and time of the domain's deletion.
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
	// Deprecated: LastError: Error message returned if the last check failed.
	LastError *string `json:"last_error,omitempty"`
	// SpfConfig: Snippet of the SPF record to register in the DNS zone.
	SpfConfig string `json:"spf_config"`
	// DkimConfig: DKIM public key to record in the DNS zone.
	DkimConfig string `json:"dkim_config"`
	// Statistics: Domain's statistics.
	Statistics *DomainStatistics `json:"statistics"`
	// Region:
	Region scw.Region `json:"region"`
}

// CancelEmailRequest:
type CancelEmailRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// EmailID: ID of the email to cancel.
	EmailID string `json:"-"`
}

// CheckDomainRequest:
type CheckDomainRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// DomainID: ID of the domain to check.
	DomainID string `json:"-"`
}

// CreateDomainRequest:
type CreateDomainRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID: ID of the project to which the domain belongs.
	ProjectID string `json:"project_id"`
	// DomainName: Fully qualified domain dame.
	DomainName string `json:"domain_name"`
	// AcceptTos: Accept Scaleway's Terms of Service.
	AcceptTos bool `json:"accept_tos"`
}

// CreateEmailRequest:
type CreateEmailRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// From: Sender information. Must be from a checked domain declared in the Project.
	From *CreateEmailRequestAddress `json:"from"`
	// To: An array of the primary recipient's information.
	To []*CreateEmailRequestAddress `json:"to"`
	// Cc: An array of the carbon copy recipient's information.
	Cc []*CreateEmailRequestAddress `json:"cc"`
	// Bcc: An array of the blind carbon copy recipient's information.
	Bcc []*CreateEmailRequestAddress `json:"bcc"`
	// Subject: Subject of the email.
	Subject string `json:"subject"`
	// Text: Text content.
	Text string `json:"text"`
	// HTML: HTML content.
	HTML string `json:"html"`
	// ProjectID: ID of the Project in which to create the email.
	ProjectID string `json:"project_id"`
	// Attachments: Array of attachments.
	Attachments []*CreateEmailRequestAttachment `json:"attachments"`
	// SendBefore: Maximum date to deliver the email.
	SendBefore *time.Time `json:"send_before,omitempty"`
}

// CreateEmailResponse:
type CreateEmailResponse struct {
	// Emails: Single page of emails matching the requested criteria.
	Emails []*Email `json:"emails"`
}

// DomainLastStatus:
type DomainLastStatus struct {
	// DomainID: The id of the domain.
	DomainID string `json:"domain_id"`
	// DomainName: The domain name (example.com).
	DomainName string `json:"domain_name"`
	// SpfRecord: The SPF record verification data.
	SpfRecord *DomainLastStatusSpfRecord `json:"spf_record"`
	// DkimRecord: The DKIM record verification data.
	DkimRecord *DomainLastStatusDkimRecord `json:"dkim_record"`
}

// GetDomainLastStatusRequest:
type GetDomainLastStatusRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// DomainID: ID of the domain to delete.
	DomainID string `json:"-"`
}

// GetDomainRequest:
type GetDomainRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// DomainID: ID of the domain.
	DomainID string `json:"-"`
}

// GetEmailRequest:
type GetEmailRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// EmailID: ID of the email to retrieve.
	EmailID string `json:"-"`
}

// GetStatisticsRequest:
type GetStatisticsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// ProjectID: (Optional) Number of emails for this Project.
	ProjectID *string `json:"-"`
	// DomainID: (Optional) Number of emails sent from this domain (must be coherent with the `project_id` and the `organization_id`).
	DomainID *string `json:"-"`
	// Since: (Optional) Number of emails created after this date.
	Since *time.Time `json:"-"`
	// Until: (Optional) Number of emails created before this date.
	Until *time.Time `json:"-"`
	// MailFrom: (Optional) Number of emails sent with this sender's email address.
	MailFrom *string `json:"-"`
}

// ListDomainsRequest:
type ListDomainsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Page: Requested page number. Value must be greater or equal to 1.
	Page *int32 `json:"-"`
	// PageSize: Page size.
	PageSize *uint32 `json:"-"`
	// ProjectID:
	ProjectID *string `json:"-"`
	// Status:
	Status []DomainStatus `json:"-"`
	// OrganizationID:
	OrganizationID *string `json:"-"`
	// Name:
	Name *string `json:"-"`
}

// ListDomainsResponse:
type ListDomainsResponse struct {
	// TotalCount: Number of domains that match the request (without pagination).
	TotalCount uint32 `json:"total_count"`
	// Domains:
	Domains []*Domain `json:"domains"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDomainsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDomainsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDomainsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Domains = append(r.Domains, results.Domains...)
	r.TotalCount += uint32(len(results.Domains))
	return uint32(len(results.Domains)), nil
}

// ListEmailsRequest:
type ListEmailsRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// Page:
	Page *int32 `json:"-"`
	// PageSize:
	PageSize *uint32 `json:"-"`
	// ProjectID: (Optional) ID of the Project in which to list the emails.
	ProjectID *string `json:"-"`
	// DomainID: (Optional) ID of the domain for which to list the emails.
	DomainID *string `json:"-"`
	// MessageID: (Optional) ID of the message for which to list the emails.
	MessageID *string `json:"-"`
	// Since: (Optional) List emails created after this date.
	Since *time.Time `json:"-"`
	// Until: (Optional) List emails created before this date.
	Until *time.Time `json:"-"`
	// MailFrom: (Optional) List emails sent with this sender's email address.
	MailFrom *string `json:"-"`
	// Deprecated: MailTo: List emails sent to this recipient's email address.
	MailTo *string `json:"-"`
	// MailRcpt: (Optional) List emails sent to this recipient's email address.
	MailRcpt *string `json:"-"`
	// Statuses: (Optional) List emails with any of these statuses.
	Statuses []EmailStatus `json:"-"`
	// Subject: (Optional) List emails with this subject.
	Subject *string `json:"-"`
	// Search: (Optional) List emails by searching to all fields.
	Search *string `json:"-"`
	// OrderBy: (Optional) List emails corresponding to specific criteria.
	OrderBy ListEmailsRequestOrderBy `json:"-"`
	// Flags: (Optional) List emails containing only specific flags.
	Flags []EmailFlag `json:"-"`
}

// ListEmailsResponse:
type ListEmailsResponse struct {
	// TotalCount: Number of emails matching the requested criteria.
	TotalCount uint32 `json:"total_count"`
	// Emails: Single page of emails matching the requested criteria.
	Emails []*Email `json:"emails"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListEmailsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListEmailsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListEmailsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Emails = append(r.Emails, results.Emails...)
	r.TotalCount += uint32(len(results.Emails))
	return uint32(len(results.Emails)), nil
}

// RevokeDomainRequest:
type RevokeDomainRequest struct {
	// Region:
	Region scw.Region `json:"-"`
	// DomainID: ID of the domain to delete.
	DomainID string `json:"-"`
}

// Statistics:
type Statistics struct {
	// TotalCount: Total number of emails matching the requested criteria.
	TotalCount uint32 `json:"total_count"`
	// NewCount: Number of emails still in the `new` transient state. This means emails received from the API but not yet processed.
	NewCount uint32 `json:"new_count"`
	// SendingCount: Number of emails still in the `sending` transient state. This means emails received from the API but not yet in their final status.
	SendingCount uint32 `json:"sending_count"`
	// SentCount: Number of emails in the final `sent` state. This means emails that have been delivered to the target mail system.
	SentCount uint32 `json:"sent_count"`
	// FailedCount: Number of emails in the final `failed` state. This means emails that have been refused by the target mail system with a final error status.
	FailedCount uint32 `json:"failed_count"`
	// CanceledCount: Number of emails in the final `canceled` state. This means emails that have been canceled upon request.
	CanceledCount uint32 `json:"canceled_count"`
}

// Transactional Email is a platform that allows you to send transactional emails. Unlike marketing emails, (which are mass distributions of the same marketing message to multiple recipients), transactional emails are personalized emails sent to individuals in response to events they have triggered (e.g. password recovery, billing information, delivery updates, etc).
//
// With Transactional Email, you can:
// - register a domain from which to send emails
// - send transactional emails via an SMTP relay
// - send transactional emails via a REST API
// - view your email's delivery status
// - cancel emails that you no longer wish to send
//
// (switchcolumn)
// (switchcolumn)
//
// ## Concepts
//
// Refer to our [dedicated concepts page](https://www.scaleway.com/en/docs/managed-services/transactional-email/concepts/) to find definitions of the different terms referring to Transactional Email.
//
// (switchcolumn)
// (switchcolumn)
// ## Quickstart
//
// 1. **Configure your environment variables.**
//
//	<Message type="note">
//	  This is an optional step that seeks to simplify your usage of the API.
//	</Message>
//
//
//	```bash
//	export SCW_ACCESS_KEY="<API access key>"
//	export SCW_SECRET_KEY="<API secret key>"
//	export SCW_PROJECT_ID="<Scaleway Project ID>"
//	```
//
// 2. **Register your domain**. Run the following command to register your domain. Make sure you note down the `dkim_config` that displays in the output, as you will need it for the next steps.
//
//	```bash
//	curl -X POST "https://api.scaleway.com/transactional-email/v1alpha1/regions/$REGION/domains" \
//	    -H "X-Auth-Token: $SCW_SECRET_KEY" \
//	    -d "{\"domain_name\":\"my.domain.example.com\",\"project_id\":\"$PROJECT_ID\"}"
//	```
//
// 3. **Create an SPF and a DKIM record**.
//
//	Find out [how to create your SPF and DKIM records](https://www.scaleway.com/en/docs/managed-services/transactional-email/how-to/add-spf-dkim-records-to-your-domain/) in our documentation.
//
//	<Message type="note">
//	  The DKIM private key is discarded when your domain is revoked. You must update the corresponding DKIM record if you register that domain again.
//	</Message>
//
// 4. **Create an MX record**
//
//	<Message type="note">
//	  Adding an MX record is not mandatory but we recommend you do it to avoid getting your emails rejected.
//	</Message>
//
//	Find out [how to add an MX record](https://www.scaleway.com/en/docs/managed-services/transactional-email/how-to/add-mx-record-to-your-domain/) in our documentation.
//
// 5. **Perform a check of your domain**. Run the following commands to query your domain and then perform a check of your domain.
//
//	```bash
//	curl -X POST "https://api.scaleway.com/transactional-email/v1alpha1/regions/$REGION/domains/<domain-id>/check" \
//	    -H "X-Auth-Token: $TOKEN" \
//	    -d "{}"
//
//	# Wait a few seconds,then run
//
//	curl -X GET "https://api.scaleway.com/transactional-email/v1alpha1/regions/$REGION/domains/<domain-id>" \
//	    -H "X-Auth-Token: $TOKEN"
//	```
//
//	<Message type="note">
//	  If the check fails, `last_error` will describe what happened. Perform another check request once you have fixed the problem. To ensure good scoring, your DNS must be correctly set up before sending any emails. The domain check is performed asynchronously.
//	</Message>
//
// 6. **Send an email via the REST API**. Run the following command to send an email using the REST API.
//
//	<Message type="note">
//	  The attachment content must be base64 encoded and is limited to 2MB. In the following example, both `html` and `attachments` fields are optional.
//	</Message>
//
//	```bash
//	cat > mail.json <<EOF
//	{
//	    "from": {
//	        "name": "Alice",
//	        "email": "alice@my.domain.example.com"
//	    },
//	    "to": [
//	        {
//	        "name": "Bob",
//	        "email": "bob@caramail.fr"
//	        }
//	    ],
//	    "subject": "Some subject",
//	    "text": "Some text.",
//	    "html": "<p>Some <span style=\"font-weight:bold\">HTML</span>.</p>",
//	    "project_id": "<project-id>",
//	    "attachments": [
//	        {
//	            "name": "file.html",
//	            "type": "text/html; charset=\"utf8\"",
//	            "content": "PCFET0NUWVBFIGh0bWw+CjxodG1sIGxhbmc9ImVuIj4KICAgIDxoZWFkPgogICAgICAgIDxtZXRhIGNoYXJzZXQ9InV0Zi04Ij4KICAgICAgICA8dGl0bGU+dGl0bGU8L3RpdGxlPgogICAgPC9oZWFkPgogICAgPGJvZHk+CiAgICAgICAgPHA+SGVsbG8gPHNwYW4gc3R5bGU9ImNvbG9yOmJsdWUiPndvcmxkPC9zcGFuPiE8L3A+CiAgICA8L2JvZHk+CjwvaHRtbD4="
//	        }
//	    ]
//	}
//	EOF
//
//	curl -X POST "https://api.scaleway.com/transactional-email/v1alpha1/regions/$REGION/emails" \
//	    -H "X-Auth-Token: $TOKEN" \
//	    -d @mail.json
//	```
//
// 7. **Send an email via SMTP**.
//
//	An SMTP relay is available at `smtp.tem.scw.cloud` on ports `25`, `587`, `2587`, `465` (TLS) and `2465` (TLS).
//
//	Use your `project_id` to authenticate as username and your token as a password.
//
//	<Message type="note">
//	  It is possible to remove the dashes in UUIDs to shorten them.
//	</Message>
//
// (switchcolumn)
// <Message type="requirement">
//   - You have a [Scaleway account](https://console.scaleway.com/)
//   - You are the owner of the domain name you want to use
//   - You have your [Organization and your Project ID](https://console.scaleway.com/project/settings)
//   - You have created an [API key](https://www.scaleway.com/en/docs/identity-and-access-management/iam/how-to/create-api-keys/) and that the API key has sufficient [IAM permissions](https://www.scaleway.com/en/docs/identity-and-access-management/iam/reference-content/permission-sets/) to perform the actions described on this page.
//   - You have [installed `curl`](https://curl.se/download.html)
//
// </Message>
// (switchcolumn)
//
// ## Technical information
// ### Regions
//
// Scaleway's infrastructure spans different [regions and Availability Zones](https://www.scaleway.com/en/docs/console/my-account/reference-content/products-availability/).
//
// Transactional Email is available in the Paris region, which is represented by the following path parameter: `fr-par`.
//
// ## Technical limitations
//
// * Transactional Email's attachment types are limited to:
//   - `application/ics`
//   - `application/pdf`
//   - `image/gif`
//   - `image/jpeg`
//   - `image/jpg`
//   - `image/png`
//   - `text/calendar`
//   - `text/csv`
//   - `text/html`
//   - `text/plain`
//
// * Attachment size is limited to 2MB
// * The recursive SPF include directive is not supported
// * The subject, body and HTML of your emails must have at least 10 characters
//
// ##  Going further
//
// For more information about Transactional Email, you can check out the following pages:
//
// * [Transactional Email Documentation](https://www.scaleway.com/en/docs/managed-services/transactional-email/)
// * [Scaleway Slack Community](https://scaleway-community.slack.com/) join the #transactional-email channel
// * [Contact our support team](https://console.scaleway.com/support/tickets).
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
	return []scw.Region{scw.RegionFrPar}
}

// CreateEmail: You must specify the `region`, the sender and the recipient's information and the `project_id` to send an email from a checked domain. The subject of the email must contain at least 6 characters.
func (s *API) CreateEmail(req *CreateEmailRequest, opts ...scw.RequestOption) (*CreateEmailResponse, error) {
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
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/emails",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateEmailResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetEmail: Retrieve information about a specific email using the `email_id` and `region` parameters.
func (s *API) GetEmail(req *GetEmailRequest, opts ...scw.RequestOption) (*Email, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.EmailID) == "" {
		return nil, errors.New("field EmailID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/emails/" + fmt.Sprint(req.EmailID) + "",
	}

	var resp Email

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListEmails: Retrieve the list of emails sent from a specific domain or for a specific Project or Organization. You must specify the `region`.
func (s *API) ListEmails(req *ListEmailsRequest, opts ...scw.RequestOption) (*ListEmailsResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "domain_id", req.DomainID)
	parameter.AddToQuery(query, "message_id", req.MessageID)
	parameter.AddToQuery(query, "since", req.Since)
	parameter.AddToQuery(query, "until", req.Until)
	parameter.AddToQuery(query, "mail_from", req.MailFrom)
	parameter.AddToQuery(query, "mail_to", req.MailTo)
	parameter.AddToQuery(query, "mail_rcpt", req.MailRcpt)
	parameter.AddToQuery(query, "statuses", req.Statuses)
	parameter.AddToQuery(query, "subject", req.Subject)
	parameter.AddToQuery(query, "search", req.Search)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "flags", req.Flags)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/emails",
		Query:  query,
	}

	var resp ListEmailsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetStatistics: Get information on your emails' statuses.
func (s *API) GetStatistics(req *GetStatisticsRequest, opts ...scw.RequestOption) (*Statistics, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "domain_id", req.DomainID)
	parameter.AddToQuery(query, "since", req.Since)
	parameter.AddToQuery(query, "until", req.Until)
	parameter.AddToQuery(query, "mail_from", req.MailFrom)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/statistics",
		Query:  query,
	}

	var resp Statistics

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CancelEmail: You can cancel the sending of an email if it has not been sent yet. You must specify the `region` and the `email_id` of the email you want to cancel.
func (s *API) CancelEmail(req *CancelEmailRequest, opts ...scw.RequestOption) (*Email, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.EmailID) == "" {
		return nil, errors.New("field EmailID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/emails/" + fmt.Sprint(req.EmailID) + "/cancel",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Email

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDomain: You must specify the `region`, `project_id` and `domain_name` to register a domain in a specific Project.
func (s *API) CreateDomain(req *CreateDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
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
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDomain: Retrieve information about a specific domain using the `region` and `domain_id` parameters.
func (s *API) GetDomain(req *GetDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DomainID) == "" {
		return nil, errors.New("field DomainID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "",
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDomains: Retrieve domains in a specific project or in a specific Organization using the `region` parameter.
func (s *API) ListDomains(req *ListDomainsRequest, opts ...scw.RequestOption) (*ListDomainsResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains",
		Query:  query,
	}

	var resp ListDomainsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevokeDomain: You must specify the domain you want to delete by the `region` and `domain_id`. Deleting a domain is permanent and cannot be undone.
func (s *API) RevokeDomain(req *RevokeDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DomainID) == "" {
		return nil, errors.New("field DomainID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "/revoke",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckDomain: Perform an immediate DNS check of a domain using the `region` and `domain_id` parameters.
func (s *API) CheckDomain(req *CheckDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DomainID) == "" {
		return nil, errors.New("field DomainID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "/check",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDomainLastStatus: Display SPF and DKIM records status and potential errors, including the found records to make debugging easier.
func (s *API) GetDomainLastStatus(req *GetDomainLastStatusRequest, opts ...scw.RequestOption) (*DomainLastStatus, error) {
	var err error
	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DomainID) == "" {
		return nil, errors.New("field DomainID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "/verification",
	}

	var resp DomainLastStatus

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
