# GPT Image Python SDK for RunAPI

The GPT Image Python SDK is the language-specific package for GPT Image on RunAPI. Use this package for image generation, image editing, and creative production workflows when your application needs request bodies, task status lookup, and consistent RunAPI errors in Python.

This README is the Python package guide inside the public `gpt-image-sdk` repository. For the repository overview, start at `../README.md`; for model details, use https://runapi.ai/models/gpt-image; for API reference, use https://runapi.ai/docs#gpt-image; for SDK docs, use https://runapi.ai/docs#sdk-gpt-image.

## Install

```bash
pip install runapi-gpt-image
```

## Quick start

```python
from runapi.gpt_image import GptImageClient

client = GptImageClient()  # reads RUNAPI_API_KEY, or pass api_key="sk-..."

task = client.text_to_image.create(
    model="gpt-image-1.5",
    prompt="A futuristic cityscape at dusk, cinematic",
    aspect_ratio="1:1",
    quality="high",
)
status = client.text_to_image.get(task.id)

edit = client.edit_image.create(
    model="gpt-image-1.5",
    prompt="Transform into an oil painting",
    source_image_urls=["https://cdn.runapi.ai/public/samples/image.jpg"],
    aspect_ratio="1:1",
    quality="high",
)
```

Use `create` when you want to submit a task and return quickly, `get` when you need the latest task state, and `run` when a script should create and poll until completion:

```python
result = client.text_to_image.run(
    model="gpt-image-1.5",
    prompt="A serene mountain lake at dawn",
    aspect_ratio="3:2",
    quality="high",
)
print(result.images[0].url)
```

In web request handlers, prefer `create` plus webhook or later `get` polling so a worker is not held open.

RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.

## Language notes

Pass parameters as keyword arguments and catch the `runapi.gpt_image` error classes when building image jobs or scripts. The available resources are `text_to_image` and `edit_image`. Keep `RUNAPI_API_KEY` in the environment or your secret manager; never commit API keys or callback secrets.

## Links

- Model page: https://runapi.ai/models/gpt-image
- SDK docs: https://runapi.ai/docs#sdk-gpt-image
- Product docs: https://runapi.ai/docs#gpt-image
- Pricing and rate limits: https://runapi.ai/models/gpt-image
- Provider comparison: https://runapi.ai/providers/openai
- Full catalog: https://runapi.ai/models
- Repository: https://github.com/runapi-ai/gpt-image-sdk

## License

Licensed under the Apache License, Version 2.0.
