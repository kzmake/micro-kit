package logger

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func ExampleLogger() {
	l := New(
		WithOutput(os.Stdout),
		WithTimeFormat(time.RFC3339Nano),
		WithLevel(InfoLevel),
		WithSkipFrameCount(3),
	).WithFields(map[string]interface{}{
		"request_id": "request_id",
		"user":       "alice",
	})

	l.Debugf("%s", "string")
	l.Infof("%s", "string")
	l.Warnf("%s", "string")
	l.Errorf("%s", "string")

	// Output:
	// {"level":"info","request_id":"request_id","user":"alice","time":"0001-01-01T00:00:00Z","caller":"file:123","message":"string"}
	// {"level":"warn","request_id":"request_id","user":"alice","time":"0001-01-01T00:00:00Z","caller":"file:123","message":"string"}
	// {"level":"error","request_id":"request_id","user":"alice","time":"0001-01-01T00:00:00Z","caller":"file:123","message":"string"}
}

func TestMain(m *testing.M) {
	zerolog.TimestampFunc = func() time.Time { return time.Time{} }
	zerolog.CallerMarshalFunc = func(_ string, _ int) string { return "file:123" }

	code := m.Run()
	os.Exit(code)
}

func TestNew(t *testing.T) {
	logger := New(
		WithOutput(os.Stdout),
		WithTimeFormat(time.RFC3339Nano),
		WithLevel(TraceLevel),
		WithFields(map[string]interface{}{
			"hogehoge": "fugafuga",
		}),
	)
	require.NotNil(t, logger)
}

func TestWithFields(t *testing.T) {
	out := &bytes.Buffer{}

	l := New(WithOutput(out)).WithFields(map[string]interface{}{
		"request_id": "request_id",
		"user":       "alice",
	})
	l.Infof("%s", "string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"info", "message":"string", "request_id":"request_id", "time":"0001-01-01T00:00:00Z", "user":"alice"}`,
		out.String(),
	)
}

func TestNoop(t *testing.T) {
	out := &bytes.Buffer{}

	l := New(WithOutput(out), WithLevel(Disabled))
	l.Tracef("%s", "string")
	l.Debugf("%s", "string")
	l.Infof("%s", "string")
	l.Warnf("%s", "string")
	l.Errorf("%s", "string")

	require.Empty(t, out.String())
}

func TestTracef(t *testing.T) {
	out := &bytes.Buffer{}

	l := New(WithOutput(out))
	l.Tracef("%s", "string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"trace", "message":"string", "time":"0001-01-01T00:00:00Z"}`,
		out.String(),
	)
}

func TestTrace(t *testing.T) {
	out := &bytes.Buffer{}

	l := New(WithOutput(out))
	l.Trace("string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"trace", "message":"string", "time":"0001-01-01T00:00:00Z"}`,
		out.String(),
	)
}

func TestDebugf(t *testing.T) {
	out := &bytes.Buffer{}

	l := New(WithOutput(out))
	l.Debugf("%s", "string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"debug", "message":"string", "time":"0001-01-01T00:00:00Z"}`,
		out.String(),
	)
}

func TestDebug(t *testing.T) {
	out := &bytes.Buffer{}

	l := New(WithOutput(out))
	l.Debug("string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"debug", "message":"string", "time":"0001-01-01T00:00:00Z"}`,
		out.String(),
	)
}

func TestInfof(t *testing.T) {
	out := &bytes.Buffer{}

	l := New(WithOutput(out))
	l.Infof("%s", "string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"info", "message":"string", "time":"0001-01-01T00:00:00Z"}`,
		out.String(),
	)
}

func TestInfo(t *testing.T) {
	out := &bytes.Buffer{}

	l := New(WithOutput(out))
	l.Info("string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"info", "message":"string", "time":"0001-01-01T00:00:00Z"}`,
		out.String(),
	)
}

func TestWarnf(t *testing.T) {
	out := &bytes.Buffer{}

	l := New(WithOutput(out))
	l.Warnf("%s", "string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"warn", "message":"string", "time":"0001-01-01T00:00:00Z"}`,
		out.String(),
	)
}

func TestWarn(t *testing.T) {
	out := &bytes.Buffer{}

	l := New(WithOutput(out))
	l.Warn("string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"warn", "message":"string", "time":"0001-01-01T00:00:00Z"}`,
		out.String(),
	)
}

func TestErrorf(t *testing.T) {
	out := &bytes.Buffer{}

	l := New(WithOutput(out))
	l.Errorf("%s", "string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"error", "message":"string", "time":"0001-01-01T00:00:00Z"}`,
		out.String(),
	)
}

func TestError(t *testing.T) {
	out := &bytes.Buffer{}

	l := New(WithOutput(out))
	l.Error("string")

	require.JSONEq(t,
		`{"caller":"file:123", "level":"error", "message":"string", "time":"0001-01-01T00:00:00Z"}`,
		out.String(),
	)
}
