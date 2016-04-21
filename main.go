package main

import (
	//chn "github.com/August2008/goplay/lib/channels"
	"runtime"
	//"time"
	//"fmt"
	//"github.com/August2008/goplay/lib/arrays"
	//"github.com/August2008/goplay/lib/xml"
	//"goplay/lib/sysio"
	algo "github.com/August2008/goplay/lib/algorithms"
)

func main() {
	runtime.GOMAXPROCS(4)

	//arrays.PrintArray()
	//algo.TestMergeSort()
	//algo.TestNFactrorial()
	//algo.TestQuickSort()
	//algo.TestCountingSort()
	algo.TestShortestPath()

	//arrays.PrintArray()
	//xml.TestWrite()
	//arrays.PrintMap()
	//consts.Print()
	//branching.PrintIfElse()
	//branching.PrintLoops()

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
