package config

import (
	"fmt"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/require"
)

func ExampleConfig() {
	type HogeConfig struct {
		Log struct {
			Out   string `default:"stdout"`
			Level string `default:"info"`
		}
		FugaService struct {
			Key string `required:"true"`
		} `envconfig:"FUGA_SERVICE"`
		ServiceDiscovery struct {
			Provider string `required:"true"`
			Endpoint string `required:"true"`
		} `envconfig:"SERVICE_DISCOVERY"`
	}

	os.Clearenv()
	os.Setenv("HOGE_LOG_OUT", "stderr")
	os.Setenv("HOGE_LOG_LEVEL", "debug")
	os.Setenv("HOGE_FUGA_SERVICE_KEY", "fuga")
	os.Setenv("HOGE_SERVICE_DISCOVERY_PROVIDER", "etcd")
	os.Setenv("HOGE_SERVICE_DISCOVERY_ENDPOINT", "127.0.0.1:2379")

	conf, err := New("HOGE", &HogeConfig{})
	if err != nil {
		fmt.Printf("failed: %+v", err)
	}

	conf, ok := conf.(*HogeConfig)
	if !ok {
		fmt.Printf("failed")
	}

	fmt.Printf("%+v", conf)

	// Output:
	// &{Log:{Out:stderr Level:debug} FugaService:{Key:fuga} ServiceDiscovery:{Provider:etcd Endpoint:127.0.0.1:2379}}
}

const prefix = "TEST_CONFIG"

func TestNew_WithBasicType(t *testing.T) {
	type BasicTypeConfig struct {
		Bool          bool
		Int           int
		Float32       float32
		String        string
		UnsignedInt32 uint32
		SomePointer   *string
	}

	want := &BasicTypeConfig{
		Bool:          true,
		Int:           1234,
		Float32:       3.14159265,
		String:        "string",
		UnsignedInt32: 8080,
		SomePointer:   func(s string) *string { return &s }("string pointer"),
	}

	os.Clearenv()
	os.Setenv(prefix+"_"+"BOOL", "true")
	os.Setenv(prefix+"_"+"INT", "1234")
	os.Setenv(prefix+"_"+"FLOAT32", "3.14159265")
	os.Setenv(prefix+"_"+"STRING", "string")
	os.Setenv(prefix+"_"+"UNSIGNEDINT32", "8080")
	os.Setenv(prefix+"_"+"SOMEPOINTER", "string pointer")

	conf, err := New(prefix, &BasicTypeConfig{})
	require.NoError(t, err)

	actual, ok := conf.(*BasicTypeConfig)
	require.True(t, ok)

	require.Equal(t, want, actual)
}

func TestNew_WithTime(t *testing.T) {
	type TimeConfig struct {
		Duration time.Duration
		Datetime time.Time
	}

	tt, _ := time.Parse(time.RFC3339, "2038-01-19T03:14:08Z")
	want := &TimeConfig{
		Duration: 7 * time.Minute,
		Datetime: tt,
	}

	os.Clearenv()
	os.Setenv(prefix+"_"+"DURATION", "7m")
	os.Setenv(prefix+"_"+"DATETIME", "2038-01-19T03:14:08Z")

	conf, err := New(prefix, &TimeConfig{})
	require.NoError(t, err)

	actual, ok := conf.(*TimeConfig)
	require.True(t, ok)

	require.Equal(t, want, actual)
}

func TestNew_WithMap(t *testing.T) {
	type MapConfig struct {
		MapData map[string]int
	}

	want := &MapConfig{
		MapData: map[string]int{
			"red":   1,
			"green": 2,
			"blue":  3,
		},
	}

	os.Clearenv()
	os.Setenv(prefix+"_"+"MAPDATA", "red:1,green:2,blue:3")

	conf, err := New(prefix, &MapConfig{})
	require.NoError(t, err)

	actual, ok := conf.(*MapConfig)
	require.True(t, ok)

	require.Equal(t, want, actual)
}

type CustomType struct {
	Value *url.URL
}

func (cu *CustomType) UnmarshalBinary(data []byte) error {
	u, err := url.Parse(string(data))
	cu.Value = u
	return err
}

func TestNew_WithCustomType(t *testing.T) {
	type CustomTypeConfig struct {
		CustomTypeValue   CustomType
		CustomTypePointer *CustomType
	}

	u, _ := url.Parse("https://example.com")
	want := &CustomTypeConfig{
		CustomTypeValue:   CustomType{Value: u},
		CustomTypePointer: &CustomType{Value: u},
	}

	os.Clearenv()
	os.Setenv(prefix+"_"+"CUSTOMTYPEVALUE", "https://example.com")
	os.Setenv(prefix+"_"+"CUSTOMTYPEPOINTER", "https://example.com")

	conf, err := New(prefix, &CustomTypeConfig{})
	require.NoError(t, err)

	actual, ok := conf.(*CustomTypeConfig)
	require.True(t, ok)

	require.Equal(t, want, actual)
}

func TestNew_WithEmbedded(t *testing.T) {
	type Embedded struct {
		Bool   bool
		Int    int
		String string
	}

	type EmbeddedConfig struct {
		Embedded
	}

	want := &EmbeddedConfig{
		Embedded{
			Bool:   true,
			Int:    1234,
			String: "embdded string",
		},
	}

	os.Clearenv()
	os.Setenv(prefix+"_"+"BOOL", "true")
	os.Setenv(prefix+"_"+"INT", "1234")
	os.Setenv(prefix+"_"+"STRING", "embdded string")

	conf, err := New(prefix, &EmbeddedConfig{})
	require.NoError(t, err)

	actual, ok := conf.(*EmbeddedConfig)
	require.True(t, ok)

	require.Equal(t, want, actual)
}

func TestNew_WithEmbeddedPointer(t *testing.T) {
	type Embedded struct {
		Bool   bool
		Int    int
		String string
	}

	type EmbeddedPointerConfig struct {
		*Embedded
	}

	want := &EmbeddedPointerConfig{
		&Embedded{
			Bool:   true,
			Int:    1234,
			String: "embdded string",
		},
	}

	os.Clearenv()
	os.Setenv(prefix+"_"+"BOOL", "true")
	os.Setenv(prefix+"_"+"INT", "1234")
	os.Setenv(prefix+"_"+"STRING", "embdded string")

	conf, err := New(prefix, &EmbeddedPointerConfig{})
	require.NoError(t, err)

	actual, ok := conf.(*EmbeddedPointerConfig)
	require.True(t, ok)

	require.Equal(t, want, actual)
}

func TestNew_WithNested(t *testing.T) {
	type NestConfig struct {
		NestedConfig struct {
			Property string `envconfig:"VAR"`
		} `envconfig:"NESTED"`
		AfterNested string
	}

	want := &NestConfig{
		NestedConfig: struct {
			Property string `envconfig:"VAR"`
		}{
			Property: "nested string",
		},
		AfterNested: "after nested string",
	}

	os.Clearenv()
	os.Setenv(prefix+"_"+"NESTED_VAR", "nested string")
	os.Setenv(prefix+"_"+"AFTERNESTED", "after nested string")

	conf, err := New(prefix, &NestConfig{})
	require.NoError(t, err)

	actual, ok := conf.(*NestConfig)
	require.True(t, ok)

	require.Equal(t, want, actual)
}

func TestNew_WithNestedPointer(t *testing.T) {
	type NestPointerConfig struct {
		NestedConfig *struct {
			Property string `envconfig:"VAR"`
		} `envconfig:"NESTED"`
		AfterNested string
	}

	want := &NestPointerConfig{
		NestedConfig: &struct {
			Property string `envconfig:"VAR"`
		}{
			Property: "nested string",
		},
		AfterNested: "after nested string",
	}

	os.Clearenv()
	os.Setenv(prefix+"_"+"NESTED_VAR", "nested string")
	os.Setenv(prefix+"_"+"AFTERNESTED", "after nested string")

	conf, err := New(prefix, &NestPointerConfig{})
	require.NoError(t, err)

	actual, ok := conf.(*NestPointerConfig)
	require.True(t, ok)

	require.Equal(t, want, actual)
}

func TestNew_WithTag(t *testing.T) {
	type TagConfig struct {
		NoPrefix string `envconfig:"NO_PREFIX"`
		Default  string `default:"default string"`
		Required string `required:"true"`
	}

	want := &TagConfig{
		NoPrefix: "no prefix string",
		Default:  "default string",
		Required: "required string",
	}

	os.Clearenv()
	os.Setenv("NO_PREFIX", "no prefix string")
	os.Unsetenv(prefix + "_" + "DEFAULT")
	os.Setenv(prefix+"_"+"REQUIRED", "required string")

	conf, err := New(prefix, &TagConfig{})
	require.NoError(t, err)

	actual, ok := conf.(*TagConfig)
	require.True(t, ok)

	require.Equal(t, want, actual)
}

func TestNew_WhenParseError(t *testing.T) {
	type BasicTypeConfig struct {
		Bool bool
	}

	os.Clearenv()
	os.Setenv(prefix+"_"+"BOOL", "yes")

	_, err := New(prefix, &BasicTypeConfig{})
	require.Error(t, err)
	require.IsType(t, &envconfig.ParseError{}, err)
}
