package log

import (
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	D("num: %v", 1)
	I("string: %v", "fake str")
	W("bool: %v", true)
	E("time: %v", time.Now())
}
