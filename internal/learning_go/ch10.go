package learning_go

// Goroutines

func someProcess(val int) int {
	return 1
}

func runThingConcurrently(in <-chan int, out chan<- int) {
	go func() {
		for val := range in {
			result := someProcess(val)
			out <- result
		}
	}()
}
