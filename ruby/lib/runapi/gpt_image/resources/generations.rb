# frozen_string_literal: true

module RunApi
  module GptImage
    module Resources
      # GPT Image 1.5 text-to-image generation resource.
      class Generations
        include RunApi::Core::ResourceHelpers

        ENDPOINT = "/api/v1/gpt_image/generations"

        RESPONSE_CLASS = Types::GenerationResponse
        COMPLETED_RESPONSE_CLASS = Types::CompletedGenerationResponse

        def initialize(http)
          @http = http
        end

        # Generate an image and wait until complete.
        #
        # @param params [Hash] generation parameters
        # @return [RunApi::GptImage::Types::CompletedGenerationResponse] completed generation with images
        def run(**params)
          task = create(**params)
          poll_until_complete { get(task.id) }
        end

        # Create a text-to-image generation task.
        #
        # @param params [Hash] generation parameters
        # @return [RunApi::GptImage::Types::GenerationResponse] task creation result with id
        def create(**params)
          params = compact_params(params)
          validate_params!(params)
          request(:post, ENDPOINT, body: params)
        end

        # Get generation status by task ID.
        #
        # @param id [String] task ID
        # @return [RunApi::GptImage::Types::GenerationResponse] current generation status
        def get(id)
          request(:get, "#{ENDPOINT}/#{id}")
        end

        private

        def validate_params!(params)
          raise Core::ValidationError, "model is required" unless param(params, :model)
          raise Core::ValidationError, "prompt is required" unless param(params, :prompt)

          model = param(params, :model)
          unless Types::GENERATION_MODELS.include?(model)
            raise Core::ValidationError, "Invalid model: #{model}. Must be: #{Types::GENERATION_MODELS.join(", ")}"
          end

          raise Core::ValidationError, "aspect_ratio is required" unless param(params, :aspect_ratio)
          validate_optional!(params, :aspect_ratio, Types::ASPECT_RATIOS)
          raise Core::ValidationError, "quality is required" unless param(params, :quality)
          validate_optional!(params, :quality, Types::QUALITY_VALUES)
        end
      end
    end
  end
end
