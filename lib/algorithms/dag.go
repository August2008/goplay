package algorithms

import (
	"fmt"
	"github.com/August2008/goplay/lib/trees"
	"math"
	//"time"
)

func TopologicalSort(tasks []Vertex, al [][]int) []int {
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
	var tasks = []Vertex{
		Vertex{Name: "undershorts"},        //0
		Vertex{Name: "socks"},              //1
		Vertex{Name: "compression shorts"}, //2
		Vertex{Name: "hose"},               //3
		Vertex{Name: "cup"},                //4
		Vertex{Name: "pants"},              //5
		Vertex{Name: "skates"},             //6
		Vertex{Name: "leg pads"},           //7
		Vertex{Name: "T-shirt"},            //8
		Vertex{Name: "chest pad"},          //9
		Vertex{Name: "sweater"},            //10
		Vertex{Name: "mask"},               //11
		Vertex{Name: "catch glove"},        //12
		Vertex{Name: "blocker"},            //13
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

func TestPERT() {
	//PERT - Program Evaluation and Review Technique
	var tasks = []Vertex{
		Vertex{Name: "start", Weight: 0},                         //0
		Vertex{Name: "cut up chicken", Weight: 6},                //1
		Vertex{Name: "mix marinade", Weight: 2},                  //2
		Vertex{Name: "marinate chicken", Weight: 15},             //3
		Vertex{Name: "partially cook chicken", Weight: 4},        //4
		Vertex{Name: "chop garlic", Weight: 4},                   //5
		Vertex{Name: "chop ginger", Weight: 3},                   //6
		Vertex{Name: "add garlic, ginger", Weight: 1},            //7
		Vertex{Name: "finish cooking chicken", Weight: 2},        //8
		Vertex{Name: "remove chicken", Weight: 1},                //9
		Vertex{Name: "chop carrots", Weight: 4},                  //10
		Vertex{Name: "chop celery", Weight: 3},                   //11
		Vertex{Name: "rinse peanuts", Weight: 2},                 //12
		Vertex{Name: "cook carrots, celery, peanuts", Weight: 4}, //13
		Vertex{Name: "add back chicken", Weight: 1},              //14
		Vertex{Name: "mix cooking sauce", Weight: 3},             //15
		Vertex{Name: "add cooking sauce", Weight: 1},             //16
		Vertex{Name: "cook sauce until thick", Weight: 3},        //17
		Vertex{Name: "remove completed dish", Weight: 1},         //18
		Vertex{Name: "finish", Weight: 0},                        //19
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
	var shortest = make([]float64, n, n)
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
	fmt.Println(shortest) //min time required that is -39 in this case
	fmt.Println(pred)     //contains tasks on the way to critical path whose weights add up to -39

	var critical = make([]Vertex, 0)
	for _, v := range pred {
		if v > 0 {
			critical = append(critical, tasks[v])
		}
	}
	fmt.Println(critical) //critical path in the PERT chart
}

func TestDijkstra() {
	var inf = math.Inf(1)

	var tasks = []*Vertex{
		&Vertex{Name: "start", Weight: 0},                         //0
		&Vertex{Name: "cut up chicken", Weight: 6},                //1
		&Vertex{Name: "mix marinade", Weight: 2},                  //2
		&Vertex{Name: "marinate chicken", Weight: 15},             //3
		&Vertex{Name: "partially cook chicken", Weight: 4},        //4
		&Vertex{Name: "chop garlic", Weight: 3},                   //5
		&Vertex{Name: "chop ginger", Weight: 3},                   //6
		&Vertex{Name: "add garlic, ginger", Weight: 1},            //7
		&Vertex{Name: "finish cooking chicken", Weight: 2},        //8
		&Vertex{Name: "remove chicken", Weight: 1},                //9
		&Vertex{Name: "chop carrots", Weight: 4},                  //10
		&Vertex{Name: "chop celery", Weight: 3},                   //11
		&Vertex{Name: "rinse peanuts", Weight: 2},                 //12
		&Vertex{Name: "cook carrots, celery, peanuts", Weight: 4}, //13
		&Vertex{Name: "add back chicken", Weight: 1},              //14
		&Vertex{Name: "mix cooking sauce", Weight: 3},             //15
		&Vertex{Name: "add cooking sauce", Weight: 1},             //16
		&Vertex{Name: "cook sauce until thick", Weight: 3},        //17
		&Vertex{Name: "remove completed dish", Weight: 1},         //18
		&Vertex{Name: "finish", Weight: 0},                        //19
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
		{17}, //16
		{18}, //17
		{19}, //18
		{},   //19
	}

	var (
		dist = make([]float64, len(tasks))
		pred = make([]interface{}, len(tasks))
		h    = trees.NewMinHeap()
		s    = 0
	)
	for i, _ := range tasks {
		dist[i] = inf
	}
	dist[s] = 0

	h.Push(s, 0)
	for _, children := range al {
		for _, v := range children {
			h.Push(v, dist[v])
		}
	}

	for h.Len() > 0 {
		var (
			u        = h.Pop().(int)
			children = al[u]
		)
		for _, v := range children {
			var weight = dist[u] + tasks[v].Weight
			if weight < dist[v] {
				dist[v] = weight
				pred[v] = u
				h.Push(v, weight)
			}
		}
	}

	//fmt.Println(pred)

	for i := len(pred) - 1; i >= 1; i-- {
		var v = pred[i]
		switch v.(type) {
		case int:
			if v.(int) == s {
				dist = append(dist[:0], dist[i+1:]...)
				pred = append(pred[:0], pred[i+1:]...)
				goto done
			}
		}
	}
done:
	fmt.Println(pred)
	fmt.Println(dist)
}

func TestBellmanFord() {
	var routes = make([]Vertex, 11, 11)
	routes[0] = Vertex{Name: "Current Location"} //0
	routes[1] = Vertex{Name: "Chestnut St"}      //1
	routes[2] = Vertex{Name: "Palisades Ave"}    //2
	routes[3] = Vertex{Name: "Dorothy Ct"}       //3
	routes[4] = Vertex{Name: "Bellman Cir"}      //4
	routes[5] = Vertex{Name: "Ford Blvd"}        //5
	routes[6] = Vertex{Name: "Jamaica Ave"}      //6
	routes[7] = Vertex{Name: "Cory Ln"}          //7
	routes[8] = Vertex{Name: "Cormen St"}        //8
	routes[9] = Vertex{Name: "Barselona Blvd"}   //9
	routes[10] = Vertex{Name: "Finish"}          //10

	var edges = make(map[string][]Vertex)
	edges["Current Location"] = []Vertex{ //0
		Vertex{Name: "Chestnut St", Weight: 3},     //1
		Vertex{Name: "Palisades Ave", Weight: 1.5}, //2
	}
	edges["Chestnut St"] = []Vertex{ //1
		Vertex{Name: "Dorothy Ct", Weight: 0.8}, //3
	}
	edges["Palisades Ave"] = []Vertex{ //2
		Vertex{Name: "Dorothy Ct", Weight: 9}, //3
	}
	edges["Dorothy Ct"] = []Vertex{ //3
		Vertex{Name: "Bellman Cir", Weight: 0.7}, //4
	}
	edges["Bellman Cir"] = []Vertex{ //4
		Vertex{Name: "Ford Blvd", Weight: 0.4}, //5
	}
	edges["Ford Blvd"] = []Vertex{ //5
		Vertex{Name: "Jamaica Ave", Weight: 3.9}, //6
		Vertex{Name: "Cory Ln", Weight: 6},       //7
	}
	edges["Jamaica Ave"] = []Vertex{ //6
		Vertex{Name: "Barselona Blvd", Weight: 10}, //9
	}
	edges["Cory Ln"] = []Vertex{ //7
		Vertex{Name: "Cormen St", Weight: 4.6},    //8
		Vertex{Name: "Barselona Blvd", Weight: 5}, //9
	}
	edges["Cormen St"] = []Vertex{ //9
		Vertex{Name: "Barselona Blvd", Weight: 0.3}, //9
	}
	edges["Barselona Blvd"] = []Vertex{ //9
		Vertex{Name: "Finish", Weight: 0}, //10
	}

	var (
		source = "Current Location"
		//target = "Finish"
		inf  = 100.0
		dist = make(map[string]float64)
		pred = make(map[string]Vertex)
	)

	for _, v := range routes {
		dist[v.Name] = inf
	}
	dist[source] = 0

	fmt.Println(dist)

	for i, _ := range routes {
		var (
			u        = source
			children = edges[u]
		)
		i = i + 1
		var next Vertex
		for _, to := range children {
			var (
				//v      = to.Name
				weight = dist[u] + to.Weight
			)
			if _, ok := pred[source]; !ok || weight < dist[to.Name] {
				dist[to.Name] = weight
				pred[source] = to
				next = to
			}
		}
		source = next.Name
	}

	source = "Current Location"

	fmt.Println(source)
	for {
		if v, ok := pred[source]; ok {
			fmt.Println(v.Name)
			source = v.Name
			continue
		}
		break
	}

	fmt.Println(dist)
	fmt.Println(pred)
}
