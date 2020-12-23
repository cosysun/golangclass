package main

import (
	"fmt"
	"golangclass/chapter1/1.4/hellworld/helloworld"
)

func main() {
	// helloWorldPrint("golang")
	helloworld.Print("golang")
}

func helloWorldPrint(name string) {
	fmt.Printf("hello world, %s!\n", name)
}
