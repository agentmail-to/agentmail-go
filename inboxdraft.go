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
	shimjson "github.com/agentmail-to/agentmail-go/internal/encoding/json"
	"github.com/agentmail-to/agentmail-go/internal/requestconfig"
	"github.com/agentmail-to/agentmail-go/option"
	"github.com/agentmail-to/agentmail-go/packages/param"
	"github.com/agentmail-to/agentmail-go/packages/respjson"
)

// InboxDraftService contains methods and other services that help with interacting
// with the agentmail API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboxDraftService] method instead.
type InboxDraftService struct {
	Options []option.RequestOption
}

// NewInboxDraftService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInboxDraftService(opts ...option.RequestOption) (r InboxDraftService) {
	r = InboxDraftService{}
	r.Options = opts
	return
}

// **CLI:**
//
// ```bash
// agentmail inboxes:drafts create --inbox-id <inbox_id> --to recipient@example.com --subject "Draft subject" --text "Draft body"
// ```
func (r *InboxDraftService) New(ctx context.Context, inboxID string, body InboxDraftNewParams, opts ...option.RequestOption) (res *Draft, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/inboxes/%s/drafts", url.PathEscape(inboxID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail inboxes:drafts update --inbox-id <inbox_id> --draft-id <draft_id> --subject "Updated subject"
// ```
func (r *InboxDraftService) Update(ctx context.Context, draftID string, params InboxDraftUpdateParams, opts ...option.RequestOption) (res *Draft, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if params.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	if draftID == "" {
		err = errors.New("missing required draft_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/inboxes/%s/drafts/%s", url.PathEscape(params.InboxID), url.PathEscape(draftID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail inboxes:drafts list --inbox-id <inbox_id>
// ```
func (r *InboxDraftService) List(ctx context.Context, inboxID string, query InboxDraftListParams, opts ...option.RequestOption) (res *ListDrafts, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/inboxes/%s/drafts", url.PathEscape(inboxID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail inboxes:drafts delete --inbox-id <inbox_id> --draft-id <draft_id>
// ```
func (r *InboxDraftService) Delete(ctx context.Context, draftID string, body InboxDraftDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if body.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return err
	}
	if draftID == "" {
		err = errors.New("missing required draft_id parameter")
		return err
	}
	path := fmt.Sprintf("v0/inboxes/%s/drafts/%s", url.PathEscape(body.InboxID), url.PathEscape(draftID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// **CLI:**
//
// ```bash
// agentmail inboxes:drafts get --inbox-id <inbox_id> --draft-id <draft_id>
// ```
func (r *InboxDraftService) Get(ctx context.Context, draftID string, query InboxDraftGetParams, opts ...option.RequestOption) (res *Draft, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if query.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	if draftID == "" {
		err = errors.New("missing required draft_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/inboxes/%s/drafts/%s", url.PathEscape(query.InboxID), url.PathEscape(draftID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail inboxes:drafts get-attachment --inbox-id <inbox_id> --draft-id <draft_id> --attachment-id <attachment_id>
// ```
func (r *InboxDraftService) GetAttachment(ctx context.Context, attachmentID string, query InboxDraftGetAttachmentParams, opts ...option.RequestOption) (res *AttachmentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if query.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	if query.DraftID == "" {
		err = errors.New("missing required draft_id parameter")
		return nil, err
	}
	if attachmentID == "" {
		err = errors.New("missing required attachment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/inboxes/%s/drafts/%s/attachments/%s", url.PathEscape(query.InboxID), url.PathEscape(query.DraftID), url.PathEscape(attachmentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail inboxes:drafts send --inbox-id <inbox_id> --draft-id <draft_id>
// ```
func (r *InboxDraftService) Send(ctx context.Context, draftID string, params InboxDraftSendParams, opts ...option.RequestOption) (res *SendMessageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if params.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	if draftID == "" {
		err = errors.New("missing required draft_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/inboxes/%s/drafts/%s/send", url.PathEscape(params.InboxID), url.PathEscape(draftID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type SendMessageResponse struct {
	// ID of message.
	MessageID string `json:"message_id" api:"required"`
	// ID of thread.
	ThreadID string `json:"thread_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessageID   respjson.Field
		ThreadID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SendMessageResponse) RawJSON() string { return r.JSON.raw }
func (r *SendMessageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UpdateMessageParam struct {
	// Label or labels to add to message.
	AddLabels UpdateMessageAddLabelsUnionParam `json:"add_labels,omitzero"`
	// Label or labels to remove from message.
	RemoveLabels UpdateMessageRemoveLabelsUnionParam `json:"remove_labels,omitzero"`
	paramObj
}

func (r UpdateMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UpdateMessageAddLabelsUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u UpdateMessageAddLabelsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *UpdateMessageAddLabelsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UpdateMessageRemoveLabelsUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u UpdateMessageRemoveLabelsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *UpdateMessageRemoveLabelsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type InboxDraftNewParams struct {
	// Client ID of draft.
	ClientID param.Opt[string] `json:"client_id,omitzero"`
	// HTML body of draft.
	HTML param.Opt[string] `json:"html,omitzero"`
	// ID of message being replied to.
	InReplyTo param.Opt[string] `json:"in_reply_to,omitzero"`
	// Time at which to schedule send draft.
	SendAt param.Opt[time.Time] `json:"send_at,omitzero" format:"date-time"`
	// Subject of draft.
	Subject param.Opt[string] `json:"subject,omitzero"`
	// Plain text body of draft.
	Text param.Opt[string] `json:"text,omitzero"`
	// Attachments to include in draft.
	Attachments []SendAttachmentParam `json:"attachments,omitzero"`
	// Addresses of BCC recipients. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	Bcc []string `json:"bcc,omitzero"`
	// Addresses of CC recipients. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	Cc []string `json:"cc,omitzero"`
	// Labels of draft.
	Labels []string `json:"labels,omitzero"`
	// Reply-to addresses. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	ReplyTo []string `json:"reply_to,omitzero"`
	// Addresses of recipients. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	To []string `json:"to,omitzero"`
	paramObj
}

func (r InboxDraftNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InboxDraftNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboxDraftNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxDraftUpdateParams struct {
	// The ID of the inbox.
	InboxID string `path:"inbox_id" api:"required" json:"-"`
	// HTML body of draft.
	HTML param.Opt[string] `json:"html,omitzero"`
	// Time at which to schedule send draft.
	SendAt param.Opt[time.Time] `json:"send_at,omitzero" format:"date-time"`
	// Subject of draft.
	Subject param.Opt[string] `json:"subject,omitzero"`
	// Plain text body of draft.
	Text param.Opt[string] `json:"text,omitzero"`
	// Addresses of BCC recipients. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	Bcc []string `json:"bcc,omitzero"`
	// Addresses of CC recipients. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	Cc []string `json:"cc,omitzero"`
	// Reply-to addresses. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	ReplyTo []string `json:"reply_to,omitzero"`
	// Addresses of recipients. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	To []string `json:"to,omitzero"`
	paramObj
}

func (r InboxDraftUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow InboxDraftUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboxDraftUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxDraftListParams struct {
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

// URLQuery serializes [InboxDraftListParams]'s query parameters as `url.Values`.
func (r InboxDraftListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InboxDraftDeleteParams struct {
	// The ID of the inbox.
	InboxID string `path:"inbox_id" api:"required" json:"-"`
	paramObj
}

type InboxDraftGetParams struct {
	// The ID of the inbox.
	InboxID string `path:"inbox_id" api:"required" json:"-"`
	paramObj
}

type InboxDraftGetAttachmentParams struct {
	// The ID of the inbox.
	InboxID string `path:"inbox_id" api:"required" json:"-"`
	// ID of draft.
	DraftID string `path:"draft_id" api:"required" json:"-"`
	paramObj
}

type InboxDraftSendParams struct {
	// The ID of the inbox.
	InboxID       string `path:"inbox_id" api:"required" json:"-"`
	UpdateMessage UpdateMessageParam
	paramObj
}

func (r InboxDraftSendParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateMessage)
}
func (r *InboxDraftSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
