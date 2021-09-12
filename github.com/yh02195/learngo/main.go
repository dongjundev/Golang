package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	defer fmt.Println("I'm done")
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	if koreaAge := age + 2; koreaAge < 18 {
		return false
	}
	return true
}

func main() {
	// length, upperName := lenAndUpper("dongjun")
	// fmt.Println(length, upperName)
	//repeatMe("jun", "ho", "woong", "jin")
	fmt.Println(superAdd(1, 2, 3, 4, 5))
}
