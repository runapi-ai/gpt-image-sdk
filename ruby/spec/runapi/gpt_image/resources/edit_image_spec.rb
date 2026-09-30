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
