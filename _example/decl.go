package _example

//StructDecl is an example of an anonymous structure with a predetermined
//set of values.
var StructDecl = []struct {
	name  string
	field string
}{
	{name: "test1", field: "string"},
	{name: "test2"},
	{name: "test3"},
}

type Struct struct {
	name string
}

var AnotherStructDecl = []Struct{
	{name: "test4"},
	{name: "test5"},
}

var NonArrayStruct = struct {
	field string
}{field: "field"}

var StringArray = []string{"one", "two", "three"}

var IntDecl = 42

func FuncDecl() {}

type InterfaceDecl interface {
	do()
}
