package main

import (
	"github.com/fisache/unittestable-example/hello"
)

func helloFunc() string {
	return hello.SayHello("inki")
}

func main() {
	hello.New(&hello.Hello{
		LastName: "hwang",
	})

	helloFunc()
}
