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

// PodAPIKeyService contains methods and other services that help with interacting
// with the agentmail API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPodAPIKeyService] method instead.
type PodAPIKeyService struct {
	Options []option.RequestOption
}

// NewPodAPIKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPodAPIKeyService(opts ...option.RequestOption) (r PodAPIKeyService) {
	r = PodAPIKeyService{}
	r.Options = opts
	return
}

// **CLI:**
//
// ```bash
// agentmail pods:api-keys create --pod-id <pod_id> --name "My Key"
// ```
func (r *PodAPIKeyService) New(ctx context.Context, podID string, body PodAPIKeyNewParams, opts ...option.RequestOption) (res *PodAPIKeyNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if podID == "" {
		err = errors.New("missing required pod_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/pods/%s/api-keys", url.PathEscape(podID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail pods:api-keys list --pod-id <pod_id>
// ```
func (r *PodAPIKeyService) List(ctx context.Context, podID string, query PodAPIKeyListParams, opts ...option.RequestOption) (res *PodAPIKeyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if podID == "" {
		err = errors.New("missing required pod_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/pods/%s/api-keys", url.PathEscape(podID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail pods:api-keys delete --pod-id <pod_id> --api-key-id <api_key_id>
// ```
func (r *PodAPIKeyService) Delete(ctx context.Context, apiKeyID string, body PodAPIKeyDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if body.PodID == "" {
		err = errors.New("missing required pod_id parameter")
		return err
	}
	if apiKeyID == "" {
		err = errors.New("missing required api_key_id parameter")
		return err
	}
	path := fmt.Sprintf("v0/pods/%s/api-keys/%s", url.PathEscape(body.PodID), url.PathEscape(apiKeyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type PodAPIKeyNewResponse struct {
	// API key.
	APIKey string `json:"api_key" api:"required"`
	// ID of api key.
	APIKeyID string `json:"api_key_id" api:"required"`
	// Time at which api key was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Name of api key.
	Name string `json:"name" api:"required"`
	// Prefix of api key.
	Prefix string `json:"prefix" api:"required"`
	// Inbox ID the api key is scoped to.
	InboxID string `json:"inbox_id" api:"nullable"`
	// Granular permissions for the API key. When ommitted all permissions are granted.
	// Otherwise, only permissions set to true are granted.
	Permissions PodAPIKeyNewResponsePermissions `json:"permissions" api:"nullable"`
	// Pod ID the api key is scoped to.
	PodID string `json:"pod_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey      respjson.Field
		APIKeyID    respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Prefix      respjson.Field
		InboxID     respjson.Field
		Permissions respjson.Field
		PodID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PodAPIKeyNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PodAPIKeyNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Granular permissions for the API key. When ommitted all permissions are granted.
// Otherwise, only permissions set to true are granted.
type PodAPIKeyNewResponsePermissions struct {
	// Create API keys.
	APIKeyCreate bool `json:"api_key_create" api:"nullable"`
	// Delete API keys.
	APIKeyDelete bool `json:"api_key_delete" api:"nullable"`
	// Read API keys.
	APIKeyRead bool `json:"api_key_read" api:"nullable"`
	// Create domains.
	DomainCreate bool `json:"domain_create" api:"nullable"`
	// Delete domains.
	DomainDelete bool `json:"domain_delete" api:"nullable"`
	// Read domain details.
	DomainRead bool `json:"domain_read" api:"nullable"`
	// Update domains.
	DomainUpdate bool `json:"domain_update" api:"nullable"`
	// Create drafts.
	DraftCreate bool `json:"draft_create" api:"nullable"`
	// Delete drafts.
	DraftDelete bool `json:"draft_delete" api:"nullable"`
	// Read drafts.
	DraftRead bool `json:"draft_read" api:"nullable"`
	// Send drafts.
	DraftSend bool `json:"draft_send" api:"nullable"`
	// Update drafts.
	DraftUpdate bool `json:"draft_update" api:"nullable"`
	// Create new inboxes.
	InboxCreate bool `json:"inbox_create" api:"nullable"`
	// Delete inboxes.
	InboxDelete bool `json:"inbox_delete" api:"nullable"`
	// Read inbox details.
	InboxRead bool `json:"inbox_read" api:"nullable"`
	// Update inbox settings.
	InboxUpdate bool `json:"inbox_update" api:"nullable"`
	// Access messages labeled blocked.
	LabelBlockedRead bool `json:"label_blocked_read" api:"nullable"`
	// Access messages labeled spam.
	LabelSpamRead bool `json:"label_spam_read" api:"nullable"`
	// Access messages labeled trash.
	LabelTrashRead bool `json:"label_trash_read" api:"nullable"`
	// Create list entries.
	ListEntryCreate bool `json:"list_entry_create" api:"nullable"`
	// Delete list entries.
	ListEntryDelete bool `json:"list_entry_delete" api:"nullable"`
	// Read list entries.
	ListEntryRead bool `json:"list_entry_read" api:"nullable"`
	// Read messages.
	MessageRead bool `json:"message_read" api:"nullable"`
	// Send messages.
	MessageSend bool `json:"message_send" api:"nullable"`
	// Update message labels.
	MessageUpdate bool `json:"message_update" api:"nullable"`
	// Read metrics.
	MetricsRead bool `json:"metrics_read" api:"nullable"`
	// Create pods.
	PodCreate bool `json:"pod_create" api:"nullable"`
	// Delete pods.
	PodDelete bool `json:"pod_delete" api:"nullable"`
	// Read pods.
	PodRead bool `json:"pod_read" api:"nullable"`
	// Delete threads.
	ThreadDelete bool `json:"thread_delete" api:"nullable"`
	// Read threads.
	ThreadRead bool `json:"thread_read" api:"nullable"`
	// Create webhooks.
	WebhookCreate bool `json:"webhook_create" api:"nullable"`
	// Delete webhooks.
	WebhookDelete bool `json:"webhook_delete" api:"nullable"`
	// Read webhook configurations.
	WebhookRead bool `json:"webhook_read" api:"nullable"`
	// Update webhooks.
	WebhookUpdate bool `json:"webhook_update" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKeyCreate     respjson.Field
		APIKeyDelete     respjson.Field
		APIKeyRead       respjson.Field
		DomainCreate     respjson.Field
		DomainDelete     respjson.Field
		DomainRead       respjson.Field
		DomainUpdate     respjson.Field
		DraftCreate      respjson.Field
		DraftDelete      respjson.Field
		DraftRead        respjson.Field
		DraftSend        respjson.Field
		DraftUpdate      respjson.Field
		InboxCreate      respjson.Field
		InboxDelete      respjson.Field
		InboxRead        respjson.Field
		InboxUpdate      respjson.Field
		LabelBlockedRead respjson.Field
		LabelSpamRead    respjson.Field
		LabelTrashRead   respjson.Field
		ListEntryCreate  respjson.Field
		ListEntryDelete  respjson.Field
		ListEntryRead    respjson.Field
		MessageRead      respjson.Field
		MessageSend      respjson.Field
		MessageUpdate    respjson.Field
		MetricsRead      respjson.Field
		PodCreate        respjson.Field
		PodDelete        respjson.Field
		PodRead          respjson.Field
		ThreadDelete     respjson.Field
		ThreadRead       respjson.Field
		WebhookCreate    respjson.Field
		WebhookDelete    respjson.Field
		WebhookRead      respjson.Field
		WebhookUpdate    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PodAPIKeyNewResponsePermissions) RawJSON() string { return r.JSON.raw }
func (r *PodAPIKeyNewResponsePermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PodAPIKeyListResponse struct {
	// Ordered by `created_at` descending.
	APIKeys []PodAPIKeyListResponseAPIKey `json:"api_keys" api:"required"`
	// Number of items returned.
	Count int64 `json:"count" api:"required"`
	// Page token for pagination.
	NextPageToken string `json:"next_page_token" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKeys       respjson.Field
		Count         respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PodAPIKeyListResponse) RawJSON() string { return r.JSON.raw }
func (r *PodAPIKeyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PodAPIKeyListResponseAPIKey struct {
	// ID of api key.
	APIKeyID string `json:"api_key_id" api:"required"`
	// Time at which api key was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Name of api key.
	Name string `json:"name" api:"required"`
	// Prefix of api key.
	Prefix string `json:"prefix" api:"required"`
	// Inbox ID the api key is scoped to. If set, the key can only access resources
	// within this inbox.
	InboxID string `json:"inbox_id" api:"nullable"`
	// Granular permissions for the API key. When ommitted all permissions are granted.
	// Otherwise, only permissions set to true are granted.
	Permissions PodAPIKeyListResponseAPIKeyPermissions `json:"permissions" api:"nullable"`
	// Pod ID the api key is scoped to. If set, the key can only access resources
	// within this pod.
	PodID string `json:"pod_id" api:"nullable"`
	// Time at which api key was last used.
	UsedAt time.Time `json:"used_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKeyID    respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Prefix      respjson.Field
		InboxID     respjson.Field
		Permissions respjson.Field
		PodID       respjson.Field
		UsedAt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PodAPIKeyListResponseAPIKey) RawJSON() string { return r.JSON.raw }
func (r *PodAPIKeyListResponseAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Granular permissions for the API key. When ommitted all permissions are granted.
// Otherwise, only permissions set to true are granted.
type PodAPIKeyListResponseAPIKeyPermissions struct {
	// Create API keys.
	APIKeyCreate bool `json:"api_key_create" api:"nullable"`
	// Delete API keys.
	APIKeyDelete bool `json:"api_key_delete" api:"nullable"`
	// Read API keys.
	APIKeyRead bool `json:"api_key_read" api:"nullable"`
	// Create domains.
	DomainCreate bool `json:"domain_create" api:"nullable"`
	// Delete domains.
	DomainDelete bool `json:"domain_delete" api:"nullable"`
	// Read domain details.
	DomainRead bool `json:"domain_read" api:"nullable"`
	// Update domains.
	DomainUpdate bool `json:"domain_update" api:"nullable"`
	// Create drafts.
	DraftCreate bool `json:"draft_create" api:"nullable"`
	// Delete drafts.
	DraftDelete bool `json:"draft_delete" api:"nullable"`
	// Read drafts.
	DraftRead bool `json:"draft_read" api:"nullable"`
	// Send drafts.
	DraftSend bool `json:"draft_send" api:"nullable"`
	// Update drafts.
	DraftUpdate bool `json:"draft_update" api:"nullable"`
	// Create new inboxes.
	InboxCreate bool `json:"inbox_create" api:"nullable"`
	// Delete inboxes.
	InboxDelete bool `json:"inbox_delete" api:"nullable"`
	// Read inbox details.
	InboxRead bool `json:"inbox_read" api:"nullable"`
	// Update inbox settings.
	InboxUpdate bool `json:"inbox_update" api:"nullable"`
	// Access messages labeled blocked.
	LabelBlockedRead bool `json:"label_blocked_read" api:"nullable"`
	// Access messages labeled spam.
	LabelSpamRead bool `json:"label_spam_read" api:"nullable"`
	// Access messages labeled trash.
	LabelTrashRead bool `json:"label_trash_read" api:"nullable"`
	// Create list entries.
	ListEntryCreate bool `json:"list_entry_create" api:"nullable"`
	// Delete list entries.
	ListEntryDelete bool `json:"list_entry_delete" api:"nullable"`
	// Read list entries.
	ListEntryRead bool `json:"list_entry_read" api:"nullable"`
	// Read messages.
	MessageRead bool `json:"message_read" api:"nullable"`
	// Send messages.
	MessageSend bool `json:"message_send" api:"nullable"`
	// Update message labels.
	MessageUpdate bool `json:"message_update" api:"nullable"`
	// Read metrics.
	MetricsRead bool `json:"metrics_read" api:"nullable"`
	// Create pods.
	PodCreate bool `json:"pod_create" api:"nullable"`
	// Delete pods.
	PodDelete bool `json:"pod_delete" api:"nullable"`
	// Read pods.
	PodRead bool `json:"pod_read" api:"nullable"`
	// Delete threads.
	ThreadDelete bool `json:"thread_delete" api:"nullable"`
	// Read threads.
	ThreadRead bool `json:"thread_read" api:"nullable"`
	// Create webhooks.
	WebhookCreate bool `json:"webhook_create" api:"nullable"`
	// Delete webhooks.
	WebhookDelete bool `json:"webhook_delete" api:"nullable"`
	// Read webhook configurations.
	WebhookRead bool `json:"webhook_read" api:"nullable"`
	// Update webhooks.
	WebhookUpdate bool `json:"webhook_update" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKeyCreate     respjson.Field
		APIKeyDelete     respjson.Field
		APIKeyRead       respjson.Field
		DomainCreate     respjson.Field
		DomainDelete     respjson.Field
		DomainRead       respjson.Field
		DomainUpdate     respjson.Field
		DraftCreate      respjson.Field
		DraftDelete      respjson.Field
		DraftRead        respjson.Field
		DraftSend        respjson.Field
		DraftUpdate      respjson.Field
		InboxCreate      respjson.Field
		InboxDelete      respjson.Field
		InboxRead        respjson.Field
		InboxUpdate      respjson.Field
		LabelBlockedRead respjson.Field
		LabelSpamRead    respjson.Field
		LabelTrashRead   respjson.Field
		ListEntryCreate  respjson.Field
		ListEntryDelete  respjson.Field
		ListEntryRead    respjson.Field
		MessageRead      respjson.Field
		MessageSend      respjson.Field
		MessageUpdate    respjson.Field
		MetricsRead      respjson.Field
		PodCreate        respjson.Field
		PodDelete        respjson.Field
		PodRead          respjson.Field
		ThreadDelete     respjson.Field
		ThreadRead       respjson.Field
		WebhookCreate    respjson.Field
		WebhookDelete    respjson.Field
		WebhookRead      respjson.Field
		WebhookUpdate    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PodAPIKeyListResponseAPIKeyPermissions) RawJSON() string { return r.JSON.raw }
func (r *PodAPIKeyListResponseAPIKeyPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PodAPIKeyNewParams struct {
	// Name of api key.
	Name param.Opt[string] `json:"name,omitzero"`
	// Granular permissions for the API key. When ommitted all permissions are granted.
	// Otherwise, only permissions set to true are granted.
	Permissions PodAPIKeyNewParamsPermissions `json:"permissions,omitzero"`
	paramObj
}

func (r PodAPIKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PodAPIKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PodAPIKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Granular permissions for the API key. When ommitted all permissions are granted.
// Otherwise, only permissions set to true are granted.
type PodAPIKeyNewParamsPermissions struct {
	// Create API keys.
	APIKeyCreate param.Opt[bool] `json:"api_key_create,omitzero"`
	// Delete API keys.
	APIKeyDelete param.Opt[bool] `json:"api_key_delete,omitzero"`
	// Read API keys.
	APIKeyRead param.Opt[bool] `json:"api_key_read,omitzero"`
	// Create domains.
	DomainCreate param.Opt[bool] `json:"domain_create,omitzero"`
	// Delete domains.
	DomainDelete param.Opt[bool] `json:"domain_delete,omitzero"`
	// Read domain details.
	DomainRead param.Opt[bool] `json:"domain_read,omitzero"`
	// Update domains.
	DomainUpdate param.Opt[bool] `json:"domain_update,omitzero"`
	// Create drafts.
	DraftCreate param.Opt[bool] `json:"draft_create,omitzero"`
	// Delete drafts.
	DraftDelete param.Opt[bool] `json:"draft_delete,omitzero"`
	// Read drafts.
	DraftRead param.Opt[bool] `json:"draft_read,omitzero"`
	// Send drafts.
	DraftSend param.Opt[bool] `json:"draft_send,omitzero"`
	// Update drafts.
	DraftUpdate param.Opt[bool] `json:"draft_update,omitzero"`
	// Create new inboxes.
	InboxCreate param.Opt[bool] `json:"inbox_create,omitzero"`
	// Delete inboxes.
	InboxDelete param.Opt[bool] `json:"inbox_delete,omitzero"`
	// Read inbox details.
	InboxRead param.Opt[bool] `json:"inbox_read,omitzero"`
	// Update inbox settings.
	InboxUpdate param.Opt[bool] `json:"inbox_update,omitzero"`
	// Access messages labeled blocked.
	LabelBlockedRead param.Opt[bool] `json:"label_blocked_read,omitzero"`
	// Access messages labeled spam.
	LabelSpamRead param.Opt[bool] `json:"label_spam_read,omitzero"`
	// Access messages labeled trash.
	LabelTrashRead param.Opt[bool] `json:"label_trash_read,omitzero"`
	// Create list entries.
	ListEntryCreate param.Opt[bool] `json:"list_entry_create,omitzero"`
	// Delete list entries.
	ListEntryDelete param.Opt[bool] `json:"list_entry_delete,omitzero"`
	// Read list entries.
	ListEntryRead param.Opt[bool] `json:"list_entry_read,omitzero"`
	// Read messages.
	MessageRead param.Opt[bool] `json:"message_read,omitzero"`
	// Send messages.
	MessageSend param.Opt[bool] `json:"message_send,omitzero"`
	// Update message labels.
	MessageUpdate param.Opt[bool] `json:"message_update,omitzero"`
	// Read metrics.
	MetricsRead param.Opt[bool] `json:"metrics_read,omitzero"`
	// Create pods.
	PodCreate param.Opt[bool] `json:"pod_create,omitzero"`
	// Delete pods.
	PodDelete param.Opt[bool] `json:"pod_delete,omitzero"`
	// Read pods.
	PodRead param.Opt[bool] `json:"pod_read,omitzero"`
	// Delete threads.
	ThreadDelete param.Opt[bool] `json:"thread_delete,omitzero"`
	// Read threads.
	ThreadRead param.Opt[bool] `json:"thread_read,omitzero"`
	// Create webhooks.
	WebhookCreate param.Opt[bool] `json:"webhook_create,omitzero"`
	// Delete webhooks.
	WebhookDelete param.Opt[bool] `json:"webhook_delete,omitzero"`
	// Read webhook configurations.
	WebhookRead param.Opt[bool] `json:"webhook_read,omitzero"`
	// Update webhooks.
	WebhookUpdate param.Opt[bool] `json:"webhook_update,omitzero"`
	paramObj
}

func (r PodAPIKeyNewParamsPermissions) MarshalJSON() (data []byte, err error) {
	type shadow PodAPIKeyNewParamsPermissions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PodAPIKeyNewParamsPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PodAPIKeyListParams struct {
	// Maximum number of items to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page token for pagination.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PodAPIKeyListParams]'s query parameters as `url.Values`.
func (r PodAPIKeyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PodAPIKeyDeleteParams struct {
	// ID of pod.
	PodID string `path:"pod_id" api:"required" json:"-"`
	paramObj
}
