package main

import (
	"fmt"
	"time"
)

func main() {
	channelTimer()
}

func channelTimer() {
	// in 5 seconds
	t1 := time.NewTimer(5 * time.Second)
	defer t1.Stop()
	// time 1 seconds
	t2 := time.NewTicker(1 * time.Second)
	defer t2.Stop()
	var count int
	for {
		select {
		case <-t1.C:
			fmt.Println("Finished")
			return
		case <-t2.C:
			fmt.Println(count)
			count++
		default:
		}
	}

}
