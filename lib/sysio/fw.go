package sysio

import (
	"os"
	"time"
	"io/ioutil"
	"strings"
	"encoding/csv"
	"strconv"
	"fmt"
)

const watchedDir = "./temp"

type Invoice struct {
	Number string
	Amount float64
	PurchaseOrderNumber int
	InvoiceDate time.Time
}

func StartWatching() {
	for {
		dir, _ := os.Open(watchedDir)
		files, _ := dir.Readdir(-1)
		
		for _, file := range files {			
			fileName := watchedDir + "/" + file.Name()
			
			if _, err := os.Stat(fileName); err == nil {									
				fileHandle, _ := os.Open(fileName)				
				data, _ := ioutil.ReadAll(fileHandle)
				
				fileHandle.Close()					
				os.Remove(fileName)
										
				fmt.Printf("Processed %s\n", fileName)
				
				go func(data string) {
					reader := csv.NewReader(strings.NewReader(data))
					records, _ := reader.ReadAll()
					for _, r := range records {
						inv := new(Invoice)
						inv.Number = r[0]
						inv.Amount, _ = strconv.ParseFloat(r[1], 64)						
						fmt.Printf("%v %v\n", inv.Number, inv.Amount)					
					}
				}(string(data))		
			} 
		}
	}
}

