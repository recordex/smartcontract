#!/bin/bash

# 開発用ローカルブロックチェーンの起動
npm run node --hostname localhost &

# バックグラウンドで開発用ローカルブロックチェーンが起動されるのを待つ
sleep 5
npx hardhat --max-memory 518 compile
npm run deploylocal

# スクリプトが終了しないようにする
tail -f /dev/null
