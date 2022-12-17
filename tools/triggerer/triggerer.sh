#!/bin/bash

export TRIGGERER_GITHUB_TOKEN=FAKE
export TRIGGERER_GITHUB_OWNER=davidwashere
export TRIGGERER_GITHUB_REPO=adventofcode
export TRIGGERER_GITHUB_ACTION_FILE=leaderboard.yml
export TRIGGERER_GITHUB_REF=master
export TRIGGERER_INTERVAL=1800
export TRIGGERER_START_HOUR=9
export TRIGGERER_END_HOUR=23
export TRIGGERER_DEBUG=true

# execute in background, forever, and eva
nohup ./triggerer >> triggerer.log 2>&1 &