import { createHttpClient, type ClientOptions } from '@runapi.ai/core';
import { Generations } from './resources/generations';
import { Edits } from './resources/edits';

/**
 * GPT Image 1.5 image generation API client.
 *
 * @example
 * ```typescript
 * const client = new GptImageClient({
 *   apiKey: 'your-api-key',
 *   baseUrl: 'https://runapi.ai',
 * });
 *
 * // Text-to-image
 * const result = await client.generations.run({
 *   model: 'gpt-image/1.5-text-to-image',
 *   prompt: 'A futuristic cityscape at night',
 * });
 *
 * // Image-to-image
 * const edited = await client.edits.run({
 *   model: 'gpt-image/1.5-image-to-image',
 *   prompt: 'Transform into oil painting style',
 *   input_urls: ['https://example.com/photo.jpg'],
 * });
 * ```
 */
export class GptImageClient {
  /** Text-to-image generation operations. */
  public readonly generations: Generations;
  /** Image-to-image edit operations. */
  public readonly edits: Edits;

  constructor(options: ClientOptions = {}) {
    const http = createHttpClient(options);
    this.generations = new Generations(http);
    this.edits = new Edits(http);
  }
}
