package channels

import (	
	"os"
	"encoding/csv"	
	"strconv"
	"time"
	"fmt"
)

type Order2 struct {
	CustomerNo int
	PartNo string
	Quantity int
	UnitCost float64
	UnitPrice float64
}

type Product2 struct {
	PartNo string
	UnitCost float64
	UnitPrice float64
}

func TestETL1() {
	var start = time.Now()
	var orders = extract1()
	orders = transform1(orders)
	load1(orders)
	fmt.Println(time.Since(start))
}

func extract1() []*Order2 {
	var result = []*Order2{}
	
	var file, _ = os.Open("./data/orders.csv")
	defer file.Close()
	
	var reader = csv.NewReader(file)
	for record, err := reader.Read(); err == nil; record, err = reader.Read() {
		var order = new(Order2)
		order.CustomerNo, _ = strconv.Atoi(record[0])
		order.PartNo = record[1]
		order.Quantity, _ = strconv.Atoi(record[2])
		result = append(result, order)
	}
	return result
}

func transform1(orders []*Order2) []*Order2 {
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
	for _, order := range orders {
		time.Sleep(time.Millisecond * 3)
		var product = products[order.PartNo]
		order.UnitCost = product.UnitCost
		order.UnitPrice = product.UnitPrice
	}
	return orders
}

func load1(orders []*Order2) {
	var file, _ = os.Create("./temp/result.txt")
	defer file.Close()
	
	fmt.Fprintf(file, "%20s %15s %12s %12s %15s %15s\n", 
		"Part Number", "Quantity", "Unit Cost", "Unit Price", "Total Cost", "Total Price")	
	
	for _, order := range orders {
		time.Sleep(time.Millisecond)
		fmt.Fprintf(file, "%20s %15d %12.2f %12.2f %15.2f %15.2f\n", 
			order.PartNo, 
			order.Quantity, 			
			order.UnitCost, order.UnitPrice,
			order.UnitCost * float64(order.Quantity), 
			order.UnitPrice * float64(order.Quantity))	
	}
}