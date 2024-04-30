package main

import 
(
	"fmt"
	// "strconv"
)

func factorial(num int) int {
	return 0
}

func main(){
	var input string
	inp := []int{}
	fmt.Scanf("%s", &input)
	for i := range(input){
		fmt.Println(i)
		// append(inp, strconv.Atoi(input[i]))
	}

	fmt.Println(inp)
	fmt.Println(factorial(1))
}