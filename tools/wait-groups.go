package main

import (
	"fmt"
	"sync"
	"time"
)
const workingDuration = 24

func longRunningProcess(i int, wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(time.Duration(workingDuration) * time.Millisecond)
	fmt.Printf("nr, %d\n", i)
}


func main(){
	fmt.Println("Concurrency Patterns in Golang")
	
	wg := &sync.WaitGroup{}

	for i :=0; i< 10; i++ {
		wg.Add(1)
		go longRunningProcess(i, wg)
	}
}

