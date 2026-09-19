# GitBrain adopt kit

**Turnkey** ways to use [GitBrain](https://gitbrain.com) ([kit repo](https://github.com/gitbrainlab/gitbrain-adopt)) as durable knowledge for AI — in code, or in a chat/voice session.

> Visuals **complement** text (same stance as Haapio & Passera on contract design): skim the diagrams, then read the short rules under them.

`https://gitbrain.com` is **marketing only**. Accounts, tokens, and the wiki live on `https://app.gitbrain.com`. Agents connect at `https://mcp.gitbrain.com`.

## Two doors

| Door | Who it is for | Start here |
|------|----------------|------------|
| **Programmatic** | Apps, agents, scripts (Python / TypeScript / Go / curl) | [docs/02-programmatic.md](docs/02-programmatic.md) → [examples/](examples/) |
| **Conversational** | Claude, ChatGPT, Grok (chat or voice) | [docs/03-conversational.md](docs/03-conversational.md) → [prompts/](prompts/) |
| **Signup / site** | Account, tokens, wiki viewer | [docs/05-signup-e2e.md](docs/05-signup-e2e.md) → `https://app.gitbrain.com` |

## One-picture story

```mermaid
flowchart LR
  subgraph You
    A[App or chat session]
  end
  subgraph GitBrain
    M[MCP / REST]
    S[Session sandbox<br/>refs/sessions/…]
    C[Campus main<br/>consensus only]
  end
  A -->|Bearer host token| M
  M -->|mount + read| C
  M -->|propose only| S
  S -.->|human merge| C
```

**What transfers (delivery diagram):**

| Stage | You can do | Who holds risk |
|-------|------------|----------------|
| Mount | Read Campus tip; open a session | Host token scopes |
| Propose | Stage files on the session ref | Still a proposal — not Campus |
| Consensus merge | Human/admin only | Campus updates only here |

Agents **never** get a commit tool. Vendor Memory (ChatGPT / Claude / Grok) is **not** GitBrain.

## 5-minute path

1. **Sign up or log in** at `https://app.gitbrain.com/signup` or `https://app.gitbrain.com/login` (email magic link; form posts to `/auth/magic-link` with `intent=signup` or `intent=login`).
2. **Open the verify link.** Until Resend is configured on the app Worker, the check-email page shows a **demo / local verify link** (`/auth/verify?token=…`) instead of sending mail — use that link.
3. **Mint an agent token** at `https://app.gitbrain.com/dashboard` (or `https://app.gitbrain.com/authorize-agent`). Dashboard tokens are prefixed `gb_live_…` and shown **once**. Ops host/service tokens are a separate class.
4. Copy [`.env.example`](.env.example) → `.env` and set `GB_TOKEN` in the **host secret store** (never commit a real token; never paste it into a chat prompt).
5. Pick a door: run an [example](examples/) **or** paste a [prompt](prompts/) into Claude / ChatGPT / Grok with the token only in the host secret store.
6. Confirm with `whoami` / “who am I”, then `mount_capsule` / open session, then propose a tiny note.
7. Review on `https://app.gitbrain.com/wiki` (or `/wiki/getting-started`). Wiki may show **Demo mode** until `MCP_SERVICE_TOKEN` is set on the app Worker. If Campus should change, hand a human the `session_id` + proposal ids — do not claim main was updated.

Full swimlane: [docs/05-signup-e2e.md](docs/05-signup-e2e.md).

## Explainers (one topic each)

| # | Topic | Pattern |
|---|--------|---------|
| [00](docs/00-picture.md) | Picture of the system | Delivery diagram + swimlane |
| [01](docs/01-terms.md) | Terms compared | Comparison table |
| [02](docs/02-programmatic.md) | Code / API stacks | Flowchart + examples |
| [03](docs/03-conversational.md) | Chat & voice hosts | Swimlane per host |
| [04](docs/04-safety.md) | Safety & SoT rules | Yes/no decision cells |
| [05](docs/05-signup-e2e.md) | Signup → MCP → site wiki | Swimlane + URL map |
| [06](docs/06-panels-research.md) | Persistent panels (research only) | Registry + review-run planning; not product |

## Live endpoints (MVP)

| Surface | URL |
|---------|-----|
| Human app (home) | `https://app.gitbrain.com/` |
| Sign up | `https://app.gitbrain.com/signup` |
| Log in | `https://app.gitbrain.com/login` |
| Dashboard (auth; mint token) | `https://app.gitbrain.com/dashboard` |
| Authorize agent (auth) | `https://app.gitbrain.com/authorize-agent` |
| Wiki (auth) | `https://app.gitbrain.com/wiki` |
| Wiki: getting started (auth) | `https://app.gitbrain.com/wiki/getting-started` |
| MCP (Streamable-HTTP) | `https://mcp.gitbrain.com/mcp` |
| Health | `https://mcp.gitbrain.com/health` |

Auth forms (not GET pages): `POST /auth/magic-link` (fields: `email`, `intent` = `signup` or `login`); `GET /auth/verify?token=…`; `POST /tokens` (secret shown once).

MCP auth: `Authorization: Bearer <token>`. Tools: `whoami`, `mount_capsule`, `read_tree`, `read_blob`, `propose_capsule_changes`.

## License

Apache-2.0. See [LICENSE](LICENSE).
