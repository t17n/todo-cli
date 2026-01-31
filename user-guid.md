# ToDo CLI 使い方ガイド

このドキュメントでは、ToDo CLIアプリケーションの使い方を説明します。

## セットアップ

### 1. Dockerコンテナの起動

```sh
docker-compose up -d
```

### 2. コンテナに入る

```sh
docker-compose exec app sh
```

### 3. データディレクトリの準備

```sh
mkdir -p data
echo "[]" > data/todos.json
```

### 4. アプリケーションのビルド

```sh
make build
```

または手動でビルド:

```sh
go build -o todo cmd/main.go
```

## 基本的な使い方

### タスクの追加

```sh
./todo add "タスクの内容"
```

例:
```sh
./todo add "買い物に行く"
./todo add "レポートを書く"
./todo add "メールを返信する"
```

### タスク一覧の表示

```sh
./todo list
```

出力例:
```
ID: 1 | [ ] 買い物に行く
ID: 2 | [ ] レポートを書く
ID: 3 | [ ] メールを返信する
```

### タスクを完了にする

```sh
./todo done <タスクID>
```

例:
```sh
./todo done 1
```

完了したタスクは `[x]` マークが付き、リストの下部に移動します。

### タスクの削除

```sh
./todo delete <タスクID>
```

例:
```sh
./todo delete 2
```

### すべてのタスクをクリア

```sh
./todo clear
```

すべてのタスクが削除されます。

## コマンド一覧

| コマンド | 説明 | 使用例 |
|---------|------|--------|
| `add <内容>` | 新しいタスクを追加 | `./todo add "買い物"` |
| `list` | すべてのタスクを表示 | `./todo list` |
| `done <ID>` | タスクを完了にする | `./todo done 1` |
| `delete <ID>` | タスクを削除 | `./todo delete 1` |
| `clear` | すべてのタスクを削除 | `./todo clear` |

## Makefileコマンド

プロジェクトには以下のMakefileコマンドが用意されています:

```sh
# ビルド
make build

# ビルドしてlistコマンドを実行
make run

# ビルド成果物を削除
make clean
```

## データの永続化

- タスクデータは `data/todos.json` に保存されます
- このファイルはホストマシンとコンテナ間で共有されます
- コンテナを再起動してもデータは保持されます

## 使用例

```sh
# 1. タスクを追加
./todo add "プレゼン資料を作成"
./todo add "会議の準備"
./todo add "コードレビュー"

# 2. タスク一覧を確認
./todo list

# 3. タスクを完了にする
./todo done 1

# 4. 再度一覧を確認（完了タスクが下に移動）
./todo list

# 5. 不要なタスクを削除
./todo delete 3

# 6. 最終確認
./todo list
```

## トラブルシューティング

### JSONファイルが見つからない

```sh
mkdir -p data
echo "[]" > data/todos.json
```

### ビルドエラー

```sh
go mod tidy
make build
```

### コンテナが起動しない

```sh
docker-compose down
docker-compose up -d
```

## 終了方法

### コンテナから退出

```sh
exit
```

### コンテナを停止

```sh
docker-compose down
```