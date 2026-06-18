# frozen_string_literal: true

module RunApi
  module GptImage
    # GPT Image 1.5 generation and editing API client.
    #
    # Both aspect_ratio and quality are required for all operations.
    #
    # @example
    #   client = RunApi::GptImage::Client.new(api_key: "your-api-key")
    #
    #   # Text-to-image
    #   result = client.text_to_image.run(
    #     model: "gpt-image-1.5", prompt: "A futuristic cityscape"
    #   )
    #
    #   # Edit image
    #   edited = client.edit_image.run(
    #     model: "gpt-image-1.5",
    #     prompt: "Transform into oil painting",
    #     source_image_urls: ["https://cdn.runapi.ai/public/samples/photo.jpg"]
    #   )
    class Client < RunApi::Core::Client
      # @return [Resources::TextToImage] Text-to-image generation operations.
      attr_reader :text_to_image
      # @return [Resources::EditImage] Image edit operations.
      attr_reader :edit_image

      def initialize(api_key: nil, **options)
        super
        @text_to_image = Resources::TextToImage.new(http)
        @edit_image = Resources::EditImage.new(http)
      end
    end
  end
end
