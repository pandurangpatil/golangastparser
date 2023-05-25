package main

import "C"
import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  *int
}

func main() {
	age := 30
	var p *Person
	p = &Person{Name: "John", Age: &age}
	jsonstr := serilizeToJson(p)
	fmt.Println(jsonstr)
}

//export ExternallyCalled
func ExternallyCalled() *C.char {
	age := 30
	var p *Person
	p = &Person{Name: "John", Age: &age}
	result := serilizeToJson(p)
	return C.CString(result)
}

//export Add
func Add(a int, b int) int {
	return a + b
}

func serilizeToJson(node interface{}) string {
	objectMap := make(map[string]interface{})
	pointerType := reflect.TypeOf(node)
	valueOf := reflect.ValueOf(node).Elem()
	valueType := pointerType.Elem()
	if pointerType.Kind() == reflect.Ptr {
		// Get the type of the value pointed to by the pointer
		fmt.Println("Pointer Type:", pointerType)
		fmt.Println("Value Type:", valueType)
	} else {
		fmt.Println("Not a pointer type")
	}

	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		value := valueOf.Field(i)
		fmt.Println("field type", field)
		fmt.Println("Field name ->", field.Name)
		fmt.Println("Field value ->", value.Interface())
		if field.Type.Kind() == reflect.Ptr {
			fieldValueType := field.Type.Elem()
			fmt.Println("field value type", fieldValueType)
			fmt.Println("field pointer value", value.Elem().Interface())
			objectMap[field.Name] = value.Elem().Interface()
		} else {
			objectMap[field.Name] = value.Interface()
		}
	}
	jsonStr, _ := json.MarshalIndent(objectMap, "", "  ")
	return string(jsonStr)
}

// build

//  go build -buildmode=c-shared -o lib-sample.dylib sample.go
