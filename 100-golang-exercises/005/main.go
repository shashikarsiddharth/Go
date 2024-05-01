package main

import "fmt"

type MyIO struct{}

type IO interface {
	ReadString() string
	PrintString(input string)
}

func (io MyIO) ReadString() string {
	var input string
	fmt.Scan(&input)
	return input
}

func (io MyIO) PrintString(input string) {
	fmt.Println(input)
}

func main() {
	var io IO = MyIO{}
	inp := io.ReadString()
	io.PrintString(inp)
}
