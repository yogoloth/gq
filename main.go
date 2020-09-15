package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
	libjq "github.com/threatgrid/jq-go"
)

type config_t struct {
	verbose   bool
	engine    string
	from_type string
	to_type   string
	query     string
	filepath  string
}

func do_main(config *config_t) ([]byte, error) {
	var buffer []byte
	var input []byte
	var err error
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
		return nil, errors.New(fmt.Sprintf("read file %v error\n", err))
	}

	if config.from_type == "yaml" {
		//buffer, err = yaml.YAMLToJSON(input)
		//if err != nil {
		//	fmt.Printf("err: %v\n", err)
		//	return
		//}
		yaml.Unmarshal(input, &mid_result)
	} else {
		return nil, errors.New(fmt.Sprintf("input type is not support yet: %v\n", config.from_type))

	}

	switch config.engine {
	case "jq":
		buffer, err = json.Marshal(mid_result)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("convert mid_data err: %v\n", err))
		}
		buffer, err = jq(config.query, buffer)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("run jq err: %v\n", err))
		}
	case "libjq":
		seq, seq_err := libjq.Apply(config.query, mid_result)
		if seq_err != nil {
			return nil, errors.New(fmt.Sprintf("apply jq err: %v\n", err))
		}
		//fmt.Printf("return %v\n", string(seq[0]))
		buffer = seq[0]
		//fmt.Printf("hello %s\n", string(buffer))

	default:
		return nil, errors.New(fmt.Sprintf("no engine %s\n", config.engine))
	}

	//fmt.Println(string(data))

	if config.to_type == "yaml" {
		if j2, err := yaml.JSONToYAML(buffer); err != nil {
			return nil, errors.New(fmt.Sprintf("convert mid data to yaml err: %v\n", err))
		} else {
			return j2, nil
		}
	} else if config.to_type == "json" {
		return buffer, nil
	} else {
		return nil, errors.New(fmt.Sprintf("output type is not support yet: %v\n", config.from_type))
	}

}

func main() {
	config := parse_args()
	result, err := do_main(config)
	if err != nil {
		fmt.Printf("do work error: %v\n", err)
		return
	}
	fmt.Println(string(result))

}
