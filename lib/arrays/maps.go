package arrays

import "fmt"

// Print function
func PrintMap() {
	myMap := make(map[string]string, 0)

	myMap["first"] = "David"
	myMap["last"] = "Mumladze"
	
	val, ok := myMap["first"]
	if ok {
		fmt.Printf(val)
	} else {
		return
	}	

	fmt.Println(myMap)
}
