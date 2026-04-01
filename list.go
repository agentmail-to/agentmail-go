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

// ListService contains methods and other services that help with interacting with
// the agentmail API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewListService] method instead.
type ListService struct {
	Options []option.RequestOption
}

// NewListService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewListService(opts ...option.RequestOption) (r ListService) {
	r = ListService{}
	r.Options = opts
	return
}

// **CLI:**
//
// ```bash
// agentmail lists create --direction <direction> --type <type> --entry user@example.com
// ```
func (r *ListService) New(ctx context.Context, type_ ListNewParamsType, params ListNewParams, opts ...option.RequestOption) (res *ListNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	path := fmt.Sprintf("v0/lists/%v/%v", params.Direction, type_)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail lists retrieve --direction <direction> --type <type> --entry <entry>
// ```
func (r *ListService) Get(ctx context.Context, entry string, query ListGetParams, opts ...option.RequestOption) (res *ListGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if entry == "" {
		err = errors.New("missing required entry parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/lists/%v/%v/%s", query.Direction, query.Type, url.PathEscape(entry))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail lists list --direction <direction> --type <type>
// ```
func (r *ListService) List(ctx context.Context, type_ ListListParamsType, params ListListParams, opts ...option.RequestOption) (res *ListListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	path := fmt.Sprintf("v0/lists/%v/%v", params.Direction, type_)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// **CLI:**
//
// ```bash
// agentmail lists delete --direction <direction> --type <type> --entry <entry>
// ```
func (r *ListService) Delete(ctx context.Context, entry string, body ListDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if entry == "" {
		err = errors.New("missing required entry parameter")
		return err
	}
	path := fmt.Sprintf("v0/lists/%v/%v/%s", body.Direction, body.Type, url.PathEscape(entry))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type ListNewResponse struct {
	// Time at which entry was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Direction of list entry.
	//
	// Any of "send", "receive", "reply".
	Direction ListNewResponseDirection `json:"direction" api:"required"`
	// Email address or domain of list entry.
	Entry string `json:"entry" api:"required"`
	// Whether the entry is an email address or domain.
	//
	// Any of "email", "domain".
	EntryType ListNewResponseEntryType `json:"entry_type" api:"required"`
	// Type of list entry.
	//
	// Any of "allow", "block".
	ListType ListNewResponseListType `json:"list_type" api:"required"`
	// ID of organization.
	OrganizationID string `json:"organization_id" api:"required"`
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
		Reason         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ListNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Direction of list entry.
type ListNewResponseDirection string

const (
	ListNewResponseDirectionSend    ListNewResponseDirection = "send"
	ListNewResponseDirectionReceive ListNewResponseDirection = "receive"
	ListNewResponseDirectionReply   ListNewResponseDirection = "reply"
)

// Whether the entry is an email address or domain.
type ListNewResponseEntryType string

const (
	ListNewResponseEntryTypeEmail  ListNewResponseEntryType = "email"
	ListNewResponseEntryTypeDomain ListNewResponseEntryType = "domain"
)

// Type of list entry.
type ListNewResponseListType string

const (
	ListNewResponseListTypeAllow ListNewResponseListType = "allow"
	ListNewResponseListTypeBlock ListNewResponseListType = "block"
)

type ListGetResponse struct {
	// Time at which entry was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Direction of list entry.
	//
	// Any of "send", "receive", "reply".
	Direction ListGetResponseDirection `json:"direction" api:"required"`
	// Email address or domain of list entry.
	Entry string `json:"entry" api:"required"`
	// Whether the entry is an email address or domain.
	//
	// Any of "email", "domain".
	EntryType ListGetResponseEntryType `json:"entry_type" api:"required"`
	// Type of list entry.
	//
	// Any of "allow", "block".
	ListType ListGetResponseListType `json:"list_type" api:"required"`
	// ID of organization.
	OrganizationID string `json:"organization_id" api:"required"`
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
		Reason         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ListGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Direction of list entry.
type ListGetResponseDirection string

const (
	ListGetResponseDirectionSend    ListGetResponseDirection = "send"
	ListGetResponseDirectionReceive ListGetResponseDirection = "receive"
	ListGetResponseDirectionReply   ListGetResponseDirection = "reply"
)

// Whether the entry is an email address or domain.
type ListGetResponseEntryType string

const (
	ListGetResponseEntryTypeEmail  ListGetResponseEntryType = "email"
	ListGetResponseEntryTypeDomain ListGetResponseEntryType = "domain"
)

// Type of list entry.
type ListGetResponseListType string

const (
	ListGetResponseListTypeAllow ListGetResponseListType = "allow"
	ListGetResponseListTypeBlock ListGetResponseListType = "block"
)

type ListListResponse struct {
	// Number of items returned.
	Count int64 `json:"count" api:"required"`
	// Ordered by entry ascending.
	Entries []ListListResponseEntry `json:"entries" api:"required"`
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
func (r ListListResponse) RawJSON() string { return r.JSON.raw }
func (r *ListListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListListResponseEntry struct {
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
		Reason         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListListResponseEntry) RawJSON() string { return r.JSON.raw }
func (r *ListListResponseEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListNewParams struct {
	// Direction of list entry.
	//
	// Any of "send", "receive", "reply".
	Direction ListNewParamsDirection `path:"direction,omitzero" api:"required" json:"-"`
	// Email address or domain to add.
	Entry string `json:"entry" api:"required"`
	// Reason for adding the entry.
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r ListNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ListNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Direction of list entry.
type ListNewParamsDirection string

const (
	ListNewParamsDirectionSend    ListNewParamsDirection = "send"
	ListNewParamsDirectionReceive ListNewParamsDirection = "receive"
	ListNewParamsDirectionReply   ListNewParamsDirection = "reply"
)

// Type of list entry.
type ListNewParamsType string

const (
	ListNewParamsTypeAllow ListNewParamsType = "allow"
	ListNewParamsTypeBlock ListNewParamsType = "block"
)

type ListGetParams struct {
	// Direction of list entry.
	//
	// Any of "send", "receive", "reply".
	Direction ListGetParamsDirection `path:"direction,omitzero" api:"required" json:"-"`
	// Type of list entry.
	//
	// Any of "allow", "block".
	Type ListGetParamsType `path:"type,omitzero" api:"required" json:"-"`
	paramObj
}

// Direction of list entry.
type ListGetParamsDirection string

const (
	ListGetParamsDirectionSend    ListGetParamsDirection = "send"
	ListGetParamsDirectionReceive ListGetParamsDirection = "receive"
	ListGetParamsDirectionReply   ListGetParamsDirection = "reply"
)

// Type of list entry.
type ListGetParamsType string

const (
	ListGetParamsTypeAllow ListGetParamsType = "allow"
	ListGetParamsTypeBlock ListGetParamsType = "block"
)

type ListListParams struct {
	// Direction of list entry.
	//
	// Any of "send", "receive", "reply".
	Direction ListListParamsDirection `path:"direction,omitzero" api:"required" json:"-"`
	// Limit of number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page token for pagination.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListListParams]'s query parameters as `url.Values`.
func (r ListListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction of list entry.
type ListListParamsDirection string

const (
	ListListParamsDirectionSend    ListListParamsDirection = "send"
	ListListParamsDirectionReceive ListListParamsDirection = "receive"
	ListListParamsDirectionReply   ListListParamsDirection = "reply"
)

// Type of list entry.
type ListListParamsType string

const (
	ListListParamsTypeAllow ListListParamsType = "allow"
	ListListParamsTypeBlock ListListParamsType = "block"
)

type ListDeleteParams struct {
	// Direction of list entry.
	//
	// Any of "send", "receive", "reply".
	Direction ListDeleteParamsDirection `path:"direction,omitzero" api:"required" json:"-"`
	// Type of list entry.
	//
	// Any of "allow", "block".
	Type ListDeleteParamsType `path:"type,omitzero" api:"required" json:"-"`
	paramObj
}

// Direction of list entry.
type ListDeleteParamsDirection string

const (
	ListDeleteParamsDirectionSend    ListDeleteParamsDirection = "send"
	ListDeleteParamsDirectionReceive ListDeleteParamsDirection = "receive"
	ListDeleteParamsDirectionReply   ListDeleteParamsDirection = "reply"
)

// Type of list entry.
type ListDeleteParamsType string

const (
	ListDeleteParamsTypeAllow ListDeleteParamsType = "allow"
	ListDeleteParamsTypeBlock ListDeleteParamsType = "block"
)
