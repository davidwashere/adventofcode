#!/bin/bash

dt=$(date)

curdir=$(pwd)

echo "Current Dir: $curdir"
ls -la

echo "Making cache dir if not exist"
mkdir -p ./.cache

echo "Appending to diff.out"
echo ${dt} >> ./.cache/diff.out

