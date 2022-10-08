package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1000 * time.Millisecond)
	count := 1
	for range tick {
		fmt.Println("Tick " + fmt.Sprint(count))
		count++
	}
}
