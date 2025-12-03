#!/bin/bash

# Setup on new year:
# - Update SESSION_COOKIE value in GitHub Actions (from browser dev tools)
# - Update LEADER_URL in GitHub Actions (Grab the new year's URL from the browser)


# Generates a list of players, stars, completeing times, and diffs
#
# takes one optional argument ./leader.sh [force], when used will overwrite the cache
go run ./tools/leader $@