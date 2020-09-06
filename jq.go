package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

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