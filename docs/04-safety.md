# 04 — Safety & source of truth

Decision cells (yes/no), for busy implementers.

| Question | Answer |
|----------|--------|
| Embed token in a public prompt? | **No** |
| Embed token in Campus markdown? | **No** |
| Agent commit to `main`? | **No** |
| Treat vendor Memory as GitBrain? | **No** |
| Open one session per task? | **Yes** |
| Propose then human-merge? | **Yes** |
| Prefer `mcp.gitbrain.com` over ad-hoc workers.dev hosts? | **Yes** |

### Return shape when the agent is done

Report: `session_id`, `mount_id`, proposal ids, files touched, human merge requested (yes/no).  
Never claim Campus was updated unless an admin consensus receipt exists.
