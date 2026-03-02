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

// Create Inbox
func (r *PodInboxService) New(ctx context.Context, podID string, body PodInboxNewParams, opts ...option.RequestOption) (res *Inbox, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if podID == "" {
		err = errors.New("missing required pod_id parameter")
		return
	}
	path := fmt.Sprintf("v0/pods/%s/inboxes", url.PathEscape(podID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Inbox
func (r *PodInboxService) Get(ctx context.Context, inboxID string, query PodInboxGetParams, opts ...option.RequestOption) (res *Inbox, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if query.PodID == "" {
		err = errors.New("missing required pod_id parameter")
		return
	}
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return
	}
	path := fmt.Sprintf("v0/pods/%s/inboxes/%s", url.PathEscape(query.PodID), url.PathEscape(inboxID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Inboxes
func (r *PodInboxService) List(ctx context.Context, podID string, query PodInboxListParams, opts ...option.RequestOption) (res *ListInboxes, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if podID == "" {
		err = errors.New("missing required pod_id parameter")
		return
	}
	path := fmt.Sprintf("v0/pods/%s/inboxes", url.PathEscape(podID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete Inbox
func (r *PodInboxService) Delete(ctx context.Context, inboxID string, body PodInboxDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if body.PodID == "" {
		err = errors.New("missing required pod_id parameter")
		return
	}
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return
	}
	path := fmt.Sprintf("v0/pods/%s/inboxes/%s", url.PathEscape(body.PodID), url.PathEscape(inboxID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type PodInboxNewParams struct {
	CreateInbox CreateInboxParam
	paramObj
}

func (r PodInboxNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateInbox)
}
func (r *PodInboxNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateInbox)
}

type PodInboxGetParams struct {
	// ID of pod.
	PodID string `path:"pod_id" api:"required" json:"-"`
	paramObj
}

type PodInboxListParams struct {
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
