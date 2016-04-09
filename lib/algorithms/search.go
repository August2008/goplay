package algorithms

import (
	"fmt"
)

func (this *Planets) LinearSearch(name string) *Planet {
	for _, val := range *this {
		if val.Name == name {
			return &val
		}
	}
	return nil
}

func TestLinearSearch() {
	var planet = planets.LinearSearch("Earth")
	fmt.Println(planet)
}

func (this *Planets) BinarySearch(name string) *Planet {
	//print(&this)
	var p = 0
	var n = len(*this)
	for p <= n {
		var m = (p + n) / 2
		var planet = (*this)[m]
		fmt.Printf("m=%v, p=%v, n=%v, planet=%v\n", m, p, n, planet)
		if planet.Name == name {
			return &planet
		}
		if planet.Name > name {
			n = m - 1
		} else {
			p = m + 1
		}
		fmt.Printf("p=%v, n=%v\n", p, n)
	}
	return nil
}

func TestBinarySearch() {
	name := func(p1, p2 *Planet) bool {
		return p1.Name < p2.Name
	}
	By(name).Sort(planets)
	//println(&planets)
	var planet = planets.BinarySearch("Pluto")
	fmt.Println(planet)
}

func NFactrorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * NFactrorial(num-1)
}

func TestNFactrorial() {
	fmt.Println(NFactrorial(5))
}
