package main

import (
	"errors"
)

type IEngine interface {
	set_input(intput *Dict)
	run() ([]byte, error)
}

type EngineFactory struct {
}

func (EngineFactory) createEngine(engine_type string, query string, input *Dict) (engine IEngine, err error) {
	switch engine_type {
	case "jq":
		engine = &JqEngine{query, input}
	case "libjq":
		engine = &LibjqEngine{query, input}
	default:
		err = errors.New("jq engine config error: should by jq or libjq\n\n  ")
	}
	return
}
