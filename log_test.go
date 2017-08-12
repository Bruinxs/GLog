package log_test

import (
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
