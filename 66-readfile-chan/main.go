package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// type LineData struct {
// 	LineNo   int
// 	LineText string
// }

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

	go ReceiveLine(chLine, mu, sig)
	<-sig
	for k, v := range lineMap {
		fmt.Println(k, "-->", v)
	}
	//<-sig
	//time.Sleep(time.Second * 5)
}

func readNSend(fn string, chLine chan struct {
	lineNo int
	line   string
}, chErr chan error) {
	file, err := os.Open(fn)
	defer file.Close()
	if err != nil {
		println("--->")
		//chErr <- err
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
		//chErr <- scanner.Err()
	}
	close(chLine)
	close(chErr)
}

func ReceiveLine(chLine <-chan struct {
	lineNo int
	line   string
}, mu *sync.Mutex, sig chan<- struct{}) {
	wg := new(sync.WaitGroup)
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
	wg.Wait()
	sig <- struct{}{}
}

// read the file
// assign numbers to the line of the text
// findout number of words per a line
// gorountine and channel
