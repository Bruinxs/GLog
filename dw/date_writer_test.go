package dw

import (
	"bufio"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestDateWriter_Write(t *testing.T) {
	fn := "/tmp/dw_test/dw.test"
	dw := NewDateWriter(fn)
	if dw == nil {
		t.Errorf("resutn dw(%v) is nil", dw)
		return
	}

	num := 10
	for i := 0; i < num; i++ {
		dw.Write([]byte(strconv.Itoa(i) + "\n"))
	}
	dw.Close()

	df := time.Now().Format("2016-01-02")
	file, err := os.OpenFile("/tmp/dw_test/dw."+df+".test", os.O_RDONLY, 0666)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		file.Close()
		err = os.RemoveAll("/tmp/dw_test")
		if err != nil {
			t.Error(err)
		}
	}()

	scan := bufio.NewScanner(file)
	for i := 0; i < num; i++ {
		if !scan.Scan() {
			t.Errorf("index(%v), scan next line err", i)
			return
		}
		text := scan.Text()
		if text != strconv.Itoa(i) {
			t.Errorf("index(%v), scan text(%v) not equal (%v)", i, text, i)
			return
		}
	}
}
