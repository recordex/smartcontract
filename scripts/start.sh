#!/bin/bash

# 開発用ローカルブロックチェーンの起動
npm run node &

# バックグラウンドで開発用ローカルブロックチェーンが起動されるのを待つ
sleep 5
npm run deploylocal
