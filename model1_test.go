// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/Hanzo-AI-go"
	"github.com/stainless-sdks/Hanzo-AI-go/internal/testutil"
	"github.com/stainless-sdks/Hanzo-AI-go/option"
	"github.com/stainless-sdks/Hanzo-AI-go/shared"
)

func TestModelNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Model.New(context.TODO(), hanzoai.ModelNewParams{
		LlmParams: hanzoai.F(hanzoai.ModelNewParamsLlmParams{
			Model:                            hanzoai.F("model"),
			APIBase:                          hanzoai.F("api_base"),
			APIKey:                           hanzoai.F("api_key"),
			APIVersion:                       hanzoai.F("api_version"),
			AwsAccessKeyID:                   hanzoai.F("aws_access_key_id"),
			AwsRegionName:                    hanzoai.F("aws_region_name"),
			AwsSecretAccessKey:               hanzoai.F("aws_secret_access_key"),
			BudgetDuration:                   hanzoai.F("budget_duration"),
			ConfigurableClientsideAuthParams: hanzoai.F([]hanzoai.ModelNewParamsLlmParamsConfigurableClientsideAuthParamUnion{shared.UnionString("string")}),
			CustomLlmProvider:                hanzoai.F("custom_llm_provider"),
			InputCostPerSecond:               hanzoai.F(0.000000),
			InputCostPerToken:                hanzoai.F(0.000000),
			LlmTraceID:                       hanzoai.F("llm_trace_id"),
			MaxBudget:                        hanzoai.F(0.000000),
			MaxFileSizeMB:                    hanzoai.F(0.000000),
			MaxRetries:                       hanzoai.F(int64(0)),
			MergeReasoningContentInChoices:   hanzoai.F(true),
			ModelInfo:                        hanzoai.F[any](map[string]interface{}{}),
			Organization:                     hanzoai.F("organization"),
			OutputCostPerSecond:              hanzoai.F(0.000000),
			OutputCostPerToken:               hanzoai.F(0.000000),
			RegionName:                       hanzoai.F("region_name"),
			Rpm:                              hanzoai.F(int64(0)),
			StreamTimeout:                    hanzoai.F[hanzoai.ModelNewParamsLlmParamsStreamTimeoutUnion](shared.UnionFloat(0.000000)),
			Timeout:                          hanzoai.F[hanzoai.ModelNewParamsLlmParamsTimeoutUnion](shared.UnionFloat(0.000000)),
			Tpm:                              hanzoai.F(int64(0)),
			UseInPassThrough:                 hanzoai.F(true),
			VertexCredentials:                hanzoai.F[any](map[string]interface{}{}),
			VertexLocation:                   hanzoai.F("vertex_location"),
			VertexProject:                    hanzoai.F("vertex_project"),
			WatsonxRegionName:                hanzoai.F("watsonx_region_name"),
		}),
		ModelInfo: hanzoai.F(hanzoai.ModelInfoParam{
			ID:                  hanzoai.F("id"),
			BaseModel:           hanzoai.F("base_model"),
			CreatedAt:           hanzoai.F(time.Now()),
			CreatedBy:           hanzoai.F("created_by"),
			DBModel:             hanzoai.F(true),
			TeamID:              hanzoai.F("team_id"),
			TeamPublicModelName: hanzoai.F("team_public_model_name"),
			Tier:                hanzoai.F(hanzoai.ModelInfoTierFree),
			UpdatedAt:           hanzoai.F(time.Now()),
			UpdatedBy:           hanzoai.F("updated_by"),
		}),
		ModelName: hanzoai.F("model_name"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestModelDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Model.Delete(context.TODO(), hanzoai.ModelDeleteParams{
		ID: hanzoai.F("id"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
