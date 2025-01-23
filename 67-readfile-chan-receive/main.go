package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	lineMap = make(map[int]struct {
		Count int
		Words []string
	})
	mu = new(sync.Mutex)
)

func main() {
	chLine := make(chan struct {
		lineNo int
		line   string
	})

	chErr := make(chan error)
	sig := make(chan struct{})
	go readNSend("data.txt", chLine, chErr)

	go Receive(chLine, chErr, mu, sig)
	<-sig
	for k, v := range lineMap {
		fmt.Println(k, "-->", v)
	}

}

func readNSend(fn string, chLine chan struct {
	lineNo int
	line   string
}, chErr chan error) {
	file, err := os.Open(fn)
	defer file.Close()
	if err != nil {
		println("--->")
		chErr <- err
	}
	scanner := bufio.NewScanner(file)
	l := 0
	for scanner.Scan() {
		println("-->", scanner.Text())
		l++
		chLine <- struct {
			lineNo int
			line   string
		}{line: scanner.Text(), lineNo: l}
	}
	if scanner.Err != nil {
		fmt.Println("-----?????????")
		chErr <- errors.New("some error XXXXXX-------------->>>>>>>>>>>>>>>>>>") //scanner.Err()
	}
	close(chLine)
	close(chErr)
}

func Receive(chLine <-chan struct {
	lineNo int
	line   string
}, chErr chan error, mu *sync.Mutex, sig chan<- struct{}) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		for lineData := range chLine {
			//fmt.Println("XXXX--->>", lineData)
			wg.Add(1)
			go func() {
				words := strings.Split(lineData.line, " ")
				mu.Lock()
				lineMap[lineData.lineNo] = struct {
					Count int
					Words []string
				}{Count: len(words), Words: words}
				mu.Unlock()
				wg.Done()
			}()
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for e := range chErr {
			fmt.Println(e)
		}
		wg.Done()
	}()
	wg.Wait()
	sig <- struct{}{}
}

// read the file
// assign numbers to the line of the text
// findout number of words per a line
// gorountine and channel
