import type { AsyncTaskStatus } from '@runapi.ai/core';

// Model types
export type GptImageTextToImageModel = 'gpt-image-1.5-text-to-image';
export type GptImageEditImageModel = 'gpt-image-1.5-image-to-image';

// Aspect ratio
export type AspectRatio = '1:1' | '2:3' | '3:2';

// Quality
export type Quality = 'medium' | 'high';

// Text-to-image generation params
export interface TextToImageParams {
  model: GptImageTextToImageModel;
  prompt: string;
  aspect_ratio: AspectRatio;
  quality: Quality;
  callback_url?: string;
}

// Image-to-image edit params
export interface EditImageParams {
  model: GptImageEditImageModel;
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

export interface TextToImageResponse {
  id: string;
  status: AsyncTaskStatus;
  images?: Image[];
  error?: string;
  [key: string]: unknown;
}

export type EditImageResponse = TextToImageResponse;
