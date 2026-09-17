# curl examples

Requires: `GB_TOKEN`, optional `GB_MCP_URL` (default `https://mcp.gitbrain.com/mcp`).

```bash
export GB_TOKEN=...          # out-of-band
export GB_MCP_URL=${GB_MCP_URL:-https://mcp.gitbrain.com/mcp}
./whoami.sh
./mount_and_propose.sh
```
