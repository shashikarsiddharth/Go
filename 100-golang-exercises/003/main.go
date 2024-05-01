package main

import "fmt"

func generateSquare(input int) map[int]int {
	res := make(map[int]int)
	for i := 1; i <= input; i++ {
		res[i] = i * i
	}
	return res
}

func main() {
	var i int
	fmt.Scanf("%d", &i)
	fmt.Println(generateSquare(i))
}
