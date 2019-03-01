package _example

//StructDecl is an example of an anonymous structure with a predetermined
//set of values.
var StructDecl = []struct {
	name string
}{
	{name: "test1"},
	{name: "test2"},
	{name: "test3"},
}

var IntDecl = 42

func FuncDecl() {}

type InterfaceDecl interface {
	do()
}
