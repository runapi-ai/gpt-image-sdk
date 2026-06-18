import type { HttpClient, RequestOptions, PollingOptions } from '@runapi.ai/core';
import { compactParams } from '@runapi.ai/core';
import { pollUntilComplete } from '@runapi.ai/core/internal';
import type {
  TextToImageParams,
  TextToImageResponse,
  TaskCreateResponse,
} from '../types';

const ENDPOINT = '/api/v1/gpt_image/text_to_image';

/** GPT Image 1.5 text-to-image generation resource. */
export class TextToImage {
  constructor(private readonly http: HttpClient) {}

  /**
   * Generate an image and wait until complete.
   * @param params Generation parameters.
   * @param options Per-request and polling overrides.
   * @returns The completed generation with images.
   */
  async run(params: TextToImageParams, options?: RequestOptions & PollingOptions): Promise<TextToImageResponse> {
    const { id } = await this.create(params, options);
    return pollUntilComplete<TextToImageResponse>(() => this.get(id, options), {
      maxWaitMs: options?.maxWaitMs,
      pollIntervalMs: options?.pollIntervalMs,
    });
  }

  /**
   * Create a text-to-image task; returns immediately with a task id.
   * @param params Generation parameters.
   * @param options Per-request overrides.
   * @returns The task creation result with id.
   */
  async create(params: TextToImageParams, options?: RequestOptions): Promise<TaskCreateResponse> {
    return this.http.request<TaskCreateResponse>('POST', ENDPOINT, {
      body: compactParams(params),
      ...options,
    });
  }

  /**
   * Fetch the current status of a text-to-image task.
   * @param id The task id.
   * @param options Per-request overrides.
   * @returns The current generation status.
   */
  async get(id: string, options?: RequestOptions): Promise<TextToImageResponse> {
    return this.http.request<TextToImageResponse>('GET', `${ENDPOINT}/${id}`, {
      ...options,
    });
  }
}
