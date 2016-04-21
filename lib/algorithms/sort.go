package algorithms

/*
"theta"-notation simply focus on the dominant term, dropping low-order terms and
constant factors.

We use O-notation to indicate that a running time is never worse
than constant times of some f(n)

"omega"-notation: a function f(n) is g(n) if, once n becomes sufficiently
large, f(n) is bounded from below by some constant times g(n). We
say that “f(n) is big-omega of g(n)” or just “f(n) is omega of g(n)”
and we can write f(n)=Omega(g(n).

*/

import (
	"fmt"
	"sort"
)

// A couple of type definitions to make the units clear.
type earthMass float64
type au float64

// By is the type of a "less" function that defines the ordering of its Planet arguments.
type By func(p1, p2 *Planet) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(planets Planets) {
	ps := &planetSorter{
		planets: planets,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// planetSorter joins a By function and a slice of Planets to be sorted.
type planetSorter struct {
	planets Planets
	by      func(p1, p2 *Planet) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *planetSorter) Len() int {
	return len(s.planets)
}

// Swap is part of sort.Interface.
func (s *planetSorter) Swap(i, j int) {
	s.planets[i], s.planets[j] = s.planets[j], s.planets[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *planetSorter) Less(i, j int) bool {
	return s.by(&s.planets[i], &s.planets[j])
}

// ExampleSortKeys demonstrates a technique for sorting a struct type using programmable sort criteria.
func TestSort() {
	// Closures that order the Planet structure.
	name := func(p1, p2 *Planet) bool {
		return p1.Name < p2.Name
	}

	// Sort the planets by the various criteria.
	By(name).Sort(planets)
	fmt.Println("By name:", planets)
}

/*If moving array elements
  is particularly time-consuming—perhaps because they are large
  or stored on a slow device such as a disk—then selection sort might be
  a reasonable algorithm to use.*/
func (this *Planets) SelectionSort() {
	var arr = *this
	for i, a := range arr {
		for j, b := range arr[:i] {
			if a.Name < b.Name {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func TestSelectionSort() {
	fmt.Println(planets)
	planets.SelectionSort()
	fmt.Println(planets)
}

//Insertion sort is an excellent choice when the array starts out as “almost sorted.”
func (this *Planets) InsertionSort() {
	var a = *this
	for i := 1; i < len(a); i++ { //start from 2nd element
		key, j := a[i], i-1 //get the key of 2nd element
		for j >= 0 && a[j].Name > key.Name {
			a[j+1] = a[j] //keep pushing to the top
			j--
		}
		a[j+1] = key //insert the key at the end
	}
}

func TestInsertionSort() {
	fmt.Println(planets)
	planets.InsertionSort()
	fmt.Println(planets)
}

/*The constant factor in the asymptotic notation
  is higher than for selection and insertion sort algorithms
  but does not matter when n becomes large.*/
func (this *Planets) MergeSort() {
	var a = *this
	var b = mergeSort(a)
	for i, v := range b {
		a[i] = v
	}
}

func mergeSort(a Planets) Planets {
	if len(a) <= 1 {
		return a
	}
	var m = len(a) / 2
	var l = mergeSort(a[:m])
	var r = mergeSort(a[m:])
	return merge(l, r)
}

func merge(l, r Planets) Planets {
	a := make(Planets, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(a, r...)
		}
		if len(r) == 0 {
			return append(a, l...)
		}
		if l[0].Name <= r[0].Name {
			a = append(a, l[0])
			l = l[1:]
		} else {
			a = append(a, r[0])
			r = r[1:]
		}
	}
	return a
}

func TestMergeSort() {
	fmt.Println(planets)
	planets.MergeSort()
	fmt.Println(planets)
}

func QuickSort(a Planets) {
	quickSort(a, 0, len(a))
}

//we can use random 3 elements to choose r from and swap it
//that is to ensure running time of O(n lg n) and avoid splits in partition
//otherwise it is possible someone may give an array that produces worst splitss
func quickSort(a Planets, p, r int) {
	if p >= r-1 {
		return
	}
	var q = partition(a, p, r)
	quickSort(a, p, q)
	quickSort(a, q+1, r)
}

func partition(a Planets, p, r int) int {
	var q = p     //start of the array
	var v = r - 1 //pivot
	for u := q; u < v; u++ {
		if a[u].Name <= a[v].Name {
			a[q], a[u] = a[u], a[q] //keep shifting ones that are less then pivot to the front
			q++                     //move forward after each shift
		}
	}
	a[q], a[v] = a[v], a[q] //and replace pivot
	return q                //return next pivot
}

func TestQuickSort() {
	fmt.Println(planets)
	QuickSort(planets)
	fmt.Println(planets)
}

func CountingSort(a []int) {
	var m = getMaxRange(a)
	var equal = countEqual(a, m)   //index equal numbers by occurence
	var less = countLess(equal, m) //less contains indexes of where a elements should go in b
	var b = rearrange(a, less, m)  //sort a by copying from a using next indexes
	fmt.Println(b)
}

func getMaxRange(a []int) int {
	var m = 0
	for _, v := range a {
		if v > m {
			m = v
		}
	}
	return m + 1
}

func countEqual(a []int, m int) []int {
	var equal = make([]int, m, m)
	for _, v := range a {
		equal[v]++
	}
	return equal
}

func countLess(equal []int, m int) []int {
	var less = make([]int, m, m)
	for j := 1; j < m; j++ {
		less[j] = less[j-1] + equal[j-1]
	}
	return less
}

func rearrange(a, next []int, m int) []int {
	var n = len(a)
	var b = make([]int, n, n)
	for i := 0; i < n; i++ {
		var key = a[i]
		var index = next[key]
		b[index] = key
		next[key]++
	}
	return b
}

func TestCountingSort() {
	var a = []int{4, 1, 5, 0, 1, 6, 5, 1, 5}
	CountingSort(a)
}
