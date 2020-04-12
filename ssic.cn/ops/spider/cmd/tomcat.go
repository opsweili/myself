package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	startprefix = "/data/tomcat/tomcat-"
	startsuffix = "/bin/catalina.sh start"
	endprefix   = `ps -ef | grep -v "grep"  | grep tomcat-`
	endsuffix   = " | awk '{print $2}'" + " | xargs kill -9"
)

func init() {
	flag.Usage = usage
}

// execShell demo
func execShell(s string) {
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())
}

func isAction(s string) bool {
	str := []string{"start", "stop", "restart"}
	flag := false
	for _, v := range str {
		if s == v {
			flag = true
		}
	}
	return flag
}

func isPort(s string) bool {
	portList := []string{"8080", "8081", "8090"}
	flag := false
	for _, v := range portList {
		if s == v {
			flag = true
		}
	}
	return flag
}

func acb(portList []string, controller string) {
	for _, port := range portList {
		if isPort(port) {
			switch controller {
			case "start":
				execShell(startprefix + port + startsuffix)
			case "stop":
				execShell(endprefix + port + endsuffix)
			case "restart":
				execShell(endprefix + port + endsuffix)
				execShell(startprefix + port + startsuffix)
			}
		} else {
			flag.Usage()
		}
	}
}

// Operate demo
func Operate(s []string) {
	if len(s)-1 < 2 {
		flag.Usage()
	} else {
		action := s[1]
		port := s[2:]
		if isAction(action) {
			acb(port, action)
		} else {
			flag.Usage()
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: tomcat ( commands ... )
commands:
  start		Start Tomcat in a separate window
  stop          Tomcat, waiting up to 5 seconds for the process to end
  restart       Front Stop Tomcat After Start Tomcat
example:           
  tomcat start|stop|restart port    port may as one or many
  tomcat start|stop|restart 8080
  tomcat start|stop|restart 8080 8081 8090
`)
	flag.PrintDefaults()
}

// func main() {
// 	flag.Parse()
// 	Operate(os.Args)

// }
