package log_test

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/bruinxs/log"
)

func TestLog(t *testing.T) {
	log.Debugf("num: %v", 1)
	log.Infof("string: %v", "fake str")
	log.Warnf("bool: %v", true)
	log.Errorf("time: %v", time.Now())
	log.Debug("num: ", 1)
	log.Info("string: ", "fake str")
	log.Warn("bool: ", true)
	log.Error("time: ", time.Now())
}

func TestLogger(t *testing.T) {
	logger := log.NewLogger("test")
	buf := bytes.NewBuffer(nil)
	logger.SetOutput(buf)

	line := codeLine()
	logger.Error("error")
	logger.Debug("debug")

	out := buf.String()
	for _, want := range []string{
		fmt.Sprintf("log_test.go:%d: [error][test] error\n", line+1),
		fmt.Sprintf("log_test.go:%d: [debug][test] debug", line+2),
	} {
		if !strings.Contains(out, want) {
			t.Fatal(out)
		}
	}
}

func codeLine() int {
	_, _, line, _ := runtime.Caller(1)
	return line
}
