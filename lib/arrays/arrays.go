package arrays

import "fmt"

// Example - Without this comment, Go compiler gives warning
type Example struct {
	// Exported public field
	Text []string
	MapArray map[string][]int
}

// PrintArray is a exported functions
func PrintArray() {
	//arr1 := [3]int{1, 2, 3}
	//fmt.Println(arr1)

	//arr2 := [][]int{{1, 2, 3}, {4, 5, 6}}
	//fmt.Println(arr2)
	/*
	structarr := []Example{
		Example{
			Text: []string{"a", "b", "c"},
			MapArray: nil,
		},
	}
	//fmt.Println(structarr) 
	*/
	//slice := arr1[:] 
	//slice = append(slice, 100) 
	//fmt.Println(slice)
	
	example := new(Example)
	example.MapArray = make(map[string][]int)
	
	example.MapArray["first"] = []int{1,2,3,4,5,6,7,8,9,10}
	example.MapArray["second"] = []int{11,12,13,14,15,16,17,18,19,20}
	example.MapArray["third"] = []int{21,22,23,24,25,26,27,28,29,30}
	
	ar1 := example.MapArray["first"]	
	fmt.Println(ar1)
	
	ar2 := append(ar1[:7], ar1[8:]...) //remove 8
	fmt.Println(ar2)
	
	ar3 := append(ar2[:3], append([]int{12,13}, ar2[3:]...)...) //insert 12,13 between 3 and 4
	fmt.Println(ar3)
	
	ar4 := append(ar3[:3], ar3[5:]...) //remove 12,13
	fmt.Println(ar4)
}
