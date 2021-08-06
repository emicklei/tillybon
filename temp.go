package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

// temporary
func createGoNode() ast.Node {
	fset := token.NewFileSet() // positions are relative to fset

	src := `package foo

import (
	"fmt"
	"time"
)

var V = "var"

const C = "const"

// Foo comment
type Foo struct {

}

func (f Foo) String() string { return "Foo!" }

type Bar int

func bar() {
	fmt.Println(time.Now())
}`

	f, err := parser.ParseFile(fset, "embedded", src, parser.ParseComments)
	if err != nil {
		log.Fatalln(err)
	}
	return f
}
