"""GPT Image edit-image resource."""

from __future__ import annotations

from typing import Any

from runapi.core import Resource

from ..contract_gen import CONTRACT
from ..types import (
    CompletedEditImageResponse,
    EditImageResponse,
)


class EditImage(Resource):
    """Edit images from a prompt and source images with GPT Image models."""

    ENDPOINT = "/api/v1/gpt_image/edit_image"

    RESPONSE_CLASS = EditImageResponse
    COMPLETED_RESPONSE_CLASS = CompletedEditImageResponse

    def run(self, **params: Any) -> Any:
        """Create an edit-image task and poll until it completes.

        Args:
            **params: Edit-image parameters (model, prompt, ...).

        Returns:
            The completed edit-image response.
        """
        task = self.create(**params)
        return self._poll_until_complete(lambda: self.get(task.id))

    def create(self, **params: Any) -> Any:
        """Create an edit-image task and return immediately with an id.

        Args:
            **params: Edit-image parameters (model, prompt, ...).

        Returns:
            The task creation result with an id.
        """
        compacted = self._compact_params(params)
        self._validate_contract(CONTRACT["edit-image"], compacted)
        return self._request("post", self.ENDPOINT, body=compacted)

    def get(self, id: str) -> Any:
        """Fetch the current status of an edit-image task.

        Args:
            id: Task id.

        Returns:
            The current edit-image status.
        """
        return self._request("get", f"{self.ENDPOINT}/{id}")
