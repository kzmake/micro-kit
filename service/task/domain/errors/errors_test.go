package errors

import (
	"fmt"
	"testing"

	"golang.org/x/xerrors"

	"github.com/stretchr/testify/require"
)

func TestWrapCode(t *testing.T) {
	testcases := []struct {
		code Code
	}{
		{code: Unknown},
		{code: Unexpected},

		{code: IllegalInputBody},
		{code: IllegalInputTaskID},
		{code: IllegalInputDescription},
		{code: NotFoundTask},
		{code: DuplicateTask},
	}
	for i, tc := range testcases {
		i := i
		tc := tc

		t.Run(fmt.Sprintf("%d: %s", i, tc.code.String()), func(t *testing.T) {
			actual := WrapCode(tc.code, fmt.Errorf(""))

			require.NotNil(t, actual)
			require.IsType(t, &withCode{}, actual)
		})
	}
}

func TestWrapCode_ReturnNil_WhenNil(t *testing.T) {
	testcases := []struct {
		code Code
	}{
		{code: Unknown},
		{code: Unexpected},

		{code: IllegalInputBody},
		{code: IllegalInputTaskID},
		{code: IllegalInputDescription},
		{code: NotFoundTask},
		{code: DuplicateTask},
	}
	for i, tc := range testcases {
		i := i
		tc := tc

		t.Run(fmt.Sprintf("%d: %s", i, tc.code.String()), func(t *testing.T) {
			actual := WrapCode(tc.code, nil)

			require.Nil(t, actual)
		})
	}
}

func TestNewCode(t *testing.T) {
	testcases := []struct {
		code Code
	}{
		{code: Unknown},
		{code: Unexpected},

		{code: IllegalInputBody},
		{code: IllegalInputTaskID},
		{code: IllegalInputDescription},
		{code: NotFoundTask},
		{code: DuplicateTask},
	}
	for i, tc := range testcases {
		i := i
		tc := tc

		t.Run(fmt.Sprintf("%d: %s", i, tc.code.String()), func(t *testing.T) {
			actual := NewCode(tc.code)

			require.NotNil(t, actual)
			require.IsType(t, &withCode{}, actual)
		})
	}
}

func TestGetCode(t *testing.T) {
	testcases := []struct {
		err      error
		expected Code
	}{
		{err: NewCode(Unknown), expected: Unknown},
		{err: NewCode(Unexpected), expected: Unexpected},

		{err: WrapCode(Unknown, fmt.Errorf("")), expected: Unknown},
		{err: WrapCode(Unexpected, fmt.Errorf("")), expected: Unexpected},

		{err: WrapCode(Unknown, WrapCode(Unknown, fmt.Errorf(""))), expected: Unknown},
		{err: WrapCode(Unexpected, WrapCode(Unexpected, fmt.Errorf(""))), expected: Unexpected},

		{err: xerrors.Errorf(":%w", WrapCode(Unknown, fmt.Errorf(""))), expected: Unknown},
		{err: xerrors.Errorf(":%w", WrapCode(Unexpected, fmt.Errorf(""))), expected: Unexpected},

		{err: fmt.Errorf(""), expected: Unknown},

		{err: NewCode(IllegalInputBody), expected: IllegalInputBody},
		{err: NewCode(IllegalInputTaskID), expected: IllegalInputTaskID},
		{err: NewCode(IllegalInputDescription), expected: IllegalInputDescription},
		{err: NewCode(NotFoundTask), expected: NotFoundTask},
		{err: NewCode(DuplicateTask), expected: DuplicateTask},

		{err: WrapCode(IllegalInputBody, fmt.Errorf("")), expected: IllegalInputBody},
		{err: WrapCode(IllegalInputTaskID, fmt.Errorf("")), expected: IllegalInputTaskID},
		{err: WrapCode(IllegalInputDescription, fmt.Errorf("")), expected: IllegalInputDescription},
		{err: WrapCode(NotFoundTask, fmt.Errorf("")), expected: NotFoundTask},
		{err: WrapCode(DuplicateTask, fmt.Errorf("")), expected: DuplicateTask},
	}
	for i, tc := range testcases {
		i := i
		tc := tc

		t.Run(fmt.Sprintf("%d: %s", i, tc.err.Error()), func(t *testing.T) {
			actual := GetCode(tc.err)

			require.Equal(t, tc.expected, actual)
			require.Equal(t, tc.expected.String(), actual.String())
		})
	}
}
