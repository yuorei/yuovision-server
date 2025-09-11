# CLAUDE.md

このファイルは、このリポジトリでコードを扱うClaude Code (claude.ai/code)にガイダンスを提供します。

## プロジェクト概要

YuoVisionは、Firebase/Firestore/Cloudflare R2を使用したサーバーレス動画配信プラットフォームのAPIサーバーです。GraphQLを使用して動画のアップロード、視聴、コメント機能を提供し、クリーンアーキテクチャの原則に従います。

## アーキテクチャ

コードベースは明確な関心の分離を持つ**クリーンアーキテクチャ**に従っています：

- **`app/domain/`** - コアビジネスエンティティとドメインモデル（Comment、User、Video）
- **`app/application/`** - ビジネスロジックとユースケース
  - **`port/`** - レイヤー間のコントラクトを定義するインターface
- **`app/adapter/`** - 外部アダプター
  - **`infrastructure/`** - 外部サービス（Redis、S3、動画変換、データベースアクセス）  
  - **`presentation/`** - GraphQLリゾルバーとAPIハンドラー
- **`app/driver/`** - 設定とドライバー（データベース、ルーターセットアップ、ログ、モニタリング）

## 主要技術

- **GraphQL API** gqlgenを使用 - GraphQLスキーマからGoの型とリゾルバーを自動生成
- **Firestore** NoSQLデータベースによるメタデータ管理
- **Firebase Authentication** ユーザー認証
- **Cloudflare R2** （S3互換）動画・画像ストレージ
- **Google Cloud Pub/Sub** 非同期動画処理
- **HLSストリーミング** FFmpegで動画処理（Worker Serviceで実行）
- **New Relic** モニタリング用
- **Sentry** エラートラッキング用

## 開発コマンド

### 基本コマンド
```bash
# コードフォーマット
make fmt
./shell/fmt.sh

# リンター実行 
make lint
./shell/lint.sh

# テスト実行
make test
go test -v ./...

# データベースコード生成（sqlc）
make gen
./shell/gen_db.sh

# モック生成
./shell/gen_mock.sh
```

### 開発環境
```bash
# 開発環境セットアップ
make setup_dev
./shell/setup_dev.sh

# 開発モードで実行（.env.devを読み込み）
make dev

# 本番モードで実行（.env.prodを読み込み）
make prod

# Docker開発
make build   # コンテナビルド
make up      # サービス開始
make ps      # コンテナ状態確認
```

## コード生成

このプロジェクトでは複数のコード生成ツールを使用します：

1. **gqlgen** - GraphQLリゾルバーと型を生成（app/domain/models/models_gen.go）
2. **gomock** - テスト用モックを生成（mock/）

スキーマ変更後は常にコード生成を実行してください：
```bash
make gen           # GraphQLコード生成
./shell/gen_mock.sh # モック生成
```

## テスト

- ユニットテストはソースファイルと同じ場所に配置（*_test.go）
- モックはgomockを使用して`mock/`ディレクトリに生成
- `make test`または`go test -v ./...`でテスト実行

## ファイル構造メモ

- **`yuovision-proto/`** - Protocol buffer定義（gRPC使用時）
- **`kubernetes/`** - Kubernetesデプロイメントマニフェスト
- **`shell/`** - 開発用ユーティリティスクリプト
- **`app/driver/firebase/`** - Firestore クライアント実装

## 重要な注意事項

- コミット前には必ず`make fmt`と`make lint`を実行
- アプリケーションは開発環境と本番環境の設定両方をサポート
- 動画ファイルはFFmpegで処理され、HLSセグメントとして保存
- 画像は最適化のためWebP形式に変換
- GraphQLスキーマ変更時はgqlgenでモデルの再生成が必要