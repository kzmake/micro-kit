package errors

import (
	"github.com/cockroachdb/errors"
	"github.com/cockroachdb/errors/markers"
)

// WrapCode は err に code を wrap して error を生成します。
func WrapCode(code Code, err error) error {
	if err == nil {
		return nil
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
