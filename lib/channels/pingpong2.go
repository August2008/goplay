package channels

import "time"

func Ping2(pong chan string) {
	pong <- "Ping"
}

func Pong2(ping chan string) {
	ping <- "Pong"
}

func Play2() {
	//ctx := Context{""}
	ping := make(chan string)
	pong := make(chan string)

	go Ping2(pong)

	go ReceivePong2(pong, ping)
	go ReceivePing2(ping, pong)
}

func ReceivePing2(in chan string, out chan string) {
	for m := range in {
		println(m)
		Pong2(out)
		time.Sleep(time.Second)
	}
}

func ReceivePong2(in chan string, out chan string) {
	for m := range in {
		println(m)
		Ping2(out)
		time.Sleep(time.Second)
	}
}
