package util

func SplitLines(lines []string, seperator string) [][]string {
	retVal := make([][]string, 0)
	temp := make([]string, 0)
	for _, line := range lines {
		if line == seperator {
			retVal = append(retVal, temp)
			temp = make([]string, 0)
		}
		temp = append(temp, line)
	}
	retVal = append(retVal, temp)
	return retVal
}
