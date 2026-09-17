# 01 — Terms (comparison table)

Skim chips first, then cells — same pattern as visual contract tables.

| Term | Means | Agent can? |
|------|--------|------------|
| **Campus** | Shared tip of record (`main`) | Read |
| **Capsule** | Versioned mount of Campus knowledge | Mount (exact version) |
| **Session** | Write sandbox `refs/sessions/<id>` | Open + propose |
| **Proposal** | Staged change on the session ref | Create |
| **Consensus merge** | Human/admin accept → Campus | **No** |
| **Host token** | Bearer principal + scopes | Use (secret) |
| **Vendor Memory** | Claude/ChatGPT/Grok product memory | Not SoT |

### Dual reading (force shared understanding)

| Misread | Correct reading |
|---------|-----------------|
| “The agent committed to GitBrain” | Agent **proposed**; Campus changes only after human merge |
| “I saved it in ChatGPT Memory, so it’s in GitBrain” | Different stores. Mount GitBrain explicitly |
| “MCP can merge” | MCP has **no** merge tool by design |
