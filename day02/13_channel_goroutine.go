package main

import (
	"fmt"
	"time"
)

func remainder(msg string, timeMillis int, stopper func() bool ) <-chan string {
   ch := make(chan string)

   go func () {
	  for {
		if ch != nil {
			time.Sleep(time.Duration(timeMillis) * time.Millisecond)
			ch <- msg

			if stopper() {
				close(ch)
				fmt.Println("closing remainder")
				break
			}
		  }
	  }
	  fmt.Println("closing remainder")
   }()

   return ch
}

func main() {
    stop_flag := false  


    remainder_channel := remainder("time to stretch", 2000, func () bool {
		return stop_flag
	})
     
    count := 0
	 
    for {
	  fmt.Println("remainder received ", <- remainder_channel)

	  if count == 10 {
		stop_flag = true
		break
	  }

	  count += 1
    }
}