/** Minimal GitBrain MCP client — whoami, mount, propose. */
const MCP_URL = process.env.GB_MCP_URL ?? "https://mcp.gitbrain.com/mcp";
const TOKEN = process.env.GB_TOKEN;

async function tool(name: string, arguments_: Record<string, unknown> = {}) {
  if (!TOKEN) throw new Error("Set GB_TOKEN");
  const res = await fetch(MCP_URL, {
    method: "POST",
    headers: {
      Authorization: `Bearer ${TOKEN}`,
      "content-type": "application/json",
    },
    body: JSON.stringify({
      jsonrpc: "2.0",
      id: crypto.randomUUID(),
      method: "tools/call",
      params: { name, arguments: arguments_ },
    }),
  });
  if (!res.ok) throw new Error(`HTTP ${res.status}`);
  const data = await res.json();
  if (data.error) throw new Error(JSON.stringify(data.error));
  return data.result ?? data;
}

function findSessionId(payload: unknown): string | undefined {
  const text = JSON.stringify(payload);
  const m = text.match(/"session_id"\s*:\s*"([^"]+)"/) ||
    text.match(/"agent_session_id"\s*:\s*"([^"]+)"/);
  return m?.[1];
}

async function main() {
  console.log("whoami", await tool("whoami"));
  const mounted = await tool("mount_capsule");
  console.log("mount", mounted);
  const sessionId = findSessionId(mounted);
  if (!sessionId) throw new Error("No session_id in mount result");
  const proposed = await tool("propose_capsule_changes", {
    session_id: sessionId,
    message: "adopt-kit typescript hello",
    files: [
      {
        path: "notes/adopt-typescript.md",
        content: "# Hello from TypeScript adopt kit\n",
      },
    ],
  });
  console.log("propose", proposed);
  console.log(`Done. session_id=${sessionId} — human merge if Campus should change.`);
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
