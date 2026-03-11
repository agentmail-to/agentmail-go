// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package agentmail_test

import (
	"context"
	"os"
	"testing"

	"github.com/agentmail-to/agentmail-go"
	"github.com/agentmail-to/agentmail-go/internal/testutil"
	"github.com/agentmail-to/agentmail-go/option"
)

func TestUsage(t *testing.T) {
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
	listInboxes, err := client.Inboxes.List(context.TODO(), agentmail.InboxListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", listInboxes.Count)
}
