package main

import (
	"fmt"
	"strconv"
	"strings"
)

var LOWER_LIMIT = 2000
var UPPER_LIMIT = 3200
 
func main(){
	var result []string
	for i := LOWER_LIMIT; i <= UPPER_LIMIT; i++ {
		if i % 7 == 0 && i % 5 != 0 {
			result = append(result, strconv.Itoa(i))
		}
	}
	fmt.Println(strings.Join(result,","))
}