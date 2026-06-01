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
    #     model: "gpt-image-1.5", prompt: "A futuristic cityscape"
    #   )
    #
    #   # Edit image
    #   edited = client.edit_image.run(
    #     model: "gpt-image-1.5",
    #     prompt: "Transform into oil painting",
    #     source_image_urls: ["https://cdn.runapi.ai/public/samples/photo.jpg"]
    #   )
    class Client
      # @return [Resources::TextToImage] Text-to-image generation operations.
      attr_reader :text_to_image
      # @return [Resources::EditImage] Image edit operations.
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
