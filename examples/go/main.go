package main

import (
        "bytes"
        "encoding/json"
        "fmt"
        "io"
        "net/http"
        "os"
        "regexp"
)

func mcpURL() string {
        if v := os.Getenv("GB_MCP_URL"); v != "" {
                return v
        }
        return "https://mcp.gitbrain.com/mcp"
}

func tool(name string, args map[string]any) (any, error) {
        token := os.Getenv("GB_TOKEN")
        if token == "" {
                return nil, fmt.Errorf("set GB_TOKEN")
        }
        body := map[string]any{
                "jsonrpc": "2.0",
                "id":      1,
                "method":  "tools/call",
                "params": map[string]any{
                        "name":      name,
                        "arguments": args,
                },
        }
        b, _ := json.Marshal(body)
        req, err := http.NewRequest(http.MethodPost, mcpURL(), bytes.NewReader(b))
        if err != nil {
                return nil, err
        }
        req.Header.Set("Authorization", "Bearer "+token)
        req.Header.Set("Content-Type", "application/json")
        res, err := http.DefaultClient.Do(req)
        if err != nil {
                return nil, err
        }
        defer res.Body.Close()
        raw, _ := io.ReadAll(res.Body)
        if res.StatusCode >= 300 {
                return nil, fmt.Errorf("HTTP %s: %s", res.Status, string(raw))
        }
        var parsed map[string]any
        if err := json.Unmarshal(raw, &parsed); err != nil {
                return nil, err
        }
        if errObj, ok := parsed["error"]; ok {
                return nil, fmt.Errorf("rpc error: %v", errObj)
        }
        if result, ok := parsed["result"]; ok {
                return result, nil
        }
        return parsed, nil
}

func sessionID(v any) string {
        b, _ := json.Marshal(v)
        for _, re := range []*regexp.Regexp{
                regexp.MustCompile(`"session_id"\s*:\s*"([^"]+)"`),
                regexp.MustCompile(`"agent_session_id"\s*:\s*"([^"]+)"`),
        } {
                if m := re.FindSubmatch(b); m != nil {
                        return string(m[1])
                }
        }
        return ""
}

func main() {
        who, err := tool("whoami", map[string]any{})
        if err != nil {
                panic(err)
        }
        fmt.Printf("whoami: %v\n", who)

        mounted, err := tool("mount_capsule", map[string]any{})
        if err != nil {
                panic(err)
        }
        fmt.Printf("mount: %v\n", mounted)
        sid := sessionID(mounted)
        if sid == "" {
                panic("no session_id in mount result")
        }
        proposed, err := tool("propose_capsule_changes", map[string]any{
                "session_id": sid,
                "message":    "adopt-kit go hello",
                "files": []map[string]string{
                        {"path": "notes/adopt-go.md", "content": "# Hello from Go adopt kit\n"},
                },
        })
        if err != nil {
                panic(err)
        }
        fmt.Printf("propose: %v\n", proposed)
        fmt.Printf("Done. session_id=%s — human merge if Campus should change.\n", sid)
}
