// Package definestruct package includeExamples of defining combining structs
package definestruct

import (
	"github.com/rl4debug/mygopkg/src/examples/design-patterns/strategy"
)

// BaseStruct is simplest struct for this example, used in combining struct
type BaseStruct struct {
	Name string
}

//CompositStruct1 is the first way to define a struct (this way is not recommmended)
type CompositStruct1 struct {
	int        //implicitly define a member of struct named 'int' and with type int
	string     //similar to above
	BaseStruct //similar to above
}

// CompositStruct2 is the second way to define a struct (this way is recommended, it's clear)
type CompositStruct2 struct {
	ValueInt    int
	ValueString string
	ValueStruct BaseStruct
}

func testDeclaringStructs() {
	strategy.Test()

	//way 1
	cstruct1 := CompositStruct1{10, "string", BaseStruct{}}
	cstruct1.int = 10

	cstruct2 := CompositStruct1{
		10,
		"hehe",
		BaseStruct{},
	}

	cstruct2.int = 10
}
