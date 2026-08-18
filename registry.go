package main

type config struct {
	registry map[string]cliCommand
}

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}


