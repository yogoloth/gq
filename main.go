package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/ghodss/yaml"
	//"github.com/yogoloth/yaml"
)

type Dict map[string]interface{}

type config_t struct {
	verbose   bool
	engine    string
	from_type string
	to_type   string
	query     string
	filepath  string
}

func SplitJson(json []byte) (json_lines []string) {
	brackets := 0
	begin := 0
	end := 0

	if json[0] != '{' && json[0] != '[' {
		json_lines = append(json_lines, string(json))
		return
	}

	for i := 0; i < len(json); i++ {
		//fmt.Printf("%c-%d ", json[i], i)
		switch json[i] {
		case '{', '[':
			//fmt.Printf("%c\n", '{')
			brackets++
			if brackets == 1 {
				begin = i
			}
		case '}', ']':
			//fmt.Printf("%c\n", '}')
			brackets--
			if brackets == 0 {
				end = i
			}
		}
		if end > begin {
			//json_line := string(json[begin : end+1])
			//fmt.Printf("got %s\n", json_line)
			json_lines = append(json_lines, string(json[begin:end+1]))
			begin = 0
			end = 0
		}
	}

	return
}

func do_main(config *config_t) (output []byte, err error) {
	var buffer []byte
	var input []byte
	var engine IEngine
	var factory EngineFactory
	mid_result := Dict{}
	//mid_result := []map[string]interface{}
	out_buffer := bytes.Buffer{}
	is_array := false

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

	if is_array, _ = regexp.MatchString("^ *-", string(input)); is_array {
		//input = format_input(input)
		tmp_buff := bytes.Buffer{}
		tmp_buff.WriteString("root:\n")
		tmp_buff.Write(input)
		input = tmp_buff.Bytes()
	}

	if config.from_type == "yaml" {
		if err = yaml.Unmarshal(input, &mid_result); err != nil {
			err = errors.New(fmt.Sprintf("decode input yaml - %v\n", err))
			return
		}
	} else {
		err = errors.New(fmt.Sprintf("input type is not support yet: %v\n", config.from_type))
		return

	}

	if config.verbose == true {
		fmt.Printf("input data is :\n%v\n", mid_result)
	}

	if is_array {
		engine, err = factory.createEngine(config.engine, config.query, mid_result["root"])
	} else {
		engine, err = factory.createEngine(config.engine, config.query, mid_result)
	}

	if err != nil {
		err = errors.New(fmt.Sprintf("create engine err: %v\n", err))
		return
	}

	buffer, err = engine.run()
	if err != nil {
		err = errors.New(fmt.Sprintf("run jq err: %v\n", err))
		return
	}

	if config.to_type == "yaml" {
		//r, _ := regexp.Compile("\n")
		//js_lines := r.FindAllString(string(buffer), -1)

		js_lines := SplitJson(buffer)

		if js_lines == nil {
			err = errors.New(fmt.Sprintf("mid json parse error: %v\n,data is:\n\n%v\n", err, buffer))
			return
		}
		for i := 0; i < len(js_lines); i++ {
			var mid []byte
			if mid, err = yaml.JSONToYAML([]byte(js_lines[i])); err != nil {
				err = errors.New(fmt.Sprintf("convert mid data to yaml err: %v, %s\n", err, js_lines[i]))
				return
			} else {
				out_buffer.Write(mid)
				if i < len(js_lines)-1 {
					out_buffer.WriteString("---\n")
				}
			}
			output = out_buffer.Bytes()
		}
	} else if config.to_type == "json" {
		output = buffer
		return
	} else {
		err = errors.New(fmt.Sprintf("output type is not support yet: %v\n", config.from_type))
		return
	}
	return

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
