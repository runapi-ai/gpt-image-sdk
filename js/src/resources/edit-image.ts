import type { HttpClient, RequestOptions, PollingOptions } from '@runapi.ai/core';
import { compactParams } from '@runapi.ai/core';
import { pollUntilComplete } from '@runapi.ai/core/internal';
import type {
  EditImageParams,
  EditImageResponse,
  TaskCreateResponse,
} from '../types';

const ENDPOINT = '/api/v1/gpt_image/edit_image';

/** GPT Image 1.5 prompt-guided image editing resource. */
export class EditImage {
  constructor(private readonly http: HttpClient) {}

  /**
   * Edit an image and wait until complete.
   * @param params Edit parameters.
   * @param options Per-request and polling overrides.
   * @returns The completed edit with images.
   */
  async run(params: EditImageParams, options?: RequestOptions & PollingOptions): Promise<EditImageResponse> {
    const { id } = await this.create(params, options);
    return pollUntilComplete<EditImageResponse>(() => this.get(id, options), {
      maxWaitMs: options?.maxWaitMs,
      pollIntervalMs: options?.pollIntervalMs,
    });
  }

  /**
   * Create an image editing task; returns immediately with a task id.
   * @param params Edit parameters.
   * @param options Per-request overrides.
   * @returns The task creation result with id.
   */
  async create(params: EditImageParams, options?: RequestOptions): Promise<TaskCreateResponse> {
    return this.http.request<TaskCreateResponse>('POST', ENDPOINT, {
      body: compactParams(params),
      ...options,
    });
  }

  /**
   * Fetch the current status of an image editing task.
   * @param id The task id.
   * @param options Per-request overrides.
   * @returns The current edit status.
   */
  async get(id: string, options?: RequestOptions): Promise<EditImageResponse> {
    return this.http.request<EditImageResponse>('GET', `${ENDPOINT}/${id}`, {
      ...options,
    });
  }
}
