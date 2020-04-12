package test

import (
	"os"
	"testing"

	"ssic.cn/ops/spider/util/log"
)

// create null file
func file0() {
	fileName := "file/f1.txt"

	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		log.Debug("create file error")

	}
	log.Info("file create success")
}

func TestFile(t *testing.T) {

	// set log level
	log.SetLevel("info")

	file0()

}
