package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 20; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println("message", i)
	}
}
