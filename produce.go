package main

type Produce interface {
	Generate(chan string, interface{})
}
