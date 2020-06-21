package errors

import (
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/cockroachdb/errors/markers"
)

// WrapCode は err に code を wrap して error を生成します。
func WrapCode(code Code, err error) error {
	if err == nil {
		return NewCode(code)
	}
	return &withCode{cause: err, code: code}
}

// NewCode は code を持つ error を生成します。
func NewCode(code Code) error {
	return &withCode{cause: errors.New(""), code: code}
}

// GetCode は error から code を取得します。
func GetCode(err error) Code {
	if v, ok := markers.If(err, func(err error) (interface{}, bool) {
		if w, ok := err.(*withCode); ok {
			return w.code, true
		}
		return nil, false
	}); ok {
		return v.(Code)
	}
	return Unknown
}

type withCode struct {
	cause error
	code  Code
}

func (w *withCode) Error() string                 { return w.code.String() }
func (w *withCode) Cause() error                  { return w.cause }
func (w *withCode) Unwrap() error                 { return w.cause }
func (w *withCode) Format(s fmt.State, verb rune) { errors.FormatError(w, s, verb) }
func (w *withCode) FormatError(p errors.Printer) (next error) {
	if p.Detail() {
		p.Printf("code: %s", w.Error())
	}
	return w.cause
}
