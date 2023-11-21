#!/bin/bash

# 現在のブランチの package.json からバージョンを取得
current_version=$(jq -r '.version' package.json)

# リモートのメインブランチからデータをフェッチ
git fetch origin main

# フェッチしたメインブランチの package.json からバージョンを取得
base_version=$(git show origin/main:package.json | jq -r '.version')

# バージョンを比較
if [ "$(printf '%s\n' "$base_version" "$current_version" | sort -V | tail -n1)" = "$current_version" ]; then
    if [ "$base_version" != "$current_version" ]; then
        echo "OK: Current version $current_version is greater than base version $base_version."
    else
        echo "Error: Current version $current_version is same as base version $base_version."
        exit 1
    fi
else
    echo "Error: Current version $current_version is less than base version $base_version."
    exit 1
fi
