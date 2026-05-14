import type { HttpClient, RequestOptions, PollingOptions } from '@runapi.ai/core';
import { compactParams } from '@runapi.ai/core';
import { pollUntilComplete } from '@runapi.ai/core/internal';
import type {
  EditParams,
  EditResponse,
  TaskCreateResponse,
} from '../types';

const ENDPOINT = '/api/v1/gpt_image/edits';

export class Edits {
  constructor(private readonly http: HttpClient) {}

  async run(params: EditParams, options?: RequestOptions & PollingOptions): Promise<EditResponse> {
    const { id } = await this.create(params, options);
    return pollUntilComplete<EditResponse>(() => this.get(id, options), {
      maxWaitMs: options?.maxWaitMs,
      pollIntervalMs: options?.pollIntervalMs,
    });
  }

  async create(params: EditParams, options?: RequestOptions): Promise<TaskCreateResponse> {
    return this.http.request<TaskCreateResponse>('POST', ENDPOINT, {
      body: compactParams(params),
      ...options,
    });
  }

  async get(id: string, options?: RequestOptions): Promise<EditResponse> {
    return this.http.request<EditResponse>('GET', `${ENDPOINT}/${id}`, {
      ...options,
    });
  }
}
