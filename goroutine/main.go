package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Printf("task %d bat dau \n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("task %d ket thuc \n", id)
}

func main() {
	start := time.Now()
	go task(1)
	go task(2)
	go task(3)
	go task(4)
	time.Sleep(time.Second * 2)
	fmt.Printf("tong time hoan thanh::: %ds", time.Since(start))
}
