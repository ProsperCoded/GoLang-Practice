package main

import (
	"fmt"
	"time"
);

func ForSelectLoop2() {
	go func(){
		for {
			select {
			default: fmt.Println("working...")
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("done")
}