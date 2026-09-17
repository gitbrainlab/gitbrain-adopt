# 02 — Programmatic access

**One topic:** call GitBrain from code.

```mermaid
flowchart TD
  A[Load GB_TOKEN] --> B[whoami]
  B --> C[mount_capsule]
  C --> D[read_tree / read_blob]
  D --> E[propose_capsule_changes]
  E --> F[Return session_id + proposal ids to a human]
```

## Pick a stack

| Stack | Folder | Notes |
|-------|--------|-------|
| curl | [examples/curl](../examples/curl) | Fastest smoke |
| Python | [examples/python](../examples/python) | `httpx`, no SDK lock-in |
| TypeScript | [examples/typescript](../examples/typescript) | Node 20+, `fetch` |
| Go | [examples/go](../examples/go) | Stdlib `net/http` |

All four speak **MCP JSON-RPC** to `GB_MCP_URL` (default `https://mcp.gitbrain.com/mcp`).

## Minimal contract

```text
Authorization: Bearer $GB_TOKEN
Content-Type: application/json

POST { whoami | tools/call mount_capsule | read_tree | read_blob | propose_capsule_changes }
```

REST twins exist on the same Worker (`/api/v1/me`, `/v1/sessions`, …) — see the session Worker repo if you prefer REST. This adopt kit standardizes on MCP so every host looks the same.
