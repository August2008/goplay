package channels

import "time"

func Ping(pong chan string) {
	pong <- "Ping"
	time.Sleep(time.Second)
	Pong(pong)
}

func Pong(ping chan string) {
	ping <- "Pong"
	time.Sleep(time.Second)
	Ping(ping)
}

func Play() {
	ch := make(chan string, 1)

	go Ping(ch)
	go Receive(ch)
}

func Receive(ch chan string) {
	for {
		println(<-ch)
	}
}
