package channels

import (
	"fmt"
)

type RingBuffer struct {
    inputChannel  <-chan int
    outputChannel chan int
}

func NewRingBuffer(inputChannel <-chan int, outputChannel chan int) *RingBuffer {
    return &RingBuffer{inputChannel, outputChannel}
}

func (rb *RingBuffer) Run() {
    for v := range rb.inputChannel {
        select {
        case rb.outputChannel <- v:
        default:
            <-rb.outputChannel
            rb.outputChannel <- v
        }
    }
    close(rb.outputChannel)
}

func TestRingBuffer() {
	
    input := make(chan int) 
    output := make(chan int, 5)
    rb := NewRingBuffer(input, output)
    
    go rb.Run()

    for i := 0; i < 10; i++ {
        input <- i
    }

    close(input)

    for res := range output {
        fmt.Println(res)
    }	
}