#!/bin/bash

## バックグラウンドで開発用ローカルブロックチェーンが起動されるのを待つ
sleep 5 && npx hardhat --max-memory 518 compile && npm run deploylocal &

# 開発用ローカルブロックチェーンの起動
npm run node
