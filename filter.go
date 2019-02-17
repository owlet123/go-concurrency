package main

import (
	"fmt"
	"sync"
	"time"
)

type Filter struct {

}

// simulate filter
// this function wait 1 second and return true if given value is smaller than 50
func simulateFilter(i int) bool {
	time.Sleep(1 * time.Second)
	fmt.Println("filter", i)
	return i < 50
}

// provide filtering and send data to write channel if it fulfills a filter restriction
func (f *Filter) filtering(data int, wg *sync.WaitGroup) {
	if simulateFilter(data) {
		write <- data
	}

	wg.Done()
}

// filter data from read channel until the channel is closed,
// send to write channel if it fulfills a filter restriction and
// close write channel after filtering
func (f *Filter) Filter() {
	fmt.Println("start filtering")

	var wg sync.WaitGroup

	for data := range read {
		wg.Add(1)
		go f.filtering(data, &wg)
	}

	wg.Wait()
	close(write)
	fmt.Println("filtering finished")
}
