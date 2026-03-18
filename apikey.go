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

// Delete API Key
func (r *APIKeyService) Delete(ctx context.Context, apiKey string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if apiKey == "" {
		err = errors.New("missing required api_key parameter")
		return err
	}
	path := fmt.Sprintf("v0/api-keys/%s", url.PathEscape(apiKey))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
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

type APIKeyNewParams struct {
	// Name of api key.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r APIKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeyNewParams) UnmarshalJSON(data []byte) error {
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
