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

func TestInboxMessageGet(t *testing.T) {
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
	_, err := client.Inboxes.Messages.Get(
		context.TODO(),
		"message_id",
		agentmail.InboxMessageGetParams{
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

func TestInboxMessageUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Inboxes.Messages.Update(
		context.TODO(),
		"message_id",
		agentmail.InboxMessageUpdateParams{
			InboxID: "inbox_id",
			UpdateMessage: agentmail.UpdateMessageParam{
				AddLabels:    []string{"string"},
				RemoveLabels: []string{"string"},
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

func TestInboxMessageListWithOptionalParams(t *testing.T) {
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
	_, err := client.Inboxes.Messages.List(
		context.TODO(),
		"inbox_id",
		agentmail.InboxMessageListParams{
			After:       agentmail.Time(time.Now()),
			Ascending:   agentmail.Bool(true),
			Before:      agentmail.Time(time.Now()),
			IncludeSpam: agentmail.Bool(true),
			Labels:      []string{"string"},
			Limit:       agentmail.Int(0),
			PageToken:   agentmail.String("page_token"),
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

func TestInboxMessageForwardWithOptionalParams(t *testing.T) {
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
	_, err := client.Inboxes.Messages.Forward(
		context.TODO(),
		"message_id",
		agentmail.InboxMessageForwardParams{
			InboxID: "inbox_id",
			SendMessageRequest: agentmail.SendMessageRequestParam{
				Attachments: []agentmail.SendAttachmentParam{{
					Content:            agentmail.String("content"),
					ContentDisposition: agentmail.AttachmentContentDispositionInline,
					ContentID:          agentmail.String("content_id"),
					ContentType:        agentmail.String("content_type"),
					Filename:           agentmail.String("filename"),
					URL:                agentmail.String("url"),
				}},
				Bcc: agentmail.AddressesUnionParam{
					OfString: agentmail.String("string"),
				},
				Cc: agentmail.AddressesUnionParam{
					OfString: agentmail.String("string"),
				},
				Headers: map[string]string{
					"foo": "string",
				},
				HTML:   agentmail.String("html"),
				Labels: []string{"string"},
				ReplyTo: agentmail.AddressesUnionParam{
					OfString: agentmail.String("string"),
				},
				Subject: agentmail.String("subject"),
				Text:    agentmail.String("text"),
				To: agentmail.AddressesUnionParam{
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

func TestInboxMessageGetAttachment(t *testing.T) {
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
	_, err := client.Inboxes.Messages.GetAttachment(
		context.TODO(),
		"attachment_id",
		agentmail.InboxMessageGetAttachmentParams{
			InboxID:   "inbox_id",
			MessageID: "message_id",
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

func TestInboxMessageGetRaw(t *testing.T) {
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
	err := client.Inboxes.Messages.GetRaw(
		context.TODO(),
		"message_id",
		agentmail.InboxMessageGetRawParams{
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

func TestInboxMessageReplyWithOptionalParams(t *testing.T) {
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
	_, err := client.Inboxes.Messages.Reply(
		context.TODO(),
		"message_id",
		agentmail.InboxMessageReplyParams{
			InboxID: "inbox_id",
			Attachments: []agentmail.SendAttachmentParam{{
				Content:            agentmail.String("content"),
				ContentDisposition: agentmail.AttachmentContentDispositionInline,
				ContentID:          agentmail.String("content_id"),
				ContentType:        agentmail.String("content_type"),
				Filename:           agentmail.String("filename"),
				URL:                agentmail.String("url"),
			}},
			Bcc: agentmail.AddressesUnionParam{
				OfString: agentmail.String("string"),
			},
			Cc: agentmail.AddressesUnionParam{
				OfString: agentmail.String("string"),
			},
			Headers: map[string]string{
				"foo": "string",
			},
			HTML:     agentmail.String("html"),
			Labels:   []string{"string"},
			ReplyAll: agentmail.Bool(true),
			ReplyTo: agentmail.AddressesUnionParam{
				OfString: agentmail.String("string"),
			},
			Text: agentmail.String("text"),
			To: agentmail.AddressesUnionParam{
				OfString: agentmail.String("string"),
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

func TestInboxMessageReplyAllWithOptionalParams(t *testing.T) {
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
	_, err := client.Inboxes.Messages.ReplyAll(
		context.TODO(),
		"message_id",
		agentmail.InboxMessageReplyAllParams{
			InboxID: "inbox_id",
			Attachments: []agentmail.SendAttachmentParam{{
				Content:            agentmail.String("content"),
				ContentDisposition: agentmail.AttachmentContentDispositionInline,
				ContentID:          agentmail.String("content_id"),
				ContentType:        agentmail.String("content_type"),
				Filename:           agentmail.String("filename"),
				URL:                agentmail.String("url"),
			}},
			Headers: map[string]string{
				"foo": "string",
			},
			HTML:   agentmail.String("html"),
			Labels: []string{"string"},
			ReplyTo: agentmail.AddressesUnionParam{
				OfString: agentmail.String("string"),
			},
			Text: agentmail.String("text"),
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

func TestInboxMessageSendWithOptionalParams(t *testing.T) {
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
	_, err := client.Inboxes.Messages.Send(
		context.TODO(),
		"inbox_id",
		agentmail.InboxMessageSendParams{
			SendMessageRequest: agentmail.SendMessageRequestParam{
				Attachments: []agentmail.SendAttachmentParam{{
					Content:            agentmail.String("content"),
					ContentDisposition: agentmail.AttachmentContentDispositionInline,
					ContentID:          agentmail.String("content_id"),
					ContentType:        agentmail.String("content_type"),
					Filename:           agentmail.String("filename"),
					URL:                agentmail.String("url"),
				}},
				Bcc: agentmail.AddressesUnionParam{
					OfString: agentmail.String("string"),
				},
				Cc: agentmail.AddressesUnionParam{
					OfString: agentmail.String("string"),
				},
				Headers: map[string]string{
					"foo": "string",
				},
				HTML:   agentmail.String("html"),
				Labels: []string{"string"},
				ReplyTo: agentmail.AddressesUnionParam{
					OfString: agentmail.String("string"),
				},
				Subject: agentmail.String("subject"),
				Text:    agentmail.String("text"),
				To: agentmail.AddressesUnionParam{
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
