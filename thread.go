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

	"github.com/agentmail-to/agentmail-go/internal/apiquery"
	"github.com/agentmail-to/agentmail-go/internal/requestconfig"
	"github.com/agentmail-to/agentmail-go/option"
	"github.com/agentmail-to/agentmail-go/packages/param"
)

// ThreadService contains methods and other services that help with interacting
// with the agentmail API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreadService] method instead.
type ThreadService struct {
	Options []option.RequestOption
}

// NewThreadService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewThreadService(opts ...option.RequestOption) (r ThreadService) {
	r = ThreadService{}
	r.Options = opts
	return
}

// Get Thread
func (r *ThreadService) Get(ctx context.Context, threadID string, opts ...option.RequestOption) (res *Thread, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return
	}
	path := fmt.Sprintf("v0/threads/%s", url.PathEscape(threadID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Threads
func (r *ThreadService) List(ctx context.Context, query ThreadListParams, opts ...option.RequestOption) (res *ListThreads, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	path := "v0/threads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Attachment
func (r *ThreadService) GetAttachment(ctx context.Context, attachmentID string, query ThreadGetAttachmentParams, opts ...option.RequestOption) (res *AttachmentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if query.ThreadID == "" {
		err = errors.New("missing required thread_id parameter")
		return
	}
	if attachmentID == "" {
		err = errors.New("missing required attachment_id parameter")
		return
	}
	path := fmt.Sprintf("v0/threads/%s/attachments/%s", url.PathEscape(query.ThreadID), url.PathEscape(attachmentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ThreadListParams struct {
	// Timestamp after which to filter by.
	After param.Opt[time.Time] `query:"after,omitzero" format:"date-time" json:"-"`
	// Sort in ascending temporal order.
	Ascending param.Opt[bool] `query:"ascending,omitzero" json:"-"`
	// Timestamp before which to filter by.
	Before param.Opt[time.Time] `query:"before,omitzero" format:"date-time" json:"-"`
	// Include spam in results.
	IncludeSpam param.Opt[bool] `query:"include_spam,omitzero" json:"-"`
	// Limit of number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page token for pagination.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	// Labels to filter by.
	Labels []string `query:"labels,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ThreadListParams]'s query parameters as `url.Values`.
func (r ThreadListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ThreadGetAttachmentParams struct {
	// ID of thread.
	ThreadID string `path:"thread_id" api:"required" json:"-"`
	paramObj
}
