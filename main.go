package main

import (
	//chn "github.com/August2008/goplay/lib/channels"
	"runtime"
	//"time"
	"fmt"
	//"github.com/August2008/goplay/lib/arrays"
	//"github.com/August2008/goplay/lib/xml"
	//"goplay/lib/sysio"
	//algo "github.com/August2008/goplay/lib/algorithms"
	"github.com/August2008/goplay/lib/trees"
)

type Job struct {
	Name   string
	Weight int
}

type JobSorter struct {
	Jobs []Job
}

func (j *JobSorter) Less(i, j int) bool {
	return j.Jobs[i] < j.Jobs[j]
}

func main() {
	runtime.GOMAXPROCS(4)

	//arrays.PrintArray()
	//algo.TestMergeSort()
	//algo.TestNFactrorial()
	//algo.TestQuickSort()
	//algo.TestCountingSort()
	//algo.TestShortestPath()

	//arrays.PrintArray()
	//xml.TestWrite()
	//arrays.PrintMap()
	//consts.Print()
	//branching.PrintIfElse()
	//branching.PrintLoops()

	var h = trees.NewBHeap(20)
	h.Push(&Job{Name: "A", Weight: 3})
	h.Push(&Job{Name: "A", Weight: 2})
	h.Push(&Job{Name: "A", Weight: 4})
	h.Push(&Job{Name: "A", Weight: 1})
	h.Push(&Job{Name: "A", Weight: 9})
	h.Push(&Job{Name: "A", Weight: 5})
	h.Push(&Job{Name: "A", Weight: 6})
	h.Push(&Job{Name: "A", Weight: 8})
	h.Push(&Job{Name: "A", Weight: 7})
	h.Push(&Job{Name: "A", Weight: 1})
	h.Push(&Job{Name: "A", Weight: 1})

	for !h.Empty() {
		fmt.Println(h.ToArray())
		fmt.Println(h.Pop())
	}

	//owl := st.NewOwl()
	//fmt.Println(owl.Taxonomy())

	//ch := make(chan string)
	//go Alphabet(ch)

	//done := make(chan bool, 1)
	//go Printer(ch, done)

	//sysio.StartWatching()

	//chn.Play3()

	//chn.MutexLog()

	//chn.TestButton()

	//chn.TestCallbacks()

	//chn.TestPromise()

	//chn.TestPrimes()

	//chn.TestETL3()

	//<-done
	//time.Sleep(time.Second * 10)
	//fmt.Scanln()
	//println("Done!")
}

//Alphabet function
func Alphabet(ch chan string) {
	for i := byte('a'); i < byte('z'); i++ {
		ch <- string(i)
	}
	close(ch)
}

//Printer function
func Printer(ch chan string, done chan bool) {
	for r := range ch {
		println(r)
	}
	done <- true
}
