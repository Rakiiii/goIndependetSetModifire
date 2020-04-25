package gomodificator

import (
	betterslice "github.com/Rakiiii/goBetterSlice"
	ls "github.com/Rakiiii/goBipartitonLocalSearch"
)

//ImproveIndependetSet return graph that number in way of better independent vertex subset size and new order of vertex
func ImproveIndependetSet(graph ls.IGraph) (ls.IGraph, []int) {
	condition := true
	ch := graph
	finOrd := make([]int, graph.AmountOfVertex())
	finOrd = initOrd(finOrd)
	for condition {
		var newOrd []int
		condition, ch, newOrd = improve(graph)
		subOrd := make([]int, len(finOrd))

		for i, v := range newOrd {
			subOrd[i] = finOrd[v]
		}
		finOrd = subOrd
	}
	return ch, finOrd
}

func initOrd(ord []int) []int {
	for i := range ord {
		ord[i] = i
	}
	return ord
}

//improve return graph with once made operation of moving bigger independent subset
func improve(graph ls.IGraph) (bool, ls.IGraph, []int) {
	//subset of dependent with biggest param
	mainSet := make([]int, 0)
	//subset of independent with biggest param
	mainIndependetSet := make([]int, 0)
	//biggest diff between subsets
	bestParam := 0

	//check all subsets
	for i := graph.GetAmountOfIndependent(); i < graph.AmountOfVertex(); i++ {
		//intermediate subset of dependent
		subSet := make([]int, 1)
		//first vertex must be added to sub set any way
		subSet[0] = i

		//set of banned vertex - vertex that cannot create independent subset
		edgeSet := make([]int, 0)
		edgeSet = betterslice.AppendWithOutRepeatInt(edgeSet, graph.GetEdges(i)...)

		for j := i + 1; j < graph.AmountOfVertex(); j++ {
			//check is vertex banned
			jflag := !isInSubSet(j, edgeSet)

			//if vertex not banned
			if jflag {
				//add vertex to subset
				subSet = append(subSet, j)
				//make subset of banned vertex bigger
				//todo: to think avout better perforamnce, @AppendWithOutRepeatInt can be replaced with @append
				edgeSet = betterslice.AppendWithOutRepeatInt(edgeSet, graph.GetEdges(j)...)
			}
		}

		//subset of independent vertex that connected to vertex in dependent subset
		independetVertexSet := make([]int, 0)

		//fill independentVertexSet
		for _, vertex := range subSet {
			for _, i := range graph.GetEdges(vertex) {
				if i < graph.GetAmountOfIndependent() && !isInSubSet(i, independetVertexSet) {
					independetVertexSet = append(independetVertexSet, i)
				}
			}
		}

		//count value of transition
		param := len(subSet) - len(independetVertexSet)
		//update best value
		if param > bestParam {
			bestParam = param
			mainSet = subSet
			mainIndependetSet = independetVertexSet
		}
	}

	//if no positiov transions then stop
	if bestParam <= 0 {
		newOrd := make([]int, graph.AmountOfVertex())
		newOrd = initOrd(newOrd)
		return false, graph, newOrd
	}

	//constract new order of vertex
	newOrd := constractNewOrder(mainSet, mainIndependetSet, graph.GetAmountOfIndependent(), graph.AmountOfVertex())

	//renumber vertex with new ordet
	graph.RenumVertex(newOrd)
	//set new amount of independent subset
	graph.SetAmountOfIndependent(graph.GetAmountOfIndependent() + bestParam)
	return true, graph, newOrd
}

func constractNewOrder(dset []int, iset []int, ai int, av int) []int {
	newOrd := make([]int, av)
	counter := 0
	//fill vertex of independet set that not moved at the start
	for i := 0; i < ai; i++ {
		if !isInSubSet(i, dset) && !isInSubSet(i, iset) {
			newOrd[counter] = i
			counter++
		}
	}

	//add new independet vertex
	for _, v := range dset {
		newOrd[counter] = v
		counter++
	}

	//add old independent vertex
	for _, v := range iset {
		newOrd[counter] = v
		counter++
	}

	//set all empty placec with impocibel vertex num
	for i := counter; i < av; i++ {
		newOrd[i] = -1
	}

	//add remaining vertex
	for i := 0; i < av; i++ {
		if !isInSubSet(i, newOrd) {
			newOrd[counter] = i
			counter++
		}
	}

	return newOrd
}

func isInSubSet(vertex int, subset []int) bool {
	for _, v := range subset {
		if v == vertex {
			return true
		}
	}
	return false
}
