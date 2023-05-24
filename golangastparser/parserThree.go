package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"reflect"
)

func main() {
	fset := token.NewFileSet()
	//file, err := parser.ParseDir(fset, "/Users/pandurang/projects/golang/helloworld/", nil, 0)
	file, err := parser.ParseFile(fset, "/Users/pandurang/projects/golang/helloworld/hello.go", nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Generate the JSON representation of the AST
	astJSON := generateASTJSON(file, fset)
	//astJSON := generateASTJSONPkg(file, fset)

	// Print the JSON
	fmt.Println(astJSON)
}

func generateASTJSONPkg(node map[string]*ast.Package, fset *token.FileSet) string {
	var astJSON interface{}

	// Inspect the AST and generate the JSON representation
	//ast.Inspect(node, func(n ast.Node) bool {
	//	// Convert the AST node to a map representation
	//	fmt.Println(node)
	//	astMap := astToMap(n, fset)
	//
	//	// Set astJSON to the converted map if it's the root node
	//	if astJSON == nil {
	//		astJSON = astMap
	//	}
	//
	//	return true
	//})
	for _, element := range node {
		astJSON = astToMap(element, fset)
	}
	// Convert the map to JSON
	astJSONBytes, err := json.MarshalIndent(astJSON, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(astJSONBytes)
}

func generateASTJSON(node ast.Node, fset *token.FileSet) string {
	var astJSON interface{}

	// Inspect the AST and generate the JSON representation
	//ast.Inspect(node, func(n ast.Node) bool {
	//	// Convert the AST node to a map representation
	//	fmt.Println(node)
	//	astMap := astToMap(n, fset)
	//
	//	// Set astJSON to the converted map if it's the root node
	//	if astJSON == nil {
	//		astJSON = astMap
	//	}
	//
	//	return true
	//})
	astJSON = astToMap(node, fset)
	// Convert the map to JSON
	astJSONBytes, err := json.MarshalIndent(astJSON, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(astJSONBytes)
}

func astToMap(node ast.Node, fset *token.FileSet) map[string]interface{} {
	astMap := make(map[string]interface{})
	if node != nil {
		// Get the type of the node
		nodeType := reflect.TypeOf(node).Elem()

		// Iterate over the fields of the node
		for i := 0; i < nodeType.NumField(); i++ {
			field := nodeType.Field(i)
			fieldValue := reflect.ValueOf(node).Elem().Field(i)

			// Skip unexported fields
			if field.PkgPath != "" {
				continue
			}

			// Convert the field value to the appropriate representation
			var val interface{}
			println("field Kinde -> " + field.Type.Kind().String())
			println("fieldValue Kinde -> " + fieldValue.Kind().String())
			switch field.Type.Kind() {
			case reflect.Struct:
				if astNode, ok := fieldValue.Interface().(ast.Node); ok {
					val = astToMap(astNode, fset)
				}
			case reflect.Slice:
				if fieldValue.Type().Elem().Kind() == reflect.Struct {
					var nodeList []interface{}
					for j := 0; j < fieldValue.Len(); j++ {
						astNode := fieldValue.Index(j).Interface().(ast.Node)
						nodeList = append(nodeList, astToMap(astNode, fset))
					}
					val = nodeList
				}
			default:
				val = fieldValue.Interface()
			}

			if val != nil {
				astMap[field.Name] = val
			}
			// Add filename, line number, and column number
		}
		if pos := node.Pos(); pos.IsValid() {
			position := fset.Position(pos)
			astMap["Filename"] = position.Filename
			astMap["Line"] = position.Line
			astMap["Column"] = position.Column
		}
	}

	return astMap
}
