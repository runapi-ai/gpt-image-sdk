# frozen_string_literal: true

module RunApi
  module GptImage
    module Resources
      # GPT Image 1.5 text-to-image generation resource.
      class TextToImage
        include RunApi::Core::ResourceHelpers

        ENDPOINT = "/api/v1/gpt_image/text_to_image"

        RESPONSE_CLASS = Types::TextToImageResponse
        COMPLETED_RESPONSE_CLASS = Types::CompletedTextToImageResponse

        def initialize(http)
          @http = http
        end

        # Generate an image and wait until complete.
        #
        # @param params [Hash] generation parameters
        # @return [RunApi::GptImage::Types::CompletedTextToImageResponse] completed generation with images
        def run(options: nil, **params)
          task = create(options: options, **params)
          poll_until_complete { get(task.id, options: options) }
        end

        # Create a text-to-image generation task.
        #
        # @param params [Hash] generation parameters
        # @return [RunApi::GptImage::Types::TextToImageResponse] task creation result with id
        def create(options: nil, **params)
          params = compact_params(params)
          validate_contract!(CONTRACT["text-to-image"], params)
          request(:post, ENDPOINT, body: params, options: options)
        end

        # Get generation status by task ID.
        #
        # @param id [String] task ID
        # @return [RunApi::GptImage::Types::TextToImageResponse] current generation status
        def get(id, options: nil)
          request(:get, "#{ENDPOINT}/#{id}", options: options)
        end
      end
    end
  end
end
