package log

import (
	"fmt"
	"github.com/Bruinxs/log/dw"
	"io"
	"log"
	"os"
)

//log level
const (
	ERROR = iota
	WARNING
	INFO
	DEBUG
)

var defLog = NewLog(os.Stdout)

type Log struct {
	level int
	*log.Logger
}

func NewLog(w io.Writer) *Log {
	return &Log{
		level:  DEBUG,
		Logger: log.New(w, "", log.LstdFlags|log.Lshortfile),
	}
}

func (this *Log) Log(cd, level int, prefix, format string, v ...interface{}) error {
	if level > this.level {
		return nil
	}
	return this.Logger.Output(cd, fmt.Sprintf("%v %v", prefix, fmt.Sprintf(format, v...)))
}

func (this *Log) SetOutput(w io.Writer) {
	this.Logger.SetOutput(w)
}

func D(format string, v ...interface{}) {
	defLog.Log(3, DEBUG, "[D]", format, v...)
}

func I(format string, v ...interface{}) {
	defLog.Log(3, INFO, "[I]", format, v...)
}

func W(format string, v ...interface{}) {
	defLog.Log(3, WARNING, "[W]", format, v...)
}

func E(format string, v ...interface{}) {
	defLog.Log(3, ERROR, "[E]", format, v...)
}

func SetOutPut(filename string) {
	defLog.SetOutput(dw.NewDateWriter(filename))
}
