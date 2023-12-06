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

func LinesToIntMap(lines []string) map[int]map[int]int {
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
