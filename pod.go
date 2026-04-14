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

// PodService contains methods and other services that help with interacting with
// the agentmail API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPodService] method instead.
type PodService struct {
	Options []option.RequestOption
	Domains PodDomainService
	Drafts  PodDraftService
	Inboxes PodInboxService
	Threads PodThreadService
	Lists   PodListService
	APIKeys PodAPIKeyService
	Metrics PodMetricService
}

// NewPodService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPodService(opts ...option.RequestOption) (r PodService) {
	r = PodService{}
	r.Options = opts
	r.Domains = NewPodDomainService(opts...)
	r.Drafts = NewPodDraftService(opts...)
	r.Inboxes = NewPodInboxService(opts...)
	r.Threads = NewPodThreadService(opts...)
	r.Lists = NewPodListService(opts...)
	r.APIKeys = NewPodAPIKeyService(opts...)
	r.Metrics = NewPodMetricService(opts...)
	return
}

// **CLI:**
//
// ```bash
// agentmail pods create --client-id my-pod
// ```
func (r *PodService) New(ctx context.Context, body PodNewParams, opts ...option.RequestOption) (res *Pod, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	path := "v0/pods"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail pods list
// ```
func (r *PodService) List(ctx context.Context, query PodListParams, opts ...option.RequestOption) (res *PodListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	path := "v0/pods"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail pods delete --pod-id <pod_id>
// ```
func (r *PodService) Delete(ctx context.Context, podID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if podID == "" {
		err = errors.New("missing required pod_id parameter")
		return err
	}
	path := fmt.Sprintf("v0/pods/%s", url.PathEscape(podID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// **CLI:**
//
// ```bash
// agentmail pods get --pod-id <pod_id>
// ```
func (r *PodService) Get(ctx context.Context, podID string, opts ...option.RequestOption) (res *Pod, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if podID == "" {
		err = errors.New("missing required pod_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/pods/%s", url.PathEscape(podID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Pod struct {
	// Time at which pod was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Name of pod.
	Name string `json:"name" api:"required"`
	// ID of pod.
	PodID string `json:"pod_id" api:"required"`
	// Time at which pod was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Client ID of pod.
	ClientID string `json:"client_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Name        respjson.Field
		PodID       respjson.Field
		UpdatedAt   respjson.Field
		ClientID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Pod) RawJSON() string { return r.JSON.raw }
func (r *Pod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PodListResponse struct {
	// Number of items returned.
	Count int64 `json:"count" api:"required"`
	// Ordered by `created_at` descending.
	Pods []Pod `json:"pods" api:"required"`
	// Limit of number of items returned.
	Limit int64 `json:"limit" api:"nullable"`
	// Page token for pagination.
	NextPageToken string `json:"next_page_token" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count         respjson.Field
		Pods          respjson.Field
		Limit         respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PodListResponse) RawJSON() string { return r.JSON.raw }
func (r *PodListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PodNewParams struct {
	// Client ID of pod.
	ClientID param.Opt[string] `json:"client_id,omitzero"`
	// Name of pod.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r PodNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PodNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PodNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PodListParams struct {
	// Sort in ascending temporal order.
	Ascending param.Opt[bool] `query:"ascending,omitzero" json:"-"`
	// Limit of number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page token for pagination.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PodListParams]'s query parameters as `url.Values`.
func (r PodListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
