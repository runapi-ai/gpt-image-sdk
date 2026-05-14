import { describe, it, expect, vi, beforeEach } from 'vitest';
import { Generations } from '../../src/resources/generations';
import type { HttpClient } from '@runapi.ai/core';
import type { GenerationResponse, TaskCreateResponse } from '../../src/types';

describe('Generations', () => {
  const mockHttp: HttpClient = {
    request: vi.fn(),
  };

  beforeEach(() => {
    vi.clearAllMocks();
  });

  describe('create', () => {
    it('should send correct request for text-to-image', async () => {
      const mockResponse: TaskCreateResponse = { id: 'task-gen-123', status: 'processing' };
      vi.mocked(mockHttp.request).mockResolvedValueOnce(mockResponse);

      const generations = new Generations(mockHttp);
      const result = await generations.create({
        model: 'gpt-image/1.5-text-to-image',
        prompt: 'A beautiful landscape',
        aspect_ratio: '1:1',
        quality: 'high',
      });

      expect(mockHttp.request).toHaveBeenCalledWith(
        'POST',
        '/api/v1/gpt_image/generations',
        {
          body: {
            model: 'gpt-image/1.5-text-to-image',
            prompt: 'A beautiful landscape',
            aspect_ratio: '1:1',
            quality: 'high',
          },
        }
      );
      expect(result).toEqual(mockResponse);
    });

    it('should send required params only (no callback_url)', async () => {
      const mockResponse: TaskCreateResponse = { id: 'task-gen-456', status: 'processing' };
      vi.mocked(mockHttp.request).mockResolvedValueOnce(mockResponse);

      const generations = new Generations(mockHttp);
      await generations.create({
        model: 'gpt-image/1.5-text-to-image',
        prompt: 'Abstract art',
        aspect_ratio: '1:1',
        quality: 'medium',
      });

      expect(mockHttp.request).toHaveBeenCalledWith(
        'POST',
        '/api/v1/gpt_image/generations',
        {
          body: {
            model: 'gpt-image/1.5-text-to-image',
            prompt: 'Abstract art',
            aspect_ratio: '1:1',
            quality: 'medium',
          },
        }
      );
    });
  });

  describe('get', () => {
    it('should send correct GET request', async () => {
      const mockResponse: GenerationResponse = {
        id: 'task-gen-789',
        status: 'completed',
        images: [{ url: 'https://file.runapi.ai/result.png' }],
      };
      vi.mocked(mockHttp.request).mockResolvedValueOnce(mockResponse);

      const generations = new Generations(mockHttp);
      const result = await generations.get('task-gen-789');

      expect(mockHttp.request).toHaveBeenCalledWith(
        'GET',
        '/api/v1/gpt_image/generations/task-gen-789',
        {}
      );
      expect(result).toEqual(mockResponse);
    });
  });
});
