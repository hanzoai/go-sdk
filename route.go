// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/Hanzo-AI-go/internal/requestconfig"
	"github.com/stainless-sdks/Hanzo-AI-go/option"
)

// RouteService contains methods and other services that help with interacting with
// the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRouteService] method instead.
type RouteService struct {
	Options []option.RequestOption
}

// NewRouteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRouteService(opts ...option.RequestOption) (r *RouteService) {
	r = &RouteService{}
	r.Options = opts
	return
}

// Get a list of available routes in the FastAPI application.
func (r *RouteService) List(ctx context.Context, opts ...option.RequestOption) (res *RouteListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "routes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RouteListResponse = interface{}
