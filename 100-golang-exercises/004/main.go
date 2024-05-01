package main

import (
	"fmt"
	"strings"
)

func main() {
	var inp string
	var out []string
	fmt.Scan(&inp)
	out = strings.Split(inp, ",")
	fmt.Println(out)
}
