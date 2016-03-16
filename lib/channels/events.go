package channels

import (
	"fmt"
	"time"
)

type Button struct {
	eventListeners map[string][]chan string
} 

func MakeButton() *Button {
	var result = new(Button)
	result.eventListeners = make(map[string][]chan string)
	return result;
}

func (this *Button) AddEventListener(event string, listener chan string) {
	var listeners, ok = this.eventListeners[event]
	if ok {
		this.eventListeners[event] = append(listeners, listener)
	} else {
		this.eventListeners[event] = []chan string {listener}
	}
}

func (this *Button) RemoveEventListener(event string, removeListener chan string) {		
	var listeners, ok = this.eventListeners[event]
	if ok {
		for i, listener := range listeners {
			if listener == removeListener {
				this.eventListeners[event] = append(listeners[:i], listeners[i+1:]...)
				break
			}
		}
	}	
}

func (this *Button) TriggerEvent(event string) {
	var listeners, ok = this.eventListeners[event]
	if ok {
		for i, listener := range listeners {
			go func(i int, listener chan string) {
				msg := fmt.Sprintf("Triggered %s of %d @ %s", event, i, time.Now().Format(time.RFC3339))
				listener <- msg
			}(i, listener)
		}
	}	
}

func TestButton() {
	var listener1 = make(chan string)
	var listener2 = make(chan string)
	
	button := MakeButton()
	button.AddEventListener("click", listener1)
	button.AddEventListener("click", listener2)
		
	go func() {
		for {
			event, ok := <- listener1
			if ok {
				fmt.Printf("%s\n", event)				
			} else {
				fmt.Printf("No event")
			}
		}
	}()
	
	go func() {
		for {
			event, ok := <- listener2
			if ok {
				fmt.Printf("%s\n", event)				
			} else {
				fmt.Printf("No event")
			}
		}
	}()	
		
	go func() {
		for i:=0; i<25; i++ {
			time.Sleep(time.Second)
			button.TriggerEvent("click")
		}
	}()	
	
	time.Sleep(time.Second * 5)
	button.RemoveEventListener("click", listener1)	
}

