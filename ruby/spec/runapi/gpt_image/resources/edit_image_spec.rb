# frozen_string_literal: true

require "spec_helper"

RSpec.describe RunApi::GptImage::Resources::EditImage do
  let(:http) { instance_double(RunApi::Core::HttpClient) }
  let(:edit_image) { described_class.new(http) }
  let(:endpoint) { "/api/v1/gpt_image/edit_image" }

  describe "#create" do
    it "POSTs to the correct endpoint with required params" do
      params = {model: "gpt-image-1.5", prompt: "transform to oil painting",
                source_image_urls: ["https://cdn.runapi.ai/public/samples/photo.jpg"], aspect_ratio: "3:2", quality: "high"}
      expect(http).to receive(:request).with(:post, endpoint, body: params)
        .and_return("id" => "edit-1")

      result = edit_image.create(**params)
      expect(result).to be_a(RunApi::GptImage::Types::EditImageResponse)
      expect(result.id).to eq("edit-1")
    end

    it "includes optional aspect_ratio and quality when provided" do
      params = {model: "gpt-image-1.5", prompt: "portrait",
                source_image_urls: ["https://cdn.runapi.ai/public/samples/input.jpg"], aspect_ratio: "3:2", quality: "high"}
      expect(http).to receive(:request).with(:post, endpoint, body: params)
        .and_return("id" => "edit-2")

      result = edit_image.create(**params)
      expect(result.id).to eq("edit-2")
    end

    it "raises ValidationError when model is missing" do
      expect { edit_image.create(prompt: "test", source_image_urls: ["https://cdn.runapi.ai/public/samples/input.jpg"]) }
        .to raise_error(RunApi::Core::ValidationError, /model must be one of: gpt-image-1.5/)
    end

    it "raises ValidationError when prompt is missing" do
      expect { edit_image.create(model: "gpt-image-1.5", aspect_ratio: "3:2", quality: "high", source_image_urls: ["https://cdn.runapi.ai/public/samples/input.jpg"]) }
        .to raise_error(RunApi::Core::ValidationError, /prompt is required/)
    end

    it "raises ValidationError for invalid model" do
      expect { edit_image.create(model: "invalid-model", prompt: "test", source_image_urls: ["https://cdn.runapi.ai/public/samples/input.jpg"]) }
        .to raise_error(RunApi::Core::ValidationError, /model must be one of: gpt-image-1.5/)
    end

    it "raises ValidationError when source_image_urls is missing" do
      expect { edit_image.create(model: "gpt-image-1.5", prompt: "test", aspect_ratio: "3:2", quality: "high") }
        .to raise_error(RunApi::Core::ValidationError, /source_image_urls is required/)
    end

    it "raises ValidationError when source_image_urls is empty" do
      expect { edit_image.create(model: "gpt-image-1.5", prompt: "test", aspect_ratio: "3:2", quality: "high", source_image_urls: []) }
        .to raise_error(RunApi::Core::ValidationError, /source_image_urls is required/)
    end

    it "raises ValidationError when aspect_ratio is missing" do
      expect {
        edit_image.create(model: "gpt-image-1.5", prompt: "test",
          source_image_urls: ["https://cdn.runapi.ai/public/samples/input.jpg"], quality: "high")
      }
        .to raise_error(RunApi::Core::ValidationError, /aspect_ratio is required/)
    end

    it "raises ValidationError for invalid aspect_ratio" do
      expect {
        edit_image.create(model: "gpt-image-1.5", prompt: "test",
          source_image_urls: ["https://cdn.runapi.ai/public/samples/input.jpg"], aspect_ratio: "16:9", quality: "high")
      }
        .to raise_error(RunApi::Core::ValidationError, /aspect_ratio must be one of: 1:1, 2:3, 3:2/)
    end

    it "raises ValidationError when quality is missing" do
      expect {
        edit_image.create(model: "gpt-image-1.5", prompt: "test",
          source_image_urls: ["https://cdn.runapi.ai/public/samples/input.jpg"], aspect_ratio: "3:2")
      }
        .to raise_error(RunApi::Core::ValidationError, /quality is required/)
    end
  end

  describe "#get" do
    it "GETs the correct endpoint with task id" do
      expect(http).to receive(:request).with(:get, "#{endpoint}/edit-456")
        .and_return("id" => "edit-456", "status" => "completed",
          "images" => [{"url" => "https://file.runapi.ai/edited.png"}])

      result = edit_image.get("edit-456")
      expect(result).to be_a(RunApi::GptImage::Types::EditImageResponse)
      expect(result.id).to eq("edit-456")
      expect(result.status).to eq("completed")
      expect(result.images.first.url).to eq("https://file.runapi.ai/edited.png")
    end
  end
end
