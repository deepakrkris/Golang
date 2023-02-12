package main

import (
	"fmt"
)

func main() {
	messenger := make(chan string, 2)

	messenger<-"hello"
	messenger<-"world"

	//close(messenger)

	fmt.Println(<-messenger)
	fmt.Println(<-messenger)

	/*for message := range messenger {
		fmt.Println(message)
	}*/
}
