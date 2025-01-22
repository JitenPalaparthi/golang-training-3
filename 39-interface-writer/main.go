package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "Hello ICICI")
	fl := New("data.txt")
	fmt.Fprintln(fl, "Hello ICICI")
}

// Write(p []byte) (n int, err error)

type File struct {
	FileName string
}

func New(filename string) *File {
	return &File{FileName: filename}
}

func (f *File) Write(p []byte) (n int, err error) {
	fl, err := os.OpenFile(f.FileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return 0, err
	}
	n, err = fl.Write(p)
	if err != nil {
		fl.Close()
		return 0, err
	}
	return n, nil
}

// F/D User Group others
// 0   6     4    4
// 0777
// 4 - Read
// 2 - Write
// 1 - Execute

// 0  0  0  0  0  0  0
// R  WR E  C  A  D  RO
// 0  1  0  1  1  0  0
