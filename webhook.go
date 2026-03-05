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

// WebhookService contains methods and other services that help with interacting
// with the agentmail API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.Options = opts
	return
}

// Create Webhook
func (r *WebhookService) New(ctx context.Context, body WebhookNewParams, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	path := "v0/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Webhook
func (r *WebhookService) Get(ctx context.Context, webhookID string, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("v0/webhooks/%s", url.PathEscape(webhookID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Webhook
func (r *WebhookService) Update(ctx context.Context, webhookID string, body WebhookUpdateParams, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("v0/webhooks/%s", url.PathEscape(webhookID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List Webhooks
func (r *WebhookService) List(ctx context.Context, query WebhookListParams, opts ...option.RequestOption) (res *WebhookListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	path := "v0/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete Webhook
func (r *WebhookService) Delete(ctx context.Context, webhookID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("v0/webhooks/%s", url.PathEscape(webhookID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type EventType string

const (
	EventTypeMessageReceived   EventType = "message.received"
	EventTypeMessageSent       EventType = "message.sent"
	EventTypeMessageDelivered  EventType = "message.delivered"
	EventTypeMessageBounced    EventType = "message.bounced"
	EventTypeMessageComplained EventType = "message.complained"
	EventTypeMessageRejected   EventType = "message.rejected"
	EventTypeDomainVerified    EventType = "domain.verified"
)

type Webhook struct {
	// Time at which webhook was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Webhook is enabled.
	Enabled bool `json:"enabled" api:"required"`
	// Secret for webhook signature verification.
	Secret string `json:"secret" api:"required"`
	// Time at which webhook was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// URL of webhook endpoint.
	URL string `json:"url" api:"required"`
	// ID of webhook.
	WebhookID string `json:"webhook_id" api:"required"`
	// Client ID of webhook.
	ClientID string `json:"client_id" api:"nullable"`
	// Event types for which to send events.
	EventTypes []EventType `json:"event_types" api:"nullable"`
	// Inboxes for which to send events. Maximum 10 per webhook.
	InboxIDs []string `json:"inbox_ids" api:"nullable"`
	// Pods for which to send events. Maximum 10 per webhook.
	PodIDs []string `json:"pod_ids" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Enabled     respjson.Field
		Secret      respjson.Field
		UpdatedAt   respjson.Field
		URL         respjson.Field
		WebhookID   respjson.Field
		ClientID    respjson.Field
		EventTypes  respjson.Field
		InboxIDs    respjson.Field
		PodIDs      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Webhook) RawJSON() string { return r.JSON.raw }
func (r *Webhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListResponse struct {
	// Number of items returned.
	Count int64 `json:"count" api:"required"`
	// Ordered by `created_at` descending.
	Webhooks []Webhook `json:"webhooks" api:"required"`
	// Limit of number of items returned.
	Limit int64 `json:"limit" api:"nullable"`
	// Page token for pagination.
	NextPageToken string `json:"next_page_token" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count         respjson.Field
		Webhooks      respjson.Field
		Limit         respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewParams struct {
	// Event types for which to send events.
	EventTypes []EventType `json:"event_types,omitzero" api:"required"`
	// URL of webhook endpoint.
	URL string `json:"url" api:"required"`
	// Client ID of webhook.
	ClientID param.Opt[string] `json:"client_id,omitzero"`
	// Inboxes for which to send events. Maximum 10 per webhook.
	InboxIDs []string `json:"inbox_ids,omitzero"`
	// Pods for which to send events. Maximum 10 per webhook.
	PodIDs []string `json:"pod_ids,omitzero"`
	paramObj
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateParams struct {
	// Inbox IDs to subscribe to the webhook.
	AddInboxIDs []string `json:"add_inbox_ids,omitzero"`
	// Pod IDs to subscribe to the webhook.
	AddPodIDs []string `json:"add_pod_ids,omitzero"`
	// Inbox IDs to unsubscribe from the webhook.
	RemoveInboxIDs []string `json:"remove_inbox_ids,omitzero"`
	// Pod IDs to unsubscribe from the webhook.
	RemovePodIDs []string `json:"remove_pod_ids,omitzero"`
	paramObj
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListParams struct {
	// Sort in ascending temporal order.
	Ascending param.Opt[bool] `query:"ascending,omitzero" json:"-"`
	// Limit of number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page token for pagination.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookListParams]'s query parameters as `url.Values`.
func (r WebhookListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
