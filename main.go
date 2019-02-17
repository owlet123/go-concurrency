package main

import (
	"fmt"
)

// create channels to read and write
var read chan int
var write chan int

// create objects
var reader Reader
var filter Filter
var writer Writer

// initialize channels
func init() {
	read = make(chan int)
	write = make(chan int)
}

func main() {
	// number of read data
	rcount := 10

	// done channel to quit main
	done := make(chan bool)

	fmt.Println("start concurrency")

	// goroutines to process data: read -> filter -> write
	go reader.Read(rcount)
	go filter.Filter()
	go writer.Write(done)

	// notification that writing is done
	<- done

	fmt.Println("done")
}
