package p1

import (
	"net/http"
	"fmt"
	"testing"
	"github.com/gin-gonic/gin"
)

// Index ...
func Index(c *gin.Context) {
	c.String(http.StatusOK, "hello, index")
}

//Helloworld ...
func Helloworld(c *gin.Context) {
	c.String(http.StatusOK, "hello, world")
}



// type Programmer interface {
// 	WriteHelloWorld() Code
// }

// type GoProgrammer struct {
// 	name string
// }

// type JavaProgrammer struct {
// 	name string
// }

// type Code string

// func (g *GoProgrammer) WriteHelloWorld() Code {
// 	fmt.Println(g.name)
// 	return "fmt.Println helloworld"
// }

// func (j *JavaProgrammer) WriteHelloWorld() Code {
// 	fmt.Println(j.name)
// 	return "System.out.Println helloworld"
// }

// func echo(p Programmer) {
// 	fmt.Println(p.WriteHelloWorld())
// }

// func TestSliceInt(t *testing.T) {
// 	p := &GoProgrammer{name: "Golang:"}
// 	j := &JavaProgrammer{name: "Java:"}

// 	echo(p)
// 	echo(j)
// }
