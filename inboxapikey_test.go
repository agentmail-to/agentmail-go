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

func TestInboxAPIKeyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Inboxes.APIKeys.New(
		context.TODO(),
		"inbox_id",
		agentmail.InboxAPIKeyNewParams{
			Name: "name",
			Permissions: agentmail.InboxAPIKeyNewParamsPermissions{
				APIKeyCreate:     agentmail.Bool(true),
				APIKeyDelete:     agentmail.Bool(true),
				APIKeyRead:       agentmail.Bool(true),
				DomainCreate:     agentmail.Bool(true),
				DomainDelete:     agentmail.Bool(true),
				DomainRead:       agentmail.Bool(true),
				DomainUpdate:     agentmail.Bool(true),
				DraftCreate:      agentmail.Bool(true),
				DraftDelete:      agentmail.Bool(true),
				DraftRead:        agentmail.Bool(true),
				DraftSend:        agentmail.Bool(true),
				DraftUpdate:      agentmail.Bool(true),
				InboxCreate:      agentmail.Bool(true),
				InboxDelete:      agentmail.Bool(true),
				InboxRead:        agentmail.Bool(true),
				InboxUpdate:      agentmail.Bool(true),
				LabelBlockedRead: agentmail.Bool(true),
				LabelSpamRead:    agentmail.Bool(true),
				LabelTrashRead:   agentmail.Bool(true),
				ListEntryCreate:  agentmail.Bool(true),
				ListEntryDelete:  agentmail.Bool(true),
				ListEntryRead:    agentmail.Bool(true),
				MessageRead:      agentmail.Bool(true),
				MessageSend:      agentmail.Bool(true),
				MessageUpdate:    agentmail.Bool(true),
				MetricsRead:      agentmail.Bool(true),
				PodCreate:        agentmail.Bool(true),
				PodDelete:        agentmail.Bool(true),
				PodRead:          agentmail.Bool(true),
				ThreadDelete:     agentmail.Bool(true),
				ThreadRead:       agentmail.Bool(true),
				WebhookCreate:    agentmail.Bool(true),
				WebhookDelete:    agentmail.Bool(true),
				WebhookRead:      agentmail.Bool(true),
				WebhookUpdate:    agentmail.Bool(true),
			},
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

func TestInboxAPIKeyListWithOptionalParams(t *testing.T) {
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
	_, err := client.Inboxes.APIKeys.List(
		context.TODO(),
		"inbox_id",
		agentmail.InboxAPIKeyListParams{
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

func TestInboxAPIKeyDelete(t *testing.T) {
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
	err := client.Inboxes.APIKeys.Delete(
		context.TODO(),
		"api_key_id",
		agentmail.InboxAPIKeyDeleteParams{
			InboxID: "inbox_id",
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
