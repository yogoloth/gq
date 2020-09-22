package main

type IEngine interface {
	set_input(intput *map[string]interface{})
	run() ([]byte, error)
}
