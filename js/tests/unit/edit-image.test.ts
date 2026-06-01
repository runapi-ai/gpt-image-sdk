import { describe, it, expect, vi, beforeEach } from 'vitest';
import { EditImage } from '../../src/resources/edit-image';
import type { HttpClient } from '@runapi.ai/core';
import type { EditImageResponse, TaskCreateResponse } from '../../src/types';

describe('EditImage', () => {
  const mockHttp: HttpClient = {
    request: vi.fn(),
  };

  beforeEach(() => {
    vi.clearAllMocks();
  });

  describe('create', () => {
    it('should send correct request for image editing', async () => {
      const mockResponse: TaskCreateResponse = { id: 'task-edit-123', status: 'processing' };
      vi.mocked(mockHttp.request).mockResolvedValueOnce(mockResponse);

      const editImage = new EditImage(mockHttp);
      const result = await editImage.create({
        model: 'gpt-image-1.5',
        prompt: 'Transform into oil painting',
        source_image_urls: ['https://cdn.runapi.ai/public/samples/photo.jpg'],
        aspect_ratio: '3:2',
        quality: 'high',
      });

      expect(mockHttp.request).toHaveBeenCalledWith(
        'POST',
        '/api/v1/gpt_image/edit_image',
        {
          body: {
            model: 'gpt-image-1.5',
            prompt: 'Transform into oil painting',
            source_image_urls: ['https://cdn.runapi.ai/public/samples/photo.jpg'],
            aspect_ratio: '3:2',
            quality: 'high',
          },
        }
      );
      expect(result).toEqual(mockResponse);
    });

    it('should send multiple source_image_urls', async () => {
      const mockResponse: TaskCreateResponse = { id: 'task-edit-456', status: 'processing' };
      vi.mocked(mockHttp.request).mockResolvedValueOnce(mockResponse);

      const editImage = new EditImage(mockHttp);
      await editImage.create({
        model: 'gpt-image-1.5',
        prompt: 'Combine these images',
        source_image_urls: ['https://cdn.runapi.ai/public/samples/photo-1.jpg', 'https://cdn.runapi.ai/public/samples/photo-2.jpg'],
      });

      const call = vi.mocked(mockHttp.request).mock.calls[0];
      const body = (call[2] as any).body;
      expect(body.source_image_urls).toHaveLength(2);
    });
  });

  describe('get', () => {
    it('should send correct GET request', async () => {
      const mockResponse: EditImageResponse = {
        id: 'task-edit-789',
        status: 'completed',
        images: [{ url: 'https://file.runapi.ai/edited.png' }],
      };
      vi.mocked(mockHttp.request).mockResolvedValueOnce(mockResponse);

      const editImage = new EditImage(mockHttp);
      const result = await editImage.get('task-edit-789');

      expect(mockHttp.request).toHaveBeenCalledWith(
        'GET',
        '/api/v1/gpt_image/edit_image/task-edit-789',
        {}
      );
      expect(result).toEqual(mockResponse);
    });
  });
});
