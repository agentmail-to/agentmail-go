// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package agentmail

import (
	"context"
	"net/http"
	"slices"

	"github.com/agentmail-to/agentmail-go/internal/apijson"
	"github.com/agentmail-to/agentmail-go/internal/requestconfig"
	"github.com/agentmail-to/agentmail-go/option"
	"github.com/agentmail-to/agentmail-go/packages/param"
	"github.com/agentmail-to/agentmail-go/packages/respjson"
)

// AgentService contains methods and other services that help with interacting with
// the agentmail API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentService] method instead.
type AgentService struct {
	Options []option.RequestOption
}

// NewAgentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAgentService(opts ...option.RequestOption) (r AgentService) {
	r = AgentService{}
	r.Options = opts
	return
}

// Create a new agent organization with an inbox and API key. A 6-digit OTP is sent
// to the human's email for verification.
//
// This endpoint is idempotent. Calling it again with the same `human_email` will
// rotate the API key and resend the OTP if expired.
//
// The returned API key has limited permissions until the organization is verified
// via the verify endpoint.
//
// **CLI:**
//
// ```bash
// agentmail agent sign-up --human-email user@example.com --username my-agent
// ```
func (r *AgentService) SignUp(ctx context.Context, body AgentSignUpParams, opts ...option.RequestOption) (res *AgentSignupResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	path := "v0/agent/sign-up"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Verify an agent organization using the 6-digit OTP sent to the human's email
// during sign-up.
//
// On success, the organization is upgraded from `agent_unverified` to
// `agent_verified`, the send allowlist is removed, and free plan entitlements are
// applied.
//
// The OTP expires after 24 hours and allows a maximum of 10 attempts.
//
// **CLI:**
//
// ```bash
// agentmail agent verify --otp-code 123456
// ```
func (r *AgentService) Verify(ctx context.Context, body AgentVerifyParams, opts ...option.RequestOption) (res *AgentVerifyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.agentmail.to/")}, opts...)
	path := "v0/agent/verify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Response after successful agent sign-up.
type AgentSignupResponse struct {
	// API key for authenticating subsequent requests. Store this securely, it cannot
	// be retrieved again.
	APIKey string `json:"api_key" api:"required"`
	// ID of the auto-created inbox.
	InboxID string `json:"inbox_id" api:"required"`
	// ID of the created organization.
	OrganizationID string `json:"organization_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey         respjson.Field
		InboxID        respjson.Field
		OrganizationID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentSignupResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentSignupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response after successful agent verification.
type AgentVerifyResponse struct {
	// Whether the organization was verified.
	Verified bool `json:"verified" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Verified    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentVerifyResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentVerifyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentSignUpParams struct {
	// Email address of the human who owns the agent. A 6-digit OTP will be sent to
	// this address.
	HumanEmail string `json:"human_email" api:"required"`
	// Username for the auto-created inbox (e.g. "my-agent" creates
	// my-agent@agentmail.to).
	Username string `json:"username" api:"required"`
	paramObj
}

func (r AgentSignUpParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentSignUpParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentSignUpParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentVerifyParams struct {
	// 6-digit verification code sent to the human's email address.
	OtpCode string `json:"otp_code" api:"required"`
	paramObj
}

func (r AgentVerifyParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentVerifyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentVerifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
