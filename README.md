# YuoVison Server

[https://yuovision.yuorei.com](https://yuovision.yuorei.com)

YuoVisionは、動画の視聴とアップロードが可能な動画配信プラットフォームです。
あなたの動画をアップロードして、世界中の人々と共有しましょう。新しい発見をYuoVisionでシェアし、一緒に新しい世界を探求しましょう。

## 提供機能
- 視聴
- チャンネル登録
- アップロード
- コメント
- 広告
## 使用技術

### 言語
- Go
### API
- GraphQL
  - スキーマと型を定義することで、GraphQLではドキュメントのようにスキーマが明確になります。これにより、リゾルバーが自動生成される利点があります。また、必要な分だけデータを取得することができるという考え方もよいです。
### インフラ
- [自宅サーバー](https://www.amazon.co.jp/gp/product/B0BZVFM3R4/ref=ppx_yo_dt_b_asin_title_o07_s00?ie=UTF8&th=1)
- Proxmox VE 8.1
  -  Ubuntu22.04
- Cloudflare Tunnel
  - パブリックにルーティング可能なIPアドレスを持たずに安全にCloudflareに接続することができます。  
- ~~Kubernetes~~ 現在はリソース不足のためDockerで動かしています。
### ストリーミング
- HLS
### DB
- MongoDB
  - ドキュメント型NoSQLです。入れ子や配列も保存できて柔軟な構造とインデックスができます。 
- Redis
  - キャッシュサーバとして動かしています。動画一覧を取得した際に1分間キャッシュとして保存します。
### ストレージ
- MinIO
  - S3と互換性のあるストレージサービスです。 
### 認証
- Keycloak
  - SSOと使いたいです。
### コンテナレジストリ
- Docker Hub
### GitHub Actions
- CI
### その他
- FFmpeg
  -　動画に関する様々なことを扱うツールです。     
- [kolesa-team/go-webp](https://github.com/kolesa-team/go-webp)
  - webpに変換するライブラリです。
## 構成図
<img width="662" alt="image" src="https://github.com/yuorei/yuovision-server/assets/108039575/0b128cac-967a-4871-8812-9f172b007789">

## 工夫したところ
### ストリーミングになっている
HLS(HTTP Live Streaming)を使いました。ストリーミング
### オンプレで構築した
さまざまなOSSを駆使してVMの上に構築しました。
### アーキテクチャ
- クリーンアーキテクチャで構築しました。DBなど外部のものに依存しないような設計にしました。
### 設計
- MongoDB
  - 正規化をすることなくなるべく関連する情報を一つのドキュメントに入れました。
- GraphQL
  - スキーマ設計は1回で必要な情報を取得できるようにしました。

## ディレクトリ構成

<pre>
.
├── Dockerfile
├── Dockerfile.dev
├── Makefile
├── app
│   ├── adapter
│   │   ├── infrastructure
│   │   │   ├── comment.go
│   │   │   ├── convert_hls.go
│   │   │   ├── image.go
│   │   │   ├── infrastructure.go
│   │   │   ├── redis.go
│   │   │   ├── upload_video_for_storage.go
│   │   │   ├── user.go
│   │   │   └── video.go
│   │   └── presentation
│   │       └── resolver
│   │           ├── comment.resolvers.go
│   │           ├── node.resolvers.go
│   │           ├── resolver.go
│   │           ├── user.resolvers.go
│   │           └── video.resolvers.go
│   ├── application
│   │   ├── app.go
│   │   ├── comment.go
│   │   ├── image.go
│   │   ├── port
│   │   │   ├── comment.go
│   │   │   ├── image_port.go
│   │   │   ├── user_port.go
│   │   │   └── video_port.go
│   │   ├── usecase.go
│   │   ├── user.go
│   │   ├── video.go
│   │   ├── video_test.go
│   │   └── video_upload.go
│   ├── domain
│   │   ├── comment.go
│   │   ├── domain.go
│   │   ├── id.go
│   │   ├── models
│   │   │   └── models_gen.go
│   │   ├── user.go
│   │   ├── uuid.go
│   │   └── video.go
│   └── driver
│       ├── db
│       │   ├── mongo.go
│       │   └── mongodb
│       │       └── collection
│       │           ├── comment.go
│       │           ├── user.go
│       │           └── video.go
│       ├── redis
│       │   └── redis.go
│       └── router
│           └── router.go
├── compose.yaml
├── go.mod
├── go.sum
├── gqlgen.yml
├── graph
│   ├── generated
│   │   └── generated.go
│   └── schema
│       ├── auth.graphqls
│       ├── comment.graphqls
│       ├── node.graphqls
│       ├── user.graphqls
│       └── video.graphqls
├── kubernetes
│   ├── deployment.yaml
│   ├── secret_example.yaml
│   └── service.yaml
├── lib
│   └── pointers.go
├── main.go
├── middleware
│   └── auth.go
├── mock
│   └── video_port.go
├── public_key.pem
├── public_key.pem.sample
├── shell
│   ├── fmt.sh
│   ├── gen.sh
│   └── gen_mock.sh
└── tools.go
</pre>

### adapter
円の一番外側として定義してあります。
#### infrastructure
外部との通信であったり、DBに接続を行います。
#### presentation
受け取りと出力を行っています。
### application
ビジネスロジックを入れています。
#### port
interfaceを定義してあります。
### domain
さまざまな定義が置いてあります。
### driver
接続、設定が置いてあります。
### graph
リゾルバー以外のGraphQL関連のものが置いてあります。
### kubernetes
Kubernetesのマニフェストファイルが置いてあります。
### lib
ちょっとしたものを置いています。
### middleware
認証をしています。
### mock
gomockの生成ファイルがあります。
### shell
必要なものをまとめてあります。
