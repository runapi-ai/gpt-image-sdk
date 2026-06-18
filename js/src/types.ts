import type { AsyncTaskStatus } from '@runapi.ai/core';

/** Single model slug for all GPT Image 1.5 operations. */
export type GptImageModel = 'gpt-image-1.5';
export type GptImageTextToImageModel = GptImageModel;
export type GptImageEditImageModel = GptImageModel;

/** Output aspect ratio. Required for all GPT Image 1.5 operations. */
export type AspectRatio = '1:1' | '2:3' | '3:2';

/** Generation quality tier. Higher quality takes longer but produces finer detail. */
export type Quality = 'medium' | 'high';

/**
 * Parameters for GPT Image 1.5 text-to-image generation.
 * Both `aspect_ratio` and `quality` are required.
 */
export interface TextToImageParams {
  model: GptImageTextToImageModel;
  /** Text description of desired image. */
  prompt: string;
  /** Width-to-height ratio. Required. */
  aspect_ratio: AspectRatio;
  /** Generation quality tier. Required. */
  quality: Quality;
  /** URL for completion callback notifications. */
  callback_url?: string;
}

/**
 * Parameters for GPT Image 1.5 editing.
 * Applies prompt-guided edits to source images. Both `aspect_ratio` and
 * `quality` are required.
 */
export interface EditImageParams {
  model: GptImageEditImageModel;
  /** Text description of desired edits. */
  prompt: string;
  /** Source image URLs to edit (up to 16). */
  source_image_urls: string[];
  /** Width-to-height ratio. Required. */
  aspect_ratio: AspectRatio;
  /** Generation quality tier. Required. */
  quality: Quality;
  /** URL for completion callback notifications. */
  callback_url?: string;
}

/** Acknowledged task with its server-assigned ID. */
export interface TaskCreateResponse {
  id: string;
  status: string;
}

/** A single generated image with its CDN URL. */
export interface Image {
  url: string;
}

/**
 * Generation result for a GPT Image 1.5 task.
 * `images` is populated once `status` reaches `'completed'`.
 */
export interface TextToImageResponse {
  id: string;
  status: AsyncTaskStatus;
  images?: Image[];
  error?: string;
  [key: string]: unknown;
}

/** Edit response -- same shape as generation. */
export type EditImageResponse = TextToImageResponse;
