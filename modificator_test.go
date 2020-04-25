package gomodificator

import (
	"fmt"
	"testing"

	ls "github.com/Rakiiii/goBipartitonLocalSearch"
)

func TestIsInSubSet(t *testing.T) {
	fmt.Println("Start TestIsInSubSet")

	set := []int{1, 2, 3, 4, 5, 6}
	if !isInSubSet(2, set) || isInSubSet(7, set) {
		t.Error("Wrong check result")
	} else {
		fmt.Println("TestIsInSubSet=[ok]")
	}
}

func TestConstractNewOrder(t *testing.T) {
	fmt.Println("Start TestConstractNewOrder")

	dset := []int{6, 7, 9, 10}
	iset := []int{0, 3, 4}

	ai := 5
	av := 11
	test := constractNewOrder(dset, iset, ai, av)

	res := []int{1, 2, 6, 7, 9, 10, 0, 3, 4, 5, 8}

	checkFlag := true

	for i, v := range test {
		if v != res[i] {
			t.Error("Wrong value:", v, " at position:", i, " expected:", res[i])
			checkFlag = false
		}
	}

	if checkFlag {
		fmt.Println("TestConstractNewOrder=[ok]")
	}
}

func TestInitOrd(t *testing.T) {
	fmt.Println("Start TestInitOrd")

	res := []int{0, 1, 2, 3, 4, 5}

	test := make([]int, 6)
	test = initOrd(test)

	checkFlag := true

	for i, v := range test {
		if v != res[i] {
			t.Error("Wrong value:", v, " at position:", i, " expected:", res[i])
			checkFlag = false
		}
	}

	if checkFlag {
		fmt.Println("TestInitOrd=[ok]")
	}
}

func TestImpovement(t *testing.T) {
	fmt.Println("Start TestImprove")

	checkFlag := true

	var graph ls.Graph
	graph.ParseGraph("Testing/TestImpovement")
	graph.SetAmountOfIndependent(2)

	var resgraph ls.Graph
	resgraph.ParseGraph("Testing/TestImpovementResult")
	resgraph.SetAmountOfIndependent(3)

	res := []int{2, 3, 5, 0, 1, 4, 6, 7}

	flag, tgraph, testord := improve(&graph)

	if !flag {
		t.Error("Wrong return flag")
		checkFlag = false
	}

	for i, v := range testord {
		if v != res[i] {
			t.Error("Wrong value:", v, " at position:", i, " expected:", res[i])
			checkFlag = false
		}
		if !checkFlag {
			fmt.Println(testord)
			break
		}
	}

	if resgraph.GetAmountOfIndependent() != graph.GetAmountOfIndependent() {
		t.Error("Wrong amount of independent:", graph.GetAmountOfIndependent(), " expected:", resgraph.GetAmountOfIndependent())
		checkFlag = false
	}

	for i := 0; i < tgraph.AmountOfVertex(); i++ {
		res := resgraph.GetEdges(i)
		for j, v := range tgraph.GetEdges(i) {
			if v != res[j] {
				t.Error("Wrong value:", v, " at position:[", i, ",", j, "]", " expected:", res[i])
				checkFlag = false
			}
		}
		if !checkFlag {
			tgraph.Print()
			break
		}
	}

	if checkFlag {
		fmt.Println("TestImprove=[ok]")
	}
}

func TestImproveIndependetSet(t *testing.T) {
	fmt.Println("Start TestImproveIndependetSet")

	checkFlag := true

	var graph ls.Graph
	graph.ParseGraph("Testing/TestImpovement")
	graph.SetAmountOfIndependent(2)

	var resgraph ls.Graph
	resgraph.ParseGraph("Testing/TestImpovementResult")
	resgraph.SetAmountOfIndependent(3)

	res := []int{2, 3, 5, 0, 1, 4, 6, 7}

	_, testord := ImproveIndependetSet(&graph)

	for i, v := range testord {
		if v != res[i] {
			t.Error("Wrong value:", v, " at position:", i, " expected:", res[i])
			checkFlag = false
		}
		if !checkFlag {
			fmt.Println(testord)
			break
		}
	}

	if checkFlag {
		fmt.Println("TestImproveIndependetSet=[ok]")
	}
}
