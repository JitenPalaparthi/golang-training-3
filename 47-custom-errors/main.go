package main

import (
	"fmt"
	"os"
)

func main() {

	// err := func() error {
	// 	return errors.New("Some error")
	// }()
	// fmt.Println(err.Error())

	fo := &FileOps{FileName: "data.txt"}
	if n, err := fo.Write([]byte("Hello ICICI")); err != nil {
		//fmt.Println(err)
		cr, ok := err.(*CustomError)
		if ok {
			fmt.Println("Custom Error")
			fmt.Println("Error Code:", cr.Code)
			fmt.Println("Error Message:", cr.Msg)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("file successfully created and wrote", n, "bytes")
	}

}

type FileOps struct {
	FileName string
}

func (fo *FileOps) Write(bytes []byte) (int, error) {
	f, err := os.OpenFile(fo.FileName, os.O_WRONLY, 0644)
	if err != nil {
		return 0, &CustomError{Code: 101, Msg: "File Error:" + err.Error()}
	}
	n, err := f.Write(bytes)
	if err != nil {
		return 0, &CustomError{Code: 102, Msg: "File Error:" + err.Error()}
	}
	return n, nil
}

type CustomError struct {
	Code int
	Msg  string
}

func (c *CustomError) Error() string {
	return fmt.Sprint("Code:", c.Code, " Error Message: ", c.Msg)
}
