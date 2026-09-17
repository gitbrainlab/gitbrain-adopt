# Mint GB_TOKEN from https://app.gitbrain.com (Tokens) or out-of-band.
#!/usr/bin/env bash
set -euo pipefail
: "${GB_TOKEN:?set GB_TOKEN}"
GB_MCP_URL="${GB_MCP_URL:-https://mcp.gitbrain.com/mcp}"
auth=(-H "Authorization: Bearer $GB_TOKEN" -H 'content-type: application/json')

echo "== mount_capsule =="
MOUNT=$(curl -sS -X POST "$GB_MCP_URL" "${auth[@]}" \
  -d '{"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"mount_capsule","arguments":{}}}')
echo "$MOUNT"
# Best-effort extract session_id (jq optional)
SESSION_ID=$(echo "$MOUNT" | sed -n 's/.*"session_id"[[:space:]]*:[[:space:]]*"\([^"]*\)".*/\1/p' | head -1)
if [[ -z "${SESSION_ID}" ]]; then
  echo "Could not parse session_id — inspect mount JSON above." >&2
  exit 1
fi

echo "== propose_capsule_changes (session=$SESSION_ID) =="
curl -sS -X POST "$GB_MCP_URL" "${auth[@]}" \
  -d "$(cat <<JSON
{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"propose_capsule_changes","arguments":{
  "session_id":"$SESSION_ID",
  "message":"adopt-kit hello",
  "files":[{"path":"notes/adopt-hello.md","content":"# Hello from adopt kit\n"}]
}}}
JSON
)"
echo
echo "Hand session_id=$SESSION_ID to a human if Campus should merge."
