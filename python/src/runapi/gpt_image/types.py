"""GPT Image model lists, enums, and response models."""

from __future__ import annotations

from runapi.core import BaseModel, TaskResponse, optional, required


class Image(BaseModel):
    url = optional(str)


class TextToImageResponse(TaskResponse):
    """Task status/result for GPT Image text-to-image."""
    id = required(str)
    status = optional(str, enum=lambda: TaskResponse.Status.ALL)
    images = optional([lambda: Image])
    error = optional(str)


EditImageResponse = TextToImageResponse


class CompletedTextToImageResponse(TextToImageResponse):
    """Narrowed response from ``run()`` once polling observes completion."""

    images = required([lambda: Image])


CompletedEditImageResponse = CompletedTextToImageResponse
