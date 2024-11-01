package main

import (
	"adventofcode/utilities"
	"fmt"
	"github.com/hmdsefi/gograph"
	"github.com/hmdsefi/gograph/traverse"
	"math"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	inputLines := utilities.ReadInputFileToLines("2015/day09/input.txt")
	graph := parseInputLines(inputLines)

}

func partTwo() {

}

func parseInputLines(inputLines []string) *Graph {
	graph := &Graph {
		Nodes: make([]*Node, 0),
		Edges: make([]*Edge, 0),
	}

	for _, line := range inputLines {
		splitLine := strings.Split(line, " ")
		first := splitLine[0]
		second := splitLine[2]
		distance, _ := strconv.ParseFloat(splitLine[4], 64)

		if graph

		firstNode := new Node {
			Label: first,
		}
	}

	return graph
}

type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

func (g *Graph) ContainsNode(n string) bool {
	for _, node := range g.Nodes {
		if node.Label == n {
			return true
		}
	}

	return false
}

type Node struct {
	Label string
	Edges []*Edge
}

type Edge struct {
	Left     *Node
	Right    *Node
	Distance int
}
