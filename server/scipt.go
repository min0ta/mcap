package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(i)
		time.Sleep(3 * time.Second)
	}
}
