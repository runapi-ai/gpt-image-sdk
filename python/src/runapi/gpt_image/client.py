"""GPT Image client."""

from __future__ import annotations

from typing import Any, Optional

from runapi.core import ProviderClient

from .resources.edit_image import EditImage
from .resources.text_to_image import TextToImage


class GptImageClient(ProviderClient):
    """GPT Image text-to-image and edit-image client.

    Example::

        client = GptImageClient(api_key="sk-...")
        result = client.text_to_image.run(
            model="gpt-image-1.5", prompt="A futuristic cityscape",
            aspect_ratio="1:1", quality="high",
        )
    """

    def __init__(self, api_key: Optional[str] = None, **options: Any) -> None:
        super().__init__(api_key, **options)
        http = self._http
        self.text_to_image = TextToImage(http)
        self.edit_image = EditImage(http)
