// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package agentmail

import (
	"github.com/agentmail-to/agentmail-go/option"
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
