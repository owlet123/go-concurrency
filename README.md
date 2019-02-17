# Concurrency in Go example

This example contains three structures to show communication between the channels in goroutines.

## Structures

### Reader

The Reader structure simulates reading of a given number of data and sending them to a read channel.

### Filter 

The Filter structure simulates filtering of a read data and sending them to a write channel if it fulfills a given restriction.

### Writer

The Writer structure simulates writing a filtered data to output.


## Build and run

To build and run this example use following commands:
```
go get github.com/owlet123/concurrency
cd $GOPATH/src/github.com/owlet123/concurrency
go run *.go

```

## Contributing

Please, report any issue or create pull request.