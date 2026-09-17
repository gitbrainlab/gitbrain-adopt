# 00 — Picture of GitBrain

**One topic:** how knowledge moves. Visuals complement the rules below.

## Delivery diagram (custody of a change)

```mermaid
flowchart LR
  H[Host AI / app] -->|1 mount| Tip[Campus tip]
  H -->|2 propose| Sand[Session sandbox]
  Sand -->|3 human merge| Tip
```

Stacked responsibility (Incoterms-style, for bundles of docs):

| Until… | Custodian of the change | If this step fails |
|--------|-------------------------|-------------------|
| Mount returns | Campus (read-only tip) | Retry; check token/DNS |
| Propose returns | Session ref only | Fix files; propose again |
| Consensus receipt | Campus `main` | Human retries merge |

## Swimlane — one hop of work

```mermaid
sequenceDiagram
  participant Host as App or chat
  participant MCP as GitBrain MCP
  participant Sess as Session ref
  participant Campus as Campus main
  Host->>MCP: whoami
  Host->>MCP: mount_capsule
  MCP->>Campus: read tip
  MCP->>Sess: open sandbox
  Host->>MCP: propose_capsule_changes
  MCP->>Sess: stage Proposal
  Note over Campus: main unchanged until human merge
```

## Rules under the diagram

1. Bearer host token out-of-band (`GB_TOKEN`). Never paste tokens into Campus docs or public prompts.
2. One session per task. Keep `session_id`.
3. Writes = propose only. No agent commit.
4. Graph is a **projection** of OKF/git objects — not a second database of truth.
5. ChatGPT / Claude / Grok Memory ≠ GitBrain SoT.
