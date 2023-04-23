package main

import (
	"fmt"

	"github.com/OlePaulsen/fib/fib"
)

func main() {
	fib10 := fib.Fib(10)
	fmt.Println(fib10)
}
