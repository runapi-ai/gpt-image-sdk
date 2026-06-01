package gptimage

type TaskStatus string

type TextToImageParams struct {
	Model       string `json:"model" help:"required; model slug"`
	Prompt      string `json:"prompt" help:"required; text description of the desired image"`
	AspectRatio string `json:"aspect_ratio" help:"required; output aspect ratio"`
	Quality     string `json:"quality" help:"required; quality preset"`
	CallbackURL string `json:"callback_url,omitempty" help:"optional; webhook URL"`
}

type EditImageParams struct {
	Model           string   `json:"model" help:"required; model slug"`
	Prompt          string   `json:"prompt" help:"required; text description of the desired edit"`
	SourceImageURLs []string `json:"source_image_urls" help:"required; 1-16 source image URLs"`
	AspectRatio     string   `json:"aspect_ratio" help:"required; output aspect ratio"`
	Quality         string   `json:"quality" help:"required; quality preset"`
	CallbackURL     string   `json:"callback_url,omitempty" help:"optional; webhook URL"`
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
