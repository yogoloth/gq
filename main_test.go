package main

import (
	"io/ioutil"
	"testing"
)

var config config_t

func assertMainSuccess(t *testing.T, config *config_t, hope_file string) {
	t.Helper()
	result, err := do_main(config)
	if err != nil {
		t.Error("parse failed!\n")
	}
	hope, err := ioutil.ReadFile(hope_file)
	if err != nil {
		t.Errorf("read request %s failed!\n", hope_file)
	}
	if string(hope) != string(result) {
		t.Errorf("\nrequest: ########\n%s\n get: ########\n%s\n", string(hope), string(result))
	}
}

func TestMain(t *testing.T) {
	request_file := "sample/test.yml"
	t.Run("run yq with jq", func(t *testing.T) {
		hope_file := "sample/test.yml"
		config = config_t{false, "jq", "yaml", "yaml", ".", request_file}
		assertMainSuccess(t, &config, hope_file)
	})
	t.Run("run yq with libjq", func(t *testing.T) {
		hope_file := "sample/test.yml"
		config = config_t{false, "libjq", "yaml", "yaml", ".", request_file}
		assertMainSuccess(t, &config, hope_file)
	})
	t.Run("run yq with jq add", func(t *testing.T) {
		hope_file := "sample/test_add.yml"
		config = config_t{false, "jq", "yaml", "yaml", `.a.b.c="世界"`, request_file}
		assertMainSuccess(t, &config, hope_file)
	})
	t.Run("run yq with libjq add", func(t *testing.T) {
		hope_file := "sample/test_add.yml"
		config = config_t{false, "libjq", "yaml", "yaml", `.a.b.c="世界"`, request_file}
		assertMainSuccess(t, &config, hope_file)
	})
	t.Run("find service deployed on servers from oper project", func(t *testing.T) {
		hope_file := "sample/test_oper.yml"
		request_file = "sample/prod_default_packages.yml"
		config = config_t{false, "libjq", "yaml", "yaml", `.[][]|select(.servers[0]|capture(".*oper.*"))`, request_file}
		assertMainSuccess(t, &config, hope_file)
	})
}
