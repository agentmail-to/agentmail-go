// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package agentmail_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/agentmail-to/agentmail-go"
	"github.com/agentmail-to/agentmail-go/internal/testutil"
	"github.com/agentmail-to/agentmail-go/option"
)

func TestListNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := agentmail.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Lists.New(
		context.TODO(),
		agentmail.ListNewParamsTypeAllow,
		agentmail.ListNewParams{
			Direction: agentmail.ListNewParamsDirectionSend,
			Entry:     "entry",
			Reason:    agentmail.String("reason"),
		},
	)
	if err != nil {
		var apierr *agentmail.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestListListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := agentmail.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Lists.List(
		context.TODO(),
		agentmail.ListListParamsTypeAllow,
		agentmail.ListListParams{
			Direction: agentmail.ListListParamsDirectionSend,
			Limit:     agentmail.Int(0),
			PageToken: agentmail.String("page_token"),
		},
	)
	if err != nil {
		var apierr *agentmail.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestListDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := agentmail.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Lists.Delete(
		context.TODO(),
		"entry",
		agentmail.ListDeleteParams{
			Direction: agentmail.ListDeleteParamsDirectionSend,
			Type:      agentmail.ListDeleteParamsTypeAllow,
		},
	)
	if err != nil {
		var apierr *agentmail.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestListGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := agentmail.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Lists.Get(
		context.TODO(),
		"entry",
		agentmail.ListGetParams{
			Direction: agentmail.ListGetParamsDirectionSend,
			Type:      agentmail.ListGetParamsTypeAllow,
		},
	)
	if err != nil {
		var apierr *agentmail.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
