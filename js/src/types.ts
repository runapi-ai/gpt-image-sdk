import type { AsyncTaskStatus } from '@runapi.ai/core';

// Model types
export type GptImageGenerationModel = 'gpt-image/1.5-text-to-image';
export type GptImageEditModel = 'gpt-image/1.5-image-to-image';

// Aspect ratio
export type AspectRatio = '1:1' | '2:3' | '3:2';

// Quality
export type Quality = 'medium' | 'high';

// Text-to-image generation params
export interface GenerationParams {
  model: GptImageGenerationModel;
  prompt: string;
  aspect_ratio: AspectRatio;
  quality: Quality;
  callback_url?: string;
}

// Image-to-image edit params
export interface EditParams {
  model: GptImageEditModel;
  prompt: string;
  input_urls: string[];
  aspect_ratio: AspectRatio;
  quality: Quality;
  callback_url?: string;
}

// Response types
export interface TaskCreateResponse {
  id: string;
  status: string;
}

export interface Image {
  url: string;
}

export interface GenerationResponse {
  id: string;
  status: AsyncTaskStatus;
  images?: Image[];
  error?: string;
  [key: string]: unknown;
}

export type EditResponse = GenerationResponse;
