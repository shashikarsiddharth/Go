package main

import "fmt"

type Class interface {
	ReadString()
	PrintString()
}

type C struct{}

func ReadString(c *C) {
	fmt.Println("Hello")
}

func PrintString(c *C) {
	fmt.Println("World")
}

func Hello() {
	fmt.Println("1")
}
