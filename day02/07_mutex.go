package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type SharedVariable struct {
    lock sync.Mutex
	value int
}

func (s *SharedVariable) Increment() {
	(*s).lock.Lock()
	time.Sleep(time.Duration(200) * time.Millisecond)
	(*s).value++
	(*s).lock.Unlock()
}

func process_shared_variable (v *SharedVariable) {
	defer wg.Done()
	for i:=0 ; i<5; i++ {
		v.Increment()
		fmt.Print(v.value, "  ")
	}
}

func main() {
	wg.Add(3)
	s := SharedVariable{}

	for t:=0 ; t<3 ; t++ {
		go process_shared_variable(&s)
	}

	fmt.Println("waiting....")
	wg.Wait()
	fmt.Println("ending now....")
}
