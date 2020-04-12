package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"ssic.cn/ops/spider/action"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:   "create",
			Usage:  "create --uid=x --username=y",
			Action: (&action.User{}).Create,
			Flags: []cli.Flag{
				cli.IntFlag{Name: "uid", Usage: "--uid"},
				cli.StringFlag{Name: "username", Usage: "--username"},
			},
		},
		
		{
			Name:   "delete",
			Usage:  "delete --uid=x --username=y",
			Action: (&action.User{}).Delete,
			Flags: []cli.Flag{
				cli.IntFlag{Name: "uid", Usage: "--uid"},
				cli.StringFlag{Name: "username", Usage: "--username"},
			},
		},
		
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Print("command error :" + err.Error())
	}
}