package main

import (
	"fmt"
	"sync"
	"time"
)

type Writer struct {

}

// simulate writing
// this function wait 2 seconds and write data to output
func simulateWriting(i int) {
	time.Sleep(2 * time.Second)
	fmt.Println("write", i)
}

// provide writing
func (w *Writer) writing(data int, wg *sync.WaitGroup) {
	simulateWriting(data)

	wg.Done()
}

// read filtered data and send it to the output and
// send signal when writing is done to finish the application
func (w *Writer) Write(done chan bool) {
	fmt.Println("start writing")

	var wg sync.WaitGroup

	for data := range write {
		wg.Add(1)
		go w.writing(data, &wg)
	}

	wg.Wait()
	done <- true

	fmt.Println("writing finished")
}
