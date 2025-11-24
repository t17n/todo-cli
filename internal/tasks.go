package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// タスク構造体
type Task struct {
	ID	    int    `json:"id"`
	Title	string `json:"title"`
	Done	bool   `json:"done"`
}

// todos.json ファイルのパス
const todoFile = "data/todos.json"

// tasks.json からタスクを読み込む
func LoadTasks() ([]Task, error) {
	// タスク読み込み処理
	t, err := os.ReadFile(todoFile)
	if err != nil {
		fmt.Printf("ファイル読み込みエラー:", err)
		return nil, err
	}
	// JSONをデコード
	var tasks []Task
	err = json.Unmarshal(t, &tasks)
	if err != nil {
		fmt.Printf("JSON変換エラー:", err)
		return nil, err
	}
	// 読み込んだタスクを返す
	return tasks, nil
}

// tasks.json にタスクを書き込む
func SaveTasks(tasks []Task) error {
	// JSONにエンコード
	t, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Printf("JSON変換エラー:", err)
		return err
	}
	// ファイルに書き込み
	err = os.WriteFile(todoFile, t, 0644)
	if err != nil {
		fmt.Printf("ファイル書き込みエラー:", err)
		return err
	}
	return nil
}

// タスクを追加する
func AddTask(title string) error {
	// タスク取得
	t, err := LoadTasks()
	if err != nil {
		return err
	}
	// 新しいIDを生成(既存の最大ID + 1)
	newID := 1
	for _, task := range t {
		if task.ID >= newID {
			newID = task.ID + 1
		}
	}
	// 新しいタスクを追加
	t = append(t, Task{
		ID:    newID,
		Title: title,
		Done:  false,
	})
	// タスクを保存
	err = SaveTasks(t)
	if err != nil {
		return err
	}
	return nil
}

// タスクの一覧を取得する
func ListTasks() ([]Task, error) {
	// タスク取得
	t, err := LoadTasks()
	if err != nil {
		return nil, err
	}
	// task.Done が false のものを先に、true のものを後に並べ替え
	sort.Slice(t, func(i, j int) bool {
		if t[i].Done == t[j].Done {
			return t[i].ID < t[j].ID
		}
		return !t[i].Done && t[j].Done
	})
	// 並べ替えたタスクを返す
	return t, nil
}

// タスクを完了にする
func MarkTaskDone(id int) error {
	// タスク取得
	t, err := LoadTasks()
	if err != nil {
		return err
	}
	// 指定されたIDのタスクを完了にする
	found := false
	for i, task := range t {
		if task.ID == id {
			t[i].Done = true
			found = true
			break
		}
	}
	// IDが見つからない場合
	if !found {
		return fmt.Errorf("タスクID %d が見つかりません。", id)
	}
	// タスクを保存
	return SaveTasks(t)
}

// タスクを削除する
func DeleteTask(id int) error {
	// タスク取得
	t, err := LoadTasks()
	if err != nil {
		return err
	}
	// 指定されたIDのタスクを削除する
	found := false
	for i, task := range t {
		if task.ID == id {
			t = append(t[:i], t[i+1:]...)
			found = true
			break
		}
	}
	// IDが見つからない場合
	if !found {
		return fmt.Errorf("タスクID %d が見つかりません。", id)
	}
	// タスクを保存
	return SaveTasks(t)
}

// タスクを全て削除する
func ClearTasks() error {
	// タスクを空にして保存
	err := SaveTasks([]Task{})
	if err != nil {
		return err
	}

	return nil
}
