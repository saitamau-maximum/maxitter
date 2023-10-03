#!/bin/bash

THIS_FILE_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_DIR="$(cd "${THIS_FILE_DIR}/.." && pwd)"
FRONT_DIR="${PROJECT_DIR}/frontend"
BACK_DIR="${PROJECT_DIR}/backend"
ENV_FILE="${PROJECT_DIR}/.env"

echo "Dockerを立ち上げます..."

docker compose up -d --build

echo "Dockerを立ち上げました！"
