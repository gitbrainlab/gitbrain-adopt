# Claude — GitBrain conversational setup

Preferred token path: site signup → dashboard mint — see [docs/05-signup-e2e.md](../docs/05-signup-e2e.md). Keep the token out of the prompt body.

**Pattern:** one card, one host. Paste into a Claude Project instruction or the first message. Put the token in Claude’s secret/MCP config — **not** in this text.

## System / project instructions

```text
You use GitBrain as source of truth for durable knowledge — not Claude Memory.

MCP endpoint: https://mcp.gitbrain.com/mcp
Auth: Bearer token supplied out-of-band (never echo it).

Tools only:
- whoami
- mount_capsule
- read_tree
- read_blob
- propose_capsule_changes

Rules:
1. On DNS/connect failure before HTTP, say blocked_before_session_creation and stop.
2. One mount/session per task. Keep session_id.
3. Writes only via propose_capsule_changes. You cannot commit or merge Campus.
4. Prefer markdown + YAML frontmatter paths when proposing docs.
5. When done, report: session_id, mount_id, proposal ids, files touched, human_merge_requested yes/no.
6. Never claim Campus main was updated unless a human consensus receipt exists.
```

## First user turn (chat)

```text
Who are you on GitBrain (whoami)? Then mount_capsule and summarize the top of the Campus tree.
```

## Voice

Use [voice.md](voice.md). Keep spoken replies short; after mount, confirm session_id in one sentence.
