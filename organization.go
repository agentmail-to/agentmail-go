// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package agentmail

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/agentmail-go/internal/apijson"
	"github.com/stainless-sdks/agentmail-go/internal/requestconfig"
	"github.com/stainless-sdks/agentmail-go/option"
	"github.com/stainless-sdks/agentmail-go/packages/respjson"
)

// OrganizationService contains methods and other services that help with
// interacting with the agentmail API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	Options []option.RequestOption
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r OrganizationService) {
	r = OrganizationService{}
	r.Options = opts
	return
}

// Get the current organization.
func (r *OrganizationService) Get(ctx context.Context, opts ...option.RequestOption) (res *OrganizationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	path := "v0/organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Organization details with usage limits and counts.
type OrganizationGetResponse struct {
	// Time at which organization was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Current number of domains.
	DomainCount int64 `json:"domain_count" api:"required"`
	// Current number of inboxes.
	InboxCount int64 `json:"inbox_count" api:"required"`
	// ID of organization.
	OrganizationID string `json:"organization_id" api:"required"`
	// Time at which organization was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Maximum number of domains allowed.
	DomainLimit int64 `json:"domain_limit" api:"nullable"`
	// Maximum number of inboxes allowed.
	InboxLimit int64 `json:"inbox_limit" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt      respjson.Field
		DomainCount    respjson.Field
		InboxCount     respjson.Field
		OrganizationID respjson.Field
		UpdatedAt      respjson.Field
		DomainLimit    respjson.Field
		InboxLimit     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
