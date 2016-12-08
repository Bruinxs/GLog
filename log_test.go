package log

import (
	"os"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	D("num: %v", 1)
	I("string: %v", "fake str")
	W("bool: %v", true)
	E("time: %v", time.Now())
}

func BenchmarkDateOutPut(b *testing.B) {
	tmp := "/tmp/benchmark_dir/"
	filename := tmp + "output-%v.log"
	content := "abcdefghijklmnopqrstuvwsyz"
	DateOutPut(filename)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			D("%v", content)
			I("%v", content)
			W("%v", content)
			E("%v", content)
		}
	})
	os.RemoveAll(tmp)
}
