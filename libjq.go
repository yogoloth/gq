package main

import (
	"errors"
	"fmt"
	libjq "github.com/snowcrystall/jq-go"
	//libjq "github.com/threatgrid/jq-go"
)

type LibjqEngine struct {
	query string
	input *map[string]interface{}
}

func (e LibjqEngine) set_input(input *map[string]interface{}) {
	e.input = input
}

func (e LibjqEngine) run() (buffer []byte, err error) {
	seq_buffer, seq_err := libjq.Apply(e.query, e.input)
	if seq_err != nil {
		err = errors.New(fmt.Sprintf("apply jq err: %v\n", seq_err))
		return
	}
	buffer = seq_buffer[0]
	return
}
