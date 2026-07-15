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

// ThreadService contains methods and other services that help with interacting
// with the agentmail API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreadService] method instead.
type ThreadService struct {
	Options []option.RequestOption
}

// NewThreadService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewThreadService(opts ...option.RequestOption) (r ThreadService) {
	r = ThreadService{}
	r.Options = opts
	return
}

// Lists threads, most recent first. Pass `senders`, `recipients`, or `subject` to
// filter by substring. Filtered requests are served by search, which caps `limit`
// at 100. For relevance-ranked full-text search across senders, recipients,
// subject, and message body, use `Search Threads`.
//
// **CLI:**
//
// ```bash
// agentmail threads list
// ```
func (r *ThreadService) List(ctx context.Context, query ThreadListParams, opts ...option.RequestOption) (res *ListThreads, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	path := "v0/threads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Permanently deletes a thread and all of its messages.
//
// **CLI:**
//
// ```bash
// agentmail threads delete --thread-id <thread_id>
// ```
func (r *ThreadService) Delete(ctx context.Context, threadID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return err
	}
	path := fmt.Sprintf("v0/threads/%s", url.PathEscape(threadID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// **CLI:**
//
// ```bash
// agentmail threads get --thread-id <thread_id>
// ```
func (r *ThreadService) Get(ctx context.Context, threadID string, opts ...option.RequestOption) (res *Thread, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/threads/%s", url.PathEscape(threadID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail threads get-attachment --thread-id <thread_id> --attachment-id <attachment_id>
// ```
func (r *ThreadService) GetAttachment(ctx context.Context, attachmentID string, query ThreadGetAttachmentParams, opts ...option.RequestOption) (res *AttachmentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if query.ThreadID == "" {
		err = errors.New("missing required thread_id parameter")
		return nil, err
	}
	if attachmentID == "" {
		err = errors.New("missing required attachment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/threads/%s/attachments/%s", url.PathEscape(query.ThreadID), url.PathEscape(attachmentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Full-text search across threads in the organization, ranked by relevance. The
// query is matched against senders, recipients, and subject (substring) and the
// message body (tokenized full text). Spam, trash, blocked, and unauthenticated
// threads are always excluded. `limit` cannot exceed 100.
func (r *ThreadService) Search(ctx context.Context, query ThreadSearchParams, opts ...option.RequestOption) (res *ThreadSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	path := "v0/threads/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ThreadSearchResponse struct {
	// Number of items returned.
	Count int64 `json:"count" api:"required"`
	// Ordered by relevance, best match first.
	Threads []ThreadSearchResponseThread `json:"threads" api:"required"`
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
func (r ThreadSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *ThreadSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ThreadSearchResponseThread struct {
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
	// Matched fragments per field. Present only when the query matched an indexed
	// field.
	Highlights ThreadSearchResponseThreadHighlights `json:"highlights" api:"nullable"`
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
		Highlights        respjson.Field
		Preview           respjson.Field
		ReceivedTimestamp respjson.Field
		SentTimestamp     respjson.Field
		Subject           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ThreadSearchResponseThread) RawJSON() string { return r.JSON.raw }
func (r *ThreadSearchResponseThread) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Matched fragments per field. Present only when the query matched an indexed
// field.
type ThreadSearchResponseThreadHighlights struct {
	// Matched fragments from a sender address in the thread.
	From []string `json:"from" api:"nullable"`
	// Matched fragments from a recipient address in the thread (to, cc, or bcc).
	Recipients []string `json:"recipients" api:"nullable"`
	// Matched fragments from the subject.
	Subject []string `json:"subject" api:"nullable"`
	// Matched fragments from a message body in the thread.
	Text []string `json:"text" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		Recipients  respjson.Field
		Subject     respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ThreadSearchResponseThreadHighlights) RawJSON() string { return r.JSON.raw }
func (r *ThreadSearchResponseThreadHighlights) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ThreadListParams struct {
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
	// Include unauthenticated in results.
	IncludeUnauthenticated param.Opt[bool] `query:"include_unauthenticated,omitzero" json:"-"`
	// Limit of number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page token for pagination.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	// Labels to filter by.
	Labels []string `query:"labels,omitzero" json:"-"`
	// Filter to threads whose recipients contain this value (substring match).
	// Repeatable; all values must match.
	Recipients []string `query:"recipients,omitzero" json:"-"`
	// Filter to threads whose senders contain this value (substring match).
	// Repeatable; all values must match.
	Senders []string `query:"senders,omitzero" json:"-"`
	// Filter to threads whose subject contains this value (substring match).
	// Repeatable; all values must match.
	Subject []string `query:"subject,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ThreadListParams]'s query parameters as `url.Values`.
func (r ThreadListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ThreadGetAttachmentParams struct {
	// ID of thread.
	ThreadID string `path:"thread_id" api:"required" json:"-"`
	paramObj
}

type ThreadSearchParams struct {
	// Full-text search query. Matched against the sender, recipients, and subject
	// (substring) and the message body (tokenized full text).
	Q string `query:"q" api:"required" json:"-"`
	// Timestamp after which to filter by.
	After param.Opt[time.Time] `query:"after,omitzero" format:"date-time" json:"-"`
	// Timestamp before which to filter by.
	Before param.Opt[time.Time] `query:"before,omitzero" format:"date-time" json:"-"`
	// Limit of number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page token for pagination.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ThreadSearchParams]'s query parameters as `url.Values`.
func (r ThreadSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
