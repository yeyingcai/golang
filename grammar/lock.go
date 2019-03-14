package main

import (
	"sync"
	"time"
)

//使用协程使用互斥锁，避免竞争
var lock sync.Mutex
var res int
var total = make(map[int]int, 10)

func main() {
	for x := 1; x <= 200; x++ {
		go test(x)
	}

	time.Sleep(10 * time.Second)
}

func test(n int) {
	for i := 1; i < n; i++ {
		res *= i
	}
	lock.Lock()
	total[n] = res
	lock.Unlock()
}
