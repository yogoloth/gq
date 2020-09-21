package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

type config_t struct {
	verbose   bool
	engine    string
	from_type string
	to_type   string
	query     string
	filepath  string
}

type IEngine interface {
	set_input(intput *map[string]interface{})
	run() ([]byte, error)
}

func do_main(config *config_t) (output []byte, err error) {
	var buffer []byte
	var input []byte
	mid_result := make(map[string]interface{})

	if config.verbose {
		fmt.Printf("config: %v\n\n", config)
	}

	if config.filepath == "stdin" {
		input, err = ioutil.ReadAll(os.Stdin)
	} else {
		input, err = ioutil.ReadFile(config.filepath)
	}
	if err != nil {
		err = errors.New(fmt.Sprintf("read file %v error\n", err))
		return
	}

	if config.from_type == "yaml" {
		//buffer, err = yaml.YAMLToJSON(input)
		//if err != nil {
		//	fmt.Printf("err: %v\n", err)
		//	return
		//}
		if err = yaml.Unmarshal(input, &mid_result); err != nil {
			err = errors.New(fmt.Sprintf("decode input yaml - %v\n", err))
			return
		}
	} else {
		err = errors.New(fmt.Sprintf("input type is not support yet: %v\n", config.from_type))
		return

	}

	var engine IEngine
	switch config.engine {
	case "jq":
		engine = JqEngine{config.query, &mid_result}
	case "libjq":
		engine = LibjqEngine{config.query, &mid_result}
	default:
		fmt.Fprintf(os.Stderr, "jq engine config error: %v\n\n", os.Args)
	}
	//engine.set_input(&mid_result)
	buffer, err = engine.run()
	if err != nil {
		err = errors.New(fmt.Sprintf("run jq err: %v\n", err))
		return
	}

	if config.to_type == "yaml" {
		if output, err = yaml.JSONToYAML(buffer); err != nil {
			err = errors.New(fmt.Sprintf("convert mid data to yaml err: %v\n", err))
			return
		} else {
			return
		}
	} else if config.to_type == "json" {
		output = buffer
		return
	} else {
		err = errors.New(fmt.Sprintf("output type is not support yet: %v\n", config.from_type))
		return
	}

}

func main() {
	config := parse_args()
	result, err := do_main(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "do main error - %v\n", err)
		return
	}
	if result != nil {
		fmt.Printf("%s", string(result))
	}
}
