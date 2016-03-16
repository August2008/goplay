package channels

import (
	"fmt"
	"time"
)

func Ping3(ch chan string, msg string) {
	ch <- msg
}

func Play3() {
	//ctx := Context{""}
	ping := make(chan string)
	pong := make(chan string)

	go Ping3(pong, "Ping Start")
	go Ping3(ping, "Pong Start")

	go Receive3(ping, pong)
}

func Receive3(ping chan string, pong chan string) {
	for {
		select {
		case msg1 := <-pong:
			time.Sleep(time.Second)
			go Ping3(pong, "Ping")
			fmt.Println("Selector 1", msg1)

		case msg2 := <-ping:
			time.Sleep(time.Second)
			go Ping3(ping, "Pong")
			fmt.Println("Selector 2", msg2)

		default:
			fmt.Println("Default")
		}
	}
	fmt.Println("Received")
}
