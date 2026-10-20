package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()

			// Thời điểm goroutine bắt đầu thực thi
			goroutineStart := time.Now()
			// Gửi dữ liệu qua channel
			ch <- n
			// Thời điểm gửi thành công
			elapsed := time.Since(goroutineStart)
			fmt.Printf(
				"Goroutine %d: sent value %d, elapsed %v\n",
				n, n, elapsed,
			)
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for value := range ch {
		fmt.Printf("Main received: %d\n", value)
	}

	fmt.Printf("Total elapsed: %v\n", time.Since(start))
}
