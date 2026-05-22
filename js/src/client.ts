import { createHttpClient, type ClientOptions } from '@runapi.ai/core';
import { TextToImage } from './resources/text-to-image';
import { EditImage } from './resources/edit-image';

/**
 * GPT Image 1.5 text-to-image API client.
 *
 * @example
 * ```typescript
 * const client = new GptImageClient({
 *   apiKey: 'your-api-key',
 *   baseUrl: 'https://runapi.ai',
 * });
 *
 * // Text-to-image
 * const result = await client.textToImage.run({
 *   model: 'gpt-image-1.5-text-to-image',
 *   prompt: 'A futuristic cityscape at night',
 * });
 *
 * // Image-to-image
 * const edited = await client.editImage.run({
 *   model: 'gpt-image-1.5-image-to-image',
 *   prompt: 'Transform into oil painting style',
 *   input_urls: ['https://example.com/photo.jpg'],
 * });
 * ```
 */
export class GptImageClient {
  /** Text-to-image generation operations. */
  public readonly textToImage: TextToImage;
  /** Image-to-image edit operations. */
  public readonly editImage: EditImage;

  constructor(options: ClientOptions = {}) {
    const http = createHttpClient(options);
    this.textToImage = new TextToImage(http);
    this.editImage = new EditImage(http);
  }
}
