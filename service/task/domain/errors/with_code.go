package errors

import (
	"fmt"

	"github.com/cockroachdb/errors"
)

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
