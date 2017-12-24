// Copyright 2017 John Dengis
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logger

import (
	"io"
	"log"
	"strconv"
)

// The prefix to use for all logs in this logger.
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

// A Logger with a logging level construct.
type leveledLogger struct {
	logger *log.Logger
	level  LogLevel
}

// The internal leveled Logger to use for this logging package.
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

// Debug writes icebox logs with the configured logger at the DEBUG level.
func Debug(v ...interface{}) {
	logIfValidLevel(DEBUG, v)
}

// Info writes icebox logs with the configured logger at the INFO level.
func Info(v ...interface{}) {
	logIfValidLevel(INFO, v)
}

// Warn writes icebox logs with the configured logger at the WARN level.
func Warn(v ...interface{}) {
	logIfValidLevel(WARN, v)
}

// Error writes icebox logs with the configured logger at the ERROR level.
func Error(v ...interface{}) {
	logIfValidLevel(ERROR, v)
}

// Writes a log if the given level is greater than or equal to the
// configured level.
func logIfValidLevel(level LogLevel, v ...interface{}) {
	if iceboxLogger.level <= level {
		iceboxLogger.logger.Println(level.String(), v)
	}
}
