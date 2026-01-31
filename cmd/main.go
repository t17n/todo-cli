package main

import (
	"fmt"
	"os"

	"todo-cli/internal"
)

func main() {
	// JSONファイルのコマンドライン引数が足りない場合はusageを表示して終了
	if len(os.Args) < 2 {
		printUsage()
		return
	}
	// JSONファイルのコマンドライン引数を取得
	command := os.Args[1]
	// コマンドごとに処理を分岐
	switch command {
	// 追加する場合
	case "add":
		if len(os.Args) < 3 {
			fmt.Printf("追加するタスクのタイトルを指定してください。\n")
			return
		}
		title := os.Args[2]
		err := internal.AddTask(title)
		if err != nil {
			fmt.Printf("タスクの追加に失敗しました: %v\n", err)
			return
		}
		fmt.Println("タスクを追加しました。")
	// 一覧表示する場合
	case "list":
		tasks, err := internal.ListTasks()
		if err != nil {
			fmt.Printf("タスクの取得に失敗しました: %v\n", err)
			return
		}
		for _, task := range tasks {
			status := "未完了"
			if task.Done {
				status = "完了"
			}
			fmt.Printf("[%d] %s - %s\n", task.ID, task.Title, status)
		}
		// タスクが0件の場合はメッセージを表示
		if len(tasks) == 0 {
			fmt.Println("タスクはありません。")
			return
		}
	// 完了にする場合
	case "done":
		if len(os.Args) < 3 {
			fmt.Printf("完了するタスクのIDを指定してください。\n")
			return
		}
		var id int
		_, err := fmt.Sscanf(os.Args[2], "%d", &id)
		if err != nil {
			fmt.Printf("無効なタスクIDです: %v\n", err)
			return
		}
		err = internal.MarkTaskDone(id)
		if err != nil {
			fmt.Printf("タスクの完了に失敗しました: %v\n", err)
			return
		}
		fmt.Println("タスクを完了しました。")
	// 削除する場合
	case "delete":
		if len(os.Args) < 3 {
			fmt.Printf("削除するタスクのIDを指定してください。\n")
			return
		}
		var id int
		_, err := fmt.Sscanf(os.Args[2], "%d", &id)
		if err != nil {
			fmt.Printf("無効なタスクIDです: %v\n", err)
			return
		}
		err = internal.DeleteTask(id)
		if err != nil {
			fmt.Printf("タスクの削除に失敗しました: %v\n", err)
			return
		}
		fmt.Println("タスクを削除しました。")
	// すべてのタスクをクリアする場合
	case "clear":
		// 確認メッセージを表示
		var response string
		fmt.Print("本当にすべてのタスクを削除しますか？ (y/n): ")
		fmt.Scanln(&response)
		if response != "y" {
			fmt.Println("操作がキャンセルされました。")
			return
		}
		err := internal.ClearTasks()
		if err != nil {
			fmt.Printf("タスクのクリアに失敗しました: %v\n", err)
			return
		}
		fmt.Println("すべてのタスクを削除しました。")
	// 不明なコマンドの場合
	default:
		printUsage()
		fmt.Printf("不明なコマンドです: %s\n", command)
		return
	}
}

// 使用方法を表示する関数
func printUsage() {
	fmt.Println("使い方:")
	fmt.Println("  ./todo add <タイトル>     - タスクを追加")
	fmt.Println("  ./todo list              - タスク一覧を表示")
	fmt.Println("  ./todo done <ID>         - タスクを完了")
	fmt.Println("  ./todo delete <ID>       - タスクを削除")
	fmt.Println("  ./todo clear             - すべてのタスクを削除")
}