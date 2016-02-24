package main

import (
	"fmt"
	"github.com/miolini/fastrand"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fastrand.FastRand(64))
	}
}
