// Package gptimage provides the GPT Image 1.5 image generation API client.
//
//	client, err := gptimage.NewClient(option.WithAPIKey("sk-your-api-key"))
//	result, err := client.Generations.Run(ctx, gptimage.GenerationParams{
//	    Model: "gpt-image/1.5-text-to-image", Prompt: "A beautiful landscape",
//	})
package gptimage

import (
	"context"

	"github.com/runapi-ai/core-sdk/go/core"
	"github.com/runapi-ai/core-sdk/go/option"
)

const (
	generationsPath = "/api/v1/gpt_image/generations"
	editsPath       = "/api/v1/gpt_image/edits"
)

// Client is the GPT Image 1.5 image generation API client.
type Client struct {
	// Generations provides text-to-image generation operations.
	Generations *Generations
	// Edits provides image-to-image edit operations.
	Edits *Edits
}

// NewClient creates a GPT Image client with the given options.
func NewClient(opts ...option.ClientOption) (*Client, error) {
	resolved, err := option.ResolveClientOptions(opts...)
	if err != nil {
		return nil, err
	}
	httpClient, err := core.NewHTTPClient(resolved)
	if err != nil {
		return nil, err
	}
	return NewClientWithHTTP(httpClient), nil
}

// NewClientWithHTTP creates a GPT Image client with a pre-configured HTTP transport.
func NewClientWithHTTP(httpClient core.HTTPClient) *Client {
	return &Client{
		Generations: &Generations{http: httpClient},
		Edits:       &Edits{http: httpClient},
	}
}

// Generations handles text-to-image generation tasks.
type Generations struct{ http core.HTTPClient }

func (r *Generations) Create(ctx context.Context, params GenerationParams, opts ...option.RequestOption) (*core.TaskCreateResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.PostJSON[core.TaskCreateResponse](ctx, r.http, generationsPath, core.CompactParams(params), requestOptions)
}
func (r *Generations) Get(ctx context.Context, id string, opts ...option.RequestOption) (*GenerationResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.GetJSON[GenerationResponse](ctx, r.http, core.ResourcePath(generationsPath, id), requestOptions)
}
func (r *Generations) Run(ctx context.Context, params GenerationParams, opts ...option.RequestOption) (*GenerationResponse, error) {
	_, pollingOptions := option.ResolveRequestOptions(opts...)
	return core.RunAsync(ctx, func(ctx context.Context) (*core.TaskCreateResponse, error) { return r.Create(ctx, params, opts...) }, func(ctx context.Context, id string) (*GenerationResponse, error) { return r.Get(ctx, id, opts...) }, pollingOptions)
}

// Edits handles image-to-image edit tasks.
type Edits struct{ http core.HTTPClient }

func (r *Edits) Create(ctx context.Context, params EditParams, opts ...option.RequestOption) (*core.TaskCreateResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.PostJSON[core.TaskCreateResponse](ctx, r.http, editsPath, core.CompactParams(params), requestOptions)
}
func (r *Edits) Get(ctx context.Context, id string, opts ...option.RequestOption) (*EditResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.GetJSON[EditResponse](ctx, r.http, core.ResourcePath(editsPath, id), requestOptions)
}
func (r *Edits) Run(ctx context.Context, params EditParams, opts ...option.RequestOption) (*EditResponse, error) {
	_, pollingOptions := option.ResolveRequestOptions(opts...)
	return core.RunAsync(ctx, func(ctx context.Context) (*core.TaskCreateResponse, error) { return r.Create(ctx, params, opts...) }, func(ctx context.Context, id string) (*EditResponse, error) { return r.Get(ctx, id, opts...) }, pollingOptions)
}
