# Triggerer
GitHub actions cron schedule was not / is not triggering consistently during the time frames specified

Triggerer will execute somewhere (super secret server somewhere) and trigger a GH action via `workflow_dispatch` on a schedule

## Needs env vars for configuration

Env | Desc
--- | ---
`TRIGGERER_GITHUB_TOKEN` | A GitHub token (ie a `PAT`) that will be used when talking to the GitHub API
`TRIGGERER_GITHUB_OWNER` | The owner of the repo for which to trigger an action
`TRIGGERER_GITHUB_REPO` | The repo that contains the action to execute
`TRIGGERER_GITHUB_ACTION_FILE` | The filename of the action to execute
`TRIGGERER_GITHUB_REF` | The branch or tag to trigger the workload on
`TRIGGERER_INTERVAL` | Number of seconds to wait between triggers, on launch will trigger immediately (defaults to `1800` = 30 mins)
`TRIGGERER_START_HOUR` | Hour in CST (24 hour format) to start the 'intervals', must be `0 <= n < 24` (defaults to `9` = 9am)
`TRIGGERER_END_HOUR` | Hour in CST (24 hour format) to end the 'intervals', must be `0 <= n < 24` (defaults to `23` = 11pm)
`TRIGGERER_DEBUG` | If `true` will not POST to GitHub, will log a message instead (defaults to `false`)

_ref: https://docs.github.com/en/rest/actions/workflows?apiVersion=2022-11-28#create-a-workflow-dispatch-event_

## Building

For Linux:
```
GOOS=linux GOARCH=amd64 go build -o triggerer ./tools/triggerer
```