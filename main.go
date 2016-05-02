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

	var h = trees.NewMinHeap(20)
	h.Push(&trees.Node{Value: "A", Ordinal: 3})
	h.Push(&trees.Node{Value: "B", Ordinal: 2})
	h.Push(&trees.Node{Value: "C", Ordinal: 4})
	h.Push(&trees.Node{Value: "D", Ordinal: 7})
	h.Push(&trees.Node{Value: "E", Ordinal: 9})
	h.Push(&trees.Node{Value: "F", Ordinal: 5})
	h.Push(&trees.Node{Value: "G", Ordinal: 6})
	h.Push(&trees.Node{Value: "H", Ordinal: 8})
	h.Push(&trees.Node{Value: "I", Ordinal: 7})
	h.Push(&trees.Node{Value: "J", Ordinal: 1})
	h.Push(&trees.Node{Value: "K", Ordinal: 0})

	for h.Len() > 0 {
		//fmt.Println(h.ToArray())
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
