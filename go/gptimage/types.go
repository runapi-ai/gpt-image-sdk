package gptimage

// TaskStatus represents the lifecycle state of an async generation task
// (e.g. "pending", "processing", "completed", "failed").
type TaskStatus string

// TextToImageParams holds the inputs for a text-to-image generation request.
// Unlike some image generation APIs, both AspectRatio and Quality are required fields.
type TextToImageParams struct {
	Model       string `json:"model" help:"required; model slug"`
	Prompt      string `json:"prompt" help:"required; text description of the desired image"`
	AspectRatio string `json:"aspect_ratio" help:"required; output aspect ratio"`
	Quality     string `json:"quality" help:"required; quality preset"`
	CallbackURL string `json:"callback_url,omitempty" help:"optional; webhook URL"`
}

// EditImageParams holds the inputs for an image editing request.
// SourceImageURLs accepts 1 to 16 source image URLs to edit from.
// Both AspectRatio and Quality are required fields.
type EditImageParams struct {
	Model           string   `json:"model" help:"required; model slug"`
	Prompt          string   `json:"prompt" help:"required; text description of the desired edit"`
	SourceImageURLs []string `json:"source_image_urls" help:"required; 1-16 source image URLs"`
	AspectRatio     string   `json:"aspect_ratio" help:"required; output aspect ratio"`
	Quality         string   `json:"quality" help:"required; quality preset"`
	CallbackURL     string   `json:"callback_url,omitempty" help:"optional; webhook URL"`
}

// AsyncTaskResponse is the base response for asynchronous generation tasks.
// It implements [core.TaskResponse] so it can be used with the SDK polling helpers.
type AsyncTaskResponse struct {
	ID     string     `json:"id"`
	Status TaskStatus `json:"status"`
	Error  string     `json:"error,omitempty"`
}

// GetID returns the task identifier assigned by the server.
func (r AsyncTaskResponse) GetID() string { return r.ID }

// GetStatus returns the current lifecycle state as a string.
func (r AsyncTaskResponse) GetStatus() string { return string(r.Status) }

// GetError returns the error message when the task has failed, or empty string on success.
func (r AsyncTaskResponse) GetError() string { return r.Error }

// Image holds a single generated image URL returned by the server.
type Image struct {
	URL string `json:"url"`
}

// TextToImageResponse is the generation result for a text-to-image request.
// Images is populated once the task reaches the completed state.
type TextToImageResponse struct {
	AsyncTaskResponse
	Images []Image `json:"images,omitempty"`
}

// EditImageResponse is an alias for [TextToImageResponse] because image edits
// return the same response shape as generations.
type EditImageResponse = TextToImageResponse
