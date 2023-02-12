package main

import (
	"fmt"
	"sync"
)

type SharedVariable struct {
	value int
	r1_chn chan string
	r2_chn chan string
	finish chan string
}

func rountine_one(s *SharedVariable, wg *sync.WaitGroup) {
	defer wg.Done()

    for msg_from_two := range s.r1_chn {
		fmt.Println(msg_from_two)

		if s.value > 5 {
			s.r2_chn <- "go ahead routine two !!"
			break
		}

		s.value += 1

		fmt.Println(fmt.Sprintf("    rountine one incremented value to %d  \n", s.value))

		s.r2_chn <- "go ahead routine two !!"
	}

	close(s.r2_chn)
	fmt.Println("routine one completed")
	
}

func rountine_two(s *SharedVariable, wg *sync.WaitGroup) {
	defer wg.Done()
	s.r1_chn <- "go ahead routine one !!"

    for msg_from_one := range s.r2_chn {
		fmt.Println(msg_from_one)
		
		if s.value > 5 {
			break
		}

		s.value += 1

		fmt.Println(fmt.Sprintf("    rountine two incremented value to %d  \n", s.value))

		s.r1_chn <- "go ahead routine one !!"
	}

	close(s.r1_chn)
	fmt.Println("routine two completed")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	s := SharedVariable {
		r1_chn : make(chan string),
		r2_chn : make(chan string),
		value : 1,
	}

	go rountine_one(&s, &wg)
	go rountine_two(&s, &wg)

	fmt.Println("waiting....")
	wg.Wait()
	fmt.Println("ending now....")
}
