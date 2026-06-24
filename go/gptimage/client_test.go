package gptimage

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/runapi-ai/core-sdk/go/core"
)

type stubHTTPClient struct {
	method   string
	path     string
	body     any
	response json.RawMessage
}

func (s *stubHTTPClient) Request(_ context.Context, method, path string, opts *core.HTTPRequestOptions) (json.RawMessage, error) {
	s.method = method
	s.path = path
	if opts != nil {
		s.body = opts.Body
	}
	return s.response, nil
}

func TestTextToImageCreate(t *testing.T) {
	stub := &stubHTTPClient{
		response: json.RawMessage(`{"id":"task_gen_123","status":"processing"}`),
	}
	client := NewClientWithHTTP(stub)
	resp, err := client.TextToImage.Create(context.Background(), TextToImageParams{
		Model:       "gpt-image-1.5",
		Prompt:      "a beautiful landscape",
		AspectRatio: "1:1",
		Quality:     "high",
	})
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "POST" || stub.path != "/api/v1/gpt_image/text_to_image" {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	body := stub.body.(map[string]any)
	if body["model"] != "gpt-image-1.5" {
		t.Fatalf("unexpected model: %v", body["model"])
	}
	if body["prompt"] != "a beautiful landscape" {
		t.Fatalf("unexpected prompt: %v", body["prompt"])
	}
	if resp.ID != "task_gen_123" {
		t.Fatalf("unexpected task ID: %v", resp.ID)
	}
}

func TestTextToImageGet(t *testing.T) {
	stub := &stubHTTPClient{
		response: json.RawMessage(`{"id":"task_gen_456","status":"completed","images":[{"url":"https://file.runapi.ai/result.png"}]}`),
	}
	client := NewClientWithHTTP(stub)
	resp, err := client.TextToImage.Get(context.Background(), "task_gen_abc")
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "GET" || stub.path != "/api/v1/gpt_image/text_to_image/task_gen_abc" {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	if resp.ID != "task_gen_456" {
		t.Fatalf("unexpected ID: %v", resp.ID)
	}
	if string(resp.Status) != "completed" {
		t.Fatalf("unexpected status: %v", resp.Status)
	}
	if len(resp.Images) != 1 || resp.Images[0].URL != "https://file.runapi.ai/result.png" {
		t.Fatalf("unexpected images: %v", resp.Images)
	}
}

func TestEditImageCreate(t *testing.T) {
	stub := &stubHTTPClient{
		response: json.RawMessage(`{"id":"task_edit_123","status":"processing"}`),
	}
	client := NewClientWithHTTP(stub)
	resp, err := client.EditImage.Create(context.Background(), EditImageParams{
		Model:           "gpt-image-1.5",
		Prompt:          "transform into oil painting",
		SourceImageURLs: []string{"https://cdn.runapi.ai/public/samples/photo.jpg"},
		AspectRatio:     "3:2",
		Quality:         "high",
	})
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "POST" || stub.path != "/api/v1/gpt_image/edit_image" {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	body := stub.body.(map[string]any)
	if body["model"] != "gpt-image-1.5" {
		t.Fatalf("unexpected model: %v", body["model"])
	}
	urls, ok := body["source_image_urls"].([]any)
	if !ok || len(urls) != 1 || urls[0] != "https://cdn.runapi.ai/public/samples/photo.jpg" {
		t.Fatalf("unexpected source_image_urls: %v", body["source_image_urls"])
	}
	if resp.ID != "task_edit_123" {
		t.Fatalf("unexpected task ID: %v", resp.ID)
	}
}

func TestEditImageGet(t *testing.T) {
	stub := &stubHTTPClient{
		response: json.RawMessage(`{"id":"task_edit_456","status":"completed","images":[{"url":"https://file.runapi.ai/edited.png"}]}`),
	}
	client := NewClientWithHTTP(stub)
	resp, err := client.EditImage.Get(context.Background(), "task_edit_abc")
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "GET" || stub.path != "/api/v1/gpt_image/edit_image/task_edit_abc" {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	if resp.ID != "task_edit_456" {
		t.Fatalf("unexpected ID: %v", resp.ID)
	}
	if string(resp.Status) != "completed" {
		t.Fatalf("unexpected status: %v", resp.Status)
	}
	if len(resp.Images) != 1 || resp.Images[0].URL != "https://file.runapi.ai/edited.png" {
		t.Fatalf("unexpected images: %v", resp.Images)
	}
}
