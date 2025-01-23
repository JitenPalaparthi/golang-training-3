package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Start of main")
	defer fmt.Println("\nend of main")
	func() { //func-1
		defer fmt.Println("end of func-1")
		fmt.Println("start of func-1")
	}()
	write("data.txt", []byte("\nHello ICICI"))
	read("data.txt")
}

// defer defers the execution to the end of the call stack

func read(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	b := make([]byte, 1)

	for {
		_, err := file.Read(b)
		if err != nil {
			break
		}
		fmt.Print(string(b))
	}
	return nil
}

func write(filename string, p []byte) (int, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.Write(p)
}
