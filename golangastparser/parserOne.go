package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	//src := `
	//	package main
	//
	//	import "fmt"
	//
	//	func main() {
	//		fmt.Println("Hello, world!")
	//	}
	//`

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "/Users/pandurang/projects/golang/helloworld/hello.go", nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	ast.Print(fset, file)
}
