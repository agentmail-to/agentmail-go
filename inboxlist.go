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

// InboxListService contains methods and other services that help with interacting
// with the agentmail API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInboxListService] method instead.
type InboxListService struct {
	Options []option.RequestOption
}

// NewInboxListService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInboxListService(opts ...option.RequestOption) (r InboxListService) {
	r = InboxListService{}
	r.Options = opts
	return
}

// **CLI:**
//
// ```bash
// agentmail inboxes:lists create --inbox-id <inbox_id> --direction <direction> --type <type> --entry user@example.com
// ```
func (r *InboxListService) New(ctx context.Context, type_ InboxListNewParamsType, params InboxListNewParams, opts ...option.RequestOption) (res *InboxListNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if params.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/inboxes/%s/lists/%v/%v", url.PathEscape(params.InboxID), params.Direction, type_)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail inboxes:lists retrieve --inbox-id <inbox_id> --direction <direction> --type <type> --entry <entry>
// ```
func (r *InboxListService) Get(ctx context.Context, entry string, query InboxListGetParams, opts ...option.RequestOption) (res *InboxListGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if query.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	if entry == "" {
		err = errors.New("missing required entry parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/inboxes/%s/lists/%v/%v/%s", url.PathEscape(query.InboxID), query.Direction, query.Type, url.PathEscape(entry))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail inboxes:lists list --inbox-id <inbox_id> --direction <direction> --type <type>
// ```
func (r *InboxListService) List(ctx context.Context, type_ InboxListListParamsType, params InboxListListParams, opts ...option.RequestOption) (res *InboxListListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if params.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/inboxes/%s/lists/%v/%v", url.PathEscape(params.InboxID), params.Direction, type_)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail inboxes:lists delete --inbox-id <inbox_id> --direction <direction> --type <type> --entry <entry>
// ```
func (r *InboxListService) Delete(ctx context.Context, entry string, body InboxListDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if body.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return err
	}
	if entry == "" {
		err = errors.New("missing required entry parameter")
		return err
	}
	path := fmt.Sprintf("v0/inboxes/%s/lists/%v/%v/%s", url.PathEscape(body.InboxID), body.Direction, body.Type, url.PathEscape(entry))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type InboxListNewResponse struct {
	// Time at which entry was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Direction of list entry.
	//
	// Any of "send", "receive", "reply".
	Direction InboxListNewResponseDirection `json:"direction" api:"required"`
	// Email address or domain of list entry.
	Entry string `json:"entry" api:"required"`
	// Whether the entry is an email address or domain.
	//
	// Any of "email", "domain".
	EntryType InboxListNewResponseEntryType `json:"entry_type" api:"required"`
	// Type of list entry.
	//
	// Any of "allow", "block".
	ListType InboxListNewResponseListType `json:"list_type" api:"required"`
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
func (r InboxListNewResponse) RawJSON() string { return r.JSON.raw }
func (r *InboxListNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Direction of list entry.
type InboxListNewResponseDirection string

const (
	InboxListNewResponseDirectionSend    InboxListNewResponseDirection = "send"
	InboxListNewResponseDirectionReceive InboxListNewResponseDirection = "receive"
	InboxListNewResponseDirectionReply   InboxListNewResponseDirection = "reply"
)

// Whether the entry is an email address or domain.
type InboxListNewResponseEntryType string

const (
	InboxListNewResponseEntryTypeEmail  InboxListNewResponseEntryType = "email"
	InboxListNewResponseEntryTypeDomain InboxListNewResponseEntryType = "domain"
)

// Type of list entry.
type InboxListNewResponseListType string

const (
	InboxListNewResponseListTypeAllow InboxListNewResponseListType = "allow"
	InboxListNewResponseListTypeBlock InboxListNewResponseListType = "block"
)

type InboxListGetResponse struct {
	// Time at which entry was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Direction of list entry.
	//
	// Any of "send", "receive", "reply".
	Direction InboxListGetResponseDirection `json:"direction" api:"required"`
	// Email address or domain of list entry.
	Entry string `json:"entry" api:"required"`
	// Whether the entry is an email address or domain.
	//
	// Any of "email", "domain".
	EntryType InboxListGetResponseEntryType `json:"entry_type" api:"required"`
	// Type of list entry.
	//
	// Any of "allow", "block".
	ListType InboxListGetResponseListType `json:"list_type" api:"required"`
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
func (r InboxListGetResponse) RawJSON() string { return r.JSON.raw }
func (r *InboxListGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Direction of list entry.
type InboxListGetResponseDirection string

const (
	InboxListGetResponseDirectionSend    InboxListGetResponseDirection = "send"
	InboxListGetResponseDirectionReceive InboxListGetResponseDirection = "receive"
	InboxListGetResponseDirectionReply   InboxListGetResponseDirection = "reply"
)

// Whether the entry is an email address or domain.
type InboxListGetResponseEntryType string

const (
	InboxListGetResponseEntryTypeEmail  InboxListGetResponseEntryType = "email"
	InboxListGetResponseEntryTypeDomain InboxListGetResponseEntryType = "domain"
)

// Type of list entry.
type InboxListGetResponseListType string

const (
	InboxListGetResponseListTypeAllow InboxListGetResponseListType = "allow"
	InboxListGetResponseListTypeBlock InboxListGetResponseListType = "block"
)

type InboxListListResponse struct {
	// Number of items returned.
	Count int64 `json:"count" api:"required"`
	// Ordered by entry ascending.
	Entries []InboxListListResponseEntry `json:"entries" api:"required"`
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
func (r InboxListListResponse) RawJSON() string { return r.JSON.raw }
func (r *InboxListListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxListListResponseEntry struct {
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
func (r InboxListListResponseEntry) RawJSON() string { return r.JSON.raw }
func (r *InboxListListResponseEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxListNewParams struct {
	// The ID of the inbox.
	InboxID string `path:"inbox_id" api:"required" json:"-"`
	// Direction of list entry.
	//
	// Any of "send", "receive", "reply".
	Direction InboxListNewParamsDirection `path:"direction,omitzero" api:"required" json:"-"`
	// Email address or domain to add.
	Entry string `json:"entry" api:"required"`
	// Reason for adding the entry.
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r InboxListNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InboxListNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboxListNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Direction of list entry.
type InboxListNewParamsDirection string

const (
	InboxListNewParamsDirectionSend    InboxListNewParamsDirection = "send"
	InboxListNewParamsDirectionReceive InboxListNewParamsDirection = "receive"
	InboxListNewParamsDirectionReply   InboxListNewParamsDirection = "reply"
)

// Type of list entry.
type InboxListNewParamsType string

const (
	InboxListNewParamsTypeAllow InboxListNewParamsType = "allow"
	InboxListNewParamsTypeBlock InboxListNewParamsType = "block"
)

type InboxListGetParams struct {
	// The ID of the inbox.
	InboxID string `path:"inbox_id" api:"required" json:"-"`
	// Direction of list entry.
	//
	// Any of "send", "receive", "reply".
	Direction InboxListGetParamsDirection `path:"direction,omitzero" api:"required" json:"-"`
	// Type of list entry.
	//
	// Any of "allow", "block".
	Type InboxListGetParamsType `path:"type,omitzero" api:"required" json:"-"`
	paramObj
}

// Direction of list entry.
type InboxListGetParamsDirection string

const (
	InboxListGetParamsDirectionSend    InboxListGetParamsDirection = "send"
	InboxListGetParamsDirectionReceive InboxListGetParamsDirection = "receive"
	InboxListGetParamsDirectionReply   InboxListGetParamsDirection = "reply"
)

// Type of list entry.
type InboxListGetParamsType string

const (
	InboxListGetParamsTypeAllow InboxListGetParamsType = "allow"
	InboxListGetParamsTypeBlock InboxListGetParamsType = "block"
)

type InboxListListParams struct {
	// The ID of the inbox.
	InboxID string `path:"inbox_id" api:"required" json:"-"`
	// Direction of list entry.
	//
	// Any of "send", "receive", "reply".
	Direction InboxListListParamsDirection `path:"direction,omitzero" api:"required" json:"-"`
	// Limit of number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page token for pagination.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InboxListListParams]'s query parameters as `url.Values`.
func (r InboxListListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction of list entry.
type InboxListListParamsDirection string

const (
	InboxListListParamsDirectionSend    InboxListListParamsDirection = "send"
	InboxListListParamsDirectionReceive InboxListListParamsDirection = "receive"
	InboxListListParamsDirectionReply   InboxListListParamsDirection = "reply"
)

// Type of list entry.
type InboxListListParamsType string

const (
	InboxListListParamsTypeAllow InboxListListParamsType = "allow"
	InboxListListParamsTypeBlock InboxListListParamsType = "block"
)

type InboxListDeleteParams struct {
	// The ID of the inbox.
	InboxID string `path:"inbox_id" api:"required" json:"-"`
	// Direction of list entry.
	//
	// Any of "send", "receive", "reply".
	Direction InboxListDeleteParamsDirection `path:"direction,omitzero" api:"required" json:"-"`
	// Type of list entry.
	//
	// Any of "allow", "block".
	Type InboxListDeleteParamsType `path:"type,omitzero" api:"required" json:"-"`
	paramObj
}

// Direction of list entry.
type InboxListDeleteParamsDirection string

const (
	InboxListDeleteParamsDirectionSend    InboxListDeleteParamsDirection = "send"
	InboxListDeleteParamsDirectionReceive InboxListDeleteParamsDirection = "receive"
	InboxListDeleteParamsDirectionReply   InboxListDeleteParamsDirection = "reply"
)

// Type of list entry.
type InboxListDeleteParamsType string

const (
	InboxListDeleteParamsTypeAllow InboxListDeleteParamsType = "allow"
	InboxListDeleteParamsTypeBlock InboxListDeleteParamsType = "block"
)
