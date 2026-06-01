# frozen_string_literal: true

module RunApi
  module GptImage
    module Types
      MODELS = %w[gpt-image-1.5].freeze
      GENERATION_MODELS = MODELS
      EDIT_MODELS = MODELS

      ASPECT_RATIOS = %w[1:1 2:3 3:2].freeze
      QUALITY_VALUES = %w[medium high].freeze

      class Image < RunApi::Core::BaseModel
        optional :url, String
      end

      class TextToImageResponse < RunApi::Core::TaskResponse
        required :id, String
        optional :status, String, enum: -> { RunApi::Core::TaskResponse::Status::ALL }
        optional :images, [-> { Image }]
        optional :error, String
      end

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
