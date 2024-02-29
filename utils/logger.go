package utils

import (
	"bufio"
	"context"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var logger zerolog.Logger

func InitLogger() {
	logFile := &lumberjack.Logger{
		Filename:   "logs/server.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   true,
	}
	localLogger := zerolog.New(logFile).With().Timestamp().Logger()
	influxdbLogger := zerolog.Logger{}
	logger = localLogger
	isUsingInfluxDB := false

	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				influxClient := NewInfluxDBClient()
				if result, err := influxDBClient.Client.Ping(ctx); result == true {
					if !isUsingInfluxDB {
						logger = influxdbLogger
						logger = logger.Hook(InfluxDBHook{Client: influxClient})
						isUsingInfluxDB = true
						uploadRotatedLogsToInfluxDB(influxClient.Client)
					}
				} else {
					if isUsingInfluxDB {
						logger = localLogger
						isUsingInfluxDB = false
					}
					LogError(err.Error())
				}
				time.Sleep(time.Minute)
			}
		}
	}()

	time.Sleep(time.Second)
}

func uploadRotatedLogsToInfluxDB(influxClient influxdb2.Client) {
	org := EnvInfluxDbOrg()
	bucket := EnvInfluxDbBucket()

	writeAPI := influxClient.WriteAPIBlocking(org, bucket)

	files, err := filepath.Glob("logs/server.log*")
	if err != nil {
		LogError(err.Error())
	}

	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			LogError(err.Error())
			continue
		}
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			level := getField(line, "level")
			logTime := getField(line, "time")
			logMessage := getField(line, "message")
			point := influxdb2.NewPointWithMeasurement("logs").
				AddTag("level", level).
				AddField("message", logMessage).
				SetTime(parseTime(logTime))

			err := writeAPI.WritePoint(context.Background(), point)
			if err != nil {
				LogError(err.Error())
			}
		}
		err = file.Close()
		if err != nil {
			LogError(err.Error())
			continue
		}
		err = os.Remove(filename)
		if err != nil {
			LogError(err.Error())
		}
	}
}

func getField(line, field string) string {
	start := strings.Index(line, `"`+field+`":"`) + len(field) + 4
	if field == "message" {
		return line[start : len(line)-2]
	}
	end := strings.Index(line[start:], `","`)
	return line[start : start+end]
}

func parseTime(logTime string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, logTime)
	if err != nil {
		return time.Now()
	}
	return parsedTime
}

func LogDebug(msg string) {
	logger.Debug().Msg(msg)
}

func LogInfo(msg string) {
	logger.Info().Msg(msg)
}

func LogWarn(msg string) {
	logger.Warn().Msg(msg)
}

func LogError(msg string) {
	logger.Error().Msg(msg)
}

func LogFatal(msg string) {
	logger.Fatal().Msg(msg)
}

func LogPanic(msg string) {
	logger.Panic().Msg(msg)
}
