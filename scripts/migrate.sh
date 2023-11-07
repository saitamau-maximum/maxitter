#!/bin/bash

THIS_FILE_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_DIR="$(cd "${THIS_FILE_DIR}/.." && pwd)"
SERVER_DIR="${PROJECT_DIR}/backend"
ENV_FILE="${PROJECT_DIR}/.env"

if [ ! -f "${ENV_FILE}" ]; then
    echo "環境変数ファイルが存在しません！"
    echo "環境変数ファイルを作成してから再度実行してください。"
    exit 1
fi

cd "${SERVER_DIR}"

go run "${SERVER_DIR}/cmd/migrate/main.go" db $1 $2
