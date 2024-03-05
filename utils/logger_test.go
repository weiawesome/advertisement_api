package utils

import (
	"bytes"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"gopkg.in/natefinch/lumberjack.v2"
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

const testMessage = "test message"

func TestGetField(t *testing.T) {
	for _, lc := range logCases {
		t.Run(lc.name, func(t *testing.T) {
			result := getField(lc.line, lc.field)
			assert.Equal(t, result, lc.expected)
		})
	}
}

func TestInitLogger(t *testing.T) {
	t.Run("Case right", func(t *testing.T) {
		InitLogger()
		assert.NotNil(t, logger)
	})
}

func TestUploadRotatedLogsToInfluxDB(t *testing.T) {
	t.Run("Case error", func(t *testing.T) {
		logFile := &lumberjack.Logger{
			Filename:   "logs/server.log",
			MaxSize:    10,
			MaxBackups: 3,
			MaxAge:     30,
			Compress:   true,
		}

		logger = zerolog.New(logFile).With().Timestamp().Logger()

		logger.Info().Msg(testMessage)

		hook := NewInfluxDBHook()

		uploadRotatedLogsToInfluxDB(hook.Client)

	})
	t.Run("Case right", func(t *testing.T) {
		logFile := &lumberjack.Logger{
			Filename:   "logs/server.log",
			MaxSize:    10,
			MaxBackups: 3,
			MaxAge:     30,
			Compress:   true,
		}

		logger = zerolog.New(logFile).With().Timestamp().Logger()

		logger.Info().Msg(testMessage)

		logger = zerolog.Logger{}

		hook := NewInfluxDBHook()

		uploadRotatedLogsToInfluxDB(hook.Client)

	})
	t.Run("Case error", func(t *testing.T) {
		logger = zerolog.Logger{}

		hook := NewInfluxDBHook()

		uploadRotatedLogsToInfluxDB(hook.Client)

	})
}

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

func TestLogPanic(t *testing.T) {
	t.Run("Case for panic log", func(t *testing.T) {
		var buf bytes.Buffer
		logger = zerolog.New(&buf)

		defer func() {
			if r := recover(); r != nil {
				if !strings.Contains(buf.String(), testMessage) {
					t.Errorf("LogPanic did not log the expected message")
				}
			}
		}()

		LogPanic(testMessage)
	})
}
