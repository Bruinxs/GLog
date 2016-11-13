package dw

import (
	"fmt"
	"github.com/Bruinxs/util/gtime"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type DateWriter struct {
	dir   string
	bfn   string
	ext   string
	daily int
	file  *os.File
	lck   *sync.Mutex
}

func NewDateWriter(filename string) *DateWriter {
	filename = filepath.Clean(filename)
	dir, bfn := filepath.Split(filename)
	ext := filepath.Ext(bfn)
	if len(ext) == 0 {
		ext = ".log"
	} else {
		bfn = strings.TrimSuffix(bfn, ext)
	}

	dw := &DateWriter{
		dir:  dir,
		bfn:  bfn,
		ext:  ext,
		file: nil,
		lck:  &sync.Mutex{},
	}
	go dw.runSwitchFile()
	return dw
}

func (this *DateWriter) Write(p []byte) (n int, err error) {
	this.lck.Lock()
	defer this.lck.Unlock()
	if this.file == nil {
		err := this.createFile()
		if err != nil {
			return -1, err
		}
	}
	n, err = this.file.Write(p)
	return
}

func (this *DateWriter) Close() error {
	this.lck.Lock()
	defer this.lck.Unlock()
	if this.file != nil {
		err := this.file.Close()
		if err != nil {
			return err
		}
		this.file = nil
	}
	return nil
}

func (this *DateWriter) createDir() error {
	if len(this.dir) == 0 {
		return nil
	}
	fi, err := os.Lstat(this.dir)
	if err == nil {
		if fi.IsDir() {
			return nil
		} else {
			return fmt.Errorf("path(%v) exist but not a dir", this.dir)
		}
	}
	return os.MkdirAll(this.dir, os.ModeDir)
}

func (this *DateWriter) createFile() error {
	this.createDir()
	if this.file != nil {
		this.file.Close()
		this.file = nil
	}
	df := time.Now().Format("2016-01-02")
	fn := fmt.Sprintf("%v.%v%v", this.bfn, df, this.ext)
	file := filepath.Join(this.dir, fn)

	//find no exist file name
	_, err := os.Lstat(file)
	num := 1
	for ; err == nil; num++ {
		fn = fmt.Sprintf("%v.%v.%v%v", this.bfn, df, num, this.ext)
		file = filepath.Join(this.dir, fn)
		_, err = os.Lstat(file)
	}

	this.file, err = os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	return nil
}

func (this *DateWriter) runSwitchFile() {
	gtime.AlarmDay("00:00:00", -1, func(count int) bool {
		this.lck.Lock()
		this.createFile()
		this.lck.Unlock()
		return false
	})
}
