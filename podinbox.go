// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package agentmail

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/agentmail-to/agentmail-go/internal/apijson"
	"github.com/agentmail-to/agentmail-go/internal/apiquery"
	shimjson "github.com/agentmail-to/agentmail-go/internal/encoding/json"
	"github.com/agentmail-to/agentmail-go/internal/requestconfig"
	"github.com/agentmail-to/agentmail-go/option"
	"github.com/agentmail-to/agentmail-go/packages/param"
)

// PodInboxService contains methods and other services that help with interacting
// with the agentmail API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPodInboxService] method instead.
type PodInboxService struct {
	Options []option.RequestOption
}

// NewPodInboxService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPodInboxService(opts ...option.RequestOption) (r PodInboxService) {
	r = PodInboxService{}
	r.Options = opts
	return
}

// **CLI:**
//
// ```bash
// agentmail pods:inboxes create --pod-id <pod_id> --username myagent --domain example.com
// ```
func (r *PodInboxService) New(ctx context.Context, podID string, body PodInboxNewParams, opts ...option.RequestOption) (res *Inbox, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if podID == "" {
		err = errors.New("missing required pod_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/pods/%s/inboxes", url.PathEscape(podID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail pods:inboxes retrieve --pod-id <pod_id> --inbox-id <inbox_id>
// ```
func (r *PodInboxService) Get(ctx context.Context, inboxID string, query PodInboxGetParams, opts ...option.RequestOption) (res *Inbox, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if query.PodID == "" {
		err = errors.New("missing required pod_id parameter")
		return nil, err
	}
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/pods/%s/inboxes/%s", url.PathEscape(query.PodID), url.PathEscape(inboxID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail pods:inboxes update --pod-id <pod_id> --inbox-id <inbox_id>
// ```
func (r *PodInboxService) Update(ctx context.Context, inboxID string, params PodInboxUpdateParams, opts ...option.RequestOption) (res *Inbox, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if params.PodID == "" {
		err = errors.New("missing required pod_id parameter")
		return nil, err
	}
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/pods/%s/inboxes/%s", url.PathEscape(params.PodID), url.PathEscape(inboxID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail pods:inboxes list --pod-id <pod_id>
// ```
func (r *PodInboxService) List(ctx context.Context, podID string, query PodInboxListParams, opts ...option.RequestOption) (res *ListInboxes, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if podID == "" {
		err = errors.New("missing required pod_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/pods/%s/inboxes", url.PathEscape(podID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail pods:inboxes delete --pod-id <pod_id> --inbox-id <inbox_id>
// ```
func (r *PodInboxService) Delete(ctx context.Context, inboxID string, body PodInboxDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if body.PodID == "" {
		err = errors.New("missing required pod_id parameter")
		return err
	}
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return err
	}
	path := fmt.Sprintf("v0/pods/%s/inboxes/%s", url.PathEscape(body.PodID), url.PathEscape(inboxID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type PodInboxNewParams struct {
	CreateInbox CreateInboxParam
	paramObj
}

func (r PodInboxNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateInbox)
}
func (r *PodInboxNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PodInboxGetParams struct {
	// ID of pod.
	PodID string `path:"pod_id" api:"required" json:"-"`
	paramObj
}

type PodInboxUpdateParams struct {
	// ID of pod.
	PodID string `path:"pod_id" api:"required" json:"-"`
	// Display name: `Display Name <username@domain.com>`.
	DisplayName string `json:"display_name" api:"required"`
	paramObj
}

func (r PodInboxUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PodInboxUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PodInboxUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PodInboxListParams struct {
	// Sort in ascending temporal order.
	Ascending param.Opt[bool] `query:"ascending,omitzero" json:"-"`
	// Limit of number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page token for pagination.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PodInboxListParams]'s query parameters as `url.Values`.
func (r PodInboxListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PodInboxDeleteParams struct {
	// ID of pod.
	PodID string `path:"pod_id" api:"required" json:"-"`
	paramObj
}
