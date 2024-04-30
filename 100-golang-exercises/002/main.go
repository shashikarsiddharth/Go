package main

import (
	"fmt"
	"strconv"
	"strings"
)

func factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

func parseInput(input string) []int {
	inp := strings.Split(input, ",")
	var i []int
	for _, s := range inp {
		n, _ := strconv.Atoi(s)
		i = append(i, n)
	}
	return i
}

func main() {
	var input string
	fmt.Scan(&input)

	var res []string
	for _, i := range parseInput(input) {
		op := strconv.Itoa(factorial(i))
		res = append(res, op)
	}
	fmt.Println(strings.Join(res, ","))
}
