package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	// length, upperName := lenAndUpper("dongjun")
	// fmt.Println(length, upperName)
	repeatMe("jun", "ho", "woong", "jin")
}
