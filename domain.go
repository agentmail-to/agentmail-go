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
	"time"

	"github.com/stainless-sdks/agentmail-go/internal/apijson"
	"github.com/stainless-sdks/agentmail-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/agentmail-go/internal/encoding/json"
	"github.com/stainless-sdks/agentmail-go/internal/requestconfig"
	"github.com/stainless-sdks/agentmail-go/option"
	"github.com/stainless-sdks/agentmail-go/packages/param"
	"github.com/stainless-sdks/agentmail-go/packages/respjson"
)

// DomainService contains methods and other services that help with interacting
// with the agentmail API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainService] method instead.
type DomainService struct {
	Options []option.RequestOption
}

// NewDomainService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDomainService(opts ...option.RequestOption) (r DomainService) {
	r = DomainService{}
	r.Options = opts
	return
}

// Create Domain
func (r *DomainService) New(ctx context.Context, body DomainNewParams, opts ...option.RequestOption) (res *Domain, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	path := "v0/domains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Domain
func (r *DomainService) Get(ctx context.Context, domainID string, opts ...option.RequestOption) (res *Domain, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if domainID == "" {
		err = errors.New("missing required domain_id parameter")
		return
	}
	path := fmt.Sprintf("v0/domains/%s", url.PathEscape(domainID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Domains
func (r *DomainService) List(ctx context.Context, query DomainListParams, opts ...option.RequestOption) (res *ListDomains, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	path := "v0/domains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete Domain
func (r *DomainService) Delete(ctx context.Context, domainID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if domainID == "" {
		err = errors.New("missing required domain_id parameter")
		return
	}
	path := fmt.Sprintf("v0/domains/%s", url.PathEscape(domainID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get Zone File
func (r *DomainService) GetZoneFile(ctx context.Context, domainID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if domainID == "" {
		err = errors.New("missing required domain_id parameter")
		return
	}
	path := fmt.Sprintf("v0/domains/%s/zone-file", url.PathEscape(domainID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Verify Domain
func (r *DomainService) Verify(ctx context.Context, domainID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if domainID == "" {
		err = errors.New("missing required domain_id parameter")
		return
	}
	path := fmt.Sprintf("v0/domains/%s/verify", url.PathEscape(domainID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// The properties Domain, FeedbackEnabled are required.
type CreateDomainParam struct {
	// The name of the domain. (e.g., "example.com")
	Domain string `json:"domain" api:"required"`
	// Bounce and complaint notifications are sent to your inboxes.
	FeedbackEnabled bool `json:"feedback_enabled" api:"required"`
	paramObj
}

func (r CreateDomainParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateDomainParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateDomainParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Domain struct {
	// Time at which the domain was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The name of the domain. (e.g., " your-domain.com")
	DomainID string `json:"domain_id" api:"required"`
	// Bounce and complaint notifications are sent to your inboxes.
	FeedbackEnabled bool `json:"feedback_enabled" api:"required"`
	// A list of DNS records required to verify the domain.
	Records []DomainRecord `json:"records" api:"required"`
	// The verification status of the domain.
	//
	// Any of "NOT_STARTED", "PENDING", "INVALID", "FAILED", "VERIFYING", "VERIFIED".
	Status DomainStatus `json:"status" api:"required"`
	// Time at which the domain was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Client ID of domain.
	ClientID string `json:"client_id" api:"nullable"`
	// ID of pod.
	PodID string `json:"pod_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt       respjson.Field
		DomainID        respjson.Field
		FeedbackEnabled respjson.Field
		Records         respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		ClientID        respjson.Field
		PodID           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Domain) RawJSON() string { return r.JSON.raw }
func (r *Domain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainRecord struct {
	// The name or host of the record.
	Name string `json:"name" api:"required"`
	// The verification status of this specific record.
	//
	// Any of "MISSING", "INVALID", "VALID".
	Status string `json:"status" api:"required"`
	// The type of the DNS record.
	//
	// Any of "TXT", "CNAME", "MX".
	Type string `json:"type" api:"required"`
	// The value of the record.
	Value string `json:"value" api:"required"`
	// The priority of the MX record.
	Priority int64 `json:"priority" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		Priority    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DomainRecord) RawJSON() string { return r.JSON.raw }
func (r *DomainRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The verification status of the domain.
type DomainStatus string

const (
	DomainStatusNotStarted DomainStatus = "NOT_STARTED"
	DomainStatusPending    DomainStatus = "PENDING"
	DomainStatusInvalid    DomainStatus = "INVALID"
	DomainStatusFailed     DomainStatus = "FAILED"
	DomainStatusVerifying  DomainStatus = "VERIFYING"
	DomainStatusVerified   DomainStatus = "VERIFIED"
)

type ListDomains struct {
	// Number of items returned.
	Count int64 `json:"count" api:"required"`
	// Ordered by `created_at` descending.
	Domains []ListDomainsDomain `json:"domains" api:"required"`
	// Limit of number of items returned.
	Limit int64 `json:"limit" api:"nullable"`
	// Page token for pagination.
	NextPageToken string `json:"next_page_token" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count         respjson.Field
		Domains       respjson.Field
		Limit         respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListDomains) RawJSON() string { return r.JSON.raw }
func (r *ListDomains) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListDomainsDomain struct {
	// Time at which the domain was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The name of the domain. (e.g., " your-domain.com")
	DomainID string `json:"domain_id" api:"required"`
	// Bounce and complaint notifications are sent to your inboxes.
	FeedbackEnabled bool `json:"feedback_enabled" api:"required"`
	// Time at which the domain was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Client ID of domain.
	ClientID string `json:"client_id" api:"nullable"`
	// ID of pod.
	PodID string `json:"pod_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt       respjson.Field
		DomainID        respjson.Field
		FeedbackEnabled respjson.Field
		UpdatedAt       respjson.Field
		ClientID        respjson.Field
		PodID           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListDomainsDomain) RawJSON() string { return r.JSON.raw }
func (r *ListDomainsDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainNewParams struct {
	CreateDomain CreateDomainParam
	paramObj
}

func (r DomainNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateDomain)
}
func (r *DomainNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateDomain)
}

type DomainListParams struct {
	// Limit of number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page token for pagination.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainListParams]'s query parameters as `url.Values`.
func (r DomainListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
