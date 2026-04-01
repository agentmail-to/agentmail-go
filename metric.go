// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package agentmail

import (
	"context"
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

// MetricService contains methods and other services that help with interacting
// with the agentmail API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMetricService] method instead.
type MetricService struct {
	Options []option.RequestOption
}

// NewMetricService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMetricService(opts ...option.RequestOption) (r MetricService) {
	r = MetricService{}
	r.Options = opts
	return
}

// **CLI:**
//
// ```bash
// agentmail metrics list
// ```
func (r *MetricService) List(ctx context.Context, query MetricListParams, opts ...option.RequestOption) (res *MetricListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	path := "v0/metrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type MetricListResponse map[string][]MetricListResponseItem

type MetricListResponseItem struct {
	// Count of events in the bucket.
	Count int64 `json:"count" api:"required"`
	// Timestamp of the bucket.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *MetricListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MetricListParams struct {
	// Sort in descending order.
	Descending param.Opt[bool] `query:"descending,omitzero" json:"-"`
	// End timestamp for the query.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date-time" json:"-"`
	// Limit on number of buckets to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Period in number of seconds for the query.
	Period param.Opt[string] `query:"period,omitzero" json:"-"`
	// Start timestamp for the query.
	Start param.Opt[time.Time] `query:"start,omitzero" format:"date-time" json:"-"`
	// List of metric event types to query.
	EventTypes []MetricEventType `query:"event_types,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MetricListParams]'s query parameters as `url.Values`.
func (r MetricListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
