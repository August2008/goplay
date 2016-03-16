package channels

import (	
	"os"
	"encoding/csv"	
	"strconv"
	"time"
	"fmt"
	"sync"
)

func TestETL3() {
	var start = time.Now()
	
	var extractChannel = make(chan *Order2)
	var transformChannel = make(chan *Order2)
	var doneChannel = make(chan bool)
	
	go extract3(extractChannel)
	go transform3(extractChannel, transformChannel)
	go load3(transformChannel, doneChannel)
	
	<-doneChannel
	
	fmt.Println(time.Since(start))
}

func extract3(extractChannel chan *Order2)  {	
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

func transform3(extractChannel, transformChannel chan *Order2) {
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
	var wg sync.WaitGroup
	for order := range extractChannel {		
		wg.Add(1)
		go func() {			
			time.Sleep(time.Millisecond * 3)
			defer wg.Done()
			var product = products[order.PartNo]
			order.UnitCost = product.UnitCost
			order.UnitPrice = product.UnitPrice
			transformChannel <- order			
		}()
	}
	wg.Wait()	
	close(transformChannel)	
}

func load3(transformChannel chan *Order2, doneChannel chan bool) {
	var file, _ = os.Create("./temp/result.txt")
	defer file.Close()
	
	fmt.Fprintf(file, "%6s %20s %15s %12s %12s %15s %15s\n", 
		"Row No", "Part Number", "Quantity", "Unit Cost", "Unit Price", "Total Cost", "Total Price")	
	
	var wg sync.WaitGroup
	var row = 0;
	for order := range transformChannel {
		wg.Add(1)
		row++
		go func(row int) {
			time.Sleep(time.Millisecond)
			defer wg.Done()
			fmt.Fprintf(file, "%6d %20s %15d %12.2f %12.2f %15.2f %15.2f\n", 
				row,
				order.PartNo, 
				order.Quantity, 			
				order.UnitCost, order.UnitPrice,
				order.UnitCost * float64(order.Quantity), 
				order.UnitPrice * float64(order.Quantity))	
		}(row)
	}
	wg.Wait()
	doneChannel <- true	
}