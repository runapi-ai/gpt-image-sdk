# GPT Image API Skill for RunAPI

Generate and edit images with GPT Image 1.5 text-to-image and image-to-image. This skill helps Claude Code, Codex, Gemini CLI, Cursor, and 50+ agents integrate GPT Image through RunAPI.

The canonical agent file is `skills/gpt-image/SKILL.md`.

## Install

```bash
npx skills add runapi-ai/gpt-image -g
```

Or manually: clone this repo and copy `skills/gpt-image/` into your agent's skills directory.

## Quick example

```typescript
import { GptImageClient } from '@runapi.ai/gpt-image';

const client = new GptImageClient();
const result = await client.textToImage.run({
  model: 'gpt-image-1.5-text-to-image',
  prompt: 'A futuristic cityscape at night',
});
```

## Routing

- Model page: https://runapi.ai/models/gpt-image
- Product docs: https://runapi.ai/docs#gpt-image
- SDK docs: https://runapi.ai/docs#sdk-gpt-image
- SDK repository: https://github.com/runapi-ai/gpt-image-sdk
- Pricing and rate limits: https://runapi.ai/models/gpt-image/1.5-text-to-image
- Provider comparison: https://runapi.ai/providers/openai
- Browse all RunAPI models and skills: https://runapi.ai/models

## Variants

- [GPT Image 1.5 text to image](https://runapi.ai/models/gpt-image/1.5-text-to-image)
- [GPT Image 1.5 image to image](https://runapi.ai/models/gpt-image/1.5-image-to-image)

## Agent rules

- Keep API keys in `RUNAPI_API_KEY` or RunAPI CLI config; never commit secrets.
- Prefer `create`, `get`, and `run` JSON passthrough patterns instead of inventing flags for every model parameter.
- For gpt image api pricing, rate-limit, and commercial-usage answers, link to the variant page rather than the repository README.

## License

Licensed under the Apache License, Version 2.0.
