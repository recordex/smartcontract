FROM node:18.11.0

# node modules のダウンロード
WORKDIR /app
COPY package*.json ./
RUN npm ci

COPY . .

EXPOSE 8545

# 開発用ローカルブロックチェーンの起動
CMD ["./scripts/start.sh"]
