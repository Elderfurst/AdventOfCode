package main

import (
	"adventofcode/utilities"
	"fmt"
	"slices"
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

	paths := make([]int, 0)
	for node := range graph.nodes {
		visited := make(map[string]bool)
		paths = append(paths, getShortestPath(graph, node, visited))
	}

	fmt.Println(slices.Min(paths))
}

func partTwo() {
	inputLines := utilities.ReadInputFileToLines("2015/day09/input.txt")
	graph := parseInputLines(inputLines)

	paths := make([]int, 0)
	for node := range graph.nodes {
		visited := make(map[string]bool)
		paths = append(paths, getLongestPath(graph, node, visited))
	}

	fmt.Println(slices.Max(paths))
}

func getShortestPath(graph *Graph, node string, visited map[string]bool) int {
	visited[node] = true

	if len(visited) == len(graph.nodes) {
		return 0
	}

	paths := make([]int, 0)
	for _, edge := range graph.edges[node] {
		if !visited[edge.right] {
			paths = append(paths, edge.distance+getShortestPath(graph, edge.right, copyVisited(visited)))
		}
	}

	return slices.Min(paths)
}

func getLongestPath(graph *Graph, node string, visited map[string]bool) int {
	visited[node] = true

	if len(visited) == len(graph.nodes) {
		return 0
	}

	paths := make([]int, 0)
	for _, edge := range graph.edges[node] {
		if !visited[edge.right] {
			paths = append(paths, edge.distance+getLongestPath(graph, edge.right, copyVisited(visited)))
		}
	}

	return slices.Max(paths)
}

func parseInputLines(inputLines []string) *Graph {
	graph := &Graph{
		nodes: make(map[string]bool),
		edges: make(map[string][]*Edge),
	}

	for _, line := range inputLines {
		splitLine := strings.Split(line, " ")
		first := splitLine[0]
		second := splitLine[2]
		distance, _ := strconv.Atoi(splitLine[4])

		graph.nodes[first] = true
		graph.nodes[second] = true

		graph.edges[first] = append(graph.edges[first], &Edge{
			left:     first,
			right:    second,
			distance: distance,
		})
		graph.edges[second] = append(graph.edges[second], &Edge{
			left:     second,
			right:    first,
			distance: distance,
		})
	}

	return graph
}

func copyVisited(visited map[string]bool) map[string]bool {
	newVisited := make(map[string]bool)
	for k, v := range visited {
		newVisited[k] = v
	}
	return newVisited
}

type Graph struct {
	nodes map[string]bool
	edges map[string][]*Edge
}

type Edge struct {
	left     string
	right    string
	distance int
}
