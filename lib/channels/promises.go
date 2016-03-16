package channels

import (
	"fmt"
	"errors"
	"time"
)

type Promise struct {
	successChannel chan interface{}
	failureChannel chan error
}

func (this *PurchaseOrder) SaveWithPromise(shouldFail bool) *Promise {
	var result = new(Promise)
	
	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)
	
	go func() {
		time.Sleep(time.Second * 6)
		if shouldFail {
			result.failureChannel <- errors.New("Failed to save purchase order")
		} else {
			this.SetNumber()
			result.successChannel <- this
		}
	}()
	
	return result
}

func (this *Promise) Then(success func(interface{}) error, failure func(error)) *Promise {
	var result = new(Promise)
	
	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)
	
	var timeout = time.After(time.Second * 5)
	
	go func() {
		select {
			case obj := <- this.successChannel: 
				var newErr = success(obj)
				if newErr == nil {
					result.successChannel <- obj
				} else {
					result.failureChannel <- newErr
				}
			case err := <- this.failureChannel:
				failure(err)	
				result.failureChannel <- err
			case <- timeout:
				var err = errors.New("Promise timeout")
				failure(err)		
				result.failureChannel <- err		
		}
	}()
	
	return result 
}

func TestPromise() {
	var order = new(PurchaseOrder)
	order.Value = 25.99
	
	var promise = order.SaveWithPromise(false)
	
	promise.Then(func(obj interface{}) error {
			var order = obj.(*PurchaseOrder)
			fmt.Printf("Saved order %v\n", order)
			return nil
	}, func(err error) {
			fmt.Println(err)
	}).Then(func(obj interface{}) error {
			fmt.Println("Second promise")
			return nil
	}, func(err error) {
			fmt.Println(err)
	}).Then(func(obj interface{}) error {
			fmt.Println("Third promise")
			return nil
	}, func(err error) {
			fmt.Println(err)
	})
}

