# Go DI Sample - オニオンアーキテクチャ + 依存性注入

このプロジェクトは、Goでオニオンアーキテクチャ（クリーンアーキテクチャ）と依存性注入（DI）を実装したサンプルです。Echoフレームワークを使用したWebサーバーを例にしています。

## アーキテクチャ

```
cmd/server/          # エントリーポイント
internal/
├── domain/          # ドメイン層（ビジネスロジックの核心）
│   ├── entities/    # エンティティ
│   └── repositories/ # リポジトリインターフェース
├── usecase/         # ユースケース層（アプリケーションロジック）
├── infrastructure/  # インフラ層（外部との接続）
│   ├── repositories/ # リポジトリ実装
│   └── handlers/    # HTTPハンドラー
└── di/             # 依存性注入コンテナ
```

## 特徴

- **オニオンアーキテクチャ**: 依存関係が内側に向かう設計
- **依存性注入**: コンパイル時に依存関係を解決
- **インターフェース分離**: 各層がインターフェースに依存
- **テスタビリティ**: モックを使ったテストが容易

## セットアップ

1. 依存関係のインストール:
```bash
go mod tidy
```

2. サーバーの起動:
```bash
go run cmd/server/main.go
```

## API エンドポイント

- `GET /health` - ヘルスチェック
- `POST /api/v1/users` - ユーザー作成
- `GET /api/v1/users` - 全ユーザー取得
- `GET /api/v1/users/:id` - ユーザー取得
- `PUT /api/v1/users/:id` - ユーザー更新
- `DELETE /api/v1/users/:id` - ユーザー削除

## 使用例

### ユーザー作成
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}'
```

### ユーザー取得
```bash
curl http://localhost:8080/api/v1/users
```

## 依存性注入の仕組み

1. **Container**: すべての依存関係を管理
2. **インターフェース**: 各層の境界を定義
3. **実装**: 具体的な実装を注入
4. **初期化**: アプリケーション起動時に依存関係を解決

この設計により、各層が疎結合になり、テストや保守が容易になります。
# go-di-sample
