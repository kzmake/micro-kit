package constant

type serviceKey = string

const (
	task serviceKey = "task"
)

// Service はサービスの定義です。
var Service = struct {
	Task serviceKey
}{
	Task: task,
}
