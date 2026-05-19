package gptimage

type TaskStatus string

type TextToImageParams struct {
	Model       string `json:"model" help:"required; gpt-image/1.5-text-to-image"`
	Prompt      string `json:"prompt" help:"required; text description of the desired image"`
	AspectRatio string `json:"aspect_ratio" help:"required; 1:1, 2:3, or 3:2"`
	Quality     string `json:"quality" help:"required; medium or high"`
	CallbackURL string `json:"callback_url,omitempty" help:"optional; webhook URL"`
}

type EditImageParams struct {
	Model       string   `json:"model" help:"required; gpt-image/1.5-image-to-image"`
	Prompt      string   `json:"prompt" help:"required; text description of the desired edit"`
	InputURLs   []string `json:"input_urls" help:"required; 1-16 input image URLs"`
	AspectRatio string   `json:"aspect_ratio" help:"required; 1:1, 2:3, or 3:2"`
	Quality     string   `json:"quality" help:"required; medium or high"`
	CallbackURL string   `json:"callback_url,omitempty" help:"optional; webhook URL"`
}

type AsyncTaskResponse struct {
	ID     string     `json:"id"`
	Status TaskStatus `json:"status"`
	Error  string     `json:"error,omitempty"`
}

func (r AsyncTaskResponse) GetID() string     { return r.ID }
func (r AsyncTaskResponse) GetStatus() string { return string(r.Status) }
func (r AsyncTaskResponse) GetError() string  { return r.Error }

type Image struct {
	URL string `json:"url"`
}

type TextToImageResponse struct {
	AsyncTaskResponse
	Images []Image `json:"images,omitempty"`
}

type EditImageResponse = TextToImageResponse
