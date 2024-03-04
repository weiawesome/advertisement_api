/*
The realize of logger make other file easy to use.
There is initialization of logger instance.
It save the logs into local when connection to influxdb fail. when connection of influx recover, it will write back all logs.
Supply all function to log including debug, info, warn, error, fatal, panic.
*/

package utils

import (
	"bytes"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

type LogCase struct {
	name     string
	line     string
	field    string
	expected string
}
type TimeCase struct {
	name     string
	logTime  string
	expected time.Time
}

var (
	logCases = []LogCase{
		{
			name:     "Case right - level",
			line:     `{"time":"2021-01-01T12:00:00Z","level":"info","message":"test message"}`,
			field:    "level",
			expected: "info",
		},
		{
			name:     "Case right - message",
			line:     `{"time":"2021-01-01T12:00:00Z","level":"info","message":"test message"}`,
			field:    "message",
			expected: "test message",
		},
		{
			name:     "Case error - not exist field",
			line:     `{"time":"2021-01-01T12:00:00Z","level":"info","message":"test message"}`,
			field:    "user",
			expected: "",
		},
	}
	timeRightCases = []TimeCase{
		{
			name:     "Case right",
			logTime:  "2024-03-04T12:00:00Z",
			expected: time.Date(2024, 3, 4, 12, 0, 0, 0, time.UTC),
		},
		{
			name:     "Case right",
			logTime:  "2024-12-31T00:00:00Z",
			expected: time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		},
	}
	timeErrorCases = []TimeCase{
		{
			name:    "Case error",
			logTime: "Case error",
		},
	}
)

// get the specific content in the log file's content
func TestGetField(t *testing.T) {
	for _, lc := range logCases {
		t.Run(lc.name, func(t *testing.T) {
			result := getField(lc.line, lc.field)
			assert.Equal(t, result, lc.expected)
		})
	}
}

// parse the time in the log file's content
func TestParseTime(t *testing.T) {
	for _, trc := range timeRightCases {
		t.Run(trc.name, func(t *testing.T) {
			result := parseTime(trc.logTime)
			assert.Equal(t, trc.expected, result)
		})
	}
	for _, tec := range timeErrorCases {
		t.Run(tec.name, func(t *testing.T) {
			result := parseTime(tec.logTime)
			assert.Equal(t, time.Now(), result)
		})
	}
}

const testMessage = "test message"

func TestLogDebug(t *testing.T) {
	t.Run("Case for debug log", func(t *testing.T) {
		var buf bytes.Buffer
		logger = zerolog.New(&buf)

		LogDebug(testMessage)

		if !strings.Contains(buf.String(), testMessage) {
			t.Errorf("LogDebug did not log the expected message")
		}
	})
}

// LogInfo is to log basic information
func TestLogInfo(t *testing.T) {
	t.Run("Case for info log", func(t *testing.T) {
		var buf bytes.Buffer
		logger = zerolog.New(&buf)

		LogInfo(testMessage)

		if !strings.Contains(buf.String(), testMessage) {
			t.Errorf("LogInfo did not log the expected message")
		}
	})
}

// LogWarn is to log warning information need to focus
func TestLogWarn(t *testing.T) {
	t.Run("Case for warn log", func(t *testing.T) {
		var buf bytes.Buffer
		logger = zerolog.New(&buf)

		LogWarn(testMessage)

		if !strings.Contains(buf.String(), testMessage) {
			t.Errorf("LogWarn did not log the expected message")
		}
	})
}

// LogError is to log error information when unknown error happened
func TestLogError(t *testing.T) {
	t.Run("Case for error log", func(t *testing.T) {
		var buf bytes.Buffer
		logger = zerolog.New(&buf)

		LogError(testMessage)

		if !strings.Contains(buf.String(), testMessage) {
			t.Errorf("LogError did not log the expected message")
		}
	})
}

// LogPanic is to log panic information when the error is more serious than fatal
func TestLogPanic(t *testing.T) {
	t.Run("Case for panic log", func(t *testing.T) {
		var buf bytes.Buffer
		logger = zerolog.New(&buf)

		defer func() {
			if r := recover(); r != nil {
				// Recover from panic to allow the test to continue and check the log output
				if !strings.Contains(buf.String(), testMessage) {
					t.Errorf("LogPanic did not log the expected message")
				}
			}
		}()

		LogPanic("test message") // 假設這會引發 panic
	})
}
