package main

import (
	"fmt"
)

type TestFuncType func(s string) string

func say(testFunc TestFuncType, s string) {
	fmt.Println(testFunc(s))
}

func test1(s string) string {
	return "hello " + s
}

func main() {
	say(test1, "word")
}
