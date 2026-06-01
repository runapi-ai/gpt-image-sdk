# GPT Image API Ruby SDK for RunAPI

The gpt image api Ruby SDK is the language-specific package for GPT Image on RunAPI. Use this gpt image api package for text-to-image, image editing, and creative production flows when your application needs JSON request bodies, task status lookup, and consistent RunAPI errors in Ruby.

This gpt image api README is the Ruby package guide inside the public `gpt-image-sdk` repository. For the repository overview, start at `../README.md`; for model details, use https://runapi.ai/models/gpt-image; for API reference, use https://runapi.ai/docs#gpt-image; for SDK docs, use https://runapi.ai/docs#sdk-gpt-image.

## Install

```bash
gem install runapi-gpt-image
```

## Quick start

```ruby
require "runapi-gpt-image"

client = RunApi::GptImage::Client.new
task = client.generations.create(
  # Pass the GPT Image JSON request body from https://runapi.ai/docs#gpt-image.
)
status = client.generations.get(task.id)
```

Use `create` when you want to submit a task and return quickly, `get` when you need the latest task state, and `run` when a script should create and poll until completion. In web request handlers, prefer `create` plus webhook or later `get` polling so a worker is not held open.

## Language notes

Use Ruby keyword arguments and the `RunApi::GptImage` error classes when building image jobs, Rails workers, or scripts. The available resources include generations, and edits. Keep `RUNAPI_API_KEY` in the environment or your secret manager; never commit API keys or callback secrets.

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
