// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/hanzoai/go-sdk"
	"github.com/hanzoai/go-sdk/internal/testutil"
	"github.com/hanzoai/go-sdk/option"
)

func TestTeamModelAdd(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Team.Model.Add(context.TODO(), hanzoai.TeamModelAddParams{
		Models: hanzoai.F([]string{"string"}),
		TeamID: hanzoai.F("team_id"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamModelRemove(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Team.Model.Remove(context.TODO(), hanzoai.TeamModelRemoveParams{
		Models: hanzoai.F([]string{"string"}),
		TeamID: hanzoai.F("team_id"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
