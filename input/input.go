package input

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func Lines(filename string) []string {
	content, _ := ioutil.ReadFile(filename)
	return strings.Split(string(content), "\n")
}

func ToInts(lines []string) []int {
	digits := []int{}
	for _, r := range lines {
		d, _ := strconv.Atoi(r)
		digits = append(digits, d)
	}
	return digits
}
