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

func TestAPIKeyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.APIKeys.New(context.TODO(), agentmail.APIKeyNewParams{
		Name: "name",
		Permissions: agentmail.APIKeyNewParamsPermissions{
			CreateAPIKey:    agentmail.Bool(true),
			CreateDomain:    agentmail.Bool(true),
			CreateDraft:     agentmail.Bool(true),
			CreateInbox:     agentmail.Bool(true),
			CreateListEntry: agentmail.Bool(true),
			CreatePod:       agentmail.Bool(true),
			CreateWebhook:   agentmail.Bool(true),
			DeleteAPIKey:    agentmail.Bool(true),
			DeleteDomain:    agentmail.Bool(true),
			DeleteDraft:     agentmail.Bool(true),
			DeleteInbox:     agentmail.Bool(true),
			DeleteListEntry: agentmail.Bool(true),
			DeletePod:       agentmail.Bool(true),
			DeleteThread:    agentmail.Bool(true),
			DeleteWebhook:   agentmail.Bool(true),
			ReadAPIKey:      agentmail.Bool(true),
			ReadBlocked:     agentmail.Bool(true),
			ReadDomain:      agentmail.Bool(true),
			ReadDraft:       agentmail.Bool(true),
			ReadInbox:       agentmail.Bool(true),
			ReadListEntry:   agentmail.Bool(true),
			ReadMessage:     agentmail.Bool(true),
			ReadMetrics:     agentmail.Bool(true),
			ReadPod:         agentmail.Bool(true),
			ReadSpam:        agentmail.Bool(true),
			ReadThread:      agentmail.Bool(true),
			ReadTrash:       agentmail.Bool(true),
			ReadWebhook:     agentmail.Bool(true),
			SendDraft:       agentmail.Bool(true),
			SendMessage:     agentmail.Bool(true),
			UpdateDomain:    agentmail.Bool(true),
			UpdateDraft:     agentmail.Bool(true),
			UpdateInbox:     agentmail.Bool(true),
			UpdateMessage:   agentmail.Bool(true),
			UpdateWebhook:   agentmail.Bool(true),
		},
	})
	if err != nil {
		var apierr *agentmail.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAPIKeyListWithOptionalParams(t *testing.T) {
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
	_, err := client.APIKeys.List(context.TODO(), agentmail.APIKeyListParams{
		Ascending: agentmail.Bool(true),
		Limit:     agentmail.Int(0),
		PageToken: agentmail.String("page_token"),
	})
	if err != nil {
		var apierr *agentmail.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
