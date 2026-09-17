# curl examples

Requires: `GB_TOKEN` (mint at `https://app.gitbrain.com/dashboard` or `/authorize-agent`; dashboard tokens look like `gb_live_…`), optional `GB_MCP_URL` (default `https://mcp.gitbrain.com/mcp`).

```bash
export GB_TOKEN=...          # host secret store — never the prompt
export GB_MCP_URL=${GB_MCP_URL:-https://mcp.gitbrain.com/mcp}
./whoami.sh
./mount_and_propose.sh
```
