package main

import (
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
)

type config_t struct {
	engine    string
	from_type string
	to_type   string
	query     string
	filepath  string
}

func main() {

	config := parse_args()
	fmt.Printf("config: %v\n\n", config)

	var buffer []byte
	var input []byte
	var err error

	if config.filepath == "stdin" {
		input, err = ioutil.ReadAll(os.Stdin)
	} else {
		input, err = ioutil.ReadFile(config.filepath)
	}
	if err != nil {
		fmt.Printf("read file %v error\n", err)
		return
	}

	if config.from_type == "yaml" {
		buffer, err = yaml.YAMLToJSON(input)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
	} else {
		fmt.Printf("input type is not support yet: %v\n", config.from_type)
		return

	}

	switch config.engine {
	case "jq":
		buffer, err = jq(config.query, buffer)
		if err != nil {
			fmt.Printf("run jq err: %v\n", err)
			return
		}
	case "libjq":

	default:
		fmt.Printf("no engine %s\n", config.engine)
		return
	}

	//fmt.Println(string(data))

	if config.to_type == "yaml" {
		if j2, err := yaml.JSONToYAML(buffer); err != nil {
			fmt.Printf("err: %v\n", err)
			return
		} else {
			fmt.Println(string(j2))
		}
	} else if config.to_type == "json" {
		fmt.Println(string(buffer))
	} else {
		fmt.Printf("output type is not support yet: %v\n", config.from_type)
		return
	}
}
