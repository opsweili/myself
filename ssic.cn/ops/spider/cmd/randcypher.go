package main

import (
	"fmt"
	"math/rand"
)

var (
	length  int
	charset string
)

// NUmStr demo
const (
	NUmStr  = "0123456789"
	CharStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SpecStr = "+=-@#~,.[]()!%^*$"
)

func generatePasswd(charset string) string {

	var passwd []byte = make([]byte, length, length)
	var sourceStr string

	if charset == "num" {
		sourceStr = NUmStr

	} else if charset == "char" {
		sourceStr = CharStr

	} else if charset == "mix" {
		sourceStr = fmt.Sprintf("%s%s", NUmStr, CharStr)

	} else if charset == "advance" {
		// fmt.Println(charset)
		sourceStr = fmt.Sprintf("%s%s%s", NUmStr, CharStr, SpecStr)
	} else {
		sourceStr = NUmStr
	}

	for i := 0; i < length; i++ {
		index := rand.Intn(len(sourceStr))
		passwd[i] = sourceStr[index]
	}
	return string(passwd)
}

// func main() {

// 	app := cli.NewApp()
// 	app.Name = "randcypher"
// 	app.Usage = `
// 	  Usage: randcypher [-lt] [-l (0-16)] [-t num|char|mix|advance]
// 		-l length sheng cheng password length
// 		-t sheng cheng password charset
//     num: use number [0-9]
//     char: use charset [a-zA-Z]
//     mix: use number and charset
//     advance: use number and charset and special charset`

// 	app.Author = "weili"
// 	app.Email = "435053099@qq.com"
// 	app.Version = "0.0.1"

// 	app.Flags = []cli.Flag{
// 		cli.IntFlag{
// 			Name:        "length, l",
// 			Usage:       "sheng cheng password length",
// 			Value:       16,
// 			Destination: &length,
// 		},

// 		cli.StringFlag{
// 			Name:        "type, t",
// 			Usage:       "sheng cheng password charset",
// 			Value:       "mix",
// 			Destination: &charset,
// 		},
// 	}

// 	app.Action = func(c *cli.Context) error {
// 		// fmt.Println(length)
// 		// fmt.Println(charset)

// 		rand.Seed(time.Now().UnixNano())
// 		fmt.Println(generatePasswd(charset))
// 		return nil

// 	}

// 	_ = app.Run(os.Args)

// }
