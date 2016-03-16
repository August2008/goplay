package channels

import (
	"fmt"
)

type Order interface {
	SetNumber()	
}

type PurchaseOrder struct {
	Number int
	Value float64
}

func (this *PurchaseOrder) SaveWithCallback(callback chan *PurchaseOrder) {
	callback <- this
}

func (this *PurchaseOrder) SetNumber() {
	this.Number = 123
}

func Archive(item Order) {	
	fmt.Printf("Archived %v", item)
}

func TestCallbacks() {
	var po = new(PurchaseOrder)
	po.Value = 125.99
	
	var callback = make(chan *PurchaseOrder, 1)	
	go po.SaveWithCallback(callback)
	
	fmt.Println(<- callback)
	
	Archive(po)	
}

