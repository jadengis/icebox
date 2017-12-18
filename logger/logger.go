package logger

import (
	"io"
	"log"
	"strconv"
)

const logPrefix string = "Icebox::"

// LogLevel is an enumeration type of increasing logging severities.
// These levels can be used to limit logging output.
type LogLevel int

// Const declaration of LogLevels
const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

// String converts the given LogLevel into its string respresentation.
func (l LogLevel) String() string {
	if int(l) < len(logLevelNames) {
		return logLevelNames[l]
	}
	return "LogLevel" + strconv.Itoa(int(l))
}

// Mapping between LogLevels and string representations
var logLevelNames = []string{
	DEBUG: "DEBUG::",
	INFO:  "INFO::",
	WARN:  "WARN::",
	ERROR: "ERROR::",
}

type leveledLogger struct {
	logger *log.Logger
	level  LogLevel
}

var iceboxLogger *leveledLogger

// Init is a method for initializing this package for logging.
// The write location can be specified as any io.Writer through the out parameter.
// The flag specifies the included log information, and you can limit the minimum
// logging level with the level parameter.
func Init(out io.Writer, flag int, level LogLevel) {
	logger := log.New(out, logPrefix, flag)
	iceboxLogger = &leveledLogger{
		logger: logger,
		level:  level,
	}
}

func Debug(v ...interface{}) {
	logIfValidLevel(DEBUG, v)
}

func Info(v ...interface{}) {
	logIfValidLevel(INFO, v)
}

func Warn(v ...interface{}) {
	logIfValidLevel(WARN, v)
}

func Error(v ...interface{}) {
	logIfValidLevel(ERROR, v)
}

//
func logIfValidLevel(level LogLevel, v ...interface{}) {
	if iceboxLogger.level >= level {
		iceboxLogger.logger.Println(level.String(), v)
	}
}
