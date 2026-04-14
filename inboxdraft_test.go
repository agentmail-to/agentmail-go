// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package agentmail_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/agentmail-to/agentmail-go"
	"github.com/agentmail-to/agentmail-go/internal/testutil"
	"github.com/agentmail-to/agentmail-go/option"
)

func TestInboxDraftNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Inboxes.Drafts.New(
		context.TODO(),
		"inbox_id",
		agentmail.InboxDraftNewParams{
			Attachments: []agentmail.SendAttachmentParam{{
				Content:            agentmail.String("content"),
				ContentDisposition: agentmail.AttachmentContentDispositionInline,
				ContentID:          agentmail.String("content_id"),
				ContentType:        agentmail.String("content_type"),
				Filename:           agentmail.String("filename"),
				URL:                agentmail.String("url"),
			}},
			Bcc:       []string{"string"},
			Cc:        []string{"string"},
			ClientID:  agentmail.String("client_id"),
			HTML:      agentmail.String("html"),
			InReplyTo: agentmail.String("in_reply_to"),
			Labels:    []string{"string"},
			ReplyTo:   []string{"string"},
			SendAt:    agentmail.Time(time.Now()),
			Subject:   agentmail.String("subject"),
			Text:      agentmail.String("text"),
			To:        []string{"string"},
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

func TestInboxDraftUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Inboxes.Drafts.Update(
		context.TODO(),
		"draft_id",
		agentmail.InboxDraftUpdateParams{
			InboxID: "inbox_id",
			Bcc:     []string{"string"},
			Cc:      []string{"string"},
			HTML:    agentmail.String("html"),
			ReplyTo: []string{"string"},
			SendAt:  agentmail.Time(time.Now()),
			Subject: agentmail.String("subject"),
			Text:    agentmail.String("text"),
			To:      []string{"string"},
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

func TestInboxDraftListWithOptionalParams(t *testing.T) {
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
	_, err := client.Inboxes.Drafts.List(
		context.TODO(),
		"inbox_id",
		agentmail.InboxDraftListParams{
			After:     agentmail.Time(time.Now()),
			Ascending: agentmail.Bool(true),
			Before:    agentmail.Time(time.Now()),
			Labels:    []string{"string"},
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

func TestInboxDraftDelete(t *testing.T) {
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
	err := client.Inboxes.Drafts.Delete(
		context.TODO(),
		"draft_id",
		agentmail.InboxDraftDeleteParams{
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

func TestInboxDraftGet(t *testing.T) {
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
	_, err := client.Inboxes.Drafts.Get(
		context.TODO(),
		"draft_id",
		agentmail.InboxDraftGetParams{
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

func TestInboxDraftGetAttachment(t *testing.T) {
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
	_, err := client.Inboxes.Drafts.GetAttachment(
		context.TODO(),
		"attachment_id",
		agentmail.InboxDraftGetAttachmentParams{
			InboxID: "inbox_id",
			DraftID: "draft_id",
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

func TestInboxDraftSendWithOptionalParams(t *testing.T) {
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
	_, err := client.Inboxes.Drafts.Send(
		context.TODO(),
		"draft_id",
		agentmail.InboxDraftSendParams{
			InboxID: "inbox_id",
			UpdateMessage: agentmail.UpdateMessageParam{
				AddLabels: agentmail.UpdateMessageAddLabelsUnionParam{
					OfString: agentmail.String("string"),
				},
				RemoveLabels: agentmail.UpdateMessageRemoveLabelsUnionParam{
					OfString: agentmail.String("string"),
				},
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
