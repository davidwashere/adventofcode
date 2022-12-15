#!/bin/bash

export TRIGGERER_GITHUB_TOKEN=
export TRIGGERER_GITHUB_OWNER=davidwashere
export TRIGGERER_GITHUB_REPO=adventofcode
export TRIGGERER_GITHUB_ACTION_FILE=leaderboard.yml
export TRIGGERER_GITHUB_REF=master
export TRIGGERER_INTERVAL=20

# execute in background, forever, and eva
nohup ./triggerer >> triggerer.log 2>&1 &