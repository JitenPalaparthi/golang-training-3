package main

import (
	"fmt"
	"os"
)

func main() {
	n := New("data.txt").Write([]byte("Hello ICICI"))
	if n > 0 {
		fmt.Println("successfully writen to the file")
	}
	//os.Exit(1)
}

func New(fn string) *FileOps {
	return &FileOps{FileName: fn}
}

type FileOps struct {
	FileName string
}

func (fo *FileOps) Write(bytes []byte) int {
	f, err := os.OpenFile(fo.FileName, os.O_WRONLY, 0644)
	if err != nil {
		panic(err.Error())
		//FatalIt(err.Error())
	}
	n, err := f.Write(bytes)
	if err != nil {
		//FatalIt(err.Error())
		panic(err.Error())
	}
	return n
}

func FatalIt(msg ...any) {
	fmt.Println(msg)
	os.Exit(1)
}

// error
// panic
// fatal
