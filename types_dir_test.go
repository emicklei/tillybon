package main

import (
	"context"
	"go/ast"
	"testing"
)

func TestTypesCollector(t *testing.T) {
	n := createGoNode()
	c := new(typeCollector)
	ast.Walk(c, n)
	t.Log(c.names)
}

func TestTypesDirReadDirAll(t *testing.T) {
	td := TypesDir{GoNode: createGoNode()}
	list, _ := td.ReadDirAll(context.Background())
	if got, want := len(list), 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
