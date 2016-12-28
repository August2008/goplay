package main

import (
	//chn "github.com/dmumladze/goplay/lib/channels"
	"runtime"
	//"time"
	"fmt"
	//"github.com/dmumladze/goplay/lib/arrays"
	//"github.com/dmumladze/goplay/lib/xml"
	//"goplay/lib/sysio"
	//algo "github.com/dmumladze/goplay/lib/algorithms
	//algo "github.com/dmumladze/goplay/lib/algorithms"
	//"github.com/dmumladze/goplay/lib/trees"
	"github.com/dmumladze/goplay/lib/lists"
	//algo "github.com/dmumladze/goplay/lib/algorithms"
	//"github.com/dmumladze/goplay/lib/trees"
)

func main() {
	runtime.GOMAXPROCS(4)

	//	var s = lists.NewStack()
	//
	//	s.Push(1)
	//	s.Push(2)
	//	s.Push(3)
	//	s.Push(4)
	//	s.Push(5)
	//
	//	s.Pop()
	//	s.Push(9)
	//
	//	for s.Size() > 0 {
	//		fmt.Println(s.Pop())
	//	}

	var s = lists.NewStack()
	var infix = "1 + 2 ) * 3 - 4 ) * 5 - 6 ) ) )"
	var parens = 0

	for i := 0; i < len(infix); i++ {
		var chr = string(infix[i])
		switch chr {
		case " ":
			continue
		case ")":
			parens++
		}
		s.Push(chr)
	}

	var q = lists.NewStack()

	for {
		if s.Size() == 0 {
			if parens > 0 {
				for i := 0; i < parens; i++ {
					q.Push("(")
				}
			}
			break
		}
		var chr = s.Peek().(string)
		switch chr {
		case "*":
			q.Push("(")
			parens--
		}
		q.Push(s.Pop())
	}

	fmt.Println(parens)

	for q.Size() > 0 {
		fmt.Print(q.Pop())
	}

	//algo.TestBellmanFord()
	//algo.TestHeapsort()
	//algo.TestDijkstra()
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

	//	var h = trees.NewMinHeap(20)
	//	h.Push(&trees.Node{Value: "A", Ordinal: 3})
	//	h.Push(&trees.Node{Value: "B", Ordinal: 2})
	//	h.Push(&trees.Node{Value: "C", Ordinal: 4})
	//	h.Push(&trees.Node{Value: "D", Ordinal: 7})
	//	h.Push(&trees.Node{Value: "E", Ordinal: 9})
	//	h.Push(&trees.Node{Value: "F", Ordinal: 5})
	//	h.Push(&trees.Node{Value: "G", Ordinal: 6})
	//	h.Push(&trees.Node{Value: "H", Ordinal: 8})
	//	h.Push(&trees.Node{Value: "I", Ordinal: 7})
	//	h.Push(&trees.Node{Value: "J", Ordinal: 1})
	//	h.Push(&trees.Node{Value: "K", Ordinal: 0})
	//
	//	for h.Len() > 0 {
	//		//fmt.Println(h.ToArray())
	//		fmt.Println(h.Pop())
	//	}

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
