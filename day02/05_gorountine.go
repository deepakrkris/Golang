package main

import (
	"fmt"
	"time"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func say_hello (name string, wait_time int) {
	defer wg.Done()
	for i:=0 ; i<5; i++ {
		time.Sleep(time.Duration(wait_time) * time.Millisecond)
		fmt.Println("hello this is ", name)
	}
	fmt.Println("ending ", name)
}

func main() {
	wg.Add(3)
	for t:=0 ; t<3 ; t++ {
		go say_hello("routine no " + strconv.Itoa(t), 100 * 3)
	}

	fmt.Println("waiting....")
	wg.Wait()
	fmt.Println("ending now....")
}
