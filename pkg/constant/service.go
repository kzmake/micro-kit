package constant

type serviceKey = string

const (
	task serviceKey = "kzmake.microkit.task.v1"
)

// Service はサービスの定義です。
var Service = struct {
	Task serviceKey
}{
	Task: task,
}
