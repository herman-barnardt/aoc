package util

import (
	"strconv"
	"strings"

	"github.com/herman-barnardt/aoc/graph"
)

func SplitLines(lines []string, seperator string) [][]string {
	retVal := make([][]string, 0)
	temp := make([]string, 0)
	for _, line := range lines {
		if line == seperator {
			retVal = append(retVal, temp)
			temp = make([]string, 0)
			continue
		}
		temp = append(temp, line)
	}
	retVal = append(retVal, temp)
	return retVal
}

func LinesToMap(lines []string) map[int]map[int]string {
	retVal := make(map[int]map[int]string)
	for y, line := range lines {
		retVal[y] = make(map[int]string)
		for x, c := range strings.Split(line, "") {
			retVal[y][x] = c
		}
	}
	return retVal
}

func LinesToPointMap(lines []string) map[Point]string {
	retVal := make(map[Point]string)
	for y, line := range lines {
		for x, c := range strings.Split(line, "") {
			retVal[Point{x, y}] = c
		}
	}
	return retVal
}

func LinesToPointMapOfInts(lines []string) map[Point]int {
	retVal := make(map[Point]int)
	for y, line := range lines {
		for x, c := range strings.Split(line, "") {
			n, _ := strconv.Atoi(c)
			retVal[Point{x, y}] = n
		}
	}
	return retVal
}

func LinesToMapofInts(lines []string) map[int]map[int]int {
	retVal := make(map[int]map[int]int)
	for y, line := range lines {
		retVal[y] = make(map[int]int)
		for x, c := range strings.Split(line, "") {
			n, _ := strconv.Atoi(c)
			retVal[y][x] = n
		}
	}
	return retVal
}

func LinesToPointMapOfBasicNodes(lines []string, neighourFilter func(*graph.BasicNode) bool) map[Point]*graph.BasicNode {
	retVal := make(map[Point]*graph.BasicNode)
	for y, line := range lines {
		for x, c := range strings.Split(line, "") {
			retVal[Point{x, y}] = &graph.BasicNode{X: x, Y: y, Value: c, PossibleNeighbours: make([]*graph.BasicNode, 0), NeighbourFilter: neighourFilter}
		}
	}
	for point, node := range retVal {
		for _, neighbour := range point.GetAdjacent() {
			if _, ok := retVal[neighbour]; ok {
				node.PossibleNeighbours = append(node.PossibleNeighbours, retVal[neighbour])
			}
		}
	}
	return retVal
}

func CreatePointMapOfBasicNodes(maxX, maxY int, value string, neighourFilter func(*graph.BasicNode) bool) map[Point]*graph.BasicNode {
	retVal := make(map[Point]*graph.BasicNode)
	for y := range maxY {
		for x := range maxX {
			retVal[Point{x, y}] = &graph.BasicNode{X: x, Y: y, Value: value, PossibleNeighbours: make([]*graph.BasicNode, 0), NeighbourFilter: neighourFilter}
		}
	}
	for point, node := range retVal {
		for _, neighbour := range point.GetAdjacent() {
			if _, ok := retVal[neighbour]; ok {
				node.PossibleNeighbours = append(node.PossibleNeighbours, retVal[neighbour])
			}
		}
	}
	return retVal
}

func StringToIntSlice(line string, seperator string) []int {
	retVal := make([]int, 0)
	for _, c := range strings.Split(line, seperator) {
		n, err := strconv.Atoi(c)
		if err == nil {
			retVal = append(retVal, n)
		}
	}
	return retVal
}
