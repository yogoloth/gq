package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"testing"
)

var config config_t

func checkMain(config *config_t, hope_file string) error {
	result, err := do_main(config)
	if err != nil {
		return errors.New("parse failed!\n")
	}
	hope, err := ioutil.ReadFile(hope_file)
	if err != nil {
		errors.New(fmt.Sprintf("read request %s failed!\n", hope_file))
	}
	if string(hope) != string(result) {
		errors.New(fmt.Sprintf("\nrequest: \n%s\n get: \n%s\n", string(hope), string(result)))
	}
	return nil
}

func TestYqWithJq(t *testing.T) {
	request_file := "sample/test.yml"
	hope_file := "sample/test.yml"
	config = config_t{false, "jq", "yaml", "yaml", ".", request_file}
	if err := checkMain(&config, hope_file); err != nil {
		t.Error(err)
	}
}
func TestYqWithLibjq(t *testing.T) {
	request_file := "sample/test.yml"
	hope_file := "sample/test.yml"
	config = config_t{false, "libjq", "yaml", "yaml", `.a.b.c="世界"`, request_file}
	if err := checkMain(&config, hope_file); err != nil {
		t.Error(err)
	}
}
func TestYqWithJqAdd(t *testing.T) {
	request_file := "sample/test.yml"
	hope_file := "sample/test_add.yml"
	config = config_t{false, "jq", "yaml", "yaml", `.a.b.c="世界"`, request_file}
	if err := checkMain(&config, hope_file); err != nil {
		t.Error(err)
	}
}
func TestYqWithLibjqAdd(t *testing.T) {
	request_file := "sample/test.yml"
	hope_file := "sample/test_add.yml"
	config = config_t{false, "libjq", "yaml", "yaml", `.a.b.c="世界"`, request_file}
	if err := checkMain(&config, hope_file); err != nil {
		t.Error(err)
	}
}
