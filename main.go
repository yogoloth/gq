package main

import (
	"encoding/json"
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

func main() {

	var buffer []byte
	var input []byte
	var err error
	mid_result := make(map[string]interface{})

	config := parse_args()
	if config.verbose {
		fmt.Printf("config: %v\n\n", config)
	}

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
		//buffer, err = yaml.YAMLToJSON(input)
		//if err != nil {
		//	fmt.Printf("err: %v\n", err)
		//	return
		//}
		yaml.Unmarshal(input, &mid_result)
	} else {
		fmt.Printf("input type is not support yet: %v\n", config.from_type)
		return

	}

	switch config.engine {
	case "jq":
		buffer, err = json.Marshal(mid_result)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		buffer, err = jq(config.query, buffer)
		if err != nil {
			fmt.Printf("run jq err: %v\n", err)
			return
		}
	case "libjq":
		seq, seq_err := libjq.Apply(config.query, mid_result)
		if seq_err != nil {
			fmt.Printf("apply jq err: %v\n", err)
			return
		}
		//fmt.Printf("return %v\n", string(seq[0]))
		buffer = seq[0]
		//fmt.Printf("hello %s\n", string(buffer))

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
