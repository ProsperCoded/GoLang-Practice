package main

// Pipeline pattern is used to connect multiple stages of processing where the output of one stage is the input to the next stage.
func sliceToChannel(slice []int) chan int {
	ch := make(chan int)
	go func() {
		// ! blocks until 1 by 1 value is read from sq goroutine
		for _, v := range slice {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func sq(in <-chan int) chan int {
	out := make(chan int)
	go func() {
		// ! blocks until 1 by 1 value is sent by sliceToChannel goroutine
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func PiplelineChannel() {
	// input
	numbers := []int{1, 2, 3, 4, 5}
	// stage 1: slice to channel
	in := sliceToChannel(numbers)
	// stage 2: square
	out := sq(in)

	// stage 3: read from out channel
	for v := range out {
		println(v)
	}
}