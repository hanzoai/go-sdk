// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/Hanzo-AI-go"
	"github.com/stainless-sdks/Hanzo-AI-go/internal/testutil"
	"github.com/stainless-sdks/Hanzo-AI-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hanzoai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	response, err := client.GetHome(context.TODO())
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", response)
}
