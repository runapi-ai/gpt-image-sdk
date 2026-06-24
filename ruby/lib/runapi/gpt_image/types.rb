# frozen_string_literal: true

module RunApi
  module GptImage
    # Type definitions and constants for GPT Image 1.5.
    # Both aspect_ratio and quality are required for all operations.
    module Types
      # A single generated image with its CDN URL.
      class Image < RunApi::Core::BaseModel
        optional :url, String
      end

      # Generation result. +images+ is populated once +status+ is +"completed"+.
      class TextToImageResponse < RunApi::Core::TaskResponse
        required :id, String
        optional :status, String, enum: -> { RunApi::Core::TaskResponse::Status::ALL }
        optional :images, [-> { Image }]
        optional :error, String
      end

      # Edit response -- same shape as generation.
      EditImageResponse = TextToImageResponse

      # Narrowed response returned by `run()` methods once polling observes
      # `status: "completed"`. `images` is required so consumers never have to
      # null-check it on a successful task.
      class CompletedTextToImageResponse < TextToImageResponse
        required :images, [-> { Image }]
      end

      CompletedEditImageResponse = CompletedTextToImageResponse
    end
  end
end
