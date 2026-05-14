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

func TestGenerationsCreate(t *testing.T) {
	stub := &stubHTTPClient{
		response: json.RawMessage(`{"id":"task_gen_123","status":"processing"}`),
	}
	client := NewClientWithHTTP(stub)
	resp, err := client.Generations.Create(context.Background(), GenerationParams{
		Model:  "gpt-image/1.5-text-to-image",
		Prompt: "a beautiful landscape",
	})
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "POST" || stub.path != "/api/v1/gpt_image/generations" {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	body := stub.body.(map[string]any)
	if body["model"] != "gpt-image/1.5-text-to-image" {
		t.Fatalf("unexpected model: %v", body["model"])
	}
	if body["prompt"] != "a beautiful landscape" {
		t.Fatalf("unexpected prompt: %v", body["prompt"])
	}
	if resp.ID != "task_gen_123" {
		t.Fatalf("unexpected task ID: %v", resp.ID)
	}
}

func TestGenerationsGet(t *testing.T) {
	stub := &stubHTTPClient{
		response: json.RawMessage(`{"id":"task_gen_456","status":"completed","images":[{"url":"https://file.runapi.ai/result.png"}]}`),
	}
	client := NewClientWithHTTP(stub)
	resp, err := client.Generations.Get(context.Background(), "task_gen_abc")
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "GET" || stub.path != "/api/v1/gpt_image/generations/task_gen_abc" {
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

func TestEditsCreate(t *testing.T) {
	stub := &stubHTTPClient{
		response: json.RawMessage(`{"id":"task_edit_123","status":"processing"}`),
	}
	client := NewClientWithHTTP(stub)
	resp, err := client.Edits.Create(context.Background(), EditParams{
		Model:     "gpt-image/1.5-image-to-image",
		Prompt:    "transform into oil painting",
		InputURLs: []string{"https://example.com/photo.jpg"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "POST" || stub.path != "/api/v1/gpt_image/edits" {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	body := stub.body.(map[string]any)
	if body["model"] != "gpt-image/1.5-image-to-image" {
		t.Fatalf("unexpected model: %v", body["model"])
	}
	if resp.ID != "task_edit_123" {
		t.Fatalf("unexpected task ID: %v", resp.ID)
	}
}

func TestEditsGet(t *testing.T) {
	stub := &stubHTTPClient{
		response: json.RawMessage(`{"id":"task_edit_456","status":"completed","images":[{"url":"https://file.runapi.ai/edited.png"}]}`),
	}
	client := NewClientWithHTTP(stub)
	resp, err := client.Edits.Get(context.Background(), "task_edit_abc")
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "GET" || stub.path != "/api/v1/gpt_image/edits/task_edit_abc" {
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
