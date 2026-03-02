// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package agentmail

import (
	"context"
	"encoding/json"
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

// InboxMessageService contains methods and other services that help with
// interacting with the agentmail API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboxMessageService] method instead.
type InboxMessageService struct {
	Options []option.RequestOption
}

// NewInboxMessageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInboxMessageService(opts ...option.RequestOption) (r InboxMessageService) {
	r = InboxMessageService{}
	r.Options = opts
	return
}

// Get Message
func (r *InboxMessageService) Get(ctx context.Context, messageID string, query InboxMessageGetParams, opts ...option.RequestOption) (res *Message, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if query.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v0/inboxes/%s/messages/%s", url.PathEscape(query.InboxID), url.PathEscape(messageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Message
func (r *InboxMessageService) Update(ctx context.Context, messageID string, params InboxMessageUpdateParams, opts ...option.RequestOption) (res *Message, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if params.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v0/inboxes/%s/messages/%s", url.PathEscape(params.InboxID), url.PathEscape(messageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List Messages
func (r *InboxMessageService) List(ctx context.Context, inboxID string, query InboxMessageListParams, opts ...option.RequestOption) (res *InboxMessageListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return
	}
	path := fmt.Sprintf("v0/inboxes/%s/messages", url.PathEscape(inboxID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Forward Message
func (r *InboxMessageService) Forward(ctx context.Context, messageID string, params InboxMessageForwardParams, opts ...option.RequestOption) (res *SendMessageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if params.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v0/inboxes/%s/messages/%s/forward", url.PathEscape(params.InboxID), url.PathEscape(messageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get Attachment
func (r *InboxMessageService) GetAttachment(ctx context.Context, attachmentID string, query InboxMessageGetAttachmentParams, opts ...option.RequestOption) (res *AttachmentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if query.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return
	}
	if query.MessageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	if attachmentID == "" {
		err = errors.New("missing required attachment_id parameter")
		return
	}
	path := fmt.Sprintf("v0/inboxes/%s/messages/%s/attachments/%s", url.PathEscape(query.InboxID), url.PathEscape(query.MessageID), url.PathEscape(attachmentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get Raw Message
func (r *InboxMessageService) GetRaw(ctx context.Context, messageID string, query InboxMessageGetRawParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if query.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v0/inboxes/%s/messages/%s/raw", url.PathEscape(query.InboxID), url.PathEscape(messageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Reply To Message
func (r *InboxMessageService) Reply(ctx context.Context, messageID string, params InboxMessageReplyParams, opts ...option.RequestOption) (res *SendMessageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if params.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v0/inboxes/%s/messages/%s/reply", url.PathEscape(params.InboxID), url.PathEscape(messageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Reply All Message
func (r *InboxMessageService) ReplyAll(ctx context.Context, messageID string, params InboxMessageReplyAllParams, opts ...option.RequestOption) (res *SendMessageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if params.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return
	}
	path := fmt.Sprintf("v0/inboxes/%s/messages/%s/reply-all", url.PathEscape(params.InboxID), url.PathEscape(messageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Send Message
func (r *InboxMessageService) Send(ctx context.Context, inboxID string, body InboxMessageSendParams, opts ...option.RequestOption) (res *SendMessageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return
	}
	path := fmt.Sprintf("v0/inboxes/%s/messages/send", url.PathEscape(inboxID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AddressesUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AddressesUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *AddressesUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Content disposition of attachment.
type AttachmentContentDisposition string

const (
	AttachmentContentDispositionInline     AttachmentContentDisposition = "inline"
	AttachmentContentDispositionAttachment AttachmentContentDisposition = "attachment"
)

type AttachmentResponse struct {
	// ID of attachment.
	AttachmentID string `json:"attachment_id" api:"required"`
	// URL to download the attachment.
	DownloadURL string `json:"download_url" api:"required"`
	// Time at which the download URL expires.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
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
		DownloadURL        respjson.Field
		ExpiresAt          respjson.Field
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
func (r AttachmentResponse) RawJSON() string { return r.JSON.raw }
func (r *AttachmentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Message struct {
	// Time at which message was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Address of sender. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	From string `json:"from" api:"required"`
	// ID of inbox.
	InboxID string `json:"inbox_id" api:"required"`
	// Labels of message.
	Labels []string `json:"labels" api:"required"`
	// ID of message.
	MessageID string `json:"message_id" api:"required"`
	// Size of message in bytes.
	Size int64 `json:"size" api:"required"`
	// ID of thread.
	ThreadID string `json:"thread_id" api:"required"`
	// Time at which message was sent or drafted.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// Addresses of recipients. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	To []string `json:"to" api:"required"`
	// Time at which message was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Attachments in message.
	Attachments []AttachmentFile `json:"attachments" api:"nullable"`
	// Addresses of BCC recipients. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	Bcc []string `json:"bcc" api:"nullable"`
	// Addresses of CC recipients. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	Cc []string `json:"cc" api:"nullable"`
	// Extracted new HTML content.
	ExtractedHTML string `json:"extracted_html" api:"nullable"`
	// Extracted new text content.
	ExtractedText string `json:"extracted_text" api:"nullable"`
	// Headers in message.
	Headers map[string]string `json:"headers" api:"nullable"`
	// HTML body of message.
	HTML string `json:"html" api:"nullable"`
	// ID of message being replied to.
	InReplyTo string `json:"in_reply_to" api:"nullable"`
	// Text preview of message.
	Preview string `json:"preview" api:"nullable"`
	// IDs of previous messages in thread.
	References []string `json:"references" api:"nullable"`
	// Reply-to addresses. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	ReplyTo []string `json:"reply_to" api:"nullable"`
	// Subject of message.
	Subject string `json:"subject" api:"nullable"`
	// Plain text body of message.
	Text string `json:"text" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt     respjson.Field
		From          respjson.Field
		InboxID       respjson.Field
		Labels        respjson.Field
		MessageID     respjson.Field
		Size          respjson.Field
		ThreadID      respjson.Field
		Timestamp     respjson.Field
		To            respjson.Field
		UpdatedAt     respjson.Field
		Attachments   respjson.Field
		Bcc           respjson.Field
		Cc            respjson.Field
		ExtractedHTML respjson.Field
		ExtractedText respjson.Field
		Headers       respjson.Field
		HTML          respjson.Field
		InReplyTo     respjson.Field
		Preview       respjson.Field
		References    respjson.Field
		ReplyTo       respjson.Field
		Subject       respjson.Field
		Text          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Message) RawJSON() string { return r.JSON.raw }
func (r *Message) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendAttachmentParam struct {
	// Base64 encoded content of attachment.
	Content param.Opt[string] `json:"content,omitzero"`
	// Content ID of attachment.
	ContentID param.Opt[string] `json:"content_id,omitzero"`
	// Content type of attachment.
	ContentType param.Opt[string] `json:"content_type,omitzero"`
	// Filename of attachment.
	Filename param.Opt[string] `json:"filename,omitzero"`
	// URL to the attachment.
	URL param.Opt[string] `json:"url,omitzero"`
	// Content disposition of attachment.
	//
	// Any of "inline", "attachment".
	ContentDisposition AttachmentContentDisposition `json:"content_disposition,omitzero"`
	paramObj
}

func (r SendAttachmentParam) MarshalJSON() (data []byte, err error) {
	type shadow SendAttachmentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendAttachmentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendMessageRequestParam struct {
	// HTML body of message.
	HTML param.Opt[string] `json:"html,omitzero"`
	// Subject of message.
	Subject param.Opt[string] `json:"subject,omitzero"`
	// Plain text body of message.
	Text param.Opt[string] `json:"text,omitzero"`
	// Attachments to include in message.
	Attachments []SendAttachmentParam `json:"attachments,omitzero"`
	// Headers to include in message.
	Headers map[string]string `json:"headers,omitzero"`
	// Labels of message.
	Labels []string `json:"labels,omitzero"`
	// BCC recipient address or addresses.
	Bcc AddressesUnionParam `json:"bcc,omitzero"`
	// CC recipient address or addresses.
	Cc AddressesUnionParam `json:"cc,omitzero"`
	// Reply-to address or addresses.
	ReplyTo AddressesUnionParam `json:"reply_to,omitzero"`
	// Recipient address or addresses.
	To AddressesUnionParam `json:"to,omitzero"`
	paramObj
}

func (r SendMessageRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxMessageListResponse struct {
	// Number of items returned.
	Count int64 `json:"count" api:"required"`
	// Ordered by `timestamp` descending.
	Messages []InboxMessageListResponseMessage `json:"messages" api:"required"`
	// Limit of number of items returned.
	Limit int64 `json:"limit" api:"nullable"`
	// Page token for pagination.
	NextPageToken string `json:"next_page_token" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count         respjson.Field
		Messages      respjson.Field
		Limit         respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboxMessageListResponse) RawJSON() string { return r.JSON.raw }
func (r *InboxMessageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxMessageListResponseMessage struct {
	// Time at which message was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Address of sender. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	From string `json:"from" api:"required"`
	// ID of inbox.
	InboxID string `json:"inbox_id" api:"required"`
	// Labels of message.
	Labels []string `json:"labels" api:"required"`
	// ID of message.
	MessageID string `json:"message_id" api:"required"`
	// Size of message in bytes.
	Size int64 `json:"size" api:"required"`
	// ID of thread.
	ThreadID string `json:"thread_id" api:"required"`
	// Time at which message was sent or drafted.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// Addresses of recipients. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	To []string `json:"to" api:"required"`
	// Time at which message was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Attachments in message.
	Attachments []AttachmentFile `json:"attachments" api:"nullable"`
	// Addresses of BCC recipients. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	Bcc []string `json:"bcc" api:"nullable"`
	// Addresses of CC recipients. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	Cc []string `json:"cc" api:"nullable"`
	// Headers in message.
	Headers map[string]string `json:"headers" api:"nullable"`
	// ID of message being replied to.
	InReplyTo string `json:"in_reply_to" api:"nullable"`
	// Text preview of message.
	Preview string `json:"preview" api:"nullable"`
	// IDs of previous messages in thread.
	References []string `json:"references" api:"nullable"`
	// Subject of message.
	Subject string `json:"subject" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		From        respjson.Field
		InboxID     respjson.Field
		Labels      respjson.Field
		MessageID   respjson.Field
		Size        respjson.Field
		ThreadID    respjson.Field
		Timestamp   respjson.Field
		To          respjson.Field
		UpdatedAt   respjson.Field
		Attachments respjson.Field
		Bcc         respjson.Field
		Cc          respjson.Field
		Headers     respjson.Field
		InReplyTo   respjson.Field
		Preview     respjson.Field
		References  respjson.Field
		Subject     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboxMessageListResponseMessage) RawJSON() string { return r.JSON.raw }
func (r *InboxMessageListResponseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxMessageGetParams struct {
	// ID of inbox.
	InboxID string `path:"inbox_id" api:"required" json:"-"`
	paramObj
}

type InboxMessageUpdateParams struct {
	// ID of inbox.
	InboxID       string `path:"inbox_id" api:"required" json:"-"`
	UpdateMessage UpdateMessageParam
	paramObj
}

func (r InboxMessageUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateMessage)
}
func (r *InboxMessageUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UpdateMessage)
}

type InboxMessageListParams struct {
	// Timestamp after which to filter by.
	After param.Opt[time.Time] `query:"after,omitzero" format:"date-time" json:"-"`
	// Sort in ascending temporal order.
	Ascending param.Opt[bool] `query:"ascending,omitzero" json:"-"`
	// Timestamp before which to filter by.
	Before param.Opt[time.Time] `query:"before,omitzero" format:"date-time" json:"-"`
	// Include spam in results.
	IncludeSpam param.Opt[bool] `query:"include_spam,omitzero" json:"-"`
	// Limit of number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page token for pagination.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	// Labels to filter by.
	Labels []string `query:"labels,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxMessageListParams]'s query parameters as `url.Values`.
func (r InboxMessageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InboxMessageForwardParams struct {
	// ID of inbox.
	InboxID            string `path:"inbox_id" api:"required" json:"-"`
	SendMessageRequest SendMessageRequestParam
	paramObj
}

func (r InboxMessageForwardParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SendMessageRequest)
}
func (r *InboxMessageForwardParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SendMessageRequest)
}

type InboxMessageGetAttachmentParams struct {
	// ID of inbox.
	InboxID string `path:"inbox_id" api:"required" json:"-"`
	// ID of message.
	MessageID string `path:"message_id" api:"required" json:"-"`
	paramObj
}

type InboxMessageGetRawParams struct {
	// ID of inbox.
	InboxID string `path:"inbox_id" api:"required" json:"-"`
	paramObj
}

type InboxMessageReplyParams struct {
	// ID of inbox.
	InboxID string `path:"inbox_id" api:"required" json:"-"`
	// HTML body of message.
	HTML param.Opt[string] `json:"html,omitzero"`
	// Reply to all recipients of the original message.
	ReplyAll param.Opt[bool] `json:"reply_all,omitzero"`
	// Plain text body of message.
	Text param.Opt[string] `json:"text,omitzero"`
	// Attachments to include in message.
	Attachments []SendAttachmentParam `json:"attachments,omitzero"`
	// Headers to include in message.
	Headers map[string]string `json:"headers,omitzero"`
	// Labels of message.
	Labels []string `json:"labels,omitzero"`
	// BCC recipient address or addresses.
	Bcc AddressesUnionParam `json:"bcc,omitzero"`
	// CC recipient address or addresses.
	Cc AddressesUnionParam `json:"cc,omitzero"`
	// Reply-to address or addresses.
	ReplyTo AddressesUnionParam `json:"reply_to,omitzero"`
	// Recipient address or addresses.
	To AddressesUnionParam `json:"to,omitzero"`
	paramObj
}

func (r InboxMessageReplyParams) MarshalJSON() (data []byte, err error) {
	type shadow InboxMessageReplyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboxMessageReplyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxMessageReplyAllParams struct {
	// ID of inbox.
	InboxID string `path:"inbox_id" api:"required" json:"-"`
	// HTML body of message.
	HTML param.Opt[string] `json:"html,omitzero"`
	// Plain text body of message.
	Text param.Opt[string] `json:"text,omitzero"`
	// Attachments to include in message.
	Attachments []SendAttachmentParam `json:"attachments,omitzero"`
	// Headers to include in message.
	Headers map[string]string `json:"headers,omitzero"`
	// Labels of message.
	Labels []string `json:"labels,omitzero"`
	// Reply-to address or addresses.
	ReplyTo AddressesUnionParam `json:"reply_to,omitzero"`
	paramObj
}

func (r InboxMessageReplyAllParams) MarshalJSON() (data []byte, err error) {
	type shadow InboxMessageReplyAllParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboxMessageReplyAllParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxMessageSendParams struct {
	SendMessageRequest SendMessageRequestParam
	paramObj
}

func (r InboxMessageSendParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SendMessageRequest)
}
func (r *InboxMessageSendParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SendMessageRequest)
}
