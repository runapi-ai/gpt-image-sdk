<p align="center">
  <a href="https://github.com/runapi-ai/gpt-image">
    <h3 align="center">GPT Image API Skill for RunAPI</h3>
  </a>
</p>

<p align="center">
  Install this agent skill, inspect GPT Image fields, then run jobs through the RunAPI CLI.
</p>

<p align="center">
  <a href="https://runapi.ai/models/gpt-image"><strong>Model Reference</strong></a> · <a href="https://github.com/runapi-ai/cli"><strong>CLI</strong></a> · <a href="https://github.com/runapi-ai/gpt-image-sdk"><strong>SDK</strong></a>
</p>

<div align="center">

[![skills.sh](https://www.skills.sh/b/runapi-ai/gpt-image)](https://www.skills.sh/runapi-ai/gpt-image/gpt-image)
[![ClawHub](https://img.shields.io/badge/ClawHub-runapi--gpt--image-111827)](https://clawhub.ai/runapi-ai/runapi-gpt-image)
[![License](https://img.shields.io/github/license/runapi-ai/gpt-image)](https://github.com/runapi-ai/gpt-image/blob/main/LICENSE)

</div>
<br/>

Generate and edit images with GPT Image 1.5. This skill helps Claude Code, Codex, Gemini CLI, Cursor, and 50+ agents integrate GPT Image through RunAPI.

The canonical agent file is `skills/gpt-image/SKILL.md`.

## Install

```bash
npx skills add runapi-ai/gpt-image -g
```

Or paste this prompt to your AI agent:

```text
Install the gpt-image skill for me:

1. Clone https://github.com/runapi-ai/gpt-image
2. Copy the skills/gpt-image/ directory into your
   user-level skills directory (e.g. ~/.claude/skills/
   for Claude Code, ~/.codex/skills/ for Codex).
3. Verify that SKILL.md is present.
4. Confirm the install path when done.
```

## Quick example

```typescript
import { GptImageClient } from '@runapi.ai/gpt-image';

const client = new GptImageClient();
const result = await client.textToImage.run({
  model: 'gpt-image-1.5',
  prompt: 'A futuristic cityscape at night',
});
```

## Routing

- Model page: https://runapi.ai/models/gpt-image
- Product docs: https://runapi.ai/docs#gpt-image
- SDK docs: https://runapi.ai/docs#sdk-gpt-image
- SDK repository: https://github.com/runapi-ai/gpt-image-sdk
- Pricing and rate limits: https://runapi.ai/models/gpt-image
- Provider comparison: https://runapi.ai/providers/openai
- Browse all RunAPI models and skills: https://runapi.ai/models

## Model

- [GPT Image 1.5](https://runapi.ai/models/gpt-image)

## Agent rules

- Keep API keys in `RUNAPI_API_KEY` or RunAPI CLI config; never commit secrets.
- Prefer `create`, `get`, and `run` JSON passthrough patterns instead of inventing flags for every model parameter.
- For gpt image api pricing, rate-limit, and commercial-usage answers, link to the model page rather than the repository README.

## License

Licensed under the Apache License, Version 2.0.
