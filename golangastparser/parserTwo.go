package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()
	//file, err := parser.ParseDir(fset, "/Users/pandurang/projects/golang/helloworld/hello.go", nil, 0)
	file, err := parser.ParseFile(fset, "/Users/pandurang/projects/golang/helloworld/hello.go", nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	ast.Print(fset, file)
}
