package main

import (
	"fmt"
	"os"
)

func Hello(target string) string {
	return fmt.Sprintf("Hello, %s", target)
}

func main() {
	target := "World"
	if len(os.Args) >= 2 {
		target = os.Args[1]
	}

	fmt.Println(Hello(target))
}

