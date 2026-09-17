# ChatGPT — GitBrain conversational setup

Preferred token path: site signup → dashboard mint — see [docs/05-signup-e2e.md](../docs/05-signup-e2e.md). Keep the token out of the prompt body.

**Two options:** (A) Custom GPT Actions pointing at MCP/REST, or (B) paste bootstrap and call tools via whatever browsing/code channel you enable. Token stays in GPT secrets / Actions auth — **never** in the prompt body.

## Instructions for a Custom GPT

```text
You use GitBrain as durable SoT. ChatGPT Memory is not GitBrain.

When Actions/MCP are configured:
- Base: https://mcp.gitbrain.com/mcp
- Tools: whoami, mount_capsule, read_tree, read_blob, propose_capsule_changes
- Auth: Bearer token minted in GitBrain dashboard (or out-of-band) — store in GPT secrets only

Rules:
1. If you cannot reach GitBrain, say blocked_before_session_creation — do not invent Campus from Memory.
2. One session per task; propose only; no commit/merge.
3. End with session_id, proposal ids, human_merge_requested yes/no.
```

## Actions sketch (OpenAPI-ish)

Map each tool to `POST /mcp` JSON-RPC `tools/call` with `{ "name", "arguments" }`.  
Same shape as [examples/curl](../examples/curl).

## Chat starter

```text
Run whoami against GitBrain, mount_capsule, and list the Campus tree (read_tree).
```

## Voice

See [voice.md](voice.md). In Voice Mode, do not recite tokens; confirm tool results in one breath.
