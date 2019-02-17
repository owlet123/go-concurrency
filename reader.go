package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Reader struct {

}

// simulate reading
// this function wait 3 seconds and create random value from 0 to 100
func simulateReading() int {
	time.Sleep(3 * time.Second)
	i := rand.Intn(100)
	fmt.Println("read", i)
	return i
}

// provide reading and send data to the channel
func (r *Reader) reading(wg *sync.WaitGroup) {
	read <- simulateReading()

	wg.Done()
}

// read given number of data
// close channel after reading
func (r *Reader) Read(rcount int) {
	fmt.Println("start reading")

	var wg sync.WaitGroup

	for i := 0; i < rcount; i++ {
		wg.Add(1)
		go r.reading(&wg)
	}

	wg.Wait()
	close(read)
	fmt.Println("reading finished")
}
