import pytest

from runapi.core import config
from runapi.core.errors import AuthenticationError, ValidationError
from runapi.gpt_image import GptImageClient
from runapi.gpt_image.resources.edit_image import EditImage
from runapi.gpt_image.resources.text_to_image import TextToImage
from runapi.gpt_image.types import CompletedTextToImageResponse, TextToImageResponse


class FakeHttp:
    def __init__(self, *responses):
        self._responses = list(responses)
        self.calls = []

    def request(self, method, path, body=None, options=None):
        self.calls.append((method, path, body))
        if self._responses:
            return self._responses.pop(0)
        return {"id": "task_1", "status": "pending"}


@pytest.fixture(autouse=True)
def reset_config(monkeypatch):
    monkeypatch.delenv("RUNAPI_API_KEY", raising=False)
    monkeypatch.setattr(config, "api_key", None)
    yield


# --- auth -----------------------------------------------------------------


def test_accepts_api_key_parameter():
    assert isinstance(GptImageClient(api_key="k", http_client=FakeHttp()), GptImageClient)


def test_falls_back_to_global(monkeypatch):
    monkeypatch.setattr(config, "api_key", "global-key")
    assert isinstance(GptImageClient(http_client=FakeHttp()), GptImageClient)


def test_falls_back_to_env(monkeypatch):
    monkeypatch.setenv("RUNAPI_API_KEY", "env-key")
    assert isinstance(GptImageClient(http_client=FakeHttp()), GptImageClient)


def test_raises_without_api_key():
    with pytest.raises(AuthenticationError, match="API key is required"):
        GptImageClient()


# --- injection / accessors ------------------------------------------------


def test_uses_injected_http_client():
    fake = FakeHttp()
    client = GptImageClient(api_key="k", http_client=fake)
    assert client.text_to_image._http is fake
    assert client.edit_image._http is fake


def test_exposes_resource_accessors():
    client = GptImageClient(api_key="k", http_client=FakeHttp())
    assert isinstance(client.text_to_image, TextToImage)
    assert isinstance(client.edit_image, EditImage)


# --- request shapes -------------------------------------------------------


def test_text_to_image_create_posts_compacted_body():
    fake = FakeHttp({"id": "t1", "status": "pending"})
    client = GptImageClient(api_key="k", http_client=fake)
    result = client.text_to_image.create(
        model="gpt-image-1.5",
        prompt="hello world",
        aspect_ratio="1:1",
        quality="high",
        seed=None,
    )
    assert fake.calls == [
        (
            "post",
            "/api/v1/gpt_image/text_to_image",
            {"model": "gpt-image-1.5", "prompt": "hello world", "aspect_ratio": "1:1", "quality": "high"},
        ),
    ]
    assert isinstance(result, TextToImageResponse)


def test_text_to_image_get_fetches_by_id():
    fake = FakeHttp({"id": "t1", "status": "processing"})
    client = GptImageClient(api_key="k", http_client=fake)
    client.text_to_image.get("t1")
    assert fake.calls == [("get", "/api/v1/gpt_image/text_to_image/t1", None)]


def test_edit_image_create_posts_compacted_body():
    fake = FakeHttp({"id": "t1", "status": "pending"})
    client = GptImageClient(api_key="k", http_client=fake)
    result = client.edit_image.create(
        model="gpt-image-1.5",
        prompt="oil painting",
        source_image_urls=["https://example.com/a.jpg"],
        aspect_ratio="2:3",
        quality="medium",
    )
    assert fake.calls == [
        (
            "post",
            "/api/v1/gpt_image/edit_image",
            {
                "model": "gpt-image-1.5",
                "prompt": "oil painting",
                "source_image_urls": ["https://example.com/a.jpg"],
                "aspect_ratio": "2:3",
                "quality": "medium",
            },
        ),
    ]
    assert isinstance(result, TextToImageResponse)


def test_edit_image_get_fetches_by_id():
    fake = FakeHttp({"id": "t1", "status": "processing"})
    client = GptImageClient(api_key="k", http_client=fake)
    client.edit_image.get("t1")
    assert fake.calls == [("get", "/api/v1/gpt_image/edit_image/t1", None)]


def test_run_narrows_completed_type():
    fake = FakeHttp(
        {"id": "t1", "status": "pending"},
        {"id": "t1", "status": "completed", "images": [{"url": "https://x/y.png"}]},
    )
    client = GptImageClient(api_key="k", http_client=fake)
    result = client.text_to_image.run(
        model="gpt-image-1.5", prompt="a serene lake", aspect_ratio="1:1", quality="high"
    )
    assert isinstance(result, CompletedTextToImageResponse)
    assert result.images[0].url == "https://x/y.png"


# --- validation -----------------------------------------------------------


def test_requires_model():
    client = GptImageClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="model is required"):
        client.text_to_image.create(prompt="hi there")


def test_requires_prompt():
    client = GptImageClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="prompt is required"):
        client.text_to_image.create(model="gpt-image-1.5")


def test_rejects_unknown_model():
    client = GptImageClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="Invalid model: nope. Must be: gpt-image-1.5"):
        client.text_to_image.create(model="nope", prompt="hi there")


def test_requires_aspect_ratio():
    client = GptImageClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="aspect_ratio is required"):
        client.text_to_image.create(model="gpt-image-1.5", prompt="hi there", quality="high")


def test_rejects_invalid_aspect_ratio():
    client = GptImageClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="Invalid aspect_ratio"):
        client.text_to_image.create(
            model="gpt-image-1.5", prompt="hi there", aspect_ratio="9:16", quality="high"
        )


def test_requires_quality():
    client = GptImageClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="quality is required"):
        client.text_to_image.create(model="gpt-image-1.5", prompt="hi there", aspect_ratio="1:1")


def test_rejects_invalid_quality():
    client = GptImageClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="Invalid quality"):
        client.text_to_image.create(
            model="gpt-image-1.5", prompt="hi there", aspect_ratio="1:1", quality="low"
        )


def test_edit_requires_source_image_urls():
    client = GptImageClient(api_key="k", http_client=FakeHttp())
    with pytest.raises(ValidationError, match="source_image_urls is required for image editing"):
        client.edit_image.create(
            model="gpt-image-1.5", prompt="make it pop", aspect_ratio="1:1", quality="high"
        )
