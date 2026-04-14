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

// PodListService contains methods and other services that help with interacting
// with the agentmail API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPodListService] method instead.
type PodListService struct {
	Options []option.RequestOption
}

// NewPodListService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPodListService(opts ...option.RequestOption) (r PodListService) {
	r = PodListService{}
	r.Options = opts
	return
}

// **CLI:**
//
// ```bash
// agentmail pods:lists create --pod-id <pod_id> --direction <direction> --type <type> --entry user@example.com
// ```
func (r *PodListService) New(ctx context.Context, type_ PodListNewParamsType, params PodListNewParams, opts ...option.RequestOption) (res *PodListNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if params.PodID == "" {
		err = errors.New("missing required pod_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/pods/%s/lists/%v/%v", url.PathEscape(params.PodID), params.Direction, type_)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail pods:lists list --pod-id <pod_id> --direction <direction> --type <type>
// ```
func (r *PodListService) List(ctx context.Context, type_ PodListListParamsType, params PodListListParams, opts ...option.RequestOption) (res *PodListListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if params.PodID == "" {
		err = errors.New("missing required pod_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/pods/%s/lists/%v/%v", url.PathEscape(params.PodID), params.Direction, type_)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail pods:lists delete --pod-id <pod_id> --direction <direction> --type <type> --entry <entry>
// ```
func (r *PodListService) Delete(ctx context.Context, entry string, body PodListDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if body.PodID == "" {
		err = errors.New("missing required pod_id parameter")
		return err
	}
	if entry == "" {
		err = errors.New("missing required entry parameter")
		return err
	}
	path := fmt.Sprintf("v0/pods/%s/lists/%v/%v/%s", url.PathEscape(body.PodID), body.Direction, body.Type, url.PathEscape(entry))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// **CLI:**
//
// ```bash
// agentmail pods:lists retrieve --pod-id <pod_id> --direction <direction> --type <type> --entry <entry>
// ```
func (r *PodListService) Get(ctx context.Context, entry string, query PodListGetParams, opts ...option.RequestOption) (res *PodListGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if query.PodID == "" {
		err = errors.New("missing required pod_id parameter")
		return nil, err
	}
	if entry == "" {
		err = errors.New("missing required entry parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/pods/%s/lists/%v/%v/%s", url.PathEscape(query.PodID), query.Direction, query.Type, url.PathEscape(entry))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type PodListNewResponse struct {
	// Time at which entry was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Direction of list entry.
	//
	// Any of "send", "receive", "reply".
	Direction PodListNewResponseDirection `json:"direction" api:"required"`
	// Email address or domain of list entry.
	Entry string `json:"entry" api:"required"`
	// Whether the entry is an email address or domain.
	//
	// Any of "email", "domain".
	EntryType PodListNewResponseEntryType `json:"entry_type" api:"required"`
	// Type of list entry.
	//
	// Any of "allow", "block".
	ListType PodListNewResponseListType `json:"list_type" api:"required"`
	// ID of organization.
	OrganizationID string `json:"organization_id" api:"required"`
	// ID of pod.
	PodID string `json:"pod_id" api:"required"`
	// ID of inbox, if entry is inbox-scoped.
	InboxID string `json:"inbox_id" api:"nullable"`
	// Reason for adding the entry.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt      respjson.Field
		Direction      respjson.Field
		Entry          respjson.Field
		EntryType      respjson.Field
		ListType       respjson.Field
		OrganizationID respjson.Field
		PodID          respjson.Field
		InboxID        respjson.Field
		Reason         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PodListNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PodListNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Direction of list entry.
type PodListNewResponseDirection string

const (
	PodListNewResponseDirectionSend    PodListNewResponseDirection = "send"
	PodListNewResponseDirectionReceive PodListNewResponseDirection = "receive"
	PodListNewResponseDirectionReply   PodListNewResponseDirection = "reply"
)

// Whether the entry is an email address or domain.
type PodListNewResponseEntryType string

const (
	PodListNewResponseEntryTypeEmail  PodListNewResponseEntryType = "email"
	PodListNewResponseEntryTypeDomain PodListNewResponseEntryType = "domain"
)

// Type of list entry.
type PodListNewResponseListType string

const (
	PodListNewResponseListTypeAllow PodListNewResponseListType = "allow"
	PodListNewResponseListTypeBlock PodListNewResponseListType = "block"
)

type PodListListResponse struct {
	// Number of items returned.
	Count int64 `json:"count" api:"required"`
	// Ordered by entry ascending.
	Entries []PodListListResponseEntry `json:"entries" api:"required"`
	// Limit of number of items returned.
	Limit int64 `json:"limit" api:"nullable"`
	// Page token for pagination.
	NextPageToken string `json:"next_page_token" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count         respjson.Field
		Entries       respjson.Field
		Limit         respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PodListListResponse) RawJSON() string { return r.JSON.raw }
func (r *PodListListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PodListListResponseEntry struct {
	// Time at which entry was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Direction of list entry.
	//
	// Any of "send", "receive", "reply".
	Direction string `json:"direction" api:"required"`
	// Email address or domain of list entry.
	Entry string `json:"entry" api:"required"`
	// Whether the entry is an email address or domain.
	//
	// Any of "email", "domain".
	EntryType string `json:"entry_type" api:"required"`
	// Type of list entry.
	//
	// Any of "allow", "block".
	ListType string `json:"list_type" api:"required"`
	// ID of organization.
	OrganizationID string `json:"organization_id" api:"required"`
	// ID of pod.
	PodID string `json:"pod_id" api:"required"`
	// ID of inbox, if entry is inbox-scoped.
	InboxID string `json:"inbox_id" api:"nullable"`
	// Reason for adding the entry.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt      respjson.Field
		Direction      respjson.Field
		Entry          respjson.Field
		EntryType      respjson.Field
		ListType       respjson.Field
		OrganizationID respjson.Field
		PodID          respjson.Field
		InboxID        respjson.Field
		Reason         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PodListListResponseEntry) RawJSON() string { return r.JSON.raw }
func (r *PodListListResponseEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PodListGetResponse struct {
	// Time at which entry was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Direction of list entry.
	//
	// Any of "send", "receive", "reply".
	Direction PodListGetResponseDirection `json:"direction" api:"required"`
	// Email address or domain of list entry.
	Entry string `json:"entry" api:"required"`
	// Whether the entry is an email address or domain.
	//
	// Any of "email", "domain".
	EntryType PodListGetResponseEntryType `json:"entry_type" api:"required"`
	// Type of list entry.
	//
	// Any of "allow", "block".
	ListType PodListGetResponseListType `json:"list_type" api:"required"`
	// ID of organization.
	OrganizationID string `json:"organization_id" api:"required"`
	// ID of pod.
	PodID string `json:"pod_id" api:"required"`
	// ID of inbox, if entry is inbox-scoped.
	InboxID string `json:"inbox_id" api:"nullable"`
	// Reason for adding the entry.
	Reason string `json:"reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt      respjson.Field
		Direction      respjson.Field
		Entry          respjson.Field
		EntryType      respjson.Field
		ListType       respjson.Field
		OrganizationID respjson.Field
		PodID          respjson.Field
		InboxID        respjson.Field
		Reason         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PodListGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PodListGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Direction of list entry.
type PodListGetResponseDirection string

const (
	PodListGetResponseDirectionSend    PodListGetResponseDirection = "send"
	PodListGetResponseDirectionReceive PodListGetResponseDirection = "receive"
	PodListGetResponseDirectionReply   PodListGetResponseDirection = "reply"
)

// Whether the entry is an email address or domain.
type PodListGetResponseEntryType string

const (
	PodListGetResponseEntryTypeEmail  PodListGetResponseEntryType = "email"
	PodListGetResponseEntryTypeDomain PodListGetResponseEntryType = "domain"
)

// Type of list entry.
type PodListGetResponseListType string

const (
	PodListGetResponseListTypeAllow PodListGetResponseListType = "allow"
	PodListGetResponseListTypeBlock PodListGetResponseListType = "block"
)

type PodListNewParams struct {
	// ID of pod.
	PodID string `path:"pod_id" api:"required" json:"-"`
	// Direction of list entry.
	//
	// Any of "send", "receive", "reply".
	Direction PodListNewParamsDirection `path:"direction,omitzero" api:"required" json:"-"`
	// Email address or domain to add.
	Entry string `json:"entry" api:"required"`
	// Reason for adding the entry.
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r PodListNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PodListNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PodListNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Direction of list entry.
type PodListNewParamsDirection string

const (
	PodListNewParamsDirectionSend    PodListNewParamsDirection = "send"
	PodListNewParamsDirectionReceive PodListNewParamsDirection = "receive"
	PodListNewParamsDirectionReply   PodListNewParamsDirection = "reply"
)

// Type of list entry.
type PodListNewParamsType string

const (
	PodListNewParamsTypeAllow PodListNewParamsType = "allow"
	PodListNewParamsTypeBlock PodListNewParamsType = "block"
)

type PodListListParams struct {
	// ID of pod.
	PodID string `path:"pod_id" api:"required" json:"-"`
	// Direction of list entry.
	//
	// Any of "send", "receive", "reply".
	Direction PodListListParamsDirection `path:"direction,omitzero" api:"required" json:"-"`
	// Limit of number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page token for pagination.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PodListListParams]'s query parameters as `url.Values`.
func (r PodListListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction of list entry.
type PodListListParamsDirection string

const (
	PodListListParamsDirectionSend    PodListListParamsDirection = "send"
	PodListListParamsDirectionReceive PodListListParamsDirection = "receive"
	PodListListParamsDirectionReply   PodListListParamsDirection = "reply"
)

// Type of list entry.
type PodListListParamsType string

const (
	PodListListParamsTypeAllow PodListListParamsType = "allow"
	PodListListParamsTypeBlock PodListListParamsType = "block"
)

type PodListDeleteParams struct {
	// ID of pod.
	PodID string `path:"pod_id" api:"required" json:"-"`
	// Direction of list entry.
	//
	// Any of "send", "receive", "reply".
	Direction PodListDeleteParamsDirection `path:"direction,omitzero" api:"required" json:"-"`
	// Type of list entry.
	//
	// Any of "allow", "block".
	Type PodListDeleteParamsType `path:"type,omitzero" api:"required" json:"-"`
	paramObj
}

// Direction of list entry.
type PodListDeleteParamsDirection string

const (
	PodListDeleteParamsDirectionSend    PodListDeleteParamsDirection = "send"
	PodListDeleteParamsDirectionReceive PodListDeleteParamsDirection = "receive"
	PodListDeleteParamsDirectionReply   PodListDeleteParamsDirection = "reply"
)

// Type of list entry.
type PodListDeleteParamsType string

const (
	PodListDeleteParamsTypeAllow PodListDeleteParamsType = "allow"
	PodListDeleteParamsTypeBlock PodListDeleteParamsType = "block"
)

type PodListGetParams struct {
	// ID of pod.
	PodID string `path:"pod_id" api:"required" json:"-"`
	// Direction of list entry.
	//
	// Any of "send", "receive", "reply".
	Direction PodListGetParamsDirection `path:"direction,omitzero" api:"required" json:"-"`
	// Type of list entry.
	//
	// Any of "allow", "block".
	Type PodListGetParamsType `path:"type,omitzero" api:"required" json:"-"`
	paramObj
}

// Direction of list entry.
type PodListGetParamsDirection string

const (
	PodListGetParamsDirectionSend    PodListGetParamsDirection = "send"
	PodListGetParamsDirectionReceive PodListGetParamsDirection = "receive"
	PodListGetParamsDirectionReply   PodListGetParamsDirection = "reply"
)

// Type of list entry.
type PodListGetParamsType string

const (
	PodListGetParamsTypeAllow PodListGetParamsType = "allow"
	PodListGetParamsTypeBlock PodListGetParamsType = "block"
)
