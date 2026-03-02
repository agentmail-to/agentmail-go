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

// PodDomainService contains methods and other services that help with interacting
// with the agentmail API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPodDomainService] method instead.
type PodDomainService struct {
	Options []option.RequestOption
}

// NewPodDomainService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPodDomainService(opts ...option.RequestOption) (r PodDomainService) {
	r = PodDomainService{}
	r.Options = opts
	return
}

// Create Domain
func (r *PodDomainService) New(ctx context.Context, podID string, body PodDomainNewParams, opts ...option.RequestOption) (res *Domain, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if podID == "" {
		err = errors.New("missing required pod_id parameter")
		return
	}
	path := fmt.Sprintf("v0/pods/%s/domains", url.PathEscape(podID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List Domains
func (r *PodDomainService) List(ctx context.Context, podID string, query PodDomainListParams, opts ...option.RequestOption) (res *ListDomains, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if podID == "" {
		err = errors.New("missing required pod_id parameter")
		return
	}
	path := fmt.Sprintf("v0/pods/%s/domains", url.PathEscape(podID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete Domain
func (r *PodDomainService) Delete(ctx context.Context, domainID string, body PodDomainDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if body.PodID == "" {
		err = errors.New("missing required pod_id parameter")
		return
	}
	if domainID == "" {
		err = errors.New("missing required domain_id parameter")
		return
	}
	path := fmt.Sprintf("v0/pods/%s/domains/%s", url.PathEscape(body.PodID), url.PathEscape(domainID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type PodDomainNewParams struct {
	CreateDomain CreateDomainParam
	paramObj
}

func (r PodDomainNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateDomain)
}
func (r *PodDomainNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateDomain)
}

type PodDomainListParams struct {
	// Limit of number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page token for pagination.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PodDomainListParams]'s query parameters as `url.Values`.
func (r PodDomainListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PodDomainDeleteParams struct {
	// ID of pod.
	PodID string `path:"pod_id" api:"required" json:"-"`
	paramObj
}
