package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

func args_error(error interface{}) {
	fmt.Printf("args error! %v\n", error)
	os.Exit(1)
}

func parse_args() *config_t {
	config := config_t{}
	config.to_type = "json"
	is_yaml := false
	if path.Base(os.Args[0]) == "gq" || path.Base(os.Args[0]) == "yq" {
		config.from_type = "yaml"
	} else {
		args_error(os.Args)
	}
	flag.BoolVar(&config.verbose, "v", false, "is verbose?")
	flag.BoolVar(&is_yaml, "y", false, "output type is yaml")

	flag.StringVar(&config.engine, "engine", "libjq", "call external jq or internal libjq")
	flag.Parse()

	if is_yaml {
		config.to_type = "yaml"
	}

	config.query = flag.CommandLine.Arg(0)
	if config.query == "" {
		args_error(os.Args)
	}

	config.filepath = flag.CommandLine.Arg(1)
	if config.filepath == "" || config.filepath == "-" {
		config.filepath = "stdin"
	}

	return &config

}
