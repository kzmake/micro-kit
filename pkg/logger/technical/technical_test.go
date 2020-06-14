package technical

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	"github.com/kzmake/micro-kit/pkg/logger"
)

func TestMain(m *testing.M) {
	zerolog.TimestampFunc = func() time.Time { return time.Time{} }
	zerolog.CallerMarshalFunc = func(_ string, _ int) string { return "file:123" }

	code := m.Run()
	os.Exit(code)
}

func TestWithFields(t *testing.T) {
	out := &bytes.Buffer{}
	Logger = logger.New(logger.WithOutput(out))

	WithFields(map[string]interface{}{
		"request_id": "request_id",
		"user":       "alice",
	}).Infof("%s", "string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"info", "message":"string", "request_id":"request_id", "time":"0001-01-01T00:00:00Z", "user":"alice"}`,
		out.String(),
	)
}

func TestTracef(t *testing.T) {
	out := &bytes.Buffer{}
	Logger = logger.New(logger.WithOutput(out))

	Tracef("%s", "string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"trace", "message":"string", "time":"0001-01-01T00:00:00Z"}`,
		out.String(),
	)
}

func TestTrace(t *testing.T) {
	out := &bytes.Buffer{}
	Logger = logger.New(logger.WithOutput(out))

	Trace("string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"trace", "message":"string", "time":"0001-01-01T00:00:00Z"}`,
		out.String(),
	)
}

func TestDebugf(t *testing.T) {
	out := &bytes.Buffer{}
	Logger = logger.New(logger.WithOutput(out))

	Debugf("%s", "string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"debug", "message":"string", "time":"0001-01-01T00:00:00Z"}`,
		out.String(),
	)
}

func TestDebug(t *testing.T) {
	out := &bytes.Buffer{}
	Logger = logger.New(logger.WithOutput(out))

	Debug("string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"debug", "message":"string", "time":"0001-01-01T00:00:00Z"}`,
		out.String(),
	)
}

func TestInfof(t *testing.T) {
	out := &bytes.Buffer{}
	Logger = logger.New(logger.WithOutput(out))

	Infof("%s", "string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"info", "message":"string", "time":"0001-01-01T00:00:00Z"}`,
		out.String(),
	)
}

func TestInfo(t *testing.T) {
	out := &bytes.Buffer{}
	Logger = logger.New(logger.WithOutput(out))

	Info("string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"info", "message":"string", "time":"0001-01-01T00:00:00Z"}`,
		out.String(),
	)
}

func TestWarnf(t *testing.T) {
	out := &bytes.Buffer{}
	Logger = logger.New(logger.WithOutput(out))

	Warnf("%s", "string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"warn", "message":"string", "time":"0001-01-01T00:00:00Z"}`,
		out.String(),
	)
}

func TestWarn(t *testing.T) {
	out := &bytes.Buffer{}
	Logger = logger.New(logger.WithOutput(out))

	Warn("string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"warn", "message":"string", "time":"0001-01-01T00:00:00Z"}`,
		out.String(),
	)
}

func TestErrorf(t *testing.T) {
	out := &bytes.Buffer{}
	Logger = logger.New(logger.WithOutput(out))

	Errorf("%s", "string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"error", "message":"string", "time":"0001-01-01T00:00:00Z"}`,
		out.String(),
	)
}

func TestError(t *testing.T) {
	out := &bytes.Buffer{}
	Logger = logger.New(logger.WithOutput(out))

	Error("string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"error", "message":"string", "time":"0001-01-01T00:00:00Z"}`,
		out.String(),
	)
}
