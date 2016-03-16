package channels

import (	
	"os"
	"encoding/csv"	
	"strconv"
	"time"
	"fmt"
)

func TestETL2() {
	var start = time.Now()
	
	var extractChannel = make(chan *Order2)
	var transformChannel = make(chan *Order2)
	var doneChannel = make(chan bool)
	
	go extract2(extractChannel)
	go transform2(extractChannel, transformChannel)
	go load2(transformChannel, doneChannel)
	
	<-doneChannel
	
	fmt.Println(time.Since(start))
}

func extract2(extractChannel chan *Order2)  {	
	var file, _ = os.Open("./data/orders.csv")
	defer file.Close()
	
	var reader = csv.NewReader(file)
	for record, err := reader.Read(); err == nil; record, err = reader.Read() {
		var order = new(Order2)
		order.CustomerNo, _ = strconv.Atoi(record[0])
		order.PartNo = record[1]
		order.Quantity, _ = strconv.Atoi(record[2])
		extractChannel <- order
	}	
	close(extractChannel)
}

func transform2(extractChannel, transformChannel chan *Order2) {
	var file, _ = os.Open("./data/products.csv")
	defer file.Close()
	
	var reader = csv.NewReader(file)
	var records, _ = reader.ReadAll()
	var products = make(map[string]*Product2)
	
	for _, record := range records {
		var product = new(Product2)
		product.PartNo = record[0]
		product.UnitCost, _ = strconv.ParseFloat(record[1], 64)
		product.UnitPrice, _ = strconv.ParseFloat(record[2], 64)		
		products[product.PartNo] = product
	}
	for order := range extractChannel {
		time.Sleep(time.Millisecond * 3)
		var product = products[order.PartNo]
		order.UnitCost = product.UnitCost
		order.UnitPrice = product.UnitPrice
		transformChannel <- order
	}	
	close(transformChannel)
}

func load2(transformChannel chan *Order2, doneChannel chan bool) {
	var file, _ = os.Create("./temp/result.txt")
	defer file.Close()
	
	fmt.Fprintf(file, "%20s %15s %12s %12s %15s %15s\n", 
		"Part Number", "Quantity", "Unit Cost", "Unit Price", "Total Cost", "Total Price")	
	
	for order := range transformChannel {
		time.Sleep(time.Millisecond)
		fmt.Fprintf(file, "%20s %15d %12.2f %12.2f %15.2f %15.2f\n", 
			order.PartNo, 
			order.Quantity, 			
			order.UnitCost, order.UnitPrice,
			order.UnitCost * float64(order.Quantity), 
			order.UnitPrice * float64(order.Quantity))	
	}
	doneChannel <- true	
}