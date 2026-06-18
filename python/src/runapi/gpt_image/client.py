"""GPT Image client."""

from __future__ import annotations

from typing import Any, Optional

from runapi.core import ClientOptions, HttpClient, resolve_api_key

from .resources.edit_image import EditImage
from .resources.text_to_image import TextToImage


class GptImageClient:
    """GPT Image text-to-image and edit-image client.

    Example::

        client = GptImageClient(api_key="sk-...")
        result = client.text_to_image.run(
            model="gpt-image-1.5", prompt="A futuristic cityscape",
            aspect_ratio="1:1", quality="high",
        )
    """

    def __init__(self, api_key: Optional[str] = None, **options: Any) -> None:
        resolved_api_key = resolve_api_key(api_key)
        client_options = ClientOptions(api_key=resolved_api_key, **options)
        http = client_options.http_client or HttpClient(client_options)
        self.text_to_image = TextToImage(http)
        self.edit_image = EditImage(http)
