#!/bin/bash

dt=$(date)

curdir=$(pwd)

echo "Current Dir: $curdir"

echo "list dir"
ls -la

echo "Making cache dir if not exist"
mkdir -p .cache

echo "list dir"
ls -la 

echo "Appending to diff.out"
echo ${dt} >> .cache/diff.out

echo "list cache"
ls -la .cache

echo "dumping cache contents"
cat ./cache/diff.out