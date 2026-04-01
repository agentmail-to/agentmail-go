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

// InboxThreadService contains methods and other services that help with
// interacting with the agentmail API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboxThreadService] method instead.
type InboxThreadService struct {
	Options []option.RequestOption
}

// NewInboxThreadService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInboxThreadService(opts ...option.RequestOption) (r InboxThreadService) {
	r = InboxThreadService{}
	r.Options = opts
	return
}

// **CLI:**
//
// ```bash
// agentmail inboxes:threads retrieve --inbox-id <inbox_id> --thread-id <thread_id>
// ```
func (r *InboxThreadService) Get(ctx context.Context, threadID string, query InboxThreadGetParams, opts ...option.RequestOption) (res *Thread, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if query.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/inboxes/%s/threads/%s", url.PathEscape(query.InboxID), url.PathEscape(threadID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail inboxes:threads list --inbox-id <inbox_id>
// ```
func (r *InboxThreadService) List(ctx context.Context, inboxID string, query InboxThreadListParams, opts ...option.RequestOption) (res *ListThreads, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/inboxes/%s/threads", url.PathEscape(inboxID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Moves the thread to trash by adding a trash label to all messages. If the thread
// is already in trash, it will be permanently deleted. Use `permanent=true` to
// force permanent deletion.
//
// **CLI:**
//
// ```bash
// agentmail inboxes:threads delete --inbox-id <inbox_id> --thread-id <thread_id>
// ```
func (r *InboxThreadService) Delete(ctx context.Context, threadID string, params InboxThreadDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if params.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return err
	}
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return err
	}
	path := fmt.Sprintf("v0/inboxes/%s/threads/%s", url.PathEscape(params.InboxID), url.PathEscape(threadID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// **CLI:**
//
// ```bash
// agentmail inboxes:threads get-attachment --inbox-id <inbox_id> --thread-id <thread_id> --attachment-id <attachment_id>
// ```
func (r *InboxThreadService) GetAttachment(ctx context.Context, attachmentID string, query InboxThreadGetAttachmentParams, opts ...option.RequestOption) (res *AttachmentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if query.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	if query.ThreadID == "" {
		err = errors.New("missing required thread_id parameter")
		return nil, err
	}
	if attachmentID == "" {
		err = errors.New("missing required attachment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/inboxes/%s/threads/%s/attachments/%s", url.PathEscape(query.InboxID), url.PathEscape(query.ThreadID), url.PathEscape(attachmentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ListThreads struct {
	// Number of items returned.
	Count int64 `json:"count" api:"required"`
	// Ordered by `timestamp` descending.
	Threads []ListThreadsThread `json:"threads" api:"required"`
	// Limit of number of items returned.
	Limit int64 `json:"limit" api:"nullable"`
	// Page token for pagination.
	NextPageToken string `json:"next_page_token" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count         respjson.Field
		Threads       respjson.Field
		Limit         respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListThreads) RawJSON() string { return r.JSON.raw }
func (r *ListThreads) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListThreadsThread struct {
	// Time at which thread was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The ID of the inbox.
	InboxID string `json:"inbox_id" api:"required"`
	// Labels of thread.
	Labels []string `json:"labels" api:"required"`
	// ID of last message in thread.
	LastMessageID string `json:"last_message_id" api:"required"`
	// Number of messages in thread.
	MessageCount int64 `json:"message_count" api:"required"`
	// Recipients in thread. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	Recipients []string `json:"recipients" api:"required"`
	// Senders in thread. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	Senders []string `json:"senders" api:"required"`
	// Size of thread in bytes.
	Size int64 `json:"size" api:"required"`
	// ID of thread.
	ThreadID string `json:"thread_id" api:"required"`
	// Timestamp of last sent or received message.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// Time at which thread was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Attachments in thread.
	Attachments []AttachmentFile `json:"attachments" api:"nullable"`
	// Text preview of last message in thread.
	Preview string `json:"preview" api:"nullable"`
	// Timestamp of last received message.
	ReceivedTimestamp time.Time `json:"received_timestamp" api:"nullable" format:"date-time"`
	// Timestamp of last sent message.
	SentTimestamp time.Time `json:"sent_timestamp" api:"nullable" format:"date-time"`
	// Subject of thread.
	Subject string `json:"subject" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt         respjson.Field
		InboxID           respjson.Field
		Labels            respjson.Field
		LastMessageID     respjson.Field
		MessageCount      respjson.Field
		Recipients        respjson.Field
		Senders           respjson.Field
		Size              respjson.Field
		ThreadID          respjson.Field
		Timestamp         respjson.Field
		UpdatedAt         respjson.Field
		Attachments       respjson.Field
		Preview           respjson.Field
		ReceivedTimestamp respjson.Field
		SentTimestamp     respjson.Field
		Subject           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListThreadsThread) RawJSON() string { return r.JSON.raw }
func (r *ListThreadsThread) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Thread struct {
	// Time at which thread was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The ID of the inbox.
	InboxID string `json:"inbox_id" api:"required"`
	// Labels of thread.
	Labels []string `json:"labels" api:"required"`
	// ID of last message in thread.
	LastMessageID string `json:"last_message_id" api:"required"`
	// Number of messages in thread.
	MessageCount int64 `json:"message_count" api:"required"`
	// Messages in thread. Ordered by `timestamp` ascending.
	Messages []Message `json:"messages" api:"required"`
	// Recipients in thread. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	Recipients []string `json:"recipients" api:"required"`
	// Senders in thread. In format `username@domain.com` or
	// `Display Name <username@domain.com>`.
	Senders []string `json:"senders" api:"required"`
	// Size of thread in bytes.
	Size int64 `json:"size" api:"required"`
	// ID of thread.
	ThreadID string `json:"thread_id" api:"required"`
	// Timestamp of last sent or received message.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// Time at which thread was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Attachments in thread.
	Attachments []AttachmentFile `json:"attachments" api:"nullable"`
	// Text preview of last message in thread.
	Preview string `json:"preview" api:"nullable"`
	// Timestamp of last received message.
	ReceivedTimestamp time.Time `json:"received_timestamp" api:"nullable" format:"date-time"`
	// Timestamp of last sent message.
	SentTimestamp time.Time `json:"sent_timestamp" api:"nullable" format:"date-time"`
	// Subject of thread.
	Subject string `json:"subject" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt         respjson.Field
		InboxID           respjson.Field
		Labels            respjson.Field
		LastMessageID     respjson.Field
		MessageCount      respjson.Field
		Messages          respjson.Field
		Recipients        respjson.Field
		Senders           respjson.Field
		Size              respjson.Field
		ThreadID          respjson.Field
		Timestamp         respjson.Field
		UpdatedAt         respjson.Field
		Attachments       respjson.Field
		Preview           respjson.Field
		ReceivedTimestamp respjson.Field
		SentTimestamp     respjson.Field
		Subject           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Thread) RawJSON() string { return r.JSON.raw }
func (r *Thread) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxThreadGetParams struct {
	// The ID of the inbox.
	InboxID string `path:"inbox_id" api:"required" json:"-"`
	paramObj
}

type InboxThreadListParams struct {
	// Timestamp after which to filter by.
	After param.Opt[time.Time] `query:"after,omitzero" format:"date-time" json:"-"`
	// Sort in ascending temporal order.
	Ascending param.Opt[bool] `query:"ascending,omitzero" json:"-"`
	// Timestamp before which to filter by.
	Before param.Opt[time.Time] `query:"before,omitzero" format:"date-time" json:"-"`
	// Include blocked in results.
	IncludeBlocked param.Opt[bool] `query:"include_blocked,omitzero" json:"-"`
	// Include spam in results.
	IncludeSpam param.Opt[bool] `query:"include_spam,omitzero" json:"-"`
	// Include trash in results.
	IncludeTrash param.Opt[bool] `query:"include_trash,omitzero" json:"-"`
	// Limit of number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page token for pagination.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	// Labels to filter by.
	Labels []string `query:"labels,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxThreadListParams]'s query parameters as `url.Values`.
func (r InboxThreadListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InboxThreadDeleteParams struct {
	// The ID of the inbox.
	InboxID string `path:"inbox_id" api:"required" json:"-"`
	// If true, permanently delete the thread instead of moving to trash.
	Permanent param.Opt[bool] `query:"permanent,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxThreadDeleteParams]'s query parameters as
// `url.Values`.
func (r InboxThreadDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InboxThreadGetAttachmentParams struct {
	// The ID of the inbox.
	InboxID string `path:"inbox_id" api:"required" json:"-"`
	// ID of thread.
	ThreadID string `path:"thread_id" api:"required" json:"-"`
	paramObj
}
