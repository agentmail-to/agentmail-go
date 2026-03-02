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

	"github.com/stainless-sdks/agentmail-go/internal/apijson"
	"github.com/stainless-sdks/agentmail-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/agentmail-go/internal/encoding/json"
	"github.com/stainless-sdks/agentmail-go/internal/requestconfig"
	"github.com/stainless-sdks/agentmail-go/option"
	"github.com/stainless-sdks/agentmail-go/packages/param"
	"github.com/stainless-sdks/agentmail-go/packages/respjson"
)

// InboxService contains methods and other services that help with interacting with
// the agentmail API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboxService] method instead.
type InboxService struct {
	Options  []option.RequestOption
	Drafts   InboxDraftService
	Messages InboxMessageService
	Threads  InboxThreadService
}

// NewInboxService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInboxService(opts ...option.RequestOption) (r InboxService) {
	r = InboxService{}
	r.Options = opts
	r.Drafts = NewInboxDraftService(opts...)
	r.Messages = NewInboxMessageService(opts...)
	r.Threads = NewInboxThreadService(opts...)
	return
}

// Create Inbox
func (r *InboxService) New(ctx context.Context, body InboxNewParams, opts ...option.RequestOption) (res *Inbox, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	path := "v0/inboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Inbox
func (r *InboxService) Get(ctx context.Context, inboxID string, opts ...option.RequestOption) (res *Inbox, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return
	}
	path := fmt.Sprintf("v0/inboxes/%s", url.PathEscape(inboxID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Inbox
func (r *InboxService) Update(ctx context.Context, inboxID string, body InboxUpdateParams, opts ...option.RequestOption) (res *Inbox, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return
	}
	path := fmt.Sprintf("v0/inboxes/%s", url.PathEscape(inboxID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List Inboxes
func (r *InboxService) List(ctx context.Context, query InboxListParams, opts ...option.RequestOption) (res *ListInboxes, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	path := "v0/inboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete Inbox
func (r *InboxService) Delete(ctx context.Context, inboxID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return
	}
	path := fmt.Sprintf("v0/inboxes/%s", url.PathEscape(inboxID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// List Metrics
func (r *InboxService) ListMetrics(ctx context.Context, inboxID string, query InboxListMetricsParams, opts ...option.RequestOption) (res *ListMetrics, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return
	}
	path := fmt.Sprintf("v0/inboxes/%s/metrics", url.PathEscape(inboxID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CreateInboxParam struct {
	// Client ID of inbox.
	ClientID param.Opt[string] `json:"client_id,omitzero"`
	// Display name: `Display Name <username@domain.com>`.
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	// Domain of address. Must be verified domain. Defaults to `agentmail.to`.
	Domain param.Opt[string] `json:"domain,omitzero"`
	// Username of address. Randomly generated if not specified.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r CreateInboxParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateInboxParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateInboxParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Inbox struct {
	// Time at which inbox was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// ID of inbox.
	InboxID string `json:"inbox_id" api:"required"`
	// ID of pod.
	PodID string `json:"pod_id" api:"required"`
	// Time at which inbox was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Client ID of inbox.
	ClientID string `json:"client_id" api:"nullable"`
	// Display name: `Display Name <username@domain.com>`.
	DisplayName string `json:"display_name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		InboxID     respjson.Field
		PodID       respjson.Field
		UpdatedAt   respjson.Field
		ClientID    respjson.Field
		DisplayName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Inbox) RawJSON() string { return r.JSON.raw }
func (r *Inbox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListInboxes struct {
	// Number of items returned.
	Count int64 `json:"count" api:"required"`
	// Ordered by `created_at` descending.
	Inboxes []Inbox `json:"inboxes" api:"required"`
	// Limit of number of items returned.
	Limit int64 `json:"limit" api:"nullable"`
	// Page token for pagination.
	NextPageToken string `json:"next_page_token" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count         respjson.Field
		Inboxes       respjson.Field
		Limit         respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListInboxes) RawJSON() string { return r.JSON.raw }
func (r *ListInboxes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListMetrics struct {
	// Message metrics grouped by event type.
	Message ListMetricsMessage `json:"message" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListMetrics) RawJSON() string { return r.JSON.raw }
func (r *ListMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message metrics grouped by event type.
type ListMetricsMessage struct {
	// Timestamps when messages bounced.
	Bounced []time.Time `json:"bounced" api:"nullable" format:"date-time"`
	// Timestamps when messages received complaints.
	Complained []time.Time `json:"complained" api:"nullable" format:"date-time"`
	// Timestamps when messages were delayed.
	Delayed []time.Time `json:"delayed" api:"nullable" format:"date-time"`
	// Timestamps when messages were delivered.
	Delivered []time.Time `json:"delivered" api:"nullable" format:"date-time"`
	// Timestamps when messages were received.
	Received []time.Time `json:"received" api:"nullable" format:"date-time"`
	// Timestamps when messages were rejected.
	Rejected []time.Time `json:"rejected" api:"nullable" format:"date-time"`
	// Timestamps when messages were sent.
	Sent []time.Time `json:"sent" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bounced     respjson.Field
		Complained  respjson.Field
		Delayed     respjson.Field
		Delivered   respjson.Field
		Received    respjson.Field
		Rejected    respjson.Field
		Sent        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListMetricsMessage) RawJSON() string { return r.JSON.raw }
func (r *ListMetricsMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of metric event.
type MetricEventType string

const (
	MetricEventTypeMessageSent       MetricEventType = "message.sent"
	MetricEventTypeMessageDelivered  MetricEventType = "message.delivered"
	MetricEventTypeMessageBounced    MetricEventType = "message.bounced"
	MetricEventTypeMessageDelayed    MetricEventType = "message.delayed"
	MetricEventTypeMessageRejected   MetricEventType = "message.rejected"
	MetricEventTypeMessageComplained MetricEventType = "message.complained"
	MetricEventTypeMessageReceived   MetricEventType = "message.received"
)

type InboxNewParams struct {
	CreateInbox CreateInboxParam
	paramObj
}

func (r InboxNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateInbox)
}
func (r *InboxNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateInbox)
}

type InboxUpdateParams struct {
	// Display name: `Display Name <username@domain.com>`.
	DisplayName string `json:"display_name" api:"required"`
	paramObj
}

func (r InboxUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow InboxUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboxUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxListParams struct {
	// Limit of number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page token for pagination.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxListParams]'s query parameters as `url.Values`.
func (r InboxListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InboxListMetricsParams struct {
	// End timestamp for the metrics query range.
	EndTimestamp time.Time `query:"end_timestamp" api:"required" format:"date-time" json:"-"`
	// Start timestamp for the metrics query range.
	StartTimestamp time.Time `query:"start_timestamp" api:"required" format:"date-time" json:"-"`
	// List of metric event types to filter by.
	EventTypes []MetricEventType `query:"event_types,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxListMetricsParams]'s query parameters as `url.Values`.
func (r InboxListMetricsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
