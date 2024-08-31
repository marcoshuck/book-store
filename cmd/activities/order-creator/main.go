package main

import (
	"github.com/marcoshuck/book-store/orders"
	"os"
)

func main() {
	if err := orders.RunOrderCreatorWorker(); err != nil {
		os.Exit(1)
		return
	}
	os.Exit(0)
}
