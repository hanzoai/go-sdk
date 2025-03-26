// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/Hanzo-AI-go/internal/requestconfig"
	"github.com/stainless-sdks/Hanzo-AI-go/option"
)

// SettingService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingService] method instead.
type SettingService struct {
	Options []option.RequestOption
}

// NewSettingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSettingService(opts ...option.RequestOption) (r *SettingService) {
	r = &SettingService{}
	r.Options = opts
	return
}

// Returns a list of llm level settings
//
// This is useful for debugging and ensuring the proxy server is configured
// correctly.
//
// Response schema:
//
// ```
//
//	{
//	    "alerting": _alerting,
//	    "llm.callbacks": llm_callbacks,
//	    "llm.input_callback": llm_input_callbacks,
//	    "llm.failure_callback": llm_failure_callbacks,
//	    "llm.success_callback": llm_success_callbacks,
//	    "llm._async_success_callback": llm_async_success_callbacks,
//	    "llm._async_failure_callback": llm_async_failure_callbacks,
//	    "llm._async_input_callback": llm_async_input_callbacks,
//	    "all_llm_callbacks": all_llm_callbacks,
//	    "num_callbacks": len(all_llm_callbacks),
//	    "num_alerting": _num_alerting,
//	    "llm.request_timeout": llm.request_timeout,
//	}
//
// ```
func (r *SettingService) Get(ctx context.Context, opts ...option.RequestOption) (res *SettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SettingGetResponse = interface{}
