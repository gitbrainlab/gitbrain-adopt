# Mint GB_TOKEN from https://app.gitbrain.com/dashboard (or /authorize-agent).
# Dashboard tokens are prefixed gb_live_… — never put the token in a prompt.
#!/usr/bin/env bash
set -euo pipefail
: "${GB_TOKEN:?set GB_TOKEN}"
GB_MCP_URL="${GB_MCP_URL:-https://mcp.gitbrain.com/mcp}"
curl -sS -X POST "$GB_MCP_URL" \
  -H "Authorization: Bearer $GB_TOKEN" \
  -H 'content-type: application/json' \
  -d '{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"whoami","arguments":{}}}'
echo
