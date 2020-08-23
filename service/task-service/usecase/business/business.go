package business

import (
	"context"
)

// Manager はビジネスロジックを管理するマネージャーのIFです。
type Manager interface {
	Execute(context.Context, Task) (interface{}, error)
}

// Task はアシスタントが担当する処理の定義です。
type Task = func(ctx context.Context) (interface{}, error)

// Assistant はビジネスロジックを補佐するアシスタントの定義です。
type Assistant func(next Task) Task

// manager はビジネスロジックを管理するマネージャーの定義です。
type manager struct {
	assistants []Assistant
}

// New はアシスタントを設定したマネージャーを生成します。
func New(assistants ...Assistant) Manager {
	return &manager{append([]Assistant(nil), assistants...)}
}

// Execute は task をアシスタントと共に実行します。
func (m *manager) Execute(ctx context.Context, task Task) (interface{}, error) {
	return m.then(task)(ctx)
}

func (m *manager) then(a Task) Task {
	for i := range m.assistants {
		a = m.assistants[len(m.assistants)-1-i](a)
	}
	return a
}
