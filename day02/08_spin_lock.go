package main

import (
	"fmt"
	"sync/atomic"
	"runtime"
	"time"
)

type spinLock uint32

func (sl *spinLock) Lock() {
    for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
        runtime.Gosched() //without this it locks up on GOMAXPROCS > 1
    }
}

func (sl *spinLock) Unlock() {
    atomic.StoreUint32((*uint32)(sl), 0)
}

func SpinLock() *spinLock {
	return new(spinLock)
}

func producer(sl *spinLock) {
	sl.Lock()
	
	fmt.Println("Producer has the lock")
	time.Sleep(time.Duration(1000) * time.Millisecond)

    sl.Unlock()

	fmt.Println("Producer unlocked")
}

func consumer(sl *spinLock) {
	time.Sleep(time.Duration(200) * time.Millisecond)

	sl.Lock()
	
	fmt.Println("Consumer has the lock")

	time.Sleep(time.Duration(1000) * time.Millisecond)

    sl.Unlock()

	fmt.Println("Consumer unlocked")
}


func main() {
	sLock := SpinLock()

	go producer(sLock)
	go consumer(sLock)

	var input string
	fmt.Scanln(&input)
}
