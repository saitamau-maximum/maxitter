# Maxitter

MaximumメンバーがWeb研究部の活動として、Twitterのようなマイクロブログサービス作りを通して共同開発・Webアプリケーション開発の経験を積むことを目的としたプロジェクトです。

## 開発環境

- Docker (Docker Compose)
  - Nginx
  - MySQL (db)
  - Go (back)
  - React (front)

## 開発形態

- ブランチモデル
  - main (本番環境)
  - develop (ステージング環境)
  - feat/xxx (機能追加)
  - fix/xxx (バグ修正)
  - hotfix/xxx (緊急バグ修正)

## 開発手順

### 最初

1. リポジトリをクローンする `git clone https://saitamau-maximum/maxitter.git`
2. `./scripts/setup.sh` でビルドしてコンテナを起動する
