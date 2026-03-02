// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package agentmail

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/agentmail-to/agentmail-go/internal/apijson"
	"github.com/agentmail-to/agentmail-go/internal/apiquery"
	"github.com/agentmail-to/agentmail-go/internal/requestconfig"
	"github.com/agentmail-to/agentmail-go/option"
	"github.com/agentmail-to/agentmail-go/packages/param"
	"github.com/agentmail-to/agentmail-go/packages/respjson"
)

// DraftService contains methods and other services that help with interacting with
// the agentmail API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDraftService] method instead.
type DraftService struct {
	Options []option.RequestOption
}

// NewDraftService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDraftService(opts ...option.RequestOption) (r DraftService) {
	r = DraftService{}
	r.Options = opts
	return
}

// Get Draft
func (r *DraftService) Get(ctx context.Context, draftID string, opts ...option.RequestOption) (res *Draft, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if draftID == "" {
		err = errors.New("missing required draft_id parameter")
		return
	}
	path := fmt.Sprintf("v0/drafts/%s", url.PathEscape(draftID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Drafts
func (r *DraftService) List(ctx context.Context, query DraftListParams, opts ...option.RequestOption) (res *ListDrafts, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	path := "v0/drafts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AttachmentFile struct {
	// ID of attachment.
	AttachmentID string `json:"attachment_id" api:"required"`
	// Size of attachment in bytes.
	Size int64 `json:"size" api:"required"`
	// Content disposition of attachment.
	//
	// Any of "inline", "attachment".
	ContentDisposition AttachmentContentDisposition `json:"content_disposition" api:"nullable"`
	// Content ID of attachment.
	ContentID string `json:"content_id" api:"nullable"`
	// Content type of attachment.
	ContentType string `json:"content_type" api:"nullable"`
	// Filename of attachment.
	Filename string `json:"filename" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttachmentID       respjson.Field
		Size               respjson.Field
		ContentDisposition respjson.Field
		ContentID          respjson.Field
		ContentType        respjson.Field
		Filename           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttachmentFile) RawJSON() string { return r.JSON.raw }
func (r *AttachmentFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Draft struct {
	// Time at which draft was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// ID of draft.
	DraftID string `json:"draft_id" api:"required"`
	// ID of inbox.
	InboxID string `json:"inbox_id" api:"required"`
	// Labels of draft.
	Labels []string `json:"labels" api:"required"`
	// ID of thread.
	ThreadID string `json:"thread_id" api:"required"`
	// Time at which draft was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Attachments in draft.
	Attachments []AttachmentFile `json:"attachments" api:"nullable"`
	// Addresses of BCC recipients. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	Bcc []string `json:"bcc" api:"nullable"`
	// Addresses of CC recipients. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	Cc []string `json:"cc" api:"nullable"`
	// Client ID of draft.
	ClientID string `json:"client_id" api:"nullable"`
	// HTML body of draft.
	HTML string `json:"html" api:"nullable"`
	// ID of message being replied to.
	InReplyTo string `json:"in_reply_to" api:"nullable"`
	// Text preview of draft.
	Preview string `json:"preview" api:"nullable"`
	// IDs of previous messages in thread.
	References []string `json:"references" api:"nullable"`
	// Reply-to addresses. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	ReplyTo []string `json:"reply_to" api:"nullable"`
	// Time at which to schedule send draft.
	SendAt time.Time `json:"send_at" api:"nullable" format:"date-time"`
	// Schedule send status of draft.
	//
	// Any of "scheduled", "sending", "failed".
	SendStatus DraftSendStatus `json:"send_status" api:"nullable"`
	// Subject of draft.
	Subject string `json:"subject" api:"nullable"`
	// Plain text body of draft.
	Text string `json:"text" api:"nullable"`
	// Addresses of recipients. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	To []string `json:"to" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		DraftID     respjson.Field
		InboxID     respjson.Field
		Labels      respjson.Field
		ThreadID    respjson.Field
		UpdatedAt   respjson.Field
		Attachments respjson.Field
		Bcc         respjson.Field
		Cc          respjson.Field
		ClientID    respjson.Field
		HTML        respjson.Field
		InReplyTo   respjson.Field
		Preview     respjson.Field
		References  respjson.Field
		ReplyTo     respjson.Field
		SendAt      respjson.Field
		SendStatus  respjson.Field
		Subject     respjson.Field
		Text        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Draft) RawJSON() string { return r.JSON.raw }
func (r *Draft) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schedule send status of draft.
type DraftSendStatus string

const (
	DraftSendStatusScheduled DraftSendStatus = "scheduled"
	DraftSendStatusSending   DraftSendStatus = "sending"
	DraftSendStatusFailed    DraftSendStatus = "failed"
)

type ListDrafts struct {
	// Number of items returned.
	Count int64 `json:"count" api:"required"`
	// Ordered by `updated_at` descending.
	Drafts []ListDraftsDraft `json:"drafts" api:"required"`
	// Limit of number of items returned.
	Limit int64 `json:"limit" api:"nullable"`
	// Page token for pagination.
	NextPageToken string `json:"next_page_token" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count         respjson.Field
		Drafts        respjson.Field
		Limit         respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListDrafts) RawJSON() string { return r.JSON.raw }
func (r *ListDrafts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListDraftsDraft struct {
	// ID of draft.
	DraftID string `json:"draft_id" api:"required"`
	// ID of inbox.
	InboxID string `json:"inbox_id" api:"required"`
	// Labels of draft.
	Labels []string `json:"labels" api:"required"`
	// ID of thread.
	ThreadID string `json:"thread_id" api:"required"`
	// Time at which draft was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Attachments in draft.
	Attachments []AttachmentFile `json:"attachments" api:"nullable"`
	// Addresses of BCC recipients. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	Bcc []string `json:"bcc" api:"nullable"`
	// Addresses of CC recipients. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	Cc []string `json:"cc" api:"nullable"`
	// Text preview of draft.
	Preview string `json:"preview" api:"nullable"`
	// Time at which to schedule send draft.
	SendAt time.Time `json:"send_at" api:"nullable" format:"date-time"`
	// Schedule send status of draft.
	//
	// Any of "scheduled", "sending", "failed".
	SendStatus DraftSendStatus `json:"send_status" api:"nullable"`
	// Subject of draft.
	Subject string `json:"subject" api:"nullable"`
	// Addresses of recipients. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	To []string `json:"to" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DraftID     respjson.Field
		InboxID     respjson.Field
		Labels      respjson.Field
		ThreadID    respjson.Field
		UpdatedAt   respjson.Field
		Attachments respjson.Field
		Bcc         respjson.Field
		Cc          respjson.Field
		Preview     respjson.Field
		SendAt      respjson.Field
		SendStatus  respjson.Field
		Subject     respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListDraftsDraft) RawJSON() string { return r.JSON.raw }
func (r *ListDraftsDraft) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DraftListParams struct {
	// Timestamp after which to filter by.
	After param.Opt[time.Time] `query:"after,omitzero" format:"date-time" json:"-"`
	// Sort in ascending temporal order.
	Ascending param.Opt[bool] `query:"ascending,omitzero" json:"-"`
	// Timestamp before which to filter by.
	Before param.Opt[time.Time] `query:"before,omitzero" format:"date-time" json:"-"`
	// Limit of number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page token for pagination.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	// Labels to filter by.
	Labels []string `query:"labels,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DraftListParams]'s query parameters as `url.Values`.
func (r DraftListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
