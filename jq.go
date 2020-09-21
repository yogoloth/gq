package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
)

type JqEngine struct {
	query string
	input *map[string]interface{}
}

func (e JqEngine) set_input(input *map[string]interface{}) {
	e.input = input
}

func (e JqEngine) run() (buffer []byte, err error) {
	buffer, err = json.Marshal(e.input)
	if err != nil {
		err = errors.New(fmt.Sprintf("convert mid_data err: %v\n", err))
		return
	}
	buffer, err = jq(e.query, buffer)
	return
}

func jq(query string, data []byte) ([]byte, error) {
	var output bytes.Buffer
	cmd := exec.Command("jq", query)
	cmd.Stdout = &output
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Printf("run jq %s error: %v\n", query, err)
		return nil, err
	}
	if err = cmd.Start(); err != nil {
		fmt.Printf("run jq %s error: %v\n", query, err)
		return nil, err
	}
	stdin.Write(data)
	stdin.Close()

	if err = cmd.Wait(); err != nil {
		fmt.Printf("run jq error: %v\n", err)
		return nil, err
	}

	//fmt.Println(string(output.Bytes()))

	return output.Bytes(), nil
}
