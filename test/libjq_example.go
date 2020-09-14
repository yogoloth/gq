package main

import (
	"os"

	libjq "github.com/threatgrid/jq-go"
)

type Box struct {
	Width  int
	Height int
	Color  string
	Open   bool
}

func main() {

	//libjq.Dump(os.Stdout, "select(. != 0) | 1 / .", 1, 0, 2, 4)
	box := Box{
		Width:  10,
		Height: 20,
		Color:  "blue",
		Open:   false,
	}
	libjq.Dump(os.Stdout, ".a=1", box)
	return
}
