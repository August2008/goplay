package consts

import "fmt"

const (
	first = 1 << iota
	second
	third
)

// Print is a exported functions
func Print() {
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
}
