package errors

import (
	"fmt"
	"testing"

	"golang.org/x/xerrors"

	"github.com/stretchr/testify/require"
)

func TestWithCode(t *testing.T) {
	testcases := []struct {
		err error
	}{
		{err: fmt.Errorf("stderror")},
	}
	for i, tc := range testcases {
		i := i
		tc := tc

		t.Run(fmt.Sprintf("%d: %s", i, tc.err), func(t *testing.T) {
			code := Unexpected
			wcode := WrapCode(code, tc.err)

			require.NotNil(t, wcode)
			require.True(t, xerrors.Is(wcode, wcode))
			require.True(t, xerrors.Is(wcode, tc.err))
			require.True(t, xerrors.Is(xerrors.Errorf(": %w", wcode), tc.err))
			require.True(t, xerrors.Is(xerrors.Errorf(": %w", wcode), wcode))
			var err *withCode
			require.True(t, xerrors.As(wcode, &err))
			require.True(t, xerrors.As(xerrors.Errorf(": %w", wcode), &err))

			require.Equal(t, code.String(), wcode.(*withCode).Error())
			require.Equal(t, tc.err, wcode.(*withCode).Cause())
			require.Equal(t, tc.err, wcode.(*withCode).Unwrap())

			require.Equal(t, fmt.Sprintf("%s", tc.err), fmt.Sprintf("%s", wcode))
			require.Equal(t, fmt.Sprintf("%v", tc.err), fmt.Sprintf("%v", wcode))
			require.Equal(t,
				fmt.Sprintf("stderror\n(1) code: %s\nWraps: (2) stderror\nError types: (1) *errors.withCode (2) *errors.errorString",
					code.String()),
				fmt.Sprintf("%+v", wcode))
		})
	}
}
