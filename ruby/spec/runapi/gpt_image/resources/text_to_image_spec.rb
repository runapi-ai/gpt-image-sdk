# frozen_string_literal: true

require "spec_helper"

RSpec.describe RunApi::GptImage::Resources::TextToImage do
  let(:http) { instance_double(RunApi::Core::HttpClient) }
  let(:text_to_image) { described_class.new(http) }
  let(:endpoint) { "/api/v1/gpt_image/text_to_image" }

  describe "#create" do
    it "POSTs to the correct endpoint with required params" do
      params = {model: "gpt-image-1.5", prompt: "a futuristic cityscape",
                aspect_ratio: "1:1", quality: "high"}
      expect(http).to receive(:request).with(:post, endpoint, body: params)
        .and_return("id" => "task-1")

      result = text_to_image.create(**params)
      expect(result).to be_a(RunApi::GptImage::Types::TextToImageResponse)
      expect(result.id).to eq("task-1")
    end

    it "includes optional params when provided" do
      params = {model: "gpt-image-1.5", prompt: "portrait", aspect_ratio: "2:3", quality: "high"}
      expect(http).to receive(:request).with(:post, endpoint, body: params)
        .and_return("id" => "task-2")

      result = text_to_image.create(**params)
      expect(result.id).to eq("task-2")
    end

    it "raises ValidationError when model is missing" do
      expect { text_to_image.create(prompt: "test") }
        .to raise_error(RunApi::Core::ValidationError, /model must be one of: gpt-image-1.5/)
    end

    it "raises ValidationError when prompt is missing" do
      expect { text_to_image.create(model: "gpt-image-1.5", aspect_ratio: "1:1", quality: "high") }
        .to raise_error(RunApi::Core::ValidationError, /prompt is required/)
    end

    it "raises ValidationError for invalid model" do
      expect { text_to_image.create(model: "invalid-model", prompt: "test") }
        .to raise_error(RunApi::Core::ValidationError, /model must be one of: gpt-image-1.5/)
    end

    it "raises ValidationError when aspect_ratio is missing" do
      expect { text_to_image.create(model: "gpt-image-1.5", prompt: "test", quality: "high") }
        .to raise_error(RunApi::Core::ValidationError, /aspect_ratio is required/)
    end

    it "raises ValidationError for invalid aspect_ratio" do
      expect { text_to_image.create(model: "gpt-image-1.5", prompt: "test", aspect_ratio: "16:9", quality: "high") }
        .to raise_error(RunApi::Core::ValidationError, /aspect_ratio must be one of: 1:1, 2:3, 3:2/)
    end

    it "raises ValidationError when quality is missing" do
      expect { text_to_image.create(model: "gpt-image-1.5", prompt: "test", aspect_ratio: "1:1") }
        .to raise_error(RunApi::Core::ValidationError, /quality is required/)
    end

    it "raises ValidationError for invalid quality" do
      expect { text_to_image.create(model: "gpt-image-1.5", prompt: "test", aspect_ratio: "1:1", quality: "ultra") }
        .to raise_error(RunApi::Core::ValidationError, /quality must be one of: medium, high/)
    end
  end

  describe "#get" do
    it "GETs the correct endpoint with task id" do
      expect(http).to receive(:request).with(:get, "#{endpoint}/task-123")
        .and_return("id" => "task-123", "status" => "completed",
          "images" => [{"url" => "https://file.runapi.ai/result.png"}])

      result = text_to_image.get("task-123")
      expect(result).to be_a(RunApi::GptImage::Types::TextToImageResponse)
      expect(result.id).to eq("task-123")
      expect(result.status).to eq("completed")
      expect(result.images.first.url).to eq("https://file.runapi.ai/result.png")
    end
  end
end
