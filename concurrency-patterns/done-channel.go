package main

import (
	"fmt"
	"time"
)

// Done channel is used to signal the termination of a goroutine, mostly from parent to child goroutine.
func DoneChannel() {

	done := make( chan bool)
	go doWork(done)

	time.Sleep(5 * time.Second)
	fmt.Println("done")
}
func doWork(done <- chan bool){
for {
			select {
			case <- done: 
			// parent stops the child goroutine
				return;
			default: fmt.Println("working...")
			}
		}
}