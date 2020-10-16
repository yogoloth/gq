package main

import (
	"bytes"
	"errors"
	"fmt"
	libjq "github.com/snowcrystall/jq-go"
	//libjq "github.com/threatgrid/jq-go"
)

type LibjqEngine struct {
	query string
	input *Dict
}

func (e *LibjqEngine) set_input(input *Dict) {
	e.input = input
}

func (e *LibjqEngine) run() (buffer []byte, err error) {
	seq_buffer, seq_err := libjq.Apply(e.query, e.input)
	if seq_err != nil {
		err = errors.New(fmt.Sprintf("apply jq err: %v\n", seq_err))
		return
	}

	tmp := bytes.Buffer{}
	for _, b := range seq_buffer {
		tmp.Write(b)
		tmp.WriteByte('\n')
	}
	buffer = tmp.Bytes()

	return
}
