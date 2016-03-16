package channels

import (
	"fmt"
	"time"
)

func TestPrimes() {
	var input = make(chan int)
	go generate(input)
	for {
		var prime = <- input		
		fmt.Println(prime) 
		var output = make(chan int)
		go filter(input, output, prime)
		input = output	
	}
}

func generate(input chan int) { 
	for i:=2; ; i++ {
		input <- i	
		time.Sleep(time.Millisecond * 50)
	}
}

func filter(input, output chan int, prime int) {
	for {
		var i = <- input
		if i%prime != 0 {			
			output <- i
		} 
	}
}

