package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, "/Users/pandurang/projects/golang/helloworld/", nil, 0)
	//file, err := parser.ParseFile(fset, "/Users/pandurang/projects/golang/helloworld/hello.go", nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	for key, element := range pkgs {
		fmt.Printf("Key -> %v", key)
		ast.Inspect(element, func(node ast.Node) bool {
			if node != nil {
				pos := fset.Position(node.Pos())
				fmt.Printf("%v\t%v\n", pos, node)
			}
			return true
		})
	}
	//pkgs

}
