package main

import (
	"fmt"
	"time"
)

// ForSelectLoopExample1 demonstrates the use of a for-select loop to send data to a channel with a default case.
func ForSelectLoopExample1() {
	fmt.Println("For Select Loop")

	// record start time
	start := time.Now()

	channel1 := make(chan string, 3)
	chars := []string{"A", "B", "C", "D", "E"}

	for _, s := range chars {
		select {
		case channel1 <- s:
			fmt.Println("Sent:", s)
		default:
			fmt.Println("Channel full, could not send:", s)
		}
	}

	close(channel1)
	for msg := range channel1 {
		fmt.Println("Received:", msg)
	}

	elapsed := time.Since(start)
	fmt.Printf("Elapsed Time: %s\n", elapsed)
}