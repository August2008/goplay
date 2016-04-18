package algorithms

import (
	"fmt"
	//"time"
)

type HockeyGear struct {
	name string
}

func TopologicalSort() {
	var gear = []HockeyGear{
		HockeyGear{"undershorts"},        //0
		HockeyGear{"socks"},              //1
		HockeyGear{"compression shorts"}, //2
		HockeyGear{"hose"},               //3
		HockeyGear{"cup"},                //4
		HockeyGear{"pants"},              //5
		HockeyGear{"skates"},             //6
		HockeyGear{"leg pads"},           //7
		HockeyGear{"T-shirt"},            //8
		HockeyGear{"chest pad"},          //9
		HockeyGear{"sweater"},            //10
		HockeyGear{"mask"},               //11
		HockeyGear{"catch glove"},        //12
		HockeyGear{"blocker"},            //13
	}
	var graph = [][]int{ //define adjacency list
		{2},     //0-undershorts>compression shorts
		{3},     //1-socks>hose
		{3, 4},  //2-compression shorts>hose,cup
		{5},     //3-hose>pants
		{5},     //4-cup>pants
		{6, 10}, //5-pants>skates,sweater
		{7},     //6-skates>leg pads
		{12},    //7-leg pads>catch gloves
		{9},     //8-T-shirt>chest pad
		{10},    //9-chest pad>sweater
		{11},    //10-sweater>mask
		{12},    //11-mask>catch gloves
		{13},    //12-catch gloves>blocker
		{},      //13-blocker
	}
	fmt.Println(gear)
	var indegree = make([]int, len(gear))

	for u := 0; u < len(graph); u++ {
		var children = graph[u]
		for v := 0; v < len(children); v++ {
			indegree[children[v]]++
		}
	}
	fmt.Println(indegree)
	var next = make([]int, 0)
	for i, v := range indegree {
		if v == 0 {
			next = append(next, i)
		}
	}
	fmt.Println(next)
	var linearOrder = make([]int, 0, len(graph))
	for len(next) > 0 {
		var u = next[0]
		var children = graph[u]              //get children of u
		next = next[1:]                      //remove u from the next
		linearOrder = append(linearOrder, u) //add u at the end of linear order
		for i := 0; i < len(children); i++ { //for each v adjacent to u
			var v = children[i]
			indegree[v]-- //decrementing means we've taken care of the dependency
			if indegree[v] == 0 {
				next = append(next, v)
				var n = len(next) - 1
				next = append(next[n:], next[:n]...) //v must be at the top of next
			}
		}
	}
	fmt.Println(linearOrder)
	for _, v := range linearOrder {
		fmt.Println(gear[v])
	}
}
