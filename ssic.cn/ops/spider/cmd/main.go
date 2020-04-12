package main

import (
	"fmt"
	"os"
	"os/exec"

	// "github.com/pkg/errors"
	"github.com/pkg/errors"
	"ssic.cn/ops/spider/util/log"
)

func createNullFile01() {
	// log.Debug("debug.......")
	// log.Info("info.........")
	// log.Warn("warn........")
	// log.Error("error.....")

	_, err := os.Open("aaa")
	if err != nil {
		log.Debug(
			fmt.Sprintf("%s",
				errors.WithMessage(err, "open file error")))

	}
	//fmt.Println(f)

}

func getCmdPath(s string) {
	cmd := exec.Command(s)

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	createNullFile01()
}

// func main() {

// setting log level, error级别记录文件，其他级别走控制台
//log.SetLevel("error")
//createNullFile01()

//cmd := exec.Command("tr", "a-z", "A-Z")

// 将字符串输入流对象赋值给命令对象输入流
//cmd.Stdin = strings.NewReader("some input")

//b, err := cmd.Output()

// if err != nil {
// 	fmt.Println(err)
// 	return
// }
//fmt.Println(string(b))

// var out bytes.Buffer
// cmd.Stdout = &out

// if err := cmd.Run(); err != nil {
// 	fmt.Println(err)
// 	return
// }
// fmt.Println(out.String())

// cmd := exec.Command("ls", "/")
// cmd.Stdout = os.Stdout
// cmd.Run()

// }
