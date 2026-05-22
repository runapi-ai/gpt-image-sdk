# frozen_string_literal: true

module RunApi
  module GptImage
    # GPT Image 1.5 image generation API client.
    #
    # @example
    #   client = RunApi::GptImage::Client.new(api_key: "your-api-key")
    #
    #   # Text-to-image
    #   result = client.text_to_image.run(
    #     model: "gpt-image-1.5-text-to-image", prompt: "A futuristic cityscape"
    #   )
    #
    #   # Image-to-image
    #   edited = client.edit_image.run(
    #     model: "gpt-image-1.5-image-to-image",
    #     prompt: "Transform into oil painting",
    #     input_urls: ["https://example.com/photo.jpg"]
    #   )
    class Client
      # @return [Resources::TextToImage] Text-to-image generation operations.
      attr_reader :text_to_image
      # @return [Resources::EditImage] Image-to-image edit operations.
      attr_reader :edit_image

      def initialize(api_key: nil, **options)
        @api_key = Core::Auth.resolve_api_key(api_key)

        client_options = Core::ClientOptions.new(api_key: @api_key, **options)
        http = client_options.http_client || Core::HttpClient.new(client_options)
        @text_to_image = Resources::TextToImage.new(http)
        @edit_image = Resources::EditImage.new(http)
      end
    end
  end
end
