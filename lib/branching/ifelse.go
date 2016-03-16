package branching

import f "fmt"

// PrintIfElse function
func PrintIfElse() {
	val1 := 2
	if val1 == 5 {
		f.Println(val1)
	}

	val3 := 3
	if val2 := val1; val2 >= 5 {
		f.Println(val2)
	} else if val3 == 3 {
		f.Println(val3)
	}
}
