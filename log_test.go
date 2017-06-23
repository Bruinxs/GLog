package log_test

import (
	"os"
	"testing"
	"time"

	"github.com/bruinxs/log"
)

func TestLog(t *testing.T) {
	log.D("num: %v", 1)
	log.I("string: %v", "fake str")
	log.W("bool: %v", true)
	log.E("time: %v", time.Now())
}

func BenchmarkDateOutPut(b *testing.B) {
	tmp := "/tmp/benchmark_dir/"
	filename := tmp + "output-%v.log"
	content := "abcdefghijklmnopqrstuvwsyz"
	log.DateOutPut(filename)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.D("%v", content)
			log.I("%v", content)
			log.W("%v", content)
			log.E("%v", content)
		}
	})
	os.RemoveAll(tmp)
}
