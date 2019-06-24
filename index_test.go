package main

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
	name string
}

type JavaProgrammer struct {
	name string
}

type Code string

func (g *GoProgrammer) WriteHelloWorld() Code {
	fmt.Println(g.name)
	return "fmt.Println helloworld"
}

func (j *JavaProgrammer) WriteHelloWorld() Code {
	fmt.Println(j.name)
	return "System.out.Println helloworld"
}

func Echo(p Programmer) {
	fmt.Println(p.WriteHelloWorld())
}

func TestSliceInt(t *testing.T) {
	p := &GoProgrammer{name: "Golang:"}
	j := &JavaProgrammer{name: "Java:"}

	Echo(p)
	Echo(j)
}
