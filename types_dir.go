package main

import (
	"context"
	"go/ast"
	"go/printer"
	"go/token"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

// TypesDir implements both Node and Handle for the root directory.
type TypesDir struct {
	GoNode ast.Node
}

func (TypesDir) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = 1
	a.Mode = os.ModeDir | 0555 // ----r-x-wx
	return nil
}

func (TypesDir) Lookup(ctx context.Context, name string) (fs.Node, error) {
	if name == "hello" {
		return File{}, nil
	}
	return nil, fuse.ENOENT
}

func (t TypesDir) ReadDirAll(ctx context.Context) (list []fuse.Dirent, err error) {
	c := new(typeCollector)
	ast.Walk(c, t.GoNode)
	for _, each := range c.names {
		list = append(list, fuse.Dirent{Inode: 2, Name: each, Type: fuse.DT_File})
	}
	return
}

type typeHolder struct {
	typeSpec *ast.TypeSpec
	methods  map[string]*ast.FuncDecl
}

type typeCollector struct {
	names []string
}

func (c *typeCollector) Visit(n ast.Node) ast.Visitor {
	g, ok := n.(*ast.GenDecl)
	if ok {
		if g.Tok == token.TYPE {
			for _, each := range g.Specs {
				ts := each.(*ast.TypeSpec)
				c.names = append(c.names, ts.Name.String())
			}
		}
		return nil
	}
	f, ok := n.(*ast.FuncDecl)
	if ok {
		if f.Recv == nil {
			return nil
		}
		printer.Fprint(os.Stdout, token.NewFileSet(), f)
		//log.Println(f.Recv.List[0])
	}
	return c
}
