package channels

import (
	"fmt"
	"sync"
	"os"
	"time"
)

func Mutex() {
	mutex := new(sync.Mutex)	
	for i:=0; i<10; i++ {
		for j:=1; j<10; j++ {
			mutex.Lock()
			go func() {
				fmt.Printf("%d+%d=%d\n", i, j, i+j)
				mutex.Unlock()
			}()
		}
	}
	fmt.Scanln()
}

func MutexChan() {
	mutex := make(chan bool, 1)
	for i:=0; i<10; i++ {
		for j:=1; j<10; j++ {
			mutex <- true
			go func() {
				fmt.Printf("%d+%d=%d\n", i, j, i+j)
				<-mutex
			}()
		}
	}	
	fmt.Scanln()
}

func MutexLog() {
	mutex := make(chan bool, 1)
	logChan := make(chan string, 25)
	file, err := os.OpenFile("./temp/mutex.log", os.O_CREATE, os.FileMode(int(0664)))
	if err != nil {
		fmt.Errorf("Error creating log file")
	}
	defer file.Close()
	go func() {		
		for {
			msg, ok := <- logChan
			if ok {				
				file.WriteString(time.Now().Format(time.RFC3339) + " - " + msg)
			} else {
				break 
			}			
		}
	}()
	
	for i:=0; i<10; i++ {
		for j:=1; j<10; j++ {
			mutex <- true
			go func() {
				msg := fmt.Sprintf("%d+%d=%d\n", i, j, i+j)
				logChan <- msg
				fmt.Printf(msg)
				<-mutex
			}()
		}
	}	
	fmt.Scanln()
}

