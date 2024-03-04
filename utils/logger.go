/*
The realize of logger make other file easy to use.
There is initialization of logger instance.
It save the logs into local when connection to influxdb fail. when connection of influx recover, it will write back all logs.
Supply all function to log including debug, info, warn, error, fatal, panic.
*/

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

// InitLogger is to initialize the logger make log be available
func InitLogger() {
	// setting of the logger about size, backup, store days etc.
	logFile := &lumberjack.Logger{
		Filename:   "logs/server.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   true,
	}

	// new a local and influxdb logger instance
	localLogger := zerolog.New(logFile).With().Timestamp().Logger()
	influxdbLogger := zerolog.Logger{}

	// make logger equal to local logger at first
	logger = localLogger

	// set using influxdb mode to false at first
	isUsingInfluxDB := false

	// start to try to connect influxdb
	go func() {
		// get the context and close when program turn off
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// make an infinite loop check and make connection to influxdb
		for {
			select {
			case <-ctx.Done():
				return
			default:
				// try to connect influxdb. if success to connect then replace logger else then turn back to local logger.
				influxClient := NewInfluxDBClient()
				if result, err := influxDBClient.Client.Ping(ctx); result == true {
					// if it is a new connection, it will write all formal logs to the influxdb.
					if !isUsingInfluxDB {
						logger = influxdbLogger
						logger = logger.Hook(InfluxDBHook{Client: influxClient})
						isUsingInfluxDB = true
						uploadRotatedLogsToInfluxDB(influxClient.Client)
					}
				} else {
					// turn back to local logger and log error to connect influxdb
					if isUsingInfluxDB {
						logger = localLogger
						isUsingInfluxDB = false
					}
					LogError(err.Error())
				}
				// sleep a minute and then try to connect again.
				time.Sleep(time.Minute)
			}
		}
	}()
}

// the function is to upload all local logs into influx db
func uploadRotatedLogsToInfluxDB(influxClient influxdb2.Client) {
	// get the organization and bucket from environment
	org := EnvInfluxDbOrg()
	bucket := EnvInfluxDbBucket()

	// make a write api to write the logs
	writeAPI := influxClient.WriteAPIBlocking(org, bucket)

	// get all local log files
	files, err := filepath.Glob("logs/server.log*")
	if err != nil {
		LogError(err.Error())
	}

	// enumerate all files
	for _, filename := range files {
		// try to open the file. when open file fail, log error and continue.
		file, err := os.Open(filename)
		if err != nil {
			LogError(err.Error())
			continue
		}

		// new a scanner to read file
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			// get the line of content and get the level, time and message
			line := scanner.Text()
			level := getField(line, "level")
			logTime := getField(line, "time")
			logMessage := getField(line, "message")

			// write the data back to the influxdb and it will log error when write error
			point := influxdb2.NewPointWithMeasurement("logs").
				AddTag("level", level).
				AddField("message", logMessage).
				SetTime(parseTime(logTime))
			err := writeAPI.WritePoint(context.Background(), point)
			if err != nil {
				LogError(err.Error())
			}
		}

		// close the log file. when fail to close it, program will log error.
		err = file.Close()
		if err != nil {
			LogError(err.Error())
			continue
		}
		// remove the log file. when fail to remove it, program will log error.
		err = os.Remove(filename)
		if err != nil {
			LogError(err.Error())
		}
	}
}

// get the specific content in the log file's content
func getField(line, field string) string {
	// when field not exist
	if !strings.Contains(line, field) {
		return ""
	}

	// get the start of the specific tag
	start := strings.Index(line, `"`+field+`":"`) + len(field) + 4

	// to process special tag case (the end of the line).
	if field == "message" {
		return line[start : len(line)-2]
	}

	// get the end of the tag
	end := strings.Index(line[start:], `","`)

	return line[start : start+end]
}

// parse the time in the log file's content
func parseTime(logTime string) time.Time {
	// parse the time when fail then return time in now
	parsedTime, err := time.Parse(time.RFC3339, logTime)
	if err != nil {
		return time.Now()
	}
	return parsedTime
}

/*
The following functions make an interface to log the information.
It will be different by the importance of the message using different function.
*/

// LogDebug is to log information when development
func LogDebug(msg string) {
	logger.Debug().Msg(msg)
}

// LogInfo is to log basic information
func LogInfo(msg string) {
	logger.Info().Msg(msg)
}

// LogWarn is to log warning information need to focus
func LogWarn(msg string) {
	logger.Warn().Msg(msg)
}

// LogError is to log error information when unknown error happened
func LogError(msg string) {
	logger.Error().Msg(msg)
}

// LogFatal is to log fatal information when the error will influence server work
func LogFatal(msg string) {
	logger.Fatal().Msg(msg)
}

// LogPanic is to log panic information when the error is more serious than fatal
func LogPanic(msg string) {
	logger.Panic().Msg(msg)
}
