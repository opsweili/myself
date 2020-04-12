package test

import (
	"fmt"
	"testing"
)

// 自定义类型
// type mystring string

// 类型别名
type str = string

// func (m *mystring) say() {
// 	fmt.Println(
// 		"mystring type obj have say method")
// }

func TestType(t *testing.T) {
	// var m mystring
	// m.say()

	var s str
	s = "wei"
	fmt.Println(s)

}
