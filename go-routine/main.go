package main

import "fmt"

// import (
// 	"fmt"
// );

// func main() {
// 	fmt.Println("Go Routines")

// 	channel1 := make(chan string)
// 	channel2 := make(chan string)

// 	go func() {
// 		channel1 <- "Message from Channel 1"
// 	}()

// 	go func() {
// 		channel2 <- "Message from Channel 2"
// 	}()

// 	select {
// 	case msg1 := <-channel1:
// 		fmt.Println(msg1)

// 	case msg2 := <-channel2:
// 		fmt.Println(msg2)
// 	}

// }

func main() {

	done := make(chan bool, 2)

	go func() {
			for i := 0; i < 5; i++ {
					done <- true // safe: buffer exists
			}
	}()

	for i := 0; i < 5; i++ {
			<-done
			fmt.Println("Received signal", i)
	}
}