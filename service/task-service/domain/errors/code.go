package errors

// Code はエラーコードの定義です。
type Code int

func (c Code) String() string { return string(codes[c]) }

type code string

// error codes
const (
	Unknown Code = iota // undefine code

	IllegalInputBody
	IllegalInputTaskID
	IllegalInputDescription
	NotFoundTask
	DuplicateTask

	Unexpected
)

var codes = map[Code]code{
	IllegalInputBody:        "IllegalInputBody",
	IllegalInputTaskID:      "IllegalInputTaskID",
	IllegalInputDescription: "IllegalInputDescription",
	NotFoundTask:            "NotFoundTask",
	DuplicateTask:           "DuplicateTask",

	Unexpected: "Unexpected",
	Unknown:    "Unknown",
}
