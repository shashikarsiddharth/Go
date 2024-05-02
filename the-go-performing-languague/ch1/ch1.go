package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Purpose:To echo the command line arguments
func echo() {
	var s string
	// var sep string
	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// }
	s = strings.Join(os.Args[1:], " ")
	fmt.Println(s)
}

// Purpose:To print count of chars in string
func charCounter(input string) map[string]int {
	var char = make(map[string]int)
	for i := range input {
		char[string(input[i])]++
	}
	return char
}

// Purpose: To print duplicate input lines along with count
func printDuplicate() {
	var count = make(map[string]int)
	inp := bufio.NewScanner(os.Stdin)
	for inp.Scan() {
		if inp.Err() != nil {
			log.Printf("failed to scan input from stdin: %v", inp.Err())
		}
		count[inp.Text()]++
	}
	fmt.Println(count)
}

// Purpose: To find duplicates from a file using file handling
func findDuplicateV2(){
	var count = make(map[string]int)
		
}

func main() {
	// fmt.Println(charCounter("HelloWorld"))
	printDuplicate()
}
