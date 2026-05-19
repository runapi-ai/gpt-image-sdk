// Package gptimage provides the GPT Image 1.5 image generation API client.
//
//	client, err := gptimage.NewClient(option.WithAPIKey("sk-your-api-key"))
//	result, err := client.TextToImage.Run(ctx, gptimage.TextToImageParams{
//	    Model: "gpt-image/1.5-text-to-image", Prompt: "A beautiful landscape",
//	})
package gptimage

import (
	"context"

	"github.com/runapi-ai/core-sdk/go/core"
	"github.com/runapi-ai/core-sdk/go/option"
)

const (
	textToImagePath = "/api/v1/gpt_image/text_to_image"
	editImagePath   = "/api/v1/gpt_image/edit_image"
)

// Client is the GPT Image 1.5 image generation API client.
type Client struct {
	// TextToImage provides text-to-image generation operations.
	TextToImage *TextToImage
	// EditImage provides image-to-image edit operations.
	EditImage *EditImage
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
		TextToImage: &TextToImage{http: httpClient},
		EditImage:   &EditImage{http: httpClient},
	}
}

// TextToImage handles text-to-image generation tasks.
type TextToImage struct{ http core.HTTPClient }

func (r *TextToImage) Create(ctx context.Context, params TextToImageParams, opts ...option.RequestOption) (*core.TaskCreateResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.PostJSON[core.TaskCreateResponse](ctx, r.http, textToImagePath, core.CompactParams(params), requestOptions)
}
func (r *TextToImage) Get(ctx context.Context, id string, opts ...option.RequestOption) (*TextToImageResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.GetJSON[TextToImageResponse](ctx, r.http, core.ResourcePath(textToImagePath, id), requestOptions)
}
func (r *TextToImage) Run(ctx context.Context, params TextToImageParams, opts ...option.RequestOption) (*TextToImageResponse, error) {
	_, pollingOptions := option.ResolveRequestOptions(opts...)
	return core.RunAsync(ctx, func(ctx context.Context) (*core.TaskCreateResponse, error) { return r.Create(ctx, params, opts...) }, func(ctx context.Context, id string) (*TextToImageResponse, error) { return r.Get(ctx, id, opts...) }, pollingOptions)
}

// EditImage handles image-to-image edit tasks.
type EditImage struct{ http core.HTTPClient }

func (r *EditImage) Create(ctx context.Context, params EditImageParams, opts ...option.RequestOption) (*core.TaskCreateResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.PostJSON[core.TaskCreateResponse](ctx, r.http, editImagePath, core.CompactParams(params), requestOptions)
}
func (r *EditImage) Get(ctx context.Context, id string, opts ...option.RequestOption) (*EditImageResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.GetJSON[EditImageResponse](ctx, r.http, core.ResourcePath(editImagePath, id), requestOptions)
}
func (r *EditImage) Run(ctx context.Context, params EditImageParams, opts ...option.RequestOption) (*EditImageResponse, error) {
	_, pollingOptions := option.ResolveRequestOptions(opts...)
	return core.RunAsync(ctx, func(ctx context.Context) (*core.TaskCreateResponse, error) { return r.Create(ctx, params, opts...) }, func(ctx context.Context, id string) (*EditImageResponse, error) { return r.Get(ctx, id, opts...) }, pollingOptions)
}
