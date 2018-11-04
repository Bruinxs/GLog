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
	_log = NewLogger("global")
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
	level     Level
	spacename string
}

// NewLogger new logger instance
func NewLogger(spacename string) *Logger {
	return &Logger{
		level:     DEBUG,
		Logger:    log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		spacename: spacename,
	}
}

func (l *Logger) log(calldepth int, level Level, s string) {
	if l.level < level {
		return
	}
	err := l.Output(calldepth, fmt.Sprintf("%s[%s] %s", level, l.spacename, s))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

//SetLevel set log level
func (l *Logger) SetLevel(level Level) {
	l.level = level
}

// SetFlags sets the out put flag
func (l *Logger) SetFlags(flag int) {
	l.Logger.SetFlags(flag)
}

// SetOutput sets the output destination
func (l *Logger) SetOutput(w io.Writer) {
	l.Logger.SetOutput(w)
}

// Error output with ERROR level
func (l *Logger) Error(v ...interface{}) {
	l.log(3, ERROR, fmt.Sprint(v...))
}

// DepthError set the call depth and output with ERROR level
func (l *Logger) DepthError(calldepth int, v ...interface{}) {
	l.log(calldepth, ERROR, fmt.Sprint(v...))
}

// Errorf output with ERROR level, Arguments are handled in the manner of fmt.Sprintf
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.log(3, ERROR, fmt.Sprintf(format, v...))
}

// Warn output with WARN level
func (l *Logger) Warn(v ...interface{}) {
	l.log(3, WARN, fmt.Sprint(v...))
}

// DepthWarn set the call depth and output with WARN level
func (l *Logger) DepthWarn(calldepth int, v ...interface{}) {
	l.log(calldepth, WARN, fmt.Sprint(v...))
}

// Warnf output with WARN level, Arguments are handled in the manner of fmt.Sprintf
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.log(3, WARN, fmt.Sprintf(format, v...))
}

// Info output with INFO level
func (l *Logger) Info(v ...interface{}) {
	l.log(3, INFO, fmt.Sprint(v...))
}

// DepthInfo set the call depth and output with INFO level
func (l *Logger) DepthInfo(calldepth int, v ...interface{}) {
	l.log(calldepth, INFO, fmt.Sprint(v...))
}

// Infof output with INFO level, Arguments are handled in the manner of fmt.Sprintf
func (l *Logger) Infof(format string, v ...interface{}) {
	l.log(3, INFO, fmt.Sprintf(format, v...))
}

// Debug output with DEBUG level
func (l *Logger) Debug(v ...interface{}) {
	l.log(3, DEBUG, fmt.Sprint(v...))
}

// DepthDebug set the call depth and output with DEBUG level
func (l *Logger) DepthDebug(calldepth int, v ...interface{}) {
	l.log(calldepth, DEBUG, fmt.Sprint(v...))
}

// Debugf output with DEBUG level, Arguments are handled in the manner of fmt.Sprintf
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.log(3, DEBUG, fmt.Sprintf(format, v...))
}
