# frozen_string_literal: true

module RunApi
  module GptImage
    module Types
      GENERATION_MODELS = %w[gpt-image/1.5-text-to-image].freeze
      EDIT_MODELS = %w[gpt-image/1.5-image-to-image].freeze
      MODELS = (GENERATION_MODELS + EDIT_MODELS).freeze

      ASPECT_RATIOS = %w[1:1 2:3 3:2].freeze
      QUALITY_VALUES = %w[medium high].freeze

      class Image < RunApi::Core::BaseModel
        optional :url, String
      end

      class GenerationResponse < RunApi::Core::TaskResponse
        required :id, String
        optional :status, String, enum: -> { RunApi::Core::TaskResponse::Status::ALL }
        optional :images, [ -> { Image } ]
        optional :error, String
      end

      EditResponse = GenerationResponse

      # Narrowed response returned by `run()` methods once polling observes
      # `status: "completed"`. `images` is required so consumers never have to
      # null-check it on a successful task.
      class CompletedGenerationResponse < GenerationResponse
        required :images, [ -> { Image } ]
      end

      CompletedEditResponse = CompletedGenerationResponse
    end
  end
end
