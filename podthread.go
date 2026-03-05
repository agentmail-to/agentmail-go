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

// PodThreadService contains methods and other services that help with interacting
// with the agentmail API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPodThreadService] method instead.
type PodThreadService struct {
	Options []option.RequestOption
}

// NewPodThreadService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPodThreadService(opts ...option.RequestOption) (r PodThreadService) {
	r = PodThreadService{}
	r.Options = opts
	return
}

// Get Thread
func (r *PodThreadService) Get(ctx context.Context, threadID string, query PodThreadGetParams, opts ...option.RequestOption) (res *Thread, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if query.PodID == "" {
		err = errors.New("missing required pod_id parameter")
		return
	}
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return
	}
	path := fmt.Sprintf("v0/pods/%s/threads/%s", url.PathEscape(query.PodID), url.PathEscape(threadID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Threads
func (r *PodThreadService) List(ctx context.Context, podID string, query PodThreadListParams, opts ...option.RequestOption) (res *ListThreads, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if podID == "" {
		err = errors.New("missing required pod_id parameter")
		return
	}
	path := fmt.Sprintf("v0/pods/%s/threads", url.PathEscape(podID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Attachment
func (r *PodThreadService) GetAttachment(ctx context.Context, attachmentID string, query PodThreadGetAttachmentParams, opts ...option.RequestOption) (res *AttachmentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	if query.PodID == "" {
		err = errors.New("missing required pod_id parameter")
		return
	}
	if query.ThreadID == "" {
		err = errors.New("missing required thread_id parameter")
		return
	}
	if attachmentID == "" {
		err = errors.New("missing required attachment_id parameter")
		return
	}
	path := fmt.Sprintf("v0/pods/%s/threads/%s/attachments/%s", url.PathEscape(query.PodID), url.PathEscape(query.ThreadID), url.PathEscape(attachmentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PodThreadGetParams struct {
	// ID of pod.
	PodID string `path:"pod_id" api:"required" json:"-"`
	paramObj
}

type PodThreadListParams struct {
	// Timestamp after which to filter by.
	After param.Opt[time.Time] `query:"after,omitzero" format:"date-time" json:"-"`
	// Sort in ascending temporal order.
	Ascending param.Opt[bool] `query:"ascending,omitzero" json:"-"`
	// Timestamp before which to filter by.
	Before param.Opt[time.Time] `query:"before,omitzero" format:"date-time" json:"-"`
	// Include blocked in results.
	IncludeBlocked param.Opt[bool] `query:"include_blocked,omitzero" json:"-"`
	// Include spam in results.
	IncludeSpam param.Opt[bool] `query:"include_spam,omitzero" json:"-"`
	// Include trash in results.
	IncludeTrash param.Opt[bool] `query:"include_trash,omitzero" json:"-"`
	// Limit of number of items returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page token for pagination.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	// Labels to filter by.
	Labels []string `query:"labels,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PodThreadListParams]'s query parameters as `url.Values`.
func (r PodThreadListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PodThreadGetAttachmentParams struct {
	// ID of pod.
	PodID string `path:"pod_id" api:"required" json:"-"`
	// ID of thread.
	ThreadID string `path:"thread_id" api:"required" json:"-"`
	paramObj
}
