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

	switch config.engine {
	case "jq":
		buffer, err = json.Marshal(mid_result)
		if err != nil {
			err = errors.New(fmt.Sprintf("convert mid_data err: %v\n", err))
			return
		}
		buffer, err = jq(config.query, buffer)
		if err != nil {
			err = errors.New(fmt.Sprintf("run jq err: %v\n", err))
			return
		}
	case "libjq":
		seq, seq_err := libjq.Apply(config.query, mid_result)
		if seq_err != nil {
			err = errors.New(fmt.Sprintf("apply jq err: %v\n", err))
			return
		}
		//fmt.Printf("return %v\n", string(seq[0]))
		buffer = seq[0]
		//fmt.Printf("hello %s\n", string(buffer))

	default:
		err = errors.New(fmt.Sprintf("no engine %s\n", config.engine))
		return
	}

	//fmt.Println(string(data))

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
