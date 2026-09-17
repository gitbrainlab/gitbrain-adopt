#!/usr/bin/env python3
"""Minimal GitBrain MCP client — whoami, mount, propose."""
from __future__ import annotations

import json
import os
import re
import sys
import uuid

import httpx

MCP_URL = os.environ.get("GB_MCP_URL", "https://mcp.gitbrain.com/mcp")
TOKEN = os.environ.get("GB_TOKEN")


def rpc(client: httpx.Client, method: str, params: dict | None = None) -> dict:
    body: dict = {"jsonrpc": "2.0", "id": str(uuid.uuid4()), "method": method}
    if params is not None:
        body["params"] = params
    r = client.post(MCP_URL, json=body)
    r.raise_for_status()
    data = r.json()
    if "error" in data:
        raise RuntimeError(data["error"])
    return data.get("result", data)


def tool(client: httpx.Client, name: str, arguments: dict | None = None) -> dict:
    return rpc(client, "tools/call", {"name": name, "arguments": arguments or {}})


def find_session_id(payload: object) -> str | None:
    text = json.dumps(payload)
    for key in ("session_id", "agent_session_id"):
        m = re.search(rf'"{key}"\s*:\s*"([^"]+)"', text)
        if m:
            return m.group(1)
    return None


def main() -> int:
    if not TOKEN:
        print("Set GB_TOKEN from https://app.gitbrain.com/dashboard (or /authorize-agent). See docs/05-signup-e2e.md.", file=sys.stderr)
        return 2
    headers = {"Authorization": f"Bearer {TOKEN}", "Content-Type": "application/json"}
    with httpx.Client(headers=headers, timeout=60.0) as client:
        print("whoami:", json.dumps(tool(client, "whoami"), indent=2))
        mounted = tool(client, "mount_capsule")
        print("mount:", json.dumps(mounted, indent=2)[:2000])
        session_id = find_session_id(mounted)
        if not session_id:
            print("No session_id in mount result — inspect JSON.", file=sys.stderr)
            return 1
        proposed = tool(
            client,
            "propose_capsule_changes",
            {
                "session_id": session_id,
                "message": "adopt-kit python hello",
                "files": [
                    {
                        "path": "notes/adopt-python.md",
                        "content": "# Hello from Python adopt kit\n",
                    }
                ],
            },
        )
        print("propose:", json.dumps(proposed, indent=2)[:2000])
        print(f"Done. session_id={session_id} — human merge if Campus should change.")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
