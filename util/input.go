package util

import (
	"strconv"
	"strings"
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
