package algorithms

import (
	"fmt"
	//"time"
)

type Task struct {
	Name   string
	Weight int
}

func TopologicalSort(tasks []Task, al [][]int) []int {
	fmt.Println(tasks)
	var indegree = make([]int, len(tasks))

	for u := 0; u < len(al); u++ {
		var children = al[u] //all child vertices of u
		for v := 0; v < len(children); v++ {
			indegree[children[v]]++ //increase the count of entering edges in the vertex
		}
	}
	fmt.Println(indegree)
	var next = make([]int, 0)
	for i, v := range indegree {
		if v == 0 {
			next = append(next, i) //vertices without entering edges
		}
	}
	fmt.Println(next)
	var lo = make([]int, 0, len(al)) //linear order of tasks
	for len(next) > 0 {
		var u = next[0]
		var children = al[u]                 //get children of u
		next = next[1:]                      //remove u from the next
		lo = append(lo, u)                   //add u at the end of linear order
		for i := 0; i < len(children); i++ { //for each v adjacent to u
			var v = children[i]
			indegree[v]-- //decrementing means we've taken care of the dependency
			if indegree[v] == 0 {
				next = append(next, v)
				var n = len(next) - 1
				next = append(next[n:], next[:n]...) //v must be added at the top of next
			}
		}
	}
	for _, v := range lo {
		fmt.Println(tasks[v])
	}
	return lo
}

func TestHockeyGear() {
	var tasks = []Task{
		Task{Name: "undershorts"},        //0
		Task{Name: "socks"},              //1
		Task{Name: "compression shorts"}, //2
		Task{Name: "hose"},               //3
		Task{Name: "cup"},                //4
		Task{Name: "pants"},              //5
		Task{Name: "skates"},             //6
		Task{Name: "leg pads"},           //7
		Task{Name: "T-shirt"},            //8
		Task{Name: "chest pad"},          //9
		Task{Name: "sweater"},            //10
		Task{Name: "mask"},               //11
		Task{Name: "catch glove"},        //12
		Task{Name: "blocker"},            //13
	}
	var al = [][]int{ //define adjacency list
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
	var lo = TopologicalSort(tasks, al)
	fmt.Println(lo)
}

func TestShortestPath() {
	//PERT - Program Evaluation and Review Technique
	var tasks = []Task{
		Task{Name: "start", Weight: 0},                         //0
		Task{Name: "cut up chicken", Weight: 6},                //1
		Task{Name: "mix marinade", Weight: 2},                  //2
		Task{Name: "marinate chicken", Weight: 15},             //3
		Task{Name: "partially cook chicken", Weight: 4},        //4
		Task{Name: "chop garlic", Weight: 4},                   //5
		Task{Name: "chop ginger", Weight: 3},                   //6
		Task{Name: "add garlic, ginger", Weight: 1},            //7
		Task{Name: "finish cooking chicken", Weight: 2},        //8
		Task{Name: "remove chicken", Weight: 1},                //9
		Task{Name: "chop carrots", Weight: 4},                  //10
		Task{Name: "chop celery", Weight: 3},                   //11
		Task{Name: "rinse peanuts", Weight: 2},                 //12
		Task{Name: "cook carrots, celery, peanuts", Weight: 4}, //13
		Task{Name: "add back chicken", Weight: 1},              //14
		Task{Name: "mix cooking sauce", Weight: 3},             //15
		Task{Name: "add cooking sauce", Weight: 1},             //16
		Task{Name: "cook sauce until thick", Weight: 3},        //17
		Task{Name: "remove completed dish", Weight: 1},         //18
		Task{Name: "finish", Weight: 0},                        //19
	}
	var al = [][]int{
		{1, 2, 5, 6, 10, 11, 12, 15}, //0
		{3},  //1
		{3},  //2
		{4},  //3
		{7},  //4
		{7},  //5
		{7},  //6
		{8},  //7
		{9},  //8
		{13}, //9
		{13}, //10
		{13}, //11
		{13}, //12
		{14}, //13
		{16}, //14
		{16}, //15
		{17}, //16,
		{18}, //17,
		{19}, //18
		{},
	}
	var lo = TopologicalSort(tasks, al)
	fmt.Println(lo)

	//negate weights as we're only interested in single shorted path
	//when the weights are added up, min sum of all weights will be the shortest
	for _, v := range lo {
		tasks[v].Weight = -tasks[v].Weight
	}
	fmt.Println(tasks)

	var n = len(lo)
	var shortest = make([]int, n, n)
	var pred = make([]int, n, n)

	for i := 0; i < len(lo); i++ { //walk down the sorted dag
		var u = lo[i]
		var children = al[u] //all child vertices of u which it connects to
		for j := 0; j < len(children); j++ {
			var v = children[j]
			//relax(u, v)
			var weight = shortest[u] + tasks[v].Weight
			if weight < shortest[v] {
				shortest[v] = weight
				pred[v] = u
			}
		}
	}
	fmt.Println(shortest) //min of all is the shortest path that is -39 in this case
	fmt.Println(pred)     //contains tasks on the way to critical path whose weights add up to -39

	var critical = make([]Task, 0)
	for _, v := range pred {
		if v > 0 {
			pert = append(critical, tasks[v])
		}
	}
	fmt.Println(critical) //critical path in the PERT chart
}
