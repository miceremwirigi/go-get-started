package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Package io provides basic interfaces to I/O(input/output) primitives
	// 1) Writer, Reader & Seeker interfaces
	// 2) Functions working alongside with these interfaces

	// Writer interface - wraps the basic Write method.
	// type Wrier interface {
	// 	Write(p []byte) (n int, err error)
	// }

	file, _ := os.Create("file.txt")
	writer := io.Writer(file)
	n, err := writer.Write([]byte("Hello, World"))
	fmt.Println(n, err)
	n, err = io.WriteString(writer, "!")
	fmt.Println(n, err)
	file.Close()

	file, _ = os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	writer = io.Writer(file)
	n, err = writer.Write([]byte("\nHey"))
	fmt.Println(n, err)
	file.Close()

}
