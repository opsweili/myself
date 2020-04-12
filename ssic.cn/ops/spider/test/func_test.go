package test

import (
	"fmt"
	"testing"
)

func add(a, b int) int {
	return a + b
}

func TestFunction(t *testing.T) {
	fmt.Printf("%d \n", add(1, 2))
}
