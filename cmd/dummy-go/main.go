package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("dummy go started")

	for {
		time.Sleep(10 * time.Second)
		fmt.Println("dummy go running now")
	}
}
