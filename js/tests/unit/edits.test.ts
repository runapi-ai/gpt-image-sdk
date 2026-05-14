import { describe, it, expect, vi, beforeEach } from 'vitest';
import { Edits } from '../../src/resources/edits';
import type { HttpClient } from '@runapi.ai/core';
import type { EditResponse, TaskCreateResponse } from '../../src/types';

describe('Edits', () => {
  const mockHttp: HttpClient = {
    request: vi.fn(),
  };

  beforeEach(() => {
    vi.clearAllMocks();
  });

  describe('create', () => {
    it('should send correct request for image-to-image edit', async () => {
      const mockResponse: TaskCreateResponse = { id: 'task-edit-123', status: 'processing' };
      vi.mocked(mockHttp.request).mockResolvedValueOnce(mockResponse);

      const edits = new Edits(mockHttp);
      const result = await edits.create({
        model: 'gpt-image/1.5-image-to-image',
        prompt: 'Transform into oil painting',
        input_urls: ['https://example.com/photo.jpg'],
        aspect_ratio: '3:2',
        quality: 'high',
      });

      expect(mockHttp.request).toHaveBeenCalledWith(
        'POST',
        '/api/v1/gpt_image/edits',
        {
          body: {
            model: 'gpt-image/1.5-image-to-image',
            prompt: 'Transform into oil painting',
            input_urls: ['https://example.com/photo.jpg'],
            aspect_ratio: '3:2',
            quality: 'high',
          },
        }
      );
      expect(result).toEqual(mockResponse);
    });

    it('should send multiple input_urls', async () => {
      const mockResponse: TaskCreateResponse = { id: 'task-edit-456', status: 'processing' };
      vi.mocked(mockHttp.request).mockResolvedValueOnce(mockResponse);

      const edits = new Edits(mockHttp);
      await edits.create({
        model: 'gpt-image/1.5-image-to-image',
        prompt: 'Combine these images',
        input_urls: ['https://example.com/photo1.jpg', 'https://example.com/photo2.jpg'],
      });

      const call = vi.mocked(mockHttp.request).mock.calls[0];
      const body = (call[2] as any).body;
      expect(body.input_urls).toHaveLength(2);
    });
  });

  describe('get', () => {
    it('should send correct GET request', async () => {
      const mockResponse: EditResponse = {
        id: 'task-edit-789',
        status: 'completed',
        images: [{ url: 'https://file.runapi.ai/edited.png' }],
      };
      vi.mocked(mockHttp.request).mockResolvedValueOnce(mockResponse);

      const edits = new Edits(mockHttp);
      const result = await edits.get('task-edit-789');

      expect(mockHttp.request).toHaveBeenCalledWith(
        'GET',
        '/api/v1/gpt_image/edits/task-edit-789',
        {}
      );
      expect(result).toEqual(mockResponse);
    });
  });
});
