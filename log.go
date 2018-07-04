package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

// enum log level
const (
	ERROR Level = 1 << iota
	WARN
	INFO
	DEBUG
)

var _log *Logger

func init() {
	_log = &Logger{
		level:  DEBUG,
		Logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	}
}

// Level log level
type Level int

func (level Level) String() string {
	switch level {
	case ERROR:
		return "[error]"
	case WARN:
		return "[warn ]"
	case INFO:
		return "[info ]"
	case DEBUG:
		return "[debug]"
	}
	return ""
}

// Logger logger with level
type Logger struct {
	*log.Logger
	level Level
}

func (l *Logger) log(calldepth int, level Level, s string) {
	if l.level < level {
		return
	}
	err := l.Output(calldepth, fmt.Sprintf("%s %s", level, s))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

//SetLevel set log level
func (l *Logger) SetLevel(level Level) {
	l.level = level
}

//SetLevel set defalut log level
func SetLevel(level Level) {
	_log.SetLevel(level)
}

// SetFlags sets the out put flag
func SetFlags(flag int) {
	_log.SetFlags(flag)
}

// SetOutput sets the output destination
func SetOutput(w io.Writer) {
	_log.SetOutput(w)
}

// Error output with ERROR level
func Error(v ...interface{}) {
	_log.log(3, ERROR, fmt.Sprint(v...))
}

// DepthError set the call depth and output with ERROR level
func DepthError(calldepth int, v ...interface{}) {
	_log.log(calldepth, ERROR, fmt.Sprint(v...))
}

// Errorf output with ERROR level, Arguments are handled in the manner of fmt.Sprintf
func Errorf(format string, v ...interface{}) {
	_log.log(3, ERROR, fmt.Sprintf(format, v...))
}

// Warn output with WARN level
func Warn(v ...interface{}) {
	_log.log(3, WARN, fmt.Sprint(v...))
}

// DepthWarn set the call depth and output with WARN level
func DepthWarn(calldepth int, v ...interface{}) {
	_log.log(calldepth, WARN, fmt.Sprint(v...))
}

// Warnf output with WARN level, Arguments are handled in the manner of fmt.Sprintf
func Warnf(format string, v ...interface{}) {
	_log.log(3, WARN, fmt.Sprintf(format, v...))
}

// Info output with INFO level
func Info(v ...interface{}) {
	_log.log(3, INFO, fmt.Sprint(v...))
}

// DepthInfo set the call depth and output with INFO level
func DepthInfo(calldepth int, v ...interface{}) {
	_log.log(calldepth, INFO, fmt.Sprint(v...))
}

// Infof output with INFO level, Arguments are handled in the manner of fmt.Sprintf
func Infof(format string, v ...interface{}) {
	_log.log(3, INFO, fmt.Sprintf(format, v...))
}

// Debug output with DEBUG level
func Debug(v ...interface{}) {
	_log.log(3, DEBUG, fmt.Sprint(v...))
}

// DepthDebug set the call depth and output with DEBUG level
func DepthDebug(calldepth int, v ...interface{}) {
	_log.log(calldepth, DEBUG, fmt.Sprint(v...))
}

// Debugf output with DEBUG level, Arguments are handled in the manner of fmt.Sprintf
func Debugf(format string, v ...interface{}) {
	_log.log(3, DEBUG, fmt.Sprintf(format, v...))
}
