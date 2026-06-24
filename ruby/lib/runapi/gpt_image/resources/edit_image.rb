# frozen_string_literal: true

module RunApi
  module GptImage
    module Resources
      # GPT Image 1.5 image edit resource.
      class EditImage
        include RunApi::Core::ResourceHelpers

        ENDPOINT = "/api/v1/gpt_image/edit_image"

        RESPONSE_CLASS = Types::EditImageResponse
        COMPLETED_RESPONSE_CLASS = Types::CompletedEditImageResponse

        def initialize(http)
          @http = http
        end

        # Edit an image and wait until complete.
        #
        # @param params [Hash] edit parameters
        # @return [RunApi::GptImage::Types::CompletedEditImageResponse] completed edit with images
        def run(**params)
          task = create(**params)
          poll_until_complete { get(task.id) }
        end

        # Create an image editing task.
        #
        # @param params [Hash] edit parameters
        # @return [RunApi::GptImage::Types::EditImageResponse] task creation result with id
        def create(**params)
          params = compact_params(params)
          validate_contract!(CONTRACT["edit-image"], params)
          request(:post, ENDPOINT, body: params)
        end

        # Get edit status by task ID.
        #
        # @param id [String] task ID
        # @return [RunApi::GptImage::Types::EditImageResponse] current edit status
        def get(id)
          request(:get, "#{ENDPOINT}/#{id}")
        end
      end
    end
  end
end
