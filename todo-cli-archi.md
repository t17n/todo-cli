# ToDo CLI 設計書

---

## 1. 概要
- プロジェクト名：TODO CLI
- 目的：
  - CLIでToDo管理を簡易に行うツール
- 実行環境：
  - Docker（golang:1.22）
  - JSONファイルをローカルへ永続化

---

## 2. 全体像

### 2.1 できること（機能一覧）
- [ ] ToDoを追加（add）
- [ ] ToDo一覧を表示（list）
- [ ] ToDoを完了にする（done）
- [ ] ToDoを削除（delete）
- [ ] 全削除（clear）※任意

---

## 3. データ設計

### 3.1 JSONファイル構造
ファイルパス（ホスト側）：  
`./data/todos.json`

コンテナ内のパス：  
`/app/data/todos.json`

JSONの例：

```json
[
  {
    "id": 1,
    "title": "サンプル",
    "done": false
  }
]
```

---

## 4. 実装手順

### 4.1 Dockerfile作成
- golang:1.22 を使用
- `/app` を作業ディレクトリにする
- 開発時はローカルを bind mount して実行する方針

### 4.2 ディレクトリ構成の作成
- `/cmd` … main.go を置く（Cobra未使用なら不要）
  - Goの慣習で、実行コマンド（CLI本体）を置くフォルダ。
  - 大規模CLI（Cobraなど）でもよく使う構造。
- `/data` … todos.json を置く
- `/internal` … 読み書きロジック（tasks.go）を置く
  - Goの慣習で、実行コマンド（CLI本体）を置くフォルダ。
  - 大規模CLI（Cobraなど）でもよく使う構造。
- `project/` … プロジェクトのルートフォルダ
- `Dockerfile` … CLIをDocker上で実行するための環境定義
- `cmd/` … エントリーポイント(main.go)を置く場所
- `main.go` … CLIの入口。os.Argsを解析して処理を振り分ける
- `internal/` … アプリ内部のロジックを置く（外部に公開しない）
- `tasks.go` … ToDoの読み書き・追加・削除などのロジック
- `data/` … データ保存用フォルダ（JSONを置く）
  - コンテナ内でもホストでも同じ場所にしやすいので、永続化に使いやすい。
- `todos.json` … ToDoデータを保存する実ファイル

todo-cli/
├── Dockerfile
├── docker-compose.yml
├── todo-cli-archi.md
├── cmd/
│   └── main.go          # CLIのエントリーポイント
├── internal/
│   └── tasks.go         # ToDoの読み書きロジック
└── data/
    └── todos.json       # ToDoデータ（永続化）


### 4.3 JSON読み書き処理の実装
- LoadTasks() … todos.json を読みこむ
- SaveTasks() … todos.json に書きこむ
- ファイルが空の場合は空配列を返すようにする
- Todo struct を定義する

### 4.4 コマンドごとの処理実装
- add：タイトルを受け取り、ID採番して追加
- list：未完了→完了の順で出力
- done：ID一致の要素を done=true
- delete：ID一致の要素を削除
- clear：空配列を書き込む（任意）

### 4.5 main.go のルーティング
- os.Args を使って add / list / done / delete / clear を振り分ける
- 想定外の入力は usage を表示する



---

## 5. テスト

### 5.1 ToDoを追加（add）
- [ ] 空の todos.json から追加
- [ ] 複数回追加して ID が連番になっているか

### 5.2 ToDo一覧を表示（list）
- [ ] done=false → done=true の順で並んでいるか
- [ ] 表示形式が崩れていないか

### 5.3 ToDoを完了にする（done）
- [ ] 完了後に list で下に移動しているか
- [ ] 存在しないIDを指定した場合のメッセージ

### 5.4 ToDoを削除（delete）
- [ ] 削除後、IDの再採番は行わない（既存維持）
- [ ] 削除後に JSON が正しい構造か確認

### 5.5 全削除（clear）
- [ ] todos.json が空配列になるか
- [ ] list 実行時に「0件」と出るか

---

