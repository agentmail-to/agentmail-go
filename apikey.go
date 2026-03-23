// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package agentmail

import (
	"context"
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

// APIKeyService contains methods and other services that help with interacting
// with the agentmail API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIKeyService] method instead.
type APIKeyService struct {
	Options []option.RequestOption
}

// NewAPIKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAPIKeyService(opts ...option.RequestOption) (r APIKeyService) {
	r = APIKeyService{}
	r.Options = opts
	return
}

// Create API Key
func (r *APIKeyService) New(ctx context.Context, body APIKeyNewParams, opts ...option.RequestOption) (res *APIKeyNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	path := "v0/api-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List API Keys
func (r *APIKeyService) List(ctx context.Context, query APIKeyListParams, opts ...option.RequestOption) (res *APIKeyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	path := "v0/api-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type APIKeyNewResponse struct {
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
	Permissions APIKeyNewResponsePermissions `json:"permissions" api:"nullable"`
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
func (r APIKeyNewResponse) RawJSON() string { return r.JSON.raw }
func (r *APIKeyNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Granular permissions for the API key. When ommitted all permissions are granted.
// Otherwise, only permissions set to true are granted.
type APIKeyNewResponsePermissions struct {
	// Create API keys.
	CreateAPIKey bool `json:"create_api_key" api:"nullable"`
	// Create domains.
	CreateDomain bool `json:"create_domain" api:"nullable"`
	// Create drafts.
	CreateDraft bool `json:"create_draft" api:"nullable"`
	// Create new inboxes.
	CreateInbox bool `json:"create_inbox" api:"nullable"`
	// Create list entries.
	CreateListEntry bool `json:"create_list_entry" api:"nullable"`
	// Create pods.
	CreatePod bool `json:"create_pod" api:"nullable"`
	// Create webhooks.
	CreateWebhook bool `json:"create_webhook" api:"nullable"`
	// Delete API keys.
	DeleteAPIKey bool `json:"delete_api_key" api:"nullable"`
	// Delete domains.
	DeleteDomain bool `json:"delete_domain" api:"nullable"`
	// Delete drafts.
	DeleteDraft bool `json:"delete_draft" api:"nullable"`
	// Delete inboxes.
	DeleteInbox bool `json:"delete_inbox" api:"nullable"`
	// Delete list entries.
	DeleteListEntry bool `json:"delete_list_entry" api:"nullable"`
	// Delete pods.
	DeletePod bool `json:"delete_pod" api:"nullable"`
	// Delete threads.
	DeleteThread bool `json:"delete_thread" api:"nullable"`
	// Delete webhooks.
	DeleteWebhook bool `json:"delete_webhook" api:"nullable"`
	// Read API keys.
	ReadAPIKey bool `json:"read_api_key" api:"nullable"`
	// Access messages labeled blocked.
	ReadBlocked bool `json:"read_blocked" api:"nullable"`
	// Read domain details.
	ReadDomain bool `json:"read_domain" api:"nullable"`
	// Read drafts.
	ReadDraft bool `json:"read_draft" api:"nullable"`
	// Read inbox details.
	ReadInbox bool `json:"read_inbox" api:"nullable"`
	// Read list entries.
	ReadListEntry bool `json:"read_list_entry" api:"nullable"`
	// Read messages.
	ReadMessage bool `json:"read_message" api:"nullable"`
	// Read metrics.
	ReadMetrics bool `json:"read_metrics" api:"nullable"`
	// Read pods.
	ReadPod bool `json:"read_pod" api:"nullable"`
	// Access messages labeled spam.
	ReadSpam bool `json:"read_spam" api:"nullable"`
	// Read threads.
	ReadThread bool `json:"read_thread" api:"nullable"`
	// Access messages labeled trash.
	ReadTrash bool `json:"read_trash" api:"nullable"`
	// Read webhook configurations.
	ReadWebhook bool `json:"read_webhook" api:"nullable"`
	// Send drafts.
	SendDraft bool `json:"send_draft" api:"nullable"`
	// Send messages.
	SendMessage bool `json:"send_message" api:"nullable"`
	// Update domains.
	UpdateDomain bool `json:"update_domain" api:"nullable"`
	// Update drafts.
	UpdateDraft bool `json:"update_draft" api:"nullable"`
	// Update inbox settings.
	UpdateInbox bool `json:"update_inbox" api:"nullable"`
	// Update message labels.
	UpdateMessage bool `json:"update_message" api:"nullable"`
	// Update webhooks.
	UpdateWebhook bool `json:"update_webhook" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateAPIKey    respjson.Field
		CreateDomain    respjson.Field
		CreateDraft     respjson.Field
		CreateInbox     respjson.Field
		CreateListEntry respjson.Field
		CreatePod       respjson.Field
		CreateWebhook   respjson.Field
		DeleteAPIKey    respjson.Field
		DeleteDomain    respjson.Field
		DeleteDraft     respjson.Field
		DeleteInbox     respjson.Field
		DeleteListEntry respjson.Field
		DeletePod       respjson.Field
		DeleteThread    respjson.Field
		DeleteWebhook   respjson.Field
		ReadAPIKey      respjson.Field
		ReadBlocked     respjson.Field
		ReadDomain      respjson.Field
		ReadDraft       respjson.Field
		ReadInbox       respjson.Field
		ReadListEntry   respjson.Field
		ReadMessage     respjson.Field
		ReadMetrics     respjson.Field
		ReadPod         respjson.Field
		ReadSpam        respjson.Field
		ReadThread      respjson.Field
		ReadTrash       respjson.Field
		ReadWebhook     respjson.Field
		SendDraft       respjson.Field
		SendMessage     respjson.Field
		UpdateDomain    respjson.Field
		UpdateDraft     respjson.Field
		UpdateInbox     respjson.Field
		UpdateMessage   respjson.Field
		UpdateWebhook   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyNewResponsePermissions) RawJSON() string { return r.JSON.raw }
func (r *APIKeyNewResponsePermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyListResponse struct {
	// Ordered by `created_at` descending.
	APIKeys []APIKeyListResponseAPIKey `json:"api_keys" api:"required"`
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
func (r APIKeyListResponse) RawJSON() string { return r.JSON.raw }
func (r *APIKeyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyListResponseAPIKey struct {
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
	Permissions APIKeyListResponseAPIKeyPermissions `json:"permissions" api:"nullable"`
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
func (r APIKeyListResponseAPIKey) RawJSON() string { return r.JSON.raw }
func (r *APIKeyListResponseAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Granular permissions for the API key. When ommitted all permissions are granted.
// Otherwise, only permissions set to true are granted.
type APIKeyListResponseAPIKeyPermissions struct {
	// Create API keys.
	CreateAPIKey bool `json:"create_api_key" api:"nullable"`
	// Create domains.
	CreateDomain bool `json:"create_domain" api:"nullable"`
	// Create drafts.
	CreateDraft bool `json:"create_draft" api:"nullable"`
	// Create new inboxes.
	CreateInbox bool `json:"create_inbox" api:"nullable"`
	// Create list entries.
	CreateListEntry bool `json:"create_list_entry" api:"nullable"`
	// Create pods.
	CreatePod bool `json:"create_pod" api:"nullable"`
	// Create webhooks.
	CreateWebhook bool `json:"create_webhook" api:"nullable"`
	// Delete API keys.
	DeleteAPIKey bool `json:"delete_api_key" api:"nullable"`
	// Delete domains.
	DeleteDomain bool `json:"delete_domain" api:"nullable"`
	// Delete drafts.
	DeleteDraft bool `json:"delete_draft" api:"nullable"`
	// Delete inboxes.
	DeleteInbox bool `json:"delete_inbox" api:"nullable"`
	// Delete list entries.
	DeleteListEntry bool `json:"delete_list_entry" api:"nullable"`
	// Delete pods.
	DeletePod bool `json:"delete_pod" api:"nullable"`
	// Delete threads.
	DeleteThread bool `json:"delete_thread" api:"nullable"`
	// Delete webhooks.
	DeleteWebhook bool `json:"delete_webhook" api:"nullable"`
	// Read API keys.
	ReadAPIKey bool `json:"read_api_key" api:"nullable"`
	// Access messages labeled blocked.
	ReadBlocked bool `json:"read_blocked" api:"nullable"`
	// Read domain details.
	ReadDomain bool `json:"read_domain" api:"nullable"`
	// Read drafts.
	ReadDraft bool `json:"read_draft" api:"nullable"`
	// Read inbox details.
	ReadInbox bool `json:"read_inbox" api:"nullable"`
	// Read list entries.
	ReadListEntry bool `json:"read_list_entry" api:"nullable"`
	// Read messages.
	ReadMessage bool `json:"read_message" api:"nullable"`
	// Read metrics.
	ReadMetrics bool `json:"read_metrics" api:"nullable"`
	// Read pods.
	ReadPod bool `json:"read_pod" api:"nullable"`
	// Access messages labeled spam.
	ReadSpam bool `json:"read_spam" api:"nullable"`
	// Read threads.
	ReadThread bool `json:"read_thread" api:"nullable"`
	// Access messages labeled trash.
	ReadTrash bool `json:"read_trash" api:"nullable"`
	// Read webhook configurations.
	ReadWebhook bool `json:"read_webhook" api:"nullable"`
	// Send drafts.
	SendDraft bool `json:"send_draft" api:"nullable"`
	// Send messages.
	SendMessage bool `json:"send_message" api:"nullable"`
	// Update domains.
	UpdateDomain bool `json:"update_domain" api:"nullable"`
	// Update drafts.
	UpdateDraft bool `json:"update_draft" api:"nullable"`
	// Update inbox settings.
	UpdateInbox bool `json:"update_inbox" api:"nullable"`
	// Update message labels.
	UpdateMessage bool `json:"update_message" api:"nullable"`
	// Update webhooks.
	UpdateWebhook bool `json:"update_webhook" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateAPIKey    respjson.Field
		CreateDomain    respjson.Field
		CreateDraft     respjson.Field
		CreateInbox     respjson.Field
		CreateListEntry respjson.Field
		CreatePod       respjson.Field
		CreateWebhook   respjson.Field
		DeleteAPIKey    respjson.Field
		DeleteDomain    respjson.Field
		DeleteDraft     respjson.Field
		DeleteInbox     respjson.Field
		DeleteListEntry respjson.Field
		DeletePod       respjson.Field
		DeleteThread    respjson.Field
		DeleteWebhook   respjson.Field
		ReadAPIKey      respjson.Field
		ReadBlocked     respjson.Field
		ReadDomain      respjson.Field
		ReadDraft       respjson.Field
		ReadInbox       respjson.Field
		ReadListEntry   respjson.Field
		ReadMessage     respjson.Field
		ReadMetrics     respjson.Field
		ReadPod         respjson.Field
		ReadSpam        respjson.Field
		ReadThread      respjson.Field
		ReadTrash       respjson.Field
		ReadWebhook     respjson.Field
		SendDraft       respjson.Field
		SendMessage     respjson.Field
		UpdateDomain    respjson.Field
		UpdateDraft     respjson.Field
		UpdateInbox     respjson.Field
		UpdateMessage   respjson.Field
		UpdateWebhook   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyListResponseAPIKeyPermissions) RawJSON() string { return r.JSON.raw }
func (r *APIKeyListResponseAPIKeyPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyNewParams struct {
	// Name of api key.
	Name string `json:"name" api:"required"`
	// Granular permissions for the API key. When ommitted all permissions are granted.
	// Otherwise, only permissions set to true are granted.
	Permissions APIKeyNewParamsPermissions `json:"permissions,omitzero"`
	paramObj
}

func (r APIKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Granular permissions for the API key. When ommitted all permissions are granted.
// Otherwise, only permissions set to true are granted.
type APIKeyNewParamsPermissions struct {
	// Create API keys.
	CreateAPIKey param.Opt[bool] `json:"create_api_key,omitzero"`
	// Create domains.
	CreateDomain param.Opt[bool] `json:"create_domain,omitzero"`
	// Create drafts.
	CreateDraft param.Opt[bool] `json:"create_draft,omitzero"`
	// Create new inboxes.
	CreateInbox param.Opt[bool] `json:"create_inbox,omitzero"`
	// Create list entries.
	CreateListEntry param.Opt[bool] `json:"create_list_entry,omitzero"`
	// Create pods.
	CreatePod param.Opt[bool] `json:"create_pod,omitzero"`
	// Create webhooks.
	CreateWebhook param.Opt[bool] `json:"create_webhook,omitzero"`
	// Delete API keys.
	DeleteAPIKey param.Opt[bool] `json:"delete_api_key,omitzero"`
	// Delete domains.
	DeleteDomain param.Opt[bool] `json:"delete_domain,omitzero"`
	// Delete drafts.
	DeleteDraft param.Opt[bool] `json:"delete_draft,omitzero"`
	// Delete inboxes.
	DeleteInbox param.Opt[bool] `json:"delete_inbox,omitzero"`
	// Delete list entries.
	DeleteListEntry param.Opt[bool] `json:"delete_list_entry,omitzero"`
	// Delete pods.
	DeletePod param.Opt[bool] `json:"delete_pod,omitzero"`
	// Delete threads.
	DeleteThread param.Opt[bool] `json:"delete_thread,omitzero"`
	// Delete webhooks.
	DeleteWebhook param.Opt[bool] `json:"delete_webhook,omitzero"`
	// Read API keys.
	ReadAPIKey param.Opt[bool] `json:"read_api_key,omitzero"`
	// Access messages labeled blocked.
	ReadBlocked param.Opt[bool] `json:"read_blocked,omitzero"`
	// Read domain details.
	ReadDomain param.Opt[bool] `json:"read_domain,omitzero"`
	// Read drafts.
	ReadDraft param.Opt[bool] `json:"read_draft,omitzero"`
	// Read inbox details.
	ReadInbox param.Opt[bool] `json:"read_inbox,omitzero"`
	// Read list entries.
	ReadListEntry param.Opt[bool] `json:"read_list_entry,omitzero"`
	// Read messages.
	ReadMessage param.Opt[bool] `json:"read_message,omitzero"`
	// Read metrics.
	ReadMetrics param.Opt[bool] `json:"read_metrics,omitzero"`
	// Read pods.
	ReadPod param.Opt[bool] `json:"read_pod,omitzero"`
	// Access messages labeled spam.
	ReadSpam param.Opt[bool] `json:"read_spam,omitzero"`
	// Read threads.
	ReadThread param.Opt[bool] `json:"read_thread,omitzero"`
	// Access messages labeled trash.
	ReadTrash param.Opt[bool] `json:"read_trash,omitzero"`
	// Read webhook configurations.
	ReadWebhook param.Opt[bool] `json:"read_webhook,omitzero"`
	// Send drafts.
	SendDraft param.Opt[bool] `json:"send_draft,omitzero"`
	// Send messages.
	SendMessage param.Opt[bool] `json:"send_message,omitzero"`
	// Update domains.
	UpdateDomain param.Opt[bool] `json:"update_domain,omitzero"`
	// Update drafts.
	UpdateDraft param.Opt[bool] `json:"update_draft,omitzero"`
	// Update inbox settings.
	UpdateInbox param.Opt[bool] `json:"update_inbox,omitzero"`
	// Update message labels.
	UpdateMessage param.Opt[bool] `json:"update_message,omitzero"`
	// Update webhooks.
	UpdateWebhook param.Opt[bool] `json:"update_webhook,omitzero"`
	paramObj
}

func (r APIKeyNewParamsPermissions) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyNewParamsPermissions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeyNewParamsPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyListParams struct {
	// Sort in ascending temporal order.
	Ascending param.Opt[bool] `query:"ascending,omitzero" json:"-"`
	// Limit of number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page token for pagination.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIKeyListParams]'s query parameters as `url.Values`.
func (r APIKeyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
